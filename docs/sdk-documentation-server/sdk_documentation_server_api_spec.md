# SDK 文档服务器 API 规范

**文档版本**: 1.0  
**更新日期**: 2024-05-13  
**状态**: 初稿  

## 目录

1. [API 概述](#1-api-概述)
2. [认证与授权](#2-认证与授权)
3. [通用API端点](#3-通用api端点)
4. [AI专用API端点](#4-ai专用api端点)
5. [错误处理](#5-错误处理)
6. [限流策略](#6-限流策略)
7. [示例](#7-示例)

## 1. API 概述

SDK文档服务器提供REST风格的API，用于访问和搜索公司内部SDK的文档。API分为两类：
1. 通用API：供所有客户端使用，包括Web界面、CLI工具等
2. AI专用API：专为AI编程工具优化，提供更丰富的元数据和上下文信息

### 1.1 基本URL结构

```
https://docs.company.internal/api/v1/...
```

### 1.2 响应格式

所有API响应均使用JSON格式，包含以下标准字段：

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

## 2. 认证与授权

### 2.1 认证方式

支持以下认证方式：

1. **JWT Bearer Token**：对于Web界面和一般用户
   ```
   Authorization: Bearer <token>
   ```

2. **API密钥**：主要用于AI工具和系统集成
   ```
   X-API-Key: <api_key>
   ```

### 2.2 权限级别

- `read`: 只读权限，可搜索和查看文档
- `write`: 可上传和管理文档
- `admin`: 完全管理权限

## 3. 通用API端点

### 3.1 文档管理

#### 3.1.1 获取SDK列表

```
GET /api/v1/sdks
```

**查询参数**:
- `category` (可选): 按类别筛选
- `status` (可选): 状态筛选（active, deprecated, beta）

**响应示例**:
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
  }
}
```

#### 3.1.2 获取SDK版本列表

```
GET /api/v1/sdks/{sdkId}/versions
```

**路径参数**:
- `sdkId`: SDK标识符

**响应示例**:
```json
{
  "status": "success",
  "data": {
    "versions": [
      {
        "version": "2.3.1",
        "releaseDate": "2024-04-15",
        "status": "stable"
      },
      {
        "version": "2.3.0",
        "releaseDate": "2024-03-10",
        "status": "stable"
      },
      {
        "version": "2.2.0",
        "releaseDate": "2024-01-20",
        "status": "deprecated"
      }
    ]
  }
}
```

#### 3.1.3 上传SDK文档

```
POST /api/v1/sdks/{sdkId}/versions/{version}/docs
```

**路径参数**:
- `sdkId`: SDK标识符
- `version`: 版本号

**请求体**:
`multipart/form-data` 包含以下字段:
- `javadoc`: Javadoc ZIP文件
- `openapi`: OpenAPI规范文件 (可选)
- `metadata`: 文档元数据 (JSON格式，可选)

**响应示例**:
```json
{
  "status": "success",
  "data": {
    "docId": "12345",
    "processingStatus": "queued",
    "estimatedCompletionTime": "2024-05-13T12:30:00Z"
  }
}
```

### 3.2 文档搜索

#### 3.2.1 基本搜索

```
GET /api/v1/search
```

**查询参数**:
- `q`: 搜索关键词
- `sdkId` (可选): 特定SDK
- `version` (可选): 特定版本
- `type` (可选): 文档类型（javadoc, openapi）
- `page` (可选): 页码，默认1
- `limit` (可选): 每页结果数，默认20

**响应示例**:
```json
{
  "status": "success",
  "data": {
    "results": [
      {
        "id": "com.company.sdk.core.ImageProcessor",
        "type": "class",
        "name": "ImageProcessor",
        "package": "com.company.sdk.core",
        "sdkId": "core-sdk",
        "version": "2.3.1",
        "description": "处理图像的核心类",
        "url": "/javadoc/core-sdk/2.3.1/com/company/sdk/core/ImageProcessor.html",
        "highlights": [
          "...处理<em>图像</em>的核心类，提供..."
        ]
      }
    ],
    "total": 1,
    "page": 1,
    "limit": 20,
    "totalPages": 1
  }
}
```

#### 3.2.2 高级搜索

```
POST /api/v1/search/advanced
```

**请求体**:
```json
{
  "query": "图像处理",
  "filters": {
    "sdkId": ["core-sdk", "image-processing"],
    "type": ["class", "method"],
    "access": ["public"],
    "deprecated": false
  },
  "sort": {
    "field": "relevance",
    "order": "desc"
  },
  "page": 1,
  "limit": 20
}
```

**响应格式**与基本搜索相同。

## 4. AI专用API端点

这些API端点专为AI编程工具优化，提供更丰富的上下文和结构化数据。

### 4.1 上下文感知搜索

```
POST /api/v1/ai/search
```

**请求体**:
```json
{
  "naturalLanguageQuery": "如何处理图像并应用滤镜",
  "context": {
    "codeContext": "用户正在编写图像处理应用",
    "previousApis": ["com.company.sdk.core.ImageLoader"]
  },
  "responseOptions": {
    "includeExamples": true,
    "includeRelatedApis": true,
    "maxResults": 5
  }
}
```

**响应示例**:
```json
{
  "status": "success",
  "data": {
    "results": [
      {
        "api": {
          "id": "com.company.sdk.core.ImageProcessor",
          "type": "class",
          "name": "ImageProcessor",
          "package": "com.company.sdk.core",
          "sdkId": "core-sdk",
          "version": "2.3.1",
          "description": "处理图像的核心类",
          "url": "/javadoc/core-sdk/2.3.1/com/company/sdk/core/ImageProcessor.html"
        },
        "relevance": 0.95,
        "reason": "此类专门用于图像处理，包含应用滤镜的方法",
        "codeExample": "ImageProcessor processor = new ImageProcessor();\nImage processedImage = processor.applyFilter(image, FilterType.BLUR);",
        "relatedApis": [
          {
            "id": "com.company.sdk.core.filters.FilterType",
            "relationship": "USED_WITH",
            "description": "定义可用的滤镜类型"
          }
        ]
      }
    ],
    "totalFound": 1,
    "queryInterpretation": {
      "intent": "FIND_IMAGE_PROCESSING_FILTER_API",
      "entities": [
        {
          "type": "TASK",
          "value": "图像处理"
        },
        {
          "type": "FEATURE",
          "value": "滤镜"
        }
      ]
    }
  }
}
```

### 4.2 获取API上下文和关系

```
GET /api/v1/ai/apis/{apiId}/context
```

**路径参数**:
- `apiId`: API标识符 (如 `com.company.sdk.core.ImageProcessor`)

**查询参数**:
- `version` (可选): 特定版本
- `depth` (可选): 关系深度，默认1

**响应示例**:
```json
{
  "status": "success",
  "data": {
    "api": {
      "id": "com.company.sdk.core.ImageProcessor",
      "type": "class",
      "name": "ImageProcessor",
      "package": "com.company.sdk.core",
      "sdkId": "core-sdk",
      "version": "2.3.1",
      "description": "处理图像的核心类",
      "url": "/javadoc/core-sdk/2.3.1/com/company/sdk/core/ImageProcessor.html",
      "methods": [
        {
          "name": "applyFilter",
          "signature": "Image applyFilter(Image image, FilterType type)",
          "description": "应用指定的滤镜到图像",
          "access": "public",
          "deprecated": false
        }
      ]
    },
    "relationships": [
      {
        "type": "USES",
        "target": "com.company.sdk.core.Image",
        "description": "使用此类作为输入和输出参数"
      },
      {
        "type": "USES",
        "target": "com.company.sdk.core.filters.FilterType",
        "description": "使用此枚举指定滤镜类型"
      },
      {
        "type": "USED_BY",
        "target": "com.company.sdk.ui.ImageEditor",
        "description": "被此UI组件使用"
      }
    ],
    "usagePatterns": [
      {
        "pattern": "图像滤镜应用",
        "description": "用于将滤镜应用到图像",
        "frequency": "高",
        "codeExample": "ImageProcessor processor = new ImageProcessor();\nImage processedImage = processor.applyFilter(image, FilterType.BLUR);"
      }
    ],
    "semanticVector": [0.2, 0.5, -0.3, 0.7, ...] // 向量嵌入表示，用于相似度计算
  }
}
```

### 4.3 代码示例生成

```
POST /api/v1/ai/code-examples
```

**请求体**:
```json
{
  "apis": ["com.company.sdk.core.ImageProcessor", "com.company.sdk.core.filters.FilterType"],
  "task": "加载图像并应用模糊滤镜",
  "language": "java",
  "complexity": "simple"
}
```

**响应示例**:
```json
{
  "status": "success",
  "data": {
    "example": {
      "code": "// 导入必要的类\nimport com.company.sdk.core.Image;\nimport com.company.sdk.core.ImageLoader;\nimport com.company.sdk.core.ImageProcessor;\nimport com.company.sdk.core.filters.FilterType;\n\n// 加载图像\nImageLoader loader = new ImageLoader();\nImage image = loader.load(\"input.jpg\");\n\n// 应用模糊滤镜\nImageProcessor processor = new ImageProcessor();\nImage blurredImage = processor.applyFilter(image, FilterType.BLUR);\n\n// 保存处理后的图像\nblurredImage.save(\"output.jpg\");",
      "explanation": "此代码示例展示了如何使用ImageLoader加载图像，然后使用ImageProcessor应用模糊滤镜，最后保存处理后的图像。",
      "imports": [
        "com.company.sdk.core.Image",
        "com.company.sdk.core.ImageLoader",
        "com.company.sdk.core.ImageProcessor",
        "com.company.sdk.core.filters.FilterType"
      ]
    },
    "additionalExamples": [
      {
        "name": "应用多个滤镜",
        "summaryOnly": true,
        "url": "/api/v1/ai/code-examples/23456"
      }
    ]
  }
}
```

### 4.4 API变更检测

```
GET /api/v1/ai/sdks/{sdkId}/changes
```

**路径参数**:
- `sdkId`: SDK标识符

**查询参数**:
- `fromVersion`: 起始版本
- `toVersion`: 目标版本

**响应示例**:
```json
{
  "status": "success",
  "data": {
    "changes": [
      {
        "type": "ADDED",
        "apiId": "com.company.sdk.core.filters.AdvancedFilters",
        "description": "添加了新的高级滤镜类",
        "impactLevel": "minor"
      },
      {
        "type": "MODIFIED",
        "apiId": "com.company.sdk.core.ImageProcessor",
        "description": "applyFilter方法添加了新的参数",
        "impactLevel": "major",
        "details": {
          "before": "Image applyFilter(Image image, FilterType type)",
          "after": "Image applyFilter(Image image, FilterType type, float intensity)"
        }
      },
      {
        "type": "DEPRECATED",
        "apiId": "com.company.sdk.core.legacy.OldImageProcessor",
        "description": "标记为废弃，推荐使用ImageProcessor代替",
        "impactLevel": "major"
      }
    ],
    "summary": {
      "added": 5,
      "modified": 3,
      "deprecated": 2,
      "removed": 0
    }
  }
}
```

## 5. 错误处理

所有API错误使用标准HTTP状态码，并在响应体中提供详细信息：

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

### 5.1 常见错误码

| HTTP状态码 | 错误码 | 描述 |
|------------|--------|------|
| 400 | INVALID_REQUEST | 请求参数无效 |
| 401 | UNAUTHORIZED | 未授权访问 |
| 403 | FORBIDDEN | 权限不足 |
| 404 | RESOURCE_NOT_FOUND | 资源不存在 |
| 429 | RATE_LIMIT_EXCEEDED | 超出请求限制 |
| 500 | INTERNAL_ERROR | 服务器内部错误 |

## 6. 限流策略

为确保系统稳定性和公平使用，API实施了以下限流策略：

- 通用API：
  - 认证用户: 100请求/分钟
  - 匿名用户: 20请求/分钟

- AI专用API：
  - 基础级: 50请求/分钟
  - 专业级: 200请求/分钟
  - 企业级: 1000请求/分钟

限流信息在响应头中提供：
```
X-Rate-Limit-Limit: 100
X-Rate-Limit-Remaining: 95
X-Rate-Limit-Reset: 1589458986
```

## 7. 示例

### 7.1 使用AI API查找处理图像的类

**请求**:
```
POST /api/v1/ai/search
Content-Type: application/json
X-API-Key: api_key_here

{
  "naturalLanguageQuery": "我需要加载和编辑图像",
  "responseOptions": {
    "includeExamples": true,
    "maxResults": 3
  }
}
```

**响应**:
```json
{
  "status": "success",
  "data": {
    "results": [
      {
        "api": {
          "id": "com.company.sdk.core.ImageLoader",
          "name": "ImageLoader",
          "description": "用于加载各种格式的图像文件"
        },
        "relevance": 0.92,
        "reason": "此类专门用于图像加载，是处理图像的第一步",
        "codeExample": "ImageLoader loader = new ImageLoader();\nImage image = loader.load(\"image.jpg\");"
      },
      {
        "api": {
          "id": "com.company.sdk.core.ImageProcessor",
          "name": "ImageProcessor",
          "description": "处理图像的核心类"
        },
        "relevance": 0.85,
        "reason": "此类提供了图像处理和编辑功能",
        "codeExample": "ImageProcessor processor = new ImageProcessor();\nImage processed = processor.applyFilter(image, FilterType.SHARPEN);"
      }
    ],
    "queryInterpretation": {
      "intent": "FIND_IMAGE_HANDLING_API",
      "entities": [
        {"type": "TASK", "value": "加载图像"},
        {"type": "TASK", "value": "编辑图像"}
      ]
    }
  }
}
```

### 7.2 获取API上下文和关系

**请求**:
```
GET /api/v1/ai/apis/com.company.sdk.core.ImageProcessor/context?version=2.3.1&depth=1
X-API-Key: api_key_here
```

**响应**:
参见[4.2 获取API上下文和关系](#42-获取api上下文和关系)的示例响应。 