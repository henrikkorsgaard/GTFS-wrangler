package main
 
import (
  "log"

  "net/http"
  "fmt"

  "github.com/gorilla/mux"
)


/* 
Deploy via Docker
*/

func main() {
  r := mux.NewRouter()

  r.HandleFunc("/", appHandler)
  r.HandleFunc("/gtfs",gtfsHandler).Methods("POST")


  r.PathPrefix("/").Handler(http.FileServer(http.Dir("client/public/")))

  srv := &http.Server{
		Handler: r,
		Addr:    ":8080",
	}

  log.Fatal(srv.ListenAndServe())
}


func appHandler(w http.ResponseWriter, req *http.Request){
	http.ServeFile(w, req, "client/gtfs-wrangler.html")
}


func gtfsHandler(w http.ResponseWriter, req *http.Request){
  fmt.Println("File Upload Endpoint Hit")


  fmt.Fprintf(w, "Successfully Uploaded File\n")

}
