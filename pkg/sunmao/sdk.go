package sunmao

import (
	"fmt"
)

// layer 1

type BaseBuilder[T any] struct {
	inner  T
	setter VersionMetadataSetter
}

func (b *BaseBuilder[T]) Version(version string) T {
	b.setter.SetVersion(version)
	return b.inner
}

func (b *BaseBuilder[T]) Name(name string) T {
	b.setter.SetName(name)
	return b.inner
}

func (b *BaseBuilder[T]) Description(description string) T {
	b.setter.SetDescription(description)
	return b.inner
}

func (b *BaseBuilder[T]) Annotation(key string, value string) T {
	b.setter.SetAnnotations(key, value)
	return b.inner
}

// App

type AppBuilder struct {
	*BaseBuilder[*AppBuilder]
	application Application
}

func NewApp() *AppBuilder {
	b := &AppBuilder{
		BaseBuilder: &BaseBuilder[*AppBuilder]{},
		application: Application{
			VersionMetadata: &VersionMetadata{
				Version: "example/v1",
				Metadata: Metadata{
					Name:        "empty",
					Description: "",
					Annotations: map[string]string{},
				},
			},
			Kind: "Application",
			Spec: ApplicationSpec{
				Components: []ComponentSchema{},
			},
		}}

	b.inner = b
	b.setter = b.application
	return b
}

var componentCount = 0

func newInnerComponent[K any](builder *AppBuilder) *InnerComponentBuilder[K] {
	componentCount = componentCount + 1
	return &InnerComponentBuilder[K]{
		component: ComponentSchema{
			Id:         "component" + fmt.Sprint(componentCount),
			Type:       "",
			Properties: map[string]interface{}{},
			Traits:     []TraitSchema{},
		},
		appBuilder: builder,
	}
}

func (b *AppBuilder) NewTrait() *TraitBuilder {
	return &TraitBuilder{
		trait: TraitSchema{
			Type:       "",
			Properties: map[string]interface{}{},
		},
	}
}

func (b *AppBuilder) ValueOf() Application {
	return b.application
}

func (b *AppBuilder) Component(builder BaseComponentBuilder) *AppBuilder {
	b.component(builder)
	return b
}

func (b *AppBuilder) component(builder BaseComponentBuilder) {
	b.application.Spec.Components = append(b.application.Spec.Components, builder.ValueOf())
}

type ModuleBuilder struct {
	*BaseBuilder[*ModuleBuilder]
	module Module
}

func NewModule() *ModuleBuilder {
	b := &ModuleBuilder{
		BaseBuilder: &BaseBuilder[*ModuleBuilder]{},
		module: Module{
			Kind: "Module",
			VersionMetadata: &VersionMetadata{
				Version: "custom/v1",
				Metadata: Metadata{
					Name:        "empty",
					Description: "",
					Annotations: map[string]string{},
				},
			},
			Impl: []ComponentSchema{},
			Spec: ModuleSpec{
				StateMap:   map[string]interface{}{},
				Properties: map[string]interface{}{},
			},
		},
	}

	b.inner = b
	b.setter = b.module
	return b
}

func (b *ModuleBuilder) Impl(app Application) *ModuleBuilder {
	b.module.Impl = append(b.module.Impl, app.Spec.Components...)
	return b
}

func (b *ModuleBuilder) StateMap(s map[string]interface{}) *ModuleBuilder {
	for k, v := range s {
		b.module.Spec.StateMap[k] = v
	}

	return b
}

func (b *ModuleBuilder) Properties(p map[string]interface{}) *ModuleBuilder {
	for k, v := range p {
		b.module.Spec.Properties[k] = v
	}

	return b
}

func (b *ModuleBuilder) ValueOf() Module {
	return b.module
}

// Component

type BaseComponentBuilder interface {
	ValueOf() ComponentSchema
	_Trait(builder BaseTraitBuilder)
}

type InnerComponentBuilder[K any] struct {
	inner      K
	component  ComponentSchema
	appBuilder *AppBuilder
}

type ComponentBuilder struct {
	*InnerComponentBuilder[*ComponentBuilder]
}

func (b *AppBuilder) NewComponent() *ComponentBuilder {
	t := &ComponentBuilder{
		InnerComponentBuilder: newInnerComponent[*ComponentBuilder](b),
	}
	t.inner = t
	return t
}

func (b *InnerComponentBuilder[K]) ValueOf() ComponentSchema {
	return b.component
}

func (b *InnerComponentBuilder[K]) Id(id string) K {
	b.component.Id = id
	return b.inner
}

func (b *InnerComponentBuilder[K]) Type(t string) K {
	b.component.Type = t
	return b.inner
}

func (b *InnerComponentBuilder[K]) Properties(properties map[string]interface{}) K {
	for k, v := range properties {
		b.component.Properties[k] = v
	}
	return b.inner
}

// type hack
func (b *InnerComponentBuilder[K]) _Trait(builder BaseTraitBuilder) {
	b.component.Traits = append(b.component.Traits, builder.ValueOf())
}

func (b *InnerComponentBuilder[K]) Trait(builder BaseTraitBuilder) K {
	b._Trait(builder)
	return b.inner
}

func (b *InnerComponentBuilder[K]) Children(slots map[string][]BaseComponentBuilder) K {
	for slot := range slots {
		for _, builder := range slots[slot] {
			builder._Trait(b.appBuilder.NewTrait().
				Type("core/v1/slot").Properties(map[string]interface{}{
				"container": map[string]interface{}{
					"id":   b.component.Id,
					"slot": slot,
				},
			}))
			b.appBuilder.Component(builder)
		}
	}
	return b.inner
}

func (b *InnerComponentBuilder[K]) Style(styleSlot string, css string) K {
	b._Trait(b.appBuilder.NewTrait().Type("core/v1/style").Properties(map[string]interface{}{
		"styles": []map[string]interface{}{
			{
				"styleSlot": styleSlot,
				"style":     css,
			},
		},
	}))
	return b.inner
}

func (b *InnerComponentBuilder[K]) Hidden(when string) K {
	b._Trait(b.appBuilder.NewTrait().Type("core/v1/hidden").Properties(map[string]interface{}{
		"hidden": when,
	}))
	return b.inner
}

// Trait

type BaseTraitBuilder interface {
	ValueOf() TraitSchema
}

type TraitBuilder struct {
	trait TraitSchema
}

func (b *TraitBuilder) ValueOf() TraitSchema {
	return b.trait
}

func (b *TraitBuilder) Type(t string) *TraitBuilder {
	b.trait.Type = t
	return b
}

func (b *TraitBuilder) Properties(properties map[string]interface{}) *TraitBuilder {
	for k, v := range properties {
		b.trait.Properties[k] = v
	}
	return b
}

// Binding

type ServerHandler struct {
	Name       string `json:"name"`
	Parameters any    `json:"parameters"`
}

// layer 2

type StackComponentBuilder struct {
	*InnerComponentBuilder[*StackComponentBuilder]
}

func (b *AppBuilder) NewStack() *StackComponentBuilder {
	t := &StackComponentBuilder{
		InnerComponentBuilder: newInnerComponent[*StackComponentBuilder](b),
	}
	t.inner = t
	return t.Type("core/v1/stack")
}

type TextComponentBuilder struct {
	*InnerComponentBuilder[*TextComponentBuilder]
}

func (b *AppBuilder) NewText() *TextComponentBuilder {
	t := &TextComponentBuilder{
		InnerComponentBuilder: newInnerComponent[*TextComponentBuilder](b),
	}
	t.inner = t
	return t.Type("core/v1/text")
}

func (b *TextComponentBuilder) Content(value string) *TextComponentBuilder {
	b.Properties(map[string]interface{}{
		"value": map[string]interface{}{
			"raw": value,
		}})
	return b
}

// layer 3

// chakra-ui

type ChakraUIAppBuilder struct {
	*AppBuilder
}

func NewChakraUIApp() *ChakraUIAppBuilder {
	b := &ChakraUIAppBuilder{
		AppBuilder: NewApp(),
	}
	b.component(
		b.NewComponent().Type("chakra_ui/v1/root").Id("root"),
	)
	return b
}

func (b *ChakraUIAppBuilder) Component(builder BaseComponentBuilder) *ChakraUIAppBuilder {
	if !componentAttachedToSlot(builder) {
		builder._Trait(b.NewTrait().
			Type("core/v1/slot").Properties(map[string]interface{}{
			"container": map[string]interface{}{
				"id":   "root",
				"slot": "root",
			},
		}))
	}
	b.component(builder)
	return b
}

func (b *ChakraUIAppBuilder) NewInput() *ComponentBuilder {
	t := &ComponentBuilder{
		InnerComponentBuilder: newInnerComponent[*ComponentBuilder](b.AppBuilder),
	}
	t.inner = t
	return t.Type("chakra_ui/v1/input")
}

type ChakraTableComponentBuilder struct {
	*InnerComponentBuilder[*ChakraTableComponentBuilder]
}

func (b *ChakraUIAppBuilder) NewTable() *ChakraTableComponentBuilder {
	t := &ChakraTableComponentBuilder{
		InnerComponentBuilder: newInnerComponent[*ChakraTableComponentBuilder](b.AppBuilder),
	}
	t.inner = t
	return t.Type("chakra_ui/v1/table").Properties(map[string]interface{}{
		"rowsPerPage": 20,
		"majorKey":    "name",
		"data":        []interface{}{},
		"columns":     []interface{}{},
	})
}

func (b *ChakraTableComponentBuilder) Data(data any) *ChakraTableComponentBuilder {
	b.Properties(map[string]interface{}{
		"data": data,
	})
	return b
}

type ChakraTableColumn struct {
	Key   string `json:"key"`
	Title string `json:"title"`
	Type  string `json:"type,omitempty"`
}

func (b *ChakraTableComponentBuilder) Column(column *ChakraTableColumn) *ChakraTableComponentBuilder {
	columns := b.ValueOf().Properties["columns"].([]interface{})
	columns = append(columns, column)
	b.Properties(map[string]interface{}{
		"columns": columns,
	})
	return b
}

type ChakraButtonComponentBuilder struct {
	*InnerComponentBuilder[*ChakraButtonComponentBuilder]
}

func (b *ChakraUIAppBuilder) NewButton() *ChakraButtonComponentBuilder {
	t := &ChakraButtonComponentBuilder{
		InnerComponentBuilder: newInnerComponent[*ChakraButtonComponentBuilder](b.AppBuilder),
	}
	t.inner = t
	return t.Type("chakra_ui/v1/button")
}

func (b *ChakraButtonComponentBuilder) Content(value string) *ChakraButtonComponentBuilder {
	b.Properties(map[string]interface{}{
		"text": map[string]interface{}{
			"raw": value,
		}})
	return b
}

func (b *ChakraButtonComponentBuilder) OnClick(serverHandler *ServerHandler) *ChakraButtonComponentBuilder {
	b._Trait(b.appBuilder.NewTrait().Type("core/v1/event").Properties(map[string]interface{}{
		"handlers": []map[string]interface{}{
			{
				"type":        "onClick",
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

type ChakraLinkComponentBuilder struct {
	*InnerComponentBuilder[*ChakraLinkComponentBuilder]
}

func (b *ChakraUIAppBuilder) NewLink() *ChakraLinkComponentBuilder {
	t := &ChakraLinkComponentBuilder{
		InnerComponentBuilder: newInnerComponent[*ChakraLinkComponentBuilder](b.AppBuilder),
	}
	t.inner = t
	return t.Type("chakra_ui/v1/link")
}

func (b *ChakraLinkComponentBuilder) Content(value string) *ChakraLinkComponentBuilder {
	b.Properties(map[string]interface{}{
		"text": map[string]interface{}{
			"raw": value,
		}})
	return b
}

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
