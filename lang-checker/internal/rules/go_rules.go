package rules

import (
	"go/ast"
	"go/token"

	"github.com/liujinliang/lang-checker/internal/models"
)

// GoRule Go语言规则接口
type GoRule interface {
	Check(node ast.Node, fset *token.FileSet) []models.Issue
	Name() string
}

// FunctionLengthRule 函数长度规则
type FunctionLengthRule struct{}

func (r *FunctionLengthRule) Check(node ast.Node, fset *token.FileSet) []models.Issue {
	var issues []models.Issue
	ast.Inspect(node, func(n ast.Node) bool {
		if fn, ok := n.(*ast.FuncDecl); ok {
			start := fset.Position(fn.Pos())
			end := fset.Position(fn.End())
			if end.Line-start.Line > 50 {
				issues = append(issues, models.Issue{
					Line:     start.Line,
					Message:  "函数过长，建议拆分",
					Severity: "warning",
				})
			}
		}
		return true
	})
	return issues
}

func (r *FunctionLengthRule) Name() string {
	return "FunctionLength"
}

// CyclomaticComplexityRule 圈复杂度规则
type CyclomaticComplexityRule struct{}

func (r *CyclomaticComplexityRule) Check(node ast.Node, fset *token.FileSet) []models.Issue {
	var issues []models.Issue
	ast.Inspect(node, func(n ast.Node) bool {
		if fn, ok := n.(*ast.FuncDecl); ok {
			complexity := calculateComplexity(fn)
			if complexity > 10 {
				start := fset.Position(fn.Pos())
				issues = append(issues, models.Issue{
					Line:     start.Line,
					Message:  "函数圈复杂度过高，建议重构",
					Severity: "warning",
				})
			}
		}
		return true
	})
	return issues
}

func (r *CyclomaticComplexityRule) Name() string {
	return "CyclomaticComplexity"
}

// NamingConventionRule 命名规范规则
type NamingConventionRule struct{}

func (r *NamingConventionRule) Check(node ast.Node, fset *token.FileSet) []models.Issue {
	var issues []models.Issue
	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			if !checkFuncName(x.Name.Name) {
				pos := fset.Position(x.Pos())
				issues = append(issues, models.Issue{
					Line:     pos.Line,
					Message:  "函数命名不符合规范",
					Severity: "warning",
				})
			}
		}
		return true
	})
	return issues
}

func (r *NamingConventionRule) Name() string {
	return "NamingConvention"
}

// 辅助函数
func calculateComplexity(fn *ast.FuncDecl) int {
	complexity := 1
	ast.Inspect(fn, func(n ast.Node) bool {
		switch n.(type) {
		case *ast.IfStmt, *ast.ForStmt, *ast.RangeStmt, *ast.CaseClause, *ast.CommClause:
			complexity++
		}
		return true
	})
	return complexity
}

func checkFuncName(name string) bool {
	// 检查函数名是否符合驼峰命名规范
	if len(name) == 0 || !ast.IsExported(name) {
		return false
	}
	return true
}
