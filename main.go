// Package main provides a demonstration of the subtitle-processor library.
package main

import (
	"fmt"

	"github.com/jonathanhecl/subtitle-processor/subtitles"
)

// Version information
var version = map[string]int{
	"major": 1,
	"minor": 0,
}

func main() {
	// Display version information
	fmt.Printf("Subtitle Processor v%d.%d\n", version["major"], version["minor"])

	// Example 1: Loading and displaying SRT subtitle information
	fmt.Println("\n=== SRT Subtitle Example ===")
	srtSubtitle := subtitles.Subtitle{}
	srtSubtitle.Verbose = true
	srtSubtitle.LoadFile("./demo.srt")

	fmt.Println("Filename:", srtSubtitle.Filename)
	fmt.Println("Format:", srtSubtitle.Format)
	fmt.Println("Lines:", len(srtSubtitle.Lines))

	// Display subtitle content details
	for i := range srtSubtitle.Lines {
		fmt.Println("\nSubtitle Entry", i+1)
		fmt.Println("Sequence:", srtSubtitle.Lines[i].Seq)
		fmt.Println("Start Time:", srtSubtitle.Lines[i].Start)
		fmt.Println("End Time:", srtSubtitle.Lines[i].End)
		fmt.Println("Text:", srtSubtitle.Lines[i].Text, "Length:", len(srtSubtitle.Lines[i].Text))
	}

	// Example 2: Loading and displaying SSA subtitle information
	fmt.Println("\n=== SSA Subtitle Example ===")
	ssaSubtitle := subtitles.Subtitle{}
	ssaSubtitle.Verbose = true
	ssaSubtitle.LoadFile("./demo.ssa")

	fmt.Println("Filename:", ssaSubtitle.Filename)
	fmt.Println("Format:", ssaSubtitle.Format)
	fmt.Println("Lines:", len(ssaSubtitle.Lines))

	/*
		for i := range ssaSubtitle.Lines {
			fmt.Println("\nSubtitle Entry", i+1)
			fmt.Println("Sequence:", ssaSubtitle.Lines[i].Seq)
			fmt.Println("Start Time:", ssaSubtitle.Lines[i].Start)
			fmt.Println("End Time:", ssaSubtitle.Lines[i].End)
			fmt.Println("Text:", ssaSubtitle.Lines[i].Text, "Length:", len(ssaSubtitle.Lines[i].Text))
		}
	*/
}
