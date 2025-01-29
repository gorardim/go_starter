package task

import (
	"testing"
	"time"
)

func TestNewWorker(t *testing.T) {
	t.Run("right", func(t *testing.T) {
		w := NewWorker(2)
		for i := 0; i < 10; i++ {
			i := i
			w.Go(func() {
				t.Log(i)
				time.Sleep(time.Millisecond * 10)
			})
		}
		w.Wait()
	})

	t.Run("wrong", func(t *testing.T) {
		w := NewWorker(2)
		go func() {
			for i := 0; i < 10; i++ {
				i := i
				w.Go(func() {
					t.Log(i)
					time.Sleep(time.Millisecond * 10)
				})
			}
		}()
		w.Wait()
	})
}
