package gohelps

import (
	"strconv"
)

//Bool is struct of base methods booleans.
type Bool struct {
	base bool
}

//ToString return of string from bool.
func (b *Bool) ToString() string {
	return strconv.FormatBool(b.base)
}
