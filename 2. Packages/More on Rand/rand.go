package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(2) // Try changing this number!
	answers := []string{
		"It is certain",
		"It is decidedly so",
		"Without a doubt",
		"Yes definitely",
		"You may rely on it",
		"As I see it yes",
		"Most likely",
		"Outlook good",
		"Yes",
		"Signs point to yes",
		"Reply hazy try again",
		"Ask again later",
		"Better not tell you now",
		"Cannot predict now",
		"Concentrate and ask again",
		"Don't count on it",
		"My reply is no",
		"My sources say no",
		"Outlook not so good",
		"Very doubtful",
	}
	fmt.Println("len(answers) is", len(answers))
	fmt.Println("rand.Intn(len(answers)) is", rand.Intn(len(answers)))
	fmt.Println("answers[0] is", answers[0])
	fmt.Println("answers[1] is", answers[1])
	fmt.Println("Magic 8-Ball says:", answers[rand.Intn(len(answers))])
}
