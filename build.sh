#/usr/bin/env bash

env GOOS=linux GOARCH=amd64 go build -o YahooIntradayDownloader_linux YahooIntradayDownloader.go
env GOOS=darwin GOARCH=amd64 go build -o YahooIntradayDownloader_mac YahooIntradayDownloader.go
env GOOS=windows GOARCH=amd64 go build -o YahooIntradayDownloader_windows.exe YahooIntradayDownloader.go