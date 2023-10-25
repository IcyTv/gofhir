package base

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
)

type StructureDefinition struct {
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	Snapshot       Snapshot `json:"snapshot"`
	Kind           string   `json:"kind"`
	Type           string   `json:"type"`
	BaseDefinition string   `json:"baseDefinition"`
	Abstract       bool     `json:"abstract"`
	Status         string   `json:"status"`
}

func (sd StructureDefinition) GetResourceType() string {
	return "StructureDefinition"
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
	Code      string                           `json:"code"`
	Extension []ElementDefinitionTypeExtension `json:"extension"`
}

type ElementDefinitionTypeExtension struct {
	ValueUrl string `json:"valueUrl"`
}

type bundleElement interface {
}

type Bundle[E bundleElement] struct {
	ResourceType string           `json:"resourceType"`
	Entry        []BundleEntry[E] `json:"entry"`
}

type BundleEntry[E bundleElement] struct {
	Resource E      `json:"resource"`
	Url      string `json:"fullUrl"`
}

func (b *Bundle[E]) UnmarshalJSON(data []byte) error {
	// Only try unmarshalling entries that can be unmarshalled into the type E
	var out map[string]json.RawMessage
	if err := json.Unmarshal(data, &out); err != nil {
		return err
	}

	// If the resourceType is not Bundle, then we can't unmarshal it
	var resourceType string
	if err := json.Unmarshal(out["resourceType"], &resourceType); err != nil {
		return fmt.Errorf("failed to unmarshal resourceType: %w", err)
	}
	if resourceType != "Bundle" {
		return fmt.Errorf("expected resourceType to be Bundle, got %s", resourceType)
	}

	// Unmarshal the entries
	var entries []json.RawMessage
	if err := json.Unmarshal(out["entry"], &entries); err != nil {
		return fmt.Errorf("failed to unmarshal entries: %w", err)
	}

	expectedResourceType := reflect.TypeOf((*E)(nil)).Elem().Name()

	log.Println("expectedResourceType", expectedResourceType)

	// Unmarshal each entry into the type E
	var bundleEntries []BundleEntry[E] = make([]BundleEntry[E], 0, len(entries))
	for _, entry := range entries {
		var bundleEntry BundleEntry[E]
		if err := json.Unmarshal(entry, &bundleEntry); err != nil {
			continue
		}

		resourceType := strings.TrimPrefix(bundleEntry.Url, "http://hl7.org/fhir/")
		resourceType = strings.Split(resourceType, "/")[0]

		if resourceType != expectedResourceType {
			continue
		}

		bundleEntries = append(bundleEntries, bundleEntry)
	}

	// Set the entries
	b.Entry = bundleEntries
	b.ResourceType = resourceType

	return nil
}

type DefinitionElement interface{}

type CapabilityStatement struct {
	Name   string                    `json:"name"`
	Format []string                  `json:"format"`
	Rest   []CapabilityStatementRest `json:"rest"`
}

type CapabilityStatementRest struct {
	Mode     string                            `json:"mode"`
	Resource []CapabilityStatementRestResource `json:"resource"`
}

type CapabilityStatementRestResource struct {
	Type          string                                       `json:"type"`
	SearchInclude []string                                     `json:"searchInclude"`
	SearchParam   []CapabilityStatementRestResourceSearchParam `json:"searchParam"`
	Interaction   []CapabilityStatementRestResourceInteraction `json:"interaction"`
	Operation     []CapabilityStatementRestResourceOperation   `json:"operation"`
}

type CapabilityStatementRestResourceSearchParam struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type CapabilityStatementRestResourceInteraction struct {
	Code string `json:"code"`
}

type CapabilityStatementRestResourceOperation struct {
	Name       string `json:"name"`
	Definition string `json:"definition"`
}
