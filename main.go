package main

import "github.com/gofiber/fiber/v2"

type HelloWorld struct {
	Say string
}

func getHelloWorld(ctx *fiber.Ctx) error {
	helloWorld := HelloWorld{
		"Hello World",
	}
	return ctx.Status(fiber.StatusOK).JSON(helloWorld)
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/hello", getHelloWorld)

	app.Listen(":3000")
}
