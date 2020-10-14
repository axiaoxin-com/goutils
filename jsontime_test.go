package goutils

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"gorm.io/gorm"
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

func TestJSONTimeInGORM(t *testing.T) {
	dbname := "./db.jsontime"
	conf := DBConfig{
		DBName: dbname,
	}
	db, err := NewGormSQLite3(conf)
	if err != nil {
		t.Fatal("new test db return error:", err)
	}
	sqlDB, err := db.DB()
	defer sqlDB.Close()
	defer os.Remove(dbname)

	type jsonTimeNow struct {
		gorm.Model
		Now JSONTime
	}
	type timeNow struct {
		gorm.Model
		Now time.Time
	}

	recordNow := timeNow{}
	recordJNow := jsonTimeNow{}
	db.AutoMigrate(&recordNow, &recordJNow)

	now := time.Now()
	recordNow.Now = now
	recordJNow.Now = NewJSONTime(now)
	db.Create(&recordNow)
	db.Create(&recordJNow)

	db.First(&recordNow)
	db.First(&recordJNow)
	b, _ := json.Marshal(recordNow)
	b1, _ := json.Marshal(recordJNow)
	if string(b) == string(b1) {
		t.Error("JSONTime not work with Gorm")
	}
}
