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
		for i := 0; i+4 < len(s.Text()); i++ {
			chars := s.Text()[i : i+4]
			myMap := make(map[rune]bool)
			for _, char := range chars {
				myMap[char] = true

			}

			// You found it
			if len(myMap) == 4 {
				fmt.Println(i + 4)
				return

			}
		}
	}

}
