package main

import (
    "encoding/json"
    "io/ioutil"
    "log"
	"fmt"
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

func main() {

	questions := load_questions()

	for _, x := range questions {
		fmt.Println(x.Name)
	}
}
