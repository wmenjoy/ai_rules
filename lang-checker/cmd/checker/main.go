package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/liujinliang/lang-checker/internal/analyzer"
	"github.com/liujinliang/lang-checker/internal/models"
	"github.com/liujinliang/lang-checker/pkg/reporter"
)

var (
	path        string
	outputFile  string
	showVersion bool
	version     = "v1.0.0"
)

func init() {
	flag.StringVar(&path, "path", "", "要分析的文件或目录路径")
	flag.StringVar(&outputFile, "output", "", "报告输出文件路径 (可选)")
	flag.BoolVar(&showVersion, "version", false, "显示版本信息")
	flag.Usage = usage
}

func usage() {
	fmt.Println("AI代码质量检测工具", version)
	fmt.Println("支持语言: Go, Java")
	fmt.Println("功能: 代码质量分析 + AI生成检测")
	fmt.Println()
	fmt.Println("使用方法:")
	fmt.Printf("  %s [选项] -path <文件路径或目录路径>\n", os.Args[0])
	fmt.Println()
	fmt.Println("选项:")
	flag.PrintDefaults()
	fmt.Println()
	fmt.Println("示例:")
	fmt.Printf("  %s -path main.go\n", os.Args[0])
	fmt.Printf("  %s -path ./src\n", os.Args[0])
	fmt.Printf("  %s -path /path/to/java/project -output report.txt\n", os.Args[0])
}

func main() {
	flag.Parse()

	if showVersion {
		fmt.Printf("AI代码质量检测工具 %s\n", version)
		return
	}

	if path == "" {
		flag.Usage()
		os.Exit(1)
	}

	// 规范化路径
	path = filepath.Clean(path)

	// 创建分析器
	codeAnalyzer := analyzer.NewCodeAnalyzer()

	var metrics []*models.QualityMetrics
	var err error

	// 检查路径是否存在
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Printf("❌ 错误: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("🔍 正在分析: %s\n", path)

	// 根据路径类型进行分析
	if fileInfo.IsDir() {
		metrics, err = codeAnalyzer.AnalyzeDirectory(path)
	} else {
		metric, err := codeAnalyzer.AnalyzeFile(path)
		if err == nil {
			metrics = []*models.QualityMetrics{metric}
		}
	}

	if err != nil {
		fmt.Printf("❌ 分析失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✅ 分析完成，共处理 %d 个文件\n\n", len(metrics))

	// 生成报告
	if outputFile != "" {
		// 将标准输出重定向到文件
		f, err := os.Create(outputFile)
		if err != nil {
			fmt.Printf("❌ 创建报告文件失败: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
		oldStdout := os.Stdout
		os.Stdout = f
		reporter.GenerateReport(metrics)
		os.Stdout = oldStdout
		fmt.Printf("报告已保存到: %s\n", outputFile)
	} else {
		reporter.GenerateReport(metrics)
	}
}
