package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jesmendi/quizCapital/quiz"
)

var countries = quiz.LoadData()

func main() {

	handleRequests()

}
func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/questions", getQuestions)
	myRouter.HandleFunc("/answer/{capital}", handleResponse)

	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
