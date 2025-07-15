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

	"ai-jar-sdk-generator/src/ai"
	"ai-jar-sdk-generator/src/core"
)

// OpenAPIObject represents the root of an OpenAPI 3.0 document.
type OpenAPIObject struct {
	OpenAPI    string                 `json:"openapi"`
	Info       OpenAPIInfo            `json:"info"`
	Paths      map[string]PathItem    `json:"paths"`
	Components *Components            `json:"components,omitempty"`
	Servers    []Server               `json:"servers,omitempty"`
}

// OpenAPIInfo contains metadata about the API.
type OpenAPIInfo struct {
	Title       string  `json:"title"`
	Description string  `json:"description,omitempty"`
	Version     string  `json:"version"`
	Contact     *Contact `json:"contact,omitempty"`
}

// Contact information for the API.
type Contact struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	URL   string `json:"url,omitempty"`
}

// Server represents a server object.
type Server struct {
	URL         string `json:"url"`
	Description string `json:"description,omitempty"`
}

// PathItem describes the operations available on a single path.
type PathItem struct {
	Get    *Operation `json:"get,omitempty"`
	Post   *Operation `json:"post,omitempty"`
	Put    *Operation `json:"put,omitempty"`
	Delete *Operation `json:"delete,omitempty"`
	Patch  *Operation `json:"patch,omitempty"`
}

// Operation describes a single API operation on a path.
type Operation struct {
	Tags        []string              `json:"tags,omitempty"`
	Summary     string                `json:"summary,omitempty"`
	Description string                `json:"description,omitempty"`
	OperationID string                `json:"operationId,omitempty"`
	Parameters  []Parameter           `json:"parameters,omitempty"`
	RequestBody *RequestBody          `json:"requestBody,omitempty"`
	Responses   map[string]Response   `json:"responses"`
}

// Parameter describes a single operation parameter.
type Parameter struct {
	Name        string      `json:"name"`
	In          string      `json:"in"` // "query", "header", "path", "cookie"
	Description string      `json:"description,omitempty"`
	Required    bool        `json:"required,omitempty"`
	Schema      *Schema     `json:"schema,omitempty"`
}

// RequestBody describes a single request body.
type RequestBody struct {
	Description string               `json:"description,omitempty"`
	Content     map[string]MediaType `json:"content"`
	Required    bool                 `json:"required,omitempty"`
}

// Response describes a single response from an API Operation.
type Response struct {
	Description string               `json:"description"`
	Content     map[string]MediaType `json:"content,omitempty"`
}

// MediaType provides schema and examples for the media type identified by its key.
type MediaType struct {
	Schema *Schema `json:"schema,omitempty"`
}

// Schema represents a JSON Schema.
type Schema struct {
	Type        string             `json:"type,omitempty"`
	Format      string             `json:"format,omitempty"`
	Description string             `json:"description,omitempty"`
	Properties  map[string]*Schema `json:"properties,omitempty"`
	Required    []string           `json:"required,omitempty"`
	Items       *Schema            `json:"items,omitempty"`
	Ref         string             `json:"$ref,omitempty"`
}

// Components holds a set of reusable objects for different aspects of the OAS.
type Components struct {
	Schemas map[string]*Schema `json:"schemas,omitempty"`
}

// OpenAPIConfig holds configuration for OpenAPI generation.
type OpenAPIConfig struct {
	Title       string
	Description string
	Version     string
	ServerURL   string
}

// FormatToOpenAPI converts decompiled class information into an OpenAPI 3.0 structure.
func FormatToOpenAPI(classInfo *core.DecompiledClass, methodDocs map[string]*ai.MethodDocumentation, config OpenAPIConfig) (*OpenAPIObject, error) {
	if classInfo == nil {
		return nil, fmt.Errorf("classInfo cannot be nil")
	}

	// Initialize OpenAPI document
	doc := &OpenAPIObject{
		OpenAPI: "3.0.0",
		Info: OpenAPIInfo{
			Title:       getTitle(config.Title, classInfo.ClassName),
			Description: getDescription(config.Description, classInfo),
			Version:     getVersion(config.Version),
		},
		Paths:      make(map[string]PathItem),
		Components: &Components{
			Schemas: make(map[string]*Schema),
		},
	}

	// Add server if provided
	if config.ServerURL != "" {
		doc.Servers = []Server{
			{
				URL:         config.ServerURL,
				Description: "Generated API server",
			},
		}
	}

	// Process methods to generate paths
	err := processMethodsForPaths(doc, classInfo, methodDocs)
	if err != nil {
		return nil, fmt.Errorf("failed to process methods: %w", err)
	}

	// Generate schemas for data structures
	generateSchemas(doc, classInfo)

	return doc, nil
}

// processMethodsForPaths analyzes methods and generates OpenAPI paths.
func processMethodsForPaths(doc *OpenAPIObject, classInfo *core.DecompiledClass, methodDocs map[string]*ai.MethodDocumentation) error {
	for _, method := range classInfo.Methods {
		// Skip non-public methods
		if !isPublicMethod(method) {
			continue
		}

		// Try to extract REST endpoint information
		pathInfo := extractRESTPathInfo(method)
		if pathInfo == nil {
			continue
		}

		// Get method documentation
		methodDoc := methodDocs[method.Name]

		// Create operation
		operation := createOperation(method, methodDoc, classInfo)

		// Add to paths
		addOperationToPath(doc, pathInfo.Path, pathInfo.HTTPMethod, operation)
	}

	return nil
}

// RESTPathInfo holds extracted REST endpoint information.
type RESTPathInfo struct {
	Path       string
	HTTPMethod string
}

// extractRESTPathInfo tries to extract REST endpoint information from method annotations.
func extractRESTPathInfo(method core.DecompiledMethod) *RESTPathInfo {
	// Check for common REST annotations
	for annotation, value := range method.Annotations {
		switch annotation {
		case "RequestMapping", "GetMapping", "PostMapping", "PutMapping", "DeleteMapping", "PatchMapping":
			path := extractPathFromAnnotation(value)
			httpMethod := getHTTPMethodFromAnnotation(annotation, value)
			if path != "" && httpMethod != "" {
				return &RESTPathInfo{
					Path:       path,
					HTTPMethod: httpMethod,
				}
			}
		case "Path": // JAX-RS
			path := extractPathFromAnnotation(value)
			httpMethod := getJAXRSHTTPMethod(method)
			if path != "" && httpMethod != "" {
				return &RESTPathInfo{
					Path:       path,
					HTTPMethod: httpMethod,
				}
			}
		}
	}

	return nil
}

// extractPathFromAnnotation extracts the path value from annotation.
func extractPathFromAnnotation(annotationValue string) string {
	// Simple regex to extract path from annotation value
	// This is a simplified implementation
	pathRegex := regexp.MustCompile(`(?:value\s*=\s*)?["']([^"']+)["']`)
	matches := pathRegex.FindStringSubmatch(annotationValue)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

// getHTTPMethodFromAnnotation determines HTTP method from Spring annotation.
func getHTTPMethodFromAnnotation(annotation, value string) string {
	switch annotation {
	case "GetMapping":
		return "get"
	case "PostMapping":
		return "post"
	case "PutMapping":
		return "put"
	case "DeleteMapping":
		return "delete"
	case "PatchMapping":
		return "patch"
	case "RequestMapping":
		// Try to extract method from RequestMapping
		methodRegex := regexp.MustCompile(`method\s*=\s*RequestMethod\.([A-Z]+)`)
		matches := methodRegex.FindStringSubmatch(value)
		if len(matches) > 1 {
			return strings.ToLower(matches[1])
		}
		return "get" // Default to GET
	}
	return ""
}

// getJAXRSHTTPMethod determines HTTP method from JAX-RS annotations.
func getJAXRSHTTPMethod(method core.DecompiledMethod) string {
	for annotation := range method.Annotations {
		switch annotation {
		case "GET":
			return "get"
		case "POST":
			return "post"
		case "PUT":
			return "put"
		case "DELETE":
			return "delete"
		case "PATCH":
			return "patch"
		}
	}
	return ""
}

// createOperation creates an OpenAPI operation from method information.
func createOperation(method core.DecompiledMethod, methodDoc *ai.MethodDocumentation, classInfo *core.DecompiledClass) *Operation {
	operation := &Operation{
		Tags:        []string{classInfo.ClassName},
		Summary:     method.Name,
		OperationID: fmt.Sprintf("%s_%s", classInfo.ClassName, method.Name),
		Responses:   make(map[string]Response),
	}

	// Add description from AI documentation
	if methodDoc != nil {
		operation.Description = methodDoc.Description
		operation.Summary = methodDoc.Description
	}

	// Add parameters
	operation.Parameters = extractParameters(method, methodDoc)

	// Add request body if needed
	if needsRequestBody(method) {
		operation.RequestBody = createRequestBody(method, methodDoc)
	}

	// Add responses
	operation.Responses["200"] = createSuccessResponse(method, methodDoc)
	operation.Responses["400"] = Response{
		Description: "Bad Request",
	}
	operation.Responses["500"] = Response{
		Description: "Internal Server Error",
	}

	return operation
}

// extractParameters extracts parameters from method signature.
func extractParameters(method core.DecompiledMethod, methodDoc *ai.MethodDocumentation) []Parameter {
	var parameters []Parameter

	// Simple parameter extraction from method signature
	for i, param := range method.Parameters {
		// Parse parameter (simplified)
		parts := strings.Fields(param)
		if len(parts) >= 2 {
			paramType := parts[0]
			paramName := parts[1]

			parameter := Parameter{
				Name:     paramName,
				In:       "query", // Default to query parameter
				Required: true,
				Schema:   javaTypeToSchema(paramType),
			}

			// Add description from AI documentation
			if methodDoc != nil && i < len(methodDoc.Parameters) {
				parameter.Description = methodDoc.Parameters[i].Description
			}

			parameters = append(parameters, parameter)
		}
	}

	return parameters
}

// needsRequestBody determines if the method needs a request body.
func needsRequestBody(method core.DecompiledMethod) bool {
	// Check if method has POST, PUT, PATCH annotations
	for annotation := range method.Annotations {
		if annotation == "PostMapping" || annotation == "PutMapping" || annotation == "PatchMapping" ||
			annotation == "POST" || annotation == "PUT" || annotation == "PATCH" {
			return true
		}
	}
	return false
}

// createRequestBody creates a request body specification.
func createRequestBody(method core.DecompiledMethod, methodDoc *ai.MethodDocumentation) *RequestBody {
	return &RequestBody{
		Description: "Request body",
		Required:    true,
		Content: map[string]MediaType{
			"application/json": {
				Schema: &Schema{
					Type: "object",
				},
			},
		},
	}
}

// createSuccessResponse creates a success response specification.
func createSuccessResponse(method core.DecompiledMethod, methodDoc *ai.MethodDocumentation) Response {
	response := Response{
		Description: "Successful response",
	}

	// Add content based on return type
	if method.ReturnType != "void" {
		response.Content = map[string]MediaType{
			"application/json": {
				Schema: javaTypeToSchema(method.ReturnType),
			},
		}
	}

	// Add description from AI documentation
	if methodDoc != nil && methodDoc.ReturnValue != "" {
		response.Description = methodDoc.ReturnValue
	}

	return response
}

// javaTypeToSchema converts Java type to OpenAPI schema.
func javaTypeToSchema(javaType string) *Schema {
	switch javaType {
	case "String", "java.lang.String":
		return &Schema{Type: "string"}
	case "int", "Integer", "java.lang.Integer":
		return &Schema{Type: "integer", Format: "int32"}
	case "long", "Long", "java.lang.Long":
		return &Schema{Type: "integer", Format: "int64"}
	case "double", "Double", "java.lang.Double":
		return &Schema{Type: "number", Format: "double"}
	case "float", "Float", "java.lang.Float":
		return &Schema{Type: "number", Format: "float"}
	case "boolean", "Boolean", "java.lang.Boolean":
		return &Schema{Type: "boolean"}
	case "Date", "java.util.Date":
		return &Schema{Type: "string", Format: "date-time"}
	default:
		// For complex types, reference to components
		if strings.Contains(javaType, "List") || strings.Contains(javaType, "[]") {
			return &Schema{
				Type:  "array",
				Items: &Schema{Type: "object"},
			}
		}
		return &Schema{Type: "object"}
	}
}

// addOperationToPath adds an operation to the specified path.
func addOperationToPath(doc *OpenAPIObject, path, httpMethod string, operation *Operation) {
	pathItem, exists := doc.Paths[path]
	if !exists {
		pathItem = PathItem{}
	}

	switch httpMethod {
	case "get":
		pathItem.Get = operation
	case "post":
		pathItem.Post = operation
	case "put":
		pathItem.Put = operation
	case "delete":
		pathItem.Delete = operation
	case "patch":
		pathItem.Patch = operation
	}

	doc.Paths[path] = pathItem
}

// generateSchemas generates component schemas for data structures.
func generateSchemas(doc *OpenAPIObject, classInfo *core.DecompiledClass) {
	// Generate schema for the main class
	schema := &Schema{
		Type:        "object",
		Description: fmt.Sprintf("Schema for %s", classInfo.ClassName),
		Properties:  make(map[string]*Schema),
	}

	// Add fields as properties
	for _, field := range classInfo.Fields {
		if isPublicField(field) {
			schema.Properties[field.Name] = javaTypeToSchema(field.Type)
		}
	}

	doc.Components.Schemas[classInfo.ClassName] = schema
}

// Helper functions

func getTitle(configTitle, className string) string {
	if configTitle != "" {
		return configTitle
	}
	return fmt.Sprintf("%s API", className)
}

func getDescription(configDesc string, classInfo *core.DecompiledClass) string {
	if configDesc != "" {
		return configDesc
	}
	return fmt.Sprintf("Generated API documentation for %s", classInfo.ClassName)
}

func getVersion(configVersion string) string {
	if configVersion != "" {
		return configVersion
	}
	return "1.0.0"
}

func isPublicMethod(method core.DecompiledMethod) bool {
	for _, modifier := range method.Modifiers {
		if modifier == "public" {
			return true
		}
	}
	return false
}

// FormatToOpenAPIJSON converts the OpenAPI object to JSON string.
func FormatToOpenAPIJSON(openAPIObj *OpenAPIObject) (string, error) {
	data, err := json.MarshalIndent(openAPIObj, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal OpenAPI object to JSON: %w", err)
	}
	return string(data), nil
}