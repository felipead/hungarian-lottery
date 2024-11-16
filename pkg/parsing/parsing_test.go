package parsing

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/felipead/hungarian-lottery/pkg/lottery"
)

func TestParseLineSuccessfully(t *testing.T) {
	var err error
	picks := make([]lottery.Number, 5)

	err = ParseLine("88 28 43 72 14", picks)
	assert.NoError(t, err)
	assert.Equal(t, []lottery.Number{88, 28, 43, 72, 14}, picks)

	err = ParseLine("7 64 80 90 58", picks)
	assert.NoError(t, err)
	assert.Equal(t, []lottery.Number{7, 64, 80, 90, 58}, picks)
}

func TestParseLineFailIfSlightlyAbove90(t *testing.T) {
	picks := make([]lottery.Number, 5)

	err := ParseLine("88 28 91 72 14", picks)
	assert.ErrorIs(t, err, ErrPickedNumberOutOfRange)
}

func TestParseLineFailIfMuchAbove90(t *testing.T) {
	picks := make([]lottery.Number, 5)

	err := ParseLine("88 28 1000 72 14", picks)
	assert.ErrorIs(t, err, ErrPickedNumberOutOfRange)
}

func TestParseLineFailIfZero(t *testing.T) {
	picks := make([]lottery.Number, 5)

	err := ParseLine("88 0 16 72 14", picks)
	assert.ErrorIs(t, err, ErrPickedNumberOutOfRange)
}

func TestParseLineFailIfNegative(t *testing.T) {
	picks := make([]lottery.Number, 5)

	err := ParseLine("88 -17 16 72 14", picks)
	assert.ErrorIs(t, err, ErrPickedNumberOutOfRange)
}

func TestParseLineFailIfHasLessFields(t *testing.T) {
	picks := make([]lottery.Number, 5)

	err := ParseLine("88 28 91 72", picks)
	assert.ErrorIs(t, err, ErrInvalidQuantityOfPickedNumbers)
}

func TestParseLineFailIfHasMoreFields(t *testing.T) {
	picks := make([]lottery.Number, 5)

	err := ParseLine("88 28 91 72 16 24", picks)
	assert.ErrorIs(t, err, ErrInvalidQuantityOfPickedNumbers)
}

func TestParseLineFailIfLineIsEmpty(t *testing.T) {
	picks := make([]lottery.Number, 5)

	err := ParseLine("", picks)
	assert.ErrorIs(t, err, ErrInvalidQuantityOfPickedNumbers)
}

func TestLoadPlayerPicksFromFile(t *testing.T) {
	registry := lottery.NewRegistry()
	err := LoadPlayerPicksFromFile("testdata/1k-players.txt", registry)
	assert.NoError(t, err)

	assert.True(t, registry.HasPlayerPick(14, 12))
	assert.True(t, registry.HasPlayerPick(14, 83))
	assert.True(t, registry.HasPlayerPick(14, 73))
	assert.True(t, registry.HasPlayerPick(14, 26))
	assert.True(t, registry.HasPlayerPick(14, 32))
	assert.False(t, registry.HasPlayerPick(14, 10))

	assert.True(t, registry.HasPlayerPick(888, 11))
	assert.True(t, registry.HasPlayerPick(888, 7))
	assert.True(t, registry.HasPlayerPick(888, 24))
	assert.True(t, registry.HasPlayerPick(888, 48))
	assert.True(t, registry.HasPlayerPick(888, 29))
	assert.False(t, registry.HasPlayerPick(888, 10))

	assert.True(t, registry.HasPlayerPick(535, 7))
	assert.True(t, registry.HasPlayerPick(535, 35))
	assert.True(t, registry.HasPlayerPick(535, 65))
	assert.True(t, registry.HasPlayerPick(535, 47))
	assert.True(t, registry.HasPlayerPick(535, 11))
	assert.False(t, registry.HasPlayerPick(535, 10))
}

func TestLoadPlayerPicksFromFileNoNewlineAtEnd(t *testing.T) {
	registry := lottery.NewRegistry()
	err := LoadPlayerPicksFromFile("testdata/1k-players_no-newline-at-end.txt", registry)
	assert.NoError(t, err)

	assert.True(t, registry.HasPlayerPick(14, 12))
	assert.True(t, registry.HasPlayerPick(14, 83))
	assert.True(t, registry.HasPlayerPick(14, 73))
	assert.True(t, registry.HasPlayerPick(14, 26))
	assert.True(t, registry.HasPlayerPick(14, 32))
	assert.False(t, registry.HasPlayerPick(14, 10))

	assert.True(t, registry.HasPlayerPick(888, 11))
	assert.True(t, registry.HasPlayerPick(888, 7))
	assert.True(t, registry.HasPlayerPick(888, 24))
	assert.True(t, registry.HasPlayerPick(888, 48))
	assert.True(t, registry.HasPlayerPick(888, 29))
	assert.False(t, registry.HasPlayerPick(888, 10))

	assert.True(t, registry.HasPlayerPick(535, 7))
	assert.True(t, registry.HasPlayerPick(535, 35))
	assert.True(t, registry.HasPlayerPick(535, 65))
	assert.True(t, registry.HasPlayerPick(535, 47))
	assert.True(t, registry.HasPlayerPick(535, 11))
	assert.False(t, registry.HasPlayerPick(535, 10))
}
