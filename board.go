package main

/**
* time: to handle delays in the algorithm
* runtime: to emit events to the frontend (Wails)
* strconv: to convert integers to strings
 */
import (
	"strconv"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

/**
* Creates a size x size board initialized with zeros.
 */
func (a *App) CreateBoard(size int) [][]int {
	var board [][]int
	for range size {
		var row []int
		for range size {
			row = append(row, 0)
		}
		board = append(board, row)
	}
	return board
}

/**
* Sets cells in the board to 2 using positions from an array given
 */
func (a *App) Reescribir(matrix, arreglo [][]int) [][]int {
	for _, cordenadas := range arreglo {
		x := cordenadas[0]
		y := cordenadas[1]
		matrix[x][y] = 2

	}
	return matrix
}

/**
* Makes a copy of the matrix
* Returns the copy
 */
func (a *App) CopiarMatriz(matrix [][]int) [][]int {
	n := len(matrix)
	copia := make([][]int, n)
	for i := range matrix {
		copia[i] = make([]int, len(matrix[i]))
		copy(copia[i], matrix[i])
	}
	return copia
}

/**
* Finds all available (0) positions on the board.
* Return the list with all the positions found.
 */
func (a *App) EntradasPermitidas(matrix [][]int) [][]int {
	var lista [][]int

	for x := range matrix {
		for y := range matrix[x] {
			if matrix[x][y] == 0 {
				lista = append(lista, []int{x, y})
			}
		}
	}
	return lista
}

/**
* Returns a list of positions attacked by a queen placed at (row, col).
* It includes the same row, column, and all diagonals.
 */
func (a *App) ValidarMovimiento(board [][]int, row, col int) [][]int {
	size := len(board)
	posiciones := [][]int{}

	for c := 0; c < size; c++ {
		if c != col {
			posiciones = append(posiciones, []int{row, c})
		}
	}

	for r := 0; r < size; r++ {
		if r != row {
			posiciones = append(posiciones, []int{r, col})
		}
	}

	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		posiciones = append(posiciones, []int{i, j})
	}

	for i, j := row+1, col+1; i < size && j < size; i, j = i+1, j+1 {
		posiciones = append(posiciones, []int{i, j})
	}

	for i, j := row-1, col+1; i >= 0 && j < size; i, j = i-1, j+1 {
		posiciones = append(posiciones, []int{i, j})
	}

	for i, j := row+1, col-1; i < size && j >= 0; i, j = i+1, j-1 {
		posiciones = append(posiciones, []int{i, j})
	}

	return posiciones
}

/**
* Initializes an empty board and starts the N-Queens backtracking algorithm.
* Sends an event when the algorithm finishes with or without a solution.
 */
func (a *App) AlgoritmoNReinas(board [][]int, queens int) {
	cleanBoard := a.CreateBoard(len(board))
	encontrado := a.AlgoritmoNReinas_Aux(cleanBoard, queens)
	if !encontrado {
		runtime.EventsEmit(a.ctx, "algorithm-finished", false)
	}
}

/**
* Recursive function to try placing queens.
* At each step, it updates the board, blocks threatened positions, and tries the next queen.
* Emits events for visual updates and logs the decision-making process.
 */
func (a *App) AlgoritmoNReinas_Aux(board [][]int, queens int) bool {

	if queens == 0 {
		runtime.EventsEmit(a.ctx, "algorithm-finished", true)
		return true
	}

	posibles := a.EntradasPermitidas(board)

	if len(posibles) == 0 && queens > 0 {
		return false
	}

	for _, posicion := range posibles {
		x, y := posicion[0], posicion[1]

		nuevoTablero := a.CopiarMatriz(board)

		nuevoTablero[x][y] = 1

		bloqueos := a.ValidarMovimiento(nuevoTablero, x, y)
		nuevoTablero = a.Reescribir(nuevoTablero, bloqueos)

		runtime.EventsEmit(a.ctx, "board-update", nuevoTablero)

		runtime.EventsEmit(a.ctx, "algorithm-log", "Intentando reina en ("+strconv.Itoa(x)+", "+strconv.Itoa(y)+")")

		time.Sleep(500 * time.Millisecond)

		if a.AlgoritmoNReinas_Aux(nuevoTablero, queens-1) {
			return true
		}

		runtime.EventsEmit(a.ctx, "board-update", board)

		runtime.EventsEmit(a.ctx, "algorithm-log", "Se regresa desde ("+strconv.Itoa(x)+", "+strconv.Itoa(y)+")")

		time.Sleep(500 * time.Millisecond)
	}

	return false
}
