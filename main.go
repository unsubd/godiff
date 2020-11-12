package main

import "fmt"

func main() {
	first := []byte("HELLOWORLD")
	second := []byte("HELLOKORLD")
	indices := Diff(first, second)
	secondStr := ""
	last := 0
	for i := 0; i < len(first) && i < len(second); i++ {
		ok := indices[i]
		if ok {
			secondStr += Color(string(second[i]))
		} else {
			secondStr += string(second[i])
		}

		last = i
	}

	if last < len(second)-1 {
		secondStr += Color(string(second[last:]))
	}
	fmt.Print(string(first), "\t")
	fmt.Print(secondStr)
}
