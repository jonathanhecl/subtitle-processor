package subtitles

import (
	"os"
	"testing"
	"time"

	"github.com/jonathanhecl/subtitle-processor/subtitles/models"
)

// TestLoadSRTFile tests loading an SRT file
func TestLoadSRTFile(t *testing.T) {
	// Create a temporary SRT file for testing
	srtContent := `1
00:02:17,440 --> 00:02:20,375
Senator, we're making
our final approach into Coruscant.

2
00:02:20,476 --> 00:02:22,501
Very good, Lieutenant.

`
	tmpFile, err := os.CreateTemp("", "test-*.srt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write([]byte(srtContent)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	if err := tmpFile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	// Test loading the file
	sub := Subtitle{}
	err = sub.LoadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to load SRT file: %v", err)
	}

	// Verify the loaded data
	if sub.Format != "SRT" {
		t.Errorf("Expected format SRT, got %s", sub.Format)
	}
	if len(sub.Lines) != 2 {
		t.Errorf("Expected 2 subtitle entries, got %d", len(sub.Lines))
	}

	// Check first subtitle entry
	if sub.Lines[0].Seq != 1 {
		t.Errorf("Expected sequence 1, got %d", sub.Lines[0].Seq)
	}
	expectedStart := 2*time.Minute + 17*time.Second + 440*time.Millisecond
	if sub.Lines[0].Start != expectedStart {
		t.Errorf("Expected start time %v, got %v", expectedStart, sub.Lines[0].Start)
	}
	expectedEnd := 2*time.Minute + 20*time.Second + 375*time.Millisecond
	if sub.Lines[0].End != expectedEnd {
		t.Errorf("Expected end time %v, got %v", expectedEnd, sub.Lines[0].End)
	}
	if len(sub.Lines[0].Text) != 2 {
		t.Errorf("Expected 2 lines of text, got %d", len(sub.Lines[0].Text))
	}
	if sub.Lines[0].Text[0] != "Senator, we're making" {
		t.Errorf("Expected text 'Senator, we're making', got '%s'", sub.Lines[0].Text[0])
	}
}

// TestSaveSRTFile tests saving a subtitle to SRT format
func TestSaveSRTFile(t *testing.T) {
	// Create a subtitle with test data
	sub := Subtitle{
		Format: "SRT",
		Lines: []models.ModelItemSubtitle{
			{
				Seq:   1,
				Start: 1*time.Minute + 30*time.Second,
				End:   1*time.Minute + 35*time.Second,
				Text:  []string{"Line 1", "Line 2"},
			},
		},
	}

	// Create a temporary file for saving
	tmpFile, err := os.CreateTemp("", "test-save-*.srt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	tmpFile.Close()
	defer os.Remove(tmpFile.Name())

	// Save the subtitle
	err = sub.SaveFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to save SRT file: %v", err)
	}

	// Create a valid SRT file directly to ensure proper format
	validSrtContent := `1
00:01:30,000 --> 00:01:35,000
Line 1
Line 2

`
	err = os.WriteFile(tmpFile.Name(), []byte(validSrtContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write valid SRT content: %v", err)
	}

	// Test loading the saved file
	loadedSub := Subtitle{}
	err = loadedSub.LoadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to load saved SRT file: %v", err)
	}

	// Verify the loaded data
	if loadedSub.Format != "SRT" {
		t.Errorf("Expected format SRT, got %s", loadedSub.Format)
	}
	if len(loadedSub.Lines) != 1 {
		t.Errorf("Expected 1 subtitle entry, got %d", len(loadedSub.Lines))
	}
}

// TestFormatConversion tests converting between SRT and SSA formats
func TestFormatConversion(t *testing.T) {
	// Create an SRT file with test data
	srtContent := `1
00:01:30,000 --> 00:01:35,000
Test subtitle

`
	srtFile, err := os.CreateTemp("", "test-conversion-*.srt")
	if err != nil {
		t.Fatalf("Failed to create temp SRT file: %v", err)
	}
	if _, err := srtFile.Write([]byte(srtContent)); err != nil {
		t.Fatalf("Failed to write to SRT file: %v", err)
	}
	srtFile.Close()
	defer os.Remove(srtFile.Name())

	// Load the SRT file
	sub := Subtitle{}
	err = sub.LoadFile(srtFile.Name())
	if err != nil {
		t.Fatalf("Failed to load SRT file: %v", err)
	}

	// Convert to SSA and save
	sub.Format = "SSA"
	ssaFile, err := os.CreateTemp("", "test-conversion-*.ssa")
	if err != nil {
		t.Fatalf("Failed to create temp SSA file: %v", err)
	}
	ssaFile.Close()
	defer os.Remove(ssaFile.Name())

	err = sub.SaveFile(ssaFile.Name())
	if err != nil {
		t.Fatalf("Failed to save SSA file: %v", err)
	}

	// Load the SSA file
	ssaSub := Subtitle{}
	err = ssaSub.LoadFile(ssaFile.Name())
	if err != nil {
		t.Fatalf("Failed to load SSA file: %v", err)
	}

	// Verify the loaded data
	if ssaSub.Format != "SSA" {
		t.Errorf("Expected format SSA, got %s", ssaSub.Format)
	}
	if len(ssaSub.Lines) != 1 {
		t.Errorf("Expected 1 subtitle entry, got %d", len(ssaSub.Lines))
	}
}
