package day2

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Puzzle1() int {
	file, err := os.Open("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	finalScore := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		roundScore := getScoreForRound(scanner.Text())
		finalScore += roundScore
	}
	return finalScore
}

func getWinStatus(myMove, opponentMove Move) WinStatus {
	if myMove == opponentMove {
		return Draw
	}

	if myMove == Rock && opponentMove == Scissors {
		return Win
	}

	if myMove == Paper && opponentMove == Rock {
		return Win
	}

	if myMove == Scissors && opponentMove == Paper {
		return Win
	}

	return Loss
}

func getMovesForRound(moveString string) (opponentMove, myMove Move) {
	bothMoves := strings.Split(moveString, " ")
	if len(bothMoves) != 2 {
		log.Fatal("Not the right number of players for the round")
	}

	switch bothMoves[0] {
	case "A":
		opponentMove = Rock
	case "B":
		opponentMove = Paper
	case "C":
		opponentMove = Scissors
	default:
		opponentMove = NoMove
	}

	switch bothMoves[1] {
	case "X":
		myMove = Rock
	case "Y":
		myMove = Paper
	case "Z":
		myMove = Scissors
	default:
		myMove = NoMove
	}

	if myMove == NoMove || opponentMove == NoMove {
		log.Fatal("Somebody used an unknown gesture!")
	}
	return opponentMove, myMove
}
