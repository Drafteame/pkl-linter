package rules

import (
	"path/filepath"
	"regexp"
	"strings"

	ierrors "github.com/Drafteame/pkl-linter/internal/errors"
)

type fileNaming struct {
	name        string
	description string
}

func (s *fileNaming) Name() string {
	return s.name
}

func (s *fileNaming) Description() string {
	return s.description
}

func (s *fileNaming) Execute(file File) error {
	filePath := file.Path()
	fileName := strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))

	patterns := []string{
		`^[A-Z][a-zA-Z0-9]*$`,      // PascalCase
		`^[a-z]+[a-zA-Z0-9]*$`,     // camelCase
		`^[a-z0-9]+(-[a-z0-9]+)*$`, // kebab-case
	}

	for _, pattern := range patterns {
		if match, _ := regexp.MatchString(pattern, fileName); match {
			return nil
		}
	}

	return ierrors.LintError{
		RuleName:    s.name,
		Description: s.description,
		LineNumber:  0,
		FilePath:    filePath,
	}
}

// FileNaming checks if the file naming is correct.
func FileNaming() Rule {
	return &fileNaming{
		name:        "file-naming",
		description: "File naming should be in PascalCase, camelCase or kebab-case.",
	}
}
