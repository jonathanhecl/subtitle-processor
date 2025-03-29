// Package models provides the data structures for subtitle processing.
package models

import "time"

// Subtitle represents a subtitle file with its metadata and content.
type Subtitle struct {
	Filename string                // Path to the subtitle file
	Format   string                // Format of the subtitle file (e.g., "SRT", "SSA")
	Lines    []ModelItemSubtitle   // Collection of subtitle entries
	Verbose  bool                  // Whether to print processing information
}

// ModelItemSubtitle represents a single subtitle entry with timing and text.
type ModelItemSubtitle struct {
	Seq   int            // Sequence number of the subtitle
	Start time.Duration  // Start time of the subtitle
	End   time.Duration  // End time of the subtitle
	Text  []string       // Lines of text in the subtitle
}
