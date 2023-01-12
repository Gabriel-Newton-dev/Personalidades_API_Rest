package database

import (
	"fmt"
	"log"

	"github.com/Gabriel-Newton-dev/Personalidades_API_Rest/models"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	DB_USER := viper.Get("DB_USER")
	DB_NAME := viper.Get("DB_NAME")
	DB_PASSWORD := viper.Get("DB_PASSWORD")

	StringDeConexao := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	DB, err := gorm.Open(postgres.Open(StringDeConexao))
	if err != nil {
		log.Fatal(err)
	}
	DB.AutoMigrate(&models.Personalidade{})
}
