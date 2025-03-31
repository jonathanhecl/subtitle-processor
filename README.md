# Subtitle Processor

A Go library for loading, processing, and saving subtitle files in different formats. Currently supports SRT and SSA subtitle formats.

## Features

- Load and parse subtitle files (SRT, SSA)
- Convert between different subtitle formats
- Modify subtitle content programmatically
- Save subtitles in different formats
- Clean handling of text encoding and special characters

## Supported Formats

### SRT (SubRip Text)
The SubRip format is one of the most common subtitle formats. It consists of:
- Sequence number
- Timestamp range (start --> end)
- Text content
- Blank line separator

### SSA (Sub Station Alpha)
The Sub Station Alpha format is more complex and includes styling information:
- Script Info section
- Styles section
- Events section with dialogue entries

## Installation

```bash
go get github.com/jonathanhecl/subtitle-processor
```

## Usage

### Basic Example

```go
package main

import (
    "fmt"
    "github.com/jonathanhecl/subtitle-processor/subtitles"
)

func main() {
    // Load a subtitle file
    sub := subtitles.Subtitle{}
    sub.Verbose = true  // Enable verbose output
    err := sub.LoadFile("input.srt")
    if err != nil {
        fmt.Println("Error loading file:", err)
        return
    }
    
    fmt.Println("Loaded subtitle file:", sub.Filename)
    fmt.Println("Format:", sub.Format)
    fmt.Println("Number of subtitle entries:", len(sub.Lines))
    
    // Access subtitle content
    for i, line := range sub.Lines {
        fmt.Printf("Entry %d: %v --> %v\n", i+1, line.Start, line.End)
        fmt.Printf("Text: %v\n", line.Text)
    }
    
    // Save to a different format
    sub.Format = "SSA"  // Change format to SSA
    err = sub.SaveFile("output.ssa")
    if err != nil {
        fmt.Println("Error saving file:", err)
    }
}
```

### Format Conversion

```go
package main

import (
    "github.com/jonathanhecl/subtitle-processor/subtitles"
)

func main() {
    // Convert from SRT to SSA
    sub := subtitles.Subtitle{}
    sub.LoadFile("input.srt")
    sub.Format = "SSA"
    sub.SaveFile("output.ssa")
    
    // Convert from SSA to SRT
    sub2 := subtitles.Subtitle{}
    sub2.LoadFile("input.ssa")
    sub2.Format = "SRT"
    sub2.SaveFile("output.srt")
}
```

### Modifying Subtitles

```go
package main

import (
    "time"
    "github.com/jonathanhecl/subtitle-processor/subtitles"
)

func main() {
    sub := subtitles.Subtitle{}
    sub.LoadFile("input.srt")
    
    // Delay all subtitles by 2 seconds
    delay := 2 * time.Second
    for i := range sub.Lines {
        sub.Lines[i].Start += delay
        sub.Lines[i].End += delay
    }
    
    sub.SaveFile("delayed.srt")
}
```

## Project Structure

- `subtitles/`: Main package
  - `models/`: Data structures for subtitle processing
  - `format/`: Format-specific parsers and writers
    - `srt.go`: SRT format handler
    - `ssa.go`: SSA format handler
    - `helper.go`: Common utility functions

## License

MIT License
