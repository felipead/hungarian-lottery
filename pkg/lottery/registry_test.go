package lottery

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoPlayerPicks(t *testing.T) {
	registry := NewRegistry()

	report := registry.ProcessLotteryPicks([]Number{11, 22, 33, 44, 55})

	assert.Equal(t, 0, report.GetWinnersOf(5))
	assert.Equal(t, 0, report.GetWinnersOf(4))
	assert.Equal(t, 0, report.GetWinnersOf(3))
	assert.Equal(t, 0, report.GetWinnersOf(2))
}

func TestNoPlayerPicksMatchLotteryPicks(t *testing.T) {
	registry := NewRegistry()

	registry.RegisterPlayerPicks(1, []Number{10, 65, 17, 30, 29})
	registry.RegisterPlayerPicks(2, []Number{89, 20, 12, 15, 02})
	registry.RegisterPlayerPicks(3, []Number{30, 20, 10, 05, 01})

	report := registry.ProcessLotteryPicks([]Number{11, 22, 33, 44, 55})

	assert.Equal(t, 0, report.GetWinnersOf(5))
	assert.Equal(t, 0, report.GetWinnersOf(4))
	assert.Equal(t, 0, report.GetWinnersOf(3))
	assert.Equal(t, 0, report.GetWinnersOf(2))
}

func TestPlayerPicksMatchLotteryPicksRegardlessOfOrder(t *testing.T) {
	registry := NewRegistry()

	registry.RegisterPlayerPicks(1, []Number{00, 00, 00, 00, 00})
	registry.RegisterPlayerPicks(2, []Number{55, 44, 33, 22, 11})
	registry.RegisterPlayerPicks(3, []Number{44, 33, 22, 11, 00})
	registry.RegisterPlayerPicks(4, []Number{33, 22, 11, 00, 00})
	registry.RegisterPlayerPicks(5, []Number{22, 11, 00, 00, 00})
	registry.RegisterPlayerPicks(6, []Number{11, 00, 00, 00, 00})
	registry.RegisterPlayerPicks(7, []Number{00, 00, 00, 00, 00})

	report := registry.ProcessLotteryPicks([]Number{11, 22, 33, 44, 55})

	assert.Equal(t, 1, report.GetWinnersOf(5))
	assert.Equal(t, 1, report.GetWinnersOf(4))
	assert.Equal(t, 1, report.GetWinnersOf(3))
	assert.Equal(t, 1, report.GetWinnersOf(2))
}

func TestSomePlayerPicksMatchLotteryPicksSample1(t *testing.T) {
	registry := NewRegistry()

	registry.RegisterPlayerPicks(1, []Number{00, 22, 00, 44, 00})
	registry.RegisterPlayerPicks(2, []Number{00, 44, 00, 22, 00})
	registry.RegisterPlayerPicks(3, []Number{00, 33, 00, 00, 00})
	registry.RegisterPlayerPicks(4, []Number{00, 00, 00, 00, 00})
	registry.RegisterPlayerPicks(5, []Number{11, 22, 33, 44, 55})
	registry.RegisterPlayerPicks(6, []Number{00, 00, 00, 00, 00})
	registry.RegisterPlayerPicks(7, []Number{11, 00, 33, 44, 00})
	registry.RegisterPlayerPicks(8, []Number{00, 00, 33, 44, 00})
	registry.RegisterPlayerPicks(9, []Number{00, 00, 00, 00, 00})
	registry.RegisterPlayerPicks(10, []Number{00, 00, 00, 00, 00})
	registry.RegisterPlayerPicks(11, []Number{55, 44, 33, 22, 11})
	registry.RegisterPlayerPicks(12, []Number{44, 33, 22, 11, 00})
	registry.RegisterPlayerPicks(13, []Number{00, 22, 11, 00, 44})

	report := registry.ProcessLotteryPicks([]Number{11, 22, 33, 44, 55})

	assert.Equal(t, 2, report.GetWinnersOf(5))
	assert.Equal(t, 1, report.GetWinnersOf(4))
	assert.Equal(t, 2, report.GetWinnersOf(3))
	assert.Equal(t, 3, report.GetWinnersOf(2))
}
