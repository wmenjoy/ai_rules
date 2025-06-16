# Java代码审查规范 - 第四章质量控制标准（完整版）

## 第四章 质量控制标准

**第十二条** 代码审查应按照以下技术标准执行，确保代码质量符合要求：

### 4.1 需求设计与检查

#### 4.1.1 需求变更影响 🔴:

##### 4.1.1.1 变更应该遵守职责单一原则，避免大范围的修改

**1. 检测目标**

a. 每次提交只做一件事。
b. 单个类/方法是否承担多个职责。
c. 变更是否只影响一小块功能，影响面可控。

**2. 检测方法**

1. SonarQube（检测方法过长、类职责过多）。
2. 人工检测代码的变动情况
3. AI 自动识别

**3. 错误示例**

```java
// ❌ 职责不单一
public void processUserRequest(Request req) {
    validateRequest(req);          // 验证
    updateUserInDatabase(req);     // 数据库操作
    sendNotification(req);         // 发送通知
}

// ❌ 大范围修改，影响面过大
public class UserService {
    // 同时修改多个不相关的方法
    public void createUser() { /* 修改1 */ }
    public void updateUser() { /* 修改2 */ }
    public void deleteUser() { /* 修改3 */ }
    public void sendEmail() { /* 修改4 */ }
    public void generateReport() { /* 修改5 */ }
}
```

**4. 正确示例**

```java
// ✅ 正确：拆分成多个职责清晰的方法
public void processUserRequest(Request req) {
    requestValidator.validate(req);
    userUpdater.update(req);
    notifier.send(req);
}

// ✅ 正确：单一职责，影响面可控
public class UserService {
    public void createUser(User user) {
        // 只负责用户创建逻辑
        validateUser(user);
        userRepository.save(user);
    }
}

public class NotificationService {
    public void sendWelcomeEmail(User user) {
        // 只负责通知逻辑
        emailSender.send(user.getEmail(), welcomeTemplate);
    }
}
```

##### 4.1.1.2 变更修改应该向后兼容，不破坏现有功能

**1. 检测目标**

a. 是否删除或重构了对外接口。
b. 是否移除或更改已有字段、类、接口。
c. 是否更改了默认行为。

**2. 检测方法**

1. SonarQube（检测API兼容性）。
2. 运行回归测试。
3. API版本对比工具检测。

**3. 错误示例**

```java
// ❌ 错误：破坏性接口变更
@RestController
public class UserController {
    // 直接删除原有接口，破坏向后兼容性
    // @GetMapping("/users/{id}")
    // public User getUser(@PathVariable Long id) { ... }
    
    @GetMapping("/users/{userId}")
    public User getUserInfo(@PathVariable Long userId) {
        // 参数名变更，破坏兼容性
    }
}

// ❌ 错误：移除已有字段
public class UserResponse {
    private String name;
    // private String email; // 删除字段，破坏兼容性
    private String phone;
}

// ❌ 错误：更改方法签名
public class UserService {
    // 原方法：public User findUser(Long id)
    public Optional<User> findUser(Long id, boolean includeDeleted) {
        // 返回类型和参数都变更，破坏兼容性
    }
}
```

**4. 正确示例**

```java
// ✅ 正确：保持向后兼容
@RestController
public class UserController {
    // 保留原有接口
    @GetMapping("/users/{id}")
    public User getUser(@PathVariable Long id) {
        return getUserInfo(id);
    }
    
    // 新增接口，不影响原有功能
    @GetMapping("/v2/users/{userId}")
    public UserDetailResponse getUserInfo(@PathVariable Long userId) {
        // 新版本接口
    }
}

// ✅ 正确：使用@Deprecated标记过时方法
public class UserService {
    @Deprecated
    public User findUser(Long id) {
        return findUser(id, false);
    }
    
    public User findUser(Long id, boolean includeDeleted) {
        // 新方法实现
    }
}

// ✅ 正确：字段只增加，不删除
public class UserResponse {
    private String name;
    private String email;
    private String phone;
    private String address; // 新增字段，不影响兼容性
}
```

##### 4.1.1.3 变更应该有完善的回滚和灰度机制

**1. 检测目标**

a. 是否有功能开关控制新功能。
b. 是否有数据库变更的回滚脚本。
c. 是否支持灰度发布和快速回滚。

**2. 检测方法**

1. 检查配置文件中的功能开关。
2. 检查数据库迁移脚本的回滚版本。
3. 检查部署脚本的回滚机制。

**3. 错误示例**

```java
// ❌ 错误：没有功能开关的新功能
@Service
public class PaymentService {
    public void processPayment(Payment payment) {
        // 直接使用新的支付逻辑，无法回滚
        newPaymentProcessor.process(payment);
    }
}

// ❌ 错误：不可逆的数据库变更
-- 删除列，无法回滚
ALTER TABLE users DROP COLUMN old_field;
```

**4. 正确示例**

```java
// ✅ 正确：使用功能开关
@Service
public class PaymentService {
    @Value("${feature.new-payment-processor.enabled:false}")
    private boolean newPaymentEnabled;
    
    public void processPayment(Payment payment) {
        if (newPaymentEnabled) {
            newPaymentProcessor.process(payment);
        } else {
            oldPaymentProcessor.process(payment);
        }
    }
}

// ✅ 正确：可回滚的数据库变更
-- 迁移脚本
ALTER TABLE users ADD COLUMN new_field VARCHAR(255);

-- 回滚脚本
ALTER TABLE users DROP COLUMN new_field;
```

#### 4.1.2 设计与需求匹配检查 🔴:

##### 4.1.2.1 设计方案应该是需求匹配、成本、可行性的最佳平衡

**1. 检测目标**

a. 设计是否满足功能需求。
b. 设计是否满足非功能需求（性能、安全、可用性）。
c. 实现成本是否合理。

**2. 检测方法**

1. 需求追溯矩阵检查。
2. 架构评审和技术方案评审。
3. 成本效益分析。

**3. 错误示例**

```java
// ❌ 错误：过度设计，成本过高
public class SimpleCalculator {
    // 为简单计算器引入复杂的设计模式
    private CalculatorStrategyFactory strategyFactory;
    private CalculatorCommandInvoker commandInvoker;
    private CalculatorStateManager stateManager;
    
    public double add(double a, double b) {
        CalculatorStrategy strategy = strategyFactory.createAddStrategy();
        CalculatorCommand command = new AddCommand(strategy, a, b);
        return commandInvoker.execute(command);
    }
}

// ❌ 错误：设计不满足性能需求
@Service
public class UserService {
    public List<User> getAllUsers() {
        // 一次性加载所有用户，不支持分页
        return userRepository.findAll(); // 可能返回百万级数据
    }
}
```

**4. 正确示例**

```java
// ✅ 正确：简单需求用简单设计
public class SimpleCalculator {
    public double add(double a, double b) {
        return a + b;
    }
    
    public double subtract(double a, double b) {
        return a - b;
    }
}

// ✅ 正确：满足性能需求的分页设计
@Service
public class UserService {
    public Page<User> getUsers(Pageable pageable) {
        return userRepository.findAll(pageable);
    }
    
    public List<User> getUsersByIds(List<Long> ids) {
        if (ids.size() > 1000) {
            throw new IllegalArgumentException("批量查询不能超过1000个ID");
        }
        return userRepository.findAllById(ids);
    }
}
```

##### 4.1.2.2 代码变更应该与需求声明一致

**1. 检测目标**

a. 实现的功能是否与需求文档一致。
b. 是否有需求之外的额外实现。
c. 是否遗漏了需求中的功能点。

**2. 检测方法**

1. 需求追溯检查。
2. 功能测试用例验证。
3. 代码审查对比需求文档。

**3. 错误示例**

```java
// ❌ 错误：实现超出需求范围
// 需求：用户注册功能
@Service
public class UserRegistrationService {
    public void registerUser(UserRegistrationRequest request) {
        // 需求范围内
        User user = createUser(request);
        userRepository.save(user);
        
        // ❌ 超出需求：自动创建用户的社交媒体账号
        socialMediaService.createAccounts(user);
        
        // ❌ 超出需求：发送营销邮件
        marketingService.sendWelcomePromotion(user);
    }
}

// ❌ 错误：遗漏需求功能
// 需求：用户登录需要记录登录日志和检查账号状态
@Service
public class LoginService {
    public LoginResult login(String username, String password) {
        User user = userRepository.findByUsername(username);
        if (passwordEncoder.matches(password, user.getPassword())) {
            // ❌ 遗漏：没有记录登录日志
            // ❌ 遗漏：没有检查账号状态（锁定、禁用等）
            return LoginResult.success(user);
        }
        return LoginResult.failure("用户名或密码错误");
    }
}
```

**4. 正确示例**

```java
// ✅ 正确：严格按照需求实现
// 需求：用户注册功能，包含邮箱验证
@Service
public class UserRegistrationService {
    public void registerUser(UserRegistrationRequest request) {
        // 按需求实现用户创建
        User user = createUser(request);
        userRepository.save(user);
        
        // 按需求发送验证邮件
        emailService.sendVerificationEmail(user);
    }
}

// ✅ 正确：完整实现需求功能
// 需求：用户登录需要记录登录日志和检查账号状态
@Service
public class LoginService {
    public LoginResult login(String username, String password) {
        User user = userRepository.findByUsername(username);
        
        // 检查账号状态
        if (!user.isActive()) {
            auditService.logLoginAttempt(username, "ACCOUNT_DISABLED");
            return LoginResult.failure("账号已禁用");
        }
        
        if (passwordEncoder.matches(password, user.getPassword())) {
            // 记录成功登录日志
            auditService.logLoginAttempt(username, "SUCCESS");
            return LoginResult.success(user);
        }
        
        // 记录失败登录日志
        auditService.logLoginAttempt(username, "INVALID_PASSWORD");
        return LoginResult.failure("用户名或密码错误");
    }
}
```

#### 4.1.3 文档完整性检查 🟡:

##### 4.1.3.1 API文档与代码实现保持一致

**1. 检测目标**

a. API文档是否与实际接口一致。
b. 参数说明是否准确完整。
c. 返回值说明是否正确。

**2. 检测方法**

1. 使用Swagger/OpenAPI自动生成文档。
2. API文档与代码的一致性检查工具。
3. 人工审查API文档。

**3. 错误示例**

```java
// ❌ 错误：文档与实现不一致
/**
 * 获取用户信息
 * @param id 用户ID
 * @return 用户信息
 */
@GetMapping("/users/{id}")
public UserResponse getUser(@PathVariable Long id, 
                           @RequestParam(required = false) Boolean includeDeleted) {
    // 实际有includeDeleted参数，但文档中没有说明
    return userService.findUser(id, includeDeleted);
}

// ❌ 错误：返回值文档不准确
/**
 * 创建用户
 * @return 创建成功的用户信息
 */
@PostMapping("/users")
public ResponseEntity<ApiResponse<User>> createUser(@RequestBody CreateUserRequest request) {
    // 实际返回包装的ApiResponse，但文档说明不准确
    User user = userService.createUser(request);
    return ResponseEntity.ok(ApiResponse.success(user));
}
```

**4. 正确示例**

```java
// ✅ 正确：完整准确的API文档
/**
 * 获取用户信息
 * @param id 用户ID
 * @param includeDeleted 是否包含已删除用户，默认false
 * @return 用户信息
 */
@ApiOperation(value = "获取用户信息", notes = "根据用户ID获取用户详细信息")
@ApiResponses({
    @ApiResponse(code = 200, message = "获取成功"),
    @ApiResponse(code = 404, message = "用户不存在")
})
@GetMapping("/users/{id}")
public UserResponse getUser(
    @ApiParam(value = "用户ID", required = true) @PathVariable Long id,
    @ApiParam(value = "是否包含已删除用户", defaultValue = "false") 
    @RequestParam(required = false, defaultValue = "false") Boolean includeDeleted) {
    return userService.findUser(id, includeDeleted);
}
```

#### 4.1.4 逻辑完整性检查 🔴:

##### 4.1.4.1 业务逻辑分支覆盖完整

**1. 检测目标**

a. 正常流程是否完整实现。
b. 异常情况是否有处理机制。
c. 边界条件是否考虑周全。

**2. 检测方法**

1. 代码覆盖率测试。
2. 业务场景测试用例检查。
3. 异常路径测试。

**3. 错误示例**

```java
// ❌ 错误：缺少异常处理分支
@Service
public class OrderService {
    public void processOrder(Order order) {
        // ❌ 没有检查订单状态
        // ❌ 没有检查库存
        // ❌ 没有处理支付失败情况
        
        paymentService.charge(order.getAmount());
        inventoryService.reduceStock(order.getItems());
        order.setStatus(OrderStatus.COMPLETED);
        orderRepository.save(order);
    }
}

// ❌ 错误：边界条件未处理
public class DiscountCalculator {
    public BigDecimal calculateDiscount(BigDecimal amount, BigDecimal discountRate) {
        // ❌ 没有检查参数为null
        // ❌ 没有检查折扣率范围
        return amount.multiply(discountRate);
    }
}
```

**4. 正确示例**

```java
// ✅ 正确：完整的业务逻辑分支
@Service
public class OrderService {
    public void processOrder(Order order) {
        // 检查订单状态
        if (order.getStatus() != OrderStatus.PENDING) {
            throw new IllegalStateException("订单状态不允许处理");
        }
        
        // 检查库存
        if (!inventoryService.checkStock(order.getItems())) {
            throw new InsufficientStockException("库存不足");
        }
        
        try {
            // 处理支付
            PaymentResult result = paymentService.charge(order.getAmount());
            if (!result.isSuccess()) {
                order.setStatus(OrderStatus.PAYMENT_FAILED);
                orderRepository.save(order);
                return;
            }
            
            // 减少库存
            inventoryService.reduceStock(order.getItems());
            order.setStatus(OrderStatus.COMPLETED);
            
        } catch (PaymentException e) {
            order.setStatus(OrderStatus.PAYMENT_FAILED);
            log.error("订单支付失败: {}", order.getId(), e);
        } catch (Exception e) {
            order.setStatus(OrderStatus.FAILED);
            log.error("订单处理失败: {}", order.getId(), e);
        } finally {
            orderRepository.save(order);
        }
    }
}

// ✅ 正确：完整的边界条件检查
public class DiscountCalculator {
    public BigDecimal calculateDiscount(BigDecimal amount, BigDecimal discountRate) {
        if (amount == null || discountRate == null) {
            throw new IllegalArgumentException("金额和折扣率不能为空");
        }
        
        if (amount.compareTo(BigDecimal.ZERO) < 0) {
            throw new IllegalArgumentException("金额不能为负数");
        }
        
        if (discountRate.compareTo(BigDecimal.ZERO) < 0 || 
            discountRate.compareTo(BigDecimal.ONE) > 0) {
            throw new IllegalArgumentException("折扣率必须在0-1之间");
        }
        
        return amount.multiply(discountRate);
    }
}
```

---

## 总结

本章节详细规定了Java代码审查中需求设计与检查的具体标准，包括：

1. **需求变更影响检查**：确保变更遵循单一职责原则、向后兼容性和完善的回滚机制
2. **设计与需求匹配检查**：确保设计方案合理且与需求一致
3. **文档完整性检查**：确保API文档与代码实现保持一致
4. **逻辑完整性检查**：确保业务逻辑分支覆盖完整

每个检查项都包含了明确的检测目标、检测方法、错误示例和正确示例，为代码审查人员提供了具体的操作指导。

通过严格执行这些标准，可以有效提高Java项目的代码质量，降低系统风险，确保项目的稳定性和可维护性。