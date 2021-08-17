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

//go test -v -run ^TestIncludes
func TestIncludes(t *testing.T) {
	str := "hello my friend."
	s := NewString(str)
	is, err := s.Includes("hello")
	if err != nil {
		t.Errorf("TestIncludes(): got -> %v, want: nil", err)
	}
	log.Println(is)
}

//go test -v -run ^TestIndexOf
func TestIndexOf(t *testing.T) {
	str := "hello my friend."
	s := NewString(str)
	is, err := s.IndexOf("my")
	if err != nil {
		t.Errorf("TestIndexOf(): got -> %v, want: nil", err)
	}
	log.Println(is)
}

//go test -v -run ^TestLastIndexOf
func TestLastIndexOf(t *testing.T) {
	str := "hello my friend. this is my dog?"
	s := NewString(str)
	is, err := s.LastIndexOf("my")
	if err != nil {
		t.Errorf("TestLastIndexOf(): got -> %v, want: nil", err)
	}
	log.Println(is)
}

//go test -v -run ^TestPadEnd
func TestPadEnd(t *testing.T) {
	str := "My name is"
	s := NewString(str)
	ns, err := s.PadEnd(5, ".")
	if err != nil {
		t.Errorf("TestPadEnd(): got -> %v, want: nil", err)
	}
	log.Println(ns)
}

//go test -v -run ^TestPadStart
func TestPadStart(t *testing.T) {
	str := ">next level"
	s := NewString(str)
	ns, err := s.PadStart(5, "-")
	if err != nil {
		t.Errorf("TestPadStart(): got -> %v, want: nil", err)
	}
	log.Println(ns)
}

//go test -v -run ^TestRepeat
func TestRepeat(t *testing.T) {
	str := "hello, my friend. "
	s := NewString(str)
	ns, err := s.Repeat(3)
	if err != nil {
		t.Errorf("TestRepeat(): got -> %v, want: nil", err)
	}
	log.Println(ns)
}
