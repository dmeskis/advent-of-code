package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	// diskSpace := 70000000
	// reqUnusedSpace := 30000000
	placeholder := 0
	size, total := getSize(&rootNode, &placeholder)
	fmt.Println(*total)
	fmt.Println(size)
}

func getSize(node *Node, totalSize *int) (int, *int) {
	size := 0

	size += node.size

	for _, child := range node.children {
		s, _ := getSize(child, totalSize)
		size += s
	}

	if node.size == 0 {
		if size <= 100000 {
			*totalSize += size
		}

	}

	fmt.Println("node name: ", node.name)
	fmt.Println("node size: ", node.size)
	fmt.Println("size: ", size)
	fmt.Println("totalsize: ", *totalSize)
	fmt.Println("----")

	return size, totalSize
}
