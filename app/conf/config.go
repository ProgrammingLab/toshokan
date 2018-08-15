package conf

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Debug bool     `toml:"debug"`
	DB    DBConfig `toml:"DB"`
}

type DBConfig struct {
	DataSourceName string `toml:"dataSourceName"`
}

var config = loadConfig()

func loadConfig() *Config {
	tmp := &Config{}
	_, err := toml.DecodeFile("toshokan.toml", tmp)
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	return tmp
}

func GetConfig() *Config {
	return config
}
