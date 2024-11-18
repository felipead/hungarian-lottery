package lottery

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReportEmpty(t *testing.T) {
	report := NewReport()

	assert.Equal(t, 0, report.GetWinnersHaving(5))
	assert.Equal(t, 0, report.GetWinnersHaving(4))
	assert.Equal(t, 0, report.GetWinnersHaving(3))
	assert.Equal(t, 0, report.GetWinnersHaving(2))

	assert.Equal(t, report.String(), "0 0 0 0")
}

func TestReportIncremented(t *testing.T) {
	report := NewReport()

	report.IncrementWinnersHaving(2)
	report.IncrementWinnersHaving(2)
	report.IncrementWinnersHaving(2)
	report.IncrementWinnersHaving(2)
	report.IncrementWinnersHaving(2)
	report.IncrementWinnersHaving(2)

	report.IncrementWinnersHaving(5)
	report.IncrementWinnersHaving(5)
	report.IncrementWinnersHaving(5)

	report.IncrementWinnersHaving(4)
	report.IncrementWinnersHaving(4)
	report.IncrementWinnersHaving(4)
	report.IncrementWinnersHaving(4)

	report.IncrementWinnersHaving(3)
	report.IncrementWinnersHaving(3)
	report.IncrementWinnersHaving(3)
	report.IncrementWinnersHaving(3)
	report.IncrementWinnersHaving(3)

	// NO EFFECT
	report.IncrementWinnersHaving(1)
	report.IncrementWinnersHaving(1)
	report.IncrementWinnersHaving(1)
	report.IncrementWinnersHaving(1)
	report.IncrementWinnersHaving(1)
	report.IncrementWinnersHaving(1)

	assert.Equal(t, 6, report.GetWinnersHaving(2))
	assert.Equal(t, 5, report.GetWinnersHaving(3))
	assert.Equal(t, 4, report.GetWinnersHaving(4))
	assert.Equal(t, 3, report.GetWinnersHaving(5))

	assert.Equal(t, report.String(), "6 5 4 3")
}
