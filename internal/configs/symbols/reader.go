package symbols

import (
	"embed"
	"fmt"
	"github.com/releaseband/golang-developer-test/internal/configs/reader"
)

//go:embed symbols.txt
var symbols embed.FS

const skipSymbol = -1

func parseReels(data [][]string) ([]Symbols, error) {
	// todo: implement me
	return nil, nil
}

// ReadReels - read symbols from file
func ReadReels() ([]Symbols, error) {
	// обрати внимание, что в файле symbols.txt символы разделены через \t
	// и что в конце каждой строки есть \n
	// символ -1 нужен только для выравнивания таблицы
	data, err := reader.Read(symbols, "symbols.txt")
	if err != nil {
		return nil, fmt.Errorf("reader.Read(): %w", err)
	}

	symbols, err := parseReels(data)
	if err != nil {
		return nil, fmt.Errorf("parseReels(): %w", err)
	}

	return symbols, nil
}
