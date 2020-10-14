// sqlx 创建 db 对象的函数封装

package goutils

import (
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

// NewSqlxSQLite3 返回 sqlx sqlite3 连接实例
func NewSqlxSQLite3(conf DBConfig) (*sqlx.DB, error) {
	dsn, err := conf.SQLite3DSN()
	if err != nil {
		return nil, err
	}

	sqlxSqlite3, err := sqlx.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	sqlxSqlite3.SetMaxIdleConns(conf.MaxIdleConns)
	sqlxSqlite3.SetMaxOpenConns(conf.MaxOpenConns)
	sqlxSqlite3.SetConnMaxLifetime(time.Duration(conf.ConnMaxLifeMinutes) * time.Minute)
	return sqlxSqlite3, nil
}

// NewSqlxMySQL 返回 sqlx mysql 连接实例
func NewSqlxMySQL(conf DBConfig) (*sqlx.DB, error) {
	dsn, err := conf.MySQLDSN()
	if err != nil {
		return nil, err
	}

	sqlxMysql, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	sqlxMysql.SetMaxIdleConns(conf.MaxIdleConns)
	sqlxMysql.SetMaxOpenConns(conf.MaxOpenConns)
	sqlxMysql.SetConnMaxLifetime(time.Duration(conf.ConnMaxLifeMinutes) * time.Minute)
	return sqlxMysql, nil
}

// NewSqlxPostgres 返回 sqlx postgresql 连接实例
func NewSqlxPostgres(conf DBConfig) (*sqlx.DB, error) {
	dsn, err := conf.PostgresDSN()
	if err != nil {
		return nil, err
	}

	sqlxPostgres, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	sqlxPostgres.SetMaxIdleConns(conf.MaxIdleConns)
	sqlxPostgres.SetMaxOpenConns(conf.MaxOpenConns)
	sqlxPostgres.SetConnMaxLifetime(time.Duration(conf.ConnMaxLifeMinutes) * time.Minute)
	return sqlxPostgres, nil
}

// NewSqlxMsSQL 返回 sqlx sqlserver 连接实例
func NewSqlxMsSQL(conf DBConfig) (*sqlx.DB, error) {
	dsn, err := conf.SqlserverDSN()
	if err != nil {
		return nil, err
	}

	sqlxMssql, err := sqlx.Open("mssql", dsn)
	if err != nil {
		return nil, err
	}
	sqlxMssql.SetMaxIdleConns(conf.MaxIdleConns)
	sqlxMssql.SetMaxOpenConns(conf.MaxOpenConns)
	sqlxMssql.SetConnMaxLifetime(time.Duration(conf.ConnMaxLifeMinutes) * time.Minute)
	return sqlxMssql, nil
}

// SqlxInstances 以 sync.Map 保存 sqlx db 相关信息
// key 为小写的数据库驱动名称， value 为实例名为 key ， 具体的 db 对象为 value 的 sync.Map
// 形如： {"mysql": {"localhost": db}, "postgres": {"localhost": db}}
var SqlxInstances sync.Map

// SqlxMySQL 根据 viper 中配置的实例名称返回 sqlx 连接 mysql 的实例
func SqlxMySQL(which string) (*sqlx.DB, error) {
	mysqls, loaded := SqlxInstances.LoadOrStore("mysql", new(sync.Map))
	if loaded {
		if db, loaded := mysqls.(*sync.Map).Load(which); loaded {
			return db.(*sqlx.DB), nil
		}
	}
	// mysql 不存在 或 存在时 which 不存在 则新建 mysql db 实例存放到 map 中
	// 注意：这里依赖 viper ，必须在外部先对 viper 配置进行加载
	prefix := "mysql." + which
	conf := DBConfig{
		Host:               viper.GetString(prefix + ".host"),
		Port:               viper.GetInt(prefix + ".port"),
		Username:           viper.GetString(prefix + ".username"),
		Password:           viper.GetString(prefix + ".password"),
		DBName:             viper.GetString(prefix + ".dbname"),
		LogMode:            viper.GetBool(prefix + ".log_mode"),
		MaxIdleConns:       viper.GetInt(prefix + ".max_idle_conns"),
		MaxOpenConns:       viper.GetInt(prefix + ".max_open_conns"),
		ConnMaxLifeMinutes: viper.GetInt(prefix + ".conn_max_life_minutes"),
		ConnTimeout:        viper.GetInt(prefix + ".conn_timeout"),
		ReadTimeout:        viper.GetInt(prefix + ".read_timeout"),
		WriteTimeout:       viper.GetInt(prefix + ".write_timeout"),
	}
	db, err := NewSqlxMySQL(conf)
	if err != nil {
		return nil, err
	}
	mysqls.(*sync.Map).Store(which, db)
	SqlxInstances.Store("mysql", mysqls)
	return db, nil
}

// SqlxSQLite3 根据  viper 中配置的实例名称返回 sqlite3 实例
func SqlxSQLite3(which string) (*sqlx.DB, error) {
	sqlite3s, loaded := SqlxInstances.LoadOrStore("sqlite3", new(sync.Map))
	if loaded {
		if db, loaded := sqlite3s.(*sync.Map).Load(which); loaded {
			return db.(*sqlx.DB), nil
		}
	}
	prefix := "sqlite3." + which
	conf := DBConfig{
		DBName:             viper.GetString(prefix + ".dbname"),
		LogMode:            viper.GetBool(prefix + ".log_mode"),
		MaxIdleConns:       viper.GetInt(prefix + ".max_idle_conns"),
		MaxOpenConns:       viper.GetInt(prefix + ".max_open_conns"),
		ConnMaxLifeMinutes: viper.GetInt(prefix + ".conn_max_life_minutes"),
	}
	db, err := NewSqlxSQLite3(conf)
	if err != nil {
		return nil, err
	}
	sqlite3s.(*sync.Map).Store(which, db)
	SqlxInstances.Store("sqlite3", sqlite3s)
	return db, nil
}

// SqlxPostgres 根据 viper 中配置的实例名称返回 pg 实例
func SqlxPostgres(which string) (*sqlx.DB, error) {
	pgs, loaded := SqlxInstances.LoadOrStore("postgres", new(sync.Map))
	if loaded {
		if db, loaded := pgs.(*sync.Map).Load(which); loaded {
			return db.(*sqlx.DB), nil
		}
	}
	prefix := "postgres." + which
	conf := DBConfig{
		Host:               viper.GetString(prefix + ".host"),
		Port:               viper.GetInt(prefix + ".port"),
		Username:           viper.GetString(prefix + ".username"),
		Password:           viper.GetString(prefix + ".password"),
		DBName:             viper.GetString(prefix + ".dbname"),
		LogMode:            viper.GetBool(prefix + ".log_mode"),
		MaxIdleConns:       viper.GetInt(prefix + ".max_idle_conns"),
		MaxOpenConns:       viper.GetInt(prefix + ".max_open_conns"),
		ConnMaxLifeMinutes: viper.GetInt(prefix + ".conn_max_life_minutes"),
		DisableSSL:         viper.GetBool(prefix + ".disable_ssl"),
	}
	db, err := NewSqlxPostgres(conf)
	if err != nil {
		return nil, err
	}
	pgs.(*sync.Map).Store(which, db)
	SqlxInstances.Store("postgres", pgs)
	return db, nil
}

// SqlxMsSQL 根据 viper 中配置的实例名称返回 sqlserver 实例
func SqlxMsSQL(which string) (*sqlx.DB, error) {
	mssqls, loaded := SqlxInstances.LoadOrStore("mssql", new(sync.Map))
	if loaded {
		if db, loaded := mssqls.(*sync.Map).Load(which); loaded {
			return db.(*sqlx.DB), nil
		}
	}
	// mysql 不存在 或 存在时 which 不存在 则新建 mysql db 实例存放到 map 中
	prefix := "mssql." + which
	conf := DBConfig{
		Host:               viper.GetString(prefix + ".host"),
		Port:               viper.GetInt(prefix + ".port"),
		Username:           viper.GetString(prefix + ".username"),
		Password:           viper.GetString(prefix + ".password"),
		DBName:             viper.GetString(prefix + ".dbname"),
		LogMode:            viper.GetBool(prefix + ".log_mode"),
		MaxIdleConns:       viper.GetInt(prefix + ".max_idle_conns"),
		MaxOpenConns:       viper.GetInt(prefix + ".max_open_conns"),
		ConnMaxLifeMinutes: viper.GetInt(prefix + ".conn_max_life_minutes"),
	}
	db, err := NewSqlxMsSQL(conf)
	if err != nil {
		return nil, err
	}
	mssqls.(*sync.Map).Store(which, db)
	SqlxInstances.Store("mssql", mssqls)
	return db, nil
}

// CloseSqlxInstances 关闭全部的 Sqlx 连接并重置 SqlxInstances
func CloseSqlxInstances() {
	SqlxInstances.Range(func(ik, iv interface{}) bool {
		if m, ok := iv.(*sync.Map); ok {
			m.Range(func(k, v interface{}) bool {
				if db, ok := v.(*sqlx.DB); ok {
					db.Close()
				}
				return true
			})
		}
		return true
	})
	SqlxInstances = sync.Map{}
}
