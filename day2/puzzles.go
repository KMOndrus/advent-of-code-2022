package day2

import (
	"bufio"
	"log"
	"os"
	"strings"
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
	file, err := os.Open("day2/input.txt")
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

func determineMyMove(opponentMove Move, status WinStatus) Move {
	if status == Win {
		switch opponentMove {
		case Rock:
			return Paper
		case Paper:
			return Scissors
		case Scissors:
			return Rock
		}
	}

	if status == Loss {
		switch opponentMove {
		case Rock:
			return Scissors
		case Paper:
			return Rock
		case Scissors:
			return Paper
		}
	}

	return opponentMove
}

func getRoundDetails(moveString string) (opponentMove Move, winStatus WinStatus) {
	roundDetails := strings.Split(moveString, " ")
	if len(roundDetails) != 2 {
		log.Fatal("Not the right number of details for the round")
	}

	switch roundDetails[0] {
	case "A":
		opponentMove = Rock
	case "B":
		opponentMove = Paper
	case "C":
		opponentMove = Scissors
	default:
		opponentMove = NoMove
	}

	switch roundDetails[1] {
	case "X":
		winStatus = Loss
	case "Y":
		winStatus = Draw
	case "Z":
		winStatus = Win
	default:
		winStatus = NoStatus
	}

	if winStatus == NoStatus || opponentMove == NoMove {
		log.Fatal("Some of the round details are invalid")
	}
	return opponentMove, winStatus
}

func getScoreForRound(moveString string, puzzleNumber int) int {
	var opponentMove, myMove Move
	var winStatus WinStatus

	if puzzleNumber == 1 {
		opponentMove, myMove = getMovesForRound(moveString)
		winStatus = getWinStatus(myMove, opponentMove)
	} else {
		opponentMove, winStatus = getRoundDetails(moveString)
		myMove = determineMyMove(opponentMove, winStatus)
	}

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
