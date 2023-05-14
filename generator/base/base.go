package base

type StructureDefinition struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Snapshot    Snapshot `json:"snapshot"`
	Kind        string   `json:"kind"`
	Type        string   `json:"type"`
}

type Snapshot struct {
	Element []ElementDefinition `json:"element"`
}

type ElementDefinition struct {
	Path       string                  `json:"path"`
	Type       []ElementDefinitionType `json:"type"`
	Definition string                  `json:"definition"`
	Min        int                     `json:"min"`
	Max        string                  `json:"max"`
}

type ElementDefinitionType struct {
	Code string `json:"code"`
}

type bundleElement interface{}

type Bundle[E bundleElement] struct {
	ResourceType string           `json:"resourceType"`
	Entry        []BundleEntry[E] `json:"entry"`
}

type BundleEntry[E bundleElement] struct {
	Resource E `json:"resource"`
}
