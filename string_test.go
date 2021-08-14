package gohelps

import (
	"log"
	"testing"
)

//go test -v -run ^TestCharAt
func TestCharAt(t *testing.T) {
	str := "informeai"
	s := NewString(str)
	ru := s.CharAt(3)
	if ru == "" {
		t.Errorf("TestCharAt(): got -> %v, want: any string", ru)
	}
	log.Println(ru)
}
