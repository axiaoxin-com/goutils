// gorm 创建 db 对象的函数封装

package goutils

import (
	"strings"
	"sync"
	"time"

	"github.com/spf13/viper"
	"gorm.io/gorm"

	// need by gorm
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
)

// GormBaseModel 基础 model 定义
type GormBaseModel struct {
	ID        string   `gorm:"primary_key,column:id" json:"id"         example:"-"` // 字符串类型的 Hash 主键 ID
	CreatedAt JSONTime `gorm:"column:created_at"     json:"created_at" example:"-"` // 创建时间
	UpdatedAt JSONTime `gorm:"column:updated_at"     json:"updated_at" example:"-"` // 更新时间
}

// NewGormSQLite3 返回 gorm sqlite3 连接实例
func NewGormSQLite3(conf DBConfig) (*gorm.DB, error) {
	dsn, err := conf.SQLite3DSN()
	if err != nil {
		return nil, err
	}

	if conf.GormConfig == nil {
		conf.GormConfig = &gorm.Config{}
	}

	gormSqlite3, err := gorm.Open(sqlite.Open(dsn), conf.GormConfig)
	if err != nil {
		return nil, err
	}
	sqlite3db, err := gormSqlite3.DB()
	if err != nil {
		return nil, err
	}
	sqlite3db.SetMaxIdleConns(conf.MaxIdleConns)
	sqlite3db.SetMaxOpenConns(conf.MaxOpenConns)
	sqlite3db.SetConnMaxLifetime(time.Duration(conf.ConnMaxLifeMinutes) * time.Minute)
	return gormSqlite3, nil
}

// NewGormMySQL 返回 gorm mysql 连接实例
func NewGormMySQL(conf DBConfig) (*gorm.DB, error) {
	dsn, err := conf.MySQLDSN()
	if err != nil {
		return nil, err
	}

	if conf.GormConfig == nil {
		conf.GormConfig = &gorm.Config{}
	}

	gormMysql, err := gorm.Open(mysql.Open(dsn), conf.GormConfig)
	if err != nil {
		return nil, err
	}
	mysqldb, err := gormMysql.DB()
	if err != nil {
		return nil, err
	}
	mysqldb.SetMaxIdleConns(conf.MaxIdleConns)
	mysqldb.SetMaxOpenConns(conf.MaxOpenConns)
	mysqldb.SetConnMaxLifetime(time.Duration(conf.ConnMaxLifeMinutes) * time.Minute)
	return gormMysql, nil
}

// NewGormPostgres 返回 gorm postgresql 连接实例
func NewGormPostgres(conf DBConfig) (*gorm.DB, error) {
	dsn, err := conf.PostgresDSN()
	if err != nil {
		return nil, err
	}

	if conf.GormConfig == nil {
		conf.GormConfig = &gorm.Config{}
	}

	gormPostgres, err := gorm.Open(postgres.Open(dsn), conf.GormConfig)
	if err != nil {
		return nil, err
	}
	pg, err := gormPostgres.DB()
	pg.SetMaxIdleConns(conf.MaxIdleConns)
	pg.SetMaxOpenConns(conf.MaxOpenConns)
	pg.SetConnMaxLifetime(time.Duration(conf.ConnMaxLifeMinutes) * time.Minute)
	return gormPostgres, nil
}

// NewGormSqlserver 返回 gorm sqlserver 连接实例
func NewGormSqlserver(conf DBConfig) (*gorm.DB, error) {
	dsn, err := conf.SqlserverDSN()
	if err != nil {
		return nil, err
	}

	if conf.GormConfig == nil {
		conf.GormConfig = &gorm.Config{}
	}

	gormSqlserver, err := gorm.Open(sqlserver.Open(dsn), conf.GormConfig)
	if err != nil {
		return nil, err
	}
	sqlserverdb, err := gormSqlserver.DB()
	sqlserverdb.SetMaxIdleConns(conf.MaxIdleConns)
	sqlserverdb.SetMaxOpenConns(conf.MaxOpenConns)
	sqlserverdb.SetConnMaxLifetime(time.Duration(conf.ConnMaxLifeMinutes) * time.Minute)
	return gormSqlserver, nil
}

// LikeFieldEscape 转义 SQL 的 like 模糊查询时字段值为通配符的值
func LikeFieldEscape(value string) string {
	value = strings.Replace(value, ";", "\\;", -1)
	value = strings.Replace(value, "\"", "\\\"", -1)
	value = strings.Replace(value, "'", "\\'", -1)
	value = strings.Replace(value, "--", "\\--", -1)
	value = strings.Replace(value, "/", "\\/", -1)
	value = strings.Replace(value, "%", "\\%", -1)
	value = strings.Replace(value, "_", "\\_", -1)
	return value
}

// GormInstances 以 sync.Map 保存 gorm db 相关信息
// key 为小写的数据库驱动名称， value 为实例名为 key ， 具体的 db 对象为 value 的 sync.Map
// 形如： {"mysql": {"localhost": db}, "postgres": {"localhost": db}}
var GormInstances sync.Map

// GormMySQL 根据 viper 配置中的实例名称返回 gorm 连接 mysql 的实例
func GormMySQL(which string) (*gorm.DB, error) {
	mysqls, loaded := GormInstances.LoadOrStore("mysql", new(sync.Map))
	if loaded {
		if db, loaded := mysqls.(*sync.Map).Load(which); loaded {
			return db.(*gorm.DB), nil
		}
	}
	// mysql 不存在 或 存在时 which 不存在 则新建 mysql db 实例存放到 map 中
	// 注意：这里依赖 viper ，必须在外部先对 viper 配置进行加载
	prefix := "mysql." + which
	conf := DBConfig{
		DriverName:         "",
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
		DisableSSL:         false,
	}
	db, err := NewGormMySQL(conf)
	if err != nil {
		return nil, err
	}
	mysqls.(*sync.Map).Store(which, db)
	GormInstances.Store("mysql", mysqls)
	return db, nil
}

// GormSQLite3 根据  viper 配置中的实例名称返回 sqlite3 实例
func GormSQLite3(which string) (*gorm.DB, error) {
	sqlite3s, loaded := GormInstances.LoadOrStore("sqlite3", new(sync.Map))
	if loaded {
		if db, loaded := sqlite3s.(*sync.Map).Load(which); loaded {
			return db.(*gorm.DB), nil
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
	db, err := NewGormSQLite3(conf)
	if err != nil {
		return nil, err
	}
	sqlite3s.(*sync.Map).Store(which, db)
	GormInstances.Store("sqlite3", sqlite3s)
	return db, nil
}

// GormPostgres 根据 viper 配置中的实例名称返回 pg 实例
func GormPostgres(which string) (*gorm.DB, error) {
	pgs, loaded := GormInstances.LoadOrStore("postgres", new(sync.Map))
	if loaded {
		if db, loaded := pgs.(*sync.Map).Load(which); loaded {
			return db.(*gorm.DB), nil
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
	db, err := NewGormPostgres(conf)
	if err != nil {
		return nil, err
	}
	pgs.(*sync.Map).Store(which, db)
	GormInstances.Store("postgres", pgs)
	return db, nil
}

// GormSqlserver 根据 viper 配置中的实例名称返回 sqlserver 实例
func GormSqlserver(which string) (*gorm.DB, error) {
	sqlservers, loaded := GormInstances.LoadOrStore("sqlserver", new(sync.Map))
	if loaded {
		if db, loaded := sqlservers.(*sync.Map).Load(which); loaded {
			return db.(*gorm.DB), nil
		}
	}
	prefix := "sqlserver." + which
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
	db, err := NewGormSqlserver(conf)
	if err != nil {
		return nil, err
	}
	sqlservers.(*sync.Map).Store(which, db)
	GormInstances.Store("sqlservers", sqlservers)
	return db, nil
}

// CloseGormInstances 关闭全部的 Gorm 连接并重置 GormInstances
func CloseGormInstances() {
	GormInstances.Range(func(ik, iv interface{}) bool {
		if m, ok := iv.(*sync.Map); ok {
			m.Range(func(k, v interface{}) bool {
				if db, ok := v.(*gorm.DB); ok {
					sqlDB, _ := db.DB()
					sqlDB.Close()
				}
				return true
			})
		}
		return true
	})
	GormInstances = sync.Map{}
}

// SetGormInstances 设置 gorm db 对象到 GormInstances 中
func SetGormInstances(driverName, which string, db *gorm.DB) error {
	ins, _ := GormInstances.LoadOrStore(driverName, new(sync.Map))
	ins.(*sync.Map).Store(which, db)
	GormInstances.Store(driverName, ins)
	return nil
}
