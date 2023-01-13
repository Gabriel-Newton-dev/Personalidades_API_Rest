package main

import (
	"fmt"

	"github.com/Gabriel-Newton-dev/Personalidades_API_Rest/database"
	"github.com/Gabriel-Newton-dev/Personalidades_API_Rest/models"
	"github.com/Gabriel-Newton-dev/Personalidades_API_Rest/routes"
	"github.com/spf13/viper"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "Gabriel Newton", Historia: "Desenvolvedor Pleno"},
		{Nome: "Deise Lima", Historia: "Esposa do Gabriel"},
	}

	fmt.Println("Iniciando a nossa API")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	database.ConectaComBancoDeDados()
	routes.HandleRequest()

}
