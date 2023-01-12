package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Gabriel-Newton-dev/Personalidades_API_Rest/database"
	"github.com/Gabriel-Newton-dev/Personalidades_API_Rest/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var personalidades []models.Personalidade
	database.DB.Find(&personalidades)
	json.NewEncoder(w).Encode(personalidades)
}
