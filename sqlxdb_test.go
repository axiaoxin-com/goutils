package goutils

import (
	"os"
	"sync"
	"testing"

	"github.com/spf13/viper"
)

func TestNewSqlxQLite3(t *testing.T) {
	dbname := "./db.sqlite3"
	conf := DBConfig{
		DBName: dbname,
	}
	db, err := NewSqlxSQLite3(conf)
	if err != nil {
		t.Error("new sqlx sqlite3 return error:", err)
	}
	defer db.Close()
	defer os.Remove(dbname)
	if err := db.Ping(); err != nil {
		t.Error(err)
	}
}

func TestNewSqlxMySQL(t *testing.T) {
	conf := DBConfig{
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
	if err := db.Ping(); err != nil {
		t.Error(err)
	}
}

func TestSqlxMySQL(t *testing.T) {
	viper.SetDefault("mysql.unittest.host", "127.0.0.1")
	viper.SetDefault("mysql.unittest.port", 3306)
	viper.SetDefault("mysql.unittest.username", "root")
	viper.SetDefault("mysql.unittest.password", "roooooot")
	viper.SetDefault("mysql.unittest.dbname", "information_schema")
	if db, err := SqlxMySQL("unittest"); err != nil {
		t.Fatal(err)
	} else if db == nil {
		t.Fatal("db is nil")
	} else if err := db.Ping(); err != nil {
		t.Error(err)
	}
	defer CloseSqlxInstances()
	if _, err := SqlxMySQL("unittest"); err != nil {
		t.Error(err)
	}
	viper.SetDefault("mysql.localhost.host", "127.0.0.1")
	viper.SetDefault("mysql.localhost.port", 3306)
	viper.SetDefault("mysql.localhost.username", "root")
	viper.SetDefault("mysql.localhost.password", "roooooot")
	viper.SetDefault("mysql.localhost.dbname", "information_schema")
	if _, err := SqlxMySQL("localhost"); err != nil {
		t.Error(err)
	}
	instanceCount := 0
	mysqlCount := 0
	SqlxInstances.Range(func(k, v interface{}) bool {
		instanceCount++
		if k.(string) == "mysql" {
			t.Logf("k:%#v v:%#v\n", k, v)
			v.(*sync.Map).Range(func(kk, vv interface{}) bool {
				mysqlCount++
				t.Logf("kk:%#v vv:%#v\n", kk, vv)
				return true
			})
		}
		return true
	})
	if instanceCount != 1 {
		t.Error("instanceCount != 1, ", instanceCount)
	}
	if mysqlCount != 2 {
		t.Error("mysqlCount != 2, ", mysqlCount)
	}
}

func TestSqlxSQLite3(t *testing.T) {
	dbname := "db.sqlite3"
	viper.SetDefault("sqlite3.unittest.dbname", dbname)
	if db, err := SqlxSQLite3("unittest"); err != nil {
		t.Fatal(err)
	} else if db == nil {
		t.Fatal("db is nil")
	} else if err := db.Ping(); err != nil {
		t.Error(err)
	}
	defer CloseSqlxInstances()
	defer os.Remove(dbname)
	if _, err := SqlxSQLite3("unittest"); err != nil {
		t.Error(err)
	}
	viper.SetDefault("sqlite3.other.dbname", "other."+dbname)
	if _, err := SqlxSQLite3("other"); err != nil {
		t.Error(err)
	}
	defer os.Remove("other." + dbname)
	instanceCount := 0
	sqlite3Count := 0
	SqlxInstances.Range(func(k, v interface{}) bool {
		instanceCount++
		if k.(string) == "sqlite3" {
			t.Logf("k:%#v v:%#v\n", k, v)
			v.(*sync.Map).Range(func(kk, vv interface{}) bool {
				sqlite3Count++
				t.Logf("kk:%#v vv:%#v\n", kk, vv)
				return true
			})
		}
		return true
	})
	if instanceCount != 1 {
		t.Error("instanceCount != 1, ", instanceCount)
	}
	if sqlite3Count != 2 {
		t.Error("sqlite3Count != 2, ", sqlite3Count)
	}
}

func TestCloseSqlxInstances(t *testing.T) {
	viper.SetDefault("mysql.unittest.host", "127.0.0.1")
	viper.SetDefault("mysql.unittest.port", 3306)
	viper.SetDefault("mysql.unittest.username", "root")
	viper.SetDefault("mysql.unittest.password", "roooooot")
	viper.SetDefault("mysql.unittest.dbname", "information_schema")
	if db, err := SqlxMySQL("unittest"); err != nil {
		t.Fatal(err)
	} else if db == nil {
		t.Fatal("db is nil")
	} else if err := db.Ping(); err != nil {
		t.Error(err)
	}
	if _, loaded := SqlxInstances.Load("mysql"); !loaded {
		t.Error("mysql should be loaded")
	}
	dbname := "db.sqlite3"
	viper.SetDefault("sqlite3.unittest.dbname", dbname)
	if db, err := SqlxSQLite3("unittest"); err != nil {
		t.Fatal(err)
	} else if db == nil {
		t.Fatal("db is nil")
	} else if err := db.Ping(); err != nil {
		t.Error(err)
	}
	if _, loaded := SqlxInstances.Load("sqlite3"); !loaded {
		t.Error("sqlite3 should be loaded")
	}
	defer os.Remove(dbname)
	CloseSqlxInstances()

	if _, loaded := SqlxInstances.Load("sqlite3"); loaded {
		t.Error("sqlite3 should not be loaded")
	}
	if _, loaded := SqlxInstances.Load("mysql"); loaded {
		t.Error("mysql should not be loaded")
	}
}
