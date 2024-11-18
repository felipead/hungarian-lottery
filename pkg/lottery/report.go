package lottery

import (
	"fmt"
	"strings"
)

type Report interface {
	IncrementWinnersHaving(numbersMatching int)
	GetWinnersHaving(numbersMatching int) int
	String() string
}

type reportType struct {
	winners [NumPicks - 1]int
}

func NewReport() Report {
	return &reportType{}
}

func (r *reportType) IncrementWinnersHaving(numbersMatching int) {
	index := numbersMatching - 2
	if index >= 0 {
		r.winners[index]++
	}
}

func (r *reportType) GetWinnersHaving(numbersMatching int) int {
	index := numbersMatching - 2
	if index >= 0 {
		return r.winners[index]
	}
	panic("invalid numbers matching")
}

func (r *reportType) String() string {
	var output strings.Builder
	length := len(r.winners)

	for i := 0; i < length; i++ {
		count := r.winners[i]

		if i != length-1 {
			output.WriteString(fmt.Sprintf("%v ", count))
		} else {
			output.WriteString(fmt.Sprintf("%v", count))
		}
	}

	return output.String()
}
