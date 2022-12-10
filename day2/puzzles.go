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
)

type WinStatus int

const (
	Win WinStatus = iota
	Loss
	Draw
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

func getScoreForRound(codedRoundInput string, puzzleNumber int) int {
	var myMove Move
	var winStatus WinStatus

	firstPartOfInput, secondPartOfInput := getInputParts(codedRoundInput)
	opponentMove := getOpponentMoveFromCodedInput(firstPartOfInput)

	if puzzleNumber == 1 {
		myMove = getMyMoveFromCodedInput(secondPartOfInput)
		winStatus = determineWinStatus(myMove, opponentMove)
	} else {
		winStatus = getWinStatusFromCodedInput(secondPartOfInput)
		myMove = determineMyMove(opponentMove, winStatus)
	}

	return calculateRoundScore(myMove, winStatus)
}

func getInputParts(fullInput string) (firstPart, secondPart string) {
	bothParts := strings.Split(fullInput, " ")
	if len(bothParts) != 2 {
		log.Fatal("Incorrect input format")
	}

	firstPart = bothParts[0]
	secondPart = bothParts[1]
	return firstPart, secondPart
}

func getOpponentMoveFromCodedInput(codedInput string) Move {
	var opponentMove Move
	switch codedInput {
	case "A":
		opponentMove = Rock
	case "B":
		opponentMove = Paper
	case "C":
		opponentMove = Scissors
	default:
		log.Fatal("The opponent used an unknown gesture")
	}
	return opponentMove
}

func getMyMoveFromCodedInput(codedInput string) Move {
	var myMove Move
	switch codedInput {
	case "X":
		myMove = Rock
	case "Y":
		myMove = Paper
	case "Z":
		myMove = Scissors
	default:
		log.Fatal("I used an unknown gesture")
	}
	return myMove
}

func getWinStatusFromCodedInput(codedInput string) WinStatus {
	var winStatus WinStatus
	switch codedInput {
	case "X":
		winStatus = Loss
	case "Y":
		winStatus = Draw
	case "Z":
		winStatus = Win
	default:
		log.Fatal("An unknown win status was provided")
	}
	return winStatus
}

func determineWinStatus(myMove, opponentMove Move) WinStatus {
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
