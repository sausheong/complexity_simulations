package main

import (
	"math"
	"math/rand"
)

func volunteerExp(cost, overall float64, maxPlayers int) [][]int {
	if maxPlayers == 0 {
		maxPlayers = 50
	}
	maxn := maxPlayers // number of players
	iterations := 2000 // number of iterations in an experiment
	c := cost          // individual cost of volunteering
	a := overall       // overall cost of no player volunteering
	// all experiments, element index is the number of players
	// each element shows the number of times an player volunteers in each experiment
	vals := make([]int, maxn)
	// for every experiment, starting with an experiment with only 1 player
	for n := 3; n < maxn; n++ {
		vcount := 0
		// probability of volunteering for each player
		p := 1 - math.Pow(c/a, 1/float64(n-1))
		for i := 0; i < iterations; i++ {
			// for every player
			for count := 0; count < n; count++ {
				r := rand.Float64()
				if r <= p {
					vcount++
					break
				}
			}
		}
		vals[n] = vcount
	}
	output := make([][]int, maxn-3)
	for c, v := range vals[3:] {
		output[c] = []int{c + 3, v * 100 / iterations}
	}
	return output
}
