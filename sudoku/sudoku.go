package sudoku

import (
	"fmt"
	"strconv"
	"strings"

    "local/sudoku-sandbox/log"
)

type Sudoku struct {
	Size  int
	Dim   int
	Rows  [][]int
	Grids [][][]int
}

/*
 *    Parse map string isnto a sudoku struct.
 */
func (s *Sudoku) ParseBoard(boardMap string) {
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

func (s *Sudoku) IsValid() bool {
	// row duplicates?
	for _, row := range s.Rows {
		for i, n := range row {
			if n == 0 {
				continue
			}
			for _, x := range row[i+1:] {
				if x == n {
					log.Debug("sudoku.IsValid", "Row Duplicate\n")
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
					log.Debug("sudoku.IsValid", "Column Duplicate\n")
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
				if IsInGridAgain(grid, val, dx, dy) {
					log.Debug("sudoku.IsValid", "Grid Duplicate\n")
					return false
				}
			}
		}
	}

	return true
}

func (s *Sudoku) PrintBoard() {
	for _, r := range s.Rows {
		fmt.Println(r)
	}
}

//    private

func IsInGridAgain(grid [][]int, val int, x int, y int) bool {
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
