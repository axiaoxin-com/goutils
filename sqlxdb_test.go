package goutils

import (
	"os"
	"testing"
)

func TestNewSqlxQLite3(t *testing.T) {
	dbname := "./db.sqlite3"
	conf := SQLite3Config{
		DBName: dbname,
	}
	db, err := NewSqlxSQLite3(conf)
	if err != nil {
		t.Error("new sqlx sqlite3 return error:", err)
	}
	defer db.Close()
	defer os.Remove(dbname)
}

func TestNewSqlxMySQL(t *testing.T) {
	conf := MySQLConfig{
		Host:     "localhost",
		Port:     3306,
		Username: "root",
		Password: "roooooot",
		DBName:   "information_schema",
	}
	db, err := NewSqlxMySQL(conf)
	if err != nil {
		t.Error("new sqlx mysql return error:", err)
	}
	defer db.Close()
}
