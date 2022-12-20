package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type ropeNode struct {
	x    int
	y    int
	tail *ropeNode
}

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
	grid := [75][75]bool{}

	headNode := ropeNode{x: len(grid) / 2, y: len(grid) / 2}
	tailNode := &headNode

	for i := 0; i < 9; i++ {
		node := ropeNode{x: len(grid) / 2, y: len(grid) / 2}
		tailNode.tail = &node
		tailNode = tailNode.tail
	}

	for s.Scan() {
		fields := strings.Fields(s.Text())
		dir := fields[0]
		count, _ := strconv.Atoi(fields[1])

		for i := 0; i < count; i++ {
			moveHead(&headNode, dir)

			node := headNode
			for node.tail != nil {
				fmt.Println(node.x, node.y)
				node = *node.tail
			}
			grid[tailNode.y][tailNode.x] = true

			fmt.Println(dir, i)

			for _, col := range grid {
				line := []string{}
				for _, val := range col {
					if val == true {
						line = append(line, "x")
					} else {
						line = append(line, ".")

					}

				}
				fmt.Println(line)
			}
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

func moveHead(node *ropeNode, dir string) {
	if node.tail == nil {
		return
	}

	originalX := node.tail.x
	originalY := node.tail.y

	switch dir {
	case "R":
		node.x++

		if node.x == node.tail.x {
			break
		}

		if node.tail.x < node.x-1 {
			node.tail.y = node.y
			node.tail.x = node.x - 1
		}
	case "L":
		node.x--

		if node.x == node.tail.x {
			break
		}

		if node.tail.x > node.x+1 {
			node.tail.y = node.y
			node.tail.x = node.x + 1
		}
	case "U":
		node.y--

		if node.y == node.tail.y {
			break
		}

		if node.tail.y > node.y+1 {
			node.tail.x = node.x
			node.tail.y = node.y + 1
		}
	case "D":
		node.y++

		if node.y == node.tail.y {
			break
		}

		if node.tail.y < node.y-1 {
			node.tail.x = node.x
			node.tail.y = node.y - 1
		}
	}

	if originalX == node.tail.x && originalY == node.tail.y-1 {
		dir = "U"
	} else if originalX == node.tail.x && originalY == node.tail.y+1 {
		dir = "D"
	} else if originalY == node.tail.y && originalX == node.tail.x-1 {
		dir = "L"
	} else if originalY == node.tail.y && originalX == node.tail.x+1 {
		dir = "R"
	} else {
		return
	}

	if node.tail != nil {
		moveTail(node.tail, dir)
	}
}

func moveTail(node *ropeNode, dir string) {

	if node.tail == nil {
		return
	}

	originalX := node.tail.x
	originalY := node.tail.y

	switch dir {
	case "R":
		if node.x == node.tail.x {
			break
		}

		if node.tail.x < node.x-1 {
			node.tail.y = node.y
			node.tail.x = node.x - 1
		}
	case "L":
		if node.x == node.tail.x {
			break
		}

		if node.tail.x > node.x+1 {
			node.tail.y = node.y
			node.tail.x = node.x + 1
		}
	case "U":
		if node.y == node.tail.y {
			break
		}

		if node.tail.y > node.y+1 {
			node.tail.x = node.x
			node.tail.y = node.y + 1
		}
	case "D":
		if node.y == node.tail.y {
			break
		}

		if node.tail.y < node.y-1 {
			node.tail.x = node.x
			node.tail.y = node.y - 1
		}
	}

	if originalX == node.tail.x && originalY == node.tail.y-1 {
		dir = "U"
	} else if originalX == node.tail.x && originalY == node.tail.y+1 {
		dir = "D"
	} else if originalY == node.tail.y && originalX == node.tail.x-1 {
		dir = "L"
	} else if originalY == node.tail.y && originalX == node.tail.x+1 {
		dir = "R"
	} else {
		return
	}

	moveTail(node.tail, dir)
}
