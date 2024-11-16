package main

import (
	"bufio"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/felipead/hungarian-lottery/pkg/lottery"
	"github.com/felipead/hungarian-lottery/pkg/parsing"
)

func inputLoop(registry lottery.Registry) {
	scanner := bufio.NewScanner(os.Stdin)
	picks := make([]lottery.Number, lottery.NumPicks)

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatalf("I/O error: %v", err)
		}

		if err := parsing.ParseLine(scanner.Text(), picks); err != nil {
			log.Fatalf("could not parse input: %v", err)
		}
		report := registry.ProcessLotteryPicks(picks)
		fmt.Println(report.ToString())
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("no input file specified")
	}
	fileName := os.Args[1]
	log.Infof("loading input file %v", fileName)

	registry := lottery.NewRegistry()
	if err := parsing.LoadPlayerPicksFromFile(fileName, registry); err != nil {
		log.Fatalf("unable to load file: %v", err)
	}

	fmt.Println("READY")
	inputLoop(registry)
}
