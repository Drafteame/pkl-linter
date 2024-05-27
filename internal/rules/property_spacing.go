package rules

import (
	"strings"

	ierrors "github.com/Drafteame/pkl-linter/internal/errors"
	"github.com/Drafteame/pkl-linter/internal/files"
)

type propertySpacing struct {
	name        string
	description string
}

func (s *propertySpacing) Name() string {
	return s.name
}

func (s *propertySpacing) Description() string {
	return s.description
}

func (s *propertySpacing) Execute(file File) error {
	blankLineCount := 0
	statementCount := 0

	return file.Scan(func(info files.FileScanInfo) error {
		text := info.Text()

		// TODO: Add code to work with blocks of code {} and [].

		// Scape for comment lines
		if strings.HasPrefix(strings.TrimSpace(text), "//") {
			blankLineCount = 0
			statementCount = 0
			return nil
		}

		if strings.TrimSpace(text) == "" {
			blankLineCount++
			statementCount = 0
			return nil
		} else {
			statementCount++
			blankLineCount = 0
		}

		if blankLineCount > 1 || statementCount > 1 {
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

// PropertySpacing checks if the properties are spaced correctly.
func PropertySpacing() Rule {
	return &propertySpacing{
		name:        "property-spacing",
		description: "Object members (properties, elements, and entries) should be separated by at most one blank line.",
	}
}
