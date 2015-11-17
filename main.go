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
	
	r.POST("/user", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		// Stub an user to be populated from the body		
		u := models.User{}
		
		// Populated the user data
		json.NewDecoder(r.Body).Decode(&u)
		
		// Add an Id
		u.Id = "foo"
		
		uj, _ := json.Marshal(u)
		
		//Write Content-Type, statuscode, payload
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		fmt.Fprintf(w, "%s", uj)
	})
	
	r.DELETE("/user/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) { 
		// TODO: only write status for now
		w.WriteHeader(200)
	})
	
	// Add a handler on /test
	r.GET("/user/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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