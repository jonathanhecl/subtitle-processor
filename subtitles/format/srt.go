package format

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/jonathanhecl/subtitle-processor/subtitles/models"
)

const (
	srtTimeSeparator = " --> "
)

/*
1
00:02:17,440 --> 00:02:20,375
Senator, we're making
our final approach into Coruscant.

2
00:02:20,476 --> 00:02:22,501
Very good, Lieutenant.
*/

func ReadSRT(content string) (ret []models.ModelItemSubtitle, err error) {
	lines := strings.Split(content, "\n")
	dummy := models.ModelItemSubtitle{}
	for i := range lines {
		if len(lines[i]) > 0 {
			if dummy.Seq == 0 {
				seq, err := strconv.Atoi(lines[i])
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
					dummy.Text = append(dummy.Text, lines[i])
				}
			}
		} else if len(lines[i]) == 0 || len(lines) == i {
			if dummy.Seq > 0 && dummy.Start.Milliseconds() > 0 &&
				dummy.End.Milliseconds() > 0 && len(dummy.Text) > 0 {
				ret = append(ret, dummy)
			}
			dummy.Seq = 0
			dummy.Start = time.Duration(0)
			dummy.End = time.Duration(0)
			dummy.Text = []string{}
		}
	}
	return ret, err
}

func formatStringSRT2Duration(line string) (start time.Duration, end time.Duration, err error) {
	exp := regexp.MustCompile(`(\d*)\:(\d*)\:(\d*)[\.,\:](\d*) --> (\d*)\:(\d*)\:(\d*)[\.,\:](\d*)`)
	res := exp.Copy().FindAllStringSubmatch(line, -1)
	if len(res) == 1 {
		if len(res[0]) == 9 {
			start = (time.Duration(toInt(res[0][1])) * time.Hour) + (time.Duration(toInt(res[0][2])) * time.Minute) + (time.Duration(toInt(res[0][3])) * time.Second) + (time.Duration(toInt(res[0][4])) * time.Millisecond)
			end = (time.Duration(toInt(res[0][5])) * time.Hour) + (time.Duration(toInt(res[0][6])) * time.Minute) + (time.Duration(toInt(res[0][7])) * time.Second) + (time.Duration(toInt(res[0][8])) * time.Millisecond)
		}
		return start, end, nil
	}
	return start, end, errors.New("not time")
}
