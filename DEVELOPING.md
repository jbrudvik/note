# Developing

## Overview

- No pull requests. Just commit to main. Until there are collaborators.
- No squash commits needed.

## Before committing

```sh
$ go fmt
$ go test  # All tests must pass
```

## Testing the AppleScript

`note` uses AppleScript to drive Notes.app. The underlying AppleScript script can be tested like this:

```sh
$ osascript append_to_latest_unshared_note.txt "text to append"
```

Please note that the formatting of the text (as HTML) is set in the containing Go code.

## Releasing

Releases are manually created with these steps:

1. Ensure code is formatted (`gofmt -d .`) and tests are passing (`go test`)
2. Increment the version in `./main.go`
3. [Draft a new release](https://github.com/jbrudvik/note/releases/new) with a new tag that matches the version from the previous step
4. Build Mac binary (`go build`), then attach it `./note` to the draft release
5. Publish the release
