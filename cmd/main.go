package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/felipead/hungarian-lottery/pkg/lottery"
	"github.com/felipead/hungarian-lottery/pkg/parsing"
)

func inputLoop(registry lottery.Registry) {
	scanner := bufio.NewScanner(os.Stdin)
	picks := make([]lottery.Number, lottery.NumPicks)

	for scanner.Scan() {
		if err := parsing.ParsePicksLine(scanner.Text(), picks); err != nil {
			log.Fatal(err)
		}
		report := registry.ProcessLotteryPicks(picks)
		fmt.Println(report.ToString())
	}
}

func main() {
	if len(os.Args) < 1 {
		log.Fatal("No input file specified")
	}
	fileName := os.Args[0]

	registry := lottery.NewRegistry()
	parsing.LoadPlayerPicksFromFile(fileName, registry)
	fmt.Println("READY")
	inputLoop(registry)
}
