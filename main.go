package main

import (
    "encoding/json"
    "io/ioutil"
    "log"
	"fmt"
	"math/rand"
	"time"
)

type Single struct {
	Name string `json:"name"`
	Correct int `json:"correct"`
	Answers []string `json:"answers"`
	Question string `json:"question"`
}

func load_questions() []Single {
	content, err := ioutil.ReadFile("./config.json")
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

func learn(amount int, questions []Single) {
	for i := 0; i < amount; i++ {
		rand_num := rand.Intn(len(questions))

		picked := questions[rand_num]
		fmt.Println(picked.Question)
		fmt.Println(picked.Answers[picked.Correct][3:] + "\n")
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	questions := load_questions()
	
	learn(10, questions)
}
