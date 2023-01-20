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
		userprompt("CORRECT")
	} else {
		guesses = append(guesses, guess)
		dispGuesses()
	}
	//	fmt.Println()
}

func dispGuesses() {
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

func playGame() {
	fmt.Print("\033[H\033[2J")
	for i := 0; i < 6; i++ {
		userquery("GUESS")
		var guess string
		fmt.Scanln(&guess)
		if len(guess) != 5 {
			fmt.Print("\033[H\033[2J")
			userprompt("GUESS MUST BE 5 LETTERS LONG")
			dispGuesses()
      i -= 1
			continue
		}
		if !isValidGuess(guess) {
			fmt.Print("\033[H\033[2J")
			userprompt("INVALID GUESS")
			dispGuesses()
      i -= 1
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
	userqueryml("PLAY AGAIN", "yes OR no")
	fmt.Scanln(&playAgain)
	if playAgain == "yes" {
		generatePuzzle()
		guesses = []string{}
		playGame()
	}
}

func isValidGuess(guess string) bool {
	for _, w := range words {
		if guess == w {
			return true
		}
	}
	return false
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

//Functions to handle displaying the impractical user interface boxes

func userprompt(content string) {
	fmt.Println("┏" + strings.Repeat("━", len(content)+2) + "┓")
	fmt.Println("┃ " + content + " ┃")
	fmt.Println("┗" + strings.Repeat("━", len(content)+2) + "┛")
}
func userquery(content string) {
	fmt.Println("┏" + strings.Repeat("━", len(content)+2) + "┓")
	fmt.Println("┃ " + content + " ┃")
	fmt.Println("┣" + strings.Repeat("━", len(content)+2) + "┛")
	fmt.Print("┃ ")
}
func userqueryml(content string, contentext string) {
	if len(contentext) == len(content) {
		fmt.Println("┏" + strings.Repeat("━", len(content)+2) + "┓")
		fmt.Println("┃ " + content + " ┃")
		fmt.Println("┃ " + contentext + " ┃")
		fmt.Println("┣" + strings.Repeat("━", len(content)+2) + "┛")
		fmt.Print("┃ ")
	} else {
		if len(contentext) > len(content) {
			targlen := len(contentext)
			remainderlen := len(contentext) - len(content)
			fmt.Println("┏" + strings.Repeat("━", targlen+2) + "┓")
			fmt.Println("┃ " + content + strings.Repeat(" ", remainderlen) + " ┃")
			fmt.Println("┃ " + contentext + " ┃")
			fmt.Println("┣" + strings.Repeat("━", targlen+2) + "┛")
			fmt.Print("┃ ")
		} else {
			targlen := len(content)
			remainderlen := len(content) - len(contentext)
			fmt.Println("┏" + strings.Repeat("━", targlen+2) + "┓")
			fmt.Println("┃ " + content + " ┃")
			fmt.Println("┃ " + contentext + strings.Repeat(" ", remainderlen) + " ┃")
			fmt.Println("┣" + strings.Repeat("━", targlen+2) + "┛")
			fmt.Print("┃ ")
		}
	}
}
