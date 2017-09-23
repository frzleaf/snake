package game

import (
	"github.com/nsf/termbox-go"
	"time"
)

type CliGame struct {
	Board           *Board
	Snakes          []Snake
	IsStarted       bool
	Level           int
	Directions      []chan int
	MovingDirection []int
}

func (game *CliGame) DrawTable() {
	for i := 0; i < game.Board.Width; i ++ {
		for j := 0; j < game.Board.Height; j ++ {
			termbox.SetCell(i, j, ' ', termbox.ColorWhite, termbox.ColorBlack)
		}
	}
}

func (game *CliGame) DrawFoods() {
	for _, f := range game.Board.Foods {
		if f != nil {
			termbox.SetCell((*f)[0], (*f)[1], '*', termbox.ColorWhite, termbox.ColorBlack)
		}
	}
}

func (game *CliGame) UpdateDrawing() {
	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)
	game.DrawTable()
	game.DrawFoods()

	for _, snake := range game.Snakes {
		for point := snake.Body.Front(); point != nil; point = point.Next() {
			val := point.Value.([]int)
			termbox.SetCell(val[0], val[1], ' ', termbox.ColorBlack, termbox.ColorWhite)
		}
	}
	termbox.Flush()
}

func (game *CliGame) backgroundProcess() {
	for ; game.IsStarted; {
		for i, moving_direction := range game.MovingDirection {
			snake := game.Snakes[i];
			snake.Move(moving_direction, 1)

			head := snake.Body.Front().Value.([]int)
			if nFood := game.Board.ReachFood(head[0], head[1]); nFood > 0 {
				snake.Eat(nFood)
				game.Board.DropFood(nFood)
			}
		}
		game.UpdateDrawing()
		time.Sleep(time.Duration((50 / game.Level)) * time.Millisecond)
	}
}

func (game *CliGame) autoRun() {
	go game.backgroundProcess()

	for {
		for i, direction := range game.Directions {
			select {
			case d := <-direction:
				game.MovingDirection[i] = d
			}
		}
	}
}

func (game *CliGame) Start() {
	termbox.Init()
	defer termbox.Close()

	game.DrawTable()
	game.IsStarted = true
	moving_direction := 0

	go game.autoRun()

	loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {

		case termbox.EventKey:

			switch ev.Key {
			case termbox.KeyEsc:
				break loop;

			case termbox.KeyArrowLeft:
				game.Snakes[0].MoveLeft(1)
				moving_direction = LEFT_DIRECTION
			case termbox.KeyArrowRight:
				game.Snakes[0].MoveRight(1)
				moving_direction = RIGHT_DIRECTION
			case termbox.KeyArrowUp:
				game.Snakes[0].MoveUp(1)
				moving_direction = UP_DIRECTION
			case termbox.KeyArrowDown:
				game.Snakes[0].MoveDown(1)
				moving_direction = DOWN_DIRECTION
			}

		default:
		}

		if moving_direction != 0 {
			game.Directions[0] <- moving_direction
		}
		moving_direction = 0
	}
}

func (game *CliGame) InitDefault() {
	game.Board = &Board{Height:24, Width:80}
	game.Board.InitDefault()

	game.Level = 1
	game.Directions = []chan int{make(chan int, 1)}
	game.MovingDirection = make([]int, 1)
	player1 := Snake{}
	player1.Init(12, 40)
	game.Snakes = []Snake{
		player1,
	}
}