package config

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	configPath = flag.String("c", "", "configure file path")
)

type Config struct {
	Port string

	Logger struct {
		Level string
	}

	Mysql struct {
		Dsn             string
		MaxIdleConns    int
		MaxOpenConns    int
		ConnMaxLifetime int
	}

	Redis struct {
		Addr     string
		Password string
		DB       int
		Prefix   string
	}
}

var config Config

func Get() *Config {
	return &config
}

func Init() {
	flag.Parse()
	if *configPath == "" {
		if runtime.GOOS == "windows" {
			*configPath = ".\\config\\config.toml"
		} else {
			*configPath = "./config/config.toml"
		}
	}

	if err := Load(*configPath); err != nil {
		fmt.Printf("load server config file error, %v", err)
		panic(err)
	}

}

func Load(path string) error {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("read server config file error %v", err)
		return err
	}

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("unmarshal config file error %v", err)
		return err
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(&config); err != nil {
			fmt.Printf("change unmarshal config file error, %v", err)
		}
	})

	viper.WatchConfig()

	return nil
}
