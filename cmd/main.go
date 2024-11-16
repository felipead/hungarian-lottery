package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/felipead/hungarian-lottery/pkg/lottery"
	"github.com/felipead/hungarian-lottery/pkg/parsing"
)

func inputLoop(registry lottery.Registry, debugMode bool) {
	scanner := bufio.NewScanner(os.Stdin)
	picks := make([]lottery.Number, lottery.NumPicks)

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatalf("I/O error: %v", err)
		}

		if err := parsing.ParseLine(scanner.Text(), picks); err != nil {
			log.Fatalf("could not parse input: %v", err)
		}

		var start time.Time
		if debugMode {
			start = time.Now()
		}

		report := registry.ProcessLotteryPicks(picks)
		fmt.Println(report.ToString())

		if debugMode {
			elapsed := time.Since(start)
			log.Infof("took: %v ms", elapsed.Milliseconds())
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("no input file specified")
	}
	fileName := os.Args[1]

	debugMode := false
	if len(os.Args) > 2 {
		debugMode = os.Args[2] == "--debug"
	}

	log.Infof("loading input file %v", fileName)

	registry := lottery.NewRegistry()
	if err := parsing.LoadPlayerPicksFromFile(fileName, registry); err != nil {
		log.Fatalf("unable to load file: %v", err)
	}

	fmt.Println("READY")
	inputLoop(registry, debugMode)
}
