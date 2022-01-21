package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func init() {
	viper.SetConfigName("database") // config file name without extension
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config/") // config file path
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	database := viper.GetString("development.database")
	username := viper.GetString("development.username")
	port := viper.GetInt("development.port")
	host := viper.GetString("development.host")
	password := viper.GetString("development.password")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
	)
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	boil.SetDB(conn)
	boil.DebugMode = true

	// seed数据
	users_count := Users().CountP(conn)
	fmt.Printf("user_count:", users_count)
	if users_count == 0 {
		list := []struct {
			name string
			age  null.Int
		}{
			{"jack", null.IntFrom(19)},
			{"su", null.IntFrom(16)},
			{"mia", null.IntFrom(6)},
		}

		for _, v := range list {
			u := User{Name: v.name, Age: v.age}
			u.Insert(conn, boil.Infer())
		}
	}
}
