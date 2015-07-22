VERSION := DateTime:\t$(shell date)\n
VERSION := $(VERSION)Commit:\t\t$(shell git rev-parse HEAD)\n
VERSION := $(VERSION)Branch:\t\t$(shell git rev-parse --abbrev-ref HEAD)

LDFLAGS += -X github.com/fellah/version.version '$(VERSION)'

install: clean format vet assets
	go generate github.com/fellah/kb/assets
	go test
	go install -ldflags="$(LDFLAGS)" -gcflags "-N -l"

clean:
	rm -f $(GOPATH)/bin/kb

format:
	go fmt github.com/fellah/kb
	go fmt github.com/fellah/kb/assets
	go fmt github.com/fellah/kb/cache
	go fmt github.com/fellah/kb/markdown
	go fmt github.com/fellah/kb/web

vet:
	go vet github.com/fellah/kb
	go vet github.com/fellah/kb/assets
	go vet github.com/fellah/kb/cache
	go vet github.com/fellah/kb/markdown
	go vet github.com/fellah/kb/web
