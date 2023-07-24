package main
 
import (
  "fmt"
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/cors" // not sure we need this anymore.
  "github.com/gofiber/websocket/v2"

  "henrikkorsgaard.dk/GTFS-wrangler/gtfs"

)

type WebsocketServerResponse struct {
    Type string
    Message string 
    Payload interface{}
}

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
    fmt.Println("running")
    app := fiber.New(fiber.Config{
      BodyLimit: 100 * 1024 * 1024,
    })

    app.Static("/public","./public")

    app.Use(cors.New())
    app.Use("/ws", websocketUpgrader)

    app.Get("/ws/gtfs", websocket.New(websocketGTFSUploadHandler))
    
    app.Get("/", mainPage) 
    err := app.Listen(":3000")
    if err != nil {
      panic(err)
    }
}

func mainPage(c *fiber.Ctx) error {
  return c.SendString("Hello, World!")
}

// need to look at error handling. That's gotta be a small thesis for me to grokk by now :)
func websocketGTFSUploadHandler(c *websocket.Conn) {
  var zbytes []byte 
  
  // TODO:
  // We need to find out what happens when a file is corrupt/errors happen in parsing.
  // Do we roll back everything
  // Do we skip the file?

  for {  
    messageType, message, err := c.ReadMessage()
    if err != nil {
      if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
        fmt.Println("Read error on websocket connection")
        // we panic on this error to force us to identify all the bizarre error codes that might happen.
        // eventually, we should log these somewhere
        panic(err) 
      } 
      c.Close()
      break
    }
   
    if messageType == websocket.BinaryMessage {
      zbytes = message 

      if err = c.WriteJSON(WebsocketServerResponse{Type:"request", Message:"filename"}); err != nil {
				fmt.Println("write:", err)
        c.Close()
				break
			}
    }

    if messageType == websocket.TextMessage {
      filename := string(message)
      if len(zbytes) > 0 {
        
        messages := make(chan gtfs.GTFSLoadProgress)
	      errorChannel := make(chan error)

        go func(){
          for {
            select {
              case msg := <-messages:
                //progResponse, err := json.Marshal(msg)
                wsResponse := WebsocketServerResponse{
                  Type: "progress_info",
                  Payload: msg,
                }

                if err = c.WriteJSON(wsResponse); err != nil { // needs to be a json btw.
                  fmt.Println("write:", err)
                  c.Close()
                  break
                }
                
                if msg.Filename == filename && msg.Done {
                  return
                }
            
              case err = <-errorChannel:
                // handle errors and 
                fmt.Println(err)
                wsResponse := WebsocketServerResponse{
                  Type: "data_error",
                  Message: err.Error(),
                }
                if err = c.WriteJSON(wsResponse); err != nil { // needs to be a json btw.
                  fmt.Println("write:", err)
                }
                // do we break on these? 
                // it depends I guess
                //c.Close()
                break
            }
          }
        }()

        gtfs.NewGTFSFromZipBytes(filename,zbytes, messages, errorChannel)
      }
    }
  }
}

// need to look at error handling. That's gotta be a small thesis for me to grokk by now :)
func websocketUpgrader(c *fiber.Ctx) error {
    //From gofiber docs https://github.com/gofiber/websocket
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
}

