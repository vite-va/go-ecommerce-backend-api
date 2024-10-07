package initialize

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/vite-va/go-ecommerce-backend-api/global"
)

func LoadConfig() {
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
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}
}
