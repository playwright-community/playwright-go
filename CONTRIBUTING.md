# Contributing

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
1. Apply patch `bash scripts/apply-patch.sh`
1. Fix merge conflicts if any, otherwise ignore this step. Once you are happy you can commit the changes `cd playwright; git commit -am "apply patch" && cd ..`
1. Regenerate a new patch `bash scripts/update-patch.sh`
1. Generate go code `go generate ./...`

To adapt to the new version of Playwright's protocol and feature updates, you may need to modify the patch. Refer to the following steps:

1. Apply patch `bash scripts/apply-patch.sh`
1. `cd playwright`
1. Revert the patch`git reset HEAD~1`
1. Modify the files under `docs/src/api`, etc. as needed. Available references:
    - Protocol `packages/protocol/src/protocol.yml`, prior to v1.27.0 are available in `packages/playwright-core/src/protocol/protocol.yml`
    - [Playwright python](https://github.com/microsoft/playwright-python)
1. Commit the changes `git commit -am "apply patch"`
1. Regenerate a new patch `bash scripts/update-patch.sh`
1. Generate go code `go generate ./...`. If you updated `scripts/data/interfaces.json`, do this step again.
