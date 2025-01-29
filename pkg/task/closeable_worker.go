package task

type CloseableWorker struct {
	w    *Worker
	quit chan struct{}
}

func NewStoppableWorker(maxWorker int) *CloseableWorker {
	return &CloseableWorker{
		w:    NewWorker(maxWorker),
		quit: make(chan struct{}),
	}
}

func (w *CloseableWorker) Go(f func()) {
	w.w.Go(f)
}

func (w *CloseableWorker) Wait() {
	<-w.quit
	w.w.Wait()
}

func (w *CloseableWorker) Close() {
	close(w.quit)
}
