package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/fatih/color"
)

var (
	words       = []string{"aaa", "bbb", "ccc", "ddd", "eee", "fff"}
	correctWord = ""
	guesses     = []string{}
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generatePuzzle() {
	correctWord = words[rand.Intn(len(words))]
}

func checkGuess(guess string) {
	if guess == correctWord {
		fmt.Println("Correct! You won!")
	} else {
		guesses = append(guesses, guess)
		for _, g := range guesses {
			var word string
			for i := range g {
				if i < len(correctWord) && g[i] == correctWord[i] {
					word += color.GreenString("%c", g[i])
				} else if strings.ContainsRune(correctWord, rune(g[i])) {
					word += color.YellowString("%c", g[i])
				} else {
					word += fmt.Sprintf("%c", g[i])
				}
			}
			fmt.Println(strings.ToUpper(word))
		}
	}
	fmt.Println()
}

func playGame() {
	color.NoColor = false
	for i := 0; i < 6; i++ {
		fmt.Print("\033[H\033[2J")
		for _, g := range guesses {
			var word string
			for i := range g {
				if i < len(correctWord) && g[i] == correctWord[i] {
					word += color.GreenString("%c", g[i])
				} else if strings.ContainsRune(correctWord, rune(g[i])) {
					word += color.YellowString("%c", g[i])
				} else {
					word += fmt.Sprintf("%c", g[i])
				}
			}
			fmt.Println(strings.ToUpper(word))
		}
		fmt.Print("Guess a word: ")
		var guess string
		fmt.Scanln(&guess)
		checkGuess(guess)
		if guess == correctWord {
			break
		}
		if i == 5 {
			fmt.Printf("Sorry, you lost. The word was %s\n", correctWord)
			break
		}
	}
	var playAgain string
	for {
		fmt.Print("Do you want to play again? (yes or no)")
		fmt.Scanln(&playAgain)
		if playAgain == "yes" || playAgain == "no" {
			break
		}
		fmt.Println("Invalid input. Please enter either 'yes' or 'no'.")
	}
	if playAgain == "yes" {
		generatePuzzle()
		guesses = []string{}
		playGame()
	}
}

func main() {
	generatePuzzle()
	playGame()
}
