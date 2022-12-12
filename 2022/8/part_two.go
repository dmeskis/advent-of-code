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

	topScore := 0
	for y, row := range cols {
		for x, height := range row {
			// Skip border
			if x == 0 || y == 0 || x == len(row)-1 || y == len(row)-1 {
				continue
			}

			ny := y - 1
			nScore := 0
			for ; ny >= 0; ny-- {
				nScore++
				oth := cols[ny][x]
				if oth >= height {
					break
				}
			}

			ex := x + 1
			eScore := 0
			for ; ex < len(row); ex++ {
				eScore++
				oth := cols[y][ex]
				if oth >= height {
					break
				}
			}

			sy := y + 1
			sScore := 0
			for ; sy < len(cols); sy++ {
				sScore++
				oth := cols[sy][x]
				if oth >= height {
					break
				}
			}

			wx := x - 1
			wScore := 0
			for ; wx >= 0; wx-- {
				wScore++
				oth := cols[y][wx]
				if oth >= height {
					break
				}
			}

			score := nScore * eScore * sScore * wScore
			if score > topScore {
				topScore = score
			}
		}
	}
	fmt.Println(topScore)

	return
}
