package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/ezequiel-bugnon/brandmonitor/controllers"
	"github.com/ezequiel-bugnon/brandmonitor/frameworks"
	"github.com/ezequiel-bugnon/brandmonitor/repository"
	"github.com/ezequiel-bugnon/brandmonitor/services"
)

func main() {
	file, err := os.Open("example.csv")
	if err != nil {
		fmt.Println("Error al abrir el archivo CSV:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Leer las filas del CSV y crear instancias de Employee
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		println(record[0])
	}

	repository := repository.NewRepository("db")
	service := services.NewService(repository)
	controller := controllers.NewController(service)
	app := frameworks.NewApp(controller)
	app.Controllers()
	app.Conection()

}
