package main

import "fmt"

var version = map[string]int{
	"major": 1,
	"minor": 0,
}

func main() {
	fmt.Println(fmt.Sprintf("Subtitle Processor v%d.%d", version["major"], version["minor"]))
}
