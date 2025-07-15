## 相关文件

- `src/core/jar_processor.go` - 处理JAR包上传、解析和反编译的核心逻辑。
- `src/core/jar_processor_test.go` - `jar_processor.go`的单元测试。
- `src/ai/doc_generator.go` - 与AI模型交互生成文档的模块。
- `src/ai/doc_generator_test.go` - `doc_generator.go`的单元测试。
- `src/output/markdown_formatter.go` - 生成Markdown格式文档的模块。
- `src/output/openapi_formatter.go` - 生成OpenAPI格式文档的模块。
- `src/output/jsonschema_formatter.go` - 生成JSON Schema格式文档的模块。
- `src/api/handlers.go` - HTTP API接口处理器。
- `src/api/handlers_test.go` - `handlers.go`的单元测试。
- `src/mcp/mcp_adapter.go` - MCP集成适配器。
- `cmd/server/main.go` - 服务主入口。

### 注释

- 单元测试通常应放置在它们测试的代码文件旁边（例如，`MyComponent.tsx`和`MyComponent.test.tsx`在同一目录中）。
- 使用`go test ./...`来运行所有测试。

## 任务

- [ ] 1.0 核心功能实现：JAR包处理与基础文档生成
  - [x] 1.1 实现JAR包上传功能
    - **描述:** 开发一个模块，允许用户通过界面或API上传本地JAR文件，或者提供一个URL供系统下载JAR文件。需要处理文件大小限制、类型校验和潜在的下载错误。
    - **依赖:** 无
    - **代码示例 (Go):**
      ```go
      // src/core/jar_uploader.go
      package core

      import (
      	"io"
      	"net/http"
      	"os"
      )

      // UploadFromLocal handles uploading a JAR file from a local path.
      // It should validate the file type and size.
      func UploadFromLocal(filePath string) (io.Reader, error) {
      	// Placeholder: Implementation for local file upload
      	// e.g., open file, basic validation
      	return os.Open(filePath)
      }

      // UploadFromURL handles downloading a JAR file from a given URL.
      // It should handle potential network errors and validate the content.
      func UploadFromURL(jarURL string) (io.Reader, error) {
      	// Placeholder: Implementation for URL download
      	// e.g., http.Get, basic validation
      	resp, err := http.Get(jarURL)
      	if err != nil {
      		return nil, err
      	}
      	// Caller is responsible for closing resp.Body
      	return resp.Body, nil
      }
      ```
  - [x] 1.2 实现JAR包解析功能
    - **描述:** 解析上传的JAR文件内容，识别其结构（如MANIFEST.MF, class files, resources）。支持JDK 6+的JAR包格式。初步处理代码混淆，例如识别常见的混淆模式或标记混淆过的类/方法。
    - **依赖:** 1.1
    - **代码示例 (Go):**
      ```go
      // src/core/jar_parser.go
      package core

      import (
      	"archive/zip"
      	"io"
        "strings"
      )

      // ParsedJarInfo holds basic information extracted from a JAR.
      type ParsedJarInfo struct {
      	Manifest   map[string]string // Contents of MANIFEST.MF
      	ClassFiles []string          // List of .class file paths within the JAR
      	// ... other relevant info
      }

      // ParseJarStream analyzes a JAR file stream.
      // readerAt is required for zip.NewReader.
      func ParseJarStream(jarStream io.ReaderAt, size int64) (*ParsedJarInfo, error) {
      	// Placeholder: Implementation for JAR parsing
      	// e.g., use archive/zip to read entries
      	r, err := zip.NewReader(jarStream, size)
      	if err != nil {
      		return nil, err
      	}
      	info := &ParsedJarInfo{ClassFiles: []string{}, Manifest: make(map[string]string)}
      	for _, f := range r.File {
      		if f.Name == "META-INF/MANIFEST.MF" {
      			// TODO: parse manifest file content into info.Manifest
      		} else if strings.HasSuffix(f.Name, ".class") {
      			info.ClassFiles = append(info.ClassFiles, f.Name)
      		}
      	}
      	return info, nil
      }
      ```
  - [x] 1.3 集成反编译引擎
    - **描述:** 集成一个或多个Java反编译引擎（如CFR, Procyon, JD-Core）作为外部工具或库。该模块负责调用反编译引擎处理JAR包中的.class文件，并提取类结构、方法签名、参数、返回值和JavaDoc注释（如果存在）。
    - **依赖:** 1.2
    - **代码示例 (Go):**
      ```go
      // src/core/decompiler.go
      package core

      // DecompiledClass represents the structure of a decompiled Java class.
      type DecompiledClass struct {
      	Name        string
      	PackageName string
      	Methods     []DecompiledMethod
      	Fields      []DecompiledField
      	SourceCode  string // Raw decompiled source code
      	// ... other details like annotations, superclass, interfaces
      }
      type DecompiledMethod struct {
      	Name       string
      	Signature  string // e.g., (Ljava/lang/String;I)V
      	Parameters []string // e.g., ["String arg0", "int arg1"]
      	ReturnType string   // e.g., "void"
      	Comments   string   // JavaDoc or other comments
      }
      type DecompiledField struct {
      	Name string
      	Type string
      }

      // DecompileJar uses an external decompiler (e.g., CFR) to process class files.
      func DecompileJar(jarPath string, classFiles []string) (map[string]*DecompiledClass, error) {
      	// Placeholder: Implementation for calling a decompiler
      	// This might involve running a command-line tool and parsing its output.
      	// For each classFile in classFiles:
      	//   decompiledOutput := runDecompiler(jarPath, classFile)
      	//   parsedClass := parseDecompiledOutput(decompiledOutput)
      	//   resultMap[classFile] = parsedClass
      	return make(map[string]*DecompiledClass), nil
      }
      ```
  - [x] 1.4 实现依赖关系和调用链分析模块
    - **描述:** 分析反编译后的代码，识别类之间的依赖关系（继承、实现、组合、聚合）以及方法间的调用链。这有助于理解代码结构和功能模块间的交互。
    - **依赖:** 1.3
    - **代码示例 (Go):**
      ```go
      // src/core/dependency_analyzer.go
      package core

      // ClassDependency represents a dependency between two classes.
      type ClassDependency struct {
      	FromClass string
      	ToClass   string
      	DepType   string // e.g., "extends", "implements", "calls_method"
      }

      // MethodCall represents a call from one method to another.
      type MethodCall struct {
      	CallerClass  string
      	CallerMethod string
      	CalleeClass  string
      	CalleeMethod string
      }

      // AnalyzeDependencies processes decompiled classes to find relationships.
      func AnalyzeDependencies(decompiledClasses map[string]*DecompiledClass) ([]ClassDependency, []MethodCall, error) {
      	// Placeholder: Implementation for dependency and call chain analysis
      	// This would involve parsing source code or bytecode-like structures
      	// to identify import statements, inheritance, method invocations etc.
      	return []ClassDependency{}, []MethodCall{}, nil
      }
      ```
  - [x] 1.5 开发AI文档生成模块
    - **描述:** 构建一个模块，该模块将反编译的代码结构、方法签名、注释以及依赖分析结果作为输入，调用外部AI模型（如Qwen3, DeepSeek, OpenAI API）生成自然语言描述的文档。
    - **依赖:** 1.3, 1.4
    - **代码示例 (Go):**
      ```go
      // src/ai/doc_generator.go
      package ai

      import (
      	"//Users/liujinliang/workspace/ai/ai_rules/src/core" // Assuming DecompiledClass is here
      )

      // AIModelConfig holds configuration for an AI model.
      type AIModelConfig struct {
      	Endpoint string
      	APIKey   string
      	Model    string // e.g., "qwen3-72b-chat", "deepseek-coder"
      }

      // GenerateAIDocumentation sends code information to an AI model and gets documentation.
      func GenerateAIDocumentation(classInfo *core.DecompiledClass, config AIModelConfig) (string, error) {
      	// Placeholder: Implementation for AI documentation generation
      	// 1. Format input for the AI model (e.g., prompt engineering).
      	// 2. Make API call to the specified AI model endpoint.
      	// 3. Parse the AI's response to extract generated documentation.
      	prompt := buildPrompt(classInfo) // Helper function to create a good prompt
      	_ = prompt // avoid unused variable error for placeholder
      	// response, err := callAIModel(config, prompt)
      	// if err != nil { return "", err }
      	// return parseAIResponse(response), nil
      	return "Generated documentation for " + classInfo.Name, nil
      }

      func buildPrompt(classInfo *core.DecompiledClass) string {
        // Construct a detailed prompt including class structure, methods, fields, comments
        return "Describe the Java class " + classInfo.Name + "..."
      }
      ```
  - [x] 1.6 实现基于代码语义生成方法说明、参数描述、使用示例和异常处理说明的功能
    - **描述:** 细化AI文档生成模块，使其能够针对每个方法，根据其签名、内部逻辑（通过反编译代码理解）、以及上下文（类信息、调用关系），生成详细的方法功能说明、每个参数的用途和类型、典型的使用代码示例，以及可能抛出的异常和处理建议。
    - **依赖:** 1.5
    - **代码示例 (Go):**
      ```go
      // src/ai/semantic_doc_generator.go (can extend doc_generator.go)
      package ai

      import (
      	"//Users/liujinliang/workspace/ai/ai_rules/src/core" // Assuming DecompiledMethod is here
      )

      // MethodDocumentation holds detailed AI-generated docs for a method.
      type MethodDocumentation struct {
      	Description      string
      	Parameters       []ParamDoc
      	ReturnValue      string
      	UsageExample     string
      	ExceptionHandling string
      }
      type ParamDoc struct {
      	Name        string
      	Type        string
      	Description string
      }

      // GenerateSemanticMethodDoc generates detailed documentation for a single method.
      func GenerateSemanticMethodDoc(methodInfo *core.DecompiledMethod, classContext *core.DecompiledClass, config AIModelConfig) (*MethodDocumentation, error) {
      	// Placeholder: Implementation for semantic documentation generation
      	// 1. Create a highly specific prompt for the AI model, focusing on the method.
      	//    Include method signature, surrounding code, class context.
      	// 2. Call AI model.
      	// 3. Parse response into MethodDocumentation struct.
      	prompt := buildMethodPrompt(methodInfo, classContext)
      	_ = prompt // avoid unused variable error for placeholder
      	// aiResponse, err := callAIModel(config, prompt)
      	// if err != nil { return nil, err }
      	// return parseMethodAIResponse(aiResponse), nil
      	return &MethodDocumentation{Description: "Detailed docs for " + methodInfo.Name}, nil
      }

      func buildMethodPrompt(methodInfo *core.DecompiledMethod, classContext *core.DecompiledClass) string {
        // Construct a prompt asking for description, params, return, example, exceptions
        return "Generate detailed documentation for method " + methodInfo.Name + " in class " + classContext.Name + "..."
      }
      ```
  - [x] 1.7 开发Markdown格式文档输出模块
    - **描述:** 根据AI生成的文档内容和结构化信息，格式化输出为Markdown文件。应遵循PRD中定义的Markdown SDK文档格式规范。
    - **依赖:** 1.6
    - **代码示例 (Go):**
      ```go
      // src/output/markdown_formatter.go
      package output

      import (
      	"//Users/liujinliang/workspace/ai/ai_rules/src/ai" // Assuming MethodDocumentation is here
      	"//Users/liujinliang/workspace/ai/ai_rules/src/core"
        "bytes"
        "text/template"
      )

      // FormatToMarkdown converts class and method documentation to a Markdown string.
      func FormatToMarkdown(classInfo *core.DecompiledClass, methodDocs map[string]*ai.MethodDocumentation) (string, error) {
      	// Placeholder: Implementation for Markdown formatting
      	// Use text/template for robust Markdown generation based on a predefined template.
        const mdTemplate = `
      # Class: {{.ClassInfo.Name}}
      {{range .MethodDocs}}
      ## Method: {{.Name}}
      **Description:** {{.Description}}
      **Parameters:**
      {{range .Parameters}}
      - {{.Name}} ({{.Type}}): {{.Description}}
      {{end}}
      **Returns:** {{.ReturnValue}}
      **Example:**
      `+"```java"+`
      {{.UsageExample}}
      `+"```"+`
      **Exceptions:** {{.ExceptionHandling}}
      {{end}}
      `
        tmpl, err := template.New("markdownDoc").Parse(mdTemplate)
        if err != nil {
            return "", err
        }
        
        data := struct {
            ClassInfo  *core.DecompiledClass
            MethodDocs map[string]*ai.MethodDocumentation
        }{
            ClassInfo:  classInfo,
            MethodDocs: methodDocs,
        }

        var buf bytes.Buffer
        if err := tmpl.Execute(&buf, data); err != nil {
            return "", err
        }
        return buf.String(), nil
      }
      ```
  - [x] 1.8 开发OpenAPI 3.0规范文档输出模块
    - **描述:** 将JAR包中识别出的（通常是Web服务相关的）API接口信息，转换为OpenAPI 3.0规范的JSON或YAML文件。这需要映射Java类/方法到OpenAPI的路径、操作、参数、模式等。
    - **依赖:** 1.3, 1.4, 1.6 (for descriptions)
    - **代码示例 (Go):**
      ```go
      // src/output/openapi_formatter.go
      package output

      // OpenAPIObject represents the root of an OpenAPI 3.0 document.
      // This is a simplified version. A full implementation would use a library.
      type OpenAPIObject struct {
      	OpenAPI string `json:"openapi"`
      	Info    OpenAPIInfo `json:"info"`
      	Paths   map[string]PathItem `json:"paths"`
      	// ... other OpenAPI fields like components, servers
      }
      type OpenAPIInfo struct {
      	Title   string `json:"title"`
      	Version string `json:"version"`
      }
      type PathItem struct {
      	Get  *Operation `json:"get,omitempty"`
      	Post *Operation `json:"post,omitempty"`
      	// ... other HTTP methods
      }
      type Operation struct {
      	Summary     string `json:"summary,omitempty"`
      	Description string `json:"description,omitempty"`
      	Responses   map[string]Response `json:"responses"`
      	// ... parameters, requestBody
      }
      type Response struct {
      	Description string `json:"description"`
      	// ... content
      }

      // FormatToOpenAPI converts API information into an OpenAPI 3.0 structure.
      func FormatToOpenAPI(/* relevant API data from decompiled code */) (*OpenAPIObject, error) {
      	// Placeholder: Implementation for OpenAPI formatting
      	// This is complex and would involve mapping Java annotations (e.g., JAX-RS, Spring MVC)
      	// or conventions to OpenAPI constructs.
      	doc := &OpenAPIObject{
      		OpenAPI: "3.0.0",
      		Info:    OpenAPIInfo{Title: "Generated API", Version: "1.0.0"},
      		Paths:   make(map[string]PathItem),
      	}
      	return doc, nil
      }
      ```
  - [x] 1.9 开发JSON Schema格式文档输出模块
    - **描述:** 为JAR包中的数据结构（如POJOs, DTOs）生成JSON Schema定义。这有助于AI或其他系统理解这些数据对象的结构。
    - **依赖:** 1.3
    - **代码示例 (Go):**
      ```go
      // src/output/jsonschema_formatter.go
      package output

      // JSONSchema represents a JSON Schema definition.
      // Simplified for example.
      type JSONSchema struct {
      	SchemaType string `json:"type,omitempty"` // e.g., "object", "string", "integer"
      	Properties map[string]*JSONSchema `json:"properties,omitempty"`
      	Required   []string `json:"required,omitempty"`
      	// ... other JSON Schema fields like format, items, description
      }

      // FormatToJSONSchema converts a Java class structure to JSON Schema.
      func FormatToJSONSchema(/* decompiled class/field info */) (*JSONSchema, error) {
      	// Placeholder: Implementation for JSON Schema generation
      	// Map Java types to JSON Schema types.
      	// Handle nested objects, arrays, enums etc.
      	schema := &JSONSchema{
      		SchemaType: "object",
      		Properties: make(map[string]*JSONSchema),
      	}
      	return schema, nil
      }
      ```
  - [ ] 1.10 实现文档下载功能
    - **描述:** 提供一个API端点或UI功能，允许用户下载生成的Markdown, OpenAPI, 和JSON Schema文档。文档可以单个下载或打包下载（例如，一个ZIP文件包含所有生成的文档）。
    - **依赖:** 1.7, 1.8, 1.9
    - **代码示例 (Go):**
      ```go
      // src/api/handlers.go (or a new file like src/api/download_handler.go)
      package api

      import (
      	"net/http"
      )

      // HandleDownloadDoc serves generated documents for download.
      func HandleDownloadDoc(w http.ResponseWriter, r *http.Request) {
      	// Placeholder: Implementation for document download
      	// 1. Get document ID or type from request (e.g., /download?docId=xyz&type=markdown).
      	// 2. Retrieve the generated document content from storage or generate on-the-fly.
      	// 3. Set appropriate Content-Disposition and Content-Type headers.
      	// 4. Write document content to http.ResponseWriter.
      	docType := r.URL.Query().Get("type")
      	fileName := "document."
      	contentType := "application/octet-stream"

      	switch docType {
      	case "markdown":
      		fileName += "md"
      		contentType = "text/markdown"
      	case "openapi_json":
      		fileName += "json"
      		contentType = "application/json"
      	case "jsonschema":
      		fileName += "schema.json"
      		contentType = "application/json"
      	default:
      		http.Error(w, "Invalid document type", http.StatusBadRequest)
      		return
      	}

      	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
      	w.Header().Set("Content-Type", contentType)
      	// _, err := w.Write([]byte("Placeholder document content for " + docType))
      	// if err != nil { http.Error(w, "Failed to write document", http.StatusInternalServerError)}
      	// For actual implementation, fetch content based on docId and docType
        http.ServeContent(w, r, fileName, time.Now(), bytes.NewReader([]byte("Placeholder for "+fileName)))
      }
      ```
- [ ] 2.0 高级功能增强：提升处理能力与集成外部服务
  - [ ] 2.1 实现JAR包目录批量扫描功能
    - **描述:** 允许用户指定一个目录，系统将扫描该目录（可配置是否递归扫描子目录）下的所有JAR文件，并将它们加入到处理队列中。
    - **依赖:** 1.1 (for individual JAR processing logic)
    - **代码示例 (Go):**
      ```go
      // src/core/batch_scanner.go
      package core

      import (
      	"os"
      	"path/filepath"
      	"strings"
      )

      // ScanJarDirectory scans a directory for JAR files.
      func ScanJarDirectory(rootDir string, recursive bool) ([]string, error) {
      	var jarFiles []string
      	visitor := func(path string, info os.FileInfo, err error) error {
      		if err != nil {
      			return err // Propagate errors
      		}
      		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".jar") {
      			jarFiles = append(jarFiles, path)
      		}
      		return nil
      	}

      	if recursive {
      		err := filepath.Walk(rootDir, visitor)
      		return jarFiles, err
      	} else {
      		entries, err := os.ReadDir(rootDir)
      		if err != nil {
      			return nil, err
      		}
      		for _, entry := range entries {
      			info, _ := entry.Info() // Error can be ignored here for simplicity in example
      			if err := visitor(filepath.Join(rootDir, entry.Name()), info, nil); err != nil {
      			    // Decide how to handle individual file errors, e.g. log and continue
      			}
      		}
      		return jarFiles, nil
      	}
      }
      ```
  - [ ] 2.2 实现多线程并发分析
    - **描述:** 对批量扫描到的JAR文件，使用多线程（goroutines）进行并发的反编译和AI文档生成处理，以显著提高整体处理效率。需要考虑资源控制（如并发数限制）和错误处理。
    - **依赖:** 2.1, (1.2 to 1.6 for processing pipeline)
    - **代码示例 (Go):**
      ```go
      // src/core/concurrent_processor.go
      package core

      import (
      	"sync"
      	"//Users/liujinliang/workspace/ai/ai_rules/src/ai" // For AIModelConfig
      )

      // ProcessJarsConcurrently handles multiple JAR files using goroutines.
      func ProcessJarsConcurrently(jarPaths []string, maxConcurrency int, aiConfig ai.AIModelConfig) (map[string] /*resultType*/ interface{}, map[string]error) {
      	var wg sync.WaitGroup
      	results := make(map[string]interface{}) // Replace interface{} with actual result type
      	errors := make(map[string]error)
      	guard := make(chan struct{}, maxConcurrency) // Semaphore to limit concurrency
      	mu := sync.Mutex{} // To safely write to maps

      	for _, jarPath := range jarPaths {
      		wg.Add(1)
      		guard <- struct{}{}
      		go func(path string) {
      			defer wg.Done()
      			defer func() { <-guard }()

      			// Simulate processing pipeline for a single JAR
      			// In reality, this would call Upload, Parse, Decompile, Analyze, GenerateAIDoc, Format
      			// For example:
      			// reader, err := UploadFromLocal(path)
      			// if err != nil { mu.Lock(); errors[path] = err; mu.Unlock(); return }
      			// parsedInfo, err := ParseJarStream(reader, size) // Need to get size
      			// ... and so on

      			// Placeholder for actual processing result
      			result := "Processed: " + path
      			var errProcess error = nil // Placeholder for error

      			mu.Lock()
      			if errProcess != nil {
      				errors[path] = errProcess
      			} else {
      				results[path] = result
      			}
      			mu.Unlock()
      		}(jarPath)
      	}
      	wg.Wait()
      	return results, errors
      }
      ```
  - [ ] 2.3 开发实时扫描进度和处理状态显示
    - **描述:** 为批量处理任务提供一个用户界面（Web）或API端点，能够实时显示当前扫描进度（如已扫描文件数/总文件数）、每个JAR包的处理状态（等待、处理中、完成、失败）、以及整体任务的预计剩余时间。
    - **依赖:** 2.2
    - **代码示例 (Go):**
      ```go
      // src/api/status_handler.go
      package api

      import (
      	"encoding/json"
      	"net/http"
      	"sync"
      )

      // BatchProcessStatus holds the current status of a batch operation.
      type BatchProcessStatus struct {
      	TotalFiles     int `json:"totalFiles"`
      	ProcessedFiles int `json:"processedFiles"`
      	FailedFiles    int `json:"failedFiles"`
      	// Individual file statuses could be a map[string]string (fileName -> status)
      	FileStatuses map[string]string `json:"fileStatuses"`
      	// ... other fields like estimatedTimeRemaining
      }

      var globalStatus BatchProcessStatus
      var statusMutex sync.RWMutex

      // UpdateBatchStatus allows concurrent processors to update the global status.
      // This is a simplified example; a real system might use channels or a dedicated status manager.
      func UpdateBatchStatus(updateFunc func(status *BatchProcessStatus)) {
      	statusMutex.Lock()
      	defer statusMutex.Unlock()
      	updateFunc(&globalStatus)
      }

      // HandleBatchStatus provides the current batch processing status via API.
      func HandleBatchStatus(w http.ResponseWriter, r *http.Request) {
      	statusMutex.RLock()
      	defer statusMutex.RUnlock()
      	w.Header().Set("Content-Type", "application/json")
      	if err := json.NewEncoder(w).Encode(globalStatus); err != nil {
      		http.Error(w, "Failed to encode status", http.StatusInternalServerError)
      	}
      }
      
      // Example of how a worker might update status (conceptual)
      /*
      func workerProcessJar(jarPath string) {
          // ... processing ...
          UpdateBatchStatus(func(s *BatchProcessStatus) {
              s.ProcessedFiles++
              s.FileStatuses[jarPath] = "Completed"
          })
          // on error:
          // UpdateBatchStatus(func(s *BatchProcessStatus) {
          //     s.FailedFiles++
          //     s.FileStatuses[jarPath] = "Failed: some error"
          // })
      }
      */
      ```
  - [ ] 2.4 实现批量处理结果的统一管理和分组导出功能
    - **描述:** 对批量处理生成的文档结果，提供统一的管理界面或功能。用户可以按项目、批次或其他自定义标签对结果进行分组，并支持将选定组的文档打包导出（例如，一个ZIP文件包含该组所有JAR的Markdown文档）。
    - **依赖:** 2.3 (for status and results), 1.10 (for individual download logic)
    - **代码示例 (Go):**
      ```go
      // src/core/result_manager.go
      package core

      import (
      	"archive/zip"
      	"io"
      	"os"
      )

      // ProcessedResult represents the outcome of processing a single JAR.
      type ProcessedResult struct {
      	JarPath         string
      	MarkdownDocPath string // Path to generated Markdown file
      	OpenAPIDocPath  string // Path to generated OpenAPI file
      	// ... other document types
      	Status          string // e.g., "Success", "Failed"
      	ErrorMessage    string
      }

      // BatchResultStore stores and manages results from batch processing.
      // This could be in-memory for simplicity, or backed by a database.
      var BatchResultStore = struct {
      	Results map[string][]ProcessedResult // Key: batchId or groupName
      }{
      	Results: make(map[string][]ProcessedResult),
      }

      // ExportGroupedResults creates a ZIP archive of results for a given group.
      func ExportGroupedResults(groupName string, writer io.Writer) error {
      	// Placeholder: Implementation for exporting grouped results
      	results, ok := BatchResultStore.Results[groupName]
      	if !ok {
      		// return errors.New("group not found")
      	}
      	zipWriter := zip.NewWriter(writer)
      	defer zipWriter.Close()

      	for _, result := range results {
      		if result.Status == "Success" {
      			// Add MarkdownDocPath to zip
      			// Add OpenAPIDocPath to zip
      			// Example for one file type:
      			fw, err := zipWriter.Create(filepath.Base(result.MarkdownDocPath)) // Use appropriate path in zip
      			if err != nil { /* handle error */ continue }
      			fileData, err := os.ReadFile(result.MarkdownDocPath)
      			if err != nil { /* handle error */ continue }
      			_, err = fw.Write(fileData)
      			if err != nil { /* handle error */ }
      		}
      	}
      	return nil
      }
      ```
  - [ ] 2.5 添加文件大小、修改时间、文件名模式的过滤功能
    - **描述:** 在批量扫描JAR包时，允许用户设置过滤条件，例如：只处理小于特定大小的JAR文件，只处理在某个日期之后修改过的文件，或只处理文件名匹配特定模式（如 `*-api.jar`）的文件。
    - **依赖:** 2.1
    - **代码示例 (Go):**
      ```go
      // src/core/batch_scanner.go (extend ScanJarDirectory or add new func)
      package core

      import (
      	"os"
      	"path/filepath"
      	"regexp"
      	"strings"
      	"time"
      )

      // FilterOptions defines criteria for filtering JAR files.
      type FilterOptions struct {
      	MaxSize        int64     // Max file size in bytes (0 for no limit)
      	MinModTime     time.Time // Minimum modification time (zero value for no limit)
      	FileNamePattern string    // Regex pattern for filename (empty for no limit)
      }

      // ScanJarDirectoryWithFilters scans a directory for JAR files applying filters.
      func ScanJarDirectoryWithFilters(rootDir string, recursive bool, filters FilterOptions) ([]string, error) {
      	var jarFiles []string
      	var nameRegex *regexp.Regexp
      	if filters.FileNamePattern != "" {
      		var err error
      		nameRegex, err = regexp.Compile(filters.FileNamePattern)
      		if err != nil {
      			return nil, err // Invalid regex pattern
      		}
      	}

      	visitor := func(path string, info os.FileInfo, err error) error {
      		if err != nil { return err }
      		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".jar") {
      			// Apply filters
      			if filters.MaxSize > 0 && info.Size() > filters.MaxSize {
      				return nil // Skip file, too large
      			}
      			if !filters.MinModTime.IsZero() && info.ModTime().Before(filters.MinModTime) {
      				return nil // Skip file, too old
      			}
      			if nameRegex != nil && !nameRegex.MatchString(info.Name()) {
      				return nil // Skip file, name doesn't match
      			}
      			jarFiles = append(jarFiles, path)
      		}
      		return nil
      	}
      	// ... (rest of the walk logic from ScanJarDirectory)
      	if recursive {
      		err := filepath.Walk(rootDir, visitor)
      		return jarFiles, err
      	} else {
      		// ... non-recursive logic with visitor ...
          return jarFiles, nil // Simplified for brevity
      	}
        return jarFiles, nil
      }
      ```
  - [ ] 2.6 实现中断后继续处理未完成文件的功能
    - **描述:** 对于长时间运行的批量处理任务，如果中途意外中断（如程序崩溃、服务器重启），系统应能记录已处理和未处理的文件列表。在任务重启后，可以跳过已成功处理的文件，仅继续处理未完成或失败的文件，以节省时间和资源。
    - **依赖:** 2.2, 2.3 (for tracking status)
    - **代码示例 (Go):**
      ```go
      // src/core/resumable_processor.go
      package core

      import (
      	"encoding/json"
      	"os"
      	"//Users/liujinliang/workspace/ai/ai_rules/src/ai" // For AIModelConfig
      )

      const progressFile = "batch_progress.json"

      // BatchProgress tracks the state of a batch job for resumability.
      type BatchProgress struct {
      	AllFiles      []string          `json:"allFiles"`
      	Processed     map[string]bool   `json:"processed"` // filePath -> true if successfully processed
      	Failed        map[string]string `json:"failed"`    // filePath -> error message
      	// Add other necessary state, e.g., current batch ID
      }

      // LoadProgress loads the last saved batch progress.
      func LoadProgress() (*BatchProgress, error) {
      	data, err := os.ReadFile(progressFile)
      	if os.IsNotExist(err) {
      		return &BatchProgress{Processed: make(map[string]bool), Failed: make(map[string]string)}, nil
      	}
      	if err != nil {
      		return nil, err
      	}
      	var progress BatchProgress
      	err = json.Unmarshal(data, &progress)
      	return &progress, err
      }

      // SaveProgress saves the current batch progress.
      func (p *BatchProgress) SaveProgress() error {
      	data, err := json.MarshalIndent(p, "", "  ")
      	if err != nil {
      		return err
      	}
      	return os.WriteFile(progressFile, data, 0644)
      }

      // ProcessJarsResumable processes JARs, skipping already completed ones.
      func ProcessJarsResumable(jarPaths []string, maxConcurrency int, aiConfig ai.AIModelConfig) error {
      	progress, err := LoadProgress()
      	if err != nil {
      		// log.Printf("Error loading progress, starting fresh: %v", err)
      		progress = &BatchProgress{Processed: make(map[string]bool), Failed: make(map[string]string)}
      	}
      	progress.AllFiles = jarPaths // Update with current list, or merge if needed

      	var filesToProcess []string
      	for _, path := range jarPaths {
      		if !progress.Processed[path] { // Only process if not already successfully processed
      			filesToProcess = append(filesToProcess, path)
      		}
      	}
      	
      	// results, errors := ProcessJarsConcurrently(filesToProcess, maxConcurrency, aiConfig)
        // For each result/error from ProcessJarsConcurrently:
        //   if error:
        //     progress.Failed[filePath] = error.Error()
        //   else:
        //     progress.Processed[filePath] = true
        //     delete(progress.Failed, filePath) // Remove from failed if it was previously failed
        //   progress.SaveProgress() // Save after each file or periodically

      	return progress.SaveProgress() // Final save
      }
      ```
    - [ ] 2.7 集成Maven Central、阿里云等仓库API，用于抓取JAR包
      - **描述:** 实现与主流Maven仓库（如Maven Central, JCenter, 阿里云效Maven仓库等）的API集成，允许用户通过GroupId, ArtifactId, Version (GAV) 坐标直接从这些仓库搜索和下载JAR包进行分析。
      - **依赖:** 1.1 (for handling downloaded JARs)
      - **代码示例 (Go):**
        ```go
        // src/integrations/maven_repository.go
        package integrations

        import (
        	"fmt"
        	"io"
        	"net/http"
        	"os"
        )

        // Repository represents a Maven repository configuration.
        type Repository struct {
        	Name    string
        	BaseURL string // e.g., "https://repo1.maven.org/maven2/"
        }

        var DefaultRepositories = []Repository{
        	{Name: "Maven Central", BaseURL: "https://repo1.maven.org/maven2/"},
        	{Name: "Aliyun Maven", BaseURL: "https://maven.aliyun.com/repository/public/"},
        	// Add more repositories here
        }

        // FetchJarFromRepository downloads a JAR from a specified repository given GAV coordinates.
        func FetchJarFromRepository(repo Repository, groupID, artifactID, version, targetDir string) (string, error) {
        	// Construct the URL, e.g., com/example/my-lib/1.0/my-lib-1.0.jar
        	// This needs to handle groupID with dots (com.example -> com/example)
        	// and potentially classifiers or packaging types if more advanced.
        	jarFileName := fmt.Sprintf("%s-%s.jar", artifactID, version)
        	relativePath := fmt.Sprintf("%s/%s/%s/%s",
        		strings.ReplaceAll(groupID, ".", "/"),
        		artifactID,
        		version,
        		jarFileName,
        	)
        	jarURL := repo.BaseURL + relativePath

        	resp, err := http.Get(jarURL)
        	if err != nil {
        		return "", err
        	}
        	defer resp.Body.Close()

        	if resp.StatusCode != http.StatusOK {
        		return "", fmt.Errorf("failed to download JAR from %s: %s", jarURL, resp.Status)
        	}

        	filePath := filepath.Join(targetDir, jarFileName)
        	out, err := os.Create(filePath)
        	if err != nil {
        		return "", err
        	}
        	defer out.Close()

        	_, err = io.Copy(out, resp.Body)
        	if err != nil {
        		return "", err
        	}
        	return filePath, nil
        }
        ```
    - [ ] 2.8 实现POM文件解析，获取依赖关系和元数据
      - **描述:** 如果用户上传的是一个Maven项目（包含`pom.xml`）或者从仓库下载的JAR有关联的POM文件，系统应能解析POM文件。从中提取项目元数据（如GAV, name, description, licenses）、依赖项列表（dependencies）、父POM信息（parent）、构建配置（build plugins, properties）等。这些信息可以丰富生成的SDK文档。
      - **依赖:** 1.2 (if JAR is part of a project), 2.7 (if fetching POM with JAR)
      - **代码示例 (Go):**
        ```go
        // src/core/pom_parser.go
        package core

        import (
        	"encoding/xml"
        	"io/ioutil"
        )

        // Project is a simplified representation of a pom.xml structure.
        // This would need to be much more comprehensive for full POM parsing.
        type Project struct {
        	XMLName        xml.Name    `xml:"project"`
        	GroupID        string      `xml:"groupId"`
        	ArtifactID     string      `xml:"artifactId"`
        	Version        string      `xml:"version"`
        	Name           string      `xml:"name"`
        	Description    string      `xml:"description"`
        	Dependencies   []Dependency `xml:"dependencies>dependency"`
        	// Add more fields: Parent, Build, Properties, Licenses, etc.
        }

        type Dependency struct {
        	GroupID    string `xml:"groupId"`
        	ArtifactID string `xml:"artifactId"`
        	Version    string `xml:"version"`
        	Scope      string `xml:"scope"`
        }

        // ParsePOMFile reads and parses a pom.xml file.
        func ParsePOMFile(pomPath string) (*Project, error) {
        	xmlFile, err := ioutil.ReadFile(pomPath)
        	if err != nil {
        		return nil, err
        	}

        	var project Project
        	err = xml.Unmarshal(xmlFile, &project)
        	if err != nil {
        		return nil, err
        	}
        	return &project, nil // Return the parsed project
        }
        ```
    - [ ] 2.9 支持解析 `MANIFEST.MF` 文件，提取元数据
    - [ ] 2.10 实现对Android APK文件的初步支持 (提取 `classes.dex` 并转换为JAR)
    - [ ] 2.11 支持从URL直接下载和处理单个JAR文件

### 3. AI模型集成与配置 (Parent Task 3.0)

- **目标:** 灵活集成和配置多种AI大模型，支持用户自定义模型参数和API凭证，以适应不同的分析需求和部署环境。
- **子任务:**
  - [ ] 3.1 设计AI模型交互的核心接口 (如 `AIClient`)
    - **描述:** 定义一个通用的AI客户端接口，抽象不同AI模型（如OpenAI, Gemini, Claude等）的通用操作，例如代码解释、文档生成、API规范提取等。这样可以方便地替换和扩展AI模型。
    - **依赖:** 无
    - **代码示例 (Go):**
      ```go
      // src/ai/client_interface.go
      package ai

      import (
      	"context"
      )

      // AIClient defines the interface for interacting with an AI model.
      type AIClient interface {
      	// GenerateDocumentation analyzes code and generates Markdown documentation.
      	GenerateDocumentation(ctx context.Context, codeSnippet string, language string, promptArgs map[string]string) (string, error)

      	// GenerateOpenAPISpec analyzes code (especially API endpoints) and generates an OpenAPI specification (JSON or YAML).
      	GenerateOpenAPISpec(ctx context.Context, codeSnippet string, language string, promptArgs map[string]string) (string, error)

      	// ExplainCode provides a natural language explanation of a given code snippet.
      	ExplainCode(ctx context.Context, codeSnippet string, language string, promptArgs map[string]string) (string, error)

      	// SuggestCodeImprovements analyzes code and suggests potential improvements or refactorings.
      	SuggestCodeImprovements(ctx context.Context, codeSnippet string, language string, promptArgs map[string]string) (string, error)

      	// GetModelInfo returns information about the underlying AI model.
      	GetModelInfo() string
      }

      // PromptArguments can be used to pass additional context or instructions to the AI model for a specific task.
      // For example, for GenerateDocumentation, it might include desired documentation style or focus points.
      ```
  - [ ] 3.2 支持自定义AI模型API Key和Endpoint
    - **描述:** 允许用户在调用AI服务时，配置自己的API Key和模型服务的Endpoint。这对于使用私有部署的AI模型或特定区域的AI服务非常重要。
    - **依赖:** 1.7
    - **代码示例 (Go):**
      ```go
      // src/ai/config.go
      package ai

      import (
      	"os"
      )

      // AIModelConfig holds configuration for an AI model.
      type AIModelConfig struct {
      	APIKey    string `json:"apiKey"`
      	Endpoint  string `json:"endpoint"` // Optional, defaults to provider's public endpoint
      	ModelName string `json:"modelName"` // e.g., "gpt-3.5-turbo", "gemini-pro"
      	// Add other common parameters like temperature, maxTokens later (Task 3.4)
      }

      // LoadConfigFromEnv loads AI configuration from environment variables.
      // Example env vars: MYAPP_AI_API_KEY, MYAPP_AI_ENDPOINT, MYAPP_AI_MODEL_NAME
      func LoadConfigFromEnv(prefix string) AIModelConfig {
      	return AIModelConfig{
      		APIKey:    os.Getenv(prefix + "_AI_API_KEY"),
      		Endpoint:  os.Getenv(prefix + "_AI_ENDPOINT"),
      		ModelName: os.Getenv(prefix + "_AI_MODEL_NAME"),
      	}
      }

      // NewAIModelConfig creates a new AI model configuration.
      func NewAIModelConfig(apiKey, endpoint, modelName string) AIModelConfig {
          return AIModelConfig{
              APIKey: apiKey,
              Endpoint: endpoint,
              ModelName: modelName,
          }
      }
      ```
  - [ ] 3.3 实现多AI模型支持与切换 (如 OpenAI, Gemini, Claude, Kimi)
    - **描述:** 系统应能集成并支持多种主流AI大模型，如OpenAI的GPT系列、Google的Gemini、Anthropic的Claude以及国内的Kimi等。用户可以根据需求选择或切换不同的模型进行代码分析和文档生成。
    - **依赖:** 3.2
    - **代码示例 (Go):**
      ```go
      // src/ai/client_factory.go (renamed from client.go for clarity)
      package ai

      import (
      	"context"
      	"errors"
      	"fmt"
      )

      // NewClient creates an AIClient based on the provided configuration.
      // This acts as a factory for different AI client implementations.
      func NewClient(config AIModelConfig) (AIClient, error) {
      	switch config.ModelName { // Or use a more sophisticated provider detection based on Endpoint or ModelName patterns
      	case "gpt-3.5-turbo", "gpt-4", "gpt-4o":
      		// return NewOpenAIClient(config) // Placeholder for actual OpenAI client implementation
      		return &DummyAIClient{Provider: "OpenAI", Config: config, Info: "OpenAI GPT Model"}, nil
      	case "gemini-pro", "gemini-1.5-pro":
      		// return NewGeminiClient(config) // Placeholder for actual Gemini client implementation
      		return &DummyAIClient{Provider: "Gemini", Config: config, Info: "Google Gemini Model"}, nil
          case "claude-3-opus", "claude-3-sonnet", "claude-3-haiku":
              // return NewClaudeClient(config) // Placeholder
              return &DummyAIClient{Provider: "Anthropic", Config: config, Info: "Anthropic Claude Model"}, nil
          case "kimi-moonshot-v1-8k", "kimi-moonshot-v1-32k", "kimi-moonshot-v1-128k": // Example Kimi model names
              // return NewKimiClient(config) // Placeholder
              return &DummyAIClient{Provider: "Moonshot", Config: config, Info: "Moonshot Kimi Model"}, nil
      	default:
      		// Attempt to infer from endpoint if known, or allow registration of custom clients
      		return nil, errors.New("unsupported or unknown AI model: " + config.ModelName)
      	}
      }

      // DummyAIClient is a placeholder for actual AI client implementations.
      // It now implements the full AIClient interface from Task 3.1.
      type DummyAIClient struct {
          Provider string
          Config   AIModelConfig
          Info     string
      }

      func (c *DummyAIClient) GenerateDocumentation(ctx context.Context, codeSnippet string, language string, promptArgs map[string]string) (string, error) {
      	return fmt.Sprintf("[%s - %s] Documentation for %s code (args: %v): ...", c.Provider, c.Config.ModelName, language, promptArgs), nil
      }

      func (c *DummyAIClient) GenerateOpenAPISpec(ctx context.Context, codeSnippet string, language string, promptArgs map[string]string) (string, error) {
      	return fmt.Sprintf("[%s - %s] OpenAPI spec for %s code (args: %v): ...", c.Provider, c.Config.ModelName, language, promptArgs), nil
      }

      func (c *DummyAIClient) ExplainCode(ctx context.Context, codeSnippet string, language string, promptArgs map[string]string) (string, error) {
      	return fmt.Sprintf("[%s - %s] Explanation for %s code (args: %v): ...", c.Provider, c.Config.ModelName, language, promptArgs), nil
      }

      func (c *DummyAIClient) SuggestCodeImprovements(ctx context.Context, codeSnippet string, language string, promptArgs map[string]string) (string, error) {
      	return fmt.Sprintf("[%s - %s] Suggestions for %s code (args: %v): ...", c.Provider, c.Config.ModelName, language, promptArgs), nil
      }

      func (c *DummyAIClient) GetModelInfo() string {
      	return fmt.Sprintf("Provider: %s, Model: %s, Endpoint: %s", c.Provider, c.Config.ModelName, c.Config.Endpoint)
      }
      ```
  - [ ] 3.4 实现AI模型参数的灵活配置 (temperature, top_p, max_tokens)
    - **描述:** 用户应能调整AI模型的核心参数，如`temperature`（控制创造性与确定性）、`top_p`（控制核心词汇选择范围）和`max_tokens`（控制生成内容的最大长度），以便针对不同任务和代码复杂度微调AI的输出质量。
    - **依赖:** 3.2, 3.3
    - **代码示例 (Go):**
      ```go
      // src/ai/config.go (extend AIModelConfig)
      package ai

      // AIModelConfig (extended from Task 3.2)
      // Ensure these fields are added to the AIModelConfig struct defined in Task 3.2.
      // For clarity, the extended struct is shown here:
      /*
      type AIModelConfig struct {
          APIKey      string  `json:"apiKey"`
          Endpoint    string  `json:"endpoint"`       // Optional, defaults to provider's public endpoint
          ModelName   string  `json:"modelName"`      // e.g., "gpt-3.5-turbo", "gemini-pro"
          Temperature *float32 `json:"temperature,omitempty"` // Pointer to allow zero value (or provider default) vs. explicitly set
          TopP        *float32 `json:"topP,omitempty"`        // Pointer for similar reasons
          MaxTokens   *int     `json:"maxTokens,omitempty"`   // Pointer for similar reasons
          // Potentially add other common params: presence_penalty, frequency_penalty, stop_sequences etc.
      }
      */

      // Modify NewAIModelConfig from Task 3.2 to include these new parameters, potentially with defaults or as options.
      // func NewAIModelConfig(apiKey, endpoint, modelName string, temp *float32, topP *float32, maxTokens *int) AIModelConfig { ... }

      // Client implementations (e.g., OpenAIClient, GeminiClient from Task 3.3) would use these parameters.
      // Conceptual example within a hypothetical OpenAI client:
      // func (c *OpenAIClient) buildChatCompletionRequest(messages []openai.ChatCompletionMessage) openai.ChatCompletionRequest {
      //    request := openai.ChatCompletionRequest{
      //        Model:    c.config.ModelName,
      //        Messages: messages,
      //    }
      //    if c.config.Temperature != nil {
      //        request.Temperature = *c.config.Temperature
      //    }
      //    if c.config.TopP != nil {
      //        request.TopP = *c.config.TopP
      //    }
      //    if c.config.MaxTokens != nil {
      //        request.MaxTokens = *c.config.MaxTokens
      //    }
      //    return request
      // }
      ```
  - [ ] 3.5 支持通过配置文件或环境变量加载AI配置
    - **描述:** 除了代码中硬编码或通过命令行参数传递，系统应支持从外部配置文件（如JSON, YAML, TOML格式）或通过环境变量加载AI模型的配置信息（API Key, Endpoint, 模型名称, 参数等），方便部署和管理。
    - **依赖:** 3.2, 3.4
    - **代码示例 (Go):**
      ```go
      // src/config/loader.go
      package config

      import (
      	"encoding/json"
      	"fmt"
      	"os"
      	"path/filepath"
      	"strconv"
      	"strings"
      	"//Users/liujinliang/workspace/ai/ai_rules/src/ai" // Assuming ai.AIModelConfig is defined there
      	"gopkg.in/yaml.v3" // Example for YAML, add to go.mod: go get gopkg.in/yaml.v3
      	// Consider adding TOML support: go get github.com/BurntSushi/toml
      )

      // AppConfig holds the overall application configuration.
      type AppConfig struct {
      	AI ai.AIModelConfig `json:"ai" yaml:"ai" toml:"ai"`
      	// Add other app configurations here, e.g., BatchProcessing, Logging, DefaultOutputPath
      }

      // LoadConfig loads configuration from a file and environment variables.
      // Env vars override file settings. File path can be JSON, YAML, or TOML.
      func LoadConfig(filePath string, envPrefix string) (*AppConfig, error) {
      	config := &AppConfig{
            // Initialize with default AI config if necessary, or ensure AIModelConfig has sensible zero-values
            AI: ai.AIModelConfig{},
        }

      	// 1. Load from file (if path provided)
      	if filePath != "" {
      		data, err := os.ReadFile(filePath)
      		if err != nil {
      			return nil, fmt.Errorf("failed to read config file %s: %w", filePath, err)
      		}
      		fileExt := strings.ToLower(filepath.Ext(filePath))
      		switch fileExt {
      		case ".json":
      			err = json.Unmarshal(data, config)
      		case ".yaml", ".yml":
      			err = yaml.Unmarshal(data, config)
      		// case ".toml":
      		// 	err = toml.Unmarshal(data, config) // If TOML support is added
      		default:
      			return nil, fmt.Errorf("unsupported config file format: %s", fileExt)
      		}
      		if err != nil {
      			return nil, fmt.Errorf("failed to unmarshal config file %s: %w", filePath, err)
      		}
      	}

      	// 2. Override with environment variables
      	// Use helper from ai package (Task 3.2) for basic AI fields
      	envAIConfigBase := ai.LoadConfigFromEnv(envPrefix) 
      	if envAIConfigBase.APIKey != "" {
      		config.AI.APIKey = envAIConfigBase.APIKey
      	}
      	if envAIConfigBase.Endpoint != "" {
      		config.AI.Endpoint = envAIConfigBase.Endpoint
      	}
      	if envAIConfigBase.ModelName != "" {
      		config.AI.ModelName = envAIConfigBase.ModelName
      	}

      	// Override specific AI parameters (Temperature, TopP, MaxTokens)
      	if tempStr := os.Getenv(envPrefix + "_AI_TEMPERATURE"); tempStr != "" {
      		if tempF64, err := strconv.ParseFloat(tempStr, 32); err == nil {
      			tempF32 := float32(tempF64)
      			config.AI.Temperature = &tempF32
      		} // else: log warning about parse error?
      	}
      	if topPStr := os.Getenv(envPrefix + "_AI_TOP_P"); topPStr != "" {
      		if topPF64, err := strconv.ParseFloat(topPStr, 32); err == nil {
      			topPF32 := float32(topPF64)
      			config.AI.TopP = &topPF32
      		}
      	}
      	if maxTokensStr := os.Getenv(envPrefix + "_AI_MAX_TOKENS"); maxTokensStr != "" {
      		if maxTokensInt, err := strconv.Atoi(maxTokensStr); err == nil {
      			config.AI.MaxTokens = &maxTokensInt
      		}
      	}

      	return config, nil
      }
      ```
  - [ ] 3.6 实现AI请求的重试与超时机制
    - **描述:** 与AI模型交互时，网络波动或模型服务暂时不可用可能导致请求失败。应实现合理的重试机制（如指数退避）和超时控制，以提高系统的健壮性。
    - **依赖:** 3.1
    - **代码示例 (Go):**
      ```go
      // src/ai/retry_handler.go
      package ai

      import (
      	"context"
      	"time"
      	// "log"
      )

      // RetryableFunc is a function that can be retried.
      type RetryableFunc func(ctx context.Context) (string, error)

      // ExecuteWithRetry executes a function with a retry mechanism.
      func ExecuteWithRetry(ctx context.Context, fn RetryableFunc, maxRetries int, initialDelay time.Duration, timeoutPerTry time.Duration) (string, error) {
      	var result string
      	var err error
      	delay := initialDelay

      	for i := 0; i < maxRetries; i++ {
      		retryCtx, cancel := context.WithTimeout(ctx, timeoutPerTry)
      		defer cancel() // Ensure cancel is called on all paths

      		result, err = fn(retryCtx)
      		if err == nil {
      			return result, nil // Success
      		}

      		// log.Printf("Attempt %d failed: %v. Retrying in %v...", i+1, err, delay)

      		select {
      		case <-ctx.Done(): // Overall context cancelled
      			return "", ctx.Err()
      		case <-time.After(delay):
      			// Continue to next retry
      		}
      		delay *= 2 // Exponential backoff
      	}
      	return "", err // All retries failed
      }

      // Example usage within an AIClient method:
      // func (c *ConcreteAIClient) GenerateDocumentation(ctx context.Context, codeSnippet string, language string, promptArgs map[string]string) (string, error) {
      // 	retryable := func(innerCtx context.Context) (string, error) {
      // 		// Actual API call logic here using innerCtx
      // 		return c.actualAPICall(innerCtx, codeSnippet, language, promptArgs)
      // 	}
      // 	return ExecuteWithRetry(ctx, retryable, c.config.MaxRetries, c.config.InitialDelay, c.config.TimeoutPerTry)
      // }
      ```
  - [ ] 3.7 实现AI交互的日志记录 (请求、响应、错误)
    - **描述:** 详细记录与AI模型交互的每个环节，包括发送的请求（或其摘要，注意脱敏）、模型返回的完整响应、发生的任何错误以及重试次数等。这对于调试问题、监控模型性能和成本分析至关重要。
    - **依赖:** 3.1, 3.6
    - **代码示例 (Go):**
      ```go
      // src/ai/logging_interceptor.go
      package ai

      import (
      	"context"
      	"log"
      	"time"
      )

      // LoggingAIClientDecorator wraps an AIClient to add logging.
      type LoggingAIClientDecorator struct {
      	Wrapped AIClient
      	// Logger can be a more sophisticated logger, e.g., from zerolog, zap
      }

      // NewLoggingAIClientDecorator creates a new logging decorator.
      func NewLoggingAIClientDecorator(client AIClient) AIClient {
      	return &LoggingAIClientDecorator{Wrapped: client}
      }

      func (d *LoggingAIClientDecorator) logInteraction(methodName string, request interface{}, response interface{}, err error, startTime time.Time) {
      	duration := time.Since(startTime)
      	if err != nil {
      		log.Printf("[AI_LOG] Method: %s, Duration: %v, Error: %v, Request: %+v", methodName, duration, err, request)
      	} else {
      		log.Printf("[AI_LOG] Method: %s, Duration: %v, Success, Request: %+v, Response: %+v", methodName, duration, request, response)
      	}
      	// For production, ensure sensitive parts of request/response (like API keys in headers or PII in content) are masked.
      }

      // Implement AIClient interface methods
      func (d *LoggingAIClientDecorator) GenerateDocumentation(ctx context.Context, codeSnippet string, language string, promptArgs map[string]string) (string, error) {
      	startTime := time.Now()
      	// Create a representation of the request for logging (can be a struct or map)
      	reqDetail := map[string]interface{}{"codeSnippetLength": len(codeSnippet), "language": language, "promptArgs": promptArgs}
      	resp, err := d.Wrapped.GenerateDocumentation(ctx, codeSnippet, language, promptArgs)
      	d.logInteraction("GenerateDocumentation", reqDetail, resp, err, startTime)
      	return resp, err
      }

      func (d *LoggingAIClientDecorator) GenerateOpenAPISpec(ctx context.Context, codeSnippet string, language string, promptArgs map[string]string) (string, error) {
      	startTime := time.Now()
      	reqDetail := map[string]interface{}{"codeSnippetLength": len(codeSnippet), "language": language, "promptArgs": promptArgs}
      	resp, err := d.Wrapped.GenerateOpenAPISpec(ctx, codeSnippet, language, promptArgs)
      	d.logInteraction("GenerateOpenAPISpec", reqDetail, resp, err, startTime)
      	return resp, err
      }

      func (d *LoggingAIClientDecorator) ExplainCode(ctx context.Context, codeSnippet string, language string, promptArgs map[string]string) (string, error) {
      	startTime := time.Now()
      	reqDetail := map[string]interface{}{"codeSnippetLength": len(codeSnippet), "language": language, "promptArgs": promptArgs}
      	resp, err := d.Wrapped.ExplainCode(ctx, codeSnippet, language, promptArgs)
      	d.logInteraction("ExplainCode", reqDetail, resp, err, startTime)
      	return resp, err
      }

      func (d *LoggingAIClientDecorator) SuggestCodeImprovements(ctx context.Context, codeSnippet string, language string, promptArgs map[string]string) (string, error) {
      	startTime := time.Now()
      	reqDetail := map[string]interface{}{"codeSnippetLength": len(codeSnippet), "language": language, "promptArgs": promptArgs}
      	resp, err := d.Wrapped.SuggestCodeImprovements(ctx, codeSnippet, language, promptArgs)
      	d.logInteraction("SuggestCodeImprovements", reqDetail, resp, err, startTime)
      	return resp, err
      }

      func (d *LoggingAIClientDecorator) GetModelInfo() string {
      	// Logging for GetModelInfo might be less critical or logged differently
      	return d.Wrapped.GetModelInfo()
      }
      ```
  - [ ] 3.8 支持Prompt模板化与管理
    - **描述:** 为不同AI任务（如生成文档、解释代码、提取API）和不同编程语言设计和管理Prompt模板。模板中可以包含占位符，由系统在运行时填充具体代码、语言、用户偏好等信息。这有助于优化AI输出质量和一致性。
    - **依赖:** 3.1
    - **代码示例 (Go):**
      ```go
      // src/ai/prompt_manager.go
      package ai

      import (
      	"bytes"
      	"fmt"
      	"text/template"
      )

      // PromptManager handles loading and rendering of prompt templates.
      type PromptManager struct {
      	templates map[string]*template.Template
      }

      // NewPromptManager creates a new prompt manager.
      func NewPromptManager() *PromptManager {
      	return &PromptManager{
      		templates: make(map[string]*template.Template),
      	}
      }

      // LoadPromptTemplate loads a named prompt template string.
      func (pm *PromptManager) LoadPromptTemplate(name, templateStr string) error {
      	tmpl, err := template.New(name).Parse(templateStr)
      	if err != nil {
      		return fmt.Errorf("failed to parse prompt template %s: %w", name, err)
      	}
      	pm.templates[name] = tmpl
      	return nil
      }

      // GetPrompt renders a named prompt template with the given data.
      // Data typically includes CodeSnippet, Language, and any custom promptArgs.
      func (pm *PromptManager) GetPrompt(name string, data interface{}) (string, error) {
      	tmpl, ok := pm.templates[name]
      	if !ok {
      		return "", fmt.Errorf("prompt template %s not found", name)
      	}

      	var buf bytes.Buffer
      	if err := tmpl.Execute(&buf, data); err != nil {
      		return "", fmt.Errorf("failed to render prompt template %s: %w", name, err)
      	}
      	return buf.String(), nil
      }

      // Example: Load templates at initialization
      // func initPrompts(pm *PromptManager) {
      // 	pm.LoadPromptTemplate("GenerateMarkdownDoc_Java", 
      // 		`Analyze the following Java code and generate comprehensive Markdown documentation for it. 
      // 		Focus on public APIs, class structure, and method signatures. 
      // 		Code: {{.CodeSnippet}} 
      // 		{{if .CustomInstructions}}Custom Instructions: {{.CustomInstructions}}{{end}}`)
      // 	pm.LoadPromptTemplate("GenerateOpenAPI_Go", 
      // 		`Extract API endpoint definitions from the following Go code and generate an OpenAPI v3 specification in JSON format. 
      // 		Code: {{.CodeSnippet}}`)
      // }

      // Usage within an AIClient method (conceptual):
      // func (c *ConcreteAIClient) GenerateDocumentation(ctx context.Context, codeSnippet string, language string, promptArgs map[string]string) (string, error) {
      // 	 templateName := fmt.Sprintf("GenerateMarkdownDoc_%s", language) // e.g., GenerateMarkdownDoc_Java
      // 	 data := map[string]interface{}{
      // 		 "CodeSnippet": codeSnippet,
      // 		 "Language": language,
      // 		 "CustomInstructions": promptArgs["customInstructions"], // Example from promptArgs
      // 	 }
      // 	 prompt, err := c.promptManager.GetPrompt(templateName, data)
      // 	 if err != nil { return "", err }
      // 	 // ... use the generated prompt in the AI call ...
      // }
      ```

### 4. 输出格式与定制 (Parent Task 4.0)

- **目标:** 支持生成多种标准化的文档和代码规范格式，并允许用户进行一定程度的定制，以满足不同团队和项目的需求。
- **子任务:**
  - [ ] 4.1 支持生成Markdown格式的API文档
    - **描述:** 将AI分析提取的类、方法、参数、返回值、注释等信息，结构化地组织成易于阅读和导航的Markdown文档。应包含代码块、表格、链接等元素。
    - **依赖:** 1.7, 3.1
    - **代码示例 (Go):**
      ```go
      // src/output/markdown_generator.go
      package output

      import (
      	"fmt"
      	"strings"
      	"//Users/liujinliang/workspace/ai/ai_rules/src/core" // For ClassInfo, MethodInfo etc. from Task 1.4, 1.5
      )

      // MarkdownGenerator generates Markdown documentation from parsed code structures.
      type MarkdownGenerator struct {
      	// Configuration options for markdown generation can be added here
      	// e.g., IncludePrivateMembers bool, CodeBlockTheme string
      }

      // NewMarkdownGenerator creates a new Markdown generator.
      func NewMarkdownGenerator() *MarkdownGenerator {
      	return &MarkdownGenerator{}
      }

      // GenerateClassDoc generates Markdown for a single class.
      func (mdg *MarkdownGenerator) GenerateClassDoc(classInfo core.ClassInfo) string {
      	var sb strings.Builder

      	sb.WriteString(fmt.Sprintf("# Class: %s\n\n", classInfo.Name))
      	if classInfo.Description != "" {
      		sb.WriteString(fmt.Sprintf("%s\n\n", classInfo.Description))
      	}
      	if len(classInfo.Fields) > 0 {
      		sb.WriteString("## Fields\n\n")
      		sb.WriteString("| Name | Type | Description |\n")
      		sb.WriteString("|------|------|-------------|\n")
      		for _, field := range classInfo.Fields {
      			sb.WriteString(fmt.Sprintf("| %s | `%s` | %s |\n", field.Name, field.Type, field.Description))
      		}
      		sb.WriteString("\n")
      	}

      	if len(classInfo.Methods) > 0 {
      		sb.WriteString("## Methods\n\n")
      		for _, method := range classInfo.Methods {
      			sb.WriteString(mdg.generateMethodDoc(method))
      		}
      	}
      	return sb.String()
      }

      // generateMethodDoc generates Markdown for a single method.
      func (mdg *MarkdownGenerator) generateMethodDoc(methodInfo core.MethodInfo) string {
      	var sb strings.Builder
      	sb.WriteString(fmt.Sprintf("### Method: %s\n\n", methodInfo.Name))
      	if methodInfo.Description != "" {
      		sb.WriteString(fmt.Sprintf("%s\n\n", methodInfo.Description))
      	}
      	sb.WriteString(fmt.Sprintf("- **Signature:** `%s`\n", methodInfo.Signature))
      	if len(methodInfo.Parameters) > 0 {
      		sb.WriteString("- **Parameters:**\n")
      		for _, param := range methodInfo.Parameters {
      			sb.WriteString(fmt.Sprintf("  - `%s` (%s): %s\n", param.Name, param.Type, param.Description))
      		}
      	}
      	if methodInfo.ReturnType != "" {
      		sb.WriteString(fmt.Sprintf("- **Returns:** `%s` - %s\n", methodInfo.ReturnType, methodInfo.ReturnDescription))
      	}
      	if methodInfo.CodeSnippet != "" {
      	    sb.WriteString(fmt.Sprintf("\n```%s\n%s\n```\n", strings.ToLower(methodInfo.Language), methodInfo.CodeSnippet))
      	}
      	sb.WriteString("\n")
      	return sb.String()
      }

      // GeneratePackageDoc generates Markdown for a collection of class/interface infos (e.g., a package).
      func (mdg *MarkdownGenerator) GeneratePackageDoc(packageName string, items []interface{}) string {
          var sb strings.Builder
          sb.WriteString(fmt.Sprintf("# Package: %s\n\n", packageName))
          for _, item := range items {
              switch v := item.(type) {
              case core.ClassInfo:
                  sb.WriteString(mdg.GenerateClassDoc(v))
                  sb.WriteString("---\n\n") // Separator
              // Add cases for InterfaceInfo, EnumInfo etc. as they are defined
              }
          }
          return sb.String()
      }
      ```
  - [ ] 4.2 支持生成OpenAPI (Swagger) 规范 (JSON/YAML)
    - **描述:** 对于包含Web API接口（如Spring MVC, JAX-RS, Go Gin/Echo等）的JAR包，利用AI分析代码，自动生成符合OpenAPI v3标准的API规范文档，支持JSON和YAML两种格式。这将极大方便API的消费和集成。
    - **依赖:** 1.8, 3.1
    - **代码示例 (Go):**
      ```go
      // src/output/openapi_generator.go
      package output

      import (
      	"encoding/json"
      	"//Users/liujinliang/workspace/ai/ai_rules/src/core" // For APIDefinition etc. from Task 1.8
      	"gopkg.in/yaml.v3" // go get gopkg.in/yaml.v3
      )

      // OpenAPIObject represents the root of an OpenAPI v3 specification.
      // This is a simplified version. A full implementation would use a library
      // like go-openapi/spec or kin-openapi.
      type OpenAPIObject struct {
      	OpenAPI string               `json:"openapi" yaml:"openapi"`
      	Info    OpenAPIInfo          `json:"info" yaml:"info"`
      	Paths   map[string]PathItem `json:"paths" yaml:"paths"`
      	// Components Components         `json:"components,omitempty" yaml:"components,omitempty"`
      }

      type OpenAPIInfo struct {
      	Title   string `json:"title" yaml:"title"`
      	Version string `json:"version" yaml:"version"`
      }

      type PathItem struct {
      	Get        *Operation `json:"get,omitempty" yaml:"get,omitempty"`
      	Post       *Operation `json:"post,omitempty" yaml:"post,omitempty"`
      	Put        *Operation `json:"put,omitempty" yaml:"put,omitempty"`
      	Delete     *Operation `json:"delete,omitempty" yaml:"delete,omitempty"`
      	// Other HTTP methods: Patch, Options, Head, Trace
      	Parameters []Parameter `json:"parameters,omitempty" yaml:"parameters,omitempty"` // Common parameters for all operations in this path
      }

      type Operation struct {
      	Summary     string                 `json:"summary,omitempty" yaml:"summary,omitempty"`
      	OperationID string                 `json:"operationId,omitempty" yaml:"operationId,omitempty"`
      	Parameters  []Parameter            `json:"parameters,omitempty" yaml:"parameters,omitempty"`
      	RequestBody *RequestBody           `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
      	Responses   map[string]APIResponse `json:"responses" yaml:"responses"`
      	// Tags, Security, Servers, etc.
      }

      type Parameter struct {
      	Name        string      `json:"name" yaml:"name"`
      	In          string      `json:"in" yaml:"in"` // e.g., "query", "path", "header", "cookie"
      	Description string      `json:"description,omitempty" yaml:"description,omitempty"`
      	Required    bool        `json:"required,omitempty" yaml:"required,omitempty"`
      	Schema      *Schema     `json:"schema,omitempty" yaml:"schema,omitempty"`
      }

      type RequestBody struct {
      	Description string                 `json:"description,omitempty" yaml:"description,omitempty"`
      	Content     map[string]MediaType `json:"content" yaml:"content"` // e.g., "application/json"
      	Required    bool                   `json:"required,omitempty" yaml:"required,omitempty"`
      }

      type APIResponse struct {
      	Description string                 `json:"description" yaml:"description"`
      	Content     map[string]MediaType `json:"content,omitempty" yaml:"content,omitempty"`
      	// Headers, Links
      }

      type MediaType struct {
      	Schema *Schema `json:"schema,omitempty" yaml:"schema,omitempty"`
      }

      type Schema struct {
      	Type        string             `json:"type,omitempty" yaml:"type,omitempty"` // e.g., "string", "integer", "object", "array"
      	Format      string             `json:"format,omitempty" yaml:"format,omitempty"` // e.g., "int32", "date-time"
      	Properties  map[string]*Schema `json:"properties,omitempty" yaml:"properties,omitempty"` // For object type
      	Items       *Schema            `json:"items,omitempty" yaml:"items,omitempty"`          // For array type
      	Ref         string             `json:"$ref,omitempty" yaml:"$ref,omitempty"`           // Reference to a schema in #/components/schemas
      	// Add more fields: enum, default, example, required (list of strings for object properties)
      }

      // GenerateOpenAPISpec converts extracted APIDefinitions to an OpenAPIObject.
      // This is a highly simplified conversion.
      func GenerateOpenAPISpec(apiDefs []core.APIDefinition, title, version string) *OpenAPIObject {
      	spec := &OpenAPIObject{
      		OpenAPI: "3.0.0",
      		Info:    OpenAPIInfo{Title: title, Version: version},
      		Paths:   make(map[string]PathItem),
      	}

      	for _, apiDef := range apiDefs {
      		pathItem := spec.Paths[apiDef.Path]
      		operation := &Operation{
      			Summary:     apiDef.Description,
      			OperationID: apiDef.OperationID,
      			Responses:   make(map[string]APIResponse),
      		}

      		// Simplified: Assume all params are query params for GET, or part of request body for POST/PUT
      		// A real implementation would parse @Param, @RequestBody annotations etc.
      		for _, param := range apiDef.Parameters {
      			operation.Parameters = append(operation.Parameters, Parameter{
      				Name: param.Name,
      				In:   "query", // Placeholder
      				Schema: &Schema{Type: param.Type}, // Placeholder, map to OpenAPI types
      			})
      		}

      		// Simplified response
      		operation.Responses["200"] = APIResponse{Description: "Success"}
      		if apiDef.ReturnType != "void" && apiDef.ReturnType != "" {
      		    operation.Responses["200"] = APIResponse{
      		        Description: "Success",
      		        Content: map[string]MediaType{
      		            "application/json": {Schema: &Schema{Type: apiDef.ReturnType /* map to OpenAPI type */ }},
      		        },
      		    }
      		}


      		switch strings.ToUpper(apiDef.HTTPMethod) {
      		case "GET":
      			pathItem.Get = operation
      		case "POST":
      			pathItem.Post = operation
      			// Potentially create RequestBody based on parameters if not query/path/header
      		case "PUT":
      			pathItem.Put = operation
      		case "DELETE":
      			pathItem.Delete = operation
      		}
      		spec.Paths[apiDef.Path] = pathItem
      	}
      	return spec
      }

      // ToJSON converts the OpenAPIObject to a JSON string.
      func (spec *OpenAPIObject) ToJSON() (string, error) {
      	jsonData, err := json.MarshalIndent(spec, "", "  ")
      	if err != nil {
      		return "", err
      	}
      	return string(jsonData), nil
      }

      // ToYAML converts the OpenAPIObject to a YAML string.
      func (spec *OpenAPIObject) ToYAML() (string, error) {
      	yamlData, err := yaml.Marshal(spec)
      	if err != nil {
      		return "", err
      	}
      	return string(yamlData), nil
      }
      ```
  - [ ] 4.3 支持生成代码注释规范 (如JavaDoc, GoDoc)
    - **描述:** 基于AI对代码的理解，为缺少注释或注释不规范的类、方法、字段等自动生成符合特定语言注释规范（如JavaDoc, GoDoc, JSDoc, Python Docstrings）的注释。用户可以选择注释风格和详细程度。
    - **依赖:** 1.4, 1.5, 3.1
    - **代码示例 (Go):**
      ```go
      // src/output/comment_generator.go
      package output

      import (
      	"fmt"
      	"strings"
      	"//Users/liujinliang/workspace/ai/ai_rules/src/core" // For ClassInfo, MethodInfo etc.
      )

      // CommentStyle defines the style of comments to generate.
      type CommentStyle string

      const (
      	StyleJavaDoc CommentStyle = "JavaDoc"
      	StyleGoDoc   CommentStyle = "GoDoc"
      	// Add other styles like JSDoc, PythonDocstring
      )

      // CommentGenerator generates code comments.
      type CommentGenerator struct {
      	Style CommentStyle
      }

      // NewCommentGenerator creates a new comment generator.
      func NewCommentGenerator(style CommentStyle) *CommentGenerator {
      	return &CommentGenerator{Style: style}
      }

      // GenerateMethodComment generates a comment for a method.
      // Assumes methodInfo.AIExplanation contains AI-generated explanation for the method.
      func (cg *CommentGenerator) GenerateMethodComment(methodInfo core.MethodInfo, aiExplanation string) string {
      	var sb strings.Builder
      	description := aiExplanation // Or a more structured summary from AI
      	if description == "" {
      		description = fmt.Sprintf("Method %s does something.", methodInfo.Name) // Fallback
      	}

      	switch cg.Style {
      	case StyleJavaDoc:
      		sb.WriteString("/**\n")
      		sb.WriteString(fmt.Sprintf(" * %s\n", wrapText(description, 70, " * ")))
      		if len(methodInfo.Parameters) > 0 {
      			sb.WriteString(" *\n")
      			for _, param := range methodInfo.Parameters {
      				paramDesc := fmt.Sprintf("Parameter %s of type %s.", param.Name, param.Type) // AI should provide better desc
      				sb.WriteString(fmt.Sprintf(" * @param %s %s\n", param.Name, wrapText(paramDesc, 60, " *         ")))
      			}
      		}
      		if methodInfo.ReturnType != "" && methodInfo.ReturnType != "void" {
      			returnDesc := fmt.Sprintf("Returns %s.", methodInfo.ReturnType) // AI should provide better desc
      			sb.WriteString(" *\n")
      			sb.WriteString(fmt.Sprintf(" * @return %s\n", wrapText(returnDesc, 60, " *          ")))
      		}
      		sb.WriteString(" */")
      	case StyleGoDoc:
      		// GoDoc is typically simpler, placed above the function.
      		// It often starts with the function name.
      		sb.WriteString(fmt.Sprintf("// %s %s\n", methodInfo.Name, wrapText(description, 70, "// ")))
      		// GoDoc doesn't have formal @param or @return tags in the same way as JavaDoc.
      		// Parameter and return value explanations are usually part of the main comment block.
      		// For brevity, this example doesn't elaborate further but a real impl would integrate param/return info.
      	default:
      		return fmt.Sprintf("// %s: %s", methodInfo.Name, description) // Default simple comment
      	}
      	return sb.String()
      }

      // wrapText is a helper to wrap long lines for comments.
      func wrapText(text string, lineWidth int, prefix string) string {
      	words := strings.Fields(text)
      	if len(words) == 0 {
      		return ""
      	}
      	var lines []string
      	currentLine := prefix
      	for _, word := range words {
      		if len(currentLine)+len(word)+1 > lineWidth && len(currentLine) > len(prefix) {
      			lines = append(lines, strings.TrimSpace(currentLine))
      			currentLine = prefix
      		}
      		currentLine += word + " "
      	}
      	lines = append(lines, strings.TrimSpace(currentLine))
      	return strings.Join(lines, "\n")
      }
      ```
  - [ ] 4.4 支持自定义输出模板 (如使用Handlebars, Go Template)
    - **描述:** 允许高级用户提供自定义的模板文件（如Handlebars或Go Template格式），用于生成Markdown文档、OpenAPI规范或其他文本输出。系统会将提取的代码结构信息和AI分析结果作为上下文数据传递给模板引擎进行渲染。
    - **依赖:** 1.4, 1.5, 1.7, 1.8, 3.8 (for prompt templating concepts, can be extended to output templating)
    - **代码示例 (Go - using Go's text/template):**
      ```go
      // src/output/custom_template_renderer.go
      package output

      import (
      	"bytes"
      	"fmt"
      	"text/template"
      	"//Users/liujinliang/workspace/ai/ai_rules/src/core" // For ClassInfo, APIDefinition etc.
      )

      // TemplateRenderer renders data using custom Go templates.
      type TemplateRenderer struct {
      	// Potentially cache compiled templates
      }

      // NewTemplateRenderer creates a new template renderer.
      func NewTemplateRenderer() *TemplateRenderer {
      	return &TemplateRenderer{}
      }

      // RenderDataWithTemplate renders the given data object using the provided template string.
      // data can be ClassInfo, APIDefinition, a list of these, or any custom struct.
      func (tr *TemplateRenderer) RenderDataWithTemplate(templateName string, templateStr string, data interface{}) (string, error) {
      	tmpl, err := template.New(templateName).Parse(templateStr)
      	if err != nil {
      		return "", fmt.Errorf("failed to parse custom template '%s': %w", templateName, err)
      	}

      	var buf bytes.Buffer
      	if err := tmpl.Execute(&buf, data); err != nil {
      		return "", fmt.Errorf("failed to render custom template '%s': %w", templateName, err)
      	}
      	return buf.String(), nil
      }

      /* Example Usage:
      func main() {
      	classData := core.ClassInfo{
      		Name:        "MyService",
      		Description: "A sample service for demonstration.",
      		Methods: []core.MethodInfo{
      			{Name: "doWork", Signature: "void doWork(String task)", Description: "Performs a task."},
      		},
      	}

      	customMarkdownTemplate := `
      # {{.Name}}

      {{.Description}}

      ## Methods:
      {{range .Methods}}
      ### {{.Name}}
      Signature: `+"`{{.Signature}}`"+`
      {{.Description}}
      {{end}}
      `
      	renderer := NewTemplateRenderer()
      	output, err := renderer.RenderDataWithTemplate("customClassMarkdown", customMarkdownTemplate, classData)
      	if err != nil {
      		log.Fatal(err)
      	}
      	fmt.Println(output)
      }
      */
      ```
  - [ ] 4.5 支持输出内容的多语言翻译 (集成翻译API)
    - **描述:** 对于生成的文档、注释或代码解释，允许用户选择目标语言，系统通过集成第三方翻译服务API（如Google Translate, DeepL）将内容翻译成指定语言。需要注意API Key管理和用量控制。
    - **依赖:** 1.7, 4.1, 4.3
    - **代码示例 (Go - conceptual, actual API call omitted):**
      ```go
      // src/output/translator.go
      package output

      import (
      	"context"
      	"fmt"
      	// "cloud.google.com/go/translate" // Example: Google Translate client
      	// "golang.org/x/text/language"      // For language tags
      )

      // Translator provides text translation services.
      type Translator struct {
      	// client *translate.Client // Example client
      	apiKey string // API key for the translation service
      }

      // NewTranslator creates a new translator instance.
      // apiKey should be securely managed, e.g., via config (Task 3.5).
      func NewTranslator(apiKey string) (*Translator, error) {
      	// ctx := context.Background()
      	// client, err := translate.NewClient(ctx, option.WithAPIKey(apiKey))
      	// if err != nil {
      	// 	return nil, fmt.Errorf("failed to create translate client: %w", err)
      	// }
      	return &Translator{apiKey: apiKey /*, client: client*/}, nil
      }

      // TranslateText translates text from sourceLang to targetLang.
      // lang codes are typically ISO 639-1 (e.g., "en", "es", "zh-CN").
      func (t *Translator) TranslateText(ctx context.Context, text, sourceLang, targetLang string) (string, error) {
      	if text == "" || sourceLang == targetLang {
      		return text, nil // No translation needed
      	}

      	// Placeholder for actual translation API call
      	// target, err := language.Parse(targetLang)
      	// if err != nil {
      	// 	return "", fmt.Errorf("invalid target language: %w", err)
      	// }
      	// source, err := language.Parse(sourceLang) // Optional, some APIs auto-detect
      	// if err != nil { ... }

      	// resp, err := t.client.Translate(ctx, []string{text}, target, &translate.Options{Source: source})
      	// if err != nil {
      	// 	return "", fmt.Errorf("translation failed: %w", err)
      	// }
      	// if len(resp) == 0 {
      	// 	return "", fmt.Errorf("translation returned empty response")
      	// }
      	// return resp[0].Text, nil

      	return fmt.Sprintf("[Translated from %s to %s: %s]", sourceLang, targetLang, text), nil // Mock implementation
      }

      // Close releases resources used by the translator (e.g., client connections).
      func (t *Translator) Close() {
      	// if t.client != nil {
      	// 	t.client.Close()
      	// }
      }
      ```

### 5. 测试、部署与监控 (Parent Task 5.0)

- **目标:** 确保系统的质量、可靠性和可维护性，提供便捷的部署方案，并建立有效的监控机制。
- **子任务:**
  - [ ] 5.1 编写单元测试、集成测试和端到端测试
    - **描述:** 为核心模块（如JAR解析、代码分析、AI交互、API接口）编写单元测试。设计集成测试验证模块间的协作。开发端到端测试模拟用户真实操作流程，确保整个系统的功能正确性和稳定性。
    - **依赖:** All major functional tasks (1.x, 2.x, 3.x, 4.x)
    - **代码示例 (Go - testing package):**
      ```go
      // src/parser/jar_parser_test.go
      package parser

      import (
      	"os"
      	"path/filepath"
      	"testing"
      	"//Users/liujinliang/workspace/ai/ai_rules/src/core" // For CodeFile
      )

      func TestJarParser_Parse(t *testing.T) {
      	// Setup: Create a dummy JAR file for testing or use a pre-existing small test JAR
      	tempDir := t.TempDir()
      	dummyJarPath := filepath.Join(tempDir, "test.jar")
      	// For a real test, you'd create a valid JAR here. For simplicity, we'll mock.
      	// For example, create a zip file with a .class file inside.
      	// Or, use a known small JAR from testdata.
      	// For this example, let's assume Parse can handle a non-JAR for error checking.
      	file, _ := os.Create(dummyJarPath)
      	file.WriteString("not a jar")
      	file.Close()

      	parser := NewJarParser("cfr") // Assuming NewJarParser exists

      	tests := []struct {
      		name    string
      		jarPath string
      		wantErr bool
      		// Add checks for expected CodeFile outputs if parsing a real JAR
      	}{
      		{"valid dummy jar (mocked success)", "testdata/dummy.jar", false}, // Assuming dummy.jar exists and is valid
      		{"non-existent jar", "testdata/nonexistent.jar", true},
      		{"not a jar file", dummyJarPath, true},
      	}

      	for _, tt := range tests {
      		t.Run(tt.name, func(t *testing.T) {
      			// Adjust path for testdata if needed
      			jarToTest := tt.jarPath
      			if strings.HasPrefix(tt.jarPath, "testdata/") && tt.jarPath != dummyJarPath {
      				// This assumes testdata is relative to the test file's directory or project root.
      				// For robustness, use absolute paths or ensure testdata is correctly located.
      				// _, currentFile, _, _ := runtime.Caller(0)
      				// baseDir := filepath.Dir(currentFile)
      				// jarToTest = filepath.Join(baseDir, "testdata", filepath.Base(tt.jarPath))
      				// For simplicity, we'll assume testdata/dummy.jar is accessible.
      				// If not, this test case will fail on file open.
      			}

      			gotFiles, err := parser.Parse(jarToTest)
      			if (err != nil) != tt.wantErr {
      				t.Errorf("JarParser.Parse() error = %v, wantErr %v", err, tt.wantErr)
      				return
      			}
      			if !tt.wantErr && len(gotFiles) == 0 && tt.jarPath == "testdata/dummy.jar" {
      				// If we expect files from dummy.jar, this is a failure
      				// t.Errorf("JarParser.Parse() got no files from dummy.jar, want >0")
      			}
      			// Add more specific assertions about the content of gotFiles if applicable
      		})
      	}
      }

      // Example for an integration test (conceptual)
      // src/api/server_integration_test.go
      // package api_test
      // import (
      //  "net/http"
      //  "net/http/httptest"
      //  "testing"
      //  "//Users/liujinliang/workspace/ai/ai_rules/src/api" // Assuming SetupRouter is in api package
      // )
      // func TestUploadAndProcessIntegration(t *testing.T) {
      //  router := api.SetupRouter() // Setup your Gin router with all handlers
      //  ts := httptest.NewServer(router)
      //  defer ts.Close()
      //  // 1. Simulate JAR upload to ts.URL + "/upload"
      //  // 2. Poll ts.URL + "/tasks/{taskId}/status"
      //  // 3. Fetch results from ts.URL + "/tasks/{taskId}/results/markdown"
      //  // Assert expected outcomes at each step.
      // }
      ```
  - [ ] 5.2 设计并实现CI/CD流水线 (如GitHub Actions, Jenkins)
    - **描述:** 建立持续集成和持续部署（CI/CD）流水线。CI阶段应包括代码编译、静态代码分析、运行所有测试、构建Docker镜像。CD阶段可以将通过测试的Docker镜像自动部署到测试环境或生产环境。
    - **依赖:** 5.1, 5.3 (Dockerization)
    - **代码示例 (GitHub Actions Workflow - .github/workflows/ci-cd.yml):**
      ```yaml
      name: Go CI/CD Pipeline

      on:
        push:
          branches: [ "main", "develop" ]
        pull_request:
          branches: [ "main" ]

      jobs:
        build-and-test:
          runs-on: ubuntu-latest
          steps:
          - uses: actions/checkout@v3

          - name: Set up Go
            uses: actions/setup-go@v4
            with:
              go-version: '1.21' # Specify your Go version

          - name: Install dependencies
            run: go mod download

          - name: Vet
            run: go vet ./...

          # - name: Lint # Add a linter like golangci-lint if desired
          #   uses: golangci/golangci-lint-action@v3
          #   with:
          #     version: v1.55.2

          - name: Test with coverage
            run: go test -v -coverprofile=coverage.out ./...

          - name: Upload coverage reports to Codecov
            uses: codecov/codecov-action@v3
            with:
              token: ${{ secrets.CODECOV_TOKEN }} # Optional: if you use Codecov
              files: ./coverage.out

          # Build step (optional here, could be part of a separate deploy job)
          - name: Build
            run: go build -v -o myapp ./cmd/server # Assuming your main is in cmd/server/main.go

          # Docker build and push (example, more complex for real CD)
          # This part would typically be in a separate job triggered on merge to main/develop
          # and would require Docker Hub credentials or other registry credentials
          # - name: Log in to Docker Hub
          #   if: github.event_name != 'pull_request' && (github.ref == 'refs/heads/main' || github.ref == 'refs/heads/develop')
          #   uses: docker/login-action@v2
          #   with:
          #     username: ${{ secrets.DOCKER_USERNAME }}
          #     password: ${{ secrets.DOCKER_PASSWORD }}

          # - name: Build and push Docker image
          #   if: github.event_name != 'pull_request' && (github.ref == 'refs/heads/main' || github.ref == 'refs/heads/develop')
          #   uses: docker/build-push-action@v4
          #   with:
          #     context: .
          #     file: ./Dockerfile # Assuming you have a Dockerfile
          #     push: true
          #     tags: yourdockerhubuser/myapp:${{ github.sha }}, yourdockerhubuser/myapp:latest
      ```
  - [ ] 5.3 提供Docker镜像和Kubernetes部署配置
    - **描述:** 将整个应用打包成Docker镜像，方便快速部署和环境一致性。提供基础的Kubernetes部署配置文件（如Deployment, Service, ConfigMap），帮助用户在K8s集群中部署和管理应用。
    - **依赖:** 3.5 (for config management in K8s)
    - **代码示例 (Dockerfile & Kubernetes YAML - conceptual):**
      **Dockerfile:**
      ```dockerfile
      # Dockerfile
      FROM golang:1.21-alpine AS builder

      WORKDIR /app

      # Copy Go modules and source code
      COPY go.mod go.sum ./
      RUN go mod download

      COPY . .

      # Build the application
      # CGO_ENABLED=0 for static linking, GOOS=linux for cross-compilation if needed
      RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/ai-jar-sdk-generator ./cmd/server/main.go

      # --- Final Stage ---
      FROM alpine:latest

      WORKDIR /root/

      # Copy the Pre-built binary file from the previous stage
      COPY --from=builder /app/ai-jar-sdk-generator .

      # Copy config files if needed (or use K8s ConfigMap)
      # COPY --from=builder /app/config.yaml .

      # Expose port (if it's a server application)
      EXPOSE 8080

      # Command to run the executable
      CMD ["./ai-jar-sdk-generator"]
      ```
      **Kubernetes Deployment (deployment.yaml - simplified):**
      ```yaml
      apiVersion: apps/v1
      kind: Deployment
      metadata:
        name: ai-jar-sdk-generator
      spec:
        replicas: 1
        selector:
          matchLabels:
            app: ai-jar-sdk-generator
        template:
          metadata:
            labels:
              app: ai-jar-sdk-generator
          spec:
            containers:
            - name: generator-app
              image: yourusername/ai-jar-sdk-generator:latest # Replace with your image
              ports:
              - containerPort: 8080
              # Volume mounts for config if using ConfigMap
              # envFrom:
              # - configMapRef:
              #     name: generator-config
      ```
      **Kubernetes Service (service.yaml - simplified):**
      ```yaml
      apiVersion: v1
      kind: Service
      metadata:
        name: ai-jar-sdk-generator-svc
      spec:
        selector:
          app: ai-jar-sdk-generator
        ports:
        - protocol: TCP
          port: 80 # Port to access the service on
          targetPort: 8080 # Port the container is listening on
        type: LoadBalancer # Or ClusterIP/NodePort depending on needs
      ```
  - [ ] 5.4 实现系统监控和告警机制 (如Prometheus, Grafana)
    - **描述:** 集成Prometheus收集应用的关键指标（如API请求延迟、错误率、任务处理时间、资源使用情况）。使用Grafana创建仪表盘可视化这些指标。配置告警规则，当指标异常时（如错误率过高、任务队列积压）通过邮件或Slack等渠道发送通知。
    - **依赖:** 3.1 (API endpoints to expose metrics), 3.7 (logging for error tracking)
    - **代码示例 (Go - Prometheus client conceptual):**
      ```go
      // src/monitoring/metrics.go
      package monitoring

      import (
      	"github.com/prometheus/client_golang/prometheus"
      	"github.com/prometheus/client_golang/prometheus/promauto"
      	"github.com/prometheus/client_golang/prometheus/promhttp"
      	"net/http"
      	"time"
      )

      var (
      	HttpRequestsTotal = promauto.NewCounterVec(
      		prometheus.CounterOpts{
      			Name: "http_requests_total",
      			Help: "Total number of HTTP requests.",
      		},
      		[]string{"path", "method", "status"},
      	)
      	HttpRequestDuration = promauto.NewHistogramVec(
      		prometheus.HistogramOpts{
      			Name:    "http_request_duration_seconds",
      			Help:    "Duration of HTTP requests.",
      			Buckets: prometheus.DefBuckets, // Default buckets: .005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10
      		},
      		[]string{"path", "method"},
      	)
      	AiTaskProcessingTime = promauto.NewHistogram(
      		prometheus.HistogramOpts{
      			Name:    "ai_task_processing_duration_seconds",
      			Help:    "Duration of AI task processing.",
      			Buckets: []float64{1, 5, 10, 30, 60, 120, 300, 600}, // Custom buckets for longer tasks
      		},
      	)
      )

      // PrometheusMiddleware is a Gin middleware to record HTTP metrics.
      func PrometheusMiddleware() gin.HandlerFunc { // Assuming Gin framework
      	return func(c *gin.Context) {
      		start := time.Now()
      		c.Next() // Process request

      		duration := time.Since(start)
      		status := fmt.Sprintf("%d", c.Writer.Status())
      		path := c.FullPath() // Or c.Request.URL.Path for raw path
      		method := c.Request.Method

      		HttpRequestsTotal.WithLabelValues(path, method, status).Inc()
      		HttpRequestDuration.WithLabelValues(path, method).Observe(duration.Seconds())
      	}
      }

      // ExposeMetricsEndpoint sets up an HTTP endpoint for Prometheus to scrape.
      func ExposeMetricsEndpoint(router *gin.Engine, path string) {
      	router.GET(path, gin.WrapH(promhttp.Handler()))
      }

      /* Example usage in main.go:
      func main() {
          router := gin.Default()
          router.Use(monitoring.PrometheusMiddleware())
          monitoring.ExposeMetricsEndpoint(router, "/metrics")

          // ... your other routes ...

          router.Run(":8080")
      }
      */
      ```
  - [ ] 5.5 制定数据备份和恢复策略
    - **描述:** 对于系统产生的持久化数据（如用户上传的JARs、生成的文档、任务状态、配置文件），制定定期备份策略。备份数据应存储在安全可靠的位置（如云存储）。同时，需要有明确的数据恢复流程，以便在发生故障时能够快速恢复服务和用户数据。
    - **依赖:** 1.1 (uploaded JARs), 1.7, 4.1, 4.2 (generated outputs), 3.4 (task state), 3.5 (configs)
    - **策略文档 (Conceptual - not code, but a description of the strategy):**
      ```markdown
      ## Data Backup and Recovery Strategy

      **1. Data to Backup:**
        - User uploaded JAR files (if stored persistently beyond task processing).
        - Generated documentation (Markdown, OpenAPI specs) if stored long-term.
        - Task metadata and status (e.g., from a database like PostgreSQL or Redis snapshots if used for task queue state).
        - System configuration files (if not fully managed by K8s ConfigMaps or similar).
        - AI interaction logs (if deemed critical for audit or retraining, consider size implications).

      **2. Backup Frequency:**
        - **User Data (JARs, Generated Docs):** Daily incremental backups, weekly full backups.
        - **Task Metadata/Status:** Near real-time replication if using a DB with such capabilities (e.g., PostgreSQL streaming replication). Otherwise, hourly snapshots for task queues/DBs.
        - **Configuration Files:** Backup on every change, and daily snapshots.
        - **Logs:** If stored, depends on retention policy. Often shipped to a log aggregation system which has its own backup.

      **3. Backup Storage:**
        - Primary backups: Secure cloud storage (e.g., AWS S3, Google Cloud Storage) with versioning enabled.
        - Secondary backups (optional): Off-site storage in a different geographical region for disaster recovery.
        - Encryption: All backups must be encrypted at rest and in transit.

      **4. Retention Policy:**
        - Daily backups: Retain for 7 days.
        - Weekly full backups: Retain for 4 weeks.
        - Monthly full backups: Retain for 6 months.
        - Yearly full backups (optional): Retain for 1-3 years for compliance if needed.

      **5. Recovery Procedures:**
        - **Recovery Time Objective (RTO):** Define target time to restore service (e.g., 4 hours).
        - **Recovery Point Objective (RPO):** Define maximum acceptable data loss (e.g., 1 hour).
        - Documented step-by-step recovery process for each data type.
        - Regular testing of recovery procedures (e.g., quarterly drills) to ensure they are effective and up-to-date.
        - Automated recovery scripts where possible.

      **6. Monitoring Backups:**
        - Alerts for backup failures.
        - Regular verification of backup integrity.
      ```

### 6. 文档完善与规范制定 (Parent Task 6.0)

- **目标:** 提供全面、清晰的系统文档，制定代码和API规范，方便用户使用和开发者维护。
- **子任务:**
  - [ ] 6.1 编写用户手册和开发者指南
    - **描述:** **用户手册**应详细说明如何安装、配置和使用本工具的各项功能，包括命令行参数、API接口用法、输入输出格式、常见问题解答等。**开发者指南**应包括项目架构、模块设计、代码结构、构建和测试流程、贡献指南、编码规范等，方便新开发者快速上手和参与贡献。
    - **依赖:** All functional tasks, 5.3 (for installation/deployment info)
    - **产出物 (Conceptual - Markdown files):**
      - `docs/user_manual.md`
      - `docs/developer_guide.md`
      - `CONTRIBUTING.md`
      - `CODE_OF_CONDUCT.md`
  - [ ] 6.2 制定API接口规范和代码风格规范
    - **描述:** **API接口规范**应遵循RESTful原则或OpenAPI标准，明确请求/响应格式、错误码、认证方式等。**代码风格规范**应选择一种主流风格（如Effective Go, Google Java Style Guide），并使用自动化工具（如`gofmt`/`goimports` for Go, Checkstyle/Spotless for Java）强制执行，确保代码一致性和可读性。
    - **依赖:** 3.1 (API design), 4.2 (OpenAPI generation can help enforce/document API spec)
    - **产出物 (Conceptual - Markdown files / Config files):**
      - `docs/api_specification.md` (or link to generated OpenAPI spec)
      - `docs/coding_style_guide.md`
      - `.golangci.yml` (for Go linters)
      - `checkstyle.xml` (for Java Checkstyle)
  - [ ] 6.3 设计项目Logo和品牌标识 (可选)
    - **描述:** 为项目设计一个简洁、有意义的Logo，并制定基本的品牌视觉指南（如标准色、字体），用于文档、网站或演示材料中，提升项目专业形象。
    - **依赖:** N/A (Creative task)
    - **产出物 (Conceptual - Image files, Style guide):**
      - `assets/logo.svg` (Vector format preferred)
      - `assets/logo_icon.svg`
      - `docs/brand_guidelines.md`
  - [ ] 6.4 建立社区支持渠道 (如GitHub Issues, Discord)
    - **描述:** 设置并维护公开的社区支持渠道，如GitHub Issues用于Bug报告和功能请求，Discord服务器或Slack频道用于用户交流和快速问答。明确社区行为准则。
    - **依赖:** N/A (Community management task)
    - **产出物 (Conceptual - Links, Policy docs):**
      - Link to GitHub Issues page.
      - Link to Discord/Slack server invite.
      - `COMMUNITY_GUIDELINES.md` (or part of `CODE_OF_CONDUCT.md`)
  - [ ] 6.5 撰写项目介绍和推广材料 (如README, 博客文章)
    - **描述:** 撰写清晰、吸引人的项目`README.md`，包括项目简介、核心功能、快速开始、示例、许可证、贡献方式等。可以撰写博客文章或教程来介绍项目的特性和使用场景，吸引潜在用户和贡献者。
    - **依赖:** All functional tasks, 6.1 (user manual content can be summarized)
    - **产出物 (Conceptual - Markdown files, Blog posts):**
      - `README.md` (at project root)
      - Blog post drafts or links to published articles.

**项目时间计划 (初步估计):**

- **阶段一 (核心功能实现):** 4-6周 (完成Parent Task 1.0, 3.1, 3.2, 3.7)
- **阶段二 (高级功能与集成):** 6-8周 (完成Parent Task 2.0, 3.3-3.6, 3.8)
- **阶段三 (输出与部署):** 4-6周 (完成Parent Task 4.0, 5.0)
- **阶段四 (文档与社区):** 2-3周 (完成Parent Task 6.0)

**总计:** 约16-23周

**注:** 以上任务分解、依赖关系、代码示例和时间计划均为初步设想，具体实现中可能需要根据实际情况进行调整。代码示例主要用于概念演示，并非完整可运行代码。路径如 `//Users/liujinliang/workspace/ai/ai_rules/src/core` 仅为示例，实际应为项目内的相对或绝对路径。安全性和错误处理在示例中被简化。API密钥管理、并发处理、资源限制等重要方面需要在实际开发中详细考虑。

      - **依赖:** 2.7 (if JARs are fetched based on repo info), 2.8 (if POM is in repo)
      - **代码示例 (Go):**
        ```go
        // src/integrations/vcs_handler.go
        package integrations

        import (
        	"context"
        	"fmt"
        	"os/exec"
        	"strings"

        	"github.com/google/go-github/v57/github" // Example: using go-github library
        )

        // CloneRepository clones a Git repository to a specified directory.
        func CloneRepository(repoURL, targetDir, branchOrTag string) error {
        	cmdArgs := []string{"clone"}
        	if branchOrTag != "" {
        		cmdArgs = append(cmdArgs, "-b", branchOrTag)
        	}
        	cmdArgs = append(cmdArgs, repoURL, targetDir)
        	cmd := exec.Command("git", cmdArgs...)
        	output, err := cmd.CombinedOutput()
        	if err != nil {
        		return fmt.Errorf("git clone failed: %s\nOutput: %s", err, string(output))
        	}
        	return nil
        }

        // GetGitHubRepoInfo fetches repository metadata from GitHub API.
        // Requires a GitHub Personal Access Token for non-public or rate-limited access.
        func GetGitHubRepoInfo(owner, repoName string) (*github.Repository, error) {
        	client := github.NewClient(nil) // Or provide an http.Client with auth
        	repo, _, err := client.Repositories.Get(context.Background(), owner, repoName)
        	if err != nil {
        		return nil, err
        	}
        	return repo, nil
        }

        // Example usage:
        // func ProcessGitHubRepo(repoURL, branch, localPath string) {
        // 	 owner, name := parseGitHubURL(repoURL) // Helper to extract owner/name
        // 	 repoInfo, err := GetGitHubRepoInfo(owner, name)
        // 	 // ... use repoInfo ...
        // 	 err = CloneRepository(repoURL, localPath, branch)
        // 	 // ... find JARs or build project in localPath ...
        // }
        - [ ] 2.10 实现反编译结果与官方文档的对比验证功能
      - **描述:** 如果用户提供了JAR包对应的官方文档链接或内容，系统可以尝试将AI生成的文档（基于反编译结果）与官方文档进行对比。通过NLP技术（如文本相似度、语义匹配）找出差异点、缺失信息或潜在的过时内容，并向用户高亮提示。这有助于提高生成文档的准确性和可信度。
      - **依赖:** 1.6 (AI generated docs), 1.7 (Markdown/OpenAPI output)
      - **代码示例 (Go):**
        ```go
        // src/validation/doc_comparator.go
        package validation

        import (
        	"//Users/liujinliang/workspace/ai/ai_rules/src/ai" // For AIModelOutput
        	"//Users/liujinliang/workspace/ai/ai_rules/src/utils" // For text processing utilities
        )

        // OfficialDocSource represents the source of official documentation.
        type OfficialDocSource struct {
        	URL     string // URL to the official documentation page/API spec
        	Content string // Raw content if provided directly
        }

        // ComparisonResult highlights differences between generated and official docs.
        type ComparisonResult struct {
        	Section         string  // e.g., "Method: com.example.MyClass.doSomething"
        	SimilarityScore float64 // 0.0 to 1.0
        	Discrepancies   []string // List of identified differences
        	IsMissingInGen  bool    // True if section exists in official but not generated
        	IsMissingInOff  bool    // True if section exists in generated but not official
        }

        // CompareGeneratedWithOfficial compares AI-generated docs with official documentation.
        func CompareGeneratedWithOfficial(generatedDoc ai.AIModelOutput, officialDoc OfficialDocSource) ([]ComparisonResult, error) {
        	var results []ComparisonResult
        	// Placeholder: Complex NLP logic required here.
        	// 1. Fetch/parse officialDoc.Content or URL.
        	// 2. Structure both generatedDoc and officialDoc into comparable segments (e.g., per class, per method).
        	// 3. For each segment, calculate similarity (e.g., using TF-IDF, embeddings, cosine similarity).
        	// 4. Identify discrepancies (e.g., parameter mismatches, different return types, missing descriptions).

        	// Example conceptual comparison for a method description:
        	// genMethodDesc := generatedDoc.Methods["com.example.MyClass.doSomething"].Description
        	// offMethodDesc := parseOfficialDocForMethod("com.example.MyClass.doSomething", officialDoc)
        	// score := utils.CalculateTextSimilarity(genMethodDesc, offMethodDesc)
        	// if score < 0.8 { // Threshold for significant difference
        	// 	 results = append(results, ComparisonResult{...})
        	// }
        	return results, nil
        }
        ```
    - [ ] 2.11 提取官方示例代码和最佳实践
      - **描述:** 在处理官方文档或源码（如果可访问，如通过GitHub集成）时，系统应尝试识别和提取代码示例、用法片段和推荐的最佳实践。这些提取的内容可以作为补充信息整合到生成的SDK文档中，为开发者提供更实用的指导。
      - **依赖:** 2.9 (for source code access), 2.10 (if official docs are processed)
      - **代码示例 (Go):**
        ```go
        // src/extraction/example_extractor.go
        package extraction

        import (
        	"regexp"
        	"//Users/liujinliang/workspace/ai/ai_rules/src/validation" // For OfficialDocSource
        )

        // CodeExample represents an extracted code snippet.
        type CodeExample struct {
        	AssociatedAPI string // e.g., "com.example.MyClass.doSomething"
        	Language      string // e.g., "java", "kotlin"
        	Snippet       string
        	Source        string // Where it was extracted from (e.g., URL, file path)
        	Description   string // Optional description of the example
        }

        // ExtractCodeExamples attempts to find code examples from various sources.
        func ExtractCodeExamples(officialDoc validation.OfficialDocSource, sourceCodeDir string) ([]CodeExample, error) {
        	var examples []CodeExample
        	// Placeholder: Complex parsing and pattern matching logic.

        	// 1. From official documentation content (HTML, Markdown, etc.)
        	//    - Look for <pre><code>, ```java, etc. blocks.
        	if officialDoc.Content != "" {
        		// Regex to find markdown Java code blocks
        		// This is a simplified example and would need to be more robust
        		javaMarkdownRegex := regexp.MustCompile("```java\n([\s\S]*?)\n```")
        		matches := javaMarkdownRegex.FindAllStringSubmatch(officialDoc.Content, -1)
        		for _, match := range matches {
        			if len(match) > 1 {
        				examples = append(examples, CodeExample{
        					Language: "java",
        					Snippet:  match[1],
        					Source:   "Official Documentation",
        				})
        			}
        		}
        	}

        	// 2. From source code files (e.g., test directories, example modules in a repo)
        	if sourceCodeDir != "" {
        		// Walk the directory, find files like *Test.java, *Example.java
        		// Parse these files to extract relevant methods or snippets.
        		// This is highly language and project-structure dependent.
        	}

        	// 3. Identify best practices from text (e.g., looking for keywords like "recommended", "best practice")

        	return examples, nil
        }
        ```
- [ ] 3.0 系统架构与集成：构建稳健系统并实现协议兼容
  - [ ] 3.1 设计并实现一套完整的API接口，支持JAR包上传、任务管理、结果查询与下载
    - **描述:** 定义并实现一套RESTful或gRPC API接口，作为系统的主要交互方式。这些接口应覆盖核心功能，包括：JAR文件上传、启动分析任务、查询任务状态、获取处理进度、检索生成的文档（Markdown, OpenAPI等）、下载文档文件。
    - **依赖:** 1.1, 1.9, 1.10, 2.3
    - **代码示例 (Go - using Gin framework for REST API):**
      ```go
      // src/api/server.go
      package api

      import (
      	"net/http"
      	"//Users/liujinliang/workspace/ai/ai_rules/src/core"
      	"github.com/gin-gonic/gin"
      )

      // SetupRouter configures the API routes.
      func SetupRouter() *gin.Engine {
      	r := gin.Default()

      	// JAR Upload
      	r.POST("/upload", func(c *gin.Context) {
      		file, err := c.FormFile("jarfile")
      		if err != nil {
      			c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
      			return
      		}
      		// tempPath := "/tmp/" + file.Filename // Save to a temporary location
      		// if err := c.SaveUploadedFile(file, tempPath); err != nil { ... }
      		// taskID, err := core.SubmitJarAnalysisTask(tempPath, core.DefaultAIConfig())
      		// if err != nil { ... }
      		c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "task_id": "dummy-task-id"})
      	})

      	// Task Management
      	r.GET("/task/:task_id/status", func(c *gin.Context) {
      		taskID := c.Param("task_id")
      		// status, progress, err := core.GetTaskStatus(taskID)
      		c.JSON(http.StatusOK, gin.H{"task_id": taskID, "status": "processing", "progress": 50})
      	})

      	// Result Query & Download
      	r.GET("/task/:task_id/result/:format", func(c *gin.Context) {
      		taskID := c.Param("task_id")
      		format := c.Param("format") // e.g., "markdown", "openapi"
      		// docContent, err := core.GetTaskResult(taskID, format)
      		// if format == "download" { c.Data(...) } else { c.String(...) }
      		c.JSON(http.StatusOK, gin.H{"task_id": taskID, "format": format, "content": "dummy content"})
      	})
      	return r
      }
      ```
  - [ ] 3.2 实现异步任务处理机制，支持长耗时任务的后台执行与状态反馈
    - **描述:** JAR包分析（特别是反编译和AI处理）可能是耗时操作。系统需要一个异步任务处理机制（如使用消息队列 RabbitMQ/Kafka,或者Go的goroutines和channels配合持久化存储）。API调用应立即返回任务ID，客户端可以通过任务ID轮询或WebSocket接收任务状态更新。
    - **依赖:** 3.1
    - **代码示例 (Go - using goroutines and channels for simplicity):**
      ```go
      // src/core/task_processor.go
      package core

      import (
      	"fmt"
      	"sync"
      	"time"
      	"//Users/liujinliang/workspace/ai/ai_rules/src/ai"
      )

      type TaskStatus struct {
      	ID         string
      	Status     string // PENDING, PROCESSING, COMPLETED, FAILED
      	Progress   int    // 0-100
      	ResultPath string // Path to result if COMPLETED
      	Error      string // Error message if FAILED
      }

      var (
      	taskStore      = make(map[string]*TaskStatus)
      	taskStoreMutex = sync.RWMutex{}
      	// In a real system, use a proper job queue (e.g., Redis list, RabbitMQ)
      	jobQueue = make(chan string, 100) // Channel of JAR paths to process
      )

      // SubmitJarAnalysisTask adds a JAR to the processing queue and returns a task ID.
      func SubmitJarAnalysisTask(jarPath string, aiConfig ai.AIModelConfig) (string, error) {
      	taskStoreMutex.Lock()
      	defer taskStoreMutex.Unlock()

      	taskID := fmt.Sprintf("task-%d", time.Now().UnixNano())
      	taskStore[taskID] = &TaskStatus{ID: taskID, Status: "PENDING", Progress: 0}
      	
      	// Simulate adding to a persistent queue before starting goroutine
      	go processJarInBackground(taskID, jarPath, aiConfig) // Fire and forget
      	// jobQueue <- jarPath // This would be for a worker pool model

      	return taskID, nil
      }

      func processJarInBackground(taskID string, jarPath string, aiConfig ai.AIModelConfig) {
      	updateTaskStatus(taskID, "PROCESSING", 10, "", "")
      	// Simulate work
      	// 1. Decompile (updateTaskStatus with progress)
      	// 2. Analyze (updateTaskStatus with progress)
      	// 3. Generate docs with AI (updateTaskStatus with progress)
      	time.Sleep(10 * time.Second) // Simulate long processing

      	// resultPath, err := ActualJarProcessingFunction(jarPath, aiConfig)
      	// if err != nil {
      	// 	 updateTaskStatus(taskID, "FAILED", 100, "", err.Error())
      	// 	 return
      	// }
      	updateTaskStatus(taskID, "COMPLETED", 100, "/path/to/generated/docs/"+taskID, "")
      }

      func updateTaskStatus(taskID, status string, progress int, resultPath, errMsg string) {
      	taskStoreMutex.Lock()
      	defer taskStoreMutex.Unlock()
      	if ts, ok := taskStore[taskID]; ok {
      		ts.Status = status
      		ts.Progress = progress
      		ts.ResultPath = resultPath
      		ts.Error = errMsg
      	}
      }

      // GetTaskStatus retrieves the current status of a task.
      func GetTaskStatus(taskID string) (*TaskStatus, error) {
      	taskStoreMutex.RLock()
      	defer taskStoreMutex.RUnlock()
      	ts, ok := taskStore[taskID]
      	if !ok {
      		return nil, fmt.Errorf("task %s not found", taskID)
      	}
      	return ts, nil
      }
      ```