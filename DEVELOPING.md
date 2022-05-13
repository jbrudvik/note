# Developing

## Overview

- No pull requests. Just commit to main. Until there are collaborators.
- No squash commits needed.

## Before committing

```
$ go fmt
$ go test  # All tests must pass
```

## Releasing

Releases are manually created with these steps:

1. Ensure linting and tests are passing
2. Increment the version in `./main.go`
3. [Draft a new release](https://github.com/jbrudvik/note/releases/new) with a new tag that matches the version from the previous step
4. Build Mac binary (`go build`), then attach it `./note` to the draft release
5. Publish the release
