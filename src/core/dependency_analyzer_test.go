package core

import (
	"testing"
)

func TestAnalyzeDependencies(t *testing.T) {
	tests := []struct {
		name              string
		decompiledClasses map[string]*DecompiledClass
		expectedError     bool
		expectedDepCount  int
		expectedCallCount int
	}{
		{
			name:              "nil input",
			decompiledClasses: nil,
			expectedError:     true,
		},
		{
			name:              "empty input",
			decompiledClasses: make(map[string]*DecompiledClass),
			expectedError:     false,
			expectedDepCount:  0,
			expectedCallCount: 0,
		},
		{
			name: "simple class with inheritance",
			decompiledClasses: map[string]*DecompiledClass{
				"TestClass": {
					ClassName:   "TestClass",
					PackageName: "com.example",
					SourceCode: `package com.example;

import java.util.List;
import java.util.ArrayList;

public class TestClass extends BaseClass implements Runnable {
    private String name;
    private List<String> items;
    
    public void run() {
        System.out.println("Running");
        items.add("test");
    }
    
    public void processItems() {
        items.forEach(System.out::println);
    }
}`,
					Methods: []DecompiledMethod{
						{Name: "run", ReturnType: "void"},
						{Name: "processItems", ReturnType: "void"},
					},
					Fields: []DecompiledField{
						{Name: "name", Type: "String"},
						{Name: "items", Type: "List<String>"},
					},
				},
			},
			expectedError:     false,
			expectedDepCount:  5, // uses_field List, extends BaseClass, implements Runnable, imports List, imports ArrayList
			expectedCallCount: 3, // System.out.println, items.add, items.forEach
		},
		{
			name: "class with method calls",
			decompiledClasses: map[string]*DecompiledClass{
				"ServiceClass": {
					ClassName:   "ServiceClass",
					PackageName: "com.example.service",
					SourceCode: `package com.example.service;

public class ServiceClass {
    private DatabaseHelper dbHelper;
    
    public void saveData(String data) {
        dbHelper.connect();
        dbHelper.save(data);
        dbHelper.disconnect();
    }
    
    public String loadData(int id) {
        dbHelper.connect();
        String result = dbHelper.load(id);
        dbHelper.disconnect();
        return result;
    }
}`,
					Methods: []DecompiledMethod{
						{Name: "saveData", ReturnType: "void"},
						{Name: "loadData", ReturnType: "String"},
					},
					Fields: []DecompiledField{
						{Name: "dbHelper", Type: "DatabaseHelper"},
					},
				},
			},
			expectedError:     false,
			expectedDepCount:  1, // uses_field DatabaseHelper
			expectedCallCount: 6, // 3 calls in saveData + 3 calls in loadData
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := AnalyzeDependencies(tt.decompiledClasses)

			if tt.expectedError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if result == nil {
				t.Errorf("Expected result but got nil")
				return
			}

			if len(result.ClassDependencies) != tt.expectedDepCount {
				t.Errorf("Expected %d class dependencies, got %d", tt.expectedDepCount, len(result.ClassDependencies))
				for i, dep := range result.ClassDependencies {
					t.Logf("Dependency %d: %s -> %s (%s)", i, dep.FromClass, dep.ToClass, dep.DepType)
				}
			}

			if len(result.MethodCalls) != tt.expectedCallCount {
				t.Errorf("Expected %d method calls, got %d", tt.expectedCallCount, len(result.MethodCalls))
			}
		})
	}
}

func TestExtractImports(t *testing.T) {
	tests := []struct {
		name           string
		sourceCode     string
		expectedCount  int
		expectedImport string
	}{
		{
			name:          "no imports",
			sourceCode:    "public class Test {}",
			expectedCount: 0,
		},
		{
			name: "single import",
			sourceCode: `import java.util.List;
public class Test {}`,
			expectedCount:  1,
			expectedImport: "java.util.List",
		},
		{
			name: "multiple imports",
			sourceCode: `import java.util.List;
import java.util.ArrayList;
import static java.lang.System.out;
public class Test {}`,
			expectedCount: 3,
		},
		{
			name: "wildcard import",
			sourceCode: `import java.util.*;
public class Test {}`,
			expectedCount:  1,
			expectedImport: "java.util.*",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			imports := extractImports(tt.sourceCode)

			if len(imports) != tt.expectedCount {
				t.Errorf("Expected %d imports, got %d", tt.expectedCount, len(imports))
			}

			if tt.expectedImport != "" {
				found := false
				for _, imp := range imports {
					if imp == tt.expectedImport {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Expected import '%s' not found in %v", tt.expectedImport, imports)
				}
			}
		})
	}
}

func TestIsPrimitiveType(t *testing.T) {
	tests := []struct {
		typeName string
		expected bool
	}{
		{"int", true},
		{"boolean", true},
		{"String", true},
		{"void", true},
		{"CustomClass", false},
		{"List", false},
		{"Object", false},
	}

	for _, tt := range tests {
		t.Run(tt.typeName, func(t *testing.T) {
			result := isPrimitiveType(tt.typeName)
			if result != tt.expected {
				t.Errorf("isPrimitiveType(%s) = %v, expected %v", tt.typeName, result, tt.expected)
			}
		})
	}
}

func TestGetDependencyGraph(t *testing.T) {
	result := &DependencyAnalysisResult{
		ClassDependencies: []ClassDependency{
			{FromClass: "A", ToClass: "B", DepType: "extends"},
			{FromClass: "A", ToClass: "C", DepType: "implements"},
			{FromClass: "B", ToClass: "D", DepType: "uses_field"},
		},
	}

	graph := result.GetDependencyGraph()

	if len(graph) != 2 {
		t.Errorf("Expected 2 nodes in graph, got %d", len(graph))
	}

	if len(graph["A"]) != 2 {
		t.Errorf("Expected A to have 2 dependencies, got %d", len(graph["A"]))
	}

	if len(graph["B"]) != 1 {
		t.Errorf("Expected B to have 1 dependency, got %d", len(graph["B"]))
	}
}

func TestGetMethodCallGraph(t *testing.T) {
	result := &DependencyAnalysisResult{
		MethodCalls: []MethodCall{
			{CallerClass: "A", CallerMethod: "methodA", CalleeClass: "B", CalleeMethod: "methodB"},
			{CallerClass: "A", CallerMethod: "methodA", CalleeClass: "C", CalleeMethod: "methodC"},
			{CallerClass: "B", CallerMethod: "methodB", CalleeClass: "D", CalleeMethod: "methodD"},
		},
	}

	graph := result.GetMethodCallGraph()

	if len(graph) != 2 {
		t.Errorf("Expected 2 nodes in method call graph, got %d", len(graph))
	}

	callerA := "A.methodA"
	if len(graph[callerA]) != 2 {
		t.Errorf("Expected %s to have 2 calls, got %d", callerA, len(graph[callerA]))
	}

	callerB := "B.methodB"
	if len(graph[callerB]) != 1 {
		t.Errorf("Expected %s to have 1 call, got %d", callerB, len(graph[callerB]))
	}
}

func TestAnalyzeInheritance(t *testing.T) {
	tests := []struct {
		name               string
		sourceCode         string
		expectedParent     string
		expectedInterfaces []string
	}{
		{
			name:           "class with extends",
			sourceCode:     "public class TestClass extends BaseClass {}",
			expectedParent: "BaseClass",
		},
		{
			name:               "class with implements",
			sourceCode:         "public class TestClass implements Runnable, Serializable {}",
			expectedInterfaces: []string{"Runnable", "Serializable"},
		},
		{
			name:               "class with extends and implements",
			sourceCode:         "public class TestClass extends BaseClass implements Runnable {}",
			expectedParent:     "BaseClass",
			expectedInterfaces: []string{"Runnable"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			classInfo := &DecompiledClass{
				ClassName:  "TestClass",
				SourceCode: tt.sourceCode,
			}

			result := &DependencyAnalysisResult{
				InheritanceTree: make(map[string]string),
				InterfaceImpls:  make(map[string][]string),
			}

			analyzeInheritance(classInfo, result)

			if tt.expectedParent != "" {
				if result.InheritanceTree["TestClass"] != tt.expectedParent {
					t.Errorf("Expected parent %s, got %s", tt.expectedParent, result.InheritanceTree["TestClass"])
				}
			}

			if len(tt.expectedInterfaces) > 0 {
				interfaces := result.InterfaceImpls["TestClass"]
				if len(interfaces) != len(tt.expectedInterfaces) {
					t.Errorf("Expected %d interfaces, got %d", len(tt.expectedInterfaces), len(interfaces))
				}
				for i, expected := range tt.expectedInterfaces {
					if i < len(interfaces) && interfaces[i] != expected {
						t.Errorf("Expected interface %s at index %d, got %s", expected, i, interfaces[i])
					}
				}
			}
		})
	}
}
