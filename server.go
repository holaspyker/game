package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jesmendi/quizCapital/quiz"
)

func getQuestions(w http.ResponseWriter, r *http.Request) {
	quiz.GetNextQuestion(w)
}

func handleResponse(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	if val, ok := pathParams["capital"]; ok {
		quiz.FillResult(val)
		if len(quiz.Res) == 4 {
			quiz.GetClasificacion(w)
		} else {
			quiz.GetNextQuestion(w)
		}
	}
}
