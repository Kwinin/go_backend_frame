package server

type serverConfig struct {
	Debug          bool           `yaml:"debug"`
	Secret         string         `required:"true" yaml:"jwt_secret"`
	Statistics     int64          `required:"true" yaml:"statistics"`
	DbConfig       *DBConfig      `required:"true" yaml:"mysql"`
	SwagConfig     SwagConfig     `yaml:"swag"`
	LogLevel       string         `yaml:"log_level"`
	ResourceConfig ResourceConfig `yaml:"resource"`
}

type DBConfig struct {
	Address  string `yaml:"address"`
	User     string `default:"root" yaml:"user"`
	Password string `default:"" yaml:"password"`
	Port     uint   `default:"3306" yaml:"port"`
	DbName   string `required:"true" yaml:"db_name"`
	Charset  string `default:"utf8" yaml:"charset"`
	MaxIdle  int    `default:"10" yaml:"max_idle"`
	MaxOpen  int    `default:"50" yaml:"max_open"`
	LogMode  bool   `yaml:"log_mode"`
	Loc      string `required:"true" yaml:"loc"`
}

type SwagConfig struct {
	Host     string `yaml:"host"`
	BasePath string `default:"/api" yaml:"base_path"`
}

type ResourceConfig struct {
	FilePath string `default:"./data/file" yaml:"file_path"`
}
