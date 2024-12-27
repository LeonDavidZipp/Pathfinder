package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	a "github.com/LeonDavidZipp/Pathfinder/src/algorithm"
	p "github.com/LeonDavidZipp/Pathfinder/src/parsing"
)

func main() {
	// Parse command line flags
	inputFile := flag.String("file", "", "Path to maze input file")
	flag.Parse()

	// Validate input
	if *inputFile == "" {
		log.Fatal("Please provide input file path using -file flag")
	}

	// Read and parse maze
	data, err := p.ReadFile(*inputFile)
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	mp, err := p.ParseMap(data)
	if err != nil {
		log.Fatalf("Error parsing map: %v", err)
	}

	// Solve maze
	solution, err := a.SolveWrapper(mp)
	if err != nil {
		log.Fatalf("Error solving maze: %v", err)
	}

	// Print results
	fmt.Printf("Solution found in %d steps\n", solution.Steps)
	fmt.Printf("Path: %v\n", solution.Route)

	os.Exit(0)
}
