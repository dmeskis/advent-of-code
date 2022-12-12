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

	cols := [][]int{}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		rowAsString := strings.Split(s.Text(), "")
		row := make([]int, len(rowAsString))
		for i, n := range rowAsString {
			val, _ := strconv.Atoi(n)
			row[i] = val
		}
		cols = append(cols, row)
	}

	visible := 0
	for y, row := range cols {
		for x, height := range row {
			// Always visible on the border
			if x == 0 || y == 0 || x == len(row)-1 || y == len(row)-1 {
				visible++
				continue
			}

			ny := y - 1
			for ; ny >= 0; ny-- {
				// otha tree height
				oth := cols[ny][x]
				if oth >= height {
					break
				}
			}

			if ny < 0 {
				visible++
				continue
			}

			ex := x + 1
			for ; ex < len(row); ex++ {
				oth := cols[y][ex]
				if oth >= height {
					break
				}
			}

			if ex > len(row)-1 {
				visible++
				continue
			}

			sy := y + 1
			for ; sy < len(cols); sy++ {
				oth := cols[sy][x]
				if oth >= height {
					break
				}
			}

			if sy > len(cols)-1 {
				visible++
				continue
			}

			wx := x - 1
			for ; wx >= 0; wx-- {
				oth := cols[y][wx]
				if oth >= height {
					break
				}
			}

			if wx < 0 {
				visible++
				continue
			}
		}
	}
	// read in the numbers into a slice of slices. Basically a giant grid. Can use indices as X, Y coordinates.
	// don't forget
	fmt.Println(visible)

	return
}
