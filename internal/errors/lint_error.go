package errors

import (
	"fmt"
)

// LintError represents a linting error
type LintError struct {
	RuleName    string
	Description string
	LineNumber  int
	FilePath    string
}

func (e LintError) Error() string {
	file := ""

	if e.LineNumber == 0 {
		file = e.FilePath
	} else {
		file = fmt.Sprintf("%s:%d", e.FilePath, e.LineNumber)
	}

	return fmt.Sprintf("[%s] %s %s", e.RuleName, file, e.Description)
}
