package battle

type RoundWinner uint8

const (
	RoundNone RoundWinner = iota
	RoundTie
	RoundP1
	RoundP2
)


func Compare(p1, p2 Move) RoundWinner {
	if !p1.IsValid() || !p2.IsValid() {
		return RoundNone
	} else if p1 == p2 {
		return RoundTie
	} else if

}
