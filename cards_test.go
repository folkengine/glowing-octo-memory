package main

import (
	"testing"
)

func TestBadCards(t *testing.T) {
	expecting := "4000111111111115"
	badcards := BadCards()
	if badcards[0].Number != expecting {
		t.Errorf("Expecting %s instead of %s",
			expecting, badcards[0].Number)
	}
}

func TestGoodCards(t *testing.T) {
	expecting := "378282246310005"
	goodcards := GoodCards()
	if goodcards[0].Number != expecting {
		t.Errorf("Expecting %s instead of %s",
			expecting, goodcards[0].Number)
	}
}

func TestSampleGoodCards(t *testing.T) {
	card := SampleGoodCards()
	if !CardsContain(card, GoodCards()) {
		t.Errorf("Unable to find %v", card)
	}
}

func TestSampleBadCards(t *testing.T) {
	card := SampleBadCards()
	if !CardsContain(card, BadCards()) {
		t.Errorf("Unable to find %v", card)
	}
}
