package models

import "time"

type ModelItemSubtitle struct {
	Seq   int
	Start time.Duration
	End   time.Duration
	Text  []string
}
