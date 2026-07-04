package main

import (
	"os"

	syml "github.com/vrianta/SYml/v1"
)

var testSYaml = `
name: Joy 
city: Kolkata
age: 30
testObj:
	test1: 1
	test2: 2

st1: "our test to test lexer" #command
`

var testObj = `
testObj:
	test1: 1
	test2: 2
`

type test struct {
	name string
	city string
}

func main() {

	data, err := os.ReadFile("stringtest.yml")
	if err != nil {
		panic(err)
	}
	t := test{}
	syml.Unmarshal(data, &t)
}
