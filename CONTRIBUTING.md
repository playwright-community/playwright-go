# Contributing

## Tests

### Test coverage

For every Pull Request on GitHub and on the main branch the coverage data will get sent over to [Coveralls](https://coveralls.io/github/playwright-community/playwright-go), this is helpful for finding functions that aren't covered by tests.

### Running tests

You can use the `BROWSER` environment variable to use a different browser than Chromium for the tests and use the `HEADLESS` environment variable which is useful for debugging.

```
BROWSER=chromium HEADLESS=1 go test -v --race
```
