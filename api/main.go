package main

import (
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mm0070/acPhotoQuiz/questions"
)

func main() {
	r := gin.Default()
	r.GET("/getQuestions/:mode/:questionCount", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		count, err := strconv.Atoi(c.Param("questionCount"))
		if err != nil {
			log.Fatal("Error parsing string to int")
		}

		if count < 1 && count > 10 {
			log.Fatal("Wrong number of questions: ", count)
		}

		// dev mode to get sample data quickly without waiting
		if c.Param("mode") == "test" {
			time.Sleep(time.Millisecond * 500) // add some fake delay
			jsonData := []byte(`{"questions":[{"manufacturer":"Airbus","model":"A319","photo_url":"https://cdn.jetphotos.com/full/3/50089_1319225745.jpg","page_url":"https://www.jetphotos.com/photo/7230902","raw_data":"OK-MEL | Airbus A319-112 | CSA Czech Airlines | Alessandro Lukas | JetPhotos"}]}`)
			c.Data(200, "application/json", jsonData)
		}

		if c.Param("mode") == "fetch" {
			q := questions.PrepareQuestions(count)

			c.JSON(200, gin.H{
				"questions": q,
			})
		}

		// TODO: add a db and create 'local' mode

		log.Println("Fetching questions...", count)

	})
	r.Run()
}
