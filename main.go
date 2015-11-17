package main

import(
	// Standard library packages
	"net/http"
	
	// Third party packages
	"github.com/julienschmidt/httprouter"
	"github.com/thiagoao/GoLang-Mongodb/controllers"
)

func main() {
	// Instantiate a new router
	r := httprouter.New()
	
	uc := controllers.NewUserController()
	
	// Get a user resource
	r.GET("/user/:id", uc.GetUser)
	
	r.POST("/user", uc.CreateUser)
	
	r.DELETE("/user/:id", uc.RemoveUser)
	
	// Fire up the server
	http.ListenAndServe(":8080", r)
}