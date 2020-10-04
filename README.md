# Go Coverage

This image performs the following tasks in order:
- calculates the test coverage percentage;
- compares the percentage with the `MINIMUM_TEST_COVERAGE` env var and exits if it's lower than that;
- creates a badge in the readme file like this ![coverage-badge-do-not-edit](https://img.shields.io/badge/Coverage-80%25-green.svg?longCache=true&style=flat)

## Usage

Make sure the readme file contains the following placeholder which will be replaced with an actual badge showing
the coverage percentage with different colors:

```
![coverage-badge-do-not-edit]()
```

Then you need to run tests with these flags to produce an output file `coverage.out`:

```bash
go test -covermode=count -coverprofile=coverage.out ./..."
```

Then you can use the image:

```bash
docker run --rm \
    -e MINIMUM_COVERAGE=70 \
    -e COVERAGE_FILE=/app/coverage.out \
    -e README_FILE=/app/README.md \
    -v $(PWD):/app \
    imagename
```

