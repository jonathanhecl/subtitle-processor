package main

import (
	"fmt"

	"github.com/jonathanhecl/subtitle-processor/subtitles"
)

var version = map[string]int{
	"major": 1,
	"minor": 1,
}

func main() {
	fmt.Println(fmt.Sprintf("Subtitle Processor v%d.%d", version["major"], version["minor"]))

	s1 := subtitles.Subtitle{}
	s1.LoadFilename("./demo.srt")
	fmt.Println(s1.Filename)
	fmt.Println(s1.Format)
	fmt.Println("Lines: ", len(s1.Lines))
	/*
		for i := range s1.Lines {
			fmt.Println("Seq: ", s1.Lines[i].Seq)
			fmt.Println("Start: ", s1.Lines[i].Start)
			fmt.Println("End: ", s1.Lines[i].End)
			fmt.Println("Text: ", s1.Lines[i].Text)
		}
	*/

	//s2 := subtitle.LoadFilename("./demo.ssa")

}
