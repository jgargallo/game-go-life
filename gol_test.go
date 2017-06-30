package main_test

import (
	. "gol"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Gol", func() {

	var (
		board2Neighbours *Board
		board1Neighbour *Board
	)

	BeforeEach(func() {
		board2Neighbours = NewBoard([]Cell{{1, 1}, {0, 1}, {1, 2}})
		board1Neighbour = NewBoard([]Cell{{1, 1}, {0, 1}})
	})

	Describe("Categorizing surviving cells", func() {
		Context("Less than two alive neighbours", func() {
			It("Shoudn't be alive", func() {
				Expect(board1Neighbour.Survives(Cell{1, 1})).To(BeFalse())
			})
		})
	})

})
