package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func fuel(mass int) int {
	f := mass/3 - 2
	if f <= 0 {
		return 0
	}
	return f
}

func fuelSum(mass int) int {
	f := fuel(mass)
	sum := 0

	for f > 0 {
		sum += f
		f = fuel(f)
	}

	return sum
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(file)

	fuelCounter := 0
	sumOfFuel := 0
	for scanner.Scan() {
		mass, _ := strconv.Atoi(scanner.Text())
		sumOfFuel += fuelSum(mass)
		fuelCounter += fuel(mass)
	}

	fmt.Println("fuel count:", fuelCounter)
	fmt.Println("sum of fuel:", sumOfFuel)
}
