package tato

import (
	"errors"
	"regexp"
	"time"
)

var (
	ErrInvalidAction = errors.New("invalid action input")
	actionRegex      = regexp.MustCompile(`(\d{1,2})([SRPABD])([JFXABC])?([=\-\.+*])(\d{1,2})?`)
)

// Action in the game
type Action struct {
	Timestamp time.Time
	Player    string

	Kind         string
	Class        string
	Evaluation   string
	StartingZone string
	EndingZone   string
}

// AddZone depending on the number of digits specified.
// If there is only one digit we should consider the
// kind of the action, otherwise is start-end.
func (a *Action) AddZone(zones string) {
	if len(zones) > 1 {
		a.StartingZone = string(zones[0])
		a.EndingZone = string(zones[1])

		return
	}

	switch a.Kind {
	case "S":
		a.EndingZone = zones
	case "A":
		a.StartingZone = zones
	case "P":
		a.EndingZone = zones
	case "D":
		a.StartingZone = zones
	case "R":
		a.StartingZone = zones
	case "B":
		a.StartingZone = zones
	}

}

func Parse(input string) (*Action, error) {
	if !actionRegex.MatchString(input) {
		return &Action{}, ErrInvalidAction
	}

	parsed := actionRegex.FindAllStringSubmatch(input, -1)
	ac := &Action{
		Timestamp: time.Now(),
		Player:    parsed[0][1],

		Kind:       parsed[0][2],
		Class:      parsed[0][3],
		Evaluation: parsed[0][4],
	}

	ac.AddZone(parsed[0][5])

	return ac, nil
}
