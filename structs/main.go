package main

import (
	"fmt"
	"structs"
)

type Example struct {
	Content string
	Value   int
	_       structs.HostLayout
}

func main() {
	e := Example{
		Content: "Example with structs",
		Value:   1,
	}

	fmt.Println("Example", e)
	fmt.Printf("Specific example %#v\n", e)
}
