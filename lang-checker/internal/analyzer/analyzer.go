package analyzer

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/liujinliang/lang-checker/internal/detector"
	"github.com/liujinliang/lang-checker/internal/models"
)

// CodeAnalyzer 代码分析器
type CodeAnalyzer struct {
	goAnalyzer   *GoAnalyzer
	javaAnalyzer *JavaAnalyzer
	aiDetector   *detector.AIDetector
}

// NewCodeAnalyzer 创建新的代码分析器
func NewCodeAnalyzer() *CodeAnalyzer {
	return &CodeAnalyzer{
		goAnalyzer:   NewGoAnalyzer(),
		javaAnalyzer: NewJavaAnalyzer(),
		aiDetector:   detector.NewAIDetector(),
	}
}

// AnalyzeFile 分析单个文件
func (ca *CodeAnalyzer) AnalyzeFile(filePath string) (*models.QualityMetrics, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	contentStr := string(content)
	var metrics *models.QualityMetrics
	var analyzeErr error

	switch detectLanguage(filePath) {
	case models.Go:
		metrics, analyzeErr = ca.goAnalyzer.Analyze(contentStr, filePath)
	case models.Java:
		metrics, analyzeErr = ca.javaAnalyzer.Analyze(contentStr, filePath)
	}

	if analyzeErr != nil {
		return nil, analyzeErr
	}

	// AI检测
	aiResult := ca.aiDetector.DetectAI(contentStr)
	metrics.AIGeneratedScore = aiResult.Score
	metrics.AIIndicators = aiResult.Indicators

	// 计算质量得分
	metrics.Score = calculateQualityScore(metrics)

	return metrics, nil
}

// AnalyzeDirectory 分析目录
func (ca *CodeAnalyzer) AnalyzeDirectory(dirPath string) ([]*models.QualityMetrics, error) {
	var allMetrics []*models.QualityMetrics

	err := filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && isTargetFile(path) {
			metrics, err := ca.AnalyzeFile(path)
			if err != nil {
				return err
			}
			allMetrics = append(allMetrics, metrics)
		}
		return nil
	})

	return allMetrics, err
}

// 辅助函数
func detectLanguage(filePath string) models.Language {
	ext := strings.ToLower(filepath.Ext(filePath))
	switch ext {
	case ".go":
		return models.Go
	case ".java":
		return models.Java
	default:
		return models.Go
	}
}

func isTargetFile(path string) bool {
	return (strings.HasSuffix(path, ".go") || strings.HasSuffix(path, ".java")) &&
		!strings.Contains(path, "vendor/") && !strings.Contains(path, "target/")
}

func calculateQualityScore(metrics *models.QualityMetrics) float64 {
	score := 100.0

	// 基础扣分规则
	score -= float64(len(metrics.Issues)) * 5
	if metrics.CyclomaticComplexity > 50 {
		score -= 10
	}
	if metrics.CommentRatio < 10 {
		score -= 15
	}
	if metrics.LongFunctions > 0 {
		score -= float64(metrics.LongFunctions) * 8
	}
	if metrics.DeepNesting > 4 {
		score -= float64(metrics.DeepNesting-4) * 3
	}
	if metrics.DuplicateLines > 10 {
		score -= float64(metrics.DuplicateLines) * 0.5
	}

	// AI生成代码扣分
	if metrics.AIGeneratedScore > 70 {
		score -= 20
	} else if metrics.AIGeneratedScore > 40 {
		score -= 10
	}

	if score < 0 {
		score = 0
	}
	return score
}
