package rules

import (
	"fmt"
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
	singleLineString := `"(?:\\.|[^"\\])*"`
	rawSingleLineString := `#"(?:[^"#])*"#`
	multiLineString := `"""(?:\\.|[^"\\])*"""`
	rawMultiLineString := `#"""(?:[^"#])*"""#`

	stringPattern := fmt.Sprintf(`(?:%s|%s|%s|%s)`, singleLineString, rawSingleLineString, multiLineString, rawMultiLineString)

	interpolationRegex := regexp.MustCompile(fmt.Sprintf(`\s*=\s*(?:%s\s*\+\s*)+.*|(?:.*\s*\+\s*)+%s\s*`, stringPattern, stringPattern))

	validContentRegex := regexp.MustCompile(`\s*=\s*"\+".*|\s*=\s*".*\+".*`)

	return file.Scan(func(info files.FileScanInfo) error {
		if interpolationRegex.MatchString(info.Text()) && !validContentRegex.MatchString(info.Text()) {
			return ierrors.LintError{
				RuleName:    s.name,
				Description: s.description,
				LineNumber:  info.LineNumber(),
				FilePath:    info.FilePath(),
			}
		}

		return nil
	})
}

// StringInterpolation checks if the ancestors paths are being used with triple dots instead of multiple two dots format.
func StringInterpolation() Rule {
	return &stringInterpolation{
		name:        "string-interpolation",
		description: "String interpolation should not use + operator, instead use \"\\(var)\" operator",
	}
}
