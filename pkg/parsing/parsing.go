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

func ParsePicksLine(line string, buffer []lottery.Number) error {
	fields := strings.Fields(line)
	if len(fields) > len(buffer) {
		// TODO: better error message
		return errors.New("too many fields")
	}

	for i, field := range fields {
		parsed, err := strconv.ParseInt(field, 10, lottery.NumberBitSize)
		if err != nil {
			return err
		}
		number := lottery.Number(parsed)

		if number > lottery.MaxNumber {
			// TODO: better error message
			return errors.New("lottery number too big")
		}

		buffer[i] = number
	}
	return nil
}

func ParseLotteryRegistryFile(fileName string) lottery.Registry {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = file.Close()
	}()

	registry := lottery.NewRegistry()
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

	return registry
}
