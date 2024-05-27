package rules

import (
	"regexp"

	ierrors "github.com/Drafteame/pkl-linter/internal/errors"
	"github.com/Drafteame/pkl-linter/internal/files"
)

type stringInterpolation struct {
	name        string
	description string
}

func (s *stringInterpolation) Name() string {
	return s.name
}

func (s *stringInterpolation) Description() string {
	return s.description
}

func (s *stringInterpolation) Execute(file File) error {
	interpolationRegex := regexp.MustCompile(`".*"\s*\+\s*.*|.*\s*\+\s*".*"|.*\s*\+\s*.*`)

	return file.Scan(func(info files.FileScanInfo) error {
		if !interpolationRegex.MatchString(info.Text()) {
			return nil
		}

		return ierrors.LintError{
			RuleName:    s.name,
			Description: s.description,
			LineNumber:  info.LineNumber(),
			FilePath:    info.FilePath(),
		}
	})
}

// StringInterpolation checks if the ancestors paths are being used with triple dots instead of multiple two dots format.
func StringInterpolation() Rule {
	return &stringInterpolation{
		name:        "string-interpolation",
		description: "String interpolation should not use + operator, instead use \"\\(var)\" operator",
	}
}
