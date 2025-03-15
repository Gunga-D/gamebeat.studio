package paytable

import (
	"embed"
	"fmt"
	"github.com/releaseband/golang-developer-test/internal/configs/reader"
	"github.com/releaseband/golang-developer-test/internal/configs/symbols"
)

//go:embed pay_table.txt
var payTable embed.FS

func parsePayouts(data [][]string) (map[symbols.Symbol]Payout, error) {
	// todo: implement me
	return nil, nil
}

func ReadPayTable() (*PayTable, error) {
	data, err := reader.Read(payTable, "pay_table.txt")
	if err != nil {
		return nil, fmt.Errorf("reader.Read(): %w", err)
	}

	payouts, err := parsePayouts(data)
	if err != nil {
		return nil, fmt.Errorf("parsePayouts(): %w", err)
	}

	return NewPayTable(payouts), nil
}
