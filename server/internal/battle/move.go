package battle

import "errors"

type Move uint8

const (
	MoveNone Move = iota
	Attack
	Defend
	Special
)

var ErrInvalidMove = errors.New("invalid move")

func (m Move) IsValid() bool {
	return m == Attack || m == Defend || m == Special
}

func (m Move) String() string {
	switch m {
	case MoveNone:
		return "none"
	case Attack:
		return "attack"
	case Defend:
		return "defend"
	case Special:
		return "special"
	default:
		return "unknown"
	}
}

func ParseMove(s string) (Move, error) {
	switch s {
	case "attack":
		return Attack, nil
	case "defend":
		return Defend, nil
	case "special":
		return Special, nil
	}
	return MoveNone, ErrInvalidMove
}
