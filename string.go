package gohelps

//String is struct base for methods the object string.
type String struct {
	base string
}

//CharAt return of rune by index.
func (s String) CharAt(i int) string {
	if i > len(s.base)-1 || i < 0 {
		return ""
	}
	return string(s.base[i])
}
