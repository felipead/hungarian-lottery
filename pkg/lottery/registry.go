package lottery

import "runtime"

// Registry registers the lottery players and their picks. It also processes the lottery picks.
type Registry interface {

	// RegisterPlayer registers a player and its numeric picks.
	// The playerID is a unique, sequential number starting from 1. If these players are loaded from a file, this could
	// be the line number.
	// The picks is a slice containing NumPicks numbers.
	RegisterPlayer(playerID PlayerID, picks []Number)

	// BeReadyForProcessing carries optimizations necessary for correct and efficient processing of lottery picks.
	// Should be invoked right before start accepting lottery picks as input.
	BeReadyForProcessing()

	// ProcessLotteryPicks processes the lottery picks from input, and returns a [lottery.Report]. Does the magic.
	ProcessLotteryPicks(picks []Number) Report

	// ResetLastProcessing cleans up and resets the state of this registry from last processing of lottery picks.
	// This method MUST be invoked at the end of the last processing, or before processing a new input.
	// Since clean-up could take some time, it was designed as a separate method so that the [lottery.Report] returned
	// from [Registry.ProcessLotteryPicks] could be rendered as soon as possible.
	ResetLastProcessing()

	// HasPlayerPick determines if the player ID has picked the given number. Used for testing.
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
	// of the array. This allows for great efficiency gains when querying the result of a given lottery pick,
	// since the counts for each player can be accessed by direct array access.
	//
	// In local benchmarks, using sparse arrays was responsible from a significant reduction from ~450ms to ~30ms in
	// processing time, compared to hash maps. The downside of this approach is that much more memory is used compared
	// to hash maps, since each player must fill an index in the array, regardless if it has wins or not. Because
	// there are fewer wins than bets, the array is sparse.
	//
	playerMatches []int
}

// NewRegistryFromNumberAllocation creates a new lottery registry when the number allocations are already known.
// This allows for player picks to be efficiently put into buckets, without the wasteful overhead of array resizing
// during slice appends when the capacity of the array is not known.
func NewRegistryFromNumberAllocation(allocation []int) Registry {
	instance := registry{}

	for i := 0; i < MaxNumber; i++ {
		instance.buckets[i] = make(bucketType, 0, allocation[i])
	}

	return &instance
}

// NewRegistry is only used for testing purposes. For production, because of efficiency concerns,
// [NewRegistryFromNumberAllocation] should be used instead.
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
	//
	// We use bucket sorting. For each number N picked by the lottery, we can efficiently query
	// all players that also picked N by accessing the bucket whose index is N-1.
	//
	// We then count the player matches in a sparse array. This count can range from 0 to NumPicks (eg: 0 to 5),
	// meaning how many matches that player got from the lottery picks.
	//
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
