package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*    PROBLEM STATEMENT:
Challenge is to make a model for sodoku so that given
the input string of the boards state you can implement
a method for IsValid where the method returns false
if there are any duplicates in rows, cols, or 3x3
matrices.
Blank spots are represented in the input as zeros.
*/

type sodoku struct {
	Size  int
	Dim   int
	Rows  [][]int
	Grids [][][]int
}

func (s *sodoku) ParseBoard(boardMap string) {
	s.Size = 9
	s.Dim = 3

	// load rows
	tail := 0
	for tail < s.Size*s.Size {
		var row []int = make([]int, s.Size)
		chars := strings.Split(boardMap[tail:tail+s.Size], "")
		for i, c := range chars {
			row[i], _ = strconv.Atoi(c)
		}
		s.Rows = append(s.Rows, row)
		tail += s.Size
	}

	// load Grids
	s.Grids = make([][][]int, s.Dim)
	tail = 0
	for g := 0; g < s.Size && tail < s.Size; g++ {
		s.Grids[g] = make([][]int, s.Size, s.Size)
		for r := 0; r < s.Size; r++ {
			s.Grids[g][r] = s.Rows[r][tail : tail+s.Dim]
			//fmt.Println(s.Grids[g][r])
		}
		tail += s.Dim
	}
}

func (s *sodoku) IsInGridAgain(grid [][]int, val int, x int, y int) bool {
	for j := 0; j < len(grid); j++ {
		for r := 0; r < len(grid[0]); r++ {
			// skip x,y position of values first occurence
			if j == x && r == y {
				continue
			}
			if grid[j][r] == val {
				return true
			}
		}
	}
	return false
}

func (s *sodoku) IsValid() bool {
	// row duplicates?
	for _, row := range s.Rows {
		for i, n := range row {
			if n == 0 {
				continue
			}
			for _, x := range row[i+1:] {
				if x == n {
					fmt.Printf("Row Duplicate\n")
					return false
				}
			}
		}
	}

	// col duplicates
	for j := 0; j < s.Size; j++ { // each col
		for r := 0; r < s.Size; r++ { // each row
			cv := s.Rows[r][j]
			if cv == 0 {
				continue
			}
			for k := r + 1; k < s.Size; k++ {
				if s.Rows[k][j] == cv {
					fmt.Printf("Column Duplicate\n")
					return false
				}
			}
		}
	}

	// grid duplicates
	for gridNum := 0; gridNum < s.Dim; gridNum++ {
		grid := s.Grids[gridNum][gridNum : gridNum+s.Dim]
		for dx := 0; dx < s.Dim; dx++ {
			for dy := 0; dy < s.Dim; dy++ {
				val := grid[dx][dy]
				if val == 0 {
					continue
				}
				if s.IsInGridAgain(grid, val, dx, dy) {
					fmt.Printf("Grid Duplicate\n")
					return false
				}
			}
		}
	}

	return true
}

func (s *sodoku) PrintBoard() {
	for _, r := range s.Rows {
		fmt.Println(r)
	}
}

func main() {
	var inputValid string = "530070000600195000098000060800060003400803001700020006060000280000419005000080079"
	var inputRowDuplicate string = "530070000600195500098000060800060003400803001700020006060000280000419005000080079"
	var inputColDuplicate string = "530070000600195000098100060800060003400803001700020006060000280000419005000080079"
	var inputGridDuplicate string = "530070000600195000098010060800060003400803001700020006060000280000419005000080079"

	a := sodoku{}
	a.ParseBoard(inputValid)
	a.PrintBoard()
	fmt.Printf("IsValid %t\n\n", a.IsValid())

	b := sodoku{}
	b.ParseBoard(inputRowDuplicate)
	b.PrintBoard()
	fmt.Printf("Row Duplicate - IsValid %t\n\n", b.IsValid())

	c := sodoku{}
	c.ParseBoard(inputColDuplicate)
	c.PrintBoard()
	fmt.Printf("Col Duplicate - IsValid %t\n\n", c.IsValid())

	d := sodoku{}
	d.ParseBoard(inputGridDuplicate)
	d.PrintBoard()
	fmt.Printf("Grid Duplicate - IsValid %t\n\n", d.IsValid())

}
