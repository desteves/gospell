package datastore

import "unicode"

type Mock struct {
	Client map[rune]string
}

func (m *Mock) Open(uri string) (err error) {
	m.Client = map[rune]string{
		rune('A'): "Apple",
		rune('B'): "Banana",
		rune('C'): "Coconut",
	}
	return
}

func (m *Mock) Read(code string) (word []string, err error) {
	for _, c := range code {
		word = append(word, m.Client[unicode.ToUpper(c)])
	}
	return
}
