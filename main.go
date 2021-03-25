package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jesmendi/quizCapital/loader"
)

var countries = loader.LoadData()

func main() {
	rand.Seed(time.Now().Unix())

	q := rand.Perm(len(countries))
	q = q[:12]
	fmt.Println(q)

}
