// Package subtitles provides functionality for loading, processing, and saving subtitle files.
package subtitles

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jonathanhecl/subtitle-processor/subtitles/format"
	"github.com/jonathanhecl/subtitle-processor/subtitles/models"
)

// Subtitle represents a subtitle file with its metadata and content.
// It embeds the models.Subtitle type to provide the necessary data structure.
type Subtitle models.Subtitle

// LoadFile loads a subtitle file from the specified path and detects its format.
// Currently supported formats: SRT, SSA.
// If Verbose is set to true, it will print processing time information.
func (sub *Subtitle) LoadFile(filename string) (err error) {
	sub.Filename = filename

	start := time.Now()

	// Read the file content
	raw, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		return err
	}
	content := string(raw)

	// Standardize line breaks and ensure proper ending
	content = strings.Replace(content, "\r\n", "\n", -1) // standardize line break
	content += "\n\n"                                    // lastest line break
	
	// Try to parse as SRT format
	ret, err := format.ReadSRT(content)
	if err == nil {
		sub.Format = "SRT"
		sub.Lines = ret
	}
	
	// If not SRT, try to parse as SSA format
	if len(sub.Format) == 0 {
		ret, err := format.ReadSSA(content)
		if err == nil {
			sub.Format = "SSA"
			sub.Lines = ret
		}
	}

	// Print processing time if verbose mode is enabled
	if sub.Verbose {
		fmt.Println("Processed in ", time.Since(start).String())
	}

	return err
}

// SaveFile saves the subtitle data to a file in the specified format.
// The format is determined by the Format field of the Subtitle struct.
// Currently supported formats: SRT, SSA.
// If Verbose is set to true, it will print processing time information.
func (sub *Subtitle) SaveFile(filename string) (err error) {
	start := time.Now()

	// Check if format is specified
	if len(sub.Format) == 0 {
		return errors.New("Format not specified")
	}

	content := ""

	// Generate content based on the format
	if sub.Format == "SRT" {
		content = format.WriteSRT(&models.Subtitle{Lines: sub.Lines})
	}
	if sub.Format == "SSA" {
		content = format.WriteSSA(&models.Subtitle{Lines: sub.Lines})
	}

	// Write content to file
	err = os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Print processing time if verbose mode is enabled
	if sub.Verbose {
		fmt.Println("Processed in ", time.Since(start).String())
	}

	return err
}
