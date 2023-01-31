package sunmao

import (
	"fmt"
)

// arco-design

type ArcoAppBuilder struct {
	*AppBuilder
}

func NewArcoApp() *ArcoAppBuilder {
	b := &ArcoAppBuilder{
		AppBuilder: NewApp(),
	}
	return b
}

type ArcoTableComponentBuilder struct {
	*InnerComponentBuilder[*ArcoTableComponentBuilder]
}

func (b *ArcoAppBuilder) NewTable() *ArcoTableComponentBuilder {
	t := &ArcoTableComponentBuilder{
		InnerComponentBuilder: newInnerComponent[*ArcoTableComponentBuilder](b.AppBuilder),
	}
	t.inner = t
	return t.Type("arco/v1/table").Properties(map[string]interface{}{
		"pagination": map[string]interface{}{
			"enablePagination": true,
			"pageSize":         20,
		},
		"rowKey":  "name",
		"data":    []interface{}{},
		"columns": []interface{}{},
	})
}

func (b *ArcoTableComponentBuilder) Data(data any) *ArcoTableComponentBuilder {
	b.Properties(map[string]interface{}{
		"data": data,
	})
	return b
}

type ArcoTableColumn struct {
	Title        string           `json:"title"`
	DataIndex    string           `json:"dataIndex"`
	Type         string           `json:"type,omitempty"`
	Sorter       bool             `json:"sorter"`
	Filter       bool             `json:"filter"`
	DisplayValue string           `json:"displayValue,omitempty"`
	Module       *ModuleContainer `json:"module,omitempty"`
}

func (b *ArcoTableComponentBuilder) Column(column *ArcoTableColumn) *ArcoTableComponentBuilder {
	columns := b.ValueOf().Properties["columns"].([]interface{})
	columns = append(columns, column)
	b.Properties(map[string]interface{}{
		"columns": columns,
	})
	return b
}

func (b *ArcoTableComponentBuilder) OnRowClick(serverHandler *ServerHandler) *ArcoTableComponentBuilder {
	b._Trait(b.appBuilder.NewTrait().Type("core/v1/event").Properties(map[string]interface{}{
		"handlers": []map[string]interface{}{
			{
				"type":        "onRowClick",
				"componentId": "$utils",
				"method": map[string]interface{}{
					"name":       fmt.Sprintf("binding/v1/%v", serverHandler.Name),
					"parameters": serverHandler.Parameters,
				},
			},
		},
	}))
	return b
}

type ArcoTabsComponentBuilder struct {
	*InnerComponentBuilder[*ArcoTabsComponentBuilder]
}

func (b *ArcoAppBuilder) NewTabs() *ArcoTabsComponentBuilder {
	t := &ArcoTabsComponentBuilder{
		InnerComponentBuilder: newInnerComponent[*ArcoTabsComponentBuilder](b.AppBuilder),
	}
	t.inner = t
	return t.Type("arco/v1/tabs").Properties(map[string]interface{}{
		"type":                          "line",
		"defaultActiveTab":              0,
		"tabPosition":                   "top",
		"size":                          "default",
		"updateWhenDefaultValueChanges": false,
		"tabs":                          []any{},
	})
}

type ArcoTabsTab struct {
	Title         string `json:"title"`
	Hidden        bool   `json:"hidden"`
	DestroyOnHide bool   `json:"destroyOnHide"`
}

func (b *ArcoTabsComponentBuilder) Tab(tab *ArcoTabsTab) *ArcoTabsComponentBuilder {
	tabs := b.ValueOf().Properties["tabs"].([]any)
	tabs = append(tabs, tab)
	b.Properties(map[string]interface{}{
		"tabs": tabs,
	})
	return b
}
