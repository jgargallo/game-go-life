package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func Clear() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Print(board *Board) {
	lenCells := len(board.AliveCells)
	minI, minJ := board.AliveCells[0][0], board.AliveCells[0][1]
	maxI, maxJ := board.AliveCells[lenCells-1][0], board.AliveCells[lenCells-1][1]
	for i := minI - 10; i <= maxI+10; i++ {
		for j := minJ - 10; j <= maxJ+10; j++ {
			if board.IsAlive(Cell{i, j}) {
				fmt.Print("\u2588")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

func Play(board *Board) {
	for len(board.AliveCells) > 0 {
		Clear()
		Print(board)
		time.Sleep(500 * time.Millisecond)
		//break
		board = board.NextGeneration()
	}
}

func main() {
	pulsar := Cells{
		{0, 2}, {0, 3}, {0, 4}, {0, 8}, {0, 9}, {0, 10},
		{2, 0}, {2, 5}, {2, 7}, {2, 12},
		{3, 0}, {3, 5}, {3, 7}, {3, 12},
		{4, 0}, {4, 5}, {4, 7}, {4, 12},
		{5, 2}, {5, 3}, {5, 4}, {5, 8}, {5, 9}, {5, 10},

		{7, 2}, {7, 3}, {7, 4}, {7, 8}, {7, 9}, {7, 10},
		{8, 0}, {8, 5}, {8, 7}, {8, 12},
		{9, 0}, {9, 5}, {9, 7}, {9, 12},
		{10, 0}, {10, 5}, {10, 7}, {10, 12},
		{12, 2}, {12, 3}, {12, 4}, {12, 8}, {12, 9}, {12, 10},
	}
	board := NewBoard(pulsar)

	Play(board)
}
