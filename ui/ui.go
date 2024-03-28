package ui

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"slices"

	"github.com/fatih/color"
)
var (
	StatusCorrectPosition = 1
	StatusIncorrectPosition = 2
	StatusWrongLetter = 3
)
var alphabet = []rune("abcdefghijklmnopqrstuvwxyz")

type Char struct {
	Value rune 
	Color color.Attribute
}
type LetterGrid [6][5]Char
type LetterStatus map[rune]int

type UI struct {}

func New() *UI{
	return &UI{}
}
func (ui * UI) Render(wordToGuess string, guesses []string){
	ui.clearTerminal()
	title := `
	__        _____  ____  ____  _     _____ 
	\ \      / / _ \|  _ \|  _ \| |   | ____|
	 \ \ /\ / / | | | |_) | | | | |   |  _|  
	  \ V  V /| |_| |  _ <| |_| | |___| |___ 
	   \_/\_/  \___/|_| \_\____/|_____|_____|
	   
	`
	fmt.Println(title)
	const EmptyRune = 0
	grid, status  := ui.constructGridAndLetterStatus(wordToGuess, guesses)
	for _, row := range grid {
		for _, col := range row{
			if col.Value == EmptyRune {
				fmt.Printf("%c ", '*')
				continue
			}
			clr := color.New(col.Color)
			clr.Printf("%c ", col.Value)
	
		}
		fmt.Println()
	}
	fmt.Println("__________")
	for _, a := range alphabet{
		s, exists := status[a]
		if !exists {
			fmt.Printf("%c ", a)
			continue
		}
		clr := color.New(color.FgGreen)
			if s == StatusCorrectPosition {
			clr.Printf("%c ", a)
			continue
		}
		if s == StatusIncorrectPosition {
			clr := color.New(color.FgYellow)
			clr.Printf("%c ", a)
			continue
		}
		if s == StatusWrongLetter {
			clr := color.New(color.FgRed)
			clr.Printf("%c ", a)
			continue
		}
		
	
	} 
	fmt.Println()
}
func (ui *UI) constructGridAndLetterStatus(wordToGuess string, guesses []string)(LetterGrid, LetterStatus){
	correctWordRune := []rune(wordToGuess)
	grid := LetterGrid{}
	status := LetterStatus{}
	for parentIdx, guess := range guesses {
		for idx, c := range guess {
			char := Char{
				Value: c,
				Color: color.FgWhite,
			}
			if c == correctWordRune[idx] {
				char.Color = color.FgGreen
				ui.insertOrUpdateLetterStatus(c, status, StatusCorrectPosition)
			}else if slices.Contains(correctWordRune, c) {
				char.Color = color.FgYellow
				ui.insertOrUpdateLetterStatus(c, status, StatusIncorrectPosition)
			}else  {
				ui.insertOrUpdateLetterStatus(c, status, StatusWrongLetter)
				char.Color = color.FgRed
			}			
			grid[parentIdx][idx] = char
		}
	}
	return grid, status
}
func (ui *UI)insertOrUpdateLetterStatus(c rune, letterStatus LetterStatus, newStatus int) {
	status, exists := letterStatus[c]
	if !exists {
		letterStatus[c] = newStatus
		return
	}
	if status == StatusCorrectPosition && exists {
		return
	}
	if status == StatusIncorrectPosition && exists {
		letterStatus[c] = newStatus
		return
	}
	letterStatus[c] = StatusWrongLetter
}

func(ui * UI) runCmd(name string, arg ...string) {
    cmd := exec.Command(name, arg...)
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func(ui * UI) clearTerminal() {
    switch runtime.GOOS {
    case "darwin":
        ui.runCmd("clear")
    case "linux":
        ui.runCmd("clear")
    case "windows":
        ui.runCmd("cmd", "/c", "cls")
    default:
        ui.runCmd("clear")
    }
}