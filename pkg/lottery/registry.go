package lottery

type Registry interface {
	RegisterPlayerPicks(playerID PlayerID, picks []Number)
	ProcessLotteryPicks(picks []Number) Report
}

type registry struct {
	//
	// We are using bucket sort, or bin sort [https://en.wikipedia.org/wiki/Bucket_sort].
	// We create several buckets, or bins, one for each possible lottery number.
	// Assuming the possible lottery numbers are a relatively small set, memory footprint is manageable.
	//
	buckets []bucket
}

type bucket = map[PlayerID]bool

func NewRegistry() Registry {
	buckets := make([]bucket, MaxNumber)
	for i := 0; i < MaxNumber; i++ {
		buckets[i] = make(bucket)
	}
	return &registry{
		buckets: buckets,
	}
}

func (r *registry) RegisterPlayerPicks(playerID PlayerID, picks []Number) {
	for _, number := range picks {
		r.buckets[number][playerID] = true
	}
}

func (r *registry) ProcessLotteryPicks(picks []Number) Report {
	playerMatches := make(map[PlayerID]int)

	for _, pick := range picks {
		for playerID := range r.buckets[pick] {
			playerMatches[playerID] += 1
		}
	}

	report := NewReport()

	for _, count := range playerMatches {
		report.IncrementWinners(count)
	}

	return report
}
