package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if fi.Size() > 0 {
		s := bufio.NewScanner(os.Stdin)
		sum := 0
		total := 0
		for s.Scan() {
			if s.Text() == "" {
				if sum > total {
					total = sum
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
		fmt.Println(total)
	} else {
		fmt.Println("Nope")
	}
}
