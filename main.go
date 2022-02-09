package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("welocome to go")
	timeLimit := flag.Int("time", 10, "time limit for the game")
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of queston and answers")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("failed to open the csv file:%s\n", *csvFileName))
	}
	readLine := csv.NewReader(file)
	lines, err := readLine.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("failed to read the csv file:%s\n", *csvFileName))
	}
	Problems := lineParser(lines)
	correct := 0

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	answerChannel := make(chan string)

	for i, p := range Problems {
		fmt.Println("problem no ", i+1, ":", p.ques)
		var answer string
		go func() {
			fmt.Scanf("%s\n", &answer)
			answerChannel <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("Ooops!! you ran out of time\nyou got %v correct outs of %v\n", correct, len(Problems))
			return
		case answer := <-answerChannel:
			if answer != p.ans {
				fmt.Println("you entered wrong answer")
				break
			} else {
				fmt.Println("Correct")
				correct++
			}
		}

	}
	fmt.Printf("you got %v correct outs of %v\n", correct, len(Problems))

}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func lineParser(lines [][]string) []problem {
	prob := make([]problem, len(lines))
	for i, line := range lines {

		prob[i] = problem{
			ques: line[0],
			ans:  strings.TrimSpace(line[1]),
		}
	}
	return prob
}

type problem struct {
	ques string
	ans  string
}
