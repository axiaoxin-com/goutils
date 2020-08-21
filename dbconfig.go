package goutils

// MySQLConfig mysql 连接配置
type MySQLConfig struct {
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
}

// SQLite3Config sqlite3 连接配置
type SQLite3Config struct {
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
}

// PostgresConfig postgresql 连接配置
type PostgresConfig struct {
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
	// DisableSSL 是否关闭 ssl 模式
	DisableSSL bool
	// LogMode 是否开启打印日志模式
	LogMode bool
	// MaxIdleConns 设置空闲连接池中的最大连接数
	MaxIdleConns int
	// MaxOpenConns 设置与数据库的最大打开连接数
	MaxOpenConns int
	// ConnMaxLifeMinutes 设置可重用连接的最长时间（分钟）
	ConnMaxLifeMinutes int
}

// MsSQLConfig sqlserver 连接配置
type MsSQLConfig struct {
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
}
