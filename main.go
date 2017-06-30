package main

import (
	"fmt"
)


func main() {
	block := [][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}}
	blinker := [][]int{{-1, 1}, {-1, 2}, {-1, 3}}
	glider := [][]int{{5, 4}, {6, 4}, {7, 4}, {7, 5}, {6, 6}}
	board := NewBoard(append(block, append(blinker, glider...)...))

	fmt.Println(*board)

	fmt.Println("GAME OVER!")
}
