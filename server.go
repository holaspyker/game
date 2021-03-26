package main

import (
	"encoding/json"
	"net/http"

	"github.com/jesmendi/quizCapital/quiz"
)

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
