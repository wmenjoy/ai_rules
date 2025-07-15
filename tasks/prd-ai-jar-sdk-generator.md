# 产品需求文档 (PRD) - AI+反编译JAR包SDK文档生成器

## 1. 介绍/概述

本产品旨在解决AI系统无法识别和调用Java JAR包API的核心问题。通过AI+反编译技术，动态分析JAR包并生成AI可识别、开发者可使用的标准化SDK文档，帮助AI系统和开发者快速理解和使用任意JAR包，包括那些缺乏文档或代码混淆的JAR包。

### 解决的核心问题
- AI系统无法直接理解和调用JAR包中的API
- 许多JAR包缺乏完整的SDK文档
- 代码混淆的JAR包难以理解和使用
- 开发者需要手动分析JAR包结构和API

## 2. 目标

### 主要目标
1. **AI可用性**：生成AI系统可以理解和调用的标准化API文档
2. **开发效率**：帮助开发者快速理解和使用任意JAR包
3. **文档完整性**：为缺乏文档的JAR包生成完整的API说明
4. **混淆处理**：合理处理代码混淆的JAR包，提取可用信息
5. **动态服务**：提供实时的JAR包分析和文档生成服务

### 可衡量指标
- 支持JDK 6以上所有版本的JAR包
- 文档生成准确率 ≥ 90%
- 混淆JAR包信息提取率 ≥ 70%
- 单个JAR包处理时间 < 5分钟
- 支持多种AI模型（Qwen3、DeepSeek、OpenAI兼容）

## 3. 用户故事

### AI系统用户故事
- **作为AI系统**，我希望能够理解JAR包中的API结构，以便正确调用相关方法
- **作为AI助手**，我希望获得标准化的API文档，以便为用户提供准确的代码示例
- **作为MCP工具**，我希望能够动态获取JAR包的API信息，以便扩展功能能力

### 开发者用户故事
- **作为Java开发者**，我希望快速了解第三方JAR包的API结构，以便快速集成到项目中
- **作为API开发者**，我希望为我的JAR包生成标准文档，以便其他开发者使用
- **作为逆向工程师**，我希望分析混淆的JAR包，以便理解其功能和接口

## 4. 功能要求

### 4.1 核心功能
1. **JAR包上传与解析**
   - 系统必须支持JAR包文件上传（本地文件、URL链接）
   - 系统必须支持JDK 6及以上版本的JAR包
   - 系统必须能够处理代码混淆的JAR包

2. **反编译与分析**
   - 系统必须使用多种反编译引擎（CFR、JD-Core等）进行交叉验证
   - 系统必须提取类结构、方法签名、参数信息、返回值类型
   - 系统必须分析依赖关系和调用链
   - 系统必须保留原有注释信息（如果存在）

3. **AI文档生成**
   - 系统必须支持多种AI模型（Qwen3、DeepSeek、OpenAI兼容API）
   - 系统必须基于代码语义生成方法说明和参数描述
   - 系统必须生成使用示例和最佳实践
   - 系统必须识别和说明异常处理机制

4. **文档输出**
   - 系统必须生成Markdown格式的开发者文档
   - 系统必须生成OpenAPI 3.0规范的机器可读文档
   - 系统必须生成JSON Schema格式的AI可读文档
   - 系统必须包含完整的API信息：方法签名、参数说明、使用示例、异常处理
   - 系统必须提供文档下载功能

5. **批量处理与扫描**
   - 系统必须支持JAR包目录批量扫描
   - 系统必须支持递归扫描子目录中的JAR文件
   - 系统必须提供扫描进度和结果统计
   - 系统必须支持扫描结果的批量导出

6. **外部仓库集成**
   - 系统必须支持从Maven Repository抓取JAR包信息
   - 系统必须支持从GitHub仓库获取项目信息和文档
   - 系统必须解析Maven POM文件获取依赖信息
   - 系统必须提取GitHub README、Wiki等文档信息

7. **MCP集成**
   - 系统必须提供MCP（Model Context Protocol）接口
   - 系统必须支持动态JAR包分析请求
   - 系统必须返回结构化的API信息供AI系统调用

### 4.2 高级功能
8. **批量处理与目录扫描**
   - 系统必须支持递归扫描目录结构，识别所有JAR文件
   - 系统必须支持多线程并发分析，提高处理效率
   - 系统必须提供实时扫描进度和处理状态显示
   - 系统必须支持统一管理批量处理结果，支持分组导出
   - 系统必须支持文件大小、修改时间、文件名模式过滤
   - 系统必须支持中断后继续处理未完成的文件

9. **外部仓库集成技术**
   - 系统必须集成Maven Central、阿里云等仓库API
   - 系统必须解析POM文件获取依赖关系和元数据
   - 系统必须自动下载指定版本的JAR包进行分析
   - 系统必须使用GitHub API获取仓库信息和文档
   - 系统必须将反编译结果与官方文档进行对比验证
   - 系统必须提取官方示例代码和最佳实践

10. **实时处理服务**
   - 系统必须提供RESTful API接口
   - 系统必须支持异步处理和进度查询
   - 系统必须提供处理状态和错误信息反馈

11. **质量保证**
   - 系统必须提供文档质量评分（准确性、完整性、可读性）
   - 系统必须支持人工审核和修正机制
   - 系统必须记录处理日志和错误信息

12. **缓存与优化**
    - 系统必须缓存已处理的JAR包结果
    - 系统必须支持增量更新和版本对比
    - 系统必须优化大型JAR包的处理性能

## 5. 非目标（超出范围）

本功能将**不**包括以下内容：
- IDE插件开发（如IntelliJ IDEA、Eclipse插件）
- CI/CD流水线集成
- 除Java外的其他编程语言支持（如.NET、Python等）
- 源码级别的代码分析（仅处理编译后的JAR包）
- 实时代码执行和测试功能
- 用户权限管理和多租户支持
- 商业化的许可证管理

## 6. 设计考虑

### 6.1 SDK文档格式规范

#### 6.1.1 Markdown格式规范
```markdown
# {JAR包名称} SDK文档

## 概述
- **包名**: {package.name}
- **版本**: {version}
- **描述**: {description}
- **依赖**: {dependencies}

## 快速开始
### 安装
```xml
<dependency>
    <groupId>{groupId}</groupId>
    <artifactId>{artifactId}</artifactId>
    <version>{version}</version>
</dependency>
```

### 基本使用
```java
// 示例代码
```

## API参考
### 类: {ClassName}
**包路径**: `{package.path}`
**描述**: {class.description}

#### 构造方法
- `{constructor.signature}`
  - **参数**: {parameters}
  - **描述**: {description}
  - **示例**: {example}

#### 方法
- `{method.signature}`
  - **参数**: {parameters}
  - **返回值**: {returnType}
  - **异常**: {exceptions}
  - **描述**: {description}
  - **示例**: {example}

## 异常处理
### {ExceptionName}
- **类型**: {exception.type}
- **原因**: {exception.cause}
- **处理**: {exception.handling}

## 最佳实践
- {best.practice.1}
- {best.practice.2}

## 更新日志
- {changelog}
```

#### 6.1.2 OpenAPI 3.0格式规范
```yaml
openapi: 3.0.3
info:
  title: {JAR包名称} API
  version: {version}
  description: {description}
  contact:
    name: {maintainer}
    url: {repository.url}

servers:
  - url: {base.url}
    description: {server.description}

paths:
  /{endpoint}:
    post:
      summary: {method.summary}
      description: {method.description}
      parameters:
        - name: {param.name}
          in: {param.location}
          required: {param.required}
          schema:
            type: {param.type}
            description: {param.description}
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/{RequestModel}'
      responses:
        '200':
          description: {success.description}
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/{ResponseModel}'
        '400':
          description: {error.description}
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/ErrorResponse'

components:
  schemas:
    {ModelName}:
      type: object
      required:
        - {required.field}
      properties:
        {field.name}:
          type: {field.type}
          description: {field.description}
          example: {field.example}
```

#### 6.1.3 JSON Schema格式规范（AI可读）
```json
{
  "jarInfo": {
    "name": "{jar.name}",
    "version": "{version}",
    "groupId": "{groupId}",
    "artifactId": "{artifactId}",
    "description": "{description}",
    "dependencies": ["{dependency.list}"]
  },
  "classes": [
    {
      "className": "{full.class.name}",
      "packageName": "{package.name}",
      "classType": "class|interface|enum|annotation",
      "modifiers": ["public", "final", "abstract"],
      "description": "{ai.generated.description}",
      "superClass": "{super.class.name}",
      "interfaces": ["{interface.names}"],
      "constructors": [
        {
          "signature": "{constructor.signature}",
          "parameters": [
            {
              "name": "{param.name}",
              "type": "{param.type}",
              "description": "{param.description}"
            }
          ],
          "description": "{constructor.description}",
          "example": "{usage.example}"
        }
      ],
      "methods": [
        {
          "methodName": "{method.name}",
          "signature": "{full.method.signature}",
          "modifiers": ["public", "static", "final"],
          "returnType": "{return.type}",
          "parameters": [
            {
              "name": "{param.name}",
              "type": "{param.type}",
              "description": "{param.description}",
              "required": true
            }
          ],
          "exceptions": [
            {
              "type": "{exception.type}",
              "description": "{exception.description}"
            }
          ],
          "description": "{ai.generated.description}",
          "example": "{usage.example}",
          "complexity": "low|medium|high",
          "aiCallable": true
        }
      ],
      "fields": [
        {
          "fieldName": "{field.name}",
          "type": "{field.type}",
          "modifiers": ["public", "static", "final"],
          "description": "{field.description}",
          "defaultValue": "{default.value}"
        }
      ]
    }
  ],
  "aiMetadata": {
    "generatedAt": "{timestamp}",
    "aiModel": "{model.name}",
    "qualityScore": 0.95,
    "completeness": 0.98,
    "confidence": 0.92,
    "suggestedUsage": ["{usage.pattern.1}", "{usage.pattern.2}"],
    "relatedLibraries": ["{related.lib.1}", "{related.lib.2}"]
  }
}
```

### 6.2 系统架构
- **微服务架构**：反编译服务、分析服务、AI生成服务、文档服务独立部署
- **异步处理**：支持大型JAR包的后台处理
- **负载均衡**：支持多实例部署和水平扩展

### 6.3 用户界面
- **Web界面**：基于React的现代化SPA应用，支持拖拽上传、实时进度显示、文档预览
- **API接口**：Go实现的高性能RESTful API
- **MCP接口**：标准化的MCP协议实现

### 6.4 安全考虑
- **沙箱执行**：JAR包分析在隔离环境中进行
- **恶意代码检测**：上传文件的安全扫描
- **访问控制**：API调用频率限制

## 7. 技术考虑

### 7.1 核心技术栈
- **反编译引擎**：CFR、JD-Core、Procyon（多引擎交叉验证）
- **字节码分析**：ASM、Javassist
- **AI模型集成**：支持Qwen3、DeepSeek、OpenAI兼容API
- **文档生成**：Markdown、OpenAPI 3.0
- **后端框架**：Go Gin/Fiber（高性能HTTP框架）
- **前端框架**：React 18+（TypeScript）
- **数据存储**：Redis（缓存）、PostgreSQL（元数据）

### 7.2 性能要求
- **处理速度**：单个JAR包（<50MB）处理时间 < 5分钟
- **并发支持**：同时处理10个JAR包
- **内存使用**：单个处理任务内存占用 < 2GB
- **存储空间**：生成文档大小通常为原JAR包的2-5倍

### 7.3 兼容性
- **Java版本**：JDK 6 - JDK 21
- **JAR包类型**：标准JAR、Fat JAR、混淆JAR
- **AI模型**：Qwen3、DeepSeek、OpenAI GPT系列

## 8. 成功指标

### 8.1 技术指标
- **准确性**：生成的API文档与实际JAR包功能匹配度 ≥ 90%
- **完整性**：提取的API信息覆盖率 ≥ 95%（非混淆JAR包）
- **可读性**：生成文档的人工评分 ≥ 4.0/5.0
- **处理成功率**：JAR包处理成功率 ≥ 95%

### 8.2 业务指标
- **用户采用率**：月活跃用户数增长
- **API调用量**：MCP接口调用次数
- **文档下载量**：生成文档的下载和使用情况
- **错误率**：系统错误和用户反馈的问题数量

### 8.3 用户体验指标
- **处理时间**：用户等待时间满意度
- **文档质量**：用户对生成文档的评价
- **易用性**：界面和API的使用便利性

## 9. 技术问题解决方案

### 9.1 混淆算法识别与处理
**问题**：如何识别和处理不同类型的代码混淆？

**解决方案**：
1. **混淆类型检测**：
   - 实现启发式算法检测常见混淆模式（ProGuard、R8、Allatori等）
   - 分析类名、方法名的命名模式（如单字符、随机字符串）
   - 检测控制流混淆、字符串加密等特征

2. **多层次信息提取**：
   - **字节码层面**：使用ASM分析方法签名、参数类型、返回值
   - **反编译层面**：CFR、JD-Core交叉验证，提取可读代码结构
   - **元数据层面**：提取MANIFEST.MF、注解信息、资源文件

3. **语义恢复策略**：
   - 基于方法调用关系推断功能模块
   - 利用字符串常量和异常信息推测方法用途
   - 通过参数类型和返回值推断方法语义

### 9.2 AI模型选择与优化
**问题**：在不同场景下如何选择最适合的AI模型？

**解决方案**：
1. **模型分工策略**：
   - **Qwen3**：中文注释生成、本土化API说明
   - **DeepSeek**：代码理解、技术文档生成
   - **OpenAI GPT**：英文文档、标准化描述

2. **智能路由机制**：
   ```go
   type ModelSelector struct {
       codeComplexity float64
       languageHint   string
       docType        string
   }
   
   func (ms *ModelSelector) SelectModel() AIModel {
       if ms.languageHint == "zh" {
           return QwenModel
       }
       if ms.codeComplexity > 0.8 {
           return DeepSeekModel
       }
       return OpenAIModel
   }
   ```

3. **质量评估与回退**：
   - 实现多模型并行生成，选择最佳结果
   - 建立质量评分机制（完整性、准确性、可读性）
   - 支持模型级联：主模型失败时自动切换备用模型

### 9.3 性能优化策略
**问题**：如何优化大型JAR包（>100MB）的处理性能？

**解决方案**：
1. **分块处理架构**：
   ```go
   type JarProcessor struct {
       chunkSize    int64
       workerPool   *WorkerPool
       resultCache  *Cache
   }
   
   func (jp *JarProcessor) ProcessLargeJar(jarPath string) {
       // 按类文件分块并行处理
       chunks := jp.splitJarByClasses(jarPath)
       results := jp.processChunksParallel(chunks)
       return jp.mergeResults(results)
   }
   ```

2. **内存管理优化**：
   - 流式处理：避免将整个JAR包加载到内存
   - 垃圾回收优化：及时释放临时对象
   - 内存池：复用反编译器实例

3. **缓存策略**：
   - **类级缓存**：相同类的处理结果缓存
   - **依赖缓存**：常用库的分析结果预缓存
   - **增量处理**：只处理变更的类文件

### 9.4 依赖关系解析
**问题**：如何处理JAR包的外部依赖关系？

**解决方案**：
1. **依赖发现机制**：
   ```go
   type DependencyAnalyzer struct {
       mavenResolver *MavenResolver
       classPath     []string
   }
   
   func (da *DependencyAnalyzer) ResolveDependencies(jarPath string) []Dependency {
       // 分析import语句和方法调用
       imports := da.extractImports(jarPath)
       // 尝试从Maven仓库解析
       return da.mavenResolver.Resolve(imports)
   }
   ```

2. **依赖处理策略**：
   - **内置依赖**：直接分析JAR包内的依赖类
   - **外部依赖**：从Maven Central、本地仓库获取
   - **缺失依赖**：基于方法签名推断，标记为"推测类型"

3. **依赖图构建**：
   - 构建完整的类依赖关系图
   - 识别核心API和辅助类
   - 生成依赖关系文档

### 9.5 MCP协议实现细节
**问题**：MCP协议的具体实现细节和扩展需求？

**解决方案**：
1. **标准MCP接口**：
   ```go
   type MCPServer struct {
       tools []MCPTool
   }
   
   type JarAnalysisTool struct {
       Name        string `json:"name"`
       Description string `json:"description"`
       InputSchema Schema `json:"inputSchema"`
   }
   
   func (jat *JarAnalysisTool) Execute(params map[string]interface{}) MCPResult {
       jarPath := params["jarPath"].(string)
       return jat.analyzeJar(jarPath)
   }
   ```

2. **扩展功能**：
   - **实时分析**：支持流式返回分析进度
   - **批量处理**：支持多JAR包并行分析
   - **结果缓存**：智能缓存机制减少重复分析

3. **错误处理**：
   - 标准化错误码和错误信息
   - 支持部分失败的优雅降级
   - 详细的调试信息输出

## 10. 前端设计规范

### 10.1 技术规范
- **框架版本**：React 18.2+ with TypeScript 5.0+
- **状态管理**：Zustand（轻量级状态管理）
- **UI组件库**：Ant Design 5.0+（企业级UI组件）
- **样式方案**：Tailwind CSS + CSS Modules
- **构建工具**：Vite 4.0+（快速构建）
- **代码规范**：ESLint + Prettier + Husky

### 10.2 页面结构设计

#### 主页面布局
```typescript
interface MainLayoutProps {
  header: React.ReactNode;    // 顶部导航栏
  sidebar: React.ReactNode;   // 左侧功能菜单
  content: React.ReactNode;   // 主内容区域
  footer: React.ReactNode;    // 底部信息栏
}
```

#### 核心页面组件
1. **上传页面** (`/upload`)
   ```typescript
   interface UploadPageProps {
     onFileUpload: (file: File) => void;
     uploadProgress: number;
     supportedFormats: string[];
   }
   ```

2. **分析进度页面** (`/analysis/:taskId`)
   ```typescript
   interface AnalysisPageProps {
     taskId: string;
     progress: AnalysisProgress;
     logs: LogEntry[];
   }
   ```

3. **文档预览页面** (`/docs/:docId`)
   ```typescript
   interface DocsPageProps {
     document: GeneratedDoc;
     format: 'markdown' | 'openapi';
     downloadOptions: DownloadOption[];
   }
   ```

### 10.3 UI设计规范

#### 色彩规范
```css
:root {
  /* 主色调 */
  --primary-color: #1890ff;      /* 蓝色 - 主要操作 */
  --primary-hover: #40a9ff;      /* 蓝色悬停态 */
  --primary-active: #096dd9;     /* 蓝色激活态 */
  
  /* 辅助色 */
  --success-color: #52c41a;      /* 绿色 - 成功状态 */
  --warning-color: #faad14;      /* 橙色 - 警告状态 */
  --error-color: #ff4d4f;        /* 红色 - 错误状态 */
  
  /* 中性色 */
  --text-primary: #262626;       /* 主要文本 */
  --text-secondary: #595959;     /* 次要文本 */
  --text-disabled: #bfbfbf;      /* 禁用文本 */
  --border-color: #d9d9d9;       /* 边框颜色 */
  --background-color: #fafafa;   /* 背景色 */
}
```

#### 组件样式规范
1. **按钮规范**
   ```typescript
   interface ButtonProps {
     type: 'primary' | 'secondary' | 'danger';
     size: 'small' | 'medium' | 'large';
     loading?: boolean;
     disabled?: boolean;
   }
   ```

2. **卡片规范**
   ```css
   .card {
     border-radius: 8px;
     box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
     padding: 24px;
     background: white;
   }
   ```

3. **表单规范**
   ```css
   .form-item {
     margin-bottom: 24px;
   }
   
   .form-label {
     font-weight: 500;
     margin-bottom: 8px;
   }
   
   .form-input {
     border-radius: 6px;
     border: 1px solid var(--border-color);
     padding: 8px 12px;
   }
   ```

### 10.4 交互设计规范

#### 文件上传交互
1. **拖拽上传区域**
   - 默认状态：虚线边框，提示文本
   - 拖拽悬停：实线边框，高亮背景
   - 上传中：进度条显示，禁用交互
   - 上传完成：成功图标，文件信息展示

2. **进度反馈**
   ```typescript
   interface ProgressState {
     stage: 'uploading' | 'decompiling' | 'analyzing' | 'generating' | 'completed';
     percentage: number;
     currentTask: string;
     estimatedTime: number;
   }
   ```

#### 文档展示交互
1. **文档导航**
   - 左侧目录树：支持折叠展开
   - 右侧内容区：支持锚点跳转
   - 顶部面包屑：显示当前位置

2. **代码高亮**
   ```typescript
   interface CodeBlockProps {
     language: string;
     code: string;
     showLineNumbers: boolean;
     highlightLines?: number[];
   }
   ```

### 10.5 响应式设计

#### 断点规范
```css
/* 移动端 */
@media (max-width: 768px) {
  .container { padding: 16px; }
  .sidebar { display: none; }
}

/* 平板端 */
@media (min-width: 769px) and (max-width: 1024px) {
  .container { padding: 24px; }
  .sidebar { width: 200px; }
}

/* 桌面端 */
@media (min-width: 1025px) {
  .container { padding: 32px; }
  .sidebar { width: 240px; }
}
```

#### 组件适配
- **表格组件**：移动端自动转换为卡片布局
- **导航菜单**：移动端收缩为汉堡菜单
- **文档预览**：移动端支持左右滑动切换章节

### 10.6 性能优化规范

1. **代码分割**
   ```typescript
   // 路由级别的懒加载
   const UploadPage = lazy(() => import('./pages/UploadPage'));
   const DocsPage = lazy(() => import('./pages/DocsPage'));
   ```

2. **组件优化**
   ```typescript
   // 使用React.memo优化重渲染
   const DocumentViewer = React.memo(({ document }: DocumentViewerProps) => {
     return <div>{/* 文档内容 */}</div>;
   });
   ```

3. **资源优化**
   - 图片懒加载：使用Intersection Observer
   - 虚拟滚动：大列表性能优化
   - 缓存策略：API响应缓存，本地存储优化

## 11. 实施计划

### 第一阶段：基础架构（3-4周）
- **Go后端框架搭建**：使用Gin框架，建立项目结构，配置路由和中间件
- **React前端初始化**：使用Vite+TypeScript，配置Ant Design和Tailwind CSS
- **基础反编译功能**：集成CFR反编译器，实现基本的JAR包解析
- **数据库设计**：设计任务、结果、用户等核心数据表
- **Docker容器化**：编写Dockerfile和docker-compose配置

### 第二阶段：核心功能开发（4-5周）
- **AI集成**：接入Qwen3、DeepSeek、OpenAI等模型API
- **文档生成引擎**：实现Markdown、OpenAPI和JSON Schema格式输出
- **Web界面开发**：完成上传、进度显示、文档预览等核心页面
- **API接口开发**：实现RESTful API和异步任务处理
- **基础测试**：单元测试和集成测试

### 第三阶段：批量处理与外部集成（4-5周）
- **目录扫描功能**：递归扫描、并行处理、进度跟踪
- **Maven仓库集成**：Maven Central API集成、POM解析、自动下载
- **GitHub集成**：GitHub API集成、文档提取、信息融合
- **批量管理界面**：批量处理UI、结果聚合、分组导出
- **过滤和断点续传**：文件过滤机制、中断恢复功能

### 第四阶段：高级功能与优化（3-4周）
- **MCP协议实现**：开发标准化的MCP接口
- **性能优化**：缓存机制、并发处理、内存优化
- **质量保证系统**：文档质量评分、人工审核功能
- **安全加固**：沙箱执行、恶意代码检测、访问控制
- **监控和日志**：系统监控、错误追踪、性能分析

### 第五阶段：测试与部署（2-3周）
- **全面测试**：功能测试、性能测试、安全测试、批量处理测试
- **文档编写**：用户手册、API文档、部署指南、格式规范文档
- **生产部署**：生产环境配置、CI/CD流水线
- **用户培训**：使用指南、最佳实践分享、批量处理指导

---

**文档版本**：v1.0  
**创建日期**：2024-12-19  
**负责人**：AI助手  
**审核状态**：待审核