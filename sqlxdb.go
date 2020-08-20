// sqlx 创建 db 对象的函数封装

package goutils

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

// NewSqlxSQLite3 返回 sqlx sqlite3 连接实例
// dbname 数据库文件名（含路径）
// logMode 是否开启打印日志模式
// maxIdleConns 设置空闲连接池中的最大连接数
// maxOpenConns 设置与数据库的最大打开连接数
// connMaxLifeMinutes 设置可重用连接的最长时间（分钟）
func NewSqlxSQLite3(dbname string, maxIdleConns, maxOpenConns, connMaxLifeMinutes int) (*sqlx.DB, error) {
	sqlxSqlite3, err := sqlx.Open("sqlite3", dbname)
	if err != nil {
		return nil, err
	}
	sqlxSqlite3.SetMaxIdleConns(maxIdleConns)                                       // 设置连接池中的最大闲置连接数
	sqlxSqlite3.SetMaxOpenConns(maxOpenConns)                                       // 设置数据库的最大连接数量
	sqlxSqlite3.SetConnMaxLifetime(time.Duration(connMaxLifeMinutes) * time.Minute) // 设置连接的最大可复用时间
	return sqlxSqlite3, nil
}

// NewSqlxMySQL 返回 sqlx mysql 连接实例
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
func NewSqlxMySQL(host string, port int, username, password, dbname string, maxIdleConns, maxOpenConns, connMaxLifeMinutes, connTimeout, readTimeout, writeTimeout int) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%ds&readTimeout=%ds&writeTimeout=%ds", username, password, host, port, dbname, connTimeout, readTimeout, writeTimeout)
	sqlxMysql, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	sqlxMysql.SetMaxIdleConns(maxIdleConns)                                       // 设置连接池中的最大闲置连接数
	sqlxMysql.SetMaxOpenConns(maxOpenConns)                                       // 设置数据库的最大连接数量
	sqlxMysql.SetConnMaxLifetime(time.Duration(connMaxLifeMinutes) * time.Minute) // 设置连接的最大可复用时间
	return sqlxMysql, nil
}

// NewSqlxPostgres 返回 sqlx postgresql 连接实例
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
func NewSqlxPostgres(host string, port int, username, password, dbname string, disableSSL bool, maxIdleConns, maxOpenConns, connMaxLifeMinutes int) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", host, port, username, dbname, password)
	if disableSSL {
		dsn = dsn + " sslmode=disable"
	}
	sqlxPostgres, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	sqlxPostgres.SetMaxIdleConns(maxIdleConns)                                       // 设置连接池中的最大闲置连接数
	sqlxPostgres.SetMaxOpenConns(maxOpenConns)                                       // 设置数据库的最大连接数量
	sqlxPostgres.SetConnMaxLifetime(time.Duration(connMaxLifeMinutes) * time.Minute) // 设置连接的最大可复用时间
	return sqlxPostgres, nil
}

// NewSqlxMsSQL 返回 sqlx sqlserver 连接实例
// host 数据库 IP 地址
// port 数据库端口
// dbname 数据库名称
// usename 数据库用户名
// password 数据库密码
// logMode 是否开启打印日志模式
// maxIdleConns 设置空闲连接池中的最大连接数
// maxOpenConns 设置与数据库的最大打开连接数
// connMaxLifeMinutes 设置可重用连接的最长时间（分钟）
func NewSqlxMsSQL(host string, port int, username, password, dbname string, maxIdleConns, maxOpenConns, connMaxLifeMinutes int) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", username, password, host, port, dbname)
	sqlxMssql, err := sqlx.Open("mssql", dsn)
	if err != nil {
		return nil, err
	}
	sqlxMssql.SetMaxIdleConns(maxIdleConns)                                       // 设置连接池中的最大闲置连接数
	sqlxMssql.SetMaxOpenConns(maxOpenConns)                                       // 设置数据库的最大连接数量
	sqlxMssql.SetConnMaxLifetime(time.Duration(connMaxLifeMinutes) * time.Minute) // 设置连接的最大可复用时间
	return sqlxMssql, nil
}
