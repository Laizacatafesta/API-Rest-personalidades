package main

import (
	"fmt"

	"github.com/Laizacatafesta/API-Rest-personalidades.git/routes"
	"github.com/Laizacatafesta/API-Rest-personalidades.git/routes/models"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Anne", Historia: "Historia 1"},
		{Id: 2, Nome: "Diana", Historia: "Historia 2"},
	}

	fmt.Println("Iniciando")
	routes.HandleRequest()
}
