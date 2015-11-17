package main

import(
	// Standard library packages
	"fmt"
	"net/http"
	"encoding/json"
	
	// Third party packages
	"github.com/julienschmidt/httprouter"
	"github.com/thiagoao/GoLang-Mongodb/models"
)

func main() {
	// Instantiate a new router
	r := httprouter.New()
	
	// Add a handler on /test
	r.GET("/:id", func(w http.ResponseWriter, http *http.Request, p httprouter.Params) {
		// Stub an user example
		u := models.User{
			Name:   "Thiago Oliveira",
			Gender: "Male",
			Age:    32,
			Id:     p.ByName("id"),
		}
		
		// Marshal provided interface into JSON structure
		uj, _ := json.Marshal(u)
		
		// Write content-type, statuscode, payload
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", uj)
	})
	
	// Fire up the server
	http.ListenAndServe(":8080", r)
}