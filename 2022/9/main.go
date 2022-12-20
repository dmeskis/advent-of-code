package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fi, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err)
	}

	if !(fi.Size() > 0) {
		log.Fatal("No input")
	}

	s := bufio.NewScanner(os.Stdin)

	// don't worry about making this dynamic for now..
	grid := [1000][1000]bool{}

	// head start coordinates
	x := len(grid) / 2
	y := len(grid) / 2

	// tail start coordinates
	a := len(grid) / 2
	b := len(grid) / 2

	if x != y {
		log.Fatal("you fucked up")
	}

	for s.Scan() {
		fields := strings.Fields(s.Text())
		dir := fields[0]
		count, _ := strconv.Atoi(fields[1])

		for i := 0; i < count; i++ {
			switch dir {
			case "R":
				x++

				if x == a {
					break
				}

				if a < x-1 {
					b = y
					a = x - 1
				}
			case "L":
				x--

				if x == a {
					break
				}

				if a > x+1 {
					b = y
					a = x + 1
				}
			case "U":
				y--

				if y == b {
					break
				}

				if b > y+1 {
					a = x
					b = y + 1
				}
			case "D":
				y++

				if y == b {
					break
				}

				if b < y-1 {
					a = x
					b = y - 1
				}
			}

			grid[b][a] = true
		}
	}

	visited := 0
	for _, col := range grid {
		for _, val := range col {
			if val == true {
				visited++

			}

		}

	}

	fmt.Println(visited)

}
