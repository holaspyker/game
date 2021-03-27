package quiz

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/jesmendi/quizCapital/config"
	"github.com/jesmendi/quizCapital/helper"
)

var Countries []CountryData

var QuesData []QuestionData

type QuestionData struct {
	Country string   `json:"country"`
	Correct string   `json:"correct"`
	Capital []string `json:"capital"`
}

func init() {

	Countries = LoadData()
	GetQuestions(Countries)
}

func GetQuestions(cty []CountryData) {

	random := helper.GetRandomSlice(config.Cfg.NumberAnswer, len(cty))
	split := helper.ArrayChunk(random, config.Cfg.AnswerPerQuestion)

	QuesData = make([]QuestionData, 0)
	for _, seeds := range split {
		rand.Seed(time.Now().UnixNano())
		permutations := rand.Perm(config.Cfg.AnswerPerQuestion)
		var answerCapitals = make([]string, 0)

		for _, p := range permutations {
			var answerCapital = cty[seeds[p]].Capital
			answerCapitals = append(answerCapitals, answerCapital)
		}

		q := QuestionData{
			Country: cty[seeds[0]].Country,
			Correct: cty[seeds[0]].Capital,
			Capital: answerCapitals,
		}
		QuesData = append(QuesData, q)
	}
}

func GetNextQuestion(w http.ResponseWriter) {
	data := QuesData[len(Res)]
	data.Correct = "xxxx"
	b, err := json.Marshal(data)
	helper.SendData(w, b, err)
}
