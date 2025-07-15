package main

import (
	"fmt"
	"os"

	"codeanalyzer/internal/analyzer"
	"codeanalyzer/pkg/reporter"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	path := os.Args[1]
	codeAnalyzer := analyzer.NewCodeAnalyzer()

	var metrics []*analyzer.QualityMetrics
	var err error

	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Printf("❌ 错误: %v\n", err)
		os.Exit(1)
	}

	if fileInfo.IsDir() {
		fmt.Printf("🔍 正在分析目录: %s\n", path)
		metrics, err = codeAnalyzer.AnalyzeDirectory(path)
	} else {
		fmt.Printf("🔍 正在分析文件: %s\n", path)
		metric, err := codeAnalyzer.AnalyzeFile(path)
		if err == nil {
			metrics = []*analyzer.QualityMetrics{metric}
		}
	}

	if err != nil {
		fmt.Printf("❌ 分析失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✅ 分析完成，共处理 %d 个文件\n\n", len(metrics))
	reporter.GenerateReport(metrics)
}

func printUsage() {
	fmt.Println("AI代码质量检测工具 v2.0")
	fmt.Println("支持语言: Go, Java")
	fmt.Println("功能: 代码质量分析 + AI生成检测")
	fmt.Println()
	fmt.Println("使用方法:")
	fmt.Println("  go run main.go <文件路径或目录路径>")
	fmt.Println()
	fmt.Println("示例:")
	fmt.Println("  go run main.go main.go")
	fmt.Println("  go run main.go ./src")
	fmt.Println("  go run main.go /path/to/java/project")
}
