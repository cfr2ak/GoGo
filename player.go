package gogo

import "errors"

const (
	BLACK = iota
	WHITE = iota
)

type Player struct {
	color int
}

func NewPlayer(color int) (*Player, error) {
	if color == BLACK || color == WHITE {
		return &Player{color}, nil
	}
	return nil, errors.New("color should be one of WHITE or BLACK")
}

func (p Player) other() *Player {
	if p.color == BLACK {
		return &Player{WHITE}
	} else {
		return &Player{BLACK}
	}
}
