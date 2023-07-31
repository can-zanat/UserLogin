package config

import (
	"github.com/k0kubun/pp/v3"
)

type Config struct {
	Mysql      Mysql
	ServerPort string
}

func New() (*Config, error) {
	config := &Config{
		Mysql: Mysql{
			Host:     "localhost",
			Port:     3306,
			Database: "USERLOGINsql",
			UserName: "root",
			Password: "123456",
		}}

	return config, nil
}

func (c *Config) Print() {
	_, _ = pp.Println(c)
}
