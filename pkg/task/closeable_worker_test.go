package task

import (
	"testing"
)

func TestNewStoppableWorker(t *testing.T) {
	w := NewStoppableWorker(2)
	go func() {
		defer w.Close()
		for i := 0; i < 10; i++ {
			i := i
			w.Go(func() {
				t.Log(i)
			})
		}
	}()
	w.Wait()
}
