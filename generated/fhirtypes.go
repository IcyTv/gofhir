package generated

type Element struct {
	Id        string    `json:"id"`        // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension Extension `json:"extension"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
}

type BackboneElement struct {
	Id                string    `json:"id"`                // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension         Extension `json:"extension"`         // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	ModifierExtension Extension `json:"modifierExtension"` /*
	May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.

	Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	*/
}

type Address struct {
	Extension  Extension `json:"extension"`  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Use        string    `json:"use"`        // The purpose of this address.
	Text       string    `json:"text"`       // Specifies the entire address as it should be displayed e.g. on a postal label. This may be provided instead of or as well as the specific parts.
	State      string    `json:"state"`      // Sub-unit of a country with limited sovereignty in a federally organized country. A code may be used if codes are in common use (e.g. US 2 letter state codes).
	PostalCode string    `json:"postalCode"` // A postal code designating a region defined by the postal service.
	Country    string    `json:"country"`    // Country - a nation as commonly understood or generally accepted.
	Period     Period    `json:"period"`     // Time period when address was/is in use.
	Id         string    `json:"id"`         // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Type       string    `json:"type"`       // Distinguishes between physical addresses (those you can visit) and mailing addresses (e.g. PO Boxes and care-of addresses). Most addresses are both.
	Line       string    `json:"line"`       // This component contains the house number, apartment number, street name, street direction,  P.O. Box number, delivery hints, and similar address information.
	City       string    `json:"city"`       // The name of the city, town, suburb, village or other community or delivery center.
	District   string    `json:"district"`   // The name of the administrative area (county).
}

type Age struct {
	Extension  Extension `json:"extension"`  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Value      float64   `json:"value"`      // The value of the measured amount. The value includes an implicit precision in the presentation of the value.
	Comparator string    `json:"comparator"` // How the value should be understood and represented - whether the actual value is greater or less than the stated value due to measurement issues; e.g. if the comparator is "<" , then the real value is < stated value.
	Unit       string    `json:"unit"`       // A human-readable form of the unit.
	System     string    `json:"system"`     // The identification of the system that provides the coded form of the unit.
	Code       string    `json:"code"`       // A computer processable form of the unit in some unit representation system.
	Id         string    `json:"id"`         // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
}

type Annotation struct {
	Id        string    `json:"id"`        // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension Extension `json:"extension"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	AnnotationAuthorx
	Time string `json:"time"` // Indicates when this particular annotation was made.
	Text string `json:"text"` // The text of the annotation in markdown format.
}
type AnnotationAuthorx struct {
	AuthorReference Reference `json:"authorReference"`
	AuthorString    string    `json:"authorString"`
}

type Attachment struct {
	Size        int64     `json:"size"`        // The number of bytes of data that make up this attachment (before base64 encoding, if that is done).
	Hash        string    `json:"hash"`        // The calculated hash of the data using SHA-1. Represented using base64.
	Title       string    `json:"title"`       // A label or set of text to display in place of the data.
	Pages       int       `json:"pages"`       // The number of pages when printed.
	Extension   Extension `json:"extension"`   // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Data        string    `json:"data"`        // The actual data of the attachment - a sequence of bytes, base64 encoded.
	ContentType string    `json:"contentType"` // Identifies the type of the data in the attachment and allows a method to be chosen to interpret or render the data. Includes mime type parameters such as charset where appropriate.
	Url         string    `json:"url"`         // A location where the data can be accessed.
	Id          string    `json:"id"`          // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Duration    float64   `json:"duration"`    // The duration of the recording in seconds - for audio and video.
	Height      int       `json:"height"`      // Height of the image in pixels (photo/video).
	Width       int       `json:"width"`       // Width of the image in pixels (photo/video).
	Frames      int       `json:"frames"`      // The number of frames in a photo. This is used with a multi-page fax, or an imaging acquisition context that takes multiple slices in a single image, or an animated gif. If there is more than one frame, this SHALL have a value in order to alert interface software that a multi-frame capable rendering widget is required.
	Language    string    `json:"language"`    // The human language of the content. The value can be any valid value according to BCP 47.
	Creation    string    `json:"creation"`    // The date that the attachment was first created.
}

type Availability struct {
	AvailableTime    AvailableTime
	NotAvailableTime NotAvailableTime
	Id               string    `json:"id"`        // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension        Extension `json:"extension"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
}
type AvailableTime struct {
	AvailableStartTime string    `json:"availableStartTime"` // Opening time of day (ignored if allDay = true).
	AvailableEndTime   string    `json:"availableEndTime"`   // Closing time of day (ignored if allDay = true).
	Id                 string    `json:"id"`                 // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension          Extension `json:"extension"`          // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	DaysOfWeek         string    `json:"daysOfWeek"`         // mon | tue | wed | thu | fri | sat | sun.
	AllDay             bool      `json:"allDay"`             // Always available? i.e. 24 hour service.
}
type NotAvailableTime struct {
	During      Period    `json:"during"`      // Service not available during this period.
	Id          string    `json:"id"`          // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension   Extension `json:"extension"`   // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Description string    `json:"description"` // Reason presented to the user explaining why time not available.
}

type BackboneType struct {
	Id                string    `json:"id"`                // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension         Extension `json:"extension"`         // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	ModifierExtension Extension `json:"modifierExtension"` /*
	May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.

	Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	*/
}

type CodeableConcept struct {
	Id        string    `json:"id"`        // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension Extension `json:"extension"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Coding    Coding    `json:"coding"`    // A reference to a code defined by a terminology system.
	Text      string    `json:"text"`      // A human language representation of the concept as seen/selected/uttered by the user who entered the data and/or which represents the intended meaning of the user.
}

type CodeableReference struct {
	Id        string          `json:"id"`        // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension Extension       `json:"extension"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Concept   CodeableConcept `json:"concept"`   // A reference to a concept - e.g. the information is identified by its general class to the degree of precision found in the terminology.
	Reference Reference       `json:"reference"` // A reference to a resource the provides exact details about the information being referenced.
}

type Coding struct {
	Id           string    `json:"id"`           // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension    Extension `json:"extension"`    // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	System       string    `json:"system"`       // The identification of the code system that defines the meaning of the symbol in the code.
	Version      string    `json:"version"`      // The version of the code system which was used when choosing this code. Note that a well-maintained code system does not need the version reported, because the meaning of codes is consistent across versions. However this cannot consistently be assured, and when the meaning is not guaranteed to be consistent, the version SHOULD be exchanged.
	Code         string    `json:"code"`         // A symbol in syntax defined by the system. The symbol may be a predefined code or an expression in a syntax defined by the coding system (e.g. post-coordination).
	Display      string    `json:"display"`      // A representation of the meaning of the code in the system, following the rules of the system.
	UserSelected bool      `json:"userSelected"` // Indicates that this coding was chosen by a user directly - e.g. off a pick list of available items (codes or displays).
}

type ContactDetail struct {
	Id        string       `json:"id"`        // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension Extension    `json:"extension"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Name      string       `json:"name"`      // The name of an individual to contact.
	Telecom   ContactPoint `json:"telecom"`   // The contact details for the individual (if a name was provided) or the organization.
}

type ContactPoint struct {
	Id        string    `json:"id"`        // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension Extension `json:"extension"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	System    string    `json:"system"`    // Telecommunications form for contact point - what communications system is required to make use of the contact.
	Value     string    `json:"value"`     // The actual contact point details, in a form that is meaningful to the designated communication system (i.e. phone number or email address).
	Use       string    `json:"use"`       // Identifies the purpose for the contact point.
	Rank      int       `json:"rank"`      // Specifies a preferred order in which to use a set of contacts. ContactPoints with lower rank values are more preferred than those with higher rank values.
	Period    Period    `json:"period"`    // Time period when the contact point was/is in use.
}

type Contributor struct {
	Type      string        `json:"type"`      // The type of contributor.
	Name      string        `json:"name"`      // The name of the individual or organization responsible for the contribution.
	Contact   ContactDetail `json:"contact"`   // Contact details to assist a user in finding and communicating with the contributor.
	Id        string        `json:"id"`        // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension Extension     `json:"extension"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
}

type Count struct {
	Id         string    `json:"id"`         // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension  Extension `json:"extension"`  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Value      float64   `json:"value"`      // The value of the measured amount. The value includes an implicit precision in the presentation of the value.
	Comparator string    `json:"comparator"` // How the value should be understood and represented - whether the actual value is greater or less than the stated value due to measurement issues; e.g. if the comparator is "<" , then the real value is < stated value.
	Unit       string    `json:"unit"`       // A human-readable form of the unit.
	System     string    `json:"system"`     // The identification of the system that provides the coded form of the unit.
	Code       string    `json:"code"`       // A computer processable form of the unit in some unit representation system.
}

type DataRequirement struct {
	Id      string `json:"id"`      // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Profile string `json:"profile"` // The profile of the required data, specified as the uri of the profile definition.
	DataRequirementSubjectx
	MustSupport string `json:"mustSupport"` /*
	Indicates that specific elements of the type are referenced by the knowledge module and must be supported by the consumer in order to obtain an effective evaluation. This does not mean that a value is required for this element, only that the consuming system must understand the element and be able to provide values for it if they are available.

	The value of mustSupport SHALL be a FHIRPath resolvable on the type of the DataRequirement. The path SHALL consist only of identifiers, constant indexers, and .resolve() (see the [Simple FHIRPath Profile](fhirpath.html#simple) for full details).
	*/
	CodeFilter  CodeFilter
	DateFilter  DateFilter
	ValueFilter ValueFilter
	Limit       int       `json:"limit"`     // Specifies a maximum number of results that are required (uses the _count search parameter).
	Extension   Extension `json:"extension"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Type        string    `json:"type"`      // The type of the required data, specified as the type name of a resource. For profiles, this value is set to the type of the base resource of the profile.
	Sort        Sort
}
type DataRequirementSubjectx struct {
	SubjectCodeableConcept CodeableConcept `json:"subjectCodeableConcept"`
	SubjectReference       Reference       `json:"subjectReference"`
}
type CodeFilter struct {
	Id          string    `json:"id"`          // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension   Extension `json:"extension"`   // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Path        string    `json:"path"`        // The code-valued attribute of the filter. The specified path SHALL be a FHIRPath resolvable on the specified type of the DataRequirement, and SHALL consist only of identifiers, constant indexers, and .resolve(). The path is allowed to contain qualifiers (.) to traverse sub-elements, as well as indexers ([x]) to traverse multiple-cardinality sub-elements (see the [Simple FHIRPath Profile](fhirpath.html#simple) for full details). Note that the index must be an integer constant. The path must resolve to an element of type code, Coding, or CodeableConcept.
	SearchParam string    `json:"searchParam"` // A token parameter that refers to a search parameter defined on the specified type of the DataRequirement, and which searches on elements of type code, Coding, or CodeableConcept.
	ValueSet    string    `json:"valueSet"`    // The valueset for the code filter. The valueSet and code elements are additive. If valueSet is specified, the filter will return only those data items for which the value of the code-valued element specified in the path is a member of the specified valueset.
	Code        Coding    `json:"code"`        // The codes for the code filter. If values are given, the filter will return only those data items for which the code-valued attribute specified by the path has a value that is one of the specified codes. If codes are specified in addition to a value set, the filter returns items matching a code in the value set or one of the specified codes.
}
type DateFilter struct {
	Id          string    `json:"id"`          // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension   Extension `json:"extension"`   // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Path        string    `json:"path"`        // The date-valued attribute of the filter. The specified path SHALL be a FHIRPath resolvable on the specified type of the DataRequirement, and SHALL consist only of identifiers, constant indexers, and .resolve(). The path is allowed to contain qualifiers (.) to traverse sub-elements, as well as indexers ([x]) to traverse multiple-cardinality sub-elements (see the [Simple FHIRPath Profile](fhirpath.html#simple) for full details). Note that the index must be an integer constant. The path must resolve to an element of type date, dateTime, Period, Schedule, or Timing.
	SearchParam string    `json:"searchParam"` // A date parameter that refers to a search parameter defined on the specified type of the DataRequirement, and which searches on elements of type date, dateTime, Period, Schedule, or Timing.
	DataRequirementDateFilterValuex
}
type DataRequirementDateFilterValuex struct {
	ValueDateTime string   `json:"valueDateTime"`
	ValuePeriod   Period   `json:"valuePeriod"`
	ValueDuration Duration `json:"valueDuration"`
}
type ValueFilter struct {
	Id          string    `json:"id"`          // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension   Extension `json:"extension"`   // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Path        string    `json:"path"`        // The attribute of the filter. The specified path SHALL be a FHIRPath resolvable on the specified type of the DataRequirement, and SHALL consist only of identifiers, constant indexers, and .resolve(). The path is allowed to contain qualifiers (.) to traverse sub-elements, as well as indexers ([x]) to traverse multiple-cardinality sub-elements (see the [Simple FHIRPath Profile](fhirpath.html#simple) for full details). Note that the index must be an integer constant. The path must resolve to an element of a type that is comparable to the valueFilter.value[x] element for the filter.
	SearchParam string    `json:"searchParam"` // A search parameter defined on the specified type of the DataRequirement, and which searches on elements of a type compatible with the type of the valueFilter.value[x] for the filter.
	Comparator  string    `json:"comparator"`  // The comparator to be used to determine whether the value is matching.
	DataRequirementValueFilterValuex
}
type DataRequirementValueFilterValuex struct {
	ValueDateTime string   `json:"valueDateTime"`
	ValuePeriod   Period   `json:"valuePeriod"`
	ValueDuration Duration `json:"valueDuration"`
}
type Sort struct {
	Path      string    `json:"path"`      // The attribute of the sort. The specified path must be resolvable from the type of the required data. The path is allowed to contain qualifiers (.) to traverse sub-elements, as well as indexers ([x]) to traverse multiple-cardinality sub-elements. Note that the index must be an integer constant.
	Direction string    `json:"direction"` // The direction of the sort, ascending or descending.
	Id        string    `json:"id"`        // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension Extension `json:"extension"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
}

type DataType struct {
	Extension Extension `json:"extension"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Id        string    `json:"id"`        // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
}

type Distance struct {
	Id         string    `json:"id"`         // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension  Extension `json:"extension"`  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Value      float64   `json:"value"`      // The value of the measured amount. The value includes an implicit precision in the presentation of the value.
	Comparator string    `json:"comparator"` // How the value should be understood and represented - whether the actual value is greater or less than the stated value due to measurement issues; e.g. if the comparator is "<" , then the real value is < stated value.
	Unit       string    `json:"unit"`       // A human-readable form of the unit.
	System     string    `json:"system"`     // The identification of the system that provides the coded form of the unit.
	Code       string    `json:"code"`       // A computer processable form of the unit in some unit representation system.
}

type Dosage struct {
	Sequence           int             `json:"sequence"`           // Indicates the order in which the dosage instructions should be applied or interpreted.
	AsNeededFor        CodeableConcept `json:"asNeededFor"`        // Indicates whether the Medication is only taken based on a precondition for taking the Medication (CodeableConcept).
	Site               CodeableConcept `json:"site"`               // Body site to administer to.
	Method             CodeableConcept `json:"method"`             // Technique for administering medication.
	MaxDosePerPeriod   Ratio           `json:"maxDosePerPeriod"`   // Upper limit on medication per unit of time.
	MaxDosePerLifetime Quantity        `json:"maxDosePerLifetime"` // Upper limit on medication per lifetime of the patient.
	Id                 string          `json:"id"`                 // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ModifierExtension  Extension       `json:"modifierExtension"`  /*
	May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.

	Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	*/
	Route                    CodeableConcept `json:"route"`                    // How drug should enter body.
	Extension                Extension       `json:"extension"`                // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Timing                   Timing          `json:"timing"`                   // When medication should be administered.
	AsNeeded                 bool            `json:"asNeeded"`                 // Indicates whether the Medication is only taken when needed within a specific dosing schedule (Boolean option).
	MaxDosePerAdministration Quantity        `json:"maxDosePerAdministration"` // Upper limit on medication per administration.
	PatientInstruction       string          `json:"patientInstruction"`       // Instructions in terms that are understood by the patient or consumer.
	DoseAndRate              DoseAndRate
	Text                     string          `json:"text"`                  // Free text dosage instructions e.g. SIG.
	AdditionalInstruction    CodeableConcept `json:"additionalInstruction"` // Supplemental instructions to the patient on how to take the medication  (e.g. "with meals" or"take half to one hour before food") or warnings for the patient about the medication (e.g. "may cause drowsiness" or "avoid exposure of skin to direct sunlight or sunlamps").
}
type DoseAndRate struct {
	DosageDoseAndRateRatex
	Id        string          `json:"id"`        // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension Extension       `json:"extension"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Type      CodeableConcept `json:"type"`      // The kind of dose or rate specified, for example, ordered or calculated.
	DosageDoseAndRateDosex
}
type DosageDoseAndRateRatex struct {
	RateRatio    Ratio    `json:"rateRatio"`
	RateRange    Range    `json:"rateRange"`
	RateQuantity Quantity `json:"rateQuantity"`
}
type DosageDoseAndRateDosex struct {
	DoseRange    Range    `json:"doseRange"`
	DoseQuantity Quantity `json:"doseQuantity"`
}

type Duration struct {
	Id         string    `json:"id"`         // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension  Extension `json:"extension"`  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Value      float64   `json:"value"`      // The value of the measured amount. The value includes an implicit precision in the presentation of the value.
	Comparator string    `json:"comparator"` // How the value should be understood and represented - whether the actual value is greater or less than the stated value due to measurement issues; e.g. if the comparator is "<" , then the real value is < stated value.
	Unit       string    `json:"unit"`       // A human-readable form of the unit.
	System     string    `json:"system"`     // The identification of the system that provides the coded form of the unit.
	Code       string    `json:"code"`       // A computer processable form of the unit in some unit representation system.
}

type ElementDefinition struct {
	SliceName         string `json:"sliceName"`    // The name of this element definition slice, when slicing is working. The name must be a token with no dots or spaces. This is a unique name referring to a specific set of constraints applied to this element, used to provide a name to different slices of the same element.
	Short             string `json:"short"`        // A concise description of what this element means (e.g. for use in autogenerated summaries).
	Requirements      string `json:"requirements"` // This element is for traceability of why the element was created and why the constraints exist as they do. This may be used to point to source materials or specifications that drove the structure of this element.
	OrderMeaning      string `json:"orderMeaning"` // If present, indicates that the order of the repeating element has meaning and describes what that meaning is.  If absent, it means that the order of the element has no meaning.
	Binding           Binding
	ModifierExtension Extension `json:"modifierExtension"` /*
	May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.

	Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	*/
	ElementDefinitionPatternx
	MustHaveValue    bool `json:"mustHaveValue"` // Specifies for a primitive data type that the value of the data type cannot be replaced by an extension.
	MaxLength        int  `json:"maxLength"`     // Indicates the maximum length in characters that is permitted to be present in conformant instances and which is expected to be supported by conformant consumers that support the element. ```maxLength``` SHOULD only be used on primitive data types that have a string representation (see [http://hl7.org/fhir/StructureDefinition/structuredefinition-type-characteristics](http://hl7.org/fhir/extensions/StructureDefinition-structuredefinition-type-characteristics.html)).
	Mapping          Mapping
	Id               string    `json:"id"`               // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension        Extension `json:"extension"`        // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Definition       string    `json:"definition"`       // Provides a complete explanation of the meaning of the data element for human readability.  For the case of elements derived from existing elements (e.g. constraints), the definition SHALL be consistent with the base definition, but convey the meaning of the element in the particular context of use of the resource. (Note: The text you are reading is specified in ElementDefinition.definition).
	Min              int       `json:"min"`              // The minimum number of times this element SHALL appear in the instance.
	ContentReference string    `json:"contentReference"` // Identifies an element defined elsewhere in the definition whose content rules should be applied to the current element. ContentReferences bring across all the rules that are in the ElementDefinition for the element, including definitions, cardinality constraints, bindings, invariants etc.
	Example          Example
	Condition        string `json:"condition"` // A reference to an invariant that may make additional statements about the cardinality or value in the instance.
	Path             string `json:"path"`      // The path identifies the element and is expressed as a "."-separated list of ancestor elements, beginning with the name of the resource or extension.
	Slicing          Slicing
	Alias            string `json:"alias"` // Identifies additional names by which this element might also be known.
	Max              string `json:"max"`   // The maximum number of times this element is permitted to appear in the instance.
	ElementDefinitionDefaultValuex
	ElementDefinitionMaxValuex
	Representation string `json:"representation"` // Codes that define how this element is represented in instances, when the deviation varies from the normal case. No extensions are allowed on elements with a representation of 'xmlAttr', no matter what FHIR serialization format is used.
	Label          string `json:"label"`          // A single preferred label which is the text to display beside the element indicating its meaning or to use to prompt for the element in a user display or form.
	Base           Base
	ElementDefinitionMinValuex
	Constraint          Constraint
	IsModifierReason    string `json:"isModifierReason"`    // Explains how that element affects the interpretation of the resource or element that contains it.
	SliceIsConstraining bool   `json:"sliceIsConstraining"` // If true, indicates that this slice definition is constraining a slice definition with the same name in an inherited profile. If false, the slice is not overriding any slice in an inherited profile. If missing, the slice might or might not be overriding a slice in an inherited profile, depending on the sliceName.
	Code                Coding `json:"code"`                // A code that has the same meaning as the element in a particular terminology.
	MeaningWhenMissing  string `json:"meaningWhenMissing"`  // The Implicit meaning that is to be understood when this element is missing (e.g. 'when this element is missing, the period is ongoing').
	ElementDefinitionFixedx
	IsSummary         bool   `json:"isSummary"` // Whether the element should be included if a client requests a search with the parameter _summary=true.
	Comment           string `json:"comment"`   // Explanatory notes and implementation guidance about the data element, including notes about how to use the data properly, exceptions to proper use, etc. (Note: The text you are reading is specified in ElementDefinition.comment).
	Type              Type
	MustSupport       bool   `json:"mustSupport"`       // If true, implementations that produce or consume resources SHALL provide "support" for the element in some meaningful way. Note that this is being phased out and replaced by obligations (see below).  If false, the element may be ignored and not supported. If false, whether to populate or use the data element in any way is at the discretion of the implementation.
	IsModifier        bool   `json:"isModifier"`        // If true, the value of this element affects the interpretation of the element or resource that contains it, and the value of the element cannot be ignored. Typically, this is used for status, negation and qualification codes. The effect of this is that the element cannot be ignored by systems: they SHALL either recognize the element and process it, and/or a pre-determination has been made that it is not relevant to their particular system. When used on the root element in an extension definition, this indicates whether or not the extension is a modifier extension.
	ValueAlternatives string `json:"valueAlternatives"` // Specifies a list of extensions that can appear in place of a primitive value.
}
type Binding struct {
	Id          string    `json:"id"`          // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension   Extension `json:"extension"`   // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Strength    string    `json:"strength"`    // Indicates the degree of conformance expectations associated with this binding - that is, the degree to which the provided value set must be adhered to in the instances.
	Description string    `json:"description"` // Describes the intended use of this particular set of codes.
	ValueSet    string    `json:"valueSet"`    // Refers to the value set that identifies the set of codes the binding refers to.
	Additional  Additional
}
type Additional struct {
	Id            string       `json:"id"`            // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension     Extension    `json:"extension"`     // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Purpose       string       `json:"purpose"`       // The use of this additional binding.
	ValueSet      string       `json:"valueSet"`      // The valueSet that is being bound for the purpose.
	Documentation string       `json:"documentation"` // Documentation of the purpose of use of the bindingproviding additional information about how it is intended to be used.
	ShortDoco     string       `json:"shortDoco"`     // Concise documentation - for summary tables.
	Usage         UsageContext `json:"usage"`         // Qualifies the usage of the binding. Typically bindings are qualified by jurisdiction, but they may also be qualified by gender, workflow status, clinical domain etc. The information to decide whether a usege context applies is usually outside the resource, determined by context, and this might present challenges for validation tooling.
	Any           bool         `json:"any"`           // Whether the binding applies to all repeats, or just to any one of them. This is only relevant for elements that can repeat.
}
type ElementDefinitionPatternx struct {
	PatternBase64Binary          string                `json:"patternBase64Binary"`
	PatternBoolean               bool                  `json:"patternBoolean"`
	PatternCanonical             string                `json:"patternCanonical"`
	PatternCode                  string                `json:"patternCode"`
	PatternDate                  string                `json:"patternDate"`
	PatternDateTime              string                `json:"patternDateTime"`
	PatternDecimal               float64               `json:"patternDecimal"`
	PatternId                    string                `json:"patternId"`
	PatternInstant               string                `json:"patternInstant"`
	PatternInteger               int                   `json:"patternInteger"`
	PatternInteger64             int64                 `json:"patternInteger64"`
	PatternMarkdown              string                `json:"patternMarkdown"`
	PatternOid                   string                `json:"patternOid"`
	PatternPositiveInt           int                   `json:"patternPositiveInt"`
	PatternString                string                `json:"patternString"`
	PatternTime                  string                `json:"patternTime"`
	PatternUnsignedInt           int                   `json:"patternUnsignedInt"`
	PatternUri                   string                `json:"patternUri"`
	PatternUrl                   string                `json:"patternUrl"`
	PatternUuid                  string                `json:"patternUuid"`
	PatternAddress               Address               `json:"patternAddress"`
	PatternAge                   Age                   `json:"patternAge"`
	PatternAnnotation            Annotation            `json:"patternAnnotation"`
	PatternAttachment            Attachment            `json:"patternAttachment"`
	PatternCodeableConcept       CodeableConcept       `json:"patternCodeableConcept"`
	PatternCodeableReference     CodeableReference     `json:"patternCodeableReference"`
	PatternCoding                Coding                `json:"patternCoding"`
	PatternContactPoint          ContactPoint          `json:"patternContactPoint"`
	PatternCount                 Count                 `json:"patternCount"`
	PatternDistance              Distance              `json:"patternDistance"`
	PatternDuration              Duration              `json:"patternDuration"`
	PatternHumanName             HumanName             `json:"patternHumanName"`
	PatternIdentifier            Identifier            `json:"patternIdentifier"`
	PatternMoney                 Money                 `json:"patternMoney"`
	PatternPeriod                Period                `json:"patternPeriod"`
	PatternQuantity              Quantity              `json:"patternQuantity"`
	PatternRange                 Range                 `json:"patternRange"`
	PatternRatio                 Ratio                 `json:"patternRatio"`
	PatternRatioRange            RatioRange            `json:"patternRatioRange"`
	PatternReference             Reference             `json:"patternReference"`
	PatternSampledData           SampledData           `json:"patternSampledData"`
	PatternSignature             Signature             `json:"patternSignature"`
	PatternTiming                Timing                `json:"patternTiming"`
	PatternContactDetail         ContactDetail         `json:"patternContactDetail"`
	PatternDataRequirement       DataRequirement       `json:"patternDataRequirement"`
	PatternExpression            Expression            `json:"patternExpression"`
	PatternParameterDefinition   ParameterDefinition   `json:"patternParameterDefinition"`
	PatternRelatedArtifact       RelatedArtifact       `json:"patternRelatedArtifact"`
	PatternTriggerDefinition     TriggerDefinition     `json:"patternTriggerDefinition"`
	PatternUsageContext          UsageContext          `json:"patternUsageContext"`
	PatternAvailability          Availability          `json:"patternAvailability"`
	PatternExtendedContactDetail ExtendedContactDetail `json:"patternExtendedContactDetail"`
	PatternDosage                Dosage                `json:"patternDosage"`
	PatternMeta                  Meta                  `json:"patternMeta"`
}
type Mapping struct {
	Id        string    `json:"id"`        // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension Extension `json:"extension"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Identity  string    `json:"identity"`  // An internal reference to the definition of a mapping.
	Language  string    `json:"language"`  // Identifies the computable language in which mapping.map is expressed.
	Map       string    `json:"map"`       // Expresses what part of the target specification corresponds to this element.
	Comment   string    `json:"comment"`   // Comments that provide information about the mapping or its use.
}
type Example struct {
	Id        string    `json:"id"`        // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension Extension `json:"extension"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Label     string    `json:"label"`     // Describes the purpose of this example among the set of examples.
	ElementDefinitionExampleValuex
}
type ElementDefinitionExampleValuex struct {
	ValueBase64Binary          string                `json:"valueBase64Binary"`
	ValueBoolean               bool                  `json:"valueBoolean"`
	ValueCanonical             string                `json:"valueCanonical"`
	ValueCode                  string                `json:"valueCode"`
	ValueDate                  string                `json:"valueDate"`
	ValueDateTime              string                `json:"valueDateTime"`
	ValueDecimal               float64               `json:"valueDecimal"`
	ValueId                    string                `json:"valueId"`
	ValueInstant               string                `json:"valueInstant"`
	ValueInteger               int                   `json:"valueInteger"`
	ValueInteger64             int64                 `json:"valueInteger64"`
	ValueMarkdown              string                `json:"valueMarkdown"`
	ValueOid                   string                `json:"valueOid"`
	ValuePositiveInt           int                   `json:"valuePositiveInt"`
	ValueString                string                `json:"valueString"`
	ValueTime                  string                `json:"valueTime"`
	ValueUnsignedInt           int                   `json:"valueUnsignedInt"`
	ValueUri                   string                `json:"valueUri"`
	ValueUrl                   string                `json:"valueUrl"`
	ValueUuid                  string                `json:"valueUuid"`
	ValueAddress               Address               `json:"valueAddress"`
	ValueAge                   Age                   `json:"valueAge"`
	ValueAnnotation            Annotation            `json:"valueAnnotation"`
	ValueAttachment            Attachment            `json:"valueAttachment"`
	ValueCodeableConcept       CodeableConcept       `json:"valueCodeableConcept"`
	ValueCodeableReference     CodeableReference     `json:"valueCodeableReference"`
	ValueCoding                Coding                `json:"valueCoding"`
	ValueContactPoint          ContactPoint          `json:"valueContactPoint"`
	ValueCount                 Count                 `json:"valueCount"`
	ValueDistance              Distance              `json:"valueDistance"`
	ValueDuration              Duration              `json:"valueDuration"`
	ValueHumanName             HumanName             `json:"valueHumanName"`
	ValueIdentifier            Identifier            `json:"valueIdentifier"`
	ValueMoney                 Money                 `json:"valueMoney"`
	ValuePeriod                Period                `json:"valuePeriod"`
	ValueQuantity              Quantity              `json:"valueQuantity"`
	ValueRange                 Range                 `json:"valueRange"`
	ValueRatio                 Ratio                 `json:"valueRatio"`
	ValueRatioRange            RatioRange            `json:"valueRatioRange"`
	ValueReference             Reference             `json:"valueReference"`
	ValueSampledData           SampledData           `json:"valueSampledData"`
	ValueSignature             Signature             `json:"valueSignature"`
	ValueTiming                Timing                `json:"valueTiming"`
	ValueContactDetail         ContactDetail         `json:"valueContactDetail"`
	ValueDataRequirement       DataRequirement       `json:"valueDataRequirement"`
	ValueExpression            Expression            `json:"valueExpression"`
	ValueParameterDefinition   ParameterDefinition   `json:"valueParameterDefinition"`
	ValueRelatedArtifact       RelatedArtifact       `json:"valueRelatedArtifact"`
	ValueTriggerDefinition     TriggerDefinition     `json:"valueTriggerDefinition"`
	ValueUsageContext          UsageContext          `json:"valueUsageContext"`
	ValueAvailability          Availability          `json:"valueAvailability"`
	ValueExtendedContactDetail ExtendedContactDetail `json:"valueExtendedContactDetail"`
	ValueDosage                Dosage                `json:"valueDosage"`
	ValueMeta                  Meta                  `json:"valueMeta"`
}
type Slicing struct {
	Description   string    `json:"description"` // A human-readable text description of how the slicing works. If there is no discriminator, this is required to be present to provide whatever information is possible about how the slices can be differentiated.
	Ordered       bool      `json:"ordered"`     // If the matching elements have to occur in the same order as defined in the profile.
	Rules         string    `json:"rules"`       // Whether additional slices are allowed or not. When the slices are ordered, profile authors can also say that additional slices are only allowed at the end.
	Id            string    `json:"id"`          // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension     Extension `json:"extension"`   // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Discriminator Discriminator
}
type Discriminator struct {
	Id        string    `json:"id"`        // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension Extension `json:"extension"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Type      string    `json:"type"`      // How the element value is interpreted when discrimination is evaluated.
	Path      string    `json:"path"`      // A FHIRPath expression, using [the simple subset of FHIRPath](fhirpath.html#simple), that is used to identify the element on which discrimination is based.
}
type ElementDefinitionDefaultValuex struct {
	DefaultValueBase64Binary          string                `json:"defaultValueBase64Binary"`
	DefaultValueBoolean               bool                  `json:"defaultValueBoolean"`
	DefaultValueCanonical             string                `json:"defaultValueCanonical"`
	DefaultValueCode                  string                `json:"defaultValueCode"`
	DefaultValueDate                  string                `json:"defaultValueDate"`
	DefaultValueDateTime              string                `json:"defaultValueDateTime"`
	DefaultValueDecimal               float64               `json:"defaultValueDecimal"`
	DefaultValueId                    string                `json:"defaultValueId"`
	DefaultValueInstant               string                `json:"defaultValueInstant"`
	DefaultValueInteger               int                   `json:"defaultValueInteger"`
	DefaultValueInteger64             int64                 `json:"defaultValueInteger64"`
	DefaultValueMarkdown              string                `json:"defaultValueMarkdown"`
	DefaultValueOid                   string                `json:"defaultValueOid"`
	DefaultValuePositiveInt           int                   `json:"defaultValuePositiveInt"`
	DefaultValueString                string                `json:"defaultValueString"`
	DefaultValueTime                  string                `json:"defaultValueTime"`
	DefaultValueUnsignedInt           int                   `json:"defaultValueUnsignedInt"`
	DefaultValueUri                   string                `json:"defaultValueUri"`
	DefaultValueUrl                   string                `json:"defaultValueUrl"`
	DefaultValueUuid                  string                `json:"defaultValueUuid"`
	DefaultValueAddress               Address               `json:"defaultValueAddress"`
	DefaultValueAge                   Age                   `json:"defaultValueAge"`
	DefaultValueAnnotation            Annotation            `json:"defaultValueAnnotation"`
	DefaultValueAttachment            Attachment            `json:"defaultValueAttachment"`
	DefaultValueCodeableConcept       CodeableConcept       `json:"defaultValueCodeableConcept"`
	DefaultValueCodeableReference     CodeableReference     `json:"defaultValueCodeableReference"`
	DefaultValueCoding                Coding                `json:"defaultValueCoding"`
	DefaultValueContactPoint          ContactPoint          `json:"defaultValueContactPoint"`
	DefaultValueCount                 Count                 `json:"defaultValueCount"`
	DefaultValueDistance              Distance              `json:"defaultValueDistance"`
	DefaultValueDuration              Duration              `json:"defaultValueDuration"`
	DefaultValueHumanName             HumanName             `json:"defaultValueHumanName"`
	DefaultValueIdentifier            Identifier            `json:"defaultValueIdentifier"`
	DefaultValueMoney                 Money                 `json:"defaultValueMoney"`
	DefaultValuePeriod                Period                `json:"defaultValuePeriod"`
	DefaultValueQuantity              Quantity              `json:"defaultValueQuantity"`
	DefaultValueRange                 Range                 `json:"defaultValueRange"`
	DefaultValueRatio                 Ratio                 `json:"defaultValueRatio"`
	DefaultValueRatioRange            RatioRange            `json:"defaultValueRatioRange"`
	DefaultValueReference             Reference             `json:"defaultValueReference"`
	DefaultValueSampledData           SampledData           `json:"defaultValueSampledData"`
	DefaultValueSignature             Signature             `json:"defaultValueSignature"`
	DefaultValueTiming                Timing                `json:"defaultValueTiming"`
	DefaultValueContactDetail         ContactDetail         `json:"defaultValueContactDetail"`
	DefaultValueDataRequirement       DataRequirement       `json:"defaultValueDataRequirement"`
	DefaultValueExpression            Expression            `json:"defaultValueExpression"`
	DefaultValueParameterDefinition   ParameterDefinition   `json:"defaultValueParameterDefinition"`
	DefaultValueRelatedArtifact       RelatedArtifact       `json:"defaultValueRelatedArtifact"`
	DefaultValueTriggerDefinition     TriggerDefinition     `json:"defaultValueTriggerDefinition"`
	DefaultValueUsageContext          UsageContext          `json:"defaultValueUsageContext"`
	DefaultValueAvailability          Availability          `json:"defaultValueAvailability"`
	DefaultValueExtendedContactDetail ExtendedContactDetail `json:"defaultValueExtendedContactDetail"`
	DefaultValueDosage                Dosage                `json:"defaultValueDosage"`
	DefaultValueMeta                  Meta                  `json:"defaultValueMeta"`
}
type ElementDefinitionMaxValuex struct {
	MaxValueDate        string   `json:"maxValueDate"`
	MaxValueDateTime    string   `json:"maxValueDateTime"`
	MaxValueInstant     string   `json:"maxValueInstant"`
	MaxValueTime        string   `json:"maxValueTime"`
	MaxValueDecimal     float64  `json:"maxValueDecimal"`
	MaxValueInteger     int      `json:"maxValueInteger"`
	MaxValueInteger64   int64    `json:"maxValueInteger64"`
	MaxValuePositiveInt int      `json:"maxValuePositiveInt"`
	MaxValueUnsignedInt int      `json:"maxValueUnsignedInt"`
	MaxValueQuantity    Quantity `json:"maxValueQuantity"`
}
type Base struct {
	Min       int       `json:"min"`       // Minimum cardinality of the base element identified by the path.
	Max       string    `json:"max"`       // Maximum cardinality of the base element identified by the path.
	Id        string    `json:"id"`        // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension Extension `json:"extension"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Path      string    `json:"path"`      // The Path that identifies the base element - this matches the ElementDefinition.path for that element. Across FHIR, there is only one base definition of any element - that is, an element definition on a [StructureDefinition](structuredefinition.html#) without a StructureDefinition.base.
}
type ElementDefinitionMinValuex struct {
	MinValueDate        string   `json:"minValueDate"`
	MinValueDateTime    string   `json:"minValueDateTime"`
	MinValueInstant     string   `json:"minValueInstant"`
	MinValueTime        string   `json:"minValueTime"`
	MinValueDecimal     float64  `json:"minValueDecimal"`
	MinValueInteger     int      `json:"minValueInteger"`
	MinValueInteger64   int64    `json:"minValueInteger64"`
	MinValuePositiveInt int      `json:"minValuePositiveInt"`
	MinValueUnsignedInt int      `json:"minValueUnsignedInt"`
	MinValueQuantity    Quantity `json:"minValueQuantity"`
}
type Constraint struct {
	Extension    Extension `json:"extension"`    // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Suppress     bool      `json:"suppress"`     // If true, indicates that the warning or best practice guideline should be suppressed.
	Id           string    `json:"id"`           // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Requirements string    `json:"requirements"` // Description of why this constraint is necessary or appropriate.
	Severity     string    `json:"severity"`     // Identifies the impact constraint violation has on the conformance of the instance.
	Human        string    `json:"human"`        // Text that can be used to describe the constraint in messages identifying that the constraint has been violated.
	Expression   string    `json:"expression"`   // A [FHIRPath](fhirpath.html) expression of constraint that can be executed to see if this constraint is met.
	Source       string    `json:"source"`       // A reference to the original source of the constraint, for traceability purposes.
	Key          string    `json:"key"`          // Allows identification of which elements have their cardinalities impacted by the constraint.  Will not be referenced for constraints that do not affect cardinality.
}
type ElementDefinitionFixedx struct {
	FixedBase64Binary          string                `json:"fixedBase64Binary"`
	FixedBoolean               bool                  `json:"fixedBoolean"`
	FixedCanonical             string                `json:"fixedCanonical"`
	FixedCode                  string                `json:"fixedCode"`
	FixedDate                  string                `json:"fixedDate"`
	FixedDateTime              string                `json:"fixedDateTime"`
	FixedDecimal               float64               `json:"fixedDecimal"`
	FixedId                    string                `json:"fixedId"`
	FixedInstant               string                `json:"fixedInstant"`
	FixedInteger               int                   `json:"fixedInteger"`
	FixedInteger64             int64                 `json:"fixedInteger64"`
	FixedMarkdown              string                `json:"fixedMarkdown"`
	FixedOid                   string                `json:"fixedOid"`
	FixedPositiveInt           int                   `json:"fixedPositiveInt"`
	FixedString                string                `json:"fixedString"`
	FixedTime                  string                `json:"fixedTime"`
	FixedUnsignedInt           int                   `json:"fixedUnsignedInt"`
	FixedUri                   string                `json:"fixedUri"`
	FixedUrl                   string                `json:"fixedUrl"`
	FixedUuid                  string                `json:"fixedUuid"`
	FixedAddress               Address               `json:"fixedAddress"`
	FixedAge                   Age                   `json:"fixedAge"`
	FixedAnnotation            Annotation            `json:"fixedAnnotation"`
	FixedAttachment            Attachment            `json:"fixedAttachment"`
	FixedCodeableConcept       CodeableConcept       `json:"fixedCodeableConcept"`
	FixedCodeableReference     CodeableReference     `json:"fixedCodeableReference"`
	FixedCoding                Coding                `json:"fixedCoding"`
	FixedContactPoint          ContactPoint          `json:"fixedContactPoint"`
	FixedCount                 Count                 `json:"fixedCount"`
	FixedDistance              Distance              `json:"fixedDistance"`
	FixedDuration              Duration              `json:"fixedDuration"`
	FixedHumanName             HumanName             `json:"fixedHumanName"`
	FixedIdentifier            Identifier            `json:"fixedIdentifier"`
	FixedMoney                 Money                 `json:"fixedMoney"`
	FixedPeriod                Period                `json:"fixedPeriod"`
	FixedQuantity              Quantity              `json:"fixedQuantity"`
	FixedRange                 Range                 `json:"fixedRange"`
	FixedRatio                 Ratio                 `json:"fixedRatio"`
	FixedRatioRange            RatioRange            `json:"fixedRatioRange"`
	FixedReference             Reference             `json:"fixedReference"`
	FixedSampledData           SampledData           `json:"fixedSampledData"`
	FixedSignature             Signature             `json:"fixedSignature"`
	FixedTiming                Timing                `json:"fixedTiming"`
	FixedContactDetail         ContactDetail         `json:"fixedContactDetail"`
	FixedDataRequirement       DataRequirement       `json:"fixedDataRequirement"`
	FixedExpression            Expression            `json:"fixedExpression"`
	FixedParameterDefinition   ParameterDefinition   `json:"fixedParameterDefinition"`
	FixedRelatedArtifact       RelatedArtifact       `json:"fixedRelatedArtifact"`
	FixedTriggerDefinition     TriggerDefinition     `json:"fixedTriggerDefinition"`
	FixedUsageContext          UsageContext          `json:"fixedUsageContext"`
	FixedAvailability          Availability          `json:"fixedAvailability"`
	FixedExtendedContactDetail ExtendedContactDetail `json:"fixedExtendedContactDetail"`
	FixedDosage                Dosage                `json:"fixedDosage"`
	FixedMeta                  Meta                  `json:"fixedMeta"`
}
type Type struct {
	Id            string    `json:"id"`            // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension     Extension `json:"extension"`     // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Code          string    `json:"code"`          // URL of Data type or Resource that is a(or the) type used for this element. References are URLs that are relative to http://hl7.org/fhir/StructureDefinition e.g. "string" is a reference to http://hl7.org/fhir/StructureDefinition/string. Absolute URLs are only allowed in logical models.
	Profile       string    `json:"profile"`       // Identifies a profile structure or implementation Guide that applies to the datatype this element refers to. If any profiles are specified, then the content must conform to at least one of them. The URL can be a local reference - to a contained StructureDefinition, or a reference to another StructureDefinition or Implementation Guide by a canonical URL. When an implementation guide is specified, the type SHALL conform to at least one profile defined in the implementation guide.
	TargetProfile string    `json:"targetProfile"` // Used when the type is "Reference" or "canonical", and identifies a profile structure or implementation Guide that applies to the target of the reference this element refers to. If any profiles are specified, then the content must conform to at least one of them. The URL can be a local reference - to a contained StructureDefinition, or a reference to another StructureDefinition or Implementation Guide by a canonical URL. When an implementation guide is specified, the target resource SHALL conform to at least one profile defined in the implementation guide.
	Aggregation   string    `json:"aggregation"`   // If the type is a reference to another resource, how the resource is or can be aggregated - is it a contained resource, or a reference, and if the context is a bundle, is it included in the bundle.
	Versioning    string    `json:"versioning"`    // Whether this reference needs to be version specific or version independent, or whether either can be used.
}

type Expression struct {
	Reference   string    `json:"reference"`   // A URI that defines where the expression is found.
	Id          string    `json:"id"`          // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension   Extension `json:"extension"`   // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Description string    `json:"description"` // A brief, natural language description of the condition that effectively communicates the intended semantics.
	Name        string    `json:"name"`        // A short name assigned to the expression to allow for multiple reuse of the expression in the context where it is defined.
	Language    string    `json:"language"`    // The media type of the language for the expression.
	Expression  string    `json:"expression"`  // An expression in the specified language that returns a value.
}

type ExtendedContactDetail struct {
	Organization Reference       `json:"organization"` // This contact detail is handled/monitored by a specific organization. If the name is provided in the contact, then it is referring to the named individual within this organization.
	Period       Period          `json:"period"`       // Period that this contact was valid for usage.
	Id           string          `json:"id"`           // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension    Extension       `json:"extension"`    // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Purpose      CodeableConcept `json:"purpose"`      // The purpose/type of contact.
	Name         HumanName       `json:"name"`         // The name of an individual to contact, some types of contact detail are usually blank.
	Telecom      ContactPoint    `json:"telecom"`      // The contact details application for the purpose defined.
	Address      Address         `json:"address"`      // Address for the contact.
}

type Extension struct {
	Id        string    `json:"id"`        // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension Extension `json:"extension"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Url       string    `json:"url"`       // Source of the definition for the extension code - a logical name or a URL.
	ExtensionValuex
}
type ExtensionValuex struct {
	ValueBase64Binary          string                `json:"valueBase64Binary"`
	ValueBoolean               bool                  `json:"valueBoolean"`
	ValueCanonical             string                `json:"valueCanonical"`
	ValueCode                  string                `json:"valueCode"`
	ValueDate                  string                `json:"valueDate"`
	ValueDateTime              string                `json:"valueDateTime"`
	ValueDecimal               float64               `json:"valueDecimal"`
	ValueId                    string                `json:"valueId"`
	ValueInstant               string                `json:"valueInstant"`
	ValueInteger               int                   `json:"valueInteger"`
	ValueInteger64             int64                 `json:"valueInteger64"`
	ValueMarkdown              string                `json:"valueMarkdown"`
	ValueOid                   string                `json:"valueOid"`
	ValuePositiveInt           int                   `json:"valuePositiveInt"`
	ValueString                string                `json:"valueString"`
	ValueTime                  string                `json:"valueTime"`
	ValueUnsignedInt           int                   `json:"valueUnsignedInt"`
	ValueUri                   string                `json:"valueUri"`
	ValueUrl                   string                `json:"valueUrl"`
	ValueUuid                  string                `json:"valueUuid"`
	ValueAddress               Address               `json:"valueAddress"`
	ValueAge                   Age                   `json:"valueAge"`
	ValueAnnotation            Annotation            `json:"valueAnnotation"`
	ValueAttachment            Attachment            `json:"valueAttachment"`
	ValueCodeableConcept       CodeableConcept       `json:"valueCodeableConcept"`
	ValueCodeableReference     CodeableReference     `json:"valueCodeableReference"`
	ValueCoding                Coding                `json:"valueCoding"`
	ValueContactPoint          ContactPoint          `json:"valueContactPoint"`
	ValueCount                 Count                 `json:"valueCount"`
	ValueDistance              Distance              `json:"valueDistance"`
	ValueDuration              Duration              `json:"valueDuration"`
	ValueHumanName             HumanName             `json:"valueHumanName"`
	ValueIdentifier            Identifier            `json:"valueIdentifier"`
	ValueMoney                 Money                 `json:"valueMoney"`
	ValuePeriod                Period                `json:"valuePeriod"`
	ValueQuantity              Quantity              `json:"valueQuantity"`
	ValueRange                 Range                 `json:"valueRange"`
	ValueRatio                 Ratio                 `json:"valueRatio"`
	ValueRatioRange            RatioRange            `json:"valueRatioRange"`
	ValueReference             Reference             `json:"valueReference"`
	ValueSampledData           SampledData           `json:"valueSampledData"`
	ValueSignature             Signature             `json:"valueSignature"`
	ValueTiming                Timing                `json:"valueTiming"`
	ValueContactDetail         ContactDetail         `json:"valueContactDetail"`
	ValueDataRequirement       DataRequirement       `json:"valueDataRequirement"`
	ValueExpression            Expression            `json:"valueExpression"`
	ValueParameterDefinition   ParameterDefinition   `json:"valueParameterDefinition"`
	ValueRelatedArtifact       RelatedArtifact       `json:"valueRelatedArtifact"`
	ValueTriggerDefinition     TriggerDefinition     `json:"valueTriggerDefinition"`
	ValueUsageContext          UsageContext          `json:"valueUsageContext"`
	ValueAvailability          Availability          `json:"valueAvailability"`
	ValueExtendedContactDetail ExtendedContactDetail `json:"valueExtendedContactDetail"`
	ValueDosage                Dosage                `json:"valueDosage"`
	ValueMeta                  Meta                  `json:"valueMeta"`
}

type HumanName struct {
	Id        string    `json:"id"`        // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension Extension `json:"extension"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Given     string    `json:"given"`     // Given name.
	Prefix    string    `json:"prefix"`    // Part of the name that is acquired as a title due to academic, legal, employment or nobility status, etc. and that appears at the start of the name.
	Period    Period    `json:"period"`    // Indicates the period of time when this name was valid for the named person.
	Use       string    `json:"use"`       // Identifies the purpose for this name.
	Text      string    `json:"text"`      // Specifies the entire name as it should be displayed e.g. on an application UI. This may be provided instead of or as well as the specific parts.
	Family    string    `json:"family"`    // The part of a name that links to the genealogy. In some cultures (e.g. Eritrea) the family name of a son is the first name of his father.
	Suffix    string    `json:"suffix"`    // Part of the name that is acquired as a title due to academic, legal, employment or nobility status, etc. and that appears at the end of the name.
}

type Identifier struct {
	Id        string          `json:"id"`        // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension Extension       `json:"extension"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Use       string          `json:"use"`       // The purpose of this identifier.
	Type      CodeableConcept `json:"type"`      // A coded type for the identifier that can be used to determine which identifier to use for a specific purpose.
	System    string          `json:"system"`    // Establishes the namespace for the value - that is, an absolute URL that describes a set values that are unique.
	Value     string          `json:"value"`     // The portion of the identifier typically relevant to the user and which is unique within the context of the system.
	Period    Period          `json:"period"`    // Time period during which identifier is/was valid for use.
	Assigner  Reference       `json:"assigner"`  // Organization that issued/manages the identifier.
}

type MarketingStatus struct {
	Id                string    `json:"id"`                // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension         Extension `json:"extension"`         // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	ModifierExtension Extension `json:"modifierExtension"` /*
	May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.

	Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	*/
	Country      CodeableConcept `json:"country"`      // The country in which the marketing authorization has been granted shall be specified It should be specified using the ISO 3166 ‑ 1 alpha-2 code elements.
	Jurisdiction CodeableConcept `json:"jurisdiction"` // Where a Medicines Regulatory Agency has granted a marketing authorization for which specific provisions within a jurisdiction apply, the jurisdiction can be specified using an appropriate controlled terminology The controlled term and the controlled term identifier shall be specified.
	Status       CodeableConcept `json:"status"`       // This attribute provides information on the status of the marketing of the medicinal product See ISO/TS 20443 for more information and examples.
	DateRange    Period          `json:"dateRange"`    // The date when the Medicinal Product is placed on the market by the Marketing Authorization Holder (or where applicable, the manufacturer/distributor) in a country and/or jurisdiction shall be provided A complete date consisting of day, month and year shall be specified using the ISO 8601 date format NOTE “Placed on the market” refers to the release of the Medicinal Product into the distribution chain.
	RestoreDate  string          `json:"restoreDate"`  // The date when the Medicinal Product is placed on the market by the Marketing Authorization Holder (or where applicable, the manufacturer/distributor) in a country and/or jurisdiction shall be provided A complete date consisting of day, month and year shall be specified using the ISO 8601 date format NOTE “Placed on the market” refers to the release of the Medicinal Product into the distribution chain.
}

type Meta struct {
	VersionId   string    `json:"versionId"`   // The version specific identifier, as it appears in the version portion of the URL. This value changes when the resource is created, updated, or deleted.
	LastUpdated string    `json:"lastUpdated"` // When the resource last changed - e.g. when the version changed.
	Source      string    `json:"source"`      // A uri that identifies the source system of the resource. This provides a minimal amount of [Provenance](provenance.html#) information that can be used to track or differentiate the source of information in the resource. The source may identify another FHIR server, document, message, database, etc.
	Profile     string    `json:"profile"`     // A list of profiles (references to [StructureDefinition](structuredefinition.html#) resources) that this resource claims to conform to. The URL is a reference to [StructureDefinition.url](structuredefinition-definitions.html#StructureDefinition.url).
	Security    Coding    `json:"security"`    // Security labels applied to this resource. These tags connect specific resources to the overall security policy and infrastructure.
	Tag         Coding    `json:"tag"`         // Tags applied to this resource. Tags are intended to be used to identify and relate resources to process and workflow, and applications are not required to consider the tags when interpreting the meaning of a resource.
	Id          string    `json:"id"`          // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension   Extension `json:"extension"`   // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
}

type MonetaryComponent struct {
	Code      CodeableConcept `json:"code"`      // Codes may be used to differentiate between kinds of taxes, surcharges, discounts etc.
	Factor    float64         `json:"factor"`    // Factor used for calculating this component.
	Amount    Money           `json:"amount"`    // Explicit value amount to be used.
	Id        string          `json:"id"`        // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension Extension       `json:"extension"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Type      string          `json:"type"`      // base | surcharge | deduction | discount | tax | informational.
}

type Money struct {
	Id        string    `json:"id"`        // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension Extension `json:"extension"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Value     float64   `json:"value"`     // Numerical value (with implicit precision).
	Currency  string    `json:"currency"`  // ISO 4217 Currency Code.
}

type Narrative struct {
	Id        string    `json:"id"`        // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension Extension `json:"extension"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Status    string    `json:"status"`    // The status of the narrative - whether it's entirely generated (from just the defined data or the extensions too), or whether a human authored it and it may contain additional data.
	Div       string    `json:"div"`       // The actual narrative content, a stripped down version of XHTML.
}

type ParameterDefinition struct {
	Extension     Extension `json:"extension"`     // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Use           string    `json:"use"`           // Whether the parameter is input or output for the module.
	Min           int       `json:"min"`           // The minimum number of times this parameter SHALL appear in the request or response.
	Profile       string    `json:"profile"`       // If specified, this indicates a profile that the input data must conform to, or that the output data will conform to.
	Id            string    `json:"id"`            // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Name          string    `json:"name"`          // The name of the parameter used to allow access to the value of the parameter in evaluation contexts.
	Max           string    `json:"max"`           // The maximum number of times this element is permitted to appear in the request or response.
	Documentation string    `json:"documentation"` // A brief discussion of what the parameter is for and how it is used by the module.
	Type          string    `json:"type"`          // The type of the parameter.
}

type Period struct {
	Id        string    `json:"id"`        // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension Extension `json:"extension"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Start     string    `json:"start"`     // The start of the period. The boundary is inclusive.
	End       string    `json:"end"`       // The end of the period. If the end of the period is missing, it means no end was known or planned at the time the instance was created. The start may be in the past, and the end date in the future, which means that period is expected/planned to end at that time.
}

type PrimitiveType struct {
	Id        string    `json:"id"`        // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension Extension `json:"extension"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
}

type ProductShelfLife struct {
	Type CodeableConcept `json:"type"` // This describes the shelf life, taking into account various scenarios such as shelf life of the packaged Medicinal Product itself, shelf life after transformation where necessary and shelf life after the first opening of a bottle, etc. The shelf life type shall be specified using an appropriate controlled vocabulary The controlled term and the controlled term identifier shall be specified.
	ProductShelfLifePeriodx
	SpecialPrecautionsForStorage CodeableConcept `json:"specialPrecautionsForStorage"` // Special precautions for storage, if any, can be specified using an appropriate controlled vocabulary The controlled term and the controlled term identifier shall be specified.
	Id                           string          `json:"id"`                           // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension                    Extension       `json:"extension"`                    // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	ModifierExtension            Extension       `json:"modifierExtension"`            /*
	May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.

	Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	*/
}
type ProductShelfLifePeriodx struct {
	PeriodDuration Duration `json:"periodDuration"`
	PeriodString   string   `json:"periodString"`
}

type Quantity struct {
	Value      float64   `json:"value"`      // The value of the measured amount. The value includes an implicit precision in the presentation of the value.
	Comparator string    `json:"comparator"` // How the value should be understood and represented - whether the actual value is greater or less than the stated value due to measurement issues; e.g. if the comparator is "<" , then the real value is < stated value.
	Unit       string    `json:"unit"`       // A human-readable form of the unit.
	System     string    `json:"system"`     // The identification of the system that provides the coded form of the unit.
	Code       string    `json:"code"`       // A computer processable form of the unit in some unit representation system.
	Id         string    `json:"id"`         // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension  Extension `json:"extension"`  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
}

type Range struct {
	Id        string    `json:"id"`        // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension Extension `json:"extension"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Low       Quantity  `json:"low"`       // The low limit. The boundary is inclusive.
	High      Quantity  `json:"high"`      // The high limit. The boundary is inclusive.
}

type Ratio struct {
	Id          string    `json:"id"`          // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension   Extension `json:"extension"`   // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Numerator   Quantity  `json:"numerator"`   // The value of the numerator.
	Denominator Quantity  `json:"denominator"` // The value of the denominator.
}

type RatioRange struct {
	Id            string    `json:"id"`            // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension     Extension `json:"extension"`     // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	LowNumerator  Quantity  `json:"lowNumerator"`  // The value of the low limit numerator.
	HighNumerator Quantity  `json:"highNumerator"` // The value of the high limit numerator.
	Denominator   Quantity  `json:"denominator"`   // The value of the denominator.
}

type Reference struct {
	Reference string `json:"reference"` // A reference to a location at which the other resource is found. The reference may be a relative reference, in which case it is relative to the service base URL, or an absolute URL that resolves to the location where the resource is found. The reference may be version specific or not. If the reference is not to a FHIR RESTful server, then it should be assumed to be version specific. Internal fragment references (start with '#') refer to contained resources.
	Type      string `json:"type"`      /*
	The expected type of the target of the reference. If both Reference.type and Reference.reference are populated and Reference.reference is a FHIR URL, both SHALL be consistent.

	The type is the Canonical URL of Resource Definition that is the type this reference refers to. References are URLs that are relative to http://hl7.org/fhir/StructureDefinition/ e.g. "Patient" is a reference to http://hl7.org/fhir/StructureDefinition/Patient. Absolute URLs are only allowed for logical models (and can only be used in references in logical models, not resources).
	*/
	Identifier Identifier `json:"identifier"` // An identifier for the target resource. This is used when there is no way to reference the other resource directly, either because the entity it represents is not available through a FHIR server, or because there is no way for the author of the resource to convert a known identifier to an actual location. There is no requirement that a Reference.identifier point to something that is actually exposed as a FHIR instance, but it SHALL point to a business concept that would be expected to be exposed as a FHIR instance, and that instance would need to be of a FHIR resource type allowed by the reference.
	Display    string     `json:"display"`    // Plain text narrative that identifies the resource in addition to the resource reference.
	Id         string     `json:"id"`         // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension  Extension  `json:"extension"`  // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
}

type RelatedArtifact struct {
	Citation          string          `json:"citation"`          // A bibliographic citation for the related artifact. This text SHOULD be formatted according to an accepted citation format.
	Resource          string          `json:"resource"`          // The related artifact, such as a library, value set, profile, or other knowledge resource.
	ResourceReference Reference       `json:"resourceReference"` // The related artifact, if the artifact is not a canonical resource, or a resource reference to a canonical resource.
	PublicationStatus string          `json:"publicationStatus"` // The publication status of the artifact being referred to.
	Display           string          `json:"display"`           // A brief description of the document or knowledge resource being referenced, suitable for display to a consumer.
	Extension         Extension       `json:"extension"`         // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Type              string          `json:"type"`              // The type of relationship to the related artifact.
	Classifier        CodeableConcept `json:"classifier"`        // Provides additional classifiers of the related artifact.
	Label             string          `json:"label"`             // A short label that can be used to reference the citation from elsewhere in the containing artifact, such as a footnote index.
	Document          Attachment      `json:"document"`          // The document being referenced, represented as an attachment. This is exclusive with the resource element.
	PublicationDate   string          `json:"publicationDate"`   // The date of publication of the artifact being referred to.
	Id                string          `json:"id"`                // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
}

type SampledData struct {
	Extension    Extension `json:"extension"`    // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Origin       Quantity  `json:"origin"`       // The base quantity that a measured value of zero represents. In addition, this provides the units of the entire measurement series.
	Interval     float64   `json:"interval"`     // Amount of intervalUnits between samples, e.g. milliseconds for time-based sampling.
	LowerLimit   float64   `json:"lowerLimit"`   // The lower limit of detection of the measured points. This is needed if any of the data points have the value "L" (lower than detection limit).
	Id           string    `json:"id"`           // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	IntervalUnit string    `json:"intervalUnit"` // The measurement unit in which the sample interval is expressed.
	Factor       float64   `json:"factor"`       // A correction factor that is applied to the sampled data points before they are added to the origin.
	UpperLimit   float64   `json:"upperLimit"`   // The upper limit of detection of the measured points. This is needed if any of the data points have the value "U" (higher than detection limit).
	Dimensions   int       `json:"dimensions"`   // The number of sample points at each time point. If this value is greater than one, then the dimensions will be interlaced - all the sample points for a point in time will be recorded at once.
	CodeMap      string    `json:"codeMap"`      // Reference to ConceptMap that defines the codes used in the data.
	Offsets      string    `json:"offsets"`      // A series of data points which are decimal values separated by a single space (character u20).  The units in which the offsets are expressed are found in intervalUnit.  The absolute point at which the measurements begin SHALL be conveyed outside the scope of this datatype, e.g. Observation.effectiveDateTime for a timing offset.
	Data         string    `json:"data"`         // A series of data points which are decimal values or codes separated by a single space (character u20). The special codes "E" (error), "L" (below detection limit) and "U" (above detection limit) are also defined for used in place of decimal values.
}

type Signature struct {
	When         string    `json:"when"`         // When the digital signature was signed.
	TargetFormat string    `json:"targetFormat"` // A mime type that indicates the technical format of the target resources signed by the signature.
	SigFormat    string    `json:"sigFormat"`    // A mime type that indicates the technical format of the signature. Important mime types are application/signature+xml for X ML DigSig, application/jose for JWS, and image/* for a graphical image of a signature, etc.
	Data         string    `json:"data"`         // The base64 encoding of the Signature content. When signature is not recorded electronically this element would be empty.
	Id           string    `json:"id"`           // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension    Extension `json:"extension"`    // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Type         Coding    `json:"type"`         // An indication of the reason that the entity signed this document. This may be explicitly included as part of the signature information and can be used when determining accountability for various actions concerning the document.
	Who          Reference `json:"who"`          // A reference to an application-usable description of the identity that signed  (e.g. the signature used their private key).
	OnBehalfOf   Reference `json:"onBehalfOf"`   // A reference to an application-usable description of the identity that is represented by the signature.
}

type Timing struct {
	Event             string `json:"event"` // Identifies specific times when the event occurs.
	Repeat            Repeat
	Code              CodeableConcept `json:"code"`              // A code for the timing schedule (or just text in code.text). Some codes such as BID are ubiquitous, but many institutions define their own additional codes. If a code is provided, the code is understood to be a complete statement of whatever is specified in the structured timing data, and either the code or the data may be used to interpret the Timing, with the exception that .repeat.bounds still applies over the code (and is not contained in the code).
	Id                string          `json:"id"`                // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension         Extension       `json:"extension"`         // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	ModifierExtension Extension       `json:"modifierExtension"` /*
	May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.

	Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	*/
}
type Repeat struct {
	CountMax     int     `json:"countMax"`     // If present, indicates that the count is a range - so to perform the action between [count] and [countMax] times.
	PeriodMax    float64 `json:"periodMax"`    // If present, indicates that the period is a range from [period] to [periodMax], allowing expressing concepts such as "do this once every 3-5 days.
	DayOfWeek    string  `json:"dayOfWeek"`    // If one or more days of week is provided, then the action happens only on the specified day(s).
	Count        int     `json:"count"`        // A total count of the desired number of repetitions across the duration of the entire timing specification. If countMax is present, this element indicates the lower bound of the allowed range of count values.
	When         string  `json:"when"`         // An approximate time period during the day, potentially linked to an event of daily living that indicates when the action should occur.
	Duration     float64 `json:"duration"`     // How long this thing happens for when it happens. If durationMax is present, this element indicates the lower bound of the allowed range of the duration.
	DurationMax  float64 `json:"durationMax"`  // If present, indicates that the duration is a range - so to perform the action between [duration] and [durationMax] time length.
	DurationUnit string  `json:"durationUnit"` /*
	The units of time for the duration, in UCUM units
	Normal practice is to use the 'mo' code as a calendar month when calculating the next occurrence.
	*/
	Period    float64 `json:"period"`    // Indicates the duration of time over which repetitions are to occur; e.g. to express "3 times per day", 3 would be the frequency and "1 day" would be the period. If periodMax is present, this element indicates the lower bound of the allowed range of the period length.
	TimeOfDay string  `json:"timeOfDay"` // Specified time of day for action to take place.
	Id        string  `json:"id"`        // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	TimingRepeatBoundsx
	Frequency    int    `json:"frequency"`    // The number of times to repeat the action within the specified period. If frequencyMax is present, this element indicates the lower bound of the allowed range of the frequency.
	FrequencyMax int    `json:"frequencyMax"` // If present, indicates that the frequency is a range - so to repeat between [frequency] and [frequencyMax] times within the period or period range.
	PeriodUnit   string `json:"periodUnit"`   /*
	The units of time for the period in UCUM units
	Normal practice is to use the 'mo' code as a calendar month when calculating the next occurrence.
	*/
	Offset    int       `json:"offset"`    // The number of minutes from the event. If the event code does not indicate whether the minutes is before or after the event, then the offset is assumed to be after the event.
	Extension Extension `json:"extension"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
}
type TimingRepeatBoundsx struct {
	BoundsDuration Duration `json:"boundsDuration"`
	BoundsRange    Range    `json:"boundsRange"`
	BoundsPeriod   Period   `json:"boundsPeriod"`
}

type TriggerDefinition struct {
	TriggerDefinitionTimingx
	Data              DataRequirement `json:"data"`              // The triggering data of the event (if this is a data trigger). If more than one data is requirement is specified, then all the data requirements must be true.
	Condition         Expression      `json:"condition"`         // A boolean-valued expression that is evaluated in the context of the container of the trigger definition and returns whether or not the trigger fires.
	Id                string          `json:"id"`                // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension         Extension       `json:"extension"`         // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Name              string          `json:"name"`              // A formal name for the event. This may be an absolute URI that identifies the event formally (e.g. from a trigger registry), or a simple relative URI that identifies the event in a local context.
	SubscriptionTopic string          `json:"subscriptionTopic"` // A reference to a SubscriptionTopic resource that defines the event. If this element is provided, no other information about the trigger definition may be supplied.
	Type              string          `json:"type"`              // The type of triggering event.
	Code              CodeableConcept `json:"code"`              // A code that identifies the event.
}
type TriggerDefinitionTimingx struct {
	TimingTiming    Timing    `json:"timingTiming"`
	TimingReference Reference `json:"timingReference"`
	TimingDate      string    `json:"timingDate"`
	TimingDateTime  string    `json:"timingDateTime"`
}

type UsageContext struct {
	Id        string    `json:"id"`        // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension Extension `json:"extension"` // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Code      Coding    `json:"code"`      // A code that identifies the type of context being specified by this usage context.
	UsageContextValuex
}
type UsageContextValuex struct {
	ValueCodeableConcept CodeableConcept `json:"valueCodeableConcept"`
	ValueQuantity        Quantity        `json:"valueQuantity"`
	ValueRange           Range           `json:"valueRange"`
	ValueReference       Reference       `json:"valueReference"`
}

type VirtualServiceDetail struct {
	MaxParticipants int       `json:"maxParticipants"` // Maximum number of participants supported by the virtual service.
	SessionKey      string    `json:"sessionKey"`      // Session Key required by the virtual service.
	Id              string    `json:"id"`              // Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Extension       Extension `json:"extension"`       // May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	ChannelType     Coding    `json:"channelType"`     // The type of virtual service to connect to (i.e. Teams, Zoom, Specific VMR technology, WhatsApp).
	VirtualServiceDetailAddressx
	AdditionalInfo string `json:"additionalInfo"` // Address to see alternative connection details.
}
type VirtualServiceDetailAddressx struct {
	AddressUrl                   string                `json:"addressUrl"`
	AddressString                string                `json:"addressString"`
	AddressContactPoint          ContactPoint          `json:"addressContactPoint"`
	AddressExtendedContactDetail ExtendedContactDetail `json:"addressExtendedContactDetail"`
}
