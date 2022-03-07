package main

import (
	"example/hello/src/config"

	"github.com/gofiber/fiber/v2"
	"go.elastic.co/apm/module/apmfiber"
)

type LoginDTO struct {
	UserName string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
	Role     string `json:"role" xml:"role" form:"role"`
}

func login(c *fiber.Ctx) error {
	p := new(LoginDTO)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	return c.SendString("Hello " + p.UserName)
}

func main() {
	config.ConfigureAPM()
	app := fiber.New()
	app.Use(apmfiber.Middleware())

	app.Post("/login", login)

	app.Listen(":3000")
}
