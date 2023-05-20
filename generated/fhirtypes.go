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
	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	return nil
}

type Element struct {
	Id           string      `bson:"_id,omitempty" json:"id,omitempty"`     // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension    []Extension `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	ResourceType string      `binding:"omitempty" bson:"-" json:"resourceType"`
}

func (out *BackboneElement) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"BackboneElement\"" {
		return fmt.Errorf("resourceType is not %s", "BackboneElement")
	}
	if len(asMap["modifierExtension"]) > 0 {
		if err := go1.Unmarshal(asMap["modifierExtension"], &out.ModifierExtension); err != nil {
			return err
		}

	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
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

type BackboneElement struct {
	ModifierExtension []*Extension `bson:",omitempty" json:"modifierExtension,omitempty"` /*
	May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.

	Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	*/
	Id           *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`     // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension    []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	ResourceType string              `binding:"omitempty" bson:"-" json:"resourceType"`
}

func (out *Address) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Address\"" {
		return fmt.Errorf("resourceType is not %s", "Address")
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
	if len(asMap["text"]) > 0 {
		if err := go1.Unmarshal(asMap["text"], &out.Text); err != nil {
			return err
		}

	}
	if len(asMap["country"]) > 0 {
		if err := go1.Unmarshal(asMap["country"], &out.Country); err != nil {
			return err
		}

	}
	if len(asMap["postalCode"]) > 0 {
		if err := go1.Unmarshal(asMap["postalCode"], &out.PostalCode); err != nil {
			return err
		}

	}
	if len(asMap["type"]) > 0 {
		if err := go1.Unmarshal(asMap["type"], &out.Type); err != nil {
			return err
		}

	}
	if len(asMap["line"]) > 0 {
		if err := go1.Unmarshal(asMap["line"], &out.Line); err != nil {
			return err
		}

	}
	if len(asMap["city"]) > 0 {
		if err := go1.Unmarshal(asMap["city"], &out.City); err != nil {
			return err
		}

	}
	if len(asMap["district"]) > 0 {
		if err := go1.Unmarshal(asMap["district"], &out.District); err != nil {
			return err
		}

	}
	if len(asMap["state"]) > 0 {
		if err := go1.Unmarshal(asMap["state"], &out.State); err != nil {
			return err
		}

	}
	return nil
}

type Address struct {
	Period       Period              `bson:",omitempty" json:"period,omitempty"`     // Time period when address was/is in use.
	Id           *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`      // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension    []*Extension        `bson:",omitempty" json:"extension,omitempty"`  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Use          *string             `bson:",omitempty" json:"use,omitempty"`        // The purpose of this address.
	Text         *string             `bson:",omitempty" json:"text,omitempty"`       // Specifies the entire address as it should be displayed e.g. on a postal label. This may be provided instead of or as well as the specific parts.
	Country      *string             `bson:",omitempty" json:"country,omitempty"`    // Country - a nation as commonly understood or generally accepted.
	PostalCode   *string             `bson:",omitempty" json:"postalCode,omitempty"` // A postal code designating a region defined by the postal service.
	Type         *string             `bson:",omitempty" json:"type,omitempty"`       // Distinguishes between physical addresses (those you can visit) and mailing addresses (e.g. PO Boxes and care-of addresses). Most addresses are both.
	Line         []*string           `bson:",omitempty" json:"line,omitempty"`       // This component contains the house number, apartment number, street name, street direction,  P.O. Box number, delivery hints, and similar address information.
	City         *string             `bson:",omitempty" json:"city,omitempty"`       // The name of the city, town, suburb, village or other community or delivery center.
	District     *string             `bson:",omitempty" json:"district,omitempty"`   // The name of the administrative area (county).
	State        *string             `bson:",omitempty" json:"state,omitempty"`      // Sub-unit of a country with limited sovereignty in a federally organized country. A code may be used if codes are in common use (e.g. US 2 letter state codes).
	ResourceType string              `binding:"omitempty" bson:"-" json:"resourceType"`
}

func (out *Age) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Age\"" {
		return fmt.Errorf("resourceType is not %s", "Age")
	}
	if len(asMap["unit"]) > 0 {
		if err := go1.Unmarshal(asMap["unit"], &out.Unit); err != nil {
			return err
		}

	}
	if len(asMap["system"]) > 0 {
		if err := go1.Unmarshal(asMap["system"], &out.System); err != nil {
			return err
		}

	}
	if len(asMap["code"]) > 0 {
		if err := go1.Unmarshal(asMap["code"], &out.Code); err != nil {
			return err
		}

	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
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
	if len(asMap["comparator"]) > 0 {
		if err := go1.Unmarshal(asMap["comparator"], &out.Comparator); err != nil {
			return err
		}

	}
	return nil
}

type Age struct {
	Unit         *string             `bson:",omitempty" json:"unit,omitempty"`       // A human-readable form of the unit.
	System       *string             `bson:",omitempty" json:"system,omitempty"`     // The identification of the system that provides the coded form of the unit.
	Code         *string             `bson:",omitempty" json:"code,omitempty"`       // A computer processable form of the unit in some unit representation system.
	Id           *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`      // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension    []*Extension        `bson:",omitempty" json:"extension,omitempty"`  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Value        *float64            `bson:",omitempty" json:"value,omitempty"`      // The value of the measured amount. The value includes an implicit precision in the presentation of the value.
	Comparator   *string             `bson:",omitempty" json:"comparator,omitempty"` // How the value should be understood and represented - whether the actual value is greater or less than the stated value due to measurement issues; e.g. if the comparator is "<" , then the real value is < stated value.
	ResourceType string              `binding:"omitempty" bson:"-" json:"resourceType"`
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
	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["authorReference"], &out.AuthorReference); err == nil {
	} else if err := go1.Unmarshal(asMap["authorString"], &out.AuthorString); err == nil {
	} else {

	}
	if len(asMap["time"]) > 0 {
		if err := go1.Unmarshal(asMap["time"], &out.Time); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["text"], &out.Text); err != nil {
		return err
	}

	return nil
}

type Annotation struct {
	Id        *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`     // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	AnnotationAuthorx
	Time         *DateTime `bson:",omitempty" json:"time,omitempty"`                    // Indicates when this particular annotation was made.
	Text         *string   `binding:"required" bson:",omitempty" json:"text,omitempty"` // The text of the annotation in markdown format.
	ResourceType string    `binding:"omitempty" bson:"-" json:"resourceType"`
}
type AnnotationAuthorx struct {
	AuthorReference Reference `bson:",omitempty" json:"authorReference,omitempty"`
	AuthorString    string    `bson:",omitempty" json:"authorString,omitempty"`
}

func (out *Attachment) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Attachment\"" {
		return fmt.Errorf("resourceType is not %s", "Attachment")
	}
	if len(asMap["language"]) > 0 {
		if err := go1.Unmarshal(asMap["language"], &out.Language); err != nil {
			return err
		}

	}
	if len(asMap["hash"]) > 0 {
		if err := go1.Unmarshal(asMap["hash"], &out.Hash); err != nil {
			return err
		}

	}
	if len(asMap["width"]) > 0 {
		if err := go1.Unmarshal(asMap["width"], &out.Width); err != nil {
			return err
		}

	}
	if len(asMap["pages"]) > 0 {
		if err := go1.Unmarshal(asMap["pages"], &out.Pages); err != nil {
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
	if len(asMap["height"]) > 0 {
		if err := go1.Unmarshal(asMap["height"], &out.Height); err != nil {
			return err
		}

	}
	if len(asMap["frames"]) > 0 {
		if err := go1.Unmarshal(asMap["frames"], &out.Frames); err != nil {
			return err
		}

	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["contentType"]) > 0 {
		if err := go1.Unmarshal(asMap["contentType"], &out.ContentType); err != nil {
			return err
		}

	}
	if len(asMap["size"]) > 0 {
		if err := go1.Unmarshal(asMap["size"], &out.Size); err != nil {
			return err
		}

	}
	if len(asMap["title"]) > 0 {
		if err := go1.Unmarshal(asMap["title"], &out.Title); err != nil {
			return err
		}

	}
	if len(asMap["creation"]) > 0 {
		if err := go1.Unmarshal(asMap["creation"], &out.Creation); err != nil {
			return err
		}

	}
	if len(asMap["url"]) > 0 {
		if err := go1.Unmarshal(asMap["url"], &out.Url); err != nil {
			return err
		}

	}
	if len(asMap["duration"]) > 0 {
		if err := go1.Unmarshal(asMap["duration"], &out.Duration); err != nil {
			return err
		}

	}
	return nil
}

type Attachment struct {
	Language     *string             `bson:",omitempty" json:"language,omitempty"`    // The human language of the content. The value can be any valid value according to BCP 47.
	Hash         *Base64Binary       `bson:",omitempty" json:"hash,omitempty"`        // The calculated hash of the data using SHA-1. Represented using base64.
	Width        *int                `bson:",omitempty" json:"width,omitempty"`       // Width of the image in pixels (photo/video).
	Pages        *int                `bson:",omitempty" json:"pages,omitempty"`       // The number of pages when printed.
	Extension    []*Extension        `bson:",omitempty" json:"extension,omitempty"`   // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Data         *Base64Binary       `bson:",omitempty" json:"data,omitempty"`        // The actual data of the attachment - a sequence of bytes, base64 encoded.
	Height       *int                `bson:",omitempty" json:"height,omitempty"`      // Height of the image in pixels (photo/video).
	Frames       *int                `bson:",omitempty" json:"frames,omitempty"`      // The number of frames in a photo. This is used with a multi-page fax, or an imaging acquisition context that takes multiple slices in a single image, or an animated gif. If there is more than one frame, this SHALL have a value in order to alert interface software that a multi-frame capable rendering widget is required.
	Id           *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`       // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ContentType  *string             `bson:",omitempty" json:"contentType,omitempty"` // Identifies the type of the data in the attachment and allows a method to be chosen to interpret or render the data. Includes mime type parameters such as charset where appropriate.
	Size         *int64              `bson:",omitempty" json:"size,omitempty"`        // The number of bytes of data that make up this attachment (before base64 encoding, if that is done).
	Title        *string             `bson:",omitempty" json:"title,omitempty"`       // A label or set of text to display in place of the data.
	Creation     *DateTime           `bson:",omitempty" json:"creation,omitempty"`    // The date that the attachment was first created.
	Url          *url.URL            `bson:",omitempty" json:"url,omitempty"`         // A location where the data can be accessed.
	Duration     *float64            `bson:",omitempty" json:"duration,omitempty"`    // The duration of the recording in seconds - for audio and video.
	ResourceType string              `binding:"omitempty" bson:"-" json:"resourceType"`
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
	Id               *primitive.ObjectID           `bson:"_id,omitempty" json:"id,omitempty"`     // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension        []*Extension                  `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	AvailableTime    *AvailabilityAvailableTime    `binding:"omitempty" bson:",omitempty"`
	NotAvailableTime *AvailabilityNotAvailableTime `binding:"omitempty" bson:",omitempty"`
	ResourceType     string                        `binding:"omitempty" bson:"-" json:"resourceType"`
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
	if len(asMap["allDay"]) > 0 {
		if err := go1.Unmarshal(asMap["allDay"], &out.AllDay); err != nil {
			return err
		}

	}
	if len(asMap["availableStartTime"]) > 0 {
		if err := go1.Unmarshal(asMap["availableStartTime"], &out.AvailableStartTime); err != nil {
			return err
		}

	}
	if len(asMap["availableEndTime"]) > 0 {
		if err := go1.Unmarshal(asMap["availableEndTime"], &out.AvailableEndTime); err != nil {
			return err
		}

	}
	return nil
}

type AvailabilityAvailableTime struct {
	Id                 *string      `bson:"_id,omitempty" json:"id,omitempty"`              // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension          []*Extension `bson:",omitempty" json:"extension,omitempty"`          // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	DaysOfWeek         []*string    `bson:",omitempty" json:"daysOfWeek,omitempty"`         // mon | tue | wed | thu | fri | sat | sun.
	AllDay             *bool        `bson:",omitempty" json:"allDay,omitempty"`             // Always available? i.e. 24 hour service.
	AvailableStartTime *Time        `bson:",omitempty" json:"availableStartTime,omitempty"` // Opening time of day (ignored if allDay = true).
	AvailableEndTime   *Time        `bson:",omitempty" json:"availableEndTime,omitempty"`   // Closing time of day (ignored if allDay = true).
}

func (out *AvailabilityNotAvailableTime) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if len(asMap["description"]) > 0 {
		if err := go1.Unmarshal(asMap["description"], &out.Description); err != nil {
			return err
		}

	}
	if len(asMap["during"]) > 0 {
		if err := go1.Unmarshal(asMap["during"], &out.During); err != nil {
			return err
		}

	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
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

type AvailabilityNotAvailableTime struct {
	Description *string      `bson:",omitempty" json:"description,omitempty"` // Reason presented to the user explaining why time not available.
	During      *Period      `bson:",omitempty" json:"during,omitempty"`      // Service not available during this period.
	Id          *string      `bson:"_id,omitempty" json:"id,omitempty"`       // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension   []*Extension `bson:",omitempty" json:"extension,omitempty"`   // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
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
	Id                *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`             // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension         []*Extension        `bson:",omitempty" json:"extension,omitempty"`         // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	ModifierExtension []*Extension        `bson:",omitempty" json:"modifierExtension,omitempty"` /*
	May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.

	Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	*/
	ResourceType string `binding:"omitempty" bson:"-" json:"resourceType"`
}

func (out *CodeableConcept) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"CodeableConcept\"" {
		return fmt.Errorf("resourceType is not %s", "CodeableConcept")
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
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
	if len(asMap["text"]) > 0 {
		if err := go1.Unmarshal(asMap["text"], &out.Text); err != nil {
			return err
		}

	}
	return nil
}

type CodeableConcept struct {
	Id           *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`     // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension    []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Coding       []Coding            `bson:",omitempty" json:"coding,omitempty"`    // A reference to a code defined by a terminology system.
	Text         *string             `bson:",omitempty" json:"text,omitempty"`      // A human language representation of the concept as seen/selected/uttered by the user who entered the data and/or which represents the intended meaning of the user.
	ResourceType string              `binding:"omitempty" bson:"-" json:"resourceType"`
}

func (out *CodeableReference) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"CodeableReference\"" {
		return fmt.Errorf("resourceType is not %s", "CodeableReference")
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

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
	return nil
}

type CodeableReference struct {
	Id           *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`     // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension    []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Concept      *CodeableConcept    `bson:",omitempty" json:"concept,omitempty"`   // A reference to a concept - e.g. the information is identified by its general class to the degree of precision found in the terminology.
	Reference    *Reference          `bson:",omitempty" json:"reference,omitempty"` // A reference to a resource the provides exact details about the information being referenced.
	ResourceType string              `binding:"omitempty" bson:"-" json:"resourceType"`
}

func (out *Coding) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Coding\"" {
		return fmt.Errorf("resourceType is not %s", "Coding")
	}
	if len(asMap["userSelected"]) > 0 {
		if err := go1.Unmarshal(asMap["userSelected"], &out.UserSelected); err != nil {
			return err
		}

	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
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
	if len(asMap["version"]) > 0 {
		if err := go1.Unmarshal(asMap["version"], &out.Version); err != nil {
			return err
		}

	}
	if len(asMap["code"]) > 0 {
		if err := go1.Unmarshal(asMap["code"], &out.Code); err != nil {
			return err
		}

	}
	if len(asMap["display"]) > 0 {
		if err := go1.Unmarshal(asMap["display"], &out.Display); err != nil {
			return err
		}

	}
	return nil
}

type Coding struct {
	UserSelected *bool                `bson:",omitempty" json:"userSelected,omitempty"` // Indicates that this coding was chosen by a user directly - e.g. off a pick list of available items (codes or displays).
	Id           **primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`        // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension    []*Extension         `bson:",omitempty" json:"extension,omitempty"`    // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	System       *string              `bson:",omitempty" json:"system,omitempty"`       // The identification of the code system that defines the meaning of the symbol in the code.
	Version      *string              `bson:",omitempty" json:"version,omitempty"`      // The version of the code system which was used when choosing this code. Note that a well-maintained code system does not need the version reported, because the meaning of codes is consistent across versions. However this cannot consistently be assured, and when the meaning is not guaranteed to be consistent, the version SHOULD be exchanged.
	Code         *string              `bson:",omitempty" json:"code,omitempty"`         // A symbol in syntax defined by the system. The symbol may be a predefined code or an expression in a syntax defined by the coding system (e.g. post-coordination).
	Display      *string              `bson:",omitempty" json:"display,omitempty"`      // A representation of the meaning of the code in the system, following the rules of the system.
	ResourceType string               `binding:"omitempty" bson:"-" json:"resourceType"`
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
	if len(asMap["telecom"]) > 0 {
		if err := go1.Unmarshal(asMap["telecom"], &out.Telecom); err != nil {
			return err
		}

	}
	return nil
}

type ContactDetail struct {
	Id           *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`     // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension    []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Name         *string             `bson:",omitempty" json:"name,omitempty"`      // The name of an individual to contact.
	Telecom      []ContactPoint      `bson:",omitempty" json:"telecom,omitempty"`   // The contact details for the individual (if a name was provided) or the organization.
	ResourceType string              `binding:"omitempty" bson:"-" json:"resourceType"`
}

func (out *ContactPoint) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"ContactPoint\"" {
		return fmt.Errorf("resourceType is not %s", "ContactPoint")
	}
	if len(asMap["value"]) > 0 {
		if err := go1.Unmarshal(asMap["value"], &out.Value); err != nil {
			return err
		}

	}
	if len(asMap["use"]) > 0 {
		if err := go1.Unmarshal(asMap["use"], &out.Use); err != nil {
			return err
		}

	}
	if len(asMap["rank"]) > 0 {
		if err := go1.Unmarshal(asMap["rank"], &out.Rank); err != nil {
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
	return nil
}

type ContactPoint struct {
	Value        *string              `bson:",omitempty" json:"value,omitempty"`     // The actual contact point details, in a form that is meaningful to the designated communication system (i.e. phone number or email address).
	Use          *string              `bson:",omitempty" json:"use,omitempty"`       // Identifies the purpose for the contact point.
	Rank         *int                 `bson:",omitempty" json:"rank,omitempty"`      // Specifies a preferred order in which to use a set of contacts. ContactPoints with lower rank values are more preferred than those with higher rank values.
	Period       *Period              `bson:",omitempty" json:"period,omitempty"`    // Time period when the contact point was/is in use.
	Id           **primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`     // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension    []*Extension         `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	System       *string              `bson:",omitempty" json:"system,omitempty"`    // Telecommunications form for contact point - what communications system is required to make use of the contact.
	ResourceType string               `binding:"omitempty" bson:"-" json:"resourceType"`
}

func (out *Contributor) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Contributor\"" {
		return fmt.Errorf("resourceType is not %s", "Contributor")
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
	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["type"], &out.Type); err != nil {
		return err
	}

	if err := go1.Unmarshal(asMap["name"], &out.Name); err != nil {
		return err
	}

	return nil
}

type Contributor struct {
	Contact      []*ContactDetail    `bson:",omitempty" json:"contact,omitempty"`                 // Contact details to assist a user in finding and communicating with the contributor.
	Id           *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`                   // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension    []*Extension        `bson:",omitempty" json:"extension,omitempty"`               // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Type         *string             `binding:"required" bson:",omitempty" json:"type,omitempty"` // The type of contributor.
	Name         *string             `binding:"required" bson:",omitempty" json:"name,omitempty"` // The name of the individual or organization responsible for the contribution.
	ResourceType string              `binding:"omitempty" bson:"-" json:"resourceType"`
}

func (out *Count) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Count\"" {
		return fmt.Errorf("resourceType is not %s", "Count")
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
	if len(asMap["comparator"]) > 0 {
		if err := go1.Unmarshal(asMap["comparator"], &out.Comparator); err != nil {
			return err
		}

	}
	if len(asMap["unit"]) > 0 {
		if err := go1.Unmarshal(asMap["unit"], &out.Unit); err != nil {
			return err
		}

	}
	if len(asMap["system"]) > 0 {
		if err := go1.Unmarshal(asMap["system"], &out.System); err != nil {
			return err
		}

	}
	if len(asMap["code"]) > 0 {
		if err := go1.Unmarshal(asMap["code"], &out.Code); err != nil {
			return err
		}

	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	return nil
}

type Count struct {
	Extension    []*Extension        `bson:",omitempty" json:"extension,omitempty"`  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Value        *float64            `bson:",omitempty" json:"value,omitempty"`      // The value of the measured amount. The value includes an implicit precision in the presentation of the value.
	Comparator   *string             `bson:",omitempty" json:"comparator,omitempty"` // How the value should be understood and represented - whether the actual value is greater or less than the stated value due to measurement issues; e.g. if the comparator is "<" , then the real value is < stated value.
	Unit         *string             `bson:",omitempty" json:"unit,omitempty"`       // A human-readable form of the unit.
	System       *string             `bson:",omitempty" json:"system,omitempty"`     // The identification of the system that provides the coded form of the unit.
	Code         *string             `bson:",omitempty" json:"code,omitempty"`       // A computer processable form of the unit in some unit representation system.
	Id           *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`      // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ResourceType string              `binding:"omitempty" bson:"-" json:"resourceType"`
}

func (out *DataRequirement) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"DataRequirement\"" {
		return fmt.Errorf("resourceType is not %s", "DataRequirement")
	}
	if len(asMap["limit"]) > 0 {
		if err := go1.Unmarshal(asMap["limit"], &out.Limit); err != nil {
			return err
		}

	}
	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["subjectCodeableConcept"], &out.SubjectCodeableConcept); err == nil {
	} else if err := go1.Unmarshal(asMap["subjectReference"], &out.SubjectReference); err == nil {
	} else {

	}
	if len(asMap["codeFilter"]) > 0 {
		if err := go1.Unmarshal(asMap["codeFilter"], &out.CodeFilter); err != nil {
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
	if err := go1.Unmarshal(asMap["type"], &out.Type); err != nil {
		return err
	}

	if len(asMap["profile"]) > 0 {
		if err := go1.Unmarshal(asMap["profile"], &out.Profile); err != nil {
			return err
		}

	}
	if len(asMap["mustSupport"]) > 0 {
		if err := go1.Unmarshal(asMap["mustSupport"], &out.MustSupport); err != nil {
			return err
		}

	}
	return nil
}

type DataRequirement struct {
	Limit     *int         `bson:",omitempty" json:"limit,omitempty"`     // Specifies a maximum number of results that are required (uses the _count search parameter).
	Extension []*Extension `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	DataRequirementSubjectx
	CodeFilter  *DataRequirementCodeFilter  `binding:"omitempty" bson:",omitempty"`
	DateFilter  *DataRequirementDateFilter  `binding:"omitempty" bson:",omitempty"`
	ValueFilter *DataRequirementValueFilter `binding:"omitempty" bson:",omitempty"`
	Sort        *DataRequirementSort        `binding:"omitempty" bson:",omitempty"`
	Id          *primitive.ObjectID         `bson:"_id,omitempty" json:"id,omitempty"`                   // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Type        *string                     `binding:"required" bson:",omitempty" json:"type,omitempty"` // The type of the required data, specified as the type name of a resource. For profiles, this value is set to the type of the base resource of the profile.
	Profile     []*string                   `bson:",omitempty" json:"profile,omitempty"`                 // The profile of the required data, specified as the uri of the profile definition.
	MustSupport []*string                   `bson:",omitempty" json:"mustSupport,omitempty"`             /*
	Indicates that specific elements of the type are referenced by the knowledge module and must be supported by the consumer in order to obtain an effective evaluation. This does not mean that a value is required for this element, only that the consuming system must understand the element and be able to provide values for it if they are available.

	The value of mustSupport SHALL be a FHIRPath resolvable on the type of the DataRequirement. The path SHALL consist only of identifiers, constant indexers, and .resolve() (see the [Simple FHIRPath Profile](fhirpath.html#simple) for full details).
	*/
	ResourceType string `binding:"omitempty" bson:"-" json:"resourceType"`
}
type DataRequirementSubjectx struct {
	SubjectCodeableConcept CodeableConcept `bson:",omitempty" json:"subjectCodeableConcept,omitempty"`
	SubjectReference       Reference       `bson:",omitempty" json:"subjectReference,omitempty"`
}

func (out *DataRequirementCodeFilter) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if len(asMap["searchParam"]) > 0 {
		if err := go1.Unmarshal(asMap["searchParam"], &out.SearchParam); err != nil {
			return err
		}

	}
	if len(asMap["valueSet"]) > 0 {
		if err := go1.Unmarshal(asMap["valueSet"], &out.ValueSet); err != nil {
			return err
		}

	}
	if len(asMap["code"]) > 0 {
		if err := go1.Unmarshal(asMap["code"], &out.Code); err != nil {
			return err
		}

	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
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
	return nil
}

type DataRequirementCodeFilter struct {
	SearchParam *string      `bson:",omitempty" json:"searchParam,omitempty"` // A token parameter that refers to a search parameter defined on the specified type of the DataRequirement, and which searches on elements of type code, Coding, or CodeableConcept.
	ValueSet    *string      `bson:",omitempty" json:"valueSet,omitempty"`    // The valueset for the code filter. The valueSet and code elements are additive. If valueSet is specified, the filter will return only those data items for which the value of the code-valued element specified in the path is a member of the specified valueset.
	Code        []*Coding    `bson:",omitempty" json:"code,omitempty"`        // The codes for the code filter. If values are given, the filter will return only those data items for which the code-valued attribute specified by the path has a value that is one of the specified codes. If codes are specified in addition to a value set, the filter returns items matching a code in the value set or one of the specified codes.
	Id          *string      `bson:"_id,omitempty" json:"id,omitempty"`       // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension   []*Extension `bson:",omitempty" json:"extension,omitempty"`   // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Path        *string      `bson:",omitempty" json:"path,omitempty"`        // The code-valued attribute of the filter. The specified path SHALL be a FHIRPath resolvable on the specified type of the DataRequirement, and SHALL consist only of identifiers, constant indexers, and .resolve(). The path is allowed to contain qualifiers (.) to traverse sub-elements, as well as indexers ([x]) to traverse multiple-cardinality sub-elements (see the [Simple FHIRPath Profile](fhirpath.html#simple) for full details). Note that the index must be an integer constant. The path must resolve to an element of type code, Coding, or CodeableConcept.
}

func (out *DataRequirementDateFilter) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
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
	if len(asMap["searchParam"]) > 0 {
		if err := go1.Unmarshal(asMap["searchParam"], &out.SearchParam); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["valueDateTime"], &out.ValueDateTime); err == nil {
	} else if err := go1.Unmarshal(asMap["valuePeriod"], &out.ValuePeriod); err == nil {
	} else if err := go1.Unmarshal(asMap["valueDuration"], &out.ValueDuration); err == nil {
	} else {

	}
	return nil
}

type DataRequirementDateFilter struct {
	Id          *string      `bson:"_id,omitempty" json:"id,omitempty"`       // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension   []*Extension `bson:",omitempty" json:"extension,omitempty"`   // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Path        *string      `bson:",omitempty" json:"path,omitempty"`        // The date-valued attribute of the filter. The specified path SHALL be a FHIRPath resolvable on the specified type of the DataRequirement, and SHALL consist only of identifiers, constant indexers, and .resolve(). The path is allowed to contain qualifiers (.) to traverse sub-elements, as well as indexers ([x]) to traverse multiple-cardinality sub-elements (see the [Simple FHIRPath Profile](fhirpath.html#simple) for full details). Note that the index must be an integer constant. The path must resolve to an element of type date, dateTime, Period, Schedule, or Timing.
	SearchParam *string      `bson:",omitempty" json:"searchParam,omitempty"` // A date parameter that refers to a search parameter defined on the specified type of the DataRequirement, and which searches on elements of type date, dateTime, Period, Schedule, or Timing.
	DataRequirementDateFilterValuex
}
type DataRequirementDateFilterValuex struct {
	ValueDateTime DateTime `bson:",omitempty" json:"valueDateTime,omitempty"`
	ValuePeriod   Period   `bson:",omitempty" json:"valuePeriod,omitempty"`
	ValueDuration Duration `bson:",omitempty" json:"valueDuration,omitempty"`
}

func (out *DataRequirementValueFilter) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if len(asMap["comparator"]) > 0 {
		if err := go1.Unmarshal(asMap["comparator"], &out.Comparator); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["valueDateTime"], &out.ValueDateTime); err == nil {
	} else if err := go1.Unmarshal(asMap["valuePeriod"], &out.ValuePeriod); err == nil {
	} else if err := go1.Unmarshal(asMap["valueDuration"], &out.ValueDuration); err == nil {
	} else {

	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
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
	if len(asMap["searchParam"]) > 0 {
		if err := go1.Unmarshal(asMap["searchParam"], &out.SearchParam); err != nil {
			return err
		}

	}
	return nil
}

type DataRequirementValueFilter struct {
	Comparator *string `bson:",omitempty" json:"comparator,omitempty"` // The comparator to be used to determine whether the value is matching.
	DataRequirementValueFilterValuex
	Id          *string      `bson:"_id,omitempty" json:"id,omitempty"`       // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension   []*Extension `bson:",omitempty" json:"extension,omitempty"`   // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Path        *string      `bson:",omitempty" json:"path,omitempty"`        // The attribute of the filter. The specified path SHALL be a FHIRPath resolvable on the specified type of the DataRequirement, and SHALL consist only of identifiers, constant indexers, and .resolve(). The path is allowed to contain qualifiers (.) to traverse sub-elements, as well as indexers ([x]) to traverse multiple-cardinality sub-elements (see the [Simple FHIRPath Profile](fhirpath.html#simple) for full details). Note that the index must be an integer constant. The path must resolve to an element of a type that is comparable to the valueFilter.value[x] element for the filter.
	SearchParam *string      `bson:",omitempty" json:"searchParam,omitempty"` // A search parameter defined on the specified type of the DataRequirement, and which searches on elements of a type compatible with the type of the valueFilter.value[x] for the filter.
}
type DataRequirementValueFilterValuex struct {
	ValueDateTime DateTime `bson:",omitempty" json:"valueDateTime,omitempty"`
	ValuePeriod   Period   `bson:",omitempty" json:"valuePeriod,omitempty"`
	ValueDuration Duration `bson:",omitempty" json:"valueDuration,omitempty"`
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
	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["path"], &out.Path); err != nil {
		return err
	}

	if err := go1.Unmarshal(asMap["direction"], &out.Direction); err != nil {
		return err
	}

	return nil
}

type DataRequirementSort struct {
	Id        *string      `bson:"_id,omitempty" json:"id,omitempty"`                        // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension []*Extension `bson:",omitempty" json:"extension,omitempty"`                    // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Path      *string      `binding:"required" bson:",omitempty" json:"path,omitempty"`      // The attribute of the sort. The specified path must be resolvable from the type of the required data. The path is allowed to contain qualifiers (.) to traverse sub-elements, as well as indexers ([x]) to traverse multiple-cardinality sub-elements. Note that the index must be an integer constant.
	Direction *string      `binding:"required" bson:",omitempty" json:"direction,omitempty"` // The direction of the sort, ascending or descending.
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
	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	return nil
}

type DataType struct {
	Id           *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`     // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension    []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	ResourceType string              `binding:"omitempty" bson:"-" json:"resourceType"`
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
	if len(asMap["code"]) > 0 {
		if err := go1.Unmarshal(asMap["code"], &out.Code); err != nil {
			return err
		}

	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
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
	if len(asMap["comparator"]) > 0 {
		if err := go1.Unmarshal(asMap["comparator"], &out.Comparator); err != nil {
			return err
		}

	}
	if len(asMap["unit"]) > 0 {
		if err := go1.Unmarshal(asMap["unit"], &out.Unit); err != nil {
			return err
		}

	}
	return nil
}

type Distance struct {
	System       *string             `bson:",omitempty" json:"system,omitempty"`     // The identification of the system that provides the coded form of the unit.
	Code         *string             `bson:",omitempty" json:"code,omitempty"`       // A computer processable form of the unit in some unit representation system.
	Id           *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`      // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension    []*Extension        `bson:",omitempty" json:"extension,omitempty"`  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Value        *float64            `bson:",omitempty" json:"value,omitempty"`      // The value of the measured amount. The value includes an implicit precision in the presentation of the value.
	Comparator   *string             `bson:",omitempty" json:"comparator,omitempty"` // How the value should be understood and represented - whether the actual value is greater or less than the stated value due to measurement issues; e.g. if the comparator is "<" , then the real value is < stated value.
	Unit         *string             `bson:",omitempty" json:"unit,omitempty"`       // A human-readable form of the unit.
	ResourceType string              `binding:"omitempty" bson:"-" json:"resourceType"`
}

func (out *Dosage) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Dosage\"" {
		return fmt.Errorf("resourceType is not %s", "Dosage")
	}
	if len(asMap["maxDosePerPeriod"]) > 0 {
		if err := go1.Unmarshal(asMap["maxDosePerPeriod"], &out.MaxDosePerPeriod); err != nil {
			return err
		}

	}
	if len(asMap["sequence"]) > 0 {
		if err := go1.Unmarshal(asMap["sequence"], &out.Sequence); err != nil {
			return err
		}

	}
	if len(asMap["text"]) > 0 {
		if err := go1.Unmarshal(asMap["text"], &out.Text); err != nil {
			return err
		}

	}
	if len(asMap["doseAndRate"]) > 0 {
		if err := go1.Unmarshal(asMap["doseAndRate"], &out.DoseAndRate); err != nil {
			return err
		}

	}
	if len(asMap["maxDosePerLifetime"]) > 0 {
		if err := go1.Unmarshal(asMap["maxDosePerLifetime"], &out.MaxDosePerLifetime); err != nil {
			return err
		}

	}
	if len(asMap["additionalInstruction"]) > 0 {
		if err := go1.Unmarshal(asMap["additionalInstruction"], &out.AdditionalInstruction); err != nil {
			return err
		}

	}
	if len(asMap["patientInstruction"]) > 0 {
		if err := go1.Unmarshal(asMap["patientInstruction"], &out.PatientInstruction); err != nil {
			return err
		}

	}
	if len(asMap["method"]) > 0 {
		if err := go1.Unmarshal(asMap["method"], &out.Method); err != nil {
			return err
		}

	}
	if len(asMap["modifierExtension"]) > 0 {
		if err := go1.Unmarshal(asMap["modifierExtension"], &out.ModifierExtension); err != nil {
			return err
		}

	}
	if len(asMap["site"]) > 0 {
		if err := go1.Unmarshal(asMap["site"], &out.Site); err != nil {
			return err
		}

	}
	if len(asMap["maxDosePerAdministration"]) > 0 {
		if err := go1.Unmarshal(asMap["maxDosePerAdministration"], &out.MaxDosePerAdministration); err != nil {
			return err
		}

	}
	if len(asMap["asNeeded"]) > 0 {
		if err := go1.Unmarshal(asMap["asNeeded"], &out.AsNeeded); err != nil {
			return err
		}

	}
	if len(asMap["asNeededFor"]) > 0 {
		if err := go1.Unmarshal(asMap["asNeededFor"], &out.AsNeededFor); err != nil {
			return err
		}

	}
	if len(asMap["route"]) > 0 {
		if err := go1.Unmarshal(asMap["route"], &out.Route); err != nil {
			return err
		}

	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
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
	MaxDosePerPeriod      []*Ratio           `bson:",omitempty" json:"maxDosePerPeriod,omitempty"` // Upper limit on medication per unit of time.
	Sequence              *int               `bson:",omitempty" json:"sequence,omitempty"`         // Indicates the order in which the dosage instructions should be applied or interpreted.
	Text                  *string            `bson:",omitempty" json:"text,omitempty"`             // Free text dosage instructions e.g. SIG.
	DoseAndRate           *DosageDoseAndRate `binding:"omitempty" bson:",omitempty"`
	MaxDosePerLifetime    *Quantity          `bson:",omitempty" json:"maxDosePerLifetime,omitempty"`    // Upper limit on medication per lifetime of the patient.
	AdditionalInstruction []*CodeableConcept `bson:",omitempty" json:"additionalInstruction,omitempty"` // Supplemental instructions to the patient on how to take the medication  (e.g. "with meals" or"take half to one hour before food") or warnings for the patient about the medication (e.g. "may cause drowsiness" or "avoid exposure of skin to direct sunlight or sunlamps").
	PatientInstruction    *string            `bson:",omitempty" json:"patientInstruction,omitempty"`    // Instructions in terms that are understood by the patient or consumer.
	Method                *CodeableConcept   `bson:",omitempty" json:"method,omitempty"`                // Technique for administering medication.
	ModifierExtension     []*Extension       `bson:",omitempty" json:"modifierExtension,omitempty"`     /*
	May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.

	Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	*/
	Site                     *CodeableConcept    `bson:",omitempty" json:"site,omitempty"`                     // Body site to administer to.
	MaxDosePerAdministration *Quantity           `bson:",omitempty" json:"maxDosePerAdministration,omitempty"` // Upper limit on medication per administration.
	AsNeeded                 *bool               `bson:",omitempty" json:"asNeeded,omitempty"`                 // Indicates whether the Medication is only taken when needed within a specific dosing schedule (Boolean option).
	AsNeededFor              []*CodeableConcept  `bson:",omitempty" json:"asNeededFor,omitempty"`              // Indicates whether the Medication is only taken based on a precondition for taking the Medication (CodeableConcept).
	Route                    *CodeableConcept    `bson:",omitempty" json:"route,omitempty"`                    // How drug should enter body.
	Id                       *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`                    // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension                []*Extension        `bson:",omitempty" json:"extension,omitempty"`                // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Timing                   Timing              `bson:",omitempty" json:"timing,omitempty"`                   // When medication should be administered.
	ResourceType             string              `binding:"omitempty" bson:"-" json:"resourceType"`
}

func (out *DosageDoseAndRate) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if err := go1.Unmarshal(asMap["doseRange"], &out.DoseRange); err == nil {
	} else if err := go1.Unmarshal(asMap["doseQuantity"], &out.DoseQuantity); err == nil {
	} else {

	}
	if err := go1.Unmarshal(asMap["rateRatio"], &out.RateRatio); err == nil {
	} else if err := go1.Unmarshal(asMap["rateRange"], &out.RateRange); err == nil {
	} else if err := go1.Unmarshal(asMap["rateQuantity"], &out.RateQuantity); err == nil {
	} else {

	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
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
	return nil
}

type DosageDoseAndRate struct {
	DosageDoseAndRateDosex
	DosageDoseAndRateRatex
	Id        *string          `bson:"_id,omitempty" json:"id,omitempty"`     // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension []*Extension     `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Type      *CodeableConcept `bson:",omitempty" json:"type,omitempty"`      // The kind of dose or rate specified, for example, ordered or calculated.
}
type DosageDoseAndRateDosex struct {
	DoseRange    Range    `bson:",omitempty" json:"doseRange,omitempty"`
	DoseQuantity Quantity `bson:",omitempty" json:"doseQuantity,omitempty"`
}
type DosageDoseAndRateRatex struct {
	RateRatio    Ratio    `bson:",omitempty" json:"rateRatio,omitempty"`
	RateRange    Range    `bson:",omitempty" json:"rateRange,omitempty"`
	RateQuantity Quantity `bson:",omitempty" json:"rateQuantity,omitempty"`
}

func (out *Duration) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Duration\"" {
		return fmt.Errorf("resourceType is not %s", "Duration")
	}
	if len(asMap["code"]) > 0 {
		if err := go1.Unmarshal(asMap["code"], &out.Code); err != nil {
			return err
		}

	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
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
	if len(asMap["comparator"]) > 0 {
		if err := go1.Unmarshal(asMap["comparator"], &out.Comparator); err != nil {
			return err
		}

	}
	if len(asMap["unit"]) > 0 {
		if err := go1.Unmarshal(asMap["unit"], &out.Unit); err != nil {
			return err
		}

	}
	if len(asMap["system"]) > 0 {
		if err := go1.Unmarshal(asMap["system"], &out.System); err != nil {
			return err
		}

	}
	return nil
}

type Duration struct {
	Code         *string              `bson:",omitempty" json:"code,omitempty"`       // A computer processable form of the unit in some unit representation system.
	Id           **primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`      // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension    []*Extension         `bson:",omitempty" json:"extension,omitempty"`  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Value        *float64             `bson:",omitempty" json:"value,omitempty"`      // The value of the measured amount. The value includes an implicit precision in the presentation of the value.
	Comparator   *string              `bson:",omitempty" json:"comparator,omitempty"` // How the value should be understood and represented - whether the actual value is greater or less than the stated value due to measurement issues; e.g. if the comparator is "<" , then the real value is < stated value.
	Unit         *string              `bson:",omitempty" json:"unit,omitempty"`       // A human-readable form of the unit.
	System       *string              `bson:",omitempty" json:"system,omitempty"`     // The identification of the system that provides the coded form of the unit.
	ResourceType string               `binding:"omitempty" bson:"-" json:"resourceType"`
}

func (out *ElementDefinition) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"ElementDefinition\"" {
		return fmt.Errorf("resourceType is not %s", "ElementDefinition")
	}
	if len(asMap["comment"]) > 0 {
		if err := go1.Unmarshal(asMap["comment"], &out.Comment); err != nil {
			return err
		}

	}
	if len(asMap["max"]) > 0 {
		if err := go1.Unmarshal(asMap["max"], &out.Max); err != nil {
			return err
		}

	}
	if len(asMap["contentReference"]) > 0 {
		if err := go1.Unmarshal(asMap["contentReference"], &out.ContentReference); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["fixedBase64Binary"], &out.FixedBase64Binary); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedBoolean"], &out.FixedBoolean); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedCanonical"], &out.FixedCanonical); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedCode"], &out.FixedCode); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedDate"], &out.FixedDate); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedDateTime"], &out.FixedDateTime); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedDecimal"], &out.FixedDecimal); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedId"], &out.FixedId); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedInstant"], &out.FixedInstant); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedInteger"], &out.FixedInteger); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedInteger64"], &out.FixedInteger64); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedMarkdown"], &out.FixedMarkdown); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedOid"], &out.FixedOid); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedPositiveInt"], &out.FixedPositiveInt); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedString"], &out.FixedString); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedTime"], &out.FixedTime); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedUnsignedInt"], &out.FixedUnsignedInt); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedUri"], &out.FixedUri); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedUrl"], &out.FixedUrl); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedUuid"], &out.FixedUuid); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedAddress"], &out.FixedAddress); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedAge"], &out.FixedAge); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedAnnotation"], &out.FixedAnnotation); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedAttachment"], &out.FixedAttachment); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedCodeableConcept"], &out.FixedCodeableConcept); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedCodeableReference"], &out.FixedCodeableReference); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedCoding"], &out.FixedCoding); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedContactPoint"], &out.FixedContactPoint); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedCount"], &out.FixedCount); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedDistance"], &out.FixedDistance); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedDuration"], &out.FixedDuration); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedHumanName"], &out.FixedHumanName); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedIdentifier"], &out.FixedIdentifier); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedMoney"], &out.FixedMoney); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedPeriod"], &out.FixedPeriod); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedQuantity"], &out.FixedQuantity); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedRange"], &out.FixedRange); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedRatio"], &out.FixedRatio); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedRatioRange"], &out.FixedRatioRange); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedReference"], &out.FixedReference); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedSampledData"], &out.FixedSampledData); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedSignature"], &out.FixedSignature); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedTiming"], &out.FixedTiming); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedContactDetail"], &out.FixedContactDetail); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedDataRequirement"], &out.FixedDataRequirement); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedExpression"], &out.FixedExpression); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedParameterDefinition"], &out.FixedParameterDefinition); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedRelatedArtifact"], &out.FixedRelatedArtifact); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedTriggerDefinition"], &out.FixedTriggerDefinition); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedUsageContext"], &out.FixedUsageContext); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedAvailability"], &out.FixedAvailability); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedExtendedContactDetail"], &out.FixedExtendedContactDetail); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedDosage"], &out.FixedDosage); err == nil {
	} else if err := go1.Unmarshal(asMap["fixedMeta"], &out.FixedMeta); err == nil {
	} else {

	}
	if len(asMap["isModifierReason"]) > 0 {
		if err := go1.Unmarshal(asMap["isModifierReason"], &out.IsModifierReason); err != nil {
			return err
		}

	}
	if len(asMap["mapping"]) > 0 {
		if err := go1.Unmarshal(asMap["mapping"], &out.Mapping); err != nil {
			return err
		}

	}
	if len(asMap["representation"]) > 0 {
		if err := go1.Unmarshal(asMap["representation"], &out.Representation); err != nil {
			return err
		}

	}
	if len(asMap["sliceIsConstraining"]) > 0 {
		if err := go1.Unmarshal(asMap["sliceIsConstraining"], &out.SliceIsConstraining); err != nil {
			return err
		}

	}
	if len(asMap["code"]) > 0 {
		if err := go1.Unmarshal(asMap["code"], &out.Code); err != nil {
			return err
		}

	}
	if len(asMap["meaningWhenMissing"]) > 0 {
		if err := go1.Unmarshal(asMap["meaningWhenMissing"], &out.MeaningWhenMissing); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["patternBase64Binary"], &out.PatternBase64Binary); err == nil {
	} else if err := go1.Unmarshal(asMap["patternBoolean"], &out.PatternBoolean); err == nil {
	} else if err := go1.Unmarshal(asMap["patternCanonical"], &out.PatternCanonical); err == nil {
	} else if err := go1.Unmarshal(asMap["patternCode"], &out.PatternCode); err == nil {
	} else if err := go1.Unmarshal(asMap["patternDate"], &out.PatternDate); err == nil {
	} else if err := go1.Unmarshal(asMap["patternDateTime"], &out.PatternDateTime); err == nil {
	} else if err := go1.Unmarshal(asMap["patternDecimal"], &out.PatternDecimal); err == nil {
	} else if err := go1.Unmarshal(asMap["patternId"], &out.PatternId); err == nil {
	} else if err := go1.Unmarshal(asMap["patternInstant"], &out.PatternInstant); err == nil {
	} else if err := go1.Unmarshal(asMap["patternInteger"], &out.PatternInteger); err == nil {
	} else if err := go1.Unmarshal(asMap["patternInteger64"], &out.PatternInteger64); err == nil {
	} else if err := go1.Unmarshal(asMap["patternMarkdown"], &out.PatternMarkdown); err == nil {
	} else if err := go1.Unmarshal(asMap["patternOid"], &out.PatternOid); err == nil {
	} else if err := go1.Unmarshal(asMap["patternPositiveInt"], &out.PatternPositiveInt); err == nil {
	} else if err := go1.Unmarshal(asMap["patternString"], &out.PatternString); err == nil {
	} else if err := go1.Unmarshal(asMap["patternTime"], &out.PatternTime); err == nil {
	} else if err := go1.Unmarshal(asMap["patternUnsignedInt"], &out.PatternUnsignedInt); err == nil {
	} else if err := go1.Unmarshal(asMap["patternUri"], &out.PatternUri); err == nil {
	} else if err := go1.Unmarshal(asMap["patternUrl"], &out.PatternUrl); err == nil {
	} else if err := go1.Unmarshal(asMap["patternUuid"], &out.PatternUuid); err == nil {
	} else if err := go1.Unmarshal(asMap["patternAddress"], &out.PatternAddress); err == nil {
	} else if err := go1.Unmarshal(asMap["patternAge"], &out.PatternAge); err == nil {
	} else if err := go1.Unmarshal(asMap["patternAnnotation"], &out.PatternAnnotation); err == nil {
	} else if err := go1.Unmarshal(asMap["patternAttachment"], &out.PatternAttachment); err == nil {
	} else if err := go1.Unmarshal(asMap["patternCodeableConcept"], &out.PatternCodeableConcept); err == nil {
	} else if err := go1.Unmarshal(asMap["patternCodeableReference"], &out.PatternCodeableReference); err == nil {
	} else if err := go1.Unmarshal(asMap["patternCoding"], &out.PatternCoding); err == nil {
	} else if err := go1.Unmarshal(asMap["patternContactPoint"], &out.PatternContactPoint); err == nil {
	} else if err := go1.Unmarshal(asMap["patternCount"], &out.PatternCount); err == nil {
	} else if err := go1.Unmarshal(asMap["patternDistance"], &out.PatternDistance); err == nil {
	} else if err := go1.Unmarshal(asMap["patternDuration"], &out.PatternDuration); err == nil {
	} else if err := go1.Unmarshal(asMap["patternHumanName"], &out.PatternHumanName); err == nil {
	} else if err := go1.Unmarshal(asMap["patternIdentifier"], &out.PatternIdentifier); err == nil {
	} else if err := go1.Unmarshal(asMap["patternMoney"], &out.PatternMoney); err == nil {
	} else if err := go1.Unmarshal(asMap["patternPeriod"], &out.PatternPeriod); err == nil {
	} else if err := go1.Unmarshal(asMap["patternQuantity"], &out.PatternQuantity); err == nil {
	} else if err := go1.Unmarshal(asMap["patternRange"], &out.PatternRange); err == nil {
	} else if err := go1.Unmarshal(asMap["patternRatio"], &out.PatternRatio); err == nil {
	} else if err := go1.Unmarshal(asMap["patternRatioRange"], &out.PatternRatioRange); err == nil {
	} else if err := go1.Unmarshal(asMap["patternReference"], &out.PatternReference); err == nil {
	} else if err := go1.Unmarshal(asMap["patternSampledData"], &out.PatternSampledData); err == nil {
	} else if err := go1.Unmarshal(asMap["patternSignature"], &out.PatternSignature); err == nil {
	} else if err := go1.Unmarshal(asMap["patternTiming"], &out.PatternTiming); err == nil {
	} else if err := go1.Unmarshal(asMap["patternContactDetail"], &out.PatternContactDetail); err == nil {
	} else if err := go1.Unmarshal(asMap["patternDataRequirement"], &out.PatternDataRequirement); err == nil {
	} else if err := go1.Unmarshal(asMap["patternExpression"], &out.PatternExpression); err == nil {
	} else if err := go1.Unmarshal(asMap["patternParameterDefinition"], &out.PatternParameterDefinition); err == nil {
	} else if err := go1.Unmarshal(asMap["patternRelatedArtifact"], &out.PatternRelatedArtifact); err == nil {
	} else if err := go1.Unmarshal(asMap["patternTriggerDefinition"], &out.PatternTriggerDefinition); err == nil {
	} else if err := go1.Unmarshal(asMap["patternUsageContext"], &out.PatternUsageContext); err == nil {
	} else if err := go1.Unmarshal(asMap["patternAvailability"], &out.PatternAvailability); err == nil {
	} else if err := go1.Unmarshal(asMap["patternExtendedContactDetail"], &out.PatternExtendedContactDetail); err == nil {
	} else if err := go1.Unmarshal(asMap["patternDosage"], &out.PatternDosage); err == nil {
	} else if err := go1.Unmarshal(asMap["patternMeta"], &out.PatternMeta); err == nil {
	} else {

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
	if len(asMap["label"]) > 0 {
		if err := go1.Unmarshal(asMap["label"], &out.Label); err != nil {
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
	if len(asMap["isModifier"]) > 0 {
		if err := go1.Unmarshal(asMap["isModifier"], &out.IsModifier); err != nil {
			return err
		}

	}
	if len(asMap["isSummary"]) > 0 {
		if err := go1.Unmarshal(asMap["isSummary"], &out.IsSummary); err != nil {
			return err
		}

	}
	if len(asMap["min"]) > 0 {
		if err := go1.Unmarshal(asMap["min"], &out.Min); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["maxValueDate"], &out.MaxValueDate); err == nil {
	} else if err := go1.Unmarshal(asMap["maxValueDateTime"], &out.MaxValueDateTime); err == nil {
	} else if err := go1.Unmarshal(asMap["maxValueInstant"], &out.MaxValueInstant); err == nil {
	} else if err := go1.Unmarshal(asMap["maxValueTime"], &out.MaxValueTime); err == nil {
	} else if err := go1.Unmarshal(asMap["maxValueDecimal"], &out.MaxValueDecimal); err == nil {
	} else if err := go1.Unmarshal(asMap["maxValueInteger"], &out.MaxValueInteger); err == nil {
	} else if err := go1.Unmarshal(asMap["maxValueInteger64"], &out.MaxValueInteger64); err == nil {
	} else if err := go1.Unmarshal(asMap["maxValuePositiveInt"], &out.MaxValuePositiveInt); err == nil {
	} else if err := go1.Unmarshal(asMap["maxValueUnsignedInt"], &out.MaxValueUnsignedInt); err == nil {
	} else if err := go1.Unmarshal(asMap["maxValueQuantity"], &out.MaxValueQuantity); err == nil {
	} else {

	}
	if len(asMap["maxLength"]) > 0 {
		if err := go1.Unmarshal(asMap["maxLength"], &out.MaxLength); err != nil {
			return err
		}

	}
	if len(asMap["binding"]) > 0 {
		if err := go1.Unmarshal(asMap["binding"], &out.Binding); err != nil {
			return err
		}

	}
	if len(asMap["valueAlternatives"]) > 0 {
		if err := go1.Unmarshal(asMap["valueAlternatives"], &out.ValueAlternatives); err != nil {
			return err
		}

	}
	if len(asMap["sliceName"]) > 0 {
		if err := go1.Unmarshal(asMap["sliceName"], &out.SliceName); err != nil {
			return err
		}

	}
	if len(asMap["alias"]) > 0 {
		if err := go1.Unmarshal(asMap["alias"], &out.Alias); err != nil {
			return err
		}

	}
	if len(asMap["base"]) > 0 {
		if err := go1.Unmarshal(asMap["base"], &out.Base); err != nil {
			return err
		}

	}
	if len(asMap["type"]) > 0 {
		if err := go1.Unmarshal(asMap["type"], &out.Type); err != nil {
			return err
		}

	}
	if len(asMap["orderMeaning"]) > 0 {
		if err := go1.Unmarshal(asMap["orderMeaning"], &out.OrderMeaning); err != nil {
			return err
		}

	}
	if len(asMap["mustHaveValue"]) > 0 {
		if err := go1.Unmarshal(asMap["mustHaveValue"], &out.MustHaveValue); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["path"], &out.Path); err != nil {
		return err
	}

	if len(asMap["constraint"]) > 0 {
		if err := go1.Unmarshal(asMap["constraint"], &out.Constraint); err != nil {
			return err
		}

	}
	if len(asMap["condition"]) > 0 {
		if err := go1.Unmarshal(asMap["condition"], &out.Condition); err != nil {
			return err
		}

	}
	if len(asMap["mustSupport"]) > 0 {
		if err := go1.Unmarshal(asMap["mustSupport"], &out.MustSupport); err != nil {
			return err
		}

	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["short"]) > 0 {
		if err := go1.Unmarshal(asMap["short"], &out.Short); err != nil {
			return err
		}

	}
	if len(asMap["requirements"]) > 0 {
		if err := go1.Unmarshal(asMap["requirements"], &out.Requirements); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["defaultValueBase64Binary"], &out.DefaultValueBase64Binary); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueBoolean"], &out.DefaultValueBoolean); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueCanonical"], &out.DefaultValueCanonical); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueCode"], &out.DefaultValueCode); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueDate"], &out.DefaultValueDate); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueDateTime"], &out.DefaultValueDateTime); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueDecimal"], &out.DefaultValueDecimal); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueId"], &out.DefaultValueId); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueInstant"], &out.DefaultValueInstant); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueInteger"], &out.DefaultValueInteger); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueInteger64"], &out.DefaultValueInteger64); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueMarkdown"], &out.DefaultValueMarkdown); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueOid"], &out.DefaultValueOid); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValuePositiveInt"], &out.DefaultValuePositiveInt); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueString"], &out.DefaultValueString); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueTime"], &out.DefaultValueTime); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueUnsignedInt"], &out.DefaultValueUnsignedInt); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueUri"], &out.DefaultValueUri); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueUrl"], &out.DefaultValueUrl); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueUuid"], &out.DefaultValueUuid); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueAddress"], &out.DefaultValueAddress); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueAge"], &out.DefaultValueAge); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueAnnotation"], &out.DefaultValueAnnotation); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueAttachment"], &out.DefaultValueAttachment); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueCodeableConcept"], &out.DefaultValueCodeableConcept); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueCodeableReference"], &out.DefaultValueCodeableReference); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueCoding"], &out.DefaultValueCoding); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueContactPoint"], &out.DefaultValueContactPoint); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueCount"], &out.DefaultValueCount); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueDistance"], &out.DefaultValueDistance); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueDuration"], &out.DefaultValueDuration); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueHumanName"], &out.DefaultValueHumanName); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueIdentifier"], &out.DefaultValueIdentifier); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueMoney"], &out.DefaultValueMoney); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValuePeriod"], &out.DefaultValuePeriod); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueQuantity"], &out.DefaultValueQuantity); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueRange"], &out.DefaultValueRange); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueRatio"], &out.DefaultValueRatio); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueRatioRange"], &out.DefaultValueRatioRange); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueReference"], &out.DefaultValueReference); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueSampledData"], &out.DefaultValueSampledData); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueSignature"], &out.DefaultValueSignature); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueTiming"], &out.DefaultValueTiming); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueContactDetail"], &out.DefaultValueContactDetail); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueDataRequirement"], &out.DefaultValueDataRequirement); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueExpression"], &out.DefaultValueExpression); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueParameterDefinition"], &out.DefaultValueParameterDefinition); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueRelatedArtifact"], &out.DefaultValueRelatedArtifact); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueTriggerDefinition"], &out.DefaultValueTriggerDefinition); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueUsageContext"], &out.DefaultValueUsageContext); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueAvailability"], &out.DefaultValueAvailability); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueExtendedContactDetail"], &out.DefaultValueExtendedContactDetail); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueDosage"], &out.DefaultValueDosage); err == nil {
	} else if err := go1.Unmarshal(asMap["defaultValueMeta"], &out.DefaultValueMeta); err == nil {
	} else {

	}
	if len(asMap["example"]) > 0 {
		if err := go1.Unmarshal(asMap["example"], &out.Example); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["minValueDate"], &out.MinValueDate); err == nil {
	} else if err := go1.Unmarshal(asMap["minValueDateTime"], &out.MinValueDateTime); err == nil {
	} else if err := go1.Unmarshal(asMap["minValueInstant"], &out.MinValueInstant); err == nil {
	} else if err := go1.Unmarshal(asMap["minValueTime"], &out.MinValueTime); err == nil {
	} else if err := go1.Unmarshal(asMap["minValueDecimal"], &out.MinValueDecimal); err == nil {
	} else if err := go1.Unmarshal(asMap["minValueInteger"], &out.MinValueInteger); err == nil {
	} else if err := go1.Unmarshal(asMap["minValueInteger64"], &out.MinValueInteger64); err == nil {
	} else if err := go1.Unmarshal(asMap["minValuePositiveInt"], &out.MinValuePositiveInt); err == nil {
	} else if err := go1.Unmarshal(asMap["minValueUnsignedInt"], &out.MinValueUnsignedInt); err == nil {
	} else if err := go1.Unmarshal(asMap["minValueQuantity"], &out.MinValueQuantity); err == nil {
	} else {

	}
	return nil
}

type ElementDefinition struct {
	Comment          *string `bson:",omitempty" json:"comment,omitempty"`          // Explanatory notes and implementation guidance about the data element, including notes about how to use the data properly, exceptions to proper use, etc. (Note: The text you are reading is specified in ElementDefinition.comment).
	Max              *string `bson:",omitempty" json:"max,omitempty"`              // The maximum number of times this element is permitted to appear in the instance.
	ContentReference *string `bson:",omitempty" json:"contentReference,omitempty"` // Identifies an element defined elsewhere in the definition whose content rules should be applied to the current element. ContentReferences bring across all the rules that are in the ElementDefinition for the element, including definitions, cardinality constraints, bindings, invariants etc.
	ElementDefinitionFixedx
	IsModifierReason    *string                   `bson:",omitempty" json:"isModifierReason,omitempty"` // Explains how that element affects the interpretation of the resource or element that contains it.
	Mapping             *ElementDefinitionMapping `binding:"omitempty" bson:",omitempty"`
	Representation      []*string                 `bson:",omitempty" json:"representation,omitempty"`      // Codes that define how this element is represented in instances, when the deviation varies from the normal case. No extensions are allowed on elements with a representation of 'xmlAttr', no matter what FHIR serialization format is used.
	SliceIsConstraining *bool                     `bson:",omitempty" json:"sliceIsConstraining,omitempty"` // If true, indicates that this slice definition is constraining a slice definition with the same name in an inherited profile. If false, the slice is not overriding any slice in an inherited profile. If missing, the slice might or might not be overriding a slice in an inherited profile, depending on the sliceName.
	Code                []*Coding                 `bson:",omitempty" json:"code,omitempty"`                // A code that has the same meaning as the element in a particular terminology.
	MeaningWhenMissing  *string                   `bson:",omitempty" json:"meaningWhenMissing,omitempty"`  // The Implicit meaning that is to be understood when this element is missing (e.g. 'when this element is missing, the period is ongoing').
	ElementDefinitionPatternx
	Extension         []*Extension `bson:",omitempty" json:"extension,omitempty"`         // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	ModifierExtension []*Extension `bson:",omitempty" json:"modifierExtension,omitempty"` /*
	May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.

	Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	*/
	Label      *string                   `bson:",omitempty" json:"label,omitempty"` // A single preferred label which is the text to display beside the element indicating its meaning or to use to prompt for the element in a user display or form.
	Slicing    *ElementDefinitionSlicing `binding:"omitempty" bson:",omitempty"`
	Definition *string                   `bson:",omitempty" json:"definition,omitempty"` // Provides a complete explanation of the meaning of the data element for human readability.  For the case of elements derived from existing elements (e.g. constraints), the definition SHALL be consistent with the base definition, but convey the meaning of the element in the particular context of use of the resource. (Note: The text you are reading is specified in ElementDefinition.definition).
	IsModifier *bool                     `bson:",omitempty" json:"isModifier,omitempty"` // If true, the value of this element affects the interpretation of the element or resource that contains it, and the value of the element cannot be ignored. Typically, this is used for status, negation and qualification codes. The effect of this is that the element cannot be ignored by systems: they SHALL either recognize the element and process it, and/or a pre-determination has been made that it is not relevant to their particular system. When used on the root element in an extension definition, this indicates whether or not the extension is a modifier extension.
	IsSummary  *bool                     `bson:",omitempty" json:"isSummary,omitempty"`  // Whether the element should be included if a client requests a search with the parameter _summary=true.
	Min        *int                      `bson:",omitempty" json:"min,omitempty"`        // The minimum number of times this element SHALL appear in the instance.
	ElementDefinitionMaxValuex
	MaxLength         *int                         `bson:",omitempty" json:"maxLength,omitempty"` // Indicates the maximum length in characters that is permitted to be present in conformant instances and which is expected to be supported by conformant consumers that support the element. ```maxLength``` SHOULD only be used on primitive data types that have a string representation (see [http://hl7.org/fhir/StructureDefinition/structuredefinition-type-characteristics](http://hl7.org/fhir/extensions/StructureDefinition-structuredefinition-type-characteristics.html)).
	Binding           *ElementDefinitionBinding    `binding:"omitempty" bson:",omitempty"`
	ValueAlternatives []*string                    `bson:",omitempty" json:"valueAlternatives,omitempty"` // Specifies a list of extensions that can appear in place of a primitive value.
	SliceName         *string                      `bson:",omitempty" json:"sliceName,omitempty"`         // The name of this element definition slice, when slicing is working. The name must be a token with no dots or spaces. This is a unique name referring to a specific set of constraints applied to this element, used to provide a name to different slices of the same element.
	Alias             []*string                    `bson:",omitempty" json:"alias,omitempty"`             // Identifies additional names by which this element might also be known.
	Base              *ElementDefinitionBase       `binding:"omitempty" bson:",omitempty"`
	Type              *ElementDefinitionType       `binding:"omitempty" bson:",omitempty"`
	OrderMeaning      *string                      `bson:",omitempty" json:"orderMeaning,omitempty"`            // If present, indicates that the order of the repeating element has meaning and describes what that meaning is.  If absent, it means that the order of the element has no meaning.
	MustHaveValue     *bool                        `bson:",omitempty" json:"mustHaveValue,omitempty"`           // Specifies for a primitive data type that the value of the data type cannot be replaced by an extension.
	Path              *string                      `binding:"required" bson:",omitempty" json:"path,omitempty"` // The path identifies the element and is expressed as a "."-separated list of ancestor elements, beginning with the name of the resource or extension.
	Constraint        *ElementDefinitionConstraint `binding:"omitempty" bson:",omitempty"`
	Condition         []**primitive.ObjectID       `bson:",omitempty" json:"condition,omitempty"`    // A reference to an invariant that may make additional statements about the cardinality or value in the instance.
	MustSupport       *bool                        `bson:",omitempty" json:"mustSupport,omitempty"`  // If true, implementations that produce or consume resources SHALL provide "support" for the element in some meaningful way. Note that this is being phased out and replaced by obligations (see below).  If false, the element may be ignored and not supported. If false, whether to populate or use the data element in any way is at the discretion of the implementation.
	Id                *primitive.ObjectID          `bson:"_id,omitempty" json:"id,omitempty"`        // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Short             *string                      `bson:",omitempty" json:"short,omitempty"`        // A concise description of what this element means (e.g. for use in autogenerated summaries).
	Requirements      *string                      `bson:",omitempty" json:"requirements,omitempty"` // This element is for traceability of why the element was created and why the constraints exist as they do. This may be used to point to source materials or specifications that drove the structure of this element.
	ElementDefinitionDefaultValuex
	Example *ElementDefinitionExample `binding:"omitempty" bson:",omitempty"`
	ElementDefinitionMinValuex
	ResourceType string `binding:"omitempty" bson:"-" json:"resourceType"`
}
type ElementDefinitionFixedx struct {
	FixedBase64Binary          Base64Binary          `bson:",omitempty" json:"fixedBase64Binary,omitempty"`
	FixedBoolean               bool                  `bson:",omitempty" json:"fixedBoolean,omitempty"`
	FixedCanonical             string                `bson:",omitempty" json:"fixedCanonical,omitempty"`
	FixedCode                  string                `bson:",omitempty" json:"fixedCode,omitempty"`
	FixedDate                  Date                  `bson:",omitempty" json:"fixedDate,omitempty"`
	FixedDateTime              DateTime              `bson:",omitempty" json:"fixedDateTime,omitempty"`
	FixedDecimal               float64               `bson:",omitempty" json:"fixedDecimal,omitempty"`
	FixedId                    *primitive.ObjectID   `bson:",omitempty" json:"fixedId,omitempty"`
	FixedInstant               time.Time             `bson:",omitempty" json:"fixedInstant,omitempty"`
	FixedInteger               int                   `bson:",omitempty" json:"fixedInteger,omitempty"`
	FixedInteger64             int64                 `bson:",omitempty" json:"fixedInteger64,omitempty"`
	FixedMarkdown              string                `bson:",omitempty" json:"fixedMarkdown,omitempty"`
	FixedOid                   string                `bson:",omitempty" json:"fixedOid,omitempty"`
	FixedPositiveInt           int                   `bson:",omitempty" json:"fixedPositiveInt,omitempty"`
	FixedString                string                `bson:",omitempty" json:"fixedString,omitempty"`
	FixedTime                  Time                  `bson:",omitempty" json:"fixedTime,omitempty"`
	FixedUnsignedInt           int                   `bson:",omitempty" json:"fixedUnsignedInt,omitempty"`
	FixedUri                   string                `bson:",omitempty" json:"fixedUri,omitempty"`
	FixedUrl                   url.URL               `bson:",omitempty" json:"fixedUrl,omitempty"`
	FixedUuid                  uuid.UUID             `bson:",omitempty" json:"fixedUuid,omitempty"`
	FixedAddress               Address               `bson:",omitempty" json:"fixedAddress,omitempty"`
	FixedAge                   Age                   `bson:",omitempty" json:"fixedAge,omitempty"`
	FixedAnnotation            Annotation            `bson:",omitempty" json:"fixedAnnotation,omitempty"`
	FixedAttachment            Attachment            `bson:",omitempty" json:"fixedAttachment,omitempty"`
	FixedCodeableConcept       CodeableConcept       `bson:",omitempty" json:"fixedCodeableConcept,omitempty"`
	FixedCodeableReference     CodeableReference     `bson:",omitempty" json:"fixedCodeableReference,omitempty"`
	FixedCoding                Coding                `bson:",omitempty" json:"fixedCoding,omitempty"`
	FixedContactPoint          ContactPoint          `bson:",omitempty" json:"fixedContactPoint,omitempty"`
	FixedCount                 Count                 `bson:",omitempty" json:"fixedCount,omitempty"`
	FixedDistance              Distance              `bson:",omitempty" json:"fixedDistance,omitempty"`
	FixedDuration              Duration              `bson:",omitempty" json:"fixedDuration,omitempty"`
	FixedHumanName             HumanName             `bson:",omitempty" json:"fixedHumanName,omitempty"`
	FixedIdentifier            Identifier            `bson:",omitempty" json:"fixedIdentifier,omitempty"`
	FixedMoney                 Money                 `bson:",omitempty" json:"fixedMoney,omitempty"`
	FixedPeriod                Period                `bson:",omitempty" json:"fixedPeriod,omitempty"`
	FixedQuantity              Quantity              `bson:",omitempty" json:"fixedQuantity,omitempty"`
	FixedRange                 Range                 `bson:",omitempty" json:"fixedRange,omitempty"`
	FixedRatio                 Ratio                 `bson:",omitempty" json:"fixedRatio,omitempty"`
	FixedRatioRange            RatioRange            `bson:",omitempty" json:"fixedRatioRange,omitempty"`
	FixedReference             Reference             `bson:",omitempty" json:"fixedReference,omitempty"`
	FixedSampledData           SampledData           `bson:",omitempty" json:"fixedSampledData,omitempty"`
	FixedSignature             Signature             `bson:",omitempty" json:"fixedSignature,omitempty"`
	FixedTiming                Timing                `bson:",omitempty" json:"fixedTiming,omitempty"`
	FixedContactDetail         ContactDetail         `bson:",omitempty" json:"fixedContactDetail,omitempty"`
	FixedDataRequirement       DataRequirement       `bson:",omitempty" json:"fixedDataRequirement,omitempty"`
	FixedExpression            Expression            `bson:",omitempty" json:"fixedExpression,omitempty"`
	FixedParameterDefinition   ParameterDefinition   `bson:",omitempty" json:"fixedParameterDefinition,omitempty"`
	FixedRelatedArtifact       RelatedArtifact       `bson:",omitempty" json:"fixedRelatedArtifact,omitempty"`
	FixedTriggerDefinition     TriggerDefinition     `bson:",omitempty" json:"fixedTriggerDefinition,omitempty"`
	FixedUsageContext          UsageContext          `bson:",omitempty" json:"fixedUsageContext,omitempty"`
	FixedAvailability          Availability          `bson:",omitempty" json:"fixedAvailability,omitempty"`
	FixedExtendedContactDetail ExtendedContactDetail `bson:",omitempty" json:"fixedExtendedContactDetail,omitempty"`
	FixedDosage                Dosage                `bson:",omitempty" json:"fixedDosage,omitempty"`
	FixedMeta                  Meta                  `bson:",omitempty" json:"fixedMeta,omitempty"`
}

func (out *ElementDefinitionMapping) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
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

	if len(asMap["language"]) > 0 {
		if err := go1.Unmarshal(asMap["language"], &out.Language); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["map"], &out.Map); err != nil {
		return err
	}

	if len(asMap["comment"]) > 0 {
		if err := go1.Unmarshal(asMap["comment"], &out.Comment); err != nil {
			return err
		}

	}
	return nil
}

type ElementDefinitionMapping struct {
	Id        *string              `bson:"_id,omitempty" json:"id,omitempty"`                       // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension []*Extension         `bson:",omitempty" json:"extension,omitempty"`                   // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Identity  **primitive.ObjectID `binding:"required" bson:",omitempty" json:"identity,omitempty"` // An internal reference to the definition of a mapping.
	Language  *string              `bson:",omitempty" json:"language,omitempty"`                    // Identifies the computable language in which mapping.map is expressed.
	Map       *string              `binding:"required" bson:",omitempty" json:"map,omitempty"`      // Expresses what part of the target specification corresponds to this element.
	Comment   *string              `bson:",omitempty" json:"comment,omitempty"`                     // Comments that provide information about the mapping or its use.
}
type ElementDefinitionPatternx struct {
	PatternBase64Binary          Base64Binary          `bson:",omitempty" json:"patternBase64Binary,omitempty"`
	PatternBoolean               bool                  `bson:",omitempty" json:"patternBoolean,omitempty"`
	PatternCanonical             string                `bson:",omitempty" json:"patternCanonical,omitempty"`
	PatternCode                  string                `bson:",omitempty" json:"patternCode,omitempty"`
	PatternDate                  Date                  `bson:",omitempty" json:"patternDate,omitempty"`
	PatternDateTime              DateTime              `bson:",omitempty" json:"patternDateTime,omitempty"`
	PatternDecimal               float64               `bson:",omitempty" json:"patternDecimal,omitempty"`
	PatternId                    *primitive.ObjectID   `bson:",omitempty" json:"patternId,omitempty"`
	PatternInstant               time.Time             `bson:",omitempty" json:"patternInstant,omitempty"`
	PatternInteger               int                   `bson:",omitempty" json:"patternInteger,omitempty"`
	PatternInteger64             int64                 `bson:",omitempty" json:"patternInteger64,omitempty"`
	PatternMarkdown              string                `bson:",omitempty" json:"patternMarkdown,omitempty"`
	PatternOid                   string                `bson:",omitempty" json:"patternOid,omitempty"`
	PatternPositiveInt           int                   `bson:",omitempty" json:"patternPositiveInt,omitempty"`
	PatternString                string                `bson:",omitempty" json:"patternString,omitempty"`
	PatternTime                  Time                  `bson:",omitempty" json:"patternTime,omitempty"`
	PatternUnsignedInt           int                   `bson:",omitempty" json:"patternUnsignedInt,omitempty"`
	PatternUri                   string                `bson:",omitempty" json:"patternUri,omitempty"`
	PatternUrl                   url.URL               `bson:",omitempty" json:"patternUrl,omitempty"`
	PatternUuid                  uuid.UUID             `bson:",omitempty" json:"patternUuid,omitempty"`
	PatternAddress               Address               `bson:",omitempty" json:"patternAddress,omitempty"`
	PatternAge                   Age                   `bson:",omitempty" json:"patternAge,omitempty"`
	PatternAnnotation            Annotation            `bson:",omitempty" json:"patternAnnotation,omitempty"`
	PatternAttachment            Attachment            `bson:",omitempty" json:"patternAttachment,omitempty"`
	PatternCodeableConcept       CodeableConcept       `bson:",omitempty" json:"patternCodeableConcept,omitempty"`
	PatternCodeableReference     CodeableReference     `bson:",omitempty" json:"patternCodeableReference,omitempty"`
	PatternCoding                Coding                `bson:",omitempty" json:"patternCoding,omitempty"`
	PatternContactPoint          ContactPoint          `bson:",omitempty" json:"patternContactPoint,omitempty"`
	PatternCount                 Count                 `bson:",omitempty" json:"patternCount,omitempty"`
	PatternDistance              Distance              `bson:",omitempty" json:"patternDistance,omitempty"`
	PatternDuration              Duration              `bson:",omitempty" json:"patternDuration,omitempty"`
	PatternHumanName             HumanName             `bson:",omitempty" json:"patternHumanName,omitempty"`
	PatternIdentifier            Identifier            `bson:",omitempty" json:"patternIdentifier,omitempty"`
	PatternMoney                 Money                 `bson:",omitempty" json:"patternMoney,omitempty"`
	PatternPeriod                Period                `bson:",omitempty" json:"patternPeriod,omitempty"`
	PatternQuantity              Quantity              `bson:",omitempty" json:"patternQuantity,omitempty"`
	PatternRange                 Range                 `bson:",omitempty" json:"patternRange,omitempty"`
	PatternRatio                 Ratio                 `bson:",omitempty" json:"patternRatio,omitempty"`
	PatternRatioRange            RatioRange            `bson:",omitempty" json:"patternRatioRange,omitempty"`
	PatternReference             Reference             `bson:",omitempty" json:"patternReference,omitempty"`
	PatternSampledData           SampledData           `bson:",omitempty" json:"patternSampledData,omitempty"`
	PatternSignature             Signature             `bson:",omitempty" json:"patternSignature,omitempty"`
	PatternTiming                Timing                `bson:",omitempty" json:"patternTiming,omitempty"`
	PatternContactDetail         ContactDetail         `bson:",omitempty" json:"patternContactDetail,omitempty"`
	PatternDataRequirement       DataRequirement       `bson:",omitempty" json:"patternDataRequirement,omitempty"`
	PatternExpression            Expression            `bson:",omitempty" json:"patternExpression,omitempty"`
	PatternParameterDefinition   ParameterDefinition   `bson:",omitempty" json:"patternParameterDefinition,omitempty"`
	PatternRelatedArtifact       RelatedArtifact       `bson:",omitempty" json:"patternRelatedArtifact,omitempty"`
	PatternTriggerDefinition     TriggerDefinition     `bson:",omitempty" json:"patternTriggerDefinition,omitempty"`
	PatternUsageContext          UsageContext          `bson:",omitempty" json:"patternUsageContext,omitempty"`
	PatternAvailability          Availability          `bson:",omitempty" json:"patternAvailability,omitempty"`
	PatternExtendedContactDetail ExtendedContactDetail `bson:",omitempty" json:"patternExtendedContactDetail,omitempty"`
	PatternDosage                Dosage                `bson:",omitempty" json:"patternDosage,omitempty"`
	PatternMeta                  Meta                  `bson:",omitempty" json:"patternMeta,omitempty"`
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
	if len(asMap["ordered"]) > 0 {
		if err := go1.Unmarshal(asMap["ordered"], &out.Ordered); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["rules"], &out.Rules); err != nil {
		return err
	}

	return nil
}

type ElementDefinitionSlicing struct {
	Id            *string                                `bson:"_id,omitempty" json:"id,omitempty"`     // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension     []*Extension                           `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Discriminator *ElementDefinitionSlicingDiscriminator `binding:"omitempty" bson:",omitempty"`
	Description   *string                                `bson:",omitempty" json:"description,omitempty"`              // A human-readable text description of how the slicing works. If there is no discriminator, this is required to be present to provide whatever information is possible about how the slices can be differentiated.
	Ordered       *bool                                  `bson:",omitempty" json:"ordered,omitempty"`                  // If the matching elements have to occur in the same order as defined in the profile.
	Rules         *string                                `binding:"required" bson:",omitempty" json:"rules,omitempty"` // Whether additional slices are allowed or not. When the slices are ordered, profile authors can also say that additional slices are only allowed at the end.
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
	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["type"], &out.Type); err != nil {
		return err
	}

	if err := go1.Unmarshal(asMap["path"], &out.Path); err != nil {
		return err
	}

	return nil
}

type ElementDefinitionSlicingDiscriminator struct {
	Id        *string      `bson:"_id,omitempty" json:"id,omitempty"`                   // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension []*Extension `bson:",omitempty" json:"extension,omitempty"`               // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Type      *string      `binding:"required" bson:",omitempty" json:"type,omitempty"` // How the element value is interpreted when discrimination is evaluated.
	Path      *string      `binding:"required" bson:",omitempty" json:"path,omitempty"` // A FHIRPath expression, using [the simple subset of FHIRPath](fhirpath.html#simple), that is used to identify the element on which discrimination is based.
}
type ElementDefinitionMaxValuex struct {
	MaxValueDate        Date      `bson:",omitempty" json:"maxValueDate,omitempty"`
	MaxValueDateTime    DateTime  `bson:",omitempty" json:"maxValueDateTime,omitempty"`
	MaxValueInstant     time.Time `bson:",omitempty" json:"maxValueInstant,omitempty"`
	MaxValueTime        Time      `bson:",omitempty" json:"maxValueTime,omitempty"`
	MaxValueDecimal     float64   `bson:",omitempty" json:"maxValueDecimal,omitempty"`
	MaxValueInteger     int       `bson:",omitempty" json:"maxValueInteger,omitempty"`
	MaxValueInteger64   int64     `bson:",omitempty" json:"maxValueInteger64,omitempty"`
	MaxValuePositiveInt int       `bson:",omitempty" json:"maxValuePositiveInt,omitempty"`
	MaxValueUnsignedInt int       `bson:",omitempty" json:"maxValueUnsignedInt,omitempty"`
	MaxValueQuantity    Quantity  `bson:",omitempty" json:"maxValueQuantity,omitempty"`
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
	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["strength"], &out.Strength); err != nil {
		return err
	}

	if len(asMap["description"]) > 0 {
		if err := go1.Unmarshal(asMap["description"], &out.Description); err != nil {
			return err
		}

	}
	if len(asMap["valueSet"]) > 0 {
		if err := go1.Unmarshal(asMap["valueSet"], &out.ValueSet); err != nil {
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
	Id          *string                             `bson:"_id,omitempty" json:"id,omitempty"`                       // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension   []*Extension                        `bson:",omitempty" json:"extension,omitempty"`                   // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Strength    *string                             `binding:"required" bson:",omitempty" json:"strength,omitempty"` // Indicates the degree of conformance expectations associated with this binding - that is, the degree to which the provided value set must be adhered to in the instances.
	Description *string                             `bson:",omitempty" json:"description,omitempty"`                 // Describes the intended use of this particular set of codes.
	ValueSet    *string                             `bson:",omitempty" json:"valueSet,omitempty"`                    // Refers to the value set that identifies the set of codes the binding refers to.
	Additional  *ElementDefinitionBindingAdditional `binding:"omitempty" bson:",omitempty"`
}

func (out *ElementDefinitionBindingAdditional) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["purpose"], &out.Purpose); err != nil {
		return err
	}

	if err := go1.Unmarshal(asMap["valueSet"], &out.ValueSet); err != nil {
		return err
	}

	if len(asMap["documentation"]) > 0 {
		if err := go1.Unmarshal(asMap["documentation"], &out.Documentation); err != nil {
			return err
		}

	}
	if len(asMap["shortDoco"]) > 0 {
		if err := go1.Unmarshal(asMap["shortDoco"], &out.ShortDoco); err != nil {
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
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	return nil
}

type ElementDefinitionBindingAdditional struct {
	Extension     []*Extension    `bson:",omitempty" json:"extension,omitempty"`                   // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Purpose       *string         `binding:"required" bson:",omitempty" json:"purpose,omitempty"`  // The use of this additional binding.
	ValueSet      *string         `binding:"required" bson:",omitempty" json:"valueSet,omitempty"` // The valueSet that is being bound for the purpose.
	Documentation *string         `bson:",omitempty" json:"documentation,omitempty"`               // Documentation of the purpose of use of the bindingproviding additional information about how it is intended to be used.
	ShortDoco     *string         `bson:",omitempty" json:"shortDoco,omitempty"`                   // Concise documentation - for summary tables.
	Usage         []*UsageContext `bson:",omitempty" json:"usage,omitempty"`                       // Qualifies the usage of the binding. Typically bindings are qualified by jurisdiction, but they may also be qualified by gender, workflow status, clinical domain etc. The information to decide whether a usege context applies is usually outside the resource, determined by context, and this might present challenges for validation tooling.
	Any           *bool           `bson:",omitempty" json:"any,omitempty"`                         // Whether the binding applies to all repeats, or just to any one of them. This is only relevant for elements that can repeat.
	Id            *string         `bson:"_id,omitempty" json:"id,omitempty"`                       // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
}

func (out *ElementDefinitionBase) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if err := go1.Unmarshal(asMap["max"], &out.Max); err != nil {
		return err
	}

	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
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

	if err := go1.Unmarshal(asMap["min"], &out.Min); err != nil {
		return err
	}

	return nil
}

type ElementDefinitionBase struct {
	Max       *string      `binding:"required" bson:",omitempty" json:"max,omitempty"`  // Maximum cardinality of the base element identified by the path.
	Id        *string      `bson:"_id,omitempty" json:"id,omitempty"`                   // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension []*Extension `bson:",omitempty" json:"extension,omitempty"`               // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Path      *string      `binding:"required" bson:",omitempty" json:"path,omitempty"` // The Path that identifies the base element - this matches the ElementDefinition.path for that element. Across FHIR, there is only one base definition of any element - that is, an element definition on a [StructureDefinition](structuredefinition.html#) without a StructureDefinition.base.
	Min       *int         `binding:"required" bson:",omitempty" json:"min,omitempty"`  // Minimum cardinality of the base element identified by the path.
}

func (out *ElementDefinitionType) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if len(asMap["profile"]) > 0 {
		if err := go1.Unmarshal(asMap["profile"], &out.Profile); err != nil {
			return err
		}

	}
	if len(asMap["targetProfile"]) > 0 {
		if err := go1.Unmarshal(asMap["targetProfile"], &out.TargetProfile); err != nil {
			return err
		}

	}
	if len(asMap["aggregation"]) > 0 {
		if err := go1.Unmarshal(asMap["aggregation"], &out.Aggregation); err != nil {
			return err
		}

	}
	if len(asMap["versioning"]) > 0 {
		if err := go1.Unmarshal(asMap["versioning"], &out.Versioning); err != nil {
			return err
		}

	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
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

	return nil
}

type ElementDefinitionType struct {
	Profile       []*string    `bson:",omitempty" json:"profile,omitempty"`                 // Identifies a profile structure or implementation Guide that applies to the datatype this element refers to. If any profiles are specified, then the content must conform to at least one of them. The URL can be a local reference - to a contained StructureDefinition, or a reference to another StructureDefinition or Implementation Guide by a canonical URL. When an implementation guide is specified, the type SHALL conform to at least one profile defined in the implementation guide.
	TargetProfile []*string    `bson:",omitempty" json:"targetProfile,omitempty"`           // Used when the type is "Reference" or "canonical", and identifies a profile structure or implementation Guide that applies to the target of the reference this element refers to. If any profiles are specified, then the content must conform to at least one of them. The URL can be a local reference - to a contained StructureDefinition, or a reference to another StructureDefinition or Implementation Guide by a canonical URL. When an implementation guide is specified, the target resource SHALL conform to at least one profile defined in the implementation guide.
	Aggregation   []*string    `bson:",omitempty" json:"aggregation,omitempty"`             // If the type is a reference to another resource, how the resource is or can be aggregated - is it a contained resource, or a reference, and if the context is a bundle, is it included in the bundle.
	Versioning    *string      `bson:",omitempty" json:"versioning,omitempty"`              // Whether this reference needs to be version specific or version independent, or whether either can be used.
	Id            *string      `bson:"_id,omitempty" json:"id,omitempty"`                   // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension     []*Extension `bson:",omitempty" json:"extension,omitempty"`               // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Code          *string      `binding:"required" bson:",omitempty" json:"code,omitempty"` // URL of Data type or Resource that is a(or the) type used for this element. References are URLs that are relative to http://hl7.org/fhir/StructureDefinition e.g. "string" is a reference to http://hl7.org/fhir/StructureDefinition/string. Absolute URLs are only allowed in logical models.
}

func (out *ElementDefinitionConstraint) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if len(asMap["suppress"]) > 0 {
		if err := go1.Unmarshal(asMap["suppress"], &out.Suppress); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["human"], &out.Human); err != nil {
		return err
	}

	if len(asMap["expression"]) > 0 {
		if err := go1.Unmarshal(asMap["expression"], &out.Expression); err != nil {
			return err
		}

	}
	if len(asMap["source"]) > 0 {
		if err := go1.Unmarshal(asMap["source"], &out.Source); err != nil {
			return err
		}

	}
	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["key"], &out.Key); err != nil {
		return err
	}

	if len(asMap["requirements"]) > 0 {
		if err := go1.Unmarshal(asMap["requirements"], &out.Requirements); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["severity"], &out.Severity); err != nil {
		return err
	}

	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	return nil
}

type ElementDefinitionConstraint struct {
	Suppress     *bool                `bson:",omitempty" json:"suppress,omitempty"`                    // If true, indicates that the warning or best practice guideline should be suppressed.
	Human        *string              `binding:"required" bson:",omitempty" json:"human,omitempty"`    // Text that can be used to describe the constraint in messages identifying that the constraint has been violated.
	Expression   *string              `bson:",omitempty" json:"expression,omitempty"`                  // A [FHIRPath](fhirpath.html) expression of constraint that can be executed to see if this constraint is met.
	Source       *string              `bson:",omitempty" json:"source,omitempty"`                      // A reference to the original source of the constraint, for traceability purposes.
	Extension    []*Extension         `bson:",omitempty" json:"extension,omitempty"`                   // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Key          **primitive.ObjectID `binding:"required" bson:",omitempty" json:"key,omitempty"`      // Allows identification of which elements have their cardinalities impacted by the constraint.  Will not be referenced for constraints that do not affect cardinality.
	Requirements *string              `bson:",omitempty" json:"requirements,omitempty"`                // Description of why this constraint is necessary or appropriate.
	Severity     *string              `binding:"required" bson:",omitempty" json:"severity,omitempty"` // Identifies the impact constraint violation has on the conformance of the instance.
	Id           *string              `bson:"_id,omitempty" json:"id,omitempty"`                       // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
}
type ElementDefinitionDefaultValuex struct {
	DefaultValueBase64Binary          Base64Binary          `bson:",omitempty" json:"defaultValueBase64Binary,omitempty"`
	DefaultValueBoolean               bool                  `bson:",omitempty" json:"defaultValueBoolean,omitempty"`
	DefaultValueCanonical             string                `bson:",omitempty" json:"defaultValueCanonical,omitempty"`
	DefaultValueCode                  string                `bson:",omitempty" json:"defaultValueCode,omitempty"`
	DefaultValueDate                  Date                  `bson:",omitempty" json:"defaultValueDate,omitempty"`
	DefaultValueDateTime              DateTime              `bson:",omitempty" json:"defaultValueDateTime,omitempty"`
	DefaultValueDecimal               float64               `bson:",omitempty" json:"defaultValueDecimal,omitempty"`
	DefaultValueId                    *primitive.ObjectID   `bson:",omitempty" json:"defaultValueId,omitempty"`
	DefaultValueInstant               time.Time             `bson:",omitempty" json:"defaultValueInstant,omitempty"`
	DefaultValueInteger               int                   `bson:",omitempty" json:"defaultValueInteger,omitempty"`
	DefaultValueInteger64             int64                 `bson:",omitempty" json:"defaultValueInteger64,omitempty"`
	DefaultValueMarkdown              string                `bson:",omitempty" json:"defaultValueMarkdown,omitempty"`
	DefaultValueOid                   string                `bson:",omitempty" json:"defaultValueOid,omitempty"`
	DefaultValuePositiveInt           int                   `bson:",omitempty" json:"defaultValuePositiveInt,omitempty"`
	DefaultValueString                string                `bson:",omitempty" json:"defaultValueString,omitempty"`
	DefaultValueTime                  Time                  `bson:",omitempty" json:"defaultValueTime,omitempty"`
	DefaultValueUnsignedInt           int                   `bson:",omitempty" json:"defaultValueUnsignedInt,omitempty"`
	DefaultValueUri                   string                `bson:",omitempty" json:"defaultValueUri,omitempty"`
	DefaultValueUrl                   url.URL               `bson:",omitempty" json:"defaultValueUrl,omitempty"`
	DefaultValueUuid                  uuid.UUID             `bson:",omitempty" json:"defaultValueUuid,omitempty"`
	DefaultValueAddress               Address               `bson:",omitempty" json:"defaultValueAddress,omitempty"`
	DefaultValueAge                   Age                   `bson:",omitempty" json:"defaultValueAge,omitempty"`
	DefaultValueAnnotation            Annotation            `bson:",omitempty" json:"defaultValueAnnotation,omitempty"`
	DefaultValueAttachment            Attachment            `bson:",omitempty" json:"defaultValueAttachment,omitempty"`
	DefaultValueCodeableConcept       CodeableConcept       `bson:",omitempty" json:"defaultValueCodeableConcept,omitempty"`
	DefaultValueCodeableReference     CodeableReference     `bson:",omitempty" json:"defaultValueCodeableReference,omitempty"`
	DefaultValueCoding                Coding                `bson:",omitempty" json:"defaultValueCoding,omitempty"`
	DefaultValueContactPoint          ContactPoint          `bson:",omitempty" json:"defaultValueContactPoint,omitempty"`
	DefaultValueCount                 Count                 `bson:",omitempty" json:"defaultValueCount,omitempty"`
	DefaultValueDistance              Distance              `bson:",omitempty" json:"defaultValueDistance,omitempty"`
	DefaultValueDuration              Duration              `bson:",omitempty" json:"defaultValueDuration,omitempty"`
	DefaultValueHumanName             HumanName             `bson:",omitempty" json:"defaultValueHumanName,omitempty"`
	DefaultValueIdentifier            Identifier            `bson:",omitempty" json:"defaultValueIdentifier,omitempty"`
	DefaultValueMoney                 Money                 `bson:",omitempty" json:"defaultValueMoney,omitempty"`
	DefaultValuePeriod                Period                `bson:",omitempty" json:"defaultValuePeriod,omitempty"`
	DefaultValueQuantity              Quantity              `bson:",omitempty" json:"defaultValueQuantity,omitempty"`
	DefaultValueRange                 Range                 `bson:",omitempty" json:"defaultValueRange,omitempty"`
	DefaultValueRatio                 Ratio                 `bson:",omitempty" json:"defaultValueRatio,omitempty"`
	DefaultValueRatioRange            RatioRange            `bson:",omitempty" json:"defaultValueRatioRange,omitempty"`
	DefaultValueReference             Reference             `bson:",omitempty" json:"defaultValueReference,omitempty"`
	DefaultValueSampledData           SampledData           `bson:",omitempty" json:"defaultValueSampledData,omitempty"`
	DefaultValueSignature             Signature             `bson:",omitempty" json:"defaultValueSignature,omitempty"`
	DefaultValueTiming                Timing                `bson:",omitempty" json:"defaultValueTiming,omitempty"`
	DefaultValueContactDetail         ContactDetail         `bson:",omitempty" json:"defaultValueContactDetail,omitempty"`
	DefaultValueDataRequirement       DataRequirement       `bson:",omitempty" json:"defaultValueDataRequirement,omitempty"`
	DefaultValueExpression            Expression            `bson:",omitempty" json:"defaultValueExpression,omitempty"`
	DefaultValueParameterDefinition   ParameterDefinition   `bson:",omitempty" json:"defaultValueParameterDefinition,omitempty"`
	DefaultValueRelatedArtifact       RelatedArtifact       `bson:",omitempty" json:"defaultValueRelatedArtifact,omitempty"`
	DefaultValueTriggerDefinition     TriggerDefinition     `bson:",omitempty" json:"defaultValueTriggerDefinition,omitempty"`
	DefaultValueUsageContext          UsageContext          `bson:",omitempty" json:"defaultValueUsageContext,omitempty"`
	DefaultValueAvailability          Availability          `bson:",omitempty" json:"defaultValueAvailability,omitempty"`
	DefaultValueExtendedContactDetail ExtendedContactDetail `bson:",omitempty" json:"defaultValueExtendedContactDetail,omitempty"`
	DefaultValueDosage                Dosage                `bson:",omitempty" json:"defaultValueDosage,omitempty"`
	DefaultValueMeta                  Meta                  `bson:",omitempty" json:"defaultValueMeta,omitempty"`
}

func (out *ElementDefinitionExample) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["label"], &out.Label); err != nil {
		return err
	}

	if err := go1.Unmarshal(asMap["valueBase64Binary"], &out.ValueBase64Binary); err == nil {
	} else if err := go1.Unmarshal(asMap["valueBoolean"], &out.ValueBoolean); err == nil {
	} else if err := go1.Unmarshal(asMap["valueCanonical"], &out.ValueCanonical); err == nil {
	} else if err := go1.Unmarshal(asMap["valueCode"], &out.ValueCode); err == nil {
	} else if err := go1.Unmarshal(asMap["valueDate"], &out.ValueDate); err == nil {
	} else if err := go1.Unmarshal(asMap["valueDateTime"], &out.ValueDateTime); err == nil {
	} else if err := go1.Unmarshal(asMap["valueDecimal"], &out.ValueDecimal); err == nil {
	} else if err := go1.Unmarshal(asMap["valueId"], &out.ValueId); err == nil {
	} else if err := go1.Unmarshal(asMap["valueInstant"], &out.ValueInstant); err == nil {
	} else if err := go1.Unmarshal(asMap["valueInteger"], &out.ValueInteger); err == nil {
	} else if err := go1.Unmarshal(asMap["valueInteger64"], &out.ValueInteger64); err == nil {
	} else if err := go1.Unmarshal(asMap["valueMarkdown"], &out.ValueMarkdown); err == nil {
	} else if err := go1.Unmarshal(asMap["valueOid"], &out.ValueOid); err == nil {
	} else if err := go1.Unmarshal(asMap["valuePositiveInt"], &out.ValuePositiveInt); err == nil {
	} else if err := go1.Unmarshal(asMap["valueString"], &out.ValueString); err == nil {
	} else if err := go1.Unmarshal(asMap["valueTime"], &out.ValueTime); err == nil {
	} else if err := go1.Unmarshal(asMap["valueUnsignedInt"], &out.ValueUnsignedInt); err == nil {
	} else if err := go1.Unmarshal(asMap["valueUri"], &out.ValueUri); err == nil {
	} else if err := go1.Unmarshal(asMap["valueUrl"], &out.ValueUrl); err == nil {
	} else if err := go1.Unmarshal(asMap["valueUuid"], &out.ValueUuid); err == nil {
	} else if err := go1.Unmarshal(asMap["valueAddress"], &out.ValueAddress); err == nil {
	} else if err := go1.Unmarshal(asMap["valueAge"], &out.ValueAge); err == nil {
	} else if err := go1.Unmarshal(asMap["valueAnnotation"], &out.ValueAnnotation); err == nil {
	} else if err := go1.Unmarshal(asMap["valueAttachment"], &out.ValueAttachment); err == nil {
	} else if err := go1.Unmarshal(asMap["valueCodeableConcept"], &out.ValueCodeableConcept); err == nil {
	} else if err := go1.Unmarshal(asMap["valueCodeableReference"], &out.ValueCodeableReference); err == nil {
	} else if err := go1.Unmarshal(asMap["valueCoding"], &out.ValueCoding); err == nil {
	} else if err := go1.Unmarshal(asMap["valueContactPoint"], &out.ValueContactPoint); err == nil {
	} else if err := go1.Unmarshal(asMap["valueCount"], &out.ValueCount); err == nil {
	} else if err := go1.Unmarshal(asMap["valueDistance"], &out.ValueDistance); err == nil {
	} else if err := go1.Unmarshal(asMap["valueDuration"], &out.ValueDuration); err == nil {
	} else if err := go1.Unmarshal(asMap["valueHumanName"], &out.ValueHumanName); err == nil {
	} else if err := go1.Unmarshal(asMap["valueIdentifier"], &out.ValueIdentifier); err == nil {
	} else if err := go1.Unmarshal(asMap["valueMoney"], &out.ValueMoney); err == nil {
	} else if err := go1.Unmarshal(asMap["valuePeriod"], &out.ValuePeriod); err == nil {
	} else if err := go1.Unmarshal(asMap["valueQuantity"], &out.ValueQuantity); err == nil {
	} else if err := go1.Unmarshal(asMap["valueRange"], &out.ValueRange); err == nil {
	} else if err := go1.Unmarshal(asMap["valueRatio"], &out.ValueRatio); err == nil {
	} else if err := go1.Unmarshal(asMap["valueRatioRange"], &out.ValueRatioRange); err == nil {
	} else if err := go1.Unmarshal(asMap["valueReference"], &out.ValueReference); err == nil {
	} else if err := go1.Unmarshal(asMap["valueSampledData"], &out.ValueSampledData); err == nil {
	} else if err := go1.Unmarshal(asMap["valueSignature"], &out.ValueSignature); err == nil {
	} else if err := go1.Unmarshal(asMap["valueTiming"], &out.ValueTiming); err == nil {
	} else if err := go1.Unmarshal(asMap["valueContactDetail"], &out.ValueContactDetail); err == nil {
	} else if err := go1.Unmarshal(asMap["valueDataRequirement"], &out.ValueDataRequirement); err == nil {
	} else if err := go1.Unmarshal(asMap["valueExpression"], &out.ValueExpression); err == nil {
	} else if err := go1.Unmarshal(asMap["valueParameterDefinition"], &out.ValueParameterDefinition); err == nil {
	} else if err := go1.Unmarshal(asMap["valueRelatedArtifact"], &out.ValueRelatedArtifact); err == nil {
	} else if err := go1.Unmarshal(asMap["valueTriggerDefinition"], &out.ValueTriggerDefinition); err == nil {
	} else if err := go1.Unmarshal(asMap["valueUsageContext"], &out.ValueUsageContext); err == nil {
	} else if err := go1.Unmarshal(asMap["valueAvailability"], &out.ValueAvailability); err == nil {
	} else if err := go1.Unmarshal(asMap["valueExtendedContactDetail"], &out.ValueExtendedContactDetail); err == nil {
	} else if err := go1.Unmarshal(asMap["valueDosage"], &out.ValueDosage); err == nil {
	} else if err := go1.Unmarshal(asMap["valueMeta"], &out.ValueMeta); err == nil {
	} else {
		return fmt.Errorf("could not unmarshal %s into any of the possible types", "value[x]")
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	return nil
}

type ElementDefinitionExample struct {
	Extension []*Extension `bson:",omitempty" json:"extension,omitempty"`                // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Label     *string      `binding:"required" bson:",omitempty" json:"label,omitempty"` // Describes the purpose of this example among the set of examples.
	ElementDefinitionExampleValuex
	Id *string `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
}
type ElementDefinitionExampleValuex struct {
	ValueBase64Binary          Base64Binary          `bson:",omitempty" json:"valueBase64Binary,omitempty"`
	ValueBoolean               bool                  `bson:",omitempty" json:"valueBoolean,omitempty"`
	ValueCanonical             string                `bson:",omitempty" json:"valueCanonical,omitempty"`
	ValueCode                  string                `bson:",omitempty" json:"valueCode,omitempty"`
	ValueDate                  Date                  `bson:",omitempty" json:"valueDate,omitempty"`
	ValueDateTime              DateTime              `bson:",omitempty" json:"valueDateTime,omitempty"`
	ValueDecimal               float64               `bson:",omitempty" json:"valueDecimal,omitempty"`
	ValueId                    *primitive.ObjectID   `bson:",omitempty" json:"valueId,omitempty"`
	ValueInstant               time.Time             `bson:",omitempty" json:"valueInstant,omitempty"`
	ValueInteger               int                   `bson:",omitempty" json:"valueInteger,omitempty"`
	ValueInteger64             int64                 `bson:",omitempty" json:"valueInteger64,omitempty"`
	ValueMarkdown              string                `bson:",omitempty" json:"valueMarkdown,omitempty"`
	ValueOid                   string                `bson:",omitempty" json:"valueOid,omitempty"`
	ValuePositiveInt           int                   `bson:",omitempty" json:"valuePositiveInt,omitempty"`
	ValueString                string                `bson:",omitempty" json:"valueString,omitempty"`
	ValueTime                  Time                  `bson:",omitempty" json:"valueTime,omitempty"`
	ValueUnsignedInt           int                   `bson:",omitempty" json:"valueUnsignedInt,omitempty"`
	ValueUri                   string                `bson:",omitempty" json:"valueUri,omitempty"`
	ValueUrl                   url.URL               `bson:",omitempty" json:"valueUrl,omitempty"`
	ValueUuid                  uuid.UUID             `bson:",omitempty" json:"valueUuid,omitempty"`
	ValueAddress               Address               `bson:",omitempty" json:"valueAddress,omitempty"`
	ValueAge                   Age                   `bson:",omitempty" json:"valueAge,omitempty"`
	ValueAnnotation            Annotation            `bson:",omitempty" json:"valueAnnotation,omitempty"`
	ValueAttachment            Attachment            `bson:",omitempty" json:"valueAttachment,omitempty"`
	ValueCodeableConcept       CodeableConcept       `bson:",omitempty" json:"valueCodeableConcept,omitempty"`
	ValueCodeableReference     CodeableReference     `bson:",omitempty" json:"valueCodeableReference,omitempty"`
	ValueCoding                Coding                `bson:",omitempty" json:"valueCoding,omitempty"`
	ValueContactPoint          ContactPoint          `bson:",omitempty" json:"valueContactPoint,omitempty"`
	ValueCount                 Count                 `bson:",omitempty" json:"valueCount,omitempty"`
	ValueDistance              Distance              `bson:",omitempty" json:"valueDistance,omitempty"`
	ValueDuration              Duration              `bson:",omitempty" json:"valueDuration,omitempty"`
	ValueHumanName             HumanName             `bson:",omitempty" json:"valueHumanName,omitempty"`
	ValueIdentifier            Identifier            `bson:",omitempty" json:"valueIdentifier,omitempty"`
	ValueMoney                 Money                 `bson:",omitempty" json:"valueMoney,omitempty"`
	ValuePeriod                Period                `bson:",omitempty" json:"valuePeriod,omitempty"`
	ValueQuantity              Quantity              `bson:",omitempty" json:"valueQuantity,omitempty"`
	ValueRange                 Range                 `bson:",omitempty" json:"valueRange,omitempty"`
	ValueRatio                 Ratio                 `bson:",omitempty" json:"valueRatio,omitempty"`
	ValueRatioRange            RatioRange            `bson:",omitempty" json:"valueRatioRange,omitempty"`
	ValueReference             Reference             `bson:",omitempty" json:"valueReference,omitempty"`
	ValueSampledData           SampledData           `bson:",omitempty" json:"valueSampledData,omitempty"`
	ValueSignature             Signature             `bson:",omitempty" json:"valueSignature,omitempty"`
	ValueTiming                Timing                `bson:",omitempty" json:"valueTiming,omitempty"`
	ValueContactDetail         ContactDetail         `bson:",omitempty" json:"valueContactDetail,omitempty"`
	ValueDataRequirement       DataRequirement       `bson:",omitempty" json:"valueDataRequirement,omitempty"`
	ValueExpression            Expression            `bson:",omitempty" json:"valueExpression,omitempty"`
	ValueParameterDefinition   ParameterDefinition   `bson:",omitempty" json:"valueParameterDefinition,omitempty"`
	ValueRelatedArtifact       RelatedArtifact       `bson:",omitempty" json:"valueRelatedArtifact,omitempty"`
	ValueTriggerDefinition     TriggerDefinition     `bson:",omitempty" json:"valueTriggerDefinition,omitempty"`
	ValueUsageContext          UsageContext          `bson:",omitempty" json:"valueUsageContext,omitempty"`
	ValueAvailability          Availability          `bson:",omitempty" json:"valueAvailability,omitempty"`
	ValueExtendedContactDetail ExtendedContactDetail `bson:",omitempty" json:"valueExtendedContactDetail,omitempty"`
	ValueDosage                Dosage                `bson:",omitempty" json:"valueDosage,omitempty"`
	ValueMeta                  Meta                  `bson:",omitempty" json:"valueMeta,omitempty"`
}
type ElementDefinitionMinValuex struct {
	MinValueDate        Date      `bson:",omitempty" json:"minValueDate,omitempty"`
	MinValueDateTime    DateTime  `bson:",omitempty" json:"minValueDateTime,omitempty"`
	MinValueInstant     time.Time `bson:",omitempty" json:"minValueInstant,omitempty"`
	MinValueTime        Time      `bson:",omitempty" json:"minValueTime,omitempty"`
	MinValueDecimal     float64   `bson:",omitempty" json:"minValueDecimal,omitempty"`
	MinValueInteger     int       `bson:",omitempty" json:"minValueInteger,omitempty"`
	MinValueInteger64   int64     `bson:",omitempty" json:"minValueInteger64,omitempty"`
	MinValuePositiveInt int       `bson:",omitempty" json:"minValuePositiveInt,omitempty"`
	MinValueUnsignedInt int       `bson:",omitempty" json:"minValueUnsignedInt,omitempty"`
	MinValueQuantity    Quantity  `bson:",omitempty" json:"minValueQuantity,omitempty"`
}

func (out *Expression) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Expression\"" {
		return fmt.Errorf("resourceType is not %s", "Expression")
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
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
	if len(asMap["name"]) > 0 {
		if err := go1.Unmarshal(asMap["name"], &out.Name); err != nil {
			return err
		}

	}
	if len(asMap["language"]) > 0 {
		if err := go1.Unmarshal(asMap["language"], &out.Language); err != nil {
			return err
		}

	}
	if len(asMap["expression"]) > 0 {
		if err := go1.Unmarshal(asMap["expression"], &out.Expression); err != nil {
			return err
		}

	}
	if len(asMap["reference"]) > 0 {
		if err := go1.Unmarshal(asMap["reference"], &out.Reference); err != nil {
			return err
		}

	}
	return nil
}

type Expression struct {
	Id           **primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`       // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension    []*Extension         `bson:",omitempty" json:"extension,omitempty"`   // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Description  *string              `bson:",omitempty" json:"description,omitempty"` // A brief, natural language description of the condition that effectively communicates the intended semantics.
	Name         *string              `bson:",omitempty" json:"name,omitempty"`        // A short name assigned to the expression to allow for multiple reuse of the expression in the context where it is defined.
	Language     *string              `bson:",omitempty" json:"language,omitempty"`    // The media type of the language for the expression.
	Expression   *string              `bson:",omitempty" json:"expression,omitempty"`  // An expression in the specified language that returns a value.
	Reference    *string              `bson:",omitempty" json:"reference,omitempty"`   // A URI that defines where the expression is found.
	ResourceType string               `binding:"omitempty" bson:"-" json:"resourceType"`
}

func (out *ExtendedContactDetail) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"ExtendedContactDetail\"" {
		return fmt.Errorf("resourceType is not %s", "ExtendedContactDetail")
	}
	if len(asMap["address"]) > 0 {
		if err := go1.Unmarshal(asMap["address"], &out.Address); err != nil {
			return err
		}

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
	return nil
}

type ExtendedContactDetail struct {
	Address      *Address             `bson:",omitempty" json:"address,omitempty"`      // Address for the contact.
	Organization *Reference           `bson:",omitempty" json:"organization,omitempty"` // This contact detail is handled/monitored by a specific organization. If the name is provided in the contact, then it is referring to the named individual within this organization.
	Period       *Period              `bson:",omitempty" json:"period,omitempty"`       // Period that this contact was valid for usage.
	Id           **primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`        // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension    []*Extension         `bson:",omitempty" json:"extension,omitempty"`    // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Purpose      *CodeableConcept     `bson:",omitempty" json:"purpose,omitempty"`      // The purpose/type of contact.
	Name         []*HumanName         `bson:",omitempty" json:"name,omitempty"`         // The name of an individual to contact, some types of contact detail are usually blank.
	Telecom      []*ContactPoint      `bson:",omitempty" json:"telecom,omitempty"`      // The contact details application for the purpose defined.
	ResourceType string               `binding:"omitempty" bson:"-" json:"resourceType"`
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
	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["url"], &out.Url); err != nil {
		return err
	}

	if err := go1.Unmarshal(asMap["valueBase64Binary"], &out.ValueBase64Binary); err == nil {
	} else if err := go1.Unmarshal(asMap["valueBoolean"], &out.ValueBoolean); err == nil {
	} else if err := go1.Unmarshal(asMap["valueCanonical"], &out.ValueCanonical); err == nil {
	} else if err := go1.Unmarshal(asMap["valueCode"], &out.ValueCode); err == nil {
	} else if err := go1.Unmarshal(asMap["valueDate"], &out.ValueDate); err == nil {
	} else if err := go1.Unmarshal(asMap["valueDateTime"], &out.ValueDateTime); err == nil {
	} else if err := go1.Unmarshal(asMap["valueDecimal"], &out.ValueDecimal); err == nil {
	} else if err := go1.Unmarshal(asMap["valueId"], &out.ValueId); err == nil {
	} else if err := go1.Unmarshal(asMap["valueInstant"], &out.ValueInstant); err == nil {
	} else if err := go1.Unmarshal(asMap["valueInteger"], &out.ValueInteger); err == nil {
	} else if err := go1.Unmarshal(asMap["valueInteger64"], &out.ValueInteger64); err == nil {
	} else if err := go1.Unmarshal(asMap["valueMarkdown"], &out.ValueMarkdown); err == nil {
	} else if err := go1.Unmarshal(asMap["valueOid"], &out.ValueOid); err == nil {
	} else if err := go1.Unmarshal(asMap["valuePositiveInt"], &out.ValuePositiveInt); err == nil {
	} else if err := go1.Unmarshal(asMap["valueString"], &out.ValueString); err == nil {
	} else if err := go1.Unmarshal(asMap["valueTime"], &out.ValueTime); err == nil {
	} else if err := go1.Unmarshal(asMap["valueUnsignedInt"], &out.ValueUnsignedInt); err == nil {
	} else if err := go1.Unmarshal(asMap["valueUri"], &out.ValueUri); err == nil {
	} else if err := go1.Unmarshal(asMap["valueUrl"], &out.ValueUrl); err == nil {
	} else if err := go1.Unmarshal(asMap["valueUuid"], &out.ValueUuid); err == nil {
	} else if err := go1.Unmarshal(asMap["valueAddress"], &out.ValueAddress); err == nil {
	} else if err := go1.Unmarshal(asMap["valueAge"], &out.ValueAge); err == nil {
	} else if err := go1.Unmarshal(asMap["valueAnnotation"], &out.ValueAnnotation); err == nil {
	} else if err := go1.Unmarshal(asMap["valueAttachment"], &out.ValueAttachment); err == nil {
	} else if err := go1.Unmarshal(asMap["valueCodeableConcept"], &out.ValueCodeableConcept); err == nil {
	} else if err := go1.Unmarshal(asMap["valueCodeableReference"], &out.ValueCodeableReference); err == nil {
	} else if err := go1.Unmarshal(asMap["valueCoding"], &out.ValueCoding); err == nil {
	} else if err := go1.Unmarshal(asMap["valueContactPoint"], &out.ValueContactPoint); err == nil {
	} else if err := go1.Unmarshal(asMap["valueCount"], &out.ValueCount); err == nil {
	} else if err := go1.Unmarshal(asMap["valueDistance"], &out.ValueDistance); err == nil {
	} else if err := go1.Unmarshal(asMap["valueDuration"], &out.ValueDuration); err == nil {
	} else if err := go1.Unmarshal(asMap["valueHumanName"], &out.ValueHumanName); err == nil {
	} else if err := go1.Unmarshal(asMap["valueIdentifier"], &out.ValueIdentifier); err == nil {
	} else if err := go1.Unmarshal(asMap["valueMoney"], &out.ValueMoney); err == nil {
	} else if err := go1.Unmarshal(asMap["valuePeriod"], &out.ValuePeriod); err == nil {
	} else if err := go1.Unmarshal(asMap["valueQuantity"], &out.ValueQuantity); err == nil {
	} else if err := go1.Unmarshal(asMap["valueRange"], &out.ValueRange); err == nil {
	} else if err := go1.Unmarshal(asMap["valueRatio"], &out.ValueRatio); err == nil {
	} else if err := go1.Unmarshal(asMap["valueRatioRange"], &out.ValueRatioRange); err == nil {
	} else if err := go1.Unmarshal(asMap["valueReference"], &out.ValueReference); err == nil {
	} else if err := go1.Unmarshal(asMap["valueSampledData"], &out.ValueSampledData); err == nil {
	} else if err := go1.Unmarshal(asMap["valueSignature"], &out.ValueSignature); err == nil {
	} else if err := go1.Unmarshal(asMap["valueTiming"], &out.ValueTiming); err == nil {
	} else if err := go1.Unmarshal(asMap["valueContactDetail"], &out.ValueContactDetail); err == nil {
	} else if err := go1.Unmarshal(asMap["valueDataRequirement"], &out.ValueDataRequirement); err == nil {
	} else if err := go1.Unmarshal(asMap["valueExpression"], &out.ValueExpression); err == nil {
	} else if err := go1.Unmarshal(asMap["valueParameterDefinition"], &out.ValueParameterDefinition); err == nil {
	} else if err := go1.Unmarshal(asMap["valueRelatedArtifact"], &out.ValueRelatedArtifact); err == nil {
	} else if err := go1.Unmarshal(asMap["valueTriggerDefinition"], &out.ValueTriggerDefinition); err == nil {
	} else if err := go1.Unmarshal(asMap["valueUsageContext"], &out.ValueUsageContext); err == nil {
	} else if err := go1.Unmarshal(asMap["valueAvailability"], &out.ValueAvailability); err == nil {
	} else if err := go1.Unmarshal(asMap["valueExtendedContactDetail"], &out.ValueExtendedContactDetail); err == nil {
	} else if err := go1.Unmarshal(asMap["valueDosage"], &out.ValueDosage); err == nil {
	} else if err := go1.Unmarshal(asMap["valueMeta"], &out.ValueMeta); err == nil {
	} else {

	}
	return nil
}

type Extension struct {
	Id        **primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`                  // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension []*Extension         `bson:",omitempty" json:"extension,omitempty"`              // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Url       *string              `binding:"required" bson:",omitempty" json:"url,omitempty"` // Source of the definition for the extension code - a logical name or a URL.
	ExtensionValuex
	ResourceType string `binding:"omitempty" bson:"-" json:"resourceType"`
}
type ExtensionValuex struct {
	ValueBase64Binary          Base64Binary          `bson:",omitempty" json:"valueBase64Binary,omitempty"`
	ValueBoolean               bool                  `bson:",omitempty" json:"valueBoolean,omitempty"`
	ValueCanonical             string                `bson:",omitempty" json:"valueCanonical,omitempty"`
	ValueCode                  string                `bson:",omitempty" json:"valueCode,omitempty"`
	ValueDate                  Date                  `bson:",omitempty" json:"valueDate,omitempty"`
	ValueDateTime              DateTime              `bson:",omitempty" json:"valueDateTime,omitempty"`
	ValueDecimal               float64               `bson:",omitempty" json:"valueDecimal,omitempty"`
	ValueId                    *primitive.ObjectID   `bson:",omitempty" json:"valueId,omitempty"`
	ValueInstant               time.Time             `bson:",omitempty" json:"valueInstant,omitempty"`
	ValueInteger               int                   `bson:",omitempty" json:"valueInteger,omitempty"`
	ValueInteger64             int64                 `bson:",omitempty" json:"valueInteger64,omitempty"`
	ValueMarkdown              string                `bson:",omitempty" json:"valueMarkdown,omitempty"`
	ValueOid                   string                `bson:",omitempty" json:"valueOid,omitempty"`
	ValuePositiveInt           int                   `bson:",omitempty" json:"valuePositiveInt,omitempty"`
	ValueString                string                `bson:",omitempty" json:"valueString,omitempty"`
	ValueTime                  Time                  `bson:",omitempty" json:"valueTime,omitempty"`
	ValueUnsignedInt           int                   `bson:",omitempty" json:"valueUnsignedInt,omitempty"`
	ValueUri                   string                `bson:",omitempty" json:"valueUri,omitempty"`
	ValueUrl                   url.URL               `bson:",omitempty" json:"valueUrl,omitempty"`
	ValueUuid                  uuid.UUID             `bson:",omitempty" json:"valueUuid,omitempty"`
	ValueAddress               Address               `bson:",omitempty" json:"valueAddress,omitempty"`
	ValueAge                   Age                   `bson:",omitempty" json:"valueAge,omitempty"`
	ValueAnnotation            Annotation            `bson:",omitempty" json:"valueAnnotation,omitempty"`
	ValueAttachment            Attachment            `bson:",omitempty" json:"valueAttachment,omitempty"`
	ValueCodeableConcept       CodeableConcept       `bson:",omitempty" json:"valueCodeableConcept,omitempty"`
	ValueCodeableReference     CodeableReference     `bson:",omitempty" json:"valueCodeableReference,omitempty"`
	ValueCoding                Coding                `bson:",omitempty" json:"valueCoding,omitempty"`
	ValueContactPoint          ContactPoint          `bson:",omitempty" json:"valueContactPoint,omitempty"`
	ValueCount                 Count                 `bson:",omitempty" json:"valueCount,omitempty"`
	ValueDistance              Distance              `bson:",omitempty" json:"valueDistance,omitempty"`
	ValueDuration              Duration              `bson:",omitempty" json:"valueDuration,omitempty"`
	ValueHumanName             HumanName             `bson:",omitempty" json:"valueHumanName,omitempty"`
	ValueIdentifier            Identifier            `bson:",omitempty" json:"valueIdentifier,omitempty"`
	ValueMoney                 Money                 `bson:",omitempty" json:"valueMoney,omitempty"`
	ValuePeriod                Period                `bson:",omitempty" json:"valuePeriod,omitempty"`
	ValueQuantity              Quantity              `bson:",omitempty" json:"valueQuantity,omitempty"`
	ValueRange                 Range                 `bson:",omitempty" json:"valueRange,omitempty"`
	ValueRatio                 Ratio                 `bson:",omitempty" json:"valueRatio,omitempty"`
	ValueRatioRange            RatioRange            `bson:",omitempty" json:"valueRatioRange,omitempty"`
	ValueReference             Reference             `bson:",omitempty" json:"valueReference,omitempty"`
	ValueSampledData           SampledData           `bson:",omitempty" json:"valueSampledData,omitempty"`
	ValueSignature             Signature             `bson:",omitempty" json:"valueSignature,omitempty"`
	ValueTiming                Timing                `bson:",omitempty" json:"valueTiming,omitempty"`
	ValueContactDetail         ContactDetail         `bson:",omitempty" json:"valueContactDetail,omitempty"`
	ValueDataRequirement       DataRequirement       `bson:",omitempty" json:"valueDataRequirement,omitempty"`
	ValueExpression            Expression            `bson:",omitempty" json:"valueExpression,omitempty"`
	ValueParameterDefinition   ParameterDefinition   `bson:",omitempty" json:"valueParameterDefinition,omitempty"`
	ValueRelatedArtifact       RelatedArtifact       `bson:",omitempty" json:"valueRelatedArtifact,omitempty"`
	ValueTriggerDefinition     TriggerDefinition     `bson:",omitempty" json:"valueTriggerDefinition,omitempty"`
	ValueUsageContext          UsageContext          `bson:",omitempty" json:"valueUsageContext,omitempty"`
	ValueAvailability          Availability          `bson:",omitempty" json:"valueAvailability,omitempty"`
	ValueExtendedContactDetail ExtendedContactDetail `bson:",omitempty" json:"valueExtendedContactDetail,omitempty"`
	ValueDosage                Dosage                `bson:",omitempty" json:"valueDosage,omitempty"`
	ValueMeta                  Meta                  `bson:",omitempty" json:"valueMeta,omitempty"`
}

func (out *HumanName) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"HumanName\"" {
		return fmt.Errorf("resourceType is not %s", "HumanName")
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["family"]) > 0 {
		if err := go1.Unmarshal(asMap["family"], &out.Family); err != nil {
			return err
		}

	}
	if len(asMap["suffix"]) > 0 {
		if err := go1.Unmarshal(asMap["suffix"], &out.Suffix); err != nil {
			return err
		}

	}
	if len(asMap["period"]) > 0 {
		if err := go1.Unmarshal(asMap["period"], &out.Period); err != nil {
			return err
		}

	}
	if len(asMap["use"]) > 0 {
		if err := go1.Unmarshal(asMap["use"], &out.Use); err != nil {
			return err
		}

	}
	if len(asMap["text"]) > 0 {
		if err := go1.Unmarshal(asMap["text"], &out.Text); err != nil {
			return err
		}

	}
	if len(asMap["given"]) > 0 {
		if err := go1.Unmarshal(asMap["given"], &out.Given); err != nil {
			return err
		}

	}
	if len(asMap["prefix"]) > 0 {
		if err := go1.Unmarshal(asMap["prefix"], &out.Prefix); err != nil {
			return err
		}

	}
	return nil
}

type HumanName struct {
	Id           **primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`     // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension    []*Extension         `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Family       *string              `bson:",omitempty" json:"family,omitempty"`    // The part of a name that links to the genealogy. In some cultures (e.g. Eritrea) the family name of a son is the first name of his father.
	Suffix       []*string            `bson:",omitempty" json:"suffix,omitempty"`    // Part of the name that is acquired as a title due to academic, legal, employment or nobility status, etc. and that appears at the end of the name.
	Period       *Period              `bson:",omitempty" json:"period,omitempty"`    // Indicates the period of time when this name was valid for the named person.
	Use          *string              `bson:",omitempty" json:"use,omitempty"`       // Identifies the purpose for this name.
	Text         *string              `bson:",omitempty" json:"text,omitempty"`      // Specifies the entire name as it should be displayed e.g. on an application UI. This may be provided instead of or as well as the specific parts.
	Given        []*string            `bson:",omitempty" json:"given,omitempty"`     // Given name.
	Prefix       []*string            `bson:",omitempty" json:"prefix,omitempty"`    // Part of the name that is acquired as a title due to academic, legal, employment or nobility status, etc. and that appears at the start of the name.
	ResourceType string               `binding:"omitempty" bson:"-" json:"resourceType"`
}

func (out *Identifier) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Identifier\"" {
		return fmt.Errorf("resourceType is not %s", "Identifier")
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
	if len(asMap["value"]) > 0 {
		if err := go1.Unmarshal(asMap["value"], &out.Value); err != nil {
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

type Identifier struct {
	Assigner     *Reference           `bson:",omitempty" json:"assigner,omitempty"`  // Organization that issued/manages the identifier.
	Id           **primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`     // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension    []*Extension         `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Use          *string              `bson:",omitempty" json:"use,omitempty"`       // The purpose of this identifier.
	Type         *CodeableConcept     `bson:",omitempty" json:"type,omitempty"`      // A coded type for the identifier that can be used to determine which identifier to use for a specific purpose.
	System       *string              `bson:",omitempty" json:"system,omitempty"`    // Establishes the namespace for the value - that is, an absolute URL that describes a set values that are unique.
	Value        *string              `bson:",omitempty" json:"value,omitempty"`     // The portion of the identifier typically relevant to the user and which is unique within the context of the system.
	Period       *Period              `bson:",omitempty" json:"period,omitempty"`    // Time period during which identifier is/was valid for use.
	ResourceType string               `binding:"omitempty" bson:"-" json:"resourceType"`
}

func (out *MarketingStatus) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"MarketingStatus\"" {
		return fmt.Errorf("resourceType is not %s", "MarketingStatus")
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
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
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
	return nil
}

type MarketingStatus struct {
	Status            *CodeableConcept    `binding:"required" bson:",omitempty" json:"status,omitempty"` // This attribute provides information on the status of the marketing of the medicinal product See ISO/TS 20443 for more information and examples.
	DateRange         *Period             `bson:",omitempty" json:"dateRange,omitempty"`                 // The date when the Medicinal Product is placed on the market by the Marketing Authorization Holder (or where applicable, the manufacturer/distributor) in a country and/or jurisdiction shall be provided A complete date consisting of day, month and year shall be specified using the ISO 8601 date format NOTE “Placed on the market” refers to the release of the Medicinal Product into the distribution chain.
	RestoreDate       *DateTime           `bson:",omitempty" json:"restoreDate,omitempty"`               // The date when the Medicinal Product is placed on the market by the Marketing Authorization Holder (or where applicable, the manufacturer/distributor) in a country and/or jurisdiction shall be provided A complete date consisting of day, month and year shall be specified using the ISO 8601 date format NOTE “Placed on the market” refers to the release of the Medicinal Product into the distribution chain.
	Id                *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`                     // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension         []*Extension        `bson:",omitempty" json:"extension,omitempty"`                 // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	ModifierExtension []*Extension        `bson:",omitempty" json:"modifierExtension,omitempty"`         /*
	May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.

	Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	*/
	Country      *CodeableConcept `bson:",omitempty" json:"country,omitempty"`      // The country in which the marketing authorization has been granted shall be specified It should be specified using the ISO 3166 ‑ 1 alpha-2 code elements.
	Jurisdiction *CodeableConcept `bson:",omitempty" json:"jurisdiction,omitempty"` // Where a Medicines Regulatory Agency has granted a marketing authorization for which specific provisions within a jurisdiction apply, the jurisdiction can be specified using an appropriate controlled terminology The controlled term and the controlled term identifier shall be specified.
	ResourceType string           `binding:"omitempty" bson:"-" json:"resourceType"`
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
	if len(asMap["lastUpdated"]) > 0 {
		if err := go1.Unmarshal(asMap["lastUpdated"], &out.LastUpdated); err != nil {
			return err
		}

	}
	if len(asMap["source"]) > 0 {
		if err := go1.Unmarshal(asMap["source"], &out.Source); err != nil {
			return err
		}

	}
	if len(asMap["profile"]) > 0 {
		if err := go1.Unmarshal(asMap["profile"], &out.Profile); err != nil {
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
	Id           **primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`       // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension    []*Extension         `bson:",omitempty" json:"extension,omitempty"`   // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	VersionId    **primitive.ObjectID `bson:",omitempty" json:"versionId,omitempty"`   // The version specific identifier, as it appears in the version portion of the URL. This value changes when the resource is created, updated, or deleted.
	LastUpdated  *time.Time           `bson:",omitempty" json:"lastUpdated,omitempty"` // When the resource last changed - e.g. when the version changed.
	Source       *string              `bson:",omitempty" json:"source,omitempty"`      // A uri that identifies the source system of the resource. This provides a minimal amount of [Provenance](provenance.html#) information that can be used to track or differentiate the source of information in the resource. The source may identify another FHIR server, document, message, database, etc.
	Profile      []*string            `bson:",omitempty" json:"profile,omitempty"`     // A list of profiles (references to [StructureDefinition](structuredefinition.html#) resources) that this resource claims to conform to. The URL is a reference to [StructureDefinition.url](structuredefinition-definitions.html#StructureDefinition.url).
	Security     []*Coding            `bson:",omitempty" json:"security,omitempty"`    // Security labels applied to this resource. These tags connect specific resources to the overall security policy and infrastructure.
	Tag          []*Coding            `bson:",omitempty" json:"tag,omitempty"`         // Tags applied to this resource. Tags are intended to be used to identify and relate resources to process and workflow, and applications are not required to consider the tags when interpreting the meaning of a resource.
	ResourceType string               `binding:"omitempty" bson:"-" json:"resourceType"`
}

func (out *MonetaryComponent) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"MonetaryComponent\"" {
		return fmt.Errorf("resourceType is not %s", "MonetaryComponent")
	}
	if err := go1.Unmarshal(asMap["type"], &out.Type); err != nil {
		return err
	}

	if len(asMap["code"]) > 0 {
		if err := go1.Unmarshal(asMap["code"], &out.Code); err != nil {
			return err
		}

	}
	if len(asMap["factor"]) > 0 {
		if err := go1.Unmarshal(asMap["factor"], &out.Factor); err != nil {
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
	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	return nil
}

type MonetaryComponent struct {
	Type         *string             `binding:"required" bson:",omitempty" json:"type,omitempty"` // base | surcharge | deduction | discount | tax | informational.
	Code         *CodeableConcept    `bson:",omitempty" json:"code,omitempty"`                    // Codes may be used to differentiate between kinds of taxes, surcharges, discounts etc.
	Factor       *float64            `bson:",omitempty" json:"factor,omitempty"`                  // Factor used for calculating this component.
	Amount       *Money              `bson:",omitempty" json:"amount,omitempty"`                  // Explicit value amount to be used.
	Id           *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`                   // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension    []*Extension        `bson:",omitempty" json:"extension,omitempty"`               // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	ResourceType string              `binding:"omitempty" bson:"-" json:"resourceType"`
}

func (out *Money) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Money\"" {
		return fmt.Errorf("resourceType is not %s", "Money")
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
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
	if len(asMap["currency"]) > 0 {
		if err := go1.Unmarshal(asMap["currency"], &out.Currency); err != nil {
			return err
		}

	}
	return nil
}

type Money struct {
	Id           **primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`     // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension    []*Extension         `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Value        *float64             `bson:",omitempty" json:"value,omitempty"`     // Numerical value (with implicit precision).
	Currency     *string              `bson:",omitempty" json:"currency,omitempty"`  // ISO 4217 Currency Code.
	ResourceType string               `binding:"omitempty" bson:"-" json:"resourceType"`
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
	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["status"], &out.Status); err != nil {
		return err
	}

	if err := go1.Unmarshal(asMap["div"], &out.Div); err != nil {
		return err
	}

	return nil
}

type Narrative struct {
	Id           *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`                     // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension    []*Extension        `bson:",omitempty" json:"extension,omitempty"`                 // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Status       *string             `binding:"required" bson:",omitempty" json:"status,omitempty"` // The status of the narrative - whether it's entirely generated (from just the defined data or the extensions too), or whether a human authored it and it may contain additional data.
	Div          *string             `binding:"required" bson:",omitempty" json:"div,omitempty"`    // The actual narrative content, a stripped down version of XHTML.
	ResourceType string              `binding:"omitempty" bson:"-" json:"resourceType"`
}

func (out *ParameterDefinition) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"ParameterDefinition\"" {
		return fmt.Errorf("resourceType is not %s", "ParameterDefinition")
	}
	if err := go1.Unmarshal(asMap["type"], &out.Type); err != nil {
		return err
	}

	if len(asMap["profile"]) > 0 {
		if err := go1.Unmarshal(asMap["profile"], &out.Profile); err != nil {
			return err
		}

	}
	if len(asMap["documentation"]) > 0 {
		if err := go1.Unmarshal(asMap["documentation"], &out.Documentation); err != nil {
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
	if err := go1.Unmarshal(asMap["use"], &out.Use); err != nil {
		return err
	}

	if len(asMap["min"]) > 0 {
		if err := go1.Unmarshal(asMap["min"], &out.Min); err != nil {
			return err
		}

	}
	if len(asMap["max"]) > 0 {
		if err := go1.Unmarshal(asMap["max"], &out.Max); err != nil {
			return err
		}

	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	return nil
}

type ParameterDefinition struct {
	Type          *string              `binding:"required" bson:",omitempty" json:"type,omitempty"` // The type of the parameter.
	Profile       *string              `bson:",omitempty" json:"profile,omitempty"`                 // If specified, this indicates a profile that the input data must conform to, or that the output data will conform to.
	Documentation *string              `bson:",omitempty" json:"documentation,omitempty"`           // A brief discussion of what the parameter is for and how it is used by the module.
	Extension     []*Extension         `bson:",omitempty" json:"extension,omitempty"`               // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Name          *string              `bson:",omitempty" json:"name,omitempty"`                    // The name of the parameter used to allow access to the value of the parameter in evaluation contexts.
	Use           *string              `binding:"required" bson:",omitempty" json:"use,omitempty"`  // Whether the parameter is input or output for the module.
	Min           *int                 `bson:",omitempty" json:"min,omitempty"`                     // The minimum number of times this parameter SHALL appear in the request or response.
	Max           *string              `bson:",omitempty" json:"max,omitempty"`                     // The maximum number of times this element is permitted to appear in the request or response.
	Id            **primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`                   // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ResourceType  string               `binding:"omitempty" bson:"-" json:"resourceType"`
}

func (out *Period) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Period\"" {
		return fmt.Errorf("resourceType is not %s", "Period")
	}
	if len(asMap["start"]) > 0 {
		if err := go1.Unmarshal(asMap["start"], &out.Start); err != nil {
			return err
		}

	}
	if len(asMap["end"]) > 0 {
		if err := go1.Unmarshal(asMap["end"], &out.End); err != nil {
			return err
		}

	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
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

type Period struct {
	Start        *DateTime            `bson:",omitempty" json:"start,omitempty"`     // The start of the period. The boundary is inclusive.
	End          *DateTime            `bson:",omitempty" json:"end,omitempty"`       // The end of the period. If the end of the period is missing, it means no end was known or planned at the time the instance was created. The start may be in the past, and the end date in the future, which means that period is expected/planned to end at that time.
	Id           **primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`     // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension    []*Extension         `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	ResourceType string               `binding:"omitempty" bson:"-" json:"resourceType"`
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
	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	return nil
}

type PrimitiveType struct {
	Id           *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`     // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension    []*Extension        `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	ResourceType string              `binding:"omitempty" bson:"-" json:"resourceType"`
}

func (out *ProductShelfLife) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"ProductShelfLife\"" {
		return fmt.Errorf("resourceType is not %s", "ProductShelfLife")
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
	if err := go1.Unmarshal(asMap["periodDuration"], &out.PeriodDuration); err == nil {
	} else if err := go1.Unmarshal(asMap["periodString"], &out.PeriodString); err == nil {
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
	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	return nil
}

type ProductShelfLife struct {
	ModifierExtension []*Extension `bson:",omitempty" json:"modifierExtension,omitempty"` /*
	May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.

	Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	*/
	Type *CodeableConcept `bson:",omitempty" json:"type,omitempty"` // This describes the shelf life, taking into account various scenarios such as shelf life of the packaged Medicinal Product itself, shelf life after transformation where necessary and shelf life after the first opening of a bottle, etc. The shelf life type shall be specified using an appropriate controlled vocabulary The controlled term and the controlled term identifier shall be specified.
	ProductShelfLifePeriodx
	SpecialPrecautionsForStorage []*CodeableConcept  `bson:",omitempty" json:"specialPrecautionsForStorage,omitempty"` // Special precautions for storage, if any, can be specified using an appropriate controlled vocabulary The controlled term and the controlled term identifier shall be specified.
	Id                           *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`                        // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension                    []*Extension        `bson:",omitempty" json:"extension,omitempty"`                    // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	ResourceType                 string              `binding:"omitempty" bson:"-" json:"resourceType"`
}
type ProductShelfLifePeriodx struct {
	PeriodDuration Duration `bson:",omitempty" json:"periodDuration,omitempty"`
	PeriodString   string   `bson:",omitempty" json:"periodString,omitempty"`
}

func (out *Quantity) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Quantity\"" {
		return fmt.Errorf("resourceType is not %s", "Quantity")
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
	if len(asMap["comparator"]) > 0 {
		if err := go1.Unmarshal(asMap["comparator"], &out.Comparator); err != nil {
			return err
		}

	}
	if len(asMap["unit"]) > 0 {
		if err := go1.Unmarshal(asMap["unit"], &out.Unit); err != nil {
			return err
		}

	}
	if len(asMap["system"]) > 0 {
		if err := go1.Unmarshal(asMap["system"], &out.System); err != nil {
			return err
		}

	}
	if len(asMap["code"]) > 0 {
		if err := go1.Unmarshal(asMap["code"], &out.Code); err != nil {
			return err
		}

	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	return nil
}

type Quantity struct {
	Extension    []*Extension         `bson:",omitempty" json:"extension,omitempty"`  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Value        *float64             `bson:",omitempty" json:"value,omitempty"`      // The value of the measured amount. The value includes an implicit precision in the presentation of the value.
	Comparator   *string              `bson:",omitempty" json:"comparator,omitempty"` // How the value should be understood and represented - whether the actual value is greater or less than the stated value due to measurement issues; e.g. if the comparator is "<" , then the real value is < stated value.
	Unit         *string              `bson:",omitempty" json:"unit,omitempty"`       // A human-readable form of the unit.
	System       *string              `bson:",omitempty" json:"system,omitempty"`     // The identification of the system that provides the coded form of the unit.
	Code         *string              `bson:",omitempty" json:"code,omitempty"`       // A computer processable form of the unit in some unit representation system.
	Id           **primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`      // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ResourceType string               `binding:"omitempty" bson:"-" json:"resourceType"`
}

func (out *Range) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Range\"" {
		return fmt.Errorf("resourceType is not %s", "Range")
	}
	if len(asMap["high"]) > 0 {
		if err := go1.Unmarshal(asMap["high"], &out.High); err != nil {
			return err
		}

	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
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
	return nil
}

type Range struct {
	High         *Quantity            `bson:",omitempty" json:"high,omitempty"`      // The high limit. The boundary is inclusive.
	Id           **primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`     // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension    []*Extension         `bson:",omitempty" json:"extension,omitempty"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Low          *Quantity            `bson:",omitempty" json:"low,omitempty"`       // The low limit. The boundary is inclusive.
	ResourceType string               `binding:"omitempty" bson:"-" json:"resourceType"`
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
	Id           **primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`       // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension    []*Extension         `bson:",omitempty" json:"extension,omitempty"`   // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Numerator    *Quantity            `bson:",omitempty" json:"numerator,omitempty"`   // The value of the numerator.
	Denominator  *Quantity            `bson:",omitempty" json:"denominator,omitempty"` // The value of the denominator.
	ResourceType string               `binding:"omitempty" bson:"-" json:"resourceType"`
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
	Id            **primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`         // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension     []*Extension         `bson:",omitempty" json:"extension,omitempty"`     // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	LowNumerator  *Quantity            `bson:",omitempty" json:"lowNumerator,omitempty"`  // The value of the low limit numerator.
	HighNumerator *Quantity            `bson:",omitempty" json:"highNumerator,omitempty"` // The value of the high limit numerator.
	Denominator   *Quantity            `bson:",omitempty" json:"denominator,omitempty"`   // The value of the denominator.
	ResourceType  string               `binding:"omitempty" bson:"-" json:"resourceType"`
}

func (out *Reference) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Reference\"" {
		return fmt.Errorf("resourceType is not %s", "Reference")
	}
	if len(asMap["type"]) > 0 {
		if err := go1.Unmarshal(asMap["type"], &out.Type); err != nil {
			return err
		}

	}
	if len(asMap["identifier"]) > 0 {
		if err := go1.Unmarshal(asMap["identifier"], &out.Identifier); err != nil {
			return err
		}

	}
	if len(asMap["display"]) > 0 {
		if err := go1.Unmarshal(asMap["display"], &out.Display); err != nil {
			return err
		}

	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
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
	return nil
}

type Reference struct {
	Type *string `bson:",omitempty" json:"type,omitempty"` /*
	The expected type of the target of the reference. If both Reference.type and Reference.reference are populated and Reference.reference is a FHIR URL, both SHALL be consistent.

	The type is the Canonical URL of Resource Definition that is the type this reference refers to. References are URLs that are relative to http://hl7.org/fhir/StructureDefinition/ e.g. "Patient" is a reference to http://hl7.org/fhir/StructureDefinition/Patient. Absolute URLs are only allowed for logical models (and can only be used in references in logical models, not resources).
	*/
	Identifier   *Identifier          `bson:",omitempty" json:"identifier,omitempty"` // An identifier for the target resource. This is used when there is no way to reference the other resource directly, either because the entity it represents is not available through a FHIR server, or because there is no way for the author of the resource to convert a known identifier to an actual location. There is no requirement that a Reference.identifier point to something that is actually exposed as a FHIR instance, but it SHALL point to a business concept that would be expected to be exposed as a FHIR instance, and that instance would need to be of a FHIR resource type allowed by the reference.
	Display      *string              `bson:",omitempty" json:"display,omitempty"`    // Plain text narrative that identifies the resource in addition to the resource reference.
	Id           **primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`      // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension    []*Extension         `bson:",omitempty" json:"extension,omitempty"`  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Reference    *string              `bson:",omitempty" json:"reference,omitempty"`  // A reference to a location at which the other resource is found. The reference may be a relative reference, in which case it is relative to the service base URL, or an absolute URL that resolves to the location where the resource is found. The reference may be version specific or not. If the reference is not to a FHIR RESTful server, then it should be assumed to be version specific. Internal fragment references (start with '#') refer to contained resources.
	ResourceType string               `binding:"omitempty" bson:"-" json:"resourceType"`
}

func (out *RelatedArtifact) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"RelatedArtifact\"" {
		return fmt.Errorf("resourceType is not %s", "RelatedArtifact")
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
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

	if len(asMap["classifier"]) > 0 {
		if err := go1.Unmarshal(asMap["classifier"], &out.Classifier); err != nil {
			return err
		}

	}
	if len(asMap["display"]) > 0 {
		if err := go1.Unmarshal(asMap["display"], &out.Display); err != nil {
			return err
		}

	}
	if len(asMap["citation"]) > 0 {
		if err := go1.Unmarshal(asMap["citation"], &out.Citation); err != nil {
			return err
		}

	}
	if len(asMap["resource"]) > 0 {
		if err := go1.Unmarshal(asMap["resource"], &out.Resource); err != nil {
			return err
		}

	}
	if len(asMap["publicationDate"]) > 0 {
		if err := go1.Unmarshal(asMap["publicationDate"], &out.PublicationDate); err != nil {
			return err
		}

	}
	if len(asMap["label"]) > 0 {
		if err := go1.Unmarshal(asMap["label"], &out.Label); err != nil {
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
	if len(asMap["publicationStatus"]) > 0 {
		if err := go1.Unmarshal(asMap["publicationStatus"], &out.PublicationStatus); err != nil {
			return err
		}

	}
	return nil
}

type RelatedArtifact struct {
	Id                **primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`                   // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension         []*Extension         `bson:",omitempty" json:"extension,omitempty"`               // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Type              *string              `binding:"required" bson:",omitempty" json:"type,omitempty"` // The type of relationship to the related artifact.
	Classifier        []*CodeableConcept   `bson:",omitempty" json:"classifier,omitempty"`              // Provides additional classifiers of the related artifact.
	Display           *string              `bson:",omitempty" json:"display,omitempty"`                 // A brief description of the document or knowledge resource being referenced, suitable for display to a consumer.
	Citation          *string              `bson:",omitempty" json:"citation,omitempty"`                // A bibliographic citation for the related artifact. This text SHOULD be formatted according to an accepted citation format.
	Resource          *string              `bson:",omitempty" json:"resource,omitempty"`                // The related artifact, such as a library, value set, profile, or other knowledge resource.
	PublicationDate   *Date                `bson:",omitempty" json:"publicationDate,omitempty"`         // The date of publication of the artifact being referred to.
	Label             *string              `bson:",omitempty" json:"label,omitempty"`                   // A short label that can be used to reference the citation from elsewhere in the containing artifact, such as a footnote index.
	Document          *Attachment          `bson:",omitempty" json:"document,omitempty"`                // The document being referenced, represented as an attachment. This is exclusive with the resource element.
	ResourceReference *Reference           `bson:",omitempty" json:"resourceReference,omitempty"`       // The related artifact, if the artifact is not a canonical resource, or a resource reference to a canonical resource.
	PublicationStatus *string              `bson:",omitempty" json:"publicationStatus,omitempty"`       // The publication status of the artifact being referred to.
	ResourceType      string               `binding:"omitempty" bson:"-" json:"resourceType"`
}

func (out *SampledData) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"SampledData\"" {
		return fmt.Errorf("resourceType is not %s", "SampledData")
	}
	if len(asMap["offsets"]) > 0 {
		if err := go1.Unmarshal(asMap["offsets"], &out.Offsets); err != nil {
			return err
		}

	}
	if len(asMap["data"]) > 0 {
		if err := go1.Unmarshal(asMap["data"], &out.Data); err != nil {
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
	if len(asMap["factor"]) > 0 {
		if err := go1.Unmarshal(asMap["factor"], &out.Factor); err != nil {
			return err
		}

	}
	if len(asMap["lowerLimit"]) > 0 {
		if err := go1.Unmarshal(asMap["lowerLimit"], &out.LowerLimit); err != nil {
			return err
		}

	}
	if len(asMap["upperLimit"]) > 0 {
		if err := go1.Unmarshal(asMap["upperLimit"], &out.UpperLimit); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["dimensions"], &out.Dimensions); err != nil {
		return err
	}

	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["intervalUnit"], &out.IntervalUnit); err != nil {
		return err
	}

	if len(asMap["codeMap"]) > 0 {
		if err := go1.Unmarshal(asMap["codeMap"], &out.CodeMap); err != nil {
			return err
		}

	}
	return nil
}

type SampledData struct {
	Offsets      *string              `bson:",omitempty" json:"offsets,omitempty"`                         // A series of data points which are decimal values separated by a single space (character u20).  The units in which the offsets are expressed are found in intervalUnit.  The absolute point at which the measurements begin SHALL be conveyed outside the scope of this datatype, e.g. Observation.effectiveDateTime for a timing offset.
	Data         *string              `bson:",omitempty" json:"data,omitempty"`                            // A series of data points which are decimal values or codes separated by a single space (character u20). The special codes "E" (error), "L" (below detection limit) and "U" (above detection limit) are also defined for used in place of decimal values.
	Origin       *Quantity            `binding:"required" bson:",omitempty" json:"origin,omitempty"`       // The base quantity that a measured value of zero represents. In addition, this provides the units of the entire measurement series.
	Interval     *float64             `bson:",omitempty" json:"interval,omitempty"`                        // Amount of intervalUnits between samples, e.g. milliseconds for time-based sampling.
	Factor       *float64             `bson:",omitempty" json:"factor,omitempty"`                          // A correction factor that is applied to the sampled data points before they are added to the origin.
	LowerLimit   *float64             `bson:",omitempty" json:"lowerLimit,omitempty"`                      // The lower limit of detection of the measured points. This is needed if any of the data points have the value "L" (lower than detection limit).
	UpperLimit   *float64             `bson:",omitempty" json:"upperLimit,omitempty"`                      // The upper limit of detection of the measured points. This is needed if any of the data points have the value "U" (higher than detection limit).
	Dimensions   *int                 `binding:"required" bson:",omitempty" json:"dimensions,omitempty"`   // The number of sample points at each time point. If this value is greater than one, then the dimensions will be interlaced - all the sample points for a point in time will be recorded at once.
	Id           **primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`                           // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension    []*Extension         `bson:",omitempty" json:"extension,omitempty"`                       // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	IntervalUnit *string              `binding:"required" bson:",omitempty" json:"intervalUnit,omitempty"` // The measurement unit in which the sample interval is expressed.
	CodeMap      *string              `bson:",omitempty" json:"codeMap,omitempty"`                         // Reference to ConceptMap that defines the codes used in the data.
	ResourceType string               `binding:"omitempty" bson:"-" json:"resourceType"`
}

func (out *Signature) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Signature\"" {
		return fmt.Errorf("resourceType is not %s", "Signature")
	}
	if len(asMap["onBehalfOf"]) > 0 {
		if err := go1.Unmarshal(asMap["onBehalfOf"], &out.OnBehalfOf); err != nil {
			return err
		}

	}
	if len(asMap["targetFormat"]) > 0 {
		if err := go1.Unmarshal(asMap["targetFormat"], &out.TargetFormat); err != nil {
			return err
		}

	}
	if len(asMap["sigFormat"]) > 0 {
		if err := go1.Unmarshal(asMap["sigFormat"], &out.SigFormat); err != nil {
			return err
		}

	}
	if len(asMap["data"]) > 0 {
		if err := go1.Unmarshal(asMap["data"], &out.Data); err != nil {
			return err
		}

	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["type"]) > 0 {
		if err := go1.Unmarshal(asMap["type"], &out.Type); err != nil {
			return err
		}

	}
	if len(asMap["when"]) > 0 {
		if err := go1.Unmarshal(asMap["when"], &out.When); err != nil {
			return err
		}

	}
	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["who"]) > 0 {
		if err := go1.Unmarshal(asMap["who"], &out.Who); err != nil {
			return err
		}

	}
	return nil
}

type Signature struct {
	OnBehalfOf   *Reference           `bson:",omitempty" json:"onBehalfOf,omitempty"`   // A reference to an application-usable description of the identity that is represented by the signature.
	TargetFormat *string              `bson:",omitempty" json:"targetFormat,omitempty"` // A mime type that indicates the technical format of the target resources signed by the signature.
	SigFormat    *string              `bson:",omitempty" json:"sigFormat,omitempty"`    // A mime type that indicates the technical format of the signature. Important mime types are application/signature+xml for X ML DigSig, application/jose for JWS, and image/* for a graphical image of a signature, etc.
	Data         *Base64Binary        `bson:",omitempty" json:"data,omitempty"`         // The base64 encoding of the Signature content. When signature is not recorded electronically this element would be empty.
	Id           **primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`        // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Type         []*Coding            `bson:",omitempty" json:"type,omitempty"`         // An indication of the reason that the entity signed this document. This may be explicitly included as part of the signature information and can be used when determining accountability for various actions concerning the document.
	When         *time.Time           `bson:",omitempty" json:"when,omitempty"`         // When the digital signature was signed.
	Extension    []*Extension         `bson:",omitempty" json:"extension,omitempty"`    // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Who          *Reference           `bson:",omitempty" json:"who,omitempty"`          // A reference to an application-usable description of the identity that signed  (e.g. the signature used their private key).
	ResourceType string               `binding:"omitempty" bson:"-" json:"resourceType"`
}

func (out *Timing) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"Timing\"" {
		return fmt.Errorf("resourceType is not %s", "Timing")
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
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	return nil
}

type Timing struct {
	Extension         []*Extension `bson:",omitempty" json:"extension,omitempty"`         // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	ModifierExtension []*Extension `bson:",omitempty" json:"modifierExtension,omitempty"` /*
	May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.

	Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	*/
	Event        []*DateTime          `bson:",omitempty" json:"event,omitempty"` // Identifies specific times when the event occurs.
	Repeat       *TimingRepeat        `binding:"omitempty" bson:",omitempty"`
	Code         *CodeableConcept     `bson:",omitempty" json:"code,omitempty"`  // A code for the timing schedule (or just text in code.text). Some codes such as BID are ubiquitous, but many institutions define their own additional codes. If a code is provided, the code is understood to be a complete statement of whatever is specified in the structured timing data, and either the code or the data may be used to interpret the Timing, with the exception that .repeat.bounds still applies over the code (and is not contained in the code).
	Id           **primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ResourceType string               `binding:"omitempty" bson:"-" json:"resourceType"`
}

func (out *TimingRepeat) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if len(asMap["count"]) > 0 {
		if err := go1.Unmarshal(asMap["count"], &out.Count); err != nil {
			return err
		}

	}
	if len(asMap["periodUnit"]) > 0 {
		if err := go1.Unmarshal(asMap["periodUnit"], &out.PeriodUnit); err != nil {
			return err
		}

	}
	if len(asMap["duration"]) > 0 {
		if err := go1.Unmarshal(asMap["duration"], &out.Duration); err != nil {
			return err
		}

	}
	if len(asMap["durationUnit"]) > 0 {
		if err := go1.Unmarshal(asMap["durationUnit"], &out.DurationUnit); err != nil {
			return err
		}

	}
	if len(asMap["period"]) > 0 {
		if err := go1.Unmarshal(asMap["period"], &out.Period); err != nil {
			return err
		}

	}
	if len(asMap["when"]) > 0 {
		if err := go1.Unmarshal(asMap["when"], &out.When); err != nil {
			return err
		}

	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["frequency"]) > 0 {
		if err := go1.Unmarshal(asMap["frequency"], &out.Frequency); err != nil {
			return err
		}

	}
	if len(asMap["frequencyMax"]) > 0 {
		if err := go1.Unmarshal(asMap["frequencyMax"], &out.FrequencyMax); err != nil {
			return err
		}

	}
	if len(asMap["timeOfDay"]) > 0 {
		if err := go1.Unmarshal(asMap["timeOfDay"], &out.TimeOfDay); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["boundsDuration"], &out.BoundsDuration); err == nil {
	} else if err := go1.Unmarshal(asMap["boundsRange"], &out.BoundsRange); err == nil {
	} else if err := go1.Unmarshal(asMap["boundsPeriod"], &out.BoundsPeriod); err == nil {
	} else {

	}
	if len(asMap["durationMax"]) > 0 {
		if err := go1.Unmarshal(asMap["durationMax"], &out.DurationMax); err != nil {
			return err
		}

	}
	if len(asMap["dayOfWeek"]) > 0 {
		if err := go1.Unmarshal(asMap["dayOfWeek"], &out.DayOfWeek); err != nil {
			return err
		}

	}
	if len(asMap["offset"]) > 0 {
		if err := go1.Unmarshal(asMap["offset"], &out.Offset); err != nil {
			return err
		}

	}
	if len(asMap["countMax"]) > 0 {
		if err := go1.Unmarshal(asMap["countMax"], &out.CountMax); err != nil {
			return err
		}

	}
	if len(asMap["periodMax"]) > 0 {
		if err := go1.Unmarshal(asMap["periodMax"], &out.PeriodMax); err != nil {
			return err
		}

	}
	return nil
}

type TimingRepeat struct {
	Count      *int    `bson:",omitempty" json:"count,omitempty"`      // A total count of the desired number of repetitions across the duration of the entire timing specification. If countMax is present, this element indicates the lower bound of the allowed range of count values.
	PeriodUnit *string `bson:",omitempty" json:"periodUnit,omitempty"` /*
	The units of time for the period in UCUM units
	Normal practice is to use the 'mo' code as a calendar month when calculating the next occurrence.
	*/
	Duration     *float64 `bson:",omitempty" json:"duration,omitempty"`     // How long this thing happens for when it happens. If durationMax is present, this element indicates the lower bound of the allowed range of the duration.
	DurationUnit *string  `bson:",omitempty" json:"durationUnit,omitempty"` /*
	The units of time for the duration, in UCUM units
	Normal practice is to use the 'mo' code as a calendar month when calculating the next occurrence.
	*/
	Period       *float64     `bson:",omitempty" json:"period,omitempty"`       // Indicates the duration of time over which repetitions are to occur; e.g. to express "3 times per day", 3 would be the frequency and "1 day" would be the period. If periodMax is present, this element indicates the lower bound of the allowed range of the period length.
	When         []*string    `bson:",omitempty" json:"when,omitempty"`         // An approximate time period during the day, potentially linked to an event of daily living that indicates when the action should occur.
	Id           *string      `bson:"_id,omitempty" json:"id,omitempty"`        // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension    []*Extension `bson:",omitempty" json:"extension,omitempty"`    // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Frequency    *int         `bson:",omitempty" json:"frequency,omitempty"`    // The number of times to repeat the action within the specified period. If frequencyMax is present, this element indicates the lower bound of the allowed range of the frequency.
	FrequencyMax *int         `bson:",omitempty" json:"frequencyMax,omitempty"` // If present, indicates that the frequency is a range - so to repeat between [frequency] and [frequencyMax] times within the period or period range.
	TimeOfDay    []*Time      `bson:",omitempty" json:"timeOfDay,omitempty"`    // Specified time of day for action to take place.
	TimingRepeatBoundsx
	DurationMax *float64  `bson:",omitempty" json:"durationMax,omitempty"` // If present, indicates that the duration is a range - so to perform the action between [duration] and [durationMax] time length.
	DayOfWeek   []*string `bson:",omitempty" json:"dayOfWeek,omitempty"`   // If one or more days of week is provided, then the action happens only on the specified day(s).
	Offset      *int      `bson:",omitempty" json:"offset,omitempty"`      // The number of minutes from the event. If the event code does not indicate whether the minutes is before or after the event, then the offset is assumed to be after the event.
	CountMax    *int      `bson:",omitempty" json:"countMax,omitempty"`    // If present, indicates that the count is a range - so to perform the action between [count] and [countMax] times.
	PeriodMax   *float64  `bson:",omitempty" json:"periodMax,omitempty"`   // If present, indicates that the period is a range from [period] to [periodMax], allowing expressing concepts such as "do this once every 3-5 days.
}
type TimingRepeatBoundsx struct {
	BoundsDuration Duration `bson:",omitempty" json:"boundsDuration,omitempty"`
	BoundsRange    Range    `bson:",omitempty" json:"boundsRange,omitempty"`
	BoundsPeriod   Period   `bson:",omitempty" json:"boundsPeriod,omitempty"`
}

func (out *TriggerDefinition) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"TriggerDefinition\"" {
		return fmt.Errorf("resourceType is not %s", "TriggerDefinition")
	}
	if len(asMap["extension"]) > 0 {
		if err := go1.Unmarshal(asMap["extension"], &out.Extension); err != nil {
			return err
		}

	}
	if len(asMap["subscriptionTopic"]) > 0 {
		if err := go1.Unmarshal(asMap["subscriptionTopic"], &out.SubscriptionTopic); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["timingTiming"], &out.TimingTiming); err == nil {
	} else if err := go1.Unmarshal(asMap["timingReference"], &out.TimingReference); err == nil {
	} else if err := go1.Unmarshal(asMap["timingDate"], &out.TimingDate); err == nil {
	} else if err := go1.Unmarshal(asMap["timingDateTime"], &out.TimingDateTime); err == nil {
	} else {

	}
	if len(asMap["condition"]) > 0 {
		if err := go1.Unmarshal(asMap["condition"], &out.Condition); err != nil {
			return err
		}

	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

	}
	if err := go1.Unmarshal(asMap["type"], &out.Type); err != nil {
		return err
	}

	if len(asMap["name"]) > 0 {
		if err := go1.Unmarshal(asMap["name"], &out.Name); err != nil {
			return err
		}

	}
	if len(asMap["code"]) > 0 {
		if err := go1.Unmarshal(asMap["code"], &out.Code); err != nil {
			return err
		}

	}
	if len(asMap["data"]) > 0 {
		if err := go1.Unmarshal(asMap["data"], &out.Data); err != nil {
			return err
		}

	}
	return nil
}

type TriggerDefinition struct {
	Extension         []*Extension `bson:",omitempty" json:"extension,omitempty"`         // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	SubscriptionTopic *string      `bson:",omitempty" json:"subscriptionTopic,omitempty"` // A reference to a SubscriptionTopic resource that defines the event. If this element is provided, no other information about the trigger definition may be supplied.
	TriggerDefinitionTimingx
	Condition    *Expression          `bson:",omitempty" json:"condition,omitempty"`               // A boolean-valued expression that is evaluated in the context of the container of the trigger definition and returns whether or not the trigger fires.
	Id           **primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`                   // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Type         *string              `binding:"required" bson:",omitempty" json:"type,omitempty"` // The type of triggering event.
	Name         *string              `bson:",omitempty" json:"name,omitempty"`                    // A formal name for the event. This may be an absolute URI that identifies the event formally (e.g. from a trigger registry), or a simple relative URI that identifies the event in a local context.
	Code         *CodeableConcept     `bson:",omitempty" json:"code,omitempty"`                    // A code that identifies the event.
	Data         []*DataRequirement   `bson:",omitempty" json:"data,omitempty"`                    // The triggering data of the event (if this is a data trigger). If more than one data is requirement is specified, then all the data requirements must be true.
	ResourceType string               `binding:"omitempty" bson:"-" json:"resourceType"`
}
type TriggerDefinitionTimingx struct {
	TimingTiming    Timing    `bson:",omitempty" json:"timingTiming,omitempty"`
	TimingReference Reference `bson:",omitempty" json:"timingReference,omitempty"`
	TimingDate      Date      `bson:",omitempty" json:"timingDate,omitempty"`
	TimingDateTime  DateTime  `bson:",omitempty" json:"timingDateTime,omitempty"`
}

func (out *UsageContext) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"UsageContext\"" {
		return fmt.Errorf("resourceType is not %s", "UsageContext")
	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
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

	if err := go1.Unmarshal(asMap["valueCodeableConcept"], &out.ValueCodeableConcept); err == nil {
	} else if err := go1.Unmarshal(asMap["valueQuantity"], &out.ValueQuantity); err == nil {
	} else if err := go1.Unmarshal(asMap["valueRange"], &out.ValueRange); err == nil {
	} else if err := go1.Unmarshal(asMap["valueReference"], &out.ValueReference); err == nil {
	} else {
		return fmt.Errorf("could not unmarshal %s into any of the possible types", "value[x]")
	}
	return nil
}

type UsageContext struct {
	Id        **primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`                   // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension []*Extension         `bson:",omitempty" json:"extension,omitempty"`               // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Code      *Coding              `binding:"required" bson:",omitempty" json:"code,omitempty"` // A code that identifies the type of context being specified by this usage context.
	UsageContextValuex
	ResourceType string `binding:"omitempty" bson:"-" json:"resourceType"`
}
type UsageContextValuex struct {
	ValueCodeableConcept CodeableConcept `bson:",omitempty" json:"valueCodeableConcept,omitempty"`
	ValueQuantity        Quantity        `bson:",omitempty" json:"valueQuantity,omitempty"`
	ValueRange           Range           `bson:",omitempty" json:"valueRange,omitempty"`
	ValueReference       Reference       `bson:",omitempty" json:"valueReference,omitempty"`
}

func (out *VirtualServiceDetail) UnmarshalJSON(data []byte) error {
	var asMap map[string]go1.RawMessage
	if err := go1.Unmarshal(data, &asMap); err != nil {
		return err
	}
	if resource, ok := asMap["resourceType"]; ok && string(resource) != "\"VirtualServiceDetail\"" {
		return fmt.Errorf("resourceType is not %s", "VirtualServiceDetail")
	}
	if len(asMap["sessionKey"]) > 0 {
		if err := go1.Unmarshal(asMap["sessionKey"], &out.SessionKey); err != nil {
			return err
		}

	}
	if len(asMap["id"]) > 0 {
		if err := go1.Unmarshal(asMap["id"], &out.Id); err != nil {
			return err
		}

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
	if err := go1.Unmarshal(asMap["addressUrl"], &out.AddressUrl); err == nil {
	} else if err := go1.Unmarshal(asMap["addressString"], &out.AddressString); err == nil {
	} else if err := go1.Unmarshal(asMap["addressContactPoint"], &out.AddressContactPoint); err == nil {
	} else if err := go1.Unmarshal(asMap["addressExtendedContactDetail"], &out.AddressExtendedContactDetail); err == nil {
	} else {

	}
	if len(asMap["additionalInfo"]) > 0 {
		if err := go1.Unmarshal(asMap["additionalInfo"], &out.AdditionalInfo); err != nil {
			return err
		}

	}
	if len(asMap["maxParticipants"]) > 0 {
		if err := go1.Unmarshal(asMap["maxParticipants"], &out.MaxParticipants); err != nil {
			return err
		}

	}
	return nil
}

type VirtualServiceDetail struct {
	SessionKey  *string             `bson:",omitempty" json:"sessionKey,omitempty"`  // Session Key required by the virtual service.
	Id          *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`       // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension   []*Extension        `bson:",omitempty" json:"extension,omitempty"`   // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	ChannelType *Coding             `bson:",omitempty" json:"channelType,omitempty"` // The type of virtual service to connect to (i.e. Teams, Zoom, Specific VMR technology, WhatsApp).
	VirtualServiceDetailAddressx
	AdditionalInfo  []*url.URL `bson:",omitempty" json:"additionalInfo,omitempty"`  // Address to see alternative connection details.
	MaxParticipants *int       `bson:",omitempty" json:"maxParticipants,omitempty"` // Maximum number of participants supported by the virtual service.
	ResourceType    string     `binding:"omitempty" bson:"-" json:"resourceType"`
}
type VirtualServiceDetailAddressx struct {
	AddressUrl                   url.URL               `bson:",omitempty" json:"addressUrl,omitempty"`
	AddressString                string                `bson:",omitempty" json:"addressString,omitempty"`
	AddressContactPoint          ContactPoint          `bson:",omitempty" json:"addressContactPoint,omitempty"`
	AddressExtendedContactDetail ExtendedContactDetail `bson:",omitempty" json:"addressExtendedContactDetail,omitempty"`
}
