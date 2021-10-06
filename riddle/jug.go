package riddle

import "fmt"

type Jug struct {
	size   int64
	amount int64
}

func (o *Jug) IsEmpty() bool {
	return o.amount == 0
}

func (o *Jug) IsFull() bool {
	return o.amount == o.size
}

func (o *Jug) Remaining() int64 {
	return o.size - o.amount
}

func (o *Jug) Fill() {
	o.amount = o.size
}

func (o *Jug) Empty() {
	if o.amount == 0 {
		panic("Jug is already empty")
	}
	o.amount = 0
}

// Pour pours the volume passed as parameter into the Jug.
// If the volume exceeds the remaining capacity of the Jug
// then the remainder is returned.
func (o *Jug) Pour(volume int64) int64 {
	if volume <= 0 {
		panic(fmt.Sprintf("Invalid volume: %d", volume))
	}
	if volume > o.Remaining() {
		remainder := volume - o.Remaining()
		o.Fill()
		return remainder
	}
	o.amount += volume
	return 0
}

func (o *Jug) Set(volume int64) {
	if volume < 0 || volume > o.size {
		panic(fmt.Sprintf("Invalid volume: %d", volume))
	}
	o.amount = volume
}

func (o *Jug) Clone() *Jug {
	return NewJug(o.size, o.amount)
}

func (o *Jug) Equals(other Jug) bool{
	return o.amount == other.amount && o.size == other.size
}

func (o *Jug) Amount() int64 {
	return o.amount
}

func NewJug(size, amount int64) *Jug {
	return &Jug{
		size:   size,
		amount: amount,
	}
}
