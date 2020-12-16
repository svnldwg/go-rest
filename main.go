package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Dish struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

// let's declare a global array
// that we can then populate in our main function
// to simulate a database
var Dishes []Dish

func returnAllDishes(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllDishes")
	json.NewEncoder(w).Encode(Dishes)
}

func returnSingleDish(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, dish := range Dishes {
		if dish.Id == key {
			json.NewEncoder(w).Encode(dish)
		}
	}
}

func createNewDish(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)

	var dish Dish
	json.Unmarshal(reqBody, &dish)
	Dishes = append(Dishes, dish)

	json.NewEncoder(w).Encode(dish)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/dishes", returnAllDishes)
	myRouter.HandleFunc("/dish", createNewDish).Methods("POST")
	myRouter.HandleFunc("/dish/{id}", returnSingleDish)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Dishes = []Dish{
		Dish{Id: "1", Title: "Tacos", Desc: "Maistacos, Hackfleisch, Tomaten, Guacamole"},
		Dish{Id: "2", Title: "Tamales", Desc: "Auch lecker"},
	}

	handleRequests()
}
