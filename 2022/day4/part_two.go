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
		set1 := []int{}
		set2 := []int{}
		for _, num := range x {
			val, _ := strconv.Atoi(num)
			set1 = append(set1, val)
		}

		for _, num := range y {
			val, _ := strconv.Atoi(num)
			set2 = append(set2, val)
		}

		if set1[0] >= set2[0] && set1[0] <= set2[1] {
			sum += 1
		} else if set1[1] >= set2[0] && set1[1] <= set2[1] {
			sum += 1
		} else if set2[0] >= set1[0] && set2[0] <= set1[1] {
			sum += 1
		} else if set2[1] >= set1[0] && set2[1] <= set1[1] {
			sum += 1
		}
	}

	fmt.Println(sum)
}
