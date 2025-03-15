package reader

import (
	"embed"
	"encoding/csv"
	"fmt"
)

func Read(fs embed.FS, fileName string) ([][]string, error) {
	file, err := fs.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("fs.Open(%s): %w", fileName, err)
	}

	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			_, _ = fmt.Println(closeErr)
		}
	}()

	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.Comment = '#'

	resp, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("reader.ReadAll(): %w", err)
	}

	return resp, nil
}
