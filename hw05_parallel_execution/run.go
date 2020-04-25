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
	if N <= 0 || len(tasks) == 0 {
		return nil
	}

	if M <= 0 || M > len(tasks) {
		SimpleRun(tasks, N)
		return nil
	}

	tasksCh := make(chan Task, len(tasks))
	errorsCh := make(chan error, M)

	var wg sync.WaitGroup
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			worker(tasksCh, errorsCh)
		}()
	}

	for _, task := range tasks {
		tasksCh <- task
	}
	close(tasksCh)

	wg.Wait()

	close(errorsCh)
	if len(errorsCh) == M {
		return ErrErrorsLimitExceeded
	}
	return nil
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
