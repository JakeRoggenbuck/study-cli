package main

import (
	"flag"
)

type Args struct {
	learn bool
	quiz bool
	amount int
}

func parse_args() Args {
	learn_flag := flag.Bool("learn", false, "Set to learn mode")
	quiz_flag := flag.Bool("quiz", false, "Set to quiz mode")
	amount := flag.Int("amount", 10, "Set amount of questions")

	flag.Parse()

	return Args { *learn_flag, *quiz_flag, *amount }
}
