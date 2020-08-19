// sqlx 创建db对象的函数封装

package goutils

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

// NewSqlxSQLite3 return sqlx sqlite3 db instance
// dbname is dbfile path
// logMode show detailed log
// maxIdleConns sets the maximum number of connections in the idle connection pool
// maxOpenConns sets the maximum number of open connections to the database.
// connMaxLifeMinutes sets the maximum amount of time(minutes) a connection may be reused
func NewSqlxSQLite3(dbname string, maxIdleConns, maxOpenConns, connMaxLifeMinutes int) (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", dbname)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(maxIdleConns)                                       // 设置连接池中的最大闲置连接数
	db.SetMaxOpenConns(maxOpenConns)                                       // 设置数据库的最大连接数量
	db.SetConnMaxLifetime(time.Duration(connMaxLifeMinutes) * time.Minute) // 设置连接的最大可复用时间
	return db, nil
}

// NewSqlxMySQL return sqlx mysql db instance
// host is database's host
// port is database's port
// dbname is database's dbname
// usename is database's username
// password is database's password
// logMode show detailed log
// maxIdleConns sets the maximum number of connections in the idle connection pool
// maxOpenConns sets the maximum number of open connections to the database.
// connMaxLifeMinutes sets the maximum amount of time(minutes) a connection may be reused
// timeout conn timeout, readtimeout and writetimeout is x3
func NewSqlxMySQL(host string, port int, username, password, dbname string, maxIdleConns, maxOpenConns, connMaxLifeMinutes, timeout int) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%ds&readTimeout=%ds&writeTimeout=%ds", username, password, host, port, dbname, timeout, timeout*5, timeout*5)
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(maxIdleConns)                                       // 设置连接池中的最大闲置连接数
	db.SetMaxOpenConns(maxOpenConns)                                       // 设置数据库的最大连接数量
	db.SetConnMaxLifetime(time.Duration(connMaxLifeMinutes) * time.Minute) // 设置连接的最大可复用时间
	return db, nil
}

// NewSqlxPostgres return sqlx postgresql db instance
// host is database's host
// port is database's port
// dbname is database's dbname
// usename is database's username
// sslmode ssl is disable or not
// password is database's password
// logMode show detailed log
// maxIdleConns sets the maximum number of connections in the idle connection pool
// maxOpenConns sets the maximum number of open connections to the database.
// connMaxLifeMinutes sets the maximum amount of time(minutes) a connection may be reused
func NewSqlxPostgres(host string, port int, username, password, dbname, sslmode string, maxIdleConns, maxOpenConns, connMaxLifeMinutes int) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s", host, port, username, dbname, password, sslmode)
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(maxIdleConns)                                       // 设置连接池中的最大闲置连接数
	db.SetMaxOpenConns(maxOpenConns)                                       // 设置数据库的最大连接数量
	db.SetConnMaxLifetime(time.Duration(connMaxLifeMinutes) * time.Minute) // 设置连接的最大可复用时间
	return db, nil
}

// NewSqlxMsSQL return sqlx sqlserver db instance
// host is database's host
// port is database's port
// dbname is database's dbname
// usename is database's username
// password is database's password
// logMode show detailed log
// maxIdleConns sets the maximum number of connections in the idle connection pool
// maxOpenConns sets the maximum number of open connections to the database.
// connMaxLifeMinutes sets the maximum amount of time(minutes) a connection may be reused
func NewSqlxMsSQL(host string, port int, username, password, dbname string, maxIdleConns, maxOpenConns, connMaxLifeMinutes int) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", username, password, host, port, dbname)
	db, err := sqlx.Open("mssql", dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(maxIdleConns)                                       // 设置连接池中的最大闲置连接数
	db.SetMaxOpenConns(maxOpenConns)                                       // 设置数据库的最大连接数量
	db.SetConnMaxLifetime(time.Duration(connMaxLifeMinutes) * time.Minute) // 设置连接的最大可复用时间
	return db, nil
}
