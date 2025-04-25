# Golang 代码规范

## 1. 项目结构
- 遵循Go Modules标准，项目根目录包含go.mod。
- 推荐项目结构：
  ```
  project/
  ├── cmd/                  # 应用程序入口点
  │   └── server/           # 服务器入口
  │       └── main.go       # 主函数
  ├── pkg/                  # 可重用代码包
  │   ├── models/           # 数据模型
  │   ├── database/         # 数据库配置
  │   ├── handlers/         # HTTP处理器
  │   ├── middleware/       # 中间件
  │   └── utils/            # 工具函数
  ├── internal/             # 内部代码
  │   ├── config/           # 配置管理
  │   └── service/          # 业务逻辑
  ├── api/                  # API定义
  ├── web/                  # 前端资源
  ├── scripts/              # 构建脚本
  ├── docs/                 # 文档
  ├── go.mod
  ├── go.sum
  ├── .gitignore
  └── README.md
  ```

## 2. 命名规范
- 包名全部小写，简短有意义，不使用下划线。
- 文件名小写，单词间下划线分隔。
- 类型名、导出函数/变量用大驼峰（PascalCase），非导出用小驼峰（camelCase）。
- 常量用CamelCase。
- 错误变量以Err或Error开头。
- 接口名以-er结尾，如Reader、Writer。
- 避免使用缩写，除非是众所周知的（如 HTTP, URL, ID）。

## 3. 代码风格
- 遵循Go官方风格，使用gofmt自动格式化。
- 每行不超过120字符。
- 注释规范，导出类型/函数需有注释，推荐使用GoDoc风格。
- 导入顺序：标准库、第三方、本地包。
- 推荐使用Go 1.20+版本。
### 3.1 完整文件标注
在文件顶部添加注释块：
```go
/**
 * [AI-ASSISTED]
 * 生成工具: {工具名称} {版本}
 * 生成日期: {YYYY-MM-DD}
 * 贡献程度: {完全生成|部分生成|辅助编写|重构优化}
 * 人工修改: {无|轻微|中度|大量}
 * 责任人: {开发者姓名}
 */
```

### 3.2 代码块标注
对于 AI 生成或修改的代码块，使用注释标记起止：
```go
/* [AI-BLOCK-START] - 生成工具: {工具名称} {版本} */
generatedCode := "AI generated code"
/* [AI-BLOCK-END] */
```

### 3.3 单行标注
对于单行 AI 生成或修改的代码，在行尾添加注释：
```go
result := complexCalculation() // [AI-LINE: {工具名称} {版本}]
```
## 4. 依赖管理
- 使用go.mod/go.sum统一管理依赖。
- 固定依赖版本，避免随意replace。
- 推荐使用私有代理或镜像源。

## 5. 文档规范
- 每个包、导出类型、函数需有GoDoc注释。
- 项目需包含README.md、CHANGELOG.md、/docs目录。
- 推荐使用godoc、Swagger/OpenAPI生成文档。

## 6. 编码规范
- 推荐使用go vet、golint、staticcheck等工具静态检查。
- 单元测试放在*_test.go文件，测试函数以Test开头。
- 避免魔法数字，使用常量。
- 错误处理要规范，避免忽略 `error`。优先使用 `errors.Is` 和 `errors.As` 进行错误判断和解包。在适当情况下定义和使用自定义错误类型。
- 推荐使用context传递上下文。
- 推荐使用标准库 `log` 或结构化日志库（如 `zap`, `logrus`）。日志应包含足够上下文信息（如请求ID、时间戳）。
- 推荐使用泛型和新特性提升代码复用。

## 7. 最佳实践

### 代码组织
- 每个包职责单一，避免大文件和大函数。
- 公共逻辑抽取为独立包，便于复用和维护。
- 目录结构清晰，便于扩展和查找。

### 性能优化
- 优先使用切片和map等高效数据结构。
- 合理使用goroutine和channel，避免资源泄漏。
- 尽量减少全局变量，优先使用局部变量。
- 使用sync包实现并发安全。

### 安全性
- 对所有外部输入（如API请求参数、配置文件、数据库读取内容）进行严格的有效性检查。
- 错误处理要规范，优先返回 error 而不是 panic，避免 panic 导致程序崩溃或信息泄漏。
- 注意并发场景下的数据竞争，使用 `sync` 包或 channel 保证并发安全。
- 避免硬编码敏感信息（如密码、API密钥），使用配置管理或环境变量。
- 考虑使用中间件处理认证、授权、日志记录等横切关注点。
- 根据需要实现速率限制，防止滥用。

### 测试
- 单元测试覆盖核心逻辑，推荐使用Go自带testing框架。
- 自动化集成测试，保证主要功能不被破坏。
- 关键路径增加断言和日志，便于定位问题。

### 设计模式
- 适当使用工厂、单例、装饰器等设计模式提升可维护性。
- 复杂模块建议先画流程图或伪代码。

## 8. API 设计最佳实践
- 遵循 RESTful 设计原则。
- 使用 Go 1.22+ 的 `net/http.ServeMux` 进行路由，利用其模式匹配能力。
  - 支持路径变量 `/users/{id}` 和通配符 `/files/{path...}`
  - 支持方法匹配 `GET /users` 和 `POST /users`
  - 支持优先级路由 `/users/latest` 优先于 `/users/{id}`
- 中间件链实现：
  ```go
  mux := http.NewServeMux()
  handler := loggingMiddleware(authMiddleware(mux))
  ```
- 正确处理不同的 HTTP 方法 (GET, POST, PUT, DELETE 等)。
- 接口路径使用小写字母和连字符 `-`。
- 使用 DTO (Data Transfer Object) 封装请求和响应体。
- API 版本控制（如 `/v1/users`）。
- 提供清晰的错误响应格式：
  ```go
  {
    "error": {
      "code": "invalid_request",
      "message": "Invalid user ID format"
    }
  }
  ```

## 11. 工具链推荐
- 构建：go build/go install
- 依赖管理：go mod
- 静态检查：go vet、golint、staticcheck
- 格式化：gofmt、goimports
- 文档：godoc、Swagger/OpenAPI
- 单元测试：testing、testif
## 9. AI 代码标注规范



## 10. Git 规范
- 分支命名: `feature/ai-{feature_name}` 或遵循项目自身规范并体现 AI 辅助。
- Commit Message: `[AI-ASSISTED] {commit_message}` 或类似标识。
- Pull Request 标题: `[AI] {pr_title}` 或类似标识。