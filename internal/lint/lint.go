package lint

import (
	"runtime"
	"sync"

	"github.com/Drafteame/pkl-linter/internal/files"
	"github.com/Drafteame/pkl-linter/internal/rules"
)

type Rule interface {
	Execute(file rules.File) error
}

func ExecRules(filePath string, rules ...Rule) []error {
	var wg sync.WaitGroup
	sem := make(chan struct{}, runtime.NumCPU())
	errChan := make(chan error, len(rules))

	for _, rule := range rules {
		wg.Add(1)
		sem <- struct{}{} // Acquire semaphore slot

		go func(rule Rule) {
			defer wg.Done()
			defer func() { <-sem }() // Release semaphore slot

			if err := rule.Execute(files.NewFile(filePath)); err != nil {
				errChan <- err
			}
		}(rule)
	}

	wg.Wait()
	close(errChan)

	// Collect all errors from the error channel
	var errs []error

	for err := range errChan {
		if err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}
