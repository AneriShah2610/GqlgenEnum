// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package enum

import (
	"fmt"
	"io"
	"strconv"
)

type Node interface {
	IsNode()
}

type ToDo struct {
	ID       string `json:"id"`
	Text     string `json:"text"`
	State    State  `json:"state"`
	Verified bool   `json:"verified"`
}

func (ToDo) IsNode() {}

type TodoInput struct {
	Text string `json:"text"`
}

type State string

const (
	StateNotYet State = "NOT_YET"
	StateDone   State = "DONE"
)

func (e State) IsValid() bool {
	switch e {
	case StateNotYet, StateDone:
		return true
	}
	return false
}

func (e State) String() string {
	return string(e)
}

func (e *State) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = State(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid State", str)
	}
	return nil
}

func (e State) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
