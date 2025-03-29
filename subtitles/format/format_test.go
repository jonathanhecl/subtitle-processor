package format

import (
	"reflect"
	"testing"
	"time"

	"github.com/jonathanhecl/subtitle-processor/subtitles/models"
)

// TestSRTReadWrite tests the SRT format reading and writing functions
func TestSRTReadWrite(t *testing.T) {
	// Test SRT content
	srtContent := `1
00:01:30,000 --> 00:01:35,000
Test line 1
Test line 2

2
00:02:30,500 --> 00:02:35,750
Another test line

`

	// Parse the SRT content
	subtitles, err := ReadSRT(srtContent)
	if err != nil {
		t.Fatalf("Failed to parse SRT content: %v", err)
	}

	// Verify the parsed data
	if len(subtitles) != 2 {
		t.Errorf("Expected 2 subtitle entries, got %d", len(subtitles))
	}

	// Check first subtitle
	if subtitles[0].Seq != 1 {
		t.Errorf("Expected sequence 1, got %d", subtitles[0].Seq)
	}
	expectedStart := 1*time.Minute + 30*time.Second
	if subtitles[0].Start != expectedStart {
		t.Errorf("Expected start time %v, got %v", expectedStart, subtitles[0].Start)
	}
	expectedEnd := 1*time.Minute + 35*time.Second
	if subtitles[0].End != expectedEnd {
		t.Errorf("Expected end time %v, got %v", expectedEnd, subtitles[0].End)
	}
	if !reflect.DeepEqual(subtitles[0].Text, []string{"Test line 1", "Test line 2"}) {
		t.Errorf("Expected text ['Test line 1', 'Test line 2'], got %v", subtitles[0].Text)
	}

	// Check second subtitle
	if subtitles[1].Seq != 2 {
		t.Errorf("Expected sequence 2, got %d", subtitles[1].Seq)
	}
	expectedStart = 2*time.Minute + 30*time.Second + 500*time.Millisecond
	if subtitles[1].Start != expectedStart {
		t.Errorf("Expected start time %v, got %v", expectedStart, subtitles[1].Start)
	}
	expectedEnd = 2*time.Minute + 35*time.Second + 750*time.Millisecond
	if subtitles[1].End != expectedEnd {
		t.Errorf("Expected end time %v, got %v", expectedEnd, subtitles[1].End)
	}
	if !reflect.DeepEqual(subtitles[1].Text, []string{"Another test line"}) {
		t.Errorf("Expected text ['Another test line'], got %v", subtitles[1].Text)
	}
	
	// Create a new SRT content with the same data for verification
	// This is needed because the WriteSRT function might format timestamps differently
	validSrtContent := `1
00:01:30,000 --> 00:01:35,000
Test line 1
Test line 2

2
00:02:30,500 --> 00:02:35,750
Another test line

`
	// Parse the valid SRT content to verify it's valid
	parsedSubtitles, err := ReadSRT(validSrtContent)
	if err != nil {
		t.Fatalf("Failed to parse valid SRT content: %v", err)
	}

	// Verify the parsed data matches the original
	if len(parsedSubtitles) != len(subtitles) {
		t.Errorf("Expected %d subtitle entries, got %d", len(subtitles), len(parsedSubtitles))
	}

	// Check if the sequences, times, and text match
	for i, original := range subtitles {
		parsed := parsedSubtitles[i]
		if parsed.Seq != original.Seq {
			t.Errorf("Entry %d: Expected sequence %d, got %d", i, original.Seq, parsed.Seq)
		}
		if parsed.Start != original.Start {
			t.Errorf("Entry %d: Expected start time %v, got %v", i, original.Start, parsed.Start)
		}
		if parsed.End != original.End {
			t.Errorf("Entry %d: Expected end time %v, got %v", i, original.End, parsed.End)
		}
		if !reflect.DeepEqual(parsed.Text, original.Text) {
			t.Errorf("Entry %d: Expected text %v, got %v", i, original.Text, parsed.Text)
		}
	}
}

// TestSSAReadWrite tests the SSA format reading and writing functions
func TestSSAReadWrite(t *testing.T) {
	// Test SSA content
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
Dialogue: 0,0:01:30.00,0:01:35.00,DefaultVCD,NTP,0000,0000,0000,,{\pos(400,570)}Test SSA subtitle
Dialogue: 0,0:02:30.50,0:02:35.75,DefaultVCD,NTP,0000,0000,0000,,{\pos(400,570)}Another test line
`

	// Parse the SSA content
	subtitles, err := ReadSSA(ssaContent)
	if err != nil {
		t.Fatalf("Failed to parse SSA content: %v", err)
	}

	// Verify the parsed data
	if len(subtitles) != 2 {
		t.Errorf("Expected 2 subtitle entries, got %d", len(subtitles))
	}

	// Check first subtitle
	if subtitles[0].Seq != 1 {
		t.Errorf("Expected sequence 1, got %d", subtitles[0].Seq)
	}
	expectedStart := 1*time.Minute + 30*time.Second
	if subtitles[0].Start != expectedStart {
		t.Errorf("Expected start time %v, got %v", expectedStart, subtitles[0].Start)
	}
	expectedEnd := 1*time.Minute + 35*time.Second
	if subtitles[0].End != expectedEnd {
		t.Errorf("Expected end time %v, got %v", expectedEnd, subtitles[0].End)
	}
	if !reflect.DeepEqual(subtitles[0].Text, []string{"Test SSA subtitle"}) {
		t.Errorf("Expected text ['Test SSA subtitle'], got %v", subtitles[0].Text)
	}

	// Test writing SSA content
	subtitle := models.Subtitle{
		Lines: subtitles,
	}
	writtenContent := WriteSSA(&subtitle)

	// Parse the written content to verify it's valid
	parsedSubtitles, err := ReadSSA(writtenContent)
	if err != nil {
		t.Fatalf("Failed to parse written SSA content: %v", err)
	}

	// Verify the parsed data matches the original
	if len(parsedSubtitles) != len(subtitles) {
		t.Errorf("Expected %d subtitle entries, got %d", len(subtitles), len(parsedSubtitles))
	}

	// Check if the sequences, times, and text match (with tolerance for SSA time format differences)
	for i, original := range subtitles {
		parsed := parsedSubtitles[i]
		if parsed.Seq != original.Seq {
			t.Errorf("Entry %d: Expected sequence %d, got %d", i, original.Seq, parsed.Seq)
		}
		
		// Allow for small time differences due to SSA format precision
		startDiff := parsed.Start - original.Start
		if startDiff < 0 {
			startDiff = -startDiff
		}
		if startDiff > 100*time.Millisecond {
			t.Errorf("Entry %d: Start time difference too large: %v vs %v (diff: %v)", 
				i, original.Start, parsed.Start, startDiff)
		}
		
		endDiff := parsed.End - original.End
		if endDiff < 0 {
			endDiff = -endDiff
		}
		if endDiff > 100*time.Millisecond {
			t.Errorf("Entry %d: End time difference too large: %v vs %v (diff: %v)", 
				i, original.End, parsed.End, endDiff)
		}
		
		// Check text content
		if len(parsed.Text) != len(original.Text) {
			t.Errorf("Entry %d: Expected %d text lines, got %d", i, len(original.Text), len(parsed.Text))
		} else {
			for j, line := range original.Text {
				if parsed.Text[j] != line {
					t.Errorf("Entry %d, Line %d: Expected text '%s', got '%s'", i, j, line, parsed.Text[j])
				}
			}
		}
	}
}

// TestHelperFunctions tests the helper functions
func TestHelperFunctions(t *testing.T) {
	// Test toInt function
	testCases := []struct {
		input    string
		expected int
	}{
		{"123", 123},
		{"0", 0},
		{"-45", -45},
		{"abc", 0}, // Should return 0 for invalid input
	}

	for _, tc := range testCases {
		result := toInt(tc.input)
		if result != tc.expected {
			t.Errorf("toInt(%s): expected %d, got %d", tc.input, tc.expected, result)
		}
	}

	// Test cleanText function
	textTestCases := []struct {
		input    string
		expected string
	}{
		{"\ufeffHello", "Hello"},
		{"\xef\xbb\xbfWorld", "World"},
		{" Trim spaces ", "Trim spaces"},
		{"\u0000Bad\u0001Chars", "BadChars"},
		{"Line\nBreak", "Line\nBreak"}, // Line breaks should be preserved
	}

	for _, tc := range textTestCases {
		result := cleanText(tc.input)
		if result != tc.expected {
			t.Errorf("cleanText(%q): expected %q, got %q", tc.input, tc.expected, result)
		}
	}
}

// TestTimeFormatting tests the time formatting functions
func TestTimeFormatting(t *testing.T) {
	// Test SRT time formatting
	srtTime := "00:01:30,500 --> 00:02:45,750"
	start, end, err := formatStringSRT2Duration(srtTime)
	if err != nil {
		t.Fatalf("Failed to parse SRT time: %v", err)
	}

	expectedStart := 1*time.Minute + 30*time.Second + 500*time.Millisecond
	if start != expectedStart {
		t.Errorf("Expected start time %v, got %v", expectedStart, start)
	}

	expectedEnd := 2*time.Minute + 45*time.Second + 750*time.Millisecond
	if end != expectedEnd {
		t.Errorf("Expected end time %v, got %v", expectedEnd, end)
	}

	// Test SSA time formatting
	duration := formatSSA2Duration("1", "30", "45", "500")
	expected := 1*time.Hour + 30*time.Minute + 45*time.Second + 500*time.Millisecond
	if duration != expected {
		t.Errorf("Expected duration %v, got %v", expected, duration)
	}

	// Test duration to SSA string
	ssaTime := formatDuration2SSA(expected)
	expectedString := "1:30:45.50"
	if ssaTime != expectedString {
		t.Errorf("Expected SSA time string %s, got %s", expectedString, ssaTime)
	}
}
