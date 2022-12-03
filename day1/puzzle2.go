package day1

import (
	"bufio"
	"container/heap"
	"log"
	"os"
	"strconv"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func Puzzle2() *IntHeap {
	file, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	highestCalories := &IntHeap{0, 0, 0}
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
			heap.Push(highestCalories, currentCalories)
			heap.Pop(highestCalories)
			currentCalories = 0
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return highestCalories
}
