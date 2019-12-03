package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func run(input []string, noun, verb int) int {
	input[1] = strconv.Itoa(noun)
	input[2] = strconv.Itoa(verb)

	for i := 0; i < len(input); {
		if input[i] == "99" {
			break // end
		}

		result := 0
		idx1, idx2, pos := toInt(input[i+1]), toInt(input[i+2]), toInt(input[i+3])

		switch input[i] {
		case "1":
			// add
			result = toInt(input[idx1]) + toInt(input[idx2])
		case "2":
			// multiply
			result = toInt(input[idx1]) * toInt(input[idx2])
		}

		input[pos] = strconv.Itoa(result)
		i += 4
	}
	return toInt(input[0])
}

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	program := string(b)

	fmt.Println(run(strings.Split(program, ","), 12, 2))

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if run(strings.Split(program, ","), i, j) == 19690720 {
				fmt.Printf("%02d%02d\n", i , j)
				break
			}
		}
	}
}

func toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
