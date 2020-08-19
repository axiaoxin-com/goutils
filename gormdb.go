// gorm 创建 db 对象的函数封装

package goutils

import (
	"fmt"
	"strings"
	"time"

	"github.com/jinzhu/gorm"

	// need by gorm
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// NewGormSQLite3 return gorm sqlite3 db instance
// dbname is dbfile path
// logMode show detailed log
// maxIdleConns sets the maximum number of connections in the idle connection pool
// maxOpenConns sets the maximum number of open connections to the database.
// connMaxLifeMinutes sets the maximum amount of time(minutes) a connection may be reused
func NewGormSQLite3(dbname string, logMode bool, maxIdleConns, maxOpenConns, connMaxLifeMinutes int) (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", dbname)
	if err != nil {
		return nil, err
	}
	db.LogMode(logMode)
	db.DB().SetMaxIdleConns(maxIdleConns)                                       // 设置连接池中的最大闲置连接数
	db.DB().SetMaxOpenConns(maxOpenConns)                                       // 设置数据库的最大连接数量
	db.DB().SetConnMaxLifetime(time.Duration(connMaxLifeMinutes) * time.Minute) // 设置连接的最大可复用时间
	return db, nil
}

// NewGormMySQL return gorm mysql db instance
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
func NewGormMySQL(host string, port int, username, password, dbname string, logMode bool, maxIdleConns, maxOpenConns, connMaxLifeMinutes, timeout int) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%ds&readTimeout=%ds&writeTimeout=%ds", username, password, host, port, dbname, timeout, timeout*5, timeout*5)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8 auto_increment=1")
	db.LogMode(logMode)
	db.DB().SetMaxIdleConns(maxIdleConns)                                       // 设置连接池中的最大闲置连接数
	db.DB().SetMaxOpenConns(maxOpenConns)                                       // 设置数据库的最大连接数量
	db.DB().SetConnMaxLifetime(time.Duration(connMaxLifeMinutes) * time.Minute) // 设置连接的最大可复用时间
	return db, nil
}

// NewGormPostgres return gorm postgresql db instance
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
func NewGormPostgres(host string, port int, username, password, dbname, sslmode string, logMode bool, maxIdleConns, maxOpenConns, connMaxLifeMinutes int) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s", host, port, username, dbname, password, sslmode)
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	db.LogMode(logMode)
	db.DB().SetMaxIdleConns(maxIdleConns)                                       // 设置连接池中的最大闲置连接数
	db.DB().SetMaxOpenConns(maxOpenConns)                                       // 设置数据库的最大连接数量
	db.DB().SetConnMaxLifetime(time.Duration(connMaxLifeMinutes) * time.Minute) // 设置连接的最大可复用时间
	return db, nil
}

// NewGormMsSQL return gorm sqlserver db instance
// host is database's host
// port is database's port
// dbname is database's dbname
// usename is database's username
// password is database's password
// logMode show detailed log
// maxIdleConns sets the maximum number of connections in the idle connection pool
// maxOpenConns sets the maximum number of open connections to the database.
// connMaxLifeMinutes sets the maximum amount of time(minutes) a connection may be reused
func NewGormMsSQL(host string, port int, username, password, dbname string, logMode bool, maxIdleConns, maxOpenConns, connMaxLifeMinutes int) (*gorm.DB, error) {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", username, password, host, port, dbname)
	db, err := gorm.Open("mssql", dsn)
	if err != nil {
		return nil, err
	}
	db.LogMode(logMode)
	db.DB().SetMaxIdleConns(maxIdleConns)                                       // 设置连接池中的最大闲置连接数
	db.DB().SetMaxOpenConns(maxOpenConns)                                       // 设置数据库的最大连接数量
	db.DB().SetConnMaxLifetime(time.Duration(connMaxLifeMinutes) * time.Minute) // 设置连接的最大可复用时间
	return db, nil
}

// GormMySQLLikeFieldEscape 转义Gorm MySQL的like模糊查询时字段值为通配符的值
func GormMySQLLikeFieldEscape(value string) string {
	value = strings.Replace(value, ";", "\\;", -1)
	value = strings.Replace(value, "\"", "\\\"", -1)
	value = strings.Replace(value, "'", "\\'", -1)
	value = strings.Replace(value, "--", "\\--", -1)
	value = strings.Replace(value, "/", "\\/", -1)
	value = strings.Replace(value, "%", "\\%", -1)
	value = strings.Replace(value, "_", "\\_", -1)
	return value
}
