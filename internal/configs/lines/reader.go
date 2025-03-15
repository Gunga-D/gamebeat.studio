package lines

import (
	"embed"
	"fmt"
	"github.com/releaseband/golang-developer-test/internal/configs/reader"
)

//go:embed lines.txt
var lines embed.FS

func parseLine(data []string) (*Line, error) {
	// todo: implement me
	return nil, nil
}

func ReadLines() (Lines, error) {
	data, err := reader.Read(lines, "lines.txt")
	if err != nil {
		return nil, fmt.Errorf("reader.Read(): %w", err)
	}

	resp := make([]Line, len(data))
	for i, str := range data {
		line, err := parseLine(str)
		if err != nil {
			return nil, fmt.Errorf("parseLines(): %w", err)
		}

		resp[i] = *line
	}

	return resp, nil
}
