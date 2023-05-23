package main
 
import (
  "github.com/gofiber/fiber/v2"

)
// Let's try out fiber
// https://gofiber.io/

func main() {
    app := fiber.New()
    app.Static("/public","./public")
    app.Get("/", mainPage)
    app.Listen(":3000")
}

func mainPage(c *fiber.Ctx) error {
  return c.SendString("Hello, World!")
}

