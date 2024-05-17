package hw05parallelexecution

import (
	"errors"
	"sync"
)

var (
	ErrErrorsLimitExceeded       = errors.New("errors limit exceeded")
	ErrIncorrectGoroutinesNumber = errors.New("goroutine number must be higher then zero")
)

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var wg sync.WaitGroup

	if n <= 0 {
		return ErrIncorrectGoroutinesNumber
	}

	errorsCount := 0
	tasksChannel := make(chan Task, len(tasks))
	limitCheck := collectErrors(m, &errorsCount)

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer func() {
				wg.Done()
			}()
			worker(tasksChannel, limitCheck)
		}()
	}

	for _, task := range tasks {
		tasksChannel <- task
	}
	close(tasksChannel)

	wg.Wait()

	if errorsCount >= m {
		return ErrErrorsLimitExceeded
	}

	return nil
}

func worker(tasks <-chan Task, limitCheck func(error) error) {
	for {
		t, ok := <-tasks

		if !ok {
			return
		}

		if taskError := t(); taskError != nil {
			if e := limitCheck(taskError); e != nil {
				return
			}
		}
	}
}

func collectErrors(limit int, errorsCount *int) func(error) error {
	var mu sync.Mutex

	return func(e error) error {
		mu.Lock()
		defer mu.Unlock()
		*errorsCount++

		if limit <= 0 || *errorsCount >= limit {
			return ErrErrorsLimitExceeded
		}

		return nil
	}
}
