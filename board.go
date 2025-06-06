package main

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


func (a *App) AlgoritmoNReinas(board [][]int, queens int) {

}


func (a *App) ValidarMovimiento(board [][]int, row, col int) []int {
	size := len(board)
	posiciones := []int{}

	for c := range size { //Movimientos verticales 
		if c != col {
			posiciones = append(posiciones, row, c)
		}
	}

	for r := range size { //Movimientos horizontales
		if r != row {
			posiciones = append(posiciones, r, col)
		}
	}

	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 { // Movimiento diagonal superior izquierda
		posiciones = append(posiciones, i, j)
	}

	for i, j := row+1, col+1; i < size && j < size; i, j = i+1, j+1 { // Movimiento diagonal inferior derecha
		posiciones = append(posiciones, i, j)
	}

	for i, j := row-1, col+1; i >= 0 && j < size; i, j = i-1, j+1 {  // Movimiento diagonal superior derecha
		posiciones = append(posiciones, i, j)
	}

	for i, j := row+1, col-1; i < size && j >= 0; i, j = i+1, j-1 { // Movimiento diagonal inferior izquierda
		posiciones = append(posiciones, i, j)
	}
	
	return posiciones
}