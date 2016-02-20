package main

import (
	"math/rand"
	"time"
)

// the grid
var grid [WIDTH * WIDTH]string
var neighbourliness int
var limit int

// initialize the grid with the initial population
func populate(nb, races, vacancy, lm int) {
	rand.Seed(time.Now().Unix())
	neighbourliness = nb
	limit = lm
	v := float64(vacancy) / 10.0
	for n, _ := range grid {
		rnd := rand.Float64()
		// v % of the grid are empty
		if rnd > v {
			rnd2 := rand.Intn(races)
			switch rnd2 {
			case 0:
				grid[n] = "#FF5050"
			case 1:
				grid[n] = "#0099FF"
			case 2:
				grid[n] = "#FFCC00"
			case 3:
				grid[n] = "#009900"
			case 4:
				grid[n] = "#CC00CC"
			case 5:
				grid[n] = "#FF9900"
			}
		} else {
			grid[n] = "white"
		}
	}
}

// every tick processes all cells in the grid
func tick() {
	// check every cell
	for cellNumber, cell := range grid {
		// if cell is empty, go to the next cell
		if cell == "white" {
			continue
		}
		// find all the cell's neighbours
		neighbours := findNeighboursIndex(cellNumber)
		// count of the neighbours that are the same as the cell
		sameCount := 0
		// for every neighbour
		for _, neighbour := range neighbours {
			// if the cell is empty, go to the next neighbour
			if grid[neighbour] == "white" {
				continue
			}
			// if the neighbour is the same, increment sameCount
			if grid[neighbour] == cell {
				sameCount++
			}
		}
		// if there are 2 or less neighbours that are the same
		// as this cell
		if sameCount <= neighbourliness || sameCount > limit {
			empty := findEmpty()
			e := findRandomEmpty(empty)
			grid[e] = cell
			grid[cellNumber] = "white"
		}
	}
}

// find the index of a random empty cell in the grid
func findRandomEmpty(empty []int) int {
	r := rand.Intn(len(empty))
	return empty[r]
}

// find all cells that are empty in the grid
func findEmpty() (empty []int) {
	for n, cell := range grid {
		if cell == "white" {
			empty = append(empty, n)
		}
	}
	return
}
