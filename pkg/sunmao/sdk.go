package sunmao

import (
	"fmt"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

// layer 1

// App

type AppBuilder struct {
	application Application
}

func NewApp() *AppBuilder {
	return &AppBuilder{application: Application{
		Version: "example/v1",
		Kind:    "Application",
		Metadata: Metadata{
			Name:        "empty",
			Description: "",
			Annotations: map[string]string{},
		},
		Spec: ApplicationSpec{
			Components: []ComponentSchema{},
		},
	}}
}

func newInnerComponent[K any](builder *AppBuilder) *InnerComponentBuilder[K] {
	id, _ := gonanoid.Generate("abcdefghijklmn_", 6)
	return &InnerComponentBuilder[K]{
		component: ComponentSchema{
			Id:         id,
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

func (b *AppBuilder) Version(version string) *AppBuilder {
	b.application.Version = version
	return b
}

func (b *AppBuilder) Name(name string) *AppBuilder {
	b.application.Metadata.Name = name
	return b
}

func (b *AppBuilder) Description(description string) *AppBuilder {
	b.application.Metadata.Description = description
	return b
}

func (b *AppBuilder) Annotation(key string, value string) *AppBuilder {
	b.application.Metadata.Annotations[key] = value
	return b
}

func (b *AppBuilder) Component(builder BaseComponentBuilder) *AppBuilder {
	b.component(builder)
	return b
}

func (b *AppBuilder) component(builder BaseComponentBuilder) {
	b.application.Spec.Components = append(b.application.Spec.Components, builder.ValueOf())
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

type ServerHandler struct {
	Name       string `json:"name"`
	Parameters any    `json:"parameters"`
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
