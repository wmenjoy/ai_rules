# AI代码安全扫描系统软件设计书 (SDD)

## 文档信息

| 项目 | 内容 |
|------|------|
| 文档标题 | AI代码安全扫描系统软件设计书 |
| 文档版本 | v1.0 |
| 创建日期 | 2024-12-19 |
| 最后更新 | 2024-12-19 |
| 文档状态 | 草案 |
| 作者 | 系统架构师 |
| 审阅者 | 技术总监、高级工程师 |
| 批准者 | 技术委员会 |

## 1. 引言

### 1.1 目的
本文档描述了AI代码安全扫描系统的软件架构设计、组件设计、接口设计和数据设计，为系统实现提供详细的技术指导。

### 1.2 范围
本设计涵盖系统的所有软件组件，包括：
- 微服务架构设计
- AI智能体系统设计
- 数据存储和处理设计
- 用户界面设计
- 集成接口设计

### 1.3 定义和缩写

| 术语 | 定义 |
|------|------|
| DDD | 领域驱动设计(Domain-Driven Design) |
| CQRS | 命令查询职责分离(Command Query Responsibility Segregation) |
| Event Sourcing | 事件溯源 |
| Microservices | 微服务架构 |
| API Gateway | API网关 |
| Service Mesh | 服务网格 |
| Container | 容器 |
| Orchestration | 编排 |

### 1.4 参考文献
- 《微服务架构设计模式》
- 《领域驱动设计》
- 《软件架构实践》
- Kubernetes官方文档
- Go语言设计规范

### 1.5 概述
本文档按照架构设计、详细设计、用户界面设计的顺序组织，每个部分都包含设计原理、实现方案和关键决策。

## 2. 设计考量

### 2.1 设计假设

#### 2.1.1 技术假设
- Kubernetes集群稳定可用
- Go语言生态成熟稳定
- AI模型推理性能可接受
- 网络延迟在可控范围内

#### 2.1.2 业务假设
- 用户具备基本的DevOps知识
- 代码仓库规模在合理范围内
- 扫描频率不会过于频繁
- 用户对AI建议有一定信任度

### 2.2 设计约束

#### 2.2.1 技术约束
- 必须支持本地化部署
- 主要支持Java生态系统
- 需要与SonarQube集成
- 支持混合云架构

#### 2.2.2 性能约束
- 扫描响应时间要求严格
- 系统资源使用需要优化
- 并发处理能力有限制
- AI模型推理时间有要求

#### 2.2.3 安全约束
- 代码数据不能外泄
- 用户权限严格控制
- 系统访问需要审计
- 敏感信息必须加密

### 2.3 设计依赖

#### 2.3.1 外部依赖
- Kubernetes容器编排平台
- PostgreSQL数据库系统
- Redis缓存系统
- MinIO对象存储系统
- RabbitMQ消息队列

#### 2.3.2 内部依赖
- AI模型和算法库
- 代码解析引擎
- 安全规则库
- 报告模板系统

## 3. 架构设计

### 3.1 系统架构概览

#### 3.1.1 整体架构图

```
┌─────────────────────────────────────────────────────────────────┐
│                        用户接入层                                │
├─────────────────┬─────────────────┬─────────────────────────────┤
│   Web Portal    │   CLI Tools     │      IDE Plugins            │
│                 │                 │                             │
│ • React前端     │ • Go CLI        │ • VS Code插件               │
│ • TypeScript    │ • 批处理脚本     │ • IntelliJ插件              │
└─────────────────┴─────────────────┴─────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────────┐
│                        API网关层                                │
├─────────────────────────────────────────────────────────────────┤
│                    Kong API Gateway                            │
│                                                                 │
│ • 路由和负载均衡  • 认证和授权  • 限流和熔断  • 监控和日志        │
└─────────────────────────────────────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────────┐
│                        微服务层                                 │
├─────────────┬─────────────┬─────────────┬─────────────────────────┤
│  用户服务    │  项目服务    │  扫描服务    │      AI智能体服务        │
│             │             │             │                        │
│ • 用户管理   │ • 项目管理   │ • 扫描引擎   │ • 代码理解智能体        │
│ • 权限控制   │ • 仓库同步   │ • 规则引擎   │ • 安全检测智能体        │
│ • 认证授权   │ • 配置管理   │ • 结果处理   │ • 修复建议智能体        │
└─────────────┴─────────────┴─────────────┴─────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────────┐
│                        数据存储层                               │
├─────────────┬─────────────┬─────────────┬─────────────────────────┤
│ PostgreSQL  │   Redis     │   MinIO     │      消息队列            │
│             │             │             │                        │
│ • 业务数据   │ • 缓存数据   │ • 文件存储   │ • RabbitMQ             │
│ • 用户信息   │ • 会话数据   │ • 报告文件   │ • 异步任务队列          │
│ • 扫描结果   │ • 临时数据   │ • 日志文件   │ • 事件通知              │
└─────────────┴─────────────┴─────────────┴─────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────────┐
│                        基础设施层                               │
├─────────────┬─────────────┬─────────────┬─────────────────────────┤
│ Kubernetes  │  监控系统    │  日志系统    │      安全系统            │
│             │             │             │                        │
│ • 容器编排   │ • Prometheus │ • ELK Stack │ • TLS证书管理           │
│ • 服务发现   │ • Grafana   │ • 日志聚合   │ • 密钥管理              │
│ • 负载均衡   │ • 告警通知   │ • 日志分析   │ • 网络安全              │
└─────────────┴─────────────┴─────────────┴─────────────────────────┘
```

#### 3.1.2 架构特点

**微服务架构**
- 服务自治：每个服务独立开发、部署和扩展
- 技术多样性：不同服务可以选择最适合的技术栈
- 故障隔离：单个服务故障不影响整体系统
- 团队独立：不同团队可以独立维护各自的服务

**事件驱动架构**
- 异步通信：通过事件实现服务间的松耦合
- 可扩展性：易于添加新的事件处理器
- 可靠性：事件持久化保证消息不丢失
- 可追溯性：完整的事件历史记录

**领域驱动设计**
- 业务导向：以业务领域为中心组织代码
- 清晰边界：明确的领域边界和上下文
- 丰富模型：业务逻辑封装在领域模型中
- 统一语言：团队使用统一的业务术语

### 3.2 微服务设计

#### 3.2.1 服务拆分原则

**业务能力拆分**
- 用户管理：用户认证、授权、配置管理
- 项目管理：代码仓库、项目配置、版本控制
- 扫描服务：代码分析、规则执行、结果处理
- AI服务：智能分析、模型推理、建议生成
- 报告服务：报告生成、数据可视化、导出功能
- 通知服务：消息通知、邮件发送、Webhook

**数据一致性边界**
- 每个服务拥有独立的数据存储
- 通过事件实现最终一致性
- 避免分布式事务的复杂性
- 使用Saga模式处理跨服务事务

**团队组织结构**
- 每个服务对应一个开发团队
- 团队负责服务的全生命周期
- 明确的服务所有权和责任
- 跨团队协作通过API契约

#### 3.2.2 服务间通信

**同步通信**
- RESTful API：用于查询和简单操作
- gRPC：用于高性能的服务间调用
- GraphQL：用于复杂的数据查询

**异步通信**
- 事件发布/订阅：用于业务事件通知
- 消息队列：用于任务处理和工作流
- 流处理：用于实时数据处理

**通信模式**
```
请求/响应模式：
Client → Service A → Service B → Response

事件驱动模式：
Service A → Event Bus → Service B
                    → Service C
                    → Service D

管道模式：
Input → Service A → Service B → Service C → Output
```

### 3.3 AI智能体架构

#### 3.3.1 智能体系统设计

**多智能体协作架构**
```
┌─────────────────────────────────────────────────────────────────┐
│                    AI智能体协调器                                │
├─────────────────────────────────────────────────────────────────┤
│ • 任务分发  • 结果聚合  • 冲突解决  • 性能监控                   │
└─────────────────────────────────────────────────────────────────┘
                                │
                ┌───────────────┼───────────────┐
                ▼               ▼               ▼
┌─────────────────┐ ┌─────────────────┐ ┌─────────────────┐
│  代码理解智能体  │ │  安全检测智能体  │ │  修复建议智能体  │
├─────────────────┤ ├─────────────────┤ ├─────────────────┤
│ • 语法分析      │ │ • 漏洞检测      │ │ • 修复方案      │
│ • 语义理解      │ │ • 风险评估      │ │ • 代码生成      │
│ • 上下文分析    │ │ • 合规检查      │ │ • 最佳实践      │
│ • 模式识别      │ │ • 威胁建模      │ │ • 重构建议      │
└─────────────────┘ └─────────────────┘ └─────────────────┘
```

**智能体特性**
- 自主性：能够独立执行分配的任务
- 反应性：能够感知环境变化并做出响应
- 主动性：能够主动发起行动以达成目标
- 社会性：能够与其他智能体协作和通信

#### 3.2.2 AI模型集成

**本地模型部署**
```yaml
# AI模型部署配置
apiVersion: apps/v1
kind: Deployment
metadata:
  name: ai-model-service
spec:
  replicas: 3
  selector:
    matchLabels:
      app: ai-model
  template:
    metadata:
      labels:
        app: ai-model
    spec:
      containers:
      - name: model-server
        image: ai-scanner/model-server:latest
        resources:
          requests:
            memory: "8Gi"
            cpu: "4"
            nvidia.com/gpu: "1"
          limits:
            memory: "16Gi"
            cpu: "8"
            nvidia.com/gpu: "1"
        env:
        - name: MODEL_PATH
          value: "/models/codebert"
        - name: BATCH_SIZE
          value: "32"
        volumeMounts:
        - name: model-storage
          mountPath: /models
      volumes:
      - name: model-storage
        persistentVolumeClaim:
          claimName: model-pvc
```

**模型管理策略**
- 模型版本控制：使用Git LFS管理模型文件
- 热更新机制：支持在线模型更新
- A/B测试：新旧模型并行运行对比
- 性能监控：实时监控模型推理性能
- 降级策略：模型故障时的备用方案

### 3.4 数据架构设计

#### 3.4.1 数据存储策略

**多存储引擎**
```
┌─────────────────────────────────────────────────────────────────┐
│                        数据分层架构                             │
├─────────────────┬─────────────────┬─────────────────────────────┤
│    关系数据      │     缓存数据     │        文件数据              │
│                 │                 │                             │
│ PostgreSQL      │ Redis Cluster   │ MinIO Object Storage        │
│                 │                 │                             │
│ • 用户信息       │ • 会话缓存       │ • 代码文件                   │
│ • 项目配置       │ • 查询缓存       │ • 扫描报告                   │
│ • 扫描结果       │ • 分布式锁       │ • 日志文件                   │
│ • 权限数据       │ • 计数器         │ • 模型文件                   │
└─────────────────┴─────────────────┴─────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────────┐
│                        数据处理层                               │
├─────────────────┬─────────────────┬─────────────────────────────┤
│    流处理        │     批处理       │        搜索引擎              │
│                 │                 │                             │
│ Apache Kafka    │ Apache Spark    │ Elasticsearch               │
│                 │                 │                             │
│ • 实时事件       │ • 大数据分析     │ • 全文搜索                   │
│ • 日志流         │ • 报告生成       │ • 日志检索                   │
│ • 监控指标       │ • 数据清洗       │ • 代码搜索                   │
└─────────────────┴─────────────────┴─────────────────────────────┘
```

**数据一致性模型**
- 强一致性：用户认证、权限管理
- 最终一致性：扫描结果、统计数据
- 弱一致性：日志数据、监控指标

#### 3.4.2 数据模型设计

**核心实体关系**
```sql
-- 用户表
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL DEFAULT 'developer',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 项目表
CREATE TABLE projects (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    repository_url VARCHAR(500) NOT NULL,
    branch VARCHAR(100) DEFAULT 'main',
    owner_id UUID REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 扫描任务表
CREATE TABLE scan_tasks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    project_id UUID REFERENCES projects(id),
    task_type VARCHAR(20) NOT NULL, -- 'full', 'incremental'
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    started_at TIMESTAMP,
    completed_at TIMESTAMP,
    created_by UUID REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 扫描结果表
CREATE TABLE scan_results (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    task_id UUID REFERENCES scan_tasks(id),
    file_path VARCHAR(1000) NOT NULL,
    line_number INTEGER,
    column_number INTEGER,
    rule_id VARCHAR(100) NOT NULL,
    severity VARCHAR(20) NOT NULL,
    message TEXT NOT NULL,
    suggestion TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

**事件存储模型**
```sql
-- 事件存储表
CREATE TABLE event_store (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    aggregate_id UUID NOT NULL,
    aggregate_type VARCHAR(50) NOT NULL,
    event_type VARCHAR(100) NOT NULL,
    event_data JSONB NOT NULL,
    event_version INTEGER NOT NULL,
    occurred_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_aggregate (aggregate_id, aggregate_type),
    INDEX idx_event_type (event_type),
    INDEX idx_occurred_at (occurred_at)
);
```

### 3.5 安全架构设计

#### 3.5.1 安全层次模型

```
┌─────────────────────────────────────────────────────────────────┐
│                        应用安全层                               │
├─────────────────────────────────────────────────────────────────┤
│ • 输入验证  • 输出编码  • 会话管理  • 错误处理                   │
└─────────────────────────────────────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────────┐
│                        认证授权层                               │
├─────────────────────────────────────────────────────────────────┤
│ • JWT令牌  • RBAC权限  • OAuth2.0  • 多因素认证                  │
└─────────────────────────────────────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────────┐
│                        传输安全层                               │
├─────────────────────────────────────────────────────────────────┤
│ • TLS加密  • 证书管理  • 密钥轮换  • 安全头部                    │
└─────────────────────────────────────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────────┐
│                        基础设施安全层                           │
├─────────────────────────────────────────────────────────────────┤
│ • 网络隔离  • 防火墙  • 入侵检测  • 安全审计                    │
└─────────────────────────────────────────────────────────────────┘
```

#### 3.5.2 身份认证和授权

**JWT令牌设计**
```json
{
  "header": {
    "alg": "RS256",
    "typ": "JWT",
    "kid": "key-id-1"
  },
  "payload": {
    "sub": "user-uuid",
    "iss": "ai-scanner",
    "aud": "api-gateway",
    "exp": 1640995200,
    "iat": 1640908800,
    "roles": ["developer", "scanner"],
    "permissions": [
      "project:read",
      "scan:execute",
      "report:view"
    ]
  }
}
```

**RBAC权限模型**
```yaml
# 角色定义
roles:
  admin:
    description: "系统管理员"
    permissions:
      - "*:*"  # 所有权限
  
  security_engineer:
    description: "安全工程师"
    permissions:
      - "project:*"
      - "scan:*"
      - "report:*"
      - "rule:*"
  
  developer:
    description: "开发工程师"
    permissions:
      - "project:read"
      - "scan:execute"
      - "report:view"
      - "issue:update"
  
  viewer:
    description: "只读用户"
    permissions:
      - "project:read"
      - "report:view"

# 权限定义
permissions:
  project:
    - create
    - read
    - update
    - delete
  scan:
    - execute
    - cancel
    - schedule
  report:
    - view
    - export
    - share
  rule:
    - create
    - update
    - delete
```

## 4. 详细设计

### 4.1 用户服务设计

#### 4.1.1 服务架构

```go
// 用户服务架构
package user

import (
    "context"
    "time"
)

// 用户聚合根
type User struct {
    ID          string
    Username    string
    Email       string
    PasswordHash string
    Role        Role
    Profile     UserProfile
    CreatedAt   time.Time
    UpdatedAt   time.Time
    Version     int
}

// 用户角色
type Role struct {
    Name        string
    Permissions []Permission
}

// 用户配置
type UserProfile struct {
    DisplayName string
    Avatar      string
    Timezone    string
    Language    string
    Preferences map[string]interface{}
}

// 用户仓储接口
type UserRepository interface {
    Save(ctx context.Context, user *User) error
    FindByID(ctx context.Context, id string) (*User, error)
    FindByUsername(ctx context.Context, username string) (*User, error)
    FindByEmail(ctx context.Context, email string) (*User, error)
    Delete(ctx context.Context, id string) error
}

// 用户服务接口
type UserService interface {
    Register(ctx context.Context, req RegisterRequest) (*User, error)
    Authenticate(ctx context.Context, username, password string) (*User, error)
    UpdateProfile(ctx context.Context, userID string, profile UserProfile) error
    ChangePassword(ctx context.Context, userID, oldPassword, newPassword string) error
    GrantRole(ctx context.Context, userID string, role Role) error
}
```

#### 4.1.2 认证流程设计

```go
// 认证服务实现
type AuthService struct {
    userRepo     UserRepository
    tokenManager TokenManager
    passwordHash PasswordHasher
    logger       Logger
}

// 用户登录
func (s *AuthService) Login(ctx context.Context, req LoginRequest) (*LoginResponse, error) {
    // 1. 验证输入参数
    if err := req.Validate(); err != nil {
        return nil, NewValidationError(err)
    }
    
    // 2. 查找用户
    user, err := s.userRepo.FindByUsername(ctx, req.Username)
    if err != nil {
        s.logger.Error("failed to find user", "username", req.Username, "error", err)
        return nil, NewAuthenticationError("invalid credentials")
    }
    
    // 3. 验证密码
    if !s.passwordHash.Verify(req.Password, user.PasswordHash) {
        s.logger.Warn("password verification failed", "username", req.Username)
        return nil, NewAuthenticationError("invalid credentials")
    }
    
    // 4. 生成访问令牌
    accessToken, err := s.tokenManager.GenerateAccessToken(user)
    if err != nil {
        s.logger.Error("failed to generate access token", "userID", user.ID, "error", err)
        return nil, NewInternalError("token generation failed")
    }
    
    // 5. 生成刷新令牌
    refreshToken, err := s.tokenManager.GenerateRefreshToken(user)
    if err != nil {
        s.logger.Error("failed to generate refresh token", "userID", user.ID, "error", err)
        return nil, NewInternalError("token generation failed")
    }
    
    // 6. 记录登录事件
    s.publishEvent(UserLoggedInEvent{
        UserID:    user.ID,
        Username:  user.Username,
        Timestamp: time.Now(),
        IPAddress: req.IPAddress,
        UserAgent: req.UserAgent,
    })
    
    return &LoginResponse{
        AccessToken:  accessToken,
        RefreshToken: refreshToken,
        ExpiresIn:    s.tokenManager.AccessTokenTTL(),
        User:         user.ToDTO(),
    }, nil
}
```

### 4.2 项目服务设计

#### 4.2.1 项目管理架构

```go
// 项目聚合根
type Project struct {
    ID            string
    Name          string
    Description   string
    Repository    Repository
    Configuration ScanConfiguration
    Owner         string
    Members       []ProjectMember
    Status        ProjectStatus
    CreatedAt     time.Time
    UpdatedAt     time.Time
    Version       int
}

// 代码仓库
type Repository struct {
    URL         string
    Type        RepositoryType // git, svn, etc.
    Branch      string
    Credentials Credentials
    LastSync    time.Time
}

// 扫描配置
type ScanConfiguration struct {
    Languages     []string
    Rules         []RuleSet
    Exclusions    []string
    Schedule      ScanSchedule
    Notifications NotificationSettings
}

// 项目服务接口
type ProjectService interface {
    CreateProject(ctx context.Context, req CreateProjectRequest) (*Project, error)
    UpdateProject(ctx context.Context, projectID string, req UpdateProjectRequest) error
    DeleteProject(ctx context.Context, projectID string) error
    SyncRepository(ctx context.Context, projectID string) error
    GetProjectMembers(ctx context.Context, projectID string) ([]ProjectMember, error)
    AddMember(ctx context.Context, projectID, userID string, role MemberRole) error
}
```

#### 4.2.2 仓库同步机制

```go
// 仓库同步服务
type RepositorySyncService struct {
    gitClient    GitClient
    fileStorage  FileStorage
    eventBus     EventBus
    logger       Logger
}

// 同步仓库代码
func (s *RepositorySyncService) SyncRepository(ctx context.Context, project *Project) error {
    // 1. 检查仓库访问权限
    if err := s.gitClient.TestConnection(project.Repository); err != nil {
        return NewRepositoryError("connection failed", err)
    }
    
    // 2. 获取最新提交信息
    latestCommit, err := s.gitClient.GetLatestCommit(project.Repository)
    if err != nil {
        return NewRepositoryError("failed to get latest commit", err)
    }
    
    // 3. 检查是否需要同步
    if project.Repository.LastCommit == latestCommit.Hash {
        s.logger.Info("repository is up to date", "projectID", project.ID)
        return nil
    }
    
    // 4. 克隆或拉取代码
    localPath, err := s.gitClient.CloneOrPull(project.Repository)
    if err != nil {
        return NewRepositoryError("failed to clone repository", err)
    }
    
    // 5. 分析代码变更
    changes, err := s.analyzeChanges(project.Repository.LastCommit, latestCommit.Hash, localPath)
    if err != nil {
        return NewRepositoryError("failed to analyze changes", err)
    }
    
    // 6. 存储代码文件
    if err := s.storeCodeFiles(project.ID, localPath); err != nil {
        return NewStorageError("failed to store code files", err)
    }
    
    // 7. 发布同步完成事件
    s.eventBus.Publish(RepositorySyncedEvent{
        ProjectID:    project.ID,
        CommitHash:   latestCommit.Hash,
        Changes:      changes,
        SyncedAt:     time.Now(),
    })
    
    return nil
}
```

### 4.3 扫描服务设计

#### 4.3.1 扫描引擎架构

```go
// 扫描引擎接口
type ScanEngine interface {
    Scan(ctx context.Context, req ScanRequest) (*ScanResult, error)
    GetSupportedLanguages() []string
    GetAvailableRules() []Rule
    ValidateConfiguration(config ScanConfiguration) error
}

// 扫描请求
type ScanRequest struct {
    ProjectID     string
    ScanType      ScanType // full, incremental, targeted
    Configuration ScanConfiguration
    Files         []string
    Timeout       time.Duration
}

// 扫描结果
type ScanResult struct {
    TaskID      string
    ProjectID   string
    Status      ScanStatus
    Issues      []Issue
    Metrics     ScanMetrics
    StartedAt   time.Time
    CompletedAt time.Time
    Duration    time.Duration
}

// 问题定义
type Issue struct {
    ID          string
    RuleID      string
    Severity    Severity
    Category    Category
    File        string
    Line        int
    Column      int
    Message     string
    Description string
    Suggestion  string
    Confidence  float64
    Context     IssueContext
}
```

#### 4.3.2 多语言解析器

```go
// 代码解析器接口
type CodeParser interface {
    Parse(ctx context.Context, file string) (*AST, error)
    GetLanguage() string
    GetFileExtensions() []string
}

// Java解析器实现
type JavaParser struct {
    treeSitter *TreeSitter
    logger     Logger
}

func (p *JavaParser) Parse(ctx context.Context, file string) (*AST, error) {
    // 1. 读取文件内容
    content, err := ioutil.ReadFile(file)
    if err != nil {
        return nil, NewParseError("failed to read file", err)
    }
    
    // 2. 使用Tree-sitter解析
    tree, err := p.treeSitter.Parse(content, "java")
    if err != nil {
        return nil, NewParseError("failed to parse Java code", err)
    }
    
    // 3. 构建AST
    ast := &AST{
        Language: "java",
        File:     file,
        Root:     p.buildASTNode(tree.RootNode()),
        Imports:  p.extractImports(tree),
        Classes:  p.extractClasses(tree),
        Methods:  p.extractMethods(tree),
    }
    
    return ast, nil
}

// 提取类信息
func (p *JavaParser) extractClasses(tree *Tree) []ClassNode {
    var classes []ClassNode
    
    query := `
        (class_declaration
            name: (identifier) @class_name
            superclass: (superclass (type_identifier) @superclass)?
            interfaces: (super_interfaces (interface_type_list (type_identifier) @interface))?
            body: (class_body) @class_body
        ) @class
    `
    
    matches := p.treeSitter.Query(tree, query)
    for _, match := range matches {
        class := ClassNode{
            Name:       match.CaptureByName("class_name").Text(),
            Superclass: match.CaptureByName("superclass").Text(),
            Interfaces: p.extractInterfaces(match.CaptureByName("interface")),
            Methods:    p.extractClassMethods(match.CaptureByName("class_body")),
            Fields:     p.extractFields(match.CaptureByName("class_body")),
            Position:   match.CaptureByName("class").Position(),
        }
        classes = append(classes, class)
    }
    
    return classes
}
```

### 4.4 AI智能体服务设计

#### 4.4.1 智能体协调器

```go
// 智能体协调器
type AgentCoordinator struct {
    agents      map[string]Agent
    taskQueue   TaskQueue
    resultStore ResultStore
    logger      Logger
}

// 智能体接口
type Agent interface {
    GetCapabilities() []Capability
    Process(ctx context.Context, task Task) (*Result, error)
    GetStatus() AgentStatus
}

// 代码理解智能体
type CodeUnderstandingAgent struct {
    model      LanguageModel
    knowledge  KnowledgeBase
    cache      Cache
    logger     Logger
}

func (a *CodeUnderstandingAgent) Process(ctx context.Context, task Task) (*Result, error) {
    codeTask, ok := task.(*CodeAnalysisTask)
    if !ok {
        return nil, NewTaskError("invalid task type")
    }
    
    // 1. 预处理代码
    preprocessed, err := a.preprocessCode(codeTask.Code)
    if err != nil {
        return nil, NewProcessingError("code preprocessing failed", err)
    }
    
    // 2. 提取代码特征
    features, err := a.extractFeatures(preprocessed)
    if err != nil {
        return nil, NewProcessingError("feature extraction failed", err)
    }
    
    // 3. 模型推理
    prediction, err := a.model.Predict(ctx, features)
    if err != nil {
        return nil, NewModelError("model prediction failed", err)
    }
    
    // 4. 后处理结果
    result, err := a.postprocessResult(prediction, codeTask)
    if err != nil {
        return nil, NewProcessingError("result postprocessing failed", err)
    }
    
    return result, nil
}

// 安全检测智能体
type SecurityDetectionAgent struct {
    models     map[string]SecurityModel
    rules      RuleEngine
    knowledge  SecurityKnowledge
    logger     Logger
}

func (a *SecurityDetectionAgent) Process(ctx context.Context, task Task) (*Result, error) {
    securityTask, ok := task.(*SecurityAnalysisTask)
    if !ok {
        return nil, NewTaskError("invalid task type")
    }
    
    var allIssues []SecurityIssue
    
    // 1. 规则引擎检测
    ruleIssues, err := a.rules.Analyze(securityTask.Code)
    if err != nil {
        a.logger.Error("rule engine analysis failed", "error", err)
    } else {
        allIssues = append(allIssues, ruleIssues...)
    }
    
    // 2. AI模型检测
    for modelName, model := range a.models {
        modelIssues, err := model.DetectVulnerabilities(ctx, securityTask.Code)
        if err != nil {
            a.logger.Error("model detection failed", "model", modelName, "error", err)
            continue
        }
        allIssues = append(allIssues, modelIssues...)
    }
    
    // 3. 去重和合并
    dedupedIssues := a.deduplicateIssues(allIssues)
    
    // 4. 风险评估
    for i := range dedupedIssues {
        dedupedIssues[i].Risk = a.assessRisk(&dedupedIssues[i])
    }
    
    return &SecurityAnalysisResult{
        Issues:    dedupedIssues,
        Summary:   a.generateSummary(dedupedIssues),
        Timestamp: time.Now(),
    }, nil
}
```

#### 4.4.2 AI模型管理

```go
// 模型管理器
type ModelManager struct {
    models    map[string]Model
    registry  ModelRegistry
    loader    ModelLoader
    monitor   ModelMonitor
    logger    Logger
}

// 模型接口
type Model interface {
    Load(ctx context.Context, path string) error
    Predict(ctx context.Context, input interface{}) (interface{}, error)
    GetMetadata() ModelMetadata
    GetMetrics() ModelMetrics
    Unload() error
}

// 模型元数据
type ModelMetadata struct {
    Name        string
    Version     string
    Type        ModelType
    Language    string
    Framework   string
    InputShape  []int
    OutputShape []int
    CreatedAt   time.Time
    UpdatedAt   time.Time
}

// 模型加载器
type ModelLoader struct {
    storage ModelStorage
    cache   ModelCache
    logger  Logger
}

func (l *ModelLoader) LoadModel(ctx context.Context, name, version string) (Model, error) {
    // 1. 检查缓存
    if model, exists := l.cache.Get(name, version); exists {
        return model, nil
    }
    
    // 2. 从存储加载
    modelPath, err := l.storage.GetModelPath(name, version)
    if err != nil {
        return nil, NewModelError("failed to get model path", err)
    }
    
    // 3. 创建模型实例
    model, err := l.createModelInstance(modelPath)
    if err != nil {
        return nil, NewModelError("failed to create model instance", err)
    }
    
    // 4. 加载模型
    if err := model.Load(ctx, modelPath); err != nil {
        return nil, NewModelError("failed to load model", err)
    }
    
    // 5. 缓存模型
    l.cache.Put(name, version, model)
    
    return model, nil
}
```

## 5. 用户界面设计

### 5.1 前端架构设计

#### 5.1.1 技术栈选择

```typescript
// 前端技术栈
interface TechStack {
  framework: 'React 18+';           // 主框架
  language: 'TypeScript 4.5+';     // 开发语言
  stateManagement: 'Redux Toolkit'; // 状态管理
  routing: 'React Router v6';       // 路由管理
  uiLibrary: 'Ant Design';          // UI组件库
  styling: 'Styled Components';     // 样式方案
  bundler: 'Vite';                  // 构建工具
  testing: 'Jest + RTL';            // 测试框架
}
```

#### 5.1.2 组件架构

```typescript
// 组件层次结构
interface ComponentHierarchy {
  app: {
    layout: {
      header: 'AppHeader';
      sidebar: 'AppSidebar';
      content: 'AppContent';
      footer: 'AppFooter';
    };
    pages: {
      dashboard: 'DashboardPage';
      projects: 'ProjectsPage';
      scans: 'ScansPage';
      reports: 'ReportsPage';
      settings: 'SettingsPage';
    };
    components: {
      common: 'Button | Input | Modal | Table';
      business: 'ProjectCard | ScanResult | IssueList';
    };
  };
}

// React组件示例
import React, { useState, useEffect } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { Card, Button, Table, Tag, Space } from 'antd';
import { PlayCircleOutlined, PauseCircleOutlined } from '@ant-design/icons';

interface ScanResultsProps {
  projectId: string;
  scanId: string;
}

const ScanResults: React.FC<ScanResultsProps> = ({ projectId, scanId }) => {
  const dispatch = useDispatch();
  const { issues, loading, pagination } = useSelector(state => state.scans);
  
  const [selectedIssues, setSelectedIssues] = useState<string[]>([]);
  
  useEffect(() => {
    dispatch(fetchScanResults({ projectId, scanId }));
  }, [dispatch, projectId, scanId]);
  
  const columns = [
    {
      title: '文件',
      dataIndex: 'file',
      key: 'file',
      render: (file: string) => (
        <span className="file-path">{file}</span>
      ),
    },
    {
      title: '行号',
      dataIndex: 'line',
      key: 'line',
      width: 80,
    },
    {
      title: '严重程度',
      dataIndex: 'severity',
      key: 'severity',
      render: (severity: string) => (
        <Tag color={getSeverityColor(severity)}>
          {severity.toUpperCase()}
        </Tag>
      ),
    },
    {
      title: '问题描述',
      dataIndex: 'message',
      key: 'message',
      ellipsis: true,
    },
    {
      title: '操作',
      key: 'actions',
      render: (_, record) => (
        <Space>
          <Button size="small" onClick={() => viewIssueDetail(record.id)}>
            查看详情
          </Button>
          <Button size="small" onClick={() => markAsFixed(record.id)}>
            标记已修复
          </Button>
        </Space>
      ),
    },
  ];
  
  return (
    <Card title="扫描结果" extra={<ScanActions scanId={scanId} />}>
      <Table
        columns={columns}
        dataSource={issues}
        loading={loading}
        pagination={pagination}
        rowSelection={{
          selectedRowKeys: selectedIssues,
          onChange: setSelectedIssues,
        }}
        rowKey="id"
      />
    </Card>
  );
};

export default ScanResults;
```

### 5.2 状态管理设计

#### 5.2.1 Redux Store结构

```typescript
// 全局状态结构
interface RootState {
  auth: AuthState;
  projects: ProjectsState;
  scans: ScansState;
  reports: ReportsState;
  ui: UIState;
}

// 认证状态
interface AuthState {
  user: User | null;
  token: string | null;
  isAuthenticated: boolean;
  loading: boolean;
  error: string | null;
}

// 项目状态
interface ProjectsState {
  items: Project[];
  currentProject: Project | null;
  loading: boolean;
  error: string | null;
  pagination: PaginationState;
  filters: ProjectFilters;
}

// 扫描状态
interface ScansState {
  tasks: ScanTask[];
  currentTask: ScanTask | null;
  results: ScanResult[];
  issues: Issue[];
  loading: boolean;
  error: string | null;
  realTimeUpdates: boolean;
}

// Redux Toolkit Slice示例
import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import { scanAPI } from '../api/scanAPI';

// 异步Action
export const startScan = createAsyncThunk(
  'scans/startScan',
  async (params: StartScanParams, { rejectWithValue }) => {
    try {
      const response = await scanAPI.startScan(params);
      return response.data;
    } catch (error) {
      return rejectWithValue(error.response.data);
    }
  }
);

export const fetchScanResults = createAsyncThunk(
  'scans/fetchResults',
  async (params: FetchResultsParams) => {
    const response = await scanAPI.getResults(params);
    return response.data;
  }
);

// Slice定义
const scansSlice = createSlice({
  name: 'scans',
  initialState: {
    tasks: [],
    currentTask: null,
    results: [],
    issues: [],
    loading: false,
    error: null,
    realTimeUpdates: false,
  } as ScansState,
  reducers: {
    setCurrentTask: (state, action) => {
      state.currentTask = action.payload;
    },
    updateTaskStatus: (state, action) => {
      const { taskId, status } = action.payload;
      const task = state.tasks.find(t => t.id === taskId);
      if (task) {
        task.status = status;
      }
    },
    addIssue: (state, action) => {
      state.issues.push(action.payload);
    },
    markIssueFixed: (state, action) => {
      const issue = state.issues.find(i => i.id === action.payload);
      if (issue) {
        issue.status = 'fixed';
      }
    },
  },
  extraReducers: (builder) => {
    builder
      .addCase(startScan.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(startScan.fulfilled, (state, action) => {
        state.loading = false;
        state.tasks.push(action.payload);
        state.currentTask = action.payload;
      })
      .addCase(startScan.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload as string;
      })
      .addCase(fetchScanResults.fulfilled, (state, action) => {
        state.results = action.payload.results;
        state.issues = action.payload.issues;
      });
  },
});

export const { setCurrentTask, updateTaskStatus, addIssue, markIssueFixed } = scansSlice.actions;
export default scansSlice.reducer;
```

### 5.3 实时通信设计

#### 5.3.1 WebSocket集成

```typescript
// WebSocket管理器
class WebSocketManager {
  private ws: WebSocket | null = null;
  private reconnectAttempts = 0;
  private maxReconnectAttempts = 5;
  private reconnectInterval = 1000;
  private eventHandlers: Map<string, Function[]> = new Map();
  
  connect(url: string, token: string): Promise<void> {
    return new Promise((resolve, reject) => {
      try {
        this.ws = new WebSocket(`${url}?token=${token}`);
        
        this.ws.onopen = () => {
          console.log('WebSocket connected');
          this.reconnectAttempts = 0;
          resolve();
        };
        
        this.ws.onmessage = (event) => {
          const message = JSON.parse(event.data);
          this.handleMessage(message);
        };
        
        this.ws.onclose = () => {
          console.log('WebSocket disconnected');
          this.handleReconnect();
        };
        
        this.ws.onerror = (error) => {
          console.error('WebSocket error:', error);
          reject(error);
        };
      } catch (error) {
        reject(error);
      }
    });
  }
  
  private handleMessage(message: WebSocketMessage): void {
    const { type, data } = message;
    const handlers = this.eventHandlers.get(type) || [];
    
    handlers.forEach(handler => {
      try {
        handler(data);
      } catch (error) {
        console.error('Error handling WebSocket message:', error);
      }
    });
  }
  
  private handleReconnect(): void {
    if (this.reconnectAttempts < this.maxReconnectAttempts) {
      this.reconnectAttempts++;
      setTimeout(() => {
        console.log(`Attempting to reconnect (${this.reconnectAttempts}/${this.maxReconnectAttempts})`);
        this.connect(this.url, this.token);
      }, this.reconnectInterval * this.reconnectAttempts);
    }
  }
  
  subscribe(eventType: string, handler: Function): void {
    if (!this.eventHandlers.has(eventType)) {
      this.eventHandlers.set(eventType, []);
    }
    this.eventHandlers.get(eventType)!.push(handler);
  }
  
  unsubscribe(eventType: string, handler: Function): void {
    const handlers = this.eventHandlers.get(eventType);
    if (handlers) {
      const index = handlers.indexOf(handler);
      if (index > -1) {
        handlers.splice(index, 1);
      }
    }
  }
  
  send(message: WebSocketMessage): void {
    if (this.ws && this.ws.readyState === WebSocket.OPEN) {
      this.ws.send(JSON.stringify(message));
    }
  }
  
  disconnect(): void {
    if (this.ws) {
      this.ws.close();
      this.ws = null;
    }
  }
}

// React Hook for WebSocket
export const useWebSocket = () => {
  const dispatch = useDispatch();
  const wsManager = useRef<WebSocketManager>(new WebSocketManager());
  
  useEffect(() => {
    const token = localStorage.getItem('token');
    if (token) {
      wsManager.current.connect('ws://localhost:8080/ws', token);
      
      // 订阅扫描状态更新
      wsManager.current.subscribe('scan_status_update', (data) => {
        dispatch(updateTaskStatus(data));
      });
      
      // 订阅新问题发现
      wsManager.current.subscribe('issue_found', (data) => {
        dispatch(addIssue(data));
      });
      
      // 订阅扫描完成
      wsManager.current.subscribe('scan_completed', (data) => {
        dispatch(fetchScanResults({ scanId: data.scanId }));
      });
    }
    
    return () => {
      wsManager.current.disconnect();
    };
  }, [dispatch]);
  
  return wsManager.current;
};
```

---

**文档版本历史**

| 版本 | 日期 | 修改内容 | 修改人 |
|------|------|----------|--------|
| v1.0 | 2024-12-19 | 初始版本 | 系统架构师 |

**审阅记录**

| 审阅者 | 审阅日期 | 审阅意见 | 状态 |
|--------|----------|----------|------|
| 技术总监 | 待定 | 待审阅 | 待审阅 |
| 高级工程师 | 待定 | 待审阅 | 待审阅 |
| 技术委员会 | 待定 | 待审阅 | 待审阅 |