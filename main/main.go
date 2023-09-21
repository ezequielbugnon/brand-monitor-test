package main

import (
	"fmt"

	"github.com/ezequiel-bugnon/brandmonitor/controllers"
	"github.com/ezequiel-bugnon/brandmonitor/db"
	"github.com/ezequiel-bugnon/brandmonitor/frameworks"
	"github.com/ezequiel-bugnon/brandmonitor/repository"
	"github.com/ezequiel-bugnon/brandmonitor/services"
	"github.com/gofiber/fiber"
)

func main() {
	database, err := db.NewMongoDB()
	if err != nil {
		fmt.Println("Error database connection", err)
		return
	}

	repository := repository.NewRepository(*database)
	service := services.NewService(repository)
	controller := controllers.NewController(service)
	framework := fiber.New()
	app := frameworks.NewApp(controller, framework)
	app.Controllers()
	app.Conection()
}
