package format

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/jonathanhecl/subtitle-processor/subtitles/models"
)

/*
SSA Format Specification:

Regular Expressions for validation:
- Full SSA validation: 
\[Script Info\]\n((?:\n?.)*?\n\n)\[(.*?)[ ]?Styles\]\n((?:\n?.)*?\n\n)\[Events\]\n(?:.*\n?)(^Dialogue: (\d),(\d*):(\d*):(\d*).(\d*),(\d*):(\d*):(\d*).(\d*),(\w*),[\w ]*,\d*,\d*,\d*,[\w*]?,{\\pos\(\d*,\d*\)}(.*)$)*

- Sections validation: 
\[Script Info\]\n((?:\n?.)*?\n\n)\[(.*?)[ ]?Styles\]\n((?:\n?.)*?\n\n)\[Events\]\n((?:\n?.)*)

- Dialogue line validation: 
^Dialogue: (\d),(\d*):(\d*):(\d*).(\d*),(\d*):(\d*):(\d*).(\d*),(\w*),[\w ]*,\d*,\d*,\d*,[\w*]?,{\\pos\(\d*,\d*\)}(.*)$

Example SSA Format:
[Script Info]
; This is a Sub Station Alpha v4 script.
; For Sub Station Alpha info and downloads,
; go to http://www.eswat.demon.co.uk/
Title: Neon Genesis Evangelion - Episode 26 (neutral Spanish)
Original Script: RoRo
Script Updated By: version 2.8.01
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
Dialogue: 0,0:00:01.18,0:00:06.85,DefaultVCD, NTP,0000,0000,0000,,{\pos(400,570)}Like an angel with pity on nobody
*/

// ReadSSA parses SSA formatted subtitle content and converts it to the internal model.
// Returns an error if the content is not a valid SSA format.
func ReadSSA(content string) (ret []models.ModelItemSubtitle, err error) {
	// Clean the input content to remove any unnecessary characters
	content = cleanText(content)

	// Regular expression to match the full SSA format
	exp := regexp.MustCompile(`\[Script Info\]\n((?:\n?.)*?\n\n)\[(.*?)[ ]?Styles\]\n((?:\n?.)*?\n\n)\[Events\]\n(?:.*\n?)(^Dialogue: (\d),(\d*):(\d*):(\d*).(\d*),(\d*):(\d*):(\d*).(\d*),(\w*),[\w ]*,\d*,\d*,\d*,[\w*]?,{\\pos\(\d*,\d*\)}(.*)$)*`)
	if !exp.MatchString(content) {
		return ret, errors.New("Invalid SSA")
	}

	// Regular expression to match the sections of the SSA format
	exp = regexp.MustCompile(`\[Script Info\]\n((?:\n?.)*?\n\n)\[(.*?)[ ]?Styles\]\n((?:\n?.)*?\n\n)\[Events\]\n((?:\n?.)*)`)
	res1 := exp.FindAllStringSubmatch(content, -1)
	if len(res1) == 1 {
		if len(res1[0]) == 5 {
			// Extract the dialogue lines from the SSA content
			dialogues := res1[0][4]

			// Regular expression to match each dialogue line
			exp = regexp.MustCompile(`(?m)^Dialogue: (\d),(\d*):(\d*):(\d*).(\d*),(\d*):(\d*):(\d*).(\d*),(\w*),[\w ]*,\d*,\d*,\d*,[\w*]?,{\\pos\(\d*,\d*\)}(.*)$`)
			res2 := exp.FindAllStringSubmatch(dialogues, -1)
			if len(res2) > 0 {
				// Initialize variables to keep track of the sequence and subtitle data
				seq := 0
				dummy := models.ModelItemSubtitle{}

				// Iterate over each dialogue line and extract the relevant data
				for i := 0; i < len(res2); i++ {
					if len(res2[i]) == 12 {
						seq++
						dummy.Seq = seq
						dummy.Start = formatSSA2Duration(res2[i][2], res2[i][3], res2[i][4], res2[i][5])
						dummy.End = formatSSA2Duration(res2[i][6], res2[i][7], res2[i][8], res2[i][9])
						text := strings.Split(cleanText(res2[i][11]), "\\N")
						dummy.Text = text
					}

					// Check if the subtitle data is valid and append it to the result
					if dummy.Seq > 0 && dummy.Start.Milliseconds() > 0 &&
						dummy.End.Milliseconds() > 0 && len(dummy.Text) > 0 {
						ret = append(ret, dummy)
						dummy = models.ModelItemSubtitle{} // empty
					}
				}
			}
		}
	}
	return ret, err
}

// formatSSA2Duration converts SSA time format components (hour, minute, second, millisecond)
// to a time.Duration value.
func formatSSA2Duration(hour string, minute string, second string, millisecond string) (duration time.Duration) {
	// Convert the SSA time format components to a time.Duration value
	duration = (time.Duration(toInt(hour)) * time.Hour) + (time.Duration(toInt(minute)) * time.Minute) + (time.Duration(toInt(second)) * time.Second) + (time.Duration(toInt(millisecond)) * time.Millisecond)
	return duration
}

// WriteSSA converts subtitle data from the internal model to SSA formatted content.
// Creates a standard SSA file with default styling.
func WriteSSA(sub *models.Subtitle) (content string) {
	// Create the SSA header
	content = `[Script Info]
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
`

	// Iterate over each subtitle line and create the SSA dialogue lines
	for i := range sub.Lines {
		start := formatDuration2SSA(sub.Lines[i].Start)
		end := formatDuration2SSA(sub.Lines[i].End)
		cleanedTexts := make([]string, len(sub.Lines[i].Text))
		for j, text := range sub.Lines[i].Text {
			cleanedTexts[j] = cleanText(text)
		}
		text := strings.Join(cleanedTexts, "\\N")
		content += fmt.Sprintf("Dialogue: 0,%s,%s,DefaultVCD,NTP,0000,0000,0000,,{\\pos(400,570)}%s\n", start, end, text)
	}
	return content
}

// formatDuration2SSA converts a time.Duration to SSA time format string (h:mm:ss.cc).
func formatDuration2SSA(d time.Duration) string {
	// Extract the hour, minute, second, and millisecond components from the time.Duration
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60
	ms := int(d.Milliseconds()) % 1000

	// Format the SSA time string
	return fmt.Sprintf("%d:%02d:%02d.%02d", h, m, s, ms/10)
}
