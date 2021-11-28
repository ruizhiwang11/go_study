package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Error("Expercted deck length of 52, but got", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Error("Expected first card of Ace of Spades, but got ", d[0])
	}
	if d[len(d)-1] != "King of Clubs" {
		t.Error("Expected last card of King of Clubs, but got ", d[51])
	}
}

func TestSaveFileLoadFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 52 {
		t.Error("Expected 52 cards but got", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
