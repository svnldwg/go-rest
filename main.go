package main

import (
	"encoding/json"
    "fmt"
    "log"
    "net/http"
)

type Dish struct {
    Title string `json:"title"`
    Desc string `json:"desc"`
}

// let's declare a global array
// that we can then populate in our main function
// to simulate a database
var Dishes []Dish

func returnAllDishes(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllDishes")
    json.NewEncoder(w).Encode(Dishes)
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/dishes", returnAllDishes)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	Dishes = []Dish{
        Dish{Title: "Tacos", Desc: "Maistacos, Hackfleisch, Tomaten, Guacamole"},
        Dish{Title: "Tamales", Desc: "Auch lecker"},
	}
	
    handleRequests()
}