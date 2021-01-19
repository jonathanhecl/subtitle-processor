package format

import (
	"time"
)

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
Dialogue: Marked=0,0:00:01.18,0:00:06.85,DefaultVCD, NTP,0000,0000,0000,,{\pos(400,570)}Like an angel with pity on nobody
*/

type InfoSSA struct {
	Title string
}

type StyleSSA struct {
	Name            string
	Fontname        string
	Fontsize        int
	PrimaryColour   int
	SecondaryColour int
	OutlineColour   int
	BackColour      int
	Bold            int
	Italic          int
	Underline       int
	StrikeOut       int
	ScaleX          int
	ScaleY          int
	Spacing         int
	Angle           float32
	BorderStyle     int
	Outline         float32
	Shadow          float32
	Alignment       int
	MarginL         int
	MarginR         int
	MarginV         int
	Encoding        int
}

type ItemSSA struct {
	Start   time.Duration
	End     time.Duration
	Style   string
	Name    string
	MarginL int
	MarginR int
	MarginV int
	Effect  string
	Text    string
}

func IsSSA(content string) bool {
	return false
}
