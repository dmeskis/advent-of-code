package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fi, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err)
	}

	if !(fi.Size() > 0) {
		log.Fatal("No input")
		return
	}

	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		for i := 0; i+14 < len(s.Text()); i++ {
			chars := s.Text()[i : i+14]
			myMap := make(map[rune]bool)
			for _, char := range chars {
				myMap[char] = true

			}

			// You found it
			if len(myMap) == 14 {
				fmt.Println(i + 14)
				return

			}
		}
	}

}
