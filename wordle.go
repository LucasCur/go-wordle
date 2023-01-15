package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

var (
	words       = []string{} // words to guess from
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
		fmt.Println("┏━━━━━━━━━┓")
		fmt.Println("┃ CORRECT ┃")
		fmt.Println("┗━━━━━━━━━┛")
	} else {
		guesses = append(guesses, guess)
		if len(guesses) > 0 {
			fmt.Print("\033[H\033[2J")
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
			fmt.Println("┗━━━━━━━┛")
		}
	}
	//	fmt.Println()
}

func playGame() {
	fmt.Print("\033[H\033[2J")
	for i := 0; i < 6; i++ {
		fmt.Println("┏━━━━━━━┓")
		fmt.Println("┃ GUESS ┃")
		fmt.Println("┣━━━━━━━┛")
		fmt.Print("┃ ")
		var guess string
		fmt.Scanln(&guess)
		if len(guess) != 5 {
			fmt.Print("\033[H\033[2J")
			fmt.Println("┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
			fmt.Println("┃ Guess must be 5 letters long ┃")
			fmt.Println("┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
			continue
		}
		checkGuess(guess)
		if guess == correctWord {
			break
		}
		if i == 5 {
			fmt.Println("┏━━━━━━━━━━┓")
			fmt.Println("┃ YOU LOST ┃")
			fmt.Println("┣━━━━━━━┳━━┛")
			fmt.Println("┃ " + correctWord + " ┃")
			fmt.Println("┗━━━━━━━┛")
			break
		}
	}
	var playAgain string
	fmt.Println("┏━━━━━━━━━━━━┓")
	fmt.Println("┃ PLAY AGAIN ┃")
	fmt.Println("┃ yes OR no  ┃")
	fmt.Println("┣━━━━━━━━━━━━┛")
	fmt.Print("┃ ")
	fmt.Scanln(&playAgain)
	if playAgain == "yes" {
		generatePuzzle()
		guesses = []string{}
		playGame()
	}
}

func main() {
	file, err := ioutil.ReadFile("words.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := json.Unmarshal(file, &words); err != nil {
		fmt.Println(err)
		return
	}
	generatePuzzle()
	playGame()
}
