package main

import (
	"fmt"
	"os"
	"flag"
	"encoding/csv"
	"strings"
	"time"
)

var csvFile string
var timeLimit int

func init(){
	flag.StringVar(&csvFile, "csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.IntVar(&timeLimit, "limit", 30, "The time limit for the quiz in seconds")
}

func main(){
	flag.Parse()

	file, err := os.Open(csvFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)

	var correct int
	for i, line := range lines {
		fmt.Printf("Problem #%d: %s = \n", i+1, line[0])
		answerCh := make(chan string)

		go func(){
			var getAnswer string
			fmt.Scanf("%s\n", &getAnswer)
			answerCh <- getAnswer
		}()

		select {
		case <- timer.C:
			fmt.Printf("\nyou scored %d out of %d\n", correct, len(lines))
			return
		case getAnswer := <- answerCh:
			if strings.TrimSpace(getAnswer) == line[1] {
				correct++
			}
		}
	}
	
	fmt.Printf("you scored %d out of %d\n", correct, len(lines))
}