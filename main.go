package main

import (
	"console-wordle/dictionary"
	"console-wordle/gamemanager"
	"console-wordle/ui"
)
func main() {
	var gm = gamemanager.New(dictionary.New(), ui.New())
	gm.Start()
}

