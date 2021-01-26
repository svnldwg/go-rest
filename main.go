package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
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
	fmt.Println("Establish DB Connection")

	db, err := sql.Open("mysql", "go-rest:pass@tcp(127.0.0.1:3306)/go-rest")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

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

func updateDish(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)

	var dishUpdate Dish
	json.Unmarshal(reqBody, &dishUpdate)

	for index, dish := range Dishes {
		if dish.Id == id {
			Dishes[index].Title = dishUpdate.Title
			Dishes[index].Desc = dishUpdate.Desc
			dishUpdate = Dishes[index]
		}
	}

	json.NewEncoder(w).Encode(dishUpdate)
}

func deleteDish(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, dish := range Dishes {
		if dish.Id == id {
			// updates our Dishes array to remove the dish
			Dishes = append(Dishes[:index], Dishes[index+1:]...)
		}
	}
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
	myRouter.HandleFunc("/dish/{id}", updateDish).Methods("PUT")
	myRouter.HandleFunc("/dish/{id}", deleteDish).Methods("DELETE")
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
