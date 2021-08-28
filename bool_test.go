package gohelps

import (
	"testing"
)

//go test -v -run ^TestBoolToString
func TestBoolToString(t *testing.T) {
	val := true
	b := NewBool(val)
	str := b.ToString()
	if str == "" {
		t.Errorf("TestBoolToString(): got -> %v, want: != nil", str)
	}
}
