package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleDownloadDoc_GET(t *testing.T) {
	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedHeader map[string]string
	}{
		{
			name:           "Valid markdown download",
			queryParams:    "type=markdown&docId=test123",
			expectedStatus: http.StatusOK,
			expectedHeader: map[string]string{
				"Content-Type":        "text/markdown",
				"Content-Disposition": "attachment; filename=SampleClass.md",
			},
		},
		{
			name:           "Valid OpenAPI download",
			queryParams:    "type=openapi_json&docId=test123",
			expectedStatus: http.StatusOK,
			expectedHeader: map[string]string{
				"Content-Type":        "application/json",
				"Content-Disposition": "attachment; filename=SampleClass-openapi.json",
			},
		},
		{
			name:           "Valid JSON Schema download",
			queryParams:    "type=jsonschema&docId=test123",
			expectedStatus: http.StatusOK,
			expectedHeader: map[string]string{
				"Content-Type":        "application/json",
				"Content-Disposition": "attachment; filename=SampleClass-schema.json",
			},
		},
		{
			name:           "Packaged download",
			queryParams:    "type=all&docId=test123&packaged=true",
			expectedStatus: http.StatusOK,
			expectedHeader: map[string]string{
				"Content-Type":        "application/zip",
				"Content-Disposition": "attachment; filename=SampleClass-docs.zip",
			},
		},
		{
			name:           "Invalid document type",
			queryParams:    "type=invalid&docId=test123",
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "Missing parameters",
			queryParams:    "",
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/download?"+tt.queryParams, nil)
			w := httptest.NewRecorder()

			HandleDownloadDoc(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			if tt.expectedStatus == http.StatusOK {
				for header, expectedValue := range tt.expectedHeader {
					actualValue := w.Header().Get(header)
					if actualValue != expectedValue {
						t.Errorf("Expected header %s: %s, got %s", header, expectedValue, actualValue)
					}
				}
			}
		})
	}
}

func TestHandleDownloadDoc_POST(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    DownloadRequest
		expectedStatus int
		expectedHeader map[string]string
	}{
		{
			name: "Valid POST request for markdown",
			requestBody: DownloadRequest{
				DocID:   "test123",
				DocType: MarkdownDoc,
			},
			expectedStatus: http.StatusOK,
			expectedHeader: map[string]string{
				"Content-Type":        "text/markdown",
				"Content-Disposition": "attachment; filename=SampleClass.md",
			},
		},
		{
			name: "Valid POST request for packaged docs",
			requestBody: DownloadRequest{
				DocID:    "test123",
				DocType:  AllDocs,
				Packaged: true,
			},
			expectedStatus: http.StatusOK,
			expectedHeader: map[string]string{
				"Content-Type":        "application/zip",
				"Content-Disposition": "attachment; filename=SampleClass-docs.zip",
			},
		},
		{
			name: "Valid POST request with JAR path",
			requestBody: DownloadRequest{
				JarPath: "/path/to/test.jar",
				DocType: JSONSchemaDoc,
			},
			expectedStatus: http.StatusInternalServerError, // Will fail due to invalid JAR path
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest(http.MethodPost, "/download", bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			HandleDownloadDoc(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			if tt.expectedStatus == http.StatusOK {
				for header, expectedValue := range tt.expectedHeader {
					actualValue := w.Header().Get(header)
					if actualValue != expectedValue {
						t.Errorf("Expected header %s: %s, got %s", header, expectedValue, actualValue)
					}
				}
			}
		})
	}
}

func TestHandleDownloadDoc_InvalidMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodPut, "/download", nil)
	w := httptest.NewRecorder()

	HandleDownloadDoc(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status %d, got %d", http.StatusMethodNotAllowed, w.Code)
	}
}

func TestHandleDownloadDoc_InvalidJSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/download", strings.NewReader("invalid json"))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	HandleDownloadDoc(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestGenerateDocument(t *testing.T) {
	tests := []struct {
		name        string
		request     DownloadRequest
		expectError bool
		expectedExt string
	}{
		{
			name: "Generate markdown document",
			request: DownloadRequest{
				DocID:   "test123",
				DocType: MarkdownDoc,
			},
			expectError: false,
			expectedExt: ".md",
		},
		{
			name: "Generate OpenAPI document",
			request: DownloadRequest{
				DocID:   "test123",
				DocType: OpenAPIDoc,
			},
			expectError: false,
			expectedExt: ".json",
		},
		{
			name: "Generate JSON Schema document",
			request: DownloadRequest{
				DocID:   "test123",
				DocType: JSONSchemaDoc,
			},
			expectError: false,
			expectedExt: ".json",
		},
		{
			name: "Invalid document type",
			request: DownloadRequest{
				DocID:   "test123",
				DocType: "invalid",
			},
			expectError: true,
		},
		{
			name: "Missing DocID and JarPath",
			request: DownloadRequest{
				DocType: MarkdownDoc,
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := generateDocument(tt.request)

			if tt.expectError {
				if err == nil {
					t.Error("Expected error, but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if doc == nil {
				t.Error("Expected document, but got nil")
				return
			}

			if !strings.HasSuffix(doc.FileName, tt.expectedExt) {
				t.Errorf("Expected file extension %s, but got %s", tt.expectedExt, doc.FileName)
			}

			if len(doc.Content) == 0 {
				t.Error("Expected non-empty content")
			}
		})
	}
}

func TestSanitizeFileName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "ValidFileName",
			expected: "ValidFileName",
		},
		{
			input:    "File/With\\Invalid:Characters",
			expected: "File_With_Invalid_Characters",
		},
		{
			input:    "File*With?Special<>Characters|",
			expected: "File_With_Special__Characters_",
		},
		{
			input:    "File\"With\"Quotes",
			expected: "File_With_Quotes",
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := sanitizeFileName(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestHandleDocumentList(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		expectedStatus int
	}{
		{
			name:           "Valid GET request",
			method:         http.MethodGet,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Invalid POST request",
			method:         http.MethodPost,
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "Invalid PUT request",
			method:         http.MethodPut,
			expectedStatus: http.StatusMethodNotAllowed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/documents", nil)
			w := httptest.NewRecorder()

			HandleDocumentList(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			if tt.expectedStatus == http.StatusOK {
				contentType := w.Header().Get("Content-Type")
				if contentType != "application/json" {
					t.Errorf("Expected Content-Type application/json, got %s", contentType)
				}

				// Verify response structure
				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				if err != nil {
					t.Errorf("Failed to parse JSON response: %v", err)
				}

				if _, exists := response["documents"]; !exists {
					t.Error("Response missing 'documents' field")
				}

				if _, exists := response["total"]; !exists {
					t.Error("Response missing 'total' field")
				}
			}
		})
	}
}

func TestDocumentTypes(t *testing.T) {
	// Test DocumentType constants
	if MarkdownDoc != "markdown" {
		t.Errorf("Expected MarkdownDoc to be 'markdown', got %s", MarkdownDoc)
	}
	if OpenAPIDoc != "openapi_json" {
		t.Errorf("Expected OpenAPIDoc to be 'openapi_json', got %s", OpenAPIDoc)
	}
	if JSONSchemaDoc != "jsonschema" {
		t.Errorf("Expected JSONSchemaDoc to be 'jsonschema', got %s", JSONSchemaDoc)
	}
	if AllDocs != "all" {
		t.Errorf("Expected AllDocs to be 'all', got %s", AllDocs)
	}
}

func TestDownloadRequestStruct(t *testing.T) {
	// Test DownloadRequest struct serialization
	req := DownloadRequest{
		DocID:    "test123",
		DocType:  MarkdownDoc,
		JarPath:  "/path/to/test.jar",
		Packaged: true,
	}

	data, err := json.Marshal(req)
	if err != nil {
		t.Errorf("Failed to marshal DownloadRequest: %v", err)
	}

	var unmarshaled DownloadRequest
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Errorf("Failed to unmarshal DownloadRequest: %v", err)
	}

	if unmarshaled.DocID != req.DocID {
		t.Errorf("DocID mismatch: expected %s, got %s", req.DocID, unmarshaled.DocID)
	}
	if unmarshaled.DocType != req.DocType {
		t.Errorf("DocType mismatch: expected %s, got %s", req.DocType, unmarshaled.DocType)
	}
	if unmarshaled.JarPath != req.JarPath {
		t.Errorf("JarPath mismatch: expected %s, got %s", req.JarPath, unmarshaled.JarPath)
	}
	if unmarshaled.Packaged != req.Packaged {
		t.Errorf("Packaged mismatch: expected %t, got %t", req.Packaged, unmarshaled.Packaged)
	}
}

func TestDocumentContentStruct(t *testing.T) {
	// Test DocumentContent struct
	content := &DocumentContent{
		Content:     []byte("test content"),
		ContentType: "text/plain",
		FileName:    "test.txt",
	}

	if string(content.Content) != "test content" {
		t.Errorf("Content mismatch: expected 'test content', got %s", string(content.Content))
	}
	if content.ContentType != "text/plain" {
		t.Errorf("ContentType mismatch: expected 'text/plain', got %s", content.ContentType)
	}
	if content.FileName != "test.txt" {
		t.Errorf("FileName mismatch: expected 'test.txt', got %s", content.FileName)
	}
}

// Benchmark tests
func BenchmarkSanitizeFileName(b *testing.B) {
	fileName := "File/With\\Invalid:Characters*?<>|\""
	for i := 0; i < b.N; i++ {
		sanitizeFileName(fileName)
	}
}

func BenchmarkGenerateDocument(b *testing.B) {
	req := DownloadRequest{
		DocID:   "test123",
		DocType: MarkdownDoc,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = generateDocument(req)
	}
}
