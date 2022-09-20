package sunmao

type VersionMetadataSetter interface {
	SetVersion(v string)
	SetName(n string)
	SetDescription(d string)
	SetAnnotations(k, v string)
}

type VersionMetadata struct {
	Version  string   `json:"version"`
	Metadata Metadata `json:"metadata"`
}

func (m *VersionMetadata) SetVersion(v string)        { m.Version = v }
func (m *VersionMetadata) SetName(n string)           { m.Metadata.Name = n }
func (m *VersionMetadata) SetDescription(d string)    { m.Metadata.Description = d }
func (m *VersionMetadata) SetAnnotations(k, v string) { m.Metadata.Annotations[k] = v }

type Application struct {
	Kind             string `json:"kind"`
	*VersionMetadata `json:",inline"`
	Spec             ApplicationSpec `json:"spec"`
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

type ModuleSpec struct {
	StateMap   map[string]interface{} `json:"stateMap"`
	Properties map[string]interface{} `json:"properties"`
}

type Module struct {
	Kind             string `json:"kind"`
	*VersionMetadata `json:",inline"`
	Spec             ModuleSpec        `json:"spec"`
	Impl             []ComponentSchema `json:"impl"`
}

type ModuleContainer struct {
	Id         string         `json:"id"`
	Type       string         `json:"type"`
	Properties map[string]any `json:"properties"`
}
