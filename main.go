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
	
	r.POST("/user", uc.CreateUser)
	
	r.DELETE("/user/:id", uc.RemoveUser)
	
	// Fire up the server
	http.ListenAndServe(":8080", r)
}

func getSession() *mgo.Session {
	// Connect to your mongo
	s, err := mgo.Dial("mongodb://golang:golang@ds053944.mongolab.com:53944/golang")
	
	if err != nil {
		panic(err)
	}
	
	return s
}