package subtitles

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/jonathanhecl/subtitle-processor/subtitles/format"
	"github.com/jonathanhecl/subtitle-processor/subtitles/models"
)

type Subtitle struct {
	Filename string
	Format   string
	Lines    []models.ModelItemSubtitle
}

func (sub *Subtitle) LoadFilename(filename string) (err error) {
	sub.Filename = filename

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		return err
	}
	content := string(data)
	content = strings.Replace(content, "\r\n", "\n", -1) // standardize line break
	content += "\n"                                      // last line break
	ret, err := format.ReadSRT(content)
	if err == nil {
		sub.Format = "SRT"
		sub.Lines = ret
	}

	return err
}
