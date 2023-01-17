package routes

import (
	"log"
	"net/http"

	"github.com/Gabriel-Newton-dev/Personalidades_API_Rest/controllers"
	"github.com/gorilla/mux"
)

// gorilla Mux é um poderoso orquestrador de request(requisicao)http.
// Temos que criar uma isntancia dele r := mux.NewRouter(), vai criar um novo mapeamento de rotas.

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.BuscarPorId).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
