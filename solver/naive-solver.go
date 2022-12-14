package solver

import (
	"local/sudoku-sandbox/log"
	S "local/sudoku-sandbox/sudoku"
)

const (
	Testpuzzle string = "000260701680070090190004500820100040004602900050003028009300074040050036703018000"
	Solution   string = "435269781682571493197834562826195347374682915951743628519326874248957136763418259"
	N          int    = 9
)

func TrySolve(s *S.Sudoku, row int, col int) (bool, error) {

	// base case 1:  if 8th row and 9th column
	if row == N-1 && col == N {
		return true, nil
	}

	// base case 2: if end of a row (last col)
	//              move to beginning of next row
	if col == N {
		row++
		col = 0
		log.Debug("TrySolve", "end of row")
	}

	// Check current position for value (0 == empty).
	if s.Rows[row][col] != 0 {
		log.Debug("TrySolve", "Position already has value [r = %d, c = %d]", row, col)
		return TrySolve(s, row, col+1) // move to next col in row
	}

	// Try to fill empty position.
	for num := 1; num < N+1; num++ {

		// Insert the current num into position.
		s.Rows[row][col] = num

		// If value is valid in current position
		// we recursively call solve on next pos.
		if s.IsValid() {
			log.Debug("TrySolve - validated", "now recurse with next position")
			isSolved, err := TrySolve(s, row, col+1)
			if isSolved && err == nil {
				log.Info("TrySolve", "return true", "finish solving")
				return true, nil
			}
		}

		// If not valid then we remove the num.
		log.Debug("TrySolve - try value INVALID", "reset to 0")
		s.Rows[row][col] = 0
	}

	log.Debug("TrySolve", "return FALSE")
	return false, nil
}
