package main
 
import (
  "fmt"
  
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/cors" // not sure we need this anymore.
  "github.com/gofiber/websocket/v2"

  "henrikkorsgaard.dk/GTFS-wrangler/gtfs"

)
// Let's try out fiber
// https://gofiber.io/


// We also need to understand this:https://go.dev/blog/context
// https://pkg.go.dev/context

// We do not need the server to be able to parse GTFS from the front-end
// It will do so from the backend
// https://www.rejseplanen.info/labs/GTFS.zip
// -> if they move this, then we need to react.
// -> we still need to be able to gulp these files
// -> we still need to be able to check if this is a new download

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
    
    gtfsDir, err := gtfs.UnzipGTFSFromBytes(message)
    
    if err != nil {
      fmt.Println(err)
    }

    //Now I wonder if it makes sense to return a list of paths instead. Then we can check if there are any and then ignore emitting error from the unzipper!
    fmt.Println(gtfsDir)
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
  fmt.Println(file) // ok, we got it. Back to making tests and implementation
  return c.SendString("Hello, World!")
}

