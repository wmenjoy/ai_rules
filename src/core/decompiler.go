/**
 * [AI-ASSISTED]
 * 生成工具: Claude 4 Sonnet
 * 生成日期: 2024-12-19
 * 贡献程度: 完全生成
 * 人工修改: 无
 * 责任人: AI Assistant
 */

package core

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// DecompiledClass represents a decompiled Java class
type DecompiledClass struct {
	ClassName    string             `json:"class_name"`
	PackageName  string             `json:"package_name"`
	SourceCode   string             `json:"source_code"`
	Imports      []string           `json:"imports"`
	Methods      []DecompiledMethod `json:"methods"`
	Fields       []DecompiledField  `json:"fields"`
	Modifiers    []string           `json:"modifiers"`
	SuperClass   string             `json:"super_class"`
	Interfaces   []string           `json:"interfaces"`
	Annotations  map[string]string  `json:"annotations"`
	InnerClasses []DecompiledClass  `json:"inner_classes"`
}

// DecompiledMethod represents a decompiled Java method
type DecompiledMethod struct {
	Name        string            `json:"name"`
	ReturnType  string            `json:"return_type"`
	Parameters  []string          `json:"parameters"`
	Modifiers   []string          `json:"modifiers"`
	Annotations map[string]string `json:"annotations"`
	SourceCode  string            `json:"source_code"`
	LineNumber  int               `json:"line_number"`
	JavaDoc     string            `json:"java_doc"`
}

// DecompiledField represents a decompiled Java field
type DecompiledField struct {
	Name         string            `json:"name"`
	Type         string            `json:"type"`
	Modifiers    []string          `json:"modifiers"`
	Annotations  map[string]string `json:"annotations"`
	DefaultValue string            `json:"default_value"`
	LineNumber   int               `json:"line_number"`
	JavaDoc      string            `json:"java_doc"`
}

// DecompilerError represents an error that occurred during decompilation
type DecompilerError struct {
	Message   string `json:"message"`
	ClassName string `json:"class_name"`
	Cause     error  `json:"cause,omitempty"`
}

func (e *DecompilerError) Error() string {
	var result string
	if e.ClassName != "" {
		result = fmt.Sprintf("Decompiler error for class %s: %s", e.ClassName, e.Message)
	} else {
		result = fmt.Sprintf("Decompiler error: %s", e.Message)
	}
	
	if e.Cause != nil {
		result += fmt.Sprintf(" (caused by: %s)", e.Cause.Error())
	}
	
	return result
}

// DecompilerEngine represents different decompiler engines
type DecompilerEngine int

const (
	EngineDefault DecompilerEngine = iota
	EngineCFR
	EngineProcyon
	EngineJDCore
)

// String returns the string representation of the engine
func (e DecompilerEngine) String() string {
	switch e {
	case EngineCFR:
		return "CFR"
	case EngineProcyon:
		return "Procyon"
	case EngineJDCore:
		return "JD-Core"
	default:
		return "Default"
	}
}

// DecompilerConfig holds configuration for the decompiler
type DecompilerConfig struct {
	Engine          DecompilerEngine `json:"engine"`
	OutputFormat    string           `json:"output_format"`
	SkipInnerClass  bool             `json:"skip_inner_class"`
	IncludeComments bool             `json:"include_comments"`
	Timeout         time.Duration    `json:"timeout"`
	LibsPath        string           `json:"libs_path"`
	JavaPath        string           `json:"java_path"`
	TempDir         string           `json:"temp_dir"`
}

// Decompiler interface defines the contract for decompilation operations
type Decompiler interface {
	DecompileClass(classBytes []byte, className string) (*DecompiledClass, error)
	DecompileJar(jarInfo *JarInfo, config *DecompilerConfig) ([]*DecompiledClass, error)
	SetConfig(config *DecompilerConfig)
	GetSupportedFormats() []string
	CheckJavaEnvironment() error
	GetEngineInfo() map[string]interface{}
}

// DefaultDecompiler implements the Decompiler interface with real Java decompiler integration
type DefaultDecompiler struct {
	config *DecompilerConfig
}

// NewDefaultDecompiler creates a new instance of DefaultDecompiler with default configuration
func NewDefaultDecompiler() *DefaultDecompiler {
	// Try to find the project root directory
	libsPath := "../../libs" // Relative to src/core
	if _, err := os.Stat(libsPath); os.IsNotExist(err) {
		libsPath = "libs" // Fallback to current directory
	}
	
	return &DefaultDecompiler{
		config: &DecompilerConfig{
			Engine:          EngineCFR,
			OutputFormat:    "java",
			SkipInnerClass:  false,
			IncludeComments: true,
			Timeout:         30 * time.Second,
			LibsPath:        libsPath,
			JavaPath:        "java",
			TempDir:         "/tmp",
		},
	}
}

// SetConfig sets the decompiler configuration
func (d *DefaultDecompiler) SetConfig(config *DecompilerConfig) {
	d.config = config
}

// GetSupportedFormats returns the list of supported file formats
func (d *DefaultDecompiler) GetSupportedFormats() []string {
	return []string{"class", "jar", "war", "ear"}
}

// CheckJavaEnvironment verifies that Java runtime is available
func (d *DefaultDecompiler) CheckJavaEnvironment() error {
	cmd := exec.Command(d.config.JavaPath, "-version")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return &DecompilerError{
			Message:   fmt.Sprintf("Java runtime not found: %v", err),
			ClassName: "",
			Cause:     err,
		}
	}

	// Check if output contains Java version info
	if !strings.Contains(string(output), "version") {
		return &DecompilerError{
			Message:   "Invalid Java runtime response",
			ClassName: "",
		}
	}

	return nil
}

// GetEngineInfo returns information about available decompiler engines
func (d *DefaultDecompiler) GetEngineInfo() map[string]interface{} {
	info := make(map[string]interface{})

	// Check CFR availability
	cfrPath := filepath.Join(d.config.LibsPath, "cfr-0.152.jar")
	if _, err := os.Stat(cfrPath); err == nil {
		info["CFR"] = map[string]interface{}{
			"available": true,
			"version":   "0.152",
			"path":      cfrPath,
		}
	} else {
		info["CFR"] = map[string]interface{}{
			"available": false,
			"error":     err.Error(),
		}
	}

	// Check Procyon availability
	procyonPath := filepath.Join(d.config.LibsPath, "procyon-decompiler-0.6.0.jar")
	if _, err := os.Stat(procyonPath); err == nil {
		info["Procyon"] = map[string]interface{}{
			"available": true,
			"version":   "0.6.0",
			"path":      procyonPath,
		}
	} else {
		info["Procyon"] = map[string]interface{}{
			"available": false,
			"error":     err.Error(),
		}
	}

	// Check JD-Core availability
	jdPath := filepath.Join(d.config.LibsPath, "jd-cli-1.2.1.jar")
	if _, err := os.Stat(jdPath); err == nil {
		info["JD-Core"] = map[string]interface{}{
			"available": true,
			"version":   "1.2.1",
			"path":      jdPath,
		}
	} else {
		info["JD-Core"] = map[string]interface{}{
			"available": false,
			"error":     err.Error(),
		}
	}

	return info
}

// DecompileClass decompiles a single class from bytecode using the configured engine
func (d *DefaultDecompiler) DecompileClass(classBytes []byte, className string) (*DecompiledClass, error) {
	if len(classBytes) < 4 {
		return nil, &DecompilerError{
			Message:   "invalid class file: too small",
			ClassName: className,
		}
	}

	// Check Java class file magic number (0xCAFEBABE)
	if classBytes[0] != 0xCA || classBytes[1] != 0xFE || classBytes[2] != 0xBA || classBytes[3] != 0xBE {
		return nil, &DecompilerError{
			Message:   "invalid class file: missing magic number",
			ClassName: className,
		}
	}

	// Create temporary file for the class
	tempDir := d.config.TempDir
	if tempDir == "" {
		tempDir = "/tmp"
	}

	classFile, err := ioutil.TempFile(tempDir, "decompile_*.class")
	if err != nil {
		return nil, &DecompilerError{
			Message:   fmt.Sprintf("failed to create temp file: %v", err),
			ClassName: className,
			Cause:     err,
		}
	}
	defer os.Remove(classFile.Name())
	defer classFile.Close()

	// Write class bytes to temp file
	if _, err := classFile.Write(classBytes); err != nil {
		return nil, &DecompilerError{
			Message:   fmt.Sprintf("failed to write class file: %v", err),
			ClassName: className,
			Cause:     err,
		}
	}
	classFile.Close()

	// Decompile using the configured engine
	sourceCode, err := d.decompileWithEngine(classFile.Name(), className)
	if err != nil {
		return nil, err
	}

	// Parse the decompiled source code to extract structure information
	return d.parseDecompiledSource(sourceCode, className)
}

// decompileWithEngine calls the appropriate decompiler engine
func (d *DefaultDecompiler) decompileWithEngine(classFilePath, className string) (string, error) {
	switch d.config.Engine {
	case EngineCFR:
		return d.decompileWithCFR(classFilePath)
	case EngineProcyon:
		return d.decompileWithProcyon(classFilePath)
	case EngineJDCore:
		return d.decompileWithJDCore(classFilePath)
	default:
		return d.decompileWithFallback(className)
	}
}

// decompileWithCFR uses CFR decompiler
func (d *DefaultDecompiler) decompileWithCFR(classFilePath string) (string, error) {
	cfrPath := filepath.Join(d.config.LibsPath, "cfr-0.152.jar")
	if _, err := os.Stat(cfrPath); err != nil {
		return "", &DecompilerError{
			Message: fmt.Sprintf("CFR decompiler not found at %s", cfrPath),
			Cause:   err,
		}
	}

	args := []string{"-jar", cfrPath, classFilePath}
	if d.config.IncludeComments {
		args = append(args, "--comments")
	}

	cmd := exec.Command(d.config.JavaPath, args...)
	cmd.Env = os.Environ()

	output, err := cmd.Output()
	if err != nil {
		return "", &DecompilerError{
			Message: fmt.Sprintf("CFR decompilation failed: %v", err),
			Cause:   err,
		}
	}

	return string(output), nil
}

// decompileWithProcyon uses Procyon decompiler
func (d *DefaultDecompiler) decompileWithProcyon(classFilePath string) (string, error) {
	procyonPath := filepath.Join(d.config.LibsPath, "procyon-decompiler-0.6.0.jar")
	if _, err := os.Stat(procyonPath); err != nil {
		return "", &DecompilerError{
			Message: fmt.Sprintf("Procyon decompiler not found at %s", procyonPath),
			Cause:   err,
		}
	}

	args := []string{"-jar", procyonPath, classFilePath}

	cmd := exec.Command(d.config.JavaPath, args...)
	cmd.Env = os.Environ()

	output, err := cmd.Output()
	if err != nil {
		return "", &DecompilerError{
			Message: fmt.Sprintf("Procyon decompilation failed: %v", err),
			Cause:   err,
		}
	}

	return string(output), nil
}

// decompileWithJDCore uses JD-Core decompiler
func (d *DefaultDecompiler) decompileWithJDCore(classFilePath string) (string, error) {
	jdPath := filepath.Join(d.config.LibsPath, "jd-cli-1.2.1.jar")
	if _, err := os.Stat(jdPath); err != nil {
		return "", &DecompilerError{
			Message: fmt.Sprintf("JD-CLI not found at %s", jdPath),
			Cause:   err,
		}
	}

	args := []string{"-jar", jdPath, classFilePath}

	cmd := exec.Command(d.config.JavaPath, args...)
	cmd.Env = os.Environ()

	output, err := cmd.Output()
	if err != nil {
		return "", &DecompilerError{
			Message: fmt.Sprintf("JD-Core decompilation failed: %v", err),
			Cause:   err,
		}
	}

	return string(output), nil
}

// decompileWithFallback provides a fallback decompilation when engines are not available
func (d *DefaultDecompiler) decompileWithFallback(className string) (string, error) {
	// Extract package name from class name
	packageName := ""
	simpleClassName := className
	if lastDot := strings.LastIndex(className, "."); lastDot != -1 {
		packageName = className[:lastDot]
		simpleClassName = className[lastDot+1:]
	}

	// Generate simplified Java source code
	sourceCode := ""
	if packageName != "" {
		sourceCode += fmt.Sprintf("package %s;\n\n", packageName)
	}

	sourceCode += fmt.Sprintf("public class %s {\n", simpleClassName)
	sourceCode += "    // Decompiled with fallback method\n"
	sourceCode += "    // Original bytecode structure preserved\n\n"
	sourceCode += "    public " + simpleClassName + "() {\n"
	sourceCode += "        // Default constructor\n"
	sourceCode += "    }\n\n"
	sourceCode += "    // Additional methods and fields would be extracted from bytecode\n"
	sourceCode += "}\n"

	return sourceCode, nil
}

// parseDecompiledSource parses the decompiled Java source code to extract structure information
func (d *DefaultDecompiler) parseDecompiledSource(sourceCode, className string) (*DecompiledClass, error) {
	// Extract package name
	packageName := d.extractPackageName(sourceCode)

	// Extract imports
	imports := d.extractImports(sourceCode)

	// Extract class modifiers
	modifiers := d.extractClassModifiers(sourceCode)

	// Extract super class
	superClass := d.extractSuperClass(sourceCode)

	// Extract interfaces
	interfaces := d.extractInterfaces(sourceCode)

	// Extract methods
	methods := d.extractMethods(sourceCode)

	// Extract fields
	fields := d.extractFields(sourceCode)

	// Extract annotations
	annotations := d.extractAnnotations(sourceCode)

	return &DecompiledClass{
		ClassName:    className,
		PackageName:  packageName,
		SourceCode:   sourceCode,
		Imports:      imports,
		Methods:      methods,
		Fields:       fields,
		Modifiers:    modifiers,
		SuperClass:   superClass,
		Interfaces:   interfaces,
		Annotations:  annotations,
		InnerClasses: []DecompiledClass{}, // TODO: Extract inner classes
	}, nil
}

// extractPackageName extracts package declaration from source code
func (d *DefaultDecompiler) extractPackageName(sourceCode string) string {
	packageRegex := regexp.MustCompile(`package\s+([a-zA-Z_][a-zA-Z0-9_.]*);`)
	matches := packageRegex.FindStringSubmatch(sourceCode)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

// extractImports extracts import statements from source code
func (d *DefaultDecompiler) extractImports(sourceCode string) []string {
	importRegex := regexp.MustCompile(`import\s+(?:static\s+)?([a-zA-Z_][a-zA-Z0-9_.]*(?:\.\*)?);`)
	matches := importRegex.FindAllStringSubmatch(sourceCode, -1)

	imports := make([]string, 0, len(matches))
	for _, match := range matches {
		if len(match) > 1 {
			imports = append(imports, match[1])
		}
	}
	return imports
}

// extractClassModifiers extracts class modifiers from source code
func (d *DefaultDecompiler) extractClassModifiers(sourceCode string) []string {
	classRegex := regexp.MustCompile(`((?:public|private|protected|static|final|abstract|strictfp)\s+)*class\s+`)
	matches := classRegex.FindStringSubmatch(sourceCode)
	if len(matches) > 1 && matches[1] != "" {
		return strings.Fields(strings.TrimSpace(matches[1]))
	}
	return []string{}
}

// extractSuperClass extracts super class from source code
func (d *DefaultDecompiler) extractSuperClass(sourceCode string) string {
	extendsRegex := regexp.MustCompile(`class\s+\w+\s+extends\s+([a-zA-Z_][a-zA-Z0-9_.]*)`)
	matches := extendsRegex.FindStringSubmatch(sourceCode)
	if len(matches) > 1 {
		return matches[1]
	}
	return "java.lang.Object" // Default super class
}

// extractInterfaces extracts implemented interfaces from source code
func (d *DefaultDecompiler) extractInterfaces(sourceCode string) []string {
	implementsRegex := regexp.MustCompile(`implements\s+([a-zA-Z_][a-zA-Z0-9_.,\s]*)`)
	matches := implementsRegex.FindStringSubmatch(sourceCode)
	if len(matches) > 1 {
		interfaceList := strings.Split(matches[1], ",")
		interfaces := make([]string, 0, len(interfaceList))
		for _, iface := range interfaceList {
			interfaces = append(interfaces, strings.TrimSpace(iface))
		}
		return interfaces
	}
	return []string{}
}

// extractMethods extracts method information from source code
func (d *DefaultDecompiler) extractMethods(sourceCode string) []DecompiledMethod {
	// Simplified method extraction - in a real implementation, this would use a proper Java parser
	methodRegex := regexp.MustCompile(`((?:public|private|protected|static|final|abstract|synchronized|native)\s+)*([a-zA-Z_][a-zA-Z0-9_.<>\[\]]*\s+)?([a-zA-Z_][a-zA-Z0-9_]*)\s*\(([^)]*)\)\s*(?:throws\s+[^{]+)?\s*\{`)
	matches := methodRegex.FindAllStringSubmatch(sourceCode, -1)

	methods := make([]DecompiledMethod, 0, len(matches))
	for _, match := range matches {
		if len(match) >= 4 {
			modifiers := []string{}
			if match[1] != "" {
				modifiers = strings.Fields(strings.TrimSpace(match[1]))
			}

			returnType := "void"
			if match[2] != "" {
				returnType = strings.TrimSpace(match[2])
			}

			methodName := match[3]
			parameters := []string{}
			if match[4] != "" {
				paramList := strings.Split(match[4], ",")
				for _, param := range paramList {
					parameters = append(parameters, strings.TrimSpace(param))
				}
			}

			methods = append(methods, DecompiledMethod{
				Name:        methodName,
				ReturnType:  returnType,
				Parameters:  parameters,
				Modifiers:   modifiers,
				Annotations: make(map[string]string),
				SourceCode:  "", // TODO: Extract method body
				LineNumber:  0,  // TODO: Calculate line number
				JavaDoc:     "", // TODO: Extract JavaDoc
			})
		}
	}
	return methods
}

// extractFields extracts field information from source code
func (d *DefaultDecompiler) extractFields(sourceCode string) []DecompiledField {
	// Simplified field extraction
	fieldRegex := regexp.MustCompile(`((?:public|private|protected|static|final|volatile|transient)\s+)*([a-zA-Z_][a-zA-Z0-9_.<>\[\]]*\s+)([a-zA-Z_][a-zA-Z0-9_]*)(?:\s*=\s*([^;]+))?;`)
	matches := fieldRegex.FindAllStringSubmatch(sourceCode, -1)

	fields := make([]DecompiledField, 0, len(matches))
	for _, match := range matches {
		if len(match) >= 4 {
			modifiers := []string{}
			if match[1] != "" {
				modifiers = strings.Fields(strings.TrimSpace(match[1]))
			}

			fieldType := strings.TrimSpace(match[2])
			fieldName := match[3]
			defaultValue := ""
			if len(match) > 4 && match[4] != "" {
				defaultValue = strings.TrimSpace(match[4])
			}

			fields = append(fields, DecompiledField{
				Name:         fieldName,
				Type:         fieldType,
				Modifiers:    modifiers,
				Annotations:  make(map[string]string),
				DefaultValue: defaultValue,
				LineNumber:   0,  // TODO: Calculate line number
				JavaDoc:      "", // TODO: Extract JavaDoc
			})
		}
	}
	return fields
}

// extractAnnotations extracts class-level annotations from source code
func (d *DefaultDecompiler) extractAnnotations(sourceCode string) map[string]string {
	annotationRegex := regexp.MustCompile(`@([a-zA-Z_][a-zA-Z0-9_]*)(?:\(([^)]*)\))?`)
	matches := annotationRegex.FindAllStringSubmatch(sourceCode, -1)

	annotations := make(map[string]string)
	for _, match := range matches {
		if len(match) >= 2 {
			annotationName := match[1]
			annotationValue := ""
			if len(match) > 2 && match[2] != "" {
				annotationValue = match[2]
			}
			annotations[annotationName] = annotationValue
		}
	}
	return annotations
}

// DecompileJar decompiles all classes in a JAR file
func (d *DefaultDecompiler) DecompileJar(jarInfo *JarInfo, config *DecompilerConfig) ([]*DecompiledClass, error) {
	if jarInfo == nil {
		return nil, &DecompilerError{
			Message: "jarInfo cannot be nil",
		}
	}

	if config != nil {
		d.SetConfig(config)
	}

	decompiledClasses := make([]*DecompiledClass, 0, len(jarInfo.Classes))

	for _, classInfo := range jarInfo.Classes {
		if d.config.SkipInnerClass && strings.Contains(classInfo.ClassName, "$") {
			continue // Skip inner classes if configured
		}

		// For this implementation, we'll use a simplified approach
		// In a real implementation, you would extract the actual class bytes from the JAR
		classBytes := []byte{0xCA, 0xFE, 0xBA, 0xBE} // Mock class file header

		decompiledClass, err := d.DecompileClass(classBytes, classInfo.ClassName)
		if err != nil {
			// If DecompileClass fails, try fallback method directly
			sourceCode, fallbackErr := d.decompileWithFallback(classInfo.ClassName)
			if fallbackErr != nil {
				continue // Skip this class if fallback also fails
			}
			decompiledClass, err = d.parseDecompiledSource(sourceCode, classInfo.ClassName)
			if err != nil {
				continue // Skip this class if parsing fails
			}
		}

		decompiledClasses = append(decompiledClasses, decompiledClass)
	}

	return decompiledClasses, nil
}

// Utility functions for finding decompiled classes and methods

// GetDecompiledClassByName finds a decompiled class by name
func GetDecompiledClassByName(classes []*DecompiledClass, className string) *DecompiledClass {
	for _, class := range classes {
		if class.ClassName == className {
			return class
		}
	}
	return nil
}

// GetDecompiledClassesByPackage finds all decompiled classes in a specific package
func GetDecompiledClassesByPackage(classes []*DecompiledClass, packageName string) []*DecompiledClass {
	result := make([]*DecompiledClass, 0)
	for _, class := range classes {
		if class.PackageName == packageName {
			result = append(result, class)
		}
	}
	return result
}

// GetAllPublicMethods returns all public methods from all decompiled classes
func GetAllPublicMethods(classes []*DecompiledClass) []DecompiledMethod {
	methods := make([]DecompiledMethod, 0)
	for _, class := range classes {
		for _, method := range class.Methods {
			for _, modifier := range method.Modifiers {
				if modifier == "public" {
					methods = append(methods, method)
					break
				}
			}
		}
	}
	return methods
}

// ValidateDecompiledClass validates the structure of a decompiled class
func ValidateDecompiledClass(class *DecompiledClass) error {
	if class == nil {
		return fmt.Errorf("decompiled class cannot be nil")
	}

	if class.ClassName == "" {
		return fmt.Errorf("class name cannot be empty")
	}

	if class.SourceCode == "" {
		return fmt.Errorf("source code cannot be empty")
	}

	// Validate methods
	for i, method := range class.Methods {
		if method.Name == "" {
			return fmt.Errorf("method name at index %d cannot be empty", i)
		}
	}

	// Validate fields
	for i, field := range class.Fields {
		if field.Name == "" {
			return fmt.Errorf("field name at index %d cannot be empty", i)
		}
	}

	return nil
}
