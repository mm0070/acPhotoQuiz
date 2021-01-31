package main

import (
	"log"
	"math/rand"
	"strings"
	"time"
)

func prepareQuestions() []quizQuestion {
	var returnedQuestions int
	var questionArray []quizQuestion
	var question quizQuestion
	var data questionData
	var match bool
	numberOfQuestions := 10

	for returnedQuestions < numberOfQuestions {
		data = getPhotoInfo(rand.Intn(9223337) + 1)
		question, match = matchAircraftTypes(&data)
		if match == true {
			questionArray = append(questionArray, question)
			returnedQuestions++
		}
		// don't ddos jetphotos
		time.Sleep(time.Second)
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

func main() {
	prepareQuestions()
}
