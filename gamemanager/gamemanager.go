package gamemanager

import (
	"fmt"
	"strings"
	"time"
)

type GMUIManager interface {
	Render(string, []string)
}
type GMDictionary interface {
	DoesWordExists(word string) bool
	GetRandomWord() string
}
type GameManager struct {
	guesses     []string
	letters     map[string]struct{}
	wordToGuess string
	Dictionary  GMDictionary
	UIManager   GMUIManager
}

func New(d GMDictionary, ui GMUIManager) GameManager {
	return GameManager{
		guesses:    make([]string, 0),
		letters:    make(map[string]struct{}, 0),
		Dictionary: d,
		UIManager:  ui,
	}
}
func (gm *GameManager) Start() {
	MaxRows := 6
	MaxCols := 5
	tries := 0
	wordToGuess := gm.GuessAWord()
	for tries < MaxRows {
		gm.UIManager.Render(wordToGuess, gm.GetGuesses())
		guess := ""
		fmt.Print("Guess the word: ")
		fmt.Scanln(&guess)
		guess = strings.ToLower(guess)
		if guess == "" {
			fmt.Println("Input cannot be empty.")
			time.Sleep(time.Second * 1)
			continue
		}
		if len(guess) != MaxCols {
			fmt.Println("Input must be 5 characters.")
			time.Sleep(time.Second * 1)
			continue
		}
		if !gm.Dictionary.DoesWordExists(guess) {
			fmt.Println("Word doesn't exists in dictionary.")
			time.Sleep(time.Second * 1)
			continue
		}
		gm.Guess(guess)
		if guess == wordToGuess {
			fmt.Println("You have guessed the word correctly.")
			tries = 5
		}
		tries++
	}
	fmt.Printf("Correct work is %s", wordToGuess)
	fmt.Scan()
}
func (gm *GameManager) Guess(input string) {
	gm.guesses = append(gm.guesses, input)
}
func (gm *GameManager) SetWordToGuess(word string) {
	gm.wordToGuess = word
}
func (gm *GameManager) GetWordToGuess() string {
	return gm.wordToGuess
}
func (gm *GameManager) GetGuesses() []string {
	return gm.guesses
}
func (gm *GameManager) GuessAWord() string {
	gm.wordToGuess = gm.Dictionary.GetRandomWord()
	return gm.wordToGuess
}
