package format

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/jonathanhecl/subtitle-processor/subtitles/models"
)

/*
SRT Format Specification:

Regular Expression for validation:
(\d)\n((\d*)\:(\d*)\:(\d*)[\.,\:](\d*) --> (\d*)\:(\d*)\:(\d*)[\.,\:](\d*))\n((?:\n?.)*?)\n\n

Example SRT Format:
1
00:02:17,440 --> 00:02:20,375
Senator, we're making
our final approach into Coruscant.

2
00:02:20,476 --> 00:02:22,501
Very good, Lieutenant.
*/

// ReadSRT parses SRT formatted subtitle content and converts it to the internal model.
// Returns an error if the content is not a valid SRT format.
func ReadSRT(content string) (ret []models.ModelItemSubtitle, err error) {
	content = cleanText(content)

	// Split content into lines
	lines := strings.Split(content, "\n")

	var currentSubtitle models.ModelItemSubtitle
	var parsingText bool
	var hasSubtitles bool

	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])

		// Skip empty lines
		if line == "" {
			// If we were parsing text, this empty line indicates the end of a subtitle
			if parsingText && currentSubtitle.Seq > 0 {
				ret = append(ret, currentSubtitle)
				currentSubtitle = models.ModelItemSubtitle{}
				parsingText = false
			}
			continue
		}

		// If not parsing text, check if this is a sequence number
		if !parsingText {
			seq, err := strconv.Atoi(line)
			if err == nil {
				currentSubtitle.Seq = seq
				continue
			}

			// If not a sequence number, check if it's a timestamp
			start, end, err := formatStringSRT2Duration(line)
			if err == nil {
				currentSubtitle.Start = start
				currentSubtitle.End = end
				parsingText = true
				hasSubtitles = true
				continue
			}
		} else {
			// We're parsing text, add this line to the current subtitle
			currentSubtitle.Text = append(currentSubtitle.Text, cleanText(line))
		}
	}

	// Add the last subtitle if we were parsing one
	if parsingText && currentSubtitle.Seq > 0 {
		ret = append(ret, currentSubtitle)
	}

	// If no subtitles were found, return an error
	if !hasSubtitles || len(ret) == 0 {
		return nil, errors.New("Invalid SRT")
	}

	return ret, nil
}

// formatStringSRT2Duration parses a time range string in SRT format (00:00:00,000 --> 00:00:00,000)
// and converts it to start and end time.Duration values.
func formatStringSRT2Duration(line string) (start time.Duration, end time.Duration, err error) {
	exp := regexp.MustCompile(`(\d*)\:(\d*)\:(\d*)[\.,\:](\d*) --> (\d*)\:(\d*)\:(\d*)[\.,\:](\d*)`)
	res := exp.FindAllStringSubmatch(line, -1)
	if len(res) == 1 {
		if len(res[0]) == 9 {
			start = (time.Duration(toInt(res[0][1])) * time.Hour) + (time.Duration(toInt(res[0][2])) * time.Minute) + (time.Duration(toInt(res[0][3])) * time.Second) + (time.Duration(toInt(res[0][4])) * time.Millisecond)
			end = (time.Duration(toInt(res[0][5])) * time.Hour) + (time.Duration(toInt(res[0][6])) * time.Minute) + (time.Duration(toInt(res[0][7])) * time.Second) + (time.Duration(toInt(res[0][8])) * time.Millisecond)
		}
		return start, end, nil
	}
	return start, end, errors.New("not time")
}

// formatDuration2SRT converts a time.Duration to SRT time format string (h:mm:ss.cc).
func formatDuration2SRT(d time.Duration) string {
	// Extract the hour, minute, second, and millisecond components from the time.Duration
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60
	ms := int(d.Milliseconds()) % 1000

	// Format the SRT time string
	return fmt.Sprintf("%d:%02d:%02d.%03d", h, m, s, ms)
}

// WriteSRT converts subtitle data from the internal model to SRT formatted content.
func WriteSRT(sub *models.Subtitle) (content string) {
	for i := range sub.Lines {
		content += fmt.Sprintf("%d\n%s --> %s\n", sub.Lines[i].Seq, formatDuration2SRT(sub.Lines[i].Start), formatDuration2SRT(sub.Lines[i].End))
		for j := range sub.Lines[i].Text {
			content += cleanText(sub.Lines[i].Text[j]) + "\n"
		}
		content += "\n"
	}
	return content
}
