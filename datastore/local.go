package datastore

import (
	"fmt"
	"unicode"
)

type Local struct {
	Client map[rune]string
}

func (l *Local) Open(uri string) (err error) {
	if uri == "" {
		err = fmt.Errorf("missing uri")
	} else {
		l.Client = map[rune]string{
			rune('A'):  "Alpha",
			rune('B'):  "Bravo",
			rune('C'):  "Charlie",
			rune('D'):  "Delta",
			rune('E'):  "Echo",
			rune('F'):  "Foxtrot",
			rune('G'):  "Golf",
			rune('H'):  "Hotel",
			rune('I'):  "India",
			rune('J'):  "Juliett",
			rune('K'):  "Kilo",
			rune('L'):  "Lima",
			rune('M'):  "Mike",
			rune('N'):  "November",
			rune('O'):  "Oscar",
			rune('P'):  "Papa",
			rune('Q'):  "Quebec",
			rune('R'):  "Romeo",
			rune('S'):  "Sierra",
			rune('T'):  "Tango",
			rune('U'):  "Uniform",
			rune('V'):  "Victor",
			rune('W'):  "Whiskey",
			rune('X'):  "XRay",
			rune('Y'):  "Yankee",
			rune('Z'):  "Zulu",
			rune('1'):  "Number One",
			rune('2'):  "Number Two",
			rune('3'):  "Number Three",
			rune('4'):  "Number Four",
			rune('5'):  "Number Five",
			rune('6'):  "Number Six",
			rune('7'):  "Number Seven",
			rune('8'):  "Number Eight",
			rune('9'):  "Number Nine",
			rune('0'):  "Number Zero",
			rune('.'):  "Stop",
			rune(','):  "Comma",
			rune('-'):  "Hyphen",
			rune('/'):  "Slant",
			rune('\\'): "Fraction bar",
			rune('('):  "Brackets on",
			rune(')'):  "Brackets off",
			rune(':'):  "Colon",
			rune(';'):  "Semi-colon",
			rune('!'):  "Exclamation mark",
			rune('?'):  "Question mark",
			rune('\''): "Apostrophe",
			rune('"'):  "Quote",
		}
	}
	return
}

func (l *Local) Read(code string) (word []string, err error) {
	for _, c := range code {
		word = append(word, l.Client[unicode.ToUpper(c)])
	}
	return
}
