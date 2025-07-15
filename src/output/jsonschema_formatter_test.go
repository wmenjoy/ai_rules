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
	"strings"
	"testing"

	"ai-jar-sdk-generator/src/core"
)

func TestFormatToJSONSchema(t *testing.T) {
	// Create test data
	classInfo := &core.DecompiledClass{
		ClassName:   "User",
		PackageName: "com.example.model",
		Modifiers:   []string{"public"},
		Fields: []core.DecompiledField{
			{
				Name:        "id",
				Type:        "Long",
				Modifiers:   []string{"public"},
				Annotations: map[string]string{"NotNull": ""},
				JavaDoc:     "User ID",
			},
			{
				Name:        "name",
				Type:        "String",
				Modifiers:   []string{"public"},
				Annotations: map[string]string{"Size": "min=1, max=100"},
				JavaDoc:     "User name",
			},
			{
				Name:        "email",
				Type:        "String",
				Modifiers:   []string{"public"},
				Annotations: map[string]string{"Email": ""},
				JavaDoc:     "User email",
			},
			{
				Name:        "age",
				Type:        "int",
				Modifiers:   []string{"public"},
				Annotations: map[string]string{"Min": "0", "Max": "150"},
				JavaDoc:     "User age",
			},
			{
				Name:        "isActive",
				Type:        "boolean",
				Modifiers:   []string{"public"},
				JavaDoc:     "Whether user is active",
			},
		},
	}

	config := DefaultJSONSchemaConfig()
	schema, err := FormatToJSONSchema(classInfo, config)

	if err != nil {
		t.Fatalf("FormatToJSONSchema failed: %v", err)
	}

	// Verify basic structure
	if schema.Schema != config.SchemaVersion {
		t.Errorf("Expected schema version %s, got %s", config.SchemaVersion, schema.Schema)
	}

	if schema.Title != "User" {
		t.Errorf("Expected title 'User', got %s", schema.Title)
	}

	if schema.Type != "object" {
		t.Errorf("Expected type 'object', got %s", schema.Type)
	}

	// Verify properties
	if len(schema.Properties) != 5 {
		t.Errorf("Expected 5 properties, got %d", len(schema.Properties))
	}

	// Check id field
	idProp := schema.Properties["id"]
	if idProp == nil {
		t.Error("Expected 'id' property to exist")
	} else {
		if idProp.Type != "integer" {
			t.Errorf("Expected id type 'integer', got %s", idProp.Type)
		}
		if idProp.Format != "int64" {
			t.Errorf("Expected id format 'int64', got %s", idProp.Format)
		}
		if idProp.Description != "User ID" {
			t.Errorf("Expected id description 'User ID', got %s", idProp.Description)
		}
	}

	// Check name field with size validation
	nameProp := schema.Properties["name"]
	if nameProp == nil {
		t.Error("Expected 'name' property to exist")
	} else {
		if nameProp.Type != "string" {
			t.Errorf("Expected name type 'string', got %s", nameProp.Type)
		}
		if nameProp.MinLength == nil || *nameProp.MinLength != 1 {
			t.Errorf("Expected name minLength 1, got %v", nameProp.MinLength)
		}
		if nameProp.MaxLength == nil || *nameProp.MaxLength != 100 {
			t.Errorf("Expected name maxLength 100, got %v", nameProp.MaxLength)
		}
	}

	// Check email field with format
	emailProp := schema.Properties["email"]
	if emailProp == nil {
		t.Error("Expected 'email' property to exist")
	} else {
		if emailProp.Format != "email" {
			t.Errorf("Expected email format 'email', got %s", emailProp.Format)
		}
	}

	// Check age field with min/max validation
	ageProp := schema.Properties["age"]
	if ageProp == nil {
		t.Error("Expected 'age' property to exist")
	} else {
		if ageProp.Type != "integer" {
			t.Errorf("Expected age type 'integer', got %s", ageProp.Type)
		}
		if ageProp.Minimum == nil || *ageProp.Minimum != 0 {
			t.Errorf("Expected age minimum 0, got %v", ageProp.Minimum)
		}
		if ageProp.Maximum == nil || *ageProp.Maximum != 150 {
			t.Errorf("Expected age maximum 150, got %v", ageProp.Maximum)
		}
	}

	// Check boolean field
	isActiveProp := schema.Properties["isActive"]
	if isActiveProp == nil {
		t.Error("Expected 'isActive' property to exist")
	} else {
		if isActiveProp.Type != "boolean" {
			t.Errorf("Expected isActive type 'boolean', got %s", isActiveProp.Type)
		}
	}

	// Check required fields
	if len(schema.Required) != 1 {
		t.Errorf("Expected 1 required field, got %d", len(schema.Required))
	} else if schema.Required[0] != "id" {
		t.Errorf("Expected required field 'id', got %s", schema.Required[0])
	}
}

func TestFormatToJSONSchemaJSON(t *testing.T) {
	// Create simple test data
	classInfo := &core.DecompiledClass{
		ClassName:   "SimpleClass",
		PackageName: "com.example",
		Fields: []core.DecompiledField{
			{
				Name:      "value",
				Type:      "String",
				Modifiers: []string{"public"},
			},
		},
	}

	config := DefaultJSONSchemaConfig()
	schema, err := FormatToJSONSchema(classInfo, config)
	if err != nil {
		t.Fatalf("FormatToJSONSchema failed: %v", err)
	}

	jsonStr, err := FormatToJSONSchemaJSON(schema)
	if err != nil {
		t.Fatalf("FormatToJSONSchemaJSON failed: %v", err)
	}

	// Verify it's valid JSON
	var result map[string]interface{}
	err = json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		t.Fatalf("Generated JSON is invalid: %v", err)
	}

	// Check some basic fields
	if result["$schema"] != config.SchemaVersion {
		t.Errorf("Expected schema version %s in JSON, got %v", config.SchemaVersion, result["$schema"])
	}

	if result["title"] != "SimpleClass" {
		t.Errorf("Expected title 'SimpleClass' in JSON, got %v", result["title"])
	}

	if result["type"] != "object" {
		t.Errorf("Expected type 'object' in JSON, got %v", result["type"])
	}
}

func TestJavaTypeToJSONSchema(t *testing.T) {
	config := DefaultJSONSchemaConfig()

	tests := []struct {
		javaType     string
		expectedType string
		expectedFormat string
	}{
		{"String", "string", ""},
		{"int", "integer", "int32"},
		{"long", "integer", "int64"},
		{"double", "number", "double"},
		{"boolean", "boolean", ""},
		{"Date", "string", "date-time"},
		{"LocalDate", "string", "date"},
		{"UUID", "string", "uuid"},
		{"BigDecimal", "number", ""},
		{"BigInteger", "integer", ""},
	}

	for _, test := range tests {
		schema := javaTypeToJSONSchema(test.javaType, config)
		if schema.Type != test.expectedType {
			t.Errorf("For type %s, expected %s, got %s", test.javaType, test.expectedType, schema.Type)
		}
		if schema.Format != test.expectedFormat {
			t.Errorf("For type %s, expected format %s, got %s", test.javaType, test.expectedFormat, schema.Format)
		}
	}
}

func TestArrayAndCollectionTypes(t *testing.T) {
	config := DefaultJSONSchemaConfig()

	// Test array type
	arraySchema := javaTypeToJSONSchema("String[]", config)
	if arraySchema.Type != "array" {
		t.Errorf("Expected array type, got %s", arraySchema.Type)
	}
	if arraySchema.Items == nil || arraySchema.Items.Type != "string" {
		t.Error("Expected array items to be string type")
	}

	// Test List type
	listSchema := javaTypeToJSONSchema("List<Integer>", config)
	if listSchema.Type != "array" {
		t.Errorf("Expected array type for List, got %s", listSchema.Type)
	}
	if listSchema.Items == nil || listSchema.Items.Type != "integer" {
		t.Error("Expected List items to be integer type")
	}

	// Test Map type
	mapSchema := javaTypeToJSONSchema("Map<String, Object>", config)
	if mapSchema.Type != "object" {
		t.Errorf("Expected object type for Map, got %s", mapSchema.Type)
	}
	if mapSchema.AdditionalProperties != true {
		t.Error("Expected Map to allow additional properties")
	}
}

func TestValidationAnnotations(t *testing.T) {
	schema := &JSONSchema{Type: "string"}
	annotations := map[string]string{
		"Size":    "min=5, max=50",
		"Pattern": "regexp='^[a-zA-Z]+$'",
		"Email":   "",
	}

	applyValidationAnnotations(schema, annotations)

	if schema.MinLength == nil || *schema.MinLength != 5 {
		t.Errorf("Expected minLength 5, got %v", schema.MinLength)
	}
	if schema.MaxLength == nil || *schema.MaxLength != 50 {
		t.Errorf("Expected maxLength 50, got %v", schema.MaxLength)
	}
	if schema.Format != "email" {
		t.Errorf("Expected format 'email', got %s", schema.Format)
	}
}

func TestFormatToJSONSchemaWithNilInput(t *testing.T) {
	config := DefaultJSONSchemaConfig()
	schema, err := FormatToJSONSchema(nil, config)

	if err == nil {
		t.Error("Expected error for nil input, got nil")
	}
	if schema != nil {
		t.Error("Expected nil schema for nil input")
	}
	if !strings.Contains(err.Error(), "cannot be nil") {
		t.Errorf("Expected error message to contain 'cannot be nil', got: %s", err.Error())
	}
}

func TestJSONSchemaConfigDefaults(t *testing.T) {
	config := DefaultJSONSchemaConfig()

	if config.SchemaVersion != "http://json-schema.org/draft-07/schema#" {
		t.Errorf("Expected default schema version 'http://json-schema.org/draft-07/schema#', got %s", config.SchemaVersion)
	}
	if config.BaseURI != "#" {
		t.Errorf("Expected default base URI '#', got %s", config.BaseURI)
	}
	if !config.IncludeTitle {
		t.Error("Expected IncludeTitle to be true by default")
	}
	if config.StrictMode {
		t.Error("Expected StrictMode to be false by default")
	}
}

func TestInnerClassDefinitions(t *testing.T) {
	// Create test data with inner class
	classInfo := &core.DecompiledClass{
		ClassName:   "OuterClass",
		PackageName: "com.example",
		Fields: []core.DecompiledField{
			{
				Name:      "value",
				Type:      "String",
				Modifiers: []string{"public"},
			},
		},
		InnerClasses: []core.DecompiledClass{
			{
				ClassName: "InnerClass",
				Fields: []core.DecompiledField{
				{
					Name:      "innerValue",
					Type:      "int",
					Modifiers: []string{"public"},
				},
			},
			},
		},
	}

	config := DefaultJSONSchemaConfig()
	schema, err := FormatToJSONSchema(classInfo, config)

	if err != nil {
		t.Fatalf("FormatToJSONSchema failed: %v", err)
	}

	// Check that definitions exist
	if schema.Definitions == nil {
		t.Error("Expected definitions to exist")
	}

	if len(schema.Definitions) != 1 {
		t.Errorf("Expected 1 definition, got %d", len(schema.Definitions))
	}

	// Check inner class definition
	innerDef := schema.Definitions["InnerClass"]
	if innerDef == nil {
		t.Error("Expected InnerClass definition to exist")
	} else {
		if innerDef.Type != "object" {
			t.Errorf("Expected InnerClass type 'object', got %s", innerDef.Type)
		}
		if len(innerDef.Properties) != 1 {
			t.Errorf("Expected 1 property in InnerClass, got %d", len(innerDef.Properties))
		}
		if innerDef.Properties["innerValue"] == nil {
			t.Error("Expected innerValue property in InnerClass")
		}
	}
}

func TestGenerateMultipleSchemas(t *testing.T) {
	classes := []*core.DecompiledClass{
		{
			ClassName: "Class1",
			Fields: []core.DecompiledField{
				{Name: "field1", Type: "String", Modifiers: []string{"public"}},
			},
		},
		{
			ClassName: "Class2",
			Fields: []core.DecompiledField{
				{Name: "field2", Type: "int", Modifiers: []string{"public"}},
			},
		},
	}

	config := DefaultJSONSchemaConfig()
	schemas, err := GenerateMultipleSchemas(classes, config)

	if err != nil {
		t.Fatalf("GenerateMultipleSchemas failed: %v", err)
	}

	if len(schemas) != 2 {
		t.Errorf("Expected 2 schemas, got %d", len(schemas))
	}

	if schemas["Class1"] == nil {
		t.Error("Expected Class1 schema to exist")
	}

	if schemas["Class2"] == nil {
		t.Error("Expected Class2 schema to exist")
	}
}

func TestStrictModeRequiredFields(t *testing.T) {
	classInfo := &core.DecompiledClass{
		ClassName: "TestClass",
		Fields: []core.DecompiledField{
			{
				Name:         "field1",
				Type:         "String",
				Modifiers:    []string{"public"},
				DefaultValue: "", // No default value
			},
			{
				Name:         "field2",
				Type:         "String",
				Modifiers:    []string{"public"},
				DefaultValue: "default", // Has default value
			},
		},
	}

	// Test with strict mode
	strictConfig := JSONSchemaConfig{
		SchemaVersion: "http://json-schema.org/draft-07/schema#",
		BaseURI:       "#",
		IncludeTitle:  true,
		StrictMode:    true,
	}

	schema, err := FormatToJSONSchema(classInfo, strictConfig)
	if err != nil {
		t.Fatalf("FormatToJSONSchema failed: %v", err)
	}

	// In strict mode, field1 should be required (no default value)
	if len(schema.Required) != 1 {
		t.Errorf("Expected 1 required field in strict mode, got %d", len(schema.Required))
	} else if schema.Required[0] != "field1" {
		t.Errorf("Expected field1 to be required, got %s", schema.Required[0])
	}
}