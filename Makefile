install: clean format vet
	go generate github.com/fellah/kb/assets
	go test
	go install -gcflags "-N -l"

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

clean:
	rm -f $(GOPATH)/bin/kb
