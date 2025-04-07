package project

import "fmt"

type bot interface{ getGreeting() string }

type englishBot struct{}
type spanishBot struct{}

func FourthProject() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) { fmt.Println(b.getGreeting()) }

func (eb englishBot) getGreeting() string { return "hello" }
func (sb spanishBot) getGreeting() string { return "hola" }
