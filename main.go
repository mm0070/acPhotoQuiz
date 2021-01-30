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

type question struct {
	Manufacturer string
	Model        string
	PhotoURL     string
	RawData      string
}

type questionSet struct {
	question []question
}

func visit(photoID int) string {
	var model string
	c := colly.NewCollector(
		colly.AllowedDomains("jetphotos.com", "www.jetphotos.com"),
		colly.CacheDir("./colly_cache"),
	)

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnHTML("title", func(e *colly.HTMLElement) {
		if !strings.Contains(e.Text, "Latest aviation photos on JetPhotos") {
			// fmt.Println(e.Text)
			model = strings.Split(e.Text, "|")[1]
		}
	})

	c.OnHTML("img.large-photo__img", func(e *colly.HTMLElement) {
		fmt.Println(e)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	err := c.Visit("https://www.jetphotos.com/photo/" + strconv.Itoa(photoID))
	if err != nil {
		log.Println(err)
	}
	return model
}

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(visit(rand.Intn(9223337) + 1))
		time.Sleep(time.Second * 1)
	}

}
