package goutils

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// DBConfig 数据库配置
type DBConfig struct {
	// DriverName 数据库 driver 类型
	DriverName string
	// Host 数据库 IP 地址
	Host string
	// Port 数据库端口
	Port int
	// Username 数据库用户名
	Username string
	// Password 数据库密码
	Password string
	// DBName 数据库名称
	DBName string
	// LogMode 是否开启打印日志模式
	LogMode bool
	// MaxIdleConns 设置空闲连接池中的最大连接数
	MaxIdleConns int
	// MaxOpenConns 设置与数据库的最大打开连接数
	MaxOpenConns int
	// ConnMaxLifeMinutes 设置可重用连接的最长时间（分钟）
	ConnMaxLifeMinutes int
	// ConnTimeout 连接超时时间（秒）
	ConnTimeout int
	// ReadTimeout 读超时时间（秒）
	ReadTimeout int
	// WriteTimeout 写超时时间（秒）
	WriteTimeout int
	// DisableSSL 是否关闭 ssl 模式
	DisableSSL bool
	// gorm config
	GormConfig *gorm.Config
}

// MySQLDSN 返回 Mysql dsn 字符串
func (conf DBConfig) MySQLDSN() (string, error) {
	if conf.Username == "" || conf.Host == "" || conf.Port == 0 || conf.DBName == "" {
		return "", errors.New("DBConfig for MySQL is empty")
	}

	return fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%ds&readTimeout=%ds&writeTimeout=%ds", conf.Username, conf.Password, conf.Host, conf.Port, conf.DBName, conf.ConnTimeout, conf.ReadTimeout, conf.WriteTimeout), nil
}

// SQLite3DSN 返回 sqlite3 文件名
func (conf DBConfig) SQLite3DSN() (string, error) {
	if conf.DBName == "" {
		return "", errors.New("DBConfig for SQLite3 is empty")
	}

	return conf.DBName, nil
}

// PostgresDSN 返回 postgres dns 字符串
func (conf DBConfig) PostgresDSN() (string, error) {
	if conf.Host == "" || conf.Port == 0 || conf.Username == "" || conf.DBName == "" {
		return "", errors.New("DBConfig for Postgres is empty")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", conf.Host, conf.Port, conf.Username, conf.DBName, conf.Password)
	if conf.DisableSSL {
		dsn = dsn + " sslmode=disable"
	}
	return dsn, nil
}

// SqlserverDSN 返回 sqlserver dns 字符串
func (conf DBConfig) SqlserverDSN() (string, error) {
	if conf.Host == "" || conf.Port == 0 || conf.Username == "" || conf.DBName == "" {
		return "", errors.New("DBConfig for Sqlserver is empty")
	}
	return fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", conf.Username, conf.Password, conf.Host, conf.Port, conf.DBName), nil
}
