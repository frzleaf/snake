package game

import (
	"container/list"
)

const LEFT_DIRECTION = 1;

const RIGHT_DIRECTION = 2;

const UP_DIRECTION = 3;

const DOWN_DIRECTION = 4;

var (
	LEFT_MOVING = [2]int{-1, 0}
	RIGHT_MOVING = [2]int{1, 0}
	UP_MOVING = [2]int{0, -1}
	DOWN_MOVING = [2]int{0, 1}
)

type Snake struct {
	Body   *list.List
	Length int
	Color  rune
}

func (sn*Snake) Eat(n int) {
	for i := 0; i < n; i++ {
		back := sn.Body.Back().Value.([]int)
		add := []int{back[0], back[1]}
		sn.Body.PushBack(add)
	}
	sn.Length += n
}

func (sn *Snake) Move(direction int, step int) {
	var direction_moving [2]int

	switch direction {
	case LEFT_DIRECTION:
		direction_moving = LEFT_MOVING;
	case RIGHT_DIRECTION:
		direction_moving = RIGHT_MOVING;
	case UP_DIRECTION:
		direction_moving = UP_MOVING;
	case DOWN_DIRECTION:
		direction_moving = DOWN_MOVING;
	default:
		return
	}

	for moved, head := 0, sn.Body.Front(); moved < step; moved ++ {

		point := head.Value.([]int)
		x, y := point[0], point[1]

		newPoint := []int{x + direction_moving[0], y + direction_moving[1]}
		x, y = newPoint[0], newPoint[1]

		sn.Body.InsertBefore(newPoint, head)
		head = sn.Body.Front()
		sn.Body.Remove(sn.Body.Back())

	}
}

func (sn *Snake) MoveLeft(step int) {
	sn.Move(LEFT_DIRECTION, step)
}

func (sn *Snake) MoveRight(step int) {
	sn.Move(RIGHT_DIRECTION, step)
}

func (sn *Snake) MoveUp(step int) {
	sn.Move(UP_DIRECTION, step)
}

func (sn *Snake) MoveDown(step int) {
	sn.Move(DOWN_DIRECTION, step)
}

func (sn *Snake) Init(x, y int) {
	sn.Body = list.New()
	sn.Body.PushFront([]int{x, y})
	sn.Length = 1
}