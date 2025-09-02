package jsonTime

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const (
	DateFormat = "2006-01-02"
)

type JsonDate struct {
	time.Time
}

func (j JsonDate) MarshalJSON() ([]byte, error) {
	if j.Time.IsZero() {
		return []byte(`""`), nil
	}
	return []byte(`"` + j.Time.Format(DateFormat) + `"`), nil
}

func (j JsonDate) String() string {
	return j.Time.Format(DateFormat)
}
func (j JsonDate) IsZero() bool {
	if j.Time.IsZero() {
		return true
	}
	return false
}
func (j JsonDate) Value() (driver.Value, error) {
	var zeroTime time.Time
	if j.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return j.Time, nil
}

func (j *JsonDate) Scan(v interface{}) error {
	if v == nil {
		*j = JsonDate{Time: time.Time{}}
		return nil
	}
	if value, ok := v.(JsonDate); ok {
		now, err := time.Parse(DateFormat, value.String())
		*j = JsonDate{Time: now}
		return err
	}
	if value, ok := v.(time.Time); ok {
		*j = JsonDate{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (j *JsonDate) UnmarshalJSON(data []byte) (err error) {
	// 空值不进行解析
	if len(data) == 2 {
		*j = JsonDate{Time: time.Time{}}
		return
	}
	// 指定解析的格式
	now, err := time.Parse(`"`+DateFormat+`"`, string(data))
	*j = JsonDate{Time: now}
	return
}
