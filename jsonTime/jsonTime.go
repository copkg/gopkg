package jsonTime

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const (
	TimeFormat = "2006-01-02 15:04:05"
)

type JsonTime struct {
	time.Time
}

func (j JsonTime) MarshalJSON() ([]byte, error) {
	if j.Time.IsZero() {
		return []byte(`""`), nil
	}
	return []byte(`"` + j.Time.Format(TimeFormat) + `"`), nil
}

func (j JsonTime) String() string {

	return j.Time.Format(TimeFormat)
}
func (j JsonTime) IsZero() bool {
	if j.Time.IsZero() {
		return true
	}
	return false
}
func (j JsonTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if j.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return j.Time, nil
}

func (j *JsonTime) Scan(v interface{}) error {
	if v == nil {
		*j = JsonTime{Time: time.Time{}}
		return nil
	}
	value, ok := v.(time.Time)
	if ok {
		*j = JsonTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (j *JsonTime) UnmarshalJSON(data []byte) (err error) {
	// 空值不进行解析
	if len(data) == 2 {
		*j = JsonTime{Time: time.Time{}}
		return
	}
	// 指定解析的格式
	now, err := time.Parse(`"`+TimeFormat+`"`, string(data))
	*j = JsonTime{Time: now}
	return
}
