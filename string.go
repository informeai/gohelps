package gohelps

import (
	"errors"
)

//String is struct base for methods the object string.
type String struct {
	base string
}

//var errors
var (
	ErrIndexOutOfRange = errors.New("index out of range.")
	ErrEmptyValue      = errors.New("this empty value.")
)

//CharAt return of rune by index.
func (s String) CharAt(i int) (string, error) {
	if i > len(s.base)-1 || i < 0 {
		return "", ErrIndexOutOfRange
	}
	return string(s.base[i]), nil
}

//CharCodeAt return of byte by index.
func (s String) CharCodeAt(i int) (byte, error) {
	if i > len(s.base)-1 || i < 0 {
		return byte(0), ErrIndexOutOfRange
	}
	return s.base[i], nil
}

//Concat return of union the two string.
func (s String) Concat(str string) (string, error) {
	if len(str) == 0 {
		return "", ErrEmptyValue
	}
	return s.base + str, nil
}

//EndsWith return the bool if ends word match of string.
func (s String) EndsWith(str string) (bool, error) {
	if len(str) == 0 {
		return false, ErrEmptyValue
	}
	lastIndex := len(s.base)
	return s.base[lastIndex-len(str):] == str, nil
}
