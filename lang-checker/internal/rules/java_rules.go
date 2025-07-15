package rules

import (
	"regexp"
	"strings"

	"github.com/liujinliang/lang-checker/internal/models"
)

// JavaFunctionLengthRule Java函数长度规则
type JavaFunctionLengthRule struct{}

func (r *JavaFunctionLengthRule) CheckJava(content string) []models.Issue {
	var issues []models.Issue
	lines := strings.Split(content, "\n")
	methodPattern := regexp.MustCompile(`(?m)^\s*(?:public|private|protected)?\s*(?:static)?\s*(?:\w+\s+)*\w+\s*\([^)]*\)`)

	for i, line := range lines {
		if methodPattern.MatchString(line) {
			braceCount := 0
			methodStart := i + 1
			methodEnd := i + 1

			for j := i; j < len(lines); j++ {
				braceCount += strings.Count(lines[j], "{")
				braceCount -= strings.Count(lines[j], "}")
				if braceCount == 0 && j > i {
					methodEnd = j + 1
					break
				}
			}

			if methodEnd-methodStart > 50 {
				issues = append(issues, models.Issue{
					Line:     methodStart,
					Message:  "方法过长，建议拆分",
					Severity: "warning",
				})
			}
		}
	}
	return issues
}

func (r *JavaFunctionLengthRule) Name() string {
	return "JavaFunctionLength"
}

// JavaNamingConventionRule Java命名规范规则
type JavaNamingConventionRule struct{}

func (r *JavaNamingConventionRule) CheckJava(content string) []models.Issue {
	var issues []models.Issue
	lines := strings.Split(content, "\n")
	methodPattern := regexp.MustCompile(`(?m)^\s*(?:public|private|protected)?\s*(?:static)?\s*(?:\w+\s+)*(\w+)\s*\([^)]*\)`)

	for i, line := range lines {
		matches := methodPattern.FindStringSubmatch(line)
		if len(matches) > 1 {
			methodName := matches[1]
			if !isValidJavaMethodName(methodName) {
				issues = append(issues, models.Issue{
					Line:     i + 1,
					Message:  "方法命名不符合规范",
					Severity: "warning",
				})
			}
		}
	}
	return issues
}

func (r *JavaNamingConventionRule) Name() string {
	return "JavaNamingConvention"
}

// 辅助函数
func isValidJavaMethodName(name string) bool {
	// Java方法命名规范：小驼峰
	if len(name) == 0 {
		return false
	}
	firstChar := name[0]
	if firstChar >= 'A' && firstChar <= 'Z' {
		return false
	}
	return true
}
