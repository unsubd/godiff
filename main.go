package main

import "fmt"

func main() {
	second := []byte("HELLOKORLD")
	first := []byte("HELLOWORLD")
	indices := Diff(first, second)
	secondStr := ""

	for i := 0; i < len(first) && i < len(second); i++ {
		ok := indices[i]
		if ok {
			secondStr += Color(string(second[i]))
		} else {
			secondStr += string(second[i])
		}
	}

	fmt.Print(string(first), "\t")
	fmt.Print(secondStr)
}
