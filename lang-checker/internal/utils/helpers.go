package utils

import (
	"math"
	"regexp"
	"strings"
	"unicode"
)

// IsCapitalized 检查字符串是否以大写字母开头
func IsCapitalized(name string) bool {
	if len(name) == 0 {
		return false
	}
	return unicode.IsUpper(rune(name[0]))
}

// IsPascalCase 检查是否符合PascalCase命名规范
func IsPascalCase(name string) bool {
	if len(name) == 0 {
		return false
	}
	return unicode.IsUpper(rune(name[0])) && !strings.Contains(name, "_")
}

// IsCamelCase 检查是否符合camelCase命名规范
func IsCamelCase(name string) bool {
	if len(name) == 0 {
		return false
	}
	return unicode.IsLower(rune(name[0])) && !strings.Contains(name, "_")
}

// ExtractJavaMethodName 从方法声明中提取方法名
func ExtractJavaMethodName(line string) string {
	methodPattern := regexp.MustCompile(`(?:public|private|protected)?\s*(?:static)?\s*(?:\w+\s+)*(\w+)\s*\(`)
	matches := methodPattern.FindStringSubmatch(line)
	if len(matches) > 1 {
		return matches[1]
	}
	return "unknown"
}

// IsConstructor 检查方法是否是构造函数
func IsConstructor(name, content string) bool {
	classPattern := regexp.MustCompile(`class\s+` + name)
	return classPattern.MatchString(content)
}

// CalculateVariance 计算数字序列的方差
func CalculateVariance(numbers []int) float64 {
	if len(numbers) == 0 {
		return 0
	}

	sum := 0
	for _, n := range numbers {
		sum += n
	}
	mean := float64(sum) / float64(len(numbers))

	variance := 0.0
	for _, n := range numbers {
		variance += math.Pow(float64(n)-mean, 2)
	}

	return variance / float64(len(numbers))
}
