package sunmao

type Application struct {
	Version  string          `json:"version"`
	Kind     string          `json:"kind"`
	Metadata Metadata        `json:"metadata"`
	Spec     ApplicationSpec `json:"spec"`
}

type Metadata struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Annotations map[string]string `json:"annotations"`
}

type ApplicationSpec struct {
	Components []ComponentSchema `json:"components"`
}

type ComponentSchema struct {
	Id         string                 `json:"id"`
	Type       string                 `json:"type"`
	Properties map[string]interface{} `json:"properties"`
	Traits     []TraitSchema          `json:"traits"`
}

type TraitSchema struct {
	Type       string                 `json:"type"`
	Properties map[string]interface{} `json:"properties"`
}
