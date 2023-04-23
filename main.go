package main

import (
	"fmt"
	"strings"

	"github.com/kristof1345/cards"
)

type Hand []cards.Card

// separate them with a comma
func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}

func (h Hand) DealerString() string {
	return h[0].String() + ", **HIDDEN**"
}

func (h Hand) MinScore() int {
	score := 0
	for _, c := range h {
		score += min(int(c.Rank), 10)
	}
	return score
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	deck := cards.New(cards.Deck(3), cards.Shuffle)
	var card cards.Card
	var player, dealer Hand
	for i := 0; i < 2; i++ {
		for _, hand := range []*Hand{&player, &dealer} {
			card, deck = draw(deck)
			*hand = append(*hand, card)
		}
	}
	var input string
	for input != "s" {
		fmt.Println("Player:", player)
		fmt.Println("Dealer:", dealer.DealerString())
		fmt.Println("What will you do? (h)it, (s)tand")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			card, deck = draw(deck)
			player = append(player, card)
		}
	}
	fmt.Println("==FINAL HANDS==")
	fmt.Println("Player:", player, "\nScore:", player.MinScore())
	fmt.Println("Dealer:", dealer, "\nScore:", dealer.MinScore())
}

func draw(deck []cards.Card) (cards.Card, []cards.Card) {
	return deck[0], deck[1:]
}
