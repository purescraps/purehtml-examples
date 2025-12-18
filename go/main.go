package main

import (
	"fmt"
	"log"

	purehtml "github.com/purescraps/purehtml/go"
)

func main() {
	html := `<html><body><h1 class="title">Hello World</h1></body></html>`

	config := &purehtml.PrimitiveValueConfig{
		Selector:  []string{".title"},
		Transform: []purehtml.TransformerSpec{{Name: "trim"}},
	}

	result, err := purehtml.ExtractFromString(purehtml.DefaultBackend, html, config, "")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result) // Output: Hello World
}
