package utils

import (
	"github.com/REIS0/scuz/models"
	"github.com/REIS0/scuz/repository"
	"github.com/gofiber/fiber/v2"
)

func testRoute(c *fiber.Ctx) error {
	return c.SendString("Funcionando")
}

func SetUpRoutes(app *fiber.App) {
	app.Get("/", testRoute)
	app.Get("/users/:username", getUser)
	app.Post("/users", postUser)
}

func getUser(c *fiber.Ctx) error {
	username := c.Params("username")
	user := repository.GetUser(username)
	return c.JSON(user)
}

func postUser(c *fiber.Ctx) error {
	newUser := new(models.User)
	if err := c.BodyParser(newUser); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	repository.CreateUser(*newUser)
	return c.Status(201).SendString("User created")
}

// func Login(login string, password string) (models.User, error) {
// 	user := repository.GetUser(login)

// 	check := user.CheckPassword(password)

// 	if !check {
// 		return models.User{}, errors.New("wrong password")
// 	}

// 	return user, nil
// }
