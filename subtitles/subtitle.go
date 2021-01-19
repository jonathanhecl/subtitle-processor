package subtitles

import "time"

type modelItemSubtitle struct {
	Seq   int
	Start time.Time
	End   time.Time
	Text  string
}

type modelSubtitle struct {
	Filename string
	Format   string
	Lines    []modelItemSubtitle
}

func LoadFilename(filename string) modelSubtitle {
	n := modelSubtitle{
		Filename: filename,
		Format:   "unknown",
		Lines:    []modelItemSubtitle{},
	}
	return n
}
