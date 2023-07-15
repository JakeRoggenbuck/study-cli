package main

import (
	"flag"
)

type Args struct {
	mcq          bool
	frq          bool
	learn        bool
	learn_all    bool
	quiz         bool
	amount       int
	filename     string
	random       bool
	random_prob  int
	random_count int
}

func parse_args() Args {
	mcq_flag := flag.Bool("mcq", false, "Set to mcq mode (multiple choice question)")
	frq_flag := flag.Bool("frq", false, "Set to frq mode (free response question)")
	learn_flag := flag.Bool("learn", false, "Set to learn mode")
	learn_all_flag := flag.Bool("learn_all", false, "Set to learn all mode")
	quiz_flag := flag.Bool("quiz", false, "Set to quiz mode")
	amount := flag.Int("amount", 10, "Set amount of questions")
	filename := flag.String("filename", "questions.json", "Set filename for questions")
	random := flag.Bool("random", false, "Set to random review mode")
	random_prob := flag.Int("random_prob", 10, "Set random probability of random quizzing you")
	random_count := flag.Int("random_count", 4, "Set amount of questions to ask user if random prompt happens")

	flag.Parse()

	return Args{*mcq_flag, *frq_flag, *learn_flag, *learn_all_flag, *quiz_flag, *amount, *filename, *random, *random_prob, *random_count}
}
