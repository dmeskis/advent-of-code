package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// struct
type Node struct {
	parent   *Node // can be nil for root
	children []*Node
	size     int
	name     string // cuz fuck an inode
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
		line := s.Text()
		// starting with $ signifies a command
		split := strings.Split(line, " ")
		if strings.HasPrefix(line, "$") {
			// $ cd a
			// $ ls
			if split[1] == "cd" {
				// handle root case
				if split[2] == "/" {
					curNode = &rootNode

				}
				// assign list to curNode if any & reset
				if len(list) > 0 {
					curNode.children = list
					list = []*Node{}
				}

				if split[2] == ".." {
					curNode = curNode.parent
				} else {
					for _, child := range curNode.children {
						if child.name == split[2] {
							curNode = child
							break
						}
					}
				}

			}

			if split[1] == "ls" {
				// no-op (for now)
			}

			continue
		}

		if strings.HasPrefix(line, "dir") {
			// add dir to the child list
			node := Node{
				parent: curNode,
				name:   split[1],
			}
			list = append(list, &node)
		} else {
			// file case
			val, err := strconv.Atoi(split[0])
			if err != nil {
				log.Fatal(err)

			}
			node := Node{
				parent: curNode,
				name:   split[1],
				size:   val,
			}
			list = append(list, &node)
		}
	}

	placeholder := 0
	_, total := getSize(&rootNode, &placeholder)
	fmt.Println(*total)
}

// totalsize gets passed thru
func getSize(node *Node, totalSize *int) (int, *int) {
	size := 0
	if node.size > 0 {
		size += node.size
	} else {
		for _, child := range node.children {
			s, _ := getSize(child, totalSize)
			size += s
		}
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
