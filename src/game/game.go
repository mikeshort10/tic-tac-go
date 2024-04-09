package game

type Board [9]int

func NewBoard() Board {
	return Board {-1,-1,-1,-1,-1,-1,-1,-1,-1}
}
 
func determineRowWinner(board Board) int {
	for i := 0; i < 3; i++ {
		index := i * 3
		if board[index] != -1 && board[index] == board[index + 1] && board[index + 1] == board[index + 2] {
			return board[index]
		}
	}
	return -1
}

func determineColumnWinner(board Board) int {
	for i := 0; i < 3; i++ {
		if board[i] != -1 && board[i] == board[i + 3] && board[i + 3] == board[i + 6] {
			return board[i]
		}
	}
	return -1
}

func determineDiagonalWinner(board Board) int {
	if (board[4] == -1) {
		return -1
	}	else if board[0] == board[4] && board[4] == board[8] {
		return board[4]
	} else if board[2] == board[4] && board[4] == board[6] {
		return board[4]
	}
	return -1
} 

func DetermineWinner(board Board) int {
	rowWinner := determineRowWinner(board)
	columnWinner := determineColumnWinner(board)
	diagonalWinner := determineDiagonalWinner(board)
	if rowWinner != -1 {
		return rowWinner
	} else if columnWinner != -1 {
		return columnWinner
	} else if diagonalWinner != -1 {
		return diagonalWinner
	}
	return -1
}

func DetermineTie(board Board) bool {
	for i := 0; i < len(board); i++ {
		if board[i] == -1 {
			return false
		}
	}
	return true
}

func UpdateBoard(board Board, index int, player int) (Board, bool) {
	if board[index - 1] != -1 {
		return board, false
	}
	board[index - 1] = player
	return board, true
}

