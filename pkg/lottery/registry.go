package lottery

type Registry interface {
	RegisterPlayerPicks(playerID PlayerID, picks []Number)
	ProcessLotteryPicks(picks []Number) Report
}

type bucket = map[PlayerID]bool

type registry struct {
	buckets []bucket
}

func NewRegistry() Registry {
	//
	// We are using bucket sort, or bin sort [https://en.wikipedia.org/wiki/Bucket_sort].
	// We create several buckets, or bins, one for each possible lottery number.
	// Assuming the possible lottery numbers are small set, memory footprint is manageable.
	//
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
		report.IncrementMatches(count)
	}

	return report
}
