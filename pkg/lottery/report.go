package lottery

import (
	"fmt"
	"strings"
)

// Report tracks and reports the lottery wins. It is built by [lottery.Registry] during processing of lottery picks.
type Report interface {

	// IncrementWinnersHaving increments the number of winners having the specified amount of matches. For example,
	// if the given amount of matches is 4, that increases the number of wins for that group. However, if there's less
	// than 2 matches, that's not considered a win.
	IncrementWinnersHaving(matches int)

	// GetWinnersHaving returns the number of winners having the specified amount of matches.
	GetWinnersHaving(matches int) int

	// String formats the report for textual representation.
	String() string
}

type reportType struct {
	winners [NumPicks - 1]int
}

func NewReport() Report {
	return &reportType{}
}

func (r *reportType) IncrementWinnersHaving(matches int) {
	index := matches - 2
	if index >= 0 {
		r.winners[index]++
	}
}

func (r *reportType) GetWinnersHaving(matches int) int {
	index := matches - 2
	if index >= 0 {
		return r.winners[index]
	}
	return 0
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
