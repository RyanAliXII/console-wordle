package dictionary

import (
	"bufio"
	_ "embed"
	"math/rand"
	"strings"
)

//go:embed data.txt
var data []byte

type Dictionary struct {
	words []string
	wordsMap map[string]struct{}
}
func New() *Dictionary{
	d := Dictionary{
		words: make([]string, 0),
		wordsMap: make(map[string]struct{}, 0),
	}
	d.load()
	return &d
}
func (d *Dictionary) load() {
	fileScanner := bufio.NewScanner(strings.NewReader(string(data)))
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		word := strings.TrimSpace(fileScanner.Text())
		d.words = append(d.words, word)
		d.wordsMap[word] = struct{}{}
	}
}
func (d * Dictionary)DoesWordExists(word string) bool{
	_, exists := d.wordsMap[word]
	return exists
}
func(d * Dictionary)GetRandomWord() string {
	max := len(d.words) - 1
	min := 0
	n := rand.Intn(max - min) + min
	return d.words[n]
}
