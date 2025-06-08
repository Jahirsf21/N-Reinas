package main

import (
	"time"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"strconv"
)

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

func (a *App) Reescribir(matrix, arreglo [][]int) [][]int {
	for _, cordenadas := range arreglo {
		x := cordenadas[0]
		y := cordenadas[1]
		matrix[x][y] = 2

	}
	return matrix
}

func(a *App) CopiarMatriz(matrix [][]int) [][]int {
	n := len(matrix)
	copia := make([][]int, n)
	for i := range matrix {
		copia[i] = make([]int, len(matrix[i]))
		copy(copia[i], matrix[i])
	}
	return copia
}



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


func (a *App) AlgoritmoNReinas(board [][]int, queens int)  {

	cleanBoard := a.CreateBoard(len(board))
	a.AlgoritmoNReinas_Aux(cleanBoard,queens)
	
}

func (a *App) AlgoritmoNReinas_Aux(board [][]int, queens int) bool {

	if queens == 0 {
		runtime.EventsEmit(a.ctx, "algorithm-log", "¡Solución encontrada!")
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

		time.Sleep(1000 * time.Millisecond)

		if a.AlgoritmoNReinas_Aux(nuevoTablero, queens-1) {
			return true 
		}
		
		runtime.EventsEmit(a.ctx, "board-update", board) 

		runtime.EventsEmit(a.ctx, "algorithm-log", "Se regresa desde ("+strconv.Itoa(x)+", "+strconv.Itoa(y)+")")

		time.Sleep(1000 * time.Millisecond)
	}
	
	return false
}

