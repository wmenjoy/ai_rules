#!/bin/bash

# 在所有.go文件中更新导入路径
find . -type f -name "*.go" -exec sed -i '' 's|codeanalyzer|github.com/liujinliang/lang-checker|g' {} + 