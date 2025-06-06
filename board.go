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