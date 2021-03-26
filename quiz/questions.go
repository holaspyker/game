package quiz

import (
	"math/rand"
	"time"

	"github.com/jesmendi/quizCapital/helper"
)

type QuestionData struct {
	Country string   `json:"country"`
	Capital []string `json:"capital"`
}

func GetQuestions(cty []CountryData) []QuestionData {

	random := helper.GetRandomSlice(12, len(cty))
	split := helper.ArrayChunk(random, 3)

	ret := make([]QuestionData, 0)
	for _, s := range split {
		rand.Seed(time.Now().UnixNano())
		t := rand.Perm(3)
		q := QuestionData{
			Country: cty[s[0]].Country,
			Capital: []string{cty[s[t[0]]].Capital, cty[s[t[1]]].Capital, cty[s[t[2]]].Capital},
		}
		ret = append(ret, q)
	}
	return ret
}
