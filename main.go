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

func (h Hand) Score() int {
	minScore := h.MinScore()
	if minScore > 11 {
		return minScore
	}
	for _, c := range h {
		if c.Rank == cards.Ace {
			return minScore + 10
		}
	}
	return minScore
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
	for dealer.Score() <= 16 || (dealer.Score() == 17 && dealer.MinScore() != 17) {
		card, deck = draw(deck)
		dealer = append(dealer, card)
	}
	pScore, dScore := player.Score(), dealer.Score()
	fmt.Println("==FINAL HANDS==")
	fmt.Println("Player:", player, "\nScore:", pScore)
	fmt.Println("Dealer:", dealer, "\nScore:", dScore)
	switch {
	case pScore > 21:
		fmt.Println("You busted")
	case dScore > 21:
		fmt.Println("Dealer busted")
	case pScore > dScore:
		fmt.Println("You win!")
	case dScore > pScore:
		fmt.Println("You lose!")
	case dScore == pScore:
		fmt.Println("Draw")
	}
}

func draw(deck []cards.Card) (cards.Card, []cards.Card) {
	return deck[0], deck[1:]
}
