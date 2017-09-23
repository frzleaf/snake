package game

import (
	"math/rand"
	"time"
)

const MAX_FOOD = 10

type Board struct {
	Width  int
	Height int
	Foods  [MAX_FOOD](*[2]int)
}

func (b *Board) DropFood(n int) {

	for i, l  := 0, len(b.Foods); i < l && n > 0; i++ {
		if food := b.Foods[i]; food == nil {
			rand.Seed(time.Now().UnixNano() + int64(1<<uint(i)))
			b.Foods[i] = &[2]int{rand.Intn(b.Width), rand.Intn(b.Height)}
			n --
		}
	}
}

func (b *Board) ReachFood(x, y int) int {
	var res int = 0
	for i, f := range b.Foods {
		if f != nil {
			if x == (*f)[0] && y == (*f)[1] {
				res ++
				b.Foods[i] = nil
			}
		}
	}
	return res
}

func (b*Board) InitDefault() {
	b.DropFood(10)
}