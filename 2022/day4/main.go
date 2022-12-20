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
		return
	}

	s := bufio.NewScanner(os.Stdin)
	sum := 0

	for s.Scan() {
		ranges := strings.Split(s.Text(), ",")
		x := strings.Split(ranges[0], "-")
		y := strings.Split(ranges[1], "-")
		xint := []int{}
		yint := []int{}
		for _, num := range x {
			val, _ := strconv.Atoi(num)
			xint = append(xint, val)
		}

		for _, num := range y {
			val, _ := strconv.Atoi(num)
			yint = append(yint, val)
		}

		if xint[0] >= yint[0] && xint[1] <= yint[1] {
			sum += 1
		} else if yint[0] >= xint[0] && yint[1] <= xint[1] {
			sum += 1
		}
	}

	fmt.Println(sum)
}
