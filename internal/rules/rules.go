package rules

import (
	"slices"

	"github.com/Drafteame/pkl-linter/internal/files"
	"github.com/Drafteame/pkl-linter/internal/lo"
)

type File interface {
	Scan(line files.ScanLine) error
	Path() string
}

type Rule interface {
	Execute(file File) error
	Name() string
	Description() string
}

var AllRules = []Rule{
	StringInterpolation(),
	AncestorPaths(),
	FileNaming(),
	PropertySpacing(),
}

type rulesOptions struct {
	subset []Rule
}

type Option func(*rulesOptions)

func WithSubset(rules []Rule) Option {
	return func(opts *rulesOptions) {
		if len(rules) == 0 {
			return
		}

		opts.subset = rules
	}
}

func GetRules(ruleNames []string, opts ...Option) []Rule {
	options := rulesOptions{
		subset: AllRules,
	}

	for _, option := range opts {
		option(&options)
	}

	return lo.Filter(options.subset, func(_ int, item Rule) bool {
		return slices.Contains(ruleNames, item.Name())
	})
}

func OmitRules(ruleNames []string, opts ...Option) []Rule {
	options := rulesOptions{
		subset: AllRules,
	}

	for _, option := range opts {
		option(&options)
	}

	return lo.Filter(options.subset, func(_ int, item Rule) bool {
		return !slices.Contains(ruleNames, item.Name())
	})
}
