# Sudoku Solver in Golang

This is a simple Golang program that solves a Sudoku puzzle using a backtracking algorithm.

## How to use

1. Clone the repository to your local machine.
2. Run the program using the command `go run main.go`.
3. The program will output the initial Sudoku puzzle and then the solved puzzle.

You can modify the `board` variable in the `main()` function to change the puzzle that is being solved.

```go
board := [][]int{
    {3, 1, 0, 0, 0, 0, 0, 2, 0},
    {0, 0, 6, 1, 0, 9, 0, 0, 5},
    {0, 0, 0, 0, 8, 0, 0, 0, 0},
    {0, 2, 0, 8, 0, 4, 0, 5, 0},
    {0, 0, 4, 0, 7, 0, 0, 0, 0},
    {0, 0, 0, 0, 6, 0, 0, 0, 8},
    {0, 6, 0, 0, 0, 0, 9, 0, 0},
    {0, 0, 9, 4, 0, 5, 0, 0, 1},
    {0, 0, 0, 0, 0, 7, 0, 0, 0},
}
```

## How it works

The program solves the Sudoku puzzle using a backtracking algorithm. The algorithm works as follows:

1. Find an empty cell in the puzzle.
2. Try each number from 1 to 9 in the empty cell.
3. If the number is valid (i.e. it does not violate any of the Sudoku rules), move to the next empty cell and repeat the process.
4. If there are no valid numbers for the current cell, backtrack to the previous cell and try a different number.
5. If the entire puzzle has been filled with valid numbers, the puzzle is solved.

The program uses a recursive function `solveSudoku()` to implement the backtracking algorithm. The function takes the current state of the puzzle as a parameter and returns `true` if the puzzle is solved and `false` if the puzzle cannot be solved.

The program also includes a helper function `isValid()` that checks whether a number is valid for a particular cell in the puzzle. The function checks whether the number violates any of the Sudoku rules (i.e. the number is not already present in the same row, column, or 3x3 box as the cell).

## License

This program is released under the MIT License. See the [LICENSE](LICENSE) file for more information.