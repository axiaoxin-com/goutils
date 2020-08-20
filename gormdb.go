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

// NewGormSQLite3 返回 gorm sqlite3 连接实例
// dbname 数据库文件名（含路径）
// logMode 是否开启打印日志模式
// maxIdleConns 设置空闲连接池中的最大连接数
// maxOpenConns 设置与数据库的最大打开连接数
// connMaxLifeMinutes 设置可重用连接的最长时间（分钟）
func NewGormSQLite3(dbname string, logMode bool, maxIdleConns, maxOpenConns, connMaxLifeMinutes int) (*gorm.DB, error) {
	gormSqlite3, err := gorm.Open("sqlite3", dbname)
	if err != nil {
		return nil, err
	}
	gormSqlite3.LogMode(logMode)
	gormSqlite3.DB().SetMaxIdleConns(maxIdleConns)                                       // 设置连接池中的最大闲置连接数
	gormSqlite3.DB().SetMaxOpenConns(maxOpenConns)                                       // 设置数据库的最大连接数量
	gormSqlite3.DB().SetConnMaxLifetime(time.Duration(connMaxLifeMinutes) * time.Minute) // 设置连接的最大可复用时间
	return gormSqlite3, nil
}

// NewGormMySQL 返回 gorm mysql 连接实例
// host 数据库 IP 地址
// port 数据库端口
// dbname 数据库名称
// usename 数据库用户名
// password 数据库密码
// logMode 是否开启打印日志模式
// maxIdleConns 设置空闲连接池中的最大连接数
// maxOpenConns 设置与数据库的最大打开连接数
// connMaxLifeMinutes 设置可重用连接的最长时间（分钟）
// connTimeout 连接超时时间（秒）
// readTimeout 读超时时间（秒）
// writeTimeout 写超时时间（秒）
func NewGormMySQL(host string, port int, username, password, dbname string, logMode bool, maxIdleConns, maxOpenConns, connMaxLifeMinutes, connTimeout, readTimeout, writeTimeout int) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%ds&readTimeout=%ds&writeTimeout=%ds", username, password, host, port, dbname, connTimeout, readTimeout, writeTimeout)
	gormMysql, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	gormMysql.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8 auto_increment=1")
	gormMysql.LogMode(logMode)
	gormMysql.DB().SetMaxIdleConns(maxIdleConns)                                       // 设置连接池中的最大闲置连接数
	gormMysql.DB().SetMaxOpenConns(maxOpenConns)                                       // 设置数据库的最大连接数量
	gormMysql.DB().SetConnMaxLifetime(time.Duration(connMaxLifeMinutes) * time.Minute) // 设置连接的最大可复用时间
	return gormMysql, nil
}

// NewGormPostgres 返回 gorm postgresql 连接实例
// host 数据库 IP 地址
// port 数据库端口
// dbname 数据库名称
// usename 数据库用户名
// password 数据库密码
// disableSSL 是否关闭 ssl 模式
// logMode 是否开启打印日志模式
// maxIdleConns 设置空闲连接池中的最大连接数
// maxOpenConns 设置与数据库的最大打开连接数
// connMaxLifeMinutes 设置可重用连接的最长时间（分钟）
func NewGormPostgres(host string, port int, username, password, dbname string, disableSSL, logMode bool, maxIdleConns, maxOpenConns, connMaxLifeMinutes int) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", host, port, username, dbname, password)
	if disableSSL {
		dsn = dsn + " sslmode=disable"
	}
	gormPostgres, err := gorm.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	gormPostgres.LogMode(logMode)
	gormPostgres.DB().SetMaxIdleConns(maxIdleConns)                                       // 设置连接池中的最大闲置连接数
	gormPostgres.DB().SetMaxOpenConns(maxOpenConns)                                       // 设置数据库的最大连接数量
	gormPostgres.DB().SetConnMaxLifetime(time.Duration(connMaxLifeMinutes) * time.Minute) // 设置连接的最大可复用时间
	return gormPostgres, nil
}

// NewGormMsSQL 返回 gorm sqlserver 连接实例
// host 数据库 IP 地址
// port 数据库端口
// dbname 数据库名称
// usename 数据库用户名
// password 数据库密码
// logMode 是否开启打印日志模式
// maxIdleConns 设置空闲连接池中的最大连接数
// maxOpenConns 设置与数据库的最大打开连接数
// connMaxLifeMinutes 设置可重用连接的最长时间（分钟）
func NewGormMsSQL(host string, port int, username, password, dbname string, logMode bool, maxIdleConns, maxOpenConns, connMaxLifeMinutes int) (*gorm.DB, error) {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", username, password, host, port, dbname)
	gormMssql, err := gorm.Open("mssql", dsn)
	if err != nil {
		return nil, err
	}
	gormMssql.LogMode(logMode)
	gormMssql.DB().SetMaxIdleConns(maxIdleConns)                                       // 设置连接池中的最大闲置连接数
	gormMssql.DB().SetMaxOpenConns(maxOpenConns)                                       // 设置数据库的最大连接数量
	gormMssql.DB().SetConnMaxLifetime(time.Duration(connMaxLifeMinutes) * time.Minute) // 设置连接的最大可复用时间
	return gormMssql, nil
}

// GormMySQLLikeFieldEscape 转义 Gorm MySQL 的 like 模糊查询时字段值为通配符的值
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
