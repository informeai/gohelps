# :smile:  GoHelps
GoHelps is intended to provide a set of methods and utilities for the go language.
Based on the concept of primitive wrapper object in javascript, providing common methods to these objects,
but with the go of being. It's not javascript but golang with a little more.:+1:

### :wrench:  Install
Get package
```
go get github.com/informeai/gohelps
```
### :fire:  Usage
Create new string object.
```
str := "hello gopher!"
s := gohelps.NewString(str)
```
Use the many methods.
```
s, _ := s.Concat(" I'ts ok...")
//s -> hello gopher! I'ts ok...
```
more.
```
c,_ := s.CharAt(3)
//c -> 'l'
```
by **wellington gadelha** :punch:.
