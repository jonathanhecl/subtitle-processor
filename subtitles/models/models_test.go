package models

import (
	"testing"
	"time"
)

// TestSubtitleModel tests the Subtitle model structure
func TestSubtitleModel(t *testing.T) {
	// Create a test subtitle
	sub := Subtitle{
		Filename: "test.srt",
		Format:   "SRT",
		Verbose:  true,
		Lines: []ModelItemSubtitle{
			{
				Seq:   1,
				Start: 1*time.Minute + 30*time.Second,
				End:   1*time.Minute + 35*time.Second,
				Text:  []string{"Line 1", "Line 2"},
			},
			{
				Seq:   2,
				Start: 2*time.Minute,
				End:   2*time.Minute + 5*time.Second,
				Text:  []string{"Another line"},
			},
		},
	}

	// Verify the structure
	if sub.Filename != "test.srt" {
		t.Errorf("Expected filename 'test.srt', got '%s'", sub.Filename)
	}
	if sub.Format != "SRT" {
		t.Errorf("Expected format 'SRT', got '%s'", sub.Format)
	}
	if !sub.Verbose {
		t.Errorf("Expected Verbose to be true")
	}
	if len(sub.Lines) != 2 {
		t.Errorf("Expected 2 subtitle lines, got %d", len(sub.Lines))
	}

	// Test first subtitle line
	if sub.Lines[0].Seq != 1 {
		t.Errorf("Expected sequence 1, got %d", sub.Lines[0].Seq)
	}
	expectedStart := 1*time.Minute + 30*time.Second
	if sub.Lines[0].Start != expectedStart {
		t.Errorf("Expected start time %v, got %v", expectedStart, sub.Lines[0].Start)
	}
	expectedEnd := 1*time.Minute + 35*time.Second
	if sub.Lines[0].End != expectedEnd {
		t.Errorf("Expected end time %v, got %v", expectedEnd, sub.Lines[0].End)
	}
	if len(sub.Lines[0].Text) != 2 {
		t.Errorf("Expected 2 text lines, got %d", len(sub.Lines[0].Text))
	}
	if sub.Lines[0].Text[0] != "Line 1" {
		t.Errorf("Expected text 'Line 1', got '%s'", sub.Lines[0].Text[0])
	}
	if sub.Lines[0].Text[1] != "Line 2" {
		t.Errorf("Expected text 'Line 2', got '%s'", sub.Lines[0].Text[1])
	}

	// Test second subtitle line
	if sub.Lines[1].Seq != 2 {
		t.Errorf("Expected sequence 2, got %d", sub.Lines[1].Seq)
	}
	expectedStart = 2 * time.Minute
	if sub.Lines[1].Start != expectedStart {
		t.Errorf("Expected start time %v, got %v", expectedStart, sub.Lines[1].Start)
	}
	expectedEnd = 2*time.Minute + 5*time.Second
	if sub.Lines[1].End != expectedEnd {
		t.Errorf("Expected end time %v, got %v", expectedEnd, sub.Lines[1].End)
	}
	if len(sub.Lines[1].Text) != 1 {
		t.Errorf("Expected 1 text line, got %d", len(sub.Lines[1].Text))
	}
	if sub.Lines[1].Text[0] != "Another line" {
		t.Errorf("Expected text 'Another line', got '%s'", sub.Lines[1].Text[0])
	}
}
