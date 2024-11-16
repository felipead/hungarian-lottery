package lottery

import (
	"fmt"
	"strings"
)

type Report interface {
	ToString() string
	IncrementMatches(count int)
}

type report struct {
	matches []int
}

func NewReport() Report {
	return &report{
		matches: make([]int, NumPicks-1),
	}
}

func (r *report) IncrementMatches(count int) {
	index := count - 2
	if index >= 0 {
		r.matches[index]++
	}
}

func (r *report) ToString() string {
	matches := r.matches
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
