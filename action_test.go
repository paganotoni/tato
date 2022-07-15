package tato_test

import (
	"testing"

	"github.com/paganotoni/tato"
)

func TestParse(t *testing.T) {
	t.Run("Invalid should return an error", func(t *testing.T) {
		for _, v := range []string{
			"",
			"12",
			"AS+",     // Invalid player number
			"1+",      // Action missing
			"1S)",     // Invalid evaluation
			"12X+",    // Invalid kind
			"12DS+00", // Invalid zone
		} {
			_, err := tato.Parse(v)
			if err == nil {
				t.Fatalf("Expected error when parsing %v, got nil", v)
			}
		}
	})

	t.Run("valid", func(t *testing.T) {
		ac, err := tato.Parse("12SF+")
		if err != nil {
			t.Fatalf("Expected nil error when parsing valid input, got %v", err)
		}

		if ac.Kind != "S" {
			t.Fatalf("Expected kind S, got %v", ac.Kind)
		}

		if ac.Player != "12" {
			t.Fatalf("Expected player 12, got %v", ac.Player)
		}

		if ac.Class != "F" {
			t.Fatalf("Expected class F, got %v", ac.Class)
		}

		if ac.Evaluation != "+" {
			t.Fatalf("Expected evaluation +, got %v", ac.Evaluation)
		}
	})
}

func TestAddZone(t *testing.T) {
	t.Run("One zone", func(t *testing.T) {

		t.Run("service", func(t *testing.T) {
			a := &tato.Action{Kind: "S"}
			a.AddZone("1")

			if a.EndingZone != "1" {
				t.Errorf("Expected ending zone to be 1, got %v", a.EndingZone)
			}

			if a.StartingZone != "" {
				t.Errorf("Expected starting zone to be empty, got %v", a.StartingZone)
			}
		})

		t.Run("attack", func(t *testing.T) {
			a := &tato.Action{Kind: "A"}
			a.AddZone("1")

			if a.EndingZone != "" {
				t.Errorf("Expected ending zone to be empty, got %v", a.EndingZone)
			}

			if a.StartingZone != "1" {
				t.Errorf("Expected starting zone to be empty, got %v", a.StartingZone)
			}
		})

		t.Run("block", func(t *testing.T) {
			a := &tato.Action{Kind: "B"}
			a.AddZone("1")

			if a.EndingZone != "" {
				t.Errorf("Expected ending zone to be empty, got %v", a.EndingZone)
			}

			if a.StartingZone != "1" {
				t.Errorf("Expected starting zone to be empty, got %v", a.StartingZone)
			}
		})

		t.Run("reception", func(t *testing.T) {
			a := &tato.Action{Kind: "R"}
			a.AddZone("1")

			if a.EndingZone != "" {
				t.Errorf("Expected ending zone to be empty, got %v", a.EndingZone)
			}

			if a.StartingZone != "1" {
				t.Errorf("Expected starting zone to be 1, got %v", a.StartingZone)
			}
		})

		t.Run("pass", func(t *testing.T) {
			a := &tato.Action{Kind: "P"}
			a.AddZone("1")

			if a.EndingZone != "1" {
				t.Errorf("Expected ending zone to be 1, got %v", a.EndingZone)
			}

			if a.StartingZone != "" {
				t.Errorf("Expected starting zone to be empty, got %v", a.StartingZone)
			}
		})

		t.Run("defense", func(t *testing.T) {
			a := &tato.Action{Kind: "D"}
			a.AddZone("1")

			if a.EndingZone != "" {
				t.Errorf("Expected ending zone to be empty, got %v", a.EndingZone)
			}

			if a.StartingZone != "1" {
				t.Errorf("Expected starting zone to be 1, got %v", a.StartingZone)
			}
		})
	})

	t.Run("Two zones", func(t *testing.T) {
		a := &tato.Action{}
		a.AddZone("12")

		if a.StartingZone != "1" {
			t.Fatalf("Expected starting zone 1, got %v", a.StartingZone)
		}

		if a.EndingZone != "2" {
			t.Fatalf("Expected starting zone 2, got %v", a.EndingZone)
		}
	})
}
