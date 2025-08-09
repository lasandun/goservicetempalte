package config

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port int    `mapstructure:"port"`
	Host string `mapstructure:"host"`
}

type DatabaseConfig struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
}

type AppConfig struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
}

var Cfg AppConfig

func loadConfigFromFile(configPath string) error {
	if configPath == "" {
		configPath = "config/config.yaml" // default path
	}

	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}

	if err := viper.Unmarshal(&Cfg); err != nil {
		return fmt.Errorf("unable to decode config into struct: %w", err)
	}

	return nil
}

func LoadConfig() {
	// Define CLI flag
	pflag.String("config", "config/config.yaml", "Path to config file")
	pflag.Parse()

	// Bind CLI flag to viper key
	if err := viper.BindPFlag("config", pflag.Lookup("config")); err != nil {
		panic(err)
	}

	// Pass the config file path from viper to your config loader
	if err := loadConfigFromFile(viper.GetString("config")); err != nil {
		panic(err)
	}

	fmt.Println("Server will start at:", Cfg.Server.Host, ":", Cfg.Server.Port)
	fmt.Println("Database user:", Cfg.Database.User)
}
