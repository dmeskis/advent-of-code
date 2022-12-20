package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// sorting

func main() {
	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	top3 := []int{0, 0, 0}

	if fi.Size() > 0 {
		s := bufio.NewScanner(os.Stdin)
		sum := 0
		for s.Scan() {
			if s.Text() == "" {
				for i, val := range top3 {
					if sum > val {
						top3[i] = sum
						break
					}
				}
				sum = 0
			} else {
				val, err := strconv.Atoi(s.Text())
				if err != nil {
					panic(err)
				}
				sum += val
			}
		}

		total := 0
		for _, val := range top3 {
			total += val
		}
		fmt.Println(total)
	} else {
		log.Fatal("No input")
	}
}
