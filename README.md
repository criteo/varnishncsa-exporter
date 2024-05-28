[![Go](https://img.shields.io/github/go-mod/go-version/criteo/golang-template)](https://github.com/criteo/golang-template)
[![status](https://img.shields.io/badge/status-template-blue)](https://github.com/criteo/golang-template)
[![CI](https://github.com/criteo/golang-template/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/criteo/golang-template/actions/workflows/ci.yml)
[![GitHub](https://img.shields.io/github/license/criteo/golang-template)](https://github.com/criteo/golang-template/blob/main/LICENSE)

# Criteo Golang template

This repository contains the template used to bootstrap GitHub Golang projects at Criteo.

## How to bootstrap your own repository

Rename the Go module, and clean the example code:
```
go mod edit -module <github.com/your-org/your-awesome-project>
rm cmd/example -rf
```

Set the build details in `.goreleaser.yaml`:
```
  - id: example
    binary: example
    main: ./cmd/example
    env:
      - CGO_ENABLED=0
    goos:
      - linux
```

## Continuous Integration

Tests are run automatically on Pull Requests and Push events:
* [golangci-lint](https://golangci-lint.run/)
* `go test`
* `go build`
* [commitlint/config-conventional](https://github.com/conventional-changelog/commitlint)

## Releases

* Step 1: push your changes
```
git add main.go
git commit "feat: adding an awesome feature"
git push
```
* Step 2: merge the related PR from the GitHub web UI
* Step 3: add a new tag
```
git tags         # get the list of existing tags
git tag v0.0.2   # create a tag on the current local HEAD
git push --tags  # push only the tag to GitHub
```

You are encouraged to use [Semantic Versioning](https://semver.org/).

Once the tag is pushed, GitHub action will run GoReleaser automatically.
GoReleaser will publish a new release with an auto-generated changelog using the commits message.
