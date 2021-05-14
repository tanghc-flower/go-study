package main

import "fmt"

var aa = 3
var bb = 1
var cc = true

var(
	qq = 1
	ww = "123asd"
)

func variable()  {
	var a, c int = 1, 2
	var b string = "asd"
	fmt.Println(a, c, b)
}

func variableShorter()  {
	a, c := 1, 2
	var b string = "asd"
	fmt.Println(a, c, b)
}

func consts()  {
	const filename = "file.txt"
	const a, b  = 1, 2
}

func enums()  {
	const (
		java = iota
		_
		python = 2
		goland = 3
	)
	fmt.Println(java, python, goland)
}

func main() {
	fmt.Println("Hello world")
	variable()
	variableShorter()
	fmt.Println(aa, bb, cc)
	enums()
}


