package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

type SingleMCQ struct {
	Name     string   `json:"name"`
	Correct  int      `json:"correct"`
	Answers  []string `json:"answers"`
	Question string   `json:"question"`
}

type SingleFRQ struct {
	Name     string `json:"name"`
	Answer   string `json:"answer"`
	Question string `json:"question"`
}

type QuestionType int


const (
	FRQ QuestionType = iota + 1
	MCQ
)

type Single struct {
	FRQ []SingleFRQ
	MCQ []SingleFRQ
	Type QuestionType
}

func load_questions[T SingleFRQ | SingleMCQ](filename string) []T {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var payload []T
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return payload
}

func arg_match[T SingleFRQ | SingleMCQ](args Args, questions []T) {
	var question_arr Single
	if args.frq {
		question_arr := Single{FRQ: questions, MCQ: nil, Type: FRQ}
	} else {
		question_arr := Single{FRQ: nil, MCQ: questions, Type: MCQ}
	}


	if args.learn {
		learn_random(args.amount, questions)
	} else if args.learn_all {
		learn_linear(args.amount, question_arr)
	} else if args.quiz {
		quiz(args.amount, question_arr)
	} else if args.random {
		random(args.random_prob, args.random_count, question_arr)
	} else {
		fmt.Println("Welcome to study-cli!")
		fmt.Println("Use --help for more info")
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	args := parse_args()

	if args.frq {
		questions := load_questions[SingleFRQ](args.filename)
		arg_match[SingleFRQ](args, questions)
	} else{
		questions := load_questions[SingleMCQ](args.filename)
		arg_match[SingleMCQ](args, questions)
	}

}
