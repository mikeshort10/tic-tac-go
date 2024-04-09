package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"example.com/game"
)

var Reset = "\033[0m"
var Yellow = "\033[33m"
var Purple = "\033[35m"
var White = "\033[97m"
var Red = "\033[31m"

func getPlayerChar(player int) (string, string) {
	if player == 0 {
		return "😇", Yellow
	} else if player == 1 {
		return "😈", Purple
	}
	return "  ", White
}

func renderBoard(board game.Board) string {
	layout := ""
	for i := 0; i < 9; i += 3 {
		row := board[i : i+3]
		layout += " "
		for j := 0; j < 3; j++ {
			tic, _ := getPlayerChar(row[j])
			layout += tic
			if j == 2 {
				layout += " "
			} else {
				layout += " | "
			}
		}
		layout += "\n"
	}

	return layout
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	player := 0
	winner := -1
	board := game.NewBoard()

	clear()
	for winner == -1 {
		var rawIndex string

		playerChar, color := getPlayerChar(player)

		fmt.Println(renderBoard(board))
		fmt.Print(color, "Player ", playerChar, " enter a space to play: ", Reset)
		fmt.Scanln(&rawIndex)
		clear()

		if rawIndex == "exit" {
			return
		}

		index, err := strconv.Atoi(rawIndex)

		if err != nil {
			fmt.Println(Red, rawIndex, "is not a valid number", Reset)
			continue
		} else if index < 1 || index > 9 {
			fmt.Println(Red, "The number must be between 1 and 9", Reset)
			continue
		}

		var success bool
		board, success = game.UpdateBoard(board, index, player)

		if !success {
			fmt.Println(Red, "Space", index, "is taken. Choose another.", Reset)
			continue
		}

		winner = game.DetermineWinner(board)

		if winner != -1 {
			fmt.Println(color, "Player", playerChar, "wins!", Reset)
			break
		} else if game.DetermineTie(board) {
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
