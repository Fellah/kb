install: format
	go vet
	go generate kb/assets
	go test
	go install -gcflags "-N -l"

format:
	go fmt
	go fmt kb/assets
	go fmt kb/web
