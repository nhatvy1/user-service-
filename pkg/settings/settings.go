package settings

type Config struct {
	Server   ServerSetting `mapstructure:"server"`
	Database DbSetting     `mapstructure:"database"`
	Redis    RedisSetting  `mapstructure:"redis"`
}
type ServerSetting struct {
	Name string `mapstructure:"name"`
	Port int    `mapstructure:"port"`
}

type DbSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"dbname"`
	MaxIdle  int    `mapstructure:"max_idle_conns"`
	MaxOpen  int    `mapstructure:"max_open_conns"`
	MaxLife  int    `mapstructure:"conn_max_lifetime"`
}

type RedisSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type LogSetting struct {
	Level      string `mapstructure:"log_level"`
	Filename   string `mapstructure:"file_log_name"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
}
