package sequence

type Int32 struct {
    start, step int32
    queue       chan int32
    running     bool
}

func NewInt32(start, step int32) (ai *Int32) {
    ai = &Int32{
        start:   start,
        step:    step,
        running: true,
        queue:   make(chan int32, 4),
    }
    go ai.process()
    return
}

func (ai *Int32) process() {
    defer func() { recover() }()
    for i := ai.start; ai.running; i = i + ai.step {
        ai.queue <- i
    }
}

func (ai *Int32) Next() int32 {
    return <-ai.queue
}

func (ai *Int32) Close() {
    ai.running = false
    close(ai.queue)
}
