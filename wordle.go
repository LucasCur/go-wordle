package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	words       = []string{"help", "test", "true"}
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
		if len(guesses) > 0 {
			fmt.Println("┏━━━━━━━━━━━━━┓")
			fmt.Println("┃ OLD GUESSES ┃")
			fmt.Println("┣━━━━━━━┳━━━━━┛")
		}
		for _, g := range guesses {
			var word string
			for i := range g {
				if i < len(correctWord) && g[i] == correctWord[i] {
					word += "\033[32m" + strings.ToUpper(string(g[i])) + "\033[0m"
				} else if strings.ContainsRune(correctWord, rune(g[i])) {
					word += "\033[33m" + strings.ToUpper(string(g[i])) + "\033[0m"
				} else {
					word += strings.ToUpper(string(g[i]))
				}
			}
			fmt.Println("┃ " + word + " ┃")
		}
		if len(guesses) > 0 {
			fmt.Println("┗━━━━━━━┛\n")
		}
	}
	//	fmt.Println()
}

func playGame() {
	for i := 0; i < 6; i++ {
		fmt.Print("\033[H\033[2J")
		fmt.Print("┃ Guess a word: ")
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
	fmt.Print("Do you want to play again? (yes or no)")
	fmt.Scanln(&playAgain)
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
