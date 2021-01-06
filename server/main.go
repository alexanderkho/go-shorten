package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
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

func beepController(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("boop\n"))
}

func postURLController(res http.ResponseWriter, req *http.Request) {
	ok := isValidURL("walmart.com")
	res.Write([]byte(strconv.FormatBool(ok) + "\n"))
}

func isValidURL(str string) bool {
	fmt.Println(str)
	u, err := url.ParseRequestURI(str)
	fmt.Println(err)
	fmt.Println(u.Host)
	return err == nil && u.Host != ""
}

func main() {
	fmt.Println("Server started!")
	urls = initDBClient()
	router := mux.NewRouter()
	router.HandleFunc("/beep", beepController).Methods("GET")
	router.HandleFunc("/url", postURLController).Methods("POST")
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
