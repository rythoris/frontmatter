# Frontmatter

`frontmatter` is a simple command-line utility for working with markdown frontmatter.

```
Usage: frontmatter [--format FORMAT] [--content] FILE

Positional arguments:
  FILE                   markdown file path

Options:
  --format FORMAT, -f FORMAT
                         frontmatter output format (possible values: 'yaml', 'json') [default: json]
  --content, -c          only print content of the file
  --help, -h             display this help and exit
```

`frontmatter` wouldn't be possible without [`adrg/frontmatter`](https://github.com/adrg/frontmatter) library. Consider showing them some love by starring their repository.

## Installation

```console
go install -v github.com/rythoris/frontmatter@latest
```

## Examples

```sh
# extract 'title' from test.md
frontmatter -f json ./test.md | jq -r '.title'

# extract content of the test.md
frontmatter -c ./test.md
```

## FAQ

### Why are you using [`json-iterator/go`](https://github.com/json-iterator/go) instead of [`encoding/json`](https://pkg.go.dev/encoding/json)

Because apperantly `encoding/json` doesn't not support unmarshalling to `map[interface {}]interface {}` type.

## License

This project is licensed under [MIT License](https://opensource.org/license/mit/). See [LICENSE](./LICENSE) for more details.
