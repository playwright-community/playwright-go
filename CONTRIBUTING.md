# Contributing

## Code style
The Go code is linted with [golangci-lint](https://golangci-lint.run/) and formatted with [gofumpt](https://github.com/mvdan/gofumpt). Please configure your editor to run the tools while developing and make sure to run the tools before committing any code.

## Tests

### Test coverage

For every Pull Request on GitHub and on the main branch the coverage data will get sent over to [Coveralls](https://coveralls.io/github/playwright-community/playwright-go), this is helpful for finding functions that aren't covered by tests.

### Running tests

You can use the `BROWSER` environment variable to use a different browser than Chromium for the tests and use the `HEADLESS` environment variable which is useful for debugging.

```
BROWSER=chromium HEADLESS=1 go test -v --race ./...
```

### Roll

1. Find out to which upstream version you want to roll, and change the value of `playwrightCliVersion` in the **run.go** to the new version.
1. Download current version of Playwright driver `go run scripts/install-browsers/main.go`
1. Apply patch `bash scripts/apply-patch.sh`
1. Fix merge conflicts if any, otherwise ignore this step. Once you are happy you can commit the changes `cd playwright; git commit -am "apply patch" && cd ..`
1. Regenerate a new patch `bash scripts/update-patch.sh`
1. Generate go code `go generate ./...`

To adapt to the new version of Playwright's protocol and feature updates, you may need to modify the patch. Refer to the following steps:

1. Apply patch `bash scripts/apply-patch.sh`
1. `cd playwright`
1. Revert the patch`git reset HEAD~1`
1. Modify the files under `docs/src/api`, etc. as needed. Available references:
    - Protocol `packages/protocol/src/protocol.yml`
    - [Playwright python](https://github.com/microsoft/playwright-python)
1. Commit the changes `git commit -am "apply patch"`
1. Regenerate a new patch `bash scripts/update-patch.sh`
1. Generate go code `go generate ./...`.
