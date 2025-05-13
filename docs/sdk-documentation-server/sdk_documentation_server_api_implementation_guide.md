# SDK文档服务器API规范实施指南

**文档编号**: API-IMPL-SDKDOCS-2024-001  
**版本**: 1.0.0  
**状态**: 初稿  
**创建日期**: 2024-05-15  
**最后更新**: 2024-05-15  
**编制人**: 技术文档团队  
**审核人**: 待定  

## 文档修订记录

| 版本   | 日期       | 修订人 | 描述       | 审核人 |
|-------|------------|-------|------------|-------|
| 1.0.0 | 2024-05-15 | 技术文档团队 | 初始版本  | 待定   |

## 目录

1. [简介](#1-简介)
2. [API设计原则](#2-api设计原则)
3. [通用API端点实施](#3-通用api端点实施)
4. [AI专用API端点实施](#4-ai专用api端点实施)
5. [安全与认证实施](#5-安全与认证实施)
6. [错误处理实施](#6-错误处理实施)
7. [API文档生成](#7-api文档生成)
8. [测试规范](#8-测试规范)
9. [部署与版本管理](#9-部署与版本管理)
10. [附录](#10-附录)

## 1. 简介

### 1.1 文档目的

本实施指南为SDK文档服务器的API开发提供详细技术指导，确保API实现符合[API规范](sdk_documentation_server_api_spec.md)中的要求，并遵循最佳实践。

### 1.2 适用范围

本文档适用于负责实现SDK文档服务器API的开发人员、测试人员和技术架构师。

### 1.3 相关文档

- [SDK文档服务器API规范](sdk_documentation_server_api_spec.md)
- [SDK文档服务器需求规格说明书](sdk_documentation_server_requirements.md)
- [SDK文档服务器系统架构](sdk_documentation_server_architecture.md)
- [SDK文档服务器实施计划](sdk_documentation_server_implementation_plan.md)

### 1.4 术语与缩略语

| 术语/缩略语 | 定义 |
|------------|------|
| API | 应用程序接口 |
| REST | 表述性状态转移，一种API设计风格 |
| JWT | JSON Web Token，一种用于认证的令牌格式 |
| JSON | JavaScript对象表示法，一种数据交换格式 |
| HTTP | 超文本传输协议 |
| CRUD | 创建、读取、更新、删除操作 |

## 2. API设计原则

### 2.1 RESTful设计准则

实现SDK文档服务器API时，应遵循以下RESTful设计准则：

1. **资源为中心**：API端点应代表资源，而非操作
   - 推荐：`/api/v1/sdks/{sdkId}`
   - 不推荐：`/api/v1/getSdk?id={sdkId}`

2. **HTTP方法语义**：正确使用HTTP方法表示操作
   - GET：读取资源
   - POST：创建资源
   - PUT：完全替换资源
   - PATCH：部分更新资源
   - DELETE：删除资源

3. **URL命名规范**：
   - 使用名词复数形式表示资源集合：`/sdks`
   - 使用单数形式表示具体资源：`/sdks/{sdkId}`
   - 使用连字符（-）连接多个单词：`/api-versions`
   - 全部使用小写字母

4. **资源层次结构**：通过URL路径表示资源间的层次关系
   - `/api/v1/sdks/{sdkId}/versions/{version}/docs`

5. **查询参数使用**：
   - 用于筛选、排序、分页等操作
   - 不应改变资源的标识或状态

### 2.2 API版本控制

SDK文档服务器API版本控制策略如下：

1. **URL路径版本控制**：
   - 在URL路径中包含版本信息：`/api/v1/...`
   - 主版本号变更表示不兼容的API变更

2. **版本兼容性要求**：
   - 同一主版本内的更新应保持向后兼容
   - 添加新字段不视为破坏性变更
   - 修改字段类型或删除字段视为破坏性变更，需增加主版本号

3. **版本过渡期**：
   - 新API版本发布后，旧版本应至少维护6个月
   - 废弃通知应提前3个月发出
   - 文档中明确标注废弃的API和字段

### 2.3 响应格式标准

所有API响应必须遵循以下JSON格式标准：

```json
{
  "status": "success|error",
  "data": { ... },
  "message": "操作成功或错误说明",
  "metadata": {
    "timestamp": "2024-05-13T12:00:00Z",
    "version": "1.0"
  }
}
```

规范要点：

1. **状态字段**：必须为"success"或"error"
2. **数据字段**：
   - 成功响应时包含请求的数据
   - 错误响应时可为null或包含错误详情
3. **消息字段**：提供人类可读的响应说明
4. **元数据字段**：
   - 必须包含timestamp（ISO 8601格式）
   - 必须包含API版本
   - 可选包含请求ID、分页信息等

## 3. 通用API端点实施

### 3.1 文档管理API实现

#### 3.1.1 获取SDK列表 (`GET /api/v1/sdks`)

**请求处理流程**：

1. 验证请求权限（需要`read`权限）
2. 处理查询参数（category, status）
3. 查询数据库获取符合条件的SDK列表
4. 格式化响应数据

**代码示例**：

```java
@RestController
@RequestMapping("/api/v1/sdks")
public class SdkController {

    @Autowired
    private SdkService sdkService;
    
    @GetMapping
    public ResponseEntity<ApiResponse> getSdks(
            @RequestParam(required = false) String category,
            @RequestParam(required = false) String status) {
        
        // 权限检查
        SecurityUtils.checkPermission("read");
        
        // 查询服务
        List<SdkDTO> sdks = sdkService.findSdks(category, status);
        
        // 构建响应
        ApiResponse response = ApiResponse.success(
            Map.of("sdks", sdks, "total", sdks.size()),
            "SDK列表获取成功"
        );
        
        return ResponseEntity.ok(response);
    }
}
```

**数据模型**：

```java
@Data
public class SdkDTO {
    private String id;
    private String name;
    private String description;
    private String latestVersion;
    private String category;
    private String status;
}
```

#### 3.1.2 获取SDK版本列表 (`GET /api/v1/sdks/{sdkId}/versions`)

**请求处理流程**：

1. 验证请求权限（需要`read`权限）
2. 验证sdkId是否存在
3. 查询数据库获取该SDK的版本列表
4. 格式化响应数据

**实现注意事项**：
- 确保SDK存在性检查
- 添加版本排序（通常按发布日期降序）
- 包含版本状态信息（stable, deprecated等）

#### 3.1.3 上传SDK文档 (`POST /api/v1/sdks/{sdkId}/versions/{version}/docs`)

**请求处理流程**：

1. 验证请求权限（需要`write`权限）
2. 验证sdkId和version参数
3. 处理上传的文件（Javadoc ZIP, OpenAPI文件等）
4. 解析并存储文档
5. 提交索引任务
6. 返回处理状态和估计完成时间

**实现注意事项**：
- 文件上传大小限制（建议最大50MB）
- 支持异步处理大型文档
- 文档处理状态跟踪
- 文档格式验证
- 支持多部分请求处理

### 3.2 搜索API实现

#### 3.2.1 基本搜索 (`GET /api/v1/search`)

**请求处理流程**：

1. 验证请求权限（需要`read`权限）
2. 解析搜索参数（q, sdkId, version, type, page, limit）
3. 构建Elasticsearch查询
4. 执行搜索并获取结果
5. 格式化搜索结果（包括高亮处理）
6. 返回分页结果

**实现注意事项**：
- 搜索性能优化（确保响应时间<1秒）
- 结果相关性排序
- 搜索词高亮处理
- 空结果处理策略
- 支持分页和限制结果数量

**Elasticsearch查询示例**：

```java
SearchSourceBuilder searchSourceBuilder = new SearchSourceBuilder();
BoolQueryBuilder boolQuery = QueryBuilders.boolQuery();

// 添加全文搜索条件
if (StringUtils.hasText(query)) {
    boolQuery.must(QueryBuilders.multiMatchQuery(query)
        .field("name", 3.0f)
        .field("description", 2.0f)
        .field("content", 1.0f)
        .type(MultiMatchQueryBuilder.Type.BEST_FIELDS));
}

// 添加过滤条件
if (StringUtils.hasText(sdkId)) {
    boolQuery.filter(QueryBuilders.termQuery("sdkId", sdkId));
}

if (StringUtils.hasText(version)) {
    boolQuery.filter(QueryBuilders.termQuery("version", version));
}

if (StringUtils.hasText(type)) {
    boolQuery.filter(QueryBuilders.termQuery("type", type));
}

// 设置分页
searchSourceBuilder.query(boolQuery);
searchSourceBuilder.from((page - 1) * limit);
searchSourceBuilder.size(limit);

// 设置高亮
HighlightBuilder highlightBuilder = new HighlightBuilder();
highlightBuilder.field(new HighlightBuilder.Field("description").fragmentSize(150).numOfFragments(3));
highlightBuilder.field(new HighlightBuilder.Field("content").fragmentSize(150).numOfFragments(3));
highlightBuilder.preTags("<em>");
highlightBuilder.postTags("</em>");
searchSourceBuilder.highlighter(highlightBuilder);

// 执行搜索
SearchRequest searchRequest = new SearchRequest(indexName);
searchRequest.source(searchSourceBuilder);
SearchResponse searchResponse = client.search(searchRequest, RequestOptions.DEFAULT);
```

#### 3.2.2 高级搜索 (`POST /api/v1/search/advanced`)

**请求处理流程**：

1. 验证请求权限（需要`read`权限）
2. 解析JSON请求体
3. 构建复杂的Elasticsearch查询
4. 执行搜索并获取结果
5. 应用自定义排序
6. 返回格式化的分页结果

**高级过滤实现**：
- 支持多字段组合过滤
- 实现自定义排序逻辑
- 支持复杂布尔逻辑（AND, OR, NOT）
- 日期范围过滤

## 4. AI专用API端点实施

### 4.1 上下文感知搜索 (`POST /api/v1/ai/search`)

**请求处理流程**：

1. 验证API密钥权限
2. 解析自然语言查询和上下文信息
3. 处理查询意图识别
4. 构建语义搜索查询
5. 获取结果并增强上下文信息
6. 返回结构化响应

**实现注意事项**：
- 使用向量嵌入模型处理自然语言查询
- 考虑上下文（之前使用的API、代码上下文等）
- 结果中包含相关API推荐
- 针对查询意图提供代码示例
- 返回包含相关度评分和推荐理由

**示例实现框架**：

```java
@RestController
@RequestMapping("/api/v1/ai")
public class AiSearchController {

    @Autowired
    private AiSearchService aiSearchService;
    @Autowired
    private QueryInterpretationService queryInterpretationService;
    @Autowired
    private RelatedApisService relatedApisService;
    @Autowired
    private CodeExampleService codeExampleService;
    
    @PostMapping("/search")
    public ResponseEntity<ApiResponse> contextAwareSearch(@RequestBody AiSearchRequest request) {
        // 权限验证
        SecurityUtils.validateApiKey();
        
        // 解析查询意图
        QueryInterpretation interpretation = queryInterpretationService.interpretQuery(
            request.getNaturalLanguageQuery(), 
            request.getContext()
        );
        
        // 执行语义搜索
        List<ApiSearchResult> results = aiSearchService.searchWithContext(
            interpretation,
            request.getContext(),
            request.getResponseOptions()
        );
        
        // 增强结果
        if (request.getResponseOptions().isIncludeExamples()) {
            results.forEach(result -> 
                result.setCodeExample(codeExampleService.generateExample(result.getApi()))
            );
        }
        
        if (request.getResponseOptions().isIncludeRelatedApis()) {
            results.forEach(result -> 
                result.setRelatedApis(relatedApisService.findRelatedApis(result.getApi().getId()))
            );
        }
        
        // 构建响应
        Map<String, Object> responseData = new HashMap<>();
        responseData.put("results", results);
        responseData.put("totalFound", results.size());
        responseData.put("queryInterpretation", interpretation);
        
        return ResponseEntity.ok(ApiResponse.success(responseData, "搜索成功"));
    }
}
```

### 4.2 获取API上下文和关系 (`GET /api/v1/ai/apis/{apiId}/context`)

**请求处理流程**：

1. 验证API密钥权限
2. 验证并解析apiId参数
3. 获取API基本信息
4. 获取指定深度的API关系图
5. 识别常见使用模式
6. 生成语义向量表示
7. 返回综合信息

**实现注意事项**：
- API依赖关系分析实现
- 使用模式识别算法
- 高效关系图查询（考虑图数据库或优化的关系查询）
- 向量嵌入模型集成
- 缓存常用API上下文信息 

### 4.3 代码示例生成 (`POST /api/v1/ai/code-examples`)

**请求处理流程**：

1. 验证API密钥权限
2. 解析请求参数（APIs列表、任务描述、编程语言、复杂度）
3. 获取指定API的完整信息
4. 根据任务生成代码示例
5. 验证代码示例的正确性
6. 返回代码示例和相关信息

**实现注意事项**：
- 组合多个API的代码生成
- 代码示例质量验证机制
- 针对不同复杂度的模板选择
- 包含必要的导入语句
- 提供解释性注释
- 维护常用代码模式库

**代码生成服务示例**：

```java
@Service
public class CodeExampleService {

    @Autowired
    private ApiRepository apiRepository;
    @Autowired
    private CodeTemplateRepository templateRepository;
    @Autowired
    private CodeVerifier codeVerifier;
    
    public CodeExample generateExample(List<String> apiIds, String task, 
                                      String language, String complexity) {
        // 获取API信息
        List<ApiEntity> apis = apiRepository.findAllByIdIn(apiIds);
        
        // 选择合适的代码模板
        CodeTemplate template = templateRepository.findBestMatch(
            apis, task, language, complexity);
        
        // 生成代码示例
        String code = renderTemplate(template, apis, task);
        
        // 提取必要的导入语句
        List<String> imports = extractImports(apis, language);
        
        // 验证代码
        boolean isValid = codeVerifier.verify(code, language);
        if (!isValid) {
            code = fixCodeIssues(code, language);
        }
        
        // 创建并返回结果
        CodeExample example = new CodeExample();
        example.setCode(code);
        example.setExplanation(generateExplanation(apis, task));
        example.setImports(imports);
        
        return example;
    }
    
    // 其他辅助方法...
}
```

### 4.4 API变更检测 (`GET /api/v1/ai/sdks/{sdkId}/changes`)

**请求处理流程**：

1. 验证API密钥权限
2. 验证sdkId和版本参数
3. 检索两个版本的API定义
4. 比较API定义找出变更
5. 分析变更的影响级别
6. 返回结构化的变更信息

**实现注意事项**：
- 高效的API差异比较算法
- 变更类型分类（添加、修改、废弃、删除）
- 变更影响级别评估（major, minor, patch）
- 提供详细的变更说明
- 统计变更摘要
- 考虑API签名、参数、返回类型和异常的变化

## 5. 安全与认证实施

### 5.1 JWT认证实现

**认证流程**：

1. 用户提交凭证（用户名/密码）
2. 验证凭证正确性
3. 生成包含用户信息和权限的JWT
4. 设置令牌过期时间和刷新机制
5. 返回JWT给客户端
6. 客户端后续请求中包含JWT
7. 服务器验证JWT有效性和权限

**实现示例**：

```java
@Component
public class JwtTokenProvider {

    @Value("${security.jwt.token.secret-key}")
    private String secretKey;
    
    @Value("${security.jwt.token.expire-length}")
    private long validityInMilliseconds;
    
    @PostConstruct
    protected void init() {
        secretKey = Base64.getEncoder().encodeToString(secretKey.getBytes());
    }
    
    public String createToken(String username, Set<String> roles) {
        Claims claims = Jwts.claims().setSubject(username);
        claims.put("roles", roles);
        
        Date now = new Date();
        Date validity = new Date(now.getTime() + validityInMilliseconds);
        
        return Jwts.builder()
            .setClaims(claims)
            .setIssuedAt(now)
            .setExpiration(validity)
            .signWith(SignatureAlgorithm.HS256, secretKey)
            .compact();
    }
    
    public Authentication getAuthentication(String token) {
        UserDetails userDetails = userDetailsService.loadUserByUsername(getUsername(token));
        return new UsernamePasswordAuthenticationToken(userDetails, "", userDetails.getAuthorities());
    }
    
    public String getUsername(String token) {
        return Jwts.parser().setSigningKey(secretKey).parseClaimsJws(token).getBody().getSubject();
    }
    
    public boolean validateToken(String token) {
        try {
            Jws<Claims> claims = Jwts.parser().setSigningKey(secretKey).parseClaimsJws(token);
            return !claims.getBody().getExpiration().before(new Date());
        } catch (JwtException | IllegalArgumentException e) {
            throw new InvalidJwtAuthenticationException("过期或无效的JWT令牌");
        }
    }
}
```

### 5.2 API密钥认证实现

**实现流程**：

1. 为AI工具生成唯一的API密钥
2. 存储API密钥与权限级别的映射
3. 验证请求头中的API密钥
4. 实施请求限制策略
5. 记录API使用情况
6. 提供API密钥管理接口

**实现注意事项**：
- 安全存储API密钥（加密或哈希）
- 密钥轮换机制
- 密钥撤销机制
- 使用情况监控和报告
- 不同级别的访问控制

### 5.3 安全最佳实践

1. **输入验证**：
   - 对所有客户端输入进行严格验证
   - 使用参数绑定和验证注解
   - 实施内容类型验证

2. **防止常见攻击**：
   - SQL注入：使用参数化查询
   - XSS攻击：输出编码和CSP
   - CSRF：使用CSRF令牌
   - 请求伪造：验证源和目标

3. **敏感数据保护**：
   - 传输中加密（TLS/SSL）
   - 存储时加密敏感数据
   - 最小化收集敏感信息
   - 实施访问控制和审计

4. **日志和监控**：
   - 记录认证尝试（成功和失败）
   - 记录敏感操作
   - 实施异常监控
   - 设置安全警报

## 6. 错误处理实施

### 6.1 标准错误响应

所有API错误应返回一致的格式：

```json
{
  "status": "error",
  "code": "ERROR_CODE",
  "message": "人类可读的错误描述",
  "details": {
    "field1": "错误详情1",
    "field2": "错误详情2"
  },
  "metadata": {
    "timestamp": "2024-05-13T12:00:00Z",
    "requestId": "a1b2c3d4"
  }
}
```

### 6.2 全局异常处理

实现全局异常处理器，确保所有异常都转换为标准格式：

```java
@RestControllerAdvice
public class GlobalExceptionHandler {

    @ExceptionHandler(ResourceNotFoundException.class)
    public ResponseEntity<ApiResponse> handleResourceNotFoundException(ResourceNotFoundException ex) {
        ApiResponse response = ApiResponse.error(
            "RESOURCE_NOT_FOUND",
            ex.getMessage(),
            ex.getDetails()
        );
        return ResponseEntity.status(HttpStatus.NOT_FOUND).body(response);
    }
    
    @ExceptionHandler(InvalidRequestException.class)
    public ResponseEntity<ApiResponse> handleInvalidRequestException(InvalidRequestException ex) {
        ApiResponse response = ApiResponse.error(
            "INVALID_REQUEST",
            ex.getMessage(),
            ex.getDetails()
        );
        return ResponseEntity.status(HttpStatus.BAD_REQUEST).body(response);
    }
    
    @ExceptionHandler(AuthenticationException.class)
    public ResponseEntity<ApiResponse> handleAuthenticationException(AuthenticationException ex) {
        ApiResponse response = ApiResponse.error(
            "UNAUTHORIZED",
            ex.getMessage(),
            null
        );
        return ResponseEntity.status(HttpStatus.UNAUTHORIZED).body(response);
    }
    
    @ExceptionHandler(AccessDeniedException.class)
    public ResponseEntity<ApiResponse> handleAccessDeniedException(AccessDeniedException ex) {
        ApiResponse response = ApiResponse.error(
            "FORBIDDEN",
            ex.getMessage(),
            null
        );
        return ResponseEntity.status(HttpStatus.FORBIDDEN).body(response);
    }
    
    @ExceptionHandler(RateLimitExceededException.class)
    public ResponseEntity<ApiResponse> handleRateLimitExceededException(RateLimitExceededException ex) {
        ApiResponse response = ApiResponse.error(
            "RATE_LIMIT_EXCEEDED",
            ex.getMessage(),
            ex.getDetails()
        );
        return ResponseEntity.status(HttpStatus.TOO_MANY_REQUESTS)
            .header("X-Rate-Limit-Retry-After-Seconds", String.valueOf(ex.getRetryAfterSeconds()))
            .body(response);
    }
    
    @ExceptionHandler(Exception.class)
    public ResponseEntity<ApiResponse> handleGenericException(Exception ex) {
        // 记录未处理的异常
        log.error("未处理的异常", ex);
        
        ApiResponse response = ApiResponse.error(
            "INTERNAL_ERROR",
            "服务器内部错误，请稍后再试",
            null
        );
        return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR).body(response);
    }
}
```

### 6.3 业务异常类

定义清晰的业务异常层次结构：

```java
// 基础API异常
public abstract class ApiException extends RuntimeException {
    private final String errorCode;
    private final Map<String, Object> details;
    
    // 构造函数、getter等
}

// 资源未找到异常
public class ResourceNotFoundException extends ApiException {
    public ResourceNotFoundException(String message) {
        super("RESOURCE_NOT_FOUND", message, null);
    }
    
    public ResourceNotFoundException(String message, Map<String, Object> details) {
        super("RESOURCE_NOT_FOUND", message, details);
    }
}

// 其他具体异常类...
```

## 7. API文档生成

### 7.1 OpenAPI规范文档

使用SpringDoc或Springfox自动生成OpenAPI规范文档：

**Maven依赖**：

```xml
<dependency>
    <groupId>org.springdoc</groupId>
    <artifactId>springdoc-openapi-ui</artifactId>
    <version>1.6.9</version>
</dependency>
```

**配置类**：

```java
@Configuration
public class OpenApiConfig {

    @Bean
    public OpenAPI sdkDocumentationServerOpenAPI() {
        return new OpenAPI()
            .info(new Info()
                .title("SDK文档服务器API")
                .description("提供SDK文档管理和搜索的API")
                .version("v1")
                .contact(new Contact()
                    .name("开发团队")
                    .email("dev@company.internal")
                    .url("https://docs.company.internal")))
            .servers(List.of(
                new Server().url("https://docs.company.internal").description("生产环境"),
                new Server().url("https://test-docs.company.internal").description("测试环境")))
            .components(new Components()
                .securitySchemes(Map.of(
                    "bearerAuth", new SecurityScheme()
                        .type(SecurityScheme.Type.HTTP)
                        .scheme("bearer")
                        .bearerFormat("JWT"),
                    "apiKeyAuth", new SecurityScheme()
                        .type(SecurityScheme.Type.APIKEY)
                        .in(SecurityScheme.In.HEADER)
                        .name("X-API-Key")
                )));
    }
}
```

**控制器注解**：

```java
@RestController
@RequestMapping("/api/v1/sdks")
@Tag(name = "SDK管理", description = "SDK文档管理API")
public class SdkController {

    @GetMapping
    @Operation(
        summary = "获取SDK列表",
        description = "返回所有可用的SDK列表，支持按类别和状态筛选",
        security = { @SecurityRequirement(name = "bearerAuth") }
    )
    @ApiResponses({
        @ApiResponse(responseCode = "200", description = "成功获取SDK列表"),
        @ApiResponse(responseCode = "401", description = "未授权访问"),
        @ApiResponse(responseCode = "403", description = "权限不足")
    })
    public ResponseEntity<ApiResponse> getSdks(
            @Parameter(description = "按类别筛选") @RequestParam(required = false) String category,
            @Parameter(description = "按状态筛选") @RequestParam(required = false) String status) {
        // 方法实现...
    }
    
    // 其他方法...
}
```

### 7.2 API文档页面

配置Swagger UI为开发人员提供交互式API文档：

```java
@Configuration
public class SwaggerUiConfig {

    @Bean
    public WebMvcConfigurer webMvcConfigurer() {
        return new WebMvcConfigurer() {
            @Override
            public void addViewControllers(ViewControllerRegistry registry) {
                registry.addViewController("/swagger-ui/")
                    .setViewName("forward:/swagger-ui/index.html");
                registry.addViewController("/api-docs")
                    .setViewName("forward:/api-docs/index.html");
            }
        };
    }
}
```

## 8. 测试规范

### 8.1 单元测试

**控制器测试示例**：

```java
@WebMvcTest(SdkController.class)
@ExtendWith(SpringExtension.class)
public class SdkControllerTest {

    @Autowired
    private MockMvc mockMvc;
    
    @MockBean
    private SdkService sdkService;
    
    @Test
    public void testGetSdks() throws Exception {
        // 准备测试数据
        List<SdkDTO> sdks = Arrays.asList(
            new SdkDTO("sdk1", "SDK 1", "描述1", "1.0.0", "分类1", "active"),
            new SdkDTO("sdk2", "SDK 2", "描述2", "2.0.0", "分类2", "active")
        );
        
        when(sdkService.findSdks(any(), any())).thenReturn(sdks);
        
        // 执行测试
        mockMvc.perform(get("/api/v1/sdks")
                .contentType(MediaType.APPLICATION_JSON))
            .andExpect(status().isOk())
            .andExpect(jsonPath("$.status").value("success"))
            .andExpect(jsonPath("$.data.sdks").isArray())
            .andExpect(jsonPath("$.data.sdks.length()").value(2))
            .andExpect(jsonPath("$.data.sdks[0].id").value("sdk1"))
            .andExpect(jsonPath("$.data.total").value(2));
        
        // 验证服务调用
        verify(sdkService).findSdks(any(), any());
    }
    
    // 其他测试方法...
}
```

### 8.2 集成测试

**API集成测试示例**：

```java
@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
public class SdkApiIntegrationTest {

    @Autowired
    private TestRestTemplate restTemplate;
    
    @Autowired
    private SdkRepository sdkRepository;
    
    @BeforeEach
    public void setup() {
        // 设置测试数据
        sdkRepository.deleteAll();
        sdkRepository.save(new SdkEntity("sdk1", "SDK 1", "描述1", "分类1", "active"));
    }
    
    @Test
    public void testGetSdks() {
        // 执行API调用
        ResponseEntity<String> response = restTemplate.exchange(
            "/api/v1/sdks",
            HttpMethod.GET,
            null,
            String.class
        );
        
        // 验证响应
        assertThat(response.getStatusCode()).isEqualTo(HttpStatus.OK);
        
        JsonNode root = parseJson(response.getBody());
        assertThat(root.path("status").asText()).isEqualTo("success");
        assertThat(root.path("data").path("sdks").isArray()).isTrue();
        assertThat(root.path("data").path("sdks").size()).isEqualTo(1);
        assertThat(root.path("data").path("sdks").get(0).path("id").asText()).isEqualTo("sdk1");
    }
    
    // 辅助方法
    private JsonNode parseJson(String json) {
        try {
            return new ObjectMapper().readTree(json);
        } catch (Exception e) {
            throw new RuntimeException("解析JSON失败", e);
        }
    }
    
    // 其他测试方法...
}
```

### 8.3 性能测试

**使用JMeter进行API性能测试**：

1. 创建测试计划，针对关键API端点
2. 设置测试场景：
   - 正常负载：50并发用户，持续5分钟
   - 峰值负载：100并发用户，持续2分钟
   - 耐久测试：75并发用户，持续30分钟
3. 设置性能指标：
   - 响应时间阈值：P95 < 1秒
   - 吞吐量：至少100请求/秒
   - 错误率：低于1%
4. 自动化测试执行
5. 分析结果并优化性能

## 9. 部署与版本管理

### 9.1 API版本发布流程

新API版本发布流程如下：

1. **规划阶段**：
   - 确定API变更范围
   - 评估兼容性影响
   - 确定版本号变更

2. **开发阶段**：
   - 实现新功能和变更
   - 撰写API文档
   - 编写单元和集成测试

3. **验证阶段**：
   - 执行完整测试套件
   - 性能测试验证
   - API契约验证

4. **发布阶段**：
   - 部署新版本
   - 更新API文档门户
   - 发布变更通知
   - 监控初期使用情况

5. **维护阶段**：
   - 处理反馈和问题
   - 规划后续改进

### 9.2 API废弃策略

当需要废弃API功能时，应遵循以下策略：

1. **标记废弃**：
   - 在文档中明确标记为废弃
   - 添加废弃通知和预计移除日期
   - 提供替代方案说明

2. **通知流程**：
   - 提前3个月发出电子邮件通知
   - 在API响应中添加废弃警告头
   - 在API文档门户显示醒目通知

3. **过渡期**：
   - 维持废弃API至少6个月
   - 提供迁移工具和指南
   - 持续监控废弃API的使用情况

4. **最终移除**：
   - 再次通知用户即将移除
   - 按计划移除废弃功能
   - 确保404响应包含说明性消息

## 10. 附录

### 10.1 API响应示例

**成功响应示例**：

```json
{
  "status": "success",
  "data": {
    "sdks": [
      {
        "id": "core-sdk",
        "name": "Core SDK",
        "description": "核心功能SDK",
        "latestVersion": "2.3.1",
        "category": "基础框架",
        "status": "active"
      },
      {
        "id": "image-processing",
        "name": "Image Processing SDK",
        "description": "图像处理SDK",
        "latestVersion": "1.5.0",
        "category": "媒体处理",
        "status": "active"
      }
    ],
    "total": 2
  },
  "message": "成功获取SDK列表",
  "metadata": {
    "timestamp": "2024-05-13T12:00:00Z",
    "version": "1.0"
  }
}
```

**错误响应示例**：

```json
{
  "status": "error",
  "code": "RESOURCE_NOT_FOUND",
  "message": "请求的SDK不存在",
  "details": {
    "sdkId": "non-existent-sdk"
  },
  "metadata": {
    "timestamp": "2024-05-13T12:00:00Z",
    "requestId": "a1b2c3d4"
  }
}
```

### 10.2 API端点清单

| 类别 | 端点 | 方法 | 描述 | 权限 |
|------|------|------|------|------|
| **文档管理** | /api/v1/sdks | GET | 获取SDK列表 | read |
|  | /api/v1/sdks/{sdkId}/versions | GET | 获取SDK版本列表 | read |
|  | /api/v1/sdks/{sdkId}/versions/{version}/docs | POST | 上传SDK文档 | write |
| **搜索功能** | /api/v1/search | GET | 基本搜索 | read |
|  | /api/v1/search/advanced | POST | 高级搜索 | read |
| **AI专用API** | /api/v1/ai/search | POST | 上下文感知搜索 | API密钥 |
|  | /api/v1/ai/apis/{apiId}/context | GET | 获取API上下文和关系 | API密钥 |
|  | /api/v1/ai/code-examples | POST | 代码示例生成 | API密钥 |
|  | /api/v1/ai/sdks/{sdkId}/changes | GET | API变更检测 | API密钥 |

### 10.3 相关技术文档链接

- [Spring Boot官方文档](https://spring.io/projects/spring-boot)
- [Spring Security参考](https://docs.spring.io/spring-security/reference/index.html)
- [OpenAPI规范](https://swagger.io/specification/)
- [JWT标准](https://jwt.io/introduction)
- [Elasticsearch指南](https://www.elastic.co/guide/index.html)
- [Spring Data Elasticsearch](https://docs.spring.io/spring-data/elasticsearch/docs/current/reference/html/)
- [JUnit 5用户指南](https://junit.org/junit5/docs/current/user-guide/) 