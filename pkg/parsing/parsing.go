package parsing

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/felipead/hungarian-lottery/pkg/lottery"
)

func LoadPlayerPicksFromFile(fileName string, registry lottery.Registry) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = file.Close()
	}()

	scanner := bufio.NewScanner(file)
	picks := make([]lottery.Number, lottery.NumPicks)
	playerID := 1

	for scanner.Scan() {
		if err := ParsePicksLine(scanner.Text(), picks); err != nil {
			log.Fatal(err)
		}

		registry.RegisterPlayerPicks(playerID, picks)
		playerID++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func ParsePicksLine(line string, picks []lottery.Number) error {
	fields := strings.Fields(line)
	if len(fields) != len(picks) {
		// TODO: better error message
		return errors.New("invalid number of many fields")
	}

	for i, field := range fields {
		parsed, err := strconv.ParseInt(field, 10, lottery.NumberBitSize)
		if err != nil {
			return err
		}
		pick := lottery.Number(parsed)

		if pick > lottery.MaxNumber {
			// TODO: better error message
			return errors.New("lottery pick is too big")
		}

		picks[i] = pick
	}
	return nil
}
