package questions

import (
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type questionData struct {
	AircraftType string
	PhotoURL     string
	PageURL      string
	TitleRawData string
}

func getPhotoInfo(photoID int) questionData {
	var q questionData

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
			// no hit on random ID, redirected to /new
			return
		}
		q.TitleRawData = e.Text
		q.AircraftType = strings.TrimSpace(strings.Split(e.Text, "|")[1])
	})

	// get image URL
	c.OnHTML("img.large-photo__img", func(e *colly.HTMLElement) {
		if q.PhotoURL == "" {
			q.PhotoURL = e.Attr("src")
		}
	})

	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting", r.URL)
	// })

	q.PageURL = "https://www.jetphotos.com/photo/" + strconv.Itoa(photoID)
	err := c.Visit(q.PageURL)
	if err != nil {
		log.Println(err)
	}

	return q
}
