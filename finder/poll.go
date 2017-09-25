package finder

import (
	"log"
	"sync"
)

type Pool struct {
	size        int
	currentSize int

	queue   chan Tasker
	queueWG sync.WaitGroup

	result   chan int
	resultWG sync.WaitGroup

	Total int

	log *log.Logger
}

func NewPool(concurrency int, logger *log.Logger) *Pool {
	p := &Pool{
		size:   concurrency,
		queue:  make(chan Tasker, concurrency),
		result: make(chan int, concurrency),
		log:    logger,
	}
	p.resultWG.Add(1)
	go func() {
		defer p.resultWG.Done()
		for i := range p.result {
			p.Total += i
		}
	}()
	return p
}

func (p *Pool) Put(task Tasker) {
	if p.currentSize < p.size {
		p.currentSize++
		p.queueWG.Add(1)
		go func() {
			defer p.queueWG.Done()
			for task := range p.queue {
				count, err := task.Run()
				if err != nil {
					p.log.Printf("Error for %s: %s", task.GetSource(), err)
				} else {
					p.log.Printf("Count for %s: %d", task.GetSource(), count)
					if count > 0 {
						p.result <- count
					}
				}
			}
		}()
	}
	p.queue <- task
}

func (p *Pool) StopPoolAndWait() {
	p.currentSize = 0

	close(p.queue)
	p.queueWG.Wait()

	close(p.result)
	p.resultWG.Wait()
}
