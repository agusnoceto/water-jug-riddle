package riddle

type Action int

const (
	Start Action = iota
	FillX
	FillY
	EmptyX
	EmptyY
	TransferX
	TransferY
)

func (o Action) String() string {
	switch o {
	case Start:
		return "Start:\t\t\t"
	case FillX:
		return "Fill Jug X:\t\t"
	case FillY:
		return "Fill Jug Y:\t\t"
	case EmptyX:
		return "Empty Jug X:\t\t"
	case EmptyY:
		return "Empty Jug Y:\t\t"
	case TransferX:
		return "Transfer Jug X into Y:\t"
	case TransferY:
		return "Transfer Jug Y into X:13\t"
	default:
		return ""
	}
}

func Actions() []Action {
	return []Action{FillX, FillY, EmptyX, EmptyY, TransferX, TransferY}
}
