// Package speller contains the interface and implementations to retrieve the code words
package speller

type Speller interface {
	GetCodeWords(input string) ([]string, error)
}
