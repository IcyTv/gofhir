package generated

import (
	"encoding/base64"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// func GenerateMongoVersion(resourceType string, id primitive.ObjectID) {
// 	coll := db.FhirDatabase.Collection(resourceType)
// 	revColl := db.FhirDatabase.Collection(resourceType + "_rev")

// 	// What we want to do:
// 	// 1. Find the current version of the resource (in the resource collection)
// 	// 2. Insert a new version of the resource (in the resource_rev collection) (automatically increment the version and add a backreference to the new document)
// 	// 3. Update the current version of the resource (in the resource collection) (automatically increment the version and add a backreference to the new document)

// 	// Some important notes:
// 	// - The version field is a nested field in the "meta" structure, so "meta.version" == <documentversion>

// 	c

// 	mongo.Pipeline{}
// }

type PrimitiveExtension struct {
	Id        string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension *Extension `bson:"extension,omitempty" json:"extension,omitempty"`
}

// A date, or partial date (e.g. just year or year + month) as used in human communication. The format is YYYY, YYYY-MM, or YYYY-MM-DD, e.g. 2018, 1973-06, or 1905-08-23. There SHALL be no timezone offset. Dates SHALL be valid dates.
type Date time.Time

func (d *Date) UnmarshalJSON(b []byte) error {
	asString := string(b)

	if asString == "null" {
		return nil
	}

	t, err := time.Parse("2006-01-02", asString[1:len(asString)-1])
	if err != nil {
		t, err = time.Parse("2006-01", asString[1:len(asString)-1])
		if err != nil {
			t, err = time.Parse("2006", asString[1:len(asString)-1])
			if err != nil {
				return err
			}
		}
	}

	*d = Date(t)
	return nil
}

func (d *Date) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(*d).Format("2006-01-02")), nil
}

// A time during the day, in the format hh:mm:ss. There is no date specified. Seconds must be provided due to schema type constraints but may be zero-filled and may be ignored at receiver discretion. The time "24:00" SHALL NOT be used. A timezone offset SHALL NOT be present. Times can be converted to a Duration since midnight.
type Time time.Time

func (o *Time) UnmarshalJSON(b []byte) error {
	asString := string(b)

	if asString == "null" {
		return nil
	}

	t, err := time.Parse("15:04:05.999999999", asString[1:len(asString)-1])
	if err != nil {
		t, err = time.Parse("15:04:05", asString[1:len(asString)-1])
		if err != nil {
			t, err = time.Parse("15:04", asString[1:len(asString)-1])
			if err != nil {
				return err
			}
		}
	}

	*o = Time(t)
	return nil
}

func (t *Time) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(*t).Format("15:04:05.999999999")), nil
}

// A date, date-time or partial date (e.g. just year or year + month) as used in human communication. The format is YYYY, YYYY-MM, YYYY-MM-DD or YYYY-MM-DDThh:mm:ss+zz:zz, e.g. 2018, 1973-06, 1905-08-23, 2015-02-07T13:28:17-05:00 or 2017-01-01T00:00:00.000Z. If hours and minutes are specified, a timezone offset SHALL be populated. Actual timezone codes can be sent using the Timezone Code extension, if desired. Seconds must be provided due to schema type constraints but may be zero-filled and may be ignored at receiver discretion. Milliseconds are optionally allowed. Dates SHALL be valid dates. The time "24:00" is not allowed. Leap Seconds are allowed - see below

// type DateTime time.Time

// func (d *DateTime) UnmarshalJSON(b []byte) error {
// 	log.Println("Unmarshalling DateTime", b)
// 	asString := string(b)

// 	if asString == "null" || asString == "" || asString == "\"\"" {
// 		return nil
// 	}

// 	//TODO this is stupid...
// 	t, err := time.Parse("2006-01-02T15:04:05.999999999-07:00", asString[1:len(asString)-1])
// 	if err != nil {
// 		t, err = time.Parse("2006-01-02T15:04:05-07:00", asString[1:len(asString)-1])
// 		if err != nil {
// 			t, err = time.Parse("2006-01-02T15:04:05.999999999Z", asString[1:len(asString)-1])
// 			if err != nil {
// 				t, err = time.Parse("2006-01-02T15:04:05Z", asString[1:len(asString)-1])
// 				if err != nil {
// 					t, err = time.Parse("2006-01-02T15:04:05.999999999", asString[1:len(asString)-1])
// 					if err != nil {
// 						t, err = time.Parse("2006-01-02T15:04:05", asString[1:len(asString)-1])
// 						if err != nil {
// 							t, err = time.Parse("2006-01-02T15:04", asString[1:len(asString)-1])
// 							if err != nil {
// 								t, err = time.Parse("2006-01-02", asString[1:len(asString)-1])
// 								if err != nil {
// 									t, err = time.Parse("2006-01", asString[1:len(asString)-1])
// 									if err != nil {
// 										t, err = time.Parse("2006", asString[1:len(asString)-1])
// 										if err != nil {
// 											return err
// 										}
// 									}
// 								}
// 							}
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}

// 	*d = DateTime(t)
// 	return nil
// }

// func (d *DateTime) MarshalJSON() ([]byte, error) {
// 	return []byte(time.Time(*d).Format("2006-01-02T15:04:05.999999999-07:00")), nil
// }

// func (d *DateTime) UnmarshalBSON(b []byte) error {
// 	// TODO this is stupid, but test data basically demands it...
// 	var AsString string
// 	err := bson.Unmarshal(b, &AsString)
// 	log.Println("Unmarshalling DateTime as string", AsString, err)
// 	if err == nil {
// 		log.Printf("Unmarshalling DateTime as string '%s'", AsString)
// 		t, err := time.Parse("2006-01-02", AsString)
// 		if err == nil {
// 			*d = DateTime(t)
// 			return nil
// 		}
// 	}

// 	var t primitive.DateTime
// 	err = bson.Unmarshal(b, &t)
// 	if err != nil {
// 		return err
// 	}

// 	*d = DateTime(t.Time())
// 	return nil
// }

// func (d *DateTime) MarshalBSON() ([]byte, error) {
// 	return bson.Marshal(primitive.NewDateTimeFromTime(time.Time(*d)))
// }

type DateTime string

type Base64Binary []byte

func (o Base64Binary) UnmarshalJSON(b []byte) error {
	_, err := base64.StdEncoding.Decode(o, b)
	return err
}

func (o Base64Binary) MarshalJSON() ([]byte, error) {
	return []byte(base64.StdEncoding.EncodeToString(o)), nil
}

func CreateCollections(db *mongo.Database) error {
	err := othersCreateCollections(db)
	if err != nil {
		return err
	}
	err = resourcesCreateCollections(db)
	if err != nil {
		return err
	}

	return nil
}

const (
	PreferReturnMinimal = "return=minimal"
	PreferReturnRepres  = "return=representation"
	PreferReturnOper    = "return=operationOutcome"
)

type FHIRHeader struct {
	Prefer string `form:"default=return=minimal"`
}
