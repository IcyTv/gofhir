package interpreter

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var dateTimeTypeSpec = newAnyTypeSpec("DateTime")
var dateTimeTypeInfo = NewSimpleTypeInfo(namespaceNameString, NewString("DateTime"), anyTypeNameString)

var timeZoneOffsetRegexp = regexp.MustCompile("^([+-])(\\d{1,2})(?::(\\d{1,2}))$")
var dateTimeRegexp = regexp.MustCompile("^(\\d(?:\\d(?:\\d[1-9]|[1-9]0)|[1-9]00)|[1-9]000)(?:-(0[1-9]|1[0-2])(?:-(0[1-9]|[1-2]\\d|3[0-1]))?)?(?:T(?:([01]\\d|2[0-3])(?::([0-5]\\d)(?::([0-5]\\d|60)(?:\\.(\\d+))?)?)(Z|[+-](?:(?:0\\d|1[0-3]):[0-5]\\d|14:00))?)?)?$")

type dateTimeType struct {
	temporalType
	value time.Time
}

type DateTimeAccessor interface {
	DateTemporalAccessor
	Hour() int
	Minute() int
	Second() int
	Nanosecond() int
	Location() *time.Location
}

func NewDateTime(value time.Time) DateTimeAccessor {
	return NewDateTimeWithSource(value, nil)
}

func NewDateTimeWithSource(value time.Time, source interface{}) DateTimeAccessor {
	return newDateTime(value, NanoTimePrecision, source)
}
func NewDateTimeYMDHMSNWithPrecision(year, month, day, hour, minute, second, nanosecond int, loc *time.Location, precision DateTimePrecisions) DateTimeAccessor {
	return NewDateTimeYMDHMSNWithPrecisionAndSource(year, month, day, hour, minute, second, nanosecond, loc, precision, nil)
}

func NewDateTimeYMDHMSNWithPrecisionAndSource(year, month, day, hour, minute, second, nanosecond int, loc *time.Location, precision DateTimePrecisions, source interface{}) DateTimeAccessor {
	if precision < YearDatePrecision || precision > NanoTimePrecision {
		log.Fatalf("Invalid dat/time precision: %d", precision)
	}

	if precision < MonthDatePrecision {
		month = 1
	}
	if precision < DayDatePrecision {
		day = 1
	}
	if precision < HourTimePrecision {
		hour = 0
	}
	if precision < MinuteTimePrecision {
		minute = 0
	}
	if precision < SecondTimePrecision {
		second = 0
	}
	if precision < NanoTimePrecision {
		nanosecond = 0
	}

	return newDateTime(time.Date(year, time.Month(month), day, hour, minute, second, nanosecond, loc), precision, source)
}

func ParseDateTime(value string) (DateTimeAccessor, error) {
	return ParseDateTimeWithSource(value, nil)
}

func ParseDateTimeWithSource(value string, source interface{}) (DateTimeAccessor, error) {
	parts := dateTimeRegexp.FindStringSubmatch(value)
	if parts == nil {
		return nil, fmt.Errorf("Not a valid fluent date/time string: %s", value)
	}
	return newDateTimeFromParts(parts, source), nil
}

func newDateTimeFromParts(parts []string, source interface{}) DateTimeAccessor {
	year, _ := strconv.Atoi(parts[1])
	precision := YearDatePrecision

	month := 1
	if parts[2] != "" {
		month, _ = strconv.Atoi(parts[2])
		precision = MonthDatePrecision
	}

	day := 1
	if parts[3] != "" {
		day, _ = strconv.Atoi(parts[3])
		precision = DayDatePrecision
	}

	hour := 0
	if parts[4] != "" {
		hour, _ = strconv.Atoi(parts[4])
		precision = HourTimePrecision
	}

	minute := 0
	if parts[5] != "" {
		minute, _ = strconv.Atoi(parts[5])
		precision = MinuteTimePrecision
	}

	second := 0
	if parts[6] != "" {
		second, _ = strconv.Atoi(parts[6])
		precision = SecondTimePrecision
	}

	nanosecond := 0
	if parts[7] != "" {
		nanosecond = parseNanosecond(parts[7])
		precision = NanoTimePrecision
	}

	location := mustEvalLocation(parts[8])
	value := time.Date(year, time.Month(month), day, hour, minute, second, nanosecond, location)

	return newDateTime(value, precision, source)
}

func newDateTime(value time.Time, precision DateTimePrecisions, source interface{}) DateTimeAccessor {
	return &dateTimeType{
		temporalType{
			baseAnyType{source},
			precision,
		},
		value,
	}
}

func mustEvalLocation(value string) *time.Location {
	if value == "" {
		return time.Local
	}
	if value == "Z" {
		return time.UTC
	}

	parts := timeZoneOffsetRegexp.FindStringSubmatch(value)
	if parts == nil {
		log.Fatalf("Invalid time zone offset: %s", value)
	}

	hours, _ := strconv.Atoi(parts[2])
	offset := hours * 60 * 60

	if parts[3] != "" {
		minutes, _ := strconv.Atoi(parts[3])
		offset += minutes * 60
	}

	if parts[1] == "-" {
		offset = -offset
	}

	if offset == 0 {
		return time.UTC
	}

	return time.FixedZone(fmt.Sprint(offset), offset)
}

func (t *dateTimeType) DataType() DataTypes {
	return DateTimeDataType
}

func (t *dateTimeType) Date() DateAccessor {
	precision := t.precision
	if precision > DayDatePrecision {
		precision = DayDatePrecision
	}
	return NewDateYMDWithPrecision(t.Year(), t.Month(), t.Day(), precision)
}

func (t *dateTimeType) DateTime() DateTimeAccessor {
	return t
}

func (t *dateTimeType) Time() time.Time {
	return t.value
}

func (t *dateTimeType) Year() int {
	return t.value.Year()
}

func (t *dateTimeType) Month() int {
	return int(t.value.Month())
}

func (t *dateTimeType) Day() int {
	return t.value.Day()
}

func (t *dateTimeType) Hour() int {
	return t.value.Hour()
}

func (t *dateTimeType) Minute() int {
	return t.value.Minute()
}

func (t *dateTimeType) Second() int {
	return t.value.Second()
}

func (t *dateTimeType) Nanosecond() int {
	return t.value.Nanosecond()
}

func (t *dateTimeType) Location() *time.Location {
	return t.value.Location()
}

func (t *dateTimeType) TypeSpec() TypeSpecAccessor {
	return dateTimeTypeSpec
}

func (t *dateTimeType) TypeInfo() TypeInfoAccessor {
	return dateTimeTypeInfo
}

func (t *dateTimeType) LowestPrecision() DateTimePrecisions {
	return YearDatePrecision
}

func (t *dateTimeType) Equal(node interface{}) bool {
	if o, ok := node.(DateTimeAccessor); !ok {
		return false
	} else {
		return TemporalPrecisionEqual(t, o) && dateTimeValueEqual(t, o)
	}
}

func (t *dateTimeType) Equivalent(node interface{}) bool {
	if o, ok := node.(DateTimeAccessor); !ok {
		return false
	} else {
		return dateTimeValueEqual(t, o)
	}
}
func dateTimeValueEqual(dt1 DateTemporalAccessor, dt2 DateTemporalAccessor) bool {
	return dt1.Time().Equal(dt2.Time())
}

func (t *dateTimeType) Compare(comparator Comparator) (int, OperatorStatus) {
	if !TypeEqual(t, comparator) {
		return -1, Inconvertible
	}

	o := comparator.(DateTimeAccessor)
	if !TemporalPrecisionEqual(t, o) {
		return -1, Empty
	}

	left, right := t.value.UnixNano(), o.Time().UnixNano()
	if left > right {
		return 1, Evaluated
	} else if left < right {
		return -1, Evaluated
	}

	return 0, Evaluated
}

func (t *dateTimeType) Add(quantity QuantityAccessor) (TemporalAccessor, error) {
	value, precision, err := quantityDateTimePrecision(quantity)
	if err != nil {
		return nil, err
	}

	res, err := addQuantityTemporalDuration(t, value, precision)
	if err != nil {
		return nil, err
	}

	return NewDateTimeYMDHMSNWithPrecisionAndSource(res.Year(), int(res.Month()), res.Day(), res.Hour(), res.Minute(), res.Second(), res.Nanosecond(), res.Location(), precision, nil), nil
}

func (t *dateTimeType) String() string {
	var b strings.Builder
	b.Grow(39)

	writeStringBuilderInt(&b, t.value.Year(), 4)
	if t.precision >= MonthDatePrecision {
		b.WriteByte('-')
		writeStringBuilderInt(&b, int(t.value.Month()), 2)
	}
	if t.precision >= DayDatePrecision {
		b.WriteByte('-')
		writeStringBuilderInt(&b, t.value.Day(), 2)
	}
	if t.precision >= HourTimePrecision {
		b.WriteByte('T')
		writeStringBuilderInt(&b, t.value.Hour(), 2)
	}
	if t.precision >= MinuteTimePrecision {
		b.WriteByte(':')
		writeStringBuilderInt(&b, t.value.Minute(), 2)
	}
	if t.precision >= SecondTimePrecision {
		b.WriteByte(':')
		writeStringBuilderInt(&b, t.value.Second(), 2)
	}
	if t.precision >= NanoTimePrecision {
		b.WriteByte('.')
		writeStringBuilderInt(&b, t.value.Nanosecond(), 9)
	}
	if t.precision >= HourTimePrecision {
		_, offset := t.value.Zone()
		if offset >= 0 {
			b.WriteByte('+')
		} else {
			b.WriteByte('-')
			offset = -offset
		}
		writeStringBuilderInt(&b, offset/3600, 2)
		b.WriteByte(':')
		writeStringBuilderInt(&b, (offset%3600)/60, 2)
	}

	return b.String()
}
