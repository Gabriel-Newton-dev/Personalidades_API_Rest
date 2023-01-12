package main

import (
	"fmt"

	"github.com/Gabriel-Newton-dev/Personalidades_API_Rest/database"
	"github.com/Gabriel-Newton-dev/Personalidades_API_Rest/routes"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("Iniciando a nossa API")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	database.ConectaComBancoDeDados()
	routes.HandleRequest()

}
