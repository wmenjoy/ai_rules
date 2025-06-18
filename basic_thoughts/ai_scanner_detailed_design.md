# AI代码安全扫描系统详细设计文档

## 文档信息

| 项目 | 内容 |
|------|------|
| 文档标题 | AI代码安全扫描系统详细设计文档 |
| 文档版本 | v1.0 |
| 创建日期 | 2024-12-19 |
| 最后更新 | 2024-12-19 |
| 文档状态 | 草案 |
| 作者 | 系统架构师 |
| 审阅者 | 开发团队负责人 |
| 批准者 | 技术委员会 |

## 1. 引言

### 1.1 目的
本文档详细描述AI代码安全扫描系统中各个组件、类和方法的具体实现设计，为开发人员提供详细的实现指导。

### 1.2 范围
本文档涵盖系统中所有核心组件的详细设计，包括：
- 类结构和接口定义
- 方法实现逻辑
- 数据结构设计
- 算法实现
- 异常处理机制
- 性能优化策略

### 1.3 参考文献
- AI代码安全扫描系统需求规格说明书 v1.0
- AI代码安全扫描系统软件设计文档 v1.0
- AI代码安全扫描系统模块设计文档 v1.0
- Go语言编程规范
- 微服务设计模式

## 2. 用户服务详细设计

### 2.1 用户管理服务

#### 2.1.1 UserService 类设计

```go
// UserService 用户管理服务实现
type UserService struct {
    userRepo        UserRepository
    roleRepo        RoleRepository
    passwordHasher  PasswordHasher
    tokenManager    TokenManager
    validator       Validator
    eventBus        EventBus
    logger          Logger
    config          UserConfig
    metrics         UserMetrics
}

// UserConfig 用户服务配置
type UserConfig struct {
    PasswordPolicy    PasswordPolicy    `yaml:"password_policy"`
    TokenConfig       TokenConfig       `yaml:"token_config"`
    LoginAttempts     LoginAttempts     `yaml:"login_attempts"`
    SessionConfig     SessionConfig     `yaml:"session_config"`
}

type PasswordPolicy struct {
    MinLength        int    `yaml:"min_length"`
    RequireUppercase bool   `yaml:"require_uppercase"`
    RequireLowercase bool   `yaml:"require_lowercase"`
    RequireNumbers   bool   `yaml:"require_numbers"`
    RequireSymbols   bool   `yaml:"require_symbols"`
    MaxAge           int    `yaml:"max_age_days"`
}

type TokenConfig struct {
    AccessTokenTTL  time.Duration `yaml:"access_token_ttl"`
    RefreshTokenTTL time.Duration `yaml:"refresh_token_ttl"`
    SigningKey      string        `yaml:"signing_key"`
    Issuer          string        `yaml:"issuer"`
}

type LoginAttempts struct {
    MaxAttempts     int           `yaml:"max_attempts"`
    LockoutDuration time.Duration `yaml:"lockout_duration"`
    ResetWindow     time.Duration `yaml:"reset_window"`
}
```

#### 2.1.2 用户注册实现

```go
// Register 用户注册
func (s *UserService) Register(ctx context.Context, req RegisterRequest) (*User, error) {
    // 记录方法调用
    s.logger.Info("user registration started", "username", req.Username)
    s.metrics.IncRegisterAttempts()
    
    // 开始计时
    start := time.Now()
    defer func() {
        s.metrics.RecordRegisterDuration(time.Since(start))
    }()
    
    // 1. 输入验证
    if err := s.validateRegisterRequest(req); err != nil {
        s.logger.Warn("registration validation failed", "error", err)
        s.metrics.IncRegisterFailures("validation")
        return nil, NewValidationError("invalid registration request", err)
    }
    
    // 2. 检查用户名唯一性
    if exists, err := s.checkUsernameExists(ctx, req.Username); err != nil {
        s.logger.Error("failed to check username existence", "error", err)
        s.metrics.IncRegisterFailures("database")
        return nil, NewInternalError("database error", err)
    } else if exists {
        s.logger.Warn("username already exists", "username", req.Username)
        s.metrics.IncRegisterFailures("conflict")
        return nil, NewConflictError("username already exists")
    }
    
    // 3. 检查邮箱唯一性
    if exists, err := s.checkEmailExists(ctx, req.Email); err != nil {
        s.logger.Error("failed to check email existence", "error", err)
        s.metrics.IncRegisterFailures("database")
        return nil, NewInternalError("database error", err)
    } else if exists {
        s.logger.Warn("email already exists", "email", req.Email)
        s.metrics.IncRegisterFailures("conflict")
        return nil, NewConflictError("email already exists")
    }
    
    // 4. 密码强度验证
    if err := s.validatePasswordStrength(req.Password); err != nil {
        s.logger.Warn("password strength validation failed", "error", err)
        s.metrics.IncRegisterFailures("password_policy")
        return nil, NewValidationError("password does not meet policy requirements", err)
    }
    
    // 5. 密码哈希
    passwordHash, err := s.passwordHasher.Hash(req.Password)
    if err != nil {
        s.logger.Error("password hashing failed", "error", err)
        s.metrics.IncRegisterFailures("hashing")
        return nil, NewInternalError("password processing failed", err)
    }
    
    // 6. 创建用户对象
    user := &User{
        ID:           s.generateUserID(),
        Username:     req.Username,
        Email:        req.Email,
        PasswordHash: passwordHash,
        DisplayName:  req.DisplayName,
        Status:       UserStatusActive,
        Roles:        []Role{s.getDefaultRole()},
        Preferences:  s.getDefaultPreferences(),
        CreatedAt:    time.Now(),
        UpdatedAt:    time.Now(),
        Version:      1,
    }
    
    // 7. 保存用户到数据库
    if err := s.userRepo.Save(ctx, user); err != nil {
        s.logger.Error("failed to save user", "error", err)
        s.metrics.IncRegisterFailures("database")
        return nil, NewInternalError("failed to create user", err)
    }
    
    // 8. 发布用户注册事件
    event := UserRegisteredEvent{
        UserID:    user.ID,
        Username:  user.Username,
        Email:     user.Email,
        Timestamp: time.Now(),
    }
    
    if err := s.eventBus.Publish(ctx, "user.registered", event); err != nil {
        s.logger.Error("failed to publish user registered event", "error", err)
        // 不返回错误，因为用户已经创建成功
    }
    
    // 9. 记录成功指标
    s.metrics.IncRegisterSuccesses()
    s.logger.Info("user registration completed", "userID", user.ID, "username", user.Username)
    
    // 10. 返回用户信息（不包含敏感数据）
    return s.sanitizeUser(user), nil
}

// validateRegisterRequest 验证注册请求
func (s *UserService) validateRegisterRequest(req RegisterRequest) error {
    var errors []string
    
    // 用户名验证
    if len(req.Username) < 3 || len(req.Username) > 50 {
        errors = append(errors, "username must be between 3 and 50 characters")
    }
    
    if !s.isValidUsername(req.Username) {
        errors = append(errors, "username contains invalid characters")
    }
    
    // 邮箱验证
    if !s.isValidEmail(req.Email) {
        errors = append(errors, "invalid email format")
    }
    
    // 密码验证
    if len(req.Password) < s.config.PasswordPolicy.MinLength {
        errors = append(errors, fmt.Sprintf("password must be at least %d characters", s.config.PasswordPolicy.MinLength))
    }
    
    // 显示名称验证
    if len(req.DisplayName) > 100 {
        errors = append(errors, "display name must not exceed 100 characters")
    }
    
    if len(errors) > 0 {
        return fmt.Errorf("validation errors: %s", strings.Join(errors, ", "))
    }
    
    return nil
}

// validatePasswordStrength 验证密码强度
func (s *UserService) validatePasswordStrength(password string) error {
    policy := s.config.PasswordPolicy
    var errors []string
    
    if policy.RequireUppercase && !regexp.MustMatch(`[A-Z]`, password) {
        errors = append(errors, "password must contain at least one uppercase letter")
    }
    
    if policy.RequireLowercase && !regexp.MustMatch(`[a-z]`, password) {
        errors = append(errors, "password must contain at least one lowercase letter")
    }
    
    if policy.RequireNumbers && !regexp.MustMatch(`[0-9]`, password) {
        errors = append(errors, "password must contain at least one number")
    }
    
    if policy.RequireSymbols && !regexp.MustMatch(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`, password) {
        errors = append(errors, "password must contain at least one special character")
    }
    
    // 检查常见弱密码
    if s.isCommonPassword(password) {
        errors = append(errors, "password is too common")
    }
    
    if len(errors) > 0 {
        return fmt.Errorf("password policy violations: %s", strings.Join(errors, ", "))
    }
    
    return nil
}
```

#### 2.1.3 用户认证实现

```go
// Login 用户登录
func (s *UserService) Login(ctx context.Context, req LoginRequest) (*LoginResponse, error) {
    s.logger.Info("user login attempt", "username", req.Username)
    s.metrics.IncLoginAttempts()
    
    start := time.Now()
    defer func() {
        s.metrics.RecordLoginDuration(time.Since(start))
    }()
    
    // 1. 检查登录尝试次数
    if locked, err := s.checkAccountLockout(ctx, req.Username); err != nil {
        s.logger.Error("failed to check account lockout", "error", err)
        return nil, NewInternalError("authentication service unavailable", err)
    } else if locked {
        s.logger.Warn("account locked due to too many failed attempts", "username", req.Username)
        s.metrics.IncLoginFailures("locked")
        return nil, NewAuthenticationError("account temporarily locked")
    }
    
    // 2. 查找用户
    user, err := s.userRepo.FindByUsername(ctx, req.Username)
    if err != nil {
        if errors.Is(err, ErrUserNotFound) {
            s.logger.Warn("login attempt with non-existent username", "username", req.Username)
            s.recordFailedLoginAttempt(ctx, req.Username)
            s.metrics.IncLoginFailures("invalid_credentials")
            return nil, NewAuthenticationError("invalid credentials")
        }
        s.logger.Error("failed to find user", "error", err)
        s.metrics.IncLoginFailures("database")
        return nil, NewInternalError("authentication service error", err)
    }
    
    // 3. 检查用户状态
    if user.Status != UserStatusActive {
        s.logger.Warn("login attempt with inactive user", "username", req.Username, "status", user.Status)
        s.metrics.IncLoginFailures("inactive_user")
        return nil, NewAuthenticationError("account is not active")
    }
    
    // 4. 验证密码
    if !s.passwordHasher.Verify(req.Password, user.PasswordHash) {
        s.logger.Warn("login attempt with invalid password", "username", req.Username)
        s.recordFailedLoginAttempt(ctx, req.Username)
        s.metrics.IncLoginFailures("invalid_credentials")
        return nil, NewAuthenticationError("invalid credentials")
    }
    
    // 5. 检查密码是否需要更新
    passwordAge := time.Since(user.PasswordUpdatedAt)
    passwordExpired := s.config.PasswordPolicy.MaxAge > 0 && 
        passwordAge > time.Duration(s.config.PasswordPolicy.MaxAge)*24*time.Hour
    
    // 6. 生成访问令牌
    accessToken, err := s.tokenManager.GenerateAccessToken(user)
    if err != nil {
        s.logger.Error("failed to generate access token", "error", err)
        s.metrics.IncLoginFailures("token_generation")
        return nil, NewInternalError("token generation failed", err)
    }
    
    // 7. 生成刷新令牌
    refreshToken, err := s.tokenManager.GenerateRefreshToken(user)
    if err != nil {
        s.logger.Error("failed to generate refresh token", "error", err)
        s.metrics.IncLoginFailures("token_generation")
        return nil, NewInternalError("token generation failed", err)
    }
    
    // 8. 更新用户最后登录时间
    user.LastLoginAt = time.Now()
    if err := s.userRepo.Update(ctx, user); err != nil {
        s.logger.Error("failed to update user last login time", "error", err)
        // 不返回错误，因为登录已经成功
    }
    
    // 9. 清除失败登录记录
    s.clearFailedLoginAttempts(ctx, req.Username)
    
    // 10. 发布登录成功事件
    event := UserLoggedInEvent{
        UserID:    user.ID,
        Username:  user.Username,
        IPAddress: s.getClientIP(ctx),
        UserAgent: s.getUserAgent(ctx),
        Timestamp: time.Now(),
    }
    
    if err := s.eventBus.Publish(ctx, "user.logged_in", event); err != nil {
        s.logger.Error("failed to publish user logged in event", "error", err)
    }
    
    // 11. 记录成功指标
    s.metrics.IncLoginSuccesses()
    s.logger.Info("user login successful", "userID", user.ID, "username", user.Username)
    
    // 12. 构建响应
    response := &LoginResponse{
        AccessToken:     accessToken,
        RefreshToken:    refreshToken,
        TokenType:       "Bearer",
        ExpiresIn:       int(s.config.TokenConfig.AccessTokenTTL.Seconds()),
        User:            s.sanitizeUser(user),
        PasswordExpired: passwordExpired,
    }
    
    return response, nil
}

// recordFailedLoginAttempt 记录失败的登录尝试
func (s *UserService) recordFailedLoginAttempt(ctx context.Context, username string) {
    key := fmt.Sprintf("failed_login:%s", username)
    
    // 获取当前失败次数
    attempts, err := s.cache.Get(ctx, key)
    if err != nil {
        attempts = 0
    }
    
    // 增加失败次数
    newAttempts := attempts.(int) + 1
    
    // 设置缓存，使用重置窗口作为TTL
    if err := s.cache.Set(ctx, key, newAttempts, s.config.LoginAttempts.ResetWindow); err != nil {
        s.logger.Error("failed to record login attempt", "error", err)
    }
    
    // 如果达到最大尝试次数，设置锁定
    if newAttempts >= s.config.LoginAttempts.MaxAttempts {
        lockKey := fmt.Sprintf("account_locked:%s", username)
        if err := s.cache.Set(ctx, lockKey, true, s.config.LoginAttempts.LockoutDuration); err != nil {
            s.logger.Error("failed to set account lockout", "error", err)
        }
    }
}

// checkAccountLockout 检查账户是否被锁定
func (s *UserService) checkAccountLockout(ctx context.Context, username string) (bool, error) {
    key := fmt.Sprintf("account_locked:%s", username)
    
    locked, err := s.cache.Get(ctx, key)
    if err != nil {
        if errors.Is(err, ErrCacheKeyNotFound) {
            return false, nil
        }
        return false, err
    }
    
    return locked.(bool), nil
}
```

### 2.2 权限管理服务

#### 2.2.1 AuthService 类设计

```go
// AuthService 权限管理服务
type AuthService struct {
    userRepo     UserRepository
    roleRepo     RoleRepository
    permRepo     PermissionRepository
    tokenManager TokenManager
    cache        CacheManager
    logger       Logger
    config       AuthConfig
    metrics      AuthMetrics
}

// AuthConfig 权限服务配置
type AuthConfig struct {
    CacheConfig      CacheConfig      `yaml:"cache_config"`
    PermissionConfig PermissionConfig `yaml:"permission_config"`
    AuditConfig      AuditConfig      `yaml:"audit_config"`
}

type CacheConfig struct {
    PermissionCacheTTL time.Duration `yaml:"permission_cache_ttl"`
    RoleCacheTTL       time.Duration `yaml:"role_cache_ttl"`
    UserCacheTTL       time.Duration `yaml:"user_cache_ttl"`
}

type PermissionConfig struct {
    DefaultDenyAll     bool     `yaml:"default_deny_all"`
    SuperAdminRoles    []string `yaml:"super_admin_roles"`
    SystemPermissions  []string `yaml:"system_permissions"`
}

type AuditConfig struct {
    EnableAudit        bool     `yaml:"enable_audit"`
    AuditPermissions   []string `yaml:"audit_permissions"`
    AuditRetentionDays int      `yaml:"audit_retention_days"`
}
```

#### 2.2.2 权限检查实现

```go
// CheckPermission 检查用户权限
func (s *AuthService) CheckPermission(ctx context.Context, userID, resource, action string) (bool, error) {
    s.logger.Debug("checking permission", "userID", userID, "resource", resource, "action", action)
    s.metrics.IncPermissionChecks()
    
    start := time.Now()
    defer func() {
        s.metrics.RecordPermissionCheckDuration(time.Since(start))
    }()
    
    // 1. 参数验证
    if userID == "" || resource == "" || action == "" {
        s.logger.Warn("invalid permission check parameters")
        s.metrics.IncPermissionCheckFailures("invalid_params")
        return false, NewValidationError("invalid permission check parameters")
    }
    
    // 2. 检查缓存
    cacheKey := fmt.Sprintf("permission:%s:%s:%s", userID, resource, action)
    if cached, err := s.cache.Get(ctx, cacheKey); err == nil {
        s.metrics.IncPermissionCacheHits()
        return cached.(bool), nil
    }
    s.metrics.IncPermissionCacheMisses()
    
    // 3. 获取用户信息
    user, err := s.getUserWithRoles(ctx, userID)
    if err != nil {
        s.logger.Error("failed to get user", "userID", userID, "error", err)
        s.metrics.IncPermissionCheckFailures("user_not_found")
        return false, err
    }
    
    // 4. 检查用户状态
    if user.Status != UserStatusActive {
        s.logger.Warn("permission check for inactive user", "userID", userID, "status", user.Status)
        s.metrics.IncPermissionCheckFailures("inactive_user")
        return false, nil
    }
    
    // 5. 检查超级管理员权限
    if s.isSuperAdmin(user) {
        s.logger.Debug("super admin access granted", "userID", userID)
        s.cachePermissionResult(ctx, cacheKey, true)
        s.metrics.IncPermissionCheckSuccesses("super_admin")
        return true, nil
    }
    
    // 6. 检查角色权限
    hasPermission, err := s.checkRolePermissions(ctx, user.Roles, resource, action)
    if err != nil {
        s.logger.Error("failed to check role permissions", "error", err)
        s.metrics.IncPermissionCheckFailures("role_check_error")
        return false, err
    }
    
    // 7. 检查直接权限（用户级别权限）
    if !hasPermission {
        hasPermission, err = s.checkDirectPermissions(ctx, userID, resource, action)
        if err != nil {
            s.logger.Error("failed to check direct permissions", "error", err)
            s.metrics.IncPermissionCheckFailures("direct_check_error")
            return false, err
        }
    }
    
    // 8. 缓存结果
    s.cachePermissionResult(ctx, cacheKey, hasPermission)
    
    // 9. 记录审计日志
    if s.config.AuditConfig.EnableAudit && s.shouldAuditPermission(resource, action) {
        s.auditPermissionCheck(ctx, userID, resource, action, hasPermission)
    }
    
    // 10. 记录指标
    if hasPermission {
        s.metrics.IncPermissionCheckSuccesses("granted")
    } else {
        s.metrics.IncPermissionCheckFailures("denied")
    }
    
    s.logger.Debug("permission check completed", "userID", userID, "resource", resource, "action", action, "granted", hasPermission)
    
    return hasPermission, nil
}

// checkRolePermissions 检查角色权限
func (s *AuthService) checkRolePermissions(ctx context.Context, roles []Role, resource, action string) (bool, error) {
    for _, role := range roles {
        // 获取角色权限
        permissions, err := s.getRolePermissions(ctx, role.ID)
        if err != nil {
            s.logger.Error("failed to get role permissions", "roleID", role.ID, "error", err)
            continue
        }
        
        // 检查权限匹配
        for _, permission := range permissions {
            if s.matchPermission(permission, resource, action) {
                if permission.Effect == PermissionEffectAllow {
                    return true, nil
                } else if permission.Effect == PermissionEffectDeny {
                    // 显式拒绝优先级最高
                    return false, nil
                }
            }
        }
    }
    
    return false, nil
}

// matchPermission 匹配权限
func (s *AuthService) matchPermission(permission Permission, resource, action string) bool {
    // 精确匹配
    if permission.Resource == resource && permission.Action == action {
        return true
    }
    
    // 通配符匹配
    if permission.Resource == "*" || permission.Action == "*" {
        return true
    }
    
    // 资源层级匹配
    if s.matchResourceHierarchy(permission.Resource, resource) {
        if permission.Action == action || permission.Action == "*" {
            return true
        }
    }
    
    // 动作组匹配
    if s.matchActionGroup(permission.Action, action) {
        if permission.Resource == resource || permission.Resource == "*" {
            return true
        }
    }
    
    return false
}

// matchResourceHierarchy 匹配资源层级
func (s *AuthService) matchResourceHierarchy(permissionResource, requestResource string) bool {
    // 支持层级资源匹配，如 "project:*" 匹配 "project:123"
    if strings.HasSuffix(permissionResource, ":*") {
        prefix := strings.TrimSuffix(permissionResource, ":*")
        return strings.HasPrefix(requestResource, prefix+":")
    }
    
    // 支持路径匹配，如 "/api/*" 匹配 "/api/users"
    if strings.HasSuffix(permissionResource, "/*") {
        prefix := strings.TrimSuffix(permissionResource, "/*")
        return strings.HasPrefix(requestResource, prefix+"/")
    }
    
    return false
}

// matchActionGroup 匹配动作组
func (s *AuthService) matchActionGroup(permissionAction, requestAction string) bool {
    // 定义动作组
    actionGroups := map[string][]string{
        "read":  {"get", "list", "view", "download"},
        "write": {"create", "update", "delete", "upload"},
        "admin": {"manage", "configure", "grant", "revoke"},
    }
    
    if actions, exists := actionGroups[permissionAction]; exists {
        for _, action := range actions {
            if action == requestAction {
                return true
            }
        }
    }
    
    return false
}
```

## 3. 项目服务详细设计

### 3.1 ProjectService 类设计

```go
// ProjectService 项目管理服务
type ProjectService struct {
    projectRepo   ProjectRepository
    configRepo    ConfigRepository
    memberRepo    MemberRepository
    repoService   RepositoryService
    authService   AuthService
    eventBus      EventBus
    validator     Validator
    logger        Logger
    config        ProjectConfig
    metrics       ProjectMetrics
}

// ProjectConfig 项目服务配置
type ProjectConfig struct {
    DefaultConfig     DefaultProjectConfig `yaml:"default_config"`
    RepositoryConfig  RepositoryConfig     `yaml:"repository_config"`
    SyncConfig        SyncConfig           `yaml:"sync_config"`
    ValidationConfig  ValidationConfig     `yaml:"validation_config"`
}

type DefaultProjectConfig struct {
    Languages         []string          `yaml:"languages"`
    ScanRules         []string          `yaml:"scan_rules"`
    Exclusions        []string          `yaml:"exclusions"`
    NotificationRules []NotificationRule `yaml:"notification_rules"`
}

type RepositoryConfig struct {
    SupportedTypes    []string      `yaml:"supported_types"`
    ConnectionTimeout time.Duration `yaml:"connection_timeout"`
    CloneTimeout      time.Duration `yaml:"clone_timeout"`
    MaxRepoSize       int64         `yaml:"max_repo_size_mb"`
}

type SyncConfig struct {
    AutoSyncEnabled   bool          `yaml:"auto_sync_enabled"`
    SyncInterval      time.Duration `yaml:"sync_interval"`
    MaxConcurrentSync int           `yaml:"max_concurrent_sync"`
    RetryAttempts     int           `yaml:"retry_attempts"`
}
```

### 3.2 项目创建实现

```go
// CreateProject 创建项目
func (s *ProjectService) CreateProject(ctx context.Context, req CreateProjectRequest) (*Project, error) {
    s.logger.Info("creating project", "name", req.Name, "owner", req.OwnerID)
    s.metrics.IncProjectCreations()
    
    start := time.Now()
    defer func() {
        s.metrics.RecordProjectCreationDuration(time.Since(start))
    }()
    
    // 1. 验证请求
    if err := s.validateCreateProjectRequest(req); err != nil {
        s.logger.Warn("project creation validation failed", "error", err)
        s.metrics.IncProjectCreationFailures("validation")
        return nil, NewValidationError("invalid project creation request", err)
    }
    
    // 2. 检查权限
    if !s.authService.HasPermission(ctx, req.OwnerID, "project", "create") {
        s.logger.Warn("insufficient permissions for project creation", "userID", req.OwnerID)
        s.metrics.IncProjectCreationFailures("permission")
        return nil, NewPermissionError("insufficient permissions to create project")
    }
    
    // 3. 检查项目名称唯一性
    if exists, err := s.checkProjectNameExists(ctx, req.Name, req.OwnerID); err != nil {
        s.logger.Error("failed to check project name existence", "error", err)
        s.metrics.IncProjectCreationFailures("database")
        return nil, NewInternalError("database error", err)
    } else if exists {
        s.logger.Warn("project name already exists", "name", req.Name, "owner", req.OwnerID)
        s.metrics.IncProjectCreationFailures("conflict")
        return nil, NewConflictError("project name already exists")
    }
    
    // 4. 验证仓库连接
    if err := s.validateRepository(ctx, req.Repository); err != nil {
        s.logger.Warn("repository validation failed", "error", err)
        s.metrics.IncProjectCreationFailures("repository")
        return nil, NewRepositoryError("repository validation failed", err)
    }
    
    // 5. 创建项目对象
    project := &Project{
        ID:          s.generateProjectID(),
        Name:        req.Name,
        Description: req.Description,
        Repository:  req.Repository,
        OwnerID:     req.OwnerID,
        Status:      ProjectStatusActive,
        Settings:    s.createDefaultSettings(),
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
        Version:     1,
    }
    
    // 6. 开始数据库事务
    tx, err := s.db.BeginTx(ctx, nil)
    if err != nil {
        s.logger.Error("failed to begin transaction", "error", err)
        s.metrics.IncProjectCreationFailures("database")
        return nil, NewInternalError("database transaction failed", err)
    }
    defer tx.Rollback()
    
    // 7. 保存项目
    if err := s.projectRepo.SaveWithTx(ctx, tx, project); err != nil {
        s.logger.Error("failed to save project", "error", err)
        s.metrics.IncProjectCreationFailures("database")
        return nil, NewInternalError("failed to save project", err)
    }
    
    // 8. 创建默认配置
    config := s.createDefaultConfiguration(project.ID)
    if err := s.configRepo.SaveWithTx(ctx, tx, config); err != nil {
        s.logger.Error("failed to save project configuration", "error", err)
        s.metrics.IncProjectCreationFailures("database")
        return nil, NewInternalError("failed to save project configuration", err)
    }
    
    // 9. 添加项目所有者为管理员
    member := &ProjectMember{
        ProjectID: project.ID,
        UserID:    req.OwnerID,
        Role:      MemberRoleAdmin,
        AddedAt:   time.Now(),
        AddedBy:   req.OwnerID,
    }
    
    if err := s.memberRepo.SaveWithTx(ctx, tx, member); err != nil {
        s.logger.Error("failed to add project owner as member", "error", err)
        s.metrics.IncProjectCreationFailures("database")
        return nil, NewInternalError("failed to add project member", err)
    }
    
    // 10. 提交事务
    if err := tx.Commit(); err != nil {
        s.logger.Error("failed to commit transaction", "error", err)
        s.metrics.IncProjectCreationFailures("database")
        return nil, NewInternalError("database transaction commit failed", err)
    }
    
    // 11. 异步初始化仓库
    go s.initializeRepositoryAsync(ctx, project)
    
    // 12. 发布项目创建事件
    event := ProjectCreatedEvent{
        ProjectID:   project.ID,
        ProjectName: project.Name,
        OwnerID:     project.OwnerID,
        Repository:  project.Repository,
        Timestamp:   time.Now(),
    }
    
    if err := s.eventBus.Publish(ctx, "project.created", event); err != nil {
        s.logger.Error("failed to publish project created event", "error", err)
    }
    
    // 13. 记录成功指标
    s.metrics.IncProjectCreationSuccesses()
    s.logger.Info("project created successfully", "projectID", project.ID, "name", project.Name)
    
    return project, nil
}

// validateRepository 验证仓库
func (s *ProjectService) validateRepository(ctx context.Context, repo Repository) error {
    // 1. 检查仓库类型支持
    if !s.isSupportedRepositoryType(repo.Type) {
        return fmt.Errorf("unsupported repository type: %s", repo.Type)
    }
    
    // 2. 验证URL格式
    if err := s.validateRepositoryURL(repo.URL); err != nil {
        return fmt.Errorf("invalid repository URL: %w", err)
    }
    
    // 3. 测试连接
    ctx, cancel := context.WithTimeout(ctx, s.config.RepositoryConfig.ConnectionTimeout)
    defer cancel()
    
    if err := s.repoService.TestConnection(ctx, repo); err != nil {
        return fmt.Errorf("repository connection test failed: %w", err)
    }
    
    // 4. 检查仓库大小
    size, err := s.repoService.GetRepositorySize(ctx, repo)
    if err != nil {
        s.logger.Warn("failed to get repository size", "error", err)
        // 不阻止项目创建，只记录警告
    } else if size > s.config.RepositoryConfig.MaxRepoSize*1024*1024 {
        return fmt.Errorf("repository size (%d MB) exceeds maximum allowed size (%d MB)", 
            size/(1024*1024), s.config.RepositoryConfig.MaxRepoSize)
    }
    
    // 5. 检查分支存在性
    if repo.Branch != "" {
        branches, err := s.repoService.ListBranches(ctx, repo)
        if err != nil {
            return fmt.Errorf("failed to list repository branches: %w", err)
        }
        
        branchExists := false
        for _, branch := range branches {
            if branch.Name == repo.Branch {
                branchExists = true
                break
            }
        }
        
        if !branchExists {
            return fmt.Errorf("specified branch '%s' does not exist", repo.Branch)
        }
    }
    
    return nil
}

// initializeRepositoryAsync 异步初始化仓库
func (s *ProjectService) initializeRepositoryAsync(ctx context.Context, project *Project) {
    s.logger.Info("initializing repository", "projectID", project.ID)
    
    defer func() {
        if r := recover(); r != nil {
            s.logger.Error("repository initialization panicked", "projectID", project.ID, "error", r)
        }
    }()
    
    // 1. 克隆仓库
    localPath, err := s.repoService.CloneRepository(ctx, project.Repository)
    if err != nil {
        s.logger.Error("failed to clone repository", "projectID", project.ID, "error", err)
        s.updateProjectStatus(ctx, project.ID, ProjectStatusError, err.Error())
        return
    }
    
    // 2. 分析代码结构
    analysis, err := s.analyzeCodeStructure(ctx, localPath)
    if err != nil {
        s.logger.Error("failed to analyze code structure", "projectID", project.ID, "error", err)
        // 不阻止初始化，继续执行
    }
    
    // 3. 更新项目信息
    updates := map[string]interface{}{
        "local_path":     localPath,
        "code_analysis":  analysis,
        "initialized_at": time.Now(),
        "status":         ProjectStatusReady,
    }
    
    if err := s.projectRepo.UpdateFields(ctx, project.ID, updates); err != nil {
        s.logger.Error("failed to update project after initialization", "projectID", project.ID, "error", err)
        return
    }
    
    // 4. 发布初始化完成事件
    event := ProjectInitializedEvent{
        ProjectID:   project.ID,
        LocalPath:   localPath,
        Analysis:    analysis,
        Timestamp:   time.Now(),
    }
    
    if err := s.eventBus.Publish(ctx, "project.initialized", event); err != nil {
        s.logger.Error("failed to publish project initialized event", "error", err)
    }
    
    s.logger.Info("repository initialization completed", "projectID", project.ID)
}
```

## 4. 扫描引擎详细设计

### 4.1 ScanEngine 类设计

```go
// ScanEngine 扫描引擎实现
type ScanEngine struct {
    taskRepo      ScanTaskRepository
    resultRepo    ScanResultRepository
    projectRepo   ProjectRepository
    ruleEngine    RuleEngine
    codeAnalyzer  CodeAnalyzer
    aiService     AIService
    scheduler     TaskScheduler
    eventBus      EventBus
    logger        Logger
    config        ScanEngineConfig
    metrics       ScanEngineMetrics
    
    // 并发控制
    semaphore     chan struct{}
    workerPool    *WorkerPool
}

// ScanEngineConfig 扫描引擎配置
type ScanEngineConfig struct {
    MaxConcurrentScans int           `yaml:"max_concurrent_scans"`
    ScanTimeout        time.Duration `yaml:"scan_timeout"`
    ChunkSize          int           `yaml:"chunk_size"`
    RetryAttempts      int           `yaml:"retry_attempts"`
    RetryDelay         time.Duration `yaml:"retry_delay"`
    
    AnalyzerConfig     AnalyzerConfig     `yaml:"analyzer_config"`
    RuleEngineConfig   RuleEngineConfig   `yaml:"rule_engine_config"`
    AIServiceConfig    AIServiceConfig    `yaml:"ai_service_config"`
}

type AnalyzerConfig struct {
    SupportedLanguages []string          `yaml:"supported_languages"`
    ParserTimeout      time.Duration     `yaml:"parser_timeout"`
    MaxFileSize        int64             `yaml:"max_file_size_mb"`
    ExcludePatterns    []string          `yaml:"exclude_patterns"`
}

type RuleEngineConfig struct {
    RuleSetPath        string            `yaml:"rule_set_path"`
    CustomRulesEnabled bool              `yaml:"custom_rules_enabled"`
    RuleCacheSize      int               `yaml:"rule_cache_size"`
    RuleCacheTTL       time.Duration     `yaml:"rule_cache_ttl"`
}

type AIServiceConfig struct {
    Enabled            bool              `yaml:"enabled"`
    ModelEndpoint      string            `yaml:"model_endpoint"`
    RequestTimeout     time.Duration     `yaml:"request_timeout"`
    MaxRetries         int               `yaml:"max_retries"`
    BatchSize          int               `yaml:"batch_size"`
}
```

### 4.2 扫描任务执行实现

```go
// StartScan 启动扫描任务
func (e *ScanEngine) StartScan(ctx context.Context, req ScanRequest) (*ScanTask, error) {
    e.logger.Info("starting scan", "projectID", req.ProjectID, "type", req.ScanType)
    e.metrics.IncScanRequests()
    
    // 1. 验证扫描请求
    if err := e.validateScanRequest(req); err != nil {
        e.logger.Warn("scan request validation failed", "error", err)
        e.metrics.IncScanRequestFailures("validation")
        return nil, NewValidationError("invalid scan request", err)
    }
    
    // 2. 检查并发限制
    if e.getCurrentScanCount() >= e.config.MaxConcurrentScans {
        e.logger.Warn("maximum concurrent scans reached", "current", e.getCurrentScanCount(), "max", e.config.MaxConcurrentScans)
        e.metrics.IncScanRequestFailures("concurrency_limit")
        return nil, NewResourceError("maximum concurrent scans reached")
    }
    
    // 3. 获取项目信息
    project, err := e.projectRepo.FindByID(ctx, req.ProjectID)
    if err != nil {
        e.logger.Error("failed to get project", "projectID", req.ProjectID, "error", err)
        e.metrics.IncScanRequestFailures("project_not_found")
        return nil, NewNotFoundError("project not found", err)
    }
    
    // 4. 检查项目状态
    if project.Status != ProjectStatusReady {
        e.logger.Warn("project not ready for scanning", "projectID", req.ProjectID, "status", project.Status)
        e.metrics.IncScanRequestFailures("project_not_ready")
        return nil, NewStateError("project is not ready for scanning")
    }
    
    // 5. 创建扫描任务
    task := &ScanTask{
        ID:          e.generateTaskID(),
        ProjectID:   req.ProjectID,
        Type:        req.ScanType,
        Status:      ScanStatusPending,
        Config:      e.mergeScanConfig(project.Configuration, req.Configuration),
        Priority:    req.Priority,
        RequestedBy: req.RequestedBy,
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
    }
    
    // 6. 保存任务到数据库
    if err := e.taskRepo.Save(ctx, task); err != nil {
        e.logger.Error("failed to save scan task", "error", err)
        e.metrics.IncScanRequestFailures("database")
        return nil, NewInternalError("failed to create scan task", err)
    }
    
    // 7. 提交任务到调度器
    if err := e.scheduler.ScheduleTask(ctx, task); err != nil {
        e.logger.Error("failed to schedule scan task", "taskID", task.ID, "error", err)
        e.updateTaskStatus(ctx, task.ID, ScanStatusFailed, "failed to schedule task")
        e.metrics.IncScanRequestFailures("scheduler")
        return nil, NewInternalError("failed to schedule scan task", err)
    }
    
    // 8. 发布任务创建事件
    event := ScanTaskCreatedEvent{
        TaskID:      task.ID,
        ProjectID:   task.ProjectID,
        ScanType:    task.Type,
        RequestedBy: task.RequestedBy,
        Timestamp:   time.Now(),
    }
    
    if err := e.eventBus.Publish(ctx, "scan.task.created", event); err != nil {
        e.logger.Error("failed to publish scan task created event", "error", err)
    }
    
    e.metrics.IncScanRequestSuccesses()
    e.logger.Info("scan task created", "taskID", task.ID, "projectID", req.ProjectID)
    
    return task, nil
}

// ExecuteScan 执行扫描任务
func (e *ScanEngine) ExecuteScan(ctx context.Context, task *ScanTask) error {
    e.logger.Info("executing scan task", "taskID", task.ID, "projectID", task.ProjectID)
    
    // 设置超时上下文
    ctx, cancel := context.WithTimeout(ctx, e.config.ScanTimeout)
    defer cancel()
    
    // 更新任务状态为运行中
    if err := e.updateTaskStatus(ctx, task.ID, ScanStatusRunning, ""); err != nil {
        e.logger.Error("failed to update task status to running", "taskID", task.ID, "error", err)
        return err
    }
    
    // 发布任务开始事件
    e.publishScanEvent(ctx, "scan.task.started", task.ID, task.ProjectID, nil)
    
    // 执行扫描的主要逻辑
    result, err := e.performScan(ctx, task)
    if err != nil {
        e.logger.Error("scan execution failed", "taskID", task.ID, "error", err)
        e.updateTaskStatus(ctx, task.ID, ScanStatusFailed, err.Error())
        e.publishScanEvent(ctx, "scan.task.failed", task.ID, task.ProjectID, map[string]interface{}{"error": err.Error()})
        return err
    }
    
    // 保存扫描结果
    if err := e.resultRepo.Save(ctx, result); err != nil {
        e.logger.Error("failed to save scan result", "taskID", task.ID, "error", err)
        e.updateTaskStatus(ctx, task.ID, ScanStatusFailed, "failed to save results")
        return err
    }
    
    // 更新任务状态为完成
    if err := e.updateTaskStatus(ctx, task.ID, ScanStatusCompleted, ""); err != nil {
        e.logger.Error("failed to update task status to completed", "taskID", task.ID, "error", err)
        return err
    }
    
    // 发布任务完成事件
    e.publishScanEvent(ctx, "scan.task.completed", task.ID, task.ProjectID, map[string]interface{}{
        "issues_count": len(result.Issues),
        "duration":     result.Duration,
    })
    
    e.logger.Info("scan task completed", "taskID", task.ID, "issues", len(result.Issues))
    return nil
}

// performScan 执行具体的扫描逻辑
func (e *ScanEngine) performScan(ctx context.Context, task *ScanTask) (*ScanResult, error) {
    startTime := time.Now()
    
    // 1. 获取项目信息
    project, err := e.projectRepo.FindByID(ctx, task.ProjectID)
    if err != nil {
        return nil, fmt.Errorf("failed to get project: %w", err)
    }
    
    // 2. 获取要扫描的文件列表
    files, err := e.getFilesToScan(ctx, project, task)
    if err != nil {
        return nil, fmt.Errorf("failed to get files to scan: %w", err)
    }
    
    e.logger.Info("files to scan", "taskID", task.ID, "count", len(files))
    
    // 3. 初始化扫描结果
    result := &ScanResult{
        TaskID:    task.ID,
        ProjectID: task.ProjectID,
        StartTime: startTime,
        Files:     files,
        Issues:    make([]Issue, 0),
        Metrics:   make(map[string]interface{}),
    }
    
    // 4. 分批处理文件
    totalFiles := len(files)
    processedFiles := 0
    
    for i := 0; i < totalFiles; i += e.config.ChunkSize {
        end := i + e.config.ChunkSize
        if end > totalFiles {
            end = totalFiles
        }
        
        chunk := files[i:end]
        
        // 处理文件块
        chunkIssues, err := e.processFileChunk(ctx, chunk, task.Config)
        if err != nil {
            e.logger.Error("failed to process file chunk", "taskID", task.ID, "chunk", i/e.config.ChunkSize, "error", err)
            // 继续处理其他块，不中断整个扫描
            continue
        }
        
        result.Issues = append(result.Issues, chunkIssues...)
        processedFiles += len(chunk)
        
        // 更新进度
        progress := float64(processedFiles) / float64(totalFiles)
        e.updateTaskProgress(ctx, task.ID, progress)
        
        // 发布进度事件
        e.publishScanEvent(ctx, "scan.task.progress", task.ID, task.ProjectID, map[string]interface{}{
            "progress":        progress,
            "processed_files": processedFiles,
            "total_files":     totalFiles,
        })
        
        // 检查上下文取消
        select {
        case <-ctx.Done():
            return nil, ctx.Err()
        default:
        }
    }
    
    // 5. 生成扫描摘要
    result.Summary = e.generateScanSummary(result.Issues)
    result.EndTime = time.Now()
    result.Duration = result.EndTime.Sub(result.StartTime)
    
    // 6. 计算扫描指标
    result.Metrics = e.calculateScanMetrics(result)
    
    e.logger.Info("scan completed", "taskID", task.ID, "duration", result.Duration, "issues", len(result.Issues))
    
    return result, nil
}

// processFileChunk 处理文件块
func (e *ScanEngine) processFileChunk(ctx context.Context, files []string, config ScanConfig) ([]Issue, error) {
    var allIssues []Issue
    
    // 使用工作池并发处理文件
    jobs := make(chan string, len(files))
    results := make(chan []Issue, len(files))
    errors := make(chan error, len(files))
    
    // 启动工作协程
    numWorkers := min(len(files), e.workerPool.Size())
    for i := 0; i < numWorkers; i++ {
        go func() {
            for file := range jobs {
                issues, err := e.processFile(ctx, file, config)
                if err != nil {
                    errors <- err
                } else {
                    results <- issues
                }
            }
        }()
    }
    
    // 发送任务
    for _, file := range files {
        jobs <- file
    }
    close(jobs)
    
    // 收集结果
    for i := 0; i < len(files); i++ {
        select {
        case issues := <-results:
            allIssues = append(allIssues, issues...)
        case err := <-errors:
            e.logger.Error("file processing error", "error", err)
            // 继续处理其他文件
        case <-ctx.Done():
            return nil, ctx.Err()
        }
    }
    
    return allIssues, nil
}

// processFile 处理单个文件
func (e *ScanEngine) processFile(ctx context.Context, file string, config ScanConfig) ([]Issue, error) {
    var allIssues []Issue
    
    // 1. 代码分析
    analysis, err := e.codeAnalyzer.AnalyzeFile(ctx, file)
    if err != nil {
        return nil, fmt.Errorf("code analysis failed for %s: %w", file, err)
    }
    
    // 2. 规则引擎检测
    ruleIssues, err := e.ruleEngine.ExecuteRules(ctx, analysis, config.Rules)
    if err != nil {
        e.logger.Error("rule engine execution failed", "file", file, "error", err)
    } else {
        allIssues = append(allIssues, ruleIssues...)
    }
    
    // 3. AI检测（如果启用）
    if e.config.AIServiceConfig.Enabled && config.EnableAI {
        aiIssues, err := e.aiService.AnalyzeCode(ctx, analysis)
        if err != nil {
            e.logger.Error("AI analysis failed", "file", file, "error", err)
        } else {
            allIssues = append(allIssues, aiIssues...)
        }
    }
    
    // 4. 去重和合并问题
    deduplicatedIssues := e.deduplicateIssues(allIssues)
    
    // 5. 应用过滤器
    filteredIssues := e.applyFilters(deduplicatedIssues, config.Filters)
    
    return filteredIssues, nil
}
```

## 5. AI智能体详细设计

### 5.1 AIService 类设计

```go
// AIService AI智能体服务
type AIService struct {
    modelManager    ModelManager
    featureExtractor FeatureExtractor
    preprocessor    CodePreprocessor
    postprocessor   ResultPostprocessor
    cache          CacheManager
    logger         Logger
    config         AIServiceConfig
    metrics        AIServiceMetrics
    
    // 模型实例
    codeModel      Model
    securityModel  Model
    qualityModel   Model
}

// AIServiceConfig AI服务配置
type AIServiceConfig struct {
    ModelConfig       ModelConfig       `yaml:"model_config"`
    ProcessingConfig  ProcessingConfig  `yaml:"processing_config"`
    CacheConfig       CacheConfig       `yaml:"cache_config"`
    PerformanceConfig PerformanceConfig `yaml:"performance_config"`
}

type ModelConfig struct {
    CodeModelPath     string            `yaml:"code_model_path"`
    SecurityModelPath string            `yaml:"security_model_path"`
    QualityModelPath  string            `yaml:"quality_model_path"`
    ModelTimeout      time.Duration     `yaml:"model_timeout"`
    MaxTokens         int               `yaml:"max_tokens"`
    Temperature       float32           `yaml:"temperature"`
    TopP              float32           `yaml:"top_p"`
}

type ProcessingConfig struct {
    MaxFileSize       int64             `yaml:"max_file_size_mb"`
    ChunkSize         int               `yaml:"chunk_size"`
    OverlapSize       int               `yaml:"overlap_size"`
    MinConfidence     float64           `yaml:"min_confidence"`
    EnablePreprocess  bool              `yaml:"enable_preprocess"`
    EnablePostprocess bool              `yaml:"enable_postprocess"`
}

type PerformanceConfig struct {
    MaxConcurrentRequests int           `yaml:"max_concurrent_requests"`
    RequestTimeout        time.Duration `yaml:"request_timeout"`
    RetryAttempts         int           `yaml:"retry_attempts"`
    RetryDelay            time.Duration `yaml:"retry_delay"`
}
```

### 5.2 代码分析实现

```go
// AnalyzeCode AI代码分析
func (s *AIService) AnalyzeCode(ctx context.Context, analysis *CodeAnalysis) ([]Issue, error) {
    s.logger.Info("starting AI code analysis", "file", analysis.FilePath)
    s.metrics.IncAIAnalysisRequests()
    
    start := time.Now()
    defer func() {
        s.metrics.RecordAIAnalysisDuration(time.Since(start))
    }()
    
    // 1. 预处理代码
    preprocessedCode, err := s.preprocessor.Process(ctx, analysis)
    if err != nil {
        s.logger.Error("code preprocessing failed", "error", err)
        s.metrics.IncAIAnalysisFailures("preprocessing")
        return nil, fmt.Errorf("code preprocessing failed: %w", err)
    }
    
    // 2. 特征提取
    features, err := s.featureExtractor.Extract(ctx, preprocessedCode)
    if err != nil {
        s.logger.Error("feature extraction failed", "error", err)
        s.metrics.IncAIAnalysisFailures("feature_extraction")
        return nil, fmt.Errorf("feature extraction failed: %w", err)
    }
    
    // 3. 检查缓存
    cacheKey := s.generateCacheKey(features)
    if cached, err := s.cache.Get(ctx, cacheKey); err == nil {
        s.metrics.IncAICacheHits()
        return cached.([]Issue), nil
    }
    s.metrics.IncAICacheMisses()
    
    // 4. 并发执行多个模型分析
    var issues []Issue
    var mu sync.Mutex
    var wg sync.WaitGroup
    
    // 安全漏洞检测
    wg.Add(1)
    go func() {
        defer wg.Done()
        securityIssues, err := s.analyzeSecurityVulnerabilities(ctx, features)
        if err != nil {
            s.logger.Error("security analysis failed", "error", err)
            return
        }
        mu.Lock()
        issues = append(issues, securityIssues...)
        mu.Unlock()
    }()
    
    // 代码质量检测
    wg.Add(1)
    go func() {
        defer wg.Done()
        qualityIssues, err := s.analyzeCodeQuality(ctx, features)
        if err != nil {
            s.logger.Error("quality analysis failed", "error", err)
            return
        }
        mu.Lock()
        issues = append(issues, qualityIssues...)
        mu.Unlock()
    }()
    
    // 逻辑错误检测
    wg.Add(1)
    go func() {
        defer wg.Done()
        logicIssues, err := s.analyzeLogicErrors(ctx, features)
        if err != nil {
            s.logger.Error("logic analysis failed", "error", err)
            return
        }
        mu.Lock()
        issues = append(issues, logicIssues...)
        mu.Unlock()
    }()
    
    // 等待所有分析完成
    wg.Wait()
    
    // 5. 后处理结果
    processedIssues, err := s.postprocessor.Process(ctx, issues, analysis)
    if err != nil {
        s.logger.Error("result postprocessing failed", "error", err)
        s.metrics.IncAIAnalysisFailures("postprocessing")
        return nil, fmt.Errorf("result postprocessing failed: %w", err)
    }
    
    // 6. 缓存结果
    if err := s.cache.Set(ctx, cacheKey, processedIssues, s.config.CacheConfig.TTL); err != nil {
        s.logger.Error("failed to cache AI analysis result", "error", err)
    }
    
    s.metrics.IncAIAnalysisSuccesses()
    s.logger.Info("AI code analysis completed", "file", analysis.FilePath, "issues", len(processedIssues))
    
    return processedIssues, nil
}

// analyzeSecurityVulnerabilities 安全漏洞分析
func (s *AIService) analyzeSecurityVulnerabilities(ctx context.Context, features *CodeFeatures) ([]Issue, error) {
    // 1. 构建安全分析提示
    prompt := s.buildSecurityAnalysisPrompt(features)
    
    // 2. 调用安全模型
    response, err := s.securityModel.Predict(ctx, prompt)
    if err != nil {
        return nil, fmt.Errorf("security model prediction failed: %w", err)
    }
    
    // 3. 解析模型响应
    vulnerabilities, err := s.parseSecurityResponse(response)
    if err != nil {
        return nil, fmt.Errorf("failed to parse security response: %w", err)
    }
    
    // 4. 转换为Issue格式
    var issues []Issue
    for _, vuln := range vulnerabilities {
        if vuln.Confidence >= s.config.ProcessingConfig.MinConfidence {
            issue := Issue{
                ID:          s.generateIssueID(),
                Type:        IssueTypeSecurity,
                Severity:    s.mapSeverity(vuln.Severity),
                Title:       vuln.Title,
                Description: vuln.Description,
                Location: Location{
                    File:   features.FilePath,
                    Line:   vuln.Line,
                    Column: vuln.Column,
                },
                Rule:        vuln.Rule,
                Confidence:  vuln.Confidence,
                Source:      "AI-Security",
                CreatedAt:   time.Now(),
            }
            issues = append(issues, issue)
        }
    }
    
    return issues, nil
}

// buildSecurityAnalysisPrompt 构建安全分析提示
func (s *AIService) buildSecurityAnalysisPrompt(features *CodeFeatures) string {
    var prompt strings.Builder
    
    prompt.WriteString("Analyze the following code for security vulnerabilities:\n\n")
    prompt.WriteString("Language: " + features.Language + "\n")
    prompt.WriteString("File: " + features.FilePath + "\n\n")
    prompt.WriteString("Code:\n```" + features.Language + "\n")
    prompt.WriteString(features.Content)
    prompt.WriteString("\n```\n\n")
    
    prompt.WriteString("Please identify potential security issues including but not limited to:\n")
    prompt.WriteString("- SQL Injection vulnerabilities\n")
    prompt.WriteString("- Cross-Site Scripting (XSS)\n")
    prompt.WriteString("- Authentication and authorization flaws\n")
    prompt.WriteString("- Input validation issues\n")
    prompt.WriteString("- Cryptographic weaknesses\n")
    prompt.WriteString("- Information disclosure\n")
    prompt.WriteString("- Buffer overflows\n")
    prompt.WriteString("- Race conditions\n\n")
    
    prompt.WriteString("For each issue found, provide:\n")
    prompt.WriteString("1. Issue type and severity (Critical/High/Medium/Low)\n")
    prompt.WriteString("2. Line number and description\n")
    prompt.WriteString("3. Potential impact\n")
    prompt.WriteString("4. Recommended fix\n")
    prompt.WriteString("5. Confidence level (0.0-1.0)\n\n")
    
    prompt.WriteString("Format the response as JSON array of vulnerability objects.")
    
    return prompt.String()
}

// parseSecurityResponse 解析安全分析响应
func (s *AIService) parseSecurityResponse(response string) ([]SecurityVulnerability, error) {
    // 清理响应文本
    cleanResponse := s.cleanModelResponse(response)
    
    // 解析JSON
    var vulnerabilities []SecurityVulnerability
    if err := json.Unmarshal([]byte(cleanResponse), &vulnerabilities); err != nil {
        // 如果JSON解析失败，尝试使用正则表达式提取信息
        return s.extractVulnerabilitiesFromText(response)
    }
    
    return vulnerabilities, nil
}

// SecurityVulnerability 安全漏洞结构
type SecurityVulnerability struct {
    Type        string  `json:"type"`
    Severity    string  `json:"severity"`
    Title       string  `json:"title"`
    Description string  `json:"description"`
    Line        int     `json:"line"`
    Column      int     `json:"column"`
    Impact      string  `json:"impact"`
    Fix         string  `json:"fix"`
    Rule        string  `json:"rule"`
    Confidence  float64 `json:"confidence"`
}
```

## 6. 数据访问层详细设计

### 6.1 Repository接口设计

```go
// UserRepository 用户数据访问接口
type UserRepository interface {
    // 基础CRUD操作
    Save(ctx context.Context, user *User) error
    SaveWithTx(ctx context.Context, tx *sql.Tx, user *User) error
    FindByID(ctx context.Context, id string) (*User, error)
    FindByUsername(ctx context.Context, username string) (*User, error)
    FindByEmail(ctx context.Context, email string) (*User, error)
    Update(ctx context.Context, user *User) error
    UpdateFields(ctx context.Context, id string, fields map[string]interface{}) error
    Delete(ctx context.Context, id string) error
    
    // 查询操作
    List(ctx context.Context, filter UserFilter, pagination Pagination) ([]*User, int64, error)
    Search(ctx context.Context, query string, pagination Pagination) ([]*User, int64, error)
    
    // 批量操作
    BatchSave(ctx context.Context, users []*User) error
    BatchUpdate(ctx context.Context, updates []UserUpdate) error
    
    // 统计操作
    Count(ctx context.Context, filter UserFilter) (int64, error)
    CountByStatus(ctx context.Context, status UserStatus) (int64, error)
    
    // 关联查询
    FindWithRoles(ctx context.Context, id string) (*User, error)
    FindUsersInRole(ctx context.Context, roleID string) ([]*User, error)
}

// UserFilter 用户过滤条件
type UserFilter struct {
    Status    *UserStatus `json:"status,omitempty"`
    Role      *string     `json:"role,omitempty"`
    CreatedAt *TimeRange  `json:"created_at,omitempty"`
    UpdatedAt *TimeRange  `json:"updated_at,omitempty"`
    Keywords  *string     `json:"keywords,omitempty"`
}

// UserUpdate 用户更新结构
type UserUpdate struct {
    ID     string                 `json:"id"`
    Fields map[string]interface{} `json:"fields"`
}
```

### 6.2 PostgreSQL实现

```go
// PostgreSQLUserRepository PostgreSQL用户仓储实现
type PostgreSQLUserRepository struct {
    db      *sql.DB
    logger  Logger
    metrics RepositoryMetrics
}

// Save 保存用户
func (r *PostgreSQLUserRepository) Save(ctx context.Context, user *User) error {
    return r.SaveWithTx(ctx, nil, user)
}

// SaveWithTx 在事务中保存用户
func (r *PostgreSQLUserRepository) SaveWithTx(ctx context.Context, tx *sql.Tx, user *User) error {
    r.logger.Debug("saving user", "userID", user.ID)
    r.metrics.IncRepositoryOperations("user", "save")
    
    start := time.Now()
    defer func() {
        r.metrics.RecordRepositoryDuration("user", "save", time.Since(start))
    }()
    
    query := `
        INSERT INTO users (
            id, username, email, password_hash, display_name, 
            status, created_at, updated_at, version
        ) VALUES (
            $1, $2, $3, $4, $5, $6, $7, $8, $9
        )
        ON CONFLICT (id) DO UPDATE SET
            username = EXCLUDED.username,
            email = EXCLUDED.email,
            password_hash = EXCLUDED.password_hash,
            display_name = EXCLUDED.display_name,
            status = EXCLUDED.status,
            updated_at = EXCLUDED.updated_at,
            version = users.version + 1
    `
    
    var err error
    if tx != nil {
        _, err = tx.ExecContext(ctx, query,
            user.ID, user.Username, user.Email, user.PasswordHash,
            user.DisplayName, user.Status, user.CreatedAt, user.UpdatedAt, user.Version)
    } else {
        _, err = r.db.ExecContext(ctx, query,
            user.ID, user.Username, user.Email, user.PasswordHash,
            user.DisplayName, user.Status, user.CreatedAt, user.UpdatedAt, user.Version)
    }
    
    if err != nil {
        r.logger.Error("failed to save user", "userID", user.ID, "error", err)
        r.metrics.IncRepositoryErrors("user", "save")
        return fmt.Errorf("failed to save user: %w", err)
    }
    
    // 保存用户角色关联
    if err := r.saveUserRoles(ctx, tx, user.ID, user.Roles); err != nil {
        r.logger.Error("failed to save user roles", "userID", user.ID, "error", err)
        return fmt.Errorf("failed to save user roles: %w", err)
    }
    
    r.metrics.IncRepositorySuccesses("user", "save")
    r.logger.Debug("user saved successfully", "userID", user.ID)
    
    return nil
}

// FindByID 根据ID查找用户
func (r *PostgreSQLUserRepository) FindByID(ctx context.Context, id string) (*User, error) {
    r.logger.Debug("finding user by ID", "userID", id)
    r.metrics.IncRepositoryOperations("user", "find_by_id")
    
    start := time.Now()
    defer func() {
        r.metrics.RecordRepositoryDuration("user", "find_by_id", time.Since(start))
    }()
    
    query := `
        SELECT 
            id, username, email, password_hash, display_name,
            status, last_login_at, password_updated_at,
            created_at, updated_at, version
        FROM users 
        WHERE id = $1 AND deleted_at IS NULL
    `
    
    var user User
    var lastLoginAt, passwordUpdatedAt sql.NullTime
    
    err := r.db.QueryRowContext(ctx, query, id).Scan(
        &user.ID, &user.Username, &user.Email, &user.PasswordHash,
        &user.DisplayName, &user.Status, &lastLoginAt, &passwordUpdatedAt,
        &user.CreatedAt, &user.UpdatedAt, &user.Version,
    )
    
    if err != nil {
        if err == sql.ErrNoRows {
            r.metrics.IncRepositoryMisses("user", "find_by_id")
            return nil, ErrUserNotFound
        }
        r.logger.Error("failed to find user by ID", "userID", id, "error", err)
        r.metrics.IncRepositoryErrors("user", "find_by_id")
        return nil, fmt.Errorf("failed to find user: %w", err)
    }
    
    // 处理可空字段
    if lastLoginAt.Valid {
        user.LastLoginAt = lastLoginAt.Time
    }
    if passwordUpdatedAt.Valid {
        user.PasswordUpdatedAt = passwordUpdatedAt.Time
    }
    
    r.metrics.IncRepositoryHits("user", "find_by_id")
    r.logger.Debug("user found", "userID", id)
    
    return &user, nil
}

// List 列出用户
func (r *PostgreSQLUserRepository) List(ctx context.Context, filter UserFilter, pagination Pagination) ([]*User, int64, error) {
    r.logger.Debug("listing users", "filter", filter, "pagination", pagination)
    r.metrics.IncRepositoryOperations("user", "list")
    
    start := time.Now()
    defer func() {
        r.metrics.RecordRepositoryDuration("user", "list", time.Since(start))
    }()
    
    // 构建查询条件
    whereClause, args := r.buildWhereClause(filter)
    
    // 查询总数
    countQuery := "SELECT COUNT(*) FROM users WHERE " + whereClause
    var total int64
    if err := r.db.QueryRowContext(ctx, countQuery, args...).Scan(&total); err != nil {
        r.logger.Error("failed to count users", "error", err)
        r.metrics.IncRepositoryErrors("user", "list")
        return nil, 0, fmt.Errorf("failed to count users: %w", err)
    }
    
    // 构建分页查询
    query := `
        SELECT 
            id, username, email, display_name, status,
            last_login_at, created_at, updated_at, version
        FROM users 
        WHERE ` + whereClause + `
        ORDER BY ` + pagination.OrderBy + ` ` + pagination.Order + `
        LIMIT $` + fmt.Sprintf("%d", len(args)+1) + ` OFFSET $` + fmt.Sprintf("%d", len(args)+2)
    
    args = append(args, pagination.Limit, pagination.Offset)
    
    rows, err := r.db.QueryContext(ctx, query, args...)
    if err != nil {
        r.logger.Error("failed to query users", "error", err)
        r.metrics.IncRepositoryErrors("user", "list")
        return nil, 0, fmt.Errorf("failed to query users: %w", err)
    }
    defer rows.Close()
    
    var users []*User
    for rows.Next() {
        var user User
        var lastLoginAt sql.NullTime
        
        if err := rows.Scan(
            &user.ID, &user.Username, &user.Email, &user.DisplayName,
            &user.Status, &lastLoginAt, &user.CreatedAt, &user.UpdatedAt, &user.Version,
        ); err != nil {
            r.logger.Error("failed to scan user row", "error", err)
            continue
        }
        
        if lastLoginAt.Valid {
            user.LastLoginAt = lastLoginAt.Time
        }
        
        users = append(users, &user)
    }
    
    if err := rows.Err(); err != nil {
        r.logger.Error("error iterating user rows", "error", err)
        r.metrics.IncRepositoryErrors("user", "list")
        return nil, 0, fmt.Errorf("error iterating users: %w", err)
    }
    
    r.metrics.IncRepositorySuccesses("user", "list")
    r.logger.Debug("users listed", "count", len(users), "total", total)
    
    return users, total, nil
}

// buildWhereClause 构建WHERE子句
func (r *PostgreSQLUserRepository) buildWhereClause(filter UserFilter) (string, []interface{}) {
    var conditions []string
    var args []interface{}
    argIndex := 1
    
    // 默认条件：未删除
    conditions = append(conditions, "deleted_at IS NULL")
    
    // 状态过滤
    if filter.Status != nil {
        conditions = append(conditions, fmt.Sprintf("status = $%d", argIndex))
        args = append(args, *filter.Status)
        argIndex++
    }
    
    // 时间范围过滤
    if filter.CreatedAt != nil {
        if filter.CreatedAt.Start != nil {
            conditions = append(conditions, fmt.Sprintf("created_at >= $%d", argIndex))
            args = append(args, *filter.CreatedAt.Start)
            argIndex++
        }
        if filter.CreatedAt.End != nil {
            conditions = append(conditions, fmt.Sprintf("created_at <= $%d", argIndex))
            args = append(args, *filter.CreatedAt.End)
            argIndex++
        }
    }
    
    // 关键词搜索
    if filter.Keywords != nil && *filter.Keywords != "" {
        searchCondition := fmt.Sprintf(
            "(username ILIKE $%d OR email ILIKE $%d OR display_name ILIKE $%d)",
            argIndex, argIndex, argIndex,
        )
        conditions = append(conditions, searchCondition)
        searchTerm := "%" + *filter.Keywords + "%"
        args = append(args, searchTerm)
        argIndex++
    }
    
    whereClause := strings.Join(conditions, " AND ")
    return whereClause, args
}
```