package compare

import (
	"fmt"
	"testing"
)

func TestSortFunc(t *testing.T) {
	t.Error("AC")
	fmt.Println(1 << 31)
	myTexas := RandomCards()
	for i := 0; i < 9; i = i + 2 {
		fmt.Printf("%s%s\n", string(myTexas.Cards[i]), string(myTexas.Cards[i+1]))
	}
	fmt.Println(myTexas.IsStraightFlush())
	fmt.Println(myTexas.IsFourOfAKind())
	fmt.Println(myTexas.IsFullHouse())
	fmt.Println(myTexas.IsStraight())
	fmt.Println(myTexas.IsFlush())
	fmt.Println(myTexas.IsTrips())
	fmt.Println(myTexas.IsPairs())
}

func TestIsGreaterPairs(t *testing.T) {
	t.Error("AC")
	fmt.Println(1 << 31)
	myTexasA := NewTexas([10]rune{'D', 'Q', 'H', 'K', 'D', 'K', 'H', '6', 'S', 'T'})
	for i := 0; i < 9; i = i + 2 {
		fmt.Printf("%s%s\n", string(myTexasA.Cards[i]), string(myTexasA.Cards[i+1]))
	}
	fmt.Println("------")
	myTexasB := NewTexas([10]rune{'S', 'A', 'C', '2', 'S', '2', 'C', '4', 'C', '5'})
	for i := 0; i < 9; i = i + 2 {
		fmt.Printf("%s%s\n", string(myTexasB.Cards[i]), string(myTexasB.Cards[i+1]))
	}
	fmt.Println(CardValueKindMap[myTexasA.CardsValueKind])
	fmt.Println(CardValueKindMap[myTexasB.CardsValueKind])
	fmt.Println(myTexasA.IsGreater(myTexasB))
}
