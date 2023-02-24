# note

A CLI to append to the latest unshared Apple Notes note

[![Build](https://github.com/jbrudvik/note/actions/workflows/build.yml/badge.svg)](https://github.com/jbrudvik/note/actions/workflows/build.yml) [![Go Reference](https://pkg.go.dev/badge/github.com/jbrudvik/note.svg)](https://pkg.go.dev/github.com/jbrudvik/note)

## Usage

```sh
$ note

NAME:
   note - Append to latest Apple Notes note

USAGE:
   note [global options] "Text to append"

VERSION:
   vX.Y.Z

DESCRIPTION:
   Ignores shared notes. Formats as new line by default.

GLOBAL OPTIONS:
   --bulleted, -b  Format text as bulleted list item (default: false)
   --help, -h      show help (default: false)
   --version, -v   print the version (default: false)
```

## Install

### Required dependencies

- [Go 1.20](https://go.dev/doc/install)

### Install note

```sh
$ go install github.com/jbrudvik/note@latest
```
