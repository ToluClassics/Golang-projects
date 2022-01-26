package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func main() {
	// get the filename from the terminal; flag name is "csv",
	// default file is "problems.csv"
	var csvFilename string
	flag.StringVar(&csvFilename, "csv", "problems.csv", "a csv file with questions and answer in two columns")
	flag.Parse()

	fmt.Println(reflect.TypeOf(csvFilename))

	all_lines := read_csv(csvFilename)

	problems := parseLines(all_lines)

	correct := 0
	// loop through all the problems
	for i, p := range problems {
		fmt.Printf("Problem Number %d=%s \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)

		if answer == p.a {
			correct += 1
			fmt.Printf("Correct!")
		}
	}

	fmt.Printf("You Scored %d out of %d", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	//declare variable to return which is a list of struc problem
	// and length of lines
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}

	return ret
}

// create a struct for the incoming file format
type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Printf(msg)
	os.Exit(1)
}

func read_csv(fileName string) [][]string {
	// read csv file using the retrieved filename
	file, err := os.Open(fileName)

	//print an error message that the file doesnt open
	if err != nil {
		msg := fmt.Sprintf("We could not open file:: %s \n", fileName)
		exit(msg)
	} else {
		fmt.Printf("Successfully opened file:: %s\n", fileName)
	}

	// read the csv file
	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		exit("Failed to parse CSV file")
	}

	return lines
}
