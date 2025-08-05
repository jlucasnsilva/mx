test:
	go test -timeout 10s ./...

test-verbose:
	go test -v -timeout 10s ./...

coverage-test:
	touch coverage.out
	go test -timeout 10s -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
