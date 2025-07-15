package analyzer

import (
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/liujinliang/lang-checker/internal/models"
	"github.com/liujinliang/lang-checker/internal/rules"
)

// GoAnalyzer Go代码分析器
type GoAnalyzer struct {
	fileSet *token.FileSet
	rules   []rules.GoRule
}

// GoRule Go语言规则接口
type GoRule interface {
	Check(node ast.Node, fset *token.FileSet) []models.Issue
	Name() string
}

// NewGoAnalyzer 创建新的Go分析器
func NewGoAnalyzer() *GoAnalyzer {
	return &GoAnalyzer{
		fileSet: token.NewFileSet(),
		rules: []rules.GoRule{
			&rules.FunctionLengthRule{},
			&rules.CyclomaticComplexityRule{},
			&rules.NamingConventionRule{},
		},
	}
}

// Analyze 分析Go代码
func (ga *GoAnalyzer) Analyze(content string, filePath string) (*models.QualityMetrics, error) {
	node, err := parser.ParseFile(ga.fileSet, filePath, content, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	metrics := &models.QualityMetrics{
		FilePath: filePath,
		Language: models.Go,
	}

	// 基础指标计算
	metrics.FunctionCount = countFunctions(node)
	metrics.CyclomaticComplexity = calculateTotalComplexity(node)
	metrics.LongFunctions = countLongFunctions(node, ga.fileSet)
	metrics.DeepNesting = detectDeepNesting(node)

	// 应用规则检查
	for _, rule := range ga.rules {
		issues := rule.Check(node, ga.fileSet)
		metrics.Issues = append(metrics.Issues, issues...)
	}

	return metrics, nil
}

// 辅助函数
func countFunctions(node ast.Node) int {
	count := 0
	ast.Inspect(node, func(n ast.Node) bool {
		if _, ok := n.(*ast.FuncDecl); ok {
			count++
		}
		return true
	})
	return count
}

func calculateTotalComplexity(node ast.Node) int {
	total := 0
	ast.Inspect(node, func(n ast.Node) bool {
		switch n.(type) {
		case *ast.IfStmt, *ast.ForStmt, *ast.RangeStmt, *ast.CaseClause, *ast.CommClause:
			total++
		}
		return true
	})
	return total
}

func countLongFunctions(node ast.Node, fset *token.FileSet) int {
	count := 0
	ast.Inspect(node, func(n ast.Node) bool {
		if fn, ok := n.(*ast.FuncDecl); ok {
			start := fset.Position(fn.Pos())
			end := fset.Position(fn.End())
			if end.Line-start.Line > 50 {
				count++
			}
		}
		return true
	})
	return count
}

func detectDeepNesting(node ast.Node) int {
	maxDepth := 0
	currentDepth := 0

	ast.Inspect(node, func(n ast.Node) bool {
		switch n.(type) {
		case *ast.IfStmt, *ast.ForStmt, *ast.RangeStmt, *ast.SelectStmt, *ast.SwitchStmt:
			currentDepth++
			if currentDepth > maxDepth {
				maxDepth = currentDepth
			}
		}
		return true
	})

	return maxDepth
}
