package runtime

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mitchellh/hashstructure"
	"github.com/yuyz0112/sunmao-ui-go-binding/pkg/sunmao"
)

type Runtime struct {
	e                        *echo.Echo
	conns                    map[int]*websocket.Conn
	appBuilder               *sunmao.AppBuilder
	moduleBuilders           []*sunmao.ModuleBuilder
	reloadWhenWsDisconnected bool
	handlers                 map[string]func(m *Message, connId int) error
	hooks                    map[string]func(connId int) error
	uiDir                    string
	patchDir                 string
	websocketMutex           sync.Mutex
}

func New(uiDir string, patchDir string) *Runtime {
	e := echo.New()

	r := &Runtime{
		e:                        e,
		conns:                    map[int]*websocket.Conn{},
		reloadWhenWsDisconnected: true,
		handlers:                 map[string]func(m *Message, connId int) error{},
		hooks:                    map[string]func(connId int) error{},
		uiDir:                    uiDir,
		patchDir:                 patchDir,
	}

	return r
}

var (
	upgrader = websocket.Upgrader{}
)

type Message struct {
	Type    string         `json:"type"`
	Handler string         `json:"handler"`
	Params  any            `json:"params"`
	Store   map[string]any `json:"store"`
}

type DeltaBody struct {
	Delta map[string]interface{} `json:"delta"`
}

func (r *Runtime) formatUiOptions() (*string, error) {
	handlers := []string{}
	for k := range r.handlers {
		handlers = append(handlers, k)
	}

	modules := make([]any, len(r.moduleBuilders))
	for i, b := range r.moduleBuilders {
		modules[i] = b.ValueOf()
	}

	var patchFile PatchFile
	patchFileBuf, err := os.ReadFile(fmt.Sprintf("%v/app.patch.json", r.patchDir))
	if err == nil {
		err = json.Unmarshal(patchFileBuf, &patchFile)
		if err != nil {
			return nil, err
		}
	}

	appBase := patchFile.Base
	appBaseHash, err := hashstructure.Hash(appBase, nil)
	currentAppJson, _ := json.Marshal(r.appBuilder.ValueOf())
	var currentApp sunmao.Application

	json.Unmarshal(currentAppJson, &currentApp)
	currentAppHash, err := hashstructure.Hash(currentApp, nil)

	hasBaseAppChanged := appBaseHash != currentAppHash
	var applicationBase *sunmao.Application = nil
	// only when base application has changed, send base application to front end
	if hasBaseAppChanged {
		applicationBase = appBase
	}

	modulesPatch := map[string]interface{}{}
	modulesPatchBuf, err := os.ReadFile(fmt.Sprintf("%v/modules.patch.json", r.patchDir))
	if err == nil {
		err = json.Unmarshal(modulesPatchBuf, &modulesPatch)
		if err != nil {
			return nil, err
		}
	}

	optionsBuf, err := json.Marshal(map[string]interface{}{
		"application":              currentApp,
		"modules":                  modules,
		"applicationPatch":         patchFile.Patch,
		"applicationBase":          applicationBase,
		"modulesPatch":             modulesPatch,
		"reloadWhenWsDisconnected": r.reloadWhenWsDisconnected,
		"handlers":                 handlers,
	})
	if err != nil {
		return nil, err
	}

	s := string(optionsBuf)
	return &s, nil
}

type PatchFile struct {
	Patch *map[string]interface{} `json:"patch"`
	Base  *sunmao.Application     `json:"base"`
}

func (r *Runtime) Run() {
	if r.appBuilder == nil {
		log.Fatalln("please load app before run")
	}

	os.MkdirAll(r.patchDir, os.ModePerm)

	r.e.Use(middleware.Gzip())

	r.e.Static("/assets", fmt.Sprintf("%v/dist/assets", r.uiDir))

	r.e.GET("/", func(c echo.Context) error {
		buf, err := os.ReadFile(fmt.Sprintf("%v/dist/index.html", r.uiDir))
		if err != nil {
			return err
		}

		options, err := r.formatUiOptions()
		if err != nil {
			return err
		}

		html := strings.Replace(string(buf),
			"/* APPLICATION */",
			fmt.Sprintf("options = Object.assign(options, %v)", *options), 1)
		return c.HTML(http.StatusOK, html)
	})

	r.e.GET("/editor", func(c echo.Context) error {
		buf, err := os.ReadFile(fmt.Sprintf("%v/dist/editor.html", r.uiDir))
		if err != nil {
			return err
		}

		options, err := r.formatUiOptions()
		if err != nil {
			return err
		}

		html := strings.Replace(string(buf),
			"/* APPLICATION */",
			fmt.Sprintf("options = Object.assign(options, %v)", *options), 1)
		return c.HTML(http.StatusOK, html)
	})

	r.e.PUT("/sunmao-binding-patch/app", func(c echo.Context) error {
		b := &DeltaBody{}
		if err := c.Bind(b); err != nil {
			return err
		}

		baseApp := r.appBuilder.ValueOf()
		patchFile := PatchFile{
			Base:  &baseApp,
			Patch: &b.Delta,
		}

		fileContent, err := json.MarshalIndent(patchFile, "", "\t")
		if err != nil {
			return err
		}
		// save app.patch.json
		writePatchErr := os.WriteFile(fmt.Sprintf("%v/app.patch.json", r.patchDir), fileContent, os.ModePerm)
		if writePatchErr != nil {
			return writePatchErr
		}

		return c.String(http.StatusOK, "ok")
	})

	r.e.PUT("/sunmao-binding-patch/modules", func(c echo.Context) error {
		b := &DeltaBody{}
		if err := c.Bind(b); err != nil {
			return err
		}

		delta, err := json.MarshalIndent(b.Delta, "", "\t")
		if err != nil {
			return err
		}

		err = os.WriteFile(fmt.Sprintf("%v/modules.patch.json", r.patchDir), delta, os.ModePerm)
		if err != nil {
			return err
		}

		return c.String(http.StatusOK, "ok")
	})

	r.e.GET("/sunmao-binding-patch/app/visualize", func(c echo.Context) error {
		appBuf, err := json.Marshal(r.appBuilder.ValueOf())
		if err != nil {
			return err
		}

		appPatchStr := "{}"
		appPatchBuf, err := os.ReadFile(fmt.Sprintf("%v/app.patch.json", r.patchDir))
		if err == nil {
			appPatchStr = string(appPatchBuf)
		}

		html := fmt.Sprintf(`
<!DOCTYPE html>
<html>
    <head>
        <script type='text/javascript' src="https://cdn.jsdelivr.net/npm/jsondiffpatch/dist/jsondiffpatch.umd.min.js"></script>
        <link rel="stylesheet" href="https://benjamine.github.io/jsondiffpatch/demo/style.css" type="text/css" />
        <link rel="stylesheet" href="https://benjamine.github.io/jsondiffpatch/formatters-styles/html.css" type="text/css" />
    </head>
    <body>
        <div id="visual"></div>
        <script>
            var left = %v;
            var delta = %v;
            document.getElementById('visual').innerHTML = jsondiffpatch.formatters.html.format(delta, left);
        </script>
    </body>
</html>
`, string(appBuf), appPatchStr)

		return c.HTML(http.StatusOK, html)
	})

	connId := 0

	r.e.GET("/ws", func(c echo.Context) error {
		ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			return err
		}
		connId++
		r.conns[connId] = ws
		defer func() {
			delete(r.conns, connId)
			ws.Close()
		}()

		connectedHook, ok := r.hooks["connected"]
		if ok {
			connectedHook(connId)
		}

		for {
			_, msgBytes, err := ws.ReadMessage()
			if err != nil {
				if strings.Contains(err.Error(), "close 1001") {
					disconnectedHook, ok := r.hooks["disconnected"]
					if ok {
						disconnectedHook(connId)
					}

					break
				} else {
					c.Logger().Error(err)
				}
			}

			msg := &Message{}

			err = json.Unmarshal(msgBytes, msg)
			if err != nil {
				// ignore
			}

			if msg.Type == "Action" {
				handler, ok := r.handlers[msg.Handler]
				if ok {
					handler(msg, connId)
				}
			}
		}

		return nil
	})

	r.e.Logger.Fatal(r.e.Start(":8999"))
}

func (r *Runtime) LoadApp(builder *sunmao.AppBuilder) error {
	r.appBuilder = builder
	return nil
}

func (r *Runtime) LoadModule(builder ...*sunmao.ModuleBuilder) error {
	r.moduleBuilders = builder
	return nil
}

func (r *Runtime) Handle(handler string, fn func(m *Message, connId int) error) {
	r.handlers[handler] = fn
}

func (r *Runtime) On(hook string, fn func(connId int) error) {
	r.hooks[hook] = fn
}

type ExecuteTarget struct {
	Id         string
	Method     string
	Parameters any
}

// maybe this is a bad idea, but currently we let connId == nil to represent broadcasting
func (r *Runtime) Execute(target *ExecuteTarget, connId *int) error {
	for id, ws := range r.conns {
		if connId != nil && id != *connId {
			continue
		}

		msg, err := json.Marshal(map[string]interface{}{
			"type":        "UiMethod",
			"componentId": target.Id,
			"name":        target.Method,
			"parameters":  target.Parameters,
		})
		if err != nil {
			return err
		}

		r.websocketMutex.Lock()
		defer r.websocketMutex.Unlock()

		err = ws.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return err
		}
	}
	return nil
}

type ServerState struct {
	r         *Runtime
	initState any
	Id        string
}

func (r *Runtime) NewServerState(id string, initState any) *ServerState {
	return &ServerState{
		r:         r,
		initState: initState,
		Id:        id,
	}
}

func (s *ServerState) AsComponent() sunmao.BaseComponentBuilder {
	t := s.r.appBuilder.NewComponent().Type("core/v1/dummy").Id(s.Id).
		Trait(
			s.r.appBuilder.NewTrait().Type("core/v1/state").
				Properties(map[string]interface{}{
					"key":          "state",
					"initialValue": s.initState,
				}))
	return t
}

func (s *ServerState) SetState(newState any, connId *int) error {
	return s.r.Execute(&ExecuteTarget{
		Id:     s.Id,
		Method: "setValue",
		Parameters: map[string]interface{}{
			"key":   "state",
			"value": newState,
		},
	}, connId)
}
