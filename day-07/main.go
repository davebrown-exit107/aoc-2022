package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type FileSystem struct {
	root *Node
	cwd  *Node
}

type Node struct {
	nodeType string
	name     string
	size     int
	parent   *Node
	members  []*Node
}

// Change the working directory of a filesystem
func (fs *FileSystem) cd(dir string) {
	for _, path := range fs.cwd.members {
		if path.name == dir && path.nodeType == "dir" {
			fs.cwd = path
		}
	}
	// find a node with type dir and name dir
	// set it as cwd for fs
	//fs.cwd = dir
}

// Make a new node and add it as a child of the cwd
func (fs *FileSystem) mkNode(nodeType, name string, size int) {
	newNode := &Node{
		name:     name,
		nodeType: nodeType,
		size:     size,
		parent:   fs.cwd,
	}
	fs.cwd.members = append(fs.cwd.members, newNode)
}

// return the size of the current directory and all nodes under it
func (fs *FileSystem) dirSize() (dirSize int) {
	return 0
}

// Solver for Day 07 Pt 01
func DaySevenPartOne(input []string) (dirSize int) {
	fs := FileSystem{}
	for _, line := range input {
		// split on the first space to get command vs listing
		c1, c2, _ := strings.Cut(line, " ")
		if c1 == "$" {
			// we're working on a command
			if c2 == "ls" {
				// ls stuff here
				// nothing needs to actually happen here, as the cwd is already set
				// and all of the node creation happens below
				continue
			} else {
				// kind of cheating here, but we have a limited command pallete...
				// do cd stuff here
				// create a new directory node and cd into it
				_, path, _ := strings.Cut(c2, " ")
				if path == "/" {
					// create the root node
					rootNode := &Node{
						nodeType: "dir",
						name:     "/",
						size:     0,
						members:  []*Node{}}
					fs.root = rootNode
					fs.cwd = rootNode
				} else {
					fs.mkNode("dir", path, 0)
				}

			}
		}
		// we're not working on commands but files and directories
		if c1 == "dir" {
			// we're at a directory, create the node and move on
		} else {
			// again, kind of taking a short cut here and assuming if it isn't a dir, it's a file
			// we're at a file, create the node and move on
		}
	}

	return
}

/*
// Solver for Day 07 Pt 02
func DaySevenPartTwo(input []string) (numChars int) {
	return
}
*/

func main() {
	// File opening boilerplate
	readFile, err := os.Open("input.txt")
	if err != nil {
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	// AoC work starts here
	directorySize := DaySevenPartOne(fileLines)
	fmt.Printf("Sum of directory sizes: %v\n", directorySize)

	/*
		numChars = DaySevenPartTwo(fileLines)
		fmt.Printf("Characters til first message: %v\n", numChars)
	*/
}
