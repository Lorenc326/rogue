build:
	GOARCH=amd64 GOOS=windows go build -o bin/web/rogue.exe web/main.go
	GOARCH=amd64 GOOS=darwin go build -o bin/web/rogue web/main.go
	GOARCH=amd64 GOOS=windows go build -o bin/cmd/rogue.exe cmd/main.go
	GOARCH=amd64 GOOS=darwin go build -o bin/cmd/rogue cmd/main.go
api:
	go run ./web
shell:
	go run ./cmd