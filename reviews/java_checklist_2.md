# Java微服务代码评审检查清单

## 前言

本检查清单基于Java微服务架构特点，从代码规范、安全性、性能、可观测性等多个维度提供系统化的代码评审指导。每个检查点按严重程度分为三个等级：
- **Critical**: 必须修复，影响系统安全或稳定性
- **Major**: 建议修复，影响代码质量或性能
- **Minor**: 可选修复，代码规范或可读性问题

## 目录

- [Java微服务代码评审检查清单](#java微服务代码评审检查清单)
  - [前言](#前言)
  - [目录](#目录)
  - [1. 基础代码规范检查](#1-基础代码规范检查)
    - [1.1 命名规范 (Minor)](#11-命名规范-minor)
      - [1.1.1 类名使用大驼峰命名法](#111-类名使用大驼峰命名法)
      - [1.1.2 方法名和变量名使用小驼峰命名法](#112-方法名和变量名使用小驼峰命名法)
      - [1.1.3 常量使用全大写下划线分隔](#113-常量使用全大写下划线分隔)
      - [1.1.4 包名全小写，避免使用关键字](#114-包名全小写避免使用关键字)
    - [1.2 代码结构 (Major)](#12-代码结构-major)
      - [1.2.1 单一职责原则，类和方法职责明确](#121-单一职责原则类和方法职责明确)
      - [1.2.2 避免过长的方法（建议不超过50行）](#122-避免过长的方法建议不超过50行)
      - [1.2.3 避免过深的嵌套（建议不超过3层）](#123-避免过深的嵌套建议不超过3层)
      - [1.2.4 合理使用设计模式](#124-合理使用设计模式)
  - [2. 微服务特有检查](#2-微服务特有检查)
    - [2.1 线程池配置 (Critical)](#21-线程池配置-critical)
      - [2.1.1 避免使用 `Executors.newCachedThreadPool()`](#211-避免使用-executorsnewcachedthreadpool)
      - [2.1.2 自定义线程池参数（核心线程数、最大线程数、队列大小）](#212-自定义线程池参数核心线程数最大线程数队列大小)
      - [2.1.3 设置有意义的线程名称](#213-设置有意义的线程名称)
      - [2.1.4 配置合适的拒绝策略](#214-配置合适的拒绝策略)
    - [2.2 超时设置 (Critical)](#22-超时设置-critical)
      - [2.2.1 HTTP客户端设置连接超时和读取超时](#221-http客户端设置连接超时和读取超时)
      - [2.2.2 数据库连接池超时配置](#222-数据库连接池超时配置)
      - [2.2.3 缓存操作超时设置](#223-缓存操作超时设置)
      - [2.2.4 消息队列消费超时配置](#224-消息队列消费超时配置)
  - [3. 并发和线程安全检查](#3-并发和线程安全检查)
    - [3.1 共享变量线程安全 (Critical)](#31-共享变量线程安全-critical)
      - [3.1.1 Service类成员变量必须是无状态或线程安全的](#311-service类成员变量必须是无状态或线程安全的)
      - [3.1.2 Controller类避免使用可变成员变量](#312-controller类避免使用可变成员变量)
      - [3.1.3 静态变量的线程安全性检查](#313-静态变量的线程安全性检查)
      - [3.1.4 集合类的线程安全使用](#314-集合类的线程安全使用)
    - [3.2 锁的使用 (Major)](#32-锁的使用-major)
      - [3.2.1 避免在循环中获取锁](#321-避免在循环中获取锁)
      - [3.2.2 锁的粒度要适当](#322-锁的粒度要适当)
      - [3.2.3 避免死锁的发生](#323-避免死锁的发生)
      - [3.2.4 使用try-finally确保锁释放](#324-使用try-finally确保锁释放)
  - [4. 安全性检查](#4-安全性检查)
    - [4.1 输入验证 (Critical)](#41-输入验证-critical)
      - [4.1.1 所有外部输入必须进行验证](#411-所有外部输入必须进行验证)
      - [4.1.2 使用@Valid注解进行参数校验](#412-使用valid注解进行参数校验)
      - [4.1.3 防止SQL注入攻击](#413-防止sql注入攻击)
      - [4.1.4 防止XSS攻击](#414-防止xss攻击)
      - [4.1.5 文件上传安全检查](#415-文件上传安全检查)
    - [4.2 敏感信息处理 (Critical)](#42-敏感信息处理-critical)
      - [4.2.1 密码、API密钥不能硬编码](#421-密码api密钥不能硬编码)
      - [4.2.2 日志中不包含敏感信息](#422-日志中不包含敏感信息)
      - [4.2.3 数据库密码加密存储](#423-数据库密码加密存储)
      - [4.2.4 错误信息不泄露系统内部结构](#424-错误信息不泄露系统内部结构)
    - [4.3 认证授权 (Critical)](#43-认证授权-critical)
      - [4.3.1 JWT token验证机制](#431-jwt-token验证机制)
      - [4.3.2 接口权限控制](#432-接口权限控制)
      - [4.3.3 HTTPS配置](#433-https配置)
      - [4.3.4 CORS配置安全](#434-cors配置安全)
  - [5. 性能优化检查](#5-性能优化检查)
    - [5.1 数据库操作 (Major)](#51-数据库操作-major)
      - [5.1.1 避免N+1查询问题](#511-避免n1查询问题)
      - [5.1.2 合理使用索引](#512-合理使用索引)
      - [5.1.3 批量操作优化](#513-批量操作优化)
      - [5.1.4 分页查询实现](#514-分页查询实现)
      - [5.1.5 连接池配置检查](#515-连接池配置检查)
    - [5.2 缓存使用 (Major)](#52-缓存使用-major)
      - [5.2.1 缓存策略合理性](#521-缓存策略合理性)
      - [5.2.2 缓存穿透、击穿、雪崩防护](#522-缓存穿透击穿雪崩防护)
      - [5.2.3 缓存过期时间设置](#523-缓存过期时间设置)
      - [5.2.4 本地缓存线程安全性](#524-本地缓存线程安全性)
    - [5.3 内存管理 (Major)](#53-内存管理-major)
      - [5.3.1 避免内存泄漏](#531-避免内存泄漏)
      - [5.3.2 大对象处理优化](#532-大对象处理优化)
      - [5.3.3 集合使用优化](#533-集合使用优化)
      - [5.3.4 字符串拼接优化](#534-字符串拼接优化)
    - [5.4 批处理操作 (Major)](#54-批处理操作-major)
      - [5.4.1 数据库批处理](#541-数据库批处理)
      - [5.4.2 集合批处理](#542-集合批处理)
    - [5.5 异步处理 (Major)](#55-异步处理-major)
      - [5.5.1 异步方法使用](#551-异步方法使用)
    - [5.6 内存使用优化 (Major)](#56-内存使用优化-major)
      - [5.6.1 对象创建优化](#561-对象创建优化)
    - [5.7 IO操作优化 (Major)](#57-io操作优化-major)
      - [5.7.1 文件IO优化](#571-文件io优化)
    - [5.8 算法和数据结构 (Minor)](#58-算法和数据结构-minor)
      - [5.8.1 选择合适的数据结构](#581-选择合适的数据结构)
    - [5.9 JVM参数调优 (Major)](#59-jvm参数调优-major)
      - [5.9.1 堆内存配置](#591-堆内存配置)
    - [5.10 垃圾回收优化 (Major)](#510-垃圾回收优化-major)
      - [5.10.1 GC算法选择](#5101-gc算法选择)
    - [5.11 序列化优化 (Minor)](#511-序列化优化-minor)
      - [5.11.1 选择高效的序列化方式](#5111-选择高效的序列化方式)
    - [5.12 网络传输优化 (Minor)](#512-网络传输优化-minor)
      - [5.12.1 HTTP客户端优化](#5121-http客户端优化)
    - [5.13 代码层面优化 (Minor)](#513-代码层面优化-minor)
      - [5.13.1 避免不必要的计算](#5131-避免不必要的计算)
    - [5.14 资源预加载 (Minor)](#514-资源预加载-minor)
      - [5.14.1 配置和静态资源预加载](#5141-配置和静态资源预加载)
    - [5.15 懒加载策略 (Minor)](#515-懒加载策略-minor)
      - [5.15.1 合理使用懒加载](#5151-合理使用懒加载)
    - [5.16 对象池使用 (Minor)](#516-对象池使用-minor)
      - [5.16.1 合理使用对象池](#5161-合理使用对象池)
  - [6. 可观测性检查](#6-可观测性检查)
    - [6.1 日志规范 (Major)](#61-日志规范-major)
      - [6.1.1 禁止使用System.out.println()](#611-禁止使用systemoutprintln)
      - [6.1.2 使用统一的日志框架](#612-使用统一的日志框架)
      - [6.1.3 日志级别配置正确](#613-日志级别配置正确)
      - [6.1.4 包含TraceId进行链路追踪](#614-包含traceid进行链路追踪)
      - [6.1.5 敏感信息脱敏处理](#615-敏感信息脱敏处理)
    - [6.2 监控集成 (Major)](#62-监控集成-major)
      - [6.2.1 Metrics指标暴露](#621-metrics指标暴露)
      - [6.2.2 健康检查端点实现](#622-健康检查端点实现)
    - [6.3 告警配置 (Major)](#63-告警配置-major)
      - [6.3.1 关键指标告警规则](#631-关键指标告警规则)
  - [7. 容错和稳定性检查](#7-容错和稳定性检查)
    - [7.1 熔断器配置 (Critical)](#71-熔断器配置-critical)
      - [7.1.1 熔断器实现和配置](#711-熔断器实现和配置)
    - [7.2 重试机制 (Major)](#72-重试机制-major)
      - [7.2.1 重试策略配置](#721-重试策略配置)
    - [7.3 限流配置 (Major)](#73-限流配置-major)
      - [7.3.1 接口限流实现](#731-接口限流实现)
    - [7.4 异常处理 (Critical)](#74-异常处理-critical)
      - [7.4.1 全局异常处理器](#741-全局异常处理器)
  - [8. 配置管理检查](#8-配置管理检查)
    - [8.1 配置外部化 (Major)](#81-配置外部化-major)
      - [检查方法](#检查方法)
      - [检查标准](#检查标准)
      - [错误示例](#错误示例)
      - [正确示例](#正确示例)
    - [8.2 环境隔离 (Critical)](#82-环境隔离-critical)
      - [检查方法](#检查方法-1)
      - [检查标准](#检查标准-1)
      - [错误示例](#错误示例-1)
      - [正确示例](#正确示例-1)
  - [9. API设计检查](#9-api设计检查)
    - [9.1 RESTful规范 (Major)](#91-restful规范-major)
      - [检查方法](#检查方法-2)
      - [检查标准](#检查标准-2)
      - [错误示例](#错误示例-2)
      - [正确示例](#正确示例-2)
    - [9.2 版本控制 (Major)](#92-版本控制-major)
      - [检查方法](#检查方法-3)
      - [检查标准](#检查标准-3)
      - [错误示例](#错误示例-3)
      - [正确示例](#正确示例-3)
    - [9.3 参数验证 (Critical)](#93-参数验证-critical)
      - [检查方法](#检查方法-4)
      - [检查标准](#检查标准-4)
      - [错误示例](#错误示例-4)
      - [正确示例](#正确示例-4)
    - [9.4 错误处理 (Critical)](#94-错误处理-critical)
      - [9.4.1 统一错误响应格式](#941-统一错误响应格式)
      - [9.4.2 HTTP状态码使用规范](#942-http状态码使用规范)
    - [9.5 文档规范 (Major)](#95-文档规范-major)
      - [9.5.1 API文档完整性](#951-api文档完整性)
      - [9.5.2 接口变更文档](#952-接口变更文档)
    - [9.6 请求验证 (Major)](#96-请求验证-major)
      - [9.6.1 输入参数验证](#961-输入参数验证)
      - [9.6.2 业务规则验证](#962-业务规则验证)
    - [9.7 响应格式 (Minor)](#97-响应格式-minor)
      - [9.7.1 统一响应结构](#971-统一响应结构)
      - [9.7.2 分页响应格式](#972-分页响应格式)
  - [10. 测试相关检查](#10-测试相关检查)
    - [10.1 单元测试 (Major)](#101-单元测试-major)
      - [10.1.1 测试覆盖率达标（建议80%+）](#1011-测试覆盖率达标建议80)
      - [10.1.2 边界条件测试](#1012-边界条件测试)
      - [10.1.3 异常情况测试](#1013-异常情况测试)
      - [10.1.4 Mock使用正确](#1014-mock使用正确)
    - [10.2 集成测试 (Major)](#102-集成测试-major)
      - [10.2.1 数据库集成测试](#1021-数据库集成测试)
      - [ß10.2.2 外部服务集成测试](#ß1022-外部服务集成测试)
      - [10.2.3 端到端测试](#1023-端到端测试)
      - [10.2.4 性能测试](#1024-性能测试)
    - [10.3 测试覆盖率 (Minor)](#103-测试覆盖率-minor)
      - [10.3.1 代码覆盖率要求](#1031-代码覆盖率要求)
      - [10.3.2 覆盖率报告分析](#1032-覆盖率报告分析)
    - [10.4 测试数据管理 (Minor)](#104-测试数据管理-minor)
      - [10.4.1 测试数据隔离](#1041-测试数据隔离)
      - [10.4.2 测试数据构建器](#1042-测试数据构建器)
    - [10.5 性能测试 (Major)](#105-性能测试-major)
      - [10.5.1 负载测试](#1051-负载测试)
      - [10.5.2 内存和资源监控](#1052-内存和资源监控)
  - [11. 部署和运维检查](#11-部署和运维检查)
    - [11.1 容器化 (Major)](#111-容器化-major)
      - [11.1.1 Dockerfile最佳实践](#1111-dockerfile最佳实践)
      - [11.1.4 健康检查状态更新](#1114-健康检查状态更新)
      - [11.1.2 镜像安全扫描](#1112-镜像安全扫描)
      - [11.1.3 多阶段构建优化](#1113-多阶段构建优化)
    - [11.2 优雅关闭 (Critical)](#112-优雅关闭-critical)
      - [11.2.1 应用停机处理](#1121-应用停机处理)
      - [11.2.2 资源清理](#1122-资源清理)
      - [11.2.3 正在处理的请求完成](#1123-正在处理的请求完成)
  - [12. 依赖管理检查](#12-依赖管理检查)
    - [12.1 依赖安全 (Critical)](#121-依赖安全-critical)
      - [12.1.1 第三方库安全漏洞扫描](#1211-第三方库安全漏洞扫描)
      - [12.1.2 依赖版本管理](#1212-依赖版本管理)
      - [12.1.3 许可证合规检查](#1213-许可证合规检查)
      - [12.1.4 依赖冲突解决](#1214-依赖冲突解决)
    - [12.2 依赖优化 (Minor)](#122-依赖优化-minor)
      - [12.2.1 移除未使用的依赖](#1221-移除未使用的依赖)
      - [12.2.2 依赖版本统一管理](#1222-依赖版本统一管理)
      - [12.2.3 传递依赖控制](#1223-传递依赖控制)
      - [12.2.4 依赖文档维护](#1224-依赖文档维护)
  - [检查清单使用说明](#检查清单使用说明)
  - [常用检查工具](#常用检查工具)

## 1. 基础代码规范检查

### 检查工具
- **静态分析工具**: SonarQube, Checkstyle, SpotBugs, PMD
- **IDE插件**: SonarLint, CheckStyle-IDEA, SpotBugs Plugin
- **AI辅助工具**: GitHub Copilot, CodeGPT, Tabnine (代码规范建议)
- **人工检查**: 代码评审(Code Review)，重点关注命名规范、代码结构
- **自动化检查**: CI/CD流水线集成质量门禁，代码提交前自动检查

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

### 检查工具
- **静态分析工具**: SonarQube (线程池配置规则), SpotBugs (并发问题检测)
- **专业工具**: Arthas (线程池监控), JProfiler (线程分析), VisualVM
- **AI辅助工具**: DeepCode, Amazon CodeGuru (性能和并发问题检测)
- **人工检查**: 架构评审，重点关注线程池配置、超时设置、资源管理
- **监控工具**: Micrometer + Prometheus (线程池指标监控)
- **测试工具**: JMeter (并发压力测试), Gatling (性能测试)

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

### 5.4 批处理操作 (Major)

#### 5.4.1 数据库批处理
- **检查方法**: 检查数据库操作是否使用批处理，查看SQL执行日志
- **检查标准**: 批量操作必须使用批处理，避免逐条执行
- **不正确实例**:
```java
// 错误示例 - 逐条插入
@Service
public class UserService {
    @Autowired
    private UserRepository userRepository;
    
    public void saveUsers(List<User> users) {
        for (User user : users) {
            userRepository.save(user);  // 错误：逐条保存
        }
    }
    
    // 错误：逐条执行SQL
    public void updateUserStatus(List<Long> userIds, String status) {
        for (Long userId : userIds) {
            jdbcTemplate.update(
                "UPDATE users SET status = ? WHERE id = ?", 
                status, userId
            );
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
    public void saveUsers(List<User> users) {
        // 正确：批量保存
        userRepository.saveAll(users);
    }
    
    @Transactional
    public void updateUserStatus(List<Long> userIds, String status) {
        // 正确：批处理更新
        String sql = "UPDATE users SET status = ? WHERE id = ?";
        List<Object[]> batchArgs = userIds.stream()
            .map(id -> new Object[]{status, id})
            .collect(Collectors.toList());
        
        jdbcTemplate.batchUpdate(sql, batchArgs);
    }
    
    // 使用MyBatis批处理
    public void batchInsertWithMyBatis(List<User> users) {
        try (SqlSession sqlSession = sqlSessionFactory.openSession(ExecutorType.BATCH)) {
            UserMapper mapper = sqlSession.getMapper(UserMapper.class);
            for (User user : users) {
                mapper.insert(user);
            }
            sqlSession.commit();
        }
    }
}
```

#### 5.4.2 集合批处理
- **检查方法**: 检查集合操作是否合理使用批处理API
- **检查标准**: 大量数据处理时使用Stream API和批处理方法
- **不正确实例**:
```java
// 错误示例 - 低效的集合操作
public class DataProcessor {
    public List<String> processData(List<String> data) {
        List<String> result = new ArrayList<>();
        for (String item : data) {
            if (item.length() > 5) {
                result.add(item.toUpperCase());  // 错误：逐个处理
            }
        }
        return result;
    }
    
    // 错误：频繁的字符串拼接
    public String concatenateStrings(List<String> strings) {
        String result = "";
        for (String str : strings) {
            result += str + ",";  // 错误：每次都创建新字符串
        }
        return result;
    }
}

// 正确示例
public class DataProcessor {
    public List<String> processData(List<String> data) {
        // 正确：使用Stream API批处理
        return data.stream()
            .filter(item -> item.length() > 5)
            .map(String::toUpperCase)
            .collect(Collectors.toList());
    }
    
    // 正确：使用StringBuilder
    public String concatenateStrings(List<String> strings) {
        return String.join(",", strings);
    }
    
    // 正确：并行处理大数据集
    public List<ProcessedData> processLargeDataset(List<RawData> rawData) {
        return rawData.parallelStream()
            .map(this::processItem)
            .collect(Collectors.toList());
    }
    
    // 正确：分批处理
    public void processBatches(List<Data> allData, int batchSize) {
        for (int i = 0; i < allData.size(); i += batchSize) {
            int end = Math.min(i + batchSize, allData.size());
            List<Data> batch = allData.subList(i, end);
            processBatch(batch);
        }
    }
}
```

### 5.5 异步处理 (Major)

#### 5.5.1 异步方法使用
- **检查方法**: 检查是否正确使用@Async注解和CompletableFuture
- **检查标准**: 耗时操作必须异步执行，避免阻塞主线程
- **不正确实例**:
```java
// 错误示例 - 同步执行耗时操作
@RestController
public class OrderController {
    @Autowired
    private EmailService emailService;
    
    @Autowired
    private InventoryService inventoryService;
    
    @PostMapping("/orders")
    public ResponseEntity<Order> createOrder(@RequestBody Order order) {
        // 错误：同步执行耗时操作
        Order savedOrder = orderService.save(order);
        emailService.sendConfirmationEmail(order);  // 阻塞
        inventoryService.updateInventory(order);    // 阻塞
        
        return ResponseEntity.ok(savedOrder);
    }
}

// 错误示例 - 不正确的异步配置
@Service
public class NotificationService {
    // 错误：没有配置线程池
    @Async
    public void sendNotification(String message) {
        // 耗时操作
    }
    
    // 错误：在同一个类中调用异步方法
    public void processOrder(Order order) {
        this.sendNotification("Order processed");  // 不会异步执行
    }
}

// 正确示例
@RestController
public class OrderController {
    @Autowired
    private EmailService emailService;
    
    @Autowired
    private InventoryService inventoryService;
    
    @PostMapping("/orders")
    public ResponseEntity<Order> createOrder(@RequestBody Order order) {
        // 正确：保存订单（同步）
        Order savedOrder = orderService.save(order);
        
        // 正确：异步执行非关键操作
        CompletableFuture.allOf(
            emailService.sendConfirmationEmailAsync(order),
            inventoryService.updateInventoryAsync(order)
        ).exceptionally(throwable -> {
            log.error("Async operations failed", throwable);
            return null;
        });
        
        return ResponseEntity.ok(savedOrder);
    }
}

@Service
public class EmailService {
    @Async("taskExecutor")
    public CompletableFuture<Void> sendConfirmationEmailAsync(Order order) {
        try {
            // 发送邮件逻辑
            sendEmail(order.getCustomerEmail(), "Order Confirmation", buildEmailContent(order));
            return CompletableFuture.completedFuture(null);
        } catch (Exception e) {
            return CompletableFuture.failedFuture(e);
        }
    }
}

// 正确：异步配置
@Configuration
@EnableAsync
public class AsyncConfig {
    @Bean("taskExecutor")
    public TaskExecutor taskExecutor() {
        ThreadPoolTaskExecutor executor = new ThreadPoolTaskExecutor();
        executor.setCorePoolSize(5);
        executor.setMaxPoolSize(10);
        executor.setQueueCapacity(100);
        executor.setThreadNamePrefix("async-");
        executor.setRejectedExecutionHandler(new ThreadPoolExecutor.CallerRunsPolicy());
        executor.initialize();
        return executor;
    }
}
```

### 5.6 内存使用优化 (Major)

#### 5.6.1 对象创建优化
- **检查方法**: 使用内存分析工具检查对象创建频率和大小
- **检查标准**: 避免不必要的对象创建，重用对象
- **不正确实例**:
```java
// 错误示例 - 频繁创建对象
public class StringProcessor {
    public String processStrings(List<String> strings) {
        String result = "";
        for (String str : strings) {
            result = result + str + ",";  // 错误：每次创建新String对象
        }
        return result;
    }
    
    // 错误：在循环中创建对象
    public List<Date> generateDates(int count) {
        List<Date> dates = new ArrayList<>();
        for (int i = 0; i < count; i++) {
            SimpleDateFormat sdf = new SimpleDateFormat("yyyy-MM-dd");  // 错误：重复创建
            dates.add(sdf.parse("2024-01-" + (i + 1)));
        }
        return dates;
    }
}

// 正确示例
public class StringProcessor {
    // 正确：使用StringBuilder
    public String processStrings(List<String> strings) {
        StringBuilder sb = new StringBuilder();
        for (String str : strings) {
            sb.append(str).append(",");
        }
        return sb.toString();
    }
    
    // 正确：重用对象
    private static final ThreadLocal<SimpleDateFormat> DATE_FORMAT = 
        ThreadLocal.withInitial(() -> new SimpleDateFormat("yyyy-MM-dd"));
    
    public List<Date> generateDates(int count) {
        List<Date> dates = new ArrayList<>(count);  // 预分配容量
        SimpleDateFormat sdf = DATE_FORMAT.get();
        
        for (int i = 0; i < count; i++) {
            try {
                dates.add(sdf.parse("2024-01-" + (i + 1)));
            } catch (ParseException e) {
                throw new RuntimeException(e);
            }
        }
        return dates;
    }
    
    // 正确：使用对象池
    private final ObjectPool<StringBuilder> stringBuilderPool = 
        new GenericObjectPool<>(new StringBuilderFactory());
    
    public String buildString(List<String> parts) {
        StringBuilder sb = null;
        try {
            sb = stringBuilderPool.borrowObject();
            for (String part : parts) {
                sb.append(part);
            }
            return sb.toString();
        } finally {
            if (sb != null) {
                sb.setLength(0);  // 清空内容
                stringBuilderPool.returnObject(sb);
            }
        }
    }
}
```

### 5.7 IO操作优化 (Major)

#### 5.7.1 文件IO优化
- **检查方法**: 检查文件读写操作是否使用缓冲和NIO
- **检查标准**: 大文件操作使用NIO，小文件使用缓冲IO
- **不正确实例**:
```java
// 错误示例 - 低效的文件操作
public class FileProcessor {
    // 错误：逐字节读取
    public String readFile(String filePath) throws IOException {
        FileInputStream fis = new FileInputStream(filePath);
        StringBuilder content = new StringBuilder();
        int b;
        while ((b = fis.read()) != -1) {  // 错误：逐字节读取
            content.append((char) b);
        }
        fis.close();
        return content.toString();
    }
    
    // 错误：没有使用缓冲
    public void writeFile(String filePath, String content) throws IOException {
        FileOutputStream fos = new FileOutputStream(filePath);
        fos.write(content.getBytes());  // 错误：一次性写入大量数据
        fos.close();
    }
}

// 正确示例
public class FileProcessor {
    // 正确：使用缓冲读取
    public String readFile(String filePath) throws IOException {
        try (BufferedReader reader = Files.newBufferedReader(Paths.get(filePath))) {
            return reader.lines().collect(Collectors.joining("\n"));
        }
    }
    
    // 正确：使用NIO处理大文件
    public void processLargeFile(String inputPath, String outputPath) throws IOException {
        try (FileChannel inputChannel = FileChannel.open(Paths.get(inputPath), StandardOpenOption.READ);
             FileChannel outputChannel = FileChannel.open(Paths.get(outputPath), 
                 StandardOpenOption.CREATE, StandardOpenOption.WRITE)) {
            
            long size = inputChannel.size();
            long position = 0;
            
            while (position < size) {
                long transferred = inputChannel.transferTo(position, size - position, outputChannel);
                position += transferred;
            }
        }
    }
    
    // 正确：使用内存映射处理超大文件
    public void processHugeFile(String filePath) throws IOException {
        try (RandomAccessFile file = new RandomAccessFile(filePath, "r");
             FileChannel channel = file.getChannel()) {
            
            long fileSize = channel.size();
            long mapSize = Math.min(fileSize, 1024 * 1024 * 100); // 100MB chunks
            
            for (long position = 0; position < fileSize; position += mapSize) {
                long size = Math.min(mapSize, fileSize - position);
                MappedByteBuffer buffer = channel.map(FileChannel.MapMode.READ_ONLY, position, size);
                processBuffer(buffer);
            }
        }
    }
}
```

### 5.8 算法和数据结构 (Minor)

#### 5.8.1 选择合适的数据结构
- **检查方法**: 检查代码中数据结构的选择是否合理
- **检查标准**: 根据使用场景选择最优的数据结构
- **不正确实例**:
```java
// 错误示例 - 数据结构选择不当
public class DataManager {
    // 错误：频繁查找使用List
    private List<User> users = new ArrayList<>();
    
    public User findUserById(String id) {
        for (User user : users) {  // O(n)时间复杂度
            if (user.getId().equals(id)) {
                return user;
            }
        }
        return null;
    }
    
    // 错误：需要去重但使用List
    public List<String> getUniqueItems(List<String> items) {
        List<String> unique = new ArrayList<>();
        for (String item : items) {
            if (!unique.contains(item)) {  // O(n)查找
                unique.add(item);
            }
        }
        return unique;
    }
    
    // 错误：频繁插入删除使用ArrayList
    public void processQueue(List<Task> tasks) {
        while (!tasks.isEmpty()) {
            Task task = tasks.remove(0);  // O(n)操作
            processTask(task);
        }
    }
}

// 正确示例
public class DataManager {
    // 正确：频繁查找使用Map
    private Map<String, User> users = new ConcurrentHashMap<>();
    
    public User findUserById(String id) {
        return users.get(id);  // O(1)时间复杂度
    }
    
    // 正确：去重使用Set
    public Set<String> getUniqueItems(List<String> items) {
        return new HashSet<>(items);  // O(n)时间复杂度
    }
    
    // 正确：队列操作使用LinkedList或Queue
    public void processQueue(Queue<Task> tasks) {
        while (!tasks.isEmpty()) {
            Task task = tasks.poll();  // O(1)操作
            processTask(task);
        }
    }
    
    // 正确：需要排序的场景使用TreeSet
    private Set<String> sortedItems = new TreeSet<>();
    
    // 正确：LRU缓存使用LinkedHashMap
    private Map<String, Object> lruCache = new LinkedHashMap<String, Object>(16, 0.75f, true) {
        @Override
        protected boolean removeEldestEntry(Map.Entry<String, Object> eldest) {
            return size() > 100;
        }
    };
}
```

### 5.9 JVM参数调优 (Major)

#### 5.9.1 堆内存配置
- **检查方法**: 检查JVM启动参数配置
- **检查标准**: 根据应用特点合理配置堆内存大小
- **不正确实例**:
```bash
# 错误示例 - JVM参数配置不当
# 错误：堆内存过小
java -Xms128m -Xmx256m -jar app.jar

# 错误：新生代配置不合理
java -Xms2g -Xmx2g -Xmn1800m -jar app.jar  # 新生代过大

# 错误：没有配置GC参数
java -Xms1g -Xmx1g -jar app.jar

# 正确示例
# 正确：合理的内存配置
java -Xms2g -Xmx2g \
     -Xmn800m \
     -XX:MetaspaceSize=256m \
     -XX:MaxMetaspaceSize=512m \
     -XX:+UseG1GC \
     -XX:MaxGCPauseMillis=200 \
     -XX:+PrintGCDetails \
     -XX:+PrintGCTimeStamps \
     -Xloggc:gc.log \
     -jar app.jar

# 针对不同场景的配置
# 高并发低延迟场景
java -Xms4g -Xmx4g \
     -XX:+UseZGC \
     -XX:+UnlockExperimentalVMOptions \
     -jar app.jar

# 大数据处理场景
java -Xms8g -Xmx8g \
     -XX:+UseParallelGC \
     -XX:ParallelGCThreads=8 \
     -jar app.jar
```

### 5.10 垃圾回收优化 (Major)

#### 5.10.1 GC算法选择
- **检查方法**: 分析GC日志，监控GC性能指标
- **检查标准**: 根据应用特点选择合适的GC算法
- **不正确实例**:
```java
// 错误示例 - 产生大量垃圾对象
public class DataProcessor {
    public String processData(List<String> data) {
        String result = "";
        for (String item : data) {
            result += item + ",";  // 产生大量临时String对象
        }
        return result;
    }
    
    // 错误：频繁创建大对象
    public void processLargeData() {
        for (int i = 0; i < 1000; i++) {
            byte[] largeArray = new byte[1024 * 1024];  // 每次创建1MB数组
            processArray(largeArray);
        }
    }
}

// 正确示例
public class DataProcessor {
    // 正确：减少对象创建
    public String processData(List<String> data) {
        return String.join(",", data);
    }
    
    // 正确：重用大对象
    private final byte[] reusableBuffer = new byte[1024 * 1024];
    
    public void processLargeData() {
        for (int i = 0; i < 1000; i++) {
            Arrays.fill(reusableBuffer, (byte) 0);  // 重用数组
            processArray(reusableBuffer);
        }
    }
    
    // 正确：使用对象池
    private final ObjectPool<StringBuilder> stringBuilderPool = 
        new GenericObjectPool<>(new StringBuilderFactory());
    
    public String buildString(List<String> parts) {
        StringBuilder sb = null;
        try {
            sb = stringBuilderPool.borrowObject();
            for (String part : parts) {
                sb.append(part);
            }
            return sb.toString();
        } finally {
            if (sb != null) {
                sb.setLength(0);
                stringBuilderPool.returnObject(sb);
            }
        }
    }
}
```

### 5.11 序列化优化 (Minor)

#### 5.11.1 选择高效的序列化方式
- **检查方法**: 检查序列化方式的选择和配置
- **检查标准**: 根据性能要求选择合适的序列化框架
- **不正确实例**:
```java
// 错误示例 - 低效的序列化
public class DataSerializer {
    // 错误：使用Java原生序列化
    public byte[] serialize(Object obj) throws IOException {
        ByteArrayOutputStream baos = new ByteArrayOutputStream();
        ObjectOutputStream oos = new ObjectOutputStream(baos);
        oos.writeObject(obj);  // 效率低，体积大
        return baos.toByteArray();
    }
    
    // 错误：没有配置Jackson优化
    private ObjectMapper mapper = new ObjectMapper();
    
    public String toJson(Object obj) throws JsonProcessingException {
        return mapper.writeValueAsString(obj);  // 默认配置，性能不佳
    }
}

// 正确示例
public class DataSerializer {
    // 正确：使用高效的序列化框架
    private final ObjectMapper optimizedMapper;
    
    public DataSerializer() {
        this.optimizedMapper = new ObjectMapper()
            .configure(DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES, false)
            .configure(SerializationFeature.FAIL_ON_EMPTY_BEANS, false)
            .setSerializationInclusion(JsonInclude.Include.NON_NULL)
            .registerModule(new JavaTimeModule());
    }
    
    public String toJson(Object obj) throws JsonProcessingException {
        return optimizedMapper.writeValueAsString(obj);
    }
    
    // 正确：使用Protobuf进行高性能序列化
    public byte[] serializeWithProtobuf(UserProto.User user) {
        return user.toByteArray();
    }
    
    public UserProto.User deserializeWithProtobuf(byte[] data) throws InvalidProtocolBufferException {
        return UserProto.User.parseFrom(data);
    }
    
    // 正确：使用Kryo进行快速序列化
    private final ThreadLocal<Kryo> kryoThreadLocal = ThreadLocal.withInitial(() -> {
        Kryo kryo = new Kryo();
        kryo.setReferences(false);
        kryo.setRegistrationRequired(false);
        return kryo;
    });
    
    public byte[] serializeWithKryo(Object obj) {
        Kryo kryo = kryoThreadLocal.get();
        ByteArrayOutputStream baos = new ByteArrayOutputStream();
        Output output = new Output(baos);
        kryo.writeObject(output, obj);
        output.close();
        return baos.toByteArray();
    }
}
```

### 5.12 网络传输优化 (Minor)

#### 5.12.1 HTTP客户端优化
- **检查方法**: 检查HTTP客户端的配置和使用方式
- **检查标准**: 使用连接池，配置合理的超时时间
- **不正确实例**:
```java
// 错误示例 - HTTP客户端配置不当
@Service
public class ApiClient {
    // 错误：每次创建新的HTTP客户端
    public String callApi(String url) throws IOException {
        HttpURLConnection connection = (HttpURLConnection) new URL(url).openConnection();
        connection.setRequestMethod("GET");
        // 没有设置超时时间
        
        try (BufferedReader reader = new BufferedReader(
                new InputStreamReader(connection.getInputStream()))) {
            return reader.lines().collect(Collectors.joining("\n"));
        }
    }
    
    // 错误：使用默认的RestTemplate
    @Autowired
    private RestTemplate restTemplate;
    
    public String getData(String url) {
        return restTemplate.getForObject(url, String.class);  // 没有连接池
    }
}

// 正确示例
@Service
public class ApiClient {
    private final RestTemplate restTemplate;
    private final WebClient webClient;
    
    public ApiClient() {
        // 正确：配置连接池和超时
        HttpComponentsClientHttpRequestFactory factory = new HttpComponentsClientHttpRequestFactory();
        factory.setConnectTimeout(5000);
        factory.setReadTimeout(10000);
        
        PoolingHttpClientConnectionManager connectionManager = new PoolingHttpClientConnectionManager();
        connectionManager.setMaxTotal(100);
        connectionManager.setDefaultMaxPerRoute(20);
        
        CloseableHttpClient httpClient = HttpClients.custom()
            .setConnectionManager(connectionManager)
            .build();
        
        factory.setHttpClient(httpClient);
        this.restTemplate = new RestTemplate(factory);
        
        // 正确：配置WebClient连接池
        ConnectionProvider connectionProvider = ConnectionProvider.builder("custom")
            .maxConnections(100)
            .maxIdleTime(Duration.ofSeconds(30))
            .maxLifeTime(Duration.ofMinutes(5))
            .build();
        
        HttpClient httpClientReactive = HttpClient.create(connectionProvider)
            .option(ChannelOption.CONNECT_TIMEOUT_MILLIS, 5000)
            .responseTimeout(Duration.ofSeconds(10));
        
        this.webClient = WebClient.builder()
            .clientConnector(new ReactorClientHttpConnector(httpClientReactive))
            .build();
    }
    
    public String getData(String url) {
        return restTemplate.getForObject(url, String.class);
    }
    
    public Mono<String> getDataAsync(String url) {
        return webClient.get()
            .uri(url)
            .retrieve()
            .bodyToMono(String.class);
    }
}
```

### 5.13 代码层面优化 (Minor)

#### 5.13.1 避免不必要的计算
- **检查方法**: 检查代码中是否有重复计算或可以缓存的计算
- **检查标准**: 将重复计算提取到循环外，使用缓存存储计算结果
- **不正确实例**:
```java
// 错误示例 - 重复计算
public class Calculator {
    public double calculateArea(List<Circle> circles) {
        double totalArea = 0;
        for (Circle circle : circles) {
            // 错误：重复计算PI
            double area = Math.PI * circle.getRadius() * circle.getRadius();
            totalArea += area;
        }
        return totalArea;
    }
    
    // 错误：在循环中进行复杂计算
    public List<String> formatNumbers(List<Integer> numbers) {
        List<String> formatted = new ArrayList<>();
        for (Integer number : numbers) {
            // 错误：每次都创建DateFormat
            DecimalFormat df = new DecimalFormat("#,###.00");
            formatted.add(df.format(number));
        }
        return formatted;
    }
    
    // 错误：重复的字符串操作
    public String buildQuery(String table, List<String> columns, String condition) {
        String query = "SELECT ";
        for (int i = 0; i < columns.size(); i++) {
            query += columns.get(i);
            if (i < columns.size() - 1) {
                query += ", ";  // 重复的字符串拼接
            }
        }
        query += " FROM " + table + " WHERE " + condition;
        return query;
    }
}

// 正确示例
public class Calculator {
    private static final double PI = Math.PI;  // 缓存常量
    
    public double calculateArea(List<Circle> circles) {
        double totalArea = 0;
        for (Circle circle : circles) {
            // 正确：使用缓存的PI值
            double radius = circle.getRadius();
            double area = PI * radius * radius;
            totalArea += area;
        }
        return totalArea;
    }
    
    // 正确：重用对象
    private static final ThreadLocal<DecimalFormat> DECIMAL_FORMAT = 
        ThreadLocal.withInitial(() -> new DecimalFormat("#,###.00"));
    
    public List<String> formatNumbers(List<Integer> numbers) {
        DecimalFormat df = DECIMAL_FORMAT.get();
        return numbers.stream()
            .map(df::format)
            .collect(Collectors.toList());
    }
    
    // 正确：使用StringBuilder和预计算
    public String buildQuery(String table, List<String> columns, String condition) {
        StringBuilder query = new StringBuilder("SELECT ");
        query.append(String.join(", ", columns));
        query.append(" FROM ").append(table);
        query.append(" WHERE ").append(condition);
        return query.toString();
    }
    
    // 正确：使用缓存避免重复计算
    private final Map<String, String> queryCache = new ConcurrentHashMap<>();
    
    public String getCachedQuery(String table, List<String> columns, String condition) {
        String key = table + "|" + String.join(",", columns) + "|" + condition;
        return queryCache.computeIfAbsent(key, k -> buildQuery(table, columns, condition));
    }
}
```

### 5.14 资源预加载 (Minor)

#### 5.14.1 配置和静态资源预加载
- **检查方法**: 检查应用启动时是否预加载必要的资源
- **检查标准**: 在应用启动时预加载配置、缓存等资源
- **不正确实例**:
```java
// 错误示例 - 懒加载导致首次访问慢
@Service
public class ConfigService {
    private Map<String, String> configCache;
    
    // 错误：首次调用时才加载配置
    public String getConfig(String key) {
        if (configCache == null) {
            loadConfig();  // 首次访问时加载，导致延迟
        }
        return configCache.get(key);
    }
    
    private void loadConfig() {
        // 加载配置的耗时操作
        configCache = loadFromDatabase();
    }
}

// 错误示例 - 没有预热缓存
@Service
public class UserService {
    @Autowired
    private RedisTemplate<String, Object> redisTemplate;
    
    public User getUser(Long id) {
        String key = "user:" + id;
        User user = (User) redisTemplate.opsForValue().get(key);
        if (user == null) {
            user = userRepository.findById(id);  // 缓存未命中时查询数据库
            redisTemplate.opsForValue().set(key, user, 1, TimeUnit.HOURS);
        }
        return user;
    }
}

// 正确示例
@Service
public class ConfigService {
    private final Map<String, String> configCache = new ConcurrentHashMap<>();
    
    // 正确：应用启动时预加载
    @PostConstruct
    public void init() {
        loadConfig();
        log.info("Configuration loaded: {} items", configCache.size());
    }
    
    private void loadConfig() {
        Map<String, String> configs = loadFromDatabase();
        configCache.putAll(configs);
    }
    
    public String getConfig(String key) {
        return configCache.get(key);
    }
    
    // 正确：定时刷新配置
    @Scheduled(fixedRate = 300000)  // 5分钟刷新一次
    public void refreshConfig() {
        try {
            loadConfig();
            log.debug("Configuration refreshed");
        } catch (Exception e) {
            log.error("Failed to refresh configuration", e);
        }
    }
}

@Service
public class UserService {
    @Autowired
    private RedisTemplate<String, Object> redisTemplate;
    
    @Autowired
    private UserRepository userRepository;
    
    // 正确：应用启动时预热热点数据
    @PostConstruct
    public void warmUpCache() {
        CompletableFuture.runAsync(() -> {
            try {
                List<User> hotUsers = userRepository.findHotUsers(100);
                for (User user : hotUsers) {
                    String key = "user:" + user.getId();
                    redisTemplate.opsForValue().set(key, user, 1, TimeUnit.HOURS);
                }
                log.info("Cache warmed up with {} hot users", hotUsers.size());
            } catch (Exception e) {
                log.error("Failed to warm up cache", e);
            }
        });
    }
    
    public User getUser(Long id) {
        String key = "user:" + id;
        User user = (User) redisTemplate.opsForValue().get(key);
        if (user == null) {
            user = userRepository.findById(id);
            if (user != null) {
                redisTemplate.opsForValue().set(key, user, 1, TimeUnit.HOURS);
            }
        }
        return user;
    }
}
```

### 5.15 懒加载策略 (Minor)

#### 5.15.1 合理使用懒加载
- **检查方法**: 检查是否在合适的场景使用懒加载
- **检查标准**: 对于大对象或不常用的资源使用懒加载
- **不正确实例**:
```java
// 错误示例 - 不合理的懒加载
@Service
public class ReportService {
    // 错误：频繁使用的对象使用懒加载
    private volatile ObjectMapper objectMapper;
    
    public String generateReport(Object data) {
        if (objectMapper == null) {
            synchronized (this) {
                if (objectMapper == null) {
                    objectMapper = new ObjectMapper();  // 每次都要检查
                }
            }
        }
        return objectMapper.writeValueAsString(data);
    }
    
    // 错误：小对象使用懒加载增加复杂度
    private volatile List<String> statusList;
    
    public List<String> getStatusList() {
        if (statusList == null) {
            synchronized (this) {
                if (statusList == null) {
                    statusList = Arrays.asList("ACTIVE", "INACTIVE");  // 简单对象不需要懒加载
                }
            }
        }
        return statusList;
    }
}

// 正确示例
@Service
public class ReportService {
    // 正确：频繁使用的对象直接初始化
    private final ObjectMapper objectMapper = new ObjectMapper();
    
    // 正确：简单对象直接初始化
    private final List<String> statusList = Arrays.asList("ACTIVE", "INACTIVE");
    
    // 正确：大对象或昂贵资源使用懒加载
    private volatile ExpensiveResource expensiveResource;
    
    public ExpensiveResource getExpensiveResource() {
        if (expensiveResource == null) {
            synchronized (this) {
                if (expensiveResource == null) {
                    expensiveResource = createExpensiveResource();
                }
            }
        }
        return expensiveResource;
    }
    
    // 正确：使用Supplier实现懒加载
    private final Supplier<List<ComplexData>> complexDataSupplier = 
        Suppliers.memoize(this::loadComplexData);
    
    public List<ComplexData> getComplexData() {
        return complexDataSupplier.get();
    }
    
    private List<ComplexData> loadComplexData() {
        // 加载复杂数据的耗时操作
        return complexDataRepository.findAll();
    }
    
    // 正确：JPA懒加载配置
    @Entity
    public class Order {
        @OneToMany(fetch = FetchType.LAZY, mappedBy = "order")
        private List<OrderItem> items;  // 大集合使用懒加载
        
        @ManyToOne(fetch = FetchType.EAGER)
        private Customer customer;  // 常用关联使用立即加载
    }
}
```

### 5.16 对象池使用 (Minor)

#### 5.16.1 合理使用对象池
- **检查方法**: 检查是否在合适的场景使用对象池
- **检查标准**: 对于创建成本高的对象使用对象池
- **不正确实例**:
```java
// 错误示例 - 不合理的对象池使用
public class StringProcessor {
    // 错误：为简单对象创建对象池
    private final ObjectPool<String> stringPool = new GenericObjectPool<>(new StringFactory());
    
    public String processString(String input) {
        String str = null;
        try {
            str = stringPool.borrowObject();  // String对象池没有意义
            return str + input;
        } finally {
            if (str != null) {
                stringPool.returnObject(str);
            }
        }
    }
    
    // 错误：没有正确配置对象池
    private final ObjectPool<StringBuilder> builderPool = new GenericObjectPool<>(new StringBuilderFactory());
    
    public String buildString(List<String> parts) {
        StringBuilder sb = null;
        try {
            sb = builderPool.borrowObject();
            for (String part : parts) {
                sb.append(part);
            }
            return sb.toString();
        } finally {
            if (sb != null) {
                // 错误：没有清理对象状态
                builderPool.returnObject(sb);
            }
        }
    }
}

// 正确示例
public class StringProcessor {
    // 正确：为创建成本高的对象使用对象池
    private final ObjectPool<StringBuilder> stringBuilderPool;
    private final ObjectPool<MessageDigest> digestPool;
    
    public StringProcessor() {
        // 正确：配置对象池参数
        GenericObjectPoolConfig<StringBuilder> config = new GenericObjectPoolConfig<>();
        config.setMaxTotal(10);
        config.setMaxIdle(5);
        config.setMinIdle(2);
        config.setTestOnBorrow(true);
        config.setTestOnReturn(true);
        
        this.stringBuilderPool = new GenericObjectPool<>(new StringBuilderFactory(), config);
        
        // 正确：为昂贵对象使用对象池
        this.digestPool = new GenericObjectPool<>(new MessageDigestFactory());
    }
    
    public String buildString(List<String> parts) {
        StringBuilder sb = null;
        try {
            sb = stringBuilderPool.borrowObject();
            for (String part : parts) {
                sb.append(part);
            }
            return sb.toString();
        } catch (Exception e) {
            throw new RuntimeException("Failed to build string", e);
        } finally {
            if (sb != null) {
                // 正确：清理对象状态后归还
                sb.setLength(0);
                try {
                    stringBuilderPool.returnObject(sb);
                } catch (Exception e) {
                    log.warn("Failed to return StringBuilder to pool", e);
                }
            }
        }
    }
    
    public String calculateHash(String input) {
        MessageDigest digest = null;
        try {
            digest = digestPool.borrowObject();
            byte[] hash = digest.digest(input.getBytes());
            return bytesToHex(hash);
        } catch (Exception e) {
            throw new RuntimeException("Failed to calculate hash", e);
        } finally {
            if (digest != null) {
                // 正确：重置对象状态
                digest.reset();
                try {
                    digestPool.returnObject(digest);
                } catch (Exception e) {
                    log.warn("Failed to return MessageDigest to pool", e);
                }
            }
        }
    }
    
    // 正确：实现对象工厂
    private static class StringBuilderFactory extends BasePooledObjectFactory<StringBuilder> {
        @Override
        public StringBuilder create() {
            return new StringBuilder(256);
        }
        
        @Override
        public PooledObject<StringBuilder> wrap(StringBuilder obj) {
            return new DefaultPooledObject<>(obj);
        }
        
        @Override
        public boolean validateObject(PooledObject<StringBuilder> p) {
            return p.getObject() != null;
        }
        
        @Override
        public void passivateObject(PooledObject<StringBuilder> p) {
            p.getObject().setLength(0);  // 清理状态
        }
    }
    
    private static class MessageDigestFactory extends BasePooledObjectFactory<MessageDigest> {
        @Override
        public MessageDigest create() throws Exception {
            return MessageDigest.getInstance("SHA-256");
        }
        
        @Override
        public PooledObject<MessageDigest> wrap(MessageDigest obj) {
            return new DefaultPooledObject<>(obj);
        }
        
        @Override
        public void passivateObject(PooledObject<MessageDigest> p) {
            p.getObject().reset();
        }
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

### 9.4 错误处理 (Critical)

#### 9.4.1 统一错误响应格式
- **检查方法**: 检查API错误响应是否使用统一格式
- **检查标准**: 所有API错误响应必须使用统一的格式和状态码
- **不正确实例**:
```java
// 错误示例 - 不统一的错误响应
@RestController
public class UserController {
    
    @GetMapping("/users/{id}")
    public ResponseEntity<?> getUser(@PathVariable Long id) {
        try {
            User user = userService.findById(id);
            return ResponseEntity.ok(user);
        } catch (UserNotFoundException e) {
            // 错误：不同的错误格式
            return ResponseEntity.notFound().build();
        } catch (Exception e) {
            // 错误：直接返回异常信息
            return ResponseEntity.status(500).body(e.getMessage());
        }
    }
    
    @PostMapping("/users")
    public ResponseEntity<?> createUser(@RequestBody User user) {
        try {
            User created = userService.create(user);
            return ResponseEntity.ok(created);
        } catch (ValidationException e) {
            // 错误：不同的错误格式
            Map<String, String> errors = new HashMap<>();
            errors.put("error", e.getMessage());
            return ResponseEntity.badRequest().body(errors);
        }
    }
}

// 正确示例 - 统一错误响应格式
@RestController
public class UserController {
    
    @GetMapping("/users/{id}")
    public ResponseEntity<ApiResponse<User>> getUser(@PathVariable Long id) {
        User user = userService.findById(id);
        return ResponseEntity.ok(ApiResponse.success(user));
    }
    
    @PostMapping("/users")
    public ResponseEntity<ApiResponse<User>> createUser(@Valid @RequestBody CreateUserRequest request) {
        User created = userService.create(request);
        return ResponseEntity.ok(ApiResponse.success(created));
    }
}

// 统一错误响应格式
@Data
@Builder
public class ApiResponse<T> {
    private boolean success;
    private T data;
    private ErrorInfo error;
    
    public static <T> ApiResponse<T> success(T data) {
        return ApiResponse.<T>builder()
            .success(true)
            .data(data)
            .build();
    }
    
    public static <T> ApiResponse<T> error(String code, String message) {
        return ApiResponse.<T>builder()
            .success(false)
            .error(ErrorInfo.builder()
                .code(code)
                .message(message)
                .timestamp(Instant.now())
                .build())
            .build();
    }
}

@Data
@Builder
public class ErrorInfo {
    private String code;
    private String message;
    private Instant timestamp;
    private List<FieldError> fieldErrors;
}

// 全局异常处理器
@RestControllerAdvice
public class GlobalExceptionHandler {
    
    @ExceptionHandler(UserNotFoundException.class)
    public ResponseEntity<ApiResponse<Void>> handleUserNotFound(UserNotFoundException e) {
        return ResponseEntity.status(HttpStatus.NOT_FOUND)
            .body(ApiResponse.error("USER_NOT_FOUND", e.getMessage()));
    }
    
    @ExceptionHandler(MethodArgumentNotValidException.class)
    public ResponseEntity<ApiResponse<Void>> handleValidation(MethodArgumentNotValidException e) {
        List<FieldError> fieldErrors = e.getBindingResult().getFieldErrors().stream()
            .map(error -> FieldError.builder()
                .field(error.getField())
                .message(error.getDefaultMessage())
                .rejectedValue(error.getRejectedValue())
                .build())
            .collect(Collectors.toList());
        
        ErrorInfo errorInfo = ErrorInfo.builder()
            .code("VALIDATION_FAILED")
            .message("Request validation failed")
            .timestamp(Instant.now())
            .fieldErrors(fieldErrors)
            .build();
        
        return ResponseEntity.badRequest()
            .body(ApiResponse.<Void>builder()
                .success(false)
                .error(errorInfo)
                .build());
    }
    
    @ExceptionHandler(Exception.class)
    public ResponseEntity<ApiResponse<Void>> handleGeneral(Exception e) {
        log.error("Unexpected error occurred", e);
        return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR)
            .body(ApiResponse.error("INTERNAL_ERROR", "An unexpected error occurred"));
    }
}
```

#### 9.4.2 HTTP状态码使用规范
- **检查方法**: 检查API返回的HTTP状态码是否符合RESTful规范
- **检查标准**: 正确使用2xx、4xx、5xx状态码，避免滥用200状态码
- **不正确实例**:
```java
// 错误示例 - 状态码使用不当
@RestController
public class OrderController {
    
    @GetMapping("/orders/{id}")
    public ResponseEntity<ApiResponse<Order>> getOrder(@PathVariable Long id) {
        try {
            Order order = orderService.findById(id);
            return ResponseEntity.ok(ApiResponse.success(order));
        } catch (OrderNotFoundException e) {
            // 错误：应该返回404，而不是200
            return ResponseEntity.ok(ApiResponse.error("ORDER_NOT_FOUND", e.getMessage()));
        }
    }
    
    @PostMapping("/orders")
    public ResponseEntity<ApiResponse<Order>> createOrder(@RequestBody CreateOrderRequest request) {
        try {
            Order order = orderService.create(request);
            // 错误：创建成功应该返回201，而不是200
            return ResponseEntity.ok(ApiResponse.success(order));
        } catch (InsufficientStockException e) {
            // 错误：业务逻辑错误应该返回400，而不是500
            return ResponseEntity.status(500).body(ApiResponse.error("STOCK_ERROR", e.getMessage()));
        }
    }
    
    @DeleteMapping("/orders/{id}")
    public ResponseEntity<ApiResponse<Void>> deleteOrder(@PathVariable Long id) {
        orderService.delete(id);
        // 错误：删除成功应该返回204，而不是200
        return ResponseEntity.ok(ApiResponse.success(null));
    }
}

// 正确示例 - 正确使用HTTP状态码
@RestController
public class OrderController {
    
    @GetMapping("/orders/{id}")
    public ResponseEntity<ApiResponse<Order>> getOrder(@PathVariable Long id) {
        Order order = orderService.findById(id);
        return ResponseEntity.ok(ApiResponse.success(order));  // 200 OK
    }
    
    @PostMapping("/orders")
    public ResponseEntity<ApiResponse<Order>> createOrder(@Valid @RequestBody CreateOrderRequest request) {
        Order order = orderService.create(request);
        return ResponseEntity.status(HttpStatus.CREATED)  // 201 Created
            .body(ApiResponse.success(order));
    }
    
    @PutMapping("/orders/{id}")
    public ResponseEntity<ApiResponse<Order>> updateOrder(
            @PathVariable Long id, 
            @Valid @RequestBody UpdateOrderRequest request) {
        Order order = orderService.update(id, request);
        return ResponseEntity.ok(ApiResponse.success(order));  // 200 OK
    }
    
    @DeleteMapping("/orders/{id}")
    public ResponseEntity<Void> deleteOrder(@PathVariable Long id) {
        orderService.delete(id);
        return ResponseEntity.noContent().build();  // 204 No Content
    }
    
    @GetMapping("/orders")
    public ResponseEntity<ApiResponse<PageResult<Order>>> getOrders(
            @RequestParam(defaultValue = "0") int page,
            @RequestParam(defaultValue = "20") int size) {
        PageResult<Order> orders = orderService.findAll(page, size);
        return ResponseEntity.ok(ApiResponse.success(orders));  // 200 OK
    }
}

// 异常处理器中正确使用状态码
@RestControllerAdvice
public class GlobalExceptionHandler {
    
    @ExceptionHandler(EntityNotFoundException.class)
    public ResponseEntity<ApiResponse<Void>> handleNotFound(EntityNotFoundException e) {
        return ResponseEntity.status(HttpStatus.NOT_FOUND)  // 404 Not Found
            .body(ApiResponse.error("RESOURCE_NOT_FOUND", e.getMessage()));
    }
    
    @ExceptionHandler(MethodArgumentNotValidException.class)
    public ResponseEntity<ApiResponse<Void>> handleValidation(MethodArgumentNotValidException e) {
        return ResponseEntity.status(HttpStatus.BAD_REQUEST)  // 400 Bad Request
            .body(ApiResponse.error("VALIDATION_FAILED", "Request validation failed"));
    }
    
    @ExceptionHandler(InsufficientStockException.class)
    public ResponseEntity<ApiResponse<Void>> handleBusinessLogicError(InsufficientStockException e) {
        return ResponseEntity.status(HttpStatus.BAD_REQUEST)  // 400 Bad Request
            .body(ApiResponse.error("INSUFFICIENT_STOCK", e.getMessage()));
    }
    
    @ExceptionHandler(AccessDeniedException.class)
    public ResponseEntity<ApiResponse<Void>> handleAccessDenied(AccessDeniedException e) {
        return ResponseEntity.status(HttpStatus.FORBIDDEN)  // 403 Forbidden
            .body(ApiResponse.error("ACCESS_DENIED", "Access denied"));
    }
    
    @ExceptionHandler(AuthenticationException.class)
    public ResponseEntity<ApiResponse<Void>> handleAuthentication(AuthenticationException e) {
        return ResponseEntity.status(HttpStatus.UNAUTHORIZED)  // 401 Unauthorized
            .body(ApiResponse.error("AUTHENTICATION_FAILED", "Authentication failed"));
    }
    
    @ExceptionHandler(Exception.class)
    public ResponseEntity<ApiResponse<Void>> handleGeneral(Exception e) {
        log.error("Unexpected error occurred", e);
        return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR)  // 500 Internal Server Error
            .body(ApiResponse.error("INTERNAL_ERROR", "An unexpected error occurred"));
    }
}
```

### 9.5 文档规范 (Major)

#### 9.5.1 API文档完整性
- **检查方法**: 检查是否使用Swagger/OpenAPI生成完整的API文档
- **检查标准**: 所有公开API必须有完整的文档，包括参数、响应、示例
- **不正确实例**:
```java
// 错误示例 - 缺少API文档
@RestController
public class ProductController {
    
    @GetMapping("/products")
    public ResponseEntity<List<Product>> getProducts(
            @RequestParam String category,
            @RequestParam Integer page,
            @RequestParam Integer size) {
        // 缺少文档说明
        return ResponseEntity.ok(productService.findByCategory(category, page, size));
    }
    
    @PostMapping("/products")
    public ResponseEntity<Product> createProduct(@RequestBody Product product) {
        // 缺少参数说明和响应说明
        return ResponseEntity.ok(productService.create(product));
    }
}

// 正确示例 - 完整的API文档
@RestController
@RequestMapping("/api/v1/products")
@Tag(name = "Product", description = "产品管理API")
public class ProductController {
    
    @Operation(
        summary = "获取产品列表",
        description = "根据分类和分页参数获取产品列表",
        responses = {
            @ApiResponse(
                responseCode = "200",
                description = "获取成功",
                content = @Content(
                    mediaType = "application/json",
                    schema = @Schema(implementation = PageResult.class)
                )
            ),
            @ApiResponse(
                responseCode = "400",
                description = "参数错误",
                content = @Content(
                    mediaType = "application/json",
                    schema = @Schema(implementation = ApiResponse.class)
                )
            )
        }
    )
    @GetMapping
    public ResponseEntity<ApiResponse<PageResult<Product>>> getProducts(
            @Parameter(description = "产品分类", required = true, example = "electronics")
            @RequestParam String category,
            
            @Parameter(description = "页码，从0开始", example = "0")
            @RequestParam(defaultValue = "0") Integer page,
            
            @Parameter(description = "每页大小，最大100", example = "20")
            @RequestParam(defaultValue = "20") @Max(100) Integer size,
            
            @Parameter(description = "排序字段", example = "name")
            @RequestParam(defaultValue = "id") String sortBy,
            
            @Parameter(description = "排序方向", example = "asc")
            @RequestParam(defaultValue = "asc") String sortDir) {
        
        PageResult<Product> products = productService.findByCategory(
            category, page, size, sortBy, sortDir);
        return ResponseEntity.ok(ApiResponse.success(products));
    }
    
    @Operation(
        summary = "创建产品",
        description = "创建新的产品信息",
        requestBody = @io.swagger.v3.oas.annotations.parameters.RequestBody(
            description = "产品信息",
            required = true,
            content = @Content(
                mediaType = "application/json",
                schema = @Schema(implementation = CreateProductRequest.class),
                examples = @ExampleObject(
                    name = "创建产品示例",
                    value = """
                    {
                        "name": "iPhone 15",
                        "category": "electronics",
                        "price": 999.99,
                        "description": "Latest iPhone model",
                        "stock": 100
                    }
                    """
                )
            )
        ),
        responses = {
            @ApiResponse(
                responseCode = "201",
                description = "创建成功",
                content = @Content(
                    mediaType = "application/json",
                    schema = @Schema(implementation = Product.class)
                )
            ),
            @ApiResponse(
                responseCode = "400",
                description = "参数验证失败",
                content = @Content(
                    mediaType = "application/json",
                    schema = @Schema(implementation = ApiResponse.class)
                )
            )
        }
    )
    @PostMapping
    public ResponseEntity<ApiResponse<Product>> createProduct(
            @Valid @RequestBody CreateProductRequest request) {
        Product product = productService.create(request);
        return ResponseEntity.status(HttpStatus.CREATED)
            .body(ApiResponse.success(product));
    }
    
    @Operation(
        summary = "获取产品详情",
        description = "根据产品ID获取详细信息"
    )
    @GetMapping("/{id}")
    public ResponseEntity<ApiResponse<Product>> getProduct(
            @Parameter(description = "产品ID", required = true, example = "1")
            @PathVariable Long id) {
        Product product = productService.findById(id);
        return ResponseEntity.ok(ApiResponse.success(product));
    }
}

// DTO类也需要完整的文档
@Schema(description = "创建产品请求")
@Data
public class CreateProductRequest {
    
    @Schema(description = "产品名称", required = true, example = "iPhone 15")
    @NotBlank(message = "产品名称不能为空")
    @Size(max = 100, message = "产品名称长度不能超过100")
    private String name;
    
    @Schema(description = "产品分类", required = true, example = "electronics")
    @NotBlank(message = "产品分类不能为空")
    private String category;
    
    @Schema(description = "产品价格", required = true, example = "999.99")
    @NotNull(message = "产品价格不能为空")
    @DecimalMin(value = "0.01", message = "产品价格必须大于0")
    private BigDecimal price;
    
    @Schema(description = "产品描述", example = "Latest iPhone model")
    @Size(max = 500, message = "产品描述长度不能超过500")
    private String description;
    
    @Schema(description = "库存数量", required = true, example = "100")
    @NotNull(message = "库存数量不能为空")
    @Min(value = 0, message = "库存数量不能为负数")
    private Integer stock;
}
```

#### 9.5.2 接口变更文档
- **检查方法**: 检查API变更是否有详细的变更日志和迁移指南
- **检查标准**: 每次API变更必须记录变更内容、影响范围、迁移方案
- **不正确实例**:
```java
// 错误示例 - 直接修改现有API，没有版本控制和文档
@RestController
public class UserController {
    
    @GetMapping("/users/{id}")
    public ResponseEntity<User> getUser(@PathVariable Long id) {
        // 直接修改返回格式，破坏向后兼容性
        User user = userService.findById(id);
        // 新增字段但没有文档说明
        user.setLastLoginTime(userService.getLastLoginTime(id));
        return ResponseEntity.ok(user);
    }
}

// 正确示例 - 版本化API和完整的变更文档
@RestController
@RequestMapping("/api/v2/users")
public class UserV2Controller {
    
    @Operation(
        summary = "获取用户信息 (v2)",
        description = """
        获取用户详细信息，v2版本相比v1版本的变更：
        
        **新增字段：**
        - lastLoginTime: 最后登录时间
        - profileCompleteness: 资料完整度百分比
        - preferences: 用户偏好设置
        
        **字段变更：**
        - email: 现在总是返回，v1版本中可能为null
        - createdAt: 格式从timestamp改为ISO 8601字符串
        
        **废弃字段：**
        - isActive: 使用status字段替代
        
        **迁移指南：**
        1. 更新客户端以处理新增字段
        2. 将isActive字段的使用改为status字段
        3. 更新日期解析逻辑以支持ISO 8601格式
        """,
        responses = {
            @ApiResponse(
                responseCode = "200",
                description = "获取成功",
                content = @Content(
                    mediaType = "application/json",
                    schema = @Schema(implementation = UserV2Response.class)
                )
            )
        }
    )
    @GetMapping("/{id}")
    public ResponseEntity<ApiResponse<UserV2Response>> getUser(
            @Parameter(description = "用户ID", required = true)
            @PathVariable Long id) {
        UserV2Response user = userService.findByIdV2(id);
        return ResponseEntity.ok(ApiResponse.success(user));
    }
}

// 变更日志文档
/**
 * API变更日志
 * 
 * ## v2.1.0 (2024-01-15)
 * 
 * ### 新增
 * - POST /api/v2/users/{id}/preferences - 更新用户偏好设置
 * - GET /api/v2/users/{id}/activity - 获取用户活动记录
 * 
 * ### 变更
 * - GET /api/v2/users/{id} 响应中新增 `preferences` 字段
 * - PUT /api/v2/users/{id} 请求体中 `email` 字段现在支持验证
 * 
 * ### 废弃
 * - GET /api/v1/users/{id} 将在v3.0.0中移除，请迁移到v2版本
 * 
 * ### 修复
 * - 修复分页查询中总数计算错误的问题
 * 
 * ### 迁移指南
 * 1. 如果使用v1 API，建议尽快迁移到v2
 * 2. 新的preferences字段为可选，现有客户端可以忽略
 * 3. 更新邮箱时需要通过邮箱验证流程
 * 
 * ## v2.0.0 (2023-12-01)
 * 
 * ### 重大变更
 * - 所有日期字段格式从timestamp改为ISO 8601字符串
 * - 移除已废弃的isActive字段，使用status字段
 * - 错误响应格式标准化
 * 
 * ### 迁移指南
 * 详见：https://docs.example.com/api/migration/v1-to-v2
 */
```

### 9.6 请求验证 (Major)

#### 9.6.1 输入参数验证
- **检查方法**: 检查所有API接口是否对输入参数进行充分验证
- **检查标准**: 使用Bean Validation注解，自定义验证器，验证所有必要字段
- **不正确实例**:
```java
// 错误示例 - 缺少输入验证
@RestController
public class UserController {
    
    @PostMapping("/users")
    public ResponseEntity<User> createUser(@RequestBody CreateUserRequest request) {
        // 错误：没有验证输入参数
        User user = userService.create(request);
        return ResponseEntity.ok(user);
    }
    
    @GetMapping("/users/{id}")
    public ResponseEntity<User> getUser(@PathVariable String id) {
        // 错误：没有验证ID格式
        User user = userService.findById(Long.parseLong(id));
        return ResponseEntity.ok(user);
    }
}
```

- **正确实例**:
```java
// 正确示例 - 完整的输入验证
@RestController
@RequestMapping("/api/v1/users")
@Validated
public class UserController {
    
    @PostMapping
    public ResponseEntity<ApiResponse<User>> createUser(
            @Valid @RequestBody CreateUserRequest request) {
        User user = userService.create(request);
        return ResponseEntity.status(HttpStatus.CREATED)
            .body(ApiResponse.success(user));
    }
    
    @GetMapping("/{id}")
    public ResponseEntity<ApiResponse<User>> getUser(
            @PathVariable @Positive(message = "用户ID必须为正数") Long id) {
        User user = userService.findById(id);
        return ResponseEntity.ok(ApiResponse.success(user));
    }
}

// 完整的请求DTO验证
public class CreateUserRequest {
    
    @NotBlank(message = "用户名不能为空")
    @Size(min = 3, max = 20, message = "用户名长度必须在3-20个字符之间")
    @Pattern(regexp = "^[a-zA-Z0-9_]+$", message = "用户名只能包含字母、数字和下划线")
    private String username;
    
    @NotBlank(message = "邮箱不能为空")
    @Email(message = "邮箱格式不正确")
    private String email;
    
    @NotBlank(message = "密码不能为空")
    @Size(min = 8, max = 20, message = "密码长度必须在8-20个字符之间")
    private String password;
    
    @NotNull(message = "年龄不能为空")
    @Min(value = 18, message = "年龄不能小于18岁")
    @Max(value = 120, message = "年龄不能大于120岁")
    private Integer age;
    
    // getters and setters
}
```

#### 9.6.2 业务规则验证
- **检查方法**: 检查是否实现了复杂的业务规则验证
- **检查标准**: 使用分组验证、自定义验证器实现业务逻辑验证
- **正确实例**:
```java
// 自定义验证注解
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
    private UserRepository userRepository;
    
    @Override
    public boolean isValid(String email, ConstraintValidatorContext context) {
        if (email == null) {
            return true;
        }
        return !userRepository.existsByEmail(email);
    }
}
```

### 9.7 响应格式 (Minor)

#### 9.7.1 统一响应结构
- **检查方法**: 检查所有API是否使用统一的响应格式
- **检查标准**: 定义标准响应结构，包含状态、数据、错误信息等字段
- **不正确实例**:
```java
// 错误示例 - 响应格式不统一
@RestController
public class ProductController {
    
    @GetMapping("/products/{id}")
    public ResponseEntity<Product> getProduct(@PathVariable Long id) {
        // 错误：直接返回实体对象
        Product product = productService.findById(id);
        return ResponseEntity.ok(product);
    }
    
    @PostMapping("/products")
    public ResponseEntity<String> createProduct(@RequestBody Product product) {
        // 错误：返回字符串，格式不统一
        productService.create(product);
        return ResponseEntity.ok("Product created successfully");
    }
}
```

- **正确实例**:
```java
// 正确示例 - 统一的响应格式
@RestController
@RequestMapping("/api/v1/products")
public class ProductController {
    
    @GetMapping("/{id}")
    public ResponseEntity<ApiResponse<Product>> getProduct(@PathVariable Long id) {
        Product product = productService.findById(id);
        return ResponseEntity.ok(ApiResponse.success(product));
    }
    
    @PostMapping
    public ResponseEntity<ApiResponse<Product>> createProduct(@Valid @RequestBody CreateProductRequest request) {
        Product product = productService.create(request);
        return ResponseEntity.status(HttpStatus.CREATED)
            .body(ApiResponse.success(product));
    }
}

// 统一响应结构
@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class ApiResponse<T> {
    private boolean success;
    private T data;
    private ErrorInfo error;
    private String message;
    private long timestamp;
    private String requestId;
    
    public static <T> ApiResponse<T> success(T data) {
        return ApiResponse.<T>builder()
            .success(true)
            .data(data)
            .timestamp(System.currentTimeMillis())
            .build();
    }
    
    public static <T> ApiResponse<T> error(String code, String message) {
        return ApiResponse.<T>builder()
            .success(false)
            .error(ErrorInfo.builder()
                .code(code)
                .message(message)
                .build())
            .timestamp(System.currentTimeMillis())
            .build();
    }
}
```

#### 9.7.2 分页响应格式
- **检查方法**: 检查分页接口是否使用统一的分页响应格式
- **检查标准**: 包含分页元数据，如总数、页码、页大小等信息
- **正确实例**:
```java
// 完整的分页结果类
@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class PageResult<T> {
    private List<T> content;
    private PageInfo pageInfo;
    
    @Data
    @Builder
    @NoArgsConstructor
    @AllArgsConstructor
    public static class PageInfo {
        private int page;
        private int size;
        private long totalElements;
        private int totalPages;
        private boolean first;
        private boolean last;
        private boolean hasNext;
        private boolean hasPrevious;
    }
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

### 10.3 测试覆盖率 (Minor)

#### 10.3.1 代码覆盖率要求
- **检查方法**: 使用JaCoCo等工具检查代码覆盖率
- **检查标准**: 单元测试覆盖率≥80%，集成测试覆盖率≥70%，关键业务逻辑覆盖率≥90%
- **不正确实例**:
```xml
<!-- 错误示例 - 覆盖率要求过低 -->
<plugin>
    <groupId>org.jacoco</groupId>
    <artifactId>jacoco-maven-plugin</artifactId>
    <configuration>
        <rules>
            <rule>
                <element>BUNDLE</element>
                <limits>
                    <limit>
                        <counter>LINE</counter>
                        <value>COVEREDRATIO</value>
                        <minimum>0.50</minimum> <!-- 覆盖率要求过低 -->
                    </limit>
                </limits>
            </rule>
        </rules>
    </configuration>
</plugin>
```

- **正确实例**:
```xml
<!-- 正确示例 - 合理的覆盖率配置 -->
<plugin>
    <groupId>org.jacoco</groupId>
    <artifactId>jacoco-maven-plugin</artifactId>
    <version>0.8.7</version>
    <executions>
        <execution>
            <goals>
                <goal>prepare-agent</goal>
            </goals>
        </execution>
        <execution>
            <id>report</id>
            <phase>test</phase>
            <goals>
                <goal>report</goal>
            </goals>
        </execution>
        <execution>
            <id>check</id>
            <goals>
                <goal>check</goal>
            </goals>
            <configuration>
                <rules>
                    <rule>
                        <element>BUNDLE</element>
                        <limits>
                            <limit>
                                <counter>LINE</counter>
                                <value>COVEREDRATIO</value>
                                <minimum>0.80</minimum>
                            </limit>
                            <limit>
                                <counter>BRANCH</counter>
                                <value>COVEREDRATIO</value>
                                <minimum>0.75</minimum>
                            </limit>
                        </limits>
                    </rule>
                    <rule>
                        <element>CLASS</element>
                        <excludes>
                            <exclude>*.*Application</exclude>
                            <exclude>*.config.*</exclude>
                            <exclude>*.dto.*</exclude>
                            <exclude>*.entity.*</exclude>
                        </excludes>
                        <limits>
                            <limit>
                                <counter>LINE</counter>
                                <value>COVEREDRATIO</value>
                                <minimum>0.85</minimum>
                            </limit>
                        </limits>
                    </rule>
                </rules>
            </configuration>
        </execution>
    </executions>
</plugin>
```

#### 10.3.2 覆盖率报告分析
- **检查方法**: 定期分析覆盖率报告，识别未覆盖的关键代码
- **检查标准**: 关注分支覆盖率、异常处理覆盖率、边界条件覆盖率
- **不正确实例**:
```java
// 错误示例 - 缺少边界条件测试
@Test
void testCalculateDiscount() {
    // 只测试正常情况，缺少边界条件
    BigDecimal result = discountService.calculateDiscount(
        new BigDecimal("100"), CustomerType.VIP);
    assertEquals(new BigDecimal("90.00"), result);
}
```

- **正确实例**:
```java
// 正确示例 - 全面的边界条件测试
@ParameterizedTest
@CsvSource({
    "0, REGULAR, 0.00",           // 边界：金额为0
    "0.01, REGULAR, 0.01",       // 边界：最小金额
    "99.99, REGULAR, 99.99",     // 边界：折扣临界点
    "100.00, REGULAR, 95.00",    // 边界：折扣起始点
    "1000.00, VIP, 850.00",      // 正常：VIP折扣
    "10000.00, PREMIUM, 8000.00" // 边界：最大折扣
})
void testCalculateDiscount_AllScenarios(String amount, CustomerType type, String expected) {
    BigDecimal result = discountService.calculateDiscount(
        new BigDecimal(amount), type);
    assertEquals(new BigDecimal(expected), result);
}

@Test
void testCalculateDiscount_NullAmount_ThrowsException() {
    // 测试异常情况
    assertThrows(IllegalArgumentException.class, () -> 
        discountService.calculateDiscount(null, CustomerType.REGULAR));
}

@Test
void testCalculateDiscount_NegativeAmount_ThrowsException() {
    // 测试负数边界
    assertThrows(IllegalArgumentException.class, () -> 
        discountService.calculateDiscount(new BigDecimal("-1"), CustomerType.REGULAR));
}

@Test
void testCalculateDiscount_ExtremelyLargeAmount_HandledCorrectly() {
    // 测试极大值边界
    BigDecimal largeAmount = new BigDecimal("999999999.99");
    BigDecimal result = discountService.calculateDiscount(largeAmount, CustomerType.VIP);
    
    assertThat(result).isPositive();
    assertThat(result).isLessThanOrEqualTo(largeAmount);
}
```

### 10.4 测试数据管理 (Minor)

#### 10.4.1 测试数据隔离
- **检查方法**: 确保测试之间数据隔离，避免测试相互影响
- **检查标准**: 每个测试使用独立的数据集，测试后清理数据
- **不正确实例**:
```java
// 错误示例 - 测试数据污染
@SpringBootTest
class UserServiceTest {
    
    @Autowired
    private UserService userService;
    
    @Test
    void testCreateUser() {
        User user = new User("john@example.com", "John", 25);
        userService.createUser(user);
        // 没有清理数据，影响后续测试
    }
    
    @Test
    void testFindUserByEmail() {
        // 依赖前一个测试的数据，测试不独立
        User user = userService.findByEmail("john@example.com");
        assertNotNull(user);
    }
}
```

- **正确实例**:
```java
// 正确示例 - 测试数据隔离
@SpringBootTest
@Transactional
@Rollback
class UserServiceTest {
    
    @Autowired
    private UserService userService;
    
    @Autowired
    private UserRepository userRepository;
    
    @Autowired
    private TestEntityManager entityManager;
    
    @BeforeEach
    void setUp() {
        // 每个测试前清理数据
        userRepository.deleteAll();
        entityManager.flush();
        entityManager.clear();
    }
    
    @Test
    void testCreateUser() {
        // 使用独立的测试数据
        User user = createTestUser("john@example.com", "John", 25);
        User savedUser = userService.createUser(user);
        
        assertThat(savedUser.getId()).isNotNull();
        assertThat(savedUser.getEmail()).isEqualTo("john@example.com");
    }
    
    @Test
    void testFindUserByEmail() {
        // 创建测试专用数据
        User testUser = createTestUser("jane@example.com", "Jane", 30);
        userService.createUser(testUser);
        
        User foundUser = userService.findByEmail("jane@example.com");
        assertThat(foundUser).isNotNull();
        assertThat(foundUser.getName()).isEqualTo("Jane");
    }
    
    @Test
    void testFindUserByEmail_NotFound() {
        // 测试不存在的数据
        User user = userService.findByEmail("nonexistent@example.com");
        assertThat(user).isNull();
    }
    
    private User createTestUser(String email, String name, int age) {
        return User.builder()
                .email(email)
                .name(name)
                .age(age)
                .createdAt(LocalDateTime.now())
                .build();
    }
}
```

#### 10.4.2 测试数据构建器
- **检查方法**: 使用Builder模式或工厂方法创建测试数据
- **检查标准**: 测试数据创建简洁、可读性强、易于维护
- **不正确实例**:
```java
// 错误示例 - 硬编码测试数据
@Test
void testOrderProcessing() {
    User user = new User();
    user.setId(1L);
    user.setEmail("test@example.com");
    user.setName("Test User");
    user.setAge(25);
    user.setCreatedAt(LocalDateTime.now());
    
    Product product = new Product();
    product.setId(1L);
    product.setName("Test Product");
    product.setPrice(new BigDecimal("99.99"));
    product.setStock(10);
    
    Order order = new Order();
    order.setUser(user);
    order.addItem(product, 2);
    // 大量重复的数据设置代码
}
```

- **正确实例**:
```java
// 正确示例 - 使用测试数据构建器
public class TestDataBuilder {
    
    public static UserBuilder aUser() {
        return new UserBuilder();
    }
    
    public static ProductBuilder aProduct() {
        return new ProductBuilder();
    }
    
    public static OrderBuilder anOrder() {
        return new OrderBuilder();
    }
    
    public static class UserBuilder {
        private String email = "test@example.com";
        private String name = "Test User";
        private int age = 25;
        private LocalDateTime createdAt = LocalDateTime.now();
        
        public UserBuilder withEmail(String email) {
            this.email = email;
            return this;
        }
        
        public UserBuilder withName(String name) {
            this.name = name;
            return this;
        }
        
        public UserBuilder withAge(int age) {
            this.age = age;
            return this;
        }
        
        public UserBuilder createdAt(LocalDateTime createdAt) {
            this.createdAt = createdAt;
            return this;
        }
        
        public User build() {
            return User.builder()
                    .email(email)
                    .name(name)
                    .age(age)
                    .createdAt(createdAt)
                    .build();
        }
    }
    
    public static class ProductBuilder {
        private String name = "Test Product";
        private BigDecimal price = new BigDecimal("99.99");
        private int stock = 10;
        private String category = "Electronics";
        
        public ProductBuilder withName(String name) {
            this.name = name;
            return this;
        }
        
        public ProductBuilder withPrice(BigDecimal price) {
            this.price = price;
            return this;
        }
        
        public ProductBuilder withStock(int stock) {
            this.stock = stock;
            return this;
        }
        
        public ProductBuilder inCategory(String category) {
            this.category = category;
            return this;
        }
        
        public Product build() {
            return Product.builder()
                    .name(name)
                    .price(price)
                    .stock(stock)
                    .category(category)
                    .build();
        }
    }
    
    public static class OrderBuilder {
        private User user = aUser().build();
        private List<OrderItem> items = new ArrayList<>();
        private OrderStatus status = OrderStatus.PENDING;
        private LocalDateTime orderDate = LocalDateTime.now();
        
        public OrderBuilder forUser(User user) {
            this.user = user;
            return this;
        }
        
        public OrderBuilder withItem(Product product, int quantity) {
            items.add(new OrderItem(product, quantity));
            return this;
        }
        
        public OrderBuilder withStatus(OrderStatus status) {
            this.status = status;
            return this;
        }
        
        public OrderBuilder orderedAt(LocalDateTime orderDate) {
            this.orderDate = orderDate;
            return this;
        }
        
        public Order build() {
            Order order = Order.builder()
                    .user(user)
                    .status(status)
                    .orderDate(orderDate)
                    .build();
            items.forEach(item -> order.addItem(item.getProduct(), item.getQuantity()));
            return order;
        }
    }
}

// 使用测试数据构建器的测试
@Test
void testOrderProcessing() {
    User user = aUser()
            .withEmail("customer@example.com")
            .withName("John Doe")
            .build();
    
    Product product = aProduct()
            .withName("Laptop")
            .withPrice(new BigDecimal("1299.99"))
            .withStock(5)
            .build();
    
    Order order = anOrder()
            .forUser(user)
            .withItem(product, 1)
            .withStatus(OrderStatus.PENDING)
            .build();
    
    Order processedOrder = orderService.processOrder(order);
    
    assertThat(processedOrder.getStatus()).isEqualTo(OrderStatus.CONFIRMED);
    assertThat(processedOrder.getTotalAmount()).isEqualTo(new BigDecimal("1299.99"));
}
```

### 10.5 性能测试 (Major)

#### 10.5.1 负载测试
- **检查方法**: 使用JMeter、Gatling或自定义负载测试
- **检查标准**: 验证系统在预期负载下的性能表现
- **不正确实例**:
```java
// 错误示例 - 简单的性能测试
@Test
void testPerformance() {
    long start = System.currentTimeMillis();
    userService.createUser(new User("test@example.com", "Test", 25));
    long duration = System.currentTimeMillis() - start;
    assertTrue(duration < 1000); // 过于简单
}
```

- **正确实例**:
```java
// 正确示例 - 完整的负载测试
@SpringBootTest
@Testcontainers
class UserServiceLoadTest {
    
    @Container
    static MySQLContainer<?> mysql = new MySQLContainer<>("mysql:8.0")
            .withDatabaseName("testdb")
            .withUsername("test")
            .withPassword("test");
    
    @Autowired
    private UserService userService;
    
    @Test
    void loadTest_CreateUsers_HandlesExpectedLoad() throws InterruptedException {
        // 负载测试参数
        int threadCount = 20;
        int requestsPerThread = 100;
        int totalRequests = threadCount * requestsPerThread;
        
        CountDownLatch startLatch = new CountDownLatch(1);
        CountDownLatch endLatch = new CountDownLatch(threadCount);
        
        AtomicInteger successCount = new AtomicInteger(0);
        AtomicInteger errorCount = new AtomicInteger(0);
        List<Long> responseTimes = Collections.synchronizedList(new ArrayList<>());
        
        ExecutorService executor = Executors.newFixedThreadPool(threadCount);
        
        // 启动负载测试线程
        for (int t = 0; t < threadCount; t++) {
            final int threadId = t;
            executor.submit(() -> {
                try {
                    startLatch.await(); // 等待统一开始
                    
                    for (int i = 0; i < requestsPerThread; i++) {
                        long start = System.nanoTime();
                        try {
                            User user = User.builder()
                                    .email(String.format("load_test_%d_%d@example.com", threadId, i))
                                    .name(String.format("Load Test User %d_%d", threadId, i))
                                    .age(25 + (i % 50))
                                    .build();
                            
                            userService.createUser(user);
                            successCount.incrementAndGet();
                            
                        } catch (Exception e) {
                            errorCount.incrementAndGet();
                        } finally {
                            long responseTime = System.nanoTime() - start;
                            responseTimes.add(responseTime);
                        }
                    }
                } catch (InterruptedException e) {
                    Thread.currentThread().interrupt();
                } finally {
                    endLatch.countDown();
                }
            });
        }
        
        // 开始测试
        long testStart = System.currentTimeMillis();
        startLatch.countDown();
        
        // 等待测试完成
        boolean completed = endLatch.await(60, TimeUnit.SECONDS);
        long testDuration = System.currentTimeMillis() - testStart;
        
        executor.shutdown();
        
        // 验证测试结果
        assertThat(completed).isTrue();
        
        // 计算性能指标
        double throughput = (double) successCount.get() / (testDuration / 1000.0);
        double errorRate = (double) errorCount.get() / totalRequests;
        
        LongSummaryStatistics stats = responseTimes.stream()
                .mapToLong(Long::longValue)
                .summaryStatistics();
        
        double avgResponseTimeMs = stats.getAverage() / 1_000_000.0;
        double maxResponseTimeMs = stats.getMax() / 1_000_000.0;
        
        // 性能断言
        assertThat(throughput).isGreaterThan(50.0); // 吞吐量 > 50 TPS
        assertThat(errorRate).isLessThan(0.01);     // 错误率 < 1%
        assertThat(avgResponseTimeMs).isLessThan(100.0); // 平均响应时间 < 100ms
        assertThat(maxResponseTimeMs).isLessThan(500.0); // 最大响应时间 < 500ms
        
        // 输出性能报告
        System.out.printf("负载测试结果:\n");
        System.out.printf("总请求数: %d\n", totalRequests);
        System.out.printf("成功请求数: %d\n", successCount.get());
        System.out.printf("失败请求数: %d\n", errorCount.get());
        System.out.printf("吞吐量: %.2f TPS\n", throughput);
        System.out.printf("错误率: %.2f%%\n", errorRate * 100);
        System.out.printf("平均响应时间: %.2f ms\n", avgResponseTimeMs);
        System.out.printf("最大响应时间: %.2f ms\n", maxResponseTimeMs);
    }
    
    @Test
    void stressTest_FindUsers_HandlesHighConcurrency() throws InterruptedException {
        // 先创建测试数据
        List<User> testUsers = new ArrayList<>();
        for (int i = 0; i < 1000; i++) {
            User user = User.builder()
                    .email(String.format("stress_test_%d@example.com", i))
                    .name(String.format("Stress Test User %d", i))
                    .age(20 + (i % 60))
                    .build();
            testUsers.add(userService.createUser(user));
        }
        
        // 压力测试参数
        int threadCount = 50;
        int requestsPerThread = 200;
        
        CountDownLatch startLatch = new CountDownLatch(1);
        CountDownLatch endLatch = new CountDownLatch(threadCount);
        
        AtomicInteger successCount = new AtomicInteger(0);
        AtomicInteger errorCount = new AtomicInteger(0);
        
        ExecutorService executor = Executors.newFixedThreadPool(threadCount);
        Random random = new Random();
        
        for (int t = 0; t < threadCount; t++) {
            executor.submit(() -> {
                try {
                    startLatch.await();
                    
                    for (int i = 0; i < requestsPerThread; i++) {
                        try {
                            // 随机查询用户
                            User randomUser = testUsers.get(random.nextInt(testUsers.size()));
                            User foundUser = userService.findByEmail(randomUser.getEmail());
                            
                            if (foundUser != null) {
                                successCount.incrementAndGet();
                            } else {
                                errorCount.incrementAndGet();
                            }
                        } catch (Exception e) {
                            errorCount.incrementAndGet();
                        }
                    }
                } catch (InterruptedException e) {
                    Thread.currentThread().interrupt();
                } finally {
                    endLatch.countDown();
                }
            });
        }
        
        long testStart = System.currentTimeMillis();
        startLatch.countDown();
        
        boolean completed = endLatch.await(120, TimeUnit.SECONDS);
        long testDuration = System.currentTimeMillis() - testStart;
        
        executor.shutdown();
        
        assertThat(completed).isTrue();
        
        int totalRequests = threadCount * requestsPerThread;
        double throughput = (double) successCount.get() / (testDuration / 1000.0);
        double errorRate = (double) errorCount.get() / totalRequests;
        
        // 压力测试断言
        assertThat(throughput).isGreaterThan(200.0); // 查询吞吐量 > 200 TPS
        assertThat(errorRate).isLessThan(0.005);     // 错误率 < 0.5%
        
        System.out.printf("压力测试结果:\n");
        System.out.printf("查询吞吐量: %.2f TPS\n", throughput);
        System.out.printf("错误率: %.2f%%\n", errorRate * 100);
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

#### 10.5.2 内存和资源监控
- **检查方法**: 监控测试过程中的内存使用、CPU使用率、数据库连接等
- **检查标准**: 资源使用在合理范围内，无内存泄漏
- **正确实例**:
```java
@Test
void memoryLeakTest_RepeatedOperations_NoMemoryLeak() {
    MemoryMXBean memoryBean = ManagementFactory.getMemoryMXBean();
    Runtime runtime = Runtime.getRuntime();
    
    // 记录初始内存状态
    runtime.gc();
    long initialMemory = memoryBean.getHeapMemoryUsage().getUsed();
    
    // 执行重复操作
    for (int cycle = 0; cycle < 10; cycle++) {
        List<User> users = new ArrayList<>();
        
        // 创建大量对象
        for (int i = 0; i < 1000; i++) {
            User user = User.builder()
                    .email(String.format("memory_test_%d_%d@example.com", cycle, i))
                    .name(String.format("Memory Test User %d_%d", cycle, i))
                    .age(25)
                    .build();
            users.add(userService.createUser(user));
        }
        
        // 处理数据
        users.forEach(user -> {
            userService.findByEmail(user.getEmail());
            userService.updateUser(user.getId(), user);
        });
        
        // 清理数据
        users.forEach(user -> userService.deleteUser(user.getId()));
        
        // 强制垃圾回收
        runtime.gc();
        
        // 检查内存使用
        long currentMemory = memoryBean.getHeapMemoryUsage().getUsed();
        long memoryIncrease = currentMemory - initialMemory;
        
        // 内存增长不应超过50MB
        assertThat(memoryIncrease).isLessThan(50 * 1024 * 1024);
        
        System.out.printf("Cycle %d: Memory increase: %d MB\n", 
                cycle, memoryIncrease / (1024 * 1024));
    }
}
```

## 11. 部署和运维检查

### 11.1 容器化 (Major)

#### 11.1.1 Dockerfile最佳实践
- **检查方法**: 检查Dockerfile的构建效率和安全性
- **检查标准**: 使用多阶段构建、选择合适的基础镜像、最小化镜像层数
- **不正确实例**:
```dockerfile
# 错误示例 - 低效的Dockerfile
FROM openjdk:11
COPY . /app
WORKDIR /app
RUN apt-get update && apt-get install -y maven
RUN mvn clean package
EXPOSE 8080
CMD ["java", "-jar", "target/app.jar"]

# 正确示例 - 优化的多阶段构建
# 构建阶段
FROM maven:3.8.4-openjdk-11-slim AS builder
WORKDIR /app

# 先复制依赖文件，利用Docker层缓存
COPY pom.xml .
RUN mvn dependency:go-offline -B

# 复制源码并构建
COPY src ./src
RUN mvn clean package -DskipTests

# 运行阶段
FROM openjdk:11-jre-slim

# 创建非root用户
RUN groupadd -r appuser && useradd -r -g appuser appuser

# 安装必要的工具
RUN apt-get update && apt-get install -y \
    curl \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# 复制构建产物
COPY --from=builder /app/target/*.jar app.jar

# 设置JVM参数
ENV JAVA_OPTS="-Xms512m -Xmx1024m -XX:+UseG1GC"

# 切换到非root用户
USER appuser

# 健康检查
HEALTHCHECK --interval=30s --timeout=3s --start-period=60s --retries=3 \
    CMD curl -f http://localhost:8080/actuator/health || exit 1

EXPOSE 8080
CMD ["sh", "-c", "java $JAVA_OPTS -jar app.jar"]
```

#### 11.1.4 健康检查状态更新
- **检查方法**: 检查容器健康检查的实现和状态更新机制
- **检查标准**: 实现准确的健康检查，及时更新容器状态
- **不正确实例**:
```dockerfile
# 错误示例 - 简单的健康检查
FROM openjdk:11-jre-slim
COPY app.jar /app.jar
EXPOSE 8080

# 简单的端口检查，不能反映应用真实状态
HEALTHCHECK --interval=30s --timeout=3s --retries=3 \
    CMD curl -f http://localhost:8080 || exit 1

CMD ["java", "-jar", "/app.jar"]

# 正确示例 - 完整的健康检查机制
FROM openjdk:11-jre-slim

# 安装健康检查工具
RUN apt-get update && apt-get install -y \
    curl \
    jq \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY app.jar app.jar
COPY healthcheck.sh healthcheck.sh
RUN chmod +x healthcheck.sh

# 设置健康检查环境变量
ENV HEALTH_CHECK_URL=http://localhost:8080/actuator/health
ENV HEALTH_CHECK_TIMEOUT=10
ENV HEALTH_CHECK_RETRIES=3

# 详细的健康检查脚本
HEALTHCHECK --interval=30s --timeout=10s --start-period=60s --retries=3 \
    CMD ./healthcheck.sh

EXPOSE 8080
CMD ["java", "-jar", "app.jar"]
```

```bash
#!/bin/bash
# healthcheck.sh - 详细的健康检查脚本

set -e

# 配置参数
HEALTH_URL=${HEALTH_CHECK_URL:-"http://localhost:8080/actuator/health"}
TIMEOUT=${HEALTH_CHECK_TIMEOUT:-10}
RETRIES=${HEALTH_CHECK_RETRIES:-3}
LOG_FILE="/tmp/healthcheck.log"

# 日志函数
log() {
    echo "$(date '+%Y-%m-%d %H:%M:%S') [HEALTHCHECK] $1" | tee -a "$LOG_FILE"
}

# 检查应用端口
check_port() {
    local port=8080
    if ! nc -z localhost $port; then
        log "ERROR: Port $port is not accessible"
        return 1
    fi
    log "INFO: Port $port is accessible"
    return 0
}

# 检查健康端点
check_health_endpoint() {
    local response
    local http_code
    local health_status
    
    # 发送健康检查请求
    response=$(curl -s -w "%{http_code}" --max-time "$TIMEOUT" "$HEALTH_URL" 2>/dev/null || echo "000")
    http_code=${response: -3}
    
    if [ "$http_code" != "200" ]; then
        log "ERROR: Health endpoint returned HTTP $http_code"
        return 1
    fi
    
    # 解析健康状态
    health_status=$(echo "${response%???}" | jq -r '.status // "UNKNOWN"' 2>/dev/null || echo "UNKNOWN")
    
    if [ "$health_status" != "UP" ]; then
        log "ERROR: Application health status is $health_status"
        log "INFO: Health response: ${response%???}"
        return 1
    fi
    
    log "INFO: Application health status is UP"
    return 0
}

# 检查关键依赖
check_dependencies() {
    local deps_url="${HEALTH_URL}/components"
    local response
    local http_code
    
    response=$(curl -s -w "%{http_code}" --max-time "$TIMEOUT" "$deps_url" 2>/dev/null || echo "000")
    http_code=${response: -3}
    
    if [ "$http_code" = "200" ]; then
        # 检查数据库状态
        local db_status=$(echo "${response%???}" | jq -r '.components.db.status // "UNKNOWN"' 2>/dev/null || echo "UNKNOWN")
        if [ "$db_status" != "UP" ]; then
            log "WARNING: Database health status is $db_status"
        fi
        
        # 检查Redis状态
        local redis_status=$(echo "${response%???}" | jq -r '.components.redis.status // "UNKNOWN"' 2>/dev/null || echo "UNKNOWN")
        if [ "$redis_status" != "UP" ]; then
            log "WARNING: Redis health status is $redis_status"
        fi
    fi
}

# 主健康检查逻辑
main() {
    log "Starting health check..."
    
    # 1. 检查端口
    if ! check_port; then
        exit 1
    fi
    
    # 2. 检查健康端点（带重试）
    local attempt=1
    while [ $attempt -le $RETRIES ]; do
        if check_health_endpoint; then
            break
        fi
        
        if [ $attempt -eq $RETRIES ]; then
            log "ERROR: Health check failed after $RETRIES attempts"
            exit 1
        fi
        
        log "WARNING: Health check attempt $attempt failed, retrying..."
        sleep 2
        attempt=$((attempt + 1))
    done
    
    # 3. 检查依赖状态（非阻塞）
    check_dependencies
    
    log "Health check passed"
    exit 0
}

# 执行健康检查
main "$@"
```

#### 11.1.2 镜像安全扫描
- **检查方法**: 使用安全扫描工具检查镜像漏洞
- **检查标准**: 定期扫描基础镜像和应用镜像的安全漏洞
- **不正确实例**:
```dockerfile
# 错误示例 - 使用不安全的基础镜像
FROM ubuntu:latest
RUN apt-get update && apt-get install -y openjdk-11-jdk
# 使用latest标签，不确定版本
# 安装了完整的JDK而不是JRE

# 正确示例 - 安全的基础镜像配置
FROM openjdk:11.0.16-jre-slim-bullseye

# 更新系统包并清理缓存
RUN apt-get update && apt-get upgrade -y \
    && apt-get install -y --no-install-recommends \
        curl \
        ca-certificates \
    && rm -rf /var/lib/apt/lists/* \
    && apt-get clean

# 创建应用目录和用户
RUN groupadd -r appgroup && useradd -r -g appgroup appuser
RUN mkdir -p /app && chown appuser:appgroup /app

# 设置安全的文件权限
USER appuser
WORKDIR /app

# 复制应用文件
COPY --chown=appuser:appgroup target/*.jar app.jar

# 镜像安全扫描命令示例:
# docker run --rm -v /var/run/docker.sock:/var/run/docker.sock \
#   aquasec/trivy image myapp:latest
# 
# 或使用Snyk:
# snyk container test myapp:latest
```

#### 11.1.3 多阶段构建优化
- **检查方法**: 检查是否使用多阶段构建减少镜像大小
- **检查标准**: 分离构建环境和运行环境，只包含运行时必需的文件
- **不正确实例**:
```dockerfile
# 错误示例 - 单阶段构建，镜像过大
FROM maven:3.8.4-openjdk-11
WORKDIR /app
COPY . .
RUN mvn clean package
EXPOSE 8080
CMD ["java", "-jar", "target/app.jar"]
# 包含了Maven、源码等不必要的文件

# 正确示例 - 多阶段构建优化
# 第一阶段：构建
FROM maven:3.8.4-openjdk-11-slim AS build
WORKDIR /app

# 复制依赖文件
COPY pom.xml .
COPY .mvn .mvn
COPY mvnw .

# 下载依赖
RUN ./mvnw dependency:go-offline -B

# 复制源码并构建
COPY src ./src
RUN ./mvnw clean package -DskipTests

# 第二阶段：运行时
FROM openjdk:11-jre-slim AS runtime

# 创建非特权用户
RUN addgroup --system appgroup && adduser --system --group appuser

# 创建应用目录
WORKDIR /app

# 只复制必要的JAR文件
COPY --from=build /app/target/*.jar app.jar

# 切换到非特权用户
USER appuser

# 设置JVM参数
ENV JAVA_OPTS="-Xms256m -Xmx512m -XX:+UseContainerSupport"

EXPOSE 8080

# 健康检查
HEALTHCHECK --interval=30s --timeout=10s --start-period=60s --retries=3 \
    CMD curl -f http://localhost:8080/actuator/health || exit 1

CMD ["sh", "-c", "java $JAVA_OPTS -jar app.jar"]

# 镜像大小对比:
# 单阶段构建: ~800MB
# 多阶段构建: ~200MB
```

### 11.2 优雅关闭 (Critical)

#### 11.2.1 应用停机处理
- **检查方法**: 检查应用对SIGTERM信号的处理
- **检查标准**: 正确处理关闭信号，完成正在处理的请求
- **不正确实例**:
```java
// 错误示例 - 没有优雅关闭机制
@SpringBootApplication
public class Application {
    public static void main(String[] args) {
        SpringApplication.run(Application.class, args);
        // 没有关闭钩子，强制终止可能导致数据丢失
    }
}

// 正确示例 - 实现优雅关闭
@SpringBootApplication
public class Application {
    
    private static final Logger logger = LoggerFactory.getLogger(Application.class);
    
    public static void main(String[] args) {
        SpringApplication app = new SpringApplication(Application.class);
        app.setRegisterShutdownHook(true);
        
        ConfigurableApplicationContext context = app.run(args);
        
        // 添加JVM关闭钩子
        Runtime.getRuntime().addShutdownHook(new Thread(() -> {
            logger.info("Received shutdown signal, starting graceful shutdown...");
            context.close();
            logger.info("Application shutdown completed.");
        }));
    }
}

// 配置优雅关闭
@Configuration
public class GracefulShutdownConfig {
    
    @Bean
    public GracefulShutdown gracefulShutdown() {
        return new GracefulShutdown();
    }
    
    @Bean
    public ConfigurableServletWebServerFactory webServerFactory(GracefulShutdown gracefulShutdown) {
        TomcatServletWebServerFactory factory = new TomcatServletWebServerFactory();
        factory.addConnectorCustomizers(gracefulShutdown);
        return factory;
    }
    
    private static class GracefulShutdown implements TomcatConnectorCustomizer, ApplicationListener<ContextClosedEvent> {
        
        private static final Logger logger = LoggerFactory.getLogger(GracefulShutdown.class);
        private volatile Connector connector;
        
        @Override
        public void customize(Connector connector) {
            this.connector = connector;
        }
        
        @Override
        public void onApplicationEvent(ContextClosedEvent event) {
            if (connector != null) {
                connector.pause();
                Executor executor = connector.getProtocolHandler().getExecutor();
                if (executor instanceof ThreadPoolExecutor) {
                    try {
                        ThreadPoolExecutor threadPoolExecutor = (ThreadPoolExecutor) executor;
                        threadPoolExecutor.shutdown();
                        if (!threadPoolExecutor.awaitTermination(30, TimeUnit.SECONDS)) {
                            logger.warn("Tomcat thread pool did not shut down gracefully within 30 seconds");
                            threadPoolExecutor.shutdownNow();
                        }
                    } catch (InterruptedException ex) {
                        Thread.currentThread().interrupt();
                    }
                }
            }
        }
    }
}

# application.yml配置
server:
  shutdown: graceful
  
spring:
  lifecycle:
    timeout-per-shutdown-phase: 30s
```

#### 11.2.2 资源清理
- **检查方法**: 检查应用关闭时的资源清理逻辑
- **检查标准**: 正确关闭数据库连接、缓存连接、文件句柄等资源
- **不正确实例**:
```java
// 错误示例 - 资源清理不完整
@Component
public class ResourceManager {
    
    private RedisTemplate<String, Object> redisTemplate;
    private DataSource dataSource;
    
    @PreDestroy
    public void cleanup() {
        // 只是简单的日志，没有实际清理资源
        logger.info("Cleaning up resources...");
    }
}

// 正确示例 - 完整的资源清理
@Component
public class ResourceManager {
    
    private static final Logger logger = LoggerFactory.getLogger(ResourceManager.class);
    
    @Autowired
    private RedisTemplate<String, Object> redisTemplate;
    
    @Autowired
    private DataSource dataSource;
    
    @Autowired
    private ScheduledExecutorService scheduler;
    
    @Autowired
    private CacheManager cacheManager;
    
    @PreDestroy
    public void cleanup() {
        logger.info("Starting resource cleanup...");
        
        // 1. 停止定时任务
        if (scheduler != null && !scheduler.isShutdown()) {
            logger.info("Shutting down scheduler...");
            scheduler.shutdown();
            try {
                if (!scheduler.awaitTermination(10, TimeUnit.SECONDS)) {
                    logger.warn("Scheduler did not terminate gracefully");
                    scheduler.shutdownNow();
                }
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
                scheduler.shutdownNow();
            }
        }
        
        // 2. 清理缓存
        if (cacheManager != null) {
            logger.info("Clearing caches...");
            cacheManager.getCacheNames().forEach(cacheName -> {
                Cache cache = cacheManager.getCache(cacheName);
                if (cache != null) {
                    cache.clear();
                }
            });
        }
        
        // 3. 关闭Redis连接
        if (redisTemplate != null) {
            logger.info("Closing Redis connections...");
            try {
                RedisConnectionFactory connectionFactory = redisTemplate.getConnectionFactory();
                if (connectionFactory instanceof DisposableBean) {
                    ((DisposableBean) connectionFactory).destroy();
                }
            } catch (Exception e) {
                logger.error("Error closing Redis connections", e);
            }
        }
        
        // 4. 关闭数据库连接池
        if (dataSource instanceof HikariDataSource) {
            logger.info("Closing database connection pool...");
            ((HikariDataSource) dataSource).close();
        }
        
        logger.info("Resource cleanup completed");
    }
}
```

#### 11.2.3 正在处理的请求完成
- **检查方法**: 检查是否等待正在处理的请求完成
- **检查标准**: 在关闭前等待活跃请求处理完成，避免数据丢失
- **不正确实例**:
```java
// 错误示例 - 立即关闭，不等待请求完成
@RestController
public class UserController {
    
    @PostMapping("/users")
    public ResponseEntity<User> createUser(@RequestBody User user) {
        // 长时间处理逻辑
        processUser(user);  // 可能需要10秒
        return ResponseEntity.ok(user);
    }
    
    // 没有优雅关闭机制，可能在处理过程中被强制终止
}

// 正确示例 - 等待请求完成的优雅关闭
@RestController
public class UserController {
    
    private static final Logger logger = LoggerFactory.getLogger(UserController.class);
    private final AtomicInteger activeRequests = new AtomicInteger(0);
    private volatile boolean shutdownRequested = false;
    
    @PostMapping("/users")
    public ResponseEntity<User> createUser(@RequestBody User user) {
        if (shutdownRequested) {
            return ResponseEntity.status(HttpStatus.SERVICE_UNAVAILABLE)
                    .body(null);
        }
        
        activeRequests.incrementAndGet();
        try {
            logger.info("Processing user creation, active requests: {}", activeRequests.get());
            processUser(user);
            return ResponseEntity.ok(user);
        } finally {
            activeRequests.decrementAndGet();
            logger.info("Completed user creation, active requests: {}", activeRequests.get());
        }
    }
    
    @EventListener
    public void handleShutdown(ContextClosedEvent event) {
        shutdownRequested = true;
        logger.info("Shutdown requested, waiting for {} active requests to complete", 
                   activeRequests.get());
        
        // 等待活跃请求完成
        long startTime = System.currentTimeMillis();
        while (activeRequests.get() > 0) {
            try {
                Thread.sleep(100);
                if (System.currentTimeMillis() - startTime > 30000) {
                    logger.warn("Timeout waiting for requests to complete, {} requests still active", 
                               activeRequests.get());
                    break;
                }
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
                break;
            }
        }
        
        logger.info("All requests completed or timeout reached");
    }
}

// 使用拦截器统一管理请求状态
@Component
public class GracefulShutdownInterceptor implements HandlerInterceptor {
    
    private final AtomicInteger activeRequests = new AtomicInteger(0);
    private volatile boolean shutdownRequested = false;
    
    @Override
    public boolean preHandle(HttpServletRequest request, HttpServletResponse response, 
                           Object handler) throws Exception {
        if (shutdownRequested) {
            response.setStatus(HttpStatus.SERVICE_UNAVAILABLE.value());
            response.getWriter().write("Service is shutting down");
            return false;
        }
        
        activeRequests.incrementAndGet();
        request.setAttribute("requestStartTime", System.currentTimeMillis());
        return true;
    }
    
    @Override
    public void afterCompletion(HttpServletRequest request, HttpServletResponse response, 
                              Object handler, Exception ex) throws Exception {
        activeRequests.decrementAndGet();
        Long startTime = (Long) request.getAttribute("requestStartTime");
        if (startTime != null) {
            long duration = System.currentTimeMillis() - startTime;
            logger.debug("Request completed in {}ms, active requests: {}", 
                        duration, activeRequests.get());
        }
    }
    
    @EventListener
    public void handleShutdown(ContextClosedEvent event) {
        shutdownRequested = true;
        
        logger.info("Shutdown requested, waiting for {} active requests to complete", 
                   activeRequests.get());
        
        // 等待活跃请求完成
        long startTime = System.currentTimeMillis();
        while (activeRequests.get() > 0) {
            try {
                Thread.sleep(100);
                if (System.currentTimeMillis() - startTime > 30000) {
                    logger.warn("Timeout waiting for requests to complete, {} requests still active", 
                               activeRequests.get());
                    break;
                }
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
                break;
            }
        }
        
        logger.info("All requests completed or timeout reached");
    }
}
```

### 11.3 资源限制 (Major)

#### 11.3.1 容器资源限制
- **检查方法**: 检查Docker容器的CPU和内存限制配置
- **检查标准**: 合理设置资源限制，防止资源耗尽影响其他服务
- **不正确实例**:
```yaml
# 错误示例 - 没有资源限制的Docker Compose
version: '3.8'
services:
  app:
    image: myapp:latest
    ports:
      - "8080:8080"
    # 没有设置资源限制，可能消耗所有系统资源
```

- **正确实例**:
```yaml
# 正确示例 - 合理的资源限制配置
version: '3.8'
services:
  app:
    image: myapp:latest
    ports:
      - "8080:8080"
    deploy:
      resources:
        limits:
          cpus: '1.0'        # 限制CPU使用
          memory: 1G         # 限制内存使用
        reservations:
          cpus: '0.5'        # 保留CPU资源
          memory: 512M       # 保留内存资源
    environment:
      - JAVA_OPTS=-Xms512m -Xmx768m -XX:+UseContainerSupport
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:8080/actuator/health"]
      interval: 30s
      timeout: 10s
      retries: 3
      start_period: 60s

# Kubernetes资源限制示例
apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp
spec:
  replicas: 3
  selector:
    matchLabels:
      app: myapp
  template:
    metadata:
      labels:
        app: myapp
    spec:
      containers:
      - name: myapp
        image: myapp:latest
        ports:
        - containerPort: 8080
        resources:
          requests:
            memory: "512Mi"
            cpu: "500m"
          limits:
            memory: "1Gi"
            cpu: "1000m"
        env:
        - name: JAVA_OPTS
          value: "-Xms512m -Xmx768m -XX:+UseContainerSupport -XX:MaxRAMPercentage=75.0"
        livenessProbe:
          httpGet:
            path: /actuator/health
            port: 8080
          initialDelaySeconds: 60
          periodSeconds: 30
          timeoutSeconds: 10
          failureThreshold: 3
        readinessProbe:
          httpGet:
            path: /actuator/health/readiness
            port: 8080
          initialDelaySeconds: 30
          periodSeconds: 10
          timeoutSeconds: 5
          failureThreshold: 3
```

#### 11.3.2 JVM内存配置
- **检查方法**: 检查JVM堆内存和非堆内存的配置
- **检查标准**: 根据容器资源限制合理配置JVM参数
- **不正确实例**:
```dockerfile
# 错误示例 - JVM内存配置不当
FROM openjdk:11-jre-slim
COPY app.jar /app.jar

# 固定的内存配置，不考虑容器限制
ENV JAVA_OPTS="-Xms2g -Xmx4g"

CMD ["sh", "-c", "java $JAVA_OPTS -jar /app.jar"]
```

- **正确实例**:
```dockerfile
# 正确示例 - 自适应的JVM内存配置
FROM openjdk:11-jre-slim

# 安装内存分析工具
RUN apt-get update && apt-get install -y \
    procps \
    && rm -rf /var/lib/apt/lists/*

COPY app.jar /app.jar
COPY jvm-config.sh /jvm-config.sh
RUN chmod +x /jvm-config.sh

# 动态JVM配置
ENV JAVA_OPTS_BASE="-XX:+UseContainerSupport -XX:+UseG1GC -XX:+PrintGCDetails -XX:+PrintGCTimeStamps"
ENV JAVA_OPTS_MEMORY="-XX:MaxRAMPercentage=75.0 -XX:InitialRAMPercentage=50.0"
ENV JAVA_OPTS_GC="-XX:MaxGCPauseMillis=200 -XX:G1HeapRegionSize=16m"
ENV JAVA_OPTS_MONITORING="-XX:+HeapDumpOnOutOfMemoryError -XX:HeapDumpPath=/tmp/heapdump.hprof"

CMD ["sh", "-c", "./jvm-config.sh && java $JAVA_OPTS -jar /app.jar"]
```

```bash
#!/bin/bash
# jvm-config.sh - 动态JVM配置脚本

# 获取容器内存限制
MEMORY_LIMIT=$(cat /sys/fs/cgroup/memory/memory.limit_in_bytes 2>/dev/null || echo "0")
if [ "$MEMORY_LIMIT" = "9223372036854775807" ] || [ "$MEMORY_LIMIT" = "0" ]; then
    # 如果没有内存限制，使用系统内存
    MEMORY_LIMIT=$(free -b | awk '/^Mem:/{print $2}')
fi

# 转换为MB
MEMORY_LIMIT_MB=$((MEMORY_LIMIT / 1024 / 1024))

echo "Container memory limit: ${MEMORY_LIMIT_MB}MB"

# 根据内存大小动态配置JVM参数
if [ $MEMORY_LIMIT_MB -le 512 ]; then
    # 小内存容器
    HEAP_SIZE="-Xms128m -Xmx384m"
    GC_OPTS="-XX:+UseSerialGC"
elif [ $MEMORY_LIMIT_MB -le 1024 ]; then
    # 中等内存容器
    HEAP_SIZE="-Xms256m -Xmx768m"
    GC_OPTS="-XX:+UseG1GC -XX:MaxGCPauseMillis=200"
else
    # 大内存容器
    HEAP_SIZE="-XX:InitialRAMPercentage=50.0 -XX:MaxRAMPercentage=75.0"
    GC_OPTS="-XX:+UseG1GC -XX:MaxGCPauseMillis=100 -XX:G1HeapRegionSize=16m"
fi

# 组合所有JVM参数
export JAVA_OPTS="$JAVA_OPTS_BASE $HEAP_SIZE $GC_OPTS $JAVA_OPTS_MONITORING"

echo "JVM Options: $JAVA_OPTS"

# 输出内存信息
echo "System Memory Info:"
free -h
echo "Container Limits:"
cat /sys/fs/cgroup/memory/memory.limit_in_bytes 2>/dev/null || echo "No memory limit"
cat /sys/fs/cgroup/cpu/cpu.cfs_quota_us 2>/dev/null || echo "No CPU quota"
```

#### 11.3.3 数据库连接池配置
- **检查方法**: 检查数据库连接池的大小和超时配置
- **检查标准**: 根据应用负载和数据库容量合理配置连接池
- **不正确实例**:
```yaml
# 错误示例 - 连接池配置不当
spring:
  datasource:
    hikari:
      maximum-pool-size: 100  # 连接池过大
      minimum-idle: 50        # 最小空闲连接过多
      connection-timeout: 60000  # 超时时间过长
      # 缺少其他重要配置
```

- **正确实例**:
```yaml
# 正确示例 - 合理的连接池配置
spring:
  datasource:
    hikari:
      # 连接池大小配置
      maximum-pool-size: 20           # 最大连接数
      minimum-idle: 5                 # 最小空闲连接
      
      # 超时配置
      connection-timeout: 20000       # 连接超时 20秒
      idle-timeout: 300000           # 空闲超时 5分钟
      max-lifetime: 1200000          # 连接最大生命周期 20分钟
      
      # 连接验证
      validation-timeout: 5000        # 验证超时 5秒
      connection-test-query: SELECT 1 # 连接测试查询
      
      # 连接池名称和JMX
      pool-name: HikariCP-Pool
      register-mbeans: true
      
      # 泄漏检测
      leak-detection-threshold: 60000 # 连接泄漏检测阈值 60秒
      
      # 数据源属性
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

# 监控配置
management:
  endpoints:
    web:
      exposure:
        include: health,info,metrics,hikaricp
  endpoint:
    health:
      show-details: always
  metrics:
    export:
      prometheus:
        enabled: true
```

```java
// 连接池监控配置
@Configuration
@EnableConfigurationProperties(DataSourceProperties.class)
public class DataSourceConfig {
    
    private static final Logger logger = LoggerFactory.getLogger(DataSourceConfig.class);
    
    @Bean
    @ConfigurationProperties("spring.datasource.hikari")
    public HikariConfig hikariConfig() {
        HikariConfig config = new HikariConfig();
        
        // 动态调整连接池大小
        int availableProcessors = Runtime.getRuntime().availableProcessors();
        int maxPoolSize = Math.max(10, availableProcessors * 2);
        config.setMaximumPoolSize(maxPoolSize);
        config.setMinimumIdle(Math.max(2, maxPoolSize / 4));
        
        logger.info("Configured HikariCP with max pool size: {}, min idle: {}", 
                   maxPoolSize, config.getMinimumIdle());
        
        return config;
    }
    
    @Bean
    public DataSource dataSource(HikariConfig hikariConfig) {
        return new HikariDataSource(hikariConfig);
    }
    
    // 连接池健康检查
    @Component
    public static class HikariHealthIndicator implements HealthIndicator {
        
        @Autowired
        private DataSource dataSource;
        
        @Override
        public Health health() {
            try {
                if (dataSource instanceof HikariDataSource) {
                    HikariDataSource hikariDataSource = (HikariDataSource) dataSource;
                    HikariPoolMXBean poolMXBean = hikariDataSource.getHikariPoolMXBean();
                    
                    int activeConnections = poolMXBean.getActiveConnections();
                    int idleConnections = poolMXBean.getIdleConnections();
                    int totalConnections = poolMXBean.getTotalConnections();
                    int threadsAwaitingConnection = poolMXBean.getThreadsAwaitingConnection();
                    
                    Health.Builder builder = Health.up()
                            .withDetail("active", activeConnections)
                            .withDetail("idle", idleConnections)
                            .withDetail("total", totalConnections)
                            .withDetail("awaiting", threadsAwaitingConnection)
                            .withDetail("max", hikariDataSource.getMaximumPoolSize());
                    
                    // 检查连接池状态
                    if (threadsAwaitingConnection > 0) {
                        builder.status("WARN")
                               .withDetail("warning", "Threads waiting for connections");
                    }
                    
                    if (activeConnections >= hikariDataSource.getMaximumPoolSize() * 0.9) {
                        builder.status("WARN")
                               .withDetail("warning", "Connection pool nearly exhausted");
                    }
                    
                    return builder.build();
                }
                
                return Health.up().build();
            } catch (Exception e) {
                return Health.down(e).build();
            }
        }
    }
    
    // 连接池监控指标
    @Component
    public static class HikariMetrics {
        
        private final MeterRegistry meterRegistry;
        private final HikariDataSource dataSource;
        
        public HikariMetrics(MeterRegistry meterRegistry, DataSource dataSource) {
            this.meterRegistry = meterRegistry;
            this.dataSource = (HikariDataSource) dataSource;
            registerMetrics();
        }
        
        private void registerMetrics() {
            HikariPoolMXBean poolMXBean = dataSource.getHikariPoolMXBean();
            
            Gauge.builder("hikari.connections.active")
                 .description("Active connections")
                 .register(meterRegistry, poolMXBean, HikariPoolMXBean::getActiveConnections);
                 
            Gauge.builder("hikari.connections.idle")
                 .description("Idle connections")
                 .register(meterRegistry, poolMXBean, HikariPoolMXBean::getIdleConnections);
                 
            Gauge.builder("hikari.connections.total")
                 .description("Total connections")
                 .register(meterRegistry, poolMXBean, HikariPoolMXBean::getTotalConnections);
                 
            Gauge.builder("hikari.connections.awaiting")
                 .description("Threads awaiting connections")
                 .register(meterRegistry, poolMXBean, HikariPoolMXBean::getThreadsAwaitingConnection);
        }
    }
}

// 连接池动态调整
@Component
public class ConnectionPoolManager {
    
    private static final Logger logger = LoggerFactory.getLogger(ConnectionPoolManager.class);
    
    @Autowired
    private HikariDataSource dataSource;
    
    @Scheduled(fixedRate = 60000) // 每分钟检查一次
    public void adjustPoolSize() {
        try {
            HikariPoolMXBean poolMXBean = dataSource.getHikariPoolMXBean();
            
            int activeConnections = poolMXBean.getActiveConnections();
            int totalConnections = poolMXBean.getTotalConnections();
            int maxPoolSize = dataSource.getMaximumPoolSize();
            
            // 如果活跃连接数超过总连接数的80%，考虑增加连接池大小
            if (activeConnections > totalConnections * 0.8 && maxPoolSize < 50) {
                int newMaxSize = Math.min(maxPoolSize + 5, 50);
                dataSource.setMaximumPoolSize(newMaxSize);
                logger.info("Increased connection pool size to: {}", newMaxSize);
            }
            
            // 如果活跃连接数长期低于总连接数的30%，考虑减少连接池大小
            if (activeConnections < totalConnections * 0.3 && maxPoolSize > 10) {
                int newMaxSize = Math.max(maxPoolSize - 2, 10);
                dataSource.setMaximumPoolSize(newMaxSize);
                logger.info("Decreased connection pool size to: {}", newMaxSize);
            }
            
        } catch (Exception e) {
            logger.error("Error adjusting connection pool size", e);
        }
    }
}
```

### 11.4 健康检查状态更新
- **检查方法**: 验证健康检查端点能正确反映应用状态
- **检查标准**: 健康检查应包含关键组件状态，响应时间<1秒
- **实现示例**: 参考前面章节的健康检查配置
```

## 12. 中间件技术检查点

### 12.1 Tomcat配置检查

#### 12.1.1 线程池配置 (Critical)
- **检查方法**: 检查Tomcat线程池的核心线程数、最大线程数、队列容量配置
- **检查标准**: 根据应用负载合理配置线程池参数，避免资源浪费或性能瓶颈
- **不正确实例**:
```java
// 错误示例 - 线程池配置不当
@Bean
public ServletWebServerFactory servletContainer() {
    TomcatServletWebServerFactory tomcat = new TomcatServletWebServerFactory();
    tomcat.addConnectorCustomizers(connector -> {
        // 配置过小，可能导致性能瓶颈
        connector.setProperty("maxThreads", "10");
        connector.setProperty("minSpareThreads", "1");
        // 缺少队列配置
    });
    return tomcat;
}
```

- **正确实例**:
```java
// 正确示例 - 合理的线程池配置
@Configuration
public class TomcatConfig {
    
    @Bean
    public ServletWebServerFactory servletContainer() {
        TomcatServletWebServerFactory tomcat = new TomcatServletWebServerFactory();
        tomcat.addConnectorCustomizers(this::customizeConnector);
        return tomcat;
    }
    
    private void customizeConnector(Connector connector) {
        Http11NioProtocol protocol = (Http11NioProtocol) connector.getProtocolHandler();
        
        // 线程池配置
        protocol.setMaxThreads(200);           // 最大线程数
        protocol.setMinSpareThreads(10);       // 最小空闲线程数
        protocol.setMaxConnections(8192);      // 最大连接数
        protocol.setAcceptCount(100);          // 等待队列长度
        
        // 连接超时配置
        protocol.setConnectionTimeout(20000);   // 连接超时 20秒
        protocol.setKeepAliveTimeout(60000);    // Keep-Alive超时 60秒
        protocol.setMaxKeepAliveRequests(100);  // 最大Keep-Alive请求数
        
        // 性能优化
        protocol.setTcpNoDelay(true);
        protocol.setCompression("on");
        protocol.setCompressionMinSize(2048);
        protocol.setCompressibleMimeType("text/html,text/xml,text/plain,text/css,text/javascript,application/javascript,application/json");
    }
    
    // 线程池监控
    @Bean
    public ThreadPoolTaskExecutor tomcatExecutor() {
        ThreadPoolTaskExecutor executor = new ThreadPoolTaskExecutor();
        executor.setCorePoolSize(10);
        executor.setMaxPoolSize(200);
        executor.setQueueCapacity(100);
        executor.setThreadNamePrefix("tomcat-exec-");
        executor.setRejectedExecutionHandler(new ThreadPoolExecutor.CallerRunsPolicy());
        executor.initialize();
        return executor;
    }
    
    // 优雅停机配置
    @Bean
    public GracefulShutdown gracefulShutdown() {
        return new GracefulShutdown();
    }
    
    @Bean
    public WebServerFactoryCustomizer<TomcatServletWebServerFactory> tomcatCustomizer(GracefulShutdown gracefulShutdown) {
        return factory -> factory.addConnectorCustomizers(gracefulShutdown);
    }
    
    private static class GracefulShutdown implements TomcatConnectorCustomizer, ApplicationListener<ContextClosedEvent> {
        
        private static final Logger logger = LoggerFactory.getLogger(GracefulShutdown.class);
        private volatile Connector connector;
        
        @Override
        public void customize(Connector connector) {
            this.connector = connector;
        }
        
        @Override
        public void onApplicationEvent(ContextClosedEvent event) {
            if (connector != null) {
                connector.pause();
                Executor executor = connector.getProtocolHandler().getExecutor();
                if (executor instanceof ThreadPoolExecutor) {
                    try {
                        ThreadPoolExecutor threadPoolExecutor = (ThreadPoolExecutor) executor;
                        threadPoolExecutor.shutdown();
                        if (!threadPoolExecutor.awaitTermination(30, TimeUnit.SECONDS)) {
                            logger.warn("Tomcat线程池未在30秒内关闭，强制关闭");
                            threadPoolExecutor.shutdownNow();
                        }
                    } catch (InterruptedException ex) {
                        Thread.currentThread().interrupt();
                    }
                }
            }
        }
    }
}
```

#### 12.1.2 安全配置 (Critical)
- **检查方法**: 检查Tomcat安全相关配置
- **检查标准**: 禁用不必要的HTTP方法，删除默认应用，配置安全头
- **不正确实例**:
```xml
<!-- 错误示例 - 不安全的server.xml配置 -->
<Connector port="8080" protocol="HTTP/1.1"
           connectionTimeout="20000"
           redirectPort="8443" />
<!-- 缺少安全配置，允许所有HTTP方法 -->
```

- **正确实例**:
```java
// 正确示例 - 安全配置
@Configuration
public class TomcatSecurityConfig {
    
    @Bean
    public WebServerFactoryCustomizer<TomcatServletWebServerFactory> tomcatSecurityCustomizer() {
        return factory -> {
            factory.addContextCustomizers(context -> {
                // 禁用不安全的HTTP方法
                SecurityConstraint securityConstraint = new SecurityConstraint();
                securityConstraint.setUserConstraint("CONFIDENTIAL");
                SecurityCollection collection = new SecurityCollection();
                collection.addPattern("/*");
                collection.addMethod("TRACE");
                collection.addMethod("OPTIONS");
                securityConstraint.addCollection(collection);
                context.addConstraint(securityConstraint);
                
                // 配置安全头
                context.addFilterDef(createSecurityHeadersFilter());
                FilterMap filterMap = new FilterMap();
                filterMap.setFilterName("securityHeaders");
                filterMap.addURLPattern("/*");
                context.addFilterMap(filterMap);
            });
        };
    }
    
    private FilterDef createSecurityHeadersFilter() {
        FilterDef filterDef = new FilterDef();
        filterDef.setFilterName("securityHeaders");
        filterDef.setFilterClass(SecurityHeadersFilter.class.getName());
        return filterDef;
    }
    
    // 安全头过滤器
    public static class SecurityHeadersFilter implements Filter {
        
        @Override
        public void doFilter(ServletRequest request, ServletResponse response, FilterChain chain)
                throws IOException, ServletException {
            
            HttpServletResponse httpResponse = (HttpServletResponse) response;
            
            // 添加安全头
            httpResponse.setHeader("X-Content-Type-Options", "nosniff");
            httpResponse.setHeader("X-Frame-Options", "DENY");
            httpResponse.setHeader("X-XSS-Protection", "1; mode=block");
            httpResponse.setHeader("Strict-Transport-Security", "max-age=31536000; includeSubDomains");
            httpResponse.setHeader("Content-Security-Policy", "default-src 'self'");
            httpResponse.setHeader("Referrer-Policy", "strict-origin-when-cross-origin");
            
            chain.doFilter(request, response);
        }
    }
}
```

### 12.2 Spring Boot配置检查

#### 12.2.1 自动配置机制 (Critical)
- **检查方法**: 检查自动配置的使用和自定义配置
- **检查标准**: 正确理解和使用Spring Boot自动配置，避免配置冲突
- **不正确实例**:
```java
// 错误示例 - 不当的自动配置覆盖
@Configuration
public class BadConfig {
    
    // 错误：完全覆盖自动配置，丢失默认配置
    @Bean
    @Primary
    public DataSource dataSource() {
        HikariDataSource dataSource = new HikariDataSource();
        dataSource.setJdbcUrl("jdbc:mysql://localhost:3306/test");
        // 缺少其他重要配置
        return dataSource;
    }
}
```

- **正确实例**:
```java
// 正确示例 - 合理使用自动配置
@Configuration
@EnableConfigurationProperties({DatabaseProperties.class})
public class DataSourceConfig {
    
    // 使用@ConditionalOnMissingBean避免冲突
    @Bean
    @ConditionalOnMissingBean
    @ConfigurationProperties("spring.datasource.hikari")
    public HikariConfig hikariConfig() {
        return new HikariConfig();
    }
    
    @Bean
    @ConditionalOnMissingBean
    public DataSource dataSource(HikariConfig hikariConfig, DatabaseProperties properties) {
        // 基于自动配置进行定制
        hikariConfig.setJdbcUrl(properties.getUrl());
        hikariConfig.setUsername(properties.getUsername());
        hikariConfig.setPassword(properties.getPassword());
        
        // 添加自定义配置
        hikariConfig.setConnectionTestQuery("SELECT 1");
        hikariConfig.setMaximumPoolSize(properties.getMaxPoolSize());
        
        return new HikariDataSource(hikariConfig);
    }
    
    // 自定义配置属性
    @ConfigurationProperties(prefix = "app.database")
    @Data
    public static class DatabaseProperties {
        private String url;
        private String username;
        private String password;
        private int maxPoolSize = 20;
        private boolean enableMetrics = true;
    }
}
```

#### 12.2.2 Profile配置 (Important)
- **检查方法**: 检查环境配置的隔离和管理
- **检查标准**: 不同环境使用不同Profile，配置外部化
- **不正确实例**:
```yaml
# 错误示例 - 所有环境混在一起
spring:
  datasource:
    url: jdbc:mysql://localhost:3306/test  # 硬编码本地环境
    username: root
    password: password123  # 密码硬编码
  redis:
    host: localhost  # 所有环境都用localhost
    port: 6379
```

- **正确实例**:
```yaml
# application.yml - 通用配置
spring:
  application:
    name: my-application
  profiles:
    active: @spring.profiles.active@  # 通过Maven/Gradle注入

management:
  endpoints:
    web:
      exposure:
        include: health,info,metrics
  endpoint:
    health:
      show-details: when-authorized

---
# application-dev.yml - 开发环境
spring:
  config:
    activate:
      on-profile: dev
  datasource:
    url: jdbc:h2:mem:testdb
    driver-class-name: org.h2.Driver
    username: sa
    password: 
  h2:
    console:
      enabled: true
  redis:
    host: localhost
    port: 6379
    database: 0

logging:
  level:
    com.example: DEBUG

---
# application-prod.yml - 生产环境
spring:
  config:
    activate:
      on-profile: prod
  datasource:
    url: ${DB_URL}
    username: ${DB_USERNAME}
    password: ${DB_PASSWORD}
    hikari:
      maximum-pool-size: 20
      minimum-idle: 5
  redis:
    host: ${REDIS_HOST}
    port: ${REDIS_PORT}
    password: ${REDIS_PASSWORD}
    ssl: true

logging:
  level:
    root: INFO
    com.example: INFO
  file:
    name: /var/log/app/application.log
```

```java
// Profile特定的配置类
@Configuration
@Profile("prod")
public class ProductionConfig {
    
    @Bean
    public HealthIndicator customHealthIndicator() {
        return new CustomHealthIndicator();
    }
    
    @Bean
    public MeterRegistryCustomizer<MeterRegistry> metricsCommonTags() {
        return registry -> registry.config().commonTags("environment", "production");
    }
}

@Configuration
@Profile("dev")
public class DevelopmentConfig {
    
    @Bean
    @Primary
    public Clock testClock() {
        return Clock.fixed(Instant.parse("2023-01-01T00:00:00Z"), ZoneOffset.UTC);
    }
}
```

#### 12.2.3 Actuator安全配置 (Important)
- **检查方法**: 检查Actuator端点的安全配置
- **检查标准**: 敏感端点需要认证，生产环境限制暴露的端点
- **不正确实例**:
```yaml
# 错误示例 - 暴露所有端点且无安全控制
management:
  endpoints:
    web:
      exposure:
        include: "*"  # 暴露所有端点，包括敏感信息
  endpoint:
    health:
      show-details: always  # 总是显示详细信息
```

- **正确实例**:
```yaml
# 正确示例 - 安全的Actuator配置
management:
  endpoints:
    web:
      exposure:
        include: health,info,metrics,prometheus
      base-path: /actuator
  endpoint:
    health:
      show-details: when-authorized
      show-components: when-authorized
    info:
      enabled: true
  security:
    enabled: true
  server:
    port: 8081  # 使用不同端口
```

```java
// Actuator安全配置
@Configuration
@EnableWebSecurity
public class ActuatorSecurityConfig {
    
    @Bean
    @Order(1)
    public SecurityFilterChain actuatorSecurityFilterChain(HttpSecurity http) throws Exception {
        return http
            .securityMatcher(EndpointRequest.toAnyEndpoint())
            .authorizeHttpRequests(auth -> auth
                .requestMatchers(EndpointRequest.to("health", "info")).permitAll()
                .requestMatchers(EndpointRequest.to("metrics", "prometheus")).hasRole("MONITOR")
                .anyRequest().hasRole("ACTUATOR_ADMIN")
            )
            .httpBasic(Customizer.withDefaults())
            .build();
    }
    
    @Bean
    public UserDetailsService actuatorUserDetailsService() {
        UserDetails monitor = User.builder()
            .username("monitor")
            .password("{noop}monitor-password")
            .roles("MONITOR")
            .build();
            
        UserDetails admin = User.builder()
            .username("admin")
            .password("{noop}admin-password")
            .roles("ACTUATOR_ADMIN", "MONITOR")
            .build();
            
        return new InMemoryUserDetailsManager(monitor, admin);
    }
    
    // 自定义健康检查指标
    @Component
    public class CustomHealthIndicator implements HealthIndicator {
        
        @Override
        public Health health() {
            // 检查外部依赖
            boolean databaseUp = checkDatabase();
            boolean redisUp = checkRedis();
            boolean externalServiceUp = checkExternalService();
            
            if (databaseUp && redisUp && externalServiceUp) {
                return Health.up()
                    .withDetail("database", "UP")
                    .withDetail("redis", "UP")
                    .withDetail("external-service", "UP")
                    .build();
            } else {
                return Health.down()
                    .withDetail("database", databaseUp ? "UP" : "DOWN")
                    .withDetail("redis", redisUp ? "UP" : "DOWN")
                    .withDetail("external-service", externalServiceUp ? "UP" : "DOWN")
                    .build();
            }
        }
        
        private boolean checkDatabase() {
            // 实现数据库健康检查逻辑
            return true;
        }
        
        private boolean checkRedis() {
            // 实现Redis健康检查逻辑
            return true;
        }
        
        private boolean checkExternalService() {
            // 实现外部服务健康检查逻辑
            return true;
        }
    }
}
```

### 12.3 Redis配置检查

#### 12.3.1 连接池配置 (Critical)
- **检查方法**: 检查Redis连接池的配置参数
- **检查标准**: 根据应用负载合理配置连接池大小和超时参数
- **不正确实例**:
```yaml
# 错误示例 - Redis连接池配置不当
spring:
  redis:
    host: localhost
    port: 6379
    lettuce:
      pool:
        max-active: 1000  # 连接池过大
        max-idle: 500     # 空闲连接过多
        min-idle: 100     # 最小空闲连接过多
        # 缺少超时配置
```

- **正确实例**:
```yaml
# 正确示例 - 合理的Redis配置
spring:
  redis:
    host: ${REDIS_HOST:localhost}
    port: ${REDIS_PORT:6379}
    password: ${REDIS_PASSWORD:}
    database: ${REDIS_DATABASE:0}
    timeout: 2000ms
    
    lettuce:
      pool:
        max-active: 20      # 最大连接数
        max-idle: 10        # 最大空闲连接
        min-idle: 2         # 最小空闲连接
        max-wait: 2000ms    # 最大等待时间
      shutdown-timeout: 100ms
      
    # 集群配置（如果使用集群）
    cluster:
      nodes: ${REDIS_CLUSTER_NODES:}
      max-redirects: 3
      
    # 哨兵配置（如果使用哨兵）
    sentinel:
      master: ${REDIS_SENTINEL_MASTER:}
      nodes: ${REDIS_SENTINEL_NODES:}
```

```java
// Redis配置类
@Configuration
@EnableCaching
public class RedisConfig {
    
    @Bean
    public LettuceConnectionFactory redisConnectionFactory() {
        LettuceClientConfiguration clientConfig = LettuceClientConfiguration.builder()
            .commandTimeout(Duration.ofSeconds(2))
            .shutdownTimeout(Duration.ofMillis(100))
            .build();
            
        return new LettuceConnectionFactory(
            new RedisStandaloneConfiguration("localhost", 6379), 
            clientConfig
        );
    }
    
    @Bean
    public RedisTemplate<String, Object> redisTemplate(LettuceConnectionFactory connectionFactory) {
        RedisTemplate<String, Object> template = new RedisTemplate<>();
        template.setConnectionFactory(connectionFactory);
        
        // 使用Jackson2JsonRedisSerializer进行序列化
        Jackson2JsonRedisSerializer<Object> jackson2JsonRedisSerializer = 
            new Jackson2JsonRedisSerializer<>(Object.class);
        ObjectMapper om = new ObjectMapper();
        om.setVisibility(PropertyAccessor.ALL, JsonAutoDetect.Visibility.ANY);
        om.activateDefaultTyping(LaissezFaireSubTypeValidator.instance, ObjectMapper.DefaultTyping.NON_FINAL);
        jackson2JsonRedisSerializer.setObjectMapper(om);
        
        // 设置序列化器
        template.setKeySerializer(new StringRedisSerializer());
        template.setHashKeySerializer(new StringRedisSerializer());
        template.setValueSerializer(jackson2JsonRedisSerializer);
        template.setHashValueSerializer(jackson2JsonRedisSerializer);
        
        template.afterPropertiesSet();
        return template;
    }
    
    @Bean
    public CacheManager cacheManager(LettuceConnectionFactory connectionFactory) {
        RedisCacheConfiguration config = RedisCacheConfiguration.defaultCacheConfig()
            .entryTtl(Duration.ofMinutes(10))  // 默认过期时间
            .serializeKeysWith(RedisSerializationContext.SerializationPair.fromSerializer(new StringRedisSerializer()))
            .serializeValuesWith(RedisSerializationContext.SerializationPair.fromSerializer(new GenericJackson2JsonRedisSerializer()))
            .disableCachingNullValues();
            
        return RedisCacheManager.builder(connectionFactory)
            .cacheDefaults(config)
            .build();
    }
    
    // Redis健康检查
    @Component
    public static class RedisHealthIndicator implements HealthIndicator {
        
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
                } else {
                    return Health.down()
                        .withDetail("redis", "Ping failed")
                        .withDetail("ping", result)
                        .build();
                }
            } catch (Exception e) {
                return Health.down(e)
                    .withDetail("redis", "Connection failed")
                    .build();
            }
        }
    }
}
```

#### 12.3.2 缓存策略配置 (Important)
- **检查方法**: 检查缓存的过期策略、淘汰策略和一致性保证
- **检查标准**: 合理设置TTL，实现缓存穿透、击穿、雪崩防护
- **不正确实例**:
```java
// 错误示例 - 缺乏缓存保护机制
@Service
public class BadCacheService {
    
    @Autowired
    private RedisTemplate<String, Object> redisTemplate;
    
    public Object getData(String key) {
        Object data = redisTemplate.opsForValue().get(key);
        if (data == null) {
            // 直接查询数据库，可能导致缓存击穿
            data = queryFromDatabase(key);
            redisTemplate.opsForValue().set(key, data); // 没有设置过期时间
        }
        return data;
    }
}
```

- **正确实例**:
```java
// 正确示例 - 完善的缓存策略
@Service
public class CacheService {
    
    private static final Logger logger = LoggerFactory.getLogger(CacheService.class);
    private static final String LOCK_PREFIX = "lock:";
    private static final String NULL_VALUE = "NULL";
    private static final Duration DEFAULT_TTL = Duration.ofMinutes(30);
    private static final Duration NULL_TTL = Duration.ofMinutes(5);
    
    @Autowired
    private RedisTemplate<String, Object> redisTemplate;
    
    @Autowired
    private RedissonClient redissonClient;
    
    // 防止缓存穿透、击穿、雪崩的缓存方法
    public Object getDataWithProtection(String key) {
        // 1. 先从缓存获取
        Object data = redisTemplate.opsForValue().get(key);
        
        if (data != null) {
            // 缓存命中，检查是否为空值标记
            if (NULL_VALUE.equals(data)) {
                return null; // 防止缓存穿透
            }
            return data;
        }
        
        // 2. 缓存未命中，使用分布式锁防止缓存击穿
        String lockKey = LOCK_PREFIX + key;
        RLock lock = redissonClient.getLock(lockKey);
        
        try {
            // 尝试获取锁，最多等待1秒，锁定10秒
            if (lock.tryLock(1, 10, TimeUnit.SECONDS)) {
                try {
                    // 双重检查，防止重复查询
                    data = redisTemplate.opsForValue().get(key);
                    if (data != null) {
                        return NULL_VALUE.equals(data) ? null : data;
                    }
                    
                    // 查询数据库
                    data = queryFromDatabase(key);
                    
                    if (data != null) {
                        // 设置随机过期时间，防止缓存雪崩
                        Duration ttl = DEFAULT_TTL.plus(Duration.ofMinutes(ThreadLocalRandom.current().nextInt(10)));
                        redisTemplate.opsForValue().set(key, data, ttl);
                    } else {
                        // 缓存空值，防止缓存穿透
                        redisTemplate.opsForValue().set(key, NULL_VALUE, NULL_TTL);
                    }
                    
                    return data;
                } finally {
                    lock.unlock();
                }
            } else {
                // 获取锁失败，等待一段时间后重试
                Thread.sleep(50);
                return getDataWithProtection(key);
            }
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
            logger.error("获取分布式锁被中断", e);
            return queryFromDatabase(key); // 降级到直接查询数据库
        }
    }
    
    // 批量预热缓存，防止缓存雪崩
    @PostConstruct
    public void warmUpCache() {
        CompletableFuture.runAsync(() -> {
            try {
                List<String> hotKeys = getHotKeys();
                for (String key : hotKeys) {
                    Object data = queryFromDatabase(key);
                    if (data != null) {
                        Duration ttl = DEFAULT_TTL.plus(Duration.ofMinutes(ThreadLocalRandom.current().nextInt(10)));
                        redisTemplate.opsForValue().set(key, data, ttl);
                    }
                    Thread.sleep(10); // 避免对数据库造成压力
                }
                logger.info("缓存预热完成，预热了{}个热点数据", hotKeys.size());
            } catch (Exception e) {
                logger.error("缓存预热失败", e);
            }
        });
    }
    
    // 缓存更新策略
    public void updateCache(String key, Object newData) {
        try {
            // 先更新数据库
            updateDatabase(key, newData);
            
            // 再删除缓存（Cache Aside模式）
            redisTemplate.delete(key);
            
            logger.info("更新缓存: key={}", key);
        } catch (Exception e) {
            logger.error("更新缓存失败: key={}", key, e);
            // 可以考虑重试机制或者发送到消息队列
        }
    }
    
    private Object queryFromDatabase(String key) {
        // 模拟数据库查询
        logger.info("从数据库查询数据: key={}", key);
        return "data_" + key;
    }
    
    private void updateDatabase(String key, Object data) {
        // 模拟数据库更新
        logger.info("更新数据库: key={}, data={}", key, data);
    }
    
    private List<String> getHotKeys() {
        // 获取热点数据的key列表
        return Arrays.asList("hot_key_1", "hot_key_2", "hot_key_3");
    }
}
```

### 12.4 Kafka配置检查

#### 12.4.1 生产者可靠性配置 (Critical)
- **检查方法**: 检查Kafka生产者的可靠性配置参数
- **检查标准**: 确保消息不丢失，配置合适的重试和确认机制
- **不正确实例**:
```java
// 错误示例 - 生产者配置不当
@Bean
public ProducerFactory<String, Object> producerFactory() {
    Map<String, Object> props = new HashMap<>();
    props.put(ProducerConfig.BOOTSTRAP_SERVERS_CONFIG, "kafka:9092");
    props.put(ProducerConfig.KEY_SERIALIZER_CLASS_CONFIG, StringSerializer.class);
    props.put(ProducerConfig.VALUE_SERIALIZER_CLASS_CONFIG, JsonSerializer.class);
    // 缺少可靠性配置，可能导致消息丢失
    return new DefaultKafkaProducerFactory<>(props);
}
```

- **正确实例**:
```java
// 正确示例 - 可靠的生产者配置
@Configuration
public class KafkaProducerConfig {
    
    @Bean
    public ProducerFactory<String, Object> producerFactory() {
        Map<String, Object> props = new HashMap<>();
        
        // 基本配置
        props.put(ProducerConfig.BOOTSTRAP_SERVERS_CONFIG, "${kafka.bootstrap-servers}");
        props.put(ProducerConfig.KEY_SERIALIZER_CLASS_CONFIG, StringSerializer.class);
        props.put(ProducerConfig.VALUE_SERIALIZER_CLASS_CONFIG, JsonSerializer.class);
        
        // 可靠性配置
        props.put(ProducerConfig.ACKS_CONFIG, "all");                    // 等待所有副本确认
        props.put(ProducerConfig.RETRIES_CONFIG, 3);                     // 重试次数
        props.put(ProducerConfig.RETRY_BACKOFF_MS_CONFIG, 1000);          // 重试间隔
        props.put(ProducerConfig.ENABLE_IDEMPOTENCE_CONFIG, true);        // 启用幂等性
        props.put(ProducerConfig.MAX_IN_FLIGHT_REQUESTS_PER_CONNECTION, 1); // 限制未确认请求数
        
        // 性能配置
        props.put(ProducerConfig.BATCH_SIZE_CONFIG, 16384);               // 批次大小
        props.put(ProducerConfig.LINGER_MS_CONFIG, 10);                   // 等待时间
        props.put(ProducerConfig.BUFFER_MEMORY_CONFIG, 33554432);         // 缓冲区大小
        props.put(ProducerConfig.COMPRESSION_TYPE_CONFIG, "snappy");       // 压缩类型
        
        // 超时配置
        props.put(ProducerConfig.REQUEST_TIMEOUT_MS_CONFIG, 30000);       // 请求超时
        props.put(ProducerConfig.DELIVERY_TIMEOUT_MS_CONFIG, 120000);     // 投递超时
        
        return new DefaultKafkaProducerFactory<>(props);
    }
    
    @Bean
    public KafkaTemplate<String, Object> kafkaTemplate(ProducerFactory<String, Object> producerFactory) {
        KafkaTemplate<String, Object> template = new KafkaTemplate<>(producerFactory);
        
        // 设置默认主题
        template.setDefaultTopic("default-topic");
        
        // 设置生产者监听器
        template.setProducerListener(new ProducerListener<String, Object>() {
            @Override
            public void onSuccess(ProducerRecord<String, Object> producerRecord, RecordMetadata recordMetadata) {
                logger.debug("消息发送成功: topic={}, partition={}, offset={}", 
                    recordMetadata.topic(), recordMetadata.partition(), recordMetadata.offset());
            }
            
            @Override
            public void onError(ProducerRecord<String, Object> producerRecord, RecordMetadata recordMetadata, Exception exception) {
                logger.error("消息发送失败: topic={}, key={}", 
                    producerRecord.topic(), producerRecord.key(), exception);
            }
        });
        
        return template;
    }
    
    // 消息发送服务
    @Service
    public static class MessageProducer {
        
        private static final Logger logger = LoggerFactory.getLogger(MessageProducer.class);
        
        @Autowired
        private KafkaTemplate<String, Object> kafkaTemplate;
        
        public void sendMessage(String topic, String key, Object message) {
            try {
                ListenableFuture<SendResult<String, Object>> future = 
                    kafkaTemplate.send(topic, key, message);
                    
                future.addCallback(
                    result -> logger.info("消息发送成功: topic={}, key={}, offset={}", 
                        topic, key, result.getRecordMetadata().offset()),
                    failure -> logger.error("消息发送失败: topic={}, key={}", 
                        topic, key, failure)
                );
            } catch (Exception e) {
                logger.error("发送消息异常: topic={}, key={}", topic, key, e);
                throw new MessageSendException("Failed to send message", e);
            }
        }
        
        // 同步发送（用于关键消息）
        public void sendMessageSync(String topic, String key, Object message) {
            try {
                SendResult<String, Object> result = kafkaTemplate.send(topic, key, message).get(10, TimeUnit.SECONDS);
                logger.info("同步发送成功: topic={}, key={}, offset={}", 
                    topic, key, result.getRecordMetadata().offset());
            } catch (Exception e) {
                logger.error("同步发送失败: topic={}, key={}", topic, key, e);
                throw new MessageSendException("Failed to send message synchronously", e);
            }
        }
    }
}
```

#### 12.4.2 消费者配置与错误处理 (Critical)
- **检查方法**: 检查Kafka消费者的配置和错误处理机制
- **检查标准**: 确保消息不重复处理，实现死信队列和重试机制
- **不正确实例**:
```java
// 错误示例 - 消费者配置不当
@KafkaListener(topics = "my-topic")
public void handleMessage(String message) {
    // 直接处理消息，没有错误处理
    processMessage(message);
    // 没有手动提交offset，可能导致消息重复或丢失
}
```

- **正确实例**:
```java
// 正确示例 - 完善的消费者配置
@Configuration
public class KafkaConsumerConfig {
    
    @Bean
    public ConsumerFactory<String, Object> consumerFactory() {
        Map<String, Object> props = new HashMap<>();
        
        // 基本配置
        props.put(ConsumerConfig.BOOTSTRAP_SERVERS_CONFIG, "${kafka.bootstrap-servers}");
        props.put(ConsumerConfig.GROUP_ID_CONFIG, "${kafka.consumer.group-id}");
        props.put(ConsumerConfig.KEY_DESERIALIZER_CLASS_CONFIG, StringDeserializer.class);
        props.put(ConsumerConfig.VALUE_DESERIALIZER_CLASS_CONFIG, JsonDeserializer.class);
        
        // 可靠性配置
        props.put(ConsumerConfig.AUTO_OFFSET_RESET_CONFIG, "earliest");   // 从最早的消息开始消费
        props.put(ConsumerConfig.ENABLE_AUTO_COMMIT_CONFIG, false);       // 禁用自动提交
        props.put(ConsumerConfig.ISOLATION_LEVEL_CONFIG, "read_committed"); // 只读取已提交的消息
        
        // 性能配置
        props.put(ConsumerConfig.MAX_POLL_RECORDS_CONFIG, 500);           // 每次拉取的最大记录数
        props.put(ConsumerConfig.MAX_POLL_INTERVAL_MS_CONFIG, 300000);    // 最大轮询间隔
        props.put(ConsumerConfig.SESSION_TIMEOUT_MS_CONFIG, 30000);       // 会话超时
        props.put(ConsumerConfig.HEARTBEAT_INTERVAL_MS_CONFIG, 10000);    // 心跳间隔
        
        // 反序列化配置
        props.put(JsonDeserializer.TRUSTED_PACKAGES, "com.example.model");
        props.put(JsonDeserializer.VALUE_DEFAULT_TYPE, "com.example.model.Message");
        
        return new DefaultKafkaConsumerFactory<>(props);
    }
    
    @Bean
    public ConcurrentKafkaListenerContainerFactory<String, Object> kafkaListenerContainerFactory(
            ConsumerFactory<String, Object> consumerFactory) {
        
        ConcurrentKafkaListenerContainerFactory<String, Object> factory = 
            new ConcurrentKafkaListenerContainerFactory<>();
        factory.setConsumerFactory(consumerFactory);
        
        // 并发配置
        factory.setConcurrency(3);
        factory.getContainerProperties().setAckMode(AckMode.MANUAL_IMMEDIATE);
        
        // 错误处理配置
        factory.setCommonErrorHandler(createErrorHandler());
        
        // 重试配置
        factory.setRetryTemplate(createRetryTemplate());
        factory.setRecoveryCallback(createRecoveryCallback());
        
        return factory;
    }
    
    private CommonErrorHandler createErrorHandler() {
        // 创建死信队列发布器
        DeadLetterPublishingRecoverer recoverer = new DeadLetterPublishingRecoverer(
            kafkaTemplate(),
            (record, exception) -> {
                // 根据异常类型决定死信队列主题
                if (exception instanceof DeserializationException) {
                    return new TopicPartition(record.topic() + ".DLT.deserialization", record.partition());
                } else {
                    return new TopicPartition(record.topic() + ".DLT", record.partition());
                }
            }
        );
        
        // 配置重试策略
        FixedBackOff backOff = new FixedBackOff(1000L, 3); // 重试3次，间隔1秒
        
        return new DefaultErrorHandler(recoverer, backOff);
    }
    
    private RetryTemplate createRetryTemplate() {
        RetryTemplate retryTemplate = new RetryTemplate();
        
        // 重试策略
        FixedBackOffPolicy backOffPolicy = new FixedBackOffPolicy();
        backOffPolicy.setBackOffPeriod(1000L);
        retryTemplate.setBackOffPolicy(backOffPolicy);
        
        // 重试次数
        SimpleRetryPolicy retryPolicy = new SimpleRetryPolicy();
        retryPolicy.setMaxAttempts(3);
        retryTemplate.setRetryPolicy(retryPolicy);
        
        return retryTemplate;
    }
    
    private RecoveryCallback<Void> createRecoveryCallback() {
        return context -> {
            ConsumerRecord<String, Object> record = 
                (ConsumerRecord<String, Object>) context.getAttribute("record");
            Exception exception = (Exception) context.getLastThrowable();
            
            logger.error("消息处理最终失败，发送到死信队列: topic={}, key={}, offset={}", 
                record.topic(), record.key(), record.offset(), exception);
            
            // 发送到死信队列或进行其他处理
            sendToDeadLetterQueue(record, exception);
            
            return null;
        };
    }
    
    @Bean
    public KafkaTemplate<String, Object> kafkaTemplate() {
        return new KafkaTemplate<>(producerFactory());
    }
    
    @Bean
    public ProducerFactory<String, Object> producerFactory() {
        // 复用生产者配置
        return new KafkaProducerConfig().producerFactory();
    }
    
    private void sendToDeadLetterQueue(ConsumerRecord<String, Object> record, Exception exception) {
        // 实现死信队列逻辑
        String dltTopic = record.topic() + ".DLT";
        kafkaTemplate().send(dltTopic, record.key(), record.value());
    }
}

// 消息消费者服务
@Component
public class MessageConsumer {
    
    private static final Logger logger = LoggerFactory.getLogger(MessageConsumer.class);
    
    @KafkaListener(
        topics = "${kafka.topics.user-events}",
        groupId = "${kafka.consumer.group-id}",
        containerFactory = "kafkaListenerContainerFactory"
    )
    public void handleUserEvent(
            @Payload UserEvent event,
            @Header(KafkaHeaders.RECEIVED_TOPIC) String topic,
            @Header(KafkaHeaders.RECEIVED_PARTITION_ID) int partition,
            @Header(KafkaHeaders.OFFSET) long offset,
            Acknowledgment acknowledgment) {
        
        try {
            logger.info("接收到用户事件: topic={}, partition={}, offset={}, event={}", 
                topic, partition, offset, event);
            
            // 处理业务逻辑
            processUserEvent(event);
            
            // 手动提交offset
            acknowledgment.acknowledge();
            
            logger.debug("用户事件处理完成: topic={}, partition={}, offset={}", 
                topic, partition, offset);
                
        } catch (Exception e) {
            logger.error("处理用户事件失败: topic={}, partition={}, offset={}, event={}", 
                topic, partition, offset, event, e);
            
            // 不提交offset，让重试机制处理
            throw new MessageProcessingException("Failed to process user event", e);
        }
    }
    
    @Retryable(
        value = {TransientException.class},
        maxAttempts = 3,
        backoff = @Backoff(delay = 1000, multiplier = 2)
    )
    private void processUserEvent(UserEvent event) {
        // 实现具体的业务逻辑
        if (event.getType() == UserEventType.REGISTRATION) {
            handleUserRegistration(event);
        } else if (event.getType() == UserEventType.LOGIN) {
            handleUserLogin(event);
        }
    }
    
    @Recover
    private void recover(TransientException ex, UserEvent event) {
        logger.error("用户事件处理最终失败，需要人工介入: event={}", event, ex);
        // 发送告警或记录到特殊表中
        alertService.sendAlert("User event processing failed", ex);
    }
    
    private void handleUserRegistration(UserEvent event) {
        // 处理用户注册事件
    }
    
    private void handleUserLogin(UserEvent event) {
        // 处理用户登录事件
    }
}
```

### 12.5 Elasticsearch配置检查

#### 12.5.1 索引设计与映射配置 (Critical)
- **检查方法**: 检查Elasticsearch索引的映射设计和分析器配置
- **检查标准**: 合理设计字段类型、分析器和索引策略
- **不正确实例**:
```java
// 错误示例 - 索引设计不当
public void createIndex() {
    CreateIndexRequest request = new CreateIndexRequest("products");
    // 没有设置映射，使用默认映射
    // 没有配置分片和副本
    elasticsearchClient.indices().create(request, RequestOptions.DEFAULT);
}
```

- **正确实例**:
```java
// 正确示例 - 完善的索引设计
@Configuration
public class ElasticsearchConfig {
    
    @Bean
    public RestHighLevelClient elasticsearchClient() {
        ClientConfiguration clientConfiguration = ClientConfiguration.builder()
            .connectedTo("${elasticsearch.hosts}")
            .withConnectTimeout(Duration.ofSeconds(5))
            .withSocketTimeout(Duration.ofSeconds(30))
            .withBasicAuth("${elasticsearch.username}", "${elasticsearch.password}")
            .build();
        
        return RestClients.create(clientConfiguration).rest();
    }
    
    @Component
    public static class IndexManager {
        
        private static final Logger logger = LoggerFactory.getLogger(IndexManager.class);
        
        @Autowired
        private RestHighLevelClient elasticsearchClient;
        
        @PostConstruct
        public void initializeIndices() {
            createProductIndex();
            createUserIndex();
        }
        
        public void createProductIndex() {
            String indexName = "products";
            
            try {
                if (indexExists(indexName)) {
                    logger.info("索引已存在: {}", indexName);
                    return;
                }
                
                CreateIndexRequest request = new CreateIndexRequest(indexName);
                
                // 设置分片和副本
                request.settings(Settings.builder()
                    .put("index.number_of_shards", 3)
                    .put("index.number_of_replicas", 1)
                    .put("index.refresh_interval", "1s")
                    .put("index.max_result_window", 50000)
                    // 自定义分析器
                    .put("analysis.analyzer.product_analyzer.type", "custom")
                    .put("analysis.analyzer.product_analyzer.tokenizer", "ik_max_word")
                    .putList("analysis.analyzer.product_analyzer.filter", 
                        Arrays.asList("lowercase", "stop"))
                );
                
                // 定义映射
                Map<String, Object> properties = new HashMap<>();
                
                // 产品名称 - 支持全文搜索
                Map<String, Object> nameField = new HashMap<>();
                nameField.put("type", "text");
                nameField.put("analyzer", "product_analyzer");
                nameField.put("search_analyzer", "ik_smart");
                // 添加keyword子字段用于聚合和排序
                Map<String, Object> nameKeyword = new HashMap<>();
                nameKeyword.put("type", "keyword");
                nameKeyword.put("ignore_above", 256);
                nameField.put("fields", Map.of("keyword", nameKeyword));
                
                // 产品描述
                Map<String, Object> descriptionField = new HashMap<>();
                descriptionField.put("type", "text");
                descriptionField.put("analyzer", "product_analyzer");
                
                // 价格 - 数值类型
                Map<String, Object> priceField = new HashMap<>();
                priceField.put("type", "double");
                priceField.put("index", true);
                
                // 分类 - 关键词类型
                Map<String, Object> categoryField = new HashMap<>();
                categoryField.put("type", "keyword");
                
                // 标签 - 关键词数组
                Map<String, Object> tagsField = new HashMap<>();
                tagsField.put("type", "keyword");
                
                // 创建时间
                Map<String, Object> createTimeField = new HashMap<>();
                createTimeField.put("type", "date");
                createTimeField.put("format", "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis");
                
                // 地理位置
                Map<String, Object> locationField = new HashMap<>();
                locationField.put("type", "geo_point");
                
                // 嵌套对象 - 产品属性
                Map<String, Object> attributesField = new HashMap<>();
                attributesField.put("type", "nested");
                Map<String, Object> attributeProperties = new HashMap<>();
                attributeProperties.put("name", Map.of("type", "keyword"));
                attributeProperties.put("value", Map.of("type", "text", "analyzer", "keyword"));
                attributesField.put("properties", attributeProperties);
                
                properties.put("name", nameField);
                properties.put("description", descriptionField);
                properties.put("price", priceField);
                properties.put("category", categoryField);
                properties.put("tags", tagsField);
                properties.put("createTime", createTimeField);
                properties.put("location", locationField);
                properties.put("attributes", attributesField);
                
                Map<String, Object> mapping = new HashMap<>();
                mapping.put("properties", properties);
                
                request.mapping(mapping);
                
                CreateIndexResponse response = elasticsearchClient.indices().create(request, RequestOptions.DEFAULT);
                logger.info("索引创建成功: {}, acknowledged: {}", indexName, response.isAcknowledged());
                
            } catch (IOException e) {
                logger.error("创建索引失败: {}", indexName, e);
                throw new RuntimeException("Failed to create product index", e);
            }
        }
        
        private boolean indexExists(String indexName) throws IOException {
            GetIndexRequest request = new GetIndexRequest(indexName);
            return elasticsearchClient.indices().exists(request, RequestOptions.DEFAULT);
        }
    }
}
```

#### 12.5.2 查询优化与批量操作 (Important)
- **检查方法**: 检查Elasticsearch查询的性能和批量操作的实现
- **检查标准**: 使用过滤器而非查询、合理分页、批量操作提升性能
- **不正确实例**:
```java
// 错误示例 - 查询性能差
public List<Product> searchProducts(String keyword, int page, int size) {
    SearchRequest request = new SearchRequest("products");
    SearchSourceBuilder sourceBuilder = new SearchSourceBuilder();
    
    // 错误：使用match查询进行精确匹配
    sourceBuilder.query(QueryBuilders.matchQuery("category", "electronics"));
    // 错误：深度分页性能差
    sourceBuilder.from(page * size).size(size);
    
    // 没有使用批量操作
    return executeSearch(request);
}
```

- **正确实例**:
```java
// 正确示例 - 优化的查询和批量操作
@Service
public class ProductSearchService {
    
    private static final Logger logger = LoggerFactory.getLogger(ProductSearchService.class);
    private static final int BULK_SIZE = 1000;
    
    @Autowired
    private RestHighLevelClient elasticsearchClient;
    
    // 优化的搜索方法
    public SearchResult<Product> searchProducts(ProductSearchRequest searchRequest) {
        try {
            SearchRequest request = new SearchRequest("products");
            SearchSourceBuilder sourceBuilder = new SearchSourceBuilder();
            
            // 构建复合查询
            BoolQueryBuilder boolQuery = QueryBuilders.boolQuery();
            
            // 全文搜索
            if (StringUtils.hasText(searchRequest.getKeyword())) {
                boolQuery.must(QueryBuilders.multiMatchQuery(searchRequest.getKeyword())
                    .field("name", 2.0f)  // 提升name字段权重
                    .field("description")
                    .type(MultiMatchQueryBuilder.Type.BEST_FIELDS)
                    .fuzziness(Fuzziness.AUTO));
            }
            
            // 使用过滤器进行精确匹配（性能更好）
            if (StringUtils.hasText(searchRequest.getCategory())) {
                boolQuery.filter(QueryBuilders.termQuery("category", searchRequest.getCategory()));
            }
            
            // 价格范围过滤
            if (searchRequest.getMinPrice() != null || searchRequest.getMaxPrice() != null) {
                RangeQueryBuilder priceRange = QueryBuilders.rangeQuery("price");
                if (searchRequest.getMinPrice() != null) {
                    priceRange.gte(searchRequest.getMinPrice());
                }
                if (searchRequest.getMaxPrice() != null) {
                    priceRange.lte(searchRequest.getMaxPrice());
                }
                boolQuery.filter(priceRange);
            }
            
            // 标签过滤
            if (searchRequest.getTags() != null && !searchRequest.getTags().isEmpty()) {
                boolQuery.filter(QueryBuilders.termsQuery("tags", searchRequest.getTags()));
            }
            
            // 地理位置过滤
            if (searchRequest.getLocation() != null && searchRequest.getDistance() != null) {
                boolQuery.filter(QueryBuilders.geoDistanceQuery("location")
                    .point(searchRequest.getLocation().getLat(), searchRequest.getLocation().getLon())
                    .distance(searchRequest.getDistance(), DistanceUnit.KILOMETERS));
            }
            
            sourceBuilder.query(boolQuery);
            
            // 排序
            if (searchRequest.getSortField() != null) {
                SortOrder sortOrder = searchRequest.getSortDirection() == SortDirection.DESC ? 
                    SortOrder.DESC : SortOrder.ASC;
                sourceBuilder.sort(searchRequest.getSortField(), sortOrder);
            } else {
                // 默认按相关性排序
                sourceBuilder.sort("_score", SortOrder.DESC);
            }
            
            // 分页 - 使用search_after避免深度分页问题
            if (searchRequest.getSearchAfter() != null) {
                sourceBuilder.searchAfter(searchRequest.getSearchAfter());
            } else {
                sourceBuilder.from(searchRequest.getFrom());
            }
            sourceBuilder.size(Math.min(searchRequest.getSize(), 1000)); // 限制最大返回数量
            
            // 高亮
            if (StringUtils.hasText(searchRequest.getKeyword())) {
                HighlightBuilder highlightBuilder = new HighlightBuilder()
                    .field("name")
                    .field("description")
                    .preTags("<mark>")
                    .postTags("</mark>");
                sourceBuilder.highlighter(highlightBuilder);
            }
            
            // 聚合
            if (searchRequest.isIncludeAggregations()) {
                // 分类聚合
                sourceBuilder.aggregation(
                    AggregationBuilders.terms("categories")
                        .field("category")
                        .size(10)
                );
                
                // 价格范围聚合
                sourceBuilder.aggregation(
                    AggregationBuilders.range("price_ranges")
                        .field("price")
                        .addRange("0-100", 0, 100)
                        .addRange("100-500", 100, 500)
                        .addRange("500+", 500, Double.MAX_VALUE)
                );
            }
            
            request.source(sourceBuilder);
            
            SearchResponse response = elasticsearchClient.search(request, RequestOptions.DEFAULT);
            
            return buildSearchResult(response);
            
        } catch (IOException e) {
            logger.error("搜索产品失败", e);
            throw new SearchException("Failed to search products", e);
        }
    }
    
    // 批量索引产品
    public void bulkIndexProducts(List<Product> products) {
        if (products.isEmpty()) {
            return;
        }
        
        // 分批处理
        List<List<Product>> batches = Lists.partition(products, BULK_SIZE);
        
        for (List<Product> batch : batches) {
            try {
                BulkRequest bulkRequest = new BulkRequest();
                
                for (Product product : batch) {
                    IndexRequest indexRequest = new IndexRequest("products")
                        .id(product.getId().toString())
                        .source(convertProductToMap(product), XContentType.JSON);
                    bulkRequest.add(indexRequest);
                }
                
                BulkResponse bulkResponse = elasticsearchClient.bulk(bulkRequest, RequestOptions.DEFAULT);
                
                if (bulkResponse.hasFailures()) {
                    logger.error("批量索引部分失败: {}", bulkResponse.buildFailureMessage());
                    
                    // 处理失败的文档
                    for (BulkItemResponse itemResponse : bulkResponse) {
                        if (itemResponse.isFailed()) {
                            logger.error("文档索引失败: id={}, error={}", 
                                itemResponse.getId(), itemResponse.getFailureMessage());
                        }
                    }
                } else {
                    logger.info("批量索引成功: {} 个文档", batch.size());
                }
                
            } catch (IOException e) {
                logger.error("批量索引异常", e);
                throw new IndexException("Failed to bulk index products", e);
            }
        }
    }
    
    // 批量更新产品
    public void bulkUpdateProducts(List<ProductUpdate> updates) {
        if (updates.isEmpty()) {
            return;
        }
        
        List<List<ProductUpdate>> batches = Lists.partition(updates, BULK_SIZE);
        
        for (List<ProductUpdate> batch : batches) {
            try {
                BulkRequest bulkRequest = new BulkRequest();
                
                for (ProductUpdate update : batch) {
                    UpdateRequest updateRequest = new UpdateRequest("products", update.getId())
                        .doc(convertUpdateToMap(update), XContentType.JSON)
                        .docAsUpsert(false)
                        .retryOnConflict(3);
                    bulkRequest.add(updateRequest);
                }
                
                BulkResponse bulkResponse = elasticsearchClient.bulk(bulkRequest, RequestOptions.DEFAULT);
                
                if (bulkResponse.hasFailures()) {
                    logger.error("批量更新部分失败: {}", bulkResponse.buildFailureMessage());
                } else {
                    logger.info("批量更新成功: {} 个文档", batch.size());
                }
                
            } catch (IOException e) {
                logger.error("批量更新异常", e);
                throw new UpdateException("Failed to bulk update products", e);
            }
        }
    }
    
    private SearchResult<Product> buildSearchResult(SearchResponse response) {
        List<Product> products = new ArrayList<>();
        Map<String, List<String>> highlights = new HashMap<>();
        
        for (SearchHit hit : response.getHits()) {
            Product product = convertMapToProduct(hit.getSourceAsMap());
            products.add(product);
            
            // 处理高亮
            if (hit.getHighlightFields() != null && !hit.getHighlightFields().isEmpty()) {
                Map<String, HighlightField> highlightFields = hit.getHighlightFields();
                List<String> highlightTexts = new ArrayList<>();
                for (HighlightField field : highlightFields.values()) {
                    for (Text fragment : field.getFragments()) {
                        highlightTexts.add(fragment.string());
                    }
                }
                highlights.put(product.getId().toString(), highlightTexts);
            }
        }
        
        // 处理聚合结果
        Map<String, Object> aggregations = new HashMap<>();
        if (response.getAggregations() != null) {
            for (Aggregation aggregation : response.getAggregations()) {
                aggregations.put(aggregation.getName(), parseAggregation(aggregation));
            }
        }
        
        return SearchResult.<Product>builder()
            .items(products)
            .total(response.getHits().getTotalHits().value)
            .highlights(highlights)
            .aggregations(aggregations)
            .took(response.getTook().millis())
            .build();
    }
    
    private Map<String, Object> convertProductToMap(Product product) {
        // 实现Product到Map的转换
        return new HashMap<>();
    }
    
    private Product convertMapToProduct(Map<String, Object> map) {
        // 实现Map到Product的转换
        return new Product();
    }
    
    private Map<String, Object> convertUpdateToMap(ProductUpdate update) {
        // 实现ProductUpdate到Map的转换
        return new HashMap<>();
    }
    
    private Object parseAggregation(Aggregation aggregation) {
        // 解析聚合结果
        return null;
    }
}
```

### 12.6 MySQL配置检查

#### 12.6.1 连接池与事务管理 (Critical)
- **检查方法**: 检查MySQL连接池配置和事务管理策略
- **检查标准**: 合理配置连接池参数，正确使用事务传播和隔离级别
- **不正确实例**:
```yaml
# 错误示例 - 连接池配置不当
spring:
  datasource:
    url: jdbc:mysql://localhost:3306/mydb
    username: root
    password: password
    # 使用默认连接池配置，可能导致性能问题
```

- **正确实例**:
```yaml
# 正确示例 - 优化的数据源配置
spring:
  datasource:
    url: jdbc:mysql://${DB_HOST:localhost}:${DB_PORT:3306}/${DB_NAME:mydb}?useSSL=${DB_SSL:false}&serverTimezone=UTC&characterEncoding=UTF-8&useUnicode=true&allowPublicKeyRetrieval=true&rewriteBatchedStatements=true
    username: ${DB_USERNAME:root}
    password: ${DB_PASSWORD:password}
    driver-class-name: com.mysql.cj.jdbc.Driver
    
    # HikariCP连接池配置
    hikari:
      pool-name: HikariCP-Pool
      minimum-idle: 10                    # 最小空闲连接数
      maximum-pool-size: 50               # 最大连接数
      idle-timeout: 600000                # 连接最大空闲时间（毫秒）
      max-lifetime: 1800000               # 连接最大生命周期（毫秒）
      connection-timeout: 30000           # 获取连接超时时间（毫秒）
      validation-timeout: 5000            # 验证超时时间（毫秒）
      connection-test-query: SELECT 1     # 连接测试查询
      leak-detection-threshold: 60000     # 连接泄漏检测阈值（毫秒）
      
      # 连接属性
      data-source-properties:
        cachePrepStmts: true              # 启用预编译语句缓存
        prepStmtCacheSize: 250            # 预编译语句缓存大小
        prepStmtCacheSqlLimit: 2048       # 预编译语句最大长度
        useServerPrepStmts: true          # 使用服务器端预编译语句
        useLocalSessionState: true        # 使用本地会话状态
        rewriteBatchedStatements: true    # 重写批量语句
        cacheResultSetMetadata: true      # 缓存结果集元数据
        cacheServerConfiguration: true    # 缓存服务器配置
        elideSetAutoCommits: true         # 省略自动提交设置
        maintainTimeStats: false          # 不维护时间统计

  # JPA配置
  jpa:
    hibernate:
      ddl-auto: validate                  # 生产环境使用validate
      naming:
        physical-strategy: org.hibernate.boot.model.naming.PhysicalNamingStrategyStandardImpl
    properties:
      hibernate:
        dialect: org.hibernate.dialect.MySQL8Dialect
        format_sql: false                 # 生产环境关闭SQL格式化
        show_sql: false                   # 生产环境关闭SQL显示
        use_sql_comments: false           # 生产环境关闭SQL注释
        jdbc:
          batch_size: 50                  # 批量操作大小
          fetch_size: 50                  # 获取大小
          time_zone: UTC                  # 时区设置
        connection:
          provider_disables_autocommit: true # 禁用自动提交
        cache:
          use_second_level_cache: true    # 启用二级缓存
          use_query_cache: true           # 启用查询缓存
          region:
            factory_class: org.hibernate.cache.jcache.JCacheRegionFactory
```

```java
// 数据源配置类
@Configuration
@EnableTransactionManagement
@EnableJpaRepositories(basePackages = "com.example.repository")
public class DatabaseConfig {
    
    private static final Logger logger = LoggerFactory.getLogger(DatabaseConfig.class);
    
    @Bean
    @Primary
    @ConfigurationProperties("spring.datasource.hikari")
    public HikariConfig hikariConfig() {
        return new HikariConfig();
    }
    
    @Bean
    @Primary
    public DataSource dataSource(HikariConfig hikariConfig) {
        // 添加连接池监控
        hikariConfig.setMetricRegistry(new MetricRegistry());
        hikariConfig.setHealthCheckRegistry(new HealthCheckRegistry());
        
        // 添加连接池事件监听器
        hikariConfig.addDataSourceProperty("connectionInitSql", "SET NAMES utf8mb4");
        
        HikariDataSource dataSource = new HikariDataSource(hikariConfig);
        
        // 添加连接池监控
        dataSource.setMetricRegistry(meterRegistry());
        
        return dataSource;
    }
    
    @Bean
    public PlatformTransactionManager transactionManager(DataSource dataSource) {
        DataSourceTransactionManager txManager = new DataSourceTransactionManager(dataSource);
        
        // 设置默认超时（秒）
        txManager.setDefaultTimeout(30);
        
        // 设置事务同步
        txManager.setTransactionSynchronization(AbstractPlatformTransactionManager.SYNCHRONIZATION_ON_ACTUAL_TRANSACTION);
        
        return txManager;
    }
    
    @Bean
    public MeterRegistry meterRegistry() {
        return new SimpleMeterRegistry();
    }
    
    // 数据库健康检查
    @Component
    public static class DatabaseHealthIndicator implements HealthIndicator {
        
        @Autowired
        private DataSource dataSource;
        
        @Override
        public Health health() {
            try (Connection connection = dataSource.getConnection()) {
                if (connection.isValid(5)) {
                    return Health.up()
                        .withDetail("database", "MySQL")
                        .withDetail("validationQuery", "SELECT 1")
                        .build();
                } else {
                    return Health.down()
                        .withDetail("database", "MySQL")
                        .withDetail("error", "Connection validation failed")
                        .build();
                }
            } catch (SQLException e) {
                return Health.down(e)
                    .withDetail("database", "MySQL")
                    .withDetail("error", e.getMessage())
                    .build();
            }
        }
    }
    
    // 连接池监控
    @Component
    public static class HikariMetrics {
        
        @Autowired
        private DataSource dataSource;
        
        @Autowired
        private MeterRegistry meterRegistry;
        
        @PostConstruct
        public void bindMetrics() {
            if (dataSource instanceof HikariDataSource) {
                HikariDataSource hikariDataSource = (HikariDataSource) dataSource;
                HikariConfigMXBean hikariConfigMXBean = hikariDataSource.getHikariConfigMXBean();
                HikariPoolMXBean hikariPoolMXBean = hikariDataSource.getHikariPoolMXBean();
                
                // 注册连接池指标
                Gauge.builder("hikari.connections.active")
                    .description("Active connections")
                    .register(meterRegistry, hikariPoolMXBean, HikariPoolMXBean::getActiveConnections);
                    
                Gauge.builder("hikari.connections.idle")
                    .description("Idle connections")
                    .register(meterRegistry, hikariPoolMXBean, HikariPoolMXBean::getIdleConnections);
                    
                Gauge.builder("hikari.connections.pending")
                    .description("Pending connections")
                    .register(meterRegistry, hikariPoolMXBean, HikariPoolMXBean::getThreadsAwaitingConnection);
                    
                Gauge.builder("hikari.connections.total")
                    .description("Total connections")
                    .register(meterRegistry, hikariPoolMXBean, HikariPoolMXBean::getTotalConnections);
            }
        }
    }
}
```

#### 12.6.2 查询优化与批量操作 (Important)
- **检查方法**: 检查SQL查询的性能和批量操作的实现
- **检查标准**: 合理使用索引、避免N+1查询、使用批量操作
- **不正确实例**:
```java
// 错误示例 - 查询性能差
@Service
public class BadUserService {
    
    @Autowired
    private UserRepository userRepository;
    
    @Autowired
    private OrderRepository orderRepository;
    
    // N+1查询问题
    public List<UserDto> getUsersWithOrders() {
        List<User> users = userRepository.findAll();
        return users.stream()
            .map(user -> {
                List<Order> orders = orderRepository.findByUserId(user.getId()); // N+1查询
                return new UserDto(user, orders);
            })
            .collect(Collectors.toList());
    }
    
    // 逐条插入，性能差
    public void createUsers(List<UserDto> userDtos) {
        for (UserDto dto : userDtos) {
            User user = new User(dto);
            userRepository.save(user); // 逐条保存
        }
    }
}
```

- **正确实例**:
```java
// 正确示例 - 优化的查询和批量操作
@Service
@Transactional(readOnly = true)
public class UserService {
    
    private static final Logger logger = LoggerFactory.getLogger(UserService.class);
    private static final int BATCH_SIZE = 1000;
    
    @Autowired
    private UserRepository userRepository;
    
    @Autowired
    private OrderRepository orderRepository;
    
    @Autowired
    private EntityManager entityManager;
    
    // 使用JOIN FETCH避免N+1查询
    public List<UserDto> getUsersWithOrders() {
        List<User> users = userRepository.findAllWithOrders();
        return users.stream()
            .map(user -> new UserDto(user, user.getOrders()))
            .collect(Collectors.toList());
    }
    
    // 分页查询避免内存溢出
    public Page<UserDto> getUsersWithOrdersPaged(Pageable pageable) {
        Page<User> users = userRepository.findAllWithOrders(pageable);
        return users.map(user -> new UserDto(user, user.getOrders()));
    }
    
    // 使用Specification进行动态查询
    public Page<User> searchUsers(UserSearchCriteria criteria, Pageable pageable) {
        Specification<User> spec = Specification.where(null);
        
        if (StringUtils.hasText(criteria.getName())) {
            spec = spec.and((root, query, cb) -> 
                cb.like(cb.lower(root.get("name")), "%" + criteria.getName().toLowerCase() + "%"));
        }
        
        if (StringUtils.hasText(criteria.getEmail())) {
            spec = spec.and((root, query, cb) -> 
                cb.like(cb.lower(root.get("email")), "%" + criteria.getEmail().toLowerCase() + "%"));
        }
        
        if (criteria.getStatus() != null) {
            spec = spec.and((root, query, cb) -> 
                cb.equal(root.get("status"), criteria.getStatus()));
        }
        
        if (criteria.getCreateTimeFrom() != null) {
            spec = spec.and((root, query, cb) -> 
                cb.greaterThanOrEqualTo(root.get("createTime"), criteria.getCreateTimeFrom()));
        }
        
        if (criteria.getCreateTimeTo() != null) {
            spec = spec.and((root, query, cb) -> 
                cb.lessThanOrEqualTo(root.get("createTime"), criteria.getCreateTimeTo()));
        }
        
        return userRepository.findAll(spec, pageable);
    }
    
    // 批量创建用户
    @Transactional
    public void batchCreateUsers(List<UserDto> userDtos) {
        if (userDtos.isEmpty()) {
            return;
        }
        
        List<User> users = userDtos.stream()
            .map(User::new)
            .collect(Collectors.toList());
        
        // 分批处理，避免内存溢出
        List<List<User>> batches = Lists.partition(users, BATCH_SIZE);
        
        for (List<User> batch : batches) {
            batchInsertUsers(batch);
            entityManager.flush();
            entityManager.clear(); // 清理一级缓存
        }
        
        logger.info("批量创建用户完成: {} 个用户", users.size());
    }
    
    private void batchInsertUsers(List<User> users) {
        // 使用JPA批量保存
        userRepository.saveAll(users);
    }
    
    // 使用原生SQL进行批量更新
    @Transactional
    public int batchUpdateUserStatus(List<Long> userIds, UserStatus newStatus) {
        if (userIds.isEmpty()) {
            return 0;
        }
        
        String sql = "UPDATE users SET status = :status, update_time = :updateTime WHERE id IN (:userIds)";
        
        Query query = entityManager.createNativeQuery(sql)
            .setParameter("status", newStatus.name())
            .setParameter("updateTime", LocalDateTime.now())
            .setParameter("userIds", userIds);
        
        int updatedCount = query.executeUpdate();
        logger.info("批量更新用户状态: {} 个用户更新为 {}", updatedCount, newStatus);
        
        return updatedCount;
    }
    
    // 使用JPQL进行复杂查询
    public List<UserStatistics> getUserStatistics(LocalDate startDate, LocalDate endDate) {
        String jpql = """
            SELECT new com.example.dto.UserStatistics(
                u.status,
                COUNT(u),
                AVG(SIZE(u.orders)),
                SUM(o.amount)
            )
            FROM User u
            LEFT JOIN u.orders o
            WHERE u.createTime BETWEEN :startDate AND :endDate
            GROUP BY u.status
            ORDER BY COUNT(u) DESC
            """;
        
        return entityManager.createQuery(jpql, UserStatistics.class)
            .setParameter("startDate", startDate.atStartOfDay())
            .setParameter("endDate", endDate.atTime(23, 59, 59))
            .getResultList();
    }
    
    // 使用存储过程
    public void callUserCleanupProcedure(int daysOld) {
        StoredProcedureQuery query = entityManager
            .createStoredProcedureQuery("cleanup_inactive_users")
            .registerStoredProcedureParameter("days_old", Integer.class, ParameterMode.IN)
            .registerStoredProcedureParameter("deleted_count", Integer.class, ParameterMode.OUT)
            .setParameter("days_old", daysOld);
        
        query.execute();
        
        Integer deletedCount = (Integer) query.getOutputParameterValue("deleted_count");
        logger.info("清理非活跃用户: 删除了 {} 个用户", deletedCount);
    }
    
    // 事务传播示例
    @Transactional(propagation = Propagation.REQUIRES_NEW)
    public void createUserAuditLog(Long userId, String action) {
        // 使用新事务记录审计日志，即使主事务回滚也会保存
        UserAuditLog log = new UserAuditLog(userId, action, LocalDateTime.now());
        entityManager.persist(log);
    }
    
    // 只读事务优化
    @Transactional(readOnly = true)
    public Optional<User> findUserById(Long id) {
        return userRepository.findById(id);
    }
    
    // 超时控制
    @Transactional(timeout = 10) // 10秒超时
    public void longRunningOperation() {
        // 长时间运行的操作
    }
}

// Repository接口优化
@Repository
public interface UserRepository extends JpaRepository<User, Long>, JpaSpecificationExecutor<User> {
    
    // 使用JOIN FETCH避免N+1查询
    @Query("SELECT DISTINCT u FROM User u LEFT JOIN FETCH u.orders WHERE u.status = :status")
    List<User> findAllWithOrders(@Param("status") UserStatus status);
    
    // 分页查询with JOIN FETCH
    @Query(value = "SELECT DISTINCT u FROM User u LEFT JOIN FETCH u.orders",
           countQuery = "SELECT COUNT(u) FROM User u")
    Page<User> findAllWithOrders(Pageable pageable);
    
    // 使用索引字段查询
    @Query("SELECT u FROM User u WHERE u.email = :email")
    Optional<User> findByEmail(@Param("email") String email);
    
    // 批量查询
    @Query("SELECT u FROM User u WHERE u.id IN :ids")
    List<User> findByIdIn(@Param("ids") List<Long> ids);
    
    // 统计查询
    @Query("SELECT COUNT(u) FROM User u WHERE u.status = :status AND u.createTime >= :since")
    long countByStatusAndCreateTimeAfter(@Param("status") UserStatus status, @Param("since") LocalDateTime since);
    
    // 原生SQL查询（复杂查询）
    @Query(value = """
        SELECT u.*, 
               COUNT(o.id) as order_count,
               COALESCE(SUM(o.amount), 0) as total_amount
        FROM users u
        LEFT JOIN orders o ON u.id = o.user_id
        WHERE u.create_time >= :startDate
        GROUP BY u.id
        HAVING COUNT(o.id) > :minOrderCount
        ORDER BY total_amount DESC
        """, nativeQuery = true)
    List<Object[]> findActiveUsersWithStatistics(
        @Param("startDate") LocalDateTime startDate,
        @Param("minOrderCount") int minOrderCount
    );
    
    // 批量更新
    @Modifying
    @Query("UPDATE User u SET u.lastLoginTime = :loginTime WHERE u.id = :userId")
    int updateLastLoginTime(@Param("userId") Long userId, @Param("loginTime") LocalDateTime loginTime);
    
    // 批量删除
    @Modifying
    @Query("DELETE FROM User u WHERE u.status = :status AND u.createTime < :before")
    int deleteByStatusAndCreateTimeBefore(@Param("status") UserStatus status, @Param("before") LocalDateTime before);
}
```

## 13. 依赖管理检查

### 12.1 依赖安全 (Critical)

#### 12.1.1 第三方库安全漏洞扫描
- **检查方法**: 使用OWASP Dependency Check、Snyk等工具扫描依赖漏洞
- **检查标准**: 定期扫描并及时更新有安全漏洞的依赖
- **不正确实例**:
```xml
<!-- 错误示例 - 使用有安全漏洞的依赖版本 -->
<dependencies>
    <!-- 使用过时版本，存在已知安全漏洞 -->
    <dependency>
        <groupId>org.apache.struts</groupId>
        <artifactId>struts2-core</artifactId>
        <version>2.3.16</version>  <!-- 存在CVE-2017-5638等严重漏洞 -->
    </dependency>
    
    <dependency>
        <groupId>com.fasterxml.jackson.core</groupId>
        <artifactId>jackson-databind</artifactId>
        <version>2.9.8</version>  <!-- 存在反序列化漏洞 -->
    </dependency>
    
    <dependency>
        <groupId>org.springframework</groupId>
        <artifactId>spring-core</artifactId>
        <version>4.3.18.RELEASE</version>  <!-- 过时版本 -->
    </dependency>
</dependencies>

<!-- 正确示例 - 使用安全的依赖版本 -->
<dependencies>
    <!-- 使用最新稳定版本，及时修复安全漏洞 -->
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-web</artifactId>
        <version>2.7.14</version>  <!-- 使用最新稳定版本 -->
    </dependency>
    
    <dependency>
        <groupId>com.fasterxml.jackson.core</groupId>
        <artifactId>jackson-databind</artifactId>
        <version>2.15.2</version>  <!-- 修复了已知漏洞的版本 -->
    </dependency>
    
    <!-- 排除有漏洞的传递依赖 -->
    <dependency>
        <groupId>org.apache.httpcomponents</groupId>
        <artifactId>httpclient</artifactId>
        <version>4.5.14</version>
        <exclusions>
            <exclusion>
                <groupId>commons-logging</groupId>
                <artifactId>commons-logging</artifactId>
            </exclusion>
        </exclusions>
    </dependency>
</dependencies>

<!-- 添加安全扫描插件 -->
<build>
    <plugins>
        <!-- OWASP依赖检查插件 -->
        <plugin>
            <groupId>org.owasp</groupId>
            <artifactId>dependency-check-maven</artifactId>
            <version>8.4.0</version>
            <configuration>
                <failBuildOnCVSS>7</failBuildOnCVSS>  <!-- CVSS评分>=7时构建失败 -->
                <suppressionFile>owasp-suppressions.xml</suppressionFile>
            </configuration>
            <executions>
                <execution>
                    <goals>
                        <goal>check</goal>
                    </goals>
                </execution>
            </executions>
        </plugin>
        
        <!-- Snyk漏洞扫描 -->
        <plugin>
            <groupId>io.snyk</groupId>
            <artifactId>snyk-maven-plugin</artifactId>
            <version>2.2.0</version>
            <executions>
                <execution>
                    <id>snyk-test</id>
                    <goals>
                        <goal>test</goal>
                    </goals>
                </execution>
            </executions>
        </plugin>
    </plugins>
</build>
```

#### 12.1.2 依赖版本管理
- **检查方法**: 检查依赖版本的一致性和更新策略
- **检查标准**: 使用版本管理策略，避免版本冲突
- **不正确实例**:
```xml
<!-- 错误示例 - 版本管理混乱 -->
<dependencies>
    <!-- 同一个库的不同版本 -->
    <dependency>
        <groupId>org.springframework</groupId>
        <artifactId>spring-core</artifactId>
        <version>5.3.21</version>
    </dependency>
    
    <dependency>
        <groupId>org.springframework</groupId>
        <artifactId>spring-context</artifactId>
        <version>5.2.15.RELEASE</version>  <!-- 版本不一致 -->
    </dependency>
    
    <!-- 使用SNAPSHOT版本在生产环境 -->
    <dependency>
        <groupId>com.example</groupId>
        <artifactId>custom-lib</artifactId>
        <version>1.0-SNAPSHOT</version>  <!-- 不稳定版本 -->
    </dependency>
    
    <!-- 版本范围过于宽泛 -->
    <dependency>
        <groupId>commons-lang</groupId>
        <artifactId>commons-lang</artifactId>
        <version>[2.0,)</version>  <!-- 可能引入不兼容版本 -->
    </dependency>
</dependencies>

<!-- 正确示例 - 统一版本管理 -->
<properties>
    <!-- 统一定义版本号 -->
    <spring.version>5.3.23</spring.version>
    <jackson.version>2.15.2</jackson.version>
    <junit.version>5.9.3</junit.version>
    <mockito.version>4.11.0</mockito.version>
</properties>

<dependencyManagement>
    <dependencies>
        <!-- 使用BOM统一管理Spring Boot依赖版本 -->
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-dependencies</artifactId>
            <version>2.7.14</version>
            <type>pom</type>
            <scope>import</scope>
        </dependency>
        
        <!-- 统一管理其他依赖版本 -->
        <dependency>
            <groupId>com.fasterxml.jackson.core</groupId>
            <artifactId>jackson-databind</artifactId>
            <version>${jackson.version}</version>
        </dependency>
        
        <dependency>
            <groupId>org.junit.jupiter</groupId>
            <artifactId>junit-jupiter</artifactId>
            <version>${junit.version}</version>
        </dependency>
    </dependencies>
</dependencyManagement>

<dependencies>
    <!-- 不指定版本，由dependencyManagement管理 -->
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-web</artifactId>
    </dependency>
    
    <dependency>
        <groupId>com.fasterxml.jackson.core</groupId>
        <artifactId>jackson-databind</artifactId>
    </dependency>
    
    <!-- 测试依赖 -->
    <dependency>
        <groupId>org.junit.jupiter</groupId>
        <artifactId>junit-jupiter</artifactId>
        <scope>test</scope>
    </dependency>
</dependencies>

<!-- 版本检查插件 -->
<build>
    <plugins>
        <plugin>
            <groupId>org.codehaus.mojo</groupId>
            <artifactId>versions-maven-plugin</artifactId>
            <version>2.16.0</version>
            <configuration>
                <generateBackupPoms>false</generateBackupPoms>
            </configuration>
        </plugin>
    </plugins>
</build>
```

#### 12.1.3 许可证合规检查
- **检查方法**: 检查第三方依赖的许可证兼容性
- **检查标准**: 确保所有依赖的许可证符合项目要求
- **不正确实例**:
```xml
<!-- 错误示例 - 没有许可证检查 -->
<dependencies>
    <!-- 使用GPL许可证的库，可能与商业项目冲突 -->
    <dependency>
        <groupId>some.gpl</groupId>
        <artifactId>gpl-library</artifactId>
        <version>1.0.0</version>
    </dependency>
    
    <!-- 使用未知许可证的库 -->
    <dependency>
        <groupId>unknown.license</groupId>
        <artifactId>mystery-lib</artifactId>
        <version>2.0.0</version>
    </dependency>
</dependencies>

<!-- 正确示例 - 许可证合规管理 -->
<build>
    <plugins>
        <!-- 许可证检查插件 -->
        <plugin>
            <groupId>org.codehaus.mojo</groupId>
            <artifactId>license-maven-plugin</artifactId>
            <version>2.2.0</version>
            <configuration>
                <licenseMerges>
                    <licenseMerge>Apache License, Version 2.0|Apache 2|Apache License 2.0</licenseMerge>
                    <licenseMerge>MIT License|MIT|The MIT License</licenseMerge>
                </licenseMerges>
                <failOnMissing>true</failOnMissing>
                <failOnBlacklist>true</failOnBlacklist>
                <includedLicenses>
                    <includedLicense>Apache License, Version 2.0</includedLicense>
                    <includedLicense>MIT License</includedLicense>
                    <includedLicense>BSD License</includedLicense>
                    <includedLicense>Eclipse Public License</includedLicense>
                </includedLicenses>
                <excludedLicenses>
                    <excludedLicense>GNU General Public License</excludedLicense>
                    <excludedLicense>GNU Lesser General Public License</excludedLicense>
                </excludedLicenses>
            </configuration>
            <executions>
                <execution>
                    <id>check-licenses</id>
                    <phase>verify</phase>
                    <goals>
                        <goal>check-third-party</goal>
                    </goals>
                </execution>
            </executions>
        </plugin>
        
        <!-- 生成许可证报告 -->
        <plugin>
            <groupId>org.codehaus.mojo</groupId>
            <artifactId>license-maven-plugin</artifactId>
            <version>2.2.0</version>
            <executions>
                <execution>
                    <id>generate-license-report</id>
                    <phase>prepare-package</phase>
                    <goals>
                        <goal>aggregate-third-party-report</goal>
                    </goals>
                </execution>
            </executions>
        </plugin>
    </plugins>
</build>

<!-- 许可证白名单配置文件 src/license/THIRD-PARTY.properties -->
# 允许的许可证列表
Apache License, Version 2.0=true
MIT License=true
BSD License=true
Eclipse Public License=true

# 禁止的许可证列表
GNU General Public License=false
GNU Lesser General Public License=false
```

#### 12.1.4 依赖冲突解决
- **检查方法**: 检查并解决依赖冲突问题
- **检查标准**: 确保没有版本冲突，使用合适的冲突解决策略
- **不正确实例**:
```xml
<!-- 错误示例 - 依赖冲突未解决 -->
<dependencies>
    <dependency>
        <groupId>org.springframework</groupId>
        <artifactId>spring-core</artifactId>
        <version>5.3.21</version>
    </dependency>
    
    <!-- 这个依赖可能引入不同版本的spring-core -->
    <dependency>
        <groupId>org.springframework.security</groupId>
        <artifactId>spring-security-core</artifactId>
        <version>5.6.10</version>
        <!-- 可能依赖spring-core 5.3.18 -->
    </dependency>
    
    <!-- 传递依赖冲突 -->
    <dependency>
        <groupId>com.example</groupId>
        <artifactId>library-a</artifactId>
        <version>1.0.0</version>
        <!-- 依赖commons-lang 2.6 -->
    </dependency>
    
    <dependency>
        <groupId>com.example</groupId>
        <artifactId>library-b</artifactId>
        <version>2.0.0</version>
        <!-- 依赖commons-lang 3.12.0 -->
    </dependency>
</dependencies>

<!-- 正确示例 - 解决依赖冲突 -->
<dependencyManagement>
    <dependencies>
        <!-- 使用BOM统一版本管理 -->
        <dependency>
            <groupId>org.springframework</groupId>
            <artifactId>spring-framework-bom</artifactId>
            <version>5.3.23</version>
            <type>pom</type>
            <scope>import</scope>
        </dependency>
        
        <!-- 强制指定冲突依赖的版本 -->
        <dependency>
            <groupId>commons-lang</groupId>
            <artifactId>commons-lang</artifactId>
            <version>2.6</version>
        </dependency>
        
        <dependency>
            <groupId>org.apache.commons</groupId>
            <artifactId>commons-lang3</artifactId>
            <version>3.12.0</version>
        </dependency>
    </dependencies>
</dependencyManagement>

<dependencies>
    <!-- Spring依赖版本由BOM管理 -->
    <dependency>
        <groupId>org.springframework</groupId>
        <artifactId>spring-core</artifactId>
    </dependency>
    
    <dependency>
        <groupId>org.springframework.security</groupId>
        <artifactId>spring-security-core</artifactId>
    </dependency>
    
    <!-- 排除冲突的传递依赖 -->
    <dependency>
        <groupId>com.example</groupId>
        <artifactId>library-a</artifactId>
        <version>1.0.0</version>
        <exclusions>
            <exclusion>
                <groupId>commons-lang</groupId>
                <artifactId>commons-lang</artifactId>
            </exclusion>
        </exclusions>
    </dependency>
    
    <dependency>
        <groupId>com.example</groupId>
        <artifactId>library-b</artifactId>
        <version>2.0.0</version>
        <exclusions>
            <exclusion>
                <groupId>org.apache.commons</groupId>
                <artifactId>commons-lang3</artifactId>
            </exclusion>
        </exclusions>
    </dependency>
    
    <!-- 显式添加需要的版本 -->
    <dependency>
        <groupId>org.apache.commons</groupId>
        <artifactId>commons-lang3</artifactId>
    </dependency>
</dependencies>

<!-- 依赖分析插件 -->
<build>
    <plugins>
        <!-- 依赖分析插件 -->
        <plugin>
            <groupId>org.apache.maven.plugins</groupId>
            <artifactId>maven-dependency-plugin</artifactId>
            <version>3.6.0</version>
            <executions>
                <execution>
                    <id>analyze-dependencies</id>
                    <phase>verify</phase>
                    <goals>
                        <goal>analyze-only</goal>
                    </goals>
                    <configuration>
                        <failOnWarning>true</failOnWarning>
                        <ignoreNonCompile>true</ignoreNonCompile>
                    </configuration>
                </execution>
            </executions>
        </plugin>
        
        <!-- 依赖树分析 -->
        <!-- mvn dependency:tree -Dverbose -->
        <!-- mvn dependency:analyze-duplicate -->
    </plugins>
</build>
```

### 12.2 依赖优化 (Minor)

#### 12.2.1 移除未使用的依赖
- **检查方法**: 使用Maven dependency插件检查未使用的依赖
- **检查标准**: 定期清理未使用的依赖，减少项目复杂度
- **不正确实例**:
```xml
<!-- 错误示例 - 包含大量未使用的依赖 -->
<dependencies>
    <!-- 添加了但从未使用的依赖 -->
    <dependency>
        <groupId>org.apache.commons</groupId>
        <artifactId>commons-math3</artifactId>
        <version>3.6.1</version>
    </dependency>
    
    <!-- 过时的依赖，已被其他依赖替代 -->
    <dependency>
        <groupId>commons-logging</groupId>
        <artifactId>commons-logging</artifactId>
        <version>1.2</version>
    </dependency>
    
    <!-- 测试依赖但scope错误 -->
    <dependency>
        <groupId>org.mockito</groupId>
        <artifactId>mockito-core</artifactId>
        <version>4.11.0</version>
        <!-- 缺少scope=test -->
    </dependency>
    
    <!-- 重复的功能依赖 -->
    <dependency>
        <groupId>com.google.guava</groupId>
        <artifactId>guava</artifactId>
        <version>32.1.2-jre</version>
    </dependency>
    
    <dependency>
        <groupId>org.apache.commons</groupId>
        <artifactId>commons-collections4</artifactId>
        <version>4.4</version>
        <!-- 与Guava功能重复 -->
    </dependency>
</dependencies>

<!-- 正确示例 - 精简的依赖管理 -->
<dependencies>
    <!-- 只包含实际使用的依赖 -->
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-web</artifactId>
    </dependency>
    
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-data-jpa</artifactId>
    </dependency>
    
    <!-- 正确设置scope -->
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-test</artifactId>
        <scope>test</scope>
    </dependency>
    
    <!-- 选择一个功能完整的库，避免重复 -->
    <dependency>
        <groupId>com.google.guava</groupId>
        <artifactId>guava</artifactId>
        <version>32.1.2-jre</version>
    </dependency>
</dependencies>

<!-- 依赖分析和清理 -->
<build>
    <plugins>
        <!-- 未使用依赖检查 -->
        <plugin>
            <groupId>org.apache.maven.plugins</groupId>
            <artifactId>maven-dependency-plugin</artifactId>
            <version>3.6.0</version>
            <executions>
                <execution>
                    <id>analyze-unused</id>
                    <phase>verify</phase>
                    <goals>
                        <goal>analyze-only</goal>
                    </goals>
                    <configuration>
                        <failOnWarning>false</failOnWarning>
                        <outputXML>true</outputXML>
                    </configuration>
                </execution>
            </executions>
        </plugin>
        
        <!-- 依赖范围检查 -->
        <plugin>
            <groupId>com.github.ferstl</groupId>
            <artifactId>depgraph-maven-plugin</artifactId>
            <version>4.0.2</version>
            <configuration>
                <createImage>true</createImage>
                <imageFormat>png</imageFormat>
            </configuration>
        </plugin>
    </plugins>
</build>

<!-- 使用命令检查未使用依赖 -->
<!-- mvn dependency:analyze -->
<!-- mvn dependency:analyze-dep-mgt -->
<!-- mvn dependency:purge-local-repository -->
```

#### 12.2.2 依赖版本统一管理
- **检查方法**: 检查项目中依赖版本管理的一致性
- **检查标准**: 使用统一的版本管理策略，避免版本碎片化
- **不正确实例**:
```xml
<!-- 错误示例 - 版本管理分散 -->
<dependencies>
    <!-- 版本号硬编码在各个依赖中 -->
    <dependency>
        <groupId>com.fasterxml.jackson.core</groupId>
        <artifactId>jackson-core</artifactId>
        <version>2.15.2</version>
    </dependency>
    
    <dependency>
        <groupId>com.fasterxml.jackson.core</groupId>
        <artifactId>jackson-databind</artifactId>
        <version>2.15.1</version>  <!-- 版本不一致 -->
    </dependency>
    
    <dependency>
        <groupId>com.fasterxml.jackson.datatype</groupId>
        <artifactId>jackson-datatype-jsr310</artifactId>
        <version>2.14.2</version>  <!-- 版本更不一致 -->
    </dependency>
</dependencies>

<!-- 正确示例 - 统一版本管理 -->
<properties>
    <!-- 版本号统一定义 -->
    <jackson.version>2.15.2</jackson.version>
    <spring-boot.version>2.7.14</spring-boot.version>
    <mysql.version>8.0.33</mysql.version>
    <junit.version>5.9.3</junit.version>
</properties>

<dependencyManagement>
    <dependencies>
        <!-- 使用BOM管理相关依赖版本 -->
        <dependency>
            <groupId>com.fasterxml.jackson</groupId>
            <artifactId>jackson-bom</artifactId>
            <version>${jackson.version}</version>
            <type>pom</type>
            <scope>import</scope>
        </dependency>
        
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-dependencies</artifactId>
            <version>${spring-boot.version}</version>
            <type>pom</type>
            <scope>import</scope>
        </dependency>
        
        <!-- 其他依赖版本管理 -->
        <dependency>
            <groupId>mysql</groupId>
            <artifactId>mysql-connector-java</artifactId>
            <version>${mysql.version}</version>
        </dependency>
    </dependencies>
</dependencyManagement>

<dependencies>
    <!-- 不指定版本，由dependencyManagement统一管理 -->
    <dependency>
        <groupId>com.fasterxml.jackson.core</groupId>
        <artifactId>jackson-core</artifactId>
    </dependency>
    
    <dependency>
        <groupId>com.fasterxml.jackson.core</groupId>
        <artifactId>jackson-databind</artifactId>
    </dependency>
    
    <dependency>
        <groupId>com.fasterxml.jackson.datatype</groupId>
        <artifactId>jackson-datatype-jsr310</artifactId>
    </dependency>
</dependencies>

<!-- 版本更新检查 -->
<build>
    <plugins>
        <plugin>
            <groupId>org.codehaus.mojo</groupId>
            <artifactId>versions-maven-plugin</artifactId>
            <version>2.16.0</version>
            <configuration>
                <generateBackupPoms>false</generateBackupPoms>
            </configuration>
        </plugin>
    </plugins>
</build>

<!-- 版本检查命令 -->
<!-- mvn versions:display-dependency-updates -->
<!-- mvn versions:display-plugin-updates -->
<!-- mvn versions:display-property-updates -->
```

#### 12.2.3 传递依赖控制
- **检查方法**: 检查和控制传递依赖的版本和范围
- **检查标准**: 明确管理传递依赖，避免版本冲突和安全风险
- **不正确实例**:
```xml
<!-- 错误示例 - 未控制传递依赖 -->
<dependencies>
    <!-- 直接依赖，但未控制其传递依赖 -->
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-web</artifactId>
        <version>2.7.14</version>
        <!-- 传递依赖可能包含漏洞版本 -->
    </dependency>
    
    <!-- 引入了不需要的传递依赖 -->
    <dependency>
        <groupId>com.example</groupId>
        <artifactId>some-library</artifactId>
        <version>1.0.0</version>
        <!-- 可能传递依赖了log4j 1.x等有安全问题的库 -->
    </dependency>
    
    <!-- 传递依赖版本冲突 -->
    <dependency>
        <groupId>library-a</groupId>
        <artifactId>library-a</artifactId>
        <version>1.0</version>
        <!-- 依赖jackson 2.13.0 -->
    </dependency>
    
    <dependency>
        <groupId>library-b</groupId>
        <artifactId>library-b</artifactId>
        <version>2.0</version>
        <!-- 依赖jackson 2.15.2 -->
    </dependency>
</dependencies>

<!-- 正确示例 - 控制传递依赖 -->
<dependencyManagement>
    <dependencies>
        <!-- 使用BOM控制传递依赖版本 -->
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-dependencies</artifactId>
            <version>2.7.14</version>
            <type>pom</type>
            <scope>import</scope>
        </dependency>
        
        <!-- 强制指定传递依赖版本 -->
        <dependency>
            <groupId>com.fasterxml.jackson.core</groupId>
            <artifactId>jackson-databind</artifactId>
            <version>2.15.2</version>
        </dependency>
        
        <!-- 排除有安全问题的传递依赖 -->
        <dependency>
            <groupId>org.apache.logging.log4j</groupId>
            <artifactId>log4j-core</artifactId>
            <version>2.20.0</version>
        </dependency>
    </dependencies>
</dependencyManagement>

<dependencies>
    <!-- 排除不安全的传递依赖 -->
    <dependency>
        <groupId>com.example</groupId>
        <artifactId>some-library</artifactId>
        <version>1.0.0</version>
        <exclusions>
            <!-- 排除旧版本log4j -->
            <exclusion>
                <groupId>log4j</groupId>
                <artifactId>log4j</artifactId>
            </exclusion>
            <!-- 排除commons-logging -->
            <exclusion>
                <groupId>commons-logging</groupId>
                <artifactId>commons-logging</artifactId>
            </exclusion>
            <!-- 排除旧版本jackson -->
            <exclusion>
                <groupId>com.fasterxml.jackson.core</groupId>
                <artifactId>jackson-databind</artifactId>
            </exclusion>
        </exclusions>
    </dependency>
    
    <!-- 显式添加安全版本 -->
    <dependency>
        <groupId>org.apache.logging.log4j</groupId>
        <artifactId>log4j-core</artifactId>
    </dependency>
    
    <dependency>
        <groupId>com.fasterxml.jackson.core</groupId>
        <artifactId>jackson-databind</artifactId>
    </dependency>
</dependencies>

<!-- 传递依赖分析插件 -->
<build>
    <plugins>
        <!-- 传递依赖分析 -->
        <plugin>
            <groupId>org.apache.maven.plugins</groupId>
            <artifactId>maven-dependency-plugin</artifactId>
            <version>3.6.0</version>
            <executions>
                <execution>
                    <id>analyze-transitive</id>
                    <phase>verify</phase>
                    <goals>
                        <goal>tree</goal>
                    </goals>
                    <configuration>
                        <outputFile>target/dependency-tree.txt</outputFile>
                        <verbose>true</verbose>
                    </configuration>
                </execution>
            </executions>
        </plugin>
        
        <!-- 依赖收敛检查 -->
        <plugin>
            <groupId>org.apache.maven.plugins</groupId>
            <artifactId>maven-enforcer-plugin</artifactId>
            <version>3.4.1</version>
            <executions>
                <execution>
                    <id>enforce-dependency-convergence</id>
                    <goals>
                        <goal>enforce</goal>
                    </goals>
                    <configuration>
                        <rules>
                            <dependencyConvergence/>
                            <bannedDependencies>
                                <excludes>
                                    <!-- 禁止使用有安全问题的依赖 -->
                                    <exclude>log4j:log4j:*:*:compile</exclude>
                                    <exclude>commons-logging:commons-logging:*:*:compile</exclude>
                                    <exclude>org.apache.logging.log4j:log4j-core:[2.0,2.17.0)</exclude>
                                </excludes>
                            </bannedDependencies>
                        </rules>
                    </configuration>
                </execution>
            </executions>
        </plugin>
    </plugins>
</build>

<!-- 传递依赖检查命令 -->
<!-- mvn dependency:tree -Dverbose -->
<!-- mvn dependency:analyze-duplicate -->
<!-- mvn enforcer:enforce -->
```

#### 12.2.4 依赖文档维护
- **检查方法**: 检查依赖文档的完整性和准确性
- **检查标准**: 维护完整的依赖文档，包括版本说明、使用原因、更新记录
- **不正确实例**:
```xml
<!-- 错误示例 - 缺少依赖文档 -->
<dependencies>
    <!-- 没有注释说明用途 -->
    <dependency>
        <groupId>com.google.guava</groupId>
        <artifactId>guava</artifactId>
        <version>32.1.2-jre</version>
    </dependency>
    
    <!-- 版本选择没有说明 -->
    <dependency>
        <groupId>org.apache.commons</groupId>
        <artifactId>commons-lang3</artifactId>
        <version>3.12.0</version>
    </dependency>
    
    <!-- 临时依赖没有标记 -->
    <dependency>
        <groupId>some.temp</groupId>
        <artifactId>temp-fix</artifactId>
        <version>1.0-SNAPSHOT</version>
    </dependency>
</dependencies>

<!-- 正确示例 - 完整的依赖文档 -->
<dependencies>
    <!-- 
        Google Guava工具库
        用途: 提供集合操作、缓存、字符串处理等工具类
        版本说明: 使用32.x版本支持Java 8+
        更新策略: 跟随最新稳定版本
        负责人: 架构组
        最后更新: 2024-01-15
    -->
    <dependency>
        <groupId>com.google.guava</groupId>
        <artifactId>guava</artifactId>
        <version>32.1.2-jre</version>
    </dependency>
    
    <!-- 
        Apache Commons Lang3
        用途: 字符串操作、数组操作、反射工具等
        版本说明: 3.12.0修复了CVE-2022-42003安全漏洞
        依赖原因: 替换自定义工具类，提高代码质量
        更新策略: 定期更新，关注安全补丁
        负责人: 开发组
        最后更新: 2024-01-10
    -->
    <dependency>
        <groupId>org.apache.commons</groupId>
        <artifactId>commons-lang3</artifactId>
        <version>3.12.0</version>
    </dependency>
    
    <!-- 
        临时修复依赖 - 待移除
        用途: 修复第三方库的bug
        版本说明: 基于原版本1.0修改
        移除计划: 等待官方版本1.1发布后移除
        跟踪Issue: JIRA-12345
        负责人: 张三
        创建时间: 2024-01-01
        预计移除: 2024-03-01
    -->
    <dependency>
        <groupId>some.temp</groupId>
        <artifactId>temp-fix</artifactId>
        <version>1.0-SNAPSHOT</version>
        <!-- TODO: 移除此临时依赖 -->
    </dependency>
</dependencies>

<!-- 依赖文档生成 -->
<build>
    <plugins>
        <!-- 生成依赖报告 -->
        <plugin>
            <groupId>org.apache.maven.plugins</groupId>
            <artifactId>maven-project-info-reports-plugin</artifactId>
            <version>3.4.5</version>
            <executions>
                <execution>
                    <id>generate-dependency-report</id>
                    <phase>site</phase>
                    <goals>
                        <goal>dependencies</goal>
                        <goal>dependency-info</goal>
                        <goal>dependency-management</goal>
                    </goals>
                </execution>
            </executions>
        </plugin>
        
        <!-- 许可证报告 -->
        <plugin>
            <groupId>org.codehaus.mojo</groupId>
            <artifactId>license-maven-plugin</artifactId>
            <version>2.2.0</version>
            <executions>
                <execution>
                    <id>generate-license-file</id>
                    <phase>generate-resources</phase>
                    <goals>
                        <goal>aggregate-add-third-party</goal>
                    </goals>
                    <configuration>
                        <outputDirectory>target/generated-sources/license</outputDirectory>
                        <thirdPartyFilename>THIRD-PARTY.txt</thirdPartyFilename>
                        <includeTransitiveDependencies>true</includeTransitiveDependencies>
                    </configuration>
                </execution>
            </executions>
        </plugin>
    </plugins>
</build>
```

```markdown
# 依赖管理文档模板

## 项目依赖清单

### 核心依赖

| 依赖名称 | 版本 | 用途 | 负责人 | 更新策略 | 最后更新 |
|---------|------|------|--------|----------|----------|
| spring-boot-starter-web | 2.7.14 | Web框架 | 架构组 | 跟随LTS版本 | 2024-01-15 |
| spring-boot-starter-data-jpa | 2.7.14 | 数据访问 | 架构组 | 跟随LTS版本 | 2024-01-15 |
| mysql-connector-java | 8.0.33 | MySQL驱动 | DBA组 | 跟随MySQL版本 | 2024-01-10 |

### 工具依赖

| 依赖名称 | 版本 | 用途 | 负责人 | 更新策略 | 最后更新 |
|---------|------|------|--------|----------|----------|
| guava | 32.1.2-jre | 工具类库 | 开发组 | 定期更新 | 2024-01-15 |
| commons-lang3 | 3.12.0 | 字符串工具 | 开发组 | 关注安全更新 | 2024-01-10 |
| jackson-databind | 2.15.2 | JSON处理 | 开发组 | 关注安全更新 | 2024-01-12 |

### 测试依赖

| 依赖名称 | 版本 | 用途 | 负责人 | 更新策略 | 最后更新 |
|---------|------|------|--------|----------|----------|
| spring-boot-starter-test | 2.7.14 | 测试框架 | 测试组 | 跟随主框架 | 2024-01-15 |
| testcontainers | 1.19.0 | 集成测试 | 测试组 | 定期更新 | 2024-01-08 |

### 临时依赖（待移除）

| 依赖名称 | 版本 | 用途 | 移除计划 | 跟踪Issue | 负责人 |
|---------|------|------|----------|-----------|--------|
| temp-fix-lib | 1.0-SNAPSHOT | 临时修复 | 2024-03-01 | JIRA-12345 | 张三 |

## 依赖更新记录

### 2024-01-15
- 升级Spring Boot从2.7.13到2.7.14
- 修复安全漏洞CVE-2023-xxxxx
- 负责人: 架构组

### 2024-01-12
- 升级jackson-databind从2.15.1到2.15.2
- 修复JSON解析安全问题
- 负责人: 开发组

### 2024-01-10
- 升级commons-lang3从3.11到3.12.0
- 添加新的字符串处理方法
- 负责人: 开发组

## 依赖审查计划

- **月度审查**: 每月第一周检查安全更新
- **季度审查**: 每季度评估依赖必要性
- **年度审查**: 每年进行大版本升级评估

## 依赖添加流程

1. 评估依赖必要性
2. 检查许可证兼容性
3. 进行安全扫描
4. 更新依赖文档
5. 代码评审确认
6. 添加到项目中

## 联系方式

- 架构组: architecture@company.com
- 开发组: development@company.com
- 测试组: testing@company.com
- DBA组: dba@company.com
```

<!-- 依赖文档生成命令 -->
<!-- mvn site -->
<!-- mvn license:aggregate-add-third-party -->
<!-- mvn project-info-reports:dependencies -->

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