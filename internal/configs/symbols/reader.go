package symbols

import (
	"embed"
	"errors"
	"fmt"
	"strconv"

	"github.com/releaseband/golang-developer-test/internal/configs/reader"
)

//go:embed symbols.txt
var symbols embed.FS

const skipSymbol = -1

func parseReels(data [][]string) ([]Symbols, error) {
	if len(data) == 0 {
		return nil, errors.New("no data")
	}

	res := make([]Symbols, len(data[0]))
	for _, line := range data {
		for idx, symb := range line {
			pSymb, err := strconv.Atoi(symb)
			if err != nil {
				return nil, err
			}
			if pSymb == skipSymbol {
				continue
			}
			res[idx] = append(res[idx], Symbol(pSymb))
		}
	}
	return res, nil
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
