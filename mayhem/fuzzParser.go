package fuzzParser

import "strings"
import "github.com/lucasepe/yml2dot/parser"
import fuzz "github.com/AdaLogics/go-fuzz-headers"


func mayhemit(bytes []byte) int {

    fuzzConsumer := fuzz.NewConsumer(bytes)
    var content string
    err := fuzzConsumer.CreateSlice(&content)
    if err != nil {
        return 0
    }

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