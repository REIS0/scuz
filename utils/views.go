package utils

import (
	"github.com/REIS0/scuz/models"
	"github.com/REIS0/scuz/repository"
	"github.com/gofiber/fiber/v2"
)

type Views struct {
	App         *fiber.App
	URepository *repository.UserRepository
}

// Setup
func (v *Views) setUpRoutes(app *fiber.App) {
	app.Get("/", v.testRoute)
	app.Get("/users/:username", v.getUser)
	app.Post("/users", v.postUser)
}

func (v *Views) InitRoutes() {
	v.setUpRoutes(v.App)
	v.App.Listen(":3000")
}

// ** ROTAS **
func (v *Views) testRoute(c *fiber.Ctx) error {
	return c.SendString("Funcionando")
}

func (v *Views) getUser(c *fiber.Ctx) error {
	username := c.Params("username")
	user := v.URepository.GetUser(username)
	return c.JSON(user)
}

func (v *Views) postUser(c *fiber.Ctx) error {
	newUser := new(models.User)
	if err := c.BodyParser(newUser); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	v.URepository.CreateUser(*newUser)
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
