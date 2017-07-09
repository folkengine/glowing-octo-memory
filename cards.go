package main

import (
	"io/ioutil"
	"log"
	"encoding/json"
	"time"
	"math/rand"
)

func init() {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
}

type Card struct {
	Number 		string
	Cardtype 	string
	Response 	string
}

func GoodCards() []Card {
	return openCards("cardsgood.json")
}

func BadCards() []Card {
	return openCards("cardsbad.json")
}

func openCards(filename string) []Card {
	var cards []Card
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(file, &cards)
	if err != nil {
		log.Fatal(err)
	}
	return cards
}

func SampleGoodCards() Card {
	return sampleCards(GoodCards())
}

func SampleBadCards() Card {
	return sampleCards(BadCards())
}

func sampleCards(cards []Card) Card {
	return cards[rand.Intn(len(cards))]
}

func CardsContain(card Card, cards []Card) bool {
	for _, s := range cards {
		if s == card {
			return true
		}
	}
	return false
}