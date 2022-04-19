package goutils

import (
	"os"
	"sync"
	"testing"

	"github.com/spf13/viper"
)

func TestNewGormSQLite(t *testing.T) {
	dbname := "db.sqlite"
	conf := DBConfig{
		DBName: dbname,
	}
	db, err := NewGormSQLite(conf)
	if err != nil {
		t.Error("new gorm sqlite return error:", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		t.Error(err)
	}
	defer sqlDB.Close()
	defer os.Remove(dbname)
	if err := sqlDB.Ping(); err != nil {
		t.Error(err)
	}
}

func TestNewGormMySQL(t *testing.T) {
	conf := DBConfig{
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
	sqlDB, err := db.DB()
	if err != nil {
		t.Error(err)
	}
	defer sqlDB.Close()
	if err := sqlDB.Ping(); err != nil {
		t.Error(err)
	}
}

func TestLikeFieldEscape(t *testing.T) {
	sql := `select * from t where f like "*100%*";`
	e := `select * from t where f like \"*100\%*\"\;`
	if s := LikeFieldEscape(sql); s != e {
		t.Error("like escape failed. raw:", sql, " escaped:", s, " expect:", e)
	}

}

func TestGormMySQL(t *testing.T) {
	defer viper.Reset()
	viper.Set("mysql.unittest.host", "127.0.0.1")
	viper.Set("mysql.unittest.port", 3306)
	viper.Set("mysql.unittest.username", "root")
	viper.Set("mysql.unittest.password", "roooooot")
	viper.Set("mysql.unittest.dbname", "information_schema")
	if db, err := GormMySQL("unittest"); err != nil {
		t.Fatal(err)
	} else if db == nil {
		t.Fatal("db is nil")
	} else {
		sqlDB, err := db.DB()
		if err != nil {
			t.Error(err)
		}
		if err := sqlDB.Ping(); err != nil {
			t.Error(err)
		}
	}
	defer CloseGormInstances()
	if _, err := GormMySQL("unittest"); err != nil {
		t.Error(err)
	}
	viper.Set("mysql.localhost.host", "127.0.0.1")
	viper.Set("mysql.localhost.port", 3306)
	viper.Set("mysql.localhost.username", "root")
	viper.Set("mysql.localhost.password", "roooooot")
	viper.Set("mysql.localhost.dbname", "information_schema")
	if _, err := GormMySQL("localhost"); err != nil {
		t.Error(err)
	}
	instanceCount := 0
	mysqlCount := 0
	GormInstances.Range(func(k, v interface{}) bool {
		instanceCount++
		if k.(string) == "mysql" {
			v.(*sync.Map).Range(func(kk, vv interface{}) bool {
				mysqlCount++
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

func TestGormSQLite(t *testing.T) {
	defer viper.Reset()
	dbname := "db.sqlite"
	viper.Set("sqlite.unittest.dbname", dbname)
	if db, err := GormSQLite("unittest"); err != nil {
		t.Fatal(err)
	} else if db == nil {
		t.Fatal("db is nil")
	} else {
		sqlDB, err := db.DB()
		if err != nil {
			t.Error(err)
		}
		if err := sqlDB.Ping(); err != nil {
			t.Error(err)
		}
	}

	defer CloseGormInstances()
	defer os.Remove(dbname)
	if _, err := GormSQLite("unittest"); err != nil {
		t.Error(err)
	}
	viper.Set("sqlite.other.dbname", "other."+dbname)
	if _, err := GormSQLite("other"); err != nil {
		t.Error(err)
	}
	defer os.Remove("other." + dbname)
	instanceCount := 0
	sqliteCount := 0
	GormInstances.Range(func(k, v interface{}) bool {
		instanceCount++
		if k.(string) == "sqlite" {
			v.(*sync.Map).Range(func(kk, vv interface{}) bool {
				sqliteCount++
				return true
			})
		}
		return true
	})
	if instanceCount != 1 {
		t.Error("instanceCount != 1, ", instanceCount)
	}
	if sqliteCount != 2 {
		t.Error("sqliteCount != 2, ", sqliteCount)
	}
}

func TestCloseGormInstances(t *testing.T) {
	defer viper.Reset()
	viper.Set("mysql.unittest.host", "127.0.0.1")
	viper.Set("mysql.unittest.port", 3306)
	viper.Set("mysql.unittest.username", "root")
	viper.Set("mysql.unittest.password", "roooooot")
	viper.Set("mysql.unittest.dbname", "information_schema")
	if db, err := GormMySQL("unittest"); err != nil {
		t.Fatal(err)
	} else if db == nil {
		t.Fatal("db is nil")
	} else {
		sqlDB, err := db.DB()
		if err != nil {
			t.Error(err)
		}
		if err := sqlDB.Ping(); err != nil {
			t.Error(err)
		}
	}

	if _, loaded := GormInstances.Load("mysql"); !loaded {
		t.Error("mysql should be loaded")
	}
	dbname := "db.sqlite"
	viper.Set("sqlite.unittest.dbname", dbname)
	if db, err := GormSQLite("unittest"); err != nil {
		t.Fatal(err)
	} else if db == nil {
		t.Fatal("db is nil")
	} else {
		sqlDB, err := db.DB()
		if err != nil {
			t.Error(err)
		}
		if err := sqlDB.Ping(); err != nil {
			t.Error(err)
		}
	}

	if _, loaded := GormInstances.Load("sqlite"); !loaded {
		t.Error("sqlite should be loaded")
	}
	defer os.Remove(dbname)
	CloseGormInstances()

	if _, loaded := GormInstances.Load("sqlite"); loaded {
		t.Error("sqlite should not be loaded")
	}
	if _, loaded := GormInstances.Load("mysql"); loaded {
		t.Error("mysql should not be loaded")
	}
}
