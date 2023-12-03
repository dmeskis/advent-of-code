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

	check := 20 // which cycle to check
	sum := 0    // sum of signal strenghts
	// cycle := 1 // start on cycle 1
	X := 1 // register X starts at 1

	instructions := [][]string{}

	cycles := 0 // execution counter, basically
	// isExecuting := false
	stack := []int{}

	for s.Scan() {
		fields := strings.Fields(s.Text())
		instructions = append(instructions, fields)
	}

	// i is a cycle i + 1 is da cycle
	// + 2 give sthe shit a chance to finish
	for i := 0; i < len(instructions)+2; i++ {

		if len(stack) > 0 {
			if cycles > 1 {
				asdf := len(stack) - 1
				// add the shit
				e := stack[len(stack)-1]
				s = slices.Delete(stack, i, i+1)
				// X +=

			}

		}
		// if isExecuting {
		//     // if cycle

		// }
		// check if

		fmt.Println("--- Begin cycle ", i+1, " ---")
		// val, ok := stack[i]
		// if ok {
		// 	fmt.Println("Adding ", val)
		// 	X += val
		// }

		if i+1 == check {
			fmt.Println("X: ", X)
			fmt.Println("strength: ", X*(i+1))
			sum += X * (i + 1)
			check += 40
			fmt.Println("sum: ", sum)
		}

		if i >= len(instructions) {
			continue
		}

		cmd := instructions[i][0]

		if cmd == "addx" {
			n, _ := strconv.Atoi(instructions[i][1])
			stack = append(stack, n)
		}

		fmt.Println("X: ", X)
		fmt.Println("--- End cycle ", i+1, " ---")
	}

	// now we have the shit

	fmt.Println(sum)
}
