# Java微服务代码评审检查清单

## 概述
本检查清单基于Java微服务架构特点，从代码规范、安全性、性能、可观测性等多个维度提供系统化的代码评审指导。

## 检查分级说明
- **Critical**: 必须修复，影响系统安全或稳定性
- **Major**: 建议修复，影响代码质量或性能
- **Minor**: 可选修复，代码规范或可读性问题

## 1. 基础代码规范检查

### 1.1 命名规范 (Minor)

#### 1.1.1 类名使用大驼峰命名法
- **检查方法**: 使用IDE插件(如SonarLint)或静态分析工具检查类名格式
- **检查标准**: 类名首字母大写，后续单词首字母大写，不使用下划线或连字符
- **不正确实例**:
```java
// 错误示例
class user_service { }        // 使用下划线
class userservice { }         // 全小写
class User-Service { }        // 使用连字符
class userService { }         // 首字母小写

// 正确示例
class UserService { }
class OrderManagementService { }
```

#### 1.1.2 方法名和变量名使用小驼峰命名法
- **检查方法**: 代码审查时检查方法和变量命名，使用Checkstyle规则检查
- **检查标准**: 首字母小写，后续单词首字母大写，动词开头，语义明确
- **不正确实例**:
```java
// 错误示例
public void GetUser() { }           // 首字母大写
public void get_user() { }          // 使用下划线
public void getuser() { }           // 缺少驼峰
String user_name;                   // 使用下划线
int Count;                          // 首字母大写

// 正确示例
public void getUser() { }
public void calculateTotalAmount() { }
String userName;
int totalCount;
```

#### 1.1.3 常量使用全大写下划线分隔
- **检查方法**: 搜索final static变量，检查命名格式
- **检查标准**: 全部大写字母，单词间用下划线分隔，语义明确
- **不正确实例**:
```java
// 错误示例
public static final String maxCount = "100";     // 小驼峰
public static final String MAXCOUNT = "100";     // 缺少下划线
public static final String Max_Count = "100";    // 混合大小写

// 正确示例
public static final String MAX_COUNT = "100";
public static final int DEFAULT_TIMEOUT = 30000;
public static final String API_BASE_URL = "https://api.example.com";
```

#### 1.1.4 包名全小写，避免使用关键字
- **检查方法**: 检查package声明，验证包名规范
- **检查标准**: 全小写，使用点分隔，避免Java关键字，遵循域名倒置规则
- **不正确实例**:
```java
// 错误示例
package com.example.User;           // 包含大写字母
package com.example.class;          // 使用关键字
package com.example.user_service;   // 使用下划线

// 正确示例
package com.example.user;
package com.example.service.user;
package com.example.util.validation;
```

### 1.2 代码结构 (Major)

#### 1.2.1 单一职责原则，类和方法职责明确
- **检查方法**: 代码审查时分析类和方法的职责，使用SonarQube检查复杂度
- **检查标准**: 每个类只负责一个功能领域，每个方法只做一件事
- **不正确实例**:
```java
// 错误示例 - 违反单一职责原则
public class UserService {
    public void saveUser(User user) { }
    public void sendEmail(String email) { }     // 邮件发送不属于用户服务
    public void generateReport() { }            // 报表生成不属于用户服务
    public void validatePayment() { }           // 支付验证不属于用户服务
}

// 正确示例
public class UserService {
    public void saveUser(User user) { }
    public User findUserById(Long id) { }
    public void updateUser(User user) { }
}

public class EmailService {
    public void sendEmail(String email) { }
}
```

#### 1.2.2 避免过长的方法（建议不超过50行）
- **检查方法**: 使用IDE统计方法行数，设置代码规范检查工具阈值
- **检查标准**: 方法行数不超过50行，复杂逻辑拆分为多个小方法
- **不正确实例**:
```java
// 错误示例 - 方法过长
public void processOrder(Order order) {
    // 验证订单 (10行代码)
    if (order == null) { throw new IllegalArgumentException(); }
    // ... 更多验证逻辑
    
    // 计算价格 (15行代码)
    BigDecimal totalPrice = BigDecimal.ZERO;
    // ... 复杂计算逻辑
    
    // 库存检查 (10行代码)
    // ... 库存检查逻辑
    
    // 支付处理 (15行代码)
    // ... 支付处理逻辑
    
    // 发送通知 (10行代码)
    // ... 通知逻辑
}

// 正确示例
public void processOrder(Order order) {
    validateOrder(order);
    BigDecimal totalPrice = calculateTotalPrice(order);
    checkInventory(order);
    processPayment(order, totalPrice);
    sendNotification(order);
}
```

#### 1.2.3 避免过深的嵌套（建议不超过3层）
- **检查方法**: 代码审查时检查if/for/while嵌套层数
- **检查标准**: 嵌套层数不超过3层，使用早期返回减少嵌套
- **不正确实例**:
```java
// 错误示例 - 嵌套过深
public void processUser(User user) {
    if (user != null) {
        if (user.isActive()) {
            if (user.hasPermission()) {
                if (user.getAge() >= 18) {
                    if (user.getBalance() > 0) {
                        // 处理逻辑
                    }
                }
            }
        }
    }
}

// 正确示例 - 使用早期返回
public void processUser(User user) {
    if (user == null) return;
    if (!user.isActive()) return;
    if (!user.hasPermission()) return;
    if (user.getAge() < 18) return;
    if (user.getBalance() <= 0) return;
    
    // 处理逻辑
}
```

#### 1.2.4 合理使用设计模式
- **检查方法**: 代码审查时识别设计模式使用场景，检查是否过度设计
- **检查标准**: 根据实际需求选择合适的设计模式，避免过度设计
- **不正确实例**:
```java
// 错误示例 - 过度使用设计模式
// 简单的字符串处理却使用了策略模式
public interface StringProcessor {
    String process(String input);
}

public class UpperCaseProcessor implements StringProcessor {
    public String process(String input) { return input.toUpperCase(); }
}

// 对于简单场景，直接使用方法即可
public class StringUtils {
    public static String toUpperCase(String input) {
        return input.toUpperCase();
    }
}
```

## 2. 微服务特有检查

### 2.1 线程池配置 (Critical)

#### 2.1.1 避免使用 `Executors.newCachedThreadPool()`
- **检查方法**: 搜索代码中的`Executors.newCachedThreadPool()`调用
- **检查标准**: 禁止使用无界线程池，必须使用自定义ThreadPoolExecutor
- **不正确实例**:
```java
// 错误示例 - 使用无界线程池
ExecutorService executor = Executors.newCachedThreadPool();
ExecutorService executor2 = Executors.newFixedThreadPool(Integer.MAX_VALUE);

// 正确示例
ThreadPoolExecutor executor = new ThreadPoolExecutor(
    10,                          // 核心线程数
    20,                          // 最大线程数
    60L, TimeUnit.SECONDS,       // 空闲线程存活时间
    new LinkedBlockingQueue<>(100), // 有界队列
    new ThreadFactoryBuilder().setNameFormat("user-service-%d").build(),
    new ThreadPoolExecutor.CallerRunsPolicy() // 拒绝策略
);
```

#### 2.1.2 自定义线程池参数（核心线程数、最大线程数、队列大小）
- **检查方法**: 检查@Bean配置的ThreadPoolExecutor参数设置
- **检查标准**: 根据业务场景合理设置核心线程数、最大线程数和队列大小
- **不正确实例**:
```java
// 错误示例 - 参数设置不合理
@Bean
public ThreadPoolExecutor taskExecutor() {
    return new ThreadPoolExecutor(
        1,                              // 核心线程数过小
        Integer.MAX_VALUE,              // 最大线程数无限制
        0L, TimeUnit.MILLISECONDS,      // 空闲时间为0
        new LinkedBlockingQueue<>(),    // 无界队列
        Executors.defaultThreadFactory(),
        new ThreadPoolExecutor.AbortPolicy()
    );
}

// 正确示例
@Bean
public ThreadPoolExecutor taskExecutor() {
    int corePoolSize = Runtime.getRuntime().availableProcessors();
    return new ThreadPoolExecutor(
        corePoolSize,                    // 基于CPU核数
        corePoolSize * 2,               // 最大线程数为核心数的2倍
        60L, TimeUnit.SECONDS,          // 合理的空闲时间
        new LinkedBlockingQueue<>(200), // 有界队列
        new ThreadFactoryBuilder().setNameFormat("async-task-%d").build(),
        new ThreadPoolExecutor.CallerRunsPolicy()
    );
}
```

#### 2.1.3 设置有意义的线程名称
- **检查方法**: 检查ThreadFactory配置，确保线程名称包含业务含义
- **检查标准**: 线程名称应包含服务名称和功能描述，便于问题排查
- **不正确实例**:
```java
// 错误示例 - 线程名称无意义
ThreadPoolExecutor executor = new ThreadPoolExecutor(
    10, 20, 60L, TimeUnit.SECONDS,
    new LinkedBlockingQueue<>(100),
    Executors.defaultThreadFactory(),  // 使用默认线程工厂
    new ThreadPoolExecutor.CallerRunsPolicy()
);

// 或者
ThreadFactory factory = r -> new Thread(r, "thread"); // 名称过于简单

// 正确示例
ThreadFactory factory = new ThreadFactoryBuilder()
    .setNameFormat("user-service-async-%d")
    .setDaemon(true)
    .build();

ThreadPoolExecutor executor = new ThreadPoolExecutor(
    10, 20, 60L, TimeUnit.SECONDS,
    new LinkedBlockingQueue<>(100),
    factory,
    new ThreadPoolExecutor.CallerRunsPolicy()
);
```

#### 2.1.4 配置合适的拒绝策略
- **检查方法**: 检查ThreadPoolExecutor的RejectedExecutionHandler配置
- **检查标准**: 根据业务需求选择合适的拒绝策略，避免使用AbortPolicy导致任务丢失
- **不正确实例**:
```java
// 错误示例 - 使用AbortPolicy可能导致任务丢失
ThreadPoolExecutor executor = new ThreadPoolExecutor(
    10, 20, 60L, TimeUnit.SECONDS,
    new LinkedBlockingQueue<>(100),
    threadFactory,
    new ThreadPoolExecutor.AbortPolicy() // 直接抛异常，任务丢失
);

// 正确示例 - 根据场景选择策略
// 对于重要任务，使用CallerRunsPolicy
ThreadPoolExecutor executor1 = new ThreadPoolExecutor(
    10, 20, 60L, TimeUnit.SECONDS,
    new LinkedBlockingQueue<>(100),
    threadFactory,
    new ThreadPoolExecutor.CallerRunsPolicy() // 调用者线程执行
);

// 对于可丢弃任务，使用DiscardOldestPolicy
ThreadPoolExecutor executor2 = new ThreadPoolExecutor(
    10, 20, 60L, TimeUnit.SECONDS,
    new LinkedBlockingQueue<>(100),
    threadFactory,
    new ThreadPoolExecutor.DiscardOldestPolicy() // 丢弃最老任务
);
```

### 2.2 超时设置 (Critical)

#### 2.2.1 HTTP客户端设置连接超时和读取超时
- **检查方法**: 检查RestTemplate、WebClient、Feign客户端的超时配置
- **检查标准**: 必须设置连接超时和读取超时，避免无限等待
- **不正确实例**:
```java
// 错误示例 - 未设置超时
@Bean
public RestTemplate restTemplate() {
    return new RestTemplate(); // 使用默认配置，可能导致无限等待
}

@FeignClient(name = "user-service")
public interface UserClient {
    @GetMapping("/users/{id}")
    User getUser(@PathVariable Long id); // 未配置超时
}

// 正确示例
@Bean
public RestTemplate restTemplate() {
    HttpComponentsClientHttpRequestFactory factory = 
        new HttpComponentsClientHttpRequestFactory();
    factory.setConnectTimeout(5000);     // 连接超时5秒
    factory.setReadTimeout(10000);       // 读取超时10秒
    return new RestTemplate(factory);
}

@FeignClient(name = "user-service", 
    configuration = FeignConfig.class)
public interface UserClient {
    @GetMapping("/users/{id}")
    User getUser(@PathVariable Long id);
}

@Configuration
public class FeignConfig {
    @Bean
    public Request.Options options() {
        return new Request.Options(5000, 10000); // 连接超时5秒，读取超时10秒
    }
}
```

#### 2.2.2 数据库连接池超时配置
- **检查方法**: 检查数据源配置文件中的超时参数
- **检查标准**: 配置连接超时、查询超时、连接池获取超时等参数
- **不正确实例**:
```yaml
# 错误示例 - 缺少超时配置
spring:
  datasource:
    url: jdbc:mysql://localhost:3306/test
    username: root
    password: password
    hikari:
      maximum-pool-size: 20
      # 缺少超时配置

# 正确示例
spring:
  datasource:
    url: jdbc:mysql://localhost:3306/test?connectTimeout=5000&socketTimeout=30000
    username: root
    password: password
    hikari:
      maximum-pool-size: 20
      connection-timeout: 5000      # 连接超时5秒
      idle-timeout: 300000         # 空闲超时5分钟
      max-lifetime: 1800000        # 最大生命周期30分钟
      leak-detection-threshold: 60000 # 连接泄露检测阈值
```

#### 2.2.3 缓存操作超时设置
- **检查方法**: 检查Redis、Memcached等缓存客户端的超时配置
- **检查标准**: 设置合理的连接超时和操作超时，避免缓存故障影响主流程
- **不正确实例**:
```java
// 错误示例 - 未设置Redis超时
@Configuration
public class RedisConfig {
    @Bean
    public JedisConnectionFactory jedisConnectionFactory() {
        return new JedisConnectionFactory(); // 使用默认配置
    }
}

// 正确示例
@Configuration
public class RedisConfig {
    @Bean
    public JedisConnectionFactory jedisConnectionFactory() {
        JedisPoolConfig poolConfig = new JedisPoolConfig();
        poolConfig.setMaxTotal(20);
        poolConfig.setMaxIdle(10);
        poolConfig.setTestOnBorrow(true);
        
        JedisConnectionFactory factory = new JedisConnectionFactory(poolConfig);
        factory.setHostName("localhost");
        factory.setPort(6379);
        factory.setTimeout(2000);        // 操作超时2秒
        factory.setUsePool(true);
        return factory;
    }
    
    @Bean
    public RedisTemplate<String, Object> redisTemplate() {
        RedisTemplate<String, Object> template = new RedisTemplate<>();
        template.setConnectionFactory(jedisConnectionFactory());
        // 设置超时时间
        template.setDefaultSerializer(new GenericJackson2JsonRedisSerializer());
        return template;
    }
}
```

#### 2.2.4 消息队列消费超时配置
- **检查方法**: 检查RabbitMQ、Kafka等消息队列的消费者超时配置
- **检查标准**: 设置合理的消费超时时间，避免消息积压
- **不正确实例**:
```java
// 错误示例 - 未设置消费超时
@RabbitListener(queues = "user.queue")
public void handleMessage(String message) {
    // 可能长时间运行的处理逻辑
    processLongRunningTask(message);
}

// 正确示例
@RabbitListener(queues = "user.queue", 
    containerFactory = "rabbitListenerContainerFactory")
public void handleMessage(String message) {
    processLongRunningTask(message);
}

@Bean
public SimpleRabbitListenerContainerFactory rabbitListenerContainerFactory() {
    SimpleRabbitListenerContainerFactory factory = 
        new SimpleRabbitListenerContainerFactory();
    factory.setConnectionFactory(connectionFactory());
    factory.setConcurrentConsumers(3);
    factory.setMaxConcurrentConsumers(10);
    factory.setReceiveTimeout(30000L);    // 接收超时30秒
    factory.setAcknowledgeMode(AcknowledgeMode.MANUAL);
    return factory;
}
```

## 3. 并发和线程安全检查

### 3.1 共享变量线程安全 (Critical)

#### 3.1.1 Service类成员变量必须是无状态或线程安全的
- **检查方法**: 静态分析工具检查Service类成员变量，手动审查变量类型和使用方式
- **检查标准**: Service类应该是无状态的，如有成员变量必须是线程安全的或不可变的
- **不正确实例**:
```java
// 错误示例 - Service类包含可变状态
@Service
public class UserService {
    private List<User> userCache = new ArrayList<>();  // 非线程安全的List
    private int requestCount = 0;                      // 可变计数器
    private SimpleDateFormat dateFormat = new SimpleDateFormat("yyyy-MM-dd"); // 非线程安全
    
    public void addUser(User user) {
        userCache.add(user);        // 多线程访问会出现问题
        requestCount++;             // 非原子操作
    }
}

// 正确示例
@Service
public class UserService {
    private final UserRepository userRepository;  // 不可变依赖
    private static final DateTimeFormatter DATE_FORMATTER = 
        DateTimeFormatter.ofPattern("yyyy-MM-dd");  // 线程安全的格式化器
    
    public UserService(UserRepository userRepository) {
        this.userRepository = userRepository;
    }
    
    public void addUser(User user) {
        userRepository.save(user);  // 无状态操作
    }
}
```

#### 3.1.2 Controller类避免使用可变成员变量
- **检查方法**: 检查Controller类的成员变量声明，确保没有可变状态
- **检查标准**: Controller类应该是无状态的，只能有final的依赖注入字段
- **不正确实例**:
```java
// 错误示例 - Controller包含可变状态
@RestController
public class UserController {
    private int requestCount = 0;           // 可变计数器
    private Map<String, Object> cache = new HashMap<>();  // 可变缓存
    
    @GetMapping("/users/{id}")
    public User getUser(@PathVariable Long id) {
        requestCount++;  // 多线程访问会出现竞态条件
        return userService.findById(id);
    }
}

// 正确示例
@RestController
public class UserController {
    private final UserService userService;  // 不可变依赖
    private final MeterRegistry meterRegistry;  // 使用监控组件统计
    
    public UserController(UserService userService, MeterRegistry meterRegistry) {
        this.userService = userService;
        this.meterRegistry = meterRegistry;
    }
    
    @GetMapping("/users/{id}")
    public User getUser(@PathVariable Long id) {
        meterRegistry.counter("user.requests").increment();  // 线程安全的计数
        return userService.findById(id);
    }
}
```

#### 3.1.3 静态变量的线程安全性检查
- **检查方法**: 搜索static变量声明，检查是否为线程安全类型或正确同步
- **检查标准**: 静态变量应该是不可变的或使用线程安全的类型
- **不正确实例**:
```java
// 错误示例 - 非线程安全的静态变量
public class ConfigManager {
    private static Map<String, String> config = new HashMap<>();  // 非线程安全
    private static List<String> cache = new ArrayList<>();        // 非线程安全
    private static SimpleDateFormat formatter = new SimpleDateFormat("yyyy-MM-dd"); // 非线程安全
    
    public static void updateConfig(String key, String value) {
        config.put(key, value);  // 多线程访问会出现问题
    }
}

// 正确示例
public class ConfigManager {
    private static final Map<String, String> config = new ConcurrentHashMap<>();  // 线程安全
    private static final List<String> cache = Collections.synchronizedList(new ArrayList<>());
    private static final DateTimeFormatter formatter = DateTimeFormatter.ofPattern("yyyy-MM-dd"); // 线程安全
    
    // 或者使用不可变集合
    private static final Map<String, String> defaultConfig = 
        Collections.unmodifiableMap(Map.of("key1", "value1", "key2", "value2"));
    
    public static void updateConfig(String key, String value) {
        config.put(key, value);  // ConcurrentHashMap保证线程安全
    }
}
```

#### 3.1.4 集合类的线程安全使用
- **检查方法**: 检查多线程环境下集合类的使用，特别是ArrayList、HashMap等非线程安全集合
- **检查标准**: 在多线程环境下使用线程安全的集合类或正确同步
- **不正确实例**:
```java
// 错误示例 - 多线程环境使用非线程安全集合
@Component
public class CacheManager {
    private Map<String, Object> cache = new HashMap<>();  // 非线程安全
    private List<String> history = new ArrayList<>();     // 非线程安全
    
    @Async
    public void updateCache(String key, Object value) {
        cache.put(key, value);     // 多线程访问HashMap可能导致死循环
        history.add(key);          // ArrayList在并发修改时可能丢失数据
    }
}

// 正确示例
@Component
public class CacheManager {
    private final Map<String, Object> cache = new ConcurrentHashMap<>();  // 线程安全
    private final List<String> history = new CopyOnWriteArrayList<>();    // 线程安全
    
    // 或者使用同步
    private final Map<String, Object> syncCache = Collections.synchronizedMap(new HashMap<>());
    
    @Async
    public void updateCache(String key, Object value) {
        cache.put(key, value);     // ConcurrentHashMap保证线程安全
        history.add(key);          // CopyOnWriteArrayList保证线程安全
    }
}
```

### 3.2 锁的使用 (Major)

#### 3.2.1 避免在循环中获取锁
- **检查方法**: 代码审查时检查循环内部是否有synchronized块或Lock.lock()调用
- **检查标准**: 锁应该在循环外部获取，避免频繁的锁竞争
- **不正确实例**:
```java
// 错误示例 - 循环中获取锁
public class DataProcessor {
    private final Object lock = new Object();
    private final ReentrantLock reentrantLock = new ReentrantLock();
    
    public void processData(List<String> data) {
        for (String item : data) {
            synchronized (lock) {           // 每次循环都获取锁
                processItem(item);
            }
        }
        
        // 或者
        for (String item : data) {
            reentrantLock.lock();          // 每次循环都获取锁
            try {
                processItem(item);
            } finally {
                reentrantLock.unlock();
            }
        }
    }
}

// 正确示例
public class DataProcessor {
    private final Object lock = new Object();
    private final ReentrantLock reentrantLock = new ReentrantLock();
    
    public void processData(List<String> data) {
        synchronized (lock) {              // 在循环外获取锁
            for (String item : data) {
                processItem(item);
            }
        }
        
        // 或者
        reentrantLock.lock();
        try {
            for (String item : data) {
                processItem(item);
            }
        } finally {
            reentrantLock.unlock();
        }
    }
}
```

#### 3.2.2 锁的粒度要适当
- **检查方法**: 检查synchronized块或Lock的范围，分析是否可以减小锁的粒度
- **检查标准**: 锁的粒度应该尽可能小，只保护必要的临界区
- **不正确实例**:
```java
// 错误示例 - 锁粒度过大
public class UserManager {
    private final Map<Long, User> users = new HashMap<>();
    private final Object lock = new Object();
    
    public synchronized void updateUser(Long id, String name, String email) {
        // 整个方法都被锁定，包括不需要同步的操作
        validateInput(name, email);        // 验证操作不需要锁
        logOperation("update", id);       // 日志操作不需要锁
        
        User user = users.get(id);         // 只有这部分需要锁
        if (user != null) {
            user.setName(name);
            user.setEmail(email);
            users.put(id, user);
        }
        
        sendNotification(user);            // 通知操作不需要锁
    }
}

// 正确示例
public class UserManager {
    private final Map<Long, User> users = new ConcurrentHashMap<>();  // 使用线程安全集合
    
    public void updateUser(Long id, String name, String email) {
        validateInput(name, email);        // 无需同步的操作
        logOperation("update", id);
        
        // 只对必要的操作进行同步
        User user = users.computeIfPresent(id, (key, existingUser) -> {
            existingUser.setName(name);
            existingUser.setEmail(email);
            return existingUser;
        });
        
        if (user != null) {
            sendNotification(user);        // 无需同步的操作
        }
    }
}
```

#### 3.2.3 避免死锁的发生
- **检查方法**: 检查多个锁的获取顺序，使用死锁检测工具分析
- **检查标准**: 多个锁必须按照固定顺序获取，避免循环等待
- **不正确实例**:
```java
// 错误示例 - 可能导致死锁
public class AccountManager {
    public void transfer(Account from, Account to, BigDecimal amount) {
        synchronized (from) {              // 线程1获取from锁
            synchronized (to) {            // 线程1尝试获取to锁
                from.withdraw(amount);
                to.deposit(amount);
            }
        }
    }
    
    // 如果另一个线程同时执行 transfer(to, from, amount)
    // 就会形成死锁：线程1持有from等待to，线程2持有to等待from
}

// 正确示例
public class AccountManager {
    public void transfer(Account from, Account to, BigDecimal amount) {
        // 按照账户ID排序来获取锁，避免死锁
        Account firstLock = from.getId() < to.getId() ? from : to;
        Account secondLock = from.getId() < to.getId() ? to : from;
        
        synchronized (firstLock) {
            synchronized (secondLock) {
                from.withdraw(amount);
                to.deposit(amount);
            }
        }
    }
    
    // 或者使用超时锁
    private final ReentrantLock lock1 = new ReentrantLock();
    private final ReentrantLock lock2 = new ReentrantLock();
    
    public boolean transferWithTimeout(Account from, Account to, BigDecimal amount) {
        try {
            if (lock1.tryLock(1, TimeUnit.SECONDS)) {
                try {
                    if (lock2.tryLock(1, TimeUnit.SECONDS)) {
                        try {
                            from.withdraw(amount);
                            to.deposit(amount);
                            return true;
                        } finally {
                            lock2.unlock();
                        }
                    }
                } finally {
                    lock1.unlock();
                }
            }
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
        }
        return false;
    }
}
```

#### 3.2.4 使用try-finally确保锁释放
- **检查方法**: 检查Lock.lock()调用是否在try-finally块中正确释放
- **检查标准**: 所有显式锁必须在finally块中释放，确保异常情况下也能释放锁
- **不正确实例**:
```java
// 错误示例 - 锁可能不会被释放
public class ResourceManager {
    private final ReentrantLock lock = new ReentrantLock();
    
    public void processResource() {
        lock.lock();
        processData();           // 如果这里抛出异常，锁不会被释放
        lock.unlock();           // 可能永远不会执行
    }
    
    public void anotherMethod() {
        lock.lock();
        try {
            processData();
        } catch (Exception e) {
            lock.unlock();       // 错误：在catch中释放锁
            throw e;
        }
        lock.unlock();           // 正常情况下释放锁
    }
}

// 正确示例
public class ResourceManager {
    private final ReentrantLock lock = new ReentrantLock();
    
    public void processResource() {
        lock.lock();
        try {
            processData();       // 无论是否抛出异常
        } finally {
            lock.unlock();       // 锁都会被释放
        }
    }
    
    public void processWithCondition() {
        if (lock.tryLock()) {
            try {
                processData();
            } finally {
                lock.unlock();
            }
        } else {
            // 处理获取锁失败的情况
            handleLockFailure();
        }
    }
}
```

## 4. 安全性检查

### 4.1 输入验证 (Critical)

#### 4.1.1 所有外部输入必须进行验证
- **检查方法**: 检查Controller方法的@RequestParam、@RequestBody、@PathVariable参数是否有验证注解
- **检查标准**: 所有来自外部的输入都必须进行格式、长度、范围验证
- **不正确实例**:
```java
// 错误示例 - 缺少输入验证
@RestController
public class UserController {
    @PostMapping("/users")
    public User createUser(@RequestBody User user) {  // 没有验证注解
        return userService.save(user);
    }
    
    @GetMapping("/users/{id}")
    public User getUser(@PathVariable String id) {    // 没有验证ID格式
        Long userId = Long.parseLong(id);  // 可能抛出NumberFormatException
        return userService.findById(userId);
    }
    
    @GetMapping("/search")
    public List<User> searchUsers(@RequestParam String keyword) {  // 没有验证关键字
        return userService.search(keyword);  // 可能导致SQL注入
    }
}

// 正确示例
@RestController
@Validated
public class UserController {
    @PostMapping("/users")
    public User createUser(@Valid @RequestBody User user) {  // 使用@Valid验证
        return userService.save(user);
    }
    
    @GetMapping("/users/{id}")
    public User getUser(@PathVariable @Min(1) @Max(Long.MAX_VALUE) Long id) {  // 验证ID范围
        return userService.findById(id);
    }
    
    @GetMapping("/search")
    public List<User> searchUsers(
        @RequestParam @NotBlank @Size(min = 2, max = 50) String keyword) {  // 验证关键字
        return userService.search(keyword);
    }
}

// User实体类验证
public class User {
    @NotBlank(message = "用户名不能为空")
    @Size(min = 3, max = 20, message = "用户名长度必须在3-20之间")
    private String username;
    
    @Email(message = "邮箱格式不正确")
    @NotBlank(message = "邮箱不能为空")
    private String email;
    
    @Pattern(regexp = "^1[3-9]\\d{9}$", message = "手机号格式不正确")
    private String phone;
}
```

#### 4.1.2 使用@Valid注解进行参数校验
- **检查方法**: 搜索@RequestBody、@ModelAttribute参数，检查是否使用@Valid注解
- **检查标准**: 所有复杂对象参数必须使用@Valid注解，并配置相应的验证规则
- **不正确实例**:
```java
// 错误示例 - 缺少@Valid注解
@RestController
public class OrderController {
    @PostMapping("/orders")
    public Order createOrder(@RequestBody Order order) {  // 缺少@Valid
        return orderService.save(order);
    }
    
    @PutMapping("/orders/{id}")
    public Order updateOrder(@PathVariable Long id, 
                           @RequestBody Order order) {  // 缺少@Valid
        return orderService.update(id, order);
    }
}

// 正确示例
@RestController
public class OrderController {
    @PostMapping("/orders")
    public Order createOrder(@Valid @RequestBody Order order) {  // 添加@Valid
        return orderService.save(order);
    }
    
    @PutMapping("/orders/{id}")
    public Order updateOrder(@PathVariable @Min(1) Long id, 
                           @Valid @RequestBody Order order) {  // 添加@Valid
        return orderService.update(id, order);
    }
    
    // 全局异常处理
    @ExceptionHandler(MethodArgumentNotValidException.class)
    public ResponseEntity<ErrorResponse> handleValidationException(
            MethodArgumentNotValidException ex) {
        List<String> errors = ex.getBindingResult()
            .getFieldErrors()
            .stream()
            .map(FieldError::getDefaultMessage)
            .collect(Collectors.toList());
        return ResponseEntity.badRequest().body(new ErrorResponse(errors));
    }
}
```

#### 4.1.3 防止SQL注入攻击
- **检查方法**: 搜索SQL字符串拼接、检查是否使用参数化查询或PreparedStatement
- **检查标准**: 禁止直接拼接SQL语句，必须使用参数化查询
- **不正确实例**:
```java
// 错误示例 - SQL注入风险
@Repository
public class UserRepository {
    @Autowired
    private JdbcTemplate jdbcTemplate;
    
    public List<User> findByName(String name) {
        // 直接拼接SQL，存在注入风险
        String sql = "SELECT * FROM users WHERE name = '" + name + "'";
        return jdbcTemplate.query(sql, new UserRowMapper());
    }
    
    public List<User> searchUsers(String keyword, String orderBy) {
        // 动态ORDER BY也存在注入风险
        String sql = "SELECT * FROM users WHERE name LIKE '%" + keyword + "%' ORDER BY " + orderBy;
        return jdbcTemplate.query(sql, new UserRowMapper());
    }
    
    // 使用MyBatis时的错误示例
    @Select("SELECT * FROM users WHERE name = '${name}'")
    List<User> findByNameUnsafe(@Param("name") String name);  // 使用${}存在注入风险
}

// 正确示例
@Repository
public class UserRepository {
    @Autowired
    private JdbcTemplate jdbcTemplate;
    
    public List<User> findByName(String name) {
        // 使用参数化查询
        String sql = "SELECT * FROM users WHERE name = ?";
        return jdbcTemplate.query(sql, new UserRowMapper(), name);
    }
    
    public List<User> searchUsers(String keyword, String orderBy) {
        // 验证orderBy参数，防止注入
        Set<String> allowedColumns = Set.of("name", "email", "created_time");
        if (!allowedColumns.contains(orderBy)) {
            throw new IllegalArgumentException("Invalid order by column");
        }
        
        String sql = "SELECT * FROM users WHERE name LIKE ? ORDER BY " + orderBy;
        return jdbcTemplate.query(sql, new UserRowMapper(), "%" + keyword + "%");
    }
    
    // 使用MyBatis时的正确示例
    @Select("SELECT * FROM users WHERE name = #{name}")
    List<User> findByNameSafe(@Param("name") String name);  // 使用#{}安全
    
    // JPA查询示例
    @Query("SELECT u FROM User u WHERE u.name = :name")
    List<User> findByNameJPA(@Param("name") String name);
}
```

#### 4.1.4 防止XSS攻击
- **检查方法**: 检查输出到前端的数据是否进行了HTML转义，检查富文本处理
- **检查标准**: 所有输出到HTML的用户输入必须进行转义或过滤
- **不正确实例**:
```java
// 错误示例 - 未进行XSS防护
@RestController
public class CommentController {
    @PostMapping("/comments")
    public Comment addComment(@RequestBody Comment comment) {
        // 直接保存用户输入，未进行过滤
        return commentService.save(comment);
    }
    
    @GetMapping("/comments/{id}")
    public String getCommentHtml(@PathVariable Long id) {
        Comment comment = commentService.findById(id);
        // 直接返回HTML内容，可能包含恶意脚本
        return "<div>" + comment.getContent() + "</div>";
    }
}

// 正确示例
@RestController
public class CommentController {
    @Autowired
    private HtmlSanitizer htmlSanitizer;
    
    @PostMapping("/comments")
    public Comment addComment(@Valid @RequestBody Comment comment) {
        // 对用户输入进行HTML清理
        String cleanContent = htmlSanitizer.sanitize(comment.getContent());
        comment.setContent(cleanContent);
        return commentService.save(comment);
    }
    
    @GetMapping("/comments/{id}")
    public String getCommentHtml(@PathVariable Long id) {
        Comment comment = commentService.findById(id);
        // 使用HTML转义
        String escapedContent = HtmlUtils.htmlEscape(comment.getContent());
        return "<div>" + escapedContent + "</div>";
    }
}

// HTML清理工具配置
@Component
public class HtmlSanitizer {
    private final PolicyFactory policy;
    
    public HtmlSanitizer() {
        // 配置允许的HTML标签和属性
        this.policy = Sanitizers.FORMATTING
            .and(Sanitizers.LINKS)
            .and(Sanitizers.BLOCKS)
            .and(Sanitizers.IMAGES);
    }
    
    public String sanitize(String html) {
        return policy.sanitize(html);
    }
}
```

#### 4.1.5 文件上传安全检查
- **检查方法**: 检查文件上传接口的文件类型、大小、路径验证
- **检查标准**: 验证文件类型、限制文件大小、防止路径遍历攻击
- **不正确实例**:
```java
// 错误示例 - 文件上传安全风险
@RestController
public class FileController {
    @PostMapping("/upload")
    public String uploadFile(@RequestParam("file") MultipartFile file) {
        // 未验证文件类型
        String fileName = file.getOriginalFilename();
        
        // 直接使用用户提供的文件名，存在路径遍历风险
        String filePath = "/uploads/" + fileName;
        
        try {
            // 未检查文件大小
            file.transferTo(new File(filePath));
            return "File uploaded: " + fileName;
        } catch (IOException e) {
            return "Upload failed";
        }
    }
}

// 正确示例
@RestController
public class FileController {
    private static final Set<String> ALLOWED_EXTENSIONS = 
        Set.of(".jpg", ".jpeg", ".png", ".gif", ".pdf", ".doc", ".docx");
    private static final long MAX_FILE_SIZE = 10 * 1024 * 1024; // 10MB
    private static final String UPLOAD_DIR = "/safe/uploads/";
    
    @PostMapping("/upload")
    public ResponseEntity<String> uploadFile(@RequestParam("file") MultipartFile file) {
        // 验证文件不为空
        if (file.isEmpty()) {
            return ResponseEntity.badRequest().body("文件不能为空");
        }
        
        // 验证文件大小
        if (file.getSize() > MAX_FILE_SIZE) {
            return ResponseEntity.badRequest().body("文件大小不能超过10MB");
        }
        
        String originalFilename = file.getOriginalFilename();
        if (originalFilename == null) {
            return ResponseEntity.badRequest().body("文件名不能为空");
        }
        
        // 验证文件扩展名
        String extension = getFileExtension(originalFilename).toLowerCase();
        if (!ALLOWED_EXTENSIONS.contains(extension)) {
            return ResponseEntity.badRequest().body("不支持的文件类型");
        }
        
        // 生成安全的文件名，避免路径遍历
        String safeFileName = UUID.randomUUID().toString() + extension;
        String filePath = UPLOAD_DIR + safeFileName;
        
        try {
            // 确保上传目录存在
            Files.createDirectories(Paths.get(UPLOAD_DIR));
            
            // 保存文件
            file.transferTo(new File(filePath));
            
            // 验证文件内容（可选）
            if (!isValidFileContent(filePath, extension)) {
                Files.delete(Paths.get(filePath));
                return ResponseEntity.badRequest().body("文件内容不合法");
            }
            
            return ResponseEntity.ok("文件上传成功: " + safeFileName);
        } catch (IOException e) {
            return ResponseEntity.status(500).body("文件上传失败");
        }
    }
    
    private String getFileExtension(String filename) {
        int lastDotIndex = filename.lastIndexOf('.');
        return lastDotIndex > 0 ? filename.substring(lastDotIndex) : "";
    }
    
    private boolean isValidFileContent(String filePath, String extension) {
        // 根据文件类型验证文件头部信息
        // 这里简化处理，实际应该检查文件魔数
        return true;
    }
}
```

### 4.2 敏感信息处理 (Critical)

#### 4.2.1 密码、API密钥不能硬编码
- **检查方法**: 搜索代码中的password、key、secret、token等关键字，检查是否有硬编码的敏感信息
- **检查标准**: 所有敏感信息必须通过配置文件、环境变量或密钥管理系统获取
- **不正确实例**:
```java
// 错误示例 - 硬编码敏感信息
@Service
public class PaymentService {
    private static final String API_KEY = "sk_live_abcd1234";  // 错误：硬编码API密钥
    private static final String DB_PASSWORD = "mypassword123";  // 错误：硬编码数据库密码
    
    public void processPayment() {
        String secretKey = "jwt_secret_key_123";  // 错误：硬编码JWT密钥
        // 处理逻辑
    }
}

// 正确示例 - 通过配置获取敏感信息
@Service
public class PaymentService {
    @Value("${payment.api.key}")
    private String apiKey;
    
    @Value("${jwt.secret}")
    private String jwtSecret;
    
    // 或者通过环境变量
    private String getApiKey() {
        return System.getenv("PAYMENT_API_KEY");
    }
}
```

#### 4.2.2 日志中不包含敏感信息
- **检查方法**: 检查日志输出语句，确保不记录密码、身份证号、银行卡号等敏感信息
- **检查标准**: 敏感信息必须脱敏处理或完全不记录
- **不正确实例**:
```java
// 错误示例 - 日志包含敏感信息
@Service
public class UserService {
    private static final Logger logger = LoggerFactory.getLogger(UserService.class);
    
    public void login(String username, String password) {
        logger.info("User login attempt: username={}, password={}", username, password);  // 错误：记录密码
        
        User user = userRepository.findByUsername(username);
        logger.info("User info: {}", user);  // 错误：可能包含敏感信息
    }
    
    public void updateUser(User user) {
        logger.info("Updating user: idCard={}, bankCard={}", 
                   user.getIdCard(), user.getBankCard());  // 错误：记录身份证和银行卡
    }
}

// 正确示例 - 敏感信息脱敏
@Service
public class UserService {
    private static final Logger logger = LoggerFactory.getLogger(UserService.class);
    
    public void login(String username, String password) {
        logger.info("User login attempt: username={}", username);  // 不记录密码
        
        User user = userRepository.findByUsername(username);
        logger.info("User login success: userId={}", user.getId());  // 只记录ID
    }
    
    public void updateUser(User user) {
        logger.info("Updating user: userId={}, idCard={}, bankCard={}", 
                   user.getId(), 
                   maskIdCard(user.getIdCard()),      // 脱敏处理
                   maskBankCard(user.getBankCard()));  // 脱敏处理
    }
    
    private String maskIdCard(String idCard) {
        if (idCard == null || idCard.length() < 8) return "***";
        return idCard.substring(0, 4) + "****" + idCard.substring(idCard.length() - 4);
    }
    
    private String maskBankCard(String bankCard) {
        if (bankCard == null || bankCard.length() < 8) return "***";
        return "****" + bankCard.substring(bankCard.length() - 4);
    }
}
```

#### 4.2.3 数据库密码加密存储
- **检查方法**: 检查配置文件中的数据库连接信息，确认密码是否加密
- **检查标准**: 数据库密码必须加密存储，使用Jasypt等工具进行加密
- **不正确实例**:
```yaml
# 错误示例 - application.yml中明文密码
spring:
  datasource:
    url: jdbc:mysql://localhost:3306/mydb
    username: root
    password: mypassword123  # 错误：明文密码

# 正确示例 - 使用Jasypt加密
spring:
  datasource:
    url: jdbc:mysql://localhost:3306/mydb
    username: root
    password: ENC(encrypted_password_here)  # 加密后的密码

jasypt:
  encryptor:
    password: ${JASYPT_ENCRYPTOR_PASSWORD}  # 通过环境变量获取加密密钥
```

#### 4.2.4 错误信息不泄露系统内部结构
- **检查方法**: 检查异常处理代码，确保错误信息不暴露系统内部信息
- **检查标准**: 对外错误信息应该通用化，详细错误信息只记录在日志中
- **不正确实例**:
```java
// 错误示例 - 错误信息泄露内部结构
@RestController
public class UserController {
    
    @GetMapping("/users/{id}")
    public ResponseEntity<User> getUser(@PathVariable Long id) {
        try {
            User user = userService.findById(id);
            return ResponseEntity.ok(user);
        } catch (SQLException e) {
            // 错误：直接返回数据库异常信息
            return ResponseEntity.status(500)
                .body("Database error: " + e.getMessage());
        } catch (Exception e) {
            // 错误：返回详细的堆栈信息
            return ResponseEntity.status(500)
                .body("Internal error: " + e.getStackTrace());
        }
    }
}

// 正确示例 - 通用化错误信息
@RestController
public class UserController {
    private static final Logger logger = LoggerFactory.getLogger(UserController.class);
    
    @GetMapping("/users/{id}")
    public ResponseEntity<ApiResponse> getUser(@PathVariable Long id) {
        try {
            User user = userService.findById(id);
            return ResponseEntity.ok(ApiResponse.success(user));
        } catch (UserNotFoundException e) {
            logger.warn("User not found: id={}", id);
            return ResponseEntity.status(404)
                .body(ApiResponse.error("用户不存在"));
        } catch (Exception e) {
            logger.error("Error getting user: id={}", id, e);  // 详细错误记录在日志
            return ResponseEntity.status(500)
                .body(ApiResponse.error("系统繁忙，请稍后重试"));  // 通用错误信息
        }
    }
}
```

### 4.3 认证授权 (Critical)

#### 4.3.1 JWT token验证机制
- **检查方法**: 检查JWT token的生成、验证和刷新逻辑，确保安全性
- **检查标准**: JWT必须包含过期时间、签名验证，避免在URL中传递token
- **不正确实例**:
```java
// 错误示例 - JWT实现不安全
@Service
public class JwtService {
    private static final String SECRET = "mySecret";  // 错误：硬编码密钥
    
    public String generateToken(String username) {
        // 错误：没有设置过期时间
        return Jwts.builder()
                .setSubject(username)
                .signWith(SignatureAlgorithm.HS256, SECRET)
                .compact();
    }
    
    public boolean validateToken(String token) {
        try {
            Jwts.parser().setSigningKey(SECRET).parseClaimsJws(token);
            return true;
        } catch (Exception e) {
            return false;  // 错误：没有记录验证失败的原因
        }
    }
}

// 正确示例 - 安全的JWT实现
@Service
public class JwtService {
    private static final Logger logger = LoggerFactory.getLogger(JwtService.class);
    
    @Value("${jwt.secret}")
    private String jwtSecret;
    
    @Value("${jwt.expiration:3600000}")  // 默认1小时
    private Long jwtExpiration;
    
    public String generateToken(UserDetails userDetails) {
        Map<String, Object> claims = new HashMap<>();
        claims.put("authorities", userDetails.getAuthorities());
        
        return Jwts.builder()
                .setClaims(claims)
                .setSubject(userDetails.getUsername())
                .setIssuedAt(new Date())
                .setExpiration(new Date(System.currentTimeMillis() + jwtExpiration))
                .signWith(SignatureAlgorithm.HS512, jwtSecret)
                .compact();
    }
    
    public boolean validateToken(String token, UserDetails userDetails) {
        try {
            Claims claims = Jwts.parser()
                    .setSigningKey(jwtSecret)
                    .parseClaimsJws(token)
                    .getBody();
            
            String username = claims.getSubject();
            Date expiration = claims.getExpiration();
            
            return username.equals(userDetails.getUsername()) 
                    && !expiration.before(new Date());
        } catch (ExpiredJwtException e) {
            logger.warn("JWT token expired: {}", e.getMessage());
            return false;
        } catch (JwtException e) {
            logger.warn("JWT token validation failed: {}", e.getMessage());
            return false;
        }
    }
}
```

#### 4.3.2 接口权限控制
- **检查方法**: 检查Controller方法是否有适当的权限注解，验证权限控制逻辑
- **检查标准**: 敏感接口必须有权限控制，使用Spring Security注解或自定义权限验证
- **不正确实例**:
```java
// 错误示例 - 缺少权限控制
@RestController
@RequestMapping("/admin")
public class AdminController {
    
    @DeleteMapping("/users/{id}")  // 错误：删除用户接口没有权限控制
    public ResponseEntity<Void> deleteUser(@PathVariable Long id) {
        userService.deleteUser(id);
        return ResponseEntity.ok().build();
    }
    
    @GetMapping("/sensitive-data")  // 错误：敏感数据接口没有权限控制
    public ResponseEntity<List<SensitiveData>> getSensitiveData() {
        return ResponseEntity.ok(sensitiveDataService.findAll());
    }
}

// 正确示例 - 完善的权限控制
@RestController
@RequestMapping("/admin")
@PreAuthorize("hasRole('ADMIN')")  // 整个Controller需要ADMIN角色
public class AdminController {
    
    @DeleteMapping("/users/{id}")
    @PreAuthorize("hasAuthority('USER_DELETE')")  // 需要删除用户权限
    public ResponseEntity<Void> deleteUser(@PathVariable Long id) {
        userService.deleteUser(id);
        return ResponseEntity.ok().build();
    }
    
    @GetMapping("/sensitive-data")
    @PreAuthorize("hasAuthority('SENSITIVE_DATA_READ')")  // 需要读取敏感数据权限
    public ResponseEntity<List<SensitiveData>> getSensitiveData() {
        return ResponseEntity.ok(sensitiveDataService.findAll());
    }
    
    @PostMapping("/system-config")
    @PreAuthorize("hasRole('SUPER_ADMIN')")  // 系统配置需要超级管理员权限
    public ResponseEntity<Void> updateSystemConfig(@RequestBody SystemConfig config) {
        systemConfigService.updateConfig(config);
        return ResponseEntity.ok().build();
    }
}
```

#### 4.3.3 HTTPS配置
- **检查方法**: 检查application.yml配置文件，确认HTTPS配置是否正确
- **检查标准**: 生产环境必须启用HTTPS，配置SSL证书，禁用HTTP
- **不正确实例**:
```yaml
# 错误示例 - 生产环境使用HTTP
server:
  port: 8080  # 错误：生产环境使用HTTP端口

# 正确示例 - HTTPS配置
server:
  port: 8443
  ssl:
    enabled: true
    key-store: classpath:keystore.p12
    key-store-password: ${SSL_KEYSTORE_PASSWORD}
    key-store-type: PKCS12
    key-alias: myapp
  # 重定向HTTP到HTTPS
  http:
    port: 8080
    redirect-to-https: true

# 安全头配置
security:
  headers:
    frame-options: DENY
    content-type-options: nosniff
    xss-protection: 1; mode=block
    hsts: max-age=31536000; includeSubDomains
```

#### 4.3.4 CORS配置安全
- **检查方法**: 检查CORS配置，确保不允许所有域名访问
- **检查标准**: CORS配置应该限制允许的域名、方法和头部
- **不正确实例**:
```java
// 错误示例 - 不安全的CORS配置
@Configuration
public class CorsConfig {
    
    @Bean
    public CorsConfigurationSource corsConfigurationSource() {
        CorsConfiguration configuration = new CorsConfiguration();
        configuration.setAllowedOrigins(Arrays.asList("*"));  // 错误：允许所有域名
        configuration.setAllowedMethods(Arrays.asList("*"));  // 错误：允许所有方法
        configuration.setAllowedHeaders(Arrays.asList("*"));  // 错误：允许所有头部
        configuration.setAllowCredentials(true);  // 错误：与通配符冲突
        
        UrlBasedCorsConfigurationSource source = new UrlBasedCorsConfigurationSource();
        source.registerCorsConfiguration("/**", configuration);
        return source;
    }
}

// 正确示例 - 安全的CORS配置
@Configuration
public class CorsConfig {
    
    @Value("${app.cors.allowed-origins}")
    private List<String> allowedOrigins;
    
    @Bean
    public CorsConfigurationSource corsConfigurationSource() {
        CorsConfiguration configuration = new CorsConfiguration();
        
        // 只允许指定的域名
        configuration.setAllowedOrigins(allowedOrigins);
        
        // 只允许必要的HTTP方法
        configuration.setAllowedMethods(Arrays.asList(
            "GET", "POST", "PUT", "DELETE", "OPTIONS"
        ));
        
        // 只允许必要的头部
        configuration.setAllowedHeaders(Arrays.asList(
            "Authorization", "Content-Type", "X-Requested-With"
        ));
        
        // 允许携带认证信息
        configuration.setAllowCredentials(true);
        
        // 预检请求缓存时间
        configuration.setMaxAge(3600L);
        
        UrlBasedCorsConfigurationSource source = new UrlBasedCorsConfigurationSource();
        source.registerCorsConfiguration("/api/**", configuration);
        return source;
    }
}
```

## 5. 性能优化检查

### 5.1 数据库操作 (Major)

#### 5.1.1 避免N+1查询问题
- **检查方法**: 使用SQL分析工具检查是否存在循环中的数据库查询，检查JPA的@OneToMany、@ManyToOne关联
- **检查标准**: 避免在循环中执行数据库查询，使用JOIN或批量查询优化
- **不正确实例**:
```java
// 错误示例 - N+1查询问题
@Service
public class OrderService {
    @Autowired
    private OrderRepository orderRepository;
    @Autowired
    private OrderItemRepository orderItemRepository;
    
    public List<OrderDTO> getOrdersWithItems() {
        List<Order> orders = orderRepository.findAll();
        List<OrderDTO> result = new ArrayList<>();
        
        for (Order order : orders) {  // N+1问题：每个订单都会执行一次查询
            List<OrderItem> items = orderItemRepository.findByOrderId(order.getId());
            OrderDTO dto = new OrderDTO(order, items);
            result.add(dto);
        }
        return result;
    }
    
    // JPA中的N+1问题
    @Entity
    public class Order {
        @OneToMany(mappedBy = "order")  // 懒加载会导致N+1
        private List<OrderItem> items;
    }
    
    public void printOrderItems() {
        List<Order> orders = orderRepository.findAll();
        for (Order order : orders) {
            System.out.println(order.getItems().size());  // 每次访问items都会触发查询
        }
    }
}

// 正确示例
@Service
public class OrderService {
    @Autowired
    private OrderRepository orderRepository;
    
    public List<OrderDTO> getOrdersWithItems() {
        // 使用JOIN一次性获取所有数据
        List<Order> orders = orderRepository.findAllWithItems();
        return orders.stream()
            .map(order -> new OrderDTO(order, order.getItems()))
            .collect(Collectors.toList());
    }
    
    // 或者使用批量查询
    public List<OrderDTO> getOrdersWithItemsBatch() {
        List<Order> orders = orderRepository.findAll();
        List<Long> orderIds = orders.stream()
            .map(Order::getId)
            .collect(Collectors.toList());
        
        // 批量查询所有订单项
        Map<Long, List<OrderItem>> itemsMap = orderItemRepository
            .findByOrderIdIn(orderIds)
            .stream()
            .collect(Collectors.groupingBy(item -> item.getOrder().getId()));
        
        return orders.stream()
            .map(order -> new OrderDTO(order, itemsMap.get(order.getId())))
            .collect(Collectors.toList());
    }
}

@Repository
public interface OrderRepository extends JpaRepository<Order, Long> {
    // 使用JOIN FETCH避免N+1
    @Query("SELECT o FROM Order o LEFT JOIN FETCH o.items")
    List<Order> findAllWithItems();
    
    // 使用@EntityGraph
    @EntityGraph(attributePaths = {"items", "customer"})
    List<Order> findAll();
}

// JPA实体正确配置
@Entity
public class Order {
    @OneToMany(mappedBy = "order", fetch = FetchType.LAZY)
    @BatchSize(size = 10)  // 批量加载优化
    private List<OrderItem> items;
}
```

#### 5.1.2 合理使用索引
- **检查方法**: 检查数据库查询的WHERE、ORDER BY、JOIN条件是否有对应索引
- **检查标准**: 频繁查询的字段必须建立索引，避免全表扫描
- **不正确实例**:
```java
// 错误示例 - 缺少索引优化
@Entity
@Table(name = "users")
public class User {
    @Id
    private Long id;
    
    private String email;      // 频繁查询但没有索引
    private String phone;      // 频繁查询但没有索引
    private String status;     // 频繁过滤但没有索引
    private LocalDateTime createdTime;  // 频繁排序但没有索引
}

@Repository
public interface UserRepository extends JpaRepository<User, Long> {
    // 这些查询会导致全表扫描
    User findByEmail(String email);
    List<User> findByStatus(String status);
    List<User> findByPhoneAndStatus(String phone, String status);
    
    @Query("SELECT u FROM User u WHERE u.createdTime BETWEEN :start AND :end ORDER BY u.createdTime")
    List<User> findByCreatedTimeBetween(LocalDateTime start, LocalDateTime end);
}

// 正确示例
@Entity
@Table(name = "users", indexes = {
    @Index(name = "idx_user_email", columnList = "email", unique = true),
    @Index(name = "idx_user_phone", columnList = "phone"),
    @Index(name = "idx_user_status", columnList = "status"),
    @Index(name = "idx_user_created_time", columnList = "created_time"),
    @Index(name = "idx_user_phone_status", columnList = "phone, status")  // 复合索引
})
public class User {
    @Id
    private Long id;
    
    @Column(unique = true)
    private String email;
    
    private String phone;
    private String status;
    
    @Column(name = "created_time")
    private LocalDateTime createdTime;
}

// 或者使用原生SQL创建索引
@Component
public class DatabaseIndexInitializer {
    @Autowired
    private JdbcTemplate jdbcTemplate;
    
    @PostConstruct
    public void createIndexes() {
        // 检查索引是否存在，不存在则创建
        String[] indexes = {
            "CREATE INDEX IF NOT EXISTS idx_user_email ON users(email)",
            "CREATE INDEX IF NOT EXISTS idx_user_status_created ON users(status, created_time)",
            "CREATE INDEX IF NOT EXISTS idx_order_user_status ON orders(user_id, status)"
        };
        
        for (String sql : indexes) {
            jdbcTemplate.execute(sql);
        }
    }
}
```

#### 5.1.3 批量操作优化
- **检查方法**: 检查是否在循环中执行INSERT、UPDATE、DELETE操作
- **检查标准**: 大量数据操作必须使用批量处理，避免逐条操作
- **不正确实例**:
```java
// 错误示例 - 逐条操作效率低
@Service
public class UserService {
    @Autowired
    private UserRepository userRepository;
    
    public void importUsers(List<UserDTO> userDTOs) {
        for (UserDTO dto : userDTOs) {  // 逐条保存，效率低
            User user = new User();
            user.setName(dto.getName());
            user.setEmail(dto.getEmail());
            userRepository.save(user);  // 每次都会执行SQL
        }
    }
    
    public void updateUserStatus(List<Long> userIds, String status) {
        for (Long userId : userIds) {  // 逐条更新
            User user = userRepository.findById(userId).orElse(null);
            if (user != null) {
                user.setStatus(status);
                userRepository.save(user);
            }
        }
    }
    
    public void deleteUsers(List<Long> userIds) {
        for (Long userId : userIds) {  // 逐条删除
            userRepository.deleteById(userId);
        }
    }
}

// 正确示例
@Service
public class UserService {
    @Autowired
    private UserRepository userRepository;
    @Autowired
    private JdbcTemplate jdbcTemplate;
    
    @Transactional
    public void importUsers(List<UserDTO> userDTOs) {
        List<User> users = userDTOs.stream()
            .map(dto -> {
                User user = new User();
                user.setName(dto.getName());
                user.setEmail(dto.getEmail());
                return user;
            })
            .collect(Collectors.toList());
        
        // 批量保存
        userRepository.saveAll(users);
    }
    
    @Transactional
    public void importUsersWithJdbc(List<UserDTO> userDTOs) {
        String sql = "INSERT INTO users (name, email, created_time) VALUES (?, ?, ?)";
        
        jdbcTemplate.batchUpdate(sql, new BatchPreparedStatementSetter() {
            @Override
            public void setValues(PreparedStatement ps, int i) throws SQLException {
                UserDTO dto = userDTOs.get(i);
                ps.setString(1, dto.getName());
                ps.setString(2, dto.getEmail());
                ps.setTimestamp(3, Timestamp.valueOf(LocalDateTime.now()));
            }
            
            @Override
            public int getBatchSize() {
                return userDTOs.size();
            }
        });
    }
    
    @Transactional
    public void updateUserStatus(List<Long> userIds, String status) {
        // 批量更新
        userRepository.updateStatusByIdIn(userIds, status);
    }
    
    @Transactional
    public void deleteUsers(List<Long> userIds) {
        // 批量删除
        userRepository.deleteByIdIn(userIds);
    }
}

@Repository
public interface UserRepository extends JpaRepository<User, Long> {
    @Modifying
    @Query("UPDATE User u SET u.status = :status WHERE u.id IN :ids")
    void updateStatusByIdIn(@Param("ids") List<Long> ids, @Param("status") String status);
    
    @Modifying
    @Query("DELETE FROM User u WHERE u.id IN :ids")
    void deleteByIdIn(@Param("ids") List<Long> ids);
}

// 配置JPA批量处理
@Configuration
public class JpaConfig {
    @Bean
    public DataSource dataSource() {
        HikariConfig config = new HikariConfig();
        // 其他配置...
        
        // 启用批量处理
        config.addDataSourceProperty("rewriteBatchedStatements", "true");
        config.addDataSourceProperty("useServerPrepStmts", "true");
        
        return new HikariDataSource(config);
    }
}

# application.yml中的JPA批量配置
spring:
  jpa:
    properties:
      hibernate:
        jdbc:
          batch_size: 100
          batch_versioned_data: true
        order_inserts: true
        order_updates: true
```

#### 5.1.4 分页查询实现
- **检查方法**: 检查大数据量查询是否实现了分页，检查分页参数验证
- **检查标准**: 大数据量查询必须实现分页，避免一次性加载过多数据
- **不正确实例**:
```java
// 错误示例 - 缺少分页处理
@RestController
public class UserController {
    @Autowired
    private UserService userService;
    
    @GetMapping("/users")
    public List<User> getUsers() {
        // 没有分页，可能返回数百万条记录
        return userService.findAll();
    }
    
    @GetMapping("/users/search")
    public List<User> searchUsers(@RequestParam String keyword) {
        // 搜索结果没有分页限制
        return userService.searchByKeyword(keyword);
    }
    
    @GetMapping("/orders")
    public List<Order> getOrders(@RequestParam(required = false) String status) {
        // 订单查询没有分页，数据量可能很大
        if (status != null) {
            return orderService.findByStatus(status);
        }
        return orderService.findAll();
    }
}

@Service
public class UserService {
    @Autowired
    private UserRepository userRepository;
    
    public List<User> findAll() {
        return userRepository.findAll();  // 可能返回大量数据
    }
    
    public List<User> searchByKeyword(String keyword) {
        return userRepository.findByNameContaining(keyword);  // 没有限制结果数量
    }
}

// 正确示例
@RestController
public class UserController {
    @Autowired
    private UserService userService;
    
    @GetMapping("/users")
    public Page<User> getUsers(
            @RequestParam(defaultValue = "0") @Min(0) int page,
            @RequestParam(defaultValue = "20") @Min(1) @Max(100) int size,
            @RequestParam(defaultValue = "id") String sort,
            @RequestParam(defaultValue = "asc") String direction) {
        
        // 验证排序字段
        Set<String> allowedSortFields = Set.of("id", "name", "email", "createdTime");
        if (!allowedSortFields.contains(sort)) {
            throw new IllegalArgumentException("Invalid sort field: " + sort);
        }
        
        Sort.Direction sortDirection = "desc".equalsIgnoreCase(direction) 
            ? Sort.Direction.DESC : Sort.Direction.ASC;
        
        Pageable pageable = PageRequest.of(page, size, Sort.by(sortDirection, sort));
        return userService.findAll(pageable);
    }
    
    @GetMapping("/users/search")
    public Page<User> searchUsers(
            @RequestParam @NotBlank @Size(min = 2, max = 50) String keyword,
            @RequestParam(defaultValue = "0") @Min(0) int page,
            @RequestParam(defaultValue = "20") @Min(1) @Max(50) int size) {
        
        Pageable pageable = PageRequest.of(page, size);
        return userService.searchByKeyword(keyword, pageable);
    }
    
    @GetMapping("/orders")
    public Page<Order> getOrders(
            @RequestParam(required = false) String status,
            @RequestParam(defaultValue = "0") @Min(0) int page,
            @RequestParam(defaultValue = "20") @Min(1) @Max(100) int size) {
        
        Pageable pageable = PageRequest.of(page, size, Sort.by(Sort.Direction.DESC, "createdTime"));
        
        if (status != null) {
            return orderService.findByStatus(status, pageable);
        }
        return orderService.findAll(pageable);
    }
}

@Service
public class UserService {
    @Autowired
    private UserRepository userRepository;
    
    public Page<User> findAll(Pageable pageable) {
        return userRepository.findAll(pageable);
    }
    
    public Page<User> searchByKeyword(String keyword, Pageable pageable) {
        return userRepository.findByNameContaining(keyword, pageable);
    }
    
    // 游标分页，适用于大数据量场景
    public List<User> findUsersAfter(Long lastId, int limit) {
        return userRepository.findTop20ByIdGreaterThanOrderById(lastId);
    }
}

@Repository
public interface UserRepository extends JpaRepository<User, Long> {
    Page<User> findByNameContaining(String keyword, Pageable pageable);
    
    // 游标分页查询
    List<User> findTop20ByIdGreaterThanOrderById(Long id);
    
    // 自定义分页查询
    @Query(value = "SELECT * FROM users WHERE status = :status ORDER BY created_time DESC LIMIT :limit OFFSET :offset", 
           nativeQuery = true)
    List<User> findByStatusWithPagination(@Param("status") String status, 
                                         @Param("limit") int limit, 
                                         @Param("offset") int offset);
}
```

#### 5.1.5 连接池配置检查
- **检查方法**: 检查数据库连接池配置参数，监控连接池使用情况
- **检查标准**: 合理配置连接池大小、超时时间、连接验证等参数
- **不正确实例**:
```yaml
# 错误示例 - 连接池配置不当
spring:
  datasource:
    url: jdbc:mysql://localhost:3306/mydb
    username: root
    password: password
    # 使用默认连接池配置，可能导致性能问题
    
# 没有配置连接池参数，使用默认值
# 没有配置连接验证
# 没有配置超时时间

# 正确示例
spring:
  datasource:
    url: jdbc:mysql://localhost:3306/mydb?useSSL=false&serverTimezone=UTC&rewriteBatchedStatements=true
    username: root
    password: password
    type: com.zaxxer.hikari.HikariDataSource
    hikari:
      # 连接池配置
      minimum-idle: 5                    # 最小空闲连接数
      maximum-pool-size: 20              # 最大连接池大小
      idle-timeout: 300000               # 空闲连接超时时间(5分钟)
      max-lifetime: 1800000              # 连接最大生存时间(30分钟)
      connection-timeout: 30000          # 连接超时时间(30秒)
      
      # 连接验证
      connection-test-query: SELECT 1
      validation-timeout: 5000
      
      # 连接池名称
      pool-name: MyHikariCP
      
      # 其他优化配置
      auto-commit: true
      connection-init-sql: SET NAMES utf8mb4
      
      # 泄漏检测
      leak-detection-threshold: 60000    # 连接泄漏检测阈值(1分钟)
      
      # 数据库特定配置
      data-source-properties:
        cachePrepStmts: true
        prepStmtCacheSize: 250
        prepStmtCacheSqlLimit: 2048
        useServerPrepStmts: true
        useLocalSessionState: true
        rewriteBatchedStatements: true
        cacheResultSetMetadata: true
        cacheServerConfiguration: true
        elideSetAutoCommits: true
        maintainTimeStats: false
```

### 5.2 缓存使用 (Major)

#### 5.2.1 缓存策略合理性
- **检查方法**: 检查@Cacheable、@CacheEvict注解使用，分析缓存键设计和过期策略
- **检查标准**: 缓存应用于读多写少的数据，合理设置缓存键和过期时间
- **不正确实例**:
```java
// 错误示例 - 缓存策略不当
@Service
public class UserService {
    @Cacheable("users")  // 缓存键不明确，可能冲突
    public User findById(Long id) {
        return userRepository.findById(id).orElse(null);
    }
    
    @Cacheable("users")  // 缓存频繁变化的数据
    public User getCurrentUser() {
        // 用户状态可能频繁变化，不适合缓存
        return getCurrentUserFromSession();
    }
    
    @Cacheable("userList")  // 缓存大量数据
    public List<User> findAll() {
        // 缓存所有用户数据，内存占用大
        return userRepository.findAll();
    }
    
    public void updateUser(User user) {
        userRepository.save(user);
        // 更新后没有清除缓存，数据不一致
    }
}

// 正确示例
@Service
public class UserService {
    @Cacheable(value = "users", key = "#id", unless = "#result == null")
    public User findById(Long id) {
        return userRepository.findById(id).orElse(null);
    }
    
    @Cacheable(value = "userProfiles", key = "#id", condition = "#id != null")
    public UserProfile getUserProfile(Long id) {
        // 缓存相对稳定的用户资料
        return userRepository.findUserProfileById(id);
    }
    
    @Cacheable(value = "userStats", key = "#id + '_' + #type", 
               unless = "#result == null")
    public UserStats getUserStats(Long id, String type) {
        // 使用复合键避免冲突
        return statisticsService.calculateUserStats(id, type);
    }
    
    @CacheEvict(value = "users", key = "#user.id")
    public User updateUser(User user) {
        User saved = userRepository.save(user);
        // 更新后清除对应缓存
        return saved;
    }
    
    @CacheEvict(value = {"users", "userProfiles"}, key = "#id")
    public void deleteUser(Long id) {
        userRepository.deleteById(id);
        // 删除时清除相关缓存
    }
    
    @Caching(evict = {
        @CacheEvict(value = "users", key = "#user.id"),
        @CacheEvict(value = "userStats", allEntries = true)
    })
    public void updateUserWithStats(User user) {
        userRepository.save(user);
        // 复杂的缓存清除策略
    }
}

// 缓存配置
@Configuration
@EnableCaching
public class CacheConfig {
    @Bean
    public CacheManager cacheManager() {
        RedisCacheManager.Builder builder = RedisCacheManager
            .RedisCacheManagerBuilder
            .fromConnectionFactory(redisConnectionFactory())
            .cacheDefaults(cacheConfiguration());
        
        // 不同缓存不同配置
        Map<String, RedisCacheConfiguration> cacheConfigurations = new HashMap<>();
        
        // 用户信息缓存 - 30分钟过期
        cacheConfigurations.put("users", 
            RedisCacheConfiguration.defaultCacheConfig()
                .entryTtl(Duration.ofMinutes(30))
                .serializeKeysWith(RedisSerializationContext.SerializationPair
                    .fromSerializer(new StringRedisSerializer()))
                .serializeValuesWith(RedisSerializationContext.SerializationPair
                    .fromSerializer(new GenericJackson2JsonRedisSerializer())));
        
        // 用户统计缓存 - 1小时过期
        cacheConfigurations.put("userStats",
            RedisCacheConfiguration.defaultCacheConfig()
                .entryTtl(Duration.ofHours(1)));
        
        return builder.withInitialCacheConfigurations(cacheConfigurations).build();
    }
}
```

#### 5.2.2 缓存穿透、击穿、雪崩防护
- **检查方法**: 检查缓存实现是否有防护机制，分析缓存失效策略
- **检查标准**: 实现布隆过滤器、互斥锁、随机过期时间等防护措施
- **不正确实例**:
```java
// 错误示例 - 缺少缓存防护
@Service
public class ProductService {
    @Autowired
    private RedisTemplate<String, Object> redisTemplate;
    @Autowired
    private ProductRepository productRepository;
    
    public Product getProduct(Long id) {
        String key = "product:" + id;
        Product product = (Product) redisTemplate.opsForValue().get(key);
        
        if (product == null) {
            // 缓存穿透：查询不存在的数据，每次都会查数据库
            product = productRepository.findById(id).orElse(null);
            if (product != null) {
                // 缓存雪崩：所有缓存同时过期
                redisTemplate.opsForValue().set(key, product, 3600, TimeUnit.SECONDS);
            }
            // 没有缓存空值，导致缓存穿透
        }
        return product;
    }
    
    public List<Product> getHotProducts() {
        String key = "hot:products";
        List<Product> products = (List<Product>) redisTemplate.opsForValue().get(key);
        
        if (products == null) {
            // 缓存击穿：热点数据过期，大量请求同时查数据库
            products = productRepository.findHotProducts();
            redisTemplate.opsForValue().set(key, products, 1800, TimeUnit.SECONDS);
        }
        return products;
    }
}

// 正确示例
@Service
public class ProductService {
    @Autowired
    private RedisTemplate<String, Object> redisTemplate;
    @Autowired
    private ProductRepository productRepository;
    @Autowired
    private BloomFilter<Long> productBloomFilter;
    
    private final Map<String, Object> lockMap = new ConcurrentHashMap<>();
    
    public Product getProduct(Long id) {
        // 1. 布隆过滤器防止缓存穿透
        if (!productBloomFilter.mightContain(id)) {
            return null;  // 布隆过滤器判断数据不存在
        }
        
        String key = "product:" + id;
        Product product = (Product) redisTemplate.opsForValue().get(key);
        
        if (product == null) {
            // 2. 互斥锁防止缓存击穿
            String lockKey = "lock:product:" + id;
            synchronized (lockMap.computeIfAbsent(lockKey, k -> new Object())) {
                try {
                    // 双重检查
                    product = (Product) redisTemplate.opsForValue().get(key);
                    if (product == null) {
                        product = productRepository.findById(id).orElse(null);
                        
                        if (product != null) {
                            // 3. 随机过期时间防止缓存雪崩
                            int randomExpire = 3600 + new Random().nextInt(600); // 3600-4200秒
                            redisTemplate.opsForValue().set(key, product, randomExpire, TimeUnit.SECONDS);
                        } else {
                            // 4. 缓存空值防止缓存穿透
                            redisTemplate.opsForValue().set(key, "NULL", 300, TimeUnit.SECONDS);
                        }
                    }
                } finally {
                    lockMap.remove(lockKey);
                }
            }
        }
        
        return "NULL".equals(product) ? null : product;
    }
    
    public List<Product> getHotProducts() {
        String key = "hot:products";
        List<Product> products = (List<Product>) redisTemplate.opsForValue().get(key);
        
        if (products == null) {
            // 使用分布式锁防止缓存击穿
            String lockKey = "lock:hot:products";
            Boolean lockAcquired = redisTemplate.opsForValue().setIfAbsent(lockKey, "1", 10, TimeUnit.SECONDS);
            
            if (Boolean.TRUE.equals(lockAcquired)) {
                try {
                    // 双重检查
                    products = (List<Product>) redisTemplate.opsForValue().get(key);
                    if (products == null) {
                        products = productRepository.findHotProducts();
                        
                        // 随机过期时间
                        int randomExpire = 1800 + new Random().nextInt(300);
                        redisTemplate.opsForValue().set(key, products, randomExpire, TimeUnit.SECONDS);
                    }
                } finally {
                    redisTemplate.delete(lockKey);
                }
            } else {
                // 获取锁失败，等待一段时间后重试
                try {
                    Thread.sleep(100);
                    return getHotProducts();
                } catch (InterruptedException e) {
                    Thread.currentThread().interrupt();
                    return productRepository.findHotProducts();
                }
            }
        }
        
        return products;
    }
}

// 布隆过滤器配置
@Configuration
public class BloomFilterConfig {
    @Bean
    public BloomFilter<Long> productBloomFilter() {
        // 预期插入100万个元素，误判率0.01%
        return BloomFilter.create(Funnels.longFunnel(), 1000000, 0.0001);
    }
    
    @EventListener
    public void initBloomFilter(ApplicationReadyEvent event) {
        BloomFilter<Long> filter = productBloomFilter();
        // 初始化布隆过滤器，加载所有存在的产品ID
        productRepository.findAllIds().forEach(filter::put);
    }
}
```

#### 5.2.3 缓存过期时间设置
- **检查方法**: 检查缓存TTL设置是否合理，是否有过期时间
- **检查标准**: 根据数据特性设置合理的过期时间，避免永不过期的缓存
- **不正确实例**:
```java
// 错误示例 - 缓存过期时间设置不当
@Service
public class ConfigService {
    @Autowired
    private RedisTemplate<String, Object> redisTemplate;
    
    public void cacheConfig(String key, Object value) {
        // 没有设置过期时间，永不过期
        redisTemplate.opsForValue().set(key, value);
    }
    
    public void cacheUserSession(String sessionId, User user) {
        // 过期时间过长，可能导致内存浪费
        redisTemplate.opsForValue().set("session:" + sessionId, user, 30, TimeUnit.DAYS);
    }
    
    public void cacheTemporaryData(String key, Object data) {
        // 临时数据过期时间过长
        redisTemplate.opsForValue().set(key, data, 24, TimeUnit.HOURS);
    }
    
    public void cacheFrequentData(String key, Object data) {
        // 频繁访问的数据过期时间过短
        redisTemplate.opsForValue().set(key, data, 60, TimeUnit.SECONDS);
    }
}

// 正确示例
@Service
public class ConfigService {
    @Autowired
    private RedisTemplate<String, Object> redisTemplate;
    
    // 系统配置缓存 - 相对稳定，较长过期时间
    public void cacheSystemConfig(String key, Object value) {
        redisTemplate.opsForValue().set("config:" + key, value, 12, TimeUnit.HOURS);
    }
    
    // 用户会话缓存 - 根据业务需求设置合理时间
    public void cacheUserSession(String sessionId, User user) {
        redisTemplate.opsForValue().set("session:" + sessionId, user, 2, TimeUnit.HOURS);
    }
    
    // 临时数据缓存 - 短期过期
    public void cacheTemporaryData(String key, Object data) {
        redisTemplate.opsForValue().set("temp:" + key, data, 10, TimeUnit.MINUTES);
    }
    
    // 热点数据缓存 - 适中的过期时间
    public void cacheHotData(String key, Object data) {
        redisTemplate.opsForValue().set("hot:" + key, data, 30, TimeUnit.MINUTES);
    }
    
    // 用户信息缓存 - 根据更新频率设置
    public void cacheUserInfo(Long userId, User user) {
        redisTemplate.opsForValue().set("user:" + userId, user, 1, TimeUnit.HOURS);
    }
    
    // 统计数据缓存 - 可以较长时间
    public void cacheStatistics(String key, Object stats) {
        redisTemplate.opsForValue().set("stats:" + key, stats, 6, TimeUnit.HOURS);
    }
    
    // 验证码缓存 - 短期有效
    public void cacheVerificationCode(String phone, String code) {
        redisTemplate.opsForValue().set("verify:" + phone, code, 5, TimeUnit.MINUTES);
    }
}

// 使用配置文件管理缓存过期时间
@ConfigurationProperties(prefix = "cache.ttl")
@Component
public class CacheTtlConfig {
    private int userInfo = 3600;        // 1小时
    private int userSession = 7200;     // 2小时
    private int systemConfig = 43200;   // 12小时
    private int hotData = 1800;         // 30分钟
    private int tempData = 600;         // 10分钟
    private int verifyCode = 300;       // 5分钟
    
    // getters and setters
}

# application.yml
cache:
  ttl:
    user-info: 3600
    user-session: 7200
    system-config: 43200
    hot-data: 1800
    temp-data: 600
    verify-code: 300
```

#### 5.2.4 本地缓存线程安全性
- **检查方法**: 检查本地缓存实现是否线程安全，是否正确处理并发访问
- **检查标准**: 使用线程安全的缓存实现，正确处理缓存更新和失效
- **不正确实例**:
```java
// 错误示例 - 本地缓存线程不安全
@Component
public class LocalCacheManager {
    // 使用非线程安全的HashMap
    private Map<String, Object> cache = new HashMap<>();
    private Map<String, Long> expireTime = new HashMap<>();
    
    public void put(String key, Object value, long ttl) {
        cache.put(key, value);                    // 非线程安全
        expireTime.put(key, System.currentTimeMillis() + ttl);
    }
    
    public Object get(String key) {
        Long expire = expireTime.get(key);
        if (expire != null && expire < System.currentTimeMillis()) {
            cache.remove(key);                    // 并发修改可能出现问题
            expireTime.remove(key);
            return null;
        }
        return cache.get(key);
    }
    
    // 定时清理过期数据
    @Scheduled(fixedRate = 60000)
    public void cleanExpired() {
        long now = System.currentTimeMillis();
        expireTime.entrySet().removeIf(entry -> {
            if (entry.getValue() < now) {
                cache.remove(entry.getKey());     // 并发修改异常
                return true;
            }
            return false;
        });
    }
}

// 错误示例 - 使用非线程安全的缓存库
@Component
public class GuavaCacheManager {
    // 错误使用Guava Cache
    private Map<String, Cache<String, Object>> caches = new HashMap<>();
    
    public Cache<String, Object> getCache(String cacheName) {
        return caches.computeIfAbsent(cacheName, name -> 
            CacheBuilder.newBuilder()
                .maximumSize(1000)
                .expireAfterWrite(30, TimeUnit.MINUTES)
                .build());
    }
}

// 正确示例
@Component
public class LocalCacheManager {
    // 使用线程安全的ConcurrentHashMap
    private final ConcurrentHashMap<String, CacheEntry> cache = new ConcurrentHashMap<>();
    
    public void put(String key, Object value, long ttlMillis) {
        long expireTime = System.currentTimeMillis() + ttlMillis;
        cache.put(key, new CacheEntry(value, expireTime));
    }
    
    public Object get(String key) {
        CacheEntry entry = cache.get(key);
        if (entry == null) {
            return null;
        }
        
        if (entry.isExpired()) {
            cache.remove(key);  // ConcurrentHashMap保证线程安全
            return null;
        }
        
        return entry.getValue();
    }
    
    public void remove(String key) {
        cache.remove(key);
    }
    
    public void clear() {
        cache.clear();
    }
    
    // 定时清理过期数据
    @Scheduled(fixedRate = 60000)
    public void cleanExpired() {
        long now = System.currentTimeMillis();
        cache.entrySet().removeIf(entry -> entry.getValue().isExpired(now));
    }
    
    private static class CacheEntry {
        private final Object value;
        private final long expireTime;
        
        public CacheEntry(Object value, long expireTime) {
            this.value = value;
            this.expireTime = expireTime;
        }
        
        public Object getValue() {
            return value;
        }
        
        public boolean isExpired() {
            return isExpired(System.currentTimeMillis());
        }
        
        public boolean isExpired(long currentTime) {
            return currentTime > expireTime;
        }
    }
}

// 使用Caffeine实现线程安全的本地缓存
@Configuration
public class CaffeineConfig {
    @Bean
    public CacheManager cacheManager() {
        CaffeineCacheManager cacheManager = new CaffeineCacheManager();
        cacheManager.setCaffeine(Caffeine.newBuilder()
            .maximumSize(10000)
            .expireAfterWrite(30, TimeUnit.MINUTES)
            .recordStats());
        return cacheManager;
    }
    
    @Bean
    public Cache<String, Object> localCache() {
        return Caffeine.newBuilder()
            .maximumSize(1000)
            .expireAfterWrite(10, TimeUnit.MINUTES)
            .expireAfterAccess(5, TimeUnit.MINUTES)
            .recordStats()
            .build();
    }
}

@Service
public class UserCacheService {
    @Autowired
    private Cache<String, Object> localCache;
    
    public void cacheUser(String key, User user) {
        localCache.put(key, user);  // Caffeine保证线程安全
    }
    
    public User getUser(String key) {
        return (User) localCache.getIfPresent(key);
    }
    
    public User getUserOrLoad(String key, Function<String, User> loader) {
        return (User) localCache.get(key, loader);
    }
}
```

### 5.3 内存管理 (Major)

#### 5.3.1 避免内存泄漏
- **检查方法**: 使用内存分析工具(如JProfiler、VisualVM)检查内存使用情况，查找未关闭的资源
- **检查标准**: 及时关闭资源，避免静态集合无限增长，正确使用WeakReference
- **不正确实例**:
```java
// 错误示例 - 内存泄漏
public class CacheService {
    // 错误：静态Map会一直持有对象引用，导致内存泄漏
    private static Map<String, Object> cache = new HashMap<>();
    
    public void addToCache(String key, Object value) {
        cache.put(key, value);  // 没有清理机制
    }
    
    // 错误：没有关闭资源
    public String readFile(String fileName) throws IOException {
        FileInputStream fis = new FileInputStream(fileName);
        BufferedReader reader = new BufferedReader(new InputStreamReader(fis));
        return reader.readLine();  // 没有关闭流
    }
    
    // 错误：监听器没有移除
    public void addListener(EventListener listener) {
        EventBus.getInstance().register(listener);  // 没有对应的unregister
    }
}

// 正确示例 - 避免内存泄漏
public class CacheService {
    // 使用有限大小的缓存
    private final Map<String, Object> cache = new ConcurrentHashMap<>();
    private final int MAX_CACHE_SIZE = 1000;
    
    public void addToCache(String key, Object value) {
        if (cache.size() >= MAX_CACHE_SIZE) {
            // 清理最老的条目
            String oldestKey = cache.keySet().iterator().next();
            cache.remove(oldestKey);
        }
        cache.put(key, value);
    }
    
    // 使用try-with-resources自动关闭资源
    public String readFile(String fileName) throws IOException {
        try (FileInputStream fis = new FileInputStream(fileName);
             BufferedReader reader = new BufferedReader(new InputStreamReader(fis))) {
            return reader.readLine();
        }
    }
    
    // 提供移除监听器的方法
    private final Set<EventListener> listeners = new WeakHashMap<>();
    
    public void addListener(EventListener listener) {
        listeners.add(listener);
    }
    
    public void removeListener(EventListener listener) {
        listeners.remove(listener);
    }
    
    @PreDestroy
    public void cleanup() {
        listeners.clear();
        cache.clear();
    }
}
```

#### 5.3.2 大对象处理优化
- **检查方法**: 检查代码中是否有大对象的创建和使用，分析对象大小和生命周期
- **检查标准**: 大对象应该及时释放，避免在循环中创建大对象，使用对象池复用
- **不正确实例**:
```java
// 错误示例 - 大对象处理不当
@Service
public class DataProcessService {
    
    public void processLargeData(List<String> dataList) {
        for (String data : dataList) {
            // 错误：在循环中创建大对象
            byte[] largeArray = new byte[1024 * 1024];  // 1MB数组
            String processedData = processData(data, largeArray);
            // 处理完后没有显式释放
        }
    }
    
    public List<BigObject> loadAllData() {
        // 错误：一次性加载所有数据到内存
        List<BigObject> allData = dataRepository.findAll();  // 可能有几万条记录
        return allData;
    }
    
    public String generateLargeReport() {
        // 错误：使用String拼接大量数据
        String report = "";
        for (int i = 0; i < 10000; i++) {
            report += "Line " + i + "\n";  // 每次都创建新的String对象
        }
        return report;
    }
}

// 正确示例 - 大对象优化处理
@Service
public class DataProcessService {
    // 使用对象池复用大对象
    private final ArrayBlockingQueue<byte[]> arrayPool = new ArrayBlockingQueue<>(10);
    
    @PostConstruct
    public void initPool() {
        for (int i = 0; i < 10; i++) {
            arrayPool.offer(new byte[1024 * 1024]);
        }
    }
    
    public void processLargeData(List<String> dataList) {
        byte[] reusableArray = null;
        try {
            reusableArray = arrayPool.take();  // 从池中获取
            for (String data : dataList) {
                String processedData = processData(data, reusableArray);
                // 复用同一个数组
            }
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
        } finally {
            if (reusableArray != null) {
                arrayPool.offer(reusableArray);  // 归还到池中
            }
        }
    }
    
    // 分页加载数据
    public void processAllData() {
        int pageSize = 100;
        int page = 0;
        Page<BigObject> dataPage;
        
        do {
            dataPage = dataRepository.findAll(PageRequest.of(page, pageSize));
            processDataBatch(dataPage.getContent());
            page++;
            // 处理完一批后，让GC有机会回收
            System.gc();
        } while (dataPage.hasNext());
    }
    
    // 使用StringBuilder处理大量字符串
    public String generateLargeReport() {
        StringBuilder report = new StringBuilder(10000 * 20);  // 预估容量
        for (int i = 0; i < 10000; i++) {
            report.append("Line ").append(i).append("\n");
        }
        return report.toString();
    }
}
```

#### 5.3.3 集合使用优化
- **检查方法**: 检查集合的初始化容量设置，分析集合的使用模式
- **检查标准**: 根据预期大小设置初始容量，选择合适的集合类型，避免频繁扩容
- **不正确实例**:
```java
// 错误示例 - 集合使用不当
public class CollectionService {
    
    public List<String> processData(int expectedSize) {
        // 错误：没有设置初始容量，会频繁扩容
        List<String> result = new ArrayList<>();
        for (int i = 0; i < expectedSize; i++) {
            result.add("item" + i);
        }
        return result;
    }
    
    public Map<String, Object> createMap() {
        // 错误：使用HashMap但需要保持插入顺序
        Map<String, Object> map = new HashMap<>();
        map.put("first", 1);
        map.put("second", 2);
        map.put("third", 3);
        return map;  // 顺序不确定
    }
    
    public Set<String> removeDuplicates(List<String> list) {
        // 错误：使用HashSet但需要保持顺序
        Set<String> uniqueItems = new HashSet<>();
        for (String item : list) {
            uniqueItems.add(item);
        }
        return uniqueItems;  // 丢失了原始顺序
    }
    
    // 错误：在循环中频繁查询List
    public boolean containsAny(List<String> list, List<String> targets) {
        for (String target : targets) {
            if (list.contains(target)) {  // O(n)复杂度
                return true;
            }
        }
        return false;
    }
}

// 正确示例 - 集合优化使用
public class CollectionService {
    
    public List<String> processData(int expectedSize) {
        // 设置合适的初始容量，避免扩容
        List<String> result = new ArrayList<>(expectedSize);
        for (int i = 0; i < expectedSize; i++) {
            result.add("item" + i);
        }
        return result;
    }
    
    public Map<String, Object> createMap() {
        // 使用LinkedHashMap保持插入顺序
        Map<String, Object> map = new LinkedHashMap<>();
        map.put("first", 1);
        map.put("second", 2);
        map.put("third", 3);
        return map;
    }
    
    public Set<String> removeDuplicates(List<String> list) {
        // 使用LinkedHashSet保持插入顺序
        return new LinkedHashSet<>(list);
    }
    
    // 使用Set提高查询效率
    public boolean containsAny(List<String> list, List<String> targets) {
        Set<String> listSet = new HashSet<>(list);  // O(1)查询
        for (String target : targets) {
            if (listSet.contains(target)) {
                return true;
            }
        }
        return false;
    }
    
    // 根据使用场景选择合适的集合
    public Map<String, String> createCacheMap(int maxSize) {
        // 使用LRU缓存
        return new LinkedHashMap<String, String>(16, 0.75f, true) {
            @Override
            protected boolean removeEldestEntry(Map.Entry<String, String> eldest) {
                return size() > maxSize;
            }
        };
    }
}
```

#### 5.3.4 字符串拼接优化
- **检查方法**: 搜索代码中的字符串拼接操作，特别是循环中的拼接
- **检查标准**: 大量字符串拼接使用StringBuilder，避免在循环中使用+操作符
- **不正确实例**:
```java
// 错误示例 - 字符串拼接不当
public class StringService {
    
    public String buildQuery(List<String> conditions) {
        // 错误：在循环中使用+拼接，每次都创建新对象
        String query = "SELECT * FROM table WHERE ";
        for (int i = 0; i < conditions.size(); i++) {
            if (i > 0) {
                query += " AND ";  // 每次拼接都创建新String
            }
            query += conditions.get(i);
        }
        return query;
    }
    
    public String formatMessage(String template, Object... args) {
        // 错误：手动字符串替换
        String result = template;
        for (int i = 0; i < args.length; i++) {
            result = result.replace("{" + i + "}", String.valueOf(args[i]));
        }
        return result;
    }
    
    public String generateCsv(List<List<String>> data) {
        // 错误：大量数据拼接使用+
        String csv = "";
        for (List<String> row : data) {
            for (int i = 0; i < row.size(); i++) {
                if (i > 0) csv += ",";
                csv += row.get(i);
            }
            csv += "\n";
        }
        return csv;
    }
}

// 正确示例 - 字符串拼接优化
public class StringService {
    
    public String buildQuery(List<String> conditions) {
        // 使用StringBuilder，预估容量
        StringBuilder query = new StringBuilder(100);
        query.append("SELECT * FROM table WHERE ");
        
        for (int i = 0; i < conditions.size(); i++) {
            if (i > 0) {
                query.append(" AND ");
            }
            query.append(conditions.get(i));
        }
        return query.toString();
    }
    
    public String formatMessage(String template, Object... args) {
        // 使用MessageFormat或String.format
        return MessageFormat.format(template, args);
        // 或者
        // return String.format(template, args);
    }
    
    public String generateCsv(List<List<String>> data) {
        // 使用StringJoiner或Stream API
        return data.stream()
                .map(row -> String.join(",", row))
                .collect(Collectors.joining("\n"));
    }
    
    // 对于大量数据，使用Writer直接写入
    public void generateLargeCsv(List<List<String>> data, Writer writer) throws IOException {
        for (List<String> row : data) {
            writer.write(String.join(",", row));
            writer.write("\n");
        }
    }
    
    // 使用StringBuilder的链式调用
    public String buildComplexString(String prefix, List<String> items, String suffix) {
        return new StringBuilder()
                .append(prefix)
                .append(" [")
                .append(String.join(", ", items))
                .append("] ")
                .append(suffix)
                .toString();
    }
}
```

## 6. 可观测性检查

### 6.1 日志规范 (Major)

#### 6.1.1 禁止使用System.out.println()
- **检查方法**: 搜索代码中的`System.out.println`、`System.err.println`、`printStackTrace()`
- **检查标准**: 所有日志输出必须使用日志框架，禁止直接使用控制台输出
- **不正确实例**:
```java
// 错误示例 - 使用System.out输出
@Service
public class UserService {
    public User createUser(User user) {
        System.out.println("Creating user: " + user.getName());  // 错误
        
        try {
            User saved = userRepository.save(user);
            System.out.println("User created successfully");  // 错误
            return saved;
        } catch (Exception e) {
            e.printStackTrace();  // 错误
            throw e;
        }
    }
    
    public void debugMethod() {
        System.err.println("Debug info");  // 错误
    }
}

// 正确示例
@Service
public class UserService {
    private static final Logger logger = LoggerFactory.getLogger(UserService.class);
    
    public User createUser(User user) {
        logger.info("Creating user: {}", user.getName());
        
        try {
            User saved = userRepository.save(user);
            logger.info("User created successfully with id: {}", saved.getId());
            return saved;
        } catch (Exception e) {
            logger.error("Failed to create user: {}", user.getName(), e);
            throw e;
        }
    }
    
    public void debugMethod() {
        logger.debug("Debug information");
    }
}
```

#### 6.1.2 使用统一的日志框架
- **检查方法**: 检查项目依赖是否使用SLF4J + Logback，检查Logger声明
- **检查标准**: 统一使用SLF4J作为日志门面，Logback作为实现
- **不正确实例**:
```java
// 错误示例 - 混用不同的日志框架
import java.util.logging.Logger;  // 错误：使用JUL
import org.apache.log4j.Logger;   // 错误：直接使用Log4j
import org.apache.commons.logging.Log;  // 错误：使用JCL

@Service
public class OrderService {
    // 错误的Logger声明
    private static final java.util.logging.Logger julLogger = 
        java.util.logging.Logger.getLogger(OrderService.class.getName());
    
    private static final org.apache.log4j.Logger log4jLogger = 
        org.apache.log4j.Logger.getLogger(OrderService.class);
    
    public void processOrder() {
        julLogger.info("Processing order");
        log4jLogger.debug("Order details");
    }
}

// 正确示例
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

@Service
public class OrderService {
    private static final Logger logger = LoggerFactory.getLogger(OrderService.class);
    
    public void processOrder() {
        logger.info("Processing order");
        logger.debug("Order details");
    }
}

// Maven依赖配置
<dependencies>
    <!-- SLF4J API -->
    <dependency>
        <groupId>org.slf4j</groupId>
        <artifactId>slf4j-api</artifactId>
        <version>1.7.36</version>
    </dependency>
    
    <!-- Logback实现 -->
    <dependency>
        <groupId>ch.qos.logback</groupId>
        <artifactId>logback-classic</artifactId>
        <version>1.2.12</version>
    </dependency>
    
    <!-- 桥接其他日志框架 -->
    <dependency>
        <groupId>org.slf4j</groupId>
        <artifactId>log4j-over-slf4j</artifactId>
        <version>1.7.36</version>
    </dependency>
</dependencies>
```

#### 6.1.3 日志级别配置正确
- **检查方法**: 检查logback.xml配置文件，检查代码中日志级别使用
- **检查标准**: 根据环境和重要性正确设置日志级别，避免生产环境输出DEBUG日志
- **不正确实例**:
```java
// 错误示例 - 日志级别使用不当
@Service
public class PaymentService {
    private static final Logger logger = LoggerFactory.getLogger(PaymentService.class);
    
    public void processPayment(Payment payment) {
        // 错误：重要业务操作使用DEBUG级别
        logger.debug("Processing payment for amount: {}", payment.getAmount());
        
        try {
            paymentGateway.charge(payment);
            // 错误：成功操作使用ERROR级别
            logger.error("Payment processed successfully");
        } catch (Exception e) {
            // 错误：异常使用INFO级别
            logger.info("Payment failed", e);
        }
        
        // 错误：在循环中使用INFO级别
        for (PaymentItem item : payment.getItems()) {
            logger.info("Processing item: {}", item.getName());
        }
    }
}

// 正确示例
@Service
public class PaymentService {
    private static final Logger logger = LoggerFactory.getLogger(PaymentService.class);
    
    public void processPayment(Payment payment) {
        // 正确：重要业务操作使用INFO级别
        logger.info("Processing payment for amount: {} from user: {}", 
                   payment.getAmount(), payment.getUserId());
        
        try {
            paymentGateway.charge(payment);
            // 正确：成功操作使用INFO级别
            logger.info("Payment processed successfully, transaction id: {}", 
                       payment.getTransactionId());
        } catch (PaymentException e) {
            // 正确：业务异常使用WARN级别
            logger.warn("Payment failed for user: {}, reason: {}", 
                       payment.getUserId(), e.getMessage());
        } catch (Exception e) {
            // 正确：系统异常使用ERROR级别
            logger.error("Unexpected error processing payment for user: {}", 
                        payment.getUserId(), e);
        }
        
        // 正确：循环中的详细信息使用DEBUG级别
        if (logger.isDebugEnabled()) {
            for (PaymentItem item : payment.getItems()) {
                logger.debug("Processing item: {} with price: {}", 
                           item.getName(), item.getPrice());
            }
        }
    }
}

<!-- logback-spring.xml配置 -->
<configuration>
    <springProfile name="dev">
        <root level="DEBUG">
            <appender-ref ref="CONSOLE"/>
        </root>
    </springProfile>
    
    <springProfile name="test">
        <root level="INFO">
            <appender-ref ref="CONSOLE"/>
            <appender-ref ref="FILE"/>
        </root>
    </springProfile>
    
    <springProfile name="prod">
        <root level="WARN">
            <appender-ref ref="FILE"/>
        </root>
        
        <!-- 业务日志保持INFO级别 -->
        <logger name="com.company.service" level="INFO"/>
    </springProfile>
</configuration>
```

#### 6.1.4 包含TraceId进行链路追踪
- **检查方法**: 检查日志配置是否包含TraceId，检查MDC使用
- **检查标准**: 所有日志必须包含TraceId，支持分布式链路追踪
- **不正确实例**:
```java
// 错误示例 - 缺少TraceId
@RestController
public class OrderController {
    private static final Logger logger = LoggerFactory.getLogger(OrderController.class);
    
    @PostMapping("/orders")
    public Order createOrder(@RequestBody Order order) {
        // 错误：日志中没有TraceId，无法追踪请求链路
        logger.info("Creating order for user: {}", order.getUserId());
        
        Order saved = orderService.save(order);
        logger.info("Order created with id: {}", saved.getId());
        
        return saved;
    }
}

// 正确示例
@RestController
public class OrderController {
    private static final Logger logger = LoggerFactory.getLogger(OrderController.class);
    
    @PostMapping("/orders")
    public Order createOrder(@RequestBody Order order) {
        // 正确：日志会自动包含TraceId
        logger.info("Creating order for user: {}", order.getUserId());
        
        Order saved = orderService.save(order);
        logger.info("Order created with id: {}", saved.getId());
        
        return saved;
    }
}

// TraceId配置
@Component
public class TraceIdFilter implements Filter {
    @Override
    public void doFilter(ServletRequest request, ServletResponse response, 
                        FilterChain chain) throws IOException, ServletException {
        try {
            // 生成或获取TraceId
            String traceId = getOrGenerateTraceId(request);
            MDC.put("traceId", traceId);
            
            // 设置响应头
            if (response instanceof HttpServletResponse) {
                ((HttpServletResponse) response).setHeader("X-Trace-Id", traceId);
            }
            
            chain.doFilter(request, response);
        } finally {
            MDC.clear();
        }
    }
    
    private String getOrGenerateTraceId(ServletRequest request) {
        if (request instanceof HttpServletRequest) {
            String traceId = ((HttpServletRequest) request).getHeader("X-Trace-Id");
            if (traceId != null && !traceId.isEmpty()) {
                return traceId;
            }
        }
        return UUID.randomUUID().toString().replace("-", "");
    }
}

<!-- logback配置包含TraceId -->
<configuration>
    <appender name="CONSOLE" class="ch.qos.logback.core.ConsoleAppender">
        <encoder>
            <pattern>%d{yyyy-MM-dd HH:mm:ss.SSS} [%thread] %-5level [%X{traceId:-}] %logger{36} - %msg%n</pattern>
        </encoder>
    </appender>
    
    <appender name="FILE" class="ch.qos.logback.core.rolling.RollingFileAppender">
        <encoder class="net.logstash.logback.encoder.LoggingEventCompositeJsonEncoder">
            <providers>
                <timestamp/>
                <logLevel/>
                <loggerName/>
                <mdc/>
                <message/>
                <stackTrace/>
            </providers>
        </encoder>
    </appender>
</configuration>

# Spring Cloud Sleuth自动配置
spring:
  sleuth:
    sampler:
      probability: 1.0  # 采样率
    zipkin:
      base-url: http://zipkin-server:9411
```

#### 6.1.5 敏感信息脱敏处理
- **检查方法**: 检查日志中是否包含密码、身份证、手机号等敏感信息
- **检查标准**: 所有敏感信息必须脱敏或不记录到日志中
- **不正确实例**:
```java
// 错误示例 - 敏感信息泄露
@Service
public class UserService {
    private static final Logger logger = LoggerFactory.getLogger(UserService.class);
    
    public User login(LoginRequest request) {
        // 错误：记录明文密码
        logger.info("User login attempt: username={}, password={}", 
                   request.getUsername(), request.getPassword());
        
        User user = authenticate(request);
        
        // 错误：记录完整身份证号
        logger.info("User logged in: {}, idCard={}, phone={}", 
                   user.getUsername(), user.getIdCard(), user.getPhone());
        
        return user;
    }
    
    public void updateUser(User user) {
        // 错误：记录银行卡号
        logger.info("Updating user: {}", user.toString());  // toString可能包含敏感信息
    }
}

// 正确示例
@Service
public class UserService {
    private static final Logger logger = LoggerFactory.getLogger(UserService.class);
    
    public User login(LoginRequest request) {
        // 正确：不记录密码
        logger.info("User login attempt: username={}", request.getUsername());
        
        User user = authenticate(request);
        
        // 正确：脱敏处理敏感信息
        logger.info("User logged in: {}, idCard={}, phone={}", 
                   user.getUsername(), 
                   maskIdCard(user.getIdCard()), 
                   maskPhone(user.getPhone()));
        
        return user;
    }
    
    public void updateUser(User user) {
        // 正确：只记录非敏感信息
        logger.info("Updating user: id={}, username={}", 
                   user.getId(), user.getUsername());
    }
    
    private String maskIdCard(String idCard) {
        if (idCard == null || idCard.length() < 8) {
            return "***";
        }
        return idCard.substring(0, 4) + "****" + idCard.substring(idCard.length() - 4);
    }
    
    private String maskPhone(String phone) {
        if (phone == null || phone.length() < 7) {
            return "***";
        }
        return phone.substring(0, 3) + "****" + phone.substring(phone.length() - 4);
    }
}

// 通用脱敏工具类
@Component
public class SensitiveDataMasker {
    
    public static String maskEmail(String email) {
        if (email == null || !email.contains("@")) {
            return "***";
        }
        String[] parts = email.split("@");
        String username = parts[0];
        if (username.length() <= 2) {
            return "***@" + parts[1];
        }
        return username.substring(0, 2) + "***@" + parts[1];
    }
    
    public static String maskBankCard(String bankCard) {
        if (bankCard == null || bankCard.length() < 8) {
            return "***";
        }
        return bankCard.substring(0, 4) + " **** **** " + 
               bankCard.substring(bankCard.length() - 4);
    }
    
    public static String maskPassword() {
        return "[PROTECTED]";  // 密码永远不记录
    }
}

// 自定义日志脱敏
@JsonSerialize(using = SensitiveDataSerializer.class)
public class User {
    private String username;
    
    @SensitiveData(type = SensitiveType.ID_CARD)
    private String idCard;
    
    @SensitiveData(type = SensitiveType.PHONE)
    private String phone;
    
    @SensitiveData(type = SensitiveType.EMAIL)
    private String email;
    
    // getters and setters
}

public class SensitiveDataSerializer extends JsonSerializer<Object> {
    @Override
    public void serialize(Object value, JsonGenerator gen, SerializerProvider serializers) 
            throws IOException {
        if (value == null) {
            gen.writeNull();
            return;
        }
        
        // 根据注解类型进行脱敏
        Field field = getCurrentField();
        if (field != null && field.isAnnotationPresent(SensitiveData.class)) {
            SensitiveData annotation = field.getAnnotation(SensitiveData.class);
            String maskedValue = maskByType(value.toString(), annotation.type());
            gen.writeString(maskedValue);
        } else {
            gen.writeString(value.toString());
        }
    }
}
```

### 6.2 监控集成 (Major)

#### 6.2.1 Metrics指标暴露
- **检查方法**: 检查是否集成Micrometer，检查Prometheus端点配置
- **检查标准**: 暴露应用关键指标，包括HTTP请求、数据库连接、JVM指标等
- **不正确实例**:
```java
// 错误示例 - 缺少监控指标
@RestController
public class OrderController {
    @Autowired
    private OrderService orderService;
    
    @GetMapping("/orders/{id}")
    public Order getOrder(@PathVariable Long id) {
        // 没有任何监控指标
        return orderService.findById(id);
    }
    
    @PostMapping("/orders")
    public Order createOrder(@RequestBody Order order) {
        // 没有记录业务指标
        return orderService.save(order);
    }
}

// 正确示例
@RestController
public class OrderController {
    @Autowired
    private OrderService orderService;
    @Autowired
    private MeterRegistry meterRegistry;
    
    private final Counter orderCreatedCounter;
    private final Timer orderProcessingTimer;
    
    public OrderController(MeterRegistry meterRegistry) {
        this.meterRegistry = meterRegistry;
        this.orderCreatedCounter = Counter.builder("orders.created")
            .description("Number of orders created")
            .register(meterRegistry);
        this.orderProcessingTimer = Timer.builder("orders.processing.time")
            .description("Order processing time")
            .register(meterRegistry);
    }
    
    @GetMapping("/orders/{id}")
    @Timed(name = "orders.get", description = "Time taken to get order")
    public Order getOrder(@PathVariable Long id) {
        return orderService.findById(id);
    }
    
    @PostMapping("/orders")
    public Order createOrder(@RequestBody Order order) {
        return orderProcessingTimer.recordCallable(() -> {
            Order saved = orderService.save(order);
            orderCreatedCounter.increment(
                Tags.of("status", "success", "type", order.getType()));
            return saved;
        });
    }
}

# application.yml配置
management:
  endpoints:
    web:
      exposure:
        include: health,info,metrics,prometheus
  endpoint:
    health:
      show-details: always
    metrics:
      enabled: true
    prometheus:
      enabled: true
  metrics:
    export:
      prometheus:
        enabled: true
    distribution:
      percentiles-histogram:
        http.server.requests: true
      percentiles:
        http.server.requests: 0.5, 0.95, 0.99
```

#### 6.2.2 健康检查端点实现
- **检查方法**: 检查/actuator/health端点，检查自定义健康检查器
- **检查标准**: 实现数据库、Redis、外部服务等关键依赖的健康检查
- **不正确实例**:
```java
// 错误示例 - 缺少健康检查
@SpringBootApplication
public class Application {
    public static void main(String[] args) {
        SpringApplication.run(Application.class, args);
    }
    // 没有自定义健康检查
}

// 正确示例
@Component
public class DatabaseHealthIndicator implements HealthIndicator {
    @Autowired
    private DataSource dataSource;
    
    @Override
    public Health health() {
        try (Connection connection = dataSource.getConnection()) {
            if (connection.isValid(1)) {
                return Health.up()
                    .withDetail("database", "Available")
                    .withDetail("validationQuery", "SELECT 1")
                    .build();
            }
        } catch (Exception e) {
            return Health.down()
                .withDetail("database", "Unavailable")
                .withException(e)
                .build();
        }
        return Health.down().withDetail("database", "Connection invalid").build();
    }
}

@Component
public class RedisHealthIndicator implements HealthIndicator {
    @Autowired
    private RedisTemplate<String, Object> redisTemplate;
    
    @Override
    public Health health() {
        try {
            String result = redisTemplate.execute((RedisCallback<String>) connection -> {
                return connection.ping();
            });
            
            if ("PONG".equals(result)) {
                return Health.up()
                    .withDetail("redis", "Available")
                    .withDetail("ping", result)
                    .build();
            }
        } catch (Exception e) {
            return Health.down()
                .withDetail("redis", "Unavailable")
                .withException(e)
                .build();
        }
        return Health.down().withDetail("redis", "Ping failed").build();
    }
}

@Component
public class ExternalServiceHealthIndicator implements HealthIndicator {
    @Autowired
    private RestTemplate restTemplate;
    
    @Value("${external.service.health.url}")
    private String healthUrl;
    
    @Override
    public Health health() {
        try {
            ResponseEntity<String> response = restTemplate.getForEntity(healthUrl, String.class);
            if (response.getStatusCode().is2xxSuccessful()) {
                return Health.up()
                    .withDetail("externalService", "Available")
                    .withDetail("responseTime", System.currentTimeMillis())
                    .build();
            }
        } catch (Exception e) {
            return Health.down()
                .withDetail("externalService", "Unavailable")
                .withException(e)
                .build();
        }
        return Health.down().withDetail("externalService", "Health check failed").build();
    }
}
```

### 6.3 告警配置 (Major)

#### 6.3.1 关键指标告警规则
- **检查方法**: 检查Prometheus告警规则配置，检查告警阈值设置
- **检查标准**: 配置响应时间、错误率、资源使用率等关键指标告警
- **配置示例**:
```yaml
# prometheus-alerts.yml
groups:
  - name: application-alerts
    rules:
      # HTTP请求响应时间告警
      - alert: HighResponseTime
        expr: histogram_quantile(0.95, http_request_duration_seconds_bucket{job="my-app"}) > 2
        for: 5m
        labels:
          severity: warning
        annotations:
          summary: "High response time detected"
          description: "95th percentile response time is {{ $value }}s for {{ $labels.instance }}"
      
      # 错误率告警
      - alert: HighErrorRate
        expr: rate(http_requests_total{status=~"5.."}[5m]) / rate(http_requests_total[5m]) > 0.05
        for: 2m
        labels:
          severity: critical
        annotations:
          summary: "High error rate detected"
          description: "Error rate is {{ $value | humanizePercentage }} for {{ $labels.instance }}"
      
      # JVM内存使用率告警
      - alert: HighMemoryUsage
        expr: jvm_memory_used_bytes{area="heap"} / jvm_memory_max_bytes{area="heap"} > 0.8
        for: 5m
        labels:
          severity: warning
        annotations:
          summary: "High memory usage detected"
          description: "Memory usage is {{ $value | humanizePercentage }} for {{ $labels.instance }}"
      
      # 数据库连接池告警
      - alert: DatabaseConnectionPoolExhausted
        expr: hikaricp_connections_active / hikaricp_connections_max > 0.9
        for: 2m
        labels:
          severity: critical
        annotations:
          summary: "Database connection pool nearly exhausted"
          description: "Connection pool usage is {{ $value | humanizePercentage }} for {{ $labels.instance }}"
      
      # 应用实例下线告警
      - alert: ApplicationInstanceDown
        expr: up{job="my-app"} == 0
        for: 1m
        labels:
          severity: critical
        annotations:
          summary: "Application instance is down"
          description: "{{ $labels.instance }} has been down for more than 1 minute"
```

## 7. 容错和稳定性检查

### 7.1 熔断器配置 (Critical)

#### 7.1.1 熔断器实现和配置
- **检查方法**: 检查是否使用Resilience4j或Hystrix，检查熔断器配置参数
- **检查标准**: 外部服务调用必须配置熔断器，合理设置失败阈值和恢复时间
- **不正确实例**:
```java
// 错误示例 - 缺少熔断器保护
@Service
public class PaymentService {
    @Autowired
    private RestTemplate restTemplate;
    
    public PaymentResult processPayment(PaymentRequest request) {
        // 直接调用外部服务，没有熔断保护
        String url = "http://payment-gateway/api/charge";
        return restTemplate.postForObject(url, request, PaymentResult.class);
    }
    
    public UserInfo getUserInfo(Long userId) {
        // 调用用户服务，没有降级策略
        String url = "http://user-service/api/users/" + userId;
        return restTemplate.getForObject(url, UserInfo.class);
    }
}

// 正确示例 - 使用Resilience4j
@Service
public class PaymentService {
    @Autowired
    private RestTemplate restTemplate;
    @Autowired
    private CircuitBreaker paymentCircuitBreaker;
    @Autowired
    private CircuitBreaker userServiceCircuitBreaker;
    
    @CircuitBreaker(name = "payment-gateway", fallbackMethod = "paymentFallback")
    @TimeLimiter(name = "payment-gateway")
    @Retry(name = "payment-gateway")
    public PaymentResult processPayment(PaymentRequest request) {
        String url = "http://payment-gateway/api/charge";
        return restTemplate.postForObject(url, request, PaymentResult.class);
    }
    
    public PaymentResult paymentFallback(PaymentRequest request, Exception ex) {
        logger.warn("Payment service fallback triggered for request: {}", request.getId(), ex);
        return PaymentResult.builder()
            .status("PENDING")
            .message("Payment service temporarily unavailable, please try again later")
            .build();
    }
    
    @CircuitBreaker(name = "user-service", fallbackMethod = "getUserInfoFallback")
    public UserInfo getUserInfo(Long userId) {
        String url = "http://user-service/api/users/" + userId;
        return restTemplate.getForObject(url, UserInfo.class);
    }
    
    public UserInfo getUserInfoFallback(Long userId, Exception ex) {
        logger.warn("User service fallback triggered for user: {}", userId, ex);
        return UserInfo.builder()
            .id(userId)
            .name("Unknown User")
            .build();
    }
}

# application.yml配置
resilience4j:
  circuitbreaker:
    instances:
      payment-gateway:
        failure-rate-threshold: 50          # 失败率阈值50%
        slow-call-rate-threshold: 50        # 慢调用率阈值50%
        slow-call-duration-threshold: 2s    # 慢调用时间阈值2秒
        minimum-number-of-calls: 10         # 最小调用次数
        sliding-window-size: 20             # 滑动窗口大小
        wait-duration-in-open-state: 30s    # 熔断器打开状态等待时间
        permitted-number-of-calls-in-half-open-state: 5  # 半开状态允许调用次数
      user-service:
        failure-rate-threshold: 60
        minimum-number-of-calls: 5
        sliding-window-size: 10
        wait-duration-in-open-state: 20s
  
  timelimiter:
    instances:
      payment-gateway:
        timeout-duration: 3s
      user-service:
        timeout-duration: 2s
  
  retry:
    instances:
      payment-gateway:
        max-attempts: 3
        wait-duration: 1s
        exponential-backoff-multiplier: 2
```

### 7.2 重试机制 (Major)

#### 7.2.1 重试策略配置
- **检查方法**: 检查重试注解使用，检查重试条件和次数配置
- **检查标准**: 对临时性失败实现重试，配置合理的重试次数和间隔
- **不正确实例**:
```java
// 错误示例 - 重试配置不当
@Service
public class NotificationService {
    @Autowired
    private EmailService emailService;
    
    public void sendEmail(String to, String subject, String content) {
        try {
            emailService.send(to, subject, content);
        } catch (Exception e) {
            // 错误：简单重试，没有退避策略
            for (int i = 0; i < 5; i++) {
                try {
                    Thread.sleep(1000);  // 固定间隔
                    emailService.send(to, subject, content);
                    break;
                } catch (Exception retryEx) {
                    if (i == 4) {
                        throw retryEx;  // 最后一次重试失败
                    }
                }
            }
        }
    }
    
    @Retryable(value = Exception.class, maxAttempts = 10)  // 错误：重试次数过多
    public void processMessage(Message message) {
        // 对所有异常都重试，包括业务异常
        messageProcessor.process(message);
    }
}

// 正确示例
@Service
public class NotificationService {
    private static final Logger logger = LoggerFactory.getLogger(NotificationService.class);
    
    @Autowired
    private EmailService emailService;
    
    @Retryable(
        value = {ConnectException.class, SocketTimeoutException.class},  // 只对特定异常重试
        maxAttempts = 3,
        backoff = @Backoff(
            delay = 1000,      // 初始延迟1秒
            multiplier = 2,    // 指数退避倍数
            maxDelay = 10000   // 最大延迟10秒
        )
    )
    public void sendEmail(String to, String subject, String content) {
        logger.info("Attempting to send email to: {}", to);
        emailService.send(to, subject, content);
    }
    
    @Recover
    public void recoverSendEmail(ConnectException ex, String to, String subject, String content) {
        logger.error("Failed to send email after retries, saving to dead letter queue", ex);
        deadLetterService.saveFailedEmail(to, subject, content, ex.getMessage());
    }
    
    @Retryable(
        value = {TransientDataAccessException.class},
        maxAttempts = 3,
        backoff = @Backoff(delay = 500, multiplier = 1.5)
    )
    public void processMessage(Message message) {
        // 只对数据库临时异常重试
        if (message.getType().equals("INVALID")) {
            throw new IllegalArgumentException("Invalid message type");  // 业务异常不重试
        }
        messageProcessor.process(message);
    }
    
    @Recover
    public void recoverProcessMessage(TransientDataAccessException ex, Message message) {
        logger.error("Failed to process message after retries: {}", message.getId(), ex);
        messageService.markAsFailed(message.getId(), ex.getMessage());
    }
}

# 配置重试
@Configuration
@EnableRetry
public class RetryConfig {
    
    @Bean
    public RetryTemplate retryTemplate() {
        RetryTemplate retryTemplate = new RetryTemplate();
        
        // 重试策略
        SimpleRetryPolicy retryPolicy = new SimpleRetryPolicy();
        retryPolicy.setMaxAttempts(3);
        retryTemplate.setRetryPolicy(retryPolicy);
        
        // 退避策略
        ExponentialBackOffPolicy backOffPolicy = new ExponentialBackOffPolicy();
        backOffPolicy.setInitialInterval(1000);
        backOffPolicy.setMultiplier(2.0);
        backOffPolicy.setMaxInterval(10000);
        retryTemplate.setBackOffPolicy(backOffPolicy);
        
        return retryTemplate;
    }
}
```

### 7.3 限流配置 (Major)

#### 7.3.1 接口限流实现
- **检查方法**: 检查限流注解或配置，检查限流算法实现
- **检查标准**: 对外接口必须配置限流，防止系统过载
- **不正确实例**:
```java
// 错误示例 - 缺少限流保护
@RestController
public class ApiController {
    
    @PostMapping("/api/orders")
    public Order createOrder(@RequestBody Order order) {
        // 没有限流保护，可能被恶意请求攻击
        return orderService.create(order);
    }
    
    @GetMapping("/api/export")
    public void exportData(HttpServletResponse response) {
        // 导出接口没有限流，可能消耗大量资源
        dataExportService.exportAllData(response);
    }
}

// 正确示例
@RestController
public class ApiController {
    @Autowired
    private RateLimiter rateLimiter;
    
    @PostMapping("/api/orders")
    @RateLimiter(name = "order-api", fallbackMethod = "createOrderFallback")
    public ResponseEntity<Order> createOrder(@RequestBody Order order) {
        Order created = orderService.create(order);
        return ResponseEntity.ok(created);
    }
    
    public ResponseEntity<String> createOrderFallback(Order order, Exception ex) {
        return ResponseEntity.status(HttpStatus.TOO_MANY_REQUESTS)
            .body("Too many requests, please try again later");
    }
    
    @GetMapping("/api/export")
    @RateLimiter(name = "export-api")
    public void exportData(HttpServletResponse response) {
        dataExportService.exportAllData(response);
    }
    
    // 用户级别限流
    @GetMapping("/api/user/{userId}/orders")
    public List<Order> getUserOrders(@PathVariable Long userId) {
        String key = "user:" + userId;
        if (!rateLimiter.tryAcquire(key, 10, TimeUnit.MINUTES)) {  // 每用户每10分钟最多10次
            throw new TooManyRequestsException("User request limit exceeded");
        }
        return orderService.findByUserId(userId);
    }
}

# Resilience4j限流配置
resilience4j:
  ratelimiter:
    instances:
      order-api:
        limit-for-period: 100        # 时间窗口内允许的请求数
        limit-refresh-period: 1s     # 时间窗口大小
        timeout-duration: 0s         # 获取许可的超时时间
      export-api:
        limit-for-period: 5
        limit-refresh-period: 1m
        timeout-duration: 5s

// 自定义限流器实现
@Component
public class CustomRateLimiter {
    private final RedisTemplate<String, String> redisTemplate;
    private final Map<String, com.google.common.util.concurrent.RateLimiter> localLimiters;
    
    public CustomRateLimiter(RedisTemplate<String, String> redisTemplate) {
        this.redisTemplate = redisTemplate;
        this.localLimiters = new ConcurrentHashMap<>();
    }
    
    // 令牌桶算法
    public boolean tryAcquire(String key, double permitsPerSecond) {
        com.google.common.util.concurrent.RateLimiter limiter = localLimiters
            .computeIfAbsent(key, k -> 
                com.google.common.util.concurrent.RateLimiter.create(permitsPerSecond));
        return limiter.tryAcquire();
    }
    
    // 滑动窗口算法（基于Redis）
    public boolean tryAcquireWithSlidingWindow(String key, int maxRequests, 
                                              Duration windowSize) {
        String script = 
            "local key = KEYS[1]\n" +
            "local window = tonumber(ARGV[1])\n" +
            "local limit = tonumber(ARGV[2])\n" +
            "local current = tonumber(redis.call('GET', key) or 0)\n" +
            "if current < limit then\n" +
            "  redis.call('INCR', key)\n" +
            "  redis.call('EXPIRE', key, window)\n" +
            "  return 1\n" +
            "else\n" +
            "  return 0\n" +
            "end";
        
        Long result = redisTemplate.execute(
            (RedisCallback<Long>) connection -> 
                (Long) connection.eval(
                    script.getBytes(),
                    ReturnType.INTEGER,
                    1,
                    key.getBytes(),
                    String.valueOf(windowSize.getSeconds()).getBytes(),
                    String.valueOf(maxRequests).getBytes()
                )
        );
        
        return result != null && result == 1L;
    }
}
```

### 7.4 异常处理 (Critical)

#### 7.4.1 全局异常处理器
- **检查方法**: 检查@RestControllerAdvice注解，检查异常处理方法
- **检查标准**: 实现统一的异常处理，区分业务异常和系统异常
- **不正确实例**:
```java
// 错误示例 - 异常处理不统一
@RestController
public class UserController {
    
    @GetMapping("/users/{id}")
    public User getUser(@PathVariable Long id) {
        try {
            return userService.findById(id);
        } catch (UserNotFoundException e) {
            // 错误：在Controller中处理异常
            return null;
        } catch (Exception e) {
            // 错误：异常信息直接暴露给客户端
            throw new RuntimeException(e.getMessage());
        }
    }
    
    @PostMapping("/users")
    public User createUser(@RequestBody User user) {
        // 错误：没有异常处理
        return userService.save(user);
    }
}

// 正确示例
@RestControllerAdvice
public class GlobalExceptionHandler {
    private static final Logger logger = LoggerFactory.getLogger(GlobalExceptionHandler.class);
    
    @ExceptionHandler(BusinessException.class)
    public ResponseEntity<ErrorResponse> handleBusinessException(BusinessException e) {
        logger.warn("Business exception: {}", e.getMessage());
        return ResponseEntity.badRequest()
            .body(ErrorResponse.builder()
                .code(e.getCode())
                .message(e.getMessage())
                .timestamp(LocalDateTime.now())
                .build());
    }
    
    @ExceptionHandler(ValidationException.class)
    public ResponseEntity<ErrorResponse> handleValidationException(ValidationException e) {
        logger.warn("Validation exception: {}", e.getMessage());
        return ResponseEntity.badRequest()
            .body(ErrorResponse.builder()
                .code("VALIDATION_ERROR")
                .message(e.getMessage())
                .details(e.getErrors())
                .timestamp(LocalDateTime.now())
                .build());
    }
    
    @ExceptionHandler(MethodArgumentNotValidException.class)
    public ResponseEntity<ErrorResponse> handleMethodArgumentNotValid(MethodArgumentNotValidException e) {
        List<String> errors = e.getBindingResult()
            .getFieldErrors()
            .stream()
            .map(error -> error.getField() + ": " + error.getDefaultMessage())
            .collect(Collectors.toList());
        
        return ResponseEntity.badRequest()
            .body(ErrorResponse.builder()
                .code("VALIDATION_ERROR")
                .message("Validation failed")
                .details(errors)
                .timestamp(LocalDateTime.now())
                .build());
    }
    
    @ExceptionHandler(DataAccessException.class)
    public ResponseEntity<ErrorResponse> handleDataAccessException(DataAccessException e) {
        logger.error("Database error", e);
        return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR)
            .body(ErrorResponse.builder()
                .code("DATABASE_ERROR")
                .message("Database operation failed")
                .timestamp(LocalDateTime.now())
                .build());
    }
    
    @ExceptionHandler(Exception.class)
    public ResponseEntity<ErrorResponse> handleGenericException(Exception e) {
        logger.error("Unexpected error", e);
        return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR)
            .body(ErrorResponse.builder()
                .code("INTERNAL_ERROR")
                .message("Internal server error")
                .timestamp(LocalDateTime.now())
                .build());
    }
    
    @ExceptionHandler(TooManyRequestsException.class)
    public ResponseEntity<ErrorResponse> handleTooManyRequests(TooManyRequestsException e) {
        logger.warn("Rate limit exceeded: {}", e.getMessage());
        return ResponseEntity.status(HttpStatus.TOO_MANY_REQUESTS)
            .header("Retry-After", "60")
            .body(ErrorResponse.builder()
                .code("RATE_LIMIT_EXCEEDED")
                .message("Too many requests")
                .timestamp(LocalDateTime.now())
                .build());
    }
}

// 异常响应类
@Data
@Builder
public class ErrorResponse {
    private String code;
    private String message;
    private List<String> details;
    private LocalDateTime timestamp;
    private String traceId;
    
    public static ErrorResponse of(String code, String message) {
        return ErrorResponse.builder()
            .code(code)
            .message(message)
            .timestamp(LocalDateTime.now())
            .traceId(MDC.get("traceId"))
            .build();
    }
}

// 业务异常基类
public abstract class BusinessException extends RuntimeException {
    private final String code;
    
    public BusinessException(String code, String message) {
        super(message);
        this.code = code;
    }
    
    public String getCode() {
        return code;
    }
}

// 具体业务异常
public class UserNotFoundException extends BusinessException {
    public UserNotFoundException(Long userId) {
        super("USER_NOT_FOUND", "User not found with id: " + userId);
    }
}

public class InsufficientBalanceException extends BusinessException {
    public InsufficientBalanceException(BigDecimal balance, BigDecimal required) {
        super("INSUFFICIENT_BALANCE", 
              String.format("Insufficient balance: %s, required: %s", balance, required));
    }
}
```

## 8. 配置管理检查

### 8.1 配置外部化 (Major)

#### 检查方法
- 检查是否存在硬编码的配置值
- 验证配置中心集成是否正确
- 确认环境变量使用是否规范
- 检查配置热更新机制

#### 检查标准
1. **禁止硬编码配置**: 所有配置项必须外部化
2. **配置中心使用**: 优先使用Nacos、Apollo等配置中心
3. **环境变量支持**: 支持通过环境变量覆盖配置
4. **配置热更新**: 支持运行时配置更新

#### 错误示例
```java
// ❌ 错误：硬编码配置
@Service
public class PaymentService {
    private static final String PAYMENT_URL = "https://api.payment.com/v1/pay";
    private static final int TIMEOUT = 5000;
    private static final String API_KEY = "sk_test_123456";
    
    public void processPayment() {
        // 硬编码配置，无法在不同环境使用
    }
}

// ❌ 错误：配置类缺少刷新支持
@Component
@ConfigurationProperties(prefix = "app")
public class AppConfig {
    private String name;
    private int timeout;
    // 缺少@RefreshScope，无法热更新
}
```

#### 正确示例
```java
// ✅ 正确：使用配置中心和环境变量
@Service
@RefreshScope
public class PaymentService {
    @Value("${payment.url:https://api.payment.com/v1/pay}")
    private String paymentUrl;
    
    @Value("${payment.timeout:5000}")
    private int timeout;
    
    @Value("${payment.api-key}")
    private String apiKey;
    
    public void processPayment() {
        // 使用外部化配置
    }
}

// ✅ 正确：支持热更新的配置类
@Component
@RefreshScope
@ConfigurationProperties(prefix = "app")
@Data
public class AppConfig {
    private String name;
    private int timeout;
    private Database database = new Database();
    
    @Data
    public static class Database {
        private String url;
        private String username;
        private String password;
        private int maxPoolSize = 10;
    }
}

// ✅ 正确：Nacos配置示例
@Configuration
@NacosPropertySource(dataId = "application.properties", autoRefreshed = true)
public class NacosConfig {
    
    @NacosValue(value = "${server.port:8080}", autoRefreshed = true)
    private int serverPort;
    
    @NacosValue(value = "${app.name}", autoRefreshed = true)
    private String appName;
}
```

### 8.2 环境隔离 (Critical)

#### 检查方法
- 检查不同环境的配置文件分离
- 验证敏感配置是否加密
- 确认配置版本管理机制
- 检查配置变更审计日志

#### 检查标准
1. **环境配置分离**: dev、test、prod环境配置完全分离
2. **敏感配置加密**: 密码、密钥等敏感信息必须加密
3. **配置版本管理**: 配置变更有版本记录和回滚能力
4. **变更审计**: 配置变更有完整的审计日志

#### 错误示例
```yaml
# ❌ 错误：所有环境共用一个配置文件
# application.yml
spring:
  datasource:
    url: jdbc:mysql://prod-db:3306/app  # 生产数据库
    username: root
    password: prod_password123  # 明文密码
  redis:
    host: prod-redis
    password: redis_password  # 明文密码

app:
  payment:
    api-key: sk_live_123456  # 生产API密钥
```

#### 正确示例
```yaml
# ✅ 正确：环境分离配置
# application-dev.yml
spring:
  datasource:
    url: jdbc:mysql://dev-db:3306/app
    username: ${DB_USERNAME:dev_user}
    password: ${DB_PASSWORD:ENC(encrypted_dev_password)}
  redis:
    host: dev-redis
    password: ${REDIS_PASSWORD:ENC(encrypted_redis_password)}

app:
  payment:
    api-key: ${PAYMENT_API_KEY:ENC(encrypted_dev_key)}

---
# application-prod.yml
spring:
  datasource:
    url: jdbc:mysql://prod-db:3306/app
    username: ${DB_USERNAME}
    password: ${DB_PASSWORD:ENC(encrypted_prod_password)}
  redis:
    host: prod-redis
    password: ${REDIS_PASSWORD:ENC(encrypted_redis_password)}

app:
  payment:
    api-key: ${PAYMENT_API_KEY:ENC(encrypted_prod_key)}
```

```java
// ✅ 正确：配置加密解密
@Configuration
@EnableEncryptableProperties
public class EncryptionConfig {
    
    @Bean("jasyptStringEncryptor")
    public StringEncryptor stringEncryptor() {
        PooledPBEStringEncryptor encryptor = new PooledPBEStringEncryptor();
        SimpleStringPBEConfig config = new SimpleStringPBEConfig();
        config.setPassword(System.getenv("JASYPT_ENCRYPTOR_PASSWORD"));
        config.setAlgorithm("PBEWITHHMACSHA512ANDAES_256");
        config.setKeyObtentionIterations("1000");
        config.setPoolSize("1");
        config.setProviderName("SunJCE");
        config.setSaltGeneratorClassName("org.jasypt.salt.RandomSaltGenerator");
        config.setIvGeneratorClassName("org.jasypt.iv.RandomIvGenerator");
        config.setStringOutputType("base64");
        encryptor.setConfig(config);
        return encryptor;
    }
}

// ✅ 正确：配置变更审计
@Component
@Slf4j
public class ConfigChangeListener {
    
    @EventListener
    public void handleConfigChange(RefreshEvent event) {
        log.info("Configuration refreshed: {}, changed keys: {}", 
                event.getEventDesc(), event.getKeys());
        
        // 记录配置变更审计日志
        auditService.recordConfigChange(
            event.getKeys(),
            getCurrentUser(),
            LocalDateTime.now()
        );
    }
    
    @NacosConfigListener(dataId = "application.properties")
    public void onConfigChange(String newContent) {
        log.info("Nacos config changed, new content length: {}", 
                newContent.length());
        
        // 记录Nacos配置变更
        auditService.recordNacosConfigChange(
            "application.properties",
            newContent,
            getCurrentUser(),
            LocalDateTime.now()
        );
    }
}
```

## 9. API设计检查

### 9.1 RESTful规范 (Major)

#### 检查方法
- 检查HTTP方法使用是否符合RESTful规范
- 验证URL设计是否遵循最佳实践
- 确认HTTP状态码使用是否正确
- 检查请求响应格式是否统一

#### 检查标准
1. **HTTP方法规范**: GET查询、POST创建、PUT更新、DELETE删除
2. **URL设计规范**: 使用名词、复数形式、层级关系清晰
3. **状态码规范**: 2xx成功、4xx客户端错误、5xx服务器错误
4. **格式统一**: 统一的请求响应格式和错误处理

#### 错误示例
```java
// ❌ 错误：HTTP方法使用不当
@RestController
public class UserController {
    
    // 错误：使用GET进行删除操作
    @GetMapping("/deleteUser/{id}")
    public String deleteUser(@PathVariable Long id) {
        userService.deleteUser(id);
        return "success";
    }
    
    // 错误：使用POST进行查询操作
    @PostMapping("/getUserById")
    public User getUserById(@RequestBody Map<String, Long> request) {
        return userService.findById(request.get("id"));
    }
    
    // 错误：URL设计不规范
    @GetMapping("/get_all_users_list")
    public List<User> getAllUsers() {
        return userService.findAll();
    }
    
    // 错误：状态码使用不当
    @PostMapping("/users")
    public ResponseEntity<String> createUser(@RequestBody User user) {
        try {
            userService.createUser(user);
            return ResponseEntity.ok("User created");  // 应该返回201
        } catch (Exception e) {
            return ResponseEntity.ok("Error: " + e.getMessage());  // 应该返回4xx或5xx
        }
    }
}
```

#### 正确示例
```java
// ✅ 正确：RESTful API设计
@RestController
@RequestMapping("/api/v1/users")
@Validated
public class UserController {
    
    @GetMapping
    public ResponseEntity<PageResponse<UserDTO>> getUsers(
            @RequestParam(defaultValue = "0") int page,
            @RequestParam(defaultValue = "20") int size,
            @RequestParam(required = false) String keyword) {
        
        PageResponse<UserDTO> users = userService.findUsers(page, size, keyword);
        return ResponseEntity.ok(users);
    }
    
    @GetMapping("/{id}")
    public ResponseEntity<UserDTO> getUser(@PathVariable Long id) {
        UserDTO user = userService.findById(id);
        return ResponseEntity.ok(user);
    }
    
    @PostMapping
    public ResponseEntity<UserDTO> createUser(@Valid @RequestBody CreateUserRequest request) {
        UserDTO user = userService.createUser(request);
        return ResponseEntity.status(HttpStatus.CREATED)
            .location(URI.create("/api/v1/users/" + user.getId()))
            .body(user);
    }
    
    @PutMapping("/{id}")
    public ResponseEntity<UserDTO> updateUser(
            @PathVariable Long id,
            @Valid @RequestBody UpdateUserRequest request) {
        UserDTO user = userService.updateUser(id, request);
        return ResponseEntity.ok(user);
    }
    
    @DeleteMapping("/{id}")
    public ResponseEntity<Void> deleteUser(@PathVariable Long id) {
        userService.deleteUser(id);
        return ResponseEntity.noContent().build();
    }
    
    // 子资源操作
    @GetMapping("/{userId}/orders")
    public ResponseEntity<List<OrderDTO>> getUserOrders(@PathVariable Long userId) {
        List<OrderDTO> orders = orderService.findByUserId(userId);
        return ResponseEntity.ok(orders);
    }
}

// ✅ 正确：统一响应格式
@Data
@Builder
public class ApiResponse<T> {
    private boolean success;
    private String message;
    private T data;
    private String timestamp;
    private String traceId;
    
    public static <T> ApiResponse<T> success(T data) {
        return ApiResponse.<T>builder()
            .success(true)
            .data(data)
            .timestamp(LocalDateTime.now().toString())
            .traceId(MDC.get("traceId"))
            .build();
    }
    
    public static <T> ApiResponse<T> error(String message) {
        return ApiResponse.<T>builder()
            .success(false)
            .message(message)
            .timestamp(LocalDateTime.now().toString())
            .traceId(MDC.get("traceId"))
            .build();
    }
}

@Data
public class PageResponse<T> {
    private List<T> content;
    private int page;
    private int size;
    private long totalElements;
    private int totalPages;
    private boolean first;
    private boolean last;
}
```

### 9.2 版本控制 (Major)

#### 检查方法
- 检查API版本控制策略是否明确
- 验证向后兼容性处理
- 确认废弃API的处理机制
- 检查版本文档维护情况

#### 检查标准
1. **版本控制策略**: 使用URL路径、请求头或参数进行版本控制
2. **向后兼容**: 新版本保持向后兼容，渐进式变更
3. **废弃处理**: 废弃API有明确的时间表和迁移指南
4. **文档维护**: 每个版本都有完整的API文档

#### 错误示例
```java
// ❌ 错误：没有版本控制
@RestController
@RequestMapping("/api/users")
public class UserController {
    
    // 直接修改现有API，破坏向后兼容性
    @GetMapping("/{id}")
    public UserV2DTO getUser(@PathVariable Long id) {
        // 返回格式变更，破坏了向后兼容性
        return userService.findByIdV2(id);
    }
}

// ❌ 错误：版本控制不一致
@RestController
public class OrderController {
    
    @GetMapping("/v1/orders")  // URL版本控制
    public List<Order> getOrdersV1() {
        return orderService.findAll();
    }
    
    @GetMapping("/orders")  // 请求头版本控制，不一致
    public List<Order> getOrdersV2(@RequestHeader("API-Version") String version) {
        return orderService.findAllV2();
    }
}
```

#### 正确示例
```java
// ✅ 正确：URL路径版本控制
@RestController
@RequestMapping("/api/v1/users")
public class UserV1Controller {
    
    @GetMapping("/{id}")
    public ResponseEntity<UserV1DTO> getUser(@PathVariable Long id) {
        UserV1DTO user = userService.findByIdV1(id);
        return ResponseEntity.ok(user);
    }
}

@RestController
@RequestMapping("/api/v2/users")
public class UserV2Controller {
    
    @GetMapping("/{id}")
    public ResponseEntity<UserV2DTO> getUser(@PathVariable Long id) {
        UserV2DTO user = userService.findByIdV2(id);
        return ResponseEntity.ok(user);
    }
    
    // 新增功能
    @GetMapping("/{id}/profile")
    public ResponseEntity<UserProfileDTO> getUserProfile(@PathVariable Long id) {
        UserProfileDTO profile = userService.getUserProfile(id);
        return ResponseEntity.ok(profile);
    }
}

// ✅ 正确：请求头版本控制
@RestController
@RequestMapping("/api/users")
public class UserController {
    
    @GetMapping("/{id}")
    public ResponseEntity<?> getUser(
            @PathVariable Long id,
            @RequestHeader(value = "API-Version", defaultValue = "v1") String version) {
        
        switch (version) {
            case "v1":
                return ResponseEntity.ok(userService.findByIdV1(id));
            case "v2":
                return ResponseEntity.ok(userService.findByIdV2(id));
            default:
                return ResponseEntity.badRequest()
                    .body("Unsupported API version: " + version);
        }
    }
}

// ✅ 正确：废弃API处理
@RestController
@RequestMapping("/api/v1/orders")
public class OrderV1Controller {
    
    @Deprecated
    @GetMapping
    public ResponseEntity<List<OrderDTO>> getOrders(
            HttpServletResponse response) {
        
        // 添加废弃警告头
        response.setHeader("Warning", "299 - \"Deprecated API\"");
        response.setHeader("Sunset", "2024-12-31");
        response.setHeader("Link", "</api/v2/orders>; rel=\"successor-version\"");
        
        List<OrderDTO> orders = orderService.findAllV1();
        return ResponseEntity.ok(orders);
    }
}

// ✅ 正确：版本兼容性处理
@Component
public class ApiVersionCompatibilityService {
    
    public UserDTO convertToVersion(User user, String version) {
        switch (version) {
            case "v1":
                return UserV1DTO.builder()
                    .id(user.getId())
                    .name(user.getName())
                    .email(user.getEmail())
                    .build();
                    
            case "v2":
                return UserV2DTO.builder()
                    .id(user.getId())
                    .fullName(user.getName())
                    .emailAddress(user.getEmail())
                    .profile(user.getProfile())
                    .createdAt(user.getCreatedAt())
                    .build();
                    
            default:
                throw new UnsupportedApiVersionException(version);
        }
    }
}
```

### 9.3 参数验证 (Critical)

#### 检查方法
- 检查是否使用@Valid注解进行参数验证
- 验证自定义校验器的实现
- 确认参数格式验证的完整性
- 检查业务规则验证的实现

#### 检查标准
1. **基础验证**: 使用@Valid和Bean Validation注解
2. **自定义验证**: 实现复杂业务规则的自定义校验器
3. **格式验证**: 邮箱、手机号、身份证等格式验证
4. **业务验证**: 唯一性、依赖关系等业务规则验证

#### 错误示例
```java
// ❌ 错误：缺少参数验证
@RestController
public class UserController {
    
    @PostMapping("/users")
    public ResponseEntity<User> createUser(@RequestBody User user) {
        // 没有验证，可能接收到无效数据
        if (user.getName() == null || user.getName().isEmpty()) {
            throw new IllegalArgumentException("Name is required");
        }
        if (user.getEmail() == null || !user.getEmail().contains("@")) {
            throw new IllegalArgumentException("Valid email is required");
        }
        // 手动验证，代码冗余且容易遗漏
        return ResponseEntity.ok(userService.createUser(user));
    }
}

// ❌ 错误：验证注解使用不当
public class CreateUserRequest {
    @NotNull  // 应该使用@NotBlank
    private String name;
    
    @Email  // 缺少@NotBlank，可能为空
    private String email;
    
    @Min(0)  // 年龄验证不够严格
    private Integer age;
    
    // 缺少getter/setter
}
```

#### 正确示例
```java
// ✅ 正确：完整的参数验证
@RestController
@RequestMapping("/api/v1/users")
@Validated
public class UserController {
    
    @PostMapping
    public ResponseEntity<UserDTO> createUser(
            @Valid @RequestBody CreateUserRequest request) {
        UserDTO user = userService.createUser(request);
        return ResponseEntity.status(HttpStatus.CREATED).body(user);
    }
    
    @PutMapping("/{id}")
    public ResponseEntity<UserDTO> updateUser(
            @PathVariable @Min(1) Long id,
            @Valid @RequestBody UpdateUserRequest request) {
        UserDTO user = userService.updateUser(id, request);
        return ResponseEntity.ok(user);
    }
    
    @GetMapping
    public ResponseEntity<PageResponse<UserDTO>> getUsers(
            @RequestParam(defaultValue = "0") @Min(0) int page,
            @RequestParam(defaultValue = "20") @Range(min = 1, max = 100) int size,
            @RequestParam(required = false) @Length(max = 50) String keyword) {
        
        PageResponse<UserDTO> users = userService.findUsers(page, size, keyword);
        return ResponseEntity.ok(users);
    }
}

// ✅ 正确：请求DTO验证
@Data
@Builder
public class CreateUserRequest {
    
    @NotBlank(message = "用户名不能为空")
    @Length(min = 2, max = 50, message = "用户名长度必须在2-50个字符之间")
    @Pattern(regexp = "^[a-zA-Z0-9\u4e00-\u9fa5_-]+$", message = "用户名只能包含字母、数字、中文、下划线和横线")
    private String name;
    
    @NotBlank(message = "邮箱不能为空")
    @Email(message = "邮箱格式不正确")
    @UniqueEmail  // 自定义验证注解
    private String email;
    
    @NotBlank(message = "手机号不能为空")
    @Pattern(regexp = "^1[3-9]\\d{9}$", message = "手机号格式不正确")
    private String phone;
    
    @NotNull(message = "年龄不能为空")
    @Range(min = 18, max = 120, message = "年龄必须在18-120之间")
    private Integer age;
    
    @NotNull(message = "性别不能为空")
    @EnumValue(enumClass = Gender.class, message = "性别值不正确")
    private String gender;
    
    @Valid
    @NotNull(message = "地址信息不能为空")
    private AddressRequest address;
    
    @Size(max = 5, message = "标签数量不能超过5个")
    private List<@NotBlank @Length(max = 20) String> tags;
}

@Data
public class AddressRequest {
    
    @NotBlank(message = "省份不能为空")
    private String province;
    
    @NotBlank(message = "城市不能为空")
    private String city;
    
    @NotBlank(message = "详细地址不能为空")
    @Length(max = 200, message = "详细地址不能超过200个字符")
    private String detail;
    
    @Pattern(regexp = "^\\d{6}$", message = "邮政编码格式不正确")
    private String zipCode;
}

// ✅ 正确：自定义验证注解
@Target({ElementType.FIELD})
@Retention(RetentionPolicy.RUNTIME)
@Constraint(validatedBy = UniqueEmailValidator.class)
public @interface UniqueEmail {
    String message() default "邮箱已存在";
    Class<?>[] groups() default {};
    Class<? extends Payload>[] payload() default {};
}

@Component
public class UniqueEmailValidator implements ConstraintValidator<UniqueEmail, String> {
    
    @Autowired
    private UserService userService;
    
    @Override
    public boolean isValid(String email, ConstraintValidatorContext context) {
        if (email == null || email.isEmpty()) {
            return true;  // 空值由@NotBlank处理
        }
        return !userService.existsByEmail(email);
    }
}

@Target({ElementType.FIELD})
@Retention(RetentionPolicy.RUNTIME)
@Constraint(validatedBy = EnumValueValidator.class)
public @interface EnumValue {
    String message() default "枚举值不正确";
    Class<?>[] groups() default {};
    Class<? extends Payload>[] payload() default {};
    Class<? extends Enum<?>> enumClass();
}

public class EnumValueValidator implements ConstraintValidator<EnumValue, String> {
    
    private Class<? extends Enum<?>> enumClass;
    
    @Override
    public void initialize(EnumValue annotation) {
        this.enumClass = annotation.enumClass();
    }
    
    @Override
    public boolean isValid(String value, ConstraintValidatorContext context) {
        if (value == null) {
            return true;
        }
        
        Enum<?>[] enumConstants = enumClass.getEnumConstants();
        for (Enum<?> enumConstant : enumConstants) {
            if (enumConstant.name().equals(value)) {
                return true;
            }
        }
        return false;
    }
}

// ✅ 正确：分组验证
public interface CreateGroup {}
public interface UpdateGroup {}

@Data
public class UserRequest {
    
    @Null(groups = CreateGroup.class, message = "创建时ID必须为空")
    @NotNull(groups = UpdateGroup.class, message = "更新时ID不能为空")
    private Long id;
    
    @NotBlank(groups = {CreateGroup.class, UpdateGroup.class})
    private String name;
    
    @NotBlank(groups = CreateGroup.class, message = "创建时密码不能为空")
    @Length(min = 8, groups = {CreateGroup.class, UpdateGroup.class})
    private String password;
}

@PostMapping
public ResponseEntity<UserDTO> createUser(
        @Validated(CreateGroup.class) @RequestBody UserRequest request) {
    return ResponseEntity.ok(userService.createUser(request));
}

@PutMapping("/{id}")
public ResponseEntity<UserDTO> updateUser(
        @PathVariable Long id,
        @Validated(UpdateGroup.class) @RequestBody UserRequest request) {
    return ResponseEntity.ok(userService.updateUser(id, request));
}
```

## 10. 测试相关检查

### 10.1 单元测试 (Major)

#### 10.1.1 测试覆盖率达标（建议80%+）
- **检查方法**: 使用JaCoCo等工具检查代码覆盖率，重点关注业务逻辑覆盖
- **检查标准**: 行覆盖率≥80%，分支覆盖率≥70%，核心业务逻辑100%覆盖
- **不正确实例**:
```java
// 错误示例 - 测试覆盖不充分
@Service
public class UserService {
    public User createUser(User user) {
        if (user == null) {
            throw new IllegalArgumentException("User cannot be null");
        }
        if (user.getAge() < 0) {
            throw new IllegalArgumentException("Age cannot be negative");
        }
        if (user.getAge() > 150) {
            throw new IllegalArgumentException("Age cannot exceed 150");
        }
        return userRepository.save(user);
    }
}

// 错误的测试 - 只测试正常情况
@Test
public void testCreateUser() {
    User user = new User("John", 25);
    User result = userService.createUser(user);
    assertNotNull(result);
}

// 正确示例 - 全面的测试覆盖
@ExtendWith(MockitoExtension.class)
class UserServiceTest {
    @Mock
    private UserRepository userRepository;
    
    @InjectMocks
    private UserService userService;
    
    @Test
    @DisplayName("创建用户 - 正常情况")
    void createUser_ValidUser_Success() {
        // Given
        User user = new User("John", 25);
        User savedUser = new User("John", 25);
        savedUser.setId(1L);
        when(userRepository.save(user)).thenReturn(savedUser);
        
        // When
        User result = userService.createUser(user);
        
        // Then
        assertThat(result.getId()).isEqualTo(1L);
        assertThat(result.getName()).isEqualTo("John");
        verify(userRepository).save(user);
    }
    
    @Test
    @DisplayName("创建用户 - 用户为null")
    void createUser_NullUser_ThrowsException() {
        // When & Then
        assertThatThrownBy(() -> userService.createUser(null))
            .isInstanceOf(IllegalArgumentException.class)
            .hasMessage("User cannot be null");
        
        verify(userRepository, never()).save(any());
    }
    
    @Test
    @DisplayName("创建用户 - 年龄为负数")
    void createUser_NegativeAge_ThrowsException() {
        // Given
        User user = new User("John", -1);
        
        // When & Then
        assertThatThrownBy(() -> userService.createUser(user))
            .isInstanceOf(IllegalArgumentException.class)
            .hasMessage("Age cannot be negative");
    }
    
    @Test
    @DisplayName("创建用户 - 年龄超过150")
    void createUser_AgeExceeds150_ThrowsException() {
        // Given
        User user = new User("John", 151);
        
        // When & Then
        assertThatThrownBy(() -> userService.createUser(user))
            .isInstanceOf(IllegalArgumentException.class)
            .hasMessage("Age cannot exceed 150");
    }
    
    @ParameterizedTest
    @ValueSource(ints = {0, 1, 18, 65, 150})
    @DisplayName("创建用户 - 边界年龄值")
    void createUser_BoundaryAges_Success(int age) {
        // Given
        User user = new User("John", age);
        when(userRepository.save(any())).thenReturn(user);
        
        // When & Then
        assertDoesNotThrow(() -> userService.createUser(user));
    }
}
```

#### 10.1.2 边界条件测试
- **检查方法**: 检查测试用例是否包含边界值、临界值测试
- **检查标准**: 必须测试最小值、最大值、零值、空值等边界条件
- **不正确实例**:
```java
// 错误示例 - 缺少边界条件测试
@Test
void testCalculateDiscount() {
    // 只测试中间值
    BigDecimal result = discountService.calculateDiscount(new BigDecimal("100"));
    assertThat(result).isEqualTo(new BigDecimal("10"));
}

// 正确示例 - 完整的边界条件测试
@ParameterizedTest
@CsvSource({
    "0, 0",           // 最小值
    "0.01, 0",        // 最小正值
    "99.99, 0",       // 临界值下限
    "100, 10",        // 临界值
    "100.01, 10.001", // 临界值上限
    "1000, 100",      // 中间值
    "9999.99, 999.999", // 最大值下限
    "10000, 1000"     // 最大值
})
void calculateDiscount_BoundaryValues(String amount, String expected) {
    BigDecimal result = discountService.calculateDiscount(new BigDecimal(amount));
    assertThat(result).isEqualByComparingTo(new BigDecimal(expected));
}

@Test
void calculateDiscount_NullAmount_ThrowsException() {
    assertThatThrownBy(() -> discountService.calculateDiscount(null))
        .isInstanceOf(IllegalArgumentException.class);
}

@Test
void calculateDiscount_NegativeAmount_ThrowsException() {
    assertThatThrownBy(() -> discountService.calculateDiscount(new BigDecimal("-1")))
        .isInstanceOf(IllegalArgumentException.class);
}
```

#### 10.1.3 异常情况测试
- **检查方法**: 检查是否测试了所有可能的异常场景
- **检查标准**: 每个可能抛出异常的方法都要有对应的异常测试
- **不正确实例**:
```java
// 错误示例 - 缺少异常测试
@Test
void testGetUserById() {
    User user = userService.getUserById(1L);
    assertNotNull(user);
}

// 正确示例 - 完整的异常测试
@Test
void getUserById_ExistingUser_ReturnsUser() {
    // Given
    User expectedUser = new User("John", 25);
    when(userRepository.findById(1L)).thenReturn(Optional.of(expectedUser));
    
    // When
    User result = userService.getUserById(1L);
    
    // Then
    assertThat(result).isEqualTo(expectedUser);
}

@Test
void getUserById_NonExistingUser_ThrowsUserNotFoundException() {
    // Given
    when(userRepository.findById(999L)).thenReturn(Optional.empty());
    
    // When & Then
    assertThatThrownBy(() -> userService.getUserById(999L))
        .isInstanceOf(UserNotFoundException.class)
        .hasMessage("User not found with id: 999");
}

@Test
void getUserById_DatabaseError_ThrowsDataAccessException() {
    // Given
    when(userRepository.findById(1L))
        .thenThrow(new DataAccessException("Database connection failed") {});
    
    // When & Then
    assertThatThrownBy(() -> userService.getUserById(1L))
        .isInstanceOf(DataAccessException.class);
}

@Test
void getUserById_NullId_ThrowsIllegalArgumentException() {
    // When & Then
    assertThatThrownBy(() -> userService.getUserById(null))
        .isInstanceOf(IllegalArgumentException.class)
        .hasMessage("User ID cannot be null");
}
```

#### 10.1.4 Mock使用正确
- **检查方法**: 检查Mock对象的使用是否合理，验证交互是否正确
- **检查标准**: 只Mock外部依赖，验证重要的交互，避免过度Mock
- **不正确实例**:
```java
// 错误示例 - Mock使用不当
@Test
void testProcessOrder() {
    // 错误：Mock了被测试的类本身
    OrderService orderService = mock(OrderService.class);
    when(orderService.processOrder(any())).thenReturn(true);
    
    boolean result = orderService.processOrder(new Order());
    assertTrue(result);  // 这个测试没有意义
}

// 错误：没有验证重要的交互
@Test
void testSendNotification() {
    notificationService.sendNotification("test message");
    // 没有验证是否真的调用了邮件服务
}

// 正确示例 - 正确使用Mock
@ExtendWith(MockitoExtension.class)
class OrderServiceTest {
    @Mock
    private OrderRepository orderRepository;
    
    @Mock
    private PaymentService paymentService;
    
    @Mock
    private NotificationService notificationService;
    
    @InjectMocks
    private OrderService orderService;
    
    @Test
    void processOrder_ValidOrder_Success() {
        // Given
        Order order = new Order("ORDER001", new BigDecimal("100"));
        when(paymentService.processPayment(order)).thenReturn(true);
        when(orderRepository.save(any(Order.class))).thenReturn(order);
        
        // When
        boolean result = orderService.processOrder(order);
        
        // Then
        assertTrue(result);
        
        // 验证重要的交互
        verify(paymentService).processPayment(order);
        verify(orderRepository).save(argThat(savedOrder -> 
            savedOrder.getStatus() == OrderStatus.COMPLETED));
        verify(notificationService).sendOrderConfirmation(order);
    }
    
    @Test
    void processOrder_PaymentFailed_OrderNotSaved() {
        // Given
        Order order = new Order("ORDER001", new BigDecimal("100"));
        when(paymentService.processPayment(order)).thenReturn(false);
        
        // When
        boolean result = orderService.processOrder(order);
        
        // Then
        assertFalse(result);
        
        // 验证支付失败时的行为
        verify(paymentService).processPayment(order);
        verify(orderRepository, never()).save(any());
        verify(notificationService, never()).sendOrderConfirmation(any());
    }
    
    @Test
    void processOrder_PaymentServiceThrowsException_HandledGracefully() {
        // Given
        Order order = new Order("ORDER001", new BigDecimal("100"));
        when(paymentService.processPayment(order))
            .thenThrow(new PaymentException("Payment gateway unavailable"));
        
        // When & Then
        assertThatThrownBy(() -> orderService.processOrder(order))
            .isInstanceOf(OrderProcessingException.class)
            .hasCauseInstanceOf(PaymentException.class);
        
        verify(orderRepository, never()).save(any());
    }
}
```

### 10.2 集成测试 (Major)

#### 10.2.1 数据库集成测试
- **检查方法**: 使用@DataJpaTest或TestContainers检查数据库操作
- **检查标准**: 测试真实的数据库交互，包括事务、约束、级联操作
- **不正确实例**:
```java
// 错误示例 - 使用H2内存数据库但与生产环境差异太大
@DataJpaTest
class UserRepositoryTest {
    @Autowired
    private TestEntityManager entityManager;
    
    @Autowired
    private UserRepository userRepository;
    
    @Test
    void testFindByEmail() {
        // 这个测试在H2上通过，但在MySQL上可能失败
        User user = new User("john@example.com", "John");
        entityManager.persistAndFlush(user);
        
        Optional<User> found = userRepository.findByEmail("JOHN@EXAMPLE.COM");
        assertTrue(found.isPresent());  // H2大小写不敏感，MySQL可能敏感
    }
}

// 正确示例 - 使用TestContainers测试真实数据库
@SpringBootTest
@Testcontainers
class UserRepositoryIntegrationTest {
    
    @Container
    static MySQLContainer<?> mysql = new MySQLContainer<>("mysql:8.0")
            .withDatabaseName("testdb")
            .withUsername("test")
            .withPassword("test");
    
    @Autowired
    private UserRepository userRepository;
    
    @Autowired
    private TestEntityManager entityManager;
    
    @Test
    @Transactional
    @Rollback
    void findByEmail_ExactCase_ReturnsUser() {
        // Given
        User user = new User("john@example.com", "John");
        entityManager.persistAndFlush(user);
        
        // When
        Optional<User> found = userRepository.findByEmail("john@example.com");
        
        // Then
        assertThat(found).isPresent();
        assertThat(found.get().getName()).isEqualTo("John");
    }
    
    @Test
    @Transactional
    @Rollback
    void findByEmail_DifferentCase_ReturnsEmpty() {
        // Given
        User user = new User("john@example.com", "John");
        entityManager.persistAndFlush(user);
        
        // When
        Optional<User> found = userRepository.findByEmail("JOHN@EXAMPLE.COM");
        
        // Then
        assertThat(found).isEmpty();  // MySQL区分大小写
    }
    
    @Test
    @Transactional
    @Rollback
    void saveUser_DuplicateEmail_ThrowsConstraintViolation() {
        // Given
        User user1 = new User("john@example.com", "John");
        User user2 = new User("john@example.com", "Jane");
        
        userRepository.save(user1);
        entityManager.flush();
        
        // When & Then
        assertThatThrownBy(() -> {
            userRepository.save(user2);
            entityManager.flush();
        }).isInstanceOf(DataIntegrityViolationException.class);
    }
    
    @Test
    @Transactional
    @Rollback
    void deleteUser_WithOrders_CascadeDelete() {
        // Given
        User user = new User("john@example.com", "John");
        Order order1 = new Order("ORDER001", user);
        Order order2 = new Order("ORDER002", user);
        user.addOrder(order1);
        user.addOrder(order2);
        
        User savedUser = userRepository.save(user);
        entityManager.flush();
        
        // When
        userRepository.delete(savedUser);
        entityManager.flush();
        
        // Then
        assertThat(userRepository.findById(savedUser.getId())).isEmpty();
        // 验证级联删除
        assertThat(entityManager.find(Order.class, order1.getId())).isNull();
        assertThat(entityManager.find(Order.class, order2.getId())).isNull();
    }
    
    @DynamicPropertySource
    static void configureProperties(DynamicPropertyRegistry registry) {
        registry.add("spring.datasource.url", mysql::getJdbcUrl);
        registry.add("spring.datasource.username", mysql::getUsername);
        registry.add("spring.datasource.password", mysql::getPassword);
    }
}
```

#### ß10.2.2 外部服务集成测试
- **检查方法**: 使用WireMock或TestContainers模拟外部服务
- **检查标准**: 测试网络调用、超时、重试、熔断等机制
- **不正确实例**:
```java
// 错误示例 - 直接调用真实的外部服务
@Test
void testGetUserFromExternalService() {
    // 错误：依赖真实的外部服务，测试不稳定
    ExternalUser user = externalUserService.getUser("123");
    assertNotNull(user);
}

// 正确示例 - 使用WireMock模拟外部服务
@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
class ExternalUserServiceIntegrationTest {
    
    @RegisterExtension
    static WireMockExtension wireMock = WireMockExtension.newInstance()
            .options(wireMockConfig().port(8089))
            .build();
    
    @Autowired
    private ExternalUserService externalUserService;
    
    @Test
    void getUser_ValidId_ReturnsUser() {
        // Given
        String userId = "123";
        String responseBody = """
            {
                "id": "123",
                "name": "John Doe",
                "email": "john@example.com"
            }
            """;
        
        wireMock.stubFor(get(urlEqualTo("/users/" + userId))
                .willReturn(aResponse()
                        .withStatus(200)
                        .withHeader("Content-Type", "application/json")
                        .withBody(responseBody)));
        
        // When
        ExternalUser user = externalUserService.getUser(userId);
        
        // Then
        assertThat(user.getId()).isEqualTo("123");
        assertThat(user.getName()).isEqualTo("John Doe");
        assertThat(user.getEmail()).isEqualTo("john@example.com");
        
        // 验证请求
        wireMock.verify(getRequestedFor(urlEqualTo("/users/" + userId))
                .withHeader("Authorization", matching("Bearer .*")));
    }
    
    @Test
    void getUser_ServiceUnavailable_ThrowsServiceException() {
        // Given
        String userId = "123";
        wireMock.stubFor(get(urlEqualTo("/users/" + userId))
                .willReturn(aResponse()
                        .withStatus(503)
                        .withBody("Service Unavailable")));
        
        // When & Then
        assertThatThrownBy(() -> externalUserService.getUser(userId))
                .isInstanceOf(ExternalServiceException.class)
                .hasMessageContaining("Service unavailable");
    }
    
    @Test
    void getUser_Timeout_ThrowsTimeoutException() {
        // Given
        String userId = "123";
        wireMock.stubFor(get(urlEqualTo("/users/" + userId))
                .willReturn(aResponse()
                        .withStatus(200)
                        .withFixedDelay(5000)  // 5秒延迟
                        .withBody("{}"));
        
        // When & Then
        assertThatThrownBy(() -> externalUserService.getUser(userId))
                .isInstanceOf(TimeoutException.class);
    }
    
    @Test
    void getUser_RetryOnFailure_EventuallySucceeds() {
        // Given
        String userId = "123";
        String responseBody = "{\"id\": \"123\", \"name\": \"John\"}";
        
        // 前两次失败，第三次成功
        wireMock.stubFor(get(urlEqualTo("/users/" + userId))
                .inScenario("Retry Scenario")
                .whenScenarioStateIs(Scenario.STARTED)
                .willReturn(aResponse().withStatus(500))
                .willSetStateTo("First Retry"));
        
        wireMock.stubFor(get(urlEqualTo("/users/" + userId))
                .inScenario("Retry Scenario")
                .whenScenarioStateIs("First Retry")
                .willReturn(aResponse().withStatus(500))
                .willSetStateTo("Second Retry"));
        
        wireMock.stubFor(get(urlEqualTo("/users/" + userId))
                .inScenario("Retry Scenario")
                .whenScenarioStateIs("Second Retry")
                .willReturn(aResponse()
                        .withStatus(200)
                        .withHeader("Content-Type", "application/json")
                        .withBody(responseBody)));
        
        // When
        ExternalUser user = externalUserService.getUser(userId);
        
        // Then
        assertThat(user.getId()).isEqualTo("123");
        
        // 验证重试了3次
        wireMock.verify(3, getRequestedFor(urlEqualTo("/users/" + userId)));
    }
    
    @TestConfiguration
    static class TestConfig {
        @Bean
        @Primary
        public RestTemplate restTemplate() {
            RestTemplate restTemplate = new RestTemplate();
            // 配置超时
            HttpComponentsClientHttpRequestFactory factory = 
                new HttpComponentsClientHttpRequestFactory();
            factory.setConnectTimeout(2000);
            factory.setReadTimeout(3000);
            restTemplate.setRequestFactory(factory);
            return restTemplate;
        }
    }
}
```

#### 10.2.3 端到端测试
- **检查方法**: 使用@SpringBootTest进行完整的应用程序测试
- **检查标准**: 测试完整的业务流程，包括用户界面到数据库的整个链路
- **不正确实例**:
```java
// 错误示例 - 端到端测试过于简单
@SpringBootTest
class UserControllerE2ETest {
    @Test
    void testCreateUser() {
        // 测试过于简单，没有验证完整流程
        ResponseEntity<String> response = restTemplate.postForEntity(
            "/api/users", new User("John", 25), String.class);
        assertEquals(200, response.getStatusCodeValue());
    }
}

// 正确示例 - 完整的端到端测试
@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
@Testcontainers
@Transactional
class UserManagementE2ETest {
    
    @Container
    static MySQLContainer<?> mysql = new MySQLContainer<>("mysql:8.0")
            .withDatabaseName("testdb")
            .withUsername("test")
            .withPassword("test");
    
    @Container
    static GenericContainer<?> redis = new GenericContainer<>("redis:6.2")
            .withExposedPorts(6379);
    
    @Autowired
    private TestRestTemplate restTemplate;
    
    @Autowired
    private UserRepository userRepository;
    
    @Autowired
    private RedisTemplate<String, Object> redisTemplate;
    
    @Test
    void userLifecycle_CompleteFlow_Success() {
        // 1. 创建用户
        CreateUserRequest createRequest = new CreateUserRequest(
            "john@example.com", "John Doe", 25);
        
        ResponseEntity<UserResponse> createResponse = restTemplate.postForEntity(
            "/api/users", createRequest, UserResponse.class);
        
        assertThat(createResponse.getStatusCode()).isEqualTo(HttpStatus.CREATED);
        assertThat(createResponse.getBody().getEmail()).isEqualTo("john@example.com");
        
        Long userId = createResponse.getBody().getId();
        
        // 2. 验证用户已保存到数据库
        Optional<User> savedUser = userRepository.findById(userId);
        assertThat(savedUser).isPresent();
        assertThat(savedUser.get().getEmail()).isEqualTo("john@example.com");
        
        // 3. 验证缓存已更新
        User cachedUser = (User) redisTemplate.opsForValue().get("user:" + userId);
        assertThat(cachedUser).isNotNull();
        assertThat(cachedUser.getEmail()).isEqualTo("john@example.com");
        
        // 4. 查询用户
        ResponseEntity<UserResponse> getResponse = restTemplate.getForEntity(
            "/api/users/" + userId, UserResponse.class);
        
        assertThat(getResponse.getStatusCode()).isEqualTo(HttpStatus.OK);
        assertThat(getResponse.getBody().getId()).isEqualTo(userId);
        
        // 5. 更新用户
        UpdateUserRequest updateRequest = new UpdateUserRequest(
            "John Smith", 26);
        
        ResponseEntity<UserResponse> updateResponse = restTemplate.exchange(
            "/api/users/" + userId, HttpMethod.PUT, 
            new HttpEntity<>(updateRequest), UserResponse.class);
        
        assertThat(updateResponse.getStatusCode()).isEqualTo(HttpStatus.OK);
        assertThat(updateResponse.getBody().getName()).isEqualTo("John Smith");
        assertThat(updateResponse.getBody().getAge()).isEqualTo(26);
        
        // 6. 验证数据库已更新
        User updatedUser = userRepository.findById(userId).orElseThrow();
        assertThat(updatedUser.getName()).isEqualTo("John Smith");
        assertThat(updatedUser.getAge()).isEqualTo(26);
        
        // 7. 验证缓存已失效并重新加载
        User updatedCachedUser = (User) redisTemplate.opsForValue().get("user:" + userId);
        assertThat(updatedCachedUser.getName()).isEqualTo("John Smith");
        
        // 8. 删除用户
        ResponseEntity<Void> deleteResponse = restTemplate.exchange(
            "/api/users/" + userId, HttpMethod.DELETE, null, Void.class);
        
        assertThat(deleteResponse.getStatusCode()).isEqualTo(HttpStatus.NO_CONTENT);
        
        // 9. 验证用户已从数据库删除
        assertThat(userRepository.findById(userId)).isEmpty();
        
        // 10. 验证缓存已清除
        assertThat(redisTemplate.opsForValue().get("user:" + userId)).isNull();
        
        // 11. 验证删除后查询返回404
        ResponseEntity<String> notFoundResponse = restTemplate.getForEntity(
            "/api/users/" + userId, String.class);
        assertThat(notFoundResponse.getStatusCode()).isEqualTo(HttpStatus.NOT_FOUND);
    }
    
    @Test
    void createUser_ValidationError_ReturnsBadRequest() {
        // 测试验证错误场景
        CreateUserRequest invalidRequest = new CreateUserRequest(
            "invalid-email", "", -1);  // 无效邮箱、空名称、负年龄
        
        ResponseEntity<ErrorResponse> response = restTemplate.postForEntity(
            "/api/users", invalidRequest, ErrorResponse.class);
        
        assertThat(response.getStatusCode()).isEqualTo(HttpStatus.BAD_REQUEST);
        assertThat(response.getBody().getErrors()).hasSize(3);
        assertThat(response.getBody().getErrors())
            .extracting("field")
            .containsExactlyInAnyOrder("email", "name", "age");
    }
    
    @Test
    void createUser_DuplicateEmail_ReturnsConflict() {
        // 先创建一个用户
        CreateUserRequest firstRequest = new CreateUserRequest(
            "john@example.com", "John Doe", 25);
        restTemplate.postForEntity("/api/users", firstRequest, UserResponse.class);
        
        // 尝试创建相同邮箱的用户
        CreateUserRequest duplicateRequest = new CreateUserRequest(
            "john@example.com", "Jane Doe", 30);
        
        ResponseEntity<ErrorResponse> response = restTemplate.postForEntity(
            "/api/users", duplicateRequest, ErrorResponse.class);
        
        assertThat(response.getStatusCode()).isEqualTo(HttpStatus.CONFLICT);
        assertThat(response.getBody().getMessage())
            .contains("Email already exists");
    }
    
    @DynamicPropertySource
    static void configureProperties(DynamicPropertyRegistry registry) {
        registry.add("spring.datasource.url", mysql::getJdbcUrl);
        registry.add("spring.datasource.username", mysql::getUsername);
        registry.add("spring.datasource.password", mysql::getPassword);
        registry.add("spring.redis.host", redis::getHost);
        registry.add("spring.redis.port", redis::getFirstMappedPort);
    }
}
```

#### 10.2.4 性能测试
- **检查方法**: 使用JMeter、Gatling或自定义性能测试
- **检查标准**: 验证响应时间、吞吐量、资源使用率等性能指标
- **不正确实例**:
```java
// 错误示例 - 性能测试不充分
@Test
void testPerformance() {
    long start = System.currentTimeMillis();
    userService.createUser(new User("John", 25));
    long end = System.currentTimeMillis();
    assertTrue(end - start < 1000);  // 过于简单的性能测试
}

// 正确示例 - 完整的性能测试
@SpringBootTest
@Testcontainers
class UserServicePerformanceTest {
    
    @Container
    static MySQLContainer<?> mysql = new MySQLContainer<>("mysql:8.0")
            .withDatabaseName("testdb")
            .withUsername("test")
            .withPassword("test");
    
    @Autowired
    private UserService userService;
    
    @Autowired
    private UserRepository userRepository;
    
    private MeterRegistry meterRegistry;
    
    @BeforeEach
    void setUp() {
        meterRegistry = new SimpleMeterRegistry();
        Metrics.addRegistry(meterRegistry);
    }
    
    @Test
    void createUser_SingleThread_MeetsPerformanceRequirements() {
        // 单线程性能测试
        Timer.Sample sample = Timer.start(meterRegistry);
        
        for (int i = 0; i < 100; i++) {
            User user = new User("user" + i + "@example.com", "User " + i, 25);
            userService.createUser(user);
        }
        
        sample.stop(Timer.builder("user.creation.time")
                .description("Time taken to create users")
                .register(meterRegistry));
        
        Timer timer = meterRegistry.get("user.creation.time").timer();
        
        // 验证性能指标
        assertThat(timer.count()).isEqualTo(100);
        assertThat(timer.mean(TimeUnit.MILLISECONDS)).isLessThan(50);  // 平均响应时间 < 50ms
        assertThat(timer.max(TimeUnit.MILLISECONDS)).isLessThan(200);  // 最大响应时间 < 200ms
    }
    
    @Test
    void createUser_ConcurrentLoad_HandlesHighThroughput() throws InterruptedException {
        // 并发性能测试
        int threadCount = 10;
        int requestsPerThread = 50;
        CountDownLatch latch = new CountDownLatch(threadCount);
        AtomicInteger successCount = new AtomicInteger(0);
        AtomicInteger errorCount = new AtomicInteger(0);
        
        ExecutorService executor = Executors.newFixedThreadPool(threadCount);
        
        long startTime = System.currentTimeMillis();
        
        for (int t = 0; t < threadCount; t++) {
            final int threadId = t;
            executor.submit(() -> {
                try {
                    for (int i = 0; i < requestsPerThread; i++) {
                        try {
                            User user = new User(
                                "user" + threadId + "_" + i + "@example.com",
                                "User " + threadId + "_" + i,
                                25
                            );
                            userService.createUser(user);
                            successCount.incrementAndGet();
                        } catch (Exception e) {
                            errorCount.incrementAndGet();
                        }
                    }
                } finally {
                    latch.countDown();
                }
            });
        }
        
        latch.await(30, TimeUnit.SECONDS);
        executor.shutdown();
        
        long endTime = System.currentTimeMillis();
        long totalTime = endTime - startTime;
        
        int totalRequests = threadCount * requestsPerThread;
        double throughput = (double) successCount.get() / (totalTime / 1000.0);
        double errorRate = (double) errorCount.get() / totalRequests;
        
        // 验证性能指标
        assertThat(throughput).isGreaterThan(100);  // 吞吐量 > 100 TPS
        assertThat(errorRate).isLessThan(0.01);     // 错误率 < 1%
        assertThat(successCount.get()).isEqualTo(totalRequests - errorCount.get());
        
        // 验证数据库中的记录数
        long userCount = userRepository.count();
        assertThat(userCount).isEqualTo(successCount.get());
    }
    
    @Test
    void getUserById_CachePerformance_ImprovedResponseTime() {
        // 缓存性能测试
        User user = new User("john@example.com", "John Doe", 25);
        User savedUser = userService.createUser(user);
        
        // 第一次查询（缓存未命中）
        long start1 = System.nanoTime();
        User result1 = userService.getUserById(savedUser.getId());
        long time1 = System.nanoTime() - start1;
        
        // 第二次查询（缓存命中）
        long start2 = System.nanoTime();
        User result2 = userService.getUserById(savedUser.getId());
        long time2 = System.nanoTime() - start2;
        
        // 验证缓存效果
        assertThat(result1).isEqualTo(result2);
        assertThat(time2).isLessThan(time1 / 2);  // 缓存命中时间应该显著减少
        
        // 多次查询验证缓存稳定性
        List<Long> cacheTimes = new ArrayList<>();
        for (int i = 0; i < 100; i++) {
            long start = System.nanoTime();
            userService.getUserById(savedUser.getId());
            cacheTimes.add(System.nanoTime() - start);
        }
        
        double avgCacheTime = cacheTimes.stream()
                .mapToLong(Long::longValue)
                .average()
                .orElse(0);
        
        assertThat(avgCacheTime).isLessThan(TimeUnit.MILLISECONDS.toNanos(5));  // 平均缓存响应时间 < 5ms
    }
    
    @Test
    void memoryUsage_BulkOperations_WithinLimits() {
        // 内存使用测试
        Runtime runtime = Runtime.getRuntime();
        
        // 记录初始内存使用
        runtime.gc();
        long initialMemory = runtime.totalMemory() - runtime.freeMemory();
        
        // 执行大批量操作
        List<User> users = new ArrayList<>();
        for (int i = 0; i < 10000; i++) {
            User user = new User(
                "bulk" + i + "@example.com",
                "Bulk User " + i,
                25
            );
            users.add(user);
        }
        
        userService.createUsersInBatch(users);
        
        // 记录操作后内存使用
        runtime.gc();
        long finalMemory = runtime.totalMemory() - runtime.freeMemory();
        
        long memoryIncrease = finalMemory - initialMemory;
        long memoryIncreaseInMB = memoryIncrease / (1024 * 1024);
        
        // 验证内存使用在合理范围内
        assertThat(memoryIncreaseInMB).isLessThan(100);  // 内存增长 < 100MB
        
        // 清理数据后验证内存释放
        userRepository.deleteAll();
        runtime.gc();
        
        long cleanupMemory = runtime.totalMemory() - runtime.freeMemory();
        long memoryAfterCleanup = cleanupMemory - initialMemory;
        
        assertThat(memoryAfterCleanup).isLessThan(memoryIncrease / 2);  // 清理后内存显著减少
    }
    
    @DynamicPropertySource
    static void configureProperties(DynamicPropertyRegistry registry) {
        registry.add("spring.datasource.url", mysql::getJdbcUrl);
        registry.add("spring.datasource.username", mysql::getUsername);
        registry.add("spring.datasource.password", mysql::getPassword);
        registry.add("spring.jpa.hibernate.ddl-auto", () -> "create-drop");
    }
}
```

## 11. 部署和运维检查

### 11.1 容器化 (Major)
- [ ] Dockerfile最佳实践
- [ ] 镜像安全扫描
- [ ] 多阶段构建
- [ ] 镜像大小优化

### 11.2 优雅关闭 (Critical)
- [ ] 应用停机处理
- [ ] 资源清理
- [ ] 正在处理的请求完成
- [ ] 健康检查状态更新

## 12. 依赖管理检查

### 12.1 依赖安全 (Critical)
- [ ] 第三方库安全漏洞扫描
- [ ] 依赖版本管理
- [ ] 许可证合规检查
- [ ] 依赖冲突解决

### 12.2 依赖优化 (Minor)
- [ ] 移除未使用的依赖
- [ ] 依赖版本统一管理
- [ ] 传递依赖控制
- [ ] 依赖文档维护

## 检查清单使用说明

1. **分阶段检查**: 建议按照Critical -> Major -> Minor的顺序进行检查
2. **工具辅助**: 结合SonarQube、Checkstyle、SpotBugs等工具进行自动化检查
3. **团队协作**: 建立代码评审流程，确保每个PR都经过检查
4. **持续改进**: 根据项目特点和团队经验不断完善检查清单
5. **文档记录**: 对发现的问题和解决方案进行记录，形成知识库

## 常用检查工具

- **静态代码分析**: SonarQube, Checkstyle, SpotBugs, PMD
- **依赖安全扫描**: OWASP Dependency Check, Snyk
- **性能分析**: JProfiler, VisualVM, Arthas
- **监控工具**: Prometheus, Grafana, ELK Stack
- **测试工具**: JUnit, Mockito, TestContainers

---

**文档版本**: v1.0  
**最后更新**: 2024年  
**适用范围**: Java微服务项目代码评审  
**维护团队**: 代码质量保障团队