build:
	GOARCH=amd64 GOOS=windows go build -o bin/rogue.exe web/main.go
	GOARCH=amd64 GOOS=darwin go build -o bin/rogue web/main.go
web:
	go run ./web