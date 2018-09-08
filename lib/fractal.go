package lib

import "fmt"

type Fractal struct {
	Size  int
	cells [][]bool
}

func NewFractal(size int) Fractal {
	f := Fractal{Size: size}

	f.cells = make([][]bool, size)

	for row := 0; row < size; row++ {
		f.cells[row] = make([]bool, size)

		for col := 0; col < size; col++ {
			f.cells[row][col] = false
		}
	}

	return f
}

func (f *Fractal) Plot(row, col int, value bool) {
	f.cells[row][col] = value
}

func (f *Fractal) Fill(row1, col1, row2, col2 int, value bool) {
	for row := row1; row < row2; row++ {
		for col := col1; col < col2; col ++ {
			f.Plot(row, col, value)
		}
	}
}

func (f *Fractal) Compute(row1, col1, row2, col2 int) {
	size := row2 - row1

	if size >= 3 {
		chunk := size / 3
		r1 := row1 + chunk
		r2 := row1 + 2*chunk
		c1 := col1 + chunk
		c2 := col1 + 2*chunk
		f.Fill(r1, c1, r2, c2, true)

		//f.Compute(row1, col1, row1+chunk, col1+chunk)
		//f.Compute(row1, c1, r1, c2)
		f.Compute(row1, col1, r1, c1)
		f.Compute(row1, c1, r1, c2)
		f.Compute(row1, c2, r1, col2)
		f.Compute(r1, col1, r2, c1)
		f.Compute(r1, c2, r2, col2)
		f.Compute(r2, col1, row2, c1)
		f.Compute(r2, c1, row2, c2)
		f.Compute(r2, c2, row2, col2)
	}
}

func (f *Fractal) Display() {
	for row := 0; row < f.Size; row++ {
		for col := 0; col < f.Size; col++ {
			if f.cells[row][col] {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
