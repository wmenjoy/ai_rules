# Java微服务代码评审检查清单

## 前言

本检查清单旨在提供一套系统化的Java微服务代码评审标准，帮助团队提升代码质量、可维护性和性能。每个检查点按严重程度分为三个等级：
- **严重(Critical)**: 必须修复，可能导致安全漏洞、性能问题或系统不稳定
- **重要(Major)**: 应当修复，影响代码质量和可维护性
- **次要(Minor)**: 建议修复，主要是代码风格和最佳实践

## 1. 代码结构与组织

### 1.1 架构设计
- [严重] 服务边界划分是否清晰，遵循单一职责原则
- [重要] 是否采用领域驱动设计(DDD)思想组织代码结构
- [重要] 控制器(Controller)、服务(Service)、数据访问(Repository)职责是否分离
- [次要] 是否合理使用设计模式解决特定问题

### 1.2 代码规范
- [重要] 代码是否符合Oracle Java编码规范
- [重要] 命名是否规范且有意义(类名、方法名、变量名)
- [重要] 是否避免硬编码常量，使用配置或常量类
- [次要] 代码格式是否一致(缩进、空行、括号位置等)

### 1.3 模块化与依赖
- [严重] 是否避免循环依赖
- [重要] 是否合理使用接口和抽象类
- [重要] 第三方依赖是否明确版本且避免冲突
- [次要] 是否遵循包导入顺序约定

## 2. 线程安全与并发处理

### 2.1 线程池配置
- [严重] 线程池核心线程数、最大线程数、队列容量配置是否合理
- [严重] 是否配置了拒绝策略
- [重要] 线程池是否有名称，便于监控和调试
- [重要] 线程池是否配置了合理的关闭策略

```java
// 示例：合理的线程池配置
ThreadPoolExecutor executor = new ThreadPoolExecutor(
    10,                          // 核心线程数
    20,                          // 最大线程数
    60, TimeUnit.SECONDS,        // 空闲线程存活时间
    new LinkedBlockingQueue<>(100), // 队列容量
    new ThreadFactoryBuilder().setNameFormat("service-pool-%d").build(), // 线程命名
    new ThreadPoolExecutor.CallerRunsPolicy()); // 拒绝策略
```

### 2.2 共享状态与线程安全
- [严重] 共享的实例变量是否保证了线程安全
- [严重] 是否正确使用了并发集合类
- [重要] 是否避免使用过度同步，导致性能下降
- [重要] 是否考虑了可见性、有序性和原子性问题

```java
// 反例：线程不安全的单例模式
public class UnsafeSingleton {
    private static UnsafeSingleton instance;
    
    public static UnsafeSingleton getInstance() {
        if (instance == null) {
            instance = new UnsafeSingleton(); // 线程不安全
        }
        return instance;
    }
}
```

### 2.3 并发控制
- [严重] 是否避免了死锁、活锁和饥饿问题
- [重要] 是否正确使用了并发工具类(ConcurrentHashMap, AtomicInteger等)
- [重要] 是否合理使用CompletableFuture进行异步操作
- [次要] 是否考虑使用不可变对象减少并发问题

## 3. 资源管理

### 3.1 连接池配置
- [严重] 数据库连接池配置是否合理(大小、超时、检测)
- [严重] HTTP客户端连接池配置是否合理
- [重要] 是否监控连接池状态
- [重要] 是否有连接泄漏检测机制

```java
// 示例：合理的数据库连接池配置
@Bean
public DataSource dataSource() {
    HikariConfig config = new HikariConfig();
    config.setJdbcUrl("jdbc:mysql://localhost:3306/db");
    config.setUsername("user");
    config.setPassword("pass");
    config.setMaximumPoolSize(20);
    config.setMinimumIdle(5);
    config.setIdleTimeout(300000);
    config.setConnectionTimeout(10000);
    config.setMaxLifetime(1800000);
    config.setLeakDetectionThreshold(60000);
    return new HikariDataSource(config);
}
```

### 3.2 超时配置
- [严重] 所有外部服务调用是否设置了合理的超时时间
- [严重] HTTP请求是否配置了连接超时和读取超时
- [重要] 定时任务是否有执行超时控制
- [重要] 长事务是否有超时控制

```java
// 示例：RestTemplate超时配置
@Bean
public RestTemplate restTemplate() {
    SimpleClientHttpRequestFactory factory = new SimpleClientHttpRequestFactory();
    factory.setConnectTimeout(3000); // 连接超时3秒
    factory.setReadTimeout(5000);    // 读取超时5秒
    return new RestTemplate(factory);
}
```

### 3.3 资源释放
- [严重] 是否使用try-with-resources确保资源关闭
- [严重] 自定义的资源是否在finally块中关闭
- [重要] 是否避免资源泄漏(文件句柄、数据库连接等)
- [次要] 是否实现了AutoCloseable接口以便资源自动关闭

## 4. 异常处理

### 4.1 异常设计
- [严重] 是否定义了合理的异常体系，区分业务异常和系统异常
- [重要] 是否避免捕获异常后不处理或仅打印日志
- [重要] 是否避免catch Exception等过于宽泛的异常
- [次要] 自定义异常是否包含足够的上下文信息

### 4.2 异常传播
- [严重] 是否实现了全局异常处理机制
- [重要] 异常信息是否对用户友好，不暴露敏感信息
- [重要] 是否记录了异常的完整堆栈信息
- [次要] 是否合理使用受检异常和运行时异常

```java
// 示例：全局异常处理
@RestControllerAdvice
public class GlobalExceptionHandler {
    private static final Logger log = LoggerFactory.getLogger(GlobalExceptionHandler.class);
    
    @ExceptionHandler(BusinessException.class)
    public ResponseEntity<ErrorResponse> handleBusinessException(BusinessException ex) {
        log.warn("业务异常: {}", ex.getMessage());
        return ResponseEntity.badRequest().body(new ErrorResponse(ex.getCode(), ex.getMessage()));
    }
    
    @ExceptionHandler(Exception.class)
    public ResponseEntity<ErrorResponse> handleUnknownException(Exception ex) {
        log.error("系统异常", ex);
        return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR)
            .body(new ErrorResponse("SYSTEM_ERROR", "系统异常，请稍后再试"));
    }
}
```

### 4.3 重试与降级
- [严重] 对于可重试的操作是否实现了重试机制
- [重要] 是否实现了服务降级策略
- [重要] 是否使用熔断器模式防止级联失败
- [次要] 重试次数和间隔是否可配置

## 5. 日志与监控

### 5.1 日志实践
- [严重] 是否避免在生产环境使用System.out或e.printStackTrace()
- [重要] 是否使用结构化日志（包含请求ID、用户ID等上下文信息）
- [重要] 是否使用合适的日志级别(ERROR/WARN/INFO/DEBUG)
- [重要] 是否避免记录敏感信息(密码、令牌等)

```java
// 示例：良好的日志实践
// 反例
System.out.println("User logged in: " + username);
// 正例
log.info("User logged in: {}, requestId: {}", username, MDC.get("requestId"));
```

### 5.2 监控指标
- [严重] 是否集成了指标收集框架(如Micrometer)
- [严重] 是否监控关键业务指标
- [重要] 是否监控JVM指标(内存、GC、线程等)
- [重要] 是否监控外部依赖的状态和性能

```java
// 示例：监控关键业务操作
@Component
public class OrderService {
    private final MeterRegistry registry;
    
    public OrderService(MeterRegistry registry) {
        this.registry = registry;
    }
    
    public void createOrder(Order order) {
        Timer.Sample sample = Timer.start(registry);
        try {
            // 处理订单逻辑
            registry.counter("orders.created").increment();
        } catch (Exception e) {
            registry.counter("orders.failed").increment();
            throw e;
        } finally {
            sample.stop(registry.timer("orders.processing.time"));
        }
    }
}
```

### 5.3 健康检查
- [严重] 是否实现了健康检查接口
- [严重] 健康检查是否包含关键依赖状态
- [重要] 是否设置了合理的健康检查超时
- [次要] 健康检查报告是否包含详细信息

## 6. 安全

### 6.1 输入验证
- [严重] 是否对所有外部输入进行验证
- [严重] 是否防范SQL注入攻击
- [严重] 是否防范XSS攻击
- [重要] 是否使用参数化查询而非字符串拼接

### 6.2 认证与授权
- [严重] 是否正确实现了认证机制
- [严重] 是否实现了细粒度的授权控制
- [严重] 敏感信息是否加密存储
- [重要] 是否避免在URL中包含敏感信息

### 6.3 通信安全
- [严重] 是否使用HTTPS进行通信
- [严重] 是否正确验证SSL证书
- [重要] 是否实现了CSRF防护
- [重要] 是否设置了安全HTTP头部

## 7. 性能优化

### 7.1 数据库优化
- [严重] SQL查询是否优化，有无不必要的连接或子查询
- [严重] 是否建立了必要的索引
- [重要] 是否实现了分页查询，避免大结果集
- [重要] 是否使用批处理进行批量操作

### 7.2 缓存策略
- [严重] 是否合理使用缓存减少重复计算和IO
- [重要] 缓存过期策略是否合理
- [重要] 缓存更新机制是否安全（防止缓存击穿、缓存雪崩）
- [次要] 是否监控缓存命中率

```java
// 示例：使用Spring缓存
@Service
public class ProductService {
    
    @Cacheable(value = "products", key = "#id", unless = "#result == null")
    public Product getProductById(Long id) {
        // 从数据库获取产品
        return productRepository.findById(id).orElse(null);
    }
    
    @CacheEvict(value = "products", key = "#product.id")
    public void updateProduct(Product product) {
        productRepository.save(product);
    }
}
```

### 7.3 异步处理
- [重要] 是否将耗时操作异步化
- [重要] 是否使用消息队列处理峰值负载
- [重要] 是否实现了请求合并减少调用次数
- [次要] 是否实现了结果缓存避免重复计算

## 8. 测试

### 8.1 单元测试
- [严重] 核心业务逻辑是否有单元测试覆盖
- [重要] 是否测试了边界条件和异常路径
- [重要] 测试是否独立且可重复
- [次要] 测试代码是否遵循AAA(Arrange-Act-Assert)模式

### 8.2 集成测试
- [严重] 是否验证了与外部系统的集成
- [重要] 是否使用适当的mock/stub替代外部依赖
- [重要] 是否测试了并发场景
- [次要] 是否进行了性能测试和负载测试

## 9. 微服务特定检查点

### 9.1 服务通信
- [严重] 服务间通信是否实现了熔断、超时和重试机制
- [严重] 是否实现了服务发现机制
- [重要] API契约是否明确且版本化
- [重要] 是否考虑了向后兼容性

### 9.2 数据一致性
- [严重] 是否正确处理分布式事务
- [重要] 是否使用了补偿事务或最终一致性模式
- [重要] 是否实现了幂等性API
- [次要] 是否有数据恢复机制

### 9.3 配置管理
- [严重] 敏感配置是否加密或使用外部配置服务
- [重要] 不同环境的配置是否分离
- [重要] 是否支持配置热更新
- [次要] 配置是否有文档说明

## 10. DevOps与部署

### 10.1 容器化
- [严重] Dockerfile是否合理优化（多阶段构建、最小镜像等）
- [重要] 容器是否以非root用户运行
- [重要] 是否暴露健康检查接口
- [次要] 是否设置了合理的资源限制

### 10.2 优雅停机
- [严重] 是否实现了优雅关闭机制
- [重要] 是否处理了SIGTERM信号
- [重要] 是否等待当前请求处理完成
- [重要] 是否正确关闭了连接池和其他资源

```java
// 示例：Spring Boot优雅停机配置
@Configuration
public class GracefulShutdownConfig {
    @Bean
    public GracefulShutdown gracefulShutdown() {
        return new GracefulShutdown();
    }
    
    @Bean
    public ServletWebServerFactory servletContainer(GracefulShutdown gracefulShutdown) {
        TomcatServletWebServerFactory factory = new TomcatServletWebServerFactory();
        factory.addConnectorCustomizers(gracefulShutdown);
        return factory;
    }
    
    private static class GracefulShutdown implements TomcatConnectorCustomizer, ApplicationListener<ContextClosedEvent> {
        private volatile Connector connector;
        
        @Override
        public void customize(Connector connector) {
            this.connector = connector;
        }
        
        @Override
        public void onApplicationEvent(ContextClosedEvent event) {
            this.connector.pause();
            Executor executor = this.connector.getProtocolHandler().getExecutor();
            if (executor instanceof ThreadPoolExecutor) {
                try {
                    ThreadPoolExecutor threadPoolExecutor = (ThreadPoolExecutor) executor;
                    threadPoolExecutor.shutdown();
                    if (!threadPoolExecutor.awaitTermination(30, TimeUnit.SECONDS)) {
                        log.warn("Tomcat线程池未在30秒内关闭，强制关闭");
                    }
                } catch (InterruptedException ex) {
                    Thread.currentThread().interrupt();
                }
            }
        }
    }
}
```

## 结语

本检查清单不是一成不变的，应随着项目的发展和技术的进步不断更新。代码审查的目的不仅是发现问题，更是促进团队成长和知识共享。建议将本清单与自动化工具(如SonarQube、CheckStyle、SpotBugs等)结合使用，以提高代码审查的效率和质量。

---

*本文档由Java代码审查专家创建，旨在提升Java微服务代码质量。* 