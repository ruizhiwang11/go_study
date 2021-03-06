package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven",
		"Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardValues {
		for _, value := range cardSuits {
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(path string) error {
	return ioutil.WriteFile(path, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Err," + err.Error())
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	// var memo []int
	// for index, _ := range d {
	// 	newPosition := rand.Intn(len(d) - 1)
	// 	for _, value := range memo {
	// 		if value == index || value == newPosition {
	// 			continue
	// 		}
	// 		d[index], d[newPosition] = d[newPosition], d[index]
	// 		memo = append(memo, index)
	// 		memo = append(memo, newPosition)
	// 	}
	// }
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for index := range d {
		newPosition := r.Intn(len(d) - 1)
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}
