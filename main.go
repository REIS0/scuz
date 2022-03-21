package main

import (
	"fmt"

	"github.com/REIS0/scuz/database"
	"github.com/REIS0/scuz/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	err := database.InitDatabase("scuz.db")
	if err != nil {
		fmt.Println("database error")
		return
	}

	utils.SetUpRoutes(app)

	app.Listen(":3000")
}
