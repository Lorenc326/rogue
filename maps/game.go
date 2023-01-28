package maps

import (
	"fmt"
	"os"
	"strings"
)

type GameMap [][]string

func Read(name string) *GameMap {
	asset := fmt.Sprintf("maps/templates/%s.txt", name)
	data, err := os.ReadFile(asset)
	if err != nil {
		fmt.Println("File reading error", err)
		return nil
	}

	rows := strings.Split(string(data), "\n")
	res := make([][]string, 0, len(rows))
	for _, rowStr := range rows {
		res = append(res, strings.Split(rowStr, ""))
	}

	gamemap := GameMap(res)
	return &gamemap
}
