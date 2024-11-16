package parsing

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/felipead/hungarian-lottery/pkg/lottery"
)

func TestParsePicksLineSuccessfully(t *testing.T) {
	var err error
	picks := make([]lottery.Number, 5)

	err = ParsePicksLine("88 28 43 72 14", picks)
	assert.NoError(t, err)
	assert.Equal(t, []lottery.Number{88, 28, 43, 72, 14}, picks)

	err = ParsePicksLine("7 64 80 90 58", picks)
	assert.NoError(t, err)
	assert.Equal(t, []lottery.Number{7, 64, 80, 90, 58}, picks)
}

func TestParsePicksLineFailIfSlightlyAbove90(t *testing.T) {
	picks := make([]lottery.Number, 5)

	err := ParsePicksLine("88 28 91 72 14", picks)
	assert.ErrorIs(t, err, ErrPickedNumberOutOfRange)
}

func TestParsePicksLineFailIfMuchAbove90(t *testing.T) {
	picks := make([]lottery.Number, 5)

	err := ParsePicksLine("88 28 1000 72 14", picks)
	assert.ErrorIs(t, err, ErrPickedNumberOutOfRange)
}

func TestParsePicksLineFailIfZero(t *testing.T) {
	picks := make([]lottery.Number, 5)

	err := ParsePicksLine("88 0 16 72 14", picks)
	assert.ErrorIs(t, err, ErrPickedNumberOutOfRange)
}

func TestParsePicksLineFailIfNegative(t *testing.T) {
	picks := make([]lottery.Number, 5)

	err := ParsePicksLine("88 -17 16 72 14", picks)
	assert.ErrorIs(t, err, ErrPickedNumberOutOfRange)
}

func TestParsePicksLineFailIfHasLessFields(t *testing.T) {
	picks := make([]lottery.Number, 5)

	err := ParsePicksLine("88 28 91 72", picks)
	assert.ErrorIs(t, err, ErrInvalidQuantityOfPicks)
}

func TestParsePicksLineFailIfHasMoreFields(t *testing.T) {
	picks := make([]lottery.Number, 5)

	err := ParsePicksLine("88 28 91 72 16 24", picks)
	assert.ErrorIs(t, err, ErrInvalidQuantityOfPicks)
}

func TestParsePicksLineFailIfLineIsEmpty(t *testing.T) {
	picks := make([]lottery.Number, 5)

	err := ParsePicksLine("", picks)
	assert.ErrorIs(t, err, ErrInvalidQuantityOfPicks)
}
