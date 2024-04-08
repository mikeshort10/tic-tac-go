package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func clear() {
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func getPlayerChar(player int) string {
	if player == 0 {
		return "X"
	} else if player == 1 {
		return "0"
	}
	return " "
}

func determineRowWinner(board []int) int {
	for i := 0; i < 3; i++ {
		index := i * 3
		if board[index] != -1 && board[index] == board[index + 1] && board[index + 1] == board[index + 2] {
			return board[index]
		}
	}
	return -1
}

func determineColumnWinner(board []int) int {
	for i := 0; i < 3; i++ {
		if board[i] != -1 && board[i] == board[i + 3] && board[i + 3] == board[i + 6] {
			return board[i]
		}
	}
	return -1
}

func determineDiagonalWinner(board []int) int {
	if (board[4] == -1) {
		return -1
	}	else if board[0] == board[4] && board[4] == board[8] {
		return board[4]
	} else if board[2] == board[4] && board[4] == board[6] {
		return board[4]
	}
	return -1
} 

func determineWinner(board []int) int {
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

func determineTie(board []int) bool {
	for i := 0; i < len(board); i++ {
		if board[i] == -1 {
			return false
		}
	}
	return true
}

func renderBoard(board []int) {
	for i := 0; i < 9; i+=3 {
		row := board[i:i+3]
		line := " "
		for j := 0; j < 3; j++ {
			tic := getPlayerChar(row[j])
			line += tic //+ tic + "  "
			if j == 2 {
				line += " "
			} else {
				line += " | "
			}
		}
		fmt.Println(line) 
	}
}

func main() {
	player := 0
	winner := -1
	board := []int {-1,-1,-1,-1,-1,-1,-1,-1,-1}
	
	clear()
	for winner < 0 {

		var rawIndex string
		renderBoard(board)
		fmt.Print("Player ", getPlayerChar(player), " enter a space to play: ")
		fmt.Scanln(&rawIndex)
		clear()

		index, err := strconv.Atoi(rawIndex)

		if rawIndex == "shit" {
			fmt.Println("Sucks to suck. Play something.")
			continue
		}
		if err != nil {
			fmt.Println(rawIndex, "is not a valid number")
			continue
		}	else if index < 1 || index > 9 {
			fmt.Println("The number must be between 1 and 9")
			continue
		} else if board[index - 1] != -1 {
			fmt.Println("Space", index, "is taken. Choose another.")
			continue
		}

		
		board[index - 1] = player
		
		winner = determineWinner(board)
		
		if winner != -1 {
			fmt.Println("Player", getPlayerChar(player), "wins!")
			break
		}
		isTie := determineTie(board)
		if isTie {
			fmt.Println("It's a tie!")
			break
		}
		player = (player + 1) % 2
	}

	var playAgain string
	fmt.Print("Play again? [Y/n]")
	fmt.Scanln(&playAgain)
	
	if playAgain == "y" || playAgain == "Y" {
		main()
	}
}