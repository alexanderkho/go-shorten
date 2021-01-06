package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	//"go.mongodb.org/mongo-driver/bson"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var urls *mongo.Collection

func initDBClient() *mongo.Collection {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	return client.Database("go-shorten").Collection("urls")
}

func main() {
	fmt.Println("Server started!")
	urls = initDBClient()
	router := mux.NewRouter()
	router.HandleFunc("/beep", beepController).Methods("GET")
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

func beepController(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("boop\n"))
}
