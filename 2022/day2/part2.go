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
	"X": "Lose",
	"Y": "Draw",
	"Z": "Win",
}

var beats = map[string]string{
	"Scissors": "Rock",
	"Rock":     "Paper",
	"Paper":    "Scissors",
}

var beatBy = map[string]string{
	"Rock":     "Scissors",
	"Paper":    "Rock",
	"Scissors": "Paper",
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
			approach := secondColMap[row[1]]

			if approach == "Draw" {
				score += scores[opp] + drawBonus
			} else if approach == "Lose" {
				myHand := beatBy[opp]
				score += scores[myHand] + loseBonus
			} else if approach == "Win" {
				myHand := beats[opp]
				score += scores[myHand] + winBonus
			}

		}

		fmt.Println(score)
	} else {
		log.Fatal("No input")
	}
}
