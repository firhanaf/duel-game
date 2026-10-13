package battle

import "testing"

/**
TestResolveRound	9 kombinasi (cek DamageToP1, DamageToP2, dan Winner())
TestResolveRound	input tidak valid (MoveNone, Move(99)) → ErrInvalidMove
TestMoveDamageFor	3 move → field yang benar; MoveNone → 0
TestRoundResultWinner	RoundResult{} kosong → RoundNone
*/

func TestResolveRound(t *testing.T) {
	cfg := DamageConfig{}
}
