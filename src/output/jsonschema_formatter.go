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
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"ai-jar-sdk-generator/src/core"
)

// JSONSchema represents a JSON Schema definition.
type JSONSchema struct {
	Schema               string                 `json:"$schema,omitempty"`
	ID                   string                 `json:"$id,omitempty"`
	Title                string                 `json:"title,omitempty"`
	Description          string                 `json:"description,omitempty"`
	Type                 string                 `json:"type,omitempty"`
	Format               string                 `json:"format,omitempty"`
	Properties           map[string]*JSONSchema `json:"properties,omitempty"`
	Required             []string               `json:"required,omitempty"`
	Items                *JSONSchema            `json:"items,omitempty"`
	Enum                 []interface{}          `json:"enum,omitempty"`
	Minimum              *float64               `json:"minimum,omitempty"`
	Maximum              *float64               `json:"maximum,omitempty"`
	MinLength            *int                   `json:"minLength,omitempty"`
	MaxLength            *int                   `json:"maxLength,omitempty"`
	Pattern              string                 `json:"pattern,omitempty"`
	Ref                  string                 `json:"$ref,omitempty"`
	Definitions          map[string]*JSONSchema `json:"definitions,omitempty"`
	AdditionalProperties interface{}            `json:"additionalProperties,omitempty"`
}

// JSONSchemaConfig holds configuration for JSON Schema generation.
type JSONSchemaConfig struct {
	SchemaVersion string // JSON Schema version (e.g., "http://json-schema.org/draft-07/schema#")
	BaseURI       string // Base URI for schema IDs
	IncludeTitle  bool   // Whether to include title in schemas
	StrictMode    bool   // Whether to use strict validation rules
}

// DefaultJSONSchemaConfig returns a default configuration.
func DefaultJSONSchemaConfig() JSONSchemaConfig {
	return JSONSchemaConfig{
		SchemaVersion: "http://json-schema.org/draft-07/schema#",
		BaseURI:       "#",
		IncludeTitle:  true,
		StrictMode:    false,
	}
}

// FormatToJSONSchema converts a Java class structure to JSON Schema.
func FormatToJSONSchema(classInfo *core.DecompiledClass, config JSONSchemaConfig) (*JSONSchema, error) {
	if classInfo == nil {
		return nil, fmt.Errorf("classInfo cannot be nil")
	}

	// Create root schema
	schema := &JSONSchema{
		Schema:      config.SchemaVersion,
		ID:          fmt.Sprintf("%s/%s", config.BaseURI, classInfo.ClassName),
		Type:        "object",
		Properties:  make(map[string]*JSONSchema),
		Definitions: make(map[string]*JSONSchema),
	}

	if config.IncludeTitle {
		schema.Title = classInfo.ClassName
	}

	// Add description from class information
	if classInfo.PackageName != "" {
		schema.Description = fmt.Sprintf("Schema for %s class from package %s", classInfo.ClassName, classInfo.PackageName)
	} else {
		schema.Description = fmt.Sprintf("Schema for %s class", classInfo.ClassName)
	}

	// Process fields
	err := processFieldsForSchema(schema, classInfo, config)
	if err != nil {
		return nil, fmt.Errorf("failed to process fields: %w", err)
	}

	// Process inner classes as definitions
	processInnerClassesForDefinitions(schema, classInfo, config)

	return schema, nil
}

// processFieldsForSchema processes class fields and adds them as properties.
func processFieldsForSchema(schema *JSONSchema, classInfo *core.DecompiledClass, config JSONSchemaConfig) error {
	var required []string

	for _, field := range classInfo.Fields {
		// Skip private fields unless in strict mode
		if !config.StrictMode && !isPublicField(field) {
			continue
		}

		// Skip static fields
		if isStaticField(field) {
			continue
		}

		// Convert field to JSON Schema property
		fieldSchema := javaTypeToJSONSchema(field.Type, config)

		// Add field description from JavaDoc
		if field.JavaDoc != "" {
			fieldSchema.Description = field.JavaDoc
		}

		// Add validation annotations
		applyValidationAnnotations(fieldSchema, field.Annotations)

		// Add to properties
		schema.Properties[field.Name] = fieldSchema

		// Check if field is required (not nullable, no default value)
		if isRequiredField(field, config) {
			required = append(required, field.Name)
		}
	}

	if len(required) > 0 {
		schema.Required = required
	}

	return nil
}

// processInnerClassesForDefinitions processes inner classes as schema definitions.
func processInnerClassesForDefinitions(schema *JSONSchema, classInfo *core.DecompiledClass, config JSONSchemaConfig) {
	for _, innerClass := range classInfo.InnerClasses {
		// Create schema for inner class
		innerSchema := &JSONSchema{
			Type:       "object",
			Properties: make(map[string]*JSONSchema),
		}

		if config.IncludeTitle {
			innerSchema.Title = innerClass.ClassName
		}

		innerSchema.Description = fmt.Sprintf("Schema for inner class %s", innerClass.ClassName)

		// Process inner class fields
		processFieldsForSchema(innerSchema, &innerClass, config)

		// Add to definitions
		schema.Definitions[innerClass.ClassName] = innerSchema
	}
}

// javaTypeToJSONSchema converts Java type to JSON Schema.
func javaTypeToJSONSchema(javaType string, config JSONSchemaConfig) *JSONSchema {
	// Remove generic type parameters for basic type mapping
	baseType := extractBaseType(javaType)

	switch baseType {
	// String types
	case "String", "java.lang.String":
		return &JSONSchema{Type: "string"}
	case "char", "Character", "java.lang.Character":
		return &JSONSchema{Type: "string", MinLength: intPtr(1), MaxLength: intPtr(1)}

	// Integer types
	case "byte", "Byte", "java.lang.Byte":
		return &JSONSchema{Type: "integer", Minimum: float64Ptr(-128), Maximum: float64Ptr(127)}
	case "short", "Short", "java.lang.Short":
		return &JSONSchema{Type: "integer", Minimum: float64Ptr(-32768), Maximum: float64Ptr(32767)}
	case "int", "Integer", "java.lang.Integer":
		return &JSONSchema{Type: "integer", Format: "int32"}
	case "long", "Long", "java.lang.Long":
		return &JSONSchema{Type: "integer", Format: "int64"}

	// Floating point types
	case "float", "Float", "java.lang.Float":
		return &JSONSchema{Type: "number", Format: "float"}
	case "double", "Double", "java.lang.Double":
		return &JSONSchema{Type: "number", Format: "double"}

	// Boolean type
	case "boolean", "Boolean", "java.lang.Boolean":
		return &JSONSchema{Type: "boolean"}

	// Date and time types
	case "Date", "java.util.Date":
		return &JSONSchema{Type: "string", Format: "date-time"}
	case "LocalDate", "java.time.LocalDate":
		return &JSONSchema{Type: "string", Format: "date"}
	case "LocalTime", "java.time.LocalTime":
		return &JSONSchema{Type: "string", Format: "time"}
	case "LocalDateTime", "java.time.LocalDateTime":
		return &JSONSchema{Type: "string", Format: "date-time"}
	case "Instant", "java.time.Instant":
		return &JSONSchema{Type: "string", Format: "date-time"}

	// UUID type
	case "UUID", "java.util.UUID":
		return &JSONSchema{Type: "string", Format: "uuid"}

	// BigDecimal and BigInteger
	case "BigDecimal", "java.math.BigDecimal":
		return &JSONSchema{Type: "number"}
	case "BigInteger", "java.math.BigInteger":
		return &JSONSchema{Type: "integer"}

	default:
		// Handle collections and arrays
		if isArrayType(javaType) {
			elementType := extractArrayElementType(javaType)
			return &JSONSchema{
				Type:  "array",
				Items: javaTypeToJSONSchema(elementType, config),
			}
		}

		if isCollectionType(javaType) {
			elementType := extractGenericType(javaType)
			return &JSONSchema{
				Type:  "array",
				Items: javaTypeToJSONSchema(elementType, config),
			}
		}

		if isMapType(javaType) {
			return &JSONSchema{
				Type:                 "object",
				AdditionalProperties: true,
			}
		}

		// For custom objects, create a reference or object schema
		if isCustomType(baseType) {
			return &JSONSchema{
				Ref: fmt.Sprintf("#/definitions/%s", baseType),
			}
		}

		// Default to object type
		return &JSONSchema{Type: "object"}
	}
}

// applyValidationAnnotations applies Java validation annotations to JSON Schema.
func applyValidationAnnotations(schema *JSONSchema, annotations map[string]string) {
	for annotation, value := range annotations {
		switch annotation {
		case "NotNull", "NonNull":
			// This will be handled at the property level by adding to required array
		case "Size":
			applySizeAnnotation(schema, value)
		case "Min":
			if min := parseFloatFromAnnotation(value); min != nil {
				schema.Minimum = min
			}
		case "Max":
			if max := parseFloatFromAnnotation(value); max != nil {
				schema.Maximum = max
			}
		case "Pattern":
			if pattern := parseStringFromAnnotation(value); pattern != "" {
				schema.Pattern = pattern
			}
		case "Email":
			schema.Format = "email"
		case "URL":
			schema.Format = "uri"
		}
	}
}

// applySizeAnnotation applies @Size annotation to schema.
func applySizeAnnotation(schema *JSONSchema, value string) {
	// Parse @Size(min=X, max=Y) annotation
	minRegex := regexp.MustCompile(`min\s*=\s*(\d+)`)
	maxRegex := regexp.MustCompile(`max\s*=\s*(\d+)`)

	if matches := minRegex.FindStringSubmatch(value); len(matches) > 1 {
		if schema.Type == "string" {
			if min := parseInt(matches[1]); min != nil {
				schema.MinLength = min
			}
		} else if schema.Type == "array" {
			if min := parseFloat64(matches[1]); min != nil {
				schema.Minimum = min
			}
		}
	}

	if matches := maxRegex.FindStringSubmatch(value); len(matches) > 1 {
		if schema.Type == "string" {
			if max := parseInt(matches[1]); max != nil {
				schema.MaxLength = max
			}
		} else if schema.Type == "array" {
			if max := parseFloat64(matches[1]); max != nil {
				schema.Maximum = max
			}
		}
	}
}

// Helper functions for type checking

func extractBaseType(javaType string) string {
	// Remove generic parameters
	if idx := strings.Index(javaType, "<"); idx != -1 {
		return javaType[:idx]
	}
	// Remove array brackets
	if idx := strings.Index(javaType, "["); idx != -1 {
		return javaType[:idx]
	}
	return javaType
}

func isArrayType(javaType string) bool {
	return strings.Contains(javaType, "[]")
}

func isCollectionType(javaType string) bool {
	collectionTypes := []string{"List", "Set", "Collection", "ArrayList", "LinkedList", "HashSet", "TreeSet"}
	baseType := extractBaseType(javaType)
	for _, collType := range collectionTypes {
		if strings.Contains(baseType, collType) {
			return true
		}
	}
	return false
}

func isMapType(javaType string) bool {
	mapTypes := []string{"Map", "HashMap", "TreeMap", "LinkedHashMap"}
	baseType := extractBaseType(javaType)
	for _, mapType := range mapTypes {
		if strings.Contains(baseType, mapType) {
			return true
		}
	}
	return false
}

func isCustomType(javaType string) bool {
	// Check if it's not a primitive or standard Java type
	standardTypes := []string{
		"String", "Integer", "Long", "Double", "Float", "Boolean", "Character", "Byte", "Short",
		"Date", "LocalDate", "LocalTime", "LocalDateTime", "Instant", "UUID", "BigDecimal", "BigInteger",
	}

	for _, stdType := range standardTypes {
		if strings.Contains(javaType, stdType) {
			return false
		}
	}

	// If it contains a package name or is capitalized, it's likely a custom type
	return strings.Contains(javaType, ".") || (len(javaType) > 0 && javaType[0] >= 'A' && javaType[0] <= 'Z')
}

func extractArrayElementType(javaType string) string {
	if idx := strings.Index(javaType, "["); idx != -1 {
		return javaType[:idx]
	}
	return "Object"
}

func extractGenericType(javaType string) string {
	// Extract type from List<Type>, Set<Type>, etc.
	start := strings.Index(javaType, "<")
	end := strings.LastIndex(javaType, ">")
	if start != -1 && end != -1 && end > start {
		return javaType[start+1 : end]
	}
	return "Object"
}

func isRequiredField(field core.DecompiledField, config JSONSchemaConfig) bool {
	// Check for @NotNull or @NonNull annotations
	for annotation := range field.Annotations {
		if annotation == "NotNull" || annotation == "NonNull" {
			return true
		}
	}

	// In strict mode, all fields without default values are required
	if config.StrictMode && field.DefaultValue == "" {
		return true
	}

	return false
}

func isPublicField(field core.DecompiledField) bool {
	for _, modifier := range field.Modifiers {
		if modifier == "public" {
			return true
		}
	}
	return false
}

func isStaticField(field core.DecompiledField) bool {
	for _, modifier := range field.Modifiers {
		if modifier == "static" {
			return true
		}
	}
	return false
}

// Utility functions for parsing annotations

func parseFloatFromAnnotation(value string) *float64 {
	// Simple regex to extract numeric value
	regex := regexp.MustCompile(`(\d+(?:\.\d+)?)`)
	matches := regex.FindStringSubmatch(value)
	if len(matches) > 1 {
		if f := parseFloat64(matches[1]); f != nil {
			return f
		}
	}
	return nil
}

func parseStringFromAnnotation(value string) string {
	// Extract string value from annotation
	regex := regexp.MustCompile(`["']([^"']+)["']`)
	matches := regex.FindStringSubmatch(value)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

// Helper functions for pointer creation

func intPtr(i int) *int {
	return &i
}

func float64Ptr(f float64) *float64 {
	return &f
}

func parseInt(s string) *int {
	var i int
	if _, err := fmt.Sscanf(s, "%d", &i); err == nil {
		return &i
	}
	return nil
}

func parseFloat64(s string) *float64 {
	var f float64
	if _, err := fmt.Sscanf(s, "%f", &f); err == nil {
		return &f
	}
	return nil
}

// FormatToJSONSchemaJSON converts the JSON Schema object to JSON string.
func FormatToJSONSchemaJSON(schema *JSONSchema) (string, error) {
	data, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal JSON Schema to JSON: %w", err)
	}
	return string(data), nil
}

// GenerateMultipleSchemas generates JSON schemas for multiple classes.
func GenerateMultipleSchemas(classes []*core.DecompiledClass, config JSONSchemaConfig) (map[string]*JSONSchema, error) {
	schemas := make(map[string]*JSONSchema)

	for _, class := range classes {
		schema, err := FormatToJSONSchema(class, config)
		if err != nil {
			return nil, fmt.Errorf("failed to generate schema for class %s: %w", class.ClassName, err)
		}
		schemas[class.ClassName] = schema
	}

	return schemas, nil
}
