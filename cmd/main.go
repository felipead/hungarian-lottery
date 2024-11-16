package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// get file name from command line args (first argument)

// read file, line by line

// each line number corresponds to a player ID - eg: 1 to 10M

// create 90 buckets (1 to 90), which are hash maps corresponding to the players IDs that gambled on these numbers

// iterate over the buckets, counting in a hashmap the occurrences of each player ID

const MaxLotteryNumber = 90
const LotteryNumberBitSize = 8
const NumLotteryPicks = 5

type LotteryNumber = uint8
type PlayerID = int

type Bucket = map[PlayerID]bool

func initializeBuckets() []map[PlayerID]bool {
	//
	// We are using bucket sort, or bin sort [https://en.wikipedia.org/wiki/Bucket_sort].
	// We create several buckets, or bins, one for each possible lottery number.
	// Assuming the possible lottery numbers are small set, memory footprint is manageable.
	//
	buckets := make([]Bucket, MaxLotteryNumber)
	for i := 0; i < MaxLotteryNumber; i++ {
		buckets[i] = make(Bucket)
	}
	return buckets
}

func parseLine(line string, buffer []LotteryNumber) error {
	fields := strings.Fields(line)
	if len(fields) > len(buffer) {
		// TODO: better error message
		return errors.New("too many fields")
	}

	for i, field := range fields {
		parsed, err := strconv.ParseInt(field, 10, LotteryNumberBitSize)
		if err != nil {
			return err
		}
		number := LotteryNumber(parsed)

		if number > MaxLotteryNumber {
			// TODO: better error message
			return errors.New("lottery number too big")
		}

		buffer[i] = number
	}
	return nil
}

func parseFile(fileName string, buckets []Bucket) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	playerPicks := make([]LotteryNumber, NumLotteryPicks)
	playerID := 1

	for scanner.Scan() {
		if err := parseLine(scanner.Text(), playerPicks); err != nil {
			log.Fatal(err)
		}

		for _, number := range playerPicks {
			buckets[number][playerID] = true
		}

		playerID++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func processLotteryPicks(buckets []Bucket, lotteryPicks []LotteryNumber) []int {
	playerMatches := make(map[PlayerID]int)

	for _, pick := range lotteryPicks {
		for playerID := range buckets[pick] {
			playerMatches[playerID] += 1
		}
	}

	const ResultSize = NumLotteryPicks - 1
	result := make([]int, ResultSize)

	for _, count := range playerMatches {
		index := count - 2
		if index >= 0 {
			result[index]++
		}
	}

	return result
}

func printResult(result []int) {
	for i := 0; i < len(result); i++ {
		last := len(result) - 1
		if i != last {
			fmt.Printf("%v ", result[i])
		} else {
			fmt.Printf("%v\n", result[i])
		}
	}
}

func inputLoop(buckets []Bucket) {
	scanner := bufio.NewScanner(os.Stdin)
	lotteryPicks := make([]LotteryNumber, NumLotteryPicks)

	for scanner.Scan() {
		if err := parseLine(scanner.Text(), lotteryPicks); err != nil {
			log.Fatal(err)
		}
		result := processLotteryPicks(buckets, lotteryPicks)
		printResult(result)
	}
}

func main() {
	// TODO: validate if at least one argument was given
	fileName := os.Args[0]

	buckets := initializeBuckets()
	parseFile(fileName, buckets)
	fmt.Println("READY")
	inputLoop(buckets)
}
