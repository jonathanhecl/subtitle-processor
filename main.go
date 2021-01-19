package main

import (
	"fmt"

	"github.com/jonathanhecl/subtitle-processor/subtitles"
)

var version = map[string]int{
	"major": 1,
	"minor": 0,
}

func main() {
	fmt.Println(fmt.Sprintf("Subtitle Processor v%d.%d", version["major"], version["minor"]))

	s1 := subtitles.LoadFilename("./demo.srt")
	fmt.Println(s1)

	//s2 := subtitle.LoadFilename("./demo.ssa")

}
