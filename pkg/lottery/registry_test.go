package lottery

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoPlayerPicks(t *testing.T) {
	registry := NewRegistry()

	report := registry.ProcessLotteryPicks([]Number{11, 22, 33, 44, 55})

	assert.Equal(t, report.GetWinners(5), 0)
	assert.Equal(t, report.GetWinners(4), 0)
	assert.Equal(t, report.GetWinners(3), 0)
	assert.Equal(t, report.GetWinners(2), 0)
}

func TestSomePlayerPicksMatchLotteryPicks(t *testing.T) {
	registry := NewRegistry()

	registry.RegisterPlayerPicks(1, []Number{00, 22, 00, 44, 00})
	registry.RegisterPlayerPicks(2, []Number{00, 22, 00, 44, 00})
	registry.RegisterPlayerPicks(3, []Number{00, 00, 33, 00, 00})
	registry.RegisterPlayerPicks(4, []Number{00, 00, 00, 00, 00})
	registry.RegisterPlayerPicks(5, []Number{11, 22, 33, 44, 55})

	report := registry.ProcessLotteryPicks([]Number{11, 22, 33, 44, 55})

	assert.Equal(t, report.GetWinners(5), 1)
	assert.Equal(t, report.GetWinners(4), 0)
	assert.Equal(t, report.GetWinners(3), 0)
	assert.Equal(t, report.GetWinners(2), 2)
}
