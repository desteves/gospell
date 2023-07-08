package speller

import "unicode"

type Mock struct {
	CodeWords map[rune]string
}

func NewMockProvider(uri string) (*Mock, error) {
	return &Mock{CodeWords: map[rune]string{
		rune('A'): "Apple",
		rune('B'): "Banana",
		rune('C'): "Coconut",
	}}, nil

}

func (m *Mock) GetCodeWords(input string) (word []string, err error) {

	for _, c := range input {
		word = append(word, m.CodeWords[unicode.ToUpper(c)])
	}
	return
}
