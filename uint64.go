package sequence

type UInt64 struct {
    start, step uint64
    queue       chan uint64
    running     bool
}

func NewUInt64(start, step uint64) (ai *UInt64) {
    ai = &UInt64{
        start:   start,
        step:    step,
        running: true,
        queue:   make(chan uint64, 4),
    }
    go ai.process()
    return
}

func (ai *UInt64) process() {
    defer func() { recover() }()
    for i := ai.start; ai.running; i = i + ai.step {
        ai.queue <- i
    }
}

func (ai *UInt64) Next() uint64 {
    return <-ai.queue
}

func (ai *UInt64) Close() {
    ai.running = false
    close(ai.queue)
}
