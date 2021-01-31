package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mm0070/acPhotoQuiz/questions"
)

func main() {
	r := gin.Default()
	r.GET("/getQuestions/:questionCount", func(c *gin.Context) {
		count, err := strconv.Atoi(c.Param("questionCount"))
		if err != nil {
			fmt.Println("Error parsing string to int")
			// todo: return something else than 200
		}

		// TODO: limit count to sensible numbers
		q := questions.PrepareQuestions(count)

		c.JSON(200, gin.H{
			"questions": q,
		})
	})
	r.Run("0.0.0.0:4000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
