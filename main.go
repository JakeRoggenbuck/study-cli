package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

type Single struct {
	Name     string   `json:"name"`
	Correct  int      `json:"correct"`
	Answers  []string `json:"answers"`
	Question string   `json:"question"`
}

func load_questions(filename string) []Single {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var payload []Single
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return payload
}

func main() {
	rand.Seed(time.Now().UnixNano())
	args := parse_args()

	questions := load_questions(args.filename)

	if args.learn {
		learn(args.amount, questions)
	} else if args.quiz {
		quiz(args.amount, questions)
	} else if args.random {
		random(args.random_prob, args.random_count, questions)
	} else {
		fmt.Println("Welcome to study-cli!")
		fmt.Println("Use --help for more info")
	}
}
