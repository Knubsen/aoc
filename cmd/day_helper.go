package cmd

import (
	day01 "github.com/Knubsen/aoc/days/d01"
	day02 "github.com/Knubsen/aoc/days/d02"
)

type DayFuncs struct {
	ExecPt01 func()
	ExecPt02 func()
}

var dayMap = map[string]DayFuncs{
	"d01": {ExecPt01: day01.ExecPt01, ExecPt02: day01.ExecPt02},
	"d02": {ExecPt01: day02.ExecPt01, ExecPt02: day02.ExecPt02},
}

func GetDayFuncs(day string) (DayFuncs, bool) {
	funcs, exists := dayMap[day]
	return funcs, exists
}
