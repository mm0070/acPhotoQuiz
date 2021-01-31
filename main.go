package main

import (
	"fmt"

	"github.com/mm0070/acPhotoQuiz/questions"
)

func main() {
	q := questions.PrepareQuestions()
	fmt.Println(q)
}
