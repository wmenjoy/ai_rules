package reporter

import (
	"fmt"

	"github.com/liujinliang/lang-checker/internal/models"
)

// GenerateReport 生成分析报告
func GenerateReport(metrics []*models.QualityMetrics) {
	fmt.Println("代码质量分析报告")
	fmt.Println("================")
	fmt.Println()

	for _, m := range metrics {
		fmt.Printf("文件: %s\n", m.FilePath)
		fmt.Printf("语言: %s\n", m.Language)
		fmt.Printf("总体得分: %.2f\n", m.Score)
		fmt.Printf("圈复杂度: %d\n", m.CyclomaticComplexity)
		fmt.Printf("注释率: %.2f%%\n", m.CommentRatio)
		fmt.Printf("AI生成概率: %.2f%%\n", m.AIGeneratedScore)

		if len(m.AIIndicators) > 0 {
			fmt.Println("\nAI生成特征:")
			for _, indicator := range m.AIIndicators {
				fmt.Printf("- %s\n", indicator)
			}
		}

		if len(m.Issues) > 0 {
			fmt.Println("\n发现的问题:")
			for _, issue := range m.Issues {
				fmt.Printf("- 第%d行: %s (%s)\n", issue.Line, issue.Message, issue.Severity)
				if issue.Suggestion != "" {
					fmt.Printf("  建议: %s\n", issue.Suggestion)
				}
			}
		}

		fmt.Println("\n-------------------\n")
	}
}
