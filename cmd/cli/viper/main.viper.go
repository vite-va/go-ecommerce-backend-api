package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port string `mapstructure:"port"`
	} `mapstructur:"server"`
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"databases"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./config/") //path to config
	viper.SetConfigName("local")     // file name
	viper.SetConfigType("yaml")

	// read configuration file
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fail to read configuration %w", err))
	}

	// read server configuration
	fmt.Println("Server port::", viper.GetInt("server.port"))
	fmt.Println("Security jwt key::", viper.GetString("security.jwt.key"))

	// configure structure
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}

	fmt.Println("Config Server Port::", config.Server.Port)
	for _, db := range config.Databases {
		fmt.Printf("Database %v:: host: %v, password: %v \n", db.User, db.Host, db.Password)
	}
}
