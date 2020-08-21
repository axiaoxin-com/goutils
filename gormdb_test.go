package goutils

import (
	"os"
	"testing"
)

func TestNewGormSQLite3(t *testing.T) {
	dbname := "./db.sqlite3"
	conf := SQLite3Config{
		DBName: dbname,
	}
	db, err := NewGormSQLite3(conf)
	if err != nil {
		t.Error("new gorm sqlite3 return error:", err)
	}
	defer db.Close()
	defer os.Remove(dbname)
}

func TestNewGormMySQL(t *testing.T) {
	conf := MySQLConfig{
		Host:     "localhost",
		Port:     3306,
		Username: "root",
		Password: "roooooot",
		DBName:   "information_schema",
	}
	db, err := NewGormMySQL(conf)
	if err != nil {
		t.Error("new gorm mysql return error:", err)
	}
	defer db.Close()
}
