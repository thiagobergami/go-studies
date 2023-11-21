/* https://github.com/gophercises/quiz */

package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type question struct {
	problem string
	result  int
}

type quiz struct {
	questions []question
	score     int
}

func ask(question string, result int) int {
	fmt.Println("What is the answer for ", question)
	var input int
	_, err := fmt.Scan(&input)

	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}
	if input == result {
		fmt.Println("CORRECT!")
		return 1
	}
	fmt.Println("WRONG!")
	return 0

}

func convertToArray(data [][]string) []question {
	var questionsList []question

	for _, line := range data {
		var quest question
		quest.problem = line[0]
		quest.result, _ = strconv.Atoi(line[1])

		questionsList = append(questionsList, quest)
	}

	return questionsList
}

func main() {
	f, err := os.Open("problem.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()

	var quiz quiz

	if err != nil {
		log.Fatal(err)
	}

	quiz.questions = convertToArray(data)
	for _, line := range quiz.questions {
		quiz.score += ask(line.problem, line.result)
	}

	fmt.Println("Total score: ", quiz.score)
}
