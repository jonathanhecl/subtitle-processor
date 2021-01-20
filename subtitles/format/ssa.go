package format

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/jonathanhecl/subtitle-processor/subtitles/models"
)

/*

SSA Validation

\[Script Info\]\n((?:\n?.)*?\n\n)\[(.*?)[ ]?Styles\]\n((?:\n?.)*?\n\n)\[Events\]\n(?:.*\n?)(^Dialogue: (\d),(\d*):(\d*):(\d*).(\d*),(\d*):(\d*):(\d*).(\d*),(\w*),[\w ]*,\d*,\d*,\d*,[\w*]?,{\\pos\(\d*,\d*\)}(.*)$)*

Sections: \[Script Info\]\n((?:\n?.)*?\n\n)\[(.*?)[ ]?Styles\]\n((?:\n?.)*?\n\n)\[Events\]\n((?:\n?.)*)
Dialogues: ^Dialogue: (\d),(\d*):(\d*):(\d*).(\d*),(\d*):(\d*):(\d*).(\d*),(\w*),[\w ]*,\d*,\d*,\d*,[\w*]?,{\\pos\(\d*,\d*\)}(.*)$

*/

/*
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

func ReadSSA(content string) (ret []models.ModelItemSubtitle, err error) {
	exp := regexp.MustCompile(`\[Script Info\]\n((?:\n?.)*?\n\n)\[(.*?)[ ]?Styles\]\n((?:\n?.)*?\n\n)\[Events\]\n(?:.*\n?)(^Dialogue: (\d),(\d*):(\d*):(\d*).(\d*),(\d*):(\d*):(\d*).(\d*),(\w*),[\w ]*,\d*,\d*,\d*,[\w*]?,{\\pos\(\d*,\d*\)}(.*)$)*`)
	if !exp.MatchString(content) {
		return ret, errors.New("Invalid SSA")
	}
	exp = regexp.MustCompile(`\[Script Info\]\n((?:\n?.)*?\n\n)\[(.*?)[ ]?Styles\]\n((?:\n?.)*?\n\n)\[Events\]\n((?:\n?.)*)`)
	res1 := exp.FindAllStringSubmatch(content, -1)
	if len(res1) == 1 {
		if len(res1[0]) == 5 {
			//fmt.Println(len(res1[0]))
			//fmt.Println(res1[0][4])
			dialogues := res1[0][4]
			exp = regexp.MustCompile(`(?m)^Dialogue: (\d),(\d*):(\d*):(\d*).(\d*),(\d*):(\d*):(\d*).(\d*),(\w*),[\w ]*,\d*,\d*,\d*,[\w*]?,{\\pos\(\d*,\d*\)}(.*)$`)
			res2 := exp.FindAllStringSubmatch(dialogues, -1)
			if len(res2) > 0 {
				seq := 0
				dummy := models.ModelItemSubtitle{}
				for i := 0; i < len(res2); i++ {
					if len(res2[i]) == 12 {
						seq++
						dummy.Seq = seq
						dummy.Start = formatSSA2Duration(res2[i][2], res2[i][3], res2[i][4], res2[i][5])
						dummy.End = formatSSA2Duration(res2[i][6], res2[i][7], res2[i][8], res2[i][9])
						text := strings.Split(res2[i][11], "\\N")
						dummy.Text = text
					}
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

func formatSSA2Duration(hour string, minute string, second string, millisecond string) (duration time.Duration) {
	duration = (time.Duration(toInt(hour)) * time.Hour) + (time.Duration(toInt(minute)) * time.Minute) + (time.Duration(toInt(second)) * time.Second) + (time.Duration(toInt(millisecond)) * time.Millisecond)
	return duration
}
