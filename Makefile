build:
	go build -o aetherc.exe cmd/aetherc/main.go

clean:
	go clean ./...

test:
	go test ./...
