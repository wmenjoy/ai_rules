#!/bin/bash

# 运行所有测试
go test -v ./...

# 运行基准测试
go test -bench=. ./... 