package main
 
import (
  "fmt"
  
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/cors" // not sure we need this anymore.
  "github.com/gofiber/websocket/v2"

)
// Let's try out fiber
// https://gofiber.io/


func main() {

    app := fiber.New(fiber.Config{
      BodyLimit: 100 * 1024 * 1024,
    })

    app.Static("/public","./public")

    //introduce new here
    app.Use(cors.New())
    app.Use("/ws", websocketUpgrader)

    app.Get("/ws/gtfs", websocket.New(websocketGTFSHandler))
    
    app.Get("/", mainPage) 
    app.Post("/gtfs", fileUploadHandler)
    app.Listen(":3000")
}

func mainPage(c *fiber.Ctx) error {
  fmt.Println("what")
  return c.SendString("Hello, World!")
}

// need to look at error handling. That's gotta be a small thesis for me to grokk by now :)
func websocketGTFSHandler(c *websocket.Conn) {
  
  for {
    _, message, err := c.ReadMessage()
    if err != nil {
      if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
        fmt.Println("read error:", err)
      }
      
    }
    
    fmt.Println(message)
    

  }
}

// need to look at error handling. That's gotta be a small thesis for me to grokk by now :)
func websocketUpgrader(c *fiber.Ctx) error {
    //From gofiber docs https://github.com/gofiber/websocket
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
      fmt.Println("Rat in the trap!")
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
}

func fileUploadHandler(c *fiber.Ctx) error {
  fmt.Println("ever")
  file, err := c.FormFile("file")
  if err != nil {
      panic(err)
  }
  fmt.Println(file)
  return c.SendString("Hello, World!")
}

