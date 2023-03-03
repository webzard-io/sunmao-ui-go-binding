package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/yuyz0112/sunmao-ui-go-binding/pkg/dovetail"
	"github.com/yuyz0112/sunmao-ui-go-binding/pkg/runtime"
	"github.com/yuyz0112/sunmao-ui-go-binding/pkg/sunmao"
)

func main() {
	// init the runtime
	r := runtime.New("ui", "patch")
	// init an App builder, use a lib
	app := sunmao.NewApp()
	
	d := dovetail.NewDovetailApp(app)
	
	app.Component(d.NewRoot().Id("root").Children(map[string][]sunmao.BaseComponentBuilder {
		"root": {
			d.NewButton().Id("button"),
			d.NewKubectlGetTable().Id("kubectl_get_table"),
			d.NewV2Modal().Id("modal").Children(map[string][]sunmao.BaseComponentBuilder {
				"content": {
					d.NewKubectlApplyForm().Id("kubectl_apply_form"),
				},
			}),
			d.NewKubectlGetDetail().Id("kubectl_get_detail"),
		},
	}))

	go func() {
		for {
			time.Sleep(1 * time.Second)
			// update state
			//myState.SetState(&MyState{Random: rand.Int()}, nil)
		}
	}()

	// add any server function as an API
	r.Handle("debug", func(m *runtime.Message, connId int) error {
		fmt.Println("debug >", m, "from >", connId)
		return nil
	})

	r.Handle("writeFile", func(m *runtime.Message, connId int) error {
		content, _ := json.Marshal(m.Params)
		err := os.WriteFile("test", content, 777)
		if err != nil {
			return err
		}

		// call any UI's method like an API
		r.Execute(&runtime.ExecuteTarget{
			Id:     "my_input",
			Method: "setInputValue",
			Parameters: map[string]interface{}{
				"value": time.Now().Format(time.UnixDate),
			},
		}, &connId)

		return nil
	})

	r.LoadApp(app)

	// start the runtime
	r.Run()
}
