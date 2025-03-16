package lines

import (
	"embed"
	"errors"
	"fmt"
	"strconv"

	"github.com/releaseband/golang-developer-test/internal/configs/reader"
)

//go:embed lines.txt
var lines embed.FS

func parseLine(data []string) (*Line, error) {
	if len(data) == 0 {
		return nil, errors.New("no input")
	}
	indices := make([]int, 0, len(data))
	for _, v := range data {
		pv, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		indices = append(indices, pv)
	}
	return NewLine(indices), nil
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
