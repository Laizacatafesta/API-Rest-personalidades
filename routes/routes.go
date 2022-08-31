package routes

import (
	"log"
	"net/http"

	"github.com/Laizacatafesta/API-Rest-personalidades.git/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
