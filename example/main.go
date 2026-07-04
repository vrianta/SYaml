package main

import syml "github.com/vrianta/SYml/v1"

type test struct {
	name string
	city string
}

func main() {
	t := test{}
	syml.Unmarshal(
		[]byte(`
    name: Joy 
    city: Kolkata
	testObj:
		test1: 1
		test2: 2
	st1: "our test to test lexer" #command
`), &t)
}
