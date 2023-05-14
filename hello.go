package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gofhir/generated"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//go:generate go run generator/main.go

func main() {
	file, err := os.Open("assets/test/Observation.json")
	check(err)
	defer file.Close()

	// Unmarshal the JSON
	var base generated.Observation
	err := json.NewDecoder(file).Decode(&base)
	check(err)

	fmt.Println(base)

	fmt.Println("Hello, World!")
}
