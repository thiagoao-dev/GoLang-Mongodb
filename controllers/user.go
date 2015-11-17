package controllers

import(
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/julienschmidt/httprouter"
	"githb.com/thiagoao/GoLang-Mongodb/models"
)

type (
	// UserController represents the controller for operating on the User resource
	UserController struct{}
)

func NewUserController() *UserController {
	return &UserController{}
}

// GetUser retrieves an individual user resouce
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Stub an example user
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
}

// CreateUser creates a new user resource
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Stub an user to be populated from the body
	u := models.User{}
	
	// Populate the user data
	json.NewDecoder(r.Body).Decode(&u)
	
	// Add an Id
	u.Id = "foo"
	
	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)
	
	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}

// RemoveUser removes an existing user resource
func (uc UserController) RemoveUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// TODO: only write status code now
	w.WriteHeader(200)
}