package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func worker(tasks <-chan Task, errors chan error) {
	for task := range tasks {
		err := task()
		if err != nil {
			select {
			case errors <- err:
			default:
				return
			}
		}
	}
}

// Run starts tasks in N goroutines and stops its work when receiving M errors from tasks
func Run(tasks []Task, N int, M int) error { //nolint:gocritic
	if M <= 0 {
		SimpleRun(tasks, N)
		return nil
	}

	tasksCh := make(chan Task, len(tasks))
	errorsCh := make(chan error, M)

	var wg sync.WaitGroup
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			worker(tasksCh, errorsCh)
			wg.Done()
		}()
	}

	for _, task := range tasks {
		tasksCh <- task
	}
	close(tasksCh)

	wg.Wait()

	select {
	case errorsCh <- errors.New(""):
		return nil
	default:
		return ErrErrorsLimitExceeded
	}
}

func SimpleRun(tasks []Task, N int) { //nolint:gocritic
	tasksCh := make(chan Task, len(tasks))

	var wg sync.WaitGroup
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			for task := range tasksCh {
				_ = task()
			}
			wg.Done()
		}()
	}

	for _, task := range tasks {
		tasksCh <- task
	}
	close(tasksCh)

	wg.Wait()
}
