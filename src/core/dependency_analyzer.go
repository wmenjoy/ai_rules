package core

import (
	"fmt"
	"regexp"
	"strings"
)

// ClassDependency represents a dependency relationship between classes.
type ClassDependency struct {
	FromClass string // The class that depends on another
	ToClass   string // The class being depended upon
	DepType   string // Type of dependency: "extends", "implements", "imports", "uses_field", "method_call"
}

// MethodCall represents a method call relationship.
type MethodCall struct {
	CallerClass  string // Class containing the calling method
	CallerMethod string // Method making the call
	CalleeClass  string // Class containing the called method
	CalleeMethod string // Method being called
	LineNumber   int    // Line number where the call occurs
}

// DependencyAnalysisResult contains the results of dependency analysis.
type DependencyAnalysisResult struct {
	ClassDependencies []ClassDependency        // All class-level dependencies
	MethodCalls       []MethodCall             // All method call relationships
	InheritanceTree   map[string]string        // Class -> Parent class mapping
	InterfaceImpls    map[string][]string      // Class -> Implemented interfaces mapping
	ImportStatements  map[string][]string      // Class -> Import statements mapping
}

// AnalyzeDependencies analyzes the dependency relationships in the given decompiled classes.
func AnalyzeDependencies(decompiledClasses map[string]*DecompiledClass) (*DependencyAnalysisResult, error) {
	if decompiledClasses == nil {
		return nil, fmt.Errorf("decompiledClasses cannot be nil")
	}

	result := &DependencyAnalysisResult{
		ClassDependencies: make([]ClassDependency, 0),
		MethodCalls:       make([]MethodCall, 0),
		InheritanceTree:   make(map[string]string),
		InterfaceImpls:    make(map[string][]string),
		ImportStatements:  make(map[string][]string),
	}

	// Analyze each class
	for className, classInfo := range decompiledClasses {
		if classInfo == nil {
			continue
		}

		// Extract imports from source code
		imports := extractImports(classInfo.SourceCode)
		result.ImportStatements[className] = imports

		// Analyze inheritance and interface implementation
		analyzeInheritance(classInfo, result)

		// Analyze method calls within the class
		analyzeMethodCalls(classInfo, result)

		// Analyze field usage and other dependencies
		analyzeFieldDependencies(classInfo, result)
	}

	// Build class dependencies from collected data
	buildClassDependencies(result)

	return result, nil
}

// extractImports extracts import statements from source code.
func extractImports(sourceCode string) []string {
	importRegex := regexp.MustCompile(`import\s+(?:static\s+)?([a-zA-Z_][a-zA-Z0-9_.]*(?:\.\*)?);`)
	matches := importRegex.FindAllStringSubmatch(sourceCode, -1)

	var imports []string
	for _, match := range matches {
		if len(match) > 1 {
			imports = append(imports, match[1])
		}
	}
	return imports
}

// analyzeInheritance extracts inheritance relationships from Java source code using AST parsing
func analyzeInheritance(classInfo *DecompiledClass, result *DependencyAnalysisResult) {
	// Initialize maps if nil
	if result.InheritanceTree == nil {
		result.InheritanceTree = make(map[string]string)
	}
	if result.InterfaceImpls == nil {
		result.InterfaceImpls = make(map[string][]string)
	}

	// Use AST parsing instead of regex
	inheritanceTree, interfaceImpls, err := ParseJavaInheritance(classInfo.SourceCode)
	if err != nil {
		// If AST parsing fails, log the error but don't crash
		// In a real application, you might want to use a proper logger
		return
	}

	// Merge results into the main result structure
	for className, parentClass := range inheritanceTree {
		result.InheritanceTree[className] = parentClass
	}

	for className, interfaces := range interfaceImpls {
		result.InterfaceImpls[className] = interfaces
	}
}

// analyzeMethodCalls extracts method call information from the source code.
func analyzeMethodCalls(classInfo *DecompiledClass, result *DependencyAnalysisResult) {
	// Look for method call patterns: object.method() or ClassName.method()
	methodCallRegex := regexp.MustCompile(`([a-zA-Z_][a-zA-Z0-9_]*)\.([a-zA-Z_][a-zA-Z0-9_]*)\s*\(`)
	matches := methodCallRegex.FindAllStringSubmatch(classInfo.SourceCode, -1)

	for _, match := range matches {
		if len(match) > 2 {
			calleeClass := match[1]
			calleeMethod := match[2]
			callerMethod := findContainingMethod(classInfo, match[0])

			methodCall := MethodCall{
				CallerClass:  classInfo.ClassName,
				CallerMethod: callerMethod,
				CalleeClass:  calleeClass,
				CalleeMethod: calleeMethod,
				LineNumber:   0, // Would need more sophisticated parsing to get exact line
			}
			result.MethodCalls = append(result.MethodCalls, methodCall)
		}
	}
}

// findContainingMethod attempts to find which method contains a given method call.
func findContainingMethod(classInfo *DecompiledClass, methodCallText string) string {
	// This is a simplified approach - in practice, you'd want to parse the AST
	for _, method := range classInfo.Methods {
		if strings.Contains(classInfo.SourceCode, methodCallText) {
			// More sophisticated logic would parse the AST to determine exact method boundaries
			return method.Name
		}
	}
	return "unknown"
}

// analyzeFieldDependencies analyzes field usage and type dependencies.
func analyzeFieldDependencies(classInfo *DecompiledClass, result *DependencyAnalysisResult) {
	// Use the Fields information from DecompiledClass instead of regex parsing
	for _, field := range classInfo.Fields {
		fieldType := field.Type
		// Extract base type from generics (e.g., "List<String>" -> "List")
		if idx := strings.Index(fieldType, "<"); idx != -1 {
			fieldType = fieldType[:idx]
		}
		// Remove array brackets (e.g., "String[]" -> "String")
		fieldType = strings.TrimSuffix(fieldType, "[]")
		
		if fieldType != "" && !isPrimitiveType(fieldType) {
			dependency := ClassDependency{
				FromClass: classInfo.ClassName,
				ToClass:   fieldType,
				DepType:   "uses_field",
			}
			result.ClassDependencies = append(result.ClassDependencies, dependency)
		}
	}
}

// buildClassDependencies creates ClassDependency entries from inheritance and interface data.
func buildClassDependencies(result *DependencyAnalysisResult) {
	// Add inheritance dependencies
	for child, parent := range result.InheritanceTree {
		dependency := ClassDependency{
			FromClass: child,
			ToClass:   parent,
			DepType:   "extends",
		}
		result.ClassDependencies = append(result.ClassDependencies, dependency)
	}

	// Add interface implementation dependencies
	for class, interfaces := range result.InterfaceImpls {
		for _, iface := range interfaces {
			dependency := ClassDependency{
				FromClass: class,
				ToClass:   iface,
				DepType:   "implements",
			}
			result.ClassDependencies = append(result.ClassDependencies, dependency)
		}
	}

	// Add import dependencies
	for class, imports := range result.ImportStatements {
		for _, importPath := range imports {
			// Extract class name from import path
			parts := strings.Split(importPath, ".")
			if len(parts) > 0 {
				importedClass := parts[len(parts)-1]
				if importedClass != "*" { // Skip wildcard imports
					dependency := ClassDependency{
						FromClass: class,
						ToClass:   importedClass,
						DepType:   "imports",
					}
					result.ClassDependencies = append(result.ClassDependencies, dependency)
				}
			}
		}
	}
}

// isPrimitiveType checks if a type is a Java primitive type.
func isPrimitiveType(typeName string) bool {
	primitives := map[string]bool{
		"boolean": true,
		"byte":    true,
		"char":    true,
		"short":   true,
		"int":     true,
		"long":    true,
		"float":   true,
		"double":  true,
		"void":    true,
		"String":  true, // Common case
	}
	return primitives[typeName]
}

// GetDependencyGraph returns a simplified dependency graph for visualization.
func (result *DependencyAnalysisResult) GetDependencyGraph() map[string][]string {
	graph := make(map[string][]string)

	for _, dep := range result.ClassDependencies {
		if graph[dep.FromClass] == nil {
			graph[dep.FromClass] = []string{}
		}
		graph[dep.FromClass] = append(graph[dep.FromClass], dep.ToClass)
	}

	return graph
}

// GetMethodCallGraph returns a method-level call graph.
func (result *DependencyAnalysisResult) GetMethodCallGraph() map[string][]string {
	graph := make(map[string][]string)

	for _, call := range result.MethodCalls {
		caller := fmt.Sprintf("%s.%s", call.CallerClass, call.CallerMethod)
		callee := fmt.Sprintf("%s.%s", call.CalleeClass, call.CalleeMethod)

		if graph[caller] == nil {
			graph[caller] = []string{}
		}
		graph[caller] = append(graph[caller], callee)
	}

	return graph
}
