package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	RedisAddr string   `mapstructure:"redis_addr"`
	RedisPass string   `mapstructure:"redis_pass"`
	RedisDB   int      `mapstructure:"redis_db"`
	Servers   []string `mapstructure:"servers"`
	Port      int      `mapstructure:"port"`
	RateLimit int      `mapstructure:"rate_limit"`
	Logging   bool     `mapstructure:"logging"`
}

func (conf *Config) String() string {
	return fmt.Sprintf("[port=%d|servers=%v|rate_limit=%d|logging=%v]", conf.Port, conf.Servers, conf.RateLimit, conf.Logging)
}

func Load(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("system.conf")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err == nil {
		err = viper.Unmarshal(&config)
	}
	return
}
