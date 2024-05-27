package root

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/Drafteame/pkl-linter/internal/files"
	"github.com/Drafteame/pkl-linter/internal/lint"
	"github.com/Drafteame/pkl-linter/internal/lo"
	"github.com/Drafteame/pkl-linter/internal/rules"
	"github.com/Drafteame/pkl-linter/pkg/version"
)

var rootCmd = &cobra.Command{
	Use:     "pkl-linter <path>",
	Example: "pkl-linter ./my/path --recursive",
	Version: version.Version,
	Args:    args,
	Run:     run,
}

var (
	filesLint      []string
	recursive      bool
	excludeFolders []string
	excludeFiles   []string
	applyRules     []string
	omitRules      []string
)

func init() {
	rootCmd.PersistentFlags().StringSliceVarP(&filesLint, "files", "i", []string{}, "Files to lint")
	rootCmd.PersistentFlags().BoolVarP(&recursive, "recursive", "r", false, "Recursively search for files in the given path")
	rootCmd.PersistentFlags().StringSliceVarP(&excludeFolders, "excludeFolders", "e", []string{}, "Exclude files that match the given regular expressions")
	rootCmd.PersistentFlags().StringSliceVarP(&excludeFiles, "excludeFiles", "f", []string{}, "Exclude folders that match the given regular expressions")
	rootCmd.PersistentFlags().StringSliceVarP(&applyRules, "applyRules", "a", []string{}, "Apply only the given rules")
	rootCmd.PersistentFlags().StringSliceVarP(&omitRules, "omitRules", "o", []string{}, "Omit the given rules")
}

func GetCmd() *cobra.Command {
	return rootCmd
}

func args(_ *cobra.Command, args []string) error {
	if len(filesLint) > 0 {
		return nil
	}

	if len(args) < 1 {
		return errors.New("requires a path argument")
	}

	return nil
}

func run(_ *cobra.Command, args []string) {
	if len(filesLint) > 0 {
		lintFiles(filesLint)
		return
	}

	path := args[0]

	lintPath(path)
}

func getRules() []rules.Rule {
	finalRules := rules.GetRules(applyRules, rules.WithSubset(rules.AllRules))
	finalRules = rules.OmitRules(omitRules, rules.WithSubset(finalRules))
	return finalRules
}

func lintFiles(filesLint []string) {
	var errs []error

	ruleset := getRules()

	for _, fileLint := range filesLint {
		errLint := lint.ExecRules(fileLint, lo.Map(ruleset, func(_ int, item rules.Rule) lint.Rule {
			return item
		})...)

		if errLint != nil {
			errs = append(errs, errLint...)
		}
	}

	printErrors(errs...)
}

func lintPath(path string) {
	lisFiles, errList := getPklFiles(path)
	if errList != nil {
		panic(errList)
	}

	lintFiles(lisFiles)
}

func getPklFiles(path string) ([]string, error) {
	var opts []files.SearchOption

	if recursive {
		opts = append(opts, files.WithRecursive())
	}

	if len(excludeFolders) > 0 {
		opts = append(opts, files.WithExcludeFolders(excludeFolders...))
	}

	if len(excludeFiles) > 0 {
		opts = append(opts, files.WithExcludeFiles(excludeFiles...))
	}

	return files.Search(path, "\\.pkl$", opts...)
}

func printErrors(errors ...error) {
	for _, mainErr := range errors {
		errs := unwrapAllErrors(mainErr)

		for _, err := range errs {
			fmt.Println(err)
		}
	}

	if len(errors) > 0 {
		os.Exit(1)
	}
}

func unwrapAllErrors(err error) []error {
	var errs []error

	for err != nil {
		errs = append(errs, err)
		err = errors.Unwrap(err)
	}

	return errs
}
