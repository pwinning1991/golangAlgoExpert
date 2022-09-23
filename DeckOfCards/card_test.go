package deck

import (
	"fmt"
	"math/rand"
	"testing"
)

func ExampleCard() {
	fmt.Println((Card{Rank: Ace, Suit: Heart}))
	fmt.Println((Card{Rank: Two, Suit: Spade}))
	fmt.Println((Card{Rank: Nine, Suit: Diamond}))
	fmt.Println((Card{Rank: Jack, Suit: Club}))
	fmt.Println((Card{Suit: Joker}))

	// Output:
	// Ace of Hearts
	// Two of Spades
	// Nine of Diamonds
	// Jack of Clubs
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()

	// 13 ranks * 4 suits
	if len(cards) != 13*4 {
		t.Error("Wrong number of cards in a new deck.")

	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Error("Expected Ace of Spades, got", cards[0])
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Error("Expected Ace of Spades, got", cards[0])
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(3))
	//Number of cards is suits * ranks + the number of Jokers added to the deck
	if len(cards) != 13*4+3 {
		t.Error("Wrong number of cards in the deck")

	}

	// Count the number of Jokers in our new deck and make sure it is the correct number
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}

	if count != 3 {
		t.Error("Did not get the right number of jokers")
	}

}

func TestShuffle(t *testing.T) {
	// Make shufflerand deterministic
	// First call to shuffleRand.Perm(52) should be:
	// [40 35 50 0 ...]
	shuffleRand = rand.New(rand.NewSource(0))
	orig := New()
	cards := New(Shuffle)
	first := orig[40]
	second := orig[35]
	if cards[0] != first {
		t.Errorf("Expected the first card to be %s, received %s", first, cards[0])
	}
	if cards[1] != second {
		t.Errorf("Expected the first card to be %s, received %s", second, cards[1])
	}

}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	cards := New(Filter(filter))
	for _, card := range cards {
		if card.Rank == Two || card.Rank == Three {
			t.Error("Was supposed to filter two and threes")
		}
	}

}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))
	// 13 ranks * 4 suits * 3 decks
	if len(cards) != 13*4*3 {
		t.Errorf("Expected %d cards, got %d cards", 13*3*4, len(cards))
	}

}
