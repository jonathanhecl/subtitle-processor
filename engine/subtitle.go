package main

import "time"

type ItemSubtitle struct {
	Start time.Duration
	End   time.Duration
	Text  string
}

var Subtitle []ItemSubtitle
