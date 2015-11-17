package controllers

import(
	"fmt"
	"net/http"
	"encoding/json"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/julienschmidt/httprouter"
	"github.com/thiagoao/GoLang-Mongodb/models"
)

type (
	// UserController represents the controller for operating on the User resource
	UserController struct{
		session *mgo.Session
	}
)

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

// GetUser retrieves an individual user resouce
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")
	
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}
	
	// Grab id
	oid := bson.ObjectIdHex(id)

	// Stub an example user
	u := models.User{}
	
	// Fetch user
	if err := uc.session.DB("golang").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
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
	err := json.NewDecoder(r.Body).Decode(&u)
	
	if err != nil {
		panic(err)
		w.WriteHeader(404)
		return
	}
	
	// Add an Id
	u.Id = bson.NewObjectId()
	
	// Write the user to mongo
	uc.session.DB("golang").C("users").Insert(u)
	
	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)
	
	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}

// RemoveUser removes an existing user resource
func (uc UserController) RemoveUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")
	
	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}
	
	// Grab id
	oid := bson.ObjectIdHex(id)
	
	// Remove user
	if err := uc.session.DB("golang").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}
	
	w.WriteHeader(200)
}