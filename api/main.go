package main

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mm0070/acPhotoQuiz/questions"
)

func main() {
	r := gin.Default()
	r.GET("/getQuestions/:questionCount", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		count, err := strconv.Atoi(c.Param("questionCount"))
		if err != nil {
			log.Fatal("Error parsing string to int")
		}

		if count < 1 && count > 10 {
			log.Fatal("Wrong number of questions: ", count)
		}

		q := questions.PrepareQuestions(count)

		log.Println("Fetching %i questions...", count)

		c.JSON(200, gin.H{
			"questions": q,
		})
	})
	r.Run()
}
