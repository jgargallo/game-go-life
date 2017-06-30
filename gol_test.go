package main_test

import (
	. "gol"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Gol", func() {

	var (
		board1Neighbour = NewBoard([]Cell{{1, 1}, {0, 1}})
		board2Neighbours *Board
		board3Neighbours *Board
		board4Neighbours *Board
	)

	BeforeEach(func() {
		board1Neighbour = NewBoard([]Cell{{1, 1}, {0, 1}})
		board2Neighbours = NewBoard([]Cell{{1, 1}, {0, 1}, {1, 2}})
		board3Neighbours = NewBoard([]Cell{{1, 1}, {0, 1}, {1, 2}, {1, 0}})
		board4Neighbours = NewBoard([]Cell{{1, 1}, {0, 1}, {1, 2}, {1, 0}, {2, 1}})
	})

	Describe("Categorizing surviving cells", func() {
		Context("Less than two alive neighbours", func() {
			It("Should NOT be alive", func() {
				Expect(board1Neighbour.Survives(Cell{1, 1})).To(BeFalse())
			})
		})
		Context("More than three alive neighbours", func() {
			It("Should NOT be alive", func() {
				Expect(board4Neighbours.Survives(Cell{1, 1})).To(BeFalse())
			})
		})
		Context("2 or 3 alive neighbours", func() {
			It("Should be alive", func() {
				Expect(board3Neighbours.Survives(Cell{1, 1})).To(BeTrue())
			})
		})
	})

	Describe("Categorizing resurrected cells", func() {
		Context("A dead cell with exactly 3 neighbours", func() {
			It("Should resurrect", func() {
				Expect(board2Neighbours.Resurrects(Cell{0, 2})).To(BeTrue())
			})
		})
		Context("A dead cell with NOT exactly 3 neighbours", func() {
			It("Should NOT resurrect", func() {
				Expect(board2Neighbours.Resurrects(Cell{2, 0})).To(BeFalse())
			})
		})
	})

	Describe("Next Generation", func() {
		Context("3 Alive cells in 0,1 | 1,1 | 1,2", func() {
			It("Should keep alive the 3 cells and resurrect 0,2", func() {
				nextGen := board2Neighbours.NextGeneration()
				Expect(nextGen.IsAlive(Cell{1,1})).To(BeTrue())
				Expect(nextGen.IsAlive(Cell{0,1})).To(BeTrue())
				Expect(nextGen.IsAlive(Cell{1,2})).To(BeTrue())
				Expect(nextGen.IsAlive(Cell{0,2})).To(BeTrue())
			})
		})
	})

})
