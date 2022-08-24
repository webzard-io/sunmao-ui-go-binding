package main

import (
	"encoding/json"
	"fmt"
	"github.com/yuyz0112/sunmao-ui-go-binding/pkg/runtime"
	"github.com/yuyz0112/sunmao-ui-go-binding/pkg/sunmao"
	"math/rand"
	"os"
	"time"
)

func main() {
	// init the runtime
	r := runtime.New()
	// init an App builder, use a lib
	b := sunmao.NewChakraUIApp()

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
	b.Component(b.NewTable().Data(data).Column(&sunmao.ChakraTableColumn{
		Key:   "name",
		Title: "Name",
	}).Column(&sunmao.ChakraTableColumn{
		Key:   "size",
		Title: "File Size",
	}).Column(&sunmao.ChakraTableColumn{
		Key:   "modTime",
		Title: "Modify Time",
	}))

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
			myState.SetState(&MyState{Random: rand.Int()})

			// call any UI's method like an API
			r.Execute(&runtime.ExecuteTarget{
				Id:     "my_input",
				Method: "setInputValue",
				Parameters: map[string]interface{}{
					"value": time.Now().Format(time.UnixDate),
				},
			})
		}
	}()

	// add any server function as an API
	r.Handle("debug", func(m *runtime.Message) error {
		fmt.Println("debug >", m)
		return nil
	})
	r.Handle("writeFile", func(m *runtime.Message) error {
		content, _ := json.Marshal(m.Params)
		return os.WriteFile("test", content, 777)
	})
	b.Component(b.NewButton().Content("click to debug").
		OnClick(&sunmao.ServerHandler{
			Name: "writeFile",
			Parameters: map[string]interface{}{
				// use ID to access component state
				"dynamic": "input value {{ my_input.value }}",
			},
		}))

	r.LoadApp(b.AppBuilder)

	// start the runtime
	r.Run()
}
