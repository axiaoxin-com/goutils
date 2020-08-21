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
func NewGormSQLite3(conf SQLite3Config) (*gorm.DB, error) {
	gormSqlite3, err := gorm.Open("sqlite3", conf.DBName)
	if err != nil {
		return nil, err
	}
	gormSqlite3.LogMode(conf.LogMode)
	gormSqlite3.DB().SetMaxIdleConns(conf.MaxIdleConns)
	gormSqlite3.DB().SetMaxOpenConns(conf.MaxOpenConns)
	gormSqlite3.DB().SetConnMaxLifetime(time.Duration(conf.ConnMaxLifeMinutes) * time.Minute)
	return gormSqlite3, nil
}

// NewGormMySQL 返回 gorm mysql 连接实例
func NewGormMySQL(conf MySQLConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%ds&readTimeout=%ds&writeTimeout=%ds", conf.Username, conf.Password, conf.Host, conf.Port, conf.DBName, conf.ConnTimeout, conf.ReadTimeout, conf.WriteTimeout)
	gormMysql, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	gormMysql.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8 auto_increment=1")
	gormMysql.LogMode(conf.LogMode)
	gormMysql.DB().SetMaxIdleConns(conf.MaxIdleConns)
	gormMysql.DB().SetMaxOpenConns(conf.MaxOpenConns)
	gormMysql.DB().SetConnMaxLifetime(time.Duration(conf.ConnMaxLifeMinutes) * time.Minute)
	return gormMysql, nil
}

// NewGormPostgres 返回 gorm postgresql 连接实例
func NewGormPostgres(conf PostgresConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", conf.Host, conf.Port, conf.Username, conf.DBName, conf.Password)
	if conf.DisableSSL {
		dsn = dsn + " sslmode=disable"
	}
	gormPostgres, err := gorm.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	gormPostgres.LogMode(conf.LogMode)
	gormPostgres.DB().SetMaxIdleConns(conf.MaxIdleConns)
	gormPostgres.DB().SetMaxOpenConns(conf.MaxOpenConns)
	gormPostgres.DB().SetConnMaxLifetime(time.Duration(conf.ConnMaxLifeMinutes) * time.Minute)
	return gormPostgres, nil
}

// NewGormMsSQL 返回 gorm sqlserver 连接实例
func NewGormMsSQL(conf MsSQLConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", conf.Username, conf.Password, conf.Host, conf.Port, conf.DBName)
	gormMssql, err := gorm.Open("mssql", dsn)
	if err != nil {
		return nil, err
	}
	gormMssql.LogMode(conf.LogMode)
	gormMssql.DB().SetMaxIdleConns(conf.MaxIdleConns)
	gormMssql.DB().SetMaxOpenConns(conf.MaxOpenConns)
	gormMssql.DB().SetConnMaxLifetime(time.Duration(conf.ConnMaxLifeMinutes) * time.Minute)
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
