package main

import (
	"advent-of-code-2022/day1"
	"advent-of-code-2022/day2"
	"fmt"
)

func main() {
	day1Puzzle1 := day1.Puzzle1()
	fmt.Printf("Day 1 Puzzle 1 Result: %d\n", day1Puzzle1)

	day1Puzzle2 := day1.Puzzle2()
	fmt.Printf("Day 1 Puzzle 2 Result: %v\n", day1Puzzle2)

	day2Puzzle1 := day2.Puzzle1()
	fmt.Printf("Day 2 Puzzle 1 Result: %d\n", day2Puzzle1)

	day2Puzzle2 := day2.Puzzle2()
	fmt.Printf("Day 2 Puzzle 2 Result: %d\n", day2Puzzle2)
}
