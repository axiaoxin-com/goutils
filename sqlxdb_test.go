package goutils

import (
	"os"
	"testing"
)

func TestNewSqlxQLite3(t *testing.T) {
	dbname := "./db.sqlite3"
	db, err := NewSqlxSQLite3(dbname, 10, 10, 10)
	if err != nil {
		t.Error("new sqlx sqlite3 return error:", err)
	}
	defer db.Close()
	defer os.Remove(dbname)
}

func TestNewSqlxMySQL(t *testing.T) {
	db, err := NewSqlxMySQL("localhost", 3306, "root", "roooooot", "information_schema", 10, 10, 10, 3, 3, 3)
	if err != nil {
		t.Error("new sqlx mysql return error:", err)
	}
	defer db.Close()
}
