package day2

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Puzzle2() int {
	file, err := os.Open("day2/sampleInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	finalScore := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		roundScore := getScoreForRound2(scanner.Text())
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
