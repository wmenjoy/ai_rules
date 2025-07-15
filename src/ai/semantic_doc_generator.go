/**
 * [AI-ASSISTED]
 * 生成工具: Claude 4 Sonnet
 * 生成日期: 2024-12-19
 * 贡献程度: 完全生成
 * 人工修改: 无
 * 责任人: AI Assistant
 */

package ai

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"ai-jar-sdk-generator/src/core"
)

// MethodDocumentation holds detailed AI-generated docs for a method.
type MethodDocumentation struct {
	MethodName        string      `json:"method_name"`
	Description       string      `json:"description"`
	Parameters        []ParamDoc  `json:"parameters"`
	ReturnValue       string      `json:"return_value"`
	UsageExample      string      `json:"usage_example"`
	ExceptionHandling string      `json:"exception_handling"`
	Complexity        string      `json:"complexity"`
	BestPractices     string      `json:"best_practices"`
	GeneratedAt       time.Time   `json:"generated_at"`
}

// ParamDoc represents documentation for a method parameter.
type ParamDoc struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Required    bool   `json:"required"`
	DefaultValue string `json:"default_value,omitempty"`
}

// SemanticDocumentationResult holds the complete semantic documentation for a class.
type SemanticDocumentationResult struct {
	ClassName         string                 `json:"class_name"`
	PackageName       string                 `json:"package_name"`
	MethodDocs        []MethodDocumentation  `json:"method_docs"`
	GeneratedAt       time.Time              `json:"generated_at"`
	TotalMethods      int                    `json:"total_methods"`
	DocumentedMethods int                    `json:"documented_methods"`
}

// GenerateSemanticMethodDoc generates detailed documentation for a single method.
func GenerateSemanticMethodDoc(methodInfo *core.DecompiledMethod, classContext *core.DecompiledClass, dependencyResult *core.DependencyAnalysisResult, config AIModelConfig) (*MethodDocumentation, error) {
	if methodInfo == nil {
		return nil, fmt.Errorf("method info cannot be nil")
	}
	if classContext == nil {
		return nil, fmt.Errorf("class context cannot be nil")
	}

	// Validate configuration
	if err := ValidateConfig(config); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	// Build specialized prompt for method documentation
	prompt := buildMethodPrompt(methodInfo, classContext, dependencyResult)

	// Call AI model
	aiResponse, err := callAIModel(config, prompt)
	if err != nil {
		return nil, fmt.Errorf("failed to call AI model: %w", err)
	}

	// Parse AI response into structured documentation
	methodDoc := parseMethodAIResponse(aiResponse, methodInfo.Name)
	methodDoc.GeneratedAt = time.Now()

	return methodDoc, nil
}

// GenerateSemanticClassDoc generates detailed documentation for all methods in a class.
func GenerateSemanticClassDoc(classInfo *core.DecompiledClass, dependencyResult *core.DependencyAnalysisResult, config AIModelConfig) (*SemanticDocumentationResult, error) {
	if classInfo == nil {
		return nil, fmt.Errorf("class info cannot be nil")
	}

	result := &SemanticDocumentationResult{
		ClassName:         classInfo.ClassName,
		PackageName:       classInfo.PackageName,
		MethodDocs:        make([]MethodDocumentation, 0, len(classInfo.Methods)),
		GeneratedAt:       time.Now(),
		TotalMethods:      len(classInfo.Methods),
		DocumentedMethods: 0,
	}

	// Generate documentation for each method
	for _, method := range classInfo.Methods {
		methodDoc, err := GenerateSemanticMethodDoc(&method, classInfo, dependencyResult, config)
		if err != nil {
			// Log error but continue with other methods
			fmt.Printf("Warning: Failed to generate docs for method %s: %v\n", method.Name, err)
			continue
		}
		result.MethodDocs = append(result.MethodDocs, *methodDoc)
		result.DocumentedMethods++
	}

	return result, nil
}

// buildMethodPrompt constructs a specialized prompt for method documentation.
func buildMethodPrompt(methodInfo *core.DecompiledMethod, classContext *core.DecompiledClass, dependencyResult *core.DependencyAnalysisResult) string {
	var prompt strings.Builder

	prompt.WriteString("Generate comprehensive documentation for the following Java method:\n\n")

	// Method signature and context
	prompt.WriteString(fmt.Sprintf("**Class:** %s.%s\n", classContext.PackageName, classContext.ClassName))
	prompt.WriteString(fmt.Sprintf("**Method:** %s\n", methodInfo.Name))
	prompt.WriteString(fmt.Sprintf("**Return Type:** %s\n", methodInfo.ReturnType))
	prompt.WriteString(fmt.Sprintf("**Parameters:** [%s]\n", strings.Join(methodInfo.Parameters, ", ")))
	prompt.WriteString(fmt.Sprintf("**Modifiers:** [%s]\n\n", strings.Join(methodInfo.Modifiers, ", ")))

	// Add JavaDoc if available
	if methodInfo.JavaDoc != "" {
		prompt.WriteString(fmt.Sprintf("**Existing JavaDoc:** %s\n\n", methodInfo.JavaDoc))
	}

	// Add method source code if available
	if methodInfo.SourceCode != "" {
		prompt.WriteString("**Method Source Code:**\n```java\n")
		prompt.WriteString(methodInfo.SourceCode)
		prompt.WriteString("\n```\n\n")
	}

	// Add class context
	prompt.WriteString("**Class Context:**\n")
	prompt.WriteString(fmt.Sprintf("- Super Class: %s\n", classContext.SuperClass))
	prompt.WriteString(fmt.Sprintf("- Interfaces: [%s]\n", strings.Join(classContext.Interfaces, ", ")))
	prompt.WriteString(fmt.Sprintf("- Fields: %d\n", len(classContext.Fields)))
	prompt.WriteString(fmt.Sprintf("- Methods: %d\n\n", len(classContext.Methods)))

	// Add dependency information if available
	if dependencyResult != nil {
		prompt.WriteString("**Dependencies:**\n")
		for _, dep := range dependencyResult.ClassDependencies {
			if dep.FromClass == classContext.ClassName {
				prompt.WriteString(fmt.Sprintf("- %s %s\n", dep.DepType, dep.ToClass))
			}
		}
		prompt.WriteString("\n")
	}

	// Specific documentation requirements
	prompt.WriteString("Please provide the following information in a structured format:\n\n")
	prompt.WriteString("## Method Description\n")
	prompt.WriteString("Provide a clear, concise description of what this method does, its purpose, and its role in the class.\n\n")

	prompt.WriteString("## Parameters\n")
	prompt.WriteString("For each parameter, describe its purpose, expected type, constraints, and whether it's required.\n\n")

	prompt.WriteString("## Return Value\n")
	prompt.WriteString("Describe what the method returns, including the type and meaning of the return value.\n\n")

	prompt.WriteString("## Usage Example\n")
	prompt.WriteString("Provide a practical code example showing how to use this method, including setup and context.\n\n")

	prompt.WriteString("## Exception Handling\n")
	prompt.WriteString("List potential exceptions that might be thrown and how to handle them properly.\n\n")

	prompt.WriteString("## Complexity\n")
	prompt.WriteString("Describe the time/space complexity and performance characteristics.\n\n")

	prompt.WriteString("## Best Practices\n")
	prompt.WriteString("Provide recommendations for best practices when using this method.\n\n")

	prompt.WriteString("Please format your response with clear section headers and provide detailed, actionable information.")

	return prompt.String()
}

// parseMethodAIResponse parses the AI response into a MethodDocumentation struct.
func parseMethodAIResponse(response, methodName string) *MethodDocumentation {
	methodDoc := &MethodDocumentation{
		MethodName:        methodName,
		Description:       response, // Full response as fallback
		Parameters:        make([]ParamDoc, 0),
		ReturnValue:       "",
		UsageExample:      "",
		ExceptionHandling: "",
		Complexity:        "",
		BestPractices:     "",
	}

	// Extract specific sections from the response
	methodDoc.Description = extractMethodSection(response, "description", "method description")
	methodDoc.ReturnValue = extractMethodSection(response, "return", "return value")
	methodDoc.UsageExample = extractMethodSection(response, "example", "usage")
	methodDoc.ExceptionHandling = extractMethodSection(response, "exception", "error")
	methodDoc.Complexity = extractMethodSection(response, "complexity", "performance")
	methodDoc.BestPractices = extractMethodSection(response, "best practices", "recommendation")

	// Parse parameters section
	parametersSection := extractMethodSection(response, "parameters", "parameter")
	if parametersSection != "" {
		methodDoc.Parameters = parseParametersFromText(parametersSection)
	}

	return methodDoc
}

// extractMethodSection extracts a specific section from the AI response.
func extractMethodSection(response, keyword1, keyword2 string) string {
	lines := strings.Split(response, "\n")
	var section strings.Builder
	inSection := false

	for _, line := range lines {
		lowerLine := strings.ToLower(strings.TrimSpace(line))

		// Check if this line starts a relevant section (header line containing keywords)
		if (strings.Contains(lowerLine, strings.ToLower(keyword1)) || strings.Contains(lowerLine, strings.ToLower(keyword2))) &&
			(strings.HasPrefix(lowerLine, "#") || strings.HasPrefix(lowerLine, "##")) {
			inSection = true
			continue
		}

		// Check if we've moved to a different section (starts with ## or #)
		if inSection && (strings.HasPrefix(strings.TrimSpace(line), "##") || strings.HasPrefix(strings.TrimSpace(line), "#")) {
			break
		}

		// Add line to section if we're in the right section
		if inSection {
			trimmedLine := strings.TrimSpace(line)
			if trimmedLine != "" {
				if section.Len() > 0 {
					section.WriteString("\n")
				}
				section.WriteString(trimmedLine)
			}
		}
	}

	return strings.TrimSpace(section.String())
}

// parseParametersFromText parses parameter information from text.
func parseParametersFromText(text string) []ParamDoc {
	var params []ParamDoc
	lines := strings.Split(text, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Try to parse parameter information
		// Expected formats:
		// - paramName (Type): Description
		// - paramName: Description
		// - Type paramName: Description
		if strings.Contains(line, ":") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				paramInfo := strings.TrimSpace(parts[0])
				description := strings.TrimSpace(parts[1])

				// Remove bullet points or dashes
				paramInfo = strings.TrimPrefix(paramInfo, "-")
				paramInfo = strings.TrimPrefix(paramInfo, "*")
				paramInfo = strings.TrimSpace(paramInfo)

				param := ParamDoc{
					Description: description,
					Required:    true, // Default to required
				}

				// Try to extract name and type
				if strings.Contains(paramInfo, "(") && strings.Contains(paramInfo, ")") {
					// Format: paramName (Type)
					start := strings.Index(paramInfo, "(")
					end := strings.Index(paramInfo, ")")
					if start < end {
						param.Name = strings.TrimSpace(paramInfo[:start])
						param.Type = strings.TrimSpace(paramInfo[start+1 : end])
					}
				} else {
					// Try to split by space to get type and name
					words := strings.Fields(paramInfo)
					if len(words) >= 2 {
						param.Type = words[0]
						param.Name = words[1]
					} else if len(words) == 1 {
						param.Name = words[0]
					}
				}

				// Check if parameter is optional
				if strings.Contains(strings.ToLower(description), "optional") {
					param.Required = false
				}

				if param.Name != "" {
					params = append(params, param)
				}
			}
		}
	}

	return params
}

// ExportSemanticDocToJSON exports semantic documentation to JSON format.
func ExportSemanticDocToJSON(result *SemanticDocumentationResult) ([]byte, error) {
	if result == nil {
		return nil, fmt.Errorf("documentation result cannot be nil")
	}

	return json.MarshalIndent(result, "", "  ")
}

// ValidateMethodDocumentation validates the completeness of method documentation.
func ValidateMethodDocumentation(doc *MethodDocumentation) []string {
	var issues []string

	if doc == nil {
		return []string{"documentation is nil"}
	}

	if doc.Description == "" {
		issues = append(issues, "missing method description")
	}

	if doc.ReturnValue == "" {
		issues = append(issues, "missing return value description")
	}

	if doc.UsageExample == "" {
		issues = append(issues, "missing usage example")
	}

	if len(doc.Parameters) == 0 {
		issues = append(issues, "no parameter documentation found")
	} else {
		for i, param := range doc.Parameters {
			if param.Name == "" {
				issues = append(issues, fmt.Sprintf("parameter %d missing name", i))
			}
			if param.Description == "" {
				issues = append(issues, fmt.Sprintf("parameter %s missing description", param.Name))
			}
		}
	}

	return issues
}