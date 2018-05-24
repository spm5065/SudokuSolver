package main

import "log"
import "fmt"

type SudokuBoard struct {
	squares [81]int
}

func main() {
	//Load the board
	var baseBoard SudokuBoard
	baseBoard = LoadBoard()

	//Check that it is valid
	if !ValidBoard(baseBoard) {
		log.Fatal("The board provided was invalid.")
	}

	//Create list of permutations, and one for solutions
	var permutationBoards []SudokuBoard //All boards must be valid, but not necessarily complete
	var solutionBoards []SudokuBoard

	permutationBoards = append(permutationBoards, baseBoard)

	for {
		if len(permutationBoards) == 0 {
			break
		}

		//Grab last board and remove from ones to work through
		currentBoard := permutationBoards[len(permutationBoards)-1]
		permutationBoards = permutationBoards[:len(permutationBoards)-1]

		filled := true
		//Find Empty square in this board
		for i:=0; i<81; i++ {
			if currentBoard.squares[i] == 0 {
				//fmt.Println("Ahh")
				filled = false

				//Fill empty square with all possible numbers, and check if its valid
				//if so, add it to the permutations
				for j:=1; j < 10; j++{
					newBoard := currentBoard
					newBoard.squares[i] = j

					//Check if we have broken any rules if so, skip this board
					if !ValidBoard(newBoard) {
						continue
					}

					//Board is valid so add it to the next ones we will look at
					permutationBoards = append( permutationBoards, newBoard )
					//fmt.Printf("%d boards to attempt!\n", len(permutationBoards))
				}
				break
			}
		}

		if filled {
			solutionBoards = append(solutionBoards, currentBoard)
		}
	}

	fmt.Printf("%d solutions found!\n", len(solutionBoards))
	for i := 0; i < len(solutionBoards); i++ {
		fmt.Printf("\nSolution: %d\n", i+1)
		for j := 0; j < 9; j++ {
			for k := 0; k < 9; k++ {
				fmt.Printf("%d ", solutionBoards[i].squares[9*j + k])
			}
			fmt.Println("")
		}
	}
}

func ValidBoard(board SudokuBoard) bool {
	//Check rows
	for i:=0; i<9; i++ {
		var numbersUsed [9]int
		for j:=0; j<9; j++ {
			if board.squares[9*i + j] != 0 {
				//fmt.Printf( "%d,%d value:%d\n", j, i, board.squares[9*i + j] )
				if numbersUsed[board.squares[9*i + j]-1] != 0 {
					//fmt.Printf( "Invalid at %d,%d, value:%d\n", j, i, board.squares[9*i + j] )
					return false
				}
				numbersUsed[board.squares[9*i + j]-1]++
			}
		}
	}

	//Check columns
	for i:=0; i<9; i++ {
		var numbersUsed [9]int
		for j:=0; j<9; j++ {
			if board.squares[9*j + i] != 0 {
				if numbersUsed[board.squares[9*j + i]-1] != 0 {
					//fmt.Printf( "Invalid at %d,%d, value:%d\n", j, i, board.squares[9*i + j] )
					return false
				}
				numbersUsed[board.squares[9*j + i]-1]++
			}
		}
	} 


	//Check squares
	for i:=0; i<9; i++ {
		var numbersUsed [9]int
		for j:=0; j<9; j++ {
			x:= 3*(i%3)+ j%3
			y:= 3*(i/3)+ j/3

			if board.squares[9*y + x] != 0 {
				if numbersUsed[board.squares[9*y + x]-1] != 0 {
					//fmt.Printf( "Invalid at %d,%d, value:%d\n", j, i, board.squares[9*i + j] )
					return false
				}
				numbersUsed[board.squares[9*y + x]-1]++
			}
		}
	}

	return true
}

func LoadBoard() SudokuBoard {
	var board SudokuBoard

	//Row 1
	board.squares[1] = 6
	board.squares[4] = 8
	board.squares[5] = 2
	board.squares[6] = 3
	board.squares[8] = 4
	//Row 2
	board.squares[10] = 9
	board.squares[13] = 4
	//Row 3
	board.squares[18] = 3
	board.squares[19] = 4
	board.squares[22] = 1
	//Row 4
	board.squares[30] = 7
	board.squares[31] = 5
	board.squares[34] = 9
	//Row 5
	board.squares[38] = 9
	board.squares[39] = 8
	board.squares[40] = 2
	board.squares[41] = 1
	board.squares[42] = 4
	//Row 6
	board.squares[46] = 2
	board.squares[49] = 6
	board.squares[50] = 9
	//Row 7
	board.squares[58] = 3
	board.squares[61] = 7
	board.squares[62] = 6
	//Row 8
	board.squares[67] = 9
	board.squares[70] = 4
	//Row 9
	board.squares[72] = 1
	board.squares[74] = 4
	board.squares[75] = 6
	board.squares[76] = 7
	board.squares[79] = 8

	return board
}