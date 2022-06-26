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

Releases are manually created with these steps:

1. Ensure code is formatted (`gofmt -d .`) and tests are passing (`go test`)
1. Increment the version in `main.go`
1. [Draft a new release](https://github.com/jbrudvik/note/releases/new) with a new tag that matches the version from the previous step
1. Build Mac binary (`go build`) and ensure it runs (`./note`)
1. Attach `./note` to the draft release
1. Add release notes since last release
1. Publish the release
