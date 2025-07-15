package analyzer

import (
	"regexp"
	"strings"

	"github.com/liujinliang/lang-checker/internal/models"
	"github.com/liujinliang/lang-checker/internal/rules"
)

// JavaAnalyzer Java代码分析器
type JavaAnalyzer struct {
	rules []JavaRule
}

// JavaRule Java语言规则接口
type JavaRule interface {
	CheckJava(content string) []models.Issue
	Name() string
}

// NewJavaAnalyzer 创建新的Java分析器
func NewJavaAnalyzer() *JavaAnalyzer {
	return &JavaAnalyzer{
		rules: []JavaRule{
			&rules.JavaFunctionLengthRule{},
			&rules.JavaNamingConventionRule{},
		},
	}
}

// Analyze 分析Java代码
func (ja *JavaAnalyzer) Analyze(content string, filePath string) (*models.QualityMetrics, error) {
	metrics := &models.QualityMetrics{
		FilePath: filePath,
		Language: models.Java,
	}

	// 基础指标计算
	metrics.FunctionCount = countJavaFunctions(content)
	metrics.CyclomaticComplexity = calculateJavaCyclomaticComplexity(content)
	metrics.LongFunctions = countJavaLongFunctions(content)
	metrics.DeepNesting = detectJavaDeepNesting(content)

	// 应用规则检查
	for _, rule := range ja.rules {
		issues := rule.CheckJava(content)
		metrics.Issues = append(metrics.Issues, issues...)
	}

	return metrics, nil
}

// 辅助函数
func countJavaFunctions(content string) int {
	methodPattern := regexp.MustCompile(`(?m)^\s*(?:public|private|protected)?\s*(?:static)?\s*(?:\w+\s+)*\w+\s*\([^)]*\)`)
	return len(methodPattern.FindAllString(content, -1))
}

func calculateJavaCyclomaticComplexity(content string) int {
	complexity := 1
	patterns := []string{
		`\bif\s*\(`, `\belse\b`, `\bwhile\s*\(`, `\bfor\s*\(`,
		`\bswitch\s*\(`, `\bcase\s+`, `\bcatch\s*\(`, `\b\?\s*`,
	}

	for _, pattern := range patterns {
		matches := regexp.MustCompile(pattern).FindAllString(content, -1)
		complexity += len(matches)
	}

	return complexity
}

func countJavaLongFunctions(content string) int {
	lines := strings.Split(content, "\n")
	count := 0
	methodPattern := regexp.MustCompile(`(?m)^\s*(?:public|private|protected)?\s*(?:static)?\s*(?:\w+\s+)*\w+\s*\([^)]*\)`)

	for i, line := range lines {
		if methodPattern.MatchString(line) {
			braceCount := 0
			methodStart := i
			methodEnd := i

			for j := i; j < len(lines); j++ {
				braceCount += strings.Count(lines[j], "{")
				braceCount -= strings.Count(lines[j], "}")
				if braceCount == 0 && j > i {
					methodEnd = j
					break
				}
			}

			if methodEnd-methodStart > 50 {
				count++
			}
		}
	}
	return count
}

func detectJavaDeepNesting(content string) int {
	lines := strings.Split(content, "\n")
	maxDepth := 0
	currentDepth := 0

	for _, line := range lines {
		// 检测控制结构开始
		if regexp.MustCompile(`\b(if|for|while|switch|try)\s*\(`).MatchString(line) {
			currentDepth++
			if currentDepth > maxDepth {
				maxDepth = currentDepth
			}
		}

		// 检测块结束
		openBraces := strings.Count(line, "{")
		closeBraces := strings.Count(line, "}")
		currentDepth += openBraces - closeBraces

		if currentDepth < 0 {
			currentDepth = 0
		}
	}

	return maxDepth
}
