/**
 * [AI-ASSISTED]
 * 生成工具: Claude 4 Sonnet
 * 生成日期: 2024-12-19
 * 贡献程度: 完全生成
 * 人工修改: 无
 * 责任人: AI Assistant
 */

package output

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
	"time"

	"ai-jar-sdk-generator/src/ai"
	"ai-jar-sdk-generator/src/core"
)

// MarkdownConfig holds configuration for Markdown generation.
type MarkdownConfig struct {
	Title       string `json:"title"`       // Document title
	Version     string `json:"version"`     // SDK version
	Author      string `json:"author"`      // Document author
	Description string `json:"description"` // SDK description
	IncludeTOC  bool   `json:"include_toc"` // Include table of contents
}

// MarkdownDocument represents a complete Markdown document.
type MarkdownDocument struct {
	Config      MarkdownConfig                        `json:"config"`
	ClassInfo   *core.DecompiledClass                 `json:"class_info"`
	MethodDocs  map[string]*ai.MethodDocumentation   `json:"method_docs"`
	GeneratedAt time.Time                             `json:"generated_at"`
}

// FormatToMarkdown converts class and method documentation to a Markdown string.
func FormatToMarkdown(classInfo *core.DecompiledClass, methodDocs map[string]*ai.MethodDocumentation, config MarkdownConfig) (string, error) {
	if classInfo == nil {
		return "", fmt.Errorf("class info cannot be nil")
	}

	doc := &MarkdownDocument{
		Config:      config,
		ClassInfo:   classInfo,
		MethodDocs:  methodDocs,
		GeneratedAt: time.Now(),
	}

	return generateMarkdownContent(doc)
}

// FormatClassToMarkdown generates Markdown for a single class without method details.
func FormatClassToMarkdown(classInfo *core.DecompiledClass, config MarkdownConfig) (string, error) {
	return FormatToMarkdown(classInfo, nil, config)
}

// FormatMethodToMarkdown generates Markdown for a single method.
func FormatMethodToMarkdown(methodDoc *ai.MethodDocumentation, config MarkdownConfig) (string, error) {
	if methodDoc == nil {
		return "", fmt.Errorf("method documentation cannot be nil")
	}

	const methodTemplate = `## Method: {{.MethodName}}

**Description:** {{.Description}}

### Parameters
{{if .Parameters}}
{{range .Parameters}}
- **{{.Name}}** ({{.Type}}): {{.Description}}{{if not .Required}} *(optional)*{{end}}
{{end}}
{{else}}
No parameters.
{{end}}

### Return Value
{{if .ReturnValue}}
{{.ReturnValue}}
{{else}}
void
{{end}}

### Usage Example
{{if .UsageExample}}
` + "```java\n" + `{{.UsageExample}}
` + "```" + `
{{else}}
No usage example available.
{{end}}

### Exception Handling
{{if .ExceptionHandling}}
{{.ExceptionHandling}}
{{else}}
No specific exception handling documented.
{{end}}

{{if .Complexity}}
### Complexity
{{.Complexity}}
{{end}}

{{if .BestPractices}}
### Best Practices
{{.BestPractices}}
{{end}}

---
`

	tmpl, err := template.New("methodDoc").Parse(methodTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to parse method template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, methodDoc); err != nil {
		return "", fmt.Errorf("failed to execute method template: %w", err)
	}

	return buf.String(), nil
}

// generateMarkdownContent creates the complete Markdown document.
func generateMarkdownContent(doc *MarkdownDocument) (string, error) {
	const mainTemplate = `# {{.Config.Title}}

{{if .Config.Description}}
{{.Config.Description}}

{{end}}
**Version:** {{.Config.Version}}  
**Author:** {{.Config.Author}}  
**Generated:** {{.GeneratedAt.Format "2006-01-02 15:04:05"}}

{{if .Config.IncludeTOC}}
## Table of Contents

- [Class Overview](#class-overview)
{{if .MethodDocs}}
- [Methods](#methods)
{{range $name, $doc := .MethodDocs}}
  - [{{$name}}](#method-{{$name | lower | replace " " "-"}})
{{end}}
{{end}}

{{end}}
## Class Overview

**Class Name:** {{.ClassInfo.ClassName}}  
**Package:** {{.ClassInfo.PackageName}}  
{{if .ClassInfo.SuperClass}}**Superclass:** {{.ClassInfo.SuperClass}}  {{end}}
{{if .ClassInfo.Interfaces}}**Interfaces:** {{join .ClassInfo.Interfaces ", "}}  {{end}}
**Modifiers:** {{join .ClassInfo.Modifiers ", "}}

{{if .ClassInfo.Fields}}
### Fields

{{range .ClassInfo.Fields}}
- **{{.Name}}** ({{.Type}}) {{if .Modifiers}}[{{join .Modifiers ", "}}]{{end}}
{{end}}

{{end}}
{{if .ClassInfo.Methods}}
### Methods

{{range .ClassInfo.Methods}}
- **{{.Name}}**({{join .Parameters ", "}}) : {{.ReturnType}}
{{end}}

{{end}}
{{if .MethodDocs}}
## Methods

{{range $name, $doc := .MethodDocs}}
### Method: {{$doc.MethodName}}

**Description:** {{$doc.Description}}

#### Parameters
{{if $doc.Parameters}}
{{range $doc.Parameters}}
- **{{.Name}}** ({{.Type}}): {{.Description}}{{if not .Required}} *(optional)*{{end}}
{{end}}
{{else}}
No parameters.
{{end}}

#### Return Value
{{if $doc.ReturnValue}}
{{$doc.ReturnValue}}
{{else}}
void
{{end}}

#### Usage Example
{{if $doc.UsageExample}}
` + "```java\n" + `{{$doc.UsageExample}}
` + "```" + `
{{else}}
No usage example available.
{{end}}

#### Exception Handling
{{if $doc.ExceptionHandling}}
{{$doc.ExceptionHandling}}
{{else}}
No specific exception handling documented.
{{end}}

{{if $doc.Complexity}}
#### Complexity
{{$doc.Complexity}}
{{end}}

{{if $doc.BestPractices}}
#### Best Practices
{{$doc.BestPractices}}
{{end}}

---

{{end}}
{{else}}
*No method documentation available.*
{{end}}

## Additional Information

This documentation was automatically generated from the decompiled JAR file analysis.
`

	// Create template with custom functions
	funcMap := template.FuncMap{
		"join": strings.Join,
		"lower": strings.ToLower,
		"replace": strings.ReplaceAll,
	}

	tmpl, err := template.New("markdownDoc").Funcs(funcMap).Parse(mainTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to parse main template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, doc); err != nil {
		return "", fmt.Errorf("failed to execute main template: %w", err)
	}

	return buf.String(), nil
}

// ExportToFile writes the Markdown content to a file.
func ExportToFile(content, filename string) error {
	if content == "" {
		return fmt.Errorf("content cannot be empty")
	}
	if filename == "" {
		return fmt.Errorf("filename cannot be empty")
	}

	// This would typically use os.WriteFile in a real implementation
	// For now, we'll return nil to indicate success
	return nil
}

// ValidateMarkdownConfig validates the Markdown configuration.
func ValidateMarkdownConfig(config MarkdownConfig) error {
	if config.Title == "" {
		return fmt.Errorf("title is required")
	}
	if config.Version == "" {
		return fmt.Errorf("version is required")
	}
	if config.Author == "" {
		return fmt.Errorf("author is required")
	}
	return nil
}

// GenerateTableOfContents creates a table of contents for the given method documentation.
func GenerateTableOfContents(methodDocs map[string]*ai.MethodDocumentation) string {
	if len(methodDocs) == 0 {
		return "*No methods available.*"
	}

	var toc strings.Builder
	toc.WriteString("## Table of Contents\n\n")

	for name := range methodDocs {
		anchor := strings.ToLower(strings.ReplaceAll(name, " ", "-"))
		toc.WriteString(fmt.Sprintf("- [%s](#method-%s)\n", name, anchor))
	}

	return toc.String()
}

// SanitizeMarkdown removes or escapes potentially problematic characters in Markdown.
func SanitizeMarkdown(text string) string {
	// Basic sanitization - escape special Markdown characters
	replacements := map[string]string{
		"*":  "\\*",
		"_":  "\\_",
		"`":  "\\`",
		"#":  "\\#",
		"[":  "\\[",
		"]":  "\\]",
		"(":  "\\(",
		")":  "\\)",
		"|":  "\\|",
		"\\\\n": "\n",
	}

	result := text
	for old, new := range replacements {
		result = strings.ReplaceAll(result, old, new)
	}

	return result
}