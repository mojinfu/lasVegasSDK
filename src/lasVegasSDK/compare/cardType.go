package compare

import (
	//"strings"
	//"fmt"
	"math/rand"
	"time"
)

var TypeCard [4]rune = [4]rune{'S', 'C', 'D', 'H'}
var RankCard [13]rune = [13]rune{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}
var myRandSeed *rand.Rand
var RankCardMap map[rune]int8 = map[rune]int8{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}
var CardValueKindMap map[int]string = map[int]string{
	CardValueKindStraightFlush: "StraightFlush",
	CardValueKindIsFourOfAKind: "FourOfAKind",
	CardValueKindIsFullHouse:   "FullHouse",
	CardValueKindIsFlush:       "Flush",
	CardValueKindIsStraight:    "Straight",
	CardValueKindIsTrips:       "Trips",
	CardValueKindIsTwoPairs:    "TwoPairs",
	CardValueKindIsPairs:       "Pairs",
	CardValueKindIsHigh:        "High",
}

const (
	CardValueKindStraightFlush = 1 << (8 + 4*5)
	CardValueKindIsFourOfAKind = 1 << (7 + 4*5)
	CardValueKindIsFullHouse   = 1 << (6 + 4*5)
	CardValueKindIsFlush       = 1 << (5 + 4*5)
	CardValueKindIsStraight    = 1 << (4 + 4*5)
	CardValueKindIsTrips       = 1 << (3 + 4*5)
	CardValueKindIsTwoPairs    = 1 << (2 + 4*5)
	CardValueKindIsPairs       = 1 << (1 + 4*5)
	CardValueKindIsHigh        = 1 << (0 + 4*5)
)

type Texas struct {
	Cards          [10]rune
	CardsValueKind int
}

func init() {
	myRandSeed = rand.New(rand.NewSource(time.Now().UnixNano()))
}
func RandomCards() *Texas {
	var myRune [10]rune
	var typeRuneA rune
	var RankRuneB rune
	isExist := func() bool {
		for ii := 0; myRune[ii] != 0; ii = ii + 2 {
			if typeRuneA == myRune[ii] && RankRuneB == myRune[ii+1] {
				return true
			}
		}
		return false
	}
	for i := 0; i < 9; i = i + 2 {
		typeRuneA, RankRuneB = RandomOneCards()
		for {
			if isExist() {
				typeRuneA, RankRuneB = RandomOneCards()
			} else {
				myRune[i] = typeRuneA
				myRune[i+1] = RankRuneB
				break
			}
		}
	}
	return NewTexas(myRune)
}
func RandomOneCards() (rune, rune) {
	return TypeCard[myRandSeed.Intn(3)], RankCard[myRandSeed.Intn(12)]
}
func NewTexas(myCards [10]rune) *Texas {
	myTexas := &Texas{
		Cards: myCards,
	}
	myTexas.Sort()
	// for i:=0;i<10;i++{
	// 	myTexas.CardsStr = myTexas.CardsStr + string(myTexas.Cards[i])
	// }
	myTexas.GetCardsValueKind()
	return myTexas
}
func (this *Texas) Sort() {
	var vThis int8
	var vNext int8
	var tempRuneA rune
	var tempRuneB rune
	for ii := 4; ii > 0; ii-- {
		for index := 0; index < ii; index = index + 1 {
			vThis = RankCardMap[this.Cards[index*2+1]]
			vNext = RankCardMap[this.Cards[(index+1)*2+1]]
			if vThis > vNext {
				tempRuneA = this.Cards[(index+1)*2]
				tempRuneB = this.Cards[(index+1)*2+1]
				this.Cards[(index+1)*2] = this.Cards[index*2]
				this.Cards[(index+1)*2+1] = this.Cards[index*2+1]
				this.Cards[index*2] = tempRuneA
				this.Cards[index*2+1] = tempRuneB
			}
		}
	}
}

//同花顺
func (this *Texas) IsStraightFlush() bool {
	return this.IsFlush() && this.IsStraight()
}

//四个
func (this *Texas) IsFourOfAKind() bool {
	return this.Cards[3] == this.Cards[5] &&
		this.Cards[5] == this.Cards[7] &&
		(this.Cards[7] == this.Cards[9] ||
			this.Cards[3] == this.Cards[1])
}

//葫芦
func (this *Texas) IsFullHouse() bool {
	return (this.Cards[1] == this.Cards[3] &&
		this.Cards[7] == this.Cards[9] && (this.Cards[3] == this.Cards[5] || this.Cards[5] == this.Cards[7]))
}

//同花
func (this *Texas) IsFlush() bool {
	return this.Cards[0] == this.Cards[2] &&
		this.Cards[2] == this.Cards[4] &&
		this.Cards[4] == this.Cards[6] &&
		this.Cards[6] == this.Cards[8]
}

//顺子
func (this *Texas) IsStraight() bool {
	if RankCardMap[this.Cards[1]] == RankCardMap[this.Cards[3]]-1 &&
		RankCardMap[this.Cards[3]] == RankCardMap[this.Cards[5]]-1 &&
		RankCardMap[this.Cards[5]] == RankCardMap[this.Cards[7]]-1 &&
		(RankCardMap[this.Cards[7]] == RankCardMap[this.Cards[9]]-1 || (this.Cards[7] == '5' && this.Cards[9] == 'A')) {
		return true
	}
	return false
}

//三个
func (this *Texas) IsTrips() bool {
	return (this.Cards[1] == this.Cards[5]) ||
		(this.Cards[3] == this.Cards[7]) ||
		(this.Cards[5] == this.Cards[9])
}

//俩对
func (this *Texas) IsTwoPairs() bool {
	return ((this.Cards[1] == this.Cards[3]) && (this.Cards[5] == this.Cards[7])) ||
		((this.Cards[1] == this.Cards[3]) && (this.Cards[7] == this.Cards[9])) ||
		((this.Cards[3] == this.Cards[5]) && (this.Cards[7] == this.Cards[9]))
}

//一对
func (this *Texas) IsPairs() bool {
	return (this.Cards[1] == this.Cards[3]) ||
		(this.Cards[3] == this.Cards[5]) ||
		(this.Cards[5] == this.Cards[7]) ||
		(this.Cards[7] == this.Cards[9])
}
func (this *Texas) GetCardsValueKind() {
	if this.IsStraightFlush() {
		this.CardsValueKind = CardValueKindStraightFlush
		return
	}
	if this.IsFourOfAKind() {
		this.CardsValueKind = CardValueKindIsFourOfAKind
		return
	}
	if this.IsFullHouse() {
		this.CardsValueKind = CardValueKindIsFullHouse
		return
	}
	if this.IsFlush() {
		this.CardsValueKind = CardValueKindIsFlush
		return
	}
	if this.IsStraight() {
		this.CardsValueKind = CardValueKindIsStraight
		return
	}
	if this.IsTrips() {
		this.CardsValueKind = CardValueKindIsTrips
		return
	}
	if this.IsTwoPairs() {
		this.CardsValueKind = CardValueKindIsTwoPairs
		return
	}
	if this.IsPairs() {
		this.CardsValueKind = CardValueKindIsPairs
		return
	}
	this.CardsValueKind = CardValueKindIsHigh
	return
}
