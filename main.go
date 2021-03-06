package main

import(
	// Standard library packages
	"net/http"
	
	// Third party packages
	"gopkg.in/mgo.v2"
	"github.com/julienschmidt/httprouter"
	"github.com/thiagoao/GoLang-Mongodb/controllers"
)

func main() {
	// Instantiate a new router
	r := httprouter.New()
	
	// Get a UserController instance
	uc := controllers.NewUserController(getSession())
	
	// Get a user resource
	r.GET("/user/:id", uc.GetUser)
	
	// Create a new user
	r.POST("/user", uc.CreateUser)
	
	// Remove an existing user
	r.DELETE("/user/:id", uc.RemoveUser)
	
	// Fire up the server
	http.ListenAndServe(":8080", r)
}

// getSession creates a new mongo session and panics if connection error occurs
func getSession() *mgo.Session {
	// Connect to your mongo
	s, err := mgo.Dial("mongodb://{user}:{password}@{host}:{port}/{dbname}")
	
	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	
	// Deliver session
	return s
}
