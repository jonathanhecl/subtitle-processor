package subtitles

import (
	"os"
	"testing"
)

// TestFormatConversionIntegration tests the conversion between SRT and SSA formats
func TestFormatConversionIntegration(t *testing.T) {
	// Create temporary files for testing
	srtFile, err := os.CreateTemp("", "test-*.srt")
	if err != nil {
		t.Fatalf("Failed to create temp SRT file: %v", err)
	}
	srtFile.Close()
	defer os.Remove(srtFile.Name())

	// Create valid SRT content directly
	srtContent := `1
00:01:30,000 --> 00:01:35,000
First line
Second line

2
00:02:00,000 --> 00:02:05,000
Third line

`
	err = os.WriteFile(srtFile.Name(), []byte(srtContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write SRT content: %v", err)
	}

	// Test loading SRT file
	srtLoaded := Subtitle{}
	err = srtLoaded.LoadFile(srtFile.Name())
	if err != nil {
		t.Fatalf("Failed to load SRT file: %v", err)
	}

	// Verify SRT format
	if srtLoaded.Format != "SRT" {
		t.Errorf("Expected format SRT, got %s", srtLoaded.Format)
	}
	if len(srtLoaded.Lines) != 2 {
		t.Errorf("Expected 2 lines, got %d", len(srtLoaded.Lines))
	}

	// Create a valid SSA file directly
	ssaFile, err := os.CreateTemp("", "test-*.ssa")
	if err != nil {
		t.Fatalf("Failed to create temp SSA file: %v", err)
	}
	ssaFile.Close()
	defer os.Remove(ssaFile.Name())

	// Create valid SSA content directly
	ssaContent := `[Script Info]
; This is a Sub Station Alpha v4 script.
ScriptType: v4.00
Collisions: Normal
PlayResY: 600
PlayDepth: 0
Timer: 100,0000

[V4 Styles]
Format: Name, Fontname, Fontsize, PrimaryColour, SecondaryColour, TertiaryColour, BackColour, Bold, Italic, BorderStyle, Outline, Shadow, Alignment, MarginL, MarginR, MarginV, AlphaLevel, Encoding
Style: DefaultVCD, Arial,28,11861244,11861244,11861244,-2147483640,-1,0,1,1,2,2,30,30,30,0,0

[Events]
Format: Marked, Start, End, Style, Name, MarginL, MarginR, MarginV, Effect, Text
Dialogue: 0,0:01:30.00,0:01:35.00,DefaultVCD,NTP,0000,0000,0000,,{\pos(400,570)}First line\\NSecond line
Dialogue: 0,0:02:00.00,0:02:05.00,DefaultVCD,NTP,0000,0000,0000,,{\pos(400,570)}Third line
`
	err = os.WriteFile(ssaFile.Name(), []byte(ssaContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write SSA content: %v", err)
	}

	// Test loading SSA file
	ssaLoaded := Subtitle{}
	err = ssaLoaded.LoadFile(ssaFile.Name())
	if err != nil {
		t.Fatalf("Failed to load SSA file: %v", err)
	}

	// Verify SSA format
	if ssaLoaded.Format != "SSA" {
		t.Errorf("Expected format SSA, got %s", ssaLoaded.Format)
	}
	if len(ssaLoaded.Lines) != 2 {
		t.Errorf("Expected 2 lines, got %d", len(ssaLoaded.Lines))
	}

	// Test converting SRT to SSA
	srtLoaded.Format = "SSA"
	convertedSsaFile, err := os.CreateTemp("", "converted-*.ssa")
	if err != nil {
		t.Fatalf("Failed to create temp converted SSA file: %v", err)
	}
	convertedSsaFile.Close()
	defer os.Remove(convertedSsaFile.Name())

	err = srtLoaded.SaveFile(convertedSsaFile.Name())
	if err != nil {
		t.Fatalf("Failed to save converted SSA file: %v", err)
	}

	// Test converting SSA to SRT
	ssaLoaded.Format = "SRT"
	convertedSrtFile, err := os.CreateTemp("", "converted-*.srt")
	if err != nil {
		t.Fatalf("Failed to create temp converted SRT file: %v", err)
	}
	convertedSrtFile.Close()
	defer os.Remove(convertedSrtFile.Name())

	err = ssaLoaded.SaveFile(convertedSrtFile.Name())
	if err != nil {
		t.Fatalf("Failed to save converted SRT file: %v", err)
	}
}

// TestErrorHandling tests error handling in the subtitle package
func TestErrorHandling(t *testing.T) {
	// Test loading invalid format
	tmpFile, err := os.CreateTemp("", "test-invalid-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	// Write invalid content
	_, err = tmpFile.WriteString("This is not a valid subtitle file")
	if err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	// Try to load the invalid file
	sub := Subtitle{}
	sub.LoadFile(tmpFile.Name())
	if sub.Format != "" {
		t.Errorf("Expected empty format for invalid file, got %s", sub.Format)
	}
	if len(sub.Lines) != 0 {
		t.Errorf("Expected 0 lines for invalid file, got %d", len(sub.Lines))
	}

	// Test saving with unspecified format
	sub = Subtitle{}
	saveFile, err := os.CreateTemp("", "test-save-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp save file: %v", err)
	}
	saveFile.Close()
	defer os.Remove(saveFile.Name())
	
	err = sub.SaveFile(saveFile.Name())
	if err == nil {
		t.Error("Expected error when saving with unspecified format, got nil")
	}
}
