package subtitles

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

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

	start := time.Now()

	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		return err
	}
	content := string(raw)

	content = strings.Replace(content, "\r\n", "\n", -1) // standardize line break
	content += "\n\n"                                    // lastest line break
	ret, err := format.ReadSRT(content)                  // check if it is the SRT format
	if err == nil {
		sub.Format = "SRT"
		sub.Lines = ret
	}
	if len(sub.Format) == 0 {
		ret, err := format.ReadSSA(content) // check if it is the SSA format
		if err == nil {
			sub.Format = "SSA"
			sub.Lines = ret
		}
	}

	fmt.Println("Processed in ", time.Since(start).String())

	return err
}
