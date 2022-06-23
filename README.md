# note

A command-line tool to append to the latest unshared Apple Notes note

## Usage

```sh
$ note

NAME:
   note - Append to latest Apple Notes note

USAGE:
   note [global options] "Text to append"

VERSION:
   v0.0.7

DESCRIPTION:
   Ignores shared notes. Formats as new line by default.

GLOBAL OPTIONS:
   --bulleted, -b  Format text as bulleted list item (default: false)
   --help, -h      show help (default: false)
   --version, -v   print the version (default: false)
```

## Install

### With Go

```sh
$ go install github.com/jbrudvik/note
```

### With pre-built Mac binary

1. Download [note](https://github.com/jbrudvik/note/releases/latest/download/note)
2. `$ chmod u+x note`
3. [Tell Gatekeeper to allow](https://support.apple.com/en-us/HT202491) `note`
4. In the dialog that opens, allow your terminal to control Notes.app
