package db

import (
	"fmt"

	"github.com/spf13/viper"
)

var DB interface{}

func init() {
	viper.SetConfigName("database") // config file name without extension
	viper.SetConfigType("yml")
	viper.AddConfigPath("./configs/") // config file path
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	database := viper.GetString("development.database")
	username := viper.GetString("development.username")

	fmt.Println("viper configs database: ", database, username)
}
