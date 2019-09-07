package main

import (
	"strings"
	"time"

	tm "github.com/buger/goterm"
)

var display = [5][27]bool{
	1: {8: true, 18: true},
	3: {8: true, 18: true},
}

var digits = [...][5][3]bool{
	{{true, true, true}, {true, false, true}, {true, false, true}, {true, false, true}, {true, true, true}},
	{{false, true, true}, {false, false, true}, {false, false, true}, {false, false, true}, {false, false, true}},
	{{true, true, true}, {false, false, true}, {true, true, true}, {true, false, false}, {true, true, true}},
	{{true, true, true}, {false, false, true}, {true, true, true}, {false, false, true}, {true, true, true}},
	{{true, false, true}, {true, false, true}, {true, true, true}, {false, false, true}, {false, false, true}},
	{{true, true, true}, {true, false, false}, {true, true, true}, {false, false, true}, {true, true, true}},
	{{true, true, true}, {true, false, false}, {true, true, true}, {true, false, true}, {true, true, true}},
	{{true, true, true}, {false, false, true}, {false, true, false}, {true, false, false}, {true, false, false}},
	{{true, true, true}, {true, false, true}, {true, true, true}, {true, false, true}, {true, true, true}},
	{{true, true, true}, {true, false, true}, {true, true, true}, {false, false, true}, {false, false, true}},
}

func getDisplayWithTime(time *[6]int) *[5][27]bool {
	n := display

	for di, v := range time {
		digit := digits[v]
		jmp := di * 4
		if di >= 2 && di < 4 {
			jmp += 2
		} else if di >= 4 {
			jmp += 4
		}
		for i, v1 := range digit {
			for j, v2 := range v1 {
				n[i][j+jmp] = v2
			}
		}
	}

	return &n
}

func displayTime(t *[5][27]bool) {
	var builder strings.Builder
	for _, v := range t {
		for _, w := range v {
			if w {
				builder.WriteRune('█')
			} else {
				builder.WriteRune(' ')
			}
		}
		tm.Println(builder.String())
		builder.Reset()
	}
}

func main() {
	tm.Clear()
	for {
		t := time.Now()
		h, m, s := t.Hour(), t.Minute(), t.Second()
		tm.MoveCursor(1, 1)
		displayTime(getDisplayWithTime(&[...]int{h / 10, h % 10, m / 10, m % 10, s / 10, s % 10}))
		tm.Flush()
		time.Sleep(time.Second)
	}

}
