package main

import (
    "encoding/json"
    "io/ioutil"
    "log"
	"fmt"
	"math/rand"
	"time"
	"github.com/gookit/color"
	"flag"
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

func indexOf(word string, data [4]string) int {
    for k, v := range data {
        if word == v {
            return k
        }
    }
    return -1
}

func learn(amount int, questions []Single) {
	for i := 0; i < amount; i++ {
		rand_num := rand.Intn(len(questions))

		picked := questions[rand_num]
		fmt.Println(picked.Question)
		fmt.Println(picked.Answers[picked.Correct][3:] + "\n")
	}
}

func quiz(amount int, questions []Single) {
	for i := 0; i < amount; i++ {
		fmt.Println("--------------------------------------------------------------------------------")
		rand_num := rand.Intn(len(questions))

		var input string
		picked := questions[rand_num]
		color.Blue.Println(picked.Question)
		for _, ans := range picked.Answers {
			fmt.Println(ans)
		}
		
		fmt.Print("Answer: ")
		fmt.Scan(&input)

		answers := [4]string{"A", "B", "C", "D"}

		if input == answers[picked.Correct] {
			color.Green.Print("Correct\n\n")
		} else {
			// Highlight correct and incorrect answers
			index_of_wrong_ans := indexOf(input, answers)
			index_of_right_ans := picked.Correct
			for c, ans := range picked.Answers {
				if c == index_of_wrong_ans {
					color.Red.Println(ans)
				} else if c == index_of_right_ans {
					color.Green.Println(ans)
				} else {
					fmt.Println(ans)
				}
			}
			fmt.Print("\n")

			fmt.Print("Reviewed? ")
			fmt.Scan(&input)
		}

		fmt.Print("\n\n\n")
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	questions := load_questions()
	
	learn_flag := flag.Bool("learn", false, "Set to learn mode")
	quiz_flag := flag.Bool("quiz", false, "Set to quiz mode")
	amount := flag.Int("amount", 10, "Set amount of questions")

	flag.Parse()

	if *learn_flag {
		learn(*amount, questions)
	} else if *quiz_flag {
		quiz(*amount, questions)
	}
}
