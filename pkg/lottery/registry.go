package lottery

import "runtime"

type Registry interface {
	RegisterPlayer(playerID PlayerID, picks []Number)
	BeReadyForProcessing()
	ProcessLotteryPicks(picks []Number) Report
	ResetLastProcessing()
	// TODO: improve
	HasPlayerPick(playerID PlayerID, pick Number) bool
}

type bucketType = []PlayerID

type registry struct {
	//
	// We are using bucket sort, or bin sort [https://en.wikipedia.org/wiki/Bucket_sort].
	// We create several buckets, or bins, one for each possible lottery number.
	// Assuming the possible lottery numbers are a relatively small set, memory footprint is manageable.
	//
	buckets [MaxNumber]bucketType

	totalPlayers int

	//
	// This is a sparse arrays that counts the matches for all players, where the player ID is the index
	// of the array.
	//
	playerMatches []int
}

func NewRegistryFromNumberAllocation(allocation []int) Registry {
	instance := registry{}

	for i := 0; i < MaxNumber; i++ {
		instance.buckets[i] = make(bucketType, 0, allocation[i])
	}

	return &instance
}

func NewRegistry() Registry {
	instance := registry{}

	for i := 0; i < MaxNumber; i++ {
		instance.buckets[i] = make(bucketType, 0)
	}

	return &instance
}

func (r *registry) RegisterPlayer(playerID PlayerID, picks []Number) {
	for _, pick := range picks {
		index := pick - 1
		r.buckets[index] = append(r.buckets[index], playerID)
	}
	r.totalPlayers++
}

func (r *registry) BeReadyForProcessing() {
	//
	// Since processing of lottery picks is a memory-intensive task, we invoke the Garbage Collector first to
	// clean any unused memory from previous steps (eg: file parsing).
	// In my benchmarks, this dramatically improved the performance of the first processing run by over 200%.
	// Subsequent processing runs were not affected.
	//
	runtime.GC()

	//
	// This may be a large sparse array, so we allocate it beforehand to save a
	// few milliseconds (~ 20ms in my benchmarks).
	// It is faster to reset its elements to zero at the end of processing than to
	// allocate a new array every time.
	//
	r.playerMatches = make([]int, r.totalPlayers)
}

func (r *registry) ProcessLotteryPicks(picks []Number) Report {
	for _, pick := range picks {
		index := pick - 1
		for _, playerID := range r.buckets[index] {
			r.playerMatches[playerID-1]++
		}
	}

	report := NewReport()
	for _, count := range r.playerMatches {
		report.IncrementWinnersHaving(count)
	}

	return report
}

func (r *registry) ResetLastProcessing() {
	//
	// At the end of the last processing, or before processing a new input,
	// this must be invoked.
	//
	for i := 0; i < len(r.playerMatches); i++ {
		r.playerMatches[i] = 0
	}
}

func (r *registry) HasPlayerPick(playerID PlayerID, pick Number) bool {
	index := pick - 1
	for _, id := range r.buckets[index] {
		if id == playerID {
			return true
		}
	}
	return false
}
