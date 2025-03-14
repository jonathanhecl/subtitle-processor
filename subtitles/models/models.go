package models

import "time"

type Subtitle struct {
	Filename string
	Format   string
	Lines    []ModelItemSubtitle
	Verbose  bool
}

type ModelItemSubtitle struct {
	Seq   int
	Start time.Duration
	End   time.Duration
	Text  []string
}
