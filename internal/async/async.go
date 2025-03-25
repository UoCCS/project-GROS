package async

import (
	"context"
	"sync"
)

// Future represents a value that will be available in the future.
type Future[T any] struct {
	result chan T
	err    chan error
}

// Await waits for the result of the Future. It blocks until the result is available.
func (f *Future[T]) Await() (T, error) {
	select {
	case res := <-f.result:
		return res, nil
	case err := <-f.err:
		var zero T
		return zero, err
	}
}

// Async runs a function asynchronously and returns a Future.
func Async[T any](fn func() (T, error)) *Future[T] {
	future := &Future[T]{
		result: make(chan T, 1),
		err:    make(chan error, 1),
	}

	go func() {
		defer close(future.result)
		defer close(future.err)

		res, err := fn()
		if err != nil {
			future.err <- err
			return
		}
		future.result <- res
	}()

	return future
}

// Join combines multiple Futures and waits for all of them to complete.
func Join[T any](futures ...*Future[T]) ([]T, error) {
	var wg sync.WaitGroup
	results := make([]T, len(futures))
	errs := make([]error, len(futures))

	for i, future := range futures {
		wg.Add(1)
		go func(i int, future *Future[T]) {
			defer wg.Done()
			res, err := future.Await()
			results[i] = res
			errs[i] = err
		}(i, future)
	}

	wg.Wait()

	// Check if any error occurred
	for _, err := range errs {
		if err != nil {
			return nil, err
		}
	}

	return results, nil
}

// Select waits for the first Future to complete and returns its result.
func Select[T any](ctx context.Context, futures ...*Future[T]) (T, error) {
	resultChan := make(chan T, len(futures))
	errChan := make(chan error, len(futures))

	for _, future := range futures {
		go func(f *Future[T]) {
			res, err := f.Await()
			if err != nil {
				errChan <- err
				return
			}
			resultChan <- res
		}(future)
	}

	select {
	case res := <-resultChan:
		return res, nil
	case err := <-errChan:
		var zero T
		return zero, err
	case <-ctx.Done():
		var zero T
		return zero, ctx.Err()
	}
}
