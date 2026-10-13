package battle

type RoundWinner uint8

const (
	RoundNone RoundWinner = iota
	RoundTie
	RoundP1
	RoundP2
)

var beats = map[Move]Move{
	Attack:  Defend,
	Special: Attack,
	Defend:  Special,
}

type MoveDamage struct {
	Attack  int
	Defend  int
	Special int
}

type DamageConfig struct {
	Win MoveDamage
	Tie MoveDamage
}

type RoundResult struct {
	P1Move     Move
	P2Move     Move
	DamageToP1 int
	DamageToP2 int
}

func (r RoundWinner) String() string {
	switch r {
	case RoundNone:
		return "none"
	case RoundTie:
		return "tie"
	case RoundP1:
		return "p1"
	case RoundP2:
		return "p2"
	default:
		return "unknown"
	}
}

func (d MoveDamage) For(m Move) int {
	switch m {
	case Attack:
		return d.Attack
	case Defend:
		return d.Defend
	case Special:
		return d.Special
	default:
		return 0
	}
}

func (r RoundResult) Winner() RoundWinner {
	return compare(r.P1Move, r.P2Move)
}

func compare(p1, p2 Move) RoundWinner {
	if !p1.IsValid() || !p2.IsValid() {
		return RoundNone
	}
	if p1 == p2 {
		return RoundTie
	}
	if beats[p1] == p2 {
		return RoundP1
	}
	return RoundP2
}

func ResolveRound(p1, p2 Move, cfg DamageConfig) (RoundResult, error) {
	winner := compare(p1, p2)

	if winner == RoundNone {
		return RoundResult{}, ErrInvalidMove
	}

	result := RoundResult{P1Move: p1, P2Move: p2}

	switch winner {
	case RoundTie:
		d := cfg.Tie.For(p1)
		result.DamageToP1 = d
		result.DamageToP2 = d
	case RoundP1:
		result.DamageToP2 = cfg.Win.For(p1)
	case RoundP2:
		result.DamageToP1 = cfg.Win.For(p2)
	}
	return result, nil
}
