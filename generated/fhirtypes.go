package generated

import (
	"fmt"
	uuid "github.com/google/uuid"
	go1 "github.com/json-iterator/go"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
	"net/url"
	"time"
)

func (out *Element) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Element\"" {
		return fmt.Errorf("resourceType is not %s", "Element")
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	return nil
}

type Element struct {
	Id                   string              `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension            []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	ResourceType         string              `bson:"-" json:"resourceType,omitempty"`
}

func (out *BackboneElement) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"BackboneElement\"" {
		return fmt.Errorf("resourceType is not %s", "BackboneElement")
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["modifierExtension"]) > 0 {
		if err := go1.Unmarshal(asMap["modifierExtension"], &out.ModifierExtension); err != nil {
			return err
		}

	}
	return nil
}

type BackboneElement struct {
	Id                   *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension            []*Extension        `bson:",omitempty" json:"extension,omitempty"`         // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	ModifierExtension    []*Extension        `bson:",omitempty" json:"modifierExtension,omitempty"` /*
	May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.

	Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	*/
	ResourceType string `bson:"-" json:"resourceType,omitempty"`
}

func (out *Address) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Address\"" {
		return fmt.Errorf("resourceType is not %s", "Address")
	}
	if len(asMap["text"]) > 0 {
		if err := go1.Unmarshal(asMap["text"], &out.Text); err != nil {
			return err
		}

	}
	if len(asMap["_text"]) > 0 {
		if err := go1.Unmarshal(asMap["_text"], &out.TextPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["line"]) > 0 {
		if err := go1.Unmarshal(asMap["line"], &out.Line); err != nil {
			return err
		}

	}
	if len(asMap["_line"]) > 0 {
		if err := go1.Unmarshal(asMap["_line"], &out.LinePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["district"]) > 0 {
		if err := go1.Unmarshal(asMap["district"], &out.District); err != nil {
			return err
		}

	}
	if len(asMap["_district"]) > 0 {
		if err := go1.Unmarshal(asMap["_district"], &out.DistrictPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["period"]) > 0 {
		if err := go1.Unmarshal(asMap["period"], &out.Period); err != nil {
			return err
		}

	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["use"]) > 0 {
		if err := go1.Unmarshal(asMap["use"], &out.Use); err != nil {
			return err
		}

	}
	if len(asMap["_use"]) > 0 {
		if err := go1.Unmarshal(asMap["_use"], &out.UsePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["city"]) > 0 {
		if err := go1.Unmarshal(asMap["city"], &out.City); err != nil {
			return err
		}

	}
	if len(asMap["_city"]) > 0 {
		if err := go1.Unmarshal(asMap["_city"], &out.CityPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["state"]) > 0 {
		if err := go1.Unmarshal(asMap["state"], &out.State); err != nil {
			return err
		}

	}
	if len(asMap["_state"]) > 0 {
		if err := go1.Unmarshal(asMap["_state"], &out.StatePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["postalCode"]) > 0 {
		if err := go1.Unmarshal(asMap["postalCode"], &out.PostalCode); err != nil {
			return err
		}

	}
	if len(asMap["_postalCode"]) > 0 {
		if err := go1.Unmarshal(asMap["_postalCode"], &out.PostalCodePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["country"]) > 0 {
		if err := go1.Unmarshal(asMap["country"], &out.Country); err != nil {
			return err
		}

	}
	if len(asMap["_country"]) > 0 {
		if err := go1.Unmarshal(asMap["_country"], &out.CountryPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["type"]) > 0 {
		if err := go1.Unmarshal(asMap["type"], &out.Type); err != nil {
			return err
		}

	}
	if len(asMap["_type"]) > 0 {
		if err := go1.Unmarshal(asMap["_type"], &out.TypePrimitiveExtension); err != nil {
			return err
		}
	}

	return nil
}

type Address struct {
	Text                         string              `bson:",omitempty" json:"text,omitempty"` // Specifies the entire address as it should be displayed e.g. on a postal label. This may be provided instead of or as well as the specific parts.
	TextPrimitiveExtension       *PrimitiveExtension `bson:"text_extension,omitempty" json:"_text,omitempty"`
	Line                         []string            `bson:",omitempty" json:"line,omitempty"` // This component contains the house number, apartment number, street name, street direction,  P.O. Box number, delivery hints, and similar address information.
	LinePrimitiveExtension       *PrimitiveExtension `bson:"line_extension,omitempty" json:"_line,omitempty"`
	District                     string              `bson:",omitempty" json:"district,omitempty"` // The name of the administrative area (county).
	DistrictPrimitiveExtension   *PrimitiveExtension `bson:"district_extension,omitempty" json:"_district,omitempty"`
	Period                       *Period             `bson:",omitempty" json:"period,omitempty"` // Time period when address was/is in use.
	Id                           *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension         *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Use                          string              `bson:",omitempty" json:"use,omitempty"` // The purpose of this address.
	UsePrimitiveExtension        *PrimitiveExtension `bson:"use_extension,omitempty" json:"_use,omitempty"`
	City                         string              `bson:",omitempty" json:"city,omitempty"` // The name of the city, town, suburb, village or other community or delivery center.
	CityPrimitiveExtension       *PrimitiveExtension `bson:"city_extension,omitempty" json:"_city,omitempty"`
	State                        string              `bson:",omitempty" json:"state,omitempty"` // Sub-unit of a country with limited sovereignty in a federally organized country. A code may be used if codes are in common use (e.g. US 2 letter state codes).
	StatePrimitiveExtension      *PrimitiveExtension `bson:"state_extension,omitempty" json:"_state,omitempty"`
	PostalCode                   string              `bson:",omitempty" json:"postalCode,omitempty"` // A postal code designating a region defined by the postal service.
	PostalCodePrimitiveExtension *PrimitiveExtension `bson:"postalCode_extension,omitempty" json:"_postalCode,omitempty"`
	Country                      string              `bson:",omitempty" json:"country,omitempty"` // Country - a nation as commonly understood or generally accepted.
	CountryPrimitiveExtension    *PrimitiveExtension `bson:"country_extension,omitempty" json:"_country,omitempty"`
	Extension                    []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Type                         string              `bson:",omitempty" json:"type,omitempty"`      // Distinguishes between physical addresses (those you can visit) and mailing addresses (e.g. PO Boxes and care-of addresses). Most addresses are both.
	TypePrimitiveExtension       *PrimitiveExtension `bson:"type_extension,omitempty" json:"_type,omitempty"`
	ResourceType                 string              `bson:"-" json:"resourceType,omitempty"`
}

func (out *Age) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Age\"" {
		return fmt.Errorf("resourceType is not %s", "Age")
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["value"]) > 0 {
		if err := go1.Unmarshal(asMap["value"], &out.Value); err != nil {
			return err
		}

	}
	if len(asMap["_value"]) > 0 {
		if err := go1.Unmarshal(asMap["_value"], &out.ValuePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["comparator"]) > 0 {
		if err := go1.Unmarshal(asMap["comparator"], &out.Comparator); err != nil {
			return err
		}

	}
	if len(asMap["_comparator"]) > 0 {
		if err := go1.Unmarshal(asMap["_comparator"], &out.ComparatorPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["unit"]) > 0 {
		if err := go1.Unmarshal(asMap["unit"], &out.Unit); err != nil {
			return err
		}

	}
	if len(asMap["_unit"]) > 0 {
		if err := go1.Unmarshal(asMap["_unit"], &out.UnitPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["system"]) > 0 {
		if err := go1.Unmarshal(asMap["system"], &out.System); err != nil {
			return err
		}

	}
	if len(asMap["_system"]) > 0 {
		if err := go1.Unmarshal(asMap["_system"], &out.SystemPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["code"]) > 0 {
		if err := go1.Unmarshal(asMap["code"], &out.Code); err != nil {
			return err
		}

	}
	if len(asMap["_code"]) > 0 {
		if err := go1.Unmarshal(asMap["_code"], &out.CodePrimitiveExtension); err != nil {
			return err
		}
	}

	return nil
}

type Age struct {
	Id                           *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension         *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension                    []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Value                        float64             `bson:",omitempty" json:"value,omitempty"`     // The value of the measured amount. The value includes an implicit precision in the presentation of the value.
	ValuePrimitiveExtension      *PrimitiveExtension `bson:"value_extension,omitempty" json:"_value,omitempty"`
	Comparator                   string              `bson:",omitempty" json:"comparator,omitempty"` // How the value should be understood and represented - whether the actual value is greater or less than the stated value due to measurement issues; e.g. if the comparator is "<" , then the real value is < stated value.
	ComparatorPrimitiveExtension *PrimitiveExtension `bson:"comparator_extension,omitempty" json:"_comparator,omitempty"`
	Unit                         string              `bson:",omitempty" json:"unit,omitempty"` // A human-readable form of the unit.
	UnitPrimitiveExtension       *PrimitiveExtension `bson:"unit_extension,omitempty" json:"_unit,omitempty"`
	System                       string              `bson:",omitempty" json:"system,omitempty"` // The identification of the system that provides the coded form of the unit.
	SystemPrimitiveExtension     *PrimitiveExtension `bson:"system_extension,omitempty" json:"_system,omitempty"`
	Code                         string              `bson:",omitempty" json:"code,omitempty"` // A computer processable form of the unit in some unit representation system.
	CodePrimitiveExtension       *PrimitiveExtension `bson:"code_extension,omitempty" json:"_code,omitempty"`
	ResourceType                 string              `bson:"-" json:"resourceType,omitempty"`
}

type AnnotationAuthorx struct {
	AuthorReference *Reference `binding:"omitempty" bson:",omitempty" json:"authorReference,omitempty"`
	AuthorString    string     `binding:"omitempty" bson:",omitempty" json:"authorString,omitempty"`
}

func (out *Annotation) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Annotation\"" {
		return fmt.Errorf("resourceType is not %s", "Annotation")
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["authorReference"]) > 0 {
		if err := go1.Unmarshal(asMap["authorReference"], &out.AuthorReference); err == nil {
		}
	} else if len(asMap["authorString"]) > 0 {
		if err := go1.Unmarshal(asMap["authorString"], &out.AuthorString); err == nil {
		}
	} else {

	}
	if len(asMap["time"]) > 0 {
		if err := go1.Unmarshal(asMap["time"], &out.Time); err != nil {
			return err
		}

	}
	if len(asMap["_time"]) > 0 {
		if err := go1.Unmarshal(asMap["_time"], &out.TimePrimitiveExtension); err != nil {
			return err
		}
	}

	if err := go1.Unmarshal(asMap["text"], &out.Text); err != nil {
		return err
	}
	if len(asMap["_text"]) > 0 {
		if err := go1.Unmarshal(asMap["_text"], &out.TextPrimitiveExtension); err != nil {
			return err
		}
	}

	return nil
}

type Annotation struct {
	Id                     *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension   *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension              []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	AnnotationAuthorx      `binding:"omitempty" bson:",omitempty" json:",omitempty"`
	Time                   *DateTime           `bson:",omitempty" json:"time,omitempty"` // Indicates when this particular annotation was made.
	TimePrimitiveExtension *PrimitiveExtension `bson:"time_extension,omitempty" json:"_time,omitempty"`
	Text                   string              `binding:"required" bson:",omitempty" json:"text,omitempty"` // The text of the annotation in markdown format.
	TextPrimitiveExtension *PrimitiveExtension `bson:"text_extension,omitempty" json:"_text,omitempty"`
	ResourceType           string              `bson:"-" json:"resourceType,omitempty"`
}

func (out *Attachment) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Attachment\"" {
		return fmt.Errorf("resourceType is not %s", "Attachment")
	}
	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["title"]) > 0 {
		if err := go1.Unmarshal(asMap["title"], &out.Title); err != nil {
			return err
		}

	}
	if len(asMap["_title"]) > 0 {
		if err := go1.Unmarshal(asMap["_title"], &out.TitlePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["height"]) > 0 {
		if err := go1.Unmarshal(asMap["height"], &out.Height); err != nil {
			return err
		}

	}
	if len(asMap["_height"]) > 0 {
		if err := go1.Unmarshal(asMap["_height"], &out.HeightPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["pages"]) > 0 {
		if err := go1.Unmarshal(asMap["pages"], &out.Pages); err != nil {
			return err
		}

	}
	if len(asMap["_pages"]) > 0 {
		if err := go1.Unmarshal(asMap["_pages"], &out.PagesPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["width"]) > 0 {
		if err := go1.Unmarshal(asMap["width"], &out.Width); err != nil {
			return err
		}

	}
	if len(asMap["_width"]) > 0 {
		if err := go1.Unmarshal(asMap["_width"], &out.WidthPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["url"]) > 0 {
		if err := go1.Unmarshal(asMap["url"], &out.Url); err != nil {
			return err
		}

	}
	if len(asMap["_url"]) > 0 {
		if err := go1.Unmarshal(asMap["_url"], &out.UrlPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["creation"]) > 0 {
		if err := go1.Unmarshal(asMap["creation"], &out.Creation); err != nil {
			return err
		}

	}
	if len(asMap["_creation"]) > 0 {
		if err := go1.Unmarshal(asMap["_creation"], &out.CreationPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["language"]) > 0 {
		if err := go1.Unmarshal(asMap["language"], &out.Language); err != nil {
			return err
		}

	}
	if len(asMap["_language"]) > 0 {
		if err := go1.Unmarshal(asMap["_language"], &out.LanguagePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["data"]) > 0 {
		if err := go1.Unmarshal(asMap["data"], &out.Data); err != nil {
			return err
		}

	}
	if len(asMap["_data"]) > 0 {
		if err := go1.Unmarshal(asMap["_data"], &out.DataPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["size"]) > 0 {
		if err := go1.Unmarshal(asMap["size"], &out.Size); err != nil {
			return err
		}

	}
	if len(asMap["_size"]) > 0 {
		if err := go1.Unmarshal(asMap["_size"], &out.SizePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["hash"]) > 0 {
		if err := go1.Unmarshal(asMap["hash"], &out.Hash); err != nil {
			return err
		}

	}
	if len(asMap["_hash"]) > 0 {
		if err := go1.Unmarshal(asMap["_hash"], &out.HashPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["frames"]) > 0 {
		if err := go1.Unmarshal(asMap["frames"], &out.Frames); err != nil {
			return err
		}

	}
	if len(asMap["_frames"]) > 0 {
		if err := go1.Unmarshal(asMap["_frames"], &out.FramesPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["duration"]) > 0 {
		if err := go1.Unmarshal(asMap["duration"], &out.Duration); err != nil {
			return err
		}

	}
	if len(asMap["_duration"]) > 0 {
		if err := go1.Unmarshal(asMap["_duration"], &out.DurationPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["contentType"]) > 0 {
		if err := go1.Unmarshal(asMap["contentType"], &out.ContentType); err != nil {
			return err
		}

	}
	if len(asMap["_contentType"]) > 0 {
		if err := go1.Unmarshal(asMap["_contentType"], &out.ContentTypePrimitiveExtension); err != nil {
			return err
		}
	}

	return nil
}

type Attachment struct {
	Extension                     []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Title                         string              `bson:",omitempty" json:"title,omitempty"`     // A label or set of text to display in place of the data.
	TitlePrimitiveExtension       *PrimitiveExtension `bson:"title_extension,omitempty" json:"_title,omitempty"`
	Height                        int                 `bson:",omitempty" json:"height,omitempty"` // Height of the image in pixels (photo/video).
	HeightPrimitiveExtension      *PrimitiveExtension `bson:"height_extension,omitempty" json:"_height,omitempty"`
	Pages                         int                 `bson:",omitempty" json:"pages,omitempty"` // The number of pages when printed.
	PagesPrimitiveExtension       *PrimitiveExtension `bson:"pages_extension,omitempty" json:"_pages,omitempty"`
	Width                         int                 `bson:",omitempty" json:"width,omitempty"` // Width of the image in pixels (photo/video).
	WidthPrimitiveExtension       *PrimitiveExtension `bson:"width_extension,omitempty" json:"_width,omitempty"`
	Url                           *url.URL            `bson:",omitempty" json:"url,omitempty"` // A location where the data can be accessed.
	UrlPrimitiveExtension         *PrimitiveExtension `bson:"url_extension,omitempty" json:"_url,omitempty"`
	Creation                      *DateTime           `bson:",omitempty" json:"creation,omitempty"` // The date that the attachment was first created.
	CreationPrimitiveExtension    *PrimitiveExtension `bson:"creation_extension,omitempty" json:"_creation,omitempty"`
	Language                      string              `bson:",omitempty" json:"language,omitempty"` // The human language of the content. The value can be any valid value according to BCP 47.
	LanguagePrimitiveExtension    *PrimitiveExtension `bson:"language_extension,omitempty" json:"_language,omitempty"`
	Data                          Base64Binary        `bson:",omitempty" json:"data,omitempty"` // The actual data of the attachment - a sequence of bytes, base64 encoded.
	DataPrimitiveExtension        *PrimitiveExtension `bson:"data_extension,omitempty" json:"_data,omitempty"`
	Size                          int64               `bson:",omitempty" json:"size,omitempty"` // The number of bytes of data that make up this attachment (before base64 encoding, if that is done).
	SizePrimitiveExtension        *PrimitiveExtension `bson:"size_extension,omitempty" json:"_size,omitempty"`
	Hash                          Base64Binary        `bson:",omitempty" json:"hash,omitempty"` // The calculated hash of the data using SHA-1. Represented using base64.
	HashPrimitiveExtension        *PrimitiveExtension `bson:"hash_extension,omitempty" json:"_hash,omitempty"`
	Frames                        int                 `bson:",omitempty" json:"frames,omitempty"` // The number of frames in a photo. This is used with a multi-page fax, or an imaging acquisition context that takes multiple slices in a single image, or an animated gif. If there is more than one frame, this SHALL have a value in order to alert interface software that a multi-frame capable rendering widget is required.
	FramesPrimitiveExtension      *PrimitiveExtension `bson:"frames_extension,omitempty" json:"_frames,omitempty"`
	Duration                      float64             `bson:",omitempty" json:"duration,omitempty"` // The duration of the recording in seconds - for audio and video.
	DurationPrimitiveExtension    *PrimitiveExtension `bson:"duration_extension,omitempty" json:"_duration,omitempty"`
	Id                            *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension          *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	ContentType                   string              `bson:",omitempty" json:"contentType,omitempty"` // Identifies the type of the data in the attachment and allows a method to be chosen to interpret or render the data. Includes mime type parameters such as charset where appropriate.
	ContentTypePrimitiveExtension *PrimitiveExtension `bson:"contentType_extension,omitempty" json:"_contentType,omitempty"`
	ResourceType                  string              `bson:"-" json:"resourceType,omitempty"`
}

func (out *AvailabilityAvailableTime) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["daysOfWeek"]) > 0 {
		if err := go1.Unmarshal(asMap["daysOfWeek"], &out.DaysOfWeek); err != nil {
			return err
		}

	}
	if len(asMap["_daysOfWeek"]) > 0 {
		if err := go1.Unmarshal(asMap["_daysOfWeek"], &out.DaysOfWeekPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["allDay"]) > 0 {
		if err := go1.Unmarshal(asMap["allDay"], &out.AllDay); err != nil {
			return err
		}

	}
	if len(asMap["_allDay"]) > 0 {
		if err := go1.Unmarshal(asMap["_allDay"], &out.AllDayPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["availableStartTime"]) > 0 {
		if err := go1.Unmarshal(asMap["availableStartTime"], &out.AvailableStartTime); err != nil {
			return err
		}

	}
	if len(asMap["_availableStartTime"]) > 0 {
		if err := go1.Unmarshal(asMap["_availableStartTime"], &out.AvailableStartTimePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["availableEndTime"]) > 0 {
		if err := go1.Unmarshal(asMap["availableEndTime"], &out.AvailableEndTime); err != nil {
			return err
		}

	}
	if len(asMap["_availableEndTime"]) > 0 {
		if err := go1.Unmarshal(asMap["_availableEndTime"], &out.AvailableEndTimePrimitiveExtension); err != nil {
			return err
		}
	}

	return nil
}

type AvailabilityAvailableTime struct {
	Id                                   string              `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension                 *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension                            []*Extension        `bson:",omitempty" json:"extension,omitempty"`  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	DaysOfWeek                           []string            `bson:",omitempty" json:"daysOfWeek,omitempty"` // mon | tue | wed | thu | fri | sat | sun.
	DaysOfWeekPrimitiveExtension         *PrimitiveExtension `bson:"daysOfWeek_extension,omitempty" json:"_daysOfWeek,omitempty"`
	AllDay                               bool                `bson:",omitempty" json:"allDay,omitempty"` // Always available? i.e. 24 hour service.
	AllDayPrimitiveExtension             *PrimitiveExtension `bson:"allDay_extension,omitempty" json:"_allDay,omitempty"`
	AvailableStartTime                   *Time               `bson:",omitempty" json:"availableStartTime,omitempty"` // Opening time of day (ignored if allDay = true).
	AvailableStartTimePrimitiveExtension *PrimitiveExtension `bson:"availableStartTime_extension,omitempty" json:"_availableStartTime,omitempty"`
	AvailableEndTime                     *Time               `bson:",omitempty" json:"availableEndTime,omitempty"` // Closing time of day (ignored if allDay = true).
	AvailableEndTimePrimitiveExtension   *PrimitiveExtension `bson:"availableEndTime_extension,omitempty" json:"_availableEndTime,omitempty"`
}

func (out *AvailabilityNotAvailableTime) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["description"]) > 0 {
		if err := go1.Unmarshal(asMap["description"], &out.Description); err != nil {
			return err
		}

	}
	if len(asMap["_description"]) > 0 {
		if err := go1.Unmarshal(asMap["_description"], &out.DescriptionPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["during"]) > 0 {
		if err := go1.Unmarshal(asMap["during"], &out.During); err != nil {
			return err
		}

	}
	return nil
}

type AvailabilityNotAvailableTime struct {
	Id                            string              `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension          *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension                     []*Extension        `bson:",omitempty" json:"extension,omitempty"`   // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Description                   string              `bson:",omitempty" json:"description,omitempty"` // Reason presented to the user explaining why time not available.
	DescriptionPrimitiveExtension *PrimitiveExtension `bson:"description_extension,omitempty" json:"_description,omitempty"`
	During                        *Period             `bson:",omitempty" json:"during,omitempty"` // Service not available during this period.
}

func (out *Availability) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Availability\"" {
		return fmt.Errorf("resourceType is not %s", "Availability")
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["availableTime"]) > 0 {
		if err := go1.Unmarshal(asMap["availableTime"], &out.AvailableTime); err != nil {
			return err
		}

	}
	if len(asMap["notAvailableTime"]) > 0 {
		if err := go1.Unmarshal(asMap["notAvailableTime"], &out.NotAvailableTime); err != nil {
			return err
		}

	}
	return nil
}

type Availability struct {
	Id                   *primitive.ObjectID           `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension *PrimitiveExtension           `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension            []*Extension                  `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	AvailableTime        *AvailabilityAvailableTime    `bson:",omitempty" json:"availableTime,omitempty"`
	NotAvailableTime     *AvailabilityNotAvailableTime `bson:",omitempty" json:"notAvailableTime,omitempty"`
	ResourceType         string                        `bson:"-" json:"resourceType,omitempty"`
}

func (out *BackboneType) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"BackboneType\"" {
		return fmt.Errorf("resourceType is not %s", "BackboneType")
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["modifierExtension"]) > 0 {
		if err := go1.Unmarshal(asMap["modifierExtension"], &out.ModifierExtension); err != nil {
			return err
		}

	}
	return nil
}

type BackboneType struct {
	Id                   *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension            []*Extension        `bson:",omitempty" json:"extension,omitempty"`         // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	ModifierExtension    []*Extension        `bson:",omitempty" json:"modifierExtension,omitempty"` /*
	May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.

	Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	*/
	ResourceType string `bson:"-" json:"resourceType,omitempty"`
}

func (out *CodeableConcept) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"CodeableConcept\"" {
		return fmt.Errorf("resourceType is not %s", "CodeableConcept")
	}
	if len(asMap["text"]) > 0 {
		if err := go1.Unmarshal(asMap["text"], &out.Text); err != nil {
			return err
		}

	}
	if len(asMap["_text"]) > 0 {
		if err := go1.Unmarshal(asMap["_text"], &out.TextPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["coding"]) > 0 {
		if err := go1.Unmarshal(asMap["coding"], &out.Coding); err != nil {
			return err
		}

	}
	return nil
}

type CodeableConcept struct {
	Text                   string              `bson:",omitempty" json:"text,omitempty"` // A human language representation of the concept as seen/selected/uttered by the user who entered the data and/or which represents the intended meaning of the user.
	TextPrimitiveExtension *PrimitiveExtension `bson:"text_extension,omitempty" json:"_text,omitempty"`
	Id                     *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension   *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension              []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Coding                 []*Coding           `bson:",omitempty" json:"coding,omitempty"`    // A reference to a code defined by a terminology system.
	ResourceType           string              `bson:"-" json:"resourceType,omitempty"`
}

func (out *CodeableReference) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"CodeableReference\"" {
		return fmt.Errorf("resourceType is not %s", "CodeableReference")
	}
	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["concept"]) > 0 {
		if err := go1.Unmarshal(asMap["concept"], &out.Concept); err != nil {
			return err
		}

	}
	if len(asMap["reference"]) > 0 {
		if err := go1.Unmarshal(asMap["reference"], &out.Reference); err != nil {
			return err
		}

	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	return nil
}

type CodeableReference struct {
	Extension            []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Concept              *CodeableConcept    `bson:",omitempty" json:"concept,omitempty"`   // A reference to a concept - e.g. the information is identified by its general class to the degree of precision found in the terminology.
	Reference            *Reference          `bson:",omitempty" json:"reference,omitempty"` // A reference to a resource the provides exact details about the information being referenced.
	Id                   *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`     // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	ResourceType         string              `bson:"-" json:"resourceType,omitempty"`
}

func (out *Coding) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Coding\"" {
		return fmt.Errorf("resourceType is not %s", "Coding")
	}
	if len(asMap["code"]) > 0 {
		if err := go1.Unmarshal(asMap["code"], &out.Code); err != nil {
			return err
		}

	}
	if len(asMap["_code"]) > 0 {
		if err := go1.Unmarshal(asMap["_code"], &out.CodePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["display"]) > 0 {
		if err := go1.Unmarshal(asMap["display"], &out.Display); err != nil {
			return err
		}

	}
	if len(asMap["_display"]) > 0 {
		if err := go1.Unmarshal(asMap["_display"], &out.DisplayPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["userSelected"]) > 0 {
		if err := go1.Unmarshal(asMap["userSelected"], &out.UserSelected); err != nil {
			return err
		}

	}
	if len(asMap["_userSelected"]) > 0 {
		if err := go1.Unmarshal(asMap["_userSelected"], &out.UserSelectedPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["system"]) > 0 {
		if err := go1.Unmarshal(asMap["system"], &out.System); err != nil {
			return err
		}

	}
	if len(asMap["_system"]) > 0 {
		if err := go1.Unmarshal(asMap["_system"], &out.SystemPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["version"]) > 0 {
		if err := go1.Unmarshal(asMap["version"], &out.Version); err != nil {
			return err
		}

	}
	if len(asMap["_version"]) > 0 {
		if err := go1.Unmarshal(asMap["_version"], &out.VersionPrimitiveExtension); err != nil {
			return err
		}
	}

	return nil
}

type Coding struct {
	Code                           string              `bson:",omitempty" json:"code,omitempty"` // A symbol in syntax defined by the system. The symbol may be a predefined code or an expression in a syntax defined by the coding system (e.g. post-coordination).
	CodePrimitiveExtension         *PrimitiveExtension `bson:"code_extension,omitempty" json:"_code,omitempty"`
	Display                        string              `bson:",omitempty" json:"display,omitempty"` // A representation of the meaning of the code in the system, following the rules of the system.
	DisplayPrimitiveExtension      *PrimitiveExtension `bson:"display_extension,omitempty" json:"_display,omitempty"`
	UserSelected                   bool                `bson:",omitempty" json:"userSelected,omitempty"` // Indicates that this coding was chosen by a user directly - e.g. off a pick list of available items (codes or displays).
	UserSelectedPrimitiveExtension *PrimitiveExtension `bson:"userSelected_extension,omitempty" json:"_userSelected,omitempty"`
	Id                             *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension           *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension                      []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	System                         string              `bson:",omitempty" json:"system,omitempty"`    // The identification of the code system that defines the meaning of the symbol in the code.
	SystemPrimitiveExtension       *PrimitiveExtension `bson:"system_extension,omitempty" json:"_system,omitempty"`
	Version                        string              `bson:",omitempty" json:"version,omitempty"` // The version of the code system which was used when choosing this code. Note that a well-maintained code system does not need the version reported, because the meaning of codes is consistent across versions. However this cannot consistently be assured, and when the meaning is not guaranteed to be consistent, the version SHOULD be exchanged.
	VersionPrimitiveExtension      *PrimitiveExtension `bson:"version_extension,omitempty" json:"_version,omitempty"`
	ResourceType                   string              `bson:"-" json:"resourceType,omitempty"`
}

func (out *ContactDetail) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"ContactDetail\"" {
		return fmt.Errorf("resourceType is not %s", "ContactDetail")
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["name"]) > 0 {
		if err := go1.Unmarshal(asMap["name"], &out.Name); err != nil {
			return err
		}

	}
	if len(asMap["_name"]) > 0 {
		if err := go1.Unmarshal(asMap["_name"], &out.NamePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["telecom"]) > 0 {
		if err := go1.Unmarshal(asMap["telecom"], &out.Telecom); err != nil {
			return err
		}

	}
	return nil
}

type ContactDetail struct {
	Id                     *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension   *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension              []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Name                   string              `bson:",omitempty" json:"name,omitempty"`      // The name of an individual to contact.
	NamePrimitiveExtension *PrimitiveExtension `bson:"name_extension,omitempty" json:"_name,omitempty"`
	Telecom                []*ContactPoint     `bson:",omitempty" json:"telecom,omitempty"` // The contact details for the individual (if a name was provided) or the organization.
	ResourceType           string              `bson:"-" json:"resourceType,omitempty"`
}

func (out *ContactPoint) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"ContactPoint\"" {
		return fmt.Errorf("resourceType is not %s", "ContactPoint")
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["system"]) > 0 {
		if err := go1.Unmarshal(asMap["system"], &out.System); err != nil {
			return err
		}

	}
	if len(asMap["_system"]) > 0 {
		if err := go1.Unmarshal(asMap["_system"], &out.SystemPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["value"]) > 0 {
		if err := go1.Unmarshal(asMap["value"], &out.Value); err != nil {
			return err
		}

	}
	if len(asMap["_value"]) > 0 {
		if err := go1.Unmarshal(asMap["_value"], &out.ValuePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["use"]) > 0 {
		if err := go1.Unmarshal(asMap["use"], &out.Use); err != nil {
			return err
		}

	}
	if len(asMap["_use"]) > 0 {
		if err := go1.Unmarshal(asMap["_use"], &out.UsePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["rank"]) > 0 {
		if err := go1.Unmarshal(asMap["rank"], &out.Rank); err != nil {
			return err
		}

	}
	if len(asMap["_rank"]) > 0 {
		if err := go1.Unmarshal(asMap["_rank"], &out.RankPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["period"]) > 0 {
		if err := go1.Unmarshal(asMap["period"], &out.Period); err != nil {
			return err
		}

	}
	return nil
}

type ContactPoint struct {
	Id                       *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension     *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension                []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	System                   string              `bson:",omitempty" json:"system,omitempty"`    // Telecommunications form for contact point - what communications system is required to make use of the contact.
	SystemPrimitiveExtension *PrimitiveExtension `bson:"system_extension,omitempty" json:"_system,omitempty"`
	Value                    string              `bson:",omitempty" json:"value,omitempty"` // The actual contact point details, in a form that is meaningful to the designated communication system (i.e. phone number or email address).
	ValuePrimitiveExtension  *PrimitiveExtension `bson:"value_extension,omitempty" json:"_value,omitempty"`
	Use                      string              `bson:",omitempty" json:"use,omitempty"` // Identifies the purpose for the contact point.
	UsePrimitiveExtension    *PrimitiveExtension `bson:"use_extension,omitempty" json:"_use,omitempty"`
	Rank                     int                 `bson:",omitempty" json:"rank,omitempty"` // Specifies a preferred order in which to use a set of contacts. ContactPoints with lower rank values are more preferred than those with higher rank values.
	RankPrimitiveExtension   *PrimitiveExtension `bson:"rank_extension,omitempty" json:"_rank,omitempty"`
	Period                   *Period             `bson:",omitempty" json:"period,omitempty"` // Time period when the contact point was/is in use.
	ResourceType             string              `bson:"-" json:"resourceType,omitempty"`
}

func (out *Contributor) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Contributor\"" {
		return fmt.Errorf("resourceType is not %s", "Contributor")
	}
	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["type"], &out.Type); err != nil {
		return err
	}
	if len(asMap["_type"]) > 0 {
		if err := go1.Unmarshal(asMap["_type"], &out.TypePrimitiveExtension); err != nil {
			return err
		}
	}

	if err := go1.Unmarshal(asMap["name"], &out.Name); err != nil {
		return err
	}
	if len(asMap["_name"]) > 0 {
		if err := go1.Unmarshal(asMap["_name"], &out.NamePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["contact"]) > 0 {
		if err := go1.Unmarshal(asMap["contact"], &out.Contact); err != nil {
			return err
		}

	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	return nil
}

type Contributor struct {
	Extension              []*Extension        `bson:",omitempty" json:"extension,omitempty"`               // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Type                   string              `binding:"required" bson:",omitempty" json:"type,omitempty"` // The type of contributor.
	TypePrimitiveExtension *PrimitiveExtension `bson:"type_extension,omitempty" json:"_type,omitempty"`
	Name                   string              `binding:"required" bson:",omitempty" json:"name,omitempty"` // The name of the individual or organization responsible for the contribution.
	NamePrimitiveExtension *PrimitiveExtension `bson:"name_extension,omitempty" json:"_name,omitempty"`
	Contact                []*ContactDetail    `bson:",omitempty" json:"contact,omitempty"` // Contact details to assist a user in finding and communicating with the contributor.
	Id                     *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`   // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension   *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	ResourceType           string              `bson:"-" json:"resourceType,omitempty"`
}

func (out *Count) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Count\"" {
		return fmt.Errorf("resourceType is not %s", "Count")
	}
	if len(asMap["system"]) > 0 {
		if err := go1.Unmarshal(asMap["system"], &out.System); err != nil {
			return err
		}

	}
	if len(asMap["_system"]) > 0 {
		if err := go1.Unmarshal(asMap["_system"], &out.SystemPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["code"]) > 0 {
		if err := go1.Unmarshal(asMap["code"], &out.Code); err != nil {
			return err
		}

	}
	if len(asMap["_code"]) > 0 {
		if err := go1.Unmarshal(asMap["_code"], &out.CodePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["value"]) > 0 {
		if err := go1.Unmarshal(asMap["value"], &out.Value); err != nil {
			return err
		}

	}
	if len(asMap["_value"]) > 0 {
		if err := go1.Unmarshal(asMap["_value"], &out.ValuePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["comparator"]) > 0 {
		if err := go1.Unmarshal(asMap["comparator"], &out.Comparator); err != nil {
			return err
		}

	}
	if len(asMap["_comparator"]) > 0 {
		if err := go1.Unmarshal(asMap["_comparator"], &out.ComparatorPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["unit"]) > 0 {
		if err := go1.Unmarshal(asMap["unit"], &out.Unit); err != nil {
			return err
		}

	}
	if len(asMap["_unit"]) > 0 {
		if err := go1.Unmarshal(asMap["_unit"], &out.UnitPrimitiveExtension); err != nil {
			return err
		}
	}

	return nil
}

type Count struct {
	System                       string              `bson:",omitempty" json:"system,omitempty"` // The identification of the system that provides the coded form of the unit.
	SystemPrimitiveExtension     *PrimitiveExtension `bson:"system_extension,omitempty" json:"_system,omitempty"`
	Code                         string              `bson:",omitempty" json:"code,omitempty"` // A computer processable form of the unit in some unit representation system.
	CodePrimitiveExtension       *PrimitiveExtension `bson:"code_extension,omitempty" json:"_code,omitempty"`
	Id                           *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension         *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension                    []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Value                        float64             `bson:",omitempty" json:"value,omitempty"`     // The value of the measured amount. The value includes an implicit precision in the presentation of the value.
	ValuePrimitiveExtension      *PrimitiveExtension `bson:"value_extension,omitempty" json:"_value,omitempty"`
	Comparator                   string              `bson:",omitempty" json:"comparator,omitempty"` // How the value should be understood and represented - whether the actual value is greater or less than the stated value due to measurement issues; e.g. if the comparator is "<" , then the real value is < stated value.
	ComparatorPrimitiveExtension *PrimitiveExtension `bson:"comparator_extension,omitempty" json:"_comparator,omitempty"`
	Unit                         string              `bson:",omitempty" json:"unit,omitempty"` // A human-readable form of the unit.
	UnitPrimitiveExtension       *PrimitiveExtension `bson:"unit_extension,omitempty" json:"_unit,omitempty"`
	ResourceType                 string              `bson:"-" json:"resourceType,omitempty"`
}

type DataRequirementDateFilterValuex struct {
	ValueDateTime *DateTime `binding:"omitempty" bson:",omitempty" json:"valueDateTime,omitempty"`
	ValuePeriod   *Period   `binding:"omitempty" bson:",omitempty" json:"valuePeriod,omitempty"`
	ValueDuration *Duration `binding:"omitempty" bson:",omitempty" json:"valueDuration,omitempty"`
}

func (out *DataRequirementDateFilter) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if len(asMap["searchParam"]) > 0 {
		if err := go1.Unmarshal(asMap["searchParam"], &out.SearchParam); err != nil {
			return err
		}

	}
	if len(asMap["_searchParam"]) > 0 {
		if err := go1.Unmarshal(asMap["_searchParam"], &out.SearchParamPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["valueDateTime"]) > 0 {
		if err := go1.Unmarshal(asMap["valueDateTime"], &out.ValueDateTime); err == nil {
		}
	} else if len(asMap["valuePeriod"]) > 0 {
		if err := go1.Unmarshal(asMap["valuePeriod"], &out.ValuePeriod); err == nil {
		}
	} else if len(asMap["valueDuration"]) > 0 {
		if err := go1.Unmarshal(asMap["valueDuration"], &out.ValueDuration); err == nil {
		}
	} else {

	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["path"]) > 0 {
		if err := go1.Unmarshal(asMap["path"], &out.Path); err != nil {
			return err
		}

	}
	if len(asMap["_path"]) > 0 {
		if err := go1.Unmarshal(asMap["_path"], &out.PathPrimitiveExtension); err != nil {
			return err
		}
	}

	return nil
}

type DataRequirementDateFilter struct {
	SearchParam                     string              `bson:",omitempty" json:"searchParam,omitempty"` // A date parameter that refers to a search parameter defined on the specified type of the DataRequirement, and which searches on elements of type date, dateTime, Period, Schedule, or Timing.
	SearchParamPrimitiveExtension   *PrimitiveExtension `bson:"searchParam_extension,omitempty" json:"_searchParam,omitempty"`
	DataRequirementDateFilterValuex `binding:"omitempty" bson:",omitempty" json:",omitempty"`
	Id                              string              `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension            *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension                       []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Path                            string              `bson:",omitempty" json:"path,omitempty"`      // The date-valued attribute of the filter. The specified path SHALL be a FHIRPath resolvable on the specified type of the DataRequirement, and SHALL consist only of identifiers, constant indexers, and .resolve(). The path is allowed to contain qualifiers (.) to traverse sub-elements, as well as indexers ([x]) to traverse multiple-cardinality sub-elements (see the [Simple FHIRPath Profile](fhirpath.html#simple) for full details). Note that the index must be an integer constant. The path must resolve to an element of type date, dateTime, Period, Schedule, or Timing.
	PathPrimitiveExtension          *PrimitiveExtension `bson:"path_extension,omitempty" json:"_path,omitempty"`
}
type DataRequirementValueFilterValuex struct {
	ValueDateTime *DateTime `binding:"omitempty" bson:",omitempty" json:"valueDateTime,omitempty"`
	ValuePeriod   *Period   `binding:"omitempty" bson:",omitempty" json:"valuePeriod,omitempty"`
	ValueDuration *Duration `binding:"omitempty" bson:",omitempty" json:"valueDuration,omitempty"`
}

func (out *DataRequirementValueFilter) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if len(asMap["valueDateTime"]) > 0 {
		if err := go1.Unmarshal(asMap["valueDateTime"], &out.ValueDateTime); err == nil {
		}
	} else if len(asMap["valuePeriod"]) > 0 {
		if err := go1.Unmarshal(asMap["valuePeriod"], &out.ValuePeriod); err == nil {
		}
	} else if len(asMap["valueDuration"]) > 0 {
		if err := go1.Unmarshal(asMap["valueDuration"], &out.ValueDuration); err == nil {
		}
	} else {

	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["path"]) > 0 {
		if err := go1.Unmarshal(asMap["path"], &out.Path); err != nil {
			return err
		}

	}
	if len(asMap["_path"]) > 0 {
		if err := go1.Unmarshal(asMap["_path"], &out.PathPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["searchParam"]) > 0 {
		if err := go1.Unmarshal(asMap["searchParam"], &out.SearchParam); err != nil {
			return err
		}

	}
	if len(asMap["_searchParam"]) > 0 {
		if err := go1.Unmarshal(asMap["_searchParam"], &out.SearchParamPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["comparator"]) > 0 {
		if err := go1.Unmarshal(asMap["comparator"], &out.Comparator); err != nil {
			return err
		}

	}
	if len(asMap["_comparator"]) > 0 {
		if err := go1.Unmarshal(asMap["_comparator"], &out.ComparatorPrimitiveExtension); err != nil {
			return err
		}
	}

	return nil
}

type DataRequirementValueFilter struct {
	DataRequirementValueFilterValuex `binding:"omitempty" bson:",omitempty" json:",omitempty"`
	Id                               string              `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension             *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension                        []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Path                             string              `bson:",omitempty" json:"path,omitempty"`      // The attribute of the filter. The specified path SHALL be a FHIRPath resolvable on the specified type of the DataRequirement, and SHALL consist only of identifiers, constant indexers, and .resolve(). The path is allowed to contain qualifiers (.) to traverse sub-elements, as well as indexers ([x]) to traverse multiple-cardinality sub-elements (see the [Simple FHIRPath Profile](fhirpath.html#simple) for full details). Note that the index must be an integer constant. The path must resolve to an element of a type that is comparable to the valueFilter.value[x] element for the filter.
	PathPrimitiveExtension           *PrimitiveExtension `bson:"path_extension,omitempty" json:"_path,omitempty"`
	SearchParam                      string              `bson:",omitempty" json:"searchParam,omitempty"` // A search parameter defined on the specified type of the DataRequirement, and which searches on elements of a type compatible with the type of the valueFilter.value[x] for the filter.
	SearchParamPrimitiveExtension    *PrimitiveExtension `bson:"searchParam_extension,omitempty" json:"_searchParam,omitempty"`
	Comparator                       string              `bson:",omitempty" json:"comparator,omitempty"` // The comparator to be used to determine whether the value is matching.
	ComparatorPrimitiveExtension     *PrimitiveExtension `bson:"comparator_extension,omitempty" json:"_comparator,omitempty"`
}

func (out *DataRequirementSort) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["path"], &out.Path); err != nil {
		return err
	}
	if len(asMap["_path"]) > 0 {
		if err := go1.Unmarshal(asMap["_path"], &out.PathPrimitiveExtension); err != nil {
			return err
		}
	}

	if err := go1.Unmarshal(asMap["direction"], &out.Direction); err != nil {
		return err
	}
	if len(asMap["_direction"]) > 0 {
		if err := go1.Unmarshal(asMap["_direction"], &out.DirectionPrimitiveExtension); err != nil {
			return err
		}
	}

	return nil
}

type DataRequirementSort struct {
	Id                          string              `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension        *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension                   []*Extension        `bson:",omitempty" json:"extension,omitempty"`               // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Path                        string              `binding:"required" bson:",omitempty" json:"path,omitempty"` // The attribute of the sort. The specified path must be resolvable from the type of the required data. The path is allowed to contain qualifiers (.) to traverse sub-elements, as well as indexers ([x]) to traverse multiple-cardinality sub-elements. Note that the index must be an integer constant.
	PathPrimitiveExtension      *PrimitiveExtension `bson:"path_extension,omitempty" json:"_path,omitempty"`
	Direction                   string              `binding:"required" bson:",omitempty" json:"direction,omitempty"` // The direction of the sort, ascending or descending.
	DirectionPrimitiveExtension *PrimitiveExtension `bson:"direction_extension,omitempty" json:"_direction,omitempty"`
}
type DataRequirementSubjectx struct {
	SubjectCodeableConcept *CodeableConcept `binding:"omitempty" bson:",omitempty" json:"subjectCodeableConcept,omitempty"`
	SubjectReference       *Reference       `binding:"omitempty" bson:",omitempty" json:"subjectReference,omitempty"`
}

func (out *DataRequirementCodeFilter) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["path"]) > 0 {
		if err := go1.Unmarshal(asMap["path"], &out.Path); err != nil {
			return err
		}

	}
	if len(asMap["_path"]) > 0 {
		if err := go1.Unmarshal(asMap["_path"], &out.PathPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["searchParam"]) > 0 {
		if err := go1.Unmarshal(asMap["searchParam"], &out.SearchParam); err != nil {
			return err
		}

	}
	if len(asMap["_searchParam"]) > 0 {
		if err := go1.Unmarshal(asMap["_searchParam"], &out.SearchParamPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["valueSet"]) > 0 {
		if err := go1.Unmarshal(asMap["valueSet"], &out.ValueSet); err != nil {
			return err
		}

	}
	if len(asMap["_valueSet"]) > 0 {
		if err := go1.Unmarshal(asMap["_valueSet"], &out.ValueSetPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["code"]) > 0 {
		if err := go1.Unmarshal(asMap["code"], &out.Code); err != nil {
			return err
		}

	}
	return nil
}

type DataRequirementCodeFilter struct {
	Id                            string              `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension          *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension                     []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Path                          string              `bson:",omitempty" json:"path,omitempty"`      // The code-valued attribute of the filter. The specified path SHALL be a FHIRPath resolvable on the specified type of the DataRequirement, and SHALL consist only of identifiers, constant indexers, and .resolve(). The path is allowed to contain qualifiers (.) to traverse sub-elements, as well as indexers ([x]) to traverse multiple-cardinality sub-elements (see the [Simple FHIRPath Profile](fhirpath.html#simple) for full details). Note that the index must be an integer constant. The path must resolve to an element of type code, Coding, or CodeableConcept.
	PathPrimitiveExtension        *PrimitiveExtension `bson:"path_extension,omitempty" json:"_path,omitempty"`
	SearchParam                   string              `bson:",omitempty" json:"searchParam,omitempty"` // A token parameter that refers to a search parameter defined on the specified type of the DataRequirement, and which searches on elements of type code, Coding, or CodeableConcept.
	SearchParamPrimitiveExtension *PrimitiveExtension `bson:"searchParam_extension,omitempty" json:"_searchParam,omitempty"`
	ValueSet                      string              `bson:",omitempty" json:"valueSet,omitempty"` // The valueset for the code filter. The valueSet and code elements are additive. If valueSet is specified, the filter will return only those data items for which the value of the code-valued element specified in the path is a member of the specified valueset.
	ValueSetPrimitiveExtension    *PrimitiveExtension `bson:"valueSet_extension,omitempty" json:"_valueSet,omitempty"`
	Code                          []*Coding           `bson:",omitempty" json:"code,omitempty"` // The codes for the code filter. If values are given, the filter will return only those data items for which the code-valued attribute specified by the path has a value that is one of the specified codes. If codes are specified in addition to a value set, the filter returns items matching a code in the value set or one of the specified codes.
}

func (out *DataRequirement) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"DataRequirement\"" {
		return fmt.Errorf("resourceType is not %s", "DataRequirement")
	}
	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["profile"]) > 0 {
		if err := go1.Unmarshal(asMap["profile"], &out.Profile); err != nil {
			return err
		}

	}
	if len(asMap["_profile"]) > 0 {
		if err := go1.Unmarshal(asMap["_profile"], &out.ProfilePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["mustSupport"]) > 0 {
		if err := go1.Unmarshal(asMap["mustSupport"], &out.MustSupport); err != nil {
			return err
		}

	}
	if len(asMap["_mustSupport"]) > 0 {
		if err := go1.Unmarshal(asMap["_mustSupport"], &out.MustSupportPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["dateFilter"]) > 0 {
		if err := go1.Unmarshal(asMap["dateFilter"], &out.DateFilter); err != nil {
			return err
		}

	}
	if len(asMap["valueFilter"]) > 0 {
		if err := go1.Unmarshal(asMap["valueFilter"], &out.ValueFilter); err != nil {
			return err
		}

	}
	if len(asMap["sort"]) > 0 {
		if err := go1.Unmarshal(asMap["sort"], &out.Sort); err != nil {
			return err
		}

	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if err := go1.Unmarshal(asMap["type"], &out.Type); err != nil {
		return err
	}
	if len(asMap["_type"]) > 0 {
		if err := go1.Unmarshal(asMap["_type"], &out.TypePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["subjectCodeableConcept"]) > 0 {
		if err := go1.Unmarshal(asMap["subjectCodeableConcept"], &out.SubjectCodeableConcept); err == nil {
		}
	} else if len(asMap["subjectReference"]) > 0 {
		if err := go1.Unmarshal(asMap["subjectReference"], &out.SubjectReference); err == nil {
		}
	} else {

	}
	if len(asMap["codeFilter"]) > 0 {
		if err := go1.Unmarshal(asMap["codeFilter"], &out.CodeFilter); err != nil {
			return err
		}

	}
	if len(asMap["limit"]) > 0 {
		if err := go1.Unmarshal(asMap["limit"], &out.Limit); err != nil {
			return err
		}

	}
	if len(asMap["_limit"]) > 0 {
		if err := go1.Unmarshal(asMap["_limit"], &out.LimitPrimitiveExtension); err != nil {
			return err
		}
	}

	return nil
}

type DataRequirement struct {
	Extension                 []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Profile                   []string            `bson:",omitempty" json:"profile,omitempty"`   // The profile of the required data, specified as the uri of the profile definition.
	ProfilePrimitiveExtension *PrimitiveExtension `bson:"profile_extension,omitempty" json:"_profile,omitempty"`
	MustSupport               []string            `bson:",omitempty" json:"mustSupport,omitempty"` /*
	Indicates that specific elements of the type are referenced by the knowledge module and must be supported by the consumer in order to obtain an effective evaluation. This does not mean that a value is required for this element, only that the consuming system must understand the element and be able to provide values for it if they are available.

	The value of mustSupport SHALL be a FHIRPath resolvable on the type of the DataRequirement. The path SHALL consist only of identifiers, constant indexers, and .resolve() (see the [Simple FHIRPath Profile](fhirpath.html#simple) for full details).
	*/
	MustSupportPrimitiveExtension *PrimitiveExtension         `bson:"mustSupport_extension,omitempty" json:"_mustSupport,omitempty"`
	DateFilter                    *DataRequirementDateFilter  `bson:",omitempty" json:"dateFilter,omitempty"`
	ValueFilter                   *DataRequirementValueFilter `bson:",omitempty" json:"valueFilter,omitempty"`
	Sort                          *DataRequirementSort        `bson:",omitempty" json:"sort,omitempty"`
	Id                            *primitive.ObjectID         `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension          *PrimitiveExtension         `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Type                          string                      `binding:"required" bson:",omitempty" json:"type,omitempty"` // The type of the required data, specified as the type name of a resource. For profiles, this value is set to the type of the base resource of the profile.
	TypePrimitiveExtension        *PrimitiveExtension         `bson:"type_extension,omitempty" json:"_type,omitempty"`
	DataRequirementSubjectx       `binding:"omitempty" bson:",omitempty" json:",omitempty"`
	CodeFilter                    *DataRequirementCodeFilter `bson:",omitempty" json:"codeFilter,omitempty"`
	Limit                         int                        `bson:",omitempty" json:"limit,omitempty"` // Specifies a maximum number of results that are required (uses the _count search parameter).
	LimitPrimitiveExtension       *PrimitiveExtension        `bson:"limit_extension,omitempty" json:"_limit,omitempty"`
	ResourceType                  string                     `bson:"-" json:"resourceType,omitempty"`
}

func (out *DataType) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"DataType\"" {
		return fmt.Errorf("resourceType is not %s", "DataType")
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	return nil
}

type DataType struct {
	Id                   *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension            []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	ResourceType         string              `bson:"-" json:"resourceType,omitempty"`
}

func (out *Distance) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Distance\"" {
		return fmt.Errorf("resourceType is not %s", "Distance")
	}
	if len(asMap["system"]) > 0 {
		if err := go1.Unmarshal(asMap["system"], &out.System); err != nil {
			return err
		}

	}
	if len(asMap["_system"]) > 0 {
		if err := go1.Unmarshal(asMap["_system"], &out.SystemPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["code"]) > 0 {
		if err := go1.Unmarshal(asMap["code"], &out.Code); err != nil {
			return err
		}

	}
	if len(asMap["_code"]) > 0 {
		if err := go1.Unmarshal(asMap["_code"], &out.CodePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["value"]) > 0 {
		if err := go1.Unmarshal(asMap["value"], &out.Value); err != nil {
			return err
		}

	}
	if len(asMap["_value"]) > 0 {
		if err := go1.Unmarshal(asMap["_value"], &out.ValuePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["comparator"]) > 0 {
		if err := go1.Unmarshal(asMap["comparator"], &out.Comparator); err != nil {
			return err
		}

	}
	if len(asMap["_comparator"]) > 0 {
		if err := go1.Unmarshal(asMap["_comparator"], &out.ComparatorPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["unit"]) > 0 {
		if err := go1.Unmarshal(asMap["unit"], &out.Unit); err != nil {
			return err
		}

	}
	if len(asMap["_unit"]) > 0 {
		if err := go1.Unmarshal(asMap["_unit"], &out.UnitPrimitiveExtension); err != nil {
			return err
		}
	}

	return nil
}

type Distance struct {
	System                       string              `bson:",omitempty" json:"system,omitempty"` // The identification of the system that provides the coded form of the unit.
	SystemPrimitiveExtension     *PrimitiveExtension `bson:"system_extension,omitempty" json:"_system,omitempty"`
	Code                         string              `bson:",omitempty" json:"code,omitempty"` // A computer processable form of the unit in some unit representation system.
	CodePrimitiveExtension       *PrimitiveExtension `bson:"code_extension,omitempty" json:"_code,omitempty"`
	Id                           *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension         *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension                    []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Value                        float64             `bson:",omitempty" json:"value,omitempty"`     // The value of the measured amount. The value includes an implicit precision in the presentation of the value.
	ValuePrimitiveExtension      *PrimitiveExtension `bson:"value_extension,omitempty" json:"_value,omitempty"`
	Comparator                   string              `bson:",omitempty" json:"comparator,omitempty"` // How the value should be understood and represented - whether the actual value is greater or less than the stated value due to measurement issues; e.g. if the comparator is "<" , then the real value is < stated value.
	ComparatorPrimitiveExtension *PrimitiveExtension `bson:"comparator_extension,omitempty" json:"_comparator,omitempty"`
	Unit                         string              `bson:",omitempty" json:"unit,omitempty"` // A human-readable form of the unit.
	UnitPrimitiveExtension       *PrimitiveExtension `bson:"unit_extension,omitempty" json:"_unit,omitempty"`
	ResourceType                 string              `bson:"-" json:"resourceType,omitempty"`
}

type DosageDoseAndRateDosex struct {
	DoseRange    *Range    `binding:"omitempty" bson:",omitempty" json:"doseRange,omitempty"`
	DoseQuantity *Quantity `binding:"omitempty" bson:",omitempty" json:"doseQuantity,omitempty"`
}
type DosageDoseAndRateRatex struct {
	RateRatio    *Ratio    `binding:"omitempty" bson:",omitempty" json:"rateRatio,omitempty"`
	RateRange    *Range    `binding:"omitempty" bson:",omitempty" json:"rateRange,omitempty"`
	RateQuantity *Quantity `binding:"omitempty" bson:",omitempty" json:"rateQuantity,omitempty"`
}

func (out *DosageDoseAndRate) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["type"]) > 0 {
		if err := go1.Unmarshal(asMap["type"], &out.Type); err != nil {
			return err
		}

	}
	if len(asMap["doseRange"]) > 0 {
		if err := go1.Unmarshal(asMap["doseRange"], &out.DoseRange); err == nil {
		}
	} else if len(asMap["doseQuantity"]) > 0 {
		if err := go1.Unmarshal(asMap["doseQuantity"], &out.DoseQuantity); err == nil {
		}
	} else {

	}
	if len(asMap["rateRatio"]) > 0 {
		if err := go1.Unmarshal(asMap["rateRatio"], &out.RateRatio); err == nil {
		}
	} else if len(asMap["rateRange"]) > 0 {
		if err := go1.Unmarshal(asMap["rateRange"], &out.RateRange); err == nil {
		}
	} else if len(asMap["rateQuantity"]) > 0 {
		if err := go1.Unmarshal(asMap["rateQuantity"], &out.RateQuantity); err == nil {
		}
	} else {

	}
	return nil
}

type DosageDoseAndRate struct {
	Id                     string              `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension   *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension              []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Type                   *CodeableConcept    `bson:",omitempty" json:"type,omitempty"`      // The kind of dose or rate specified, for example, ordered or calculated.
	DosageDoseAndRateDosex `binding:"omitempty" bson:",omitempty" json:",omitempty"`
	DosageDoseAndRateRatex `binding:"omitempty" bson:",omitempty" json:",omitempty"`
}

func (out *Dosage) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Dosage\"" {
		return fmt.Errorf("resourceType is not %s", "Dosage")
	}
	if len(asMap["patientInstruction"]) > 0 {
		if err := go1.Unmarshal(asMap["patientInstruction"], &out.PatientInstruction); err != nil {
			return err
		}

	}
	if len(asMap["_patientInstruction"]) > 0 {
		if err := go1.Unmarshal(asMap["_patientInstruction"], &out.PatientInstructionPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["asNeeded"]) > 0 {
		if err := go1.Unmarshal(asMap["asNeeded"], &out.AsNeeded); err != nil {
			return err
		}

	}
	if len(asMap["_asNeeded"]) > 0 {
		if err := go1.Unmarshal(asMap["_asNeeded"], &out.AsNeededPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["maxDosePerLifetime"]) > 0 {
		if err := go1.Unmarshal(asMap["maxDosePerLifetime"], &out.MaxDosePerLifetime); err != nil {
			return err
		}

	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["asNeededFor"]) > 0 {
		if err := go1.Unmarshal(asMap["asNeededFor"], &out.AsNeededFor); err != nil {
			return err
		}

	}
	if len(asMap["site"]) > 0 {
		if err := go1.Unmarshal(asMap["site"], &out.Site); err != nil {
			return err
		}

	}
	if len(asMap["method"]) > 0 {
		if err := go1.Unmarshal(asMap["method"], &out.Method); err != nil {
			return err
		}

	}
	if len(asMap["maxDosePerPeriod"]) > 0 {
		if err := go1.Unmarshal(asMap["maxDosePerPeriod"], &out.MaxDosePerPeriod); err != nil {
			return err
		}

	}
	if len(asMap["maxDosePerAdministration"]) > 0 {
		if err := go1.Unmarshal(asMap["maxDosePerAdministration"], &out.MaxDosePerAdministration); err != nil {
			return err
		}

	}
	if len(asMap["sequence"]) > 0 {
		if err := go1.Unmarshal(asMap["sequence"], &out.Sequence); err != nil {
			return err
		}

	}
	if len(asMap["_sequence"]) > 0 {
		if err := go1.Unmarshal(asMap["_sequence"], &out.SequencePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["text"]) > 0 {
		if err := go1.Unmarshal(asMap["text"], &out.Text); err != nil {
			return err
		}

	}
	if len(asMap["_text"]) > 0 {
		if err := go1.Unmarshal(asMap["_text"], &out.TextPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["additionalInstruction"]) > 0 {
		if err := go1.Unmarshal(asMap["additionalInstruction"], &out.AdditionalInstruction); err != nil {
			return err
		}

	}
	if len(asMap["route"]) > 0 {
		if err := go1.Unmarshal(asMap["route"], &out.Route); err != nil {
			return err
		}

	}
	if len(asMap["doseAndRate"]) > 0 {
		if err := go1.Unmarshal(asMap["doseAndRate"], &out.DoseAndRate); err != nil {
			return err
		}

	}
	if len(asMap["modifierExtension"]) > 0 {
		if err := go1.Unmarshal(asMap["modifierExtension"], &out.ModifierExtension); err != nil {
			return err
		}

	}
	if len(asMap["timing"]) > 0 {
		if err := go1.Unmarshal(asMap["timing"], &out.Timing); err != nil {
			return err
		}

	}
	return nil
}

type Dosage struct {
	PatientInstruction                   string              `bson:",omitempty" json:"patientInstruction,omitempty"` // Instructions in terms that are understood by the patient or consumer.
	PatientInstructionPrimitiveExtension *PrimitiveExtension `bson:"patientInstruction_extension,omitempty" json:"_patientInstruction,omitempty"`
	AsNeeded                             bool                `bson:",omitempty" json:"asNeeded,omitempty"` // Indicates whether the Medication is only taken when needed within a specific dosing schedule (Boolean option).
	AsNeededPrimitiveExtension           *PrimitiveExtension `bson:"asNeeded_extension,omitempty" json:"_asNeeded,omitempty"`
	MaxDosePerLifetime                   *Quantity           `bson:",omitempty" json:"maxDosePerLifetime,omitempty"` // Upper limit on medication per lifetime of the patient.
	Id                                   *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`              // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension                 *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension                            []*Extension        `bson:",omitempty" json:"extension,omitempty"`                // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	AsNeededFor                          []*CodeableConcept  `bson:",omitempty" json:"asNeededFor,omitempty"`              // Indicates whether the Medication is only taken based on a precondition for taking the Medication (CodeableConcept).
	Site                                 *CodeableConcept    `bson:",omitempty" json:"site,omitempty"`                     // Body site to administer to.
	Method                               *CodeableConcept    `bson:",omitempty" json:"method,omitempty"`                   // Technique for administering medication.
	MaxDosePerPeriod                     []*Ratio            `bson:",omitempty" json:"maxDosePerPeriod,omitempty"`         // Upper limit on medication per unit of time.
	MaxDosePerAdministration             *Quantity           `bson:",omitempty" json:"maxDosePerAdministration,omitempty"` // Upper limit on medication per administration.
	Sequence                             int                 `bson:",omitempty" json:"sequence,omitempty"`                 // Indicates the order in which the dosage instructions should be applied or interpreted.
	SequencePrimitiveExtension           *PrimitiveExtension `bson:"sequence_extension,omitempty" json:"_sequence,omitempty"`
	Text                                 string              `bson:",omitempty" json:"text,omitempty"` // Free text dosage instructions e.g. SIG.
	TextPrimitiveExtension               *PrimitiveExtension `bson:"text_extension,omitempty" json:"_text,omitempty"`
	AdditionalInstruction                []*CodeableConcept  `bson:",omitempty" json:"additionalInstruction,omitempty"` // Supplemental instructions to the patient on how to take the medication  (e.g. "with meals" or"take half to one hour before food") or warnings for the patient about the medication (e.g. "may cause drowsiness" or "avoid exposure of skin to direct sunlight or sunlamps").
	Route                                *CodeableConcept    `bson:",omitempty" json:"route,omitempty"`                 // How drug should enter body.
	DoseAndRate                          *DosageDoseAndRate  `bson:",omitempty" json:"doseAndRate,omitempty"`
	ModifierExtension                    []*Extension        `bson:",omitempty" json:"modifierExtension,omitempty"` /*
	May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.

	Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	*/
	Timing       *Timing `bson:",omitempty" json:"timing,omitempty"` // When medication should be administered.
	ResourceType string  `bson:"-" json:"resourceType,omitempty"`
}

func (out *Duration) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Duration\"" {
		return fmt.Errorf("resourceType is not %s", "Duration")
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["value"]) > 0 {
		if err := go1.Unmarshal(asMap["value"], &out.Value); err != nil {
			return err
		}

	}
	if len(asMap["_value"]) > 0 {
		if err := go1.Unmarshal(asMap["_value"], &out.ValuePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["comparator"]) > 0 {
		if err := go1.Unmarshal(asMap["comparator"], &out.Comparator); err != nil {
			return err
		}

	}
	if len(asMap["_comparator"]) > 0 {
		if err := go1.Unmarshal(asMap["_comparator"], &out.ComparatorPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["unit"]) > 0 {
		if err := go1.Unmarshal(asMap["unit"], &out.Unit); err != nil {
			return err
		}

	}
	if len(asMap["_unit"]) > 0 {
		if err := go1.Unmarshal(asMap["_unit"], &out.UnitPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["system"]) > 0 {
		if err := go1.Unmarshal(asMap["system"], &out.System); err != nil {
			return err
		}

	}
	if len(asMap["_system"]) > 0 {
		if err := go1.Unmarshal(asMap["_system"], &out.SystemPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["code"]) > 0 {
		if err := go1.Unmarshal(asMap["code"], &out.Code); err != nil {
			return err
		}

	}
	if len(asMap["_code"]) > 0 {
		if err := go1.Unmarshal(asMap["_code"], &out.CodePrimitiveExtension); err != nil {
			return err
		}
	}

	return nil
}

type Duration struct {
	Id                           *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension         *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension                    []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Value                        float64             `bson:",omitempty" json:"value,omitempty"`     // The value of the measured amount. The value includes an implicit precision in the presentation of the value.
	ValuePrimitiveExtension      *PrimitiveExtension `bson:"value_extension,omitempty" json:"_value,omitempty"`
	Comparator                   string              `bson:",omitempty" json:"comparator,omitempty"` // How the value should be understood and represented - whether the actual value is greater or less than the stated value due to measurement issues; e.g. if the comparator is "<" , then the real value is < stated value.
	ComparatorPrimitiveExtension *PrimitiveExtension `bson:"comparator_extension,omitempty" json:"_comparator,omitempty"`
	Unit                         string              `bson:",omitempty" json:"unit,omitempty"` // A human-readable form of the unit.
	UnitPrimitiveExtension       *PrimitiveExtension `bson:"unit_extension,omitempty" json:"_unit,omitempty"`
	System                       string              `bson:",omitempty" json:"system,omitempty"` // The identification of the system that provides the coded form of the unit.
	SystemPrimitiveExtension     *PrimitiveExtension `bson:"system_extension,omitempty" json:"_system,omitempty"`
	Code                         string              `bson:",omitempty" json:"code,omitempty"` // A computer processable form of the unit in some unit representation system.
	CodePrimitiveExtension       *PrimitiveExtension `bson:"code_extension,omitempty" json:"_code,omitempty"`
	ResourceType                 string              `bson:"-" json:"resourceType,omitempty"`
}

func (out *ElementDefinitionConstraint) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if err := go1.Unmarshal(asMap["human"], &out.Human); err != nil {
		return err
	}
	if len(asMap["_human"]) > 0 {
		if err := go1.Unmarshal(asMap["_human"], &out.HumanPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["requirements"]) > 0 {
		if err := go1.Unmarshal(asMap["requirements"], &out.Requirements); err != nil {
			return err
		}

	}
	if len(asMap["_requirements"]) > 0 {
		if err := go1.Unmarshal(asMap["_requirements"], &out.RequirementsPrimitiveExtension); err != nil {
			return err
		}
	}

	if err := go1.Unmarshal(asMap["severity"], &out.Severity); err != nil {
		return err
	}
	if len(asMap["_severity"]) > 0 {
		if err := go1.Unmarshal(asMap["_severity"], &out.SeverityPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["suppress"]) > 0 {
		if err := go1.Unmarshal(asMap["suppress"], &out.Suppress); err != nil {
			return err
		}

	}
	if len(asMap["_suppress"]) > 0 {
		if err := go1.Unmarshal(asMap["_suppress"], &out.SuppressPrimitiveExtension); err != nil {
			return err
		}
	}

	if err := go1.Unmarshal(asMap["key"], &out.Key); err != nil {
		return err
	}
	if len(asMap["_key"]) > 0 {
		if err := go1.Unmarshal(asMap["_key"], &out.KeyPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["expression"]) > 0 {
		if err := go1.Unmarshal(asMap["expression"], &out.Expression); err != nil {
			return err
		}

	}
	if len(asMap["_expression"]) > 0 {
		if err := go1.Unmarshal(asMap["_expression"], &out.ExpressionPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["source"]) > 0 {
		if err := go1.Unmarshal(asMap["source"], &out.Source); err != nil {
			return err
		}

	}
	if len(asMap["_source"]) > 0 {
		if err := go1.Unmarshal(asMap["_source"], &out.SourcePrimitiveExtension); err != nil {
			return err
		}
	}

	return nil
}

type ElementDefinitionConstraint struct {
	Human                          string              `binding:"required" bson:",omitempty" json:"human,omitempty"` // Text that can be used to describe the constraint in messages identifying that the constraint has been violated.
	HumanPrimitiveExtension        *PrimitiveExtension `bson:"human_extension,omitempty" json:"_human,omitempty"`
	Id                             string              `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension           *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension                      []*Extension        `bson:",omitempty" json:"extension,omitempty"`    // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Requirements                   string              `bson:",omitempty" json:"requirements,omitempty"` // Description of why this constraint is necessary or appropriate.
	RequirementsPrimitiveExtension *PrimitiveExtension `bson:"requirements_extension,omitempty" json:"_requirements,omitempty"`
	Severity                       string              `binding:"required" bson:",omitempty" json:"severity,omitempty"` // Identifies the impact constraint violation has on the conformance of the instance.
	SeverityPrimitiveExtension     *PrimitiveExtension `bson:"severity_extension,omitempty" json:"_severity,omitempty"`
	Suppress                       bool                `bson:",omitempty" json:"suppress,omitempty"` // If true, indicates that the warning or best practice guideline should be suppressed.
	SuppressPrimitiveExtension     *PrimitiveExtension `bson:"suppress_extension,omitempty" json:"_suppress,omitempty"`
	Key                            *primitive.ObjectID `binding:"required" bson:",omitempty" json:"key,omitempty"` // Allows identification of which elements have their cardinalities impacted by the constraint.  Will not be referenced for constraints that do not affect cardinality.
	KeyPrimitiveExtension          *PrimitiveExtension `bson:"key_extension,omitempty" json:"_key,omitempty"`
	Expression                     string              `bson:",omitempty" json:"expression,omitempty"` // A [FHIRPath](fhirpath.html) expression of constraint that can be executed to see if this constraint is met.
	ExpressionPrimitiveExtension   *PrimitiveExtension `bson:"expression_extension,omitempty" json:"_expression,omitempty"`
	Source                         string              `bson:",omitempty" json:"source,omitempty"` // A reference to the original source of the constraint, for traceability purposes.
	SourcePrimitiveExtension       *PrimitiveExtension `bson:"source_extension,omitempty" json:"_source,omitempty"`
}

func (out *ElementDefinitionBase) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["path"], &out.Path); err != nil {
		return err
	}
	if len(asMap["_path"]) > 0 {
		if err := go1.Unmarshal(asMap["_path"], &out.PathPrimitiveExtension); err != nil {
			return err
		}
	}

	if err := go1.Unmarshal(asMap["min"], &out.Min); err != nil {
		return err
	}
	if len(asMap["_min"]) > 0 {
		if err := go1.Unmarshal(asMap["_min"], &out.MinPrimitiveExtension); err != nil {
			return err
		}
	}

	if err := go1.Unmarshal(asMap["max"], &out.Max); err != nil {
		return err
	}
	if len(asMap["_max"]) > 0 {
		if err := go1.Unmarshal(asMap["_max"], &out.MaxPrimitiveExtension); err != nil {
			return err
		}
	}

	return nil
}

type ElementDefinitionBase struct {
	Id                     string              `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension   *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension              []*Extension        `bson:",omitempty" json:"extension,omitempty"`               // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Path                   string              `binding:"required" bson:",omitempty" json:"path,omitempty"` // The Path that identifies the base element - this matches the ElementDefinition.path for that element. Across FHIR, there is only one base definition of any element - that is, an element definition on a [StructureDefinition](structuredefinition.html#) without a StructureDefinition.base.
	PathPrimitiveExtension *PrimitiveExtension `bson:"path_extension,omitempty" json:"_path,omitempty"`
	Min                    int                 `binding:"required" bson:",omitempty" json:"min,omitempty"` // Minimum cardinality of the base element identified by the path.
	MinPrimitiveExtension  *PrimitiveExtension `bson:"min_extension,omitempty" json:"_min,omitempty"`
	Max                    string              `binding:"required" bson:",omitempty" json:"max,omitempty"` // Maximum cardinality of the base element identified by the path.
	MaxPrimitiveExtension  *PrimitiveExtension `bson:"max_extension,omitempty" json:"_max,omitempty"`
}
type ElementDefinitionMinValuex struct {
	MinValueDate        *Date      `binding:"omitempty" bson:",omitempty" json:"minValueDate,omitempty"`
	MinValueDateTime    *DateTime  `binding:"omitempty" bson:",omitempty" json:"minValueDateTime,omitempty"`
	MinValueInstant     *time.Time `binding:"omitempty" bson:",omitempty" json:"minValueInstant,omitempty"`
	MinValueTime        *Time      `binding:"omitempty" bson:",omitempty" json:"minValueTime,omitempty"`
	MinValueDecimal     float64    `binding:"omitempty" bson:",omitempty" json:"minValueDecimal,omitempty"`
	MinValueInteger     int        `binding:"omitempty" bson:",omitempty" json:"minValueInteger,omitempty"`
	MinValueInteger64   int64      `binding:"omitempty" bson:",omitempty" json:"minValueInteger64,omitempty"`
	MinValuePositiveInt int        `binding:"omitempty" bson:",omitempty" json:"minValuePositiveInt,omitempty"`
	MinValueUnsignedInt int        `binding:"omitempty" bson:",omitempty" json:"minValueUnsignedInt,omitempty"`
	MinValueQuantity    *Quantity  `binding:"omitempty" bson:",omitempty" json:"minValueQuantity,omitempty"`
}

func (out *ElementDefinitionBindingAdditional) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if len(asMap["shortDoco"]) > 0 {
		if err := go1.Unmarshal(asMap["shortDoco"], &out.ShortDoco); err != nil {
			return err
		}

	}
	if len(asMap["_shortDoco"]) > 0 {
		if err := go1.Unmarshal(asMap["_shortDoco"], &out.ShortDocoPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["usage"]) > 0 {
		if err := go1.Unmarshal(asMap["usage"], &out.Usage); err != nil {
			return err
		}

	}
	if len(asMap["any"]) > 0 {
		if err := go1.Unmarshal(asMap["any"], &out.Any); err != nil {
			return err
		}

	}
	if len(asMap["_any"]) > 0 {
		if err := go1.Unmarshal(asMap["_any"], &out.AnyPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["purpose"], &out.Purpose); err != nil {
		return err
	}
	if len(asMap["_purpose"]) > 0 {
		if err := go1.Unmarshal(asMap["_purpose"], &out.PurposePrimitiveExtension); err != nil {
			return err
		}
	}

	if err := go1.Unmarshal(asMap["valueSet"], &out.ValueSet); err != nil {
		return err
	}
	if len(asMap["_valueSet"]) > 0 {
		if err := go1.Unmarshal(asMap["_valueSet"], &out.ValueSetPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["documentation"]) > 0 {
		if err := go1.Unmarshal(asMap["documentation"], &out.Documentation); err != nil {
			return err
		}

	}
	if len(asMap["_documentation"]) > 0 {
		if err := go1.Unmarshal(asMap["_documentation"], &out.DocumentationPrimitiveExtension); err != nil {
			return err
		}
	}

	return nil
}

type ElementDefinitionBindingAdditional struct {
	ShortDoco                       string              `bson:",omitempty" json:"shortDoco,omitempty"` // Concise documentation - for summary tables.
	ShortDocoPrimitiveExtension     *PrimitiveExtension `bson:"shortDoco_extension,omitempty" json:"_shortDoco,omitempty"`
	Usage                           []*UsageContext     `bson:",omitempty" json:"usage,omitempty"` // Qualifies the usage of the binding. Typically bindings are qualified by jurisdiction, but they may also be qualified by gender, workflow status, clinical domain etc. The information to decide whether a usege context applies is usually outside the resource, determined by context, and this might present challenges for validation tooling.
	Any                             bool                `bson:",omitempty" json:"any,omitempty"`   // Whether the binding applies to all repeats, or just to any one of them. This is only relevant for elements that can repeat.
	AnyPrimitiveExtension           *PrimitiveExtension `bson:"any_extension,omitempty" json:"_any,omitempty"`
	Id                              string              `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension            *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension                       []*Extension        `bson:",omitempty" json:"extension,omitempty"`                  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Purpose                         string              `binding:"required" bson:",omitempty" json:"purpose,omitempty"` // The use of this additional binding.
	PurposePrimitiveExtension       *PrimitiveExtension `bson:"purpose_extension,omitempty" json:"_purpose,omitempty"`
	ValueSet                        string              `binding:"required" bson:",omitempty" json:"valueSet,omitempty"` // The valueSet that is being bound for the purpose.
	ValueSetPrimitiveExtension      *PrimitiveExtension `bson:"valueSet_extension,omitempty" json:"_valueSet,omitempty"`
	Documentation                   string              `bson:",omitempty" json:"documentation,omitempty"` // Documentation of the purpose of use of the bindingproviding additional information about how it is intended to be used.
	DocumentationPrimitiveExtension *PrimitiveExtension `bson:"documentation_extension,omitempty" json:"_documentation,omitempty"`
}

func (out *ElementDefinitionBinding) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["strength"], &out.Strength); err != nil {
		return err
	}
	if len(asMap["_strength"]) > 0 {
		if err := go1.Unmarshal(asMap["_strength"], &out.StrengthPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["description"]) > 0 {
		if err := go1.Unmarshal(asMap["description"], &out.Description); err != nil {
			return err
		}

	}
	if len(asMap["_description"]) > 0 {
		if err := go1.Unmarshal(asMap["_description"], &out.DescriptionPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["valueSet"]) > 0 {
		if err := go1.Unmarshal(asMap["valueSet"], &out.ValueSet); err != nil {
			return err
		}

	}
	if len(asMap["_valueSet"]) > 0 {
		if err := go1.Unmarshal(asMap["_valueSet"], &out.ValueSetPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["additional"]) > 0 {
		if err := go1.Unmarshal(asMap["additional"], &out.Additional); err != nil {
			return err
		}

	}
	return nil
}

type ElementDefinitionBinding struct {
	Id                            string                              `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension          *PrimitiveExtension                 `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension                     []*Extension                        `bson:",omitempty" json:"extension,omitempty"`                   // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Strength                      string                              `binding:"required" bson:",omitempty" json:"strength,omitempty"` // Indicates the degree of conformance expectations associated with this binding - that is, the degree to which the provided value set must be adhered to in the instances.
	StrengthPrimitiveExtension    *PrimitiveExtension                 `bson:"strength_extension,omitempty" json:"_strength,omitempty"`
	Description                   string                              `bson:",omitempty" json:"description,omitempty"` // Describes the intended use of this particular set of codes.
	DescriptionPrimitiveExtension *PrimitiveExtension                 `bson:"description_extension,omitempty" json:"_description,omitempty"`
	ValueSet                      string                              `bson:",omitempty" json:"valueSet,omitempty"` // Refers to the value set that identifies the set of codes the binding refers to.
	ValueSetPrimitiveExtension    *PrimitiveExtension                 `bson:"valueSet_extension,omitempty" json:"_valueSet,omitempty"`
	Additional                    *ElementDefinitionBindingAdditional `bson:",omitempty" json:"additional,omitempty"`
}
type ElementDefinitionDefaultValuex struct {
	DefaultValueBase64Binary          Base64Binary           `binding:"omitempty" bson:",omitempty" json:"defaultValueBase64Binary,omitempty"`
	DefaultValueBoolean               bool                   `binding:"omitempty" bson:",omitempty" json:"defaultValueBoolean,omitempty"`
	DefaultValueCanonical             string                 `binding:"omitempty" bson:",omitempty" json:"defaultValueCanonical,omitempty"`
	DefaultValueCode                  string                 `binding:"omitempty" bson:",omitempty" json:"defaultValueCode,omitempty"`
	DefaultValueDate                  *Date                  `binding:"omitempty" bson:",omitempty" json:"defaultValueDate,omitempty"`
	DefaultValueDateTime              *DateTime              `binding:"omitempty" bson:",omitempty" json:"defaultValueDateTime,omitempty"`
	DefaultValueDecimal               float64                `binding:"omitempty" bson:",omitempty" json:"defaultValueDecimal,omitempty"`
	DefaultValueId                    *primitive.ObjectID    `binding:"omitempty" bson:",omitempty" json:"defaultValueId,omitempty"`
	DefaultValueInstant               *time.Time             `binding:"omitempty" bson:",omitempty" json:"defaultValueInstant,omitempty"`
	DefaultValueInteger               int                    `binding:"omitempty" bson:",omitempty" json:"defaultValueInteger,omitempty"`
	DefaultValueInteger64             int64                  `binding:"omitempty" bson:",omitempty" json:"defaultValueInteger64,omitempty"`
	DefaultValueMarkdown              string                 `binding:"omitempty" bson:",omitempty" json:"defaultValueMarkdown,omitempty"`
	DefaultValueOid                   string                 `binding:"omitempty" bson:",omitempty" json:"defaultValueOid,omitempty"`
	DefaultValuePositiveInt           int                    `binding:"omitempty" bson:",omitempty" json:"defaultValuePositiveInt,omitempty"`
	DefaultValueString                string                 `binding:"omitempty" bson:",omitempty" json:"defaultValueString,omitempty"`
	DefaultValueTime                  *Time                  `binding:"omitempty" bson:",omitempty" json:"defaultValueTime,omitempty"`
	DefaultValueUnsignedInt           int                    `binding:"omitempty" bson:",omitempty" json:"defaultValueUnsignedInt,omitempty"`
	DefaultValueUri                   string                 `binding:"omitempty" bson:",omitempty" json:"defaultValueUri,omitempty"`
	DefaultValueUrl                   *url.URL               `binding:"omitempty" bson:",omitempty" json:"defaultValueUrl,omitempty"`
	DefaultValueUuid                  *uuid.UUID             `binding:"omitempty" bson:",omitempty" json:"defaultValueUuid,omitempty"`
	DefaultValueAddress               *Address               `binding:"omitempty" bson:",omitempty" json:"defaultValueAddress,omitempty"`
	DefaultValueAge                   *Age                   `binding:"omitempty" bson:",omitempty" json:"defaultValueAge,omitempty"`
	DefaultValueAnnotation            *Annotation            `binding:"omitempty" bson:",omitempty" json:"defaultValueAnnotation,omitempty"`
	DefaultValueAttachment            *Attachment            `binding:"omitempty" bson:",omitempty" json:"defaultValueAttachment,omitempty"`
	DefaultValueCodeableConcept       *CodeableConcept       `binding:"omitempty" bson:",omitempty" json:"defaultValueCodeableConcept,omitempty"`
	DefaultValueCodeableReference     *CodeableReference     `binding:"omitempty" bson:",omitempty" json:"defaultValueCodeableReference,omitempty"`
	DefaultValueCoding                *Coding                `binding:"omitempty" bson:",omitempty" json:"defaultValueCoding,omitempty"`
	DefaultValueContactPoint          *ContactPoint          `binding:"omitempty" bson:",omitempty" json:"defaultValueContactPoint,omitempty"`
	DefaultValueCount                 *Count                 `binding:"omitempty" bson:",omitempty" json:"defaultValueCount,omitempty"`
	DefaultValueDistance              *Distance              `binding:"omitempty" bson:",omitempty" json:"defaultValueDistance,omitempty"`
	DefaultValueDuration              *Duration              `binding:"omitempty" bson:",omitempty" json:"defaultValueDuration,omitempty"`
	DefaultValueHumanName             *HumanName             `binding:"omitempty" bson:",omitempty" json:"defaultValueHumanName,omitempty"`
	DefaultValueIdentifier            *Identifier            `binding:"omitempty" bson:",omitempty" json:"defaultValueIdentifier,omitempty"`
	DefaultValueMoney                 *Money                 `binding:"omitempty" bson:",omitempty" json:"defaultValueMoney,omitempty"`
	DefaultValuePeriod                *Period                `binding:"omitempty" bson:",omitempty" json:"defaultValuePeriod,omitempty"`
	DefaultValueQuantity              *Quantity              `binding:"omitempty" bson:",omitempty" json:"defaultValueQuantity,omitempty"`
	DefaultValueRange                 *Range                 `binding:"omitempty" bson:",omitempty" json:"defaultValueRange,omitempty"`
	DefaultValueRatio                 *Ratio                 `binding:"omitempty" bson:",omitempty" json:"defaultValueRatio,omitempty"`
	DefaultValueRatioRange            *RatioRange            `binding:"omitempty" bson:",omitempty" json:"defaultValueRatioRange,omitempty"`
	DefaultValueReference             *Reference             `binding:"omitempty" bson:",omitempty" json:"defaultValueReference,omitempty"`
	DefaultValueSampledData           *SampledData           `binding:"omitempty" bson:",omitempty" json:"defaultValueSampledData,omitempty"`
	DefaultValueSignature             *Signature             `binding:"omitempty" bson:",omitempty" json:"defaultValueSignature,omitempty"`
	DefaultValueTiming                *Timing                `binding:"omitempty" bson:",omitempty" json:"defaultValueTiming,omitempty"`
	DefaultValueContactDetail         *ContactDetail         `binding:"omitempty" bson:",omitempty" json:"defaultValueContactDetail,omitempty"`
	DefaultValueDataRequirement       *DataRequirement       `binding:"omitempty" bson:",omitempty" json:"defaultValueDataRequirement,omitempty"`
	DefaultValueExpression            *Expression            `binding:"omitempty" bson:",omitempty" json:"defaultValueExpression,omitempty"`
	DefaultValueParameterDefinition   *ParameterDefinition   `binding:"omitempty" bson:",omitempty" json:"defaultValueParameterDefinition,omitempty"`
	DefaultValueRelatedArtifact       *RelatedArtifact       `binding:"omitempty" bson:",omitempty" json:"defaultValueRelatedArtifact,omitempty"`
	DefaultValueTriggerDefinition     *TriggerDefinition     `binding:"omitempty" bson:",omitempty" json:"defaultValueTriggerDefinition,omitempty"`
	DefaultValueUsageContext          *UsageContext          `binding:"omitempty" bson:",omitempty" json:"defaultValueUsageContext,omitempty"`
	DefaultValueAvailability          *Availability          `binding:"omitempty" bson:",omitempty" json:"defaultValueAvailability,omitempty"`
	DefaultValueExtendedContactDetail *ExtendedContactDetail `binding:"omitempty" bson:",omitempty" json:"defaultValueExtendedContactDetail,omitempty"`
	DefaultValueDosage                *Dosage                `binding:"omitempty" bson:",omitempty" json:"defaultValueDosage,omitempty"`
	DefaultValueMeta                  *Meta                  `binding:"omitempty" bson:",omitempty" json:"defaultValueMeta,omitempty"`
}
type ElementDefinitionPatternx struct {
	PatternBase64Binary          Base64Binary           `binding:"omitempty" bson:",omitempty" json:"patternBase64Binary,omitempty"`
	PatternBoolean               bool                   `binding:"omitempty" bson:",omitempty" json:"patternBoolean,omitempty"`
	PatternCanonical             string                 `binding:"omitempty" bson:",omitempty" json:"patternCanonical,omitempty"`
	PatternCode                  string                 `binding:"omitempty" bson:",omitempty" json:"patternCode,omitempty"`
	PatternDate                  *Date                  `binding:"omitempty" bson:",omitempty" json:"patternDate,omitempty"`
	PatternDateTime              *DateTime              `binding:"omitempty" bson:",omitempty" json:"patternDateTime,omitempty"`
	PatternDecimal               float64                `binding:"omitempty" bson:",omitempty" json:"patternDecimal,omitempty"`
	PatternId                    *primitive.ObjectID    `binding:"omitempty" bson:",omitempty" json:"patternId,omitempty"`
	PatternInstant               *time.Time             `binding:"omitempty" bson:",omitempty" json:"patternInstant,omitempty"`
	PatternInteger               int                    `binding:"omitempty" bson:",omitempty" json:"patternInteger,omitempty"`
	PatternInteger64             int64                  `binding:"omitempty" bson:",omitempty" json:"patternInteger64,omitempty"`
	PatternMarkdown              string                 `binding:"omitempty" bson:",omitempty" json:"patternMarkdown,omitempty"`
	PatternOid                   string                 `binding:"omitempty" bson:",omitempty" json:"patternOid,omitempty"`
	PatternPositiveInt           int                    `binding:"omitempty" bson:",omitempty" json:"patternPositiveInt,omitempty"`
	PatternString                string                 `binding:"omitempty" bson:",omitempty" json:"patternString,omitempty"`
	PatternTime                  *Time                  `binding:"omitempty" bson:",omitempty" json:"patternTime,omitempty"`
	PatternUnsignedInt           int                    `binding:"omitempty" bson:",omitempty" json:"patternUnsignedInt,omitempty"`
	PatternUri                   string                 `binding:"omitempty" bson:",omitempty" json:"patternUri,omitempty"`
	PatternUrl                   *url.URL               `binding:"omitempty" bson:",omitempty" json:"patternUrl,omitempty"`
	PatternUuid                  *uuid.UUID             `binding:"omitempty" bson:",omitempty" json:"patternUuid,omitempty"`
	PatternAddress               *Address               `binding:"omitempty" bson:",omitempty" json:"patternAddress,omitempty"`
	PatternAge                   *Age                   `binding:"omitempty" bson:",omitempty" json:"patternAge,omitempty"`
	PatternAnnotation            *Annotation            `binding:"omitempty" bson:",omitempty" json:"patternAnnotation,omitempty"`
	PatternAttachment            *Attachment            `binding:"omitempty" bson:",omitempty" json:"patternAttachment,omitempty"`
	PatternCodeableConcept       *CodeableConcept       `binding:"omitempty" bson:",omitempty" json:"patternCodeableConcept,omitempty"`
	PatternCodeableReference     *CodeableReference     `binding:"omitempty" bson:",omitempty" json:"patternCodeableReference,omitempty"`
	PatternCoding                *Coding                `binding:"omitempty" bson:",omitempty" json:"patternCoding,omitempty"`
	PatternContactPoint          *ContactPoint          `binding:"omitempty" bson:",omitempty" json:"patternContactPoint,omitempty"`
	PatternCount                 *Count                 `binding:"omitempty" bson:",omitempty" json:"patternCount,omitempty"`
	PatternDistance              *Distance              `binding:"omitempty" bson:",omitempty" json:"patternDistance,omitempty"`
	PatternDuration              *Duration              `binding:"omitempty" bson:",omitempty" json:"patternDuration,omitempty"`
	PatternHumanName             *HumanName             `binding:"omitempty" bson:",omitempty" json:"patternHumanName,omitempty"`
	PatternIdentifier            *Identifier            `binding:"omitempty" bson:",omitempty" json:"patternIdentifier,omitempty"`
	PatternMoney                 *Money                 `binding:"omitempty" bson:",omitempty" json:"patternMoney,omitempty"`
	PatternPeriod                *Period                `binding:"omitempty" bson:",omitempty" json:"patternPeriod,omitempty"`
	PatternQuantity              *Quantity              `binding:"omitempty" bson:",omitempty" json:"patternQuantity,omitempty"`
	PatternRange                 *Range                 `binding:"omitempty" bson:",omitempty" json:"patternRange,omitempty"`
	PatternRatio                 *Ratio                 `binding:"omitempty" bson:",omitempty" json:"patternRatio,omitempty"`
	PatternRatioRange            *RatioRange            `binding:"omitempty" bson:",omitempty" json:"patternRatioRange,omitempty"`
	PatternReference             *Reference             `binding:"omitempty" bson:",omitempty" json:"patternReference,omitempty"`
	PatternSampledData           *SampledData           `binding:"omitempty" bson:",omitempty" json:"patternSampledData,omitempty"`
	PatternSignature             *Signature             `binding:"omitempty" bson:",omitempty" json:"patternSignature,omitempty"`
	PatternTiming                *Timing                `binding:"omitempty" bson:",omitempty" json:"patternTiming,omitempty"`
	PatternContactDetail         *ContactDetail         `binding:"omitempty" bson:",omitempty" json:"patternContactDetail,omitempty"`
	PatternDataRequirement       *DataRequirement       `binding:"omitempty" bson:",omitempty" json:"patternDataRequirement,omitempty"`
	PatternExpression            *Expression            `binding:"omitempty" bson:",omitempty" json:"patternExpression,omitempty"`
	PatternParameterDefinition   *ParameterDefinition   `binding:"omitempty" bson:",omitempty" json:"patternParameterDefinition,omitempty"`
	PatternRelatedArtifact       *RelatedArtifact       `binding:"omitempty" bson:",omitempty" json:"patternRelatedArtifact,omitempty"`
	PatternTriggerDefinition     *TriggerDefinition     `binding:"omitempty" bson:",omitempty" json:"patternTriggerDefinition,omitempty"`
	PatternUsageContext          *UsageContext          `binding:"omitempty" bson:",omitempty" json:"patternUsageContext,omitempty"`
	PatternAvailability          *Availability          `binding:"omitempty" bson:",omitempty" json:"patternAvailability,omitempty"`
	PatternExtendedContactDetail *ExtendedContactDetail `binding:"omitempty" bson:",omitempty" json:"patternExtendedContactDetail,omitempty"`
	PatternDosage                *Dosage                `binding:"omitempty" bson:",omitempty" json:"patternDosage,omitempty"`
	PatternMeta                  *Meta                  `binding:"omitempty" bson:",omitempty" json:"patternMeta,omitempty"`
}

func (out *ElementDefinitionMapping) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if len(asMap["comment"]) > 0 {
		if err := go1.Unmarshal(asMap["comment"], &out.Comment); err != nil {
			return err
		}

	}
	if len(asMap["_comment"]) > 0 {
		if err := go1.Unmarshal(asMap["_comment"], &out.CommentPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["identity"], &out.Identity); err != nil {
		return err
	}
	if len(asMap["_identity"]) > 0 {
		if err := go1.Unmarshal(asMap["_identity"], &out.IdentityPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["language"]) > 0 {
		if err := go1.Unmarshal(asMap["language"], &out.Language); err != nil {
			return err
		}

	}
	if len(asMap["_language"]) > 0 {
		if err := go1.Unmarshal(asMap["_language"], &out.LanguagePrimitiveExtension); err != nil {
			return err
		}
	}

	if err := go1.Unmarshal(asMap["map"], &out.Map); err != nil {
		return err
	}
	if len(asMap["_map"]) > 0 {
		if err := go1.Unmarshal(asMap["_map"], &out.MapPrimitiveExtension); err != nil {
			return err
		}
	}

	return nil
}

type ElementDefinitionMapping struct {
	Comment                    string              `bson:",omitempty" json:"comment,omitempty"` // Comments that provide information about the mapping or its use.
	CommentPrimitiveExtension  *PrimitiveExtension `bson:"comment_extension,omitempty" json:"_comment,omitempty"`
	Id                         string              `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension       *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension                  []*Extension        `bson:",omitempty" json:"extension,omitempty"`                   // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Identity                   *primitive.ObjectID `binding:"required" bson:",omitempty" json:"identity,omitempty"` // An internal reference to the definition of a mapping.
	IdentityPrimitiveExtension *PrimitiveExtension `bson:"identity_extension,omitempty" json:"_identity,omitempty"`
	Language                   string              `bson:",omitempty" json:"language,omitempty"` // Identifies the computable language in which mapping.map is expressed.
	LanguagePrimitiveExtension *PrimitiveExtension `bson:"language_extension,omitempty" json:"_language,omitempty"`
	Map                        string              `binding:"required" bson:",omitempty" json:"map,omitempty"` // Expresses what part of the target specification corresponds to this element.
	MapPrimitiveExtension      *PrimitiveExtension `bson:"map_extension,omitempty" json:"_map,omitempty"`
}

func (out *ElementDefinitionSlicingDiscriminator) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["type"], &out.Type); err != nil {
		return err
	}
	if len(asMap["_type"]) > 0 {
		if err := go1.Unmarshal(asMap["_type"], &out.TypePrimitiveExtension); err != nil {
			return err
		}
	}

	if err := go1.Unmarshal(asMap["path"], &out.Path); err != nil {
		return err
	}
	if len(asMap["_path"]) > 0 {
		if err := go1.Unmarshal(asMap["_path"], &out.PathPrimitiveExtension); err != nil {
			return err
		}
	}

	return nil
}

type ElementDefinitionSlicingDiscriminator struct {
	Id                     string              `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension   *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension              []*Extension        `bson:",omitempty" json:"extension,omitempty"`               // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Type                   string              `binding:"required" bson:",omitempty" json:"type,omitempty"` // How the element value is interpreted when discrimination is evaluated.
	TypePrimitiveExtension *PrimitiveExtension `bson:"type_extension,omitempty" json:"_type,omitempty"`
	Path                   string              `binding:"required" bson:",omitempty" json:"path,omitempty"` // A FHIRPath expression, using [the simple subset of FHIRPath](fhirpath.html#simple), that is used to identify the element on which discrimination is based.
	PathPrimitiveExtension *PrimitiveExtension `bson:"path_extension,omitempty" json:"_path,omitempty"`
}

func (out *ElementDefinitionSlicing) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["discriminator"]) > 0 {
		if err := go1.Unmarshal(asMap["discriminator"], &out.Discriminator); err != nil {
			return err
		}

	}
	if len(asMap["description"]) > 0 {
		if err := go1.Unmarshal(asMap["description"], &out.Description); err != nil {
			return err
		}

	}
	if len(asMap["_description"]) > 0 {
		if err := go1.Unmarshal(asMap["_description"], &out.DescriptionPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["ordered"]) > 0 {
		if err := go1.Unmarshal(asMap["ordered"], &out.Ordered); err != nil {
			return err
		}

	}
	if len(asMap["_ordered"]) > 0 {
		if err := go1.Unmarshal(asMap["_ordered"], &out.OrderedPrimitiveExtension); err != nil {
			return err
		}
	}

	if err := go1.Unmarshal(asMap["rules"], &out.Rules); err != nil {
		return err
	}
	if len(asMap["_rules"]) > 0 {
		if err := go1.Unmarshal(asMap["_rules"], &out.RulesPrimitiveExtension); err != nil {
			return err
		}
	}

	return nil
}

type ElementDefinitionSlicing struct {
	Id                            string                                 `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension          *PrimitiveExtension                    `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension                     []*Extension                           `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Discriminator                 *ElementDefinitionSlicingDiscriminator `bson:",omitempty" json:"discriminator,omitempty"`
	Description                   string                                 `bson:",omitempty" json:"description,omitempty"` // A human-readable text description of how the slicing works. If there is no discriminator, this is required to be present to provide whatever information is possible about how the slices can be differentiated.
	DescriptionPrimitiveExtension *PrimitiveExtension                    `bson:"description_extension,omitempty" json:"_description,omitempty"`
	Ordered                       bool                                   `bson:",omitempty" json:"ordered,omitempty"` // If the matching elements have to occur in the same order as defined in the profile.
	OrderedPrimitiveExtension     *PrimitiveExtension                    `bson:"ordered_extension,omitempty" json:"_ordered,omitempty"`
	Rules                         string                                 `binding:"required" bson:",omitempty" json:"rules,omitempty"` // Whether additional slices are allowed or not. When the slices are ordered, profile authors can also say that additional slices are only allowed at the end.
	RulesPrimitiveExtension       *PrimitiveExtension                    `bson:"rules_extension,omitempty" json:"_rules,omitempty"`
}
type ElementDefinitionExampleValuex struct {
	ValueBase64Binary          Base64Binary           `binding:"omitempty" bson:",omitempty" json:"valueBase64Binary,omitempty"`
	ValueBoolean               bool                   `binding:"omitempty" bson:",omitempty" json:"valueBoolean,omitempty"`
	ValueCanonical             string                 `binding:"omitempty" bson:",omitempty" json:"valueCanonical,omitempty"`
	ValueCode                  string                 `binding:"omitempty" bson:",omitempty" json:"valueCode,omitempty"`
	ValueDate                  *Date                  `binding:"omitempty" bson:",omitempty" json:"valueDate,omitempty"`
	ValueDateTime              *DateTime              `binding:"omitempty" bson:",omitempty" json:"valueDateTime,omitempty"`
	ValueDecimal               float64                `binding:"omitempty" bson:",omitempty" json:"valueDecimal,omitempty"`
	ValueId                    *primitive.ObjectID    `binding:"omitempty" bson:",omitempty" json:"valueId,omitempty"`
	ValueInstant               *time.Time             `binding:"omitempty" bson:",omitempty" json:"valueInstant,omitempty"`
	ValueInteger               int                    `binding:"omitempty" bson:",omitempty" json:"valueInteger,omitempty"`
	ValueInteger64             int64                  `binding:"omitempty" bson:",omitempty" json:"valueInteger64,omitempty"`
	ValueMarkdown              string                 `binding:"omitempty" bson:",omitempty" json:"valueMarkdown,omitempty"`
	ValueOid                   string                 `binding:"omitempty" bson:",omitempty" json:"valueOid,omitempty"`
	ValuePositiveInt           int                    `binding:"omitempty" bson:",omitempty" json:"valuePositiveInt,omitempty"`
	ValueString                string                 `binding:"omitempty" bson:",omitempty" json:"valueString,omitempty"`
	ValueTime                  *Time                  `binding:"omitempty" bson:",omitempty" json:"valueTime,omitempty"`
	ValueUnsignedInt           int                    `binding:"omitempty" bson:",omitempty" json:"valueUnsignedInt,omitempty"`
	ValueUri                   string                 `binding:"omitempty" bson:",omitempty" json:"valueUri,omitempty"`
	ValueUrl                   *url.URL               `binding:"omitempty" bson:",omitempty" json:"valueUrl,omitempty"`
	ValueUuid                  *uuid.UUID             `binding:"omitempty" bson:",omitempty" json:"valueUuid,omitempty"`
	ValueAddress               *Address               `binding:"omitempty" bson:",omitempty" json:"valueAddress,omitempty"`
	ValueAge                   *Age                   `binding:"omitempty" bson:",omitempty" json:"valueAge,omitempty"`
	ValueAnnotation            *Annotation            `binding:"omitempty" bson:",omitempty" json:"valueAnnotation,omitempty"`
	ValueAttachment            *Attachment            `binding:"omitempty" bson:",omitempty" json:"valueAttachment,omitempty"`
	ValueCodeableConcept       *CodeableConcept       `binding:"omitempty" bson:",omitempty" json:"valueCodeableConcept,omitempty"`
	ValueCodeableReference     *CodeableReference     `binding:"omitempty" bson:",omitempty" json:"valueCodeableReference,omitempty"`
	ValueCoding                *Coding                `binding:"omitempty" bson:",omitempty" json:"valueCoding,omitempty"`
	ValueContactPoint          *ContactPoint          `binding:"omitempty" bson:",omitempty" json:"valueContactPoint,omitempty"`
	ValueCount                 *Count                 `binding:"omitempty" bson:",omitempty" json:"valueCount,omitempty"`
	ValueDistance              *Distance              `binding:"omitempty" bson:",omitempty" json:"valueDistance,omitempty"`
	ValueDuration              *Duration              `binding:"omitempty" bson:",omitempty" json:"valueDuration,omitempty"`
	ValueHumanName             *HumanName             `binding:"omitempty" bson:",omitempty" json:"valueHumanName,omitempty"`
	ValueIdentifier            *Identifier            `binding:"omitempty" bson:",omitempty" json:"valueIdentifier,omitempty"`
	ValueMoney                 *Money                 `binding:"omitempty" bson:",omitempty" json:"valueMoney,omitempty"`
	ValuePeriod                *Period                `binding:"omitempty" bson:",omitempty" json:"valuePeriod,omitempty"`
	ValueQuantity              *Quantity              `binding:"omitempty" bson:",omitempty" json:"valueQuantity,omitempty"`
	ValueRange                 *Range                 `binding:"omitempty" bson:",omitempty" json:"valueRange,omitempty"`
	ValueRatio                 *Ratio                 `binding:"omitempty" bson:",omitempty" json:"valueRatio,omitempty"`
	ValueRatioRange            *RatioRange            `binding:"omitempty" bson:",omitempty" json:"valueRatioRange,omitempty"`
	ValueReference             *Reference             `binding:"omitempty" bson:",omitempty" json:"valueReference,omitempty"`
	ValueSampledData           *SampledData           `binding:"omitempty" bson:",omitempty" json:"valueSampledData,omitempty"`
	ValueSignature             *Signature             `binding:"omitempty" bson:",omitempty" json:"valueSignature,omitempty"`
	ValueTiming                *Timing                `binding:"omitempty" bson:",omitempty" json:"valueTiming,omitempty"`
	ValueContactDetail         *ContactDetail         `binding:"omitempty" bson:",omitempty" json:"valueContactDetail,omitempty"`
	ValueDataRequirement       *DataRequirement       `binding:"omitempty" bson:",omitempty" json:"valueDataRequirement,omitempty"`
	ValueExpression            *Expression            `binding:"omitempty" bson:",omitempty" json:"valueExpression,omitempty"`
	ValueParameterDefinition   *ParameterDefinition   `binding:"omitempty" bson:",omitempty" json:"valueParameterDefinition,omitempty"`
	ValueRelatedArtifact       *RelatedArtifact       `binding:"omitempty" bson:",omitempty" json:"valueRelatedArtifact,omitempty"`
	ValueTriggerDefinition     *TriggerDefinition     `binding:"omitempty" bson:",omitempty" json:"valueTriggerDefinition,omitempty"`
	ValueUsageContext          *UsageContext          `binding:"omitempty" bson:",omitempty" json:"valueUsageContext,omitempty"`
	ValueAvailability          *Availability          `binding:"omitempty" bson:",omitempty" json:"valueAvailability,omitempty"`
	ValueExtendedContactDetail *ExtendedContactDetail `binding:"omitempty" bson:",omitempty" json:"valueExtendedContactDetail,omitempty"`
	ValueDosage                *Dosage                `binding:"omitempty" bson:",omitempty" json:"valueDosage,omitempty"`
	ValueMeta                  *Meta                  `binding:"omitempty" bson:",omitempty" json:"valueMeta,omitempty"`
}

func (out *ElementDefinitionExample) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["label"], &out.Label); err != nil {
		return err
	}
	if len(asMap["_label"]) > 0 {
		if err := go1.Unmarshal(asMap["_label"], &out.LabelPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["valueBase64Binary"]) > 0 {
		if err := go1.Unmarshal(asMap["valueBase64Binary"], &out.ValueBase64Binary); err == nil {
		}
	} else if len(asMap["valueBoolean"]) > 0 {
		if err := go1.Unmarshal(asMap["valueBoolean"], &out.ValueBoolean); err == nil {
		}
	} else if len(asMap["valueCanonical"]) > 0 {
		if err := go1.Unmarshal(asMap["valueCanonical"], &out.ValueCanonical); err == nil {
		}
	} else if len(asMap["valueCode"]) > 0 {
		if err := go1.Unmarshal(asMap["valueCode"], &out.ValueCode); err == nil {
		}
	} else if len(asMap["valueDate"]) > 0 {
		if err := go1.Unmarshal(asMap["valueDate"], &out.ValueDate); err == nil {
		}
	} else if len(asMap["valueDateTime"]) > 0 {
		if err := go1.Unmarshal(asMap["valueDateTime"], &out.ValueDateTime); err == nil {
		}
	} else if len(asMap["valueDecimal"]) > 0 {
		if err := go1.Unmarshal(asMap["valueDecimal"], &out.ValueDecimal); err == nil {
		}
	} else if len(asMap["valueId"]) > 0 {
		if err := go1.Unmarshal(asMap["valueId"], &out.ValueId); err == nil {
		}
	} else if len(asMap["valueInstant"]) > 0 {
		if err := go1.Unmarshal(asMap["valueInstant"], &out.ValueInstant); err == nil {
		}
	} else if len(asMap["valueInteger"]) > 0 {
		if err := go1.Unmarshal(asMap["valueInteger"], &out.ValueInteger); err == nil {
		}
	} else if len(asMap["valueInteger64"]) > 0 {
		if err := go1.Unmarshal(asMap["valueInteger64"], &out.ValueInteger64); err == nil {
		}
	} else if len(asMap["valueMarkdown"]) > 0 {
		if err := go1.Unmarshal(asMap["valueMarkdown"], &out.ValueMarkdown); err == nil {
		}
	} else if len(asMap["valueOid"]) > 0 {
		if err := go1.Unmarshal(asMap["valueOid"], &out.ValueOid); err == nil {
		}
	} else if len(asMap["valuePositiveInt"]) > 0 {
		if err := go1.Unmarshal(asMap["valuePositiveInt"], &out.ValuePositiveInt); err == nil {
		}
	} else if len(asMap["valueString"]) > 0 {
		if err := go1.Unmarshal(asMap["valueString"], &out.ValueString); err == nil {
		}
	} else if len(asMap["valueTime"]) > 0 {
		if err := go1.Unmarshal(asMap["valueTime"], &out.ValueTime); err == nil {
		}
	} else if len(asMap["valueUnsignedInt"]) > 0 {
		if err := go1.Unmarshal(asMap["valueUnsignedInt"], &out.ValueUnsignedInt); err == nil {
		}
	} else if len(asMap["valueUri"]) > 0 {
		if err := go1.Unmarshal(asMap["valueUri"], &out.ValueUri); err == nil {
		}
	} else if len(asMap["valueUrl"]) > 0 {
		if err := go1.Unmarshal(asMap["valueUrl"], &out.ValueUrl); err == nil {
		}
	} else if len(asMap["valueUuid"]) > 0 {
		if err := go1.Unmarshal(asMap["valueUuid"], &out.ValueUuid); err == nil {
		}
	} else if len(asMap["valueAddress"]) > 0 {
		if err := go1.Unmarshal(asMap["valueAddress"], &out.ValueAddress); err == nil {
		}
	} else if len(asMap["valueAge"]) > 0 {
		if err := go1.Unmarshal(asMap["valueAge"], &out.ValueAge); err == nil {
		}
	} else if len(asMap["valueAnnotation"]) > 0 {
		if err := go1.Unmarshal(asMap["valueAnnotation"], &out.ValueAnnotation); err == nil {
		}
	} else if len(asMap["valueAttachment"]) > 0 {
		if err := go1.Unmarshal(asMap["valueAttachment"], &out.ValueAttachment); err == nil {
		}
	} else if len(asMap["valueCodeableConcept"]) > 0 {
		if err := go1.Unmarshal(asMap["valueCodeableConcept"], &out.ValueCodeableConcept); err == nil {
		}
	} else if len(asMap["valueCodeableReference"]) > 0 {
		if err := go1.Unmarshal(asMap["valueCodeableReference"], &out.ValueCodeableReference); err == nil {
		}
	} else if len(asMap["valueCoding"]) > 0 {
		if err := go1.Unmarshal(asMap["valueCoding"], &out.ValueCoding); err == nil {
		}
	} else if len(asMap["valueContactPoint"]) > 0 {
		if err := go1.Unmarshal(asMap["valueContactPoint"], &out.ValueContactPoint); err == nil {
		}
	} else if len(asMap["valueCount"]) > 0 {
		if err := go1.Unmarshal(asMap["valueCount"], &out.ValueCount); err == nil {
		}
	} else if len(asMap["valueDistance"]) > 0 {
		if err := go1.Unmarshal(asMap["valueDistance"], &out.ValueDistance); err == nil {
		}
	} else if len(asMap["valueDuration"]) > 0 {
		if err := go1.Unmarshal(asMap["valueDuration"], &out.ValueDuration); err == nil {
		}
	} else if len(asMap["valueHumanName"]) > 0 {
		if err := go1.Unmarshal(asMap["valueHumanName"], &out.ValueHumanName); err == nil {
		}
	} else if len(asMap["valueIdentifier"]) > 0 {
		if err := go1.Unmarshal(asMap["valueIdentifier"], &out.ValueIdentifier); err == nil {
		}
	} else if len(asMap["valueMoney"]) > 0 {
		if err := go1.Unmarshal(asMap["valueMoney"], &out.ValueMoney); err == nil {
		}
	} else if len(asMap["valuePeriod"]) > 0 {
		if err := go1.Unmarshal(asMap["valuePeriod"], &out.ValuePeriod); err == nil {
		}
	} else if len(asMap["valueQuantity"]) > 0 {
		if err := go1.Unmarshal(asMap["valueQuantity"], &out.ValueQuantity); err == nil {
		}
	} else if len(asMap["valueRange"]) > 0 {
		if err := go1.Unmarshal(asMap["valueRange"], &out.ValueRange); err == nil {
		}
	} else if len(asMap["valueRatio"]) > 0 {
		if err := go1.Unmarshal(asMap["valueRatio"], &out.ValueRatio); err == nil {
		}
	} else if len(asMap["valueRatioRange"]) > 0 {
		if err := go1.Unmarshal(asMap["valueRatioRange"], &out.ValueRatioRange); err == nil {
		}
	} else if len(asMap["valueReference"]) > 0 {
		if err := go1.Unmarshal(asMap["valueReference"], &out.ValueReference); err == nil {
		}
	} else if len(asMap["valueSampledData"]) > 0 {
		if err := go1.Unmarshal(asMap["valueSampledData"], &out.ValueSampledData); err == nil {
		}
	} else if len(asMap["valueSignature"]) > 0 {
		if err := go1.Unmarshal(asMap["valueSignature"], &out.ValueSignature); err == nil {
		}
	} else if len(asMap["valueTiming"]) > 0 {
		if err := go1.Unmarshal(asMap["valueTiming"], &out.ValueTiming); err == nil {
		}
	} else if len(asMap["valueContactDetail"]) > 0 {
		if err := go1.Unmarshal(asMap["valueContactDetail"], &out.ValueContactDetail); err == nil {
		}
	} else if len(asMap["valueDataRequirement"]) > 0 {
		if err := go1.Unmarshal(asMap["valueDataRequirement"], &out.ValueDataRequirement); err == nil {
		}
	} else if len(asMap["valueExpression"]) > 0 {
		if err := go1.Unmarshal(asMap["valueExpression"], &out.ValueExpression); err == nil {
		}
	} else if len(asMap["valueParameterDefinition"]) > 0 {
		if err := go1.Unmarshal(asMap["valueParameterDefinition"], &out.ValueParameterDefinition); err == nil {
		}
	} else if len(asMap["valueRelatedArtifact"]) > 0 {
		if err := go1.Unmarshal(asMap["valueRelatedArtifact"], &out.ValueRelatedArtifact); err == nil {
		}
	} else if len(asMap["valueTriggerDefinition"]) > 0 {
		if err := go1.Unmarshal(asMap["valueTriggerDefinition"], &out.ValueTriggerDefinition); err == nil {
		}
	} else if len(asMap["valueUsageContext"]) > 0 {
		if err := go1.Unmarshal(asMap["valueUsageContext"], &out.ValueUsageContext); err == nil {
		}
	} else if len(asMap["valueAvailability"]) > 0 {
		if err := go1.Unmarshal(asMap["valueAvailability"], &out.ValueAvailability); err == nil {
		}
	} else if len(asMap["valueExtendedContactDetail"]) > 0 {
		if err := go1.Unmarshal(asMap["valueExtendedContactDetail"], &out.ValueExtendedContactDetail); err == nil {
		}
	} else if len(asMap["valueDosage"]) > 0 {
		if err := go1.Unmarshal(asMap["valueDosage"], &out.ValueDosage); err == nil {
		}
	} else if len(asMap["valueMeta"]) > 0 {
		if err := go1.Unmarshal(asMap["valueMeta"], &out.ValueMeta); err == nil {
		}
	} else {
		return fmt.Errorf("could not unmarshal %s into any of the possible types", "value[x]")
	}
	return nil
}

type ElementDefinitionExample struct {
	Id                             string              `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension           *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension                      []*Extension        `bson:",omitempty" json:"extension,omitempty"`                // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Label                          string              `binding:"required" bson:",omitempty" json:"label,omitempty"` // Describes the purpose of this example among the set of examples.
	LabelPrimitiveExtension        *PrimitiveExtension `bson:"label_extension,omitempty" json:"_label,omitempty"`
	ElementDefinitionExampleValuex `binding:"omitempty" bson:",omitempty" json:",omitempty"`
}
type ElementDefinitionFixedx struct {
	FixedBase64Binary          Base64Binary           `binding:"omitempty" bson:",omitempty" json:"fixedBase64Binary,omitempty"`
	FixedBoolean               bool                   `binding:"omitempty" bson:",omitempty" json:"fixedBoolean,omitempty"`
	FixedCanonical             string                 `binding:"omitempty" bson:",omitempty" json:"fixedCanonical,omitempty"`
	FixedCode                  string                 `binding:"omitempty" bson:",omitempty" json:"fixedCode,omitempty"`
	FixedDate                  *Date                  `binding:"omitempty" bson:",omitempty" json:"fixedDate,omitempty"`
	FixedDateTime              *DateTime              `binding:"omitempty" bson:",omitempty" json:"fixedDateTime,omitempty"`
	FixedDecimal               float64                `binding:"omitempty" bson:",omitempty" json:"fixedDecimal,omitempty"`
	FixedId                    *primitive.ObjectID    `binding:"omitempty" bson:",omitempty" json:"fixedId,omitempty"`
	FixedInstant               *time.Time             `binding:"omitempty" bson:",omitempty" json:"fixedInstant,omitempty"`
	FixedInteger               int                    `binding:"omitempty" bson:",omitempty" json:"fixedInteger,omitempty"`
	FixedInteger64             int64                  `binding:"omitempty" bson:",omitempty" json:"fixedInteger64,omitempty"`
	FixedMarkdown              string                 `binding:"omitempty" bson:",omitempty" json:"fixedMarkdown,omitempty"`
	FixedOid                   string                 `binding:"omitempty" bson:",omitempty" json:"fixedOid,omitempty"`
	FixedPositiveInt           int                    `binding:"omitempty" bson:",omitempty" json:"fixedPositiveInt,omitempty"`
	FixedString                string                 `binding:"omitempty" bson:",omitempty" json:"fixedString,omitempty"`
	FixedTime                  *Time                  `binding:"omitempty" bson:",omitempty" json:"fixedTime,omitempty"`
	FixedUnsignedInt           int                    `binding:"omitempty" bson:",omitempty" json:"fixedUnsignedInt,omitempty"`
	FixedUri                   string                 `binding:"omitempty" bson:",omitempty" json:"fixedUri,omitempty"`
	FixedUrl                   *url.URL               `binding:"omitempty" bson:",omitempty" json:"fixedUrl,omitempty"`
	FixedUuid                  *uuid.UUID             `binding:"omitempty" bson:",omitempty" json:"fixedUuid,omitempty"`
	FixedAddress               *Address               `binding:"omitempty" bson:",omitempty" json:"fixedAddress,omitempty"`
	FixedAge                   *Age                   `binding:"omitempty" bson:",omitempty" json:"fixedAge,omitempty"`
	FixedAnnotation            *Annotation            `binding:"omitempty" bson:",omitempty" json:"fixedAnnotation,omitempty"`
	FixedAttachment            *Attachment            `binding:"omitempty" bson:",omitempty" json:"fixedAttachment,omitempty"`
	FixedCodeableConcept       *CodeableConcept       `binding:"omitempty" bson:",omitempty" json:"fixedCodeableConcept,omitempty"`
	FixedCodeableReference     *CodeableReference     `binding:"omitempty" bson:",omitempty" json:"fixedCodeableReference,omitempty"`
	FixedCoding                *Coding                `binding:"omitempty" bson:",omitempty" json:"fixedCoding,omitempty"`
	FixedContactPoint          *ContactPoint          `binding:"omitempty" bson:",omitempty" json:"fixedContactPoint,omitempty"`
	FixedCount                 *Count                 `binding:"omitempty" bson:",omitempty" json:"fixedCount,omitempty"`
	FixedDistance              *Distance              `binding:"omitempty" bson:",omitempty" json:"fixedDistance,omitempty"`
	FixedDuration              *Duration              `binding:"omitempty" bson:",omitempty" json:"fixedDuration,omitempty"`
	FixedHumanName             *HumanName             `binding:"omitempty" bson:",omitempty" json:"fixedHumanName,omitempty"`
	FixedIdentifier            *Identifier            `binding:"omitempty" bson:",omitempty" json:"fixedIdentifier,omitempty"`
	FixedMoney                 *Money                 `binding:"omitempty" bson:",omitempty" json:"fixedMoney,omitempty"`
	FixedPeriod                *Period                `binding:"omitempty" bson:",omitempty" json:"fixedPeriod,omitempty"`
	FixedQuantity              *Quantity              `binding:"omitempty" bson:",omitempty" json:"fixedQuantity,omitempty"`
	FixedRange                 *Range                 `binding:"omitempty" bson:",omitempty" json:"fixedRange,omitempty"`
	FixedRatio                 *Ratio                 `binding:"omitempty" bson:",omitempty" json:"fixedRatio,omitempty"`
	FixedRatioRange            *RatioRange            `binding:"omitempty" bson:",omitempty" json:"fixedRatioRange,omitempty"`
	FixedReference             *Reference             `binding:"omitempty" bson:",omitempty" json:"fixedReference,omitempty"`
	FixedSampledData           *SampledData           `binding:"omitempty" bson:",omitempty" json:"fixedSampledData,omitempty"`
	FixedSignature             *Signature             `binding:"omitempty" bson:",omitempty" json:"fixedSignature,omitempty"`
	FixedTiming                *Timing                `binding:"omitempty" bson:",omitempty" json:"fixedTiming,omitempty"`
	FixedContactDetail         *ContactDetail         `binding:"omitempty" bson:",omitempty" json:"fixedContactDetail,omitempty"`
	FixedDataRequirement       *DataRequirement       `binding:"omitempty" bson:",omitempty" json:"fixedDataRequirement,omitempty"`
	FixedExpression            *Expression            `binding:"omitempty" bson:",omitempty" json:"fixedExpression,omitempty"`
	FixedParameterDefinition   *ParameterDefinition   `binding:"omitempty" bson:",omitempty" json:"fixedParameterDefinition,omitempty"`
	FixedRelatedArtifact       *RelatedArtifact       `binding:"omitempty" bson:",omitempty" json:"fixedRelatedArtifact,omitempty"`
	FixedTriggerDefinition     *TriggerDefinition     `binding:"omitempty" bson:",omitempty" json:"fixedTriggerDefinition,omitempty"`
	FixedUsageContext          *UsageContext          `binding:"omitempty" bson:",omitempty" json:"fixedUsageContext,omitempty"`
	FixedAvailability          *Availability          `binding:"omitempty" bson:",omitempty" json:"fixedAvailability,omitempty"`
	FixedExtendedContactDetail *ExtendedContactDetail `binding:"omitempty" bson:",omitempty" json:"fixedExtendedContactDetail,omitempty"`
	FixedDosage                *Dosage                `binding:"omitempty" bson:",omitempty" json:"fixedDosage,omitempty"`
	FixedMeta                  *Meta                  `binding:"omitempty" bson:",omitempty" json:"fixedMeta,omitempty"`
}

func (out *ElementDefinitionType) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["code"], &out.Code); err != nil {
		return err
	}
	if len(asMap["_code"]) > 0 {
		if err := go1.Unmarshal(asMap["_code"], &out.CodePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["profile"]) > 0 {
		if err := go1.Unmarshal(asMap["profile"], &out.Profile); err != nil {
			return err
		}

	}
	if len(asMap["_profile"]) > 0 {
		if err := go1.Unmarshal(asMap["_profile"], &out.ProfilePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["targetProfile"]) > 0 {
		if err := go1.Unmarshal(asMap["targetProfile"], &out.TargetProfile); err != nil {
			return err
		}

	}
	if len(asMap["_targetProfile"]) > 0 {
		if err := go1.Unmarshal(asMap["_targetProfile"], &out.TargetProfilePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["aggregation"]) > 0 {
		if err := go1.Unmarshal(asMap["aggregation"], &out.Aggregation); err != nil {
			return err
		}

	}
	if len(asMap["_aggregation"]) > 0 {
		if err := go1.Unmarshal(asMap["_aggregation"], &out.AggregationPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["versioning"]) > 0 {
		if err := go1.Unmarshal(asMap["versioning"], &out.Versioning); err != nil {
			return err
		}

	}
	if len(asMap["_versioning"]) > 0 {
		if err := go1.Unmarshal(asMap["_versioning"], &out.VersioningPrimitiveExtension); err != nil {
			return err
		}
	}

	return nil
}

type ElementDefinitionType struct {
	Id                              string              `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension            *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension                       []*Extension        `bson:",omitempty" json:"extension,omitempty"`               // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Code                            string              `binding:"required" bson:",omitempty" json:"code,omitempty"` // URL of Data type or Resource that is a(or the) type used for this element. References are URLs that are relative to http://hl7.org/fhir/StructureDefinition e.g. "string" is a reference to http://hl7.org/fhir/StructureDefinition/string. Absolute URLs are only allowed in logical models.
	CodePrimitiveExtension          *PrimitiveExtension `bson:"code_extension,omitempty" json:"_code,omitempty"`
	Profile                         []string            `bson:",omitempty" json:"profile,omitempty"` // Identifies a profile structure or implementation Guide that applies to the datatype this element refers to. If any profiles are specified, then the content must conform to at least one of them. The URL can be a local reference - to a contained StructureDefinition, or a reference to another StructureDefinition or Implementation Guide by a canonical URL. When an implementation guide is specified, the type SHALL conform to at least one profile defined in the implementation guide.
	ProfilePrimitiveExtension       *PrimitiveExtension `bson:"profile_extension,omitempty" json:"_profile,omitempty"`
	TargetProfile                   []string            `bson:",omitempty" json:"targetProfile,omitempty"` // Used when the type is "Reference" or "canonical", and identifies a profile structure or implementation Guide that applies to the target of the reference this element refers to. If any profiles are specified, then the content must conform to at least one of them. The URL can be a local reference - to a contained StructureDefinition, or a reference to another StructureDefinition or Implementation Guide by a canonical URL. When an implementation guide is specified, the target resource SHALL conform to at least one profile defined in the implementation guide.
	TargetProfilePrimitiveExtension *PrimitiveExtension `bson:"targetProfile_extension,omitempty" json:"_targetProfile,omitempty"`
	Aggregation                     []string            `bson:",omitempty" json:"aggregation,omitempty"` // If the type is a reference to another resource, how the resource is or can be aggregated - is it a contained resource, or a reference, and if the context is a bundle, is it included in the bundle.
	AggregationPrimitiveExtension   *PrimitiveExtension `bson:"aggregation_extension,omitempty" json:"_aggregation,omitempty"`
	Versioning                      string              `bson:",omitempty" json:"versioning,omitempty"` // Whether this reference needs to be version specific or version independent, or whether either can be used.
	VersioningPrimitiveExtension    *PrimitiveExtension `bson:"versioning_extension,omitempty" json:"_versioning,omitempty"`
}
type ElementDefinitionMaxValuex struct {
	MaxValueDate        *Date      `binding:"omitempty" bson:",omitempty" json:"maxValueDate,omitempty"`
	MaxValueDateTime    *DateTime  `binding:"omitempty" bson:",omitempty" json:"maxValueDateTime,omitempty"`
	MaxValueInstant     *time.Time `binding:"omitempty" bson:",omitempty" json:"maxValueInstant,omitempty"`
	MaxValueTime        *Time      `binding:"omitempty" bson:",omitempty" json:"maxValueTime,omitempty"`
	MaxValueDecimal     float64    `binding:"omitempty" bson:",omitempty" json:"maxValueDecimal,omitempty"`
	MaxValueInteger     int        `binding:"omitempty" bson:",omitempty" json:"maxValueInteger,omitempty"`
	MaxValueInteger64   int64      `binding:"omitempty" bson:",omitempty" json:"maxValueInteger64,omitempty"`
	MaxValuePositiveInt int        `binding:"omitempty" bson:",omitempty" json:"maxValuePositiveInt,omitempty"`
	MaxValueUnsignedInt int        `binding:"omitempty" bson:",omitempty" json:"maxValueUnsignedInt,omitempty"`
	MaxValueQuantity    *Quantity  `binding:"omitempty" bson:",omitempty" json:"maxValueQuantity,omitempty"`
}

func (out *ElementDefinition) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"ElementDefinition\"" {
		return fmt.Errorf("resourceType is not %s", "ElementDefinition")
	}
	if len(asMap["constraint"]) > 0 {
		if err := go1.Unmarshal(asMap["constraint"], &out.Constraint); err != nil {
			return err
		}

	}
	if len(asMap["modifierExtension"]) > 0 {
		if err := go1.Unmarshal(asMap["modifierExtension"], &out.ModifierExtension); err != nil {
			return err
		}

	}
	if len(asMap["base"]) > 0 {
		if err := go1.Unmarshal(asMap["base"], &out.Base); err != nil {
			return err
		}

	}
	if len(asMap["sliceIsConstraining"]) > 0 {
		if err := go1.Unmarshal(asMap["sliceIsConstraining"], &out.SliceIsConstraining); err != nil {
			return err
		}

	}
	if len(asMap["_sliceIsConstraining"]) > 0 {
		if err := go1.Unmarshal(asMap["_sliceIsConstraining"], &out.SliceIsConstrainingPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["orderMeaning"]) > 0 {
		if err := go1.Unmarshal(asMap["orderMeaning"], &out.OrderMeaning); err != nil {
			return err
		}

	}
	if len(asMap["_orderMeaning"]) > 0 {
		if err := go1.Unmarshal(asMap["_orderMeaning"], &out.OrderMeaningPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["valueAlternatives"]) > 0 {
		if err := go1.Unmarshal(asMap["valueAlternatives"], &out.ValueAlternatives); err != nil {
			return err
		}

	}
	if len(asMap["_valueAlternatives"]) > 0 {
		if err := go1.Unmarshal(asMap["_valueAlternatives"], &out.ValueAlternativesPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["representation"]) > 0 {
		if err := go1.Unmarshal(asMap["representation"], &out.Representation); err != nil {
			return err
		}

	}
	if len(asMap["_representation"]) > 0 {
		if err := go1.Unmarshal(asMap["_representation"], &out.RepresentationPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["sliceName"]) > 0 {
		if err := go1.Unmarshal(asMap["sliceName"], &out.SliceName); err != nil {
			return err
		}

	}
	if len(asMap["_sliceName"]) > 0 {
		if err := go1.Unmarshal(asMap["_sliceName"], &out.SliceNamePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["short"]) > 0 {
		if err := go1.Unmarshal(asMap["short"], &out.Short); err != nil {
			return err
		}

	}
	if len(asMap["_short"]) > 0 {
		if err := go1.Unmarshal(asMap["_short"], &out.ShortPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["alias"]) > 0 {
		if err := go1.Unmarshal(asMap["alias"], &out.Alias); err != nil {
			return err
		}

	}
	if len(asMap["_alias"]) > 0 {
		if err := go1.Unmarshal(asMap["_alias"], &out.AliasPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["max"]) > 0 {
		if err := go1.Unmarshal(asMap["max"], &out.Max); err != nil {
			return err
		}

	}
	if len(asMap["_max"]) > 0 {
		if err := go1.Unmarshal(asMap["_max"], &out.MaxPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["contentReference"]) > 0 {
		if err := go1.Unmarshal(asMap["contentReference"], &out.ContentReference); err != nil {
			return err
		}

	}
	if len(asMap["_contentReference"]) > 0 {
		if err := go1.Unmarshal(asMap["_contentReference"], &out.ContentReferencePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["minValueDate"]) > 0 {
		if err := go1.Unmarshal(asMap["minValueDate"], &out.MinValueDate); err == nil {
		}
	} else if len(asMap["minValueDateTime"]) > 0 {
		if err := go1.Unmarshal(asMap["minValueDateTime"], &out.MinValueDateTime); err == nil {
		}
	} else if len(asMap["minValueInstant"]) > 0 {
		if err := go1.Unmarshal(asMap["minValueInstant"], &out.MinValueInstant); err == nil {
		}
	} else if len(asMap["minValueTime"]) > 0 {
		if err := go1.Unmarshal(asMap["minValueTime"], &out.MinValueTime); err == nil {
		}
	} else if len(asMap["minValueDecimal"]) > 0 {
		if err := go1.Unmarshal(asMap["minValueDecimal"], &out.MinValueDecimal); err == nil {
		}
	} else if len(asMap["minValueInteger"]) > 0 {
		if err := go1.Unmarshal(asMap["minValueInteger"], &out.MinValueInteger); err == nil {
		}
	} else if len(asMap["minValueInteger64"]) > 0 {
		if err := go1.Unmarshal(asMap["minValueInteger64"], &out.MinValueInteger64); err == nil {
		}
	} else if len(asMap["minValuePositiveInt"]) > 0 {
		if err := go1.Unmarshal(asMap["minValuePositiveInt"], &out.MinValuePositiveInt); err == nil {
		}
	} else if len(asMap["minValueUnsignedInt"]) > 0 {
		if err := go1.Unmarshal(asMap["minValueUnsignedInt"], &out.MinValueUnsignedInt); err == nil {
		}
	} else if len(asMap["minValueQuantity"]) > 0 {
		if err := go1.Unmarshal(asMap["minValueQuantity"], &out.MinValueQuantity); err == nil {
		}
	} else {

	}
	if len(asMap["condition"]) > 0 {
		if err := go1.Unmarshal(asMap["condition"], &out.Condition); err != nil {
			return err
		}

	}
	if len(asMap["_condition"]) > 0 {
		if err := go1.Unmarshal(asMap["_condition"], &out.ConditionPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["label"]) > 0 {
		if err := go1.Unmarshal(asMap["label"], &out.Label); err != nil {
			return err
		}

	}
	if len(asMap["_label"]) > 0 {
		if err := go1.Unmarshal(asMap["_label"], &out.LabelPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["code"]) > 0 {
		if err := go1.Unmarshal(asMap["code"], &out.Code); err != nil {
			return err
		}

	}
	if len(asMap["mustSupport"]) > 0 {
		if err := go1.Unmarshal(asMap["mustSupport"], &out.MustSupport); err != nil {
			return err
		}

	}
	if len(asMap["_mustSupport"]) > 0 {
		if err := go1.Unmarshal(asMap["_mustSupport"], &out.MustSupportPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["binding"]) > 0 {
		if err := go1.Unmarshal(asMap["binding"], &out.Binding); err != nil {
			return err
		}

	}
	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["isModifierReason"]) > 0 {
		if err := go1.Unmarshal(asMap["isModifierReason"], &out.IsModifierReason); err != nil {
			return err
		}

	}
	if len(asMap["_isModifierReason"]) > 0 {
		if err := go1.Unmarshal(asMap["_isModifierReason"], &out.IsModifierReasonPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["mustHaveValue"]) > 0 {
		if err := go1.Unmarshal(asMap["mustHaveValue"], &out.MustHaveValue); err != nil {
			return err
		}

	}
	if len(asMap["_mustHaveValue"]) > 0 {
		if err := go1.Unmarshal(asMap["_mustHaveValue"], &out.MustHaveValuePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["comment"]) > 0 {
		if err := go1.Unmarshal(asMap["comment"], &out.Comment); err != nil {
			return err
		}

	}
	if len(asMap["_comment"]) > 0 {
		if err := go1.Unmarshal(asMap["_comment"], &out.CommentPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["defaultValueBase64Binary"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueBase64Binary"], &out.DefaultValueBase64Binary); err == nil {
		}
	} else if len(asMap["defaultValueBoolean"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueBoolean"], &out.DefaultValueBoolean); err == nil {
		}
	} else if len(asMap["defaultValueCanonical"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueCanonical"], &out.DefaultValueCanonical); err == nil {
		}
	} else if len(asMap["defaultValueCode"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueCode"], &out.DefaultValueCode); err == nil {
		}
	} else if len(asMap["defaultValueDate"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueDate"], &out.DefaultValueDate); err == nil {
		}
	} else if len(asMap["defaultValueDateTime"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueDateTime"], &out.DefaultValueDateTime); err == nil {
		}
	} else if len(asMap["defaultValueDecimal"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueDecimal"], &out.DefaultValueDecimal); err == nil {
		}
	} else if len(asMap["defaultValueId"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueId"], &out.DefaultValueId); err == nil {
		}
	} else if len(asMap["defaultValueInstant"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueInstant"], &out.DefaultValueInstant); err == nil {
		}
	} else if len(asMap["defaultValueInteger"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueInteger"], &out.DefaultValueInteger); err == nil {
		}
	} else if len(asMap["defaultValueInteger64"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueInteger64"], &out.DefaultValueInteger64); err == nil {
		}
	} else if len(asMap["defaultValueMarkdown"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueMarkdown"], &out.DefaultValueMarkdown); err == nil {
		}
	} else if len(asMap["defaultValueOid"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueOid"], &out.DefaultValueOid); err == nil {
		}
	} else if len(asMap["defaultValuePositiveInt"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValuePositiveInt"], &out.DefaultValuePositiveInt); err == nil {
		}
	} else if len(asMap["defaultValueString"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueString"], &out.DefaultValueString); err == nil {
		}
	} else if len(asMap["defaultValueTime"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueTime"], &out.DefaultValueTime); err == nil {
		}
	} else if len(asMap["defaultValueUnsignedInt"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueUnsignedInt"], &out.DefaultValueUnsignedInt); err == nil {
		}
	} else if len(asMap["defaultValueUri"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueUri"], &out.DefaultValueUri); err == nil {
		}
	} else if len(asMap["defaultValueUrl"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueUrl"], &out.DefaultValueUrl); err == nil {
		}
	} else if len(asMap["defaultValueUuid"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueUuid"], &out.DefaultValueUuid); err == nil {
		}
	} else if len(asMap["defaultValueAddress"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueAddress"], &out.DefaultValueAddress); err == nil {
		}
	} else if len(asMap["defaultValueAge"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueAge"], &out.DefaultValueAge); err == nil {
		}
	} else if len(asMap["defaultValueAnnotation"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueAnnotation"], &out.DefaultValueAnnotation); err == nil {
		}
	} else if len(asMap["defaultValueAttachment"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueAttachment"], &out.DefaultValueAttachment); err == nil {
		}
	} else if len(asMap["defaultValueCodeableConcept"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueCodeableConcept"], &out.DefaultValueCodeableConcept); err == nil {
		}
	} else if len(asMap["defaultValueCodeableReference"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueCodeableReference"], &out.DefaultValueCodeableReference); err == nil {
		}
	} else if len(asMap["defaultValueCoding"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueCoding"], &out.DefaultValueCoding); err == nil {
		}
	} else if len(asMap["defaultValueContactPoint"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueContactPoint"], &out.DefaultValueContactPoint); err == nil {
		}
	} else if len(asMap["defaultValueCount"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueCount"], &out.DefaultValueCount); err == nil {
		}
	} else if len(asMap["defaultValueDistance"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueDistance"], &out.DefaultValueDistance); err == nil {
		}
	} else if len(asMap["defaultValueDuration"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueDuration"], &out.DefaultValueDuration); err == nil {
		}
	} else if len(asMap["defaultValueHumanName"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueHumanName"], &out.DefaultValueHumanName); err == nil {
		}
	} else if len(asMap["defaultValueIdentifier"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueIdentifier"], &out.DefaultValueIdentifier); err == nil {
		}
	} else if len(asMap["defaultValueMoney"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueMoney"], &out.DefaultValueMoney); err == nil {
		}
	} else if len(asMap["defaultValuePeriod"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValuePeriod"], &out.DefaultValuePeriod); err == nil {
		}
	} else if len(asMap["defaultValueQuantity"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueQuantity"], &out.DefaultValueQuantity); err == nil {
		}
	} else if len(asMap["defaultValueRange"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueRange"], &out.DefaultValueRange); err == nil {
		}
	} else if len(asMap["defaultValueRatio"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueRatio"], &out.DefaultValueRatio); err == nil {
		}
	} else if len(asMap["defaultValueRatioRange"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueRatioRange"], &out.DefaultValueRatioRange); err == nil {
		}
	} else if len(asMap["defaultValueReference"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueReference"], &out.DefaultValueReference); err == nil {
		}
	} else if len(asMap["defaultValueSampledData"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueSampledData"], &out.DefaultValueSampledData); err == nil {
		}
	} else if len(asMap["defaultValueSignature"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueSignature"], &out.DefaultValueSignature); err == nil {
		}
	} else if len(asMap["defaultValueTiming"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueTiming"], &out.DefaultValueTiming); err == nil {
		}
	} else if len(asMap["defaultValueContactDetail"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueContactDetail"], &out.DefaultValueContactDetail); err == nil {
		}
	} else if len(asMap["defaultValueDataRequirement"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueDataRequirement"], &out.DefaultValueDataRequirement); err == nil {
		}
	} else if len(asMap["defaultValueExpression"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueExpression"], &out.DefaultValueExpression); err == nil {
		}
	} else if len(asMap["defaultValueParameterDefinition"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueParameterDefinition"], &out.DefaultValueParameterDefinition); err == nil {
		}
	} else if len(asMap["defaultValueRelatedArtifact"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueRelatedArtifact"], &out.DefaultValueRelatedArtifact); err == nil {
		}
	} else if len(asMap["defaultValueTriggerDefinition"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueTriggerDefinition"], &out.DefaultValueTriggerDefinition); err == nil {
		}
	} else if len(asMap["defaultValueUsageContext"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueUsageContext"], &out.DefaultValueUsageContext); err == nil {
		}
	} else if len(asMap["defaultValueAvailability"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueAvailability"], &out.DefaultValueAvailability); err == nil {
		}
	} else if len(asMap["defaultValueExtendedContactDetail"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueExtendedContactDetail"], &out.DefaultValueExtendedContactDetail); err == nil {
		}
	} else if len(asMap["defaultValueDosage"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueDosage"], &out.DefaultValueDosage); err == nil {
		}
	} else if len(asMap["defaultValueMeta"]) > 0 {
		if err := go1.Unmarshal(asMap["defaultValueMeta"], &out.DefaultValueMeta); err == nil {
		}
	} else {

	}
	if len(asMap["meaningWhenMissing"]) > 0 {
		if err := go1.Unmarshal(asMap["meaningWhenMissing"], &out.MeaningWhenMissing); err != nil {
			return err
		}

	}
	if len(asMap["_meaningWhenMissing"]) > 0 {
		if err := go1.Unmarshal(asMap["_meaningWhenMissing"], &out.MeaningWhenMissingPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["patternBase64Binary"]) > 0 {
		if err := go1.Unmarshal(asMap["patternBase64Binary"], &out.PatternBase64Binary); err == nil {
		}
	} else if len(asMap["patternBoolean"]) > 0 {
		if err := go1.Unmarshal(asMap["patternBoolean"], &out.PatternBoolean); err == nil {
		}
	} else if len(asMap["patternCanonical"]) > 0 {
		if err := go1.Unmarshal(asMap["patternCanonical"], &out.PatternCanonical); err == nil {
		}
	} else if len(asMap["patternCode"]) > 0 {
		if err := go1.Unmarshal(asMap["patternCode"], &out.PatternCode); err == nil {
		}
	} else if len(asMap["patternDate"]) > 0 {
		if err := go1.Unmarshal(asMap["patternDate"], &out.PatternDate); err == nil {
		}
	} else if len(asMap["patternDateTime"]) > 0 {
		if err := go1.Unmarshal(asMap["patternDateTime"], &out.PatternDateTime); err == nil {
		}
	} else if len(asMap["patternDecimal"]) > 0 {
		if err := go1.Unmarshal(asMap["patternDecimal"], &out.PatternDecimal); err == nil {
		}
	} else if len(asMap["patternId"]) > 0 {
		if err := go1.Unmarshal(asMap["patternId"], &out.PatternId); err == nil {
		}
	} else if len(asMap["patternInstant"]) > 0 {
		if err := go1.Unmarshal(asMap["patternInstant"], &out.PatternInstant); err == nil {
		}
	} else if len(asMap["patternInteger"]) > 0 {
		if err := go1.Unmarshal(asMap["patternInteger"], &out.PatternInteger); err == nil {
		}
	} else if len(asMap["patternInteger64"]) > 0 {
		if err := go1.Unmarshal(asMap["patternInteger64"], &out.PatternInteger64); err == nil {
		}
	} else if len(asMap["patternMarkdown"]) > 0 {
		if err := go1.Unmarshal(asMap["patternMarkdown"], &out.PatternMarkdown); err == nil {
		}
	} else if len(asMap["patternOid"]) > 0 {
		if err := go1.Unmarshal(asMap["patternOid"], &out.PatternOid); err == nil {
		}
	} else if len(asMap["patternPositiveInt"]) > 0 {
		if err := go1.Unmarshal(asMap["patternPositiveInt"], &out.PatternPositiveInt); err == nil {
		}
	} else if len(asMap["patternString"]) > 0 {
		if err := go1.Unmarshal(asMap["patternString"], &out.PatternString); err == nil {
		}
	} else if len(asMap["patternTime"]) > 0 {
		if err := go1.Unmarshal(asMap["patternTime"], &out.PatternTime); err == nil {
		}
	} else if len(asMap["patternUnsignedInt"]) > 0 {
		if err := go1.Unmarshal(asMap["patternUnsignedInt"], &out.PatternUnsignedInt); err == nil {
		}
	} else if len(asMap["patternUri"]) > 0 {
		if err := go1.Unmarshal(asMap["patternUri"], &out.PatternUri); err == nil {
		}
	} else if len(asMap["patternUrl"]) > 0 {
		if err := go1.Unmarshal(asMap["patternUrl"], &out.PatternUrl); err == nil {
		}
	} else if len(asMap["patternUuid"]) > 0 {
		if err := go1.Unmarshal(asMap["patternUuid"], &out.PatternUuid); err == nil {
		}
	} else if len(asMap["patternAddress"]) > 0 {
		if err := go1.Unmarshal(asMap["patternAddress"], &out.PatternAddress); err == nil {
		}
	} else if len(asMap["patternAge"]) > 0 {
		if err := go1.Unmarshal(asMap["patternAge"], &out.PatternAge); err == nil {
		}
	} else if len(asMap["patternAnnotation"]) > 0 {
		if err := go1.Unmarshal(asMap["patternAnnotation"], &out.PatternAnnotation); err == nil {
		}
	} else if len(asMap["patternAttachment"]) > 0 {
		if err := go1.Unmarshal(asMap["patternAttachment"], &out.PatternAttachment); err == nil {
		}
	} else if len(asMap["patternCodeableConcept"]) > 0 {
		if err := go1.Unmarshal(asMap["patternCodeableConcept"], &out.PatternCodeableConcept); err == nil {
		}
	} else if len(asMap["patternCodeableReference"]) > 0 {
		if err := go1.Unmarshal(asMap["patternCodeableReference"], &out.PatternCodeableReference); err == nil {
		}
	} else if len(asMap["patternCoding"]) > 0 {
		if err := go1.Unmarshal(asMap["patternCoding"], &out.PatternCoding); err == nil {
		}
	} else if len(asMap["patternContactPoint"]) > 0 {
		if err := go1.Unmarshal(asMap["patternContactPoint"], &out.PatternContactPoint); err == nil {
		}
	} else if len(asMap["patternCount"]) > 0 {
		if err := go1.Unmarshal(asMap["patternCount"], &out.PatternCount); err == nil {
		}
	} else if len(asMap["patternDistance"]) > 0 {
		if err := go1.Unmarshal(asMap["patternDistance"], &out.PatternDistance); err == nil {
		}
	} else if len(asMap["patternDuration"]) > 0 {
		if err := go1.Unmarshal(asMap["patternDuration"], &out.PatternDuration); err == nil {
		}
	} else if len(asMap["patternHumanName"]) > 0 {
		if err := go1.Unmarshal(asMap["patternHumanName"], &out.PatternHumanName); err == nil {
		}
	} else if len(asMap["patternIdentifier"]) > 0 {
		if err := go1.Unmarshal(asMap["patternIdentifier"], &out.PatternIdentifier); err == nil {
		}
	} else if len(asMap["patternMoney"]) > 0 {
		if err := go1.Unmarshal(asMap["patternMoney"], &out.PatternMoney); err == nil {
		}
	} else if len(asMap["patternPeriod"]) > 0 {
		if err := go1.Unmarshal(asMap["patternPeriod"], &out.PatternPeriod); err == nil {
		}
	} else if len(asMap["patternQuantity"]) > 0 {
		if err := go1.Unmarshal(asMap["patternQuantity"], &out.PatternQuantity); err == nil {
		}
	} else if len(asMap["patternRange"]) > 0 {
		if err := go1.Unmarshal(asMap["patternRange"], &out.PatternRange); err == nil {
		}
	} else if len(asMap["patternRatio"]) > 0 {
		if err := go1.Unmarshal(asMap["patternRatio"], &out.PatternRatio); err == nil {
		}
	} else if len(asMap["patternRatioRange"]) > 0 {
		if err := go1.Unmarshal(asMap["patternRatioRange"], &out.PatternRatioRange); err == nil {
		}
	} else if len(asMap["patternReference"]) > 0 {
		if err := go1.Unmarshal(asMap["patternReference"], &out.PatternReference); err == nil {
		}
	} else if len(asMap["patternSampledData"]) > 0 {
		if err := go1.Unmarshal(asMap["patternSampledData"], &out.PatternSampledData); err == nil {
		}
	} else if len(asMap["patternSignature"]) > 0 {
		if err := go1.Unmarshal(asMap["patternSignature"], &out.PatternSignature); err == nil {
		}
	} else if len(asMap["patternTiming"]) > 0 {
		if err := go1.Unmarshal(asMap["patternTiming"], &out.PatternTiming); err == nil {
		}
	} else if len(asMap["patternContactDetail"]) > 0 {
		if err := go1.Unmarshal(asMap["patternContactDetail"], &out.PatternContactDetail); err == nil {
		}
	} else if len(asMap["patternDataRequirement"]) > 0 {
		if err := go1.Unmarshal(asMap["patternDataRequirement"], &out.PatternDataRequirement); err == nil {
		}
	} else if len(asMap["patternExpression"]) > 0 {
		if err := go1.Unmarshal(asMap["patternExpression"], &out.PatternExpression); err == nil {
		}
	} else if len(asMap["patternParameterDefinition"]) > 0 {
		if err := go1.Unmarshal(asMap["patternParameterDefinition"], &out.PatternParameterDefinition); err == nil {
		}
	} else if len(asMap["patternRelatedArtifact"]) > 0 {
		if err := go1.Unmarshal(asMap["patternRelatedArtifact"], &out.PatternRelatedArtifact); err == nil {
		}
	} else if len(asMap["patternTriggerDefinition"]) > 0 {
		if err := go1.Unmarshal(asMap["patternTriggerDefinition"], &out.PatternTriggerDefinition); err == nil {
		}
	} else if len(asMap["patternUsageContext"]) > 0 {
		if err := go1.Unmarshal(asMap["patternUsageContext"], &out.PatternUsageContext); err == nil {
		}
	} else if len(asMap["patternAvailability"]) > 0 {
		if err := go1.Unmarshal(asMap["patternAvailability"], &out.PatternAvailability); err == nil {
		}
	} else if len(asMap["patternExtendedContactDetail"]) > 0 {
		if err := go1.Unmarshal(asMap["patternExtendedContactDetail"], &out.PatternExtendedContactDetail); err == nil {
		}
	} else if len(asMap["patternDosage"]) > 0 {
		if err := go1.Unmarshal(asMap["patternDosage"], &out.PatternDosage); err == nil {
		}
	} else if len(asMap["patternMeta"]) > 0 {
		if err := go1.Unmarshal(asMap["patternMeta"], &out.PatternMeta); err == nil {
		}
	} else {

	}
	if len(asMap["isSummary"]) > 0 {
		if err := go1.Unmarshal(asMap["isSummary"], &out.IsSummary); err != nil {
			return err
		}

	}
	if len(asMap["_isSummary"]) > 0 {
		if err := go1.Unmarshal(asMap["_isSummary"], &out.IsSummaryPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["mapping"]) > 0 {
		if err := go1.Unmarshal(asMap["mapping"], &out.Mapping); err != nil {
			return err
		}

	}
	if len(asMap["slicing"]) > 0 {
		if err := go1.Unmarshal(asMap["slicing"], &out.Slicing); err != nil {
			return err
		}

	}
	if len(asMap["definition"]) > 0 {
		if err := go1.Unmarshal(asMap["definition"], &out.Definition); err != nil {
			return err
		}

	}
	if len(asMap["_definition"]) > 0 {
		if err := go1.Unmarshal(asMap["_definition"], &out.DefinitionPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["example"]) > 0 {
		if err := go1.Unmarshal(asMap["example"], &out.Example); err != nil {
			return err
		}

	}
	if len(asMap["maxLength"]) > 0 {
		if err := go1.Unmarshal(asMap["maxLength"], &out.MaxLength); err != nil {
			return err
		}

	}
	if len(asMap["_maxLength"]) > 0 {
		if err := go1.Unmarshal(asMap["_maxLength"], &out.MaxLengthPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["isModifier"]) > 0 {
		if err := go1.Unmarshal(asMap["isModifier"], &out.IsModifier); err != nil {
			return err
		}

	}
	if len(asMap["_isModifier"]) > 0 {
		if err := go1.Unmarshal(asMap["_isModifier"], &out.IsModifierPrimitiveExtension); err != nil {
			return err
		}
	}

	if err := go1.Unmarshal(asMap["path"], &out.Path); err != nil {
		return err
	}
	if len(asMap["_path"]) > 0 {
		if err := go1.Unmarshal(asMap["_path"], &out.PathPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["fixedBase64Binary"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedBase64Binary"], &out.FixedBase64Binary); err == nil {
		}
	} else if len(asMap["fixedBoolean"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedBoolean"], &out.FixedBoolean); err == nil {
		}
	} else if len(asMap["fixedCanonical"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedCanonical"], &out.FixedCanonical); err == nil {
		}
	} else if len(asMap["fixedCode"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedCode"], &out.FixedCode); err == nil {
		}
	} else if len(asMap["fixedDate"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedDate"], &out.FixedDate); err == nil {
		}
	} else if len(asMap["fixedDateTime"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedDateTime"], &out.FixedDateTime); err == nil {
		}
	} else if len(asMap["fixedDecimal"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedDecimal"], &out.FixedDecimal); err == nil {
		}
	} else if len(asMap["fixedId"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedId"], &out.FixedId); err == nil {
		}
	} else if len(asMap["fixedInstant"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedInstant"], &out.FixedInstant); err == nil {
		}
	} else if len(asMap["fixedInteger"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedInteger"], &out.FixedInteger); err == nil {
		}
	} else if len(asMap["fixedInteger64"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedInteger64"], &out.FixedInteger64); err == nil {
		}
	} else if len(asMap["fixedMarkdown"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedMarkdown"], &out.FixedMarkdown); err == nil {
		}
	} else if len(asMap["fixedOid"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedOid"], &out.FixedOid); err == nil {
		}
	} else if len(asMap["fixedPositiveInt"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedPositiveInt"], &out.FixedPositiveInt); err == nil {
		}
	} else if len(asMap["fixedString"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedString"], &out.FixedString); err == nil {
		}
	} else if len(asMap["fixedTime"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedTime"], &out.FixedTime); err == nil {
		}
	} else if len(asMap["fixedUnsignedInt"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedUnsignedInt"], &out.FixedUnsignedInt); err == nil {
		}
	} else if len(asMap["fixedUri"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedUri"], &out.FixedUri); err == nil {
		}
	} else if len(asMap["fixedUrl"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedUrl"], &out.FixedUrl); err == nil {
		}
	} else if len(asMap["fixedUuid"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedUuid"], &out.FixedUuid); err == nil {
		}
	} else if len(asMap["fixedAddress"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedAddress"], &out.FixedAddress); err == nil {
		}
	} else if len(asMap["fixedAge"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedAge"], &out.FixedAge); err == nil {
		}
	} else if len(asMap["fixedAnnotation"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedAnnotation"], &out.FixedAnnotation); err == nil {
		}
	} else if len(asMap["fixedAttachment"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedAttachment"], &out.FixedAttachment); err == nil {
		}
	} else if len(asMap["fixedCodeableConcept"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedCodeableConcept"], &out.FixedCodeableConcept); err == nil {
		}
	} else if len(asMap["fixedCodeableReference"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedCodeableReference"], &out.FixedCodeableReference); err == nil {
		}
	} else if len(asMap["fixedCoding"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedCoding"], &out.FixedCoding); err == nil {
		}
	} else if len(asMap["fixedContactPoint"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedContactPoint"], &out.FixedContactPoint); err == nil {
		}
	} else if len(asMap["fixedCount"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedCount"], &out.FixedCount); err == nil {
		}
	} else if len(asMap["fixedDistance"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedDistance"], &out.FixedDistance); err == nil {
		}
	} else if len(asMap["fixedDuration"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedDuration"], &out.FixedDuration); err == nil {
		}
	} else if len(asMap["fixedHumanName"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedHumanName"], &out.FixedHumanName); err == nil {
		}
	} else if len(asMap["fixedIdentifier"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedIdentifier"], &out.FixedIdentifier); err == nil {
		}
	} else if len(asMap["fixedMoney"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedMoney"], &out.FixedMoney); err == nil {
		}
	} else if len(asMap["fixedPeriod"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedPeriod"], &out.FixedPeriod); err == nil {
		}
	} else if len(asMap["fixedQuantity"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedQuantity"], &out.FixedQuantity); err == nil {
		}
	} else if len(asMap["fixedRange"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedRange"], &out.FixedRange); err == nil {
		}
	} else if len(asMap["fixedRatio"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedRatio"], &out.FixedRatio); err == nil {
		}
	} else if len(asMap["fixedRatioRange"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedRatioRange"], &out.FixedRatioRange); err == nil {
		}
	} else if len(asMap["fixedReference"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedReference"], &out.FixedReference); err == nil {
		}
	} else if len(asMap["fixedSampledData"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedSampledData"], &out.FixedSampledData); err == nil {
		}
	} else if len(asMap["fixedSignature"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedSignature"], &out.FixedSignature); err == nil {
		}
	} else if len(asMap["fixedTiming"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedTiming"], &out.FixedTiming); err == nil {
		}
	} else if len(asMap["fixedContactDetail"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedContactDetail"], &out.FixedContactDetail); err == nil {
		}
	} else if len(asMap["fixedDataRequirement"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedDataRequirement"], &out.FixedDataRequirement); err == nil {
		}
	} else if len(asMap["fixedExpression"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedExpression"], &out.FixedExpression); err == nil {
		}
	} else if len(asMap["fixedParameterDefinition"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedParameterDefinition"], &out.FixedParameterDefinition); err == nil {
		}
	} else if len(asMap["fixedRelatedArtifact"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedRelatedArtifact"], &out.FixedRelatedArtifact); err == nil {
		}
	} else if len(asMap["fixedTriggerDefinition"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedTriggerDefinition"], &out.FixedTriggerDefinition); err == nil {
		}
	} else if len(asMap["fixedUsageContext"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedUsageContext"], &out.FixedUsageContext); err == nil {
		}
	} else if len(asMap["fixedAvailability"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedAvailability"], &out.FixedAvailability); err == nil {
		}
	} else if len(asMap["fixedExtendedContactDetail"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedExtendedContactDetail"], &out.FixedExtendedContactDetail); err == nil {
		}
	} else if len(asMap["fixedDosage"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedDosage"], &out.FixedDosage); err == nil {
		}
	} else if len(asMap["fixedMeta"]) > 0 {
		if err := go1.Unmarshal(asMap["fixedMeta"], &out.FixedMeta); err == nil {
		}
	} else {

	}
	if len(asMap["type"]) > 0 {
		if err := go1.Unmarshal(asMap["type"], &out.Type); err != nil {
			return err
		}

	}
	if len(asMap["maxValueDate"]) > 0 {
		if err := go1.Unmarshal(asMap["maxValueDate"], &out.MaxValueDate); err == nil {
		}
	} else if len(asMap["maxValueDateTime"]) > 0 {
		if err := go1.Unmarshal(asMap["maxValueDateTime"], &out.MaxValueDateTime); err == nil {
		}
	} else if len(asMap["maxValueInstant"]) > 0 {
		if err := go1.Unmarshal(asMap["maxValueInstant"], &out.MaxValueInstant); err == nil {
		}
	} else if len(asMap["maxValueTime"]) > 0 {
		if err := go1.Unmarshal(asMap["maxValueTime"], &out.MaxValueTime); err == nil {
		}
	} else if len(asMap["maxValueDecimal"]) > 0 {
		if err := go1.Unmarshal(asMap["maxValueDecimal"], &out.MaxValueDecimal); err == nil {
		}
	} else if len(asMap["maxValueInteger"]) > 0 {
		if err := go1.Unmarshal(asMap["maxValueInteger"], &out.MaxValueInteger); err == nil {
		}
	} else if len(asMap["maxValueInteger64"]) > 0 {
		if err := go1.Unmarshal(asMap["maxValueInteger64"], &out.MaxValueInteger64); err == nil {
		}
	} else if len(asMap["maxValuePositiveInt"]) > 0 {
		if err := go1.Unmarshal(asMap["maxValuePositiveInt"], &out.MaxValuePositiveInt); err == nil {
		}
	} else if len(asMap["maxValueUnsignedInt"]) > 0 {
		if err := go1.Unmarshal(asMap["maxValueUnsignedInt"], &out.MaxValueUnsignedInt); err == nil {
		}
	} else if len(asMap["maxValueQuantity"]) > 0 {
		if err := go1.Unmarshal(asMap["maxValueQuantity"], &out.MaxValueQuantity); err == nil {
		}
	} else {

	}
	if len(asMap["requirements"]) > 0 {
		if err := go1.Unmarshal(asMap["requirements"], &out.Requirements); err != nil {
			return err
		}

	}
	if len(asMap["_requirements"]) > 0 {
		if err := go1.Unmarshal(asMap["_requirements"], &out.RequirementsPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["min"]) > 0 {
		if err := go1.Unmarshal(asMap["min"], &out.Min); err != nil {
			return err
		}

	}
	if len(asMap["_min"]) > 0 {
		if err := go1.Unmarshal(asMap["_min"], &out.MinPrimitiveExtension); err != nil {
			return err
		}
	}

	return nil
}

type ElementDefinition struct {
	Constraint        *ElementDefinitionConstraint `bson:",omitempty" json:"constraint,omitempty"`
	ModifierExtension []*Extension                 `bson:",omitempty" json:"modifierExtension,omitempty"` /*
	May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.

	Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	*/
	Base                                  *ElementDefinitionBase `bson:",omitempty" json:"base,omitempty"`
	SliceIsConstraining                   bool                   `bson:",omitempty" json:"sliceIsConstraining,omitempty"` // If true, indicates that this slice definition is constraining a slice definition with the same name in an inherited profile. If false, the slice is not overriding any slice in an inherited profile. If missing, the slice might or might not be overriding a slice in an inherited profile, depending on the sliceName.
	SliceIsConstrainingPrimitiveExtension *PrimitiveExtension    `bson:"sliceIsConstraining_extension,omitempty" json:"_sliceIsConstraining,omitempty"`
	OrderMeaning                          string                 `bson:",omitempty" json:"orderMeaning,omitempty"` // If present, indicates that the order of the repeating element has meaning and describes what that meaning is.  If absent, it means that the order of the element has no meaning.
	OrderMeaningPrimitiveExtension        *PrimitiveExtension    `bson:"orderMeaning_extension,omitempty" json:"_orderMeaning,omitempty"`
	ValueAlternatives                     []string               `bson:",omitempty" json:"valueAlternatives,omitempty"` // Specifies a list of extensions that can appear in place of a primitive value.
	ValueAlternativesPrimitiveExtension   *PrimitiveExtension    `bson:"valueAlternatives_extension,omitempty" json:"_valueAlternatives,omitempty"`
	Representation                        []string               `bson:",omitempty" json:"representation,omitempty"` // Codes that define how this element is represented in instances, when the deviation varies from the normal case. No extensions are allowed on elements with a representation of 'xmlAttr', no matter what FHIR serialization format is used.
	RepresentationPrimitiveExtension      *PrimitiveExtension    `bson:"representation_extension,omitempty" json:"_representation,omitempty"`
	SliceName                             string                 `bson:",omitempty" json:"sliceName,omitempty"` // The name of this element definition slice, when slicing is working. The name must be a token with no dots or spaces. This is a unique name referring to a specific set of constraints applied to this element, used to provide a name to different slices of the same element.
	SliceNamePrimitiveExtension           *PrimitiveExtension    `bson:"sliceName_extension,omitempty" json:"_sliceName,omitempty"`
	Short                                 string                 `bson:",omitempty" json:"short,omitempty"` // A concise description of what this element means (e.g. for use in autogenerated summaries).
	ShortPrimitiveExtension               *PrimitiveExtension    `bson:"short_extension,omitempty" json:"_short,omitempty"`
	Alias                                 []string               `bson:",omitempty" json:"alias,omitempty"` // Identifies additional names by which this element might also be known.
	AliasPrimitiveExtension               *PrimitiveExtension    `bson:"alias_extension,omitempty" json:"_alias,omitempty"`
	Max                                   string                 `bson:",omitempty" json:"max,omitempty"` // The maximum number of times this element is permitted to appear in the instance.
	MaxPrimitiveExtension                 *PrimitiveExtension    `bson:"max_extension,omitempty" json:"_max,omitempty"`
	ContentReference                      string                 `bson:",omitempty" json:"contentReference,omitempty"` // Identifies an element defined elsewhere in the definition whose content rules should be applied to the current element. ContentReferences bring across all the rules that are in the ElementDefinition for the element, including definitions, cardinality constraints, bindings, invariants etc.
	ContentReferencePrimitiveExtension    *PrimitiveExtension    `bson:"contentReference_extension,omitempty" json:"_contentReference,omitempty"`
	ElementDefinitionMinValuex            `binding:"omitempty" bson:",omitempty" json:",omitempty"`
	Condition                             []*primitive.ObjectID     `bson:",omitempty" json:"condition,omitempty"` // A reference to an invariant that may make additional statements about the cardinality or value in the instance.
	ConditionPrimitiveExtension           *PrimitiveExtension       `bson:"condition_extension,omitempty" json:"_condition,omitempty"`
	Label                                 string                    `bson:",omitempty" json:"label,omitempty"` // A single preferred label which is the text to display beside the element indicating its meaning or to use to prompt for the element in a user display or form.
	LabelPrimitiveExtension               *PrimitiveExtension       `bson:"label_extension,omitempty" json:"_label,omitempty"`
	Code                                  []*Coding                 `bson:",omitempty" json:"code,omitempty"`        // A code that has the same meaning as the element in a particular terminology.
	MustSupport                           bool                      `bson:",omitempty" json:"mustSupport,omitempty"` // If true, implementations that produce or consume resources SHALL provide "support" for the element in some meaningful way. Note that this is being phased out and replaced by obligations (see below).  If false, the element may be ignored and not supported. If false, whether to populate or use the data element in any way is at the discretion of the implementation.
	MustSupportPrimitiveExtension         *PrimitiveExtension       `bson:"mustSupport_extension,omitempty" json:"_mustSupport,omitempty"`
	Binding                               *ElementDefinitionBinding `bson:",omitempty" json:"binding,omitempty"`
	Extension                             []*Extension              `bson:",omitempty" json:"extension,omitempty"`        // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	IsModifierReason                      string                    `bson:",omitempty" json:"isModifierReason,omitempty"` // Explains how that element affects the interpretation of the resource or element that contains it.
	IsModifierReasonPrimitiveExtension    *PrimitiveExtension       `bson:"isModifierReason_extension,omitempty" json:"_isModifierReason,omitempty"`
	Id                                    *primitive.ObjectID       `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension                  *PrimitiveExtension       `bson:"id_extension,omitempty" json:"_id,omitempty"`
	MustHaveValue                         bool                      `bson:",omitempty" json:"mustHaveValue,omitempty"` // Specifies for a primitive data type that the value of the data type cannot be replaced by an extension.
	MustHaveValuePrimitiveExtension       *PrimitiveExtension       `bson:"mustHaveValue_extension,omitempty" json:"_mustHaveValue,omitempty"`
	Comment                               string                    `bson:",omitempty" json:"comment,omitempty"` // Explanatory notes and implementation guidance about the data element, including notes about how to use the data properly, exceptions to proper use, etc. (Note: The text you are reading is specified in ElementDefinition.comment).
	CommentPrimitiveExtension             *PrimitiveExtension       `bson:"comment_extension,omitempty" json:"_comment,omitempty"`
	ElementDefinitionDefaultValuex        `binding:"omitempty" bson:",omitempty" json:",omitempty"`
	MeaningWhenMissing                    string              `bson:",omitempty" json:"meaningWhenMissing,omitempty"` // The Implicit meaning that is to be understood when this element is missing (e.g. 'when this element is missing, the period is ongoing').
	MeaningWhenMissingPrimitiveExtension  *PrimitiveExtension `bson:"meaningWhenMissing_extension,omitempty" json:"_meaningWhenMissing,omitempty"`
	ElementDefinitionPatternx             `binding:"omitempty" bson:",omitempty" json:",omitempty"`
	IsSummary                             bool                      `bson:",omitempty" json:"isSummary,omitempty"` // Whether the element should be included if a client requests a search with the parameter _summary=true.
	IsSummaryPrimitiveExtension           *PrimitiveExtension       `bson:"isSummary_extension,omitempty" json:"_isSummary,omitempty"`
	Mapping                               *ElementDefinitionMapping `bson:",omitempty" json:"mapping,omitempty"`
	Slicing                               *ElementDefinitionSlicing `bson:",omitempty" json:"slicing,omitempty"`
	Definition                            string                    `bson:",omitempty" json:"definition,omitempty"` // Provides a complete explanation of the meaning of the data element for human readability.  For the case of elements derived from existing elements (e.g. constraints), the definition SHALL be consistent with the base definition, but convey the meaning of the element in the particular context of use of the resource. (Note: The text you are reading is specified in ElementDefinition.definition).
	DefinitionPrimitiveExtension          *PrimitiveExtension       `bson:"definition_extension,omitempty" json:"_definition,omitempty"`
	Example                               *ElementDefinitionExample `bson:",omitempty" json:"example,omitempty"`
	MaxLength                             int                       `bson:",omitempty" json:"maxLength,omitempty"` // Indicates the maximum length in characters that is permitted to be present in conformant instances and which is expected to be supported by conformant consumers that support the element. ```maxLength``` SHOULD only be used on primitive data types that have a string representation (see [http://hl7.org/fhir/StructureDefinition/structuredefinition-type-characteristics](http://hl7.org/fhir/extensions/StructureDefinition-structuredefinition-type-characteristics.html)).
	MaxLengthPrimitiveExtension           *PrimitiveExtension       `bson:"maxLength_extension,omitempty" json:"_maxLength,omitempty"`
	IsModifier                            bool                      `bson:",omitempty" json:"isModifier,omitempty"` // If true, the value of this element affects the interpretation of the element or resource that contains it, and the value of the element cannot be ignored. Typically, this is used for status, negation and qualification codes. The effect of this is that the element cannot be ignored by systems: they SHALL either recognize the element and process it, and/or a pre-determination has been made that it is not relevant to their particular system. When used on the root element in an extension definition, this indicates whether or not the extension is a modifier extension.
	IsModifierPrimitiveExtension          *PrimitiveExtension       `bson:"isModifier_extension,omitempty" json:"_isModifier,omitempty"`
	Path                                  string                    `binding:"required" bson:",omitempty" json:"path,omitempty"` // The path identifies the element and is expressed as a "."-separated list of ancestor elements, beginning with the name of the resource or extension.
	PathPrimitiveExtension                *PrimitiveExtension       `bson:"path_extension,omitempty" json:"_path,omitempty"`
	ElementDefinitionFixedx               `binding:"omitempty" bson:",omitempty" json:",omitempty"`
	Type                                  *ElementDefinitionType `bson:",omitempty" json:"type,omitempty"`
	ElementDefinitionMaxValuex            `binding:"omitempty" bson:",omitempty" json:",omitempty"`
	Requirements                          string              `bson:",omitempty" json:"requirements,omitempty"` // This element is for traceability of why the element was created and why the constraints exist as they do. This may be used to point to source materials or specifications that drove the structure of this element.
	RequirementsPrimitiveExtension        *PrimitiveExtension `bson:"requirements_extension,omitempty" json:"_requirements,omitempty"`
	Min                                   int                 `bson:",omitempty" json:"min,omitempty"` // The minimum number of times this element SHALL appear in the instance.
	MinPrimitiveExtension                 *PrimitiveExtension `bson:"min_extension,omitempty" json:"_min,omitempty"`
	ResourceType                          string              `bson:"-" json:"resourceType,omitempty"`
}

func (out *Expression) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Expression\"" {
		return fmt.Errorf("resourceType is not %s", "Expression")
	}
	if len(asMap["language"]) > 0 {
		if err := go1.Unmarshal(asMap["language"], &out.Language); err != nil {
			return err
		}

	}
	if len(asMap["_language"]) > 0 {
		if err := go1.Unmarshal(asMap["_language"], &out.LanguagePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["expression"]) > 0 {
		if err := go1.Unmarshal(asMap["expression"], &out.Expression); err != nil {
			return err
		}

	}
	if len(asMap["_expression"]) > 0 {
		if err := go1.Unmarshal(asMap["_expression"], &out.ExpressionPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["reference"]) > 0 {
		if err := go1.Unmarshal(asMap["reference"], &out.Reference); err != nil {
			return err
		}

	}
	if len(asMap["_reference"]) > 0 {
		if err := go1.Unmarshal(asMap["_reference"], &out.ReferencePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["description"]) > 0 {
		if err := go1.Unmarshal(asMap["description"], &out.Description); err != nil {
			return err
		}

	}
	if len(asMap["_description"]) > 0 {
		if err := go1.Unmarshal(asMap["_description"], &out.DescriptionPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["name"]) > 0 {
		if err := go1.Unmarshal(asMap["name"], &out.Name); err != nil {
			return err
		}

	}
	if len(asMap["_name"]) > 0 {
		if err := go1.Unmarshal(asMap["_name"], &out.NamePrimitiveExtension); err != nil {
			return err
		}
	}

	return nil
}

type Expression struct {
	Language                      string              `bson:",omitempty" json:"language,omitempty"` // The media type of the language for the expression.
	LanguagePrimitiveExtension    *PrimitiveExtension `bson:"language_extension,omitempty" json:"_language,omitempty"`
	Expression                    string              `bson:",omitempty" json:"expression,omitempty"` // An expression in the specified language that returns a value.
	ExpressionPrimitiveExtension  *PrimitiveExtension `bson:"expression_extension,omitempty" json:"_expression,omitempty"`
	Reference                     string              `bson:",omitempty" json:"reference,omitempty"` // A URI that defines where the expression is found.
	ReferencePrimitiveExtension   *PrimitiveExtension `bson:"reference_extension,omitempty" json:"_reference,omitempty"`
	Id                            *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension          *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension                     []*Extension        `bson:",omitempty" json:"extension,omitempty"`   // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Description                   string              `bson:",omitempty" json:"description,omitempty"` // A brief, natural language description of the condition that effectively communicates the intended semantics.
	DescriptionPrimitiveExtension *PrimitiveExtension `bson:"description_extension,omitempty" json:"_description,omitempty"`
	Name                          string              `bson:",omitempty" json:"name,omitempty"` // A short name assigned to the expression to allow for multiple reuse of the expression in the context where it is defined.
	NamePrimitiveExtension        *PrimitiveExtension `bson:"name_extension,omitempty" json:"_name,omitempty"`
	ResourceType                  string              `bson:"-" json:"resourceType,omitempty"`
}

func (out *ExtendedContactDetail) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"ExtendedContactDetail\"" {
		return fmt.Errorf("resourceType is not %s", "ExtendedContactDetail")
	}
	if len(asMap["organization"]) > 0 {
		if err := go1.Unmarshal(asMap["organization"], &out.Organization); err != nil {
			return err
		}

	}
	if len(asMap["period"]) > 0 {
		if err := go1.Unmarshal(asMap["period"], &out.Period); err != nil {
			return err
		}

	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["purpose"]) > 0 {
		if err := go1.Unmarshal(asMap["purpose"], &out.Purpose); err != nil {
			return err
		}

	}
	if len(asMap["name"]) > 0 {
		if err := go1.Unmarshal(asMap["name"], &out.Name); err != nil {
			return err
		}

	}
	if len(asMap["telecom"]) > 0 {
		if err := go1.Unmarshal(asMap["telecom"], &out.Telecom); err != nil {
			return err
		}

	}
	if len(asMap["address"]) > 0 {
		if err := go1.Unmarshal(asMap["address"], &out.Address); err != nil {
			return err
		}

	}
	return nil
}

type ExtendedContactDetail struct {
	Organization         *Reference          `bson:",omitempty" json:"organization,omitempty"` // This contact detail is handled/monitored by a specific organization. If the name is provided in the contact, then it is referring to the named individual within this organization.
	Period               *Period             `bson:",omitempty" json:"period,omitempty"`       // Period that this contact was valid for usage.
	Id                   *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`        // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension            []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Purpose              *CodeableConcept    `bson:",omitempty" json:"purpose,omitempty"`   // The purpose/type of contact.
	Name                 []*HumanName        `bson:",omitempty" json:"name,omitempty"`      // The name of an individual to contact, some types of contact detail are usually blank.
	Telecom              []*ContactPoint     `bson:",omitempty" json:"telecom,omitempty"`   // The contact details application for the purpose defined.
	Address              *Address            `bson:",omitempty" json:"address,omitempty"`   // Address for the contact.
	ResourceType         string              `bson:"-" json:"resourceType,omitempty"`
}

type ExtensionValuex struct {
	ValueBase64Binary          Base64Binary           `binding:"omitempty" bson:",omitempty" json:"valueBase64Binary,omitempty"`
	ValueBoolean               bool                   `binding:"omitempty" bson:",omitempty" json:"valueBoolean,omitempty"`
	ValueCanonical             string                 `binding:"omitempty" bson:",omitempty" json:"valueCanonical,omitempty"`
	ValueCode                  string                 `binding:"omitempty" bson:",omitempty" json:"valueCode,omitempty"`
	ValueDate                  *Date                  `binding:"omitempty" bson:",omitempty" json:"valueDate,omitempty"`
	ValueDateTime              *DateTime              `binding:"omitempty" bson:",omitempty" json:"valueDateTime,omitempty"`
	ValueDecimal               float64                `binding:"omitempty" bson:",omitempty" json:"valueDecimal,omitempty"`
	ValueId                    *primitive.ObjectID    `binding:"omitempty" bson:",omitempty" json:"valueId,omitempty"`
	ValueInstant               *time.Time             `binding:"omitempty" bson:",omitempty" json:"valueInstant,omitempty"`
	ValueInteger               int                    `binding:"omitempty" bson:",omitempty" json:"valueInteger,omitempty"`
	ValueInteger64             int64                  `binding:"omitempty" bson:",omitempty" json:"valueInteger64,omitempty"`
	ValueMarkdown              string                 `binding:"omitempty" bson:",omitempty" json:"valueMarkdown,omitempty"`
	ValueOid                   string                 `binding:"omitempty" bson:",omitempty" json:"valueOid,omitempty"`
	ValuePositiveInt           int                    `binding:"omitempty" bson:",omitempty" json:"valuePositiveInt,omitempty"`
	ValueString                string                 `binding:"omitempty" bson:",omitempty" json:"valueString,omitempty"`
	ValueTime                  *Time                  `binding:"omitempty" bson:",omitempty" json:"valueTime,omitempty"`
	ValueUnsignedInt           int                    `binding:"omitempty" bson:",omitempty" json:"valueUnsignedInt,omitempty"`
	ValueUri                   string                 `binding:"omitempty" bson:",omitempty" json:"valueUri,omitempty"`
	ValueUrl                   *url.URL               `binding:"omitempty" bson:",omitempty" json:"valueUrl,omitempty"`
	ValueUuid                  *uuid.UUID             `binding:"omitempty" bson:",omitempty" json:"valueUuid,omitempty"`
	ValueAddress               *Address               `binding:"omitempty" bson:",omitempty" json:"valueAddress,omitempty"`
	ValueAge                   *Age                   `binding:"omitempty" bson:",omitempty" json:"valueAge,omitempty"`
	ValueAnnotation            *Annotation            `binding:"omitempty" bson:",omitempty" json:"valueAnnotation,omitempty"`
	ValueAttachment            *Attachment            `binding:"omitempty" bson:",omitempty" json:"valueAttachment,omitempty"`
	ValueCodeableConcept       *CodeableConcept       `binding:"omitempty" bson:",omitempty" json:"valueCodeableConcept,omitempty"`
	ValueCodeableReference     *CodeableReference     `binding:"omitempty" bson:",omitempty" json:"valueCodeableReference,omitempty"`
	ValueCoding                *Coding                `binding:"omitempty" bson:",omitempty" json:"valueCoding,omitempty"`
	ValueContactPoint          *ContactPoint          `binding:"omitempty" bson:",omitempty" json:"valueContactPoint,omitempty"`
	ValueCount                 *Count                 `binding:"omitempty" bson:",omitempty" json:"valueCount,omitempty"`
	ValueDistance              *Distance              `binding:"omitempty" bson:",omitempty" json:"valueDistance,omitempty"`
	ValueDuration              *Duration              `binding:"omitempty" bson:",omitempty" json:"valueDuration,omitempty"`
	ValueHumanName             *HumanName             `binding:"omitempty" bson:",omitempty" json:"valueHumanName,omitempty"`
	ValueIdentifier            *Identifier            `binding:"omitempty" bson:",omitempty" json:"valueIdentifier,omitempty"`
	ValueMoney                 *Money                 `binding:"omitempty" bson:",omitempty" json:"valueMoney,omitempty"`
	ValuePeriod                *Period                `binding:"omitempty" bson:",omitempty" json:"valuePeriod,omitempty"`
	ValueQuantity              *Quantity              `binding:"omitempty" bson:",omitempty" json:"valueQuantity,omitempty"`
	ValueRange                 *Range                 `binding:"omitempty" bson:",omitempty" json:"valueRange,omitempty"`
	ValueRatio                 *Ratio                 `binding:"omitempty" bson:",omitempty" json:"valueRatio,omitempty"`
	ValueRatioRange            *RatioRange            `binding:"omitempty" bson:",omitempty" json:"valueRatioRange,omitempty"`
	ValueReference             *Reference             `binding:"omitempty" bson:",omitempty" json:"valueReference,omitempty"`
	ValueSampledData           *SampledData           `binding:"omitempty" bson:",omitempty" json:"valueSampledData,omitempty"`
	ValueSignature             *Signature             `binding:"omitempty" bson:",omitempty" json:"valueSignature,omitempty"`
	ValueTiming                *Timing                `binding:"omitempty" bson:",omitempty" json:"valueTiming,omitempty"`
	ValueContactDetail         *ContactDetail         `binding:"omitempty" bson:",omitempty" json:"valueContactDetail,omitempty"`
	ValueDataRequirement       *DataRequirement       `binding:"omitempty" bson:",omitempty" json:"valueDataRequirement,omitempty"`
	ValueExpression            *Expression            `binding:"omitempty" bson:",omitempty" json:"valueExpression,omitempty"`
	ValueParameterDefinition   *ParameterDefinition   `binding:"omitempty" bson:",omitempty" json:"valueParameterDefinition,omitempty"`
	ValueRelatedArtifact       *RelatedArtifact       `binding:"omitempty" bson:",omitempty" json:"valueRelatedArtifact,omitempty"`
	ValueTriggerDefinition     *TriggerDefinition     `binding:"omitempty" bson:",omitempty" json:"valueTriggerDefinition,omitempty"`
	ValueUsageContext          *UsageContext          `binding:"omitempty" bson:",omitempty" json:"valueUsageContext,omitempty"`
	ValueAvailability          *Availability          `binding:"omitempty" bson:",omitempty" json:"valueAvailability,omitempty"`
	ValueExtendedContactDetail *ExtendedContactDetail `binding:"omitempty" bson:",omitempty" json:"valueExtendedContactDetail,omitempty"`
	ValueDosage                *Dosage                `binding:"omitempty" bson:",omitempty" json:"valueDosage,omitempty"`
	ValueMeta                  *Meta                  `binding:"omitempty" bson:",omitempty" json:"valueMeta,omitempty"`
}

func (out *Extension) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Extension\"" {
		return fmt.Errorf("resourceType is not %s", "Extension")
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["url"], &out.Url); err != nil {
		return err
	}
	if len(asMap["_url"]) > 0 {
		if err := go1.Unmarshal(asMap["_url"], &out.UrlPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["valueBase64Binary"]) > 0 {
		if err := go1.Unmarshal(asMap["valueBase64Binary"], &out.ValueBase64Binary); err == nil {
		}
	} else if len(asMap["valueBoolean"]) > 0 {
		if err := go1.Unmarshal(asMap["valueBoolean"], &out.ValueBoolean); err == nil {
		}
	} else if len(asMap["valueCanonical"]) > 0 {
		if err := go1.Unmarshal(asMap["valueCanonical"], &out.ValueCanonical); err == nil {
		}
	} else if len(asMap["valueCode"]) > 0 {
		if err := go1.Unmarshal(asMap["valueCode"], &out.ValueCode); err == nil {
		}
	} else if len(asMap["valueDate"]) > 0 {
		if err := go1.Unmarshal(asMap["valueDate"], &out.ValueDate); err == nil {
		}
	} else if len(asMap["valueDateTime"]) > 0 {
		if err := go1.Unmarshal(asMap["valueDateTime"], &out.ValueDateTime); err == nil {
		}
	} else if len(asMap["valueDecimal"]) > 0 {
		if err := go1.Unmarshal(asMap["valueDecimal"], &out.ValueDecimal); err == nil {
		}
	} else if len(asMap["valueId"]) > 0 {
		if err := go1.Unmarshal(asMap["valueId"], &out.ValueId); err == nil {
		}
	} else if len(asMap["valueInstant"]) > 0 {
		if err := go1.Unmarshal(asMap["valueInstant"], &out.ValueInstant); err == nil {
		}
	} else if len(asMap["valueInteger"]) > 0 {
		if err := go1.Unmarshal(asMap["valueInteger"], &out.ValueInteger); err == nil {
		}
	} else if len(asMap["valueInteger64"]) > 0 {
		if err := go1.Unmarshal(asMap["valueInteger64"], &out.ValueInteger64); err == nil {
		}
	} else if len(asMap["valueMarkdown"]) > 0 {
		if err := go1.Unmarshal(asMap["valueMarkdown"], &out.ValueMarkdown); err == nil {
		}
	} else if len(asMap["valueOid"]) > 0 {
		if err := go1.Unmarshal(asMap["valueOid"], &out.ValueOid); err == nil {
		}
	} else if len(asMap["valuePositiveInt"]) > 0 {
		if err := go1.Unmarshal(asMap["valuePositiveInt"], &out.ValuePositiveInt); err == nil {
		}
	} else if len(asMap["valueString"]) > 0 {
		if err := go1.Unmarshal(asMap["valueString"], &out.ValueString); err == nil {
		}
	} else if len(asMap["valueTime"]) > 0 {
		if err := go1.Unmarshal(asMap["valueTime"], &out.ValueTime); err == nil {
		}
	} else if len(asMap["valueUnsignedInt"]) > 0 {
		if err := go1.Unmarshal(asMap["valueUnsignedInt"], &out.ValueUnsignedInt); err == nil {
		}
	} else if len(asMap["valueUri"]) > 0 {
		if err := go1.Unmarshal(asMap["valueUri"], &out.ValueUri); err == nil {
		}
	} else if len(asMap["valueUrl"]) > 0 {
		if err := go1.Unmarshal(asMap["valueUrl"], &out.ValueUrl); err == nil {
		}
	} else if len(asMap["valueUuid"]) > 0 {
		if err := go1.Unmarshal(asMap["valueUuid"], &out.ValueUuid); err == nil {
		}
	} else if len(asMap["valueAddress"]) > 0 {
		if err := go1.Unmarshal(asMap["valueAddress"], &out.ValueAddress); err == nil {
		}
	} else if len(asMap["valueAge"]) > 0 {
		if err := go1.Unmarshal(asMap["valueAge"], &out.ValueAge); err == nil {
		}
	} else if len(asMap["valueAnnotation"]) > 0 {
		if err := go1.Unmarshal(asMap["valueAnnotation"], &out.ValueAnnotation); err == nil {
		}
	} else if len(asMap["valueAttachment"]) > 0 {
		if err := go1.Unmarshal(asMap["valueAttachment"], &out.ValueAttachment); err == nil {
		}
	} else if len(asMap["valueCodeableConcept"]) > 0 {
		if err := go1.Unmarshal(asMap["valueCodeableConcept"], &out.ValueCodeableConcept); err == nil {
		}
	} else if len(asMap["valueCodeableReference"]) > 0 {
		if err := go1.Unmarshal(asMap["valueCodeableReference"], &out.ValueCodeableReference); err == nil {
		}
	} else if len(asMap["valueCoding"]) > 0 {
		if err := go1.Unmarshal(asMap["valueCoding"], &out.ValueCoding); err == nil {
		}
	} else if len(asMap["valueContactPoint"]) > 0 {
		if err := go1.Unmarshal(asMap["valueContactPoint"], &out.ValueContactPoint); err == nil {
		}
	} else if len(asMap["valueCount"]) > 0 {
		if err := go1.Unmarshal(asMap["valueCount"], &out.ValueCount); err == nil {
		}
	} else if len(asMap["valueDistance"]) > 0 {
		if err := go1.Unmarshal(asMap["valueDistance"], &out.ValueDistance); err == nil {
		}
	} else if len(asMap["valueDuration"]) > 0 {
		if err := go1.Unmarshal(asMap["valueDuration"], &out.ValueDuration); err == nil {
		}
	} else if len(asMap["valueHumanName"]) > 0 {
		if err := go1.Unmarshal(asMap["valueHumanName"], &out.ValueHumanName); err == nil {
		}
	} else if len(asMap["valueIdentifier"]) > 0 {
		if err := go1.Unmarshal(asMap["valueIdentifier"], &out.ValueIdentifier); err == nil {
		}
	} else if len(asMap["valueMoney"]) > 0 {
		if err := go1.Unmarshal(asMap["valueMoney"], &out.ValueMoney); err == nil {
		}
	} else if len(asMap["valuePeriod"]) > 0 {
		if err := go1.Unmarshal(asMap["valuePeriod"], &out.ValuePeriod); err == nil {
		}
	} else if len(asMap["valueQuantity"]) > 0 {
		if err := go1.Unmarshal(asMap["valueQuantity"], &out.ValueQuantity); err == nil {
		}
	} else if len(asMap["valueRange"]) > 0 {
		if err := go1.Unmarshal(asMap["valueRange"], &out.ValueRange); err == nil {
		}
	} else if len(asMap["valueRatio"]) > 0 {
		if err := go1.Unmarshal(asMap["valueRatio"], &out.ValueRatio); err == nil {
		}
	} else if len(asMap["valueRatioRange"]) > 0 {
		if err := go1.Unmarshal(asMap["valueRatioRange"], &out.ValueRatioRange); err == nil {
		}
	} else if len(asMap["valueReference"]) > 0 {
		if err := go1.Unmarshal(asMap["valueReference"], &out.ValueReference); err == nil {
		}
	} else if len(asMap["valueSampledData"]) > 0 {
		if err := go1.Unmarshal(asMap["valueSampledData"], &out.ValueSampledData); err == nil {
		}
	} else if len(asMap["valueSignature"]) > 0 {
		if err := go1.Unmarshal(asMap["valueSignature"], &out.ValueSignature); err == nil {
		}
	} else if len(asMap["valueTiming"]) > 0 {
		if err := go1.Unmarshal(asMap["valueTiming"], &out.ValueTiming); err == nil {
		}
	} else if len(asMap["valueContactDetail"]) > 0 {
		if err := go1.Unmarshal(asMap["valueContactDetail"], &out.ValueContactDetail); err == nil {
		}
	} else if len(asMap["valueDataRequirement"]) > 0 {
		if err := go1.Unmarshal(asMap["valueDataRequirement"], &out.ValueDataRequirement); err == nil {
		}
	} else if len(asMap["valueExpression"]) > 0 {
		if err := go1.Unmarshal(asMap["valueExpression"], &out.ValueExpression); err == nil {
		}
	} else if len(asMap["valueParameterDefinition"]) > 0 {
		if err := go1.Unmarshal(asMap["valueParameterDefinition"], &out.ValueParameterDefinition); err == nil {
		}
	} else if len(asMap["valueRelatedArtifact"]) > 0 {
		if err := go1.Unmarshal(asMap["valueRelatedArtifact"], &out.ValueRelatedArtifact); err == nil {
		}
	} else if len(asMap["valueTriggerDefinition"]) > 0 {
		if err := go1.Unmarshal(asMap["valueTriggerDefinition"], &out.ValueTriggerDefinition); err == nil {
		}
	} else if len(asMap["valueUsageContext"]) > 0 {
		if err := go1.Unmarshal(asMap["valueUsageContext"], &out.ValueUsageContext); err == nil {
		}
	} else if len(asMap["valueAvailability"]) > 0 {
		if err := go1.Unmarshal(asMap["valueAvailability"], &out.ValueAvailability); err == nil {
		}
	} else if len(asMap["valueExtendedContactDetail"]) > 0 {
		if err := go1.Unmarshal(asMap["valueExtendedContactDetail"], &out.ValueExtendedContactDetail); err == nil {
		}
	} else if len(asMap["valueDosage"]) > 0 {
		if err := go1.Unmarshal(asMap["valueDosage"], &out.ValueDosage); err == nil {
		}
	} else if len(asMap["valueMeta"]) > 0 {
		if err := go1.Unmarshal(asMap["valueMeta"], &out.ValueMeta); err == nil {
		}
	} else {

	}
	return nil
}

type Extension struct {
	Id                    *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension  *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension             []*Extension        `bson:",omitempty" json:"extension,omitempty"`              // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Url                   string              `binding:"required" bson:",omitempty" json:"url,omitempty"` // Source of the definition for the extension code - a logical name or a URL.
	UrlPrimitiveExtension *PrimitiveExtension `bson:"url_extension,omitempty" json:"_url,omitempty"`
	ExtensionValuex       `binding:"omitempty" bson:",omitempty" json:",omitempty"`
	ResourceType          string `bson:"-" json:"resourceType,omitempty"`
}

func (out *HumanName) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"HumanName\"" {
		return fmt.Errorf("resourceType is not %s", "HumanName")
	}
	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["use"]) > 0 {
		if err := go1.Unmarshal(asMap["use"], &out.Use); err != nil {
			return err
		}

	}
	if len(asMap["_use"]) > 0 {
		if err := go1.Unmarshal(asMap["_use"], &out.UsePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["text"]) > 0 {
		if err := go1.Unmarshal(asMap["text"], &out.Text); err != nil {
			return err
		}

	}
	if len(asMap["_text"]) > 0 {
		if err := go1.Unmarshal(asMap["_text"], &out.TextPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["family"]) > 0 {
		if err := go1.Unmarshal(asMap["family"], &out.Family); err != nil {
			return err
		}

	}
	if len(asMap["_family"]) > 0 {
		if err := go1.Unmarshal(asMap["_family"], &out.FamilyPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["period"]) > 0 {
		if err := go1.Unmarshal(asMap["period"], &out.Period); err != nil {
			return err
		}

	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["given"]) > 0 {
		if err := go1.Unmarshal(asMap["given"], &out.Given); err != nil {
			return err
		}

	}
	if len(asMap["_given"]) > 0 {
		if err := go1.Unmarshal(asMap["_given"], &out.GivenPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["prefix"]) > 0 {
		if err := go1.Unmarshal(asMap["prefix"], &out.Prefix); err != nil {
			return err
		}

	}
	if len(asMap["_prefix"]) > 0 {
		if err := go1.Unmarshal(asMap["_prefix"], &out.PrefixPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["suffix"]) > 0 {
		if err := go1.Unmarshal(asMap["suffix"], &out.Suffix); err != nil {
			return err
		}

	}
	if len(asMap["_suffix"]) > 0 {
		if err := go1.Unmarshal(asMap["_suffix"], &out.SuffixPrimitiveExtension); err != nil {
			return err
		}
	}

	return nil
}

type HumanName struct {
	Extension                []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Use                      string              `bson:",omitempty" json:"use,omitempty"`       // Identifies the purpose for this name.
	UsePrimitiveExtension    *PrimitiveExtension `bson:"use_extension,omitempty" json:"_use,omitempty"`
	Text                     string              `bson:",omitempty" json:"text,omitempty"` // Specifies the entire name as it should be displayed e.g. on an application UI. This may be provided instead of or as well as the specific parts.
	TextPrimitiveExtension   *PrimitiveExtension `bson:"text_extension,omitempty" json:"_text,omitempty"`
	Family                   string              `bson:",omitempty" json:"family,omitempty"` // The part of a name that links to the genealogy. In some cultures (e.g. Eritrea) the family name of a son is the first name of his father.
	FamilyPrimitiveExtension *PrimitiveExtension `bson:"family_extension,omitempty" json:"_family,omitempty"`
	Period                   *Period             `bson:",omitempty" json:"period,omitempty"` // Indicates the period of time when this name was valid for the named person.
	Id                       *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension     *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Given                    []string            `bson:",omitempty" json:"given,omitempty"` // Given name.
	GivenPrimitiveExtension  *PrimitiveExtension `bson:"given_extension,omitempty" json:"_given,omitempty"`
	Prefix                   []string            `bson:",omitempty" json:"prefix,omitempty"` // Part of the name that is acquired as a title due to academic, legal, employment or nobility status, etc. and that appears at the start of the name.
	PrefixPrimitiveExtension *PrimitiveExtension `bson:"prefix_extension,omitempty" json:"_prefix,omitempty"`
	Suffix                   []string            `bson:",omitempty" json:"suffix,omitempty"` // Part of the name that is acquired as a title due to academic, legal, employment or nobility status, etc. and that appears at the end of the name.
	SuffixPrimitiveExtension *PrimitiveExtension `bson:"suffix_extension,omitempty" json:"_suffix,omitempty"`
	ResourceType             string              `bson:"-" json:"resourceType,omitempty"`
}

func (out *Identifier) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Identifier\"" {
		return fmt.Errorf("resourceType is not %s", "Identifier")
	}
	if len(asMap["value"]) > 0 {
		if err := go1.Unmarshal(asMap["value"], &out.Value); err != nil {
			return err
		}

	}
	if len(asMap["_value"]) > 0 {
		if err := go1.Unmarshal(asMap["_value"], &out.ValuePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["period"]) > 0 {
		if err := go1.Unmarshal(asMap["period"], &out.Period); err != nil {
			return err
		}

	}
	if len(asMap["assigner"]) > 0 {
		if err := go1.Unmarshal(asMap["assigner"], &out.Assigner); err != nil {
			return err
		}

	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["use"]) > 0 {
		if err := go1.Unmarshal(asMap["use"], &out.Use); err != nil {
			return err
		}

	}
	if len(asMap["_use"]) > 0 {
		if err := go1.Unmarshal(asMap["_use"], &out.UsePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["type"]) > 0 {
		if err := go1.Unmarshal(asMap["type"], &out.Type); err != nil {
			return err
		}

	}
	if len(asMap["system"]) > 0 {
		if err := go1.Unmarshal(asMap["system"], &out.System); err != nil {
			return err
		}

	}
	if len(asMap["_system"]) > 0 {
		if err := go1.Unmarshal(asMap["_system"], &out.SystemPrimitiveExtension); err != nil {
			return err
		}
	}

	return nil
}

type Identifier struct {
	Value                    string              `bson:",omitempty" json:"value,omitempty"` // The portion of the identifier typically relevant to the user and which is unique within the context of the system.
	ValuePrimitiveExtension  *PrimitiveExtension `bson:"value_extension,omitempty" json:"_value,omitempty"`
	Period                   *Period             `bson:",omitempty" json:"period,omitempty"`   // Time period during which identifier is/was valid for use.
	Assigner                 *Reference          `bson:",omitempty" json:"assigner,omitempty"` // Organization that issued/manages the identifier.
	Id                       *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`    // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension     *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension                []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Use                      string              `bson:",omitempty" json:"use,omitempty"`       // The purpose of this identifier.
	UsePrimitiveExtension    *PrimitiveExtension `bson:"use_extension,omitempty" json:"_use,omitempty"`
	Type                     *CodeableConcept    `bson:",omitempty" json:"type,omitempty"`   // A coded type for the identifier that can be used to determine which identifier to use for a specific purpose.
	System                   string              `bson:",omitempty" json:"system,omitempty"` // Establishes the namespace for the value - that is, an absolute URL that describes a set values that are unique.
	SystemPrimitiveExtension *PrimitiveExtension `bson:"system_extension,omitempty" json:"_system,omitempty"`
	ResourceType             string              `bson:"-" json:"resourceType,omitempty"`
}

func (out *MarketingStatus) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"MarketingStatus\"" {
		return fmt.Errorf("resourceType is not %s", "MarketingStatus")
	}
	if len(asMap["modifierExtension"]) > 0 {
		if err := go1.Unmarshal(asMap["modifierExtension"], &out.ModifierExtension); err != nil {
			return err
		}

	}
	if len(asMap["country"]) > 0 {
		if err := go1.Unmarshal(asMap["country"], &out.Country); err != nil {
			return err
		}

	}
	if len(asMap["jurisdiction"]) > 0 {
		if err := go1.Unmarshal(asMap["jurisdiction"], &out.Jurisdiction); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["status"], &out.Status); err != nil {
		return err
	}

	if len(asMap["dateRange"]) > 0 {
		if err := go1.Unmarshal(asMap["dateRange"], &out.DateRange); err != nil {
			return err
		}

	}
	if len(asMap["restoreDate"]) > 0 {
		if err := go1.Unmarshal(asMap["restoreDate"], &out.RestoreDate); err != nil {
			return err
		}

	}
	if len(asMap["_restoreDate"]) > 0 {
		if err := go1.Unmarshal(asMap["_restoreDate"], &out.RestoreDatePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	return nil
}

type MarketingStatus struct {
	ModifierExtension []*Extension `bson:",omitempty" json:"modifierExtension,omitempty"` /*
	May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.

	Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	*/
	Country                       *CodeableConcept    `bson:",omitempty" json:"country,omitempty"`                   // The country in which the marketing authorization has been granted shall be specified It should be specified using the ISO 3166 ‑ 1 alpha-2 code elements.
	Jurisdiction                  *CodeableConcept    `bson:",omitempty" json:"jurisdiction,omitempty"`              // Where a Medicines Regulatory Agency has granted a marketing authorization for which specific provisions within a jurisdiction apply, the jurisdiction can be specified using an appropriate controlled terminology The controlled term and the controlled term identifier shall be specified.
	Status                        *CodeableConcept    `binding:"required" bson:",omitempty" json:"status,omitempty"` // This attribute provides information on the status of the marketing of the medicinal product See ISO/TS 20443 for more information and examples.
	DateRange                     *Period             `bson:",omitempty" json:"dateRange,omitempty"`                 // The date when the Medicinal Product is placed on the market by the Marketing Authorization Holder (or where applicable, the manufacturer/distributor) in a country and/or jurisdiction shall be provided A complete date consisting of day, month and year shall be specified using the ISO 8601 date format NOTE “Placed on the market” refers to the release of the Medicinal Product into the distribution chain.
	RestoreDate                   *DateTime           `bson:",omitempty" json:"restoreDate,omitempty"`               // The date when the Medicinal Product is placed on the market by the Marketing Authorization Holder (or where applicable, the manufacturer/distributor) in a country and/or jurisdiction shall be provided A complete date consisting of day, month and year shall be specified using the ISO 8601 date format NOTE “Placed on the market” refers to the release of the Medicinal Product into the distribution chain.
	RestoreDatePrimitiveExtension *PrimitiveExtension `bson:"restoreDate_extension,omitempty" json:"_restoreDate,omitempty"`
	Id                            *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension          *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension                     []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	ResourceType                  string              `bson:"-" json:"resourceType,omitempty"`
}

func (out *Meta) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Meta\"" {
		return fmt.Errorf("resourceType is not %s", "Meta")
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["versionId"]) > 0 {
		if err := go1.Unmarshal(asMap["versionId"], &out.VersionId); err != nil {
			return err
		}

	}
	if len(asMap["_versionId"]) > 0 {
		if err := go1.Unmarshal(asMap["_versionId"], &out.VersionIdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["lastUpdated"]) > 0 {
		if err := go1.Unmarshal(asMap["lastUpdated"], &out.LastUpdated); err != nil {
			return err
		}

	}
	if len(asMap["_lastUpdated"]) > 0 {
		if err := go1.Unmarshal(asMap["_lastUpdated"], &out.LastUpdatedPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["source"]) > 0 {
		if err := go1.Unmarshal(asMap["source"], &out.Source); err != nil {
			return err
		}

	}
	if len(asMap["_source"]) > 0 {
		if err := go1.Unmarshal(asMap["_source"], &out.SourcePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["profile"]) > 0 {
		if err := go1.Unmarshal(asMap["profile"], &out.Profile); err != nil {
			return err
		}

	}
	if len(asMap["_profile"]) > 0 {
		if err := go1.Unmarshal(asMap["_profile"], &out.ProfilePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["security"]) > 0 {
		if err := go1.Unmarshal(asMap["security"], &out.Security); err != nil {
			return err
		}

	}
	if len(asMap["tag"]) > 0 {
		if err := go1.Unmarshal(asMap["tag"], &out.Tag); err != nil {
			return err
		}

	}
	return nil
}

type Meta struct {
	Id                            *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension          *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension                     []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	VersionId                     int                 `bson:",omitempty" json:"versionId,omitempty"` // The version specific identifier, as it appears in the version portion of the URL. This value changes when the resource is created, updated, or deleted.
	VersionIdPrimitiveExtension   *PrimitiveExtension `bson:"versionId_extension,omitempty" json:"_versionId,omitempty"`
	LastUpdated                   *time.Time          `bson:",omitempty" json:"lastUpdated,omitempty"` // When the resource last changed - e.g. when the version changed.
	LastUpdatedPrimitiveExtension *PrimitiveExtension `bson:"lastUpdated_extension,omitempty" json:"_lastUpdated,omitempty"`
	Source                        string              `bson:",omitempty" json:"source,omitempty"` // A uri that identifies the source system of the resource. This provides a minimal amount of [Provenance](provenance.html#) information that can be used to track or differentiate the source of information in the resource. The source may identify another FHIR server, document, message, database, etc.
	SourcePrimitiveExtension      *PrimitiveExtension `bson:"source_extension,omitempty" json:"_source,omitempty"`
	Profile                       []string            `bson:",omitempty" json:"profile,omitempty"` // A list of profiles (references to [StructureDefinition](structuredefinition.html#) resources) that this resource claims to conform to. The URL is a reference to [StructureDefinition.url](structuredefinition-definitions.html#StructureDefinition.url).
	ProfilePrimitiveExtension     *PrimitiveExtension `bson:"profile_extension,omitempty" json:"_profile,omitempty"`
	Security                      []*Coding           `bson:",omitempty" json:"security,omitempty"` // Security labels applied to this resource. These tags connect specific resources to the overall security policy and infrastructure.
	Tag                           []*Coding           `bson:",omitempty" json:"tag,omitempty"`      // Tags applied to this resource. Tags are intended to be used to identify and relate resources to process and workflow, and applications are not required to consider the tags when interpreting the meaning of a resource.
	ResourceType                  string              `bson:"-" json:"resourceType,omitempty"`
}

func (out *MonetaryComponent) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"MonetaryComponent\"" {
		return fmt.Errorf("resourceType is not %s", "MonetaryComponent")
	}
	if len(asMap["factor"]) > 0 {
		if err := go1.Unmarshal(asMap["factor"], &out.Factor); err != nil {
			return err
		}

	}
	if len(asMap["_factor"]) > 0 {
		if err := go1.Unmarshal(asMap["_factor"], &out.FactorPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["amount"]) > 0 {
		if err := go1.Unmarshal(asMap["amount"], &out.Amount); err != nil {
			return err
		}

	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["type"], &out.Type); err != nil {
		return err
	}
	if len(asMap["_type"]) > 0 {
		if err := go1.Unmarshal(asMap["_type"], &out.TypePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["code"]) > 0 {
		if err := go1.Unmarshal(asMap["code"], &out.Code); err != nil {
			return err
		}

	}
	return nil
}

type MonetaryComponent struct {
	Factor                   float64             `bson:",omitempty" json:"factor,omitempty"` // Factor used for calculating this component.
	FactorPrimitiveExtension *PrimitiveExtension `bson:"factor_extension,omitempty" json:"_factor,omitempty"`
	Amount                   *Money              `bson:",omitempty" json:"amount,omitempty"` // Explicit value amount to be used.
	Id                       *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension     *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension                []*Extension        `bson:",omitempty" json:"extension,omitempty"`               // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Type                     string              `binding:"required" bson:",omitempty" json:"type,omitempty"` // base | surcharge | deduction | discount | tax | informational.
	TypePrimitiveExtension   *PrimitiveExtension `bson:"type_extension,omitempty" json:"_type,omitempty"`
	Code                     *CodeableConcept    `bson:",omitempty" json:"code,omitempty"` // Codes may be used to differentiate between kinds of taxes, surcharges, discounts etc.
	ResourceType             string              `bson:"-" json:"resourceType,omitempty"`
}

func (out *Money) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Money\"" {
		return fmt.Errorf("resourceType is not %s", "Money")
	}
	if len(asMap["value"]) > 0 {
		if err := go1.Unmarshal(asMap["value"], &out.Value); err != nil {
			return err
		}

	}
	if len(asMap["_value"]) > 0 {
		if err := go1.Unmarshal(asMap["_value"], &out.ValuePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["currency"]) > 0 {
		if err := go1.Unmarshal(asMap["currency"], &out.Currency); err != nil {
			return err
		}

	}
	if len(asMap["_currency"]) > 0 {
		if err := go1.Unmarshal(asMap["_currency"], &out.CurrencyPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	return nil
}

type Money struct {
	Value                      float64             `bson:",omitempty" json:"value,omitempty"` // Numerical value (with implicit precision).
	ValuePrimitiveExtension    *PrimitiveExtension `bson:"value_extension,omitempty" json:"_value,omitempty"`
	Currency                   string              `bson:",omitempty" json:"currency,omitempty"` // ISO 4217 Currency Code.
	CurrencyPrimitiveExtension *PrimitiveExtension `bson:"currency_extension,omitempty" json:"_currency,omitempty"`
	Id                         *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension       *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension                  []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	ResourceType               string              `bson:"-" json:"resourceType,omitempty"`
}

func (out *Narrative) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Narrative\"" {
		return fmt.Errorf("resourceType is not %s", "Narrative")
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["status"], &out.Status); err != nil {
		return err
	}
	if len(asMap["_status"]) > 0 {
		if err := go1.Unmarshal(asMap["_status"], &out.StatusPrimitiveExtension); err != nil {
			return err
		}
	}

	if err := go1.Unmarshal(asMap["div"], &out.Div); err != nil {
		return err
	}

	return nil
}

type Narrative struct {
	Id                       *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension     *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension                []*Extension        `bson:",omitempty" json:"extension,omitempty"`                 // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Status                   string              `binding:"required" bson:",omitempty" json:"status,omitempty"` // The status of the narrative - whether it's entirely generated (from just the defined data or the extensions too), or whether a human authored it and it may contain additional data.
	StatusPrimitiveExtension *PrimitiveExtension `bson:"status_extension,omitempty" json:"_status,omitempty"`
	Div                      string              `binding:"required" bson:",omitempty" json:"div,omitempty"` // The actual narrative content, a stripped down version of XHTML.
	ResourceType             string              `bson:"-" json:"resourceType,omitempty"`
}

func (out *ParameterDefinition) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"ParameterDefinition\"" {
		return fmt.Errorf("resourceType is not %s", "ParameterDefinition")
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["min"]) > 0 {
		if err := go1.Unmarshal(asMap["min"], &out.Min); err != nil {
			return err
		}

	}
	if len(asMap["_min"]) > 0 {
		if err := go1.Unmarshal(asMap["_min"], &out.MinPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["max"]) > 0 {
		if err := go1.Unmarshal(asMap["max"], &out.Max); err != nil {
			return err
		}

	}
	if len(asMap["_max"]) > 0 {
		if err := go1.Unmarshal(asMap["_max"], &out.MaxPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["documentation"]) > 0 {
		if err := go1.Unmarshal(asMap["documentation"], &out.Documentation); err != nil {
			return err
		}

	}
	if len(asMap["_documentation"]) > 0 {
		if err := go1.Unmarshal(asMap["_documentation"], &out.DocumentationPrimitiveExtension); err != nil {
			return err
		}
	}

	if err := go1.Unmarshal(asMap["type"], &out.Type); err != nil {
		return err
	}
	if len(asMap["_type"]) > 0 {
		if err := go1.Unmarshal(asMap["_type"], &out.TypePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["name"]) > 0 {
		if err := go1.Unmarshal(asMap["name"], &out.Name); err != nil {
			return err
		}

	}
	if len(asMap["_name"]) > 0 {
		if err := go1.Unmarshal(asMap["_name"], &out.NamePrimitiveExtension); err != nil {
			return err
		}
	}

	if err := go1.Unmarshal(asMap["use"], &out.Use); err != nil {
		return err
	}
	if len(asMap["_use"]) > 0 {
		if err := go1.Unmarshal(asMap["_use"], &out.UsePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["profile"]) > 0 {
		if err := go1.Unmarshal(asMap["profile"], &out.Profile); err != nil {
			return err
		}

	}
	if len(asMap["_profile"]) > 0 {
		if err := go1.Unmarshal(asMap["_profile"], &out.ProfilePrimitiveExtension); err != nil {
			return err
		}
	}

	return nil
}

type ParameterDefinition struct {
	Id                              *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension            *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension                       []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Min                             int                 `bson:",omitempty" json:"min,omitempty"`       // The minimum number of times this parameter SHALL appear in the request or response.
	MinPrimitiveExtension           *PrimitiveExtension `bson:"min_extension,omitempty" json:"_min,omitempty"`
	Max                             string              `bson:",omitempty" json:"max,omitempty"` // The maximum number of times this element is permitted to appear in the request or response.
	MaxPrimitiveExtension           *PrimitiveExtension `bson:"max_extension,omitempty" json:"_max,omitempty"`
	Documentation                   string              `bson:",omitempty" json:"documentation,omitempty"` // A brief discussion of what the parameter is for and how it is used by the module.
	DocumentationPrimitiveExtension *PrimitiveExtension `bson:"documentation_extension,omitempty" json:"_documentation,omitempty"`
	Type                            string              `binding:"required" bson:",omitempty" json:"type,omitempty"` // The type of the parameter.
	TypePrimitiveExtension          *PrimitiveExtension `bson:"type_extension,omitempty" json:"_type,omitempty"`
	Name                            string              `bson:",omitempty" json:"name,omitempty"` // The name of the parameter used to allow access to the value of the parameter in evaluation contexts.
	NamePrimitiveExtension          *PrimitiveExtension `bson:"name_extension,omitempty" json:"_name,omitempty"`
	Use                             string              `binding:"required" bson:",omitempty" json:"use,omitempty"` // Whether the parameter is input or output for the module.
	UsePrimitiveExtension           *PrimitiveExtension `bson:"use_extension,omitempty" json:"_use,omitempty"`
	Profile                         string              `bson:",omitempty" json:"profile,omitempty"` // If specified, this indicates a profile that the input data must conform to, or that the output data will conform to.
	ProfilePrimitiveExtension       *PrimitiveExtension `bson:"profile_extension,omitempty" json:"_profile,omitempty"`
	ResourceType                    string              `bson:"-" json:"resourceType,omitempty"`
}

func (out *Period) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Period\"" {
		return fmt.Errorf("resourceType is not %s", "Period")
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["start"]) > 0 {
		if err := go1.Unmarshal(asMap["start"], &out.Start); err != nil {
			return err
		}

	}
	if len(asMap["_start"]) > 0 {
		if err := go1.Unmarshal(asMap["_start"], &out.StartPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["end"]) > 0 {
		if err := go1.Unmarshal(asMap["end"], &out.End); err != nil {
			return err
		}

	}
	if len(asMap["_end"]) > 0 {
		if err := go1.Unmarshal(asMap["_end"], &out.EndPrimitiveExtension); err != nil {
			return err
		}
	}

	return nil
}

type Period struct {
	Id                      *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension    *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension               []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Start                   *DateTime           `bson:",omitempty" json:"start,omitempty"`     // The start of the period. The boundary is inclusive.
	StartPrimitiveExtension *PrimitiveExtension `bson:"start_extension,omitempty" json:"_start,omitempty"`
	End                     *DateTime           `bson:",omitempty" json:"end,omitempty"` // The end of the period. If the end of the period is missing, it means no end was known or planned at the time the instance was created. The start may be in the past, and the end date in the future, which means that period is expected/planned to end at that time.
	EndPrimitiveExtension   *PrimitiveExtension `bson:"end_extension,omitempty" json:"_end,omitempty"`
	ResourceType            string              `bson:"-" json:"resourceType,omitempty"`
}

func (out *PrimitiveType) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"PrimitiveType\"" {
		return fmt.Errorf("resourceType is not %s", "PrimitiveType")
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	return nil
}

type PrimitiveType struct {
	Id                   *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension            []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	ResourceType         string              `bson:"-" json:"resourceType,omitempty"`
}

type ProductShelfLifePeriodx struct {
	PeriodDuration *Duration `binding:"omitempty" bson:",omitempty" json:"periodDuration,omitempty"`
	PeriodString   string    `binding:"omitempty" bson:",omitempty" json:"periodString,omitempty"`
}

func (out *ProductShelfLife) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"ProductShelfLife\"" {
		return fmt.Errorf("resourceType is not %s", "ProductShelfLife")
	}
	if len(asMap["periodDuration"]) > 0 {
		if err := go1.Unmarshal(asMap["periodDuration"], &out.PeriodDuration); err == nil {
		}
	} else if len(asMap["periodString"]) > 0 {
		if err := go1.Unmarshal(asMap["periodString"], &out.PeriodString); err == nil {
		}
	} else {

	}
	if len(asMap["specialPrecautionsForStorage"]) > 0 {
		if err := go1.Unmarshal(asMap["specialPrecautionsForStorage"], &out.SpecialPrecautionsForStorage); err != nil {
			return err
		}

	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["modifierExtension"]) > 0 {
		if err := go1.Unmarshal(asMap["modifierExtension"], &out.ModifierExtension); err != nil {
			return err
		}

	}
	if len(asMap["type"]) > 0 {
		if err := go1.Unmarshal(asMap["type"], &out.Type); err != nil {
			return err
		}

	}
	return nil
}

type ProductShelfLife struct {
	ProductShelfLifePeriodx      `binding:"omitempty" bson:",omitempty" json:",omitempty"`
	SpecialPrecautionsForStorage []*CodeableConcept  `bson:",omitempty" json:"specialPrecautionsForStorage,omitempty"` // Special precautions for storage, if any, can be specified using an appropriate controlled vocabulary The controlled term and the controlled term identifier shall be specified.
	Id                           *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`                        // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension         *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension                    []*Extension        `bson:",omitempty" json:"extension,omitempty"`         // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	ModifierExtension            []*Extension        `bson:",omitempty" json:"modifierExtension,omitempty"` /*
	May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.

	Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	*/
	Type         *CodeableConcept `bson:",omitempty" json:"type,omitempty"` // This describes the shelf life, taking into account various scenarios such as shelf life of the packaged Medicinal Product itself, shelf life after transformation where necessary and shelf life after the first opening of a bottle, etc. The shelf life type shall be specified using an appropriate controlled vocabulary The controlled term and the controlled term identifier shall be specified.
	ResourceType string           `bson:"-" json:"resourceType,omitempty"`
}

func (out *Quantity) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Quantity\"" {
		return fmt.Errorf("resourceType is not %s", "Quantity")
	}
	if len(asMap["system"]) > 0 {
		if err := go1.Unmarshal(asMap["system"], &out.System); err != nil {
			return err
		}

	}
	if len(asMap["_system"]) > 0 {
		if err := go1.Unmarshal(asMap["_system"], &out.SystemPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["code"]) > 0 {
		if err := go1.Unmarshal(asMap["code"], &out.Code); err != nil {
			return err
		}

	}
	if len(asMap["_code"]) > 0 {
		if err := go1.Unmarshal(asMap["_code"], &out.CodePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["value"]) > 0 {
		if err := go1.Unmarshal(asMap["value"], &out.Value); err != nil {
			return err
		}

	}
	if len(asMap["_value"]) > 0 {
		if err := go1.Unmarshal(asMap["_value"], &out.ValuePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["comparator"]) > 0 {
		if err := go1.Unmarshal(asMap["comparator"], &out.Comparator); err != nil {
			return err
		}

	}
	if len(asMap["_comparator"]) > 0 {
		if err := go1.Unmarshal(asMap["_comparator"], &out.ComparatorPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["unit"]) > 0 {
		if err := go1.Unmarshal(asMap["unit"], &out.Unit); err != nil {
			return err
		}

	}
	if len(asMap["_unit"]) > 0 {
		if err := go1.Unmarshal(asMap["_unit"], &out.UnitPrimitiveExtension); err != nil {
			return err
		}
	}

	return nil
}

type Quantity struct {
	System                       string              `bson:",omitempty" json:"system,omitempty"` // The identification of the system that provides the coded form of the unit.
	SystemPrimitiveExtension     *PrimitiveExtension `bson:"system_extension,omitempty" json:"_system,omitempty"`
	Code                         string              `bson:",omitempty" json:"code,omitempty"` // A computer processable form of the unit in some unit representation system.
	CodePrimitiveExtension       *PrimitiveExtension `bson:"code_extension,omitempty" json:"_code,omitempty"`
	Id                           *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension         *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension                    []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Value                        float64             `bson:",omitempty" json:"value,omitempty"`     // The value of the measured amount. The value includes an implicit precision in the presentation of the value.
	ValuePrimitiveExtension      *PrimitiveExtension `bson:"value_extension,omitempty" json:"_value,omitempty"`
	Comparator                   string              `bson:",omitempty" json:"comparator,omitempty"` // How the value should be understood and represented - whether the actual value is greater or less than the stated value due to measurement issues; e.g. if the comparator is "<" , then the real value is < stated value.
	ComparatorPrimitiveExtension *PrimitiveExtension `bson:"comparator_extension,omitempty" json:"_comparator,omitempty"`
	Unit                         string              `bson:",omitempty" json:"unit,omitempty"` // A human-readable form of the unit.
	UnitPrimitiveExtension       *PrimitiveExtension `bson:"unit_extension,omitempty" json:"_unit,omitempty"`
	ResourceType                 string              `bson:"-" json:"resourceType,omitempty"`
}

func (out *Range) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Range\"" {
		return fmt.Errorf("resourceType is not %s", "Range")
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["low"]) > 0 {
		if err := go1.Unmarshal(asMap["low"], &out.Low); err != nil {
			return err
		}

	}
	if len(asMap["high"]) > 0 {
		if err := go1.Unmarshal(asMap["high"], &out.High); err != nil {
			return err
		}

	}
	return nil
}

type Range struct {
	Id                   *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension            []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Low                  *Quantity           `bson:",omitempty" json:"low,omitempty"`       // The low limit. The boundary is inclusive.
	High                 *Quantity           `bson:",omitempty" json:"high,omitempty"`      // The high limit. The boundary is inclusive.
	ResourceType         string              `bson:"-" json:"resourceType,omitempty"`
}

func (out *Ratio) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Ratio\"" {
		return fmt.Errorf("resourceType is not %s", "Ratio")
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["numerator"]) > 0 {
		if err := go1.Unmarshal(asMap["numerator"], &out.Numerator); err != nil {
			return err
		}

	}
	if len(asMap["denominator"]) > 0 {
		if err := go1.Unmarshal(asMap["denominator"], &out.Denominator); err != nil {
			return err
		}

	}
	return nil
}

type Ratio struct {
	Id                   *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension            []*Extension        `bson:",omitempty" json:"extension,omitempty"`   // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Numerator            *Quantity           `bson:",omitempty" json:"numerator,omitempty"`   // The value of the numerator.
	Denominator          *Quantity           `bson:",omitempty" json:"denominator,omitempty"` // The value of the denominator.
	ResourceType         string              `bson:"-" json:"resourceType,omitempty"`
}

func (out *RatioRange) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"RatioRange\"" {
		return fmt.Errorf("resourceType is not %s", "RatioRange")
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["lowNumerator"]) > 0 {
		if err := go1.Unmarshal(asMap["lowNumerator"], &out.LowNumerator); err != nil {
			return err
		}

	}
	if len(asMap["highNumerator"]) > 0 {
		if err := go1.Unmarshal(asMap["highNumerator"], &out.HighNumerator); err != nil {
			return err
		}

	}
	if len(asMap["denominator"]) > 0 {
		if err := go1.Unmarshal(asMap["denominator"], &out.Denominator); err != nil {
			return err
		}

	}
	return nil
}

type RatioRange struct {
	Id                   *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension            []*Extension        `bson:",omitempty" json:"extension,omitempty"`     // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	LowNumerator         *Quantity           `bson:",omitempty" json:"lowNumerator,omitempty"`  // The value of the low limit numerator.
	HighNumerator        *Quantity           `bson:",omitempty" json:"highNumerator,omitempty"` // The value of the high limit numerator.
	Denominator          *Quantity           `bson:",omitempty" json:"denominator,omitempty"`   // The value of the denominator.
	ResourceType         string              `bson:"-" json:"resourceType,omitempty"`
}

func (out *Reference) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Reference\"" {
		return fmt.Errorf("resourceType is not %s", "Reference")
	}
	if len(asMap["display"]) > 0 {
		if err := go1.Unmarshal(asMap["display"], &out.Display); err != nil {
			return err
		}

	}
	if len(asMap["_display"]) > 0 {
		if err := go1.Unmarshal(asMap["_display"], &out.DisplayPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["reference"]) > 0 {
		if err := go1.Unmarshal(asMap["reference"], &out.Reference); err != nil {
			return err
		}

	}
	if len(asMap["_reference"]) > 0 {
		if err := go1.Unmarshal(asMap["_reference"], &out.ReferencePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["type"]) > 0 {
		if err := go1.Unmarshal(asMap["type"], &out.Type); err != nil {
			return err
		}

	}
	if len(asMap["_type"]) > 0 {
		if err := go1.Unmarshal(asMap["_type"], &out.TypePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["identifier"]) > 0 {
		if err := go1.Unmarshal(asMap["identifier"], &out.Identifier); err != nil {
			return err
		}

	}
	return nil
}

type Reference struct {
	Display                     string              `bson:",omitempty" json:"display,omitempty"` // Plain text narrative that identifies the resource in addition to the resource reference.
	DisplayPrimitiveExtension   *PrimitiveExtension `bson:"display_extension,omitempty" json:"_display,omitempty"`
	Id                          *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension        *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension                   []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Reference                   string              `bson:",omitempty" json:"reference,omitempty"` // A reference to a location at which the other resource is found. The reference may be a relative reference, in which case it is relative to the service base URL, or an absolute URL that resolves to the location where the resource is found. The reference may be version specific or not. If the reference is not to a FHIR RESTful server, then it should be assumed to be version specific. Internal fragment references (start with '#') refer to contained resources.
	ReferencePrimitiveExtension *PrimitiveExtension `bson:"reference_extension,omitempty" json:"_reference,omitempty"`
	Type                        string              `bson:",omitempty" json:"type,omitempty"` /*
	The expected type of the target of the reference. If both Reference.type and Reference.reference are populated and Reference.reference is a FHIR URL, both SHALL be consistent.

	The type is the Canonical URL of Resource Definition that is the type this reference refers to. References are URLs that are relative to http://hl7.org/fhir/StructureDefinition/ e.g. "Patient" is a reference to http://hl7.org/fhir/StructureDefinition/Patient. Absolute URLs are only allowed for logical models (and can only be used in references in logical models, not resources).
	*/
	TypePrimitiveExtension *PrimitiveExtension `bson:"type_extension,omitempty" json:"_type,omitempty"`
	Identifier             *Identifier         `bson:",omitempty" json:"identifier,omitempty"` // An identifier for the target resource. This is used when there is no way to reference the other resource directly, either because the entity it represents is not available through a FHIR server, or because there is no way for the author of the resource to convert a known identifier to an actual location. There is no requirement that a Reference.identifier point to something that is actually exposed as a FHIR instance, but it SHALL point to a business concept that would be expected to be exposed as a FHIR instance, and that instance would need to be of a FHIR resource type allowed by the reference.
	ResourceType           string              `bson:"-" json:"resourceType,omitempty"`
}

func (out *RelatedArtifact) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"RelatedArtifact\"" {
		return fmt.Errorf("resourceType is not %s", "RelatedArtifact")
	}
	if len(asMap["display"]) > 0 {
		if err := go1.Unmarshal(asMap["display"], &out.Display); err != nil {
			return err
		}

	}
	if len(asMap["_display"]) > 0 {
		if err := go1.Unmarshal(asMap["_display"], &out.DisplayPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["citation"]) > 0 {
		if err := go1.Unmarshal(asMap["citation"], &out.Citation); err != nil {
			return err
		}

	}
	if len(asMap["_citation"]) > 0 {
		if err := go1.Unmarshal(asMap["_citation"], &out.CitationPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["resource"]) > 0 {
		if err := go1.Unmarshal(asMap["resource"], &out.Resource); err != nil {
			return err
		}

	}
	if len(asMap["_resource"]) > 0 {
		if err := go1.Unmarshal(asMap["_resource"], &out.ResourcePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["publicationStatus"]) > 0 {
		if err := go1.Unmarshal(asMap["publicationStatus"], &out.PublicationStatus); err != nil {
			return err
		}

	}
	if len(asMap["_publicationStatus"]) > 0 {
		if err := go1.Unmarshal(asMap["_publicationStatus"], &out.PublicationStatusPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["publicationDate"]) > 0 {
		if err := go1.Unmarshal(asMap["publicationDate"], &out.PublicationDate); err != nil {
			return err
		}

	}
	if len(asMap["_publicationDate"]) > 0 {
		if err := go1.Unmarshal(asMap["_publicationDate"], &out.PublicationDatePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["label"]) > 0 {
		if err := go1.Unmarshal(asMap["label"], &out.Label); err != nil {
			return err
		}

	}
	if len(asMap["_label"]) > 0 {
		if err := go1.Unmarshal(asMap["_label"], &out.LabelPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["classifier"]) > 0 {
		if err := go1.Unmarshal(asMap["classifier"], &out.Classifier); err != nil {
			return err
		}

	}
	if len(asMap["document"]) > 0 {
		if err := go1.Unmarshal(asMap["document"], &out.Document); err != nil {
			return err
		}

	}
	if len(asMap["resourceReference"]) > 0 {
		if err := go1.Unmarshal(asMap["resourceReference"], &out.ResourceReference); err != nil {
			return err
		}

	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if err := go1.Unmarshal(asMap["type"], &out.Type); err != nil {
		return err
	}
	if len(asMap["_type"]) > 0 {
		if err := go1.Unmarshal(asMap["_type"], &out.TypePrimitiveExtension); err != nil {
			return err
		}
	}

	return nil
}

type RelatedArtifact struct {
	Display                             string              `bson:",omitempty" json:"display,omitempty"` // A brief description of the document or knowledge resource being referenced, suitable for display to a consumer.
	DisplayPrimitiveExtension           *PrimitiveExtension `bson:"display_extension,omitempty" json:"_display,omitempty"`
	Citation                            string              `bson:",omitempty" json:"citation,omitempty"` // A bibliographic citation for the related artifact. This text SHOULD be formatted according to an accepted citation format.
	CitationPrimitiveExtension          *PrimitiveExtension `bson:"citation_extension,omitempty" json:"_citation,omitempty"`
	Resource                            string              `bson:",omitempty" json:"resource,omitempty"` // The related artifact, such as a library, value set, profile, or other knowledge resource.
	ResourcePrimitiveExtension          *PrimitiveExtension `bson:"resource_extension,omitempty" json:"_resource,omitempty"`
	PublicationStatus                   string              `bson:",omitempty" json:"publicationStatus,omitempty"` // The publication status of the artifact being referred to.
	PublicationStatusPrimitiveExtension *PrimitiveExtension `bson:"publicationStatus_extension,omitempty" json:"_publicationStatus,omitempty"`
	PublicationDate                     *Date               `bson:",omitempty" json:"publicationDate,omitempty"` // The date of publication of the artifact being referred to.
	PublicationDatePrimitiveExtension   *PrimitiveExtension `bson:"publicationDate_extension,omitempty" json:"_publicationDate,omitempty"`
	Extension                           []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Label                               string              `bson:",omitempty" json:"label,omitempty"`     // A short label that can be used to reference the citation from elsewhere in the containing artifact, such as a footnote index.
	LabelPrimitiveExtension             *PrimitiveExtension `bson:"label_extension,omitempty" json:"_label,omitempty"`
	Classifier                          []*CodeableConcept  `bson:",omitempty" json:"classifier,omitempty"`        // Provides additional classifiers of the related artifact.
	Document                            *Attachment         `bson:",omitempty" json:"document,omitempty"`          // The document being referenced, represented as an attachment. This is exclusive with the resource element.
	ResourceReference                   *Reference          `bson:",omitempty" json:"resourceReference,omitempty"` // The related artifact, if the artifact is not a canonical resource, or a resource reference to a canonical resource.
	Id                                  *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`             // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension                *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Type                                string              `binding:"required" bson:",omitempty" json:"type,omitempty"` // The type of relationship to the related artifact.
	TypePrimitiveExtension              *PrimitiveExtension `bson:"type_extension,omitempty" json:"_type,omitempty"`
	ResourceType                        string              `bson:"-" json:"resourceType,omitempty"`
}

func (out *SampledData) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"SampledData\"" {
		return fmt.Errorf("resourceType is not %s", "SampledData")
	}
	if len(asMap["lowerLimit"]) > 0 {
		if err := go1.Unmarshal(asMap["lowerLimit"], &out.LowerLimit); err != nil {
			return err
		}

	}
	if len(asMap["_lowerLimit"]) > 0 {
		if err := go1.Unmarshal(asMap["_lowerLimit"], &out.LowerLimitPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["upperLimit"]) > 0 {
		if err := go1.Unmarshal(asMap["upperLimit"], &out.UpperLimit); err != nil {
			return err
		}

	}
	if len(asMap["_upperLimit"]) > 0 {
		if err := go1.Unmarshal(asMap["_upperLimit"], &out.UpperLimitPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["codeMap"]) > 0 {
		if err := go1.Unmarshal(asMap["codeMap"], &out.CodeMap); err != nil {
			return err
		}

	}
	if len(asMap["_codeMap"]) > 0 {
		if err := go1.Unmarshal(asMap["_codeMap"], &out.CodeMapPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["offsets"]) > 0 {
		if err := go1.Unmarshal(asMap["offsets"], &out.Offsets); err != nil {
			return err
		}

	}
	if len(asMap["_offsets"]) > 0 {
		if err := go1.Unmarshal(asMap["_offsets"], &out.OffsetsPrimitiveExtension); err != nil {
			return err
		}
	}

	if err := go1.Unmarshal(asMap["origin"], &out.Origin); err != nil {
		return err
	}

	if len(asMap["interval"]) > 0 {
		if err := go1.Unmarshal(asMap["interval"], &out.Interval); err != nil {
			return err
		}

	}
	if len(asMap["_interval"]) > 0 {
		if err := go1.Unmarshal(asMap["_interval"], &out.IntervalPrimitiveExtension); err != nil {
			return err
		}
	}

	if err := go1.Unmarshal(asMap["intervalUnit"], &out.IntervalUnit); err != nil {
		return err
	}
	if len(asMap["_intervalUnit"]) > 0 {
		if err := go1.Unmarshal(asMap["_intervalUnit"], &out.IntervalUnitPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["factor"]) > 0 {
		if err := go1.Unmarshal(asMap["factor"], &out.Factor); err != nil {
			return err
		}

	}
	if len(asMap["_factor"]) > 0 {
		if err := go1.Unmarshal(asMap["_factor"], &out.FactorPrimitiveExtension); err != nil {
			return err
		}
	}

	if err := go1.Unmarshal(asMap["dimensions"], &out.Dimensions); err != nil {
		return err
	}
	if len(asMap["_dimensions"]) > 0 {
		if err := go1.Unmarshal(asMap["_dimensions"], &out.DimensionsPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["data"]) > 0 {
		if err := go1.Unmarshal(asMap["data"], &out.Data); err != nil {
			return err
		}

	}
	if len(asMap["_data"]) > 0 {
		if err := go1.Unmarshal(asMap["_data"], &out.DataPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	return nil
}

type SampledData struct {
	LowerLimit                     float64             `bson:",omitempty" json:"lowerLimit,omitempty"` // The lower limit of detection of the measured points. This is needed if any of the data points have the value "L" (lower than detection limit).
	LowerLimitPrimitiveExtension   *PrimitiveExtension `bson:"lowerLimit_extension,omitempty" json:"_lowerLimit,omitempty"`
	UpperLimit                     float64             `bson:",omitempty" json:"upperLimit,omitempty"` // The upper limit of detection of the measured points. This is needed if any of the data points have the value "U" (higher than detection limit).
	UpperLimitPrimitiveExtension   *PrimitiveExtension `bson:"upperLimit_extension,omitempty" json:"_upperLimit,omitempty"`
	CodeMap                        string              `bson:",omitempty" json:"codeMap,omitempty"` // Reference to ConceptMap that defines the codes used in the data.
	CodeMapPrimitiveExtension      *PrimitiveExtension `bson:"codeMap_extension,omitempty" json:"_codeMap,omitempty"`
	Offsets                        string              `bson:",omitempty" json:"offsets,omitempty"` // A series of data points which are decimal values separated by a single space (character u20).  The units in which the offsets are expressed are found in intervalUnit.  The absolute point at which the measurements begin SHALL be conveyed outside the scope of this datatype, e.g. Observation.effectiveDateTime for a timing offset.
	OffsetsPrimitiveExtension      *PrimitiveExtension `bson:"offsets_extension,omitempty" json:"_offsets,omitempty"`
	Origin                         *Quantity           `binding:"required" bson:",omitempty" json:"origin,omitempty"` // The base quantity that a measured value of zero represents. In addition, this provides the units of the entire measurement series.
	Interval                       float64             `bson:",omitempty" json:"interval,omitempty"`                  // Amount of intervalUnits between samples, e.g. milliseconds for time-based sampling.
	IntervalPrimitiveExtension     *PrimitiveExtension `bson:"interval_extension,omitempty" json:"_interval,omitempty"`
	IntervalUnit                   string              `binding:"required" bson:",omitempty" json:"intervalUnit,omitempty"` // The measurement unit in which the sample interval is expressed.
	IntervalUnitPrimitiveExtension *PrimitiveExtension `bson:"intervalUnit_extension,omitempty" json:"_intervalUnit,omitempty"`
	Factor                         float64             `bson:",omitempty" json:"factor,omitempty"` // A correction factor that is applied to the sampled data points before they are added to the origin.
	FactorPrimitiveExtension       *PrimitiveExtension `bson:"factor_extension,omitempty" json:"_factor,omitempty"`
	Dimensions                     int                 `binding:"required" bson:",omitempty" json:"dimensions,omitempty"` // The number of sample points at each time point. If this value is greater than one, then the dimensions will be interlaced - all the sample points for a point in time will be recorded at once.
	DimensionsPrimitiveExtension   *PrimitiveExtension `bson:"dimensions_extension,omitempty" json:"_dimensions,omitempty"`
	Data                           string              `bson:",omitempty" json:"data,omitempty"` // A series of data points which are decimal values or codes separated by a single space (character u20). The special codes "E" (error), "L" (below detection limit) and "U" (above detection limit) are also defined for used in place of decimal values.
	DataPrimitiveExtension         *PrimitiveExtension `bson:"data_extension,omitempty" json:"_data,omitempty"`
	Id                             *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension           *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension                      []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	ResourceType                   string              `bson:"-" json:"resourceType,omitempty"`
}

func (out *Signature) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Signature\"" {
		return fmt.Errorf("resourceType is not %s", "Signature")
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["type"]) > 0 {
		if err := go1.Unmarshal(asMap["type"], &out.Type); err != nil {
			return err
		}

	}
	if len(asMap["onBehalfOf"]) > 0 {
		if err := go1.Unmarshal(asMap["onBehalfOf"], &out.OnBehalfOf); err != nil {
			return err
		}

	}
	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["when"]) > 0 {
		if err := go1.Unmarshal(asMap["when"], &out.When); err != nil {
			return err
		}

	}
	if len(asMap["_when"]) > 0 {
		if err := go1.Unmarshal(asMap["_when"], &out.WhenPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["who"]) > 0 {
		if err := go1.Unmarshal(asMap["who"], &out.Who); err != nil {
			return err
		}

	}
	if len(asMap["targetFormat"]) > 0 {
		if err := go1.Unmarshal(asMap["targetFormat"], &out.TargetFormat); err != nil {
			return err
		}

	}
	if len(asMap["_targetFormat"]) > 0 {
		if err := go1.Unmarshal(asMap["_targetFormat"], &out.TargetFormatPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["sigFormat"]) > 0 {
		if err := go1.Unmarshal(asMap["sigFormat"], &out.SigFormat); err != nil {
			return err
		}

	}
	if len(asMap["_sigFormat"]) > 0 {
		if err := go1.Unmarshal(asMap["_sigFormat"], &out.SigFormatPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["data"]) > 0 {
		if err := go1.Unmarshal(asMap["data"], &out.Data); err != nil {
			return err
		}

	}
	if len(asMap["_data"]) > 0 {
		if err := go1.Unmarshal(asMap["_data"], &out.DataPrimitiveExtension); err != nil {
			return err
		}
	}

	return nil
}

type Signature struct {
	Id                             *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension           *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Type                           []*Coding           `bson:",omitempty" json:"type,omitempty"`       // An indication of the reason that the entity signed this document. This may be explicitly included as part of the signature information and can be used when determining accountability for various actions concerning the document.
	OnBehalfOf                     *Reference          `bson:",omitempty" json:"onBehalfOf,omitempty"` // A reference to an application-usable description of the identity that is represented by the signature.
	Extension                      []*Extension        `bson:",omitempty" json:"extension,omitempty"`  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	When                           *time.Time          `bson:",omitempty" json:"when,omitempty"`       // When the digital signature was signed.
	WhenPrimitiveExtension         *PrimitiveExtension `bson:"when_extension,omitempty" json:"_when,omitempty"`
	Who                            *Reference          `bson:",omitempty" json:"who,omitempty"`          // A reference to an application-usable description of the identity that signed  (e.g. the signature used their private key).
	TargetFormat                   string              `bson:",omitempty" json:"targetFormat,omitempty"` // A mime type that indicates the technical format of the target resources signed by the signature.
	TargetFormatPrimitiveExtension *PrimitiveExtension `bson:"targetFormat_extension,omitempty" json:"_targetFormat,omitempty"`
	SigFormat                      string              `bson:",omitempty" json:"sigFormat,omitempty"` // A mime type that indicates the technical format of the signature. Important mime types are application/signature+xml for X ML DigSig, application/jose for JWS, and image/* for a graphical image of a signature, etc.
	SigFormatPrimitiveExtension    *PrimitiveExtension `bson:"sigFormat_extension,omitempty" json:"_sigFormat,omitempty"`
	Data                           Base64Binary        `bson:",omitempty" json:"data,omitempty"` // The base64 encoding of the Signature content. When signature is not recorded electronically this element would be empty.
	DataPrimitiveExtension         *PrimitiveExtension `bson:"data_extension,omitempty" json:"_data,omitempty"`
	ResourceType                   string              `bson:"-" json:"resourceType,omitempty"`
}

type TimingRepeatBoundsx struct {
	BoundsDuration *Duration `binding:"omitempty" bson:",omitempty" json:"boundsDuration,omitempty"`
	BoundsRange    *Range    `binding:"omitempty" bson:",omitempty" json:"boundsRange,omitempty"`
	BoundsPeriod   *Period   `binding:"omitempty" bson:",omitempty" json:"boundsPeriod,omitempty"`
}

func (out *TimingRepeat) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if len(asMap["offset"]) > 0 {
		if err := go1.Unmarshal(asMap["offset"], &out.Offset); err != nil {
			return err
		}

	}
	if len(asMap["_offset"]) > 0 {
		if err := go1.Unmarshal(asMap["_offset"], &out.OffsetPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["count"]) > 0 {
		if err := go1.Unmarshal(asMap["count"], &out.Count); err != nil {
			return err
		}

	}
	if len(asMap["_count"]) > 0 {
		if err := go1.Unmarshal(asMap["_count"], &out.CountPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["durationUnit"]) > 0 {
		if err := go1.Unmarshal(asMap["durationUnit"], &out.DurationUnit); err != nil {
			return err
		}

	}
	if len(asMap["_durationUnit"]) > 0 {
		if err := go1.Unmarshal(asMap["_durationUnit"], &out.DurationUnitPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["period"]) > 0 {
		if err := go1.Unmarshal(asMap["period"], &out.Period); err != nil {
			return err
		}

	}
	if len(asMap["_period"]) > 0 {
		if err := go1.Unmarshal(asMap["_period"], &out.PeriodPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["periodUnit"]) > 0 {
		if err := go1.Unmarshal(asMap["periodUnit"], &out.PeriodUnit); err != nil {
			return err
		}

	}
	if len(asMap["_periodUnit"]) > 0 {
		if err := go1.Unmarshal(asMap["_periodUnit"], &out.PeriodUnitPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["timeOfDay"]) > 0 {
		if err := go1.Unmarshal(asMap["timeOfDay"], &out.TimeOfDay); err != nil {
			return err
		}

	}
	if len(asMap["_timeOfDay"]) > 0 {
		if err := go1.Unmarshal(asMap["_timeOfDay"], &out.TimeOfDayPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["periodMax"]) > 0 {
		if err := go1.Unmarshal(asMap["periodMax"], &out.PeriodMax); err != nil {
			return err
		}

	}
	if len(asMap["_periodMax"]) > 0 {
		if err := go1.Unmarshal(asMap["_periodMax"], &out.PeriodMaxPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["dayOfWeek"]) > 0 {
		if err := go1.Unmarshal(asMap["dayOfWeek"], &out.DayOfWeek); err != nil {
			return err
		}

	}
	if len(asMap["_dayOfWeek"]) > 0 {
		if err := go1.Unmarshal(asMap["_dayOfWeek"], &out.DayOfWeekPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["when"]) > 0 {
		if err := go1.Unmarshal(asMap["when"], &out.When); err != nil {
			return err
		}

	}
	if len(asMap["_when"]) > 0 {
		if err := go1.Unmarshal(asMap["_when"], &out.WhenPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["boundsDuration"]) > 0 {
		if err := go1.Unmarshal(asMap["boundsDuration"], &out.BoundsDuration); err == nil {
		}
	} else if len(asMap["boundsRange"]) > 0 {
		if err := go1.Unmarshal(asMap["boundsRange"], &out.BoundsRange); err == nil {
		}
	} else if len(asMap["boundsPeriod"]) > 0 {
		if err := go1.Unmarshal(asMap["boundsPeriod"], &out.BoundsPeriod); err == nil {
		}
	} else {

	}
	if len(asMap["countMax"]) > 0 {
		if err := go1.Unmarshal(asMap["countMax"], &out.CountMax); err != nil {
			return err
		}

	}
	if len(asMap["_countMax"]) > 0 {
		if err := go1.Unmarshal(asMap["_countMax"], &out.CountMaxPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["frequencyMax"]) > 0 {
		if err := go1.Unmarshal(asMap["frequencyMax"], &out.FrequencyMax); err != nil {
			return err
		}

	}
	if len(asMap["_frequencyMax"]) > 0 {
		if err := go1.Unmarshal(asMap["_frequencyMax"], &out.FrequencyMaxPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["frequency"]) > 0 {
		if err := go1.Unmarshal(asMap["frequency"], &out.Frequency); err != nil {
			return err
		}

	}
	if len(asMap["_frequency"]) > 0 {
		if err := go1.Unmarshal(asMap["_frequency"], &out.FrequencyPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["duration"]) > 0 {
		if err := go1.Unmarshal(asMap["duration"], &out.Duration); err != nil {
			return err
		}

	}
	if len(asMap["_duration"]) > 0 {
		if err := go1.Unmarshal(asMap["_duration"], &out.DurationPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["durationMax"]) > 0 {
		if err := go1.Unmarshal(asMap["durationMax"], &out.DurationMax); err != nil {
			return err
		}

	}
	if len(asMap["_durationMax"]) > 0 {
		if err := go1.Unmarshal(asMap["_durationMax"], &out.DurationMaxPrimitiveExtension); err != nil {
			return err
		}
	}

	return nil
}

type TimingRepeat struct {
	Offset                   int                 `bson:",omitempty" json:"offset,omitempty"` // The number of minutes from the event. If the event code does not indicate whether the minutes is before or after the event, then the offset is assumed to be after the event.
	OffsetPrimitiveExtension *PrimitiveExtension `bson:"offset_extension,omitempty" json:"_offset,omitempty"`
	Count                    int                 `bson:",omitempty" json:"count,omitempty"` // A total count of the desired number of repetitions across the duration of the entire timing specification. If countMax is present, this element indicates the lower bound of the allowed range of count values.
	CountPrimitiveExtension  *PrimitiveExtension `bson:"count_extension,omitempty" json:"_count,omitempty"`
	DurationUnit             string              `bson:",omitempty" json:"durationUnit,omitempty"` /*
	The units of time for the duration, in UCUM units
	Normal practice is to use the 'mo' code as a calendar month when calculating the next occurrence.
	*/
	DurationUnitPrimitiveExtension *PrimitiveExtension `bson:"durationUnit_extension,omitempty" json:"_durationUnit,omitempty"`
	Period                         float64             `bson:",omitempty" json:"period,omitempty"` // Indicates the duration of time over which repetitions are to occur; e.g. to express "3 times per day", 3 would be the frequency and "1 day" would be the period. If periodMax is present, this element indicates the lower bound of the allowed range of the period length.
	PeriodPrimitiveExtension       *PrimitiveExtension `bson:"period_extension,omitempty" json:"_period,omitempty"`
	PeriodUnit                     string              `bson:",omitempty" json:"periodUnit,omitempty"` /*
	The units of time for the period in UCUM units
	Normal practice is to use the 'mo' code as a calendar month when calculating the next occurrence.
	*/
	PeriodUnitPrimitiveExtension   *PrimitiveExtension `bson:"periodUnit_extension,omitempty" json:"_periodUnit,omitempty"`
	TimeOfDay                      []*Time             `bson:",omitempty" json:"timeOfDay,omitempty"` // Specified time of day for action to take place.
	TimeOfDayPrimitiveExtension    *PrimitiveExtension `bson:"timeOfDay_extension,omitempty" json:"_timeOfDay,omitempty"`
	PeriodMax                      float64             `bson:",omitempty" json:"periodMax,omitempty"` // If present, indicates that the period is a range from [period] to [periodMax], allowing expressing concepts such as "do this once every 3-5 days.
	PeriodMaxPrimitiveExtension    *PrimitiveExtension `bson:"periodMax_extension,omitempty" json:"_periodMax,omitempty"`
	DayOfWeek                      []string            `bson:",omitempty" json:"dayOfWeek,omitempty"` // If one or more days of week is provided, then the action happens only on the specified day(s).
	DayOfWeekPrimitiveExtension    *PrimitiveExtension `bson:"dayOfWeek_extension,omitempty" json:"_dayOfWeek,omitempty"`
	When                           []string            `bson:",omitempty" json:"when,omitempty"` // An approximate time period during the day, potentially linked to an event of daily living that indicates when the action should occur.
	WhenPrimitiveExtension         *PrimitiveExtension `bson:"when_extension,omitempty" json:"_when,omitempty"`
	Id                             string              `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension           *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension                      []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	TimingRepeatBoundsx            `binding:"omitempty" bson:",omitempty" json:",omitempty"`
	CountMax                       int                 `bson:",omitempty" json:"countMax,omitempty"` // If present, indicates that the count is a range - so to perform the action between [count] and [countMax] times.
	CountMaxPrimitiveExtension     *PrimitiveExtension `bson:"countMax_extension,omitempty" json:"_countMax,omitempty"`
	FrequencyMax                   int                 `bson:",omitempty" json:"frequencyMax,omitempty"` // If present, indicates that the frequency is a range - so to repeat between [frequency] and [frequencyMax] times within the period or period range.
	FrequencyMaxPrimitiveExtension *PrimitiveExtension `bson:"frequencyMax_extension,omitempty" json:"_frequencyMax,omitempty"`
	Frequency                      int                 `bson:",omitempty" json:"frequency,omitempty"` // The number of times to repeat the action within the specified period. If frequencyMax is present, this element indicates the lower bound of the allowed range of the frequency.
	FrequencyPrimitiveExtension    *PrimitiveExtension `bson:"frequency_extension,omitempty" json:"_frequency,omitempty"`
	Duration                       float64             `bson:",omitempty" json:"duration,omitempty"` // How long this thing happens for when it happens. If durationMax is present, this element indicates the lower bound of the allowed range of the duration.
	DurationPrimitiveExtension     *PrimitiveExtension `bson:"duration_extension,omitempty" json:"_duration,omitempty"`
	DurationMax                    float64             `bson:",omitempty" json:"durationMax,omitempty"` // If present, indicates that the duration is a range - so to perform the action between [duration] and [durationMax] time length.
	DurationMaxPrimitiveExtension  *PrimitiveExtension `bson:"durationMax_extension,omitempty" json:"_durationMax,omitempty"`
}

func (out *Timing) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Timing\"" {
		return fmt.Errorf("resourceType is not %s", "Timing")
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["modifierExtension"]) > 0 {
		if err := go1.Unmarshal(asMap["modifierExtension"], &out.ModifierExtension); err != nil {
			return err
		}

	}
	if len(asMap["event"]) > 0 {
		if err := go1.Unmarshal(asMap["event"], &out.Event); err != nil {
			return err
		}

	}
	if len(asMap["_event"]) > 0 {
		if err := go1.Unmarshal(asMap["_event"], &out.EventPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["repeat"]) > 0 {
		if err := go1.Unmarshal(asMap["repeat"], &out.Repeat); err != nil {
			return err
		}

	}
	if len(asMap["code"]) > 0 {
		if err := go1.Unmarshal(asMap["code"], &out.Code); err != nil {
			return err
		}

	}
	return nil
}

type Timing struct {
	Id                   *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension            []*Extension        `bson:",omitempty" json:"extension,omitempty"`         // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	ModifierExtension    []*Extension        `bson:",omitempty" json:"modifierExtension,omitempty"` /*
	May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.

	Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	*/
	Event                   []*DateTime         `bson:",omitempty" json:"event,omitempty"` // Identifies specific times when the event occurs.
	EventPrimitiveExtension *PrimitiveExtension `bson:"event_extension,omitempty" json:"_event,omitempty"`
	Repeat                  *TimingRepeat       `bson:",omitempty" json:"repeat,omitempty"`
	Code                    *CodeableConcept    `bson:",omitempty" json:"code,omitempty"` // A code for the timing schedule (or just text in code.text). Some codes such as BID are ubiquitous, but many institutions define their own additional codes. If a code is provided, the code is understood to be a complete statement of whatever is specified in the structured timing data, and either the code or the data may be used to interpret the Timing, with the exception that .repeat.bounds still applies over the code (and is not contained in the code).
	ResourceType            string              `bson:"-" json:"resourceType,omitempty"`
}

type TriggerDefinitionTimingx struct {
	TimingTiming    *Timing    `binding:"omitempty" bson:",omitempty" json:"timingTiming,omitempty"`
	TimingReference *Reference `binding:"omitempty" bson:",omitempty" json:"timingReference,omitempty"`
	TimingDate      *Date      `binding:"omitempty" bson:",omitempty" json:"timingDate,omitempty"`
	TimingDateTime  *DateTime  `binding:"omitempty" bson:",omitempty" json:"timingDateTime,omitempty"`
}

func (out *TriggerDefinition) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"TriggerDefinition\"" {
		return fmt.Errorf("resourceType is not %s", "TriggerDefinition")
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["data"]) > 0 {
		if err := go1.Unmarshal(asMap["data"], &out.Data); err != nil {
			return err
		}

	}
	if len(asMap["condition"]) > 0 {
		if err := go1.Unmarshal(asMap["condition"], &out.Condition); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["type"], &out.Type); err != nil {
		return err
	}
	if len(asMap["_type"]) > 0 {
		if err := go1.Unmarshal(asMap["_type"], &out.TypePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["name"]) > 0 {
		if err := go1.Unmarshal(asMap["name"], &out.Name); err != nil {
			return err
		}

	}
	if len(asMap["_name"]) > 0 {
		if err := go1.Unmarshal(asMap["_name"], &out.NamePrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["code"]) > 0 {
		if err := go1.Unmarshal(asMap["code"], &out.Code); err != nil {
			return err
		}

	}
	if len(asMap["subscriptionTopic"]) > 0 {
		if err := go1.Unmarshal(asMap["subscriptionTopic"], &out.SubscriptionTopic); err != nil {
			return err
		}

	}
	if len(asMap["_subscriptionTopic"]) > 0 {
		if err := go1.Unmarshal(asMap["_subscriptionTopic"], &out.SubscriptionTopicPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["timingTiming"]) > 0 {
		if err := go1.Unmarshal(asMap["timingTiming"], &out.TimingTiming); err == nil {
		}
	} else if len(asMap["timingReference"]) > 0 {
		if err := go1.Unmarshal(asMap["timingReference"], &out.TimingReference); err == nil {
		}
	} else if len(asMap["timingDate"]) > 0 {
		if err := go1.Unmarshal(asMap["timingDate"], &out.TimingDate); err == nil {
		}
	} else if len(asMap["timingDateTime"]) > 0 {
		if err := go1.Unmarshal(asMap["timingDateTime"], &out.TimingDateTime); err == nil {
		}
	} else {

	}
	return nil
}

type TriggerDefinition struct {
	Id                                  *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension                *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension                           []*Extension        `bson:",omitempty" json:"extension,omitempty"`               // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Data                                []*DataRequirement  `bson:",omitempty" json:"data,omitempty"`                    // The triggering data of the event (if this is a data trigger). If more than one data is requirement is specified, then all the data requirements must be true.
	Condition                           *Expression         `bson:",omitempty" json:"condition,omitempty"`               // A boolean-valued expression that is evaluated in the context of the container of the trigger definition and returns whether or not the trigger fires.
	Type                                string              `binding:"required" bson:",omitempty" json:"type,omitempty"` // The type of triggering event.
	TypePrimitiveExtension              *PrimitiveExtension `bson:"type_extension,omitempty" json:"_type,omitempty"`
	Name                                string              `bson:",omitempty" json:"name,omitempty"` // A formal name for the event. This may be an absolute URI that identifies the event formally (e.g. from a trigger registry), or a simple relative URI that identifies the event in a local context.
	NamePrimitiveExtension              *PrimitiveExtension `bson:"name_extension,omitempty" json:"_name,omitempty"`
	Code                                *CodeableConcept    `bson:",omitempty" json:"code,omitempty"`              // A code that identifies the event.
	SubscriptionTopic                   string              `bson:",omitempty" json:"subscriptionTopic,omitempty"` // A reference to a SubscriptionTopic resource that defines the event. If this element is provided, no other information about the trigger definition may be supplied.
	SubscriptionTopicPrimitiveExtension *PrimitiveExtension `bson:"subscriptionTopic_extension,omitempty" json:"_subscriptionTopic,omitempty"`
	TriggerDefinitionTimingx            `binding:"omitempty" bson:",omitempty" json:",omitempty"`
	ResourceType                        string `bson:"-" json:"resourceType,omitempty"`
}

type UsageContextValuex struct {
	ValueCodeableConcept *CodeableConcept `binding:"omitempty" bson:",omitempty" json:"valueCodeableConcept,omitempty"`
	ValueQuantity        *Quantity        `binding:"omitempty" bson:",omitempty" json:"valueQuantity,omitempty"`
	ValueRange           *Range           `binding:"omitempty" bson:",omitempty" json:"valueRange,omitempty"`
	ValueReference       *Reference       `binding:"omitempty" bson:",omitempty" json:"valueReference,omitempty"`
}

func (out *UsageContext) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"UsageContext\"" {
		return fmt.Errorf("resourceType is not %s", "UsageContext")
	}
	if err := go1.Unmarshal(asMap["code"], &out.Code); err != nil {
		return err
	}

	if len(asMap["valueCodeableConcept"]) > 0 {
		if err := go1.Unmarshal(asMap["valueCodeableConcept"], &out.ValueCodeableConcept); err == nil {
		}
	} else if len(asMap["valueQuantity"]) > 0 {
		if err := go1.Unmarshal(asMap["valueQuantity"], &out.ValueQuantity); err == nil {
		}
	} else if len(asMap["valueRange"]) > 0 {
		if err := go1.Unmarshal(asMap["valueRange"], &out.ValueRange); err == nil {
		}
	} else if len(asMap["valueReference"]) > 0 {
		if err := go1.Unmarshal(asMap["valueReference"], &out.ValueReference); err == nil {
		}
	} else {
		return fmt.Errorf("could not unmarshal %s into any of the possible types", "value[x]")
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	return nil
}

type UsageContext struct {
	Code                 *Coding `binding:"required" bson:",omitempty" json:"code,omitempty"` // A code that identifies the type of context being specified by this usage context.
	UsageContextValuex   `binding:"omitempty" bson:",omitempty" json:",omitempty"`
	Id                   *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	Extension            []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	ResourceType         string              `bson:"-" json:"resourceType,omitempty"`
}

type VirtualServiceDetailAddressx struct {
	AddressUrl                   *url.URL               `binding:"omitempty" bson:",omitempty" json:"addressUrl,omitempty"`
	AddressString                string                 `binding:"omitempty" bson:",omitempty" json:"addressString,omitempty"`
	AddressContactPoint          *ContactPoint          `binding:"omitempty" bson:",omitempty" json:"addressContactPoint,omitempty"`
	AddressExtendedContactDetail *ExtendedContactDetail `binding:"omitempty" bson:",omitempty" json:"addressExtendedContactDetail,omitempty"`
}

func (out *VirtualServiceDetail) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"VirtualServiceDetail\"" {
		return fmt.Errorf("resourceType is not %s", "VirtualServiceDetail")
	}
	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["channelType"]) > 0 {
		if err := go1.Unmarshal(asMap["channelType"], &out.ChannelType); err != nil {
			return err
		}

	}
	if len(asMap["addressUrl"]) > 0 {
		if err := go1.Unmarshal(asMap["addressUrl"], &out.AddressUrl); err == nil {
		}
	} else if len(asMap["addressString"]) > 0 {
		if err := go1.Unmarshal(asMap["addressString"], &out.AddressString); err == nil {
		}
	} else if len(asMap["addressContactPoint"]) > 0 {
		if err := go1.Unmarshal(asMap["addressContactPoint"], &out.AddressContactPoint); err == nil {
		}
	} else if len(asMap["addressExtendedContactDetail"]) > 0 {
		if err := go1.Unmarshal(asMap["addressExtendedContactDetail"], &out.AddressExtendedContactDetail); err == nil {
		}
	} else {

	}
	if len(asMap["additionalInfo"]) > 0 {
		if err := go1.Unmarshal(asMap["additionalInfo"], &out.AdditionalInfo); err != nil {
			return err
		}

	}
	if len(asMap["_additionalInfo"]) > 0 {
		if err := go1.Unmarshal(asMap["_additionalInfo"], &out.AdditionalInfoPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["maxParticipants"]) > 0 {
		if err := go1.Unmarshal(asMap["maxParticipants"], &out.MaxParticipants); err != nil {
			return err
		}

	}
	if len(asMap["_maxParticipants"]) > 0 {
		if err := go1.Unmarshal(asMap["_maxParticipants"], &out.MaxParticipantsPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["sessionKey"]) > 0 {
		if err := go1.Unmarshal(asMap["sessionKey"], &out.SessionKey); err != nil {
			return err
		}

	}
	if len(asMap["_sessionKey"]) > 0 {
		if err := go1.Unmarshal(asMap["_sessionKey"], &out.SessionKeyPrimitiveExtension); err != nil {
			return err
		}
	}

	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["_id"]) > 0 {
		if err := go1.Unmarshal(asMap["_id"], &out.IdPrimitiveExtension); err != nil {
			return err
		}
	}

	return nil
}

type VirtualServiceDetail struct {
	Extension                         []*Extension `bson:",omitempty" json:"extension,omitempty"`   // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	ChannelType                       *Coding      `bson:",omitempty" json:"channelType,omitempty"` // The type of virtual service to connect to (i.e. Teams, Zoom, Specific VMR technology, WhatsApp).
	VirtualServiceDetailAddressx      `binding:"omitempty" bson:",omitempty" json:",omitempty"`
	AdditionalInfo                    []*url.URL          `bson:",omitempty" json:"additionalInfo,omitempty"` // Address to see alternative connection details.
	AdditionalInfoPrimitiveExtension  *PrimitiveExtension `bson:"additionalInfo_extension,omitempty" json:"_additionalInfo,omitempty"`
	MaxParticipants                   int                 `bson:",omitempty" json:"maxParticipants,omitempty"` // Maximum number of participants supported by the virtual service.
	MaxParticipantsPrimitiveExtension *PrimitiveExtension `bson:"maxParticipants_extension,omitempty" json:"_maxParticipants,omitempty"`
	SessionKey                        string              `bson:",omitempty" json:"sessionKey,omitempty"` // Session Key required by the virtual service.
	SessionKeyPrimitiveExtension      *PrimitiveExtension `bson:"sessionKey_extension,omitempty" json:"_sessionKey,omitempty"`
	Id                                *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IdPrimitiveExtension              *PrimitiveExtension `bson:"id_extension,omitempty" json:"_id,omitempty"`
	ResourceType                      string              `bson:"-" json:"resourceType,omitempty"`
}
