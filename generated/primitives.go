// TODO this file is not actually generated, do we want to generate it? does it belong here?
package generated

import (
	"encoding/base64"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Date time.Time

func (d *Date) UnmarshalJSON(b []byte) error {
	log.Println("Unmarshalling date")

	// Unmarshal from string, format (YYYY(-MM(-DD)?)?)
	s := string(b)
	// FIXME this feels very, very stupid, both the quotes and the multiple parses
	date, err := time.Parse("\"2006-01-02\"", s)
	if err != nil {
		date, err = time.Parse("\"2006-01\"", s)
		if err != nil {
			date, err = time.Parse("\"2006\"", s)
			if err != nil {
				return err
			}
		}
	}
	*d = Date(date)
	return nil
}

func (d *Date) MarshalJSON() ([]byte, error) {
	log.Println("Marshalling date")
	return []byte("\"" + time.Time(*d).Format("2006-01-02") + "\""), nil
}

func (d *Date) MarshalBSON() ([]byte, error) {
	var timeMirror struct {
		primitive.DateTime `bson:"date"`
	}
	timeMirror.DateTime = primitive.NewDateTimeFromTime(time.Time(*d))
	return bson.Marshal(timeMirror)
}

func (d *Date) UnmarshalBSON(b []byte) error {
	var timeMirror struct {
		primitive.DateTime `bson:"date"`
	}
	err := bson.Unmarshal(b, &timeMirror)
	if err != nil {
		return err
	}
	*d = Date(timeMirror.Time())
	return nil
}

type Time time.Time

func (t *Time) UnmarshalJSON(b []byte) error {
	// Regex for time is ([01][0-9]|2[0-3]):[0-5][0-9]:([0-5][0-9]|60)(\.[0-9]{1,9})?
	s := string(b)
	tm, err := time.Parse("\"15:04:05\"", s)
	if err != nil {
		tm, err = time.Parse("\"15:04\"", s)
		if err != nil {
			tm, err = time.Parse("\"15\"", s)
			if err != nil {
				return err
			}
		}
	}
	*t = Time(tm)
	return nil
}

func (t *Time) MarshalJSON() ([]byte, error) {
	return []byte("\"" + time.Time(*t).Format("15:04:05") + "\""), nil
}

func (t *Time) MarshalBSON() ([]byte, error) {
	var timeMirror struct {
		primitive.DateTime `bson:"time"`
	}
	timeMirror.DateTime = primitive.NewDateTimeFromTime(time.Time(*t))
	return bson.Marshal(timeMirror)
}

func (t *Time) UnmarshalBSON(b []byte) error {
	var timeMirror struct {
		primitive.DateTime `bson:"time"`
	}
	err := bson.Unmarshal(b, &timeMirror)
	if err != nil {
		return err
	}
	*t = Time(timeMirror.Time())
	return nil
}

type DateTime time.Time

func (dt *DateTime) UnmarshalJSON(b []byte) error {
	// Regex is ([0-9]([0-9]([0-9][1-9]|[1-9]0)|[1-9]00)|[1-9]000)(-(0[1-9]|1[0-2])(-(0[1-9]|[1-2][0-9]|3[0-1])
	// (T([01][0-9]|2[0-3]):[0-5][0-9]:([0-5][0-9]|60)(\.[0-9]{1,9})?)?)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-
	// 5][0-9]|14:00)?)?)?

	s := string(b)
	tm, err := time.Parse("\"2006-01-02T15:04:05\"", s)
	if err != nil {
		tm, err = time.Parse("\"2006-01-02\"", s)
		if err != nil {
			tm, err = time.Parse("\"2006-01\"", s)
			if err != nil {
				tm, err = time.Parse("\"2006\"", s)
				if err != nil {
					return err
				}
			}
		}
	}
	*dt = DateTime(tm)
	return nil
}

func (dt *DateTime) MarshalJSON() ([]byte, error) {
	return []byte("\"" + time.Time(*dt).Format("2006-01-02T15:04:05") + "\""), nil
}

func (dt *DateTime) MarshalBSON() ([]byte, error) {
	var timeMirror struct {
		primitive.DateTime `bson:"datetime"`
	}
	timeMirror.DateTime = primitive.NewDateTimeFromTime(time.Time(*dt))
	return bson.Marshal(timeMirror)
}

func (dt *DateTime) UnmarshalBSON(b []byte) error {
	var timeMirror struct {
		primitive.DateTime `bson:"datetime"`
	}
	err := bson.Unmarshal(b, &timeMirror)
	if err != nil {
		return err
	}
	*dt = DateTime(timeMirror.Time())
	return nil
}

type Base64Binary []byte

func (b *Base64Binary) UnmarshalJSON(data []byte) error {
	// Decode base64 and remove leading and trailing quotes
	str := string(data)
	if str[0] != '"' || str[len(str)-1] != '"' {
		return errors.New("invalid base64 string")
	}

	decoded, err := base64.StdEncoding.DecodeString(str[1 : len(str)-1])
	if err != nil {
		return err
	}
	*b = Base64Binary(decoded)
	return nil
}

func (b *Base64Binary) MarshalJSON() ([]byte, error) {
	return []byte("\"" + base64.StdEncoding.EncodeToString(*b) + "\""), nil
}
