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
- [严重]  [  对于可重试的操作是否实现了重试机制
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

## 11. 中间件技术检查点

### 11.1 Tomcat配置
- [严重] 线程池配置是否合理（核心线程数、最大线程数、队列容量）
- [严重] 连接器配置是否优化（端口、协议、超时设置）
- [重要] 内存配置是否合理（JVM参数调优）
- [重要] 日志配置是否完善（级别、轮转策略、日志位置）
- [重要] 安全配置是否到位（禁用不必要的HTTP方法、删除默认应用）
- [次要] 是否配置了优雅停机策略

```java
// 示例：Tomcat线程池配置
@Bean
public ServletWebServerFactory servletContainer() {
    TomcatServletWebServerFactory tomcat = new TomcatServletWebServerFactory();
    tomcat.addConnectorCustomizers(connector -> {
        ThreadPoolExecutor executor = (ThreadPoolExecutor) connector.getProtocolHandler().getExecutor();
        executor.setCorePoolSize(10);
        executor.setMaxPoolSize(200);
        executor.setQueueCapacity(100);
    });
    return tomcat;
}

// 示例：Tomcat连接器配置
@Bean
public WebServerFactoryCustomizer<TomcatServletWebServerFactory> tomcatCustomizer() {
    return (factory) -> {
        factory.addConnectorCustomizers(connector -> {
            connector.setPort(8080);
            connector.setProperty("connectionTimeout", "20000");
            connector.setProperty("relaxedQueryChars", "[]|{}^&#x5c;&#x60;&quot;&lt;&gt;");
            connector.setProperty("maxKeepAliveRequests", "100");
            connector.setProperty("keepAliveTimeout", "60000");
        });
    };
}
```

### 11.2 Spring Boot配置
- [严重] 是否理解并正确使用自动配置机制
- [严重] Profile配置是否合理（环境隔离、配置分离）
- [重要] 是否实现了自定义健康检查指标
- [重要] Actuator端点是否安全配置
- [重要] 是否实现了配置外部化（配置服务器整合）
- [次要] 是否合理使用Spring缓存抽象

```java
// 示例：Spring Boot Profile和健康检查配置
@Configuration
@Profile("production")
public class ProductionConfig {
    
    @Bean
    public HealthIndicator databaseHealthIndicator(DataSource dataSource) {
        return new DataSourceHealthIndicator(dataSource, "SELECT 1 FROM DUAL");
    }
    
    @Bean
    public HealthIndicator redisHealthIndicator(RedisConnectionFactory redisConnectionFactory) {
        return new RedisHealthIndicator(redisConnectionFactory);
    }
}

// 示例：Actuator端点安全配置
@Configuration
public class ActuatorSecurityConfig extends WebSecurityConfigurerAdapter {
    
    @Override
    protected void configure(HttpSecurity http) throws Exception {
        http.requestMatcher(EndpointRequest.toAnyEndpoint())
            .authorizeRequests()
            .requestMatchers(EndpointRequest.to("health", "info")).permitAll()
            .anyRequest().hasRole("ACTUATOR_ADMIN")
            .and()
            .httpBasic();
    }
}
```

### 11.3 Redis配置
- [严重] 连接池配置是否合理（连接数、超时设置）
- [严重] 序列化方式是否安全高效（JSON/MessagePack/自定义）
- [重要] 是否设置了合理的缓存过期策略
- [重要] 是否实现了缓存穿透/击穿/雪崩防护措施
- [重要] 是否监控Redis性能指标（命中率、延迟、内存使用）
- [次要] 集群配置是否合理（主从复制、哨兵模式、集群模式）

```java
// 示例：Redis连接池和序列化配置
@Bean
public RedisConnectionFactory redisConnectionFactory() {
    LettuceConnectionFactory factory = new LettuceConnectionFactory();
    factory.setHostName("redis-host");
    factory.setPort(6379);
    factory.setPassword("password");
    factory.setDatabase(0);
    
    // 连接池配置
    LettucePoolingClientConfiguration.builder()
        .commandTimeout(Duration.ofMillis(500))
        .poolConfig(new GenericObjectPoolConfig<>() {{
            setMaxTotal(50);
            setMaxIdle(20);
            setMinIdle(5);
            setTestOnBorrow(true);
            setTestOnReturn(true);
            setTestWhileIdle(true);
        }})
        .build();
    
    return factory;
}

@Bean
public RedisTemplate<String, Object> redisTemplate(RedisConnectionFactory connectionFactory) {
    RedisTemplate<String, Object> template = new RedisTemplate<>();
    template.setConnectionFactory(connectionFactory);
    
    // 使用JSON序列化
    Jackson2JsonRedisSerializer<Object> serializer = new Jackson2JsonRedisSerializer<>(Object.class);
    template.setKeySerializer(new StringRedisSerializer());
    template.setValueSerializer(serializer);
    template.setHashKeySerializer(new StringRedisSerializer());
    template.setHashValueSerializer(serializer);
    
    return template;
}

// 示例：缓存击穿防护 - 使用分布式锁
public String getDataWithLock(String key) {
    String value = redisTemplate.opsForValue().get(key);
    if (value == null) {
        String lockKey = "lock:" + key;
        Boolean acquired = redisTemplate.opsForValue().setIfAbsent(lockKey, "1", Duration.ofSeconds(10));
        if (Boolean.TRUE.equals(acquired)) {
            try {
                // 双重检查
                value = redisTemplate.opsForValue().get(key);
                if (value == null) {
                    // 从数据库获取数据
                    value = loadFromDb(key);
                    redisTemplate.opsForValue().set(key, value, Duration.ofHours(1));
                }
            } finally {
                redisTemplate.delete(lockKey);
            }
        } else {
            // 等待一会重试
            Thread.sleep(50);
            return getDataWithLock(key);
        }
    }
    return value;
}
```

### 11.4 Kafka配置
- [严重] 生产者配置是否保证可靠性（幂等性、确认机制）
- [严重] 消费者配置是否合理（消费组、偏移量管理、并行处理）
- [重要] Topic设计是否合理（分区、副本因子、保留策略）
- [重要] 是否实现了完善的重试与错误处理机制（死信队列）
- [重要] 是否配置了消息监控指标（延迟、生产/消费速率）
- [次要] 序列化和反序列化方式是否高效（Avro/Protobuf/JSON）

```java
// 示例：Kafka生产者配置
@Bean
public ProducerFactory<String, MyEvent> producerFactory() {
    Map<String, Object> props = new HashMap<>();
    props.put(ProducerConfig.BOOTSTRAP_SERVERS_CONFIG, "kafka:9092");
    props.put(ProducerConfig.KEY_SERIALIZER_CLASS_CONFIG, StringSerializer.class);
    props.put(ProducerConfig.VALUE_SERIALIZER_CLASS_CONFIG, JsonSerializer.class);
    
    // 可靠性配置
    props.put(ProducerConfig.ACKS_CONFIG, "all");
    props.put(ProducerConfig.RETRIES_CONFIG, 3);
    props.put(ProducerConfig.ENABLE_IDEMPOTENCE_CONFIG, true);
    
    // 性能配置
    props.put(ProducerConfig.BATCH_SIZE_CONFIG, 16384);
    props.put(ProducerConfig.LINGER_MS_CONFIG, 10);
    props.put(ProducerConfig.COMPRESSION_TYPE_CONFIG, "snappy");
    
    return new DefaultKafkaProducerFactory<>(props);
}

// 示例：Kafka消费者配置
@Bean
public ConsumerFactory<String, MyEvent> consumerFactory() {
    Map<String, Object> props = new HashMap<>();
    props.put(ConsumerConfig.BOOTSTRAP_SERVERS_CONFIG, "kafka:9092");
    props.put(ConsumerConfig.GROUP_ID_CONFIG, "my-consumer-group");
    props.put(ConsumerConfig.KEY_DESERIALIZER_CLASS_CONFIG, StringDeserializer.class);
    props.put(ConsumerConfig.VALUE_DESERIALIZER_CLASS_CONFIG, JsonDeserializer.class);
    
    // 可靠性配置
    props.put(ConsumerConfig.AUTO_OFFSET_RESET_CONFIG, "earliest");
    props.put(ConsumerConfig.ENABLE_AUTO_COMMIT_CONFIG, false);
    props.put(ConsumerConfig.MAX_POLL_RECORDS_CONFIG, 500);
    props.put(ConsumerConfig.MAX_POLL_INTERVAL_MS_CONFIG, 300000);
    
    return new DefaultKafkaConsumerFactory<>(props);
}

// 示例：死信队列与重试配置
@Bean
public KafkaListenerContainerFactory<?> kafkaListenerContainerFactory(
        ConsumerFactory<String, MyEvent> consumerFactory) {
    ConcurrentKafkaListenerContainerFactory<String, MyEvent> factory = 
        new ConcurrentKafkaListenerContainerFactory<>();
    factory.setConsumerFactory(consumerFactory);
    factory.setConcurrency(3);
    factory.getContainerProperties().setAckMode(AckMode.MANUAL_IMMEDIATE);
    
    // 设置重试与死信
    factory.setErrorHandler(new SeekToCurrentErrorHandler(
        new DeadLetterPublishingRecoverer(kafkaTemplate, 
            (r, e) -> new TopicPartition("my-topic.DLT", r.partition())),
        new FixedBackOff(1000L, 3))); // 重试3次，间隔1秒
    
    return factory;
}
```

### 11.5 Elasticsearch配置
- [严重] 索引设计是否合理（映射、分析器、字段类型）
- [严重] 查询是否优化（过滤器、聚合、分页）
- [重要] 连接池配置是否合理（连接数、超时设置）
- [重要] 是否使用批量操作提升性能
- [重要] 分片与副本策略是否合理（分片数量、副本数）
- [次要] 是否监控ES集群状态和性能

```java
// 示例：Elasticsearch客户端配置
@Bean
public RestHighLevelClient elasticsearchClient() {
    ClientConfiguration clientConfiguration = ClientConfiguration.builder()
        .connectedTo("elasticsearch:9200")
        .withConnectTimeout(Duration.ofSeconds(5))
        .withSocketTimeout(Duration.ofSeconds(30))
        .withBasicAuth("username", "password")
        .build();
    
    return RestClients.create(clientConfiguration).rest();
}

// 示例：索引创建与映射设置
public void createProductIndex() {
    try {
        CreateIndexRequest request = new CreateIndexRequest("products");
        
        // 设置分片和副本
        request.settings(Settings.builder()
            .put("index.number_of_shards", 3)
            .put("index.number_of_replicas", 2)
            .put("index.refresh_interval", "1s")
        );
        
        // 定义映射
        Map<String, Object> properties = new HashMap<>();
        
        Map<String, Object> nameField = new HashMap<>();
        nameField.put("type", "text");
        nameField.put("analyzer", "ik_max_word");
        nameField.put("search_analyzer", "ik_smart");
        
        Map<String, Object> priceField = new HashMap<>();
        priceField.put("type", "double");
        
        Map<String, Object> createTimeField = new HashMap<>();
        createTimeField.put("type", "date");
        createTimeField.put("format", "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis");
        
        properties.put("name", nameField);
        properties.put("price", priceField);
        properties.put("createTime", createTimeField);
        
        Map<String, Object> mapping = new HashMap<>();
        mapping.put("properties", properties);
        
        request.mapping(mapping);
        
        elasticsearchClient().indices().create(request, RequestOptions.DEFAULT);
    } catch (IOException e) {
        throw new RuntimeException("Failed to create product index", e);
    }
}

// 示例：批量操作
public void bulkIndexProducts(List<Product> products) {
    BulkRequest bulkRequest = new BulkRequest();
    for (Product product : products) {
        IndexRequest indexRequest = new IndexRequest("products")
            .id(product.getId().toString())
            .source(convertProductToMap(product));
        bulkRequest.add(indexRequest);
    }
    
    try {
        BulkResponse bulkResponse = elasticsearchClient().bulk(bulkRequest, RequestOptions.DEFAULT);
        if (bulkResponse.hasFailures()) {
            log.error("Bulk indexing has failures: {}", bulkResponse.buildFailureMessage());
        }
    } catch (IOException e) {
        throw new RuntimeException("Failed to bulk index products", e);
    }
}
```

### 11.6 MySQL配置
- [严重] 连接池配置是否合理（最大连接数、超时设置）
- [严重] 事务管理是否正确（传播属性、隔离级别、超时）
- [重要] 索引设计是否优化（主键、索引选择、联合索引顺序）
- [重要] 是否合理使用批量操作提升性能
- [重要] 是否配置了慢查询监控
- [次要] 字段类型选择是否合理（数据类型、长度、字符集）

```java
// 示例：MySQL连接池配置
@Bean
public DataSource dataSource() {
    HikariConfig config = new HikariConfig();
    
    // 基本连接配置
    config.setJdbcUrl("jdbc:mysql://localhost:3306/mydb?useSSL=false&serverTimezone=UTC&characterEncoding=UTF-8");
    config.setUsername("dbuser");
    config.setPassword("dbpass");
    
    // 连接池设置
    config.setMinimumIdle(10);                   // 最小空闲连接数
    config.setMaximumPoolSize(50);               // 最大连接数
    config.setIdleTimeout(600000);               // 连接最大空闲时间（毫秒）
    config.setMaxLifetime(1800000);              // 连接最大生命周期（毫秒）
    config.setConnectionTimeout(30000);          // 获取连接超时时间（毫秒）
    
    // 连接测试
    config.setConnectionTestQuery("SELECT 1");   // 连接测试查询
    config.setValidationTimeout(5000);           // 验证超时时间（毫秒）
    
    // 泄漏检测
    config.setLeakDetectionThreshold(60000);     // 连接泄漏检测阈值
    
    return new HikariDataSource(config);
}

// 示例：事务管理配置
@Bean
public PlatformTransactionManager transactionManager(DataSource dataSource) {
    DataSourceTransactionManager txManager = new DataSourceTransactionManager(dataSource);
    // 设置默认超时（秒）
    txManager.setDefaultTimeout(30);
    return txManager;
}

// 示例：事务使用
@Service
public class UserService {
    
    private final UserRepository userRepository;
    
    public UserService(UserRepository userRepository) {
        this.userRepository = userRepository;
    }
    
    // 使用声明式事务，指定传播行为和隔离级别
    @Transactional(
        propagation = Propagation.REQUIRED,
        isolation = Isolation.READ_COMMITTED,
        timeout = 10,
        rollbackFor = Exception.class,
        noRollbackFor = NotFoundException.class
    )
    public User createUser(UserDto userDto) {
        // 业务逻辑
        return userRepository.save(new User(userDto));
    }
    
    // 批量操作示例
    @Transactional
    public void batchCreateUsers(List<UserDto> userDtos) {
        List<User> users = userDtos.stream()
            .map(User::new)
            .collect(Collectors.toList());
            
        // 分批处理，每批500条
        Lists.partition(users, 500).forEach(batch -> {
            userRepository.saveAll(batch);
        });
    }
}
```

## 结语

本检查清单不是一成不变的，应随着项目的发展和技术的进步不断更新。代码审查的目的不仅是发现问题，更是促进团队成长和知识共享。建议将本清单与自动化工具(如SonarQube、CheckStyle、SpotBugs等)结合使用，以提高代码审查的效率和质量。

---

*本文档由Java代码审查专家创建，旨在提升Java微服务代码质量。* 