package mydict

import "errors"

type Dictionary map[string]string // Dictionary는 map[string]string에 대한 alias(가명)이다.
// type에 대해서도 alias를 만들 수 있다.
// type에도 struct처럼 method를 추가할 수 있다.

var errNotFound = errors.New("not found")
var errCantUpdate = errors.New("can't update non-existing word")
var errWordExists = errors.New("that word already exists")

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// map은 key의 존재여부를 알려주는 방법이 있다.
// i, ok := m["route"] -> map에서 "route"를 호출하면 i(value)와 존재여부를 알려주는 ok(boolean)을 반환한다.

// Add a word to the dictionary
func (d Dictionary) Add(word string, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}

// Update a word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil: // 단어가 이미 존재한다는 뜻
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
