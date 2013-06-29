package sequence

type Int64 struct {
    start, step int64
    queue       chan int64
    running     bool
}

func NewInt64(start, step int64) (ai *Int64) {
    ai = &Int64{
        start:   start,
        step:    step,
        running: true,
        queue:   make(chan int64, 4),
    }
    go ai.process()
    return
}

func (ai *Int64) process() {
    defer func() { recover() }()
    for i := ai.start; ai.running; i = i + ai.step {
        ai.queue <- i
    }
}

func (ai *Int64) Next() int64 {
    return <-ai.queue
}

func (ai *Int64) Close() {
    ai.running = false
    close(ai.queue)
}
