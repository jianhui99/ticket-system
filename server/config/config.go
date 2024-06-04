package config

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf *Config

type Config struct {
	App      AppConf  `mapstructure:"app"`
	Log      LogConf  `mapstructure:"log"`
	Port     int      `mapstructure:"port"`
	Database Database `mapstructure:"database"`
	Jwt      JwtConf  `mapstructure:"jwt"`
}

type AppConf struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
	Env     string `mapstructure:"env"`
}

type JwtConf struct {
	Secret string `mapstructure:"secret"`
	Exp    int64  `mapstructure:"exp"`
}
type LogConf struct {
	Level string `mapstructure:"level"`
}

type Database struct {
	MySqlConf MySqlConf `mapstructure:"mysql"`
	RedisConf RedisConf `mapstructure:"redis"`
}

type MySqlConf struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	DbName       string `mapstructure:"db"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	MaxOpenConns int    `mapstructure:"maxOpenConns"`
	MaxIdleConns int    `mapstructure:"maxIdleConns"`
}

type RedisConf struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"poolSize"`
}

func Init(confFile string) {
	Conf = new(Config)
	v := viper.New()
	v.SetConfigFile(confFile)
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		log.Println("配置文件被修改了")
		err := v.Unmarshal(&Conf)
		if err != nil {
			panic(fmt.Errorf("unmarshal change config data,err:%v", err))
		}
	})
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("读取配置文件出错,err:%v", err))
	}
	err = v.Unmarshal(&Conf)
	if err != nil {
		panic(fmt.Errorf("unmarshal config data,err:%v", err))
	}
}
