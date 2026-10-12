package battle

import (
	"errors"
	"testing"
)

func TestParseMove(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Move
		wantErr error
	}{
		{"valid attack", "attack", Attack, nil},
		{"valid defend", "defend", Defend, nil},
		{"valid special", "special", Special, nil},
		{"empty string", "", MoveNone, ErrInvalidMove},
		{"unknown move", "fireball", MoveNone, ErrInvalidMove},
		{"uppercase rejected", "ATTACK", MoveNone, ErrInvalidMove},
		{"none label rejected", "none", MoveNone, ErrInvalidMove},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseMove(tt.input)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("ParseMove(%q) error = %v, want %v", tt.input, err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("ParseMove(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		name  string
		input Move
		want  bool
	}{
		{"valid attack move", Attack, true},
		{"valid defend move", Defend, true},
		{"valid special move", Special, true},
		{"invalid none move", MoveNone, false},
		{"invalid move", Move(99), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isValid := tt.input.IsValid()
			if isValid != tt.want {
				t.Errorf("IsValid(%q) = %v, want %v", tt.input, isValid, tt.want)
			}
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name  string
		input Move
		want  string
	}{
		{"valid attack string", Attack, "attack"},
		{"valid defend string", Defend, "defend"},
		{"valid special string", Special, "special"},
		{"valid none string", MoveNone, "none"},
		{"invalid move", Move(99), "unknown"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.String()
			if got != tt.want {
				t.Errorf("String() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestMoveRoundTrip(t *testing.T) {
	moves := []Move{
		Attack,
		Defend,
		Special,
	}

	for _, m := range moves {
		t.Run(m.String(), func(t *testing.T) {
			s := m.String()

			got, err := ParseMove(s)
			if err != nil {
				t.Fatalf("ParseMove(%q) error = %v", s, err)
			}
			if got != m {
				t.Fatalf("ParseMove(%q) = %v, want %v", s, got, m)
			}
		})
	}
}
