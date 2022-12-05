package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var firstColMap = map[string]string{
	"A": "Rock",
	"B": "Paper",
	"C": "Scissors",
}

var secondColMap = map[string]string{
	"X": "Rock",
	"Y": "Paper",
	"Z": "Scissors",
}

var scores = map[string]int{
	"Rock":     1,
	"Paper":    2,
	"Scissors": 3,
}

func main() {
	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if fi.Size() > 0 {
		s := bufio.NewScanner(os.Stdin)
		loseBonus := 0
		drawBonus := 3
		winBonus := 6
		score := 0
		for s.Scan() {
			row := strings.Fields(s.Text())

			opp := firstColMap[row[0]]
			me := secondColMap[row[1]]

			fmt.Printf("Opp played %s\n", opp)
			fmt.Printf("You played %s\n", me)

			if opp == me {
				score += scores[me] + drawBonus
			} else if opp == "Rock" {
				if me == "Paper" {
					score += scores[me] + winBonus
				} else {
					score += scores[me] + loseBonus
				}
			} else if opp == "Paper" {
				if me == "Scissors" {
					score += scores[me] + winBonus
				} else {
					score += scores[me] + loseBonus
				}
			} else if opp == "Scissors" {
				if me == "Rock" {
					score += scores[me] + winBonus
				} else {
					score += scores[me] + loseBonus
				}

			}

		}

		fmt.Println(score)
	} else {
		log.Fatal("No input")
	}
}
