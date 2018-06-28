package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 but got %v", len(d))
	}

	if d[0] != "superman is one" {
		t.Errorf("Expected the is first card is superman is one but got %v", d[0])
	}

	if d[len(d)-1] != "cyborg is four" {
		t.Errorf("Expected the last card is cyborg is four but got %v", d[len(d)-1])

	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {

	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards but got a %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
