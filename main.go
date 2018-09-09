package main

import "github.com/dbowers42/fractal/lib"

func main() {
	size := 81
	f := lib.NewFractal(size)
	f.Compute(0,0, size, size)
	f.Render("fractal.txt")
}
