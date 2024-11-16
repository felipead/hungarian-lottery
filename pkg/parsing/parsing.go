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

func LoadPlayerPicksFromFile(fileName string, registry lottery.Registry) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()

	lineNumber := 1
	playerID := 1
	scanner := bufio.NewScanner(file)
	picks := make([]lottery.Number, lottery.NumPicks)

	for scanner.Scan() {
		if err := ParseLine(scanner.Text(), picks); err != nil {
			log.Warnf("skipping line %v because it could not be parsed: %v\n", lineNumber, err)
			lineNumber++
			continue
		}

		registry.RegisterPlayerPicks(playerID, picks)
		lineNumber++
		playerID++
	}

	return scanner.Err()
}

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
