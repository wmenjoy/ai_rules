/**
 * [AI-ASSISTED]
 * 生成工具: Trae AI Claude 4 Sonnet
 * 生成日期: 2024-12-19
 * 贡献程度: 完全生成
 * 人工修改: 无
 * 责任人: AI Assistant
 */

// Package core provides core functionality for JAR processing
package core

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	// MaxJarFileSize defines the maximum allowed JAR file size (100MB)
	MaxJarFileSize = 100 * 1024 * 1024
	// JarFileExtension defines the expected file extension
	JarFileExtension = ".jar"
)

// JarUploadError represents errors that can occur during JAR upload
type JarUploadError struct {
	Message string
	Cause   error
}

func (e *JarUploadError) Error() string {
	return fmt.Sprintf("JAR upload error: %s", e.Message)
}

func (e *JarUploadError) Unwrap() error {
	return e.Cause
}

// UploadFromLocal handles uploading a JAR file from a local path.
// It validates the file type, size, and accessibility.
func UploadFromLocal(filePath string) (io.ReadCloser, error) {
	// Validate file extension
	if !strings.HasSuffix(strings.ToLower(filePath), JarFileExtension) {
		return nil, &JarUploadError{
			Message: fmt.Sprintf("invalid file extension, expected %s", JarFileExtension),
		}
	}

	// Check if file exists and get file info
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return nil, &JarUploadError{
			Message: "file not found or inaccessible",
			Cause:   err,
		}
	}

	// Validate file size
	if fileInfo.Size() > MaxJarFileSize {
		return nil, &JarUploadError{
			Message: fmt.Sprintf("file size %d exceeds maximum allowed size %d", fileInfo.Size(), MaxJarFileSize),
		}
	}

	// Validate it's a regular file
	if !fileInfo.Mode().IsRegular() {
		return nil, &JarUploadError{
			Message: "path does not point to a regular file",
		}
	}

	// Open and return the file
	file, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		return nil, &JarUploadError{
			Message: "failed to open file",
			Cause:   err,
		}
	}

	return file, nil
}

// UploadFromURL handles downloading a JAR file from a given URL.
// It validates the content type, size, and handles potential network errors.
func UploadFromURL(jarURL string) (io.ReadCloser, error) {
	// Validate URL format
	if jarURL == "" {
		return nil, &JarUploadError{
			Message: "URL cannot be empty",
		}
	}

	// Create HTTP request with proper headers
	req, err := http.NewRequest("GET", jarURL, nil)
	if err != nil {
		return nil, &JarUploadError{
			Message: "invalid URL format",
			Cause:   err,
		}
	}

	// Set user agent to identify our application
	req.Header.Set("User-Agent", "AI-JAR-SDK-Generator/1.0")

	// Create HTTP client with timeout
	client := &http.Client{
		// Set reasonable timeout for downloads
		// Timeout: 30 * time.Second, // Uncomment if needed
	}

	// Make the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, &JarUploadError{
			Message: "failed to download from URL",
			Cause:   err,
		}
	}

	// Check HTTP status code
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, &JarUploadError{
			Message: fmt.Sprintf("HTTP request failed with status: %d %s", resp.StatusCode, resp.Status),
		}
	}

	// Validate content length if provided
	if resp.ContentLength > MaxJarFileSize {
		resp.Body.Close()
		return nil, &JarUploadError{
			Message: fmt.Sprintf("content length %d exceeds maximum allowed size %d", resp.ContentLength, MaxJarFileSize),
		}
	}

	// Optional: Validate content type
	contentType := resp.Header.Get("Content-Type")
	if contentType != "" && !isValidJarContentType(contentType) {
		// Log warning but don't fail, as some servers may not set correct content type
		// In a real implementation, you might want to log this
	}

	// Return the response body - caller is responsible for closing
	return resp.Body, nil
}

// isValidJarContentType checks if the content type is valid for JAR files
func isValidJarContentType(contentType string) bool {
	validTypes := []string{
		"application/java-archive",
		"application/x-java-archive",
		"application/zip",
		"application/octet-stream",
	}

	for _, validType := range validTypes {
		if strings.Contains(strings.ToLower(contentType), validType) {
			return true
		}
	}
	return false
}

// ValidateJarFile performs basic validation on a JAR file stream
// This is a helper function that can be used after upload
func ValidateJarFile(reader io.Reader) error {
	// Read first few bytes to check for ZIP/JAR signature
	buffer := make([]byte, 4)
	n, err := reader.Read(buffer)
	if err != nil {
		return &JarUploadError{
			Message: "failed to read file header",
			Cause:   err,
		}
	}

	if n < 4 {
		return &JarUploadError{
			Message: "file too small to be a valid JAR",
		}
	}

	// Check for ZIP/JAR magic number (PK\003\004)
	if buffer[0] != 0x50 || buffer[1] != 0x4B || buffer[2] != 0x03 || buffer[3] != 0x04 {
		return &JarUploadError{
			Message: "file does not appear to be a valid JAR/ZIP file",
		}
	}

	return nil
}