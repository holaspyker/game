package quiz

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/jesmendi/quizCapital/helper"
)

type Answer struct {
	answer string
}
type Result struct {
	Answer  string
	Country string
	Capital string
	OK      bool
}

type Classification struct {
	Result   []Result
	Position int
	Correct  int
}

var class []int

var Res []Result
var sum = 0

func FillResult(val string) {

	cor := strings.ToLower(val) == strings.ToLower(QuesData[len(Res)].Correct)
	if cor {
		sum = sum + 1
	}
	c := Result{
		Answer:  val,
		Country: QuesData[len(Res)].Country,
		Capital: QuesData[len(Res)].Correct,
		OK:      cor,
	}
	Res = append(Res, c)

}

func GetClasificacion(w http.ResponseWriter) {

	position := getPosition()
	class = append(class, sum)
	data :=
		Classification{
			Result:   Res,
			Position: position,
			Correct:  sum,
		}
	clearData()
	b, err := json.Marshal(data)
	helper.SendData(w, b, err)

}

func getPosition() int {

	if sum == 0 {
		return 0
	}

	if len(class) == 0 {
		return 100
	}
	var greater = 0

	for _, j := range class {
		if sum > j {
			fmt.Println("entra en  if ")
			fmt.Println(j)
			greater = greater + 1
		}
	}
	total := 100 * float32(greater) / float32(len(class))
	return int(total)
}

func clearData() {
	Res = []Result{}
	sum = 0

}
