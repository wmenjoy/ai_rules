package models

// Language 代码语言类型
type Language string

const (
	Go   Language = "Go"
	Java Language = "Java"
)

// Issue 代码问题
type Issue struct {
	Line        int    `json:"line"`
	Column      int    `json:"column"`
	Message     string `json:"message"`
	Severity    string `json:"severity"`
	RuleID      string `json:"ruleId"`
	Suggestion  string `json:"suggestion,omitempty"`
	CodeSnippet string `json:"codeSnippet,omitempty"`
}

// AIDetectionResult AI检测结果
type AIDetectionResult struct {
	Score      float64  `json:"score"`
	Indicators []string `json:"indicators"`
}

// QualityMetrics 代码质量指标
type QualityMetrics struct {
	FilePath             string   `json:"filePath"`
	Language             Language `json:"language"`
	Issues               []Issue  `json:"issues"`
	CyclomaticComplexity int      `json:"cyclomaticComplexity"`
	CommentRatio         float64  `json:"commentRatio"`
	LongFunctions        int      `json:"longFunctions"`
	DeepNesting          int      `json:"deepNesting"`
	DuplicateLines       int      `json:"duplicateLines"`
	AIGeneratedScore     float64  `json:"aiGeneratedScore"`
	AIIndicators         []string `json:"aiIndicators"`
	Score                float64  `json:"score"`
	FunctionCount        int      `json:"functionCount"`
}
