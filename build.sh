#/usr/bin/env bash

env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o YahooIntradayDownloader_linux YahooIntradayDownloader.go
env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -o YahooIntradayDownloader_mac YahooIntradayDownloader.go
env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -o YahooIntradayDownloader_windows.exe YahooIntradayDownloader.go
