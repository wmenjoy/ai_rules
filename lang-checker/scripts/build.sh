#!/bin/bash

# 设置版本号
VERSION=$(git describe --tags --always --dirty)
BUILD_TIME=$(date -u '+%Y-%m-%d_%H:%M:%S')

# 构建命令
go build -ldflags "-X main.version=$VERSION -X main.buildTime=$BUILD_TIME" -o bin/checker cmd/checker/main.go 