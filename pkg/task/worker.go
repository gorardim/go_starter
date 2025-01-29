package task

import "sync"

type Worker struct {
	maxWorker int
	ch        chan func()
	wg        sync.WaitGroup
}

func NewWorker(maxWorker int) *Worker {
	w := &Worker{
		maxWorker: maxWorker,
		ch:        make(chan func()),
	}
	for i := 0; i < maxWorker; i++ {
		go func() {
			for f := range w.ch {
				w.done(f)
			}
		}()
	}
	return w
}

func (w *Worker) Go(f func()) {
	w.wg.Add(1)
	w.ch <- f
}

func (w *Worker) Wait() {
	w.wg.Wait()
	close(w.ch)
}

func (w *Worker) done(f func()) {
	defer w.wg.Done()
	f()
}
