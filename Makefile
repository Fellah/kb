install: format vet
	go generate kb/assets
	go test
	go install -gcflags "-N -l"

format:
	go fmt kb
	go fmt kb/assets
	go fmt kb/markdown
	go fmt kb/web

vet:
	go vet kb
	go vet kb/assets
	go vet kb/markdown
	go vet kb/web
