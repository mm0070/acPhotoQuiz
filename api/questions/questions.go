package questions

import (
	"log"
	"math/rand"
	"strings"
	"time"
)

func PrepareQuestions(numberOfQuestions int) []quizQuestion {
	var returnedQuestions int
	var questionArray []quizQuestion
	var question quizQuestion
	var data questionData
	var match bool

	for returnedQuestions < numberOfQuestions {
		rand.Seed(time.Now().UTC().UnixNano())
		data = getPhotoInfo(rand.Intn(9223337) + 1)
		question, match = matchAircraftTypes(&data)
		if match == true {
			questionArray = append(questionArray, question)
			returnedQuestions++
		}
		// don't ddos jetphotos
		// TODO: this can be a parameter
		time.Sleep(time.Millisecond * 200)
	}

	// debug lines
	for _, v := range questionArray {
		log.Println("")
		log.Println("    Found manufacturer: ", v.Manufacturer)
		log.Println("    Found model: ", v.Model)
		log.Println("    PhotoURL: ", v.PhotoURL)
		log.Println("    Jetphotos thinks this is: ", strings.Split(v.TitleRawData, "|")[1])
	}

	return questionArray
}
