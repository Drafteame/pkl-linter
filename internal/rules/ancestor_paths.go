package rules

import (
	"regexp"

	ierrors "github.com/Drafteame/pkl-linter/internal/errors"
	"github.com/Drafteame/pkl-linter/internal/files"
)

type ancestorPaths struct {
	name        string
	description string
}

func (s *ancestorPaths) Name() string {
	return s.name
}

func (s *ancestorPaths) Description() string {
	return s.description
}

func (s *ancestorPaths) Execute(file File) error {
	ancestorRegex := regexp.MustCompile(`((extends|import\*?|amends)\s"(\.\./)+.+")|(import\*?\("(\.\./)+.+"\))`)

	return file.Scan(func(info files.FileScanInfo) error {
		if !ancestorRegex.MatchString(info.Text()) {
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

// AncestorPaths checks if the ancestors paths are being used with triple dots instead of multiple two dots format.
func AncestorPaths() Rule {
	return &ancestorPaths{
		name:        "ancestor-paths",
		description: "Ancestor paths should not use multiple two dots (../../) instead of single triple dots (.../)",
	}
}
