package main

import "fmt"

func grade(score int) string {
	g := ""
	switch {
	case score < 60 && score > 0:
		g = "F"
	case score >= 60:
		g = "A"
	default:
		panic("wrong sorce")
	}
	return g
}

func main() {
	//const filename = "abc.txt"
	//bytes, err := ioutil.ReadFile(filename)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Printf("%s\n", bytes)
	//}
	//
	//
	//if file, err1 := ioutil.ReadFile(filename); err1 != nil {
	//	fmt.Println(err1)
	//} else {
	//	fmt.Printf("%s\n", file)
	//}

	fmt.Print(grade(60), grade(10),grade(10),grade(-1))
}
