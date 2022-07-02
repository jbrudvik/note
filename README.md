# note

A CLI to append to the latest unshared Apple Notes note

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

```sh
$ go install github.com/jbrudvik/note@latest
```
