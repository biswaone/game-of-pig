package main

import (
	"flag"
	"game-of-pig/simulate"
	"game-of-pig/strategyparse"
	"log"
)

func main() {
	flag.Parse()
	args := flag.Args()
	s, err := strategyparse.ParseAndValidateArgs(args)
	if err != nil {
		log.Fatal(err)
	}
	simulate.Simulate(s)
}
