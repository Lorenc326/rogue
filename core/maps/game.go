package maps

import (
	_ "embed"
	"strings"
)

//go:embed templates/default.txt
var defaultMapStr string

// go:embed templates/easy.txt
// var easyMapStr string

type GameMap [][]string

func Read(name string) *GameMap {
	rows := strings.Split(defaultMapStr, "\n")
	res := make([][]string, 0, len(rows))
	for _, rowStr := range rows {
		res = append(res, strings.Split(rowStr, ""))
	}

	gamemap := GameMap(res)
	return &gamemap
}

func (m *GameMap) String() string {
	rows := []string{}
	for i, row := range *m {
		rows[i] = strings.Join(row, "")
	}
	return strings.Join(rows, "\n")
}
