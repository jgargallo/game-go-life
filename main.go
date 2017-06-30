package main

import (
	"fmt"
)


func main() {
	block := []Cell{{1, 1}, {1, 2}, {2, 1}, {2, 2}}
	blinker := []Cell{{-1, 1}, {-1, 2}, {-1, 3}}
	glider := []Cell{{5, 4}, {6, 4}, {7, 4}, {7, 5}, {6, 6}}
	board := NewBoard(append(block, append(blinker, glider...)...))

	Play(board)

	fmt.Println("GAME OVER!")
}
