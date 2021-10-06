package riddle

import (
	"errors"
	"fmt"
)

type Step struct {
	JugX     *Jug
	JugY     *Jug
	Previous *Step
	Action   Action
}

func (o *Step) Execute() {
	switch o.Action {
	case FillX:
		o.JugX.Fill()
	case FillY:
		o.JugY.Fill()
	case EmptyX:
		o.JugX.Empty()
	case EmptyY:
		o.JugY.Empty()
	case TransferX:
		o.transfer(o.JugX, o.JugY)
	case TransferY:
		o.transfer(o.JugY, o.JugX)
	}
}

//Transfer transfers water from one jug to the other.
//If the origin Jug contained more water that the remaining capacity of the destination Jug
// then the remainder is set into the origin Jug
func (o *Step) transfer(from, to *Jug) {
	remainder := to.Pour(from.amount)
	from.Set(remainder)
}

//NextSteps returns an array with the valid next Steps of the current Step
func (o *Step) NextSteps() []Step {
	nextSteps := make([]Step, 0)

	for _, action := range Actions() {
		if err := o.validate(action); err == nil {
			nextSteps = append(nextSteps, o.next(action))
		}
	}
	return nextSteps
}

func (o *Step) validate(action Action) error {
	switch action {
	case FillX:
		if !o.JugX.IsEmpty() {
			return errors.New("jug X is not empty")
		}
	case FillY:
		if !o.JugY.IsEmpty() {
			return errors.New("jug Y is not empty")
		}
	case EmptyX:
		if o.JugX.IsEmpty() {
			return errors.New("jug X is empty")
		}
	case EmptyY:
		if o.JugY.IsEmpty() {
			return errors.New("jug Y is empty")
		}
	case TransferX:
		if o.JugX.IsEmpty() {
			return errors.New("jug X is empty")
		}
		if o.JugY.IsFull() {
			return errors.New("jug Y is full")
		}
	case TransferY:
		if o.JugY.IsEmpty() {
			return errors.New("jug Y is empty")
		}
		if o.JugX.IsFull() {
			return errors.New("jug X is full")
		}
	}
	return nil
}

func (o *Step) next(action Action) Step {
	return Step{
		JugX:     o.JugX.Clone(),
		JugY:     o.JugY.Clone(),
		Previous: o,
		Action:   action,
	}
}

func (o *Step) String() string {
	v := fmt.Sprintf("%s [JugX: %d] [JugY: %d]", o.Action, o.JugX.amount, o.JugY.amount)
	return v
}

func (o *Step) Solution() []Step {
	result := make([]Step, 0)
	return o.solution(result)

}

func (o *Step) solution(steps []Step) []Step {
	if o.Previous == nil{
		return steps
	}
	steps = append(o.Previous.solution(steps), *o)
	return steps
}

func (o *Step) Equals(other *Step) bool {
	return o.JugX.amount == other.JugX.amount &&
		o.JugY.amount == other.JugY.amount
}

func NewStep(sizeX, sizeY int64) *Step {
	return &Step{
		JugX:     NewJug(sizeX, 0),
		JugY:     NewJug(sizeY, 0),
		Previous: nil,
		Action:   Start,
	}
}
