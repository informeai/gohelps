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
