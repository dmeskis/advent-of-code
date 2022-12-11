package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Node struct {
	parent   *Node
	children []*Node
	size     int
	name     string
}

type fileTable map[string]Node

func main() {
	fi, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err)
	}

	if !(fi.Size() > 0) {
		log.Fatal("no input")

	}

	s := bufio.NewScanner(os.Stdin)

	rootNode := Node{
		name: "/",
	}

	var curNode *Node
	list := []*Node{}

	for s.Scan() {
		line := strings.Fields(s.Text())
		if line[0] == "$" {
			if line[1] == "cd" {
				// handle root case
				if line[2] == "/" {
					curNode = &rootNode

				}
				// assign list to curNode if any & reset
				if len(list) > 0 {
					curNode.children = list
					list = []*Node{}
				}

				if line[2] == ".." {
					curNode = curNode.parent
				} else {
					for _, child := range curNode.children {
						if child.name == line[2] {
							curNode = child
							break
						}
					}
				}

			}

			if line[1] == "ls" {
				// no-op (for now)
			}

			continue
		}

		if line[0] == "dir" {
			// add dir to the child list
			node := Node{
				parent: curNode,
				name:   line[1],
			}
			list = append(list, &node)
		} else {
			// file case
			val, err := strconv.Atoi(line[0])
			if err != nil {
				log.Fatal(err)

			}
			node := Node{
				parent: curNode,
				name:   line[1],
				size:   val,
			}
			list = append(list, &node)
		}
	}

	// Catch any missed children..
	if len(list) > 0 {
		curNode.children = list
		list = []*Node{}
	}

	diskSpace := 70000000
	reqUnusedSpace := 30000000

	placeholder := 0
	dirSizes := []int{}
	size, _, sizes := getSize(&rootNode, &placeholder, &dirSizes)

	// sort sizes,
	sort.Ints(*sizes)

	for _, s := range *sizes {
		if diskSpace-size+s >= reqUnusedSpace {
			fmt.Println(s)
			return

		}

	}
}

func getSize(node *Node, totalSize *int, dirSizes *[]int) (int, *int, *[]int) {
	size := 0

	size += node.size

	for _, child := range node.children {
		s, _, _ := getSize(child, totalSize, dirSizes)
		size += s
	}

	if node.size == 0 {
		if size <= 100000 {
			*totalSize += size
		}

		*dirSizes = append(*dirSizes, size)
	}

	return size, totalSize, dirSizes
}
