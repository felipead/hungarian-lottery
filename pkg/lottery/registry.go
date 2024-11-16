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
	buckets []bucketType
}

type bucketType = map[PlayerID]bool

func NewRegistry() Registry {
	buckets := make([]bucketType, MaxNumber)
	for i := 0; i < MaxNumber; i++ {
		buckets[i] = make(bucketType)
	}
	return &registry{
		buckets: buckets,
	}
}

func (r *registry) RegisterPlayerPicks(playerID PlayerID, picks []Number) {
	for _, number := range picks {
		bucket := r.buckets[number-1]
		bucket[playerID] = true
	}
}

func (r *registry) ProcessLotteryPicks(picks []Number) Report {
	playerMatches := make(map[PlayerID]int)

	for _, number := range picks {
		bucket := r.buckets[number-1]
		for playerID := range bucket {
			playerMatches[playerID] += 1
		}
	}

	report := NewReport()
	for _, count := range playerMatches {
		report.IncrementWinnersOf(count)
	}
	return report
}
