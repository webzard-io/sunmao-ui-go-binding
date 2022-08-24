package runtime

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	echo "github.com/labstack/echo/v4"
	"github.com/yuyz0112/sunmao-ui-go-binding/pkg/sunmao"
	"log"
	"net/http"
	"os"
	"strings"
)

type Runtime struct {
	e                        *echo.Echo
	conns                    map[int]*websocket.Conn
	appBuilder               *sunmao.AppBuilder
	reloadWhenWsDisconnected bool
	handlers                 map[string]func(m *Message) error
}

func New() *Runtime {
	e := echo.New()

	return &Runtime{
		e:                        e,
		conns:                    map[int]*websocket.Conn{},
		reloadWhenWsDisconnected: true,
		handlers:                 map[string]func(m *Message) error{},
	}
}

var (
	upgrader = websocket.Upgrader{}
)

func (r *Runtime) Run() {
	if r.appBuilder == nil {
		log.Fatalln("please load app before run")
	}

	r.e.Static("/assets", "ui/dist/assets")
	r.e.GET("/", func(c echo.Context) error {
		buf, err := os.ReadFile("ui/dist/index.html")
		if err != nil {
			return err
		}

		handlers := []string{}
		for k := range r.handlers {
			handlers = append(handlers, k)
		}

		optionsBuf, err := json.Marshal(map[string]interface{}{
			"application":              r.appBuilder.ValueOf(),
			"reloadWhenWsDisconnected": r.reloadWhenWsDisconnected,
			"wsUrl":                    "ws://localhost:8999/ws",
			"handlers":                 handlers,
		})
		if err != nil {
			return err
		}
		html := strings.Replace(string(buf),
			"/* APPLICATION */",
			fmt.Sprintf("options = %v", string((optionsBuf))), 1)
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

		for {
			_, msgBytes, err := ws.ReadMessage()
			if err != nil {
				c.Logger().Error(err)
			}
			msg := &Message{}
			err = json.Unmarshal(msgBytes, msg)
			if err != nil {
				// ignore
			}
			if msg.Type == "Action" {
				handler, ok := r.handlers[msg.Handler]
				if ok {
					handler(msg)
				}
			}
		}
	})

	r.e.Logger.Fatal(r.e.Start(":8999"))
}

type Message struct {
	Type    string `json:"type"`
	Handler string `json:"handler"`
	Params  any    `json:"params"`
}

func (r *Runtime) LoadApp(builder *sunmao.AppBuilder) error {
	r.appBuilder = builder
	return nil
}

func (r *Runtime) Handle(handler string, fn func(m *Message) error) {
	r.handlers[handler] = fn
}

type ExecuteTarget struct {
	Id         string
	Method     string
	Parameters any
}

func (r *Runtime) Execute(target *ExecuteTarget) error {
	for _, ws := range r.conns {
		msg, err := json.Marshal(map[string]interface{}{
			"type":        "UiMethod",
			"componentId": target.Id,
			"name":        target.Method,
			"parameters":  target.Parameters,
		})
		if err != nil {
			return err
		}
		ws.WriteMessage(websocket.TextMessage, msg)
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

func (s *ServerState) SetState(newState any) error {
	return s.r.Execute(&ExecuteTarget{
		Id:     s.Id,
		Method: "setValue",
		Parameters: map[string]interface{}{
			"key":   "state",
			"value": newState,
		},
	})
}
