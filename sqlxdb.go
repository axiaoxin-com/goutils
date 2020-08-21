// sqlx 创建 db 对象的函数封装

package goutils

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

// NewSqlxSQLite3 返回 sqlx sqlite3 连接实例
func NewSqlxSQLite3(conf SQLite3Config) (*sqlx.DB, error) {
	sqlxSqlite3, err := sqlx.Open("sqlite3", conf.DBName)
	if err != nil {
		return nil, err
	}
	sqlxSqlite3.SetMaxIdleConns(conf.MaxIdleConns)
	sqlxSqlite3.SetMaxOpenConns(conf.MaxOpenConns)
	sqlxSqlite3.SetConnMaxLifetime(time.Duration(conf.ConnMaxLifeMinutes) * time.Minute)
	return sqlxSqlite3, nil
}

// NewSqlxMySQL 返回 sqlx mysql 连接实例
func NewSqlxMySQL(conf MySQLConfig) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%ds&readTimeout=%ds&writeTimeout=%ds", conf.Username, conf.Password, conf.Host, conf.Port, conf.DBName, conf.ConnTimeout, conf.ReadTimeout, conf.WriteTimeout)
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
func NewSqlxPostgres(conf PostgresConfig) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", conf.Host, conf.Port, conf.Username, conf.DBName, conf.Password)
	if conf.DisableSSL {
		dsn = dsn + " sslmode=disable"
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
func NewSqlxMsSQL(conf MsSQLConfig) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", conf.Username, conf.Password, conf.Host, conf.Port, conf.DBName)
	sqlxMssql, err := sqlx.Open("mssql", dsn)
	if err != nil {
		return nil, err
	}
	sqlxMssql.SetMaxIdleConns(conf.MaxIdleConns)
	sqlxMssql.SetMaxOpenConns(conf.MaxOpenConns)
	sqlxMssql.SetConnMaxLifetime(time.Duration(conf.ConnMaxLifeMinutes) * time.Minute)
	return sqlxMssql, nil
}
