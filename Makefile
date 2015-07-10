VERSION = 0.0.1
DATETIME = $(shell date)
COMMIT = $(shell git rev-parse HEAD)
BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
AUTHOR = $(shell git config user.name)

LDFLAGS += -X github.com/fellah/version.version '$(VERSION)'
LDFLAGS += -X github.com/fellah/version.dateTime '$(DATETIME)'
LDFLAGS += -X github.com/fellah/version.commit '$(COMMIT)'
LDFLAGS += -X github.com/fellah/version.branch '$(BRANCH)'
LDFLAGS += -X github.com/fellah/version.author '$(AUTHOR)'

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
