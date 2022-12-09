package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//     [H]         [D]     [P]
// [W] [B]         [C] [Z] [D]
// [T] [J]     [T] [J] [D] [J]
// [H] [Z]     [H] [H] [W] [S]     [M]
// [P] [F] [R] [P] [Z] [F] [W]     [F]
// [J] [V] [T] [N] [F] [G] [Z] [S] [S]
// [C] [R] [P] [S] [V] [M] [V] [D] [Z]
// [F] [G] [H] [Z] [N] [P] [M] [N] [D]
//  1   2   3   4   5   6   7   8   9

var stacks = [9][]rune{
	{'F', 'C', 'J', 'P', 'H', 'T', 'W'}, // 1
	{'G', 'R', 'V', 'F', 'Z', 'J', 'B', 'H'},
	{'H', 'P', 'T', 'R'},
	{'Z', 'S', 'N', 'P', 'H', 'T'},
	{'N', 'V', 'F', 'Z', 'H', 'J', 'C', 'D'},
	{'P', 'M', 'G', 'F', 'W', 'D', 'Z'},
	{'M', 'V', 'Z', 'W', 'S', 'J', 'D', 'P'},
	{'N', 'D', 'S'},
	{'D', 'Z', 'S', 'F', 'M'},
}

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
		// disregard line if it doesn't start w/ move
		if !strings.HasPrefix(s.Text(), "move") {
			continue
		}

		// split on space, use 1, 3, 5
		tmpCft := strings.Split(s.Text(), " ")
		var cft [3]int
		for i, idx := range []int{1, 3, 5} {
			val, err := strconv.Atoi(tmpCft[idx])
			if err != nil {
				log.Fatal(err)
			}

			cft[i] = val
		}

		// we have our count, our from, and our To
		// just do the moving...
		// grab the slice you'll be taking
		// then assign the opposite of that slice to the stack
		from := cft[1] - 1
		count := cft[0]
		toAppend := stacks[from][len(stacks[from])-count:]

		// update stack
		stacks[from] = stacks[from][:len(stacks[from])-count]

		to := cft[2] - 1
		// append to new stacks, it's LIFO so reverse before appending
		appLen := len(toAppend)
		reversed := make([]rune, appLen)
		for i, n := range toAppend {
			j := appLen - i - 1
			reversed[j] = n
		}
		stacks[to] = append(stacks[to], reversed...)

	}

	final := ""
	for _, stack := range stacks {
		length := len(stack)
		final = fmt.Sprintf("%s%c", final, stack[length-1])
	}
	fmt.Println(final)
}
