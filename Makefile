.PHONY: test mocks

test:
	go test -v ./internal/...

mocks:
	mockery --all --dir internal/interfaces --output internal/interfaces/mocks