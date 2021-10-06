package riddle

type Queue struct {
	queue []Step
}

func (o *Queue) Push(value Step) {
	o.queue = append(o.queue, value)
}

func (o *Queue) PushAll(values []Step) {
	o.queue = append(o.queue, values...)
}

func (o *Queue) Pop() *Step {
	if len(o.queue) > 0 {
		first := o.queue[0]
		o.queue = o.queue[1:]
		return &first
	}
	panic("pop Error: Queue is empty")
}

func (o *Queue) Size() int {
	return len(o.queue)
}

func (o *Queue) IsEmpty() bool {
	return len(o.queue) == 0
}

func NewQueue() *Queue {
	return &Queue{queue: make([]Step, 0)}
}
