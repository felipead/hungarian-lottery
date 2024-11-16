package lottery

import (
	"fmt"
	"strings"
)

type Report interface {
	ToString() string
	IncrementWinnersOf(numbersMatching int)
	GetWinnersOf(numbersMatching int) int
}

type report struct {
	winners []int
}

func NewReport() Report {
	return &report{
		winners: make([]int, NumPicks-1),
	}
}

func (r *report) IncrementWinnersOf(numbersMatching int) {
	index := numbersMatching - 2
	if index >= 0 {
		r.winners[index]++
	}
}

func (r *report) GetWinnersOf(numbersMatching int) int {
	index := numbersMatching - 2
	if index >= 0 {
		return r.winners[index]
	}
	panic("invalid numbers matching")
}

func (r *report) ToString() string {
	winners := r.winners
	var output strings.Builder

	for i := 0; i < len(winners); i++ {
		count := winners[i]
		last := len(winners) - 1

		if i != last {
			output.WriteString(fmt.Sprintf("%v ", count))
		} else {
			output.WriteString(fmt.Sprintf("%v", count))
		}
	}

	return output.String()
}
