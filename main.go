package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector(
		// colly.AllowedDomains("www.jetphotos.net"),
		colly.CacheDir("./colly_cache"),
	)

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnHTML(".list__item", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	err := c.Visit("https://www.jetphotos.com/photo/10035336")
	if err != nil {
		log.Println(err)
	}
}

// https://www.jetphotos.com/photo/10035336
