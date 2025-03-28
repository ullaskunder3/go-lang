package project

import "fmt"

func FirstProject() {
	cards := newDec()
	hand, remainingDeck := deal(cards, 5)
	hand.print()
	fmt.Println("====")
	remainingDeck.print()
	fmt.Println("====")

	greeting := "hello I'm ullas"
	fmt.Println([]byte(greeting))

	cards.toString()
	fmt.Println(cards)
	cards.saveToFile("my_cards")

	newcards := newDeckFromFile("my_cards")
	fmt.Println(newcards)

	cards.shuffle()
	cards.print()

	// TEST 39 Assignment

	evenOdd := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, num := range evenOdd {
		if num%2 == 0 {
			fmt.Println("even number", num)
		} else {
			fmt.Println("Odd number", num)

		}
	}

}
