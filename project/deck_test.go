package project

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDec()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
	if d[0] != "Ace of ♠️" {
		t.Errorf("Expected first card to Ace of spades, but got %v", d[0])
	}
	if d[len(d)-1] != "Four of ♣️" {
		t.Errorf("Expected last card to Four of clubs, but got %v", d[len(d)-1])
	}
}
func TestDeckSaveAndReCheckDeck(t *testing.T) {
	os.Remove("_deck_test")
	deck := newDec()
	deck.saveToFile("_deck_test")

	loadedDeck := newDeckFromFile("_deck_test")
	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(deck))
	}
}
