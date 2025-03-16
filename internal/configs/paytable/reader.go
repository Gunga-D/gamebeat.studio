package paytable

import (
	"embed"
	"errors"
	"fmt"
	"strconv"

	"github.com/releaseband/golang-developer-test/internal/configs/reader"
	"github.com/releaseband/golang-developer-test/internal/configs/symbols"
)

//go:embed pay_table.txt
var payTable embed.FS

func parsePayouts(data [][]string) (map[symbols.Symbol]Payout, error) {
	if len(data) == 0 {
		return nil, errors.New("no data")
	}

	res := make(map[symbols.Symbol]Payout)
	for idx, p := range data {
		ps := make([]uint64, 0, len(p))
		for _, v := range p {
			pv, err := strconv.ParseUint(v, 10, 64)
			if err != nil {
				return nil, err
			}
			ps = append(ps, pv)
		}
		res[symbols.Symbol(idx)] = Payout(ps)
	}
	return res, nil
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
