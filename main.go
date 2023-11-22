package main

import (
	"octagon/route"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	route.RouteInit(app)

	app.Listen("localhost:3000")
}
