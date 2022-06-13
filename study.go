package main

import (
	"fmt"
	"github.com/gookit/color"
	"math/rand"
)

func learn(amount int, questions []Single) {
	var ready string

	for i := 0; i < amount; i++ {
		rand_num := rand.Intn(len(questions))

		picked := questions[rand_num]
		fmt.Println(picked.Question)
		color.Green.Println(picked.Answers[picked.Correct][3:] + "\n")

		fmt.Print("Reviewed? ")
		fmt.Scan(&ready)
		fmt.Print("\n")
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

		fmt.Print("\n")
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

		fmt.Print("\n\n")
	}
}

func random(random_prob int, random_count int, questions []Single) {
	rand_num := rand.Intn(100)
	quiz_user := rand_num < random_prob

	if quiz_user {
		quiz(random_count, questions)
	}
}
