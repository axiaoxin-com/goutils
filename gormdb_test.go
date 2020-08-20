package goutils

import (
	"os"
	"testing"
)

func TestNewGormSQLite3(t *testing.T) {
	dbname := "./db.sqlite3"
	db, err := NewGormSQLite3(dbname, false, 10, 10, 10)
	if err != nil {
		t.Error("new gorm sqlite3 return error:", err)
	}
	defer db.Close()
	defer os.Remove(dbname)
}

func TestNewGormMySQL(t *testing.T) {
	db, err := NewGormMySQL("localhost", 3306, "root", "roooooot", "information_schema", false, 10, 10, 10, 3, 5, 5)
	if err != nil {
		t.Error("new gorm mysql return error:", err)
	}
	defer db.Close()
}
