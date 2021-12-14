// Package merrychristmas creates an asterisk Christmas tree given a height.
// If input file is empty, the Christmas tree height is 25.
package merrychristmas

import (
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// MakeTree reads an integer from a file named merry-christmas.txt and generates a tree to happy-holidays.txt
func MakeTree() {
	height := readFile()
	writeFile(height)
}

// readFile returns an integer from merry-christmas.txt that specifies the height of the tree.
// If file is empty, the tree will default to a height of 25
func readFile() int {
	content, err := os.ReadFile("./2021-12-13/merry-christmas.txt")
	check(err)

	var height int
	if len(content) == 0 {
		height = 25
	} else {
		height, err = strconv.Atoi(string(content))
		check(err)
	}

	return height
}

// makeBranch returns the number of spaces and asterisks to print at a given level of the tree.
func makeBranch(height int, level int) (int, int) {

	spaces := height - level
	leaves := level + level - 1

	return spaces, leaves
}

// writeFile prints an asterisk Christmas tree given the height of the tree.
func writeFile(height int) {
	f, err := os.Create("./2021-12-13/happy-holidays.txt")
	check(err)

	defer f.Close()

	for i := 1; i <= height; i++ {
		spaces, leaves := makeBranch(height, i)
		s := strings.Repeat(" ", spaces)
		l := strings.Repeat("*", leaves)
		_, err = f.WriteString(s + l + s)
		check(err)

		if i != height {
			_, err = f.WriteString("\n")
			check(err)
		}
	}

}
