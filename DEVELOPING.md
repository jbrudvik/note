# Developing

## Before committing

```sh
$ go fmt
$ go test ./... # All tests must pass
```

## Manual testing

### AppleScript

`note` uses AppleScript to drive Notes.app. The underlying AppleScript script can be tested like this:

```sh
$ osascript append_to_latest_unshared_note.txt "Text to append"
```

Please note that the formatting of the text (as HTML) is set in the containing Go code.

## Releasing

1. Ensure code is formatted (`gofmt -d .`) and tests are passing (`go test ./...`)
1. [Create a new release](https://github.com/jbrudvik/note/releases/new) with new version
