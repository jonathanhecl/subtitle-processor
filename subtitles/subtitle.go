package subtitles

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jonathanhecl/subtitle-processor/subtitles/format"
	"github.com/jonathanhecl/subtitle-processor/subtitles/models"
)

type Subtitle models.Subtitle

func (sub *Subtitle) LoadFile(filename string) (err error) {
	sub.Filename = filename

	start := time.Now()

	raw, err := os.ReadFile(filename)
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

	if sub.Verbose {
		fmt.Println("Processed in ", time.Since(start).String())
	}

	return err
}

func (sub *Subtitle) SaveFile(filename string) (err error) {
	start := time.Now()

	if len(sub.Format) == 0 {
		return errors.New("Format not specified")
	}

	content := ""

	if sub.Format == "SRT" {
		content = format.WriteSRT(&models.Subtitle{Lines: sub.Lines})
	}
	if sub.Format == "SSA" {
		content = format.WriteSSA(&models.Subtitle{Lines: sub.Lines})
	}

	err = os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		log.Fatal(err)
		return err
	}

	if sub.Verbose {
		fmt.Println("Processed in ", time.Since(start).String())
	}

	return err
}
