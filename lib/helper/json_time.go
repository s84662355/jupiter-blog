package helper

import (
	"database/sql/driver"
	"fmt"
	"github.com/uniplaces/carbon"
	"strconv"
	"strings"
	"time"
)

// JSONTime format json time field by myself
type JSONTime struct {
	*carbon.Carbon
}

// MarshalJSON on JSONTime format Time field with %Y-%m-%d %H:%M:%S
func (t JSONTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.DateTimeString())
	return []byte(formatted), nil
}

// MarshalJSON on JSONTime format Time field with %Y-%m-%d %H:%M:%S
func (t *JSONTime) UnmarshalJSON(data []byte) error {
	t0 := string(data)
	t0 = strings.Replace(t0, "\"", "", -1)
	if len(t0) > 10 {
		tm2, _ := time.Parse("2006-01-02 15:04:05", t0)
		//t.Time = tm2
		t.Carbon = carbon.NewCarbon(tm2)
	} else {
		tm2, _ := time.Parse("2006-01-02", t0)
		//t.Time = tm2
		t.Carbon = carbon.NewCarbon(tm2)
	}
	return nil
}

// Value insert timestamp into mysql need this function.
func (t JSONTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Carbon.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Carbon.Time, nil
}

func (t JSONTime) String() string {
	return t.DateTimeString()
}

func (t JSONTime) Int(ss string) uint {
	dataInt, _ := strconv.Atoi(t.Format(ss))
	return uint(dataInt)
}

// Scan valueof time.Time
func (t *JSONTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = JSONTime{Carbon: carbon.NewCarbon(value)}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (t JSONTime) Create() *JSONTime {
	tt := JSONTime{}
	tt.Scan(time.Now())
	return &tt
}

//时间戳 秒
func (t *JSONTime) Unix() int64 {
	return t.Carbon.Timestamp()
}
