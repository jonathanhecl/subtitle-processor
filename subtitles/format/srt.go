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

	exp := regexp.MustCompile(`(\d+)\n((\d*)\:(\d*)\:(\d*)[\.,\:](\d*) --> (\d*)\:(\d*)\:(\d*)[\.,\:](\d*))\n((?:\n?.)*?)\n\n`)
	if !exp.MatchString(content) {
		return ret, errors.New("Invalid SRT")
	}
	lines := strings.Split(content, "\n")
	dummy := models.ModelItemSubtitle{}

	// Add a final empty line to ensure the last subtitle is processed
	lines = append(lines, "")

	for i := range lines {
		if len(lines[i]) > 0 {
			if dummy.Seq == 0 {
				cleanLine := cleanText(lines[i])
				seq, err := strconv.Atoi(cleanLine)
				if err != nil {
					return ret, err
				}
				dummy.Seq = seq
			} else if dummy.Seq > 0 {
				start, end, err := formatStringSRT2Duration(lines[i])
				if err == nil {
					dummy.Start = start
					dummy.End = end
				} else {
					dummy.Text = append(dummy.Text, cleanText(lines[i]))
				}
			}
		} else if (len(lines[i]) == 0 || i == len(lines)-1) && dummy.Seq > 0 {
			if dummy.Start.Milliseconds() > 0 && dummy.End.Milliseconds() > 0 && len(dummy.Text) > 0 {
				ret = append(ret, dummy)
			}
			dummy = models.ModelItemSubtitle{} // empty
		}
	}
	return ret, err
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

// WriteSRT converts subtitle data from the internal model to SRT formatted content.
func WriteSRT(sub *models.Subtitle) (content string) {
	for i := range sub.Lines {
		content += fmt.Sprintf("%d\n%v --> %v\n", sub.Lines[i].Seq, sub.Lines[i].Start, sub.Lines[i].End)
		for j := range sub.Lines[i].Text {
			content += cleanText(sub.Lines[i].Text[j]) + "\n"
		}
		content += "\n"
	}
	return content
}
