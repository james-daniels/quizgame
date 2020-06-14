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

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	mlines := make(map[string]string)
	for _, v := range lines {
		mlines[v[0]] = v[1]
	}

timer := time.NewTimer(time.Duration(timeLimit) * time.Second)

var correct, i int
	for key, value := range mlines {
		fmt.Printf("Problem #%d: %s = \n", i+1, key)

		answerCh := make(chan string)
		go func(){
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <- timer.C:
			fmt.Printf("\nyou scored %d out of %d\n", correct, len(lines))
			return
		case answer := <- answerCh:
			if strings.TrimSpace(answer) == value {
				correct++
			}
		}
		i++
	}
	fmt.Printf("you scored %d out of %d\n", correct, len(lines))
}