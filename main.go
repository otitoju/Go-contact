package main

import (
	"fmt"
	"time"
	"context"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
// model
type Person struct {
	ID primitive.ObjectID `json:"_id, omitempty" bson:"_id, omitempty"`
	Firstname string `json:"firstname, omitempty" bson:"firstname, omitempty"`
	Lastname  string `json:lastname, omitempty" bson:lastname, omitempty"`
}

var client *mongo.Client

// end point for creating data
func registerUsers(res http.ResponseWriter, req *http.Request){
	// response type should be in json format
	res.Header().Add("content-type", "application/json")
	var person Person
	json.NewDecoder(req.Body).Decode(&person)
	//create database name
	collection := client.Database("GoDb").Collection("people")
	//set timeout
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, person)
	json.NewEncoder(res).Encode(result)
}

func main() {
	fmt.Println("Server running")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil { 
		fmt.Println("Error")
	}
	router := mux.NewRouter()
	router.HandleFunc("/register", registerUsers).Methods("POST")
	http.ListenAndServe(":4000", router)
}