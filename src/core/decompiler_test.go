package core

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

// TestNewDefaultDecompiler tests the creation of a default decompiler
func TestNewDefaultDecompiler(t *testing.T) {
	tests := []struct {
		name   string
		config *DecompilerConfig
		want   *DecompilerConfig
	}{
		{
			name:   "with nil config",
			config: nil,
			want: &DecompilerConfig{
				Engine:          EngineCFR,
				OutputFormat:    "java",
				SkipInnerClass:  false,
				IncludeComments: true,
				Timeout:         30 * time.Second,
				LibsPath:        "../../libs",
				JavaPath:        "java",
				TempDir:         "/tmp",
			},
		},
		{
			name: "with custom config",
			config: &DecompilerConfig{
				Engine:          EngineProcyon,
				OutputFormat:    "java",
				SkipInnerClass:  true,
				IncludeComments: false,
				Timeout:         60 * time.Second,
				LibsPath:        "custom/libs",
				JavaPath:        "/usr/bin/java",
				TempDir:         "/custom/tmp",
			},
			want: &DecompilerConfig{
				Engine:          EngineProcyon,
				OutputFormat:    "java",
				SkipInnerClass:  true,
				IncludeComments: false,
				Timeout:         60 * time.Second,
				LibsPath:        "custom/libs",
				JavaPath:        "/usr/bin/java",
				TempDir:         "/custom/tmp",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decompiler := NewDefaultDecompiler()
			if tt.config != nil {
				decompiler.SetConfig(tt.config)
			}
			if decompiler == nil {
				t.Fatal("NewDefaultDecompiler returned nil")
			}
			if decompiler.config == nil {
				t.Fatal("decompiler config is nil")
			}

			// Check all config fields
			if decompiler.config.Engine != tt.want.Engine {
				t.Errorf("Engine = %v, want %v", decompiler.config.Engine, tt.want.Engine)
			}
			if decompiler.config.OutputFormat != tt.want.OutputFormat {
				t.Errorf("OutputFormat = %v, want %v", decompiler.config.OutputFormat, tt.want.OutputFormat)
			}
			if decompiler.config.SkipInnerClass != tt.want.SkipInnerClass {
				t.Errorf("SkipInnerClass = %v, want %v", decompiler.config.SkipInnerClass, tt.want.SkipInnerClass)
			}
			if decompiler.config.IncludeComments != tt.want.IncludeComments {
				t.Errorf("IncludeComments = %v, want %v", decompiler.config.IncludeComments, tt.want.IncludeComments)
			}
			if decompiler.config.Timeout != tt.want.Timeout {
				t.Errorf("Timeout = %v, want %v", decompiler.config.Timeout, tt.want.Timeout)
			}
			if decompiler.config.LibsPath != tt.want.LibsPath {
				t.Errorf("LibsPath = %v, want %v", decompiler.config.LibsPath, tt.want.LibsPath)
			}
			if decompiler.config.JavaPath != tt.want.JavaPath {
				t.Errorf("JavaPath = %v, want %v", decompiler.config.JavaPath, tt.want.JavaPath)
			}
			if decompiler.config.TempDir != tt.want.TempDir {
				t.Errorf("TempDir = %v, want %v", decompiler.config.TempDir, tt.want.TempDir)
			}
		})
	}
}

// TestSetConfig tests updating decompiler configuration
func TestSetConfig(t *testing.T) {
	decompiler := NewDefaultDecompiler()
	newConfig := &DecompilerConfig{
		Engine:          EngineJDCore,
		OutputFormat:    "java",
		SkipInnerClass:  true,
		IncludeComments: false,
		Timeout:         45 * time.Second,
		LibsPath:        "custom/libs",
		JavaPath:        "/usr/bin/java",
		TempDir:         "/custom/tmp",
	}

	decompiler.SetConfig(newConfig)

	if decompiler.config.Engine != EngineJDCore {
		t.Errorf("Engine = %v, want %v", decompiler.config.Engine, EngineJDCore)
	}
	if decompiler.config.SkipInnerClass != true {
		t.Errorf("SkipInnerClass = %v, want %v", decompiler.config.SkipInnerClass, true)
	}
	if decompiler.config.IncludeComments != false {
		t.Errorf("IncludeComments = %v, want %v", decompiler.config.IncludeComments, false)
	}
	if decompiler.config.Timeout != 45*time.Second {
		t.Errorf("Timeout = %v, want %v", decompiler.config.Timeout, 45*time.Second)
	}
}

// TestGetSupportedFormats tests supported file formats
func TestGetSupportedFormats(t *testing.T) {
	decompiler := NewDefaultDecompiler()
	formats := decompiler.GetSupportedFormats()

	expected := []string{"class", "jar", "war", "ear"}
	if len(formats) != len(expected) {
		t.Errorf("GetSupportedFormats() returned %d formats, want %d", len(formats), len(expected))
	}

	for i, format := range expected {
		if i >= len(formats) || formats[i] != format {
			t.Errorf("GetSupportedFormats()[%d] = %v, want %v", i, formats[i], format)
		}
	}
}

// TestDecompileClass tests class decompilation
func TestDecompileClass(t *testing.T) {
	tests := []struct {
		name        string
		classBytes  []byte
		className   string
		expectError bool
		errorMsg    string
	}{
		{
			name:        "too small file",
			classBytes:  []byte{0xCA, 0xFE},
			className:   "TestClass",
			expectError: true,
			errorMsg:    "invalid class file: too small",
		},
		{
			name:        "invalid magic number",
			classBytes:  []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x34},
			className:   "TestClass",
			expectError: true,
			errorMsg:    "invalid class file: missing magic number",
		},
	}

	decompiler := NewDefaultDecompiler()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := decompiler.DecompileClass(tt.classBytes, tt.className)

			if tt.expectError {
				if err == nil {
					t.Errorf("DecompileClass() expected error but got none")
					return
				}
				if tt.errorMsg != "" && err.Error() != fmt.Sprintf("Decompiler error for class %s: %s", tt.className, tt.errorMsg) {
					t.Errorf("DecompileClass() error = %v, want error containing %v", err.Error(), tt.errorMsg)
				}
				return
			}

			if err != nil {
				t.Errorf("DecompileClass() unexpected error = %v", err)
				return
			}

			if result == nil {
				t.Errorf("DecompileClass() returned nil result")
				return
			}

			if result.ClassName != tt.className {
				t.Errorf("DecompileClass() ClassName = %v, want %v", result.ClassName, tt.className)
			}

			if result.SourceCode == "" {
				t.Errorf("DecompileClass() SourceCode is empty")
			}

			if result.SuperClass != "java.lang.Object" {
				t.Errorf("DecompileClass() SuperClass = %v, want %v", result.SuperClass, "java.lang.Object")
			}
		})
	}
}

// TestDecompileJar tests JAR decompilation
func TestDecompileJar(t *testing.T) {
	tests := []struct {
		name        string
		jarInfo     *JarInfo
		expectError bool
		errorMsg    string
	}{
		{
			name:        "nil jarInfo",
			jarInfo:     nil,
			expectError: true,
			errorMsg:    "jarInfo cannot be nil",
		},
		{
			name: "valid jarInfo with classes",
			jarInfo: &JarInfo{
				Classes: []ClassInfo{
					{ClassName: "com.example.TestClass"},
					{ClassName: "com.example.AnotherClass"},
				},
				Manifest: "Main-Class: com.example.Main\n",
			},
			expectError: false,
		},
		{
			name: "jarInfo with inner classes",
			jarInfo: &JarInfo{
				Classes: []ClassInfo{
					{ClassName: "com.example.OuterClass"},
					{ClassName: "com.example.OuterClass$InnerClass"},
				},
			},
			expectError: false,
		},
		{
			name: "empty jarInfo",
			jarInfo: &JarInfo{
				Classes: []ClassInfo{},
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decompiler := NewDefaultDecompiler()
			result, err := decompiler.DecompileJar(tt.jarInfo, nil)

			if tt.expectError {
				if err == nil {
					t.Errorf("DecompileJar() expected error but got none")
					return
				}
				if tt.errorMsg != "" && err.Error() != fmt.Sprintf("Decompiler error: %s", tt.errorMsg) {
					t.Errorf("DecompileJar() error = %v, want error containing %v", err.Error(), tt.errorMsg)
				}
				return
			}

			if err != nil {
				t.Errorf("DecompileJar() unexpected error = %v", err)
				return
			}

			if result == nil {
				t.Errorf("DecompileJar() returned nil result")
				return
			}

			if tt.jarInfo != nil {
				expectedCount := 0
				for _, class := range tt.jarInfo.Classes {
					// Count non-inner classes if SkipInnerClass is false
				if !decompiler.config.SkipInnerClass || !strings.Contains(class.ClassName, "$") {
						expectedCount++
					}
				}

				if len(result) != expectedCount {
					t.Errorf("DecompileJar() returned %d classes, want %d", len(result), expectedCount)
				}
			}
		})
	}
}

// TestDecompileJarWithSkipInnerClasses tests JAR decompilation with inner class skipping
func TestDecompileJarWithSkipInnerClasses(t *testing.T) {
	jarInfo := &JarInfo{
		Classes: []ClassInfo{
			{ClassName: "com.example.OuterClass"},
			{ClassName: "com.example.OuterClass$InnerClass"},
			{ClassName: "com.example.AnotherClass"},
		},
	}

	config := &DecompilerConfig{
		Engine:         EngineCFR,
		SkipInnerClass: true,
	}

	decompiler := NewDefaultDecompiler()
	decompiler.SetConfig(config)
	result, err := decompiler.DecompileJar(jarInfo, config)

	if err != nil {
		t.Errorf("DecompileJar() unexpected error = %v", err)
		return
	}

	// Should only have 2 classes (excluding inner class)
	if len(result) != 2 {
		t.Errorf("DecompileJar() with SkipInnerClasses returned %d classes, want 2", len(result))
	}

	// Check that inner class is not present
	found := false
	for _, class := range result {
		if class.ClassName == "com.example.OuterClass$InnerClass" {
			found = true
			break
		}
	}
	if found {
		t.Errorf("DecompileJar() with SkipInnerClass should not include inner class")
	}

	// Check that outer classes are present
	outerFound := false
	anotherFound := false
	for _, class := range result {
		if class.ClassName == "com.example.OuterClass" {
			outerFound = true
		}
		if class.ClassName == "com.example.AnotherClass" {
			anotherFound = true
		}
	}
	if !outerFound {
		t.Errorf("DecompileJar() with SkipInnerClass should include outer class")
	}
	if !anotherFound {
		t.Errorf("DecompileJar() with SkipInnerClass should include another class")
	}
}

// TestDecompilerError tests custom error type
func TestDecompilerError(t *testing.T) {
	tests := []struct {
		name      string
		err       *DecompilerError
		expected  string
	}{
		{
			name: "error without cause",
			err: &DecompilerError{
				Message:   "test error",
				ClassName: "TestClass",
			},
			expected: "Decompiler error for class TestClass: test error",
		},
		{
			name: "error with cause",
			err: &DecompilerError{
				Message:   "test error",
				ClassName: "TestClass",
				Cause:     fmt.Errorf("original error"),
			},
			expected: "Decompiler error for class TestClass: test error (caused by: original error)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err.Error() != tt.expected {
				t.Errorf("DecompilerError.Error() = %v, want %v", tt.err.Error(), tt.expected)
			}
		})
	}
}

// TestUtilityFunctions tests utility functions
func TestUtilityFunctions(t *testing.T) {
	// Create test data
	decompiledClassesMap := map[string]*DecompiledClass{
		"com.example.ClassA": {
			ClassName:   "com.example.ClassA",
			PackageName: "com.example",
			Methods: []DecompiledMethod{
				{Name: "publicMethod", Modifiers: []string{"public"}},
				{Name: "privateMethod", Modifiers: []string{"private"}},
			},
		},
		"com.example.ClassB": {
			ClassName:   "com.example.ClassB",
			PackageName: "com.example",
			Methods: []DecompiledMethod{
				{Name: "anotherPublicMethod", Modifiers: []string{"public"}},
			},
		},
		"com.other.ClassC": {
			ClassName:   "com.other.ClassC",
			PackageName: "com.other",
			Methods: []DecompiledMethod{
				{Name: "otherMethod", Modifiers: []string{"protected"}},
			},
		},
	}

	// Convert map to slice for testing
	decompiledClasses := make([]*DecompiledClass, 0, len(decompiledClassesMap))
	for _, class := range decompiledClassesMap {
		decompiledClasses = append(decompiledClasses, class)
	}

	// Test GetDecompiledClassByName
	t.Run("GetDecompiledClassByName", func(t *testing.T) {
		result := GetDecompiledClassByName(decompiledClasses, "com.example.ClassA")
		if result == nil {
			t.Errorf("GetDecompiledClassByName() returned nil for existing class")
		}
		if result != nil && result.ClassName != "com.example.ClassA" {
			t.Errorf("GetDecompiledClassByName() returned wrong class")
		}

		result = GetDecompiledClassByName(decompiledClasses, "nonexistent.Class")
		if result != nil {
			t.Errorf("GetDecompiledClassByName() should return nil for nonexistent class")
		}
	})

	// Test GetDecompiledClassesByPackage
	t.Run("GetDecompiledClassesByPackage", func(t *testing.T) {
		result := GetDecompiledClassesByPackage(decompiledClasses, "com.example")
		if len(result) != 2 {
			t.Errorf("GetDecompiledClassesByPackage() returned %d classes, want 2", len(result))
		}

		result = GetDecompiledClassesByPackage(decompiledClasses, "com.other")
		if len(result) != 1 {
			t.Errorf("GetDecompiledClassesByPackage() returned %d classes, want 1", len(result))
		}

		result = GetDecompiledClassesByPackage(decompiledClasses, "nonexistent.package")
		if len(result) != 0 {
			t.Errorf("GetDecompiledClassesByPackage() should return empty slice for nonexistent package")
		}
	})

	// Test GetAllPublicMethods
	t.Run("GetAllPublicMethods", func(t *testing.T) {
		result := GetAllPublicMethods(decompiledClasses)
		if len(result) != 2 {
			t.Errorf("GetAllPublicMethods() returned %d methods, want 2", len(result))
		}

		// Check that all returned methods are public
		for _, method := range result {
			isPublic := false
			for _, modifier := range method.Modifiers {
				if modifier == "public" {
					isPublic = true
					break
				}
			}
			if !isPublic {
				t.Errorf("GetAllPublicMethods() returned non-public method: %s", method.Name)
			}
		}
	})
}

// TestValidateDecompiledClass tests decompiled class validation
func TestValidateDecompiledClass(t *testing.T) {
	tests := []struct {
		name        string
		class       *DecompiledClass
		expectError bool
		errorMsg    string
	}{
		{
			name:        "nil class",
			class:       nil,
			expectError: true,
			errorMsg:    "decompiled class cannot be nil",
		},
		{
			name: "empty class name",
			class: &DecompiledClass{
				ClassName:  "",
				SourceCode: "public class Test {}",
			},
			expectError: true,
			errorMsg:    "class name cannot be empty",
		},
		{
			name: "empty source code",
			class: &DecompiledClass{
				ClassName:  "TestClass",
				SourceCode: "",
			},
			expectError: true,
			errorMsg:    "source code cannot be empty",
		},
		{
			name: "empty method name",
			class: &DecompiledClass{
				ClassName:  "TestClass",
				SourceCode: "public class TestClass {}",
				Methods: []DecompiledMethod{
					{Name: "", ReturnType: "void"},
				},
			},
			expectError: true,
			errorMsg:    "method name at index 0 cannot be empty",
		},
		{
			name: "empty field name",
			class: &DecompiledClass{
				ClassName:  "TestClass",
				SourceCode: "public class TestClass {}",
				Fields: []DecompiledField{
					{Name: "", Type: "String"},
				},
			},
			expectError: true,
			errorMsg:    "field name at index 0 cannot be empty",
		},
		{
			name: "valid class",
			class: &DecompiledClass{
				ClassName:  "TestClass",
				SourceCode: "public class TestClass {}",
				Methods: []DecompiledMethod{
					{Name: "testMethod", ReturnType: "void"},
				},
				Fields: []DecompiledField{
					{Name: "testField", Type: "String"},
				},
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateDecompiledClass(tt.class)

			if tt.expectError {
				if err == nil {
					t.Errorf("ValidateDecompiledClass() expected error but got none")
					return
				}
				if tt.errorMsg != "" && err.Error() != tt.errorMsg {
					t.Errorf("ValidateDecompiledClass() error = %v, want %v", err.Error(), tt.errorMsg)
				}
			} else {
				if err != nil {
					t.Errorf("ValidateDecompiledClass() unexpected error = %v", err)
				}
			}
		})
	}
}

// BenchmarkDecompileClass benchmarks class decompilation
func BenchmarkDecompileClass(b *testing.B) {
	decompiler := NewDefaultDecompiler()
	classBytes := []byte{0xCA, 0xFE, 0xBA, 0xBE, 0x00, 0x00, 0x00, 0x34}
	className := "com.example.BenchmarkClass"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := decompiler.DecompileClass(classBytes, className)
		if err != nil {
			b.Fatalf("DecompileClass() error = %v", err)
		}
	}
}

// BenchmarkDecompileJar benchmarks JAR decompilation
func BenchmarkDecompileJar(b *testing.B) {
	decompiler := NewDefaultDecompiler()
	jarInfo := &JarInfo{
		Classes: []ClassInfo{
			{ClassName: "com.example.Class1"},
			{ClassName: "com.example.Class2"},
			{ClassName: "com.example.Class3"},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := decompiler.DecompileJar(jarInfo, nil)
		if err != nil {
			b.Fatalf("DecompileJar() error = %v", err)
		}
	}
}