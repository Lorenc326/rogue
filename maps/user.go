package maps

import "strings"

type UserVision [11][11]string

func (u *UserVision) String() string {
	rows := [11]string{}
	for i, row := range u {
		rows[i] = strings.Join(row[:], "")
	}
	return strings.Join(rows[:], "\n")
}
