package api

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// DocumentType represents the type of document to download
type DocumentType string

const (
	MarkdownDoc   DocumentType = "markdown"
	OpenAPIDoc    DocumentType = "openapi_json"
	JSONSchemaDoc DocumentType = "jsonschema"
	AllDocs       DocumentType = "all"
)

// DownloadRequest represents a document download request
type DownloadRequest struct {
	DocID    string       `json:"doc_id"`
	DocType  DocumentType `json:"doc_type"`
	JarPath  string       `json:"jar_path,omitempty"`
	Packaged bool         `json:"packaged,omitempty"` // Whether to package multiple docs into ZIP
}

// DocumentContent represents generated document content
type DocumentContent struct {
	Content     []byte
	ContentType string
	FileName    string
}

// HandleDownloadDoc serves generated documents for download
func HandleDownloadDoc(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetDownload(w, r)
	case http.MethodPost:
		handlePostDownload(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// handleGetDownload handles GET requests for document download
func handleGetDownload(w http.ResponseWriter, r *http.Request) {
	docType := DocumentType(r.URL.Query().Get("type"))
	docID := r.URL.Query().Get("docId"))
	jarPath := r.URL.Query().Get("jarPath"))
	packaged := r.URL.Query().Get("packaged") == "true"

	req := DownloadRequest{
		DocID:    docID,
		DocType:  docType,
		JarPath:  jarPath,
		Packaged: packaged,
	}

	processDownloadRequest(w, req)
}

// handlePostDownload handles POST requests for document download
func handlePostDownload(w http.ResponseWriter, r *http.Request) {
	var req DownloadRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	processDownloadRequest(w, req)
}

// processDownloadRequest processes the download request and serves the document
func processDownloadRequest(w http.ResponseWriter, req DownloadRequest) {
	if req.DocType == AllDocs || req.Packaged {
		servePackagedDocuments(w, req)
		return
	}

	docContent, err := generateDocument(req)
	if err != nil {
		http.Error(w, "Failed to generate document: "+err.Error(), http.StatusInternalServerError)
		return
	}

	serveDocument(w, docContent)
}

// generateDocument generates the requested document
func generateDocument(req DownloadRequest) (*DocumentContent, error) {
	// If jarPath is provided, parse the JAR file first
	var classInfo *core.ClassInfo
	var err error

	if req.JarPath != "" {
		// Parse JAR file to get class information
		classInfo, err = parseJarFile(req.JarPath)
		if err != nil {
			return nil, fmt.Errorf("failed to parse JAR file: %w", err)
		}
	} else if req.DocID != "" {
		// Retrieve class information from storage/cache using DocID
		classInfo, err = getClassInfoByID(req.DocID)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve class info: %w", err)
		}
	} else {
		return nil, fmt.Errorf("either jarPath or docID must be provided")
	}

	switch req.DocType {
	case MarkdownDoc:
		return generateMarkdownDoc(classInfo)
	case OpenAPIDoc:
		return generateOpenAPIDoc(classInfo)
	case JSONSchemaDoc:
		return generateJSONSchemaDoc(classInfo)
	default:
		return nil, fmt.Errorf("unsupported document type: %s", req.DocType)
	}
}

// generateMarkdownDoc generates a Markdown document
func generateMarkdownDoc(classInfo *core.ClassInfo) (*DocumentContent, error) {
	config := output.MarkdownConfig{
		IncludePrivateFields: false,
		IncludeMethodDetails: true,
		IncludeInheritance:   true,
	}

	content, err := output.FormatToMarkdown(classInfo, config)
	if err != nil {
		return nil, err
	}

	return &DocumentContent{
		Content:     []byte(content),
		ContentType: "text/markdown",
		FileName:    fmt.Sprintf("%s.md", sanitizeFileName(classInfo.ClassName)),
	}, nil
}

// generateOpenAPIDoc generates an OpenAPI document
func generateOpenAPIDoc(classInfo *core.ClassInfo) (*DocumentContent, error) {
	config := output.OpenAPIConfig{
		Title:   "Generated API Documentation",
		Version: "1.0.0",
	}

	openAPIObj, err := output.FormatToOpenAPI(classInfo, config)
	if err != nil {
		return nil, err
	}

	content, err := json.MarshalIndent(openAPIObj, "", "  ")
	if err != nil {
		return nil, err
	}

	return &DocumentContent{
		Content:     content,
		ContentType: "application/json",
		FileName:    fmt.Sprintf("%s-openapi.json", sanitizeFileName(classInfo.ClassName)),
	}, nil
}

// generateJSONSchemaDoc generates a JSON Schema document
func generateJSONSchemaDoc(classInfo *core.ClassInfo) (*DocumentContent, error) {
	config := output.JSONSchemaConfig{
		StrictMode: false,
	}

	schema, err := output.FormatToJSONSchema(classInfo, config)
	if err != nil {
		return nil, err
	}

	content, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		return nil, err
	}

	return &DocumentContent{
		Content:     content,
		ContentType: "application/json",
		FileName:    fmt.Sprintf("%s-schema.json", sanitizeFileName(classInfo.ClassName)),
	}, nil
}

// servePackagedDocuments creates a ZIP file containing multiple documents
func servePackagedDocuments(w http.ResponseWriter, req DownloadRequest) {
	var classInfo *core.ClassInfo
	var err error

	if req.JarPath != "" {
		classInfo, err = parseJarFile(req.JarPath)
		if err != nil {
			http.Error(w, "Failed to parse JAR file: "+err.Error(), http.StatusInternalServerError)
			return
		}
	} else if req.DocID != "" {
		classInfo, err = getClassInfoByID(req.DocID)
		if err != nil {
			http.Error(w, "Failed to retrieve class info: "+err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Either jarPath or docID must be provided", http.StatusBadRequest)
		return
	}

	// Generate all document types
	documents := make(map[string]*DocumentContent)

	// Generate Markdown
	if markdownDoc, err := generateMarkdownDoc(classInfo); err == nil {
		documents["markdown"] = markdownDoc
	}

	// Generate OpenAPI
	if openAPIDoc, err := generateOpenAPIDoc(classInfo); err == nil {
		documents["openapi"] = openAPIDoc
	}

	// Generate JSON Schema
	if jsonSchemaDoc, err := generateJSONSchemaDoc(classInfo); err == nil {
		documents["jsonschema"] = jsonSchemaDoc
	}

	if len(documents) == 0 {
		http.Error(w, "Failed to generate any documents", http.StatusInternalServerError)
		return
	}

	// Create ZIP file
	zipBuffer := new(bytes.Buffer)
	zipWriter := zip.NewWriter(zipBuffer)

	for _, doc := range documents {
		fileWriter, err := zipWriter.Create(doc.FileName)
		if err != nil {
			http.Error(w, "Failed to create ZIP file: "+err.Error(), http.StatusInternalServerError)
			return
		}
		_, err = fileWriter.Write(doc.Content)
		if err != nil {
			http.Error(w, "Failed to write to ZIP file: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	err = zipWriter.Close()
	if err != nil {
		http.Error(w, "Failed to close ZIP file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Serve ZIP file
	fileName := fmt.Sprintf("%s-docs.zip", sanitizeFileName(classInfo.ClassName))
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Length", strconv.Itoa(zipBuffer.Len()))

	_, err = io.Copy(w, zipBuffer)
	if err != nil {
		http.Error(w, "Failed to serve ZIP file: "+err.Error(), http.StatusInternalServerError)
	}
}

// serveDocument serves a single document
func serveDocument(w http.ResponseWriter, doc *DocumentContent) {
	w.Header().Set("Content-Disposition", "attachment; filename="+doc.FileName)
	w.Header().Set("Content-Type", doc.ContentType)
	w.Header().Set("Content-Length", strconv.Itoa(len(doc.Content)))

	_, err := w.Write(doc.Content)
	if err != nil {
		http.Error(w, "Failed to write document: "+err.Error(), http.StatusInternalServerError)
	}
}

// parseJarFile parses a JAR file and returns class information
func parseJarFile(jarPath string) (*core.ClassInfo, error) {
	// This is a placeholder implementation
	// In a real implementation, this would use the JAR parser from core package
	parser := &core.JarParser{}
	classInfo, err := parser.ParseJar(jarPath)
	if err != nil {
		return nil, err
	}
	return classInfo, nil
}

// getClassInfoByID retrieves class information by ID from storage/cache
func getClassInfoByID(docID string) (*core.ClassInfo, error) {
	// This is a placeholder implementation
	// In a real implementation, this would retrieve from a database or cache
	return &core.ClassInfo{
		ClassName: "SampleClass",
		PackageName: "com.example",
		Fields: []core.FieldInfo{
			{
				Name: "id",
				Type: "Long",
				Modifiers: []string{"private"},
			},
			{
				Name: "name",
				Type: "String",
				Modifiers: []string{"private"},
			},
		},
	}, nil
}

// sanitizeFileName removes invalid characters from file names
func sanitizeFileName(fileName string) string {
	// Replace invalid characters with underscores
	invalidChars := []string{"/", "\\", ":", "*", "?", "\"", "<", ">", "|"}
	result := fileName
	for _, char := range invalidChars {
		result = strings.ReplaceAll(result, char, "_")
	}
	return result
}

// HandleDocumentList returns a list of available documents
func HandleDocumentList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// This is a placeholder implementation
	// In a real implementation, this would query a database or storage system
	documents := []map[string]interface{}{
		{
			"id":          "doc1",
			"name":        "Sample Document 1",
			"jar_path":    "/path/to/sample1.jar",
			"created_at":  time.Now().Format(time.RFC3339),
			"class_name":  "SampleClass1",
			"package":     "com.example",
		},
		{
			"id":          "doc2",
			"name":        "Sample Document 2",
			"jar_path":    "/path/to/sample2.jar",
			"created_at":  time.Now().Format(time.RFC3339),
			"class_name":  "SampleClass2",
			"package":     "com.example.dto",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"documents": documents,
		"total":     len(documents),
	})
}