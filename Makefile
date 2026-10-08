.PHONY: build test vet run clean

BINARY = listkit

build:
	CGO_ENABLED=0 go build -trimpath ./cmd/$(BINARY)

test:
	go test -race ./...

vet:
	go vet ./...

run: build
	./$(BINARY) -no-auth

clean:
	rm -f $(BINARY)
