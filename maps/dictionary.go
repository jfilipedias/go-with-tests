package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you are looking for")

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}