package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"math/rand"
	"time"

)

type deck [] string

func newDeck() deck {

	cards := deck{}
	cardSuits := [] string {"Spades", "Diamonds", "Hearts","Clubs"}
	cardValues := [] string {"Ace", "Two", "Three", "Four", "Five", "Six", "Seven"}

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append (cards, cardValue+" of "+ cardSuit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveFile(filename string) error {
  return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)

}

func (d deck) shuffle () {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		p := r.Intn(len(d)-1)

		d[i], d[p] = d[p], d[i]
	}
}