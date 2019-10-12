package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func getQuestions(problemsFile string) ([][]string, error) {
	csvFile, err := os.Open(problemsFile)
	if err != nil {
		fmt.Println("Error opening the csv file for reading!")
		return nil, err
	}

	csvReader := csv.NewReader(csvFile)
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("Error parsing the provided csv file")
		return nil, err
	}

	return records, nil
}

func main() {
	problemsFile := flag.String("problems", "problems.csv", "a csv file with the quize problems")
	timerLength := flag.Int("timer", 30, "amount of time to take the quiz in seconds")
	flag.Parse()

	fmt.Print("Press enter to start the quiz!")
	fmt.Scanln()

	quizTimer := time.NewTimer(time.Duration(*timerLength) * time.Second)

	records, err := getQuestions(*problemsFile)
	if err != nil {
		fmt.Println("Error encountered while getting the questions!")
	}

	questionsCorrect := 0

	inputReader := bufio.NewReader(os.Stdin)
	answersChannel := make(chan string)

questionsLoop:
	for _, question := range records {
		fmt.Printf("What is %s? ", question[0])
		go func() {
			answer, _ := inputReader.ReadString('\n')
			answersChannel <- answer
		}()

		select {
		case <-quizTimer.C:
			fmt.Println()
			break questionsLoop
		case answer := <-answersChannel:
			if strings.TrimRight(answer, "\n") == question[1] {
				questionsCorrect++
			}
		}
	}

	fmt.Printf("You have completed the quiz.  Results %d/%d correct, way to go!\n", questionsCorrect, len(records))
}
