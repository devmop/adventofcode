package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	p := pad(5)

	for s.Scan() {
		for _, c := range s.Text() {
			switch c {
			case 'U':
				p = up[p]
			case 'D':
				p = down[p]
			case 'L':
				p = left[p]
			case 'R':
				p = right[p]
			}
		}

		fmt.Println(p)
	}
}

type pad uint

var up = []pad{0, 1, 2, 1, 4, 5, 2, 3, 4, 9, 6, 7, 8, 11}
var down = []pad{0, 3, 6, 7, 8, 5, 10, 11, 12, 9, 10, 13, 12, 13}
var left = []pad{0, 1, 2, 2, 3, 5, 5, 6, 7, 8, 10, 10, 11, 13}
var right = []pad{0, 1, 3, 4, 4, 6, 7, 8, 9, 9, 11, 12, 12, 13}
