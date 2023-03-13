package main

import (
	"flag"
	"game-of-pig/simulate"
	"game-of-pig/strategyparse"
)

func main() {
	flag.Parse()
	args := flag.Args()
	s, _ := strategyparse.ParseAndValidateArgs(args)
	simulate.Simulate(s)
}
