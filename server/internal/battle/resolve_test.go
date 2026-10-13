package battle

import (
	"errors"
	"testing"
)

/**
TestResolveRound	9 kombinasi (cek DamageToP1, DamageToP2, dan Winner())
TestResolveRound	input tidak valid (MoveNone, Move(99)) → ErrInvalidMove
TestMoveDamageFor	3 move → field yang benar; MoveNone → 0
TestRoundResultWinner	RoundResult{} kosong → RoundNone
*/

func TestResolveRound(t *testing.T) {
	cfg := DamageConfig{
		Win:  MoveDamage{Attack: 11, Defend: 22, Special: 33},
		Tie:  MoveDamage{Attack: 1, Defend: 0, Special: 3},
		Chip: MoveDamage{Attack: 5, Defend: 0, Special: 0},
	}

	tests := []struct {
		name       string
		p1, p2     Move
		wantDmgP1  int
		wantDmgP2  int
		wantWinner RoundWinner
		wantErr    error
	}{
		{"attack beats defend", Attack, Defend, 0, 11, RoundP1, nil},
		{"special beats attack", Special, Attack, 5, 33, RoundP1, nil},
		{"defend beats special", Defend, Special, 0, 22, RoundP1, nil},
		{"attack tie", Attack, Attack, 1, 1, RoundTie, nil},
		{"defend tie", Defend, Defend, 0, 0, RoundTie, nil},
		{"special tie", Special, Special, 3, 3, RoundTie, nil},
		{"defend beated by attack", Defend, Attack, 11, 0, RoundP2, nil},
		{"attack beated by special", Attack, Special, 33, 5, RoundP2, nil},
		{"special beated by defend", Special, Defend, 22, 0, RoundP2, nil},
		{"p1 invalid none", MoveNone, Attack, 0, 0, RoundNone, ErrInvalidMove},
		{"p2 invalid oor", Attack, Move(99), 0, 0, RoundNone, ErrInvalidMove},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ResolveRound(tt.p1, tt.p2, cfg)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("ResolveRound(%v,%v) error = %v, want %v", tt.p1, tt.p2, err, tt.wantErr)
			}
			if tt.wantErr != nil {
				return
			}

			if got.DamageToP1 != tt.wantDmgP1 {
				t.Errorf("ResolveRound(%v, %v) DamageToP1 = %d, want %d",
					tt.p1, tt.p2, got.DamageToP1, tt.wantDmgP1)
			}

			if got.DamageToP2 != tt.wantDmgP2 {
				t.Errorf("ResolveRound(%v, %v) DamageToP2 = %d, want %d", tt.p1, tt.p2, got.DamageToP2, tt.wantDmgP2)
			}

			if got.Winner() != tt.wantWinner {
				t.Errorf("ResolveRound(%v, %v) Winner = %v, want %v", tt.p1, tt.p2, got.Winner(), tt.wantWinner)
			}

		})
	}
}
