package main

import (
	"math/rand"
	"time"
)

var cultureGrid [WIDTH * WIDTH]int

var fdistances []int
var changes []int
var uniques []int

var MASK int = 0x00000F
var MASKARRAY []int = []int{0xFFFFF0, 0xFFFF0F, 0xFFF0FF, 0xFF0FFF, 0xF0FFFF, 0x0FFFFF}

// initialize the cultureGrid
func culturePopulate() {
	rand.Seed(time.Now().Unix())
	for n, _ := range cultureGrid {
		cultureGrid[n] = rand.Intn(0xFFFFFF)
	}
	fdistances, changes, uniques = []int{}, []int{}, []int{}
}

// every tick processes all cells in the cultureGrid
func cultureTick() {
	var change int
	rand.Seed(time.Now().Unix())
	for c := 0; c < 2000; c++ {
		r := rand.Intn(WIDTH * WIDTH)
		neighbours := findNeighboursIndex(r)
		for _, neighbour := range neighbours {
			d := diff(r, neighbour)
			probability := 1 - float64(d)/96.0
			dp := rand.Float64()
			if dp < probability {
				i := rand.Intn(6)
				if d != 0 {
					var rp int
					if rand.Intn(1) == 0 {
						replacement := extract(cultureGrid[r], uint(i))
						rp = replace(cultureGrid[neighbour], replacement, uint(i))
					} else {
						replacement := extract(cultureGrid[neighbour], uint(i))
						rp = replace(cultureGrid[r], replacement, uint(i))
					}
					cultureGrid[neighbour] = rp
					change++
				}
			}
		}
	}
	fdistances = append(fdistances, featureDistAvg())
	changes = append(changes, change/WIDTH)
	uniques = append(uniques, similarCount())
}

// total distance between traits for all features, between 2 cultures
func diff(a1, a2 int) int {
	var d int
	for i := 0; i < 5; i++ {
		d = d + traitDistance(cultureGrid[a1], cultureGrid[a2], uint(i))
	}
	return d
}

// average feature distance for the whole grid
func featureDistAvg() int {
	var count int
	var dist int
	for c, _ := range cultureGrid {
		neighbours := findNeighboursIndex(c)
		for _, neighbour := range neighbours {
			count++
			dist = dist + featureDistance(cultureGrid[c], cultureGrid[neighbour])
		}
	}
	return dist / WIDTH
}

// distance between 2 features
func featureDistance(n1, n2 int) int {
	var features int = 0
	for i := 0; i < 5; i++ {
		f1, f2 := extract(n1, uint(i)), extract(n2, uint(i))
		if f1 == f2 {
			features++
		}
	}
	return 6 - features
}

// count unique colors
func similarCount() int {
	uniques := make(map[int]int)
	for _, c := range cultureGrid {
		uniques[c] = c
	}
	return len(uniques)
}

// find the distance of 2 numbers at position pos
func traitDistance(n1, n2 int, pos uint) int {
	d := extract(n1, pos) - extract(n2, pos)
	if d < 0 {
		return d * -1
	} else {
		return d
	}
}

// extract trait for 1 feature
func extract(n int, pos uint) int {
	return (n >> (4 * pos)) & MASK
}

// replace the trait in 1 feature
func replace(n, replacement int, pos uint) int {
	i1 := n & MASKARRAY[pos]
	mask2 := replacement << (4 * pos)
	return (i1 ^ mask2)
}
