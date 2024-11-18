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

func main() {
	fileName, debugMode := parseArgs()

	log.Infof("loading input file %v", fileName)
	registry, err := parsing.LoadFile(fileName)
	if err != nil {
		log.Fatalf("unable to load file: %v", err)
	}

	registry.BeReadyForProcessing()
	fmt.Println("READY")

	inputLoop(registry, debugMode)
}

func parseArgs() (fileName string, debugMode bool) {
	if len(os.Args) < 2 {
		log.Fatalf("no input file specified")
	}

	fileName = os.Args[1]
	if len(os.Args) > 2 {
		debugMode = os.Args[2] == "--debug"
	}

	return fileName, debugMode
}

func inputLoop(registry lottery.Registry, debugMode bool) {
	scanner := bufio.NewScanner(os.Stdin)
	picks := make([]lottery.Number, lottery.NumPicks)

	for scanner.Scan() {
		line := scanner.Text()
		if err := parsing.ParseLine(line, picks); err != nil {
			log.Fatalf("could not parse input: %v — '%v'", err, line)
		}

		var start time.Time
		if debugMode {
			start = time.Now()
		}

		report := registry.ProcessLotteryPicks(picks)
		fmt.Println(report.String())

		if debugMode {
			elapsed := time.Since(start)
			log.Infof("took: %v ms", elapsed.Milliseconds())
		}

		registry.ResetLastProcessing()
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("I/O error: %v", err)
	}
}
