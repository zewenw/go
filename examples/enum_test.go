package examples

import (
	"fmt"
	"testing"
)

type ServerState int

const (
	StateIdel ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdel:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	if v, ok := stateName[ss]; ok {
		return v
	}
	return "unknown state"
}

func transition(s ServerState) ServerState {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover from panic", r)
		}
	}()
	switch s {
	case StateIdel:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdel
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknow servers state: %s", s))

	}
}

func TestEnum(t *testing.T) {
	t.Run("Enum Test", func(t *testing.T) {
		s := transition(StateIdel)
		fmt.Println(s)

		ss := transition(s)
		fmt.Println(ss)

		transition(5)
	})
}
