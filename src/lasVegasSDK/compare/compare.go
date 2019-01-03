package compare

//"strings"
//"fmt"

func (A *Texas) IsGreater(B *Texas) int8 {
	if A.CardsValueKind > B.CardsValueKind {
		return 1
	} else if A.CardsValueKind < B.CardsValueKind {
		return -1
	}
	if A.CardsValueKind == CardValueKindIsHigh {
		return A.IsGreaterHigh(B)
	}
	if A.CardsValueKind == CardValueKindIsPairs {
		return A.IsGreaterPairs(B)
	}
	if A.CardsValueKind == CardValueKindIsTwoPairs {
		return A.IsGreaterTwoPairs(B)
	}
	if A.CardsValueKind == CardValueKindIsTrips {
		return A.IsGreaterTrips(B)
	}
	if A.CardsValueKind == CardValueKindIsStraight {
		return A.IsGreaterStraight(B)
	}
	if A.CardsValueKind == CardValueKindIsFlush {
		return A.IsGreaterFlush(B)
	}
	if A.CardsValueKind == CardValueKindIsFourOfAKind {
		return A.IsGreaterFourOfAKind(B)
	}
	if A.CardsValueKind == CardValueKindStraightFlush {
		return A.IsGreaterStraightFlush(B)
	}
	return -2
}
func (A *Texas) IsGreaterHigh(B *Texas) int8 {
	if RankCardMap[A.Cards[9]] > RankCardMap[B.Cards[9]] {
		return 1
	} else if RankCardMap[A.Cards[9]] < RankCardMap[B.Cards[9]] {
		return -1
	} else if RankCardMap[A.Cards[7]] > RankCardMap[B.Cards[7]] {
		return 1
	} else if RankCardMap[A.Cards[7]] < RankCardMap[B.Cards[7]] {
		return -1
	} else if RankCardMap[A.Cards[5]] > RankCardMap[B.Cards[5]] {
		return 1
	} else if RankCardMap[A.Cards[5]] < RankCardMap[B.Cards[5]] {
		return -1
	} else if RankCardMap[A.Cards[3]] > RankCardMap[B.Cards[3]] {
		return 1
	} else if RankCardMap[A.Cards[3]] < RankCardMap[B.Cards[3]] {
		return -1
	} else if RankCardMap[A.Cards[1]] > RankCardMap[B.Cards[1]] {
		return 1
	} else if RankCardMap[A.Cards[1]] < RankCardMap[B.Cards[1]] {
		return -1
	} else {
		return 0
	}
}
func (A *Texas) IsGreaterPairs(B *Texas) int8 {
	var APairs int8
	var BPairs int8
	var APairsIndex int8
	var BPairsIndex int8
	getPairs := func(myCards *Texas) int8 {
		if myCards.Cards[1] == myCards.Cards[3] {
			return 1
		}
		if myCards.Cards[3] == myCards.Cards[5] {
			return 3
		}
		if myCards.Cards[5] == myCards.Cards[7] {
			return 5
		}
		if myCards.Cards[7] == myCards.Cards[9] {
			return 7
		}
		return -1
	}
	APairsIndex = getPairs(A)
	BPairsIndex = getPairs(B)

	APairs = RankCardMap[A.Cards[APairsIndex]]
	BPairs = RankCardMap[B.Cards[BPairsIndex]]

	if APairs > BPairs {
		return 1
	} else if APairs < BPairs {
		return -1
	}

	myAHighCard := [5]rune{A.Cards[1], A.Cards[3], A.Cards[5], A.Cards[7], A.Cards[9]}
	myBHighCard := [5]rune{B.Cards[1], B.Cards[3], B.Cards[5], B.Cards[7], B.Cards[9]}
	myAHighCard[(APairsIndex-1)/2] = 0
	myAHighCard[(APairsIndex+1)/2] = 0
	myBHighCard[(BPairsIndex-1)/2] = 0
	myBHighCard[(BPairsIndex+1)/2] = 0
	myAHighRank := getHighRank(myAHighCard)
	myBHighRank := getHighRank(myBHighCard)
	if myAHighRank > myBHighRank {
		return 1
	}
	if myAHighRank < myBHighRank {
		return -1
	}
	return 0
}

func getHighRank(myHighCard [5]rune) int {
	var myRank int = 0
	for i := 4; i >= 0; i-- {
		if myHighCard[i] == 0 {
			continue
		}
		myRank = myRank << 4
		myRank = myRank + int(RankCardMap[myHighCard[i]])
	}
	return myRank
}

func (A *Texas) IsGreaterTwoPairs(B *Texas) int8 {
	var APairs1 int8
	var BPairs1 int8
	var APairsIndex1 int8
	var BPairsIndex1 int8
	var APairs2 int8
	var BPairs2 int8
	var APairsIndex2 int8
	var BPairsIndex2 int8
	getTwoPairs := func(myCards *Texas) (int8, int8) {
		if myCards.Cards[1] != myCards.Cards[3] {
			return 3, 7
		}
		if myCards.Cards[7] != myCards.Cards[9] {
			return 1, 5
		}
		return 1, 7
	}
	APairsIndex1, APairsIndex2 = getTwoPairs(A)
	BPairsIndex1, BPairsIndex2 = getTwoPairs(B)

	APairs1 = RankCardMap[A.Cards[APairsIndex1]]
	BPairs1 = RankCardMap[B.Cards[BPairsIndex1]]
	APairs2 = RankCardMap[A.Cards[APairsIndex2]]
	BPairs2 = RankCardMap[B.Cards[BPairsIndex2]]
	if APairs2 > BPairs2 {
		return 1
	} else if APairs2 < BPairs2 {
		return -1
	} else if APairs1 > BPairs1 {
		return 1
	} else if APairs2 < BPairs2 {
		return -1
	}
	myAHighCard := [5]rune{A.Cards[1], A.Cards[3], A.Cards[5], A.Cards[7], A.Cards[9]}
	myBHighCard := [5]rune{B.Cards[1], B.Cards[3], B.Cards[5], B.Cards[7], B.Cards[9]}
	myAHighCard[(APairsIndex1-1)/2] = 0
	myAHighCard[(APairsIndex1+1)/2] = 0
	myAHighCard[(APairsIndex2-1)/2] = 0
	myAHighCard[(APairsIndex2+1)/2] = 0

	myBHighCard[(BPairsIndex1-1)/2] = 0
	myBHighCard[(BPairsIndex1+1)/2] = 0
	myBHighCard[(BPairsIndex2-1)/2] = 0
	myBHighCard[(BPairsIndex2+1)/2] = 0
	myAHighRank := getHighRank(myAHighCard)
	myBHighRank := getHighRank(myBHighCard)
	if myAHighRank > myBHighRank {
		return 1
	}
	if myAHighRank < myBHighRank {
		return -1
	}
	return 0
}

func (A *Texas) IsGreaterTrips(B *Texas) int8 {
	var ATrips int8
	var BTrips int8
	var ATripsIndex int8
	var BTripsIndex int8
	getTrips := func(myCards *Texas) int8 {
		if myCards.Cards[1] == myCards.Cards[5] {
			return 1
		}
		if myCards.Cards[3] == myCards.Cards[7] {
			return 3
		}
		if myCards.Cards[5] == myCards.Cards[9] {
			return 5
		}
		return -1
	}
	ATripsIndex = getTrips(A)
	BTripsIndex = getTrips(B)
	ATrips = RankCardMap[A.Cards[ATripsIndex]]
	BTrips = RankCardMap[B.Cards[BTripsIndex]]
	if ATrips > BTrips {
		return 1
	} else if ATrips < BTrips {
		return -1
	}
	myAHighCard := [5]rune{A.Cards[1], A.Cards[3], A.Cards[5], A.Cards[7], A.Cards[9]}
	myBHighCard := [5]rune{B.Cards[1], B.Cards[3], B.Cards[5], B.Cards[7], B.Cards[9]}
	myAHighCard[(ATripsIndex-1)/2] = 0
	myAHighCard[(ATripsIndex+1)/2] = 0
	myAHighCard[(ATripsIndex+3)/2] = 0
	myBHighCard[(BTripsIndex-1)/2] = 0
	myBHighCard[(BTripsIndex+1)/2] = 0
	myBHighCard[(BTripsIndex+3)/2] = 0
	myAHighRank := getHighRank(myAHighCard)
	myBHighRank := getHighRank(myBHighCard)
	if myAHighRank > myBHighRank {
		return 1
	}
	if myAHighRank < myBHighRank {
		return -1
	}
	return 0
}
func (A *Texas) IsGreaterStraight(B *Texas) int8 {
	if RankCardMap[A.Cards[5]] > RankCardMap[A.Cards[5]] {
		return 1
	}
	if RankCardMap[A.Cards[5]] < RankCardMap[A.Cards[5]] {
		return -1
	}
	return 0
}
func (A *Texas) IsGreaterFlush(B *Texas) int8 {
	return A.IsGreaterHigh(B)
}

func (A *Texas) IsGreaterFullHouse(B *Texas) int8 {
	var ATrips int8
	var BTrips int8
	var APair int8
	var BPair int8
	var ATripsIndex int8
	var BTripsIndex int8
	var APairIndex int8
	var BPairIndex int8
	getTripsAndPair := func(myCards *Texas) (int8, int8) {
		if myCards.Cards[1] == myCards.Cards[5] {
			return 1, 7
		}
		if myCards.Cards[5] == myCards.Cards[9] {
			return 5, 1
		}
		return -1, -1
	}
	ATripsIndex, APairIndex = getTripsAndPair(A)
	BTripsIndex, BPairIndex = getTripsAndPair(B)
	ATrips = RankCardMap[A.Cards[ATripsIndex]]
	BTrips = RankCardMap[B.Cards[BTripsIndex]]
	APair = RankCardMap[A.Cards[APairIndex]]
	BPair = RankCardMap[B.Cards[BPairIndex]]
	if ATrips > BTrips {
		return 1
	} else if ATrips < BTrips {
		return -1
	} else if APair > BPair {
		return 1
	} else if APair < BPair {
		return -1
	} else {
		return 0
	}
}

func (A *Texas) IsGreaterFourOfAKind(B *Texas) int8 {
	var AFour int8
	var BFour int8
	var AHigh int8
	var BHigh int8
	var AFourIndex int8
	var BFourIndex int8
	var AHighIndex int8
	var BHighIndex int8
	getFourAndHigh := func(myCards *Texas) (int8, int8) {
		if myCards.Cards[1] == myCards.Cards[7] {
			return 1, 9
		}
		if myCards.Cards[3] == myCards.Cards[9] {
			return 3, 1
		}
		return -1, -1
	}
	AFourIndex, AHighIndex = getFourAndHigh(A)
	BFourIndex, BHighIndex = getFourAndHigh(B)
	AFour = RankCardMap[A.Cards[AFourIndex]]
	BFour = RankCardMap[B.Cards[BFourIndex]]
	AHigh = RankCardMap[A.Cards[AHighIndex]]
	BHigh = RankCardMap[B.Cards[BHighIndex]]
	if AFour > BFour {
		return 1
	} else if AFour < BFour {
		return -1
	} else if AHigh > BHigh {
		return 1
	} else if AHigh < BHigh {
		return -1
	} else {
		return 0
	}
}

func (A *Texas) IsGreaterStraightFlush(B *Texas) int8 {
	return A.IsGreaterStraight(B)
}
