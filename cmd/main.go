package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/yuyz0112/sunmao-ui-go-binding/pkg/runtime"
	"github.com/yuyz0112/sunmao-ui-go-binding/pkg/sunmao"
)

func main() {
	// init the runtime
	r := runtime.New("ui")
	// init an App builder, use a lib
	b := sunmao.NewChakraUIApp()

	arcoB := sunmao.NewArcoApp()

	// build some UI, some CSS is good
	// more UI
	b.Component(b.NewStack().Children(map[string][]sunmao.BaseComponentBuilder{
		"content": {
			b.NewText().Content("Hello Sunmao"),
			b.NewInput().Id("my_input"),
		},
	}).Properties(map[string]interface{}{
		"direction": "vertical",
	}).Style("content", `
		width: 100%;
		padding: 1em;
		background: #f1f3f5;
	`))

	// use server dynamic data
	entries, _ := os.ReadDir(".")
	data := []map[string]interface{}{}
	for _, e := range entries {
		info, _ := e.Info()
		data = append(data, map[string]interface{}{
			"name":    e.Name(),
			"size":    info.Size(),
			"modTime": info.ModTime().Format(time.UnixDate),
		})
	}

	b.Component(arcoB.NewTabs().Id("tabs").
		Tab(&sunmao.ArcoTabsTab{Title: "Tab 1"}).
		Tab(&sunmao.ArcoTabsTab{Title: "Tab 2"}).
		Tab(&sunmao.ArcoTabsTab{Title: "Tab 3"}))

	b.Component(arcoB.NewTable().Data(data).Column(&sunmao.ArcoTableColumn{
		DataIndex:    "name",
		Title:        "Name",
		Type:         "link",
		DisplayValue: "{{ $listItem.name }} - {{ $listItem.size }}",
	}).Column(&sunmao.ArcoTableColumn{
		DataIndex: "size",
		Title:     "File Size",
	}).Column(&sunmao.ArcoTableColumn{
		DataIndex: "modTime",
		Title:     "Modify Time",
	}).Hidden("{{ tabs.activeTab != 0 }}"))

	// use server push real-time data
	type MyState struct {
		Random int `json:"random"`
	}
	myState := r.NewServerState("server_push", &MyState{})
	b.Component(myState.AsComponent())

	b.Component(b.NewText().Content("data from server {{ server_push.state.random }}"))

	go func() {
		for {
			time.Sleep(1 * time.Second)
			// update state
			//myState.SetState(&MyState{Random: rand.Int()})

			// call any UI's method like an API
			//r.Execute(&runtime.ExecuteTarget{
			//	Id:     "my_input",
			//	Method: "setInputValue",
			//	Parameters: map[string]interface{}{
			//		"value": time.Now().Format(time.UnixDate),
			//	},
			//})
		}
	}()

	r.HandleStore(func(s map[string]any, connId int) error {
		fmt.Println(s, connId)
		return nil
	})

	// add any server function as an API
	r.Handle("debug", func(m *runtime.Message, connId int) error {
		store := r.GetStore()

		type Input struct {
			Value string `json:"value"`
		}
		var input *Input

		jsonData, _ := json.Marshal(store["my_input"])
		_ = json.Unmarshal(jsonData, &input)

		fmt.Println("debug >", input, "from >", connId)
		return nil
	})

	r.Handle("writeFile", func(m *runtime.Message, connId int) error {
		content, _ := json.Marshal(m.Params)
		return os.WriteFile("test", content, 777)
	})

	b.Component(b.NewButton().Content("click to debug").
		OnClick(&sunmao.ServerHandler{
			Name: "debug",
			Parameters: map[string]interface{}{
				// use ID to access component state
				"dynamic": "input value {{ my_input.value }}",
			},
		}))

	r.LoadApp(b.AppBuilder)

	// start the runtime
	r.Run()
}
