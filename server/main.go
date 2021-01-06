package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	//"go.mongodb.org/mongo-driver/bson"
	"github.com/asaskevich/govalidator"
	"github.com/dchest/uniuri"
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
	fmt.Fprint(res, "boop")
}

type url struct {
	URL  string `json:"url"`
	UUID string `json:"uuid"`
}

//TODO: Split this up, clean up error handling
func postURLController(res http.ResponseWriter, req *http.Request) {
	var data url
	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	if govalidator.IsURL(data.URL) == false {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(res, "Invalid URL")
		return
	}
	uuid := uniuri.NewLen(5)
	data.UUID = uuid
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err2 := urls.InsertOne(ctx, data)
	if err2 != nil {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(res, "DB Error")
		return
	}
	res.WriteHeader(http.StatusAccepted)
	json.NewEncoder(res).Encode(map[string]string{"uuid": uuid})
}

func main() {
	fmt.Println("Server started!")
	urls = initDBClient()
	router := mux.NewRouter()
	router.HandleFunc("/beep", beepController).Methods("GET")
	router.HandleFunc("/url", postURLController).Methods("POST")
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
