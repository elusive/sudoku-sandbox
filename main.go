package main

import (
	"fmt"
	"io"
	"os"

	"local/sudoku-sandbox/log"
	V "local/sudoku-sandbox/solver"
	S "local/sudoku-sandbox/sudoku"
)

/*    PROBLEM STATEMENT:
Challenge is to make a model for sodoku so that given
the input string of the boards state you can implement
a method for IsValid where the method returns false
if there are any duplicates in rows, cols, or 3x3
matrices.
Blank spots are represented in the input as zeros.
*/

var out io.Writer = os.Stdout

func main() {
	/*
		   	var inputValid string = "530070000600195000098000060800060003400803001700020006060000280000419005000080079"
		   	var inputRowDuplicate string = "530070000600195500098000060800060003400803001700020006060000280000419005000080079"
		   	var inputColDuplicate string = "530070000600195000098100060800060003400803001700020006060000280000419005000080079"
		   	var inputGridDuplicate string = "53007000000195000098010060800060003400803001700020006060000280000419005000080079"

		       a := sodoku{}on of values first occurence
				if j == x && r == y {
					continue
		   	b.ParseBoard(inputRowDuplicate)
		   	b.PrintBoard()
		   	fmt.Printf("Row Duplicate - IsValid %t\n\n", b.IsValid())

		   	c := sodoku{}
		   	c.ParseBoard(inputColDuplicate)
		   	c.Printar out io.Writer = os.StdoutBoard()t
		   	fmt.Printf("Grid Duplicate - IsValid %t\n\n", d.IsValid())
	*/

	log.Init(out)
	//log.Verbose = true

	s := S.Sudoku{}
	s.ParseBoard(V.Testpuzzle)

	isSolved, err := V.TrySolve(&s, 0, 0)
	if isSolved && err == nil {
		fmt.Println("SOLVED")
		s.PrintBoard()
	} else {
		fmt.Println("NO Solution Exists")
	}
}
