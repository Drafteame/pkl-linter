# pkl-linter

`pkl-linter` is a command-line interface (CLI) tool designed to lint PKL files. It provides options to lint specific 
files or by default search for `.pkl` files on a provided directory, apply or omit certain rules, and exclude files or 
folders based on regular expressions.

## Usage

```
pkl-linter <path>
```

Or if you want to lint specific files:

```
pkl-linter --files="file1.pkl,file2.pkl"
```

### Example

```
pkl-linter ./my/path --recursive
```

## Installation

To install `pkl-linter`, you need to have Go installed on your machine. Then you can use the following command to 
install the tool:

```sh
go install github.com/yourusername/pkl-linter@latest
```

## Command-line Options

- `--files`, `-i`: Specify files to lint. Provide a list of files.
- `--recursive`, `-r`: Recursively search for files in the given path.
- `--excludeFolders`, `-e`: Exclude folders that match the given regular expressions.
- `--excludeFiles`, `-f`: Exclude files that match the given regular expressions.
- `--applyRules`, `-a`: Apply only the given rules.
- `--omitRules`, `-o`: Omit the given rules.

## Flags

### `--files`, `-i`

Use this flag to specify the files you want to lint. This takes a list of files as input.

If this flag is provided, the tool will only lint the files specified and will omit the recursive search and specified 
path.

```sh
pkl-linter --files="file1.pkl,file2.pkl"
```

### `--recursive`, `-r`

Use this flag to recursively search for files in the given path.

```sh
pkl-linter ./my/path --recursive
```

### `--excludeFolders`, `-e`

Use this flag to exclude folders that match the given regular expressions.

```sh
pkl-linter ./my/path --excludeFolders "test,docs"
```

### `--excludeFiles`, `-f`

Use this flag to exclude files that match the given regular expressions.

```sh
pkl-linter ./my/path --excludeFiles ".*_test.pkl,.*_docs.pkl"
```

### `--applyRules`, `-a`

Use this flag to apply only the specified rules.

```sh
pkl-linter ./my/path --applyRules "rule1,rule2"
```

### `--omitRules`, `-o`

Use this flag to omit certain rules.

```sh
pkl-linter ./my/path --omitRules "rule3,rule4"
```

## License

This project is licensed under the MIT License.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any changes.

## References

- [PKL Lang](https://pkl-lang.org/index.html)
