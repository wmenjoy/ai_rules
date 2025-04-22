# Java 代码规范

## 1. 项目结构
- 遵循标准Maven/Gradle目录结构。
- 推荐分层：controller、service、repository、model、util等。
- 目录结构示例：
  ```
  project/
  ├── src/
  │   ├── main/
  │   │   ├── java/com/example/project/controller
  │   │   ├── java/com/example/project/service
  │   │   ├── java/com/example/project/repository
  │   │   ├── java/com/example/project/model
  │   │   └── resources/
  │   └── test/
  ├── docs/
  ├── scripts/
  ├── Dockerfile
  ├── README.md
  └── pom.xml / build.gradle
  ```
- 推荐使用Spring Boot最佳实践，控制层、服务层、持久层、实体层分明。

## 2. 命名规范
- 包名全部小写，使用反域名。
- 类名使用大驼峰（PascalCase）。
- 接口名以I开头（如IUserService）。
- 方法、变量名使用小驼峰（camelCase）。
- 常量全大写，单词间下划线分隔。
- 数据库表、字段名小写下划线。

## 3. 代码风格
- 遵循Google Java风格指南。
- 每行不超过120字符。
- 使用4空格缩进，不使用Tab。
- 大括号不换行。
- 文件名与类名保持一致。
- 所有文件使用UTF-8编码。
- 注释规范，类/方法/复杂逻辑需有Javadoc注释。
- 创建Web服务时遵循RESTful API设计。

## 4. 依赖管理
- 推荐使用Maven或Gradle统一管理依赖。
- 依赖版本锁定，避免使用最新快照。
- 合理使用Spring Boot Starter简化依赖。

## 5. 文档与注释规范
- 必须遵守 AI 代码标注规范（见下文）。
- 每个类使用类级别Javadoc。
- 公共方法添加方法级别Javadoc。
- 项目需包含README.md、CHANGELOG.md、/docs目录。
- 推荐使用 Springdoc OpenAPI (基于 Swagger) 生成接口文档。
- 使用 `@author`、`@version`、`@since` 等 Javadoc 标签。
- 对复杂算法或业务逻辑进行详细说明。

## 6. 编码规范
- 推荐使用Checkstyle、SpotBugs等工具进行静态检查。
- 单元测试放在test目录，测试类以Test结尾。
- 避免魔法数字，使用常量。
- 错误处理要规范，优先使用自定义异常。
- 推荐使用Lombok简化代码，但需团队统一。
- 推荐使用构造器注入依赖，避免字段注入。
- 合理使用事务、日志、校验等Spring特性。
- 遵循 SOLID 原则，保持高内聚、低耦合。
- 推荐使用函数式编程风格，优先使用不可变对象。
- 遵循数据导向编程原则（分离代码与数据、使用通用数据结构等）。

## 7. 数据库规范
- 默认使用 InnoDB 存储引擎。
- 使用 UTF-8 字符集。
- 表名、列名使用小写下划线命名法。
- 每张表必须定义主键。
- 约束命名规则：
  - 主键：`pk_表名`
  - 外键：`fk_表名_引用表名`
  - 唯一键：`uk_表名_字段名`
  - 索引：`idx_表名_字段名`
- 标准列（推荐）：
  - `id`: bigint, 主键
  - `create_time`: datetime
  - `update_time`: datetime
  - `create_by`: varchar
  - `update_by`: varchar
  - `is_deleted`: tinyint (逻辑删除标志, 0:未删除, 1:已删除)

## 8. Redis 规范
- 键命名：`{service}:{entity}:{id}`，例如 `user:info:123`。
- 设置合理的 TTL (Time-To-Live)，避免内存泄漏。
- 分布式锁场景推荐使用 Redisson 或其他成熟的 Redis 锁实现。
- 处理好连接异常及重连机制。

## 9. Kafka 规范
- 主题命名：`{环境}.{服务}.{实体}.{操作}`，例如 `prod.order.order.created`。
- 消息生产端配置幂等性与合理的重试机制。
- 消费端关注偏移量提交策略（自动/手动）与错误处理。
- 配置死信队列 (DLQ) 处理无法消费的消息。

## 10. 技术栈与错误处理

## 11. 常用库推荐

### 开发框架
- **Spring生态**: Spring Boot, Spring MVC, Spring Data, Spring Security, Spring Cloud
- **微服务框架**: Spring Cloud Alibaba, Dubbo, gRPC
- **Web框架**: Micronaut, Quarkus, Vert.x

### 工具类库
- **Apache Commons**: Commons Lang, Commons IO, Commons Collections
- **Google Guava**: 集合、缓存、并发等工具
- **Lombok**: 简化Java代码
- **MapStruct**: 对象映射工具

### 测试框架
- **JUnit 5**: 单元测试框架
- **Mockito**: Mock测试框架
- **Testcontainers**: 集成测试工具
- **AssertJ**: 流式断言库

### 数据库
- **JPA实现**: Hibernate, EclipseLink
- **MyBatis**: MyBatis, MyBatis-Plus
- **连接池**: HikariCP, Druid
- **NoSQL客户端**: Jedis, Lettuce, MongoDB Driver

### 消息队列
- **RabbitMQ**: AMQP协议实现
- **Kafka**: 高吞吐量消息系统
- **RocketMQ**: 阿里开源消息中间件

### 监控与指标
- **Micrometer**: 应用指标收集
- **Prometheus**: 监控系统
- **SkyWalking**: 分布式追踪系统

### 其他
- **Jackson/Gson**: JSON处理
- **Joda-Time**: 日期时间处理
- **Netty**: 高性能网络框架
- **OkHttp**: HTTP客户端

## 12. Spring Boot 最佳实践

#### 依赖注入
- 优先使用构造器注入，避免字段注入
- 使用`@RequiredArgsConstructor`简化构造器注入
- 避免循环依赖，保持组件职责单一

#### 异常处理
- 使用`@ControllerAdvice`和`@ExceptionHandler`进行全局异常处理
- 自定义业务异常继承`RuntimeException`
- 异常信息国际化处理
- REST API返回统一错误格式

#### 配置管理
- 使用`@ConfigurationProperties`绑定配置
- 敏感信息使用环境变量或配置中心
- 多环境配置：`application-{profile}.yml`
- 配置验证：使用 `@Validated` 注解。
- 使用 Spring Boot Actuator 进行应用监控与指标暴露。

#### 日志管理
- 使用SLF4J + Logback
- 合理配置日志级别(ERROR/WARN/INFO/DEBUG)
- 添加请求ID等上下文信息
- 配置日志滚动策略和保留周期

#### 监控与指标
- 使用Spring Boot Actuator暴露健康检查、指标等端点
- 集成Prometheus收集应用指标
- 使用 Micrometer 实现自定义指标。
- 异步处理：使用 Spring 的 `@Async` 或响应式编程 (Spring WebFlux)。

#### 数据库访问
- 使用Spring Data JPA或MyBatis-Plus简化CRUD
- 合理使用事务`@Transactional`
- 批量操作使用批处理模式
- 实现读写分离（如果需要）。
- 使用 DTO (Data Transfer Object) 封装 Controller 层的请求和响应。

#### 缓存
- 使用Spring Cache抽象
- 合理设置缓存过期时间
- 注意缓存穿透、雪崩问题
- 重要操作实现双写一致性

#### 消息队列
- 使用Spring AMQP或Spring Kafka
- 实现消息幂等处理
- 死信队列处理失败消息
- 消息轨迹追踪

#### 安全
- 使用Spring Security实现认证授权
- 密码加密存储
- 接口权限控制
- CSRF防护
- 定期更新依赖库版本。
- 输入数据校验，防止 XSS/SQL 注入。
- 生产环境强制使用 HTTPS。
- 密码使用安全哈希算法存储 (如 BCrypt)。
- 遵循最小权限原则。
- 必要时实现 CORS 配置。

## 13. 测试规范
- 单元测试覆盖率建议不低于 70%。
- 单元测试重点关注 Service 层和 Util 类。
- 集成测试覆盖 Controller 层。
- 使用 JUnit 5 + Mockito + Spring Boot Test 进行测试。
- 使用 MockMvc 测试 Controller。
- 考虑使用 H2 等内存数据库进行测试。

## 14. AI 代码标注规范

### 完整文件标注
在文件顶部添加 Javadoc 注释块：
```java
/**
 * [AI-ASSISTED]
 * 生成工具: {工具名称} {版本}
 * 生成日期: {YYYY-MM-DD}
 * 贡献程度: {完全生成|部分生成|辅助编写|重构优化}
 * 人工修改: {无|轻微|中度|大量}
 * 责任人: {开发者姓名}
 */
```

### 代码块标注
对于 AI 生成或修改的代码块，使用注释标记起止：
```java
/* [AI-BLOCK-START] - 生成工具: {工具名称} {版本} */
String generatedCode = "AI generated code";
/* [AI-BLOCK-END] */
```

### 单行标注
对于单行 AI 生成或修改的代码，在行尾添加注释：
```java
result = complexCalculation(); // [AI-LINE: {工具名称} {版本}]
```

## 15. Git 规范
- 分支命名: `feature/ai-{feature_name}` 或遵循项目自身规范并体现 AI 辅助。
- Commit Message: `[AI-ASSISTED] {commit_message}` 或类似标识。
- Pull Request 标题: `[AI] {pr_title}` 或类似标识。

## 16. 性能优化
- 大数据量查询使用分页。
- 常用数据、计算结果适当使用缓存 (如 Spring Cache)。
- 数据库添加必要索引，优化 SQL 语句，避免全表扫描。
- 耗时任务考虑异步执行 (`@Async`)。
- 长耗时操作（如外部调用、文件处理）必须设置合理的超时机制。
- 资源密集型操作（高并发接口）必须实现限流和熔断机制 (如 Sentinel, Resilience4j)。
- 所有删除和修改等敏感操作考虑幂等性设计。
- API 响应根据需要设置合适的 HTTP 缓存控制头部 (Cache-Control)。

## 17. 监控与可观测性
- 实现健康检查接口 (`/actuator/health`)。
- 通过 Actuator 暴露应用指标 (`/actuator/prometheus`) 并集成 Prometheus 进行监控。
- 配置合理的日志告警规则，及时发现和响应异常。
- 考虑引入分布式追踪系统 (如 SkyWalking, Zipkin) 追踪请求链路。

## 18. 开发指南摘要 (源自 Effective Java)

### 创建和销毁对象
- 考虑静态工厂方法代替构造器。
- 多参数构造器考虑 Builder 模式。
- 使用私有构造器或枚举实现 Singleton。
- 避免创建不必要的对象。
- 优先使用 try-with-resources。

### 通用方法
- 重写 `equals` 时必须重写 `hashCode`。
- 始终重写 `toString`。
- 谨慎实现 `Cloneable`。
- 考虑实现 `Comparable`。

### 类和接口
- 最小化类和成员的可访问性。
- 公有类中使用访问方法而非公有域。
- 最小化可变性。
- 优先使用组合而非继承。
- 接口优于抽象类。

### 泛型
- 不要使用原始类型。
- 消除未检查的警告。
- 列表优于数组。
- 优先使用泛型方法。
- 利用有限制通配符提升 API 灵活性。

### Lambda 和 Stream
- Lambda 优于匿名类。
- 方法引用优于 Lambda。
- 优先使用标准函数式接口。
- 明智地使用 Stream。
- Stream 中优先使用无副作用的函数。

### 方法
- 检查参数有效性。
- 必要时进行防御性拷贝。
- 谨慎设计方法签名。
- 返回空集合或数组而非 null。
- 明智地返回 Optional。

### 通用编程
- 最小化局部变量作用域。
- for-each 优于传统 for 循环。
- 了解并使用库。
- 需要精确计算时避免使用 float 和 double。
- 接口优于反射。
- 遵循通用命名约定。

### 异常
- 只在异常情况下使用异常。
- 受检异常用于可恢复情况，运行时异常用于编程错误。
- 优先使用标准异常。
- 为异常提供详尽的失败捕获信息。
- 不要忽略异常。

### 并发
- 同步访问共享可变数据。
- 避免过度同步。
- Executor、Task、Stream 优于 Thread。
- 并发工具优于 `wait` 和 `notify`。
- 谨慎使用延迟初始化。