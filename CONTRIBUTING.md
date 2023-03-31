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

1. Find out to which upstream version you want to roll, you need to find out the upstream commit SHA.
1. `bash scripts/apply-patch.sh`
1. `cd playwright`
1. `git reset HEAD~1` this reverts the custom patches
1. `git stash`
1. checkout new new sha `git checkout <sha>`
1. apply the patch again `git stash pop` (and fix merge conflicts)
1. once you are happy you can commit the changes `git commit -am "apply patch"`
1. regenerate a new patch `bash scripts/update-patch.sh`
