package day2

import (
	"bufio"
	"log"
	"os"
)

type Move int

const (
	Rock Move = iota
	Paper
	Scissors
	NoMove
)

type WinStatus int

const (
	Win WinStatus = iota
	Loss
	Draw
	NoStatus
)

func Puzzles(puzzleNumber int) int {
	file, err := os.Open("day2/sampleInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	finalScore := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		roundScore := getScoreForRound(scanner.Text(), puzzleNumber)
		finalScore += roundScore
	}
	return finalScore
}

func getScoreForRound(moveString string, puzzleNumber int) int {
	opponentMove, myMove := getMovesForRound(moveString)
	winStatus := getWinStatus(myMove, opponentMove)
	return calculateRoundScore(myMove, winStatus)
}

func getScoreForRound2(moveString string) int {
	opponentMove, winStatus := getRoundDetails(moveString)
	myMove := determineMyMove(opponentMove, winStatus)
	return calculateRoundScore(myMove, winStatus)
}

func calculateRoundScore(myMove Move, status WinStatus) int {
	roundScore := 0

	switch myMove {
	case Rock:
		roundScore += 1
	case Paper:
		roundScore += 2
	case Scissors:
		roundScore += 3
	}

	switch status {
	case Win:
		roundScore += 6
	case Draw:
		roundScore += 3
	}

	return roundScore
}
