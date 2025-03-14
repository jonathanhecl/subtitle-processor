# Subtitle Processor


### SRT validation
![SRT validation](https://i.imgur.com/1ait2i1.png)

### SSA validation
![SSA validation](https://i.imgur.com/1igtUyX.png)


### Installation

```bash
go get github.com/jonathanhecl/subtitle-processor
```

### Usage

```go
import "github.com/jonathanhecl/subtitle-processor/subtitles"

func main() {
    sub := subtitles.Subtitle{}
    sub.LoadFile("archivo.srt")
    // Procesar subtítulos...
    sub.SaveFile("nuevo.srt")
}
```