test:
	go test ./... -cover -v

lint:
	golint ./...

.PHONY: test