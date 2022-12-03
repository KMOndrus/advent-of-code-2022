package day1

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func Puzzle1() int {
	file, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	highestCalories := 0
	currentCalories := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			servingCalories, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
			currentCalories += servingCalories
		} else {
			if currentCalories > highestCalories {
				highestCalories = currentCalories
			}
			currentCalories = 0
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return highestCalories
}
