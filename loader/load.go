package loader

import (
	//"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
)

type CountryData struct {
	CountryID string `json:"country_id"`
	Country   string `json:"country"`
	Capital   string `json:"capital"`
}

type QuestionData struct {
	Country string   `json:"country"`
	Capital []string `json:"capital"`
}

type CountryStore interface {
	LoadData()
}

type Country struct {
	Store []CountryData `json:"store"`
}

func LoadData() []CountryData {
	csvfile, err := os.Open("file/countries.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	ret := make([]CountryData, 0, 0)
	// Parse the file
	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		count := CountryData{
			CountryID: record[0],
			Country:   record[1],
			Capital:   record[2],
		}

		ret = append(ret, count)
	}

	return ret

}

func GetQuestions([]CountryData) []QuestionData {

	ret := make([]QuestionData, 0, 0)

	return ret
}
