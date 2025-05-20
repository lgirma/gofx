SHELL := /bin/bash

release_cputil:
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ./bin/cputil.exe ./licensing/cputil