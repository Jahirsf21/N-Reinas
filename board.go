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

func Reescribir(matrix [][]int, areglo [][]int) [][]int {
	for _, cordenadas := range areglo {
		x := cordenadas[0]
		y := cordenadas[1]
		matrix[x][y] = 2

	}
	return matrix
}

func EntradasPermitidas(matrix [][]int) [][]int {
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

func ValidarMovimiento(board [][]int, row, col int) [][]int {
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

func (a *App) AlgoritmoNReinas(board [][]int, queens int) (bool, [][]int) {
	board[0][0] = 1
	return AlgoritmoNReinas_Aux(board, queens)
}

func AlgoritmoNReinas_Aux(board [][]int, queens int) (bool, [][]int) {
	if queens == 0 {
		return true, board
	}

	var posibles = EntradasPermitidas(board)

	for _, posicion := range posibles {
		var x, y = posicion[0], posicion[1]

		var nuevoTablero = board
		nuevoTablero[x][y] = 1

		var bloqueos = ValidarMovimiento(nuevoTablero, x, y)
		nuevoTablero = Reescribir(nuevoTablero, bloqueos)

		var cumplido, res = AlgoritmoNReinas_Aux(nuevoTablero, queens-1)
		if cumplido {
			return true, res
		}
	}

	return false, board
}
