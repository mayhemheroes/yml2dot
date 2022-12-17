package fuzz

// import "io"
import "strings"
import "github.com/lucasepe/yml2dot/parser"


func mayhemit(bytes []byte) int {

    content := string(bytes)
    start := content[0:len(content)/2]
    end := content[len(content)/2:len(content)]

    reader := strings.NewReader(content)

    parser.Parse(reader, start, end)
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}