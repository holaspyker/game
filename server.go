package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jesmendi/quizCapital/quiz"
)

type Answer struct {
	answer string
}

var res []Answer

func getQuestions(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	data := quiz.GetQuestions(countries)

	if data != nil {
		b, err := json.Marshal(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "error marshalling data"}`))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(b)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "not found"}`))
}

func handleResponse(w http.ResponseWriter, r *http.Request) {

	pathParams := mux.Vars(r)

	if val, ok := pathParams["answer"]; ok {

		c := Answer{
			answer: val,
		}
		res = append(res, c)

		fmt.Println(res)

	}
	w.WriteHeader(http.StatusOK)
	return

	w.WriteHeader(http.StatusNotFound)

}
