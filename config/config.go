package config

type Config struct {
	File              string
	NumberAnswer      int
	AnswerPerQuestion int
}

var Cfg Config

func init() {

	Cfg = Config{
		File:         "file/countries.csv",
		NumberAnswer: 12,
	}

}
