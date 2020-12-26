package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
)

const (
	//黑桃
	Spade = 0
	//红桃
	Hearts = 1
	//梅花
	Club = 2
	//方块
	Diamond = 3
)

type Pokers []Poker

type Poker struct {
	Num    int
	Flower int
}

func (p Poker) PokerSelf() string {
	var buffer string

	switch p.Flower {
	case Spade:
		buffer += "♤"
	case Hearts:
		buffer += "♡"
	case Club:
		buffer += "♧"
	case Diamond:
		buffer += "♢"
	}
	switch p.Num {
	case 13:
		buffer += "2"
	case 12:
		buffer += "A"
	case 11:
		buffer += "K"
	case 10:
		buffer += "Q"
	case 9:
		buffer += "J"
	default:
		buffer += strconv.Itoa(p.Num + 2)
	}

	return buffer
}

func CreatePokers() (pokers Pokers) {
	for i := 1; i < 14; i++ {
		for j := 0; j < 4; j++ {
			pokers = append(pokers, Poker{
				Num:    i,
				Flower: j,
			})
		}
	}
	return
}

func (p Pokers) Print() {
	for _, i2 := range p {
		fmt.Print(i2.PrintPoker(), " ")
	}
	fmt.Println()
}

func (p Poker) PrintPoker() string {
	return p.PokerSelf()
}
func main() {
	var pokers Pokers
	pokers = CreatePokers()
	fmt.Println("洗牌前：")
	pokers.Print()
	sort.Slice(pokers, func(i, j int) bool {
		if rand.Intn(100) < 50 {
			return false
		} else {
			return true
		}
	})
	fmt.Println("洗牌后：")
	pokers.Print()
	sort.Slice(pokers, func(i, j int) bool {
		if pokers[i].Num < pokers[j].Num {
			return true
		} else {
			return false
		}
	})
	fmt.Println("排序后：")
	pokers.Print()
}
