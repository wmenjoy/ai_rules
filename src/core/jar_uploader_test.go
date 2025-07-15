/**
 * [AI-ASSISTED]
 * 生成工具: Trae AI Claude 4 Sonnet
 * 生成日期: 2024-12-19
 * 贡献程度: 完全生成
 * 人工修改: 无
 * 责任人: AI Assistant
 */

package core

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestUploadFromLocal tests the local file upload functionality
func TestUploadFromLocal(t *testing.T) {
	tests := []struct {
		name        string
		filePath    string
		setupFunc   func() (string, func()) // returns filepath and cleanup function
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid jar file",
			setupFunc: func() (string, func()) {
				// Create a temporary JAR file with valid ZIP header
				tmpDir := t.TempDir()
				jarPath := filepath.Join(tmpDir, "test.jar")

				// Create file with ZIP/JAR magic number
				jarContent := []byte{0x50, 0x4B, 0x03, 0x04}          // PK\003\004
				jarContent = append(jarContent, make([]byte, 100)...) // Add some content

				err := os.WriteFile(jarPath, jarContent, 0644)
				if err != nil {
					t.Fatalf("Failed to create test JAR file: %v", err)
				}

				return jarPath, func() {} // cleanup handled by t.TempDir()
			},
			expectError: false,
		},
		{
			name:        "non-existent file",
			filePath:    "/non/existent/file.jar",
			expectError: true,
			errorMsg:    "file not found or inaccessible",
		},
		{
			name:        "invalid extension",
			filePath:    "test.txt",
			expectError: true,
			errorMsg:    "invalid file extension",
		},
		{
			name: "file too large",
			setupFunc: func() (string, func()) {
				tmpDir := t.TempDir()
				jarPath := filepath.Join(tmpDir, "large.jar")

				// Create a file larger than MaxJarFileSize
				largeContent := make([]byte, MaxJarFileSize+1)
				err := os.WriteFile(jarPath, largeContent, 0644)
				if err != nil {
					t.Fatalf("Failed to create large test file: %v", err)
				}

				return jarPath, func() {}
			},
			expectError: true,
			errorMsg:    "exceeds maximum allowed size",
		},
		{
			name: "directory instead of file",
			setupFunc: func() (string, func()) {
				tmpDir := t.TempDir()
				dirPath := filepath.Join(tmpDir, "test.jar")
				err := os.Mkdir(dirPath, 0755)
				if err != nil {
					t.Fatalf("Failed to create test directory: %v", err)
				}
				return dirPath, func() {}
			},
			expectError: true,
			errorMsg:    "does not point to a regular file",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var filePath string
			var cleanup func()

			if tt.setupFunc != nil {
				filePath, cleanup = tt.setupFunc()
				defer cleanup()
			} else {
				filePath = tt.filePath
			}

			reader, err := UploadFromLocal(filePath)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
					if reader != nil {
						reader.Close()
					}
					return
				}
				if !strings.Contains(err.Error(), tt.errorMsg) {
					t.Errorf("Expected error message to contain '%s', got: %v", tt.errorMsg, err)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
					return
				}
				if reader == nil {
					t.Errorf("Expected reader but got nil")
					return
				}
				reader.Close()
			}
		})
	}
}

// TestUploadFromURL tests the URL download functionality
func TestUploadFromURL(t *testing.T) {
	tests := []struct {
		name        string
		setupServer func() *httptest.Server
		url         string
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid jar from URL",
			setupServer: func() *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Header().Set("Content-Type", "application/java-archive")
					w.Header().Set("Content-Length", "100")

					// Write valid JAR header
					jarContent := []byte{0x50, 0x4B, 0x03, 0x04}         // PK\003\004
					jarContent = append(jarContent, make([]byte, 96)...) // Total 100 bytes
					w.Write(jarContent)
				}))
			},
			expectError: false,
		},
		{
			name: "404 not found",
			setupServer: func() *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusNotFound)
				}))
			},
			expectError: true,
			errorMsg:    "HTTP request failed with status: 404",
		},
		{
			name: "content too large",
			setupServer: func() *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Header().Set("Content-Length", fmt.Sprintf("%d", MaxJarFileSize+1))
					w.WriteHeader(http.StatusOK)
				}))
			},
			expectError: true,
			errorMsg:    "exceeds maximum allowed size",
		},
		{
			name:        "empty URL",
			url:         "",
			expectError: true,
			errorMsg:    "URL cannot be empty",
		},
		{
			name:        "invalid URL",
			url:         "not-a-valid-url",
			expectError: true,
			errorMsg:    "failed to download from URL",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var server *httptest.Server
			var url string

			if tt.setupServer != nil {
				server = tt.setupServer()
				defer server.Close()
				url = server.URL
			} else {
				url = tt.url
			}

			reader, err := UploadFromURL(url)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
					if reader != nil {
						reader.Close()
					}
					return
				}
				if !strings.Contains(err.Error(), tt.errorMsg) {
					t.Errorf("Expected error message to contain '%s', got: %v", tt.errorMsg, err)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
					return
				}
				if reader == nil {
					t.Errorf("Expected reader but got nil")
					return
				}
				reader.Close()
			}
		})
	}
}

// TestValidateJarFile tests JAR file validation
func TestValidateJarFile(t *testing.T) {
	tests := []struct {
		name        string
		content     []byte
		expectError bool
		errorMsg    string
	}{
		{
			name:        "valid jar header",
			content:     []byte{0x50, 0x4B, 0x03, 0x04, 0x00, 0x00}, // PK\003\004 + extra
			expectError: false,
		},
		{
			name:        "invalid header",
			content:     []byte{0x00, 0x00, 0x00, 0x00},
			expectError: true,
			errorMsg:    "does not appear to be a valid JAR/ZIP file",
		},
		{
			name:        "file too small",
			content:     []byte{0x50, 0x4B},
			expectError: true,
			errorMsg:    "file too small to be a valid JAR",
		},
		{
			name:        "empty file",
			content:     []byte{},
			expectError: true,
			errorMsg:    "failed to read file header",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := bytes.NewReader(tt.content)
			err := ValidateJarFile(reader)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
					return
				}
				if !strings.Contains(err.Error(), tt.errorMsg) {
					t.Errorf("Expected error message to contain '%s', got: %v", tt.errorMsg, err)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		})
	}
}

// TestIsValidJarContentType tests content type validation
func TestIsValidJarContentType(t *testing.T) {
	tests := []struct {
		name        string
		contentType string
		expected    bool
	}{
		{"java archive", "application/java-archive", true},
		{"x-java archive", "application/x-java-archive", true},
		{"zip", "application/zip", true},
		{"octet stream", "application/octet-stream", true},
		{"with charset", "application/java-archive; charset=utf-8", true},
		{"text plain", "text/plain", false},
		{"empty", "", false},
		{"image", "image/jpeg", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidJarContentType(tt.contentType)
			if result != tt.expected {
				t.Errorf("isValidJarContentType(%q) = %v, expected %v", tt.contentType, result, tt.expected)
			}
		})
	}
}

// TestJarUploadError tests custom error type
func TestJarUploadError(t *testing.T) {
	originalErr := fmt.Errorf("original error")
	uploadErr := &JarUploadError{
		Message: "test message",
		Cause:   originalErr,
	}

	// Test Error() method
	expectedMsg := "JAR upload error: test message"
	if uploadErr.Error() != expectedMsg {
		t.Errorf("Error() = %q, expected %q", uploadErr.Error(), expectedMsg)
	}

	// Test Unwrap() method
	if uploadErr.Unwrap() != originalErr {
		t.Errorf("Unwrap() = %v, expected %v", uploadErr.Unwrap(), originalErr)
	}
}

// BenchmarkUploadFromLocal benchmarks local file upload
func BenchmarkUploadFromLocal(b *testing.B) {
	// Create a temporary JAR file
	tmpDir := b.TempDir()
	jarPath := filepath.Join(tmpDir, "benchmark.jar")

	// Create file with valid content
	jarContent := []byte{0x50, 0x4B, 0x03, 0x04}           // PK\003\004
	jarContent = append(jarContent, make([]byte, 1000)...) // Add some content

	err := os.WriteFile(jarPath, jarContent, 0644)
	if err != nil {
		b.Fatalf("Failed to create benchmark JAR file: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reader, err := UploadFromLocal(jarPath)
		if err != nil {
			b.Fatalf("Unexpected error: %v", err)
		}
		reader.Close()
	}
}
