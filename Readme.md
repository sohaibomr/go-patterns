# Go test coverage

## How to run

1. Run tests of all packages `go test -v -coverprofile=coverage.out ./...` and generate test coverage file.
2. view coverage file in html `go tool cover -html=coverage.out`
3. `go test -cover ./...` to print coverage score on console
