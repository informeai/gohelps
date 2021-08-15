package gohelps

import (
	"log"
	"testing"
)

//go test -v -run ^TestCharAt
func TestCharAt(t *testing.T) {
	str := "informeai"
	s := NewString(str)
	ru, err := s.CharAt(3)
	if err != nil {
		t.Errorf("TestCharAt(): got -> %v, want: nil", err)
	}
	log.Println(ru)
}

//go test -v -run ^TestCharCodeAt
func TestCharCodeAt(t *testing.T) {
	str := "informeai"
	s := NewString(str)
	bs, err := s.CharCodeAt(3)
	if err != nil {
		t.Errorf("TestCharCodeAt(): got -> %v, want: nil", err)
	}
	log.Println(bs)
}

//go test -v -run ^TestConcat
func TestConcat(t *testing.T) {
	str := "hello"
	s := NewString(str)
	cs, err := s.Concat(", word!")
	if err != nil {
		t.Errorf("TestConcat(): got -> %v, want: nil", err)
	}
	log.Println(cs)
}

//go test -v -run ^TestEndsWith
func TestEndsWith(t *testing.T) {
	str := "I'm sorry doctor gopher."
	s := NewString(str)
	es, err := s.EndsWith("gopher.")
	if err != nil {
		t.Errorf("TestEndsWith(): got -> %v, want: nil", err)
	}
	log.Println(es)
}
