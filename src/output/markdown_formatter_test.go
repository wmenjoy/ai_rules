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
	"strings"
	"testing"
	"time"

	"ai-jar-sdk-generator/src/ai"
	"ai-jar-sdk-generator/src/core"
)

func TestFormatToMarkdown(t *testing.T) {
	// Create test data
	classInfo := &core.DecompiledClass{
		ClassName:        "TestClass",
		PackageName: "com.example",
		SuperClass:  "Object",
		Interfaces:  []string{"Serializable", "Comparable"},
		Modifiers:   []string{"public"},
		Fields: []core.DecompiledField{
			{Name: "id", Type: "int", Modifiers: []string{}},
			{Name: "name", Type: "String", Modifiers: []string{}},
		},
		Methods: []core.DecompiledMethod{
			{Name: "TestClass", Parameters: []string{"int id", "String name"}},
		},
	}

	methodDocs := map[string]*ai.MethodDocumentation{
		"getId": {
			MethodName:        "getId",
			Description: "Returns the ID of the object",
			Parameters: []ai.ParamDoc{
				// No parameters for getter
			},
			ReturnValue:       "int - The ID value",
			UsageExample:      "int id = obj.getId();",
			ExceptionHandling: "No exceptions thrown",
			Complexity:        "O(1)",
			BestPractices:     "Use this method to access the ID safely",
		},
		"setName": {
			MethodName:        "setName",
			Description: "Sets the name of the object",
			Parameters: []ai.ParamDoc{
				{Name: "name", Type: "String", Description: "The new name", Required: true},
			},
			ReturnValue:       "void",
			UsageExample:      "obj.setName(\"New Name\");",
			ExceptionHandling: "Throws IllegalArgumentException if name is null",
			Complexity:        "O(1)",
			BestPractices:     "Validate input before setting",
		},
	}

	config := MarkdownConfig{
		Title:       "Test SDK Documentation",
		Version:     "1.0.0",
		Author:      "Test Author",
		Description: "Test SDK for demonstration",
		IncludeTOC:  true,
	}

	result, err := FormatToMarkdown(classInfo, methodDocs, config)
	if err != nil {
		t.Fatalf("FormatToMarkdown failed: %v", err)
	}

	// Verify the result contains expected content
	expectedContent := []string{
		"# Test SDK Documentation",
		"**Class Name:** TestClass",
		"**Package:** com.example",
		"## Methods",
		"### Method: getId",
		"### Method: setName",
		"Returns the ID of the object",
		"Sets the name of the object",
	}

	for _, expected := range expectedContent {
		if !strings.Contains(result, expected) {
			t.Errorf("Expected content not found: %s", expected)
		}
	}
}

func TestFormatToMarkdownWithNilClass(t *testing.T) {
	config := MarkdownConfig{
		Title:   "Test",
		Version: "1.0.0",
		Author:  "Test Author",
	}

	_, err := FormatToMarkdown(nil, nil, config)
	if err == nil {
		t.Error("Expected error for nil class info")
	}
}

func TestFormatClassToMarkdown(t *testing.T) {
	classInfo := &core.DecompiledClass{
		ClassName:    "SimpleClass",
		PackageName: "com.example",
	}

	config := MarkdownConfig{
		Title:       "Simple Class Documentation",
		Version:     "1.0.0",
		Author:      "Test Author",
		IncludeTOC:  false,
	}

	result, err := FormatClassToMarkdown(classInfo, config)
	if err != nil {
		t.Fatalf("FormatClassToMarkdown failed: %v", err)
	}

	if !strings.Contains(result, "SimpleClass") {
		t.Error("Expected class name not found in result")
	}
	if !strings.Contains(result, "com.example") {
		t.Error("Expected package name not found in result")
	}
}

func TestFormatMethodToMarkdown(t *testing.T) {
	methodDoc := &ai.MethodDocumentation{
		MethodName:        "testMethod",
		Description: "A test method",
		Parameters: []ai.ParamDoc{
			{Name: "param1", Type: "String", Description: "First parameter", Required: true},
			{Name: "param2", Type: "int", Description: "Second parameter", Required: false},
		},
		ReturnValue:       "boolean - Success status",
		UsageExample:      "boolean result = obj.testMethod(\"test\", 42);",
		ExceptionHandling: "May throw RuntimeException",
		Complexity:        "O(n)",
		BestPractices:     "Always check return value",
	}

	config := MarkdownConfig{}

	result, err := FormatMethodToMarkdown(methodDoc, config)
	if err != nil {
		t.Fatalf("FormatMethodToMarkdown failed: %v", err)
	}

	expectedContent := []string{
		"## Method: testMethod",
		"A test method",
		"### Parameters",
		"**param1** (String): First parameter",
		"**param2** (int): Second parameter *(optional)*",
		"### Return Value",
		"boolean - Success status",
		"### Usage Example",
		"```java",
		"boolean result = obj.testMethod",
		"### Exception Handling",
		"May throw RuntimeException",
		"### Complexity",
		"O(n)",
		"### Best Practices",
		"Always check return value",
	}

	for _, expected := range expectedContent {
		if !strings.Contains(result, expected) {
			t.Errorf("Expected content not found: %s", expected)
		}
	}
}

func TestFormatMethodToMarkdownWithNilMethod(t *testing.T) {
	config := MarkdownConfig{}

	_, err := FormatMethodToMarkdown(nil, config)
	if err == nil {
		t.Error("Expected error for nil method documentation")
	}
}

func TestValidateMarkdownConfig(t *testing.T) {
	tests := []struct {
		name        string
		config      MarkdownConfig
		expectError bool
	}{
		{
			name: "valid_config",
			config: MarkdownConfig{
				Title:   "Test",
				Version: "1.0.0",
				Author:  "Test Author",
			},
			expectError: false,
		},
		{
			name: "missing_title",
			config: MarkdownConfig{
				Version: "1.0.0",
				Author:  "Test Author",
			},
			expectError: true,
		},
		{
			name: "missing_version",
			config: MarkdownConfig{
				Title:  "Test",
				Author: "Test Author",
			},
			expectError: true,
		},
		{
			name: "missing_author",
			config: MarkdownConfig{
				Title:   "Test",
				Version: "1.0.0",
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateMarkdownConfig(tt.config)
			if tt.expectError && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

func TestGenerateTableOfContents(t *testing.T) {
	methodDocs := map[string]*ai.MethodDocumentation{
		"method1": {MethodName: "method1"},
		"method2": {MethodName: "method2"},
		"method3": {MethodName: "method3"},
	}

	result := GenerateTableOfContents(methodDocs)

	if !strings.Contains(result, "## Table of Contents") {
		t.Error("Expected table of contents header not found")
	}

	// Check that all methods are included
	for name := range methodDocs {
		if !strings.Contains(result, name) {
			t.Errorf("Method %s not found in table of contents", name)
		}
	}
}

func TestGenerateTableOfContentsEmpty(t *testing.T) {
	emptyDocs := map[string]*ai.MethodDocumentation{}

	result := GenerateTableOfContents(emptyDocs)

	if !strings.Contains(result, "*No methods available.*") {
		t.Error("Expected empty message not found")
	}
}

func TestSanitizeMarkdown(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "asterisk",
			input:    "This *text* has asterisks",
			expected: "This \\*text\\* has asterisks",
		},
		{
			name:     "underscore",
			input:    "This _text_ has underscores",
			expected: "This \\_text\\_ has underscores",
		},
		{
			name:     "backtick",
			input:    "This `code` has backticks",
			expected: "This \\`code\\` has backticks",
		},
		{
			name:     "hash",
			input:    "# This is a header",
			expected: "\\# This is a header",
		},
		{
			name:     "brackets",
			input:    "[link](url)",
			expected: "\\[link\\]\\(url\\)",
		},
		{
			name:     "pipe",
			input:    "table | cell",
			expected: "table \\| cell",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SanitizeMarkdown(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestExportToFile(t *testing.T) {
	tests := []struct {
		name        string
		content     string
		filename    string
		expectError bool
	}{
		{
			name:        "valid_export",
			content:     "# Test Content",
			filename:    "test.md",
			expectError: false,
		},
		{
			name:        "empty_content",
			content:     "",
			filename:    "test.md",
			expectError: true,
		},
		{
			name:        "empty_filename",
			content:     "# Test Content",
			filename:    "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ExportToFile(tt.content, tt.filename)
			if tt.expectError && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

func TestMarkdownDocumentStructure(t *testing.T) {
	// Test that MarkdownDocument can be created and has expected fields
	config := MarkdownConfig{
		Title:       "Test",
		Version:     "1.0.0",
		Author:      "Test Author",
		Description: "Test description",
		IncludeTOC:  true,
	}

	classInfo := &core.DecompiledClass{
		ClassName:    "TestClass",
		PackageName: "com.example",
	}

	methodDocs := map[string]*ai.MethodDocumentation{
		"test": {MethodName: "test", Description: "Test method"},
	}

	doc := &MarkdownDocument{
		Config:      config,
		ClassInfo:   classInfo,
		MethodDocs:  methodDocs,
		GeneratedAt: time.Now(),
	}

	if doc.Config.Title != "Test" {
		t.Error("Config not set correctly")
	}
	if doc.ClassInfo.ClassName != "TestClass" {
		t.Error("ClassInfo not set correctly")
	}
	if len(doc.MethodDocs) != 1 {
		t.Error("MethodDocs not set correctly")
	}
	if doc.GeneratedAt.IsZero() {
		t.Error("GeneratedAt not set")
	}
}
