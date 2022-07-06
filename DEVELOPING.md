# Developing

## Required dependencies

- [Go 1.18](https://go.dev/doc/install)
- [goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports): `$ go install golang.org/x/tools/cmd/goimports@latest`

## Before committing

```sh
# Format all code correctly
$ goimports -w .

# Ensure code builds succesfully
$ go build

# Ensure code is free of common mistakes
$ go vet ./...

# Ensure all tests pass
$ go test ./...
```

## Manual testing

### AppleScript

`note` uses AppleScript to drive Notes.app. The underlying AppleScript script can be tested like this:

```sh
$ osascript append_to_latest_unshared_note.txt "Text to append"
```

Please note that the formatting of the text (as HTML) is set in the containing Go code.

## Releasing

1. Ensure build is passing: [![Build](https://github.com/jbrudvik/note/actions/workflows/build.yml/badge.svg)](https://github.com/jbrudvik/gmc/actions/workflows/build.yml)
1. [Create a new release](https://github.com/jbrudvik/note/releases/new) with:
   - Version: Incremented in format: vX.Y.Z
   - Release title: note `<version-from-last-step>`
   - Release notes
