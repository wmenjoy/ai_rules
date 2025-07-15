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
	"archive/zip"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"path/filepath"
	"regexp"
	"strings"
)

// ClassInfo represents information about a Java class
type ClassInfo struct {
	ClassName    string        `json:"className"`
	PackageName  string        `json:"packageName"`
	SuperClass   string        `json:"superClass"`
	Interfaces   []string      `json:"interfaces"`
	Methods      []MethodInfo  `json:"methods"`
	Fields       []FieldInfo   `json:"fields"`
	IsInterface  bool          `json:"isInterface"`
	IsAbstract   bool          `json:"isAbstract"`
	IsPublic     bool          `json:"isPublic"`
	SourceFile   string        `json:"sourceFile"`
}

// MethodInfo represents information about a Java method
type MethodInfo struct {
	MethodName   string   `json:"methodName"`
	ReturnType   string   `json:"returnType"`
	Parameters   []string `json:"parameters"`
	Exceptions   []string `json:"exceptions"`
	IsPublic     bool     `json:"isPublic"`
	IsPrivate    bool     `json:"isPrivate"`
	IsProtected  bool     `json:"isProtected"`
	IsStatic     bool     `json:"isStatic"`
	IsFinal      bool     `json:"isFinal"`
	IsAbstract   bool     `json:"isAbstract"`
	IsNative     bool     `json:"isNative"`
	IsSynchronized bool   `json:"isSynchronized"`
}

// FieldInfo represents information about a Java field
type FieldInfo struct {
	FieldName   string `json:"fieldName"`
	FieldType   string `json:"fieldType"`
	IsPublic    bool   `json:"isPublic"`
	IsPrivate   bool   `json:"isPrivate"`
	IsProtected bool   `json:"isProtected"`
	IsStatic    bool   `json:"isStatic"`
	IsFinal     bool   `json:"isFinal"`
	IsVolatile  bool   `json:"isVolatile"`
	IsTransient bool   `json:"isTransient"`
}

// JarInfo represents the complete information extracted from a JAR file
type JarInfo struct {
	JarName     string      `json:"jarName"`
	Classes     []ClassInfo `json:"classes"`
	Manifest    string      `json:"manifest"`
	Resources   []string    `json:"resources"`
	TotalClasses int        `json:"totalClasses"`
}

// JarParseError represents errors that can occur during JAR parsing
type JarParseError struct {
	Message string
	Cause   error
}

func (e *JarParseError) Error() string {
	return fmt.Sprintf("JAR parse error: %s", e.Message)
}

func (e *JarParseError) Unwrap() error {
	return e.Cause
}

// ParseJar parses a JAR file and extracts class information
func ParseJar(reader io.ReaderAt, size int64, jarName string) (*JarInfo, error) {
	// Open the ZIP archive
	zipReader, err := zip.NewReader(reader, size)
	if err != nil {
		return nil, &JarParseError{
			Message: "failed to open JAR as ZIP archive",
			Cause:   err,
		}
	}

	jarInfo := &JarInfo{
		JarName:   jarName,
		Classes:   make([]ClassInfo, 0),
		Resources: make([]string, 0),
	}

	// Process each file in the JAR
	for _, file := range zipReader.File {
		if err := processJarEntry(file, jarInfo); err != nil {
			// Log error but continue processing other files
			// In a real implementation, you might want to use a proper logger
			continue
		}
	}

	jarInfo.TotalClasses = len(jarInfo.Classes)
	return jarInfo, nil
}

// processJarEntry processes a single entry in the JAR file
func processJarEntry(file *zip.File, jarInfo *JarInfo) error {
	fileName := file.Name

	// Handle different file types
	switch {
	case strings.HasSuffix(fileName, ".class"):
		return processClassFile(file, jarInfo)
	case fileName == "META-INF/MANIFEST.MF":
		return processManifest(file, jarInfo)
	default:
		// Add to resources list
		jarInfo.Resources = append(jarInfo.Resources, fileName)
		return nil
	}
}

// processClassFile processes a .class file and extracts class information
func processClassFile(file *zip.File, jarInfo *JarInfo) error {
	reader, err := file.Open()
	if err != nil {
		return &JarParseError{
			Message: fmt.Sprintf("failed to open class file %s", file.Name),
			Cause:   err,
		}
	}
	defer reader.Close()

	// Read the class file content
	classData, err := io.ReadAll(reader)
	if err != nil {
		return &JarParseError{
			Message: fmt.Sprintf("failed to read class file %s", file.Name),
			Cause:   err,
		}
	}

	// Parse the class file (simplified version)
	classInfo, err := parseClassFile(classData, file.Name)
	if err != nil {
		return err
	}

	jarInfo.Classes = append(jarInfo.Classes, *classInfo)
	return nil
}

// processManifest processes the MANIFEST.MF file
func processManifest(file *zip.File, jarInfo *JarInfo) error {
	reader, err := file.Open()
	if err != nil {
		return &JarParseError{
			Message: "failed to open manifest file",
			Cause:   err,
		}
	}
	defer reader.Close()

	manifestData, err := io.ReadAll(reader)
	if err != nil {
		return &JarParseError{
			Message: "failed to read manifest file",
			Cause:   err,
		}
	}

	jarInfo.Manifest = string(manifestData)
	return nil
}

// parseClassFile parses a Java class file (simplified implementation)
// Note: This is a basic implementation. A full implementation would need
// to parse the Java class file format completely.
func parseClassFile(classData []byte, fileName string) (*ClassInfo, error) {
	if len(classData) < 10 {
		return nil, &JarParseError{
			Message: fmt.Sprintf("class file %s is too small", fileName),
		}
	}

	// Check Java class file magic number (0xCAFEBABE)
	if len(classData) < 4 || 
	   classData[0] != 0xCA || classData[1] != 0xFE || 
	   classData[2] != 0xBA || classData[3] != 0xBE {
		return nil, &JarParseError{
			Message: fmt.Sprintf("invalid class file magic number in %s", fileName),
		}
	}

	// Extract class name from file path
	className := strings.TrimSuffix(filepath.Base(fileName), ".class")
	packageName := extractPackageFromPath(fileName)

	// Create basic class info
	// Note: This is a simplified implementation
	// A complete implementation would parse the constant pool and class structure
	classInfo := &ClassInfo{
		ClassName:   className,
		PackageName: packageName,
		Methods:     make([]MethodInfo, 0),
		Fields:      make([]FieldInfo, 0),
		Interfaces:  make([]string, 0),
		SourceFile:  fileName,
	}

	// Try to extract basic information using bytecode analysis
	// This is a simplified approach - real implementation would need full bytecode parsing
	if err := extractBasicClassInfo(classData, classInfo); err != nil {
		// If extraction fails, return basic info without detailed analysis
		// In a real implementation, you might want to log this
	}

	return classInfo, nil
}

// extractPackageFromPath extracts package name from class file path
func extractPackageFromPath(filePath string) string {
	// Remove .class extension and convert path separators
	classPath := strings.TrimSuffix(filePath, ".class")
	classPath = strings.ReplaceAll(classPath, "/", ".")
	classPath = strings.ReplaceAll(classPath, "\\", ".")

	// Extract package (everything except the last component)
	parts := strings.Split(classPath, ".")
	if len(parts) <= 1 {
		return "" // Default package
	}

	return strings.Join(parts[:len(parts)-1], ".")
}

// extractBasicClassInfo attempts to extract basic class information from bytecode
// This is a simplified implementation for demonstration purposes
func extractBasicClassInfo(classData []byte, classInfo *ClassInfo) error {
	// This is a placeholder for bytecode analysis
	// A real implementation would:
	// 1. Parse the constant pool
	// 2. Extract access flags
	// 3. Parse method and field tables
	// 4. Extract method signatures and field types

	// For now, we'll add some dummy data to demonstrate the structure
	// In a real implementation, this would be replaced with actual bytecode parsing

	// Check if this looks like a valid class file structure
	if len(classData) < 20 {
		return &JarParseError{
			Message: "class file too small for analysis",
		}
	}

	// Extract version information (bytes 4-7)
	// minorVersion := binary.BigEndian.Uint16(classData[4:6])
	// majorVersion := binary.BigEndian.Uint16(classData[6:8])

	// For demonstration, set some basic flags
	// In a real implementation, these would be extracted from access_flags
	classInfo.IsPublic = true // Simplified assumption

	// Add a placeholder method to show structure
	// In a real implementation, this would be extracted from the method table
	if strings.Contains(classInfo.ClassName, "Main") {
		mainMethod := MethodInfo{
			MethodName: "main",
			ReturnType: "void",
			Parameters: []string{"String[]"},
			IsPublic:   true,
			IsStatic:   true,
		}
		classInfo.Methods = append(classInfo.Methods, mainMethod)
	}

	return nil
}

// GetClassByName finds a class by its fully qualified name
func (ji *JarInfo) GetClassByName(className string) *ClassInfo {
	for i := range ji.Classes {
		fullName := ji.Classes[i].PackageName
		if fullName != "" {
			fullName += "."
		}
		fullName += ji.Classes[i].ClassName

		if fullName == className {
			return &ji.Classes[i]
		}
	}
	return nil
}

// GetClassesByPackage returns all classes in a specific package
func (ji *JarInfo) GetClassesByPackage(packageName string) []ClassInfo {
	var result []ClassInfo
	for _, class := range ji.Classes {
		if class.PackageName == packageName {
			result = append(result, class)
		}
	}
	return result
}

// GetPublicMethods returns all public methods from all classes
func (ji *JarInfo) GetPublicMethods() []MethodInfo {
	var methods []MethodInfo
	for _, class := range ji.Classes {
		for _, method := range class.Methods {
			if method.IsPublic {
				methods = append(methods, method)
			}
		}
	}
	return methods
}

// ValidateJarStructure performs basic validation on the parsed JAR structure
func (ji *JarInfo) ValidateJarStructure() error {
	if ji.JarName == "" {
		return &JarParseError{
			Message: "JAR name is empty",
		}
	}

	if len(ji.Classes) == 0 {
		return &JarParseError{
			Message: "no classes found in JAR",
		}
	}

	// Validate class names
	classNamePattern := regexp.MustCompile(`^[a-zA-Z_$][a-zA-Z0-9_$]*$`)
	for _, class := range ji.Classes {
		if !classNamePattern.MatchString(class.ClassName) {
			return &JarParseError{
				Message: fmt.Sprintf("invalid class name: %s", class.ClassName),
			}
		}
	}

	return nil
}

// GetManifestAttribute extracts a specific attribute from the manifest
func (ji *JarInfo) GetManifestAttribute(attributeName string) string {
	if ji.Manifest == "" {
		return ""
	}

	scanner := bufio.NewScanner(strings.NewReader(ji.Manifest))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, attributeName+":") {
			return strings.TrimSpace(strings.TrimPrefix(line, attributeName+":"))
		}
	}
	return ""
}

// ParseJarFromBytes is a convenience function to parse JAR from byte slice
func ParseJarFromBytes(data []byte, jarName string) (*JarInfo, error) {
	reader := bytes.NewReader(data)
	return ParseJar(reader, int64(len(data)), jarName)
}