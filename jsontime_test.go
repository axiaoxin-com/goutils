package goutils

import (
	"encoding/json"
	"testing"
	"time"
)

func TestJSONTime(t *testing.T) {

	type TimeFieldStruct struct {
		TimeField JSONTime
	}

	timeValue := "2020-08-18 09:55:00"
	timeObj, _ := time.Parse("2006-01-02 15:04:05", timeValue)
	jsonTime := NewJSONTime(timeObj)

	s := TimeFieldStruct{
		TimeField: jsonTime,
	}

	b, _ := json.Marshal(s)
	if string(b) != `{"TimeField":"2020-08-18 09:55:00"}` {
		t.Error("wrong JSONTime field format:", string(b))
	}
}
