package models

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	// User represents the structure of our resource
	User struct {
		Id         bson.ObjectId `json:"id" bson:"_id"`
		Name       string        `json:"name" bson:"name"`
		Email      string 		 `json:"email" bson:"email"`
		Gender     string        `json:"gender" bson:"gender"`
		Age        int           `json:"age" bson:"age"`
		UserAdress Address       `json:"address" bson:"address"`
	}
	
	Address struct {
		Street  string `json:"street" bson:"street"`
		Number  string `json:"number" bson:"number,omitempty"`
		ZipCode string `json:"zipcode" bson:"zipcode"`
		City    string `json:"city" bson:"city"`
		State   State  `json:"state" bson:"state"`
	}
	
	State struct {
		Id           bson.ObjectId `json:"id" bson:"_id"`
		Name         string        `json:"name" bson:"name"`
		StateCountry Country       `json:"country" bson:"country"`
	}
	
	Country struct {
		Id           bson.ObjectId `json:"id" bson:"_id"`
		Name         string        `json:"name" bson:"name"`
	}
)