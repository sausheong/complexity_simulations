package main

// width of the grid, which is a square of WIDTH * WIDTH
const WIDTH int = 36

// find the adjacent 8 neighbours, taking care of cases for corners and sides
func findNeighboursIndex(n int) (nb []int) {
	switch {
	// corner cases
	case top_left(n):
		nb = append(nb, c5(n))
		nb = append(nb, c7(n))
		nb = append(nb, c8(n))
		return
	case top_right(n):
		nb = append(nb, c4(n))
		nb = append(nb, c6(n))
		nb = append(nb, c7(n))
		return
	case bottom_left(n):
		nb = append(nb, c2(n))
		nb = append(nb, c3(n))
		nb = append(nb, c5(n))
		return
	case bottom_right(n):
		nb = append(nb, c1(n))
		nb = append(nb, c2(n))
		nb = append(nb, c4(n))
		return
		// side cases
	case top(n):
		nb = append(nb, c4(n))
		nb = append(nb, c5(n))
		nb = append(nb, c6(n))
		nb = append(nb, c7(n))
		nb = append(nb, c8(n))
		return
	case left(n):
		nb = append(nb, c2(n))
		nb = append(nb, c3(n))
		nb = append(nb, c5(n))
		nb = append(nb, c7(n))
		nb = append(nb, c8(n))
		return
	case right(n):
		nb = append(nb, c1(n))
		nb = append(nb, c2(n))
		nb = append(nb, c4(n))
		nb = append(nb, c6(n))
		nb = append(nb, c7(n))
		return
	case bottom(n):
		nb = append(nb, c1(n))
		nb = append(nb, c2(n))
		nb = append(nb, c3(n))
		nb = append(nb, c4(n))
		nb = append(nb, c5(n))
		return
		// everything else
	default:
		nb = append(nb, c1(n))
		nb = append(nb, c2(n))
		nb = append(nb, c3(n))
		nb = append(nb, c4(n))
		nb = append(nb, c5(n))
		nb = append(nb, c6(n))
		nb = append(nb, c7(n))
		nb = append(nb, c8(n))
	}
	return
}

// functions to check for corners and sides
func top_left(n int) bool     { return n == 0 }
func top_right(n int) bool    { return n == WIDTH-1 }
func bottom_left(n int) bool  { return n == WIDTH*(WIDTH-1) }
func bottom_right(n int) bool { return n == (WIDTH*WIDTH)-1 }

func top(n int) bool    { return n < WIDTH }
func left(n int) bool   { return n%WIDTH == 0 }
func right(n int) bool  { return n%WIDTH == WIDTH-1 }
func bottom(n int) bool { return n >= WIDTH*(WIDTH-1) }

// functions to get the index of the neighbours
func c1(n int) int { return n - WIDTH - 1 }
func c2(n int) int { return n - WIDTH }
func c3(n int) int { return n - WIDTH + 1 }
func c4(n int) int { return n - 1 }
func c5(n int) int { return n + 1 }
func c6(n int) int { return n + WIDTH - 1 }
func c7(n int) int { return n + WIDTH }
func c8(n int) int { return n + WIDTH + 1 }
