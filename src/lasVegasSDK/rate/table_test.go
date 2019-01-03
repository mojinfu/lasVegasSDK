package rate

import (
	"fmt"
	"testing"
)

func TestSelectGreatestCardsComposer(t *testing.T) {
	t.Error("AC")
	rateFunc := func() int8 {
		var myTable table = table{
			tableCards: [10]rune{},
			myCards:    [4]rune{'H', '7', 'S', '2'}, //输入自己的手牌（红桃7  黑桃2）  既可输出胜负几率
			//{'S', 'C', 'D', 'H'} 对应扑克花色   黑桃 spade  梅花 club  方块 diamond   红桃 heart
		}
		myTable.MakeupTable()
		// for i:=0;i<len(myTable.myCards);i=i+2  {
		// 	fmt.Printf("%s%s\n", string(myTable.myCards[i]),string(myTable.myCards[i+1]))
		// }
		// fmt.Println("---")
		// for i:=0;i<len(myTable.tableCards);i=i+2  {
		// 	fmt.Printf("%s%s\n", string(myTable.tableCards[i]),string(myTable.tableCards[i+1]))
		// }
		// fmt.Println("---")
		// for i:=0;i<len(myTable.hisCards);i=i+2  {
		// 	fmt.Printf("%s%s\n", string(myTable.hisCards[i]),string(myTable.hisCards[i+1]))
		// }
		//	fmt.Println("------")
		AA := myTable.SelectMyGreatestCardsComposer()
		// for i:=0;i<9;i=i+2{
		// 	fmt.Printf("%s%s\n", string(AA.Cards[i]),string(AA.Cards[i+1]))
		// }
		// fmt.Println("------")
		BB := myTable.SelectHisGreatestCardsComposer()
		// for i:=0;i<9;i=i+2{
		// 	fmt.Printf("%s%s\n", string(BB.Cards[i]),string(BB.Cards[i+1]))
		// }
		return AA.IsGreater(BB)
	}
	win := 0
	lose := 0
	pin := 0
	for ii := 0; ii < 100000; ii++ {
		isWIn := rateFunc()
		if isWIn > 0 {
			win++
		} else if isWIn < 0 {
			lose++
		} else {
			pin++
		}
	}
	fmt.Println("胜局次数:", win)
	fmt.Println("败局次数:", lose)
	fmt.Println("平局次数:", pin)
}
