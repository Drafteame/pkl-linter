package files

import (
	"os"
	"path/filepath"
	"regexp"
)

type searchOptions struct {
	excludeFolders []string
	excludeFiles   []string
	recursive      bool
}

type SearchOption func(*searchOptions)

func WithExcludeFiles(files ...string) SearchOption {
	return func(o *searchOptions) {
		if o.excludeFiles == nil {
			o.excludeFiles = make([]string, 0)
		}

		o.excludeFiles = append(o.excludeFiles, files...)
	}
}

func WithExcludeFolders(folders ...string) SearchOption {
	return func(o *searchOptions) {
		if o.excludeFolders == nil {
			o.excludeFolders = make([]string, 0)
		}

		o.excludeFolders = append(o.excludeFolders, folders...)
	}

}

func WithRecursive() SearchOption {
	return func(o *searchOptions) {
		o.recursive = true
	}
}

// Search searches for a file in a given path and returns a list of files that match the given file name.
func Search(rootPath, regExp string, opts ...SearchOption) ([]string, error) {
	options := searchOptions{}

	for _, opt := range opts {
		opt(&options)
	}

	return search(rootPath, regExp, make([]string, 0), options)
}

func search(rootPath, regExp string, files []string, options searchOptions) ([]string, error) {
	entries, err := os.ReadDir(rootPath)
	if err != nil {
		return nil, err
	}

	exp := regexp.MustCompile(regExp)

	for _, entry := range entries {
		path := filepath.Join(rootPath, entry.Name())

		if entry.IsDir() && options.recursive && !shouldExcludeFolder(entry.Name(), options) {
			files, err = search(path, regExp, files, options)
			if err != nil {
				return files, err
			}

			continue
		}

		if !shouldExcludeFile(entry.Name(), options) && exp.MatchString(entry.Name()) {
			files = append(files, filepath.Join(rootPath, entry.Name()))
		}
	}

	return files, nil
}

func shouldExcludeFolder(folder string, opts searchOptions) bool {
	for _, omit := range opts.excludeFolders {
		if omit == folder {
			return true
		}
	}

	return false
}

func shouldExcludeFile(file string, opts searchOptions) bool {
	for _, omit := range opts.excludeFiles {
		if omit == file {
			return true
		}
	}

	return false
}
