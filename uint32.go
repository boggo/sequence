package sequence

type UInt32 struct {
    start, step uint32
    queue       chan uint32
    running     bool
}

func NewUInt32(start, step uint32) (ai *UInt32) {
    ai = &UInt32{
        start:   start,
        step:    step,
        running: true,
        queue:   make(chan uint32, 4),
    }
    go ai.process()
    return
}

func (ai *UInt32) process() {
    defer func() { recover() }()
    for i := ai.start; ai.running; i = i + ai.step {
        ai.queue <- i
    }
}

func (ai *UInt32) Next() uint32 {
    return <-ai.queue
}

func (ai *UInt32) Close() {
    ai.running = false
    close(ai.queue)
}
