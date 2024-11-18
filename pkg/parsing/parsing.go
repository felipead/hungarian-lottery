package parsing

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/felipead/hungarian-lottery/pkg/lottery"
)

// LoadFile parses a file and fills the player picks into a new [lottery.Registry] instance.
// For efficiency purposes, first the file is traversed so that we can determine the allocation
// necessary to represent the player picks. Then, the file is read again to register the player picks.
// This was done to avoid the unnecessary overhead from resizing the underlying arrays during slice appends.
func LoadFile(fileName string) (lottery.Registry, error) {
	allocation, err := determineNumberAllocation(fileName)
	if err != nil {
		return nil, err
	}

	registry := lottery.NewRegistryFromNumberAllocation(allocation)

	if err = registerPlayers(fileName, registry); err != nil {
		return nil, err
	}

	return registry, nil
}

func determineNumberAllocation(fileName string) ([]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer func() { _ = file.Close() }()

	numberAllocation := make([]int, lottery.MaxNumber)

	scanner := bufio.NewScanner(file)
	picks := make([]lottery.Number, lottery.NumPicks)

	for scanner.Scan() {
		if err = ParseLine(scanner.Text(), picks); err != nil {
			continue
		}

		for _, pick := range picks {
			numberAllocation[pick-1]++
		}
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return numberAllocation, nil
}

func registerPlayers(fileName string, registry lottery.Registry) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	lineNumber := 1
	var playerID lottery.PlayerID = 1
	scanner := bufio.NewScanner(file)
	picks := make([]lottery.Number, lottery.NumPicks)

	for scanner.Scan() {
		if err = ParseLine(scanner.Text(), picks); err != nil {
			log.Warnf("skipping line %v because it could not be parsed: %v", lineNumber, err)
			lineNumber++
			continue
		}

		registry.RegisterPlayer(playerID, picks)
		lineNumber++
		playerID++
	}

	return scanner.Err()
}

// ParseLine parses a textual line representing the picked lottery numbers.
// The numbers must be separated by whitespace, as defined by [unicode.IsSpace].
func ParseLine(line string, picks []lottery.Number) error {
	fields := strings.Fields(line)
	if len(fields) != len(picks) {
		return ErrInvalidQuantityOfPickedNumbers
	}

	for i := 0; i < len(picks); i++ {
		field := fields[i]
		parsed, err := strconv.ParseInt(field, 10, lottery.NumberBitSize)
		if err != nil {
			if errors.Is(err, strconv.ErrRange) {
				return ErrPickedNumberOutOfRange
			}
			return err
		}
		pick := lottery.Number(parsed)

		if pick > lottery.MaxNumber || pick < 1 {
			return ErrPickedNumberOutOfRange
		}

		picks[i] = pick
	}
	return nil
}
