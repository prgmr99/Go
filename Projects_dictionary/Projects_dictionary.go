package mydict

import "errors"

type Dictionary map[string]string // Dictionary는 map[string]string에 대한 alias(가명)이다.
// type에 대해서도 alias를 만들 수 있다.
// type에도 struct처럼 method를 추가할 수 있다.

var errNotFound = errors.New("not found")

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
