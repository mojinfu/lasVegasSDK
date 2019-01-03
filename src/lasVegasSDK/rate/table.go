package rate

import (
	. "lasVegas/src/lasVegasSDK/compare"
)

type table struct {
	tableCards [10]rune
	myCards    [4]rune
	hisCards   [4]rune
}

var AllHandsSelect [21][7]int8 = [21][7]int8{
	{1, 1, 1, 1, 1, 0, 0},
	{1, 1, 1, 1, 0, 1, 0},
	{1, 1, 1, 1, 0, 0, 1},
	{1, 1, 1, 0, 1, 0, 1},
	{1, 1, 1, 0, 0, 1, 1},
	{1, 1, 1, 0, 1, 1, 0},
	{1, 1, 0, 0, 1, 1, 1},
	{1, 1, 0, 1, 0, 1, 1},
	{1, 1, 0, 1, 1, 0, 1},
	{1, 1, 0, 1, 1, 1, 0},

	{1, 0, 1, 1, 1, 1, 0},
	{1, 0, 1, 1, 1, 0, 1},
	{1, 0, 1, 1, 0, 1, 1},
	{1, 0, 1, 0, 1, 1, 1},
	{1, 0, 0, 1, 1, 1, 1},

	{0, 1, 1, 1, 1, 1, 0},
	{0, 1, 1, 1, 1, 0, 1},
	{0, 1, 1, 1, 0, 1, 1},
	{0, 1, 1, 0, 1, 1, 1},
	{0, 1, 0, 1, 1, 1, 1},
}

func (this *table) SelectMyGreatestCardsComposer() (greatestHand *Texas) {
	baseHand := NewTexas(this.tableCards)
	var tempHand *Texas
	greatestHand = baseHand
	var VSelect [7]int8
	var mySelectHandCards [10]rune
	var ii int
	for _, VSelect = range AllHandsSelect {
		ii = 0
		for index, V := range VSelect {
			if V == 1 {
				if index < 2 {
					mySelectHandCards[ii] = this.myCards[index*2]
					mySelectHandCards[ii+1] = this.myCards[index*2+1]

				} else {
					mySelectHandCards[ii] = this.tableCards[(index-2)*2]
					mySelectHandCards[ii+1] = this.tableCards[(index-2)*2+1]
				}
				ii = ii + 2
			}
		}
		tempHand = NewTexas(mySelectHandCards)
		if tempHand.IsGreater(greatestHand) > 0 {
			greatestHand = tempHand
		}
	}
	return
}
func (this *table) SelectHisGreatestCardsComposer() (greatestHand *Texas) {
	baseHand := NewTexas(this.tableCards)
	var tempHand *Texas
	greatestHand = baseHand
	var VSelect [7]int8
	var mySelectHandCards [10]rune
	var ii int
	for _, VSelect = range AllHandsSelect {
		ii = 0
		for index, V := range VSelect {
			if V == 1 {
				if index < 2 {
					mySelectHandCards[ii] = this.hisCards[index*2]
					mySelectHandCards[ii+1] = this.hisCards[index*2+1]
				} else {
					mySelectHandCards[ii] = this.tableCards[(index-2)*2]
					mySelectHandCards[ii+1] = this.tableCards[(index-2)*2+1]
				}
				ii = ii + 2
			}
		}
		tempHand = NewTexas(mySelectHandCards)
		if tempHand.IsGreater(greatestHand) > 0 {
			greatestHand = tempHand
		}
	}
	return
}
func (this *table) MakeupTable() {
	for i := 0; i < len(this.myCards); i = i + 2 {
		if this.myCards[i] == 0 {
			myTypeRune, myRankRune := this.RandomOneCardsNew()
			this.myCards[i] = myTypeRune
			this.myCards[i+1] = myRankRune
		}
	}
	for i := 0; i < len(this.tableCards); i = i + 2 {
		if this.tableCards[i] == 0 {
			myTypeRune, myRankRune := this.RandomOneCardsNew()
			this.tableCards[i] = myTypeRune
			this.tableCards[i+1] = myRankRune
		}
	}
	for i := 0; i < len(this.hisCards); i = i + 2 {
		if this.hisCards[i] == 0 {
			myTypeRune, myRankRune := this.RandomOneCardsNew()
			this.hisCards[i] = myTypeRune
			this.hisCards[i+1] = myRankRune
		}
	}
}
func (this *table) RandomOneCardsNew() (myTypeRune rune, myRankRune rune) {
	for {
		myTypeRune, myRankRune = RandomOneCards()
		if !this.IfExistOnTable(myTypeRune, myRankRune) {
			return
		}
	}
}
func (this *table) IfExistOnTable(typeRune rune, RankRune rune) bool {
	for i := 0; i < len(this.myCards); i = i + 2 {
		if this.myCards[i] == typeRune && this.myCards[i+1] == RankRune {
			return true
		}
	}
	for i := 0; i < len(this.tableCards); i = i + 2 {
		if this.tableCards[i] == typeRune && this.tableCards[i+1] == RankRune {
			return true
		}
	}
	for i := 0; i < len(this.hisCards); i = i + 2 {
		if this.hisCards[i] == typeRune && this.hisCards[i+1] == RankRune {
			return true
		}
	}
	return false
}
