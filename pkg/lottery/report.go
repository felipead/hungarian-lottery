package lottery

import (
	"fmt"
	"strings"
)

type Report interface {
	ToString() string
	IncrementWinners(numbersMatching int)
	GetWinners(numbersMatching int) int
}

type report struct {
	winners []int
}

func NewReport() Report {
	return &report{
		winners: make([]int, NumPicks-1),
	}
}

func (r *report) GetWinners(numbersMatching int) int {
	index := numbersMatching - 2
	if index >= 0 {
		return r.winners[index]
	}
	panic("invalid numbers matching")
}

func (r *report) IncrementWinners(numbersMatching int) {
	index := numbersMatching - 2
	if index >= 0 {
		r.winners[index]++
	}
}

func (r *report) ToString() string {
	matches := r.winners
	var output strings.Builder

	for i := 0; i < len(matches); i++ {
		match := matches[i]
		last := len(matches) - 1

		if i != last {
			output.WriteString(fmt.Sprintf("%v ", match))
		} else {
			output.WriteString(fmt.Sprintf("%v", match))
		}
	}

	return output.String()
}
