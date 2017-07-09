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

type Cards struct {
	Collection []Card
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