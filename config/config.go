package config

type GlobalConfig struct {
	HTTP  HTTPConfig  `yaml:"http"`
	DB    DBConfig    `yaml:"db"`
	Redis RedisConfig `yaml:"redis"`
}

type HTTPConfig struct {
	Port int `yaml:"port"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int32  `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int32  `yaml:"port"`
	Password string `yaml:"password"`
	TTL      int64  `yaml:"ttl"`
}
