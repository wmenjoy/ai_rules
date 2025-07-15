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
	"archive/zip"
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// createTestJar creates a simple test JAR file in memory
func createTestJar() ([]byte, error) {
	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)

	// Add a simple class file with valid Java class magic number
	classContent := []byte{0xCA, 0xFE, 0xBA, 0xBE} // Java class magic number
	classContent = append(classContent, make([]byte, 100)...) // Add some dummy content

	f1, err := w.Create("com/example/TestClass.class")
	if err != nil {
		return nil, err
	}
	_, err = f1.Write(classContent)
	if err != nil {
		return nil, err
	}

	// Add another class in default package
	f2, err := w.Create("MainClass.class")
	if err != nil {
		return nil, err
	}
	_, err = f2.Write(classContent)
	if err != nil {
		return nil, err
	}

	// Add manifest file
	manifestContent := `Manifest-Version: 1.0
Main-Class: MainClass
Created-By: Test
`
	f3, err := w.Create("META-INF/MANIFEST.MF")
	if err != nil {
		return nil, err
	}
	_, err = f3.Write([]byte(manifestContent))
	if err != nil {
		return nil, err
	}

	// Add a resource file
	f4, err := w.Create("resources/config.properties")
	if err != nil {
		return nil, err
	}
	_, err = f4.Write([]byte("key=value\n"))
	if err != nil {
		return nil, err
	}

	err = w.Close()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// TestParseJarFromBytes tests parsing JAR from byte slice
func TestParseJarFromBytes(t *testing.T) {
	tests := []struct {
		name        string
		setupFunc   func() ([]byte, error)
		jarName     string
		expectError bool
		errorMsg    string
		validateFunc func(*testing.T, *JarInfo)
	}{
		{
			name: "valid jar file",
			setupFunc: createTestJar,
			jarName: "test.jar",
			expectError: false,
			validateFunc: func(t *testing.T, ji *JarInfo) {
				if ji.JarName != "test.jar" {
					t.Errorf("Expected jar name 'test.jar', got '%s'", ji.JarName)
				}
				if ji.TotalClasses != 2 {
					t.Errorf("Expected 2 classes, got %d", ji.TotalClasses)
				}
				if len(ji.Classes) != 2 {
					t.Errorf("Expected 2 classes in slice, got %d", len(ji.Classes))
				}
				if ji.Manifest == "" {
					t.Error("Expected manifest content, got empty string")
				}
				if len(ji.Resources) == 0 {
					t.Error("Expected resources, got none")
				}
			},
		},
		{
			name: "invalid jar data",
			setupFunc: func() ([]byte, error) {
				return []byte("not a jar file"), nil
			},
			jarName: "invalid.jar",
			expectError: true,
			errorMsg: "failed to open JAR as ZIP archive",
		},
		{
			name: "empty data",
			setupFunc: func() ([]byte, error) {
				return []byte{}, nil
			},
			jarName: "empty.jar",
			expectError: true,
			errorMsg: "failed to open JAR as ZIP archive",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := tt.setupFunc()
			if err != nil {
				t.Fatalf("Failed to setup test data: %v", err)
			}

			jarInfo, err := ParseJarFromBytes(data, tt.jarName)

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
					return
				}
				if jarInfo == nil {
					t.Errorf("Expected jar info but got nil")
					return
				}
				if tt.validateFunc != nil {
					tt.validateFunc(t, jarInfo)
				}
			}
		})
	}
}

// TestParseClassFile tests class file parsing
func TestParseClassFile(t *testing.T) {
	tests := []struct {
		name        string
		classData   []byte
		fileName    string
		expectError bool
		errorMsg    string
		validateFunc func(*testing.T, *ClassInfo)
	}{
		{
			name: "valid class file",
			classData: append([]byte{0xCA, 0xFE, 0xBA, 0xBE}, make([]byte, 20)...),
			fileName: "com/example/TestClass.class",
			expectError: false,
			validateFunc: func(t *testing.T, ci *ClassInfo) {
				if ci.ClassName != "TestClass" {
					t.Errorf("Expected class name 'TestClass', got '%s'", ci.ClassName)
				}
				if ci.PackageName != "com.example" {
					t.Errorf("Expected package name 'com.example', got '%s'", ci.PackageName)
				}
				if ci.SourceFile != "com/example/TestClass.class" {
					t.Errorf("Expected source file 'com/example/TestClass.class', got '%s'", ci.SourceFile)
				}
			},
		},
		{
			name: "class in default package",
			classData: append([]byte{0xCA, 0xFE, 0xBA, 0xBE}, make([]byte, 20)...),
			fileName: "MainClass.class",
			expectError: false,
			validateFunc: func(t *testing.T, ci *ClassInfo) {
				if ci.ClassName != "MainClass" {
					t.Errorf("Expected class name 'MainClass', got '%s'", ci.ClassName)
				}
				if ci.PackageName != "" {
					t.Errorf("Expected empty package name, got '%s'", ci.PackageName)
				}
			},
		},
		{
			name: "invalid magic number",
			classData: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			fileName: "Invalid.class",
			expectError: true,
			errorMsg: "invalid class file magic number",
		},
		{
			name: "file too small",
			classData: []byte{0xCA, 0xFE},
			fileName: "Small.class",
			expectError: true,
			errorMsg: "class file Small.class is too small",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			classInfo, err := parseClassFile(tt.classData, tt.fileName)

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
					return
				}
				if classInfo == nil {
					t.Errorf("Expected class info but got nil")
					return
				}
				if tt.validateFunc != nil {
					tt.validateFunc(t, classInfo)
				}
			}
		})
	}
}

// TestExtractPackageFromPath tests package extraction from file paths
func TestExtractPackageFromPath(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		expected string
	}{
		{"class in package", "com/example/TestClass.class", "com.example"},
		{"class in nested package", "org/springframework/boot/Application.class", "org.springframework.boot"},
		{"class in default package", "MainClass.class", ""},
		{"single level package", "util/Helper.class", "util"},
		{"windows path separators", "com\\example\\TestClass.class", "com.example"},
		{"mixed separators", "com/example\\util/Helper.class", "com.example.util"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractPackageFromPath(tt.filePath)
			if result != tt.expected {
				t.Errorf("extractPackageFromPath(%q) = %q, expected %q", tt.filePath, result, tt.expected)
			}
		})
	}
}

// TestJarInfoMethods tests various methods on JarInfo
func TestJarInfoMethods(t *testing.T) {
	// Create test jar info
	jarInfo := &JarInfo{
		JarName: "test.jar",
		Classes: []ClassInfo{
			{
				ClassName:   "TestClass",
				PackageName: "com.example",
				Methods: []MethodInfo{
					{MethodName: "publicMethod", IsPublic: true},
					{MethodName: "privateMethod", IsPrivate: true},
				},
			},
			{
				ClassName:   "MainClass",
				PackageName: "",
				Methods: []MethodInfo{
					{MethodName: "main", IsPublic: true, IsStatic: true},
				},
			},
			{
				ClassName:   "AnotherClass",
				PackageName: "com.example",
			},
		},
		Manifest: "Manifest-Version: 1.0\nMain-Class: MainClass\n",
	}

	// Test GetClassByName
	t.Run("GetClassByName", func(t *testing.T) {
		// Test existing class with package
		class := jarInfo.GetClassByName("com.example.TestClass")
		if class == nil {
			t.Error("Expected to find com.example.TestClass")
		} else if class.ClassName != "TestClass" {
			t.Errorf("Expected class name 'TestClass', got '%s'", class.ClassName)
		}

		// Test class in default package
		class = jarInfo.GetClassByName("MainClass")
		if class == nil {
			t.Error("Expected to find MainClass")
		} else if class.ClassName != "MainClass" {
			t.Errorf("Expected class name 'MainClass', got '%s'", class.ClassName)
		}

		// Test non-existent class
		class = jarInfo.GetClassByName("NonExistent")
		if class != nil {
			t.Error("Expected nil for non-existent class")
		}
	})

	// Test GetClassesByPackage
	t.Run("GetClassesByPackage", func(t *testing.T) {
		// Test existing package
		classes := jarInfo.GetClassesByPackage("com.example")
		if len(classes) != 2 {
			t.Errorf("Expected 2 classes in com.example package, got %d", len(classes))
		}

		// Test default package
		classes = jarInfo.GetClassesByPackage("")
		if len(classes) != 1 {
			t.Errorf("Expected 1 class in default package, got %d", len(classes))
		}

		// Test non-existent package
		classes = jarInfo.GetClassesByPackage("non.existent")
		if len(classes) != 0 {
			t.Errorf("Expected 0 classes in non-existent package, got %d", len(classes))
		}
	})

	// Test GetPublicMethods
	t.Run("GetPublicMethods", func(t *testing.T) {
		methods := jarInfo.GetPublicMethods()
		if len(methods) != 2 {
			t.Errorf("Expected 2 public methods, got %d", len(methods))
		}

		// Check method names
		methodNames := make(map[string]bool)
		for _, method := range methods {
			methodNames[method.MethodName] = true
		}
		if !methodNames["publicMethod"] {
			t.Error("Expected to find 'publicMethod' in public methods")
		}
		if !methodNames["main"] {
			t.Error("Expected to find 'main' in public methods")
		}
		if methodNames["privateMethod"] {
			t.Error("Did not expect to find 'privateMethod' in public methods")
		}
	})

	// Test GetManifestAttribute
	t.Run("GetManifestAttribute", func(t *testing.T) {
		// Test existing attribute
		mainClass := jarInfo.GetManifestAttribute("Main-Class")
		if mainClass != "MainClass" {
			t.Errorf("Expected Main-Class 'MainClass', got '%s'", mainClass)
		}

		// Test existing attribute
		version := jarInfo.GetManifestAttribute("Manifest-Version")
		if version != "1.0" {
			t.Errorf("Expected Manifest-Version '1.0', got '%s'", version)
		}

		// Test non-existent attribute
		nonExistent := jarInfo.GetManifestAttribute("Non-Existent")
		if nonExistent != "" {
			t.Errorf("Expected empty string for non-existent attribute, got '%s'", nonExistent)
		}
	})
}

// TestValidateJarStructure tests JAR structure validation
func TestValidateJarStructure(t *testing.T) {
	tests := []struct {
		name        string
		jarInfo     *JarInfo
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid jar structure",
			jarInfo: &JarInfo{
				JarName: "test.jar",
				Classes: []ClassInfo{
					{ClassName: "ValidClass"},
				},
			},
			expectError: false,
		},
		{
			name: "empty jar name",
			jarInfo: &JarInfo{
				JarName: "",
				Classes: []ClassInfo{
					{ClassName: "ValidClass"},
				},
			},
			expectError: true,
			errorMsg: "JAR name is empty",
		},
		{
			name: "no classes",
			jarInfo: &JarInfo{
				JarName: "test.jar",
				Classes: []ClassInfo{},
			},
			expectError: true,
			errorMsg: "no classes found in JAR",
		},
		{
			name: "invalid class name",
			jarInfo: &JarInfo{
				JarName: "test.jar",
				Classes: []ClassInfo{
					{ClassName: "123InvalidClass"}, // starts with number
				},
			},
			expectError: true,
			errorMsg: "invalid class name",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.jarInfo.ValidateJarStructure()

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

// TestJarParseError tests custom error type
func TestJarParseError(t *testing.T) {
	originalErr := fmt.Errorf("original error")
	parseErr := &JarParseError{
		Message: "test parse error",
		Cause:   originalErr,
	}

	// Test Error() method
	expectedMsg := "JAR parse error: test parse error"
	if parseErr.Error() != expectedMsg {
		t.Errorf("Error() = %q, expected %q", parseErr.Error(), expectedMsg)
	}

	// Test Unwrap() method
	if parseErr.Unwrap() != originalErr {
		t.Errorf("Unwrap() = %v, expected %v", parseErr.Unwrap(), originalErr)
	}
}

// BenchmarkParseJarFromBytes benchmarks JAR parsing
func BenchmarkParseJarFromBytes(b *testing.B) {
	// Create test data once
	testData, err := createTestJar()
	if err != nil {
		b.Fatalf("Failed to create test JAR: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := ParseJarFromBytes(testData, "benchmark.jar")
		if err != nil {
			b.Fatalf("Unexpected error: %v", err)
		}
	}
}