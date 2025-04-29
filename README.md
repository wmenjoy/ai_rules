# 提示词模板管理系统 (Prompt Template Management System)

一个基于 Golang + MySQL + GORM 的提示词模板管理系统后端服务，提供 API 以管理提示词模板，包括模板的增删改查、分类和标签管理、导入/导出功能和统计分析。

## 功能特性

- 模板管理（创建、编辑、删除、复制）
- 模板变量支持
- 模板分类和标签管理
- 模板收藏和使用统计
- 模板的导入和导出
- 使用分析和统计

## 技术栈

- **后端**：Golang
- **数据库**：MySQL
- **ORM**：GORM
- **Web 框架**：Gin
- **配置管理**：环境变量 + .env 文件

## 项目结构

```
prompt-template-system/
├── cmd/
│   └── server/            # 应用入口
│       └── main.go
├── internal/
│   ├── api/               # API 层
│   │   ├── handlers/      # 请求处理器
│   │   ├── middleware/    # 中间件
│   │   └── routes.go      # 路由定义
│   ├── models/            # 数据模型
│   ├── repository/        # 数据仓库层
│   ├── service/           # 业务逻辑层
│   └── database/          # 数据库连接
├── pkg/
│   ├── config/            # 配置管理
│   └── utils/             # 工具函数
├── docs/
│   ├── api.md             # API 文档
│   └── db.sql             # 数据库脚本
└── go.mod                 # Go 模块定义
```

## 环境要求

- Go 1.16+
- MySQL 5.7+

## 快速开始

### 1. 克隆仓库

```bash
git clone https://github.com/yourusername/prompt-template-system.git
cd prompt-template-system
```

### 2. 配置环境变量

在项目根目录创建 `.env` 文件：

```
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=prompt_template_system
SERVER_PORT=8080
```

### 3. 初始化数据库

使用 `docs/db.sql` 文件创建数据库和表：

```bash
mysql -u root -p < docs/db.sql
```

### 4. 构建和运行

```bash
go mod tidy
go build -o server ./cmd/server
./server
```

或者直接运行：

```bash
go run ./cmd/server
```

服务将在 http://localhost:8080 运行。

## API 文档

API 文档可在 `docs/api.md` 中找到，包括以下主要端点：

- `/api/templates` - 模板管理
- `/api/categories` - 分类管理
- `/api/tags` - 标签管理
- `/api/analytics` - 统计和分析

## 开发

### 项目依赖

```bash
go mod tidy
```

### 运行测试

```bash
go test ./...
```

## 许可证

MIT