package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

type questionData struct {
	AircraftType string
	PhotoURL     string
	TitleRawData string
}

func visit(photoID int) questionData {
	var question questionData

	// instantiate collector
	c := colly.NewCollector(
		colly.AllowedDomains("jetphotos.com", "www.jetphotos.com"),
		colly.CacheDir("./colly_cache"),
	)

	// handle errors
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	// get model from <title>
	c.OnHTML("title", func(e *colly.HTMLElement) {
		if strings.Contains(e.Text, "Latest aviation photos on JetPhotos") {
			return
		}
		question.TitleRawData = e.Text
		question.AircraftType = strings.TrimSpace(strings.Split(e.Text, "|")[1])
	})

	// get image URL
	c.OnHTML("img.large-photo__img", func(e *colly.HTMLElement) {
		if question.PhotoURL == "" {
			question.PhotoURL = e.Attr("src")
		}
	})

	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting", r.URL)
	// })

	err := c.Visit("https://www.jetphotos.com/photo/" + strconv.Itoa(photoID))
	if err != nil {
		log.Println(err)
	}

	return question
}

func main() {
	numberOfQuestions := 10
	for i := 0; i < numberOfQuestions; {

		result := visit(rand.Intn(9223337) + 1)
		fmt.Println("\n\nNew question: ")
		fmt.Println(result.AircraftType)
		fmt.Println(result.PhotoURL)
		fmt.Println(result.TitleRawData)
		fmt.Println("result: ", matchAircraftTypes(&result))
		time.Sleep(time.Second * 1)
	}

}
