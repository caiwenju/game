#!/usr/bin/env bash

cd ../

# 打 mac 包
CGO_ENABLED=0
GOOS=darwin
GOARCH=amd64
go build -o build/package/mac/mota

# 打 linux 包
CGO_ENABLED=0
GOOS=linux
GOARCH=amd64
go build -o build/package/linux/mota

## 打 windows 包
CGO_ENABLED=0
GOOS=windows
GOARCH=amd64
go build -o build/package/windows/mota.exe

