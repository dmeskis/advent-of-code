package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// [a-z] = 1-26
// [A-Z] = 27-52
// :set nrformats+=alpha
var points = map[rune]int{
	'a': 1,
	'b': 2,
	'c': 3,
	'd': 4,
	'e': 5,
	'f': 6,
	'g': 7,
	'h': 8,
	'i': 9,
	'j': 10,
	'k': 11,
	'l': 12,
	'm': 13,
	'n': 14,
	'o': 15,
	'p': 16,
	'q': 17,
	'r': 18,
	's': 19,
	't': 20,
	'u': 21,
	'v': 22,
	'w': 23,
	'x': 24,
	'y': 25,
	'z': 26,
	'A': 27,
	'B': 28,
	'C': 29,
	'D': 30,
	'E': 31,
	'F': 32,
	'G': 33,
	'H': 34,
	'I': 35,
	'J': 36,
	'K': 37,
	'L': 38,
	'M': 39,
	'N': 40,
	'O': 41,
	'P': 42,
	'Q': 43,
	'R': 44,
	'S': 45,
	'T': 46,
	'U': 47,
	'V': 48,
	'W': 49,
	'X': 50,
	'Y': 51,
	'Z': 52,
}

func main() {
	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if !(fi.Size() > 0) {
		log.Fatal("No input")
		return
	}

	s := bufio.NewScanner(os.Stdin)
	sum := 0

	lines := []string{}

	for s.Scan() {
		lines = append(lines, s.Text())
		length := len(lines)
		if length > 2 {
			sets := []map[rune]bool{}

			for i := 0; i < length; i++ {
				sets = append(sets, make(map[rune]bool))
			}

			for i, line := range lines {
				for _, rune := range line {
					sets[i][rune] = true
				}

			}

			for k := range sets[0] {
				fmt.Println(k)
				if sets[0][k] && sets[1][k] && sets[2][k] == true {
					sum += points[k]
					break
				}

			}

			lines = []string{}
		}
	}

	fmt.Println(sum)

	return
}
