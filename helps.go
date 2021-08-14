package gohelps

//String return new object string.
func NewString(s string) *String {
	return &String{base: s}
}
