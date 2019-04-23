package main

import (
	"sort"
)

const UNDERPOPULATION_LIMIT = 2
const OVERPOPULATION_LIMIT = 3
const RESURRECTION = 3

type Cell [2]int
type Cells []Cell

func (slice Cells) Len() int {
	return len(slice)
}

func (slice Cells) Less(i, j int) bool {
	return slice[i][0] < slice[j][0] || slice[i][0] == slice[j][0] && slice[i][1] < slice[j][1]
}

func (slice Cells) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

type Board struct {
	AliveCells       []Cell
	neighboursOffset []Cell
}

func (board *Board) IsAlive(cell Cell) bool {
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
		x, y := cell[0]+n[0], cell[1]+n[1]
		if board.IsAlive(Cell{x, y}) {
			num++
		}
	}

	return num
}

func (board *Board) Survives(cell Cell) bool {
	return board.IsAlive(cell) &&
		(board.numAliveNeighbours(cell) == UNDERPOPULATION_LIMIT ||
			board.numAliveNeighbours(cell) == OVERPOPULATION_LIMIT)
}

func (board *Board) Resurrects(cell Cell) bool {
	return !board.IsAlive(cell) && board.numAliveNeighbours(cell) == RESURRECTION
}

func (board *Board) NextGeneration() *Board {
	cells := make(map[Cell]bool)
	for _, c := range board.AliveCells {
		if board.Survives(c) {
			cells[c] = true
		}
		for _, n := range board.neighboursOffset {
			target := Cell{c[0] + n[0], c[1] + n[1]}
			if board.Resurrects(target) {
				cells[target] = true
			}
		}
	}

	keys := make(Cells, len(cells))

	i := 0
	for k := range cells {
		keys[i] = k
		i++
	}

	return NewBoard(keys)
}

func NewBoard(cells Cells) *Board {
	sort.Sort(cells)
	board := new(Board)
	board.AliveCells = cells
	board.neighboursOffset = []Cell{
		{-1, 1}, {0, 1}, {1, 1},
		{-1, 0}, {1, 0},
		{-1, -1}, {0, -1}, {1, -1},
	}

	return board
}
