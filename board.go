package main

const UNDERPOPULATION_LIMIT = 2
const OVERPOPULATION_LIMIT = 3

type Cell [2]int

type Board struct {
	AliveCells []Cell
	neighboursOffset []Cell
}

func (board *Board) isAlive(cell Cell) bool {
	for _, c := range board.AliveCells {
		if c == cell {
			return true
		}
	}
	return false
}

func (board *Board) numAliveNeighbours(cell Cell) int {
	num := 0
	for _, n := range board.neighboursOffset {
		x, y := cell[0] + n[0], cell[1] + n[1]
		if board.isAlive(Cell{x, y}) { num++ }
	}

	return num
}

func (board *Board) Survives(cell Cell) bool {
	return board.isAlive(cell) &&
			(board.numAliveNeighbours(cell) == UNDERPOPULATION_LIMIT ||
			 board.numAliveNeighbours(cell) == OVERPOPULATION_LIMIT)
}

func NewBoard(conf []Cell) *Board {
	board := new(Board)
	board.AliveCells = conf
	board.neighboursOffset = []Cell{
		{-1, 1} , {0, 1} , {1, 1},
		{-1, 0} ,          {1, 0},
		{-1, -1}, {0, -1}, {1, -1},
	}

	return board
}
