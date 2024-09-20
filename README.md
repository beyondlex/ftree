
# File Tree Printer

A simple CLI tool based on [Cobra](https://cobra.dev/) for printing files as a tree under a specific directory. This project serves as a practice exercise for learning Go.

## Features

- Print file structure in a tree format.
- Limit the number of printed lines.
- Control the depth of directory traversal.

## Installation

To install, clone the repository and build the application:

```bash
git clone https://github.com/beyondlex/ftree
cd ftree
go build
```

## Usage

Run the CLI with:

```bash
ftree
```

Example output:

```
- ./
  - LICENSE
  - README.md
  - cmd/
    - printer.go
    - root.go
  - go.mod
  - go.sum
  - main.go
  - util/
    - file_util.go
```

**Note:** The output above can be used as the content of the [FileTree](https://starlight.astro.build/components/file-tree/) component of [Starlight](https://starlight.astro.build/), which is one of the reasons I created this CLI program :)


To print files under a specific directory with constraints:

```bash
ftree ~/Downloads -l 10 -d 3
```

This command will print files under the `~/Downloads` directory, limiting output to 10 lines and traversing up to 3 levels deep.

### Additional Examples

Print files in the parent directory:

```bash
ftree ..
```

Help:

```bash
ftree -h

Examples:
	ftree
	ftree ..
	ftree ~/Downloads
	ftree ~/Downloads -l=10 -d=1
	ftree .. -l 10 -d 1
	ftree -l 10 -d 1

Usage:
  ftree [filepath] [flags]

Flags:
  -d, --depth int   Limit depth for printing (default 8)
  -h, --help        help for ftree
  -l, --lines int   Limit lines for printing (default 50)
```

