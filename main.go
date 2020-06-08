package main

import (
	"fmt"
	"os"
	"flag"
	"encoding/csv"
	"strings"
)

var csvFile string

	type problem struct {
		question string
		answer string
	}

func init(){
	flag.StringVar(&csvFile, "csv", "problems.csv", "a csv file in the format of 'question,answer'")
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

	problems := parseLines(lines)

	var correct int

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.answer {
			correct++
		}
	}

	fmt.Printf("you scored %d out of %d\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer: strings.TrimSpace(line[1]),
		}
	}

	return ret
}

// func main(){
	// flag.Parse()

	// file, err := os.Open(csvFile)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// defer file.Close()

	// reader := csv.NewReader(file)

	// lines, err := reader.ReadAll()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// var correct int
	// for i, line := range lines {
	// 	fmt.Printf("Problem #%d: %s = \n", i+1, line[0])

	// 	var getAnswer string
	// 	fmt.Scanf("%s\n", &getAnswer)
	// 	if strings.TrimSpace(getAnswer) == line[1] {
	// 		correct++
	// 	}
	// }
	
	// fmt.Printf("you scored %d out of %d\n", correct, len(lines))

// }