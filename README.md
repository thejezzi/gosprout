# mkgo

![build](https://github.com/thejezzi/mkgo/actions/workflows/build.yml/badge.svg)

`mkgo` is a command-line tool for creating new Go projects. It can be used to create a new project
from scratch, or to clone an existing Git repository and use it as a template.

## Installation

To install `mkgo`, use `go install`:

```
go install github.com/thejezzi/mkgo/cmd/mkgo@latest
```

## Usage

To create a new project, run `mkgo` with the name of the project as an argument:

```
mkgo my-new-project
```

This will create a new directory called `my-new-project` in the current working directory, and
initialize it as a Go module.

### Flags

`mkgo` supports the following flags:

- `-path`: The path to create the project in. Defaults to the project name.
- `-template`: The template to use for the project. Defaults to `simple`.
- `-git`: The Git repository to clone from.
- `-makefile`: Create a Makefile.
- `-init-git`: Initialize a new Git repository.

### Templates

`mkgo` supports the following templates:

- `simple`: A simple project with a `main.go` file.
- `cli`: A project with a `main.go` file and a `cmd` directory.
- `git`: Clones a git repository and uses it as a template. This is useful for starting a new
  project from an existing template.

## Contributing

Contributions are welcome! Please open an issue or a pull request if you have any ideas or
suggestions.
