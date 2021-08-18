package gohelps

import (
	"errors"
	"strings"
)

//String is struct base for methods the object string.
type String struct {
	base string
}

//var errors
var (
	ErrIndexOutOfRange = errors.New("index out of range")
	ErrEmptyValue      = errors.New("this empty value")
	ErrNotFound        = errors.New("not found value")
	ErrNotPermited     = errors.New("not permited value")
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

//Includes verify if exists other string in actual string.
func (s String) Includes(str string) (bool, error) {
	if len(str) == 0 {
		return false, ErrEmptyValue
	}
	return strings.Contains(s.base, str), nil

}

//IndexOf verify if exists other string e return initial index.
func (s String) IndexOf(str string) (int, error) {
	index := strings.Index(s.base, str)
	if index == -1 {
		return -1, ErrNotFound
	}
	return index, nil
}

//LastIndexOf verify if exists other string e return final index of string.
func (s String) LastIndexOf(str string) (int, error) {
	index := strings.LastIndex(s.base, str)
	if index == -1 {
		return -1, ErrNotFound
	}
	return index, nil
}

//PadEnd return pads the current string with a given string repeated the final.
func (s String) PadEnd(count int, str string) (string, error) {
	if count <= 0 {
		return "", ErrNotPermited
	}
	for i := 0; i < count; i++ {
		s.base += str
	}
	return s.base, nil
}

//PadStart return pads the current string with a given string repeated the initial.
func (s String) PadStart(count int, str string) (string, error) {
	if count <= 0 {
		return "", ErrNotPermited
	}
	ns := ""
	for i := 0; i < count; i++ {
		ns += str
	}
	return ns + s.base, nil
}

//Repeat return new string repeated many times.
func (s String) Repeat(count int) (string, error) {
	if count <= 0 {
		return "", ErrNotPermited
	}
	ns := s.base
	for i := 0; i < count; i++ {
		ns += s.base
	}
	return ns, nil
}

//Replace return new string with replaced first ocorrency the string.
func (s String) Replace(str, old string) (string, error) {
	if len(str) == 0 {
		return "", ErrEmptyValue
	}
	return strings.Replace(s.base, str, old, 1), nil
}
