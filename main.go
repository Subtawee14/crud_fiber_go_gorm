package main

import (
	"log"

	"crud_fiber_go_gorm/configs"
	"crud_fiber_go_gorm/routes"

	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//config for customization
	app.Use(cors.New(configs.ConfigDefault))

	configs.InitDatabase()
	routes.SetupRoutes(app)
	log.Fatal(app.Listen(":8080"))
}
