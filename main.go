package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jesmendi/quizCapital/quiz"
)

var countries = quiz.LoadData()

func main() {
	rand.Seed(time.Now().Unix())

	questions := quiz.GetQuestions(countries)
	fmt.Println(questions)

}
