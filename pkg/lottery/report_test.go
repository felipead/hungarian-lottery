package lottery

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReportEmpty(t *testing.T) {
	report := NewReport()

	assert.Equal(t, 0, report.GetWinnersOf(5))
	assert.Equal(t, 0, report.GetWinnersOf(4))
	assert.Equal(t, 0, report.GetWinnersOf(3))
	assert.Equal(t, 0, report.GetWinnersOf(2))

	assert.Equal(t, report.ToString(), "0 0 0 0")
}

func TestReportIncremented(t *testing.T) {
	report := NewReport()

	report.IncrementWinnersOf(2)
	report.IncrementWinnersOf(2)
	report.IncrementWinnersOf(2)
	report.IncrementWinnersOf(2)
	report.IncrementWinnersOf(2)
	report.IncrementWinnersOf(2)

	report.IncrementWinnersOf(5)
	report.IncrementWinnersOf(5)
	report.IncrementWinnersOf(5)

	report.IncrementWinnersOf(4)
	report.IncrementWinnersOf(4)
	report.IncrementWinnersOf(4)
	report.IncrementWinnersOf(4)

	report.IncrementWinnersOf(3)
	report.IncrementWinnersOf(3)
	report.IncrementWinnersOf(3)
	report.IncrementWinnersOf(3)
	report.IncrementWinnersOf(3)

	// NO EFFECT
	report.IncrementWinnersOf(1)
	report.IncrementWinnersOf(1)
	report.IncrementWinnersOf(1)
	report.IncrementWinnersOf(1)
	report.IncrementWinnersOf(1)
	report.IncrementWinnersOf(1)

	assert.Equal(t, 6, report.GetWinnersOf(2))
	assert.Equal(t, 5, report.GetWinnersOf(3))
	assert.Equal(t, 4, report.GetWinnersOf(4))
	assert.Equal(t, 3, report.GetWinnersOf(5))

	assert.Equal(t, report.ToString(), "6 5 4 3")
}
