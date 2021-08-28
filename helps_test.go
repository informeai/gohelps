package gohelps

import (
	"log"
	"testing"
)

//go test -v -run ^TestNewString
func TestNewString(t *testing.T) {
	h := NewString("hello string :)")
	if h == nil {
		t.Errorf("TestNewString(): got -> %v, want: != nil", h)
	}
	log.Println(h)
}

//go test -v -run ^TestNewBool
func TestNewBool(t *testing.T) {
	b := NewBool(true)
	if b == nil {
		t.Errorf("TestNewBool(): got -> %v, want: != nil", b)
	}
}
