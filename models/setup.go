package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func SetupModels() *gorm.DB {
	// viper.AddConfigPath("../")
	// viper.AutomaticEnv()
	// viper_user := viper.Get("POSTGRES_USER")
	// viper_password := viper.Get("POSTGRES_PASSWORD")
	// viper_db := viper.Get("POSTGRES_DB")
	// viper_host := viper.Get("POSTGRES_HOST")
	// viper_port := viper.Get("POSTGRES_PORT")
	viper_user := "postgres"
	viper_password := "postgres"
	viper_db := "work_db"
	viper_host := "127.0.0.1"
	viper_port := "5432"

	progret_conname := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", viper_host, viper_port, viper_user, viper_db, viper_password)
	fmt.Println("conname is \t\t", progret_conname)
	db, err := gorm.Open("postgres", progret_conname)
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&Book{})

	m := Book{Author: "author1", Title: "title1"}
	db.Create(&m)

	return db
}
