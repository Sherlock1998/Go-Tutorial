package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 20 {
		t.Errorf("Expected deck length of 20. but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be ace of spades. but got %v", d[0])
	}

	if d[len(d)-1] != "Five of Clubs" {
		t.Errorf("Expected first card to be Five of Clubs. but got %v", d[0])
	}
}

func TestSaveAndLoadFuncs(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedFile := newDeckFromFile("_decktesting")

	if len(loadedFile) != 20 {
		t.Errorf("Expected deck length of 20. but got %v", len(loadedFile))
	}
	os.Remove("_decktesting")
}
