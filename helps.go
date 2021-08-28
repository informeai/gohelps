package gohelps

//NewString return new instance string.
func NewString(s string) *String {
	return &String{base: s}
}

//NewBool return instance of Bool.
func NewBool(val bool) *Bool {
	return &Bool{base: val}
}
