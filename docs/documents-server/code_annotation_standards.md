# 代码注解标准规范

**文档版本**: 1.0  
**更新日期**: 2024-05-15  
**状态**: 初稿  

## 目录

1. [概述](#1-概述)
2. [通用原则](#2-通用原则)
3. [Java 注解标准](#3-java-注解标准)
4. [REST API 注解标准](#4-rest-api-注解标准)
5. [AI 增强注解](#5-ai-增强注解)
6. [语言特定扩展](#6-语言特定扩展)
7. [代码示例](#7-代码示例)
8. [自动化验证](#8-自动化验证)

## 1. 概述

本文档定义了企业级SDK代码注解的标准规范，旨在提供一致的、机器可读的、适合AI工具理解的代码文档。标准化的注解使得文档生成过程高效且产出文档内容丰富，同时满足不同用户群体的需求。

### 1.1 目标

- 确保所有SDK代码具有完整且一致的文档注解
- 优化注解以同时满足人类开发者和AI工具的需求
- 提供明确的类型、取值范围和上下文信息
- 支持自动化的文档生成和验证流程

### 1.2 适用范围

本标准适用于公司所有SDK项目的代码开发，包括但不限于：

- Java库和框架
- RESTful API服务
- 微服务组件
- 工具包和辅助库

## 2. 通用原则

无论使用何种编程语言或框架，都应遵循以下通用原则：

### 2.1 完整性原则

每个公开API（类、方法、函数、常量等）必须包含的文档元素：

- 功能描述：清晰说明组件的用途和功能
- 参数说明：详细解释每个参数的含义、类型和约束
- 返回值：明确描述返回值的类型和含义
- 异常/错误：说明可能抛出的异常或错误情况
- 示例：提供典型用法示例

### 2.2 元数据丰富性

文档注解必须包含以下元数据：

- 版本信息：组件的当前版本
- 作者信息：开发者或团队标识
- 时间戳：创建或最后修改日期
- 弃用警告：若适用，包含替代建议
- 相关引用：关联的API或外部资源链接

### 2.3 机器可读性

为提高AI工具的理解能力：

- 使用结构化标签而非自由文本
- 标记参数类型、取值范围和约束条件
- 采用一致的术语和命名规则
- 避免歧义和模糊描述

### 2.4 上下文关联

提供足够的上下文信息：

- 在类/模块级别说明整体用途
- 描述组件在更大系统中的角色
- 注明常见的使用场景和模式
- 记录组件间的依赖关系

## 3. Java 注解标准

### 3.1 类级注解

每个公开类、接口或枚举必须包含以下Javadoc标签：

```java
/**
 * 类的主要功能描述（一句话概述）。
 * <p>
 * 详细的类描述，包括使用场景、设计意图等。
 * </p>
 *
 * @author 开发者姓名
 * @version 版本号
 * @since 引入版本
 * @apiNote 使用注意事项
 * @see 相关类的引用
 * @ai.context 此类在系统中的上下文位置和作用
 * @ai.useCases 用换行符分隔的多个使用场景
 */
```

### 3.2 方法级注解

每个公开方法必须包含以下Javadoc标签：

```java
/**
 * 方法的主要功能描述（一句话概述）。
 * <p>
 * 详细的方法描述，包括算法、复杂度、副作用等。
 * </p>
 *
 * @param paramName 参数描述，包括类型约束、取值范围
 * @return 返回值描述，包括可能的返回状态
 * @throws ExceptionType 异常描述，包括触发条件
 * @apiNote 使用注意事项
 * @implNote 实现细节说明
 * @implSpec 实现规范说明
 * @example
 * <pre>{@code
 *   // 使用示例代码
 *   Object result = method(param);
 * }</pre>
 * @ai.constraints 方法的约束条件
 * @ai.alternatives 替代方法或处理方式
 */
```

### 3.3 字段注解

公开的字段和常量应包含以下Javadoc标签：

```java
/**
 * 字段的功能描述。
 * <p>
 * 更详细的说明，包括用途和限制。
 * </p>
 *
 * @since 引入版本
 * @ai.range 值的有效范围
 * @ai.default 默认值
 */
```

## 4. REST API 注解标准

### 4.1 Spring REST控制器注解

使用Spring框架的REST API应结合使用Javadoc和SpringDoc/Swagger注解：

```java
/**
 * 控制器的功能描述。
 *
 * @author 开发者姓名
 * @version 版本号
 */
@Tag(name = "资源名称", description = "资源的详细描述")
@RestController
@RequestMapping("/api/v1/resource")
public class ResourceController {

    /**
     * 端点的功能描述。
     *
     * @param id 资源ID
     * @return 资源对象
     * @throws ResourceNotFoundException 资源不存在时
     */
    @Operation(
        summary = "操作简述",
        description = "详细操作描述，包括用例和限制",
        responses = {
            @ApiResponse(
                responseCode = "200", 
                description = "成功响应",
                content = @Content(
                    mediaType = "application/json",
                    schema = @Schema(implementation = ResourceDTO.class)
                )
            ),
            @ApiResponse(
                responseCode = "404", 
                description = "资源未找到",
                content = @Content(
                    mediaType = "application/json",
                    schema = @Schema(implementation = ErrorDTO.class)
                )
            )
        }
    )
    @GetMapping("/{id}")
    public ResponseEntity<ResourceDTO> getResource(
            @Parameter(description = "资源的唯一标识符", example = "123")
            @PathVariable Long id) {
        // 实现代码
    }
}
```

### 4.2 Model对象注解

数据模型类应包含完整的属性描述：

```java
/**
 * 资源数据传输对象。
 *
 * @author 开发者姓名
 */
@Schema(description = "资源的详细信息")
public class ResourceDTO {

    /**
     * 资源的唯一标识符。
     */
    @Schema(description = "资源的唯一标识符", example = "123")
    private Long id;

    /**
     * 资源名称。
     */
    @Schema(description = "资源名称", example = "示例资源")
    private String name;

    /**
     * 资源状态。
     */
    @Schema(
        description = "资源状态",
        example = "ACTIVE",
        allowableValues = {"ACTIVE", "INACTIVE", "PENDING"}
    )
    private String status;

    // Getter和Setter方法也应包含Javadoc
}
```

## 5. AI 增强注解

为提高文档对AI工具的友好性，添加以下自定义标签：

### 5.1 AI上下文标签

```java
/**
 * @ai.context 组件在系统中的角色和位置
 * @ai.useCases 适用场景列表，每行一个
 * @ai.constraints 使用约束和限制
 * @ai.alternatives 替代方案或组件
 * @ai.relationships 相关组件和依赖关系
 * @ai.complexity 时间/空间复杂度或性能特征
 * @ai.thread-safety 线程安全性说明
 */
```

### 5.2 AI生成代码标识

对于AI辅助生成的代码，添加标准标识：

```java
/**
 * [AI-BLOCK-START] - 生成工具: {工具名称} {版本}
 * 生成日期: {日期}
 * 贡献程度: {完全生成|部分生成|基于模板}
 * 人工修改: {有|无} 
 * 责任人: {人员名称}
 */
```

## 6. 语言特定扩展

### 6.1 Go语言注解标准

Go代码应遵循以下注释规范：

```go
// PackageName 提供了...（包的简要描述）
//
// 包的详细描述，包括主要用途和功能。
// 可以使用多行来提供完整信息。
//
// AI-Context: 包在系统中的位置和作用
// AI-UseCases:
// - 用例1
// - 用例2
package packagename

// FunctionName 执行...（函数的简要描述）
//
// 函数的详细描述，包括算法、复杂度等。
//
// 参数:
//   - param1: 参数1的描述，包括类型约束
//   - param2: 参数2的描述
//
// 返回:
//   - 返回值的描述
//
// 错误:
//   - ErrType1: 错误情况1的描述
//   - ErrType2: 错误情况2的描述
//
// 示例:
//   result, err := FunctionName(param1, param2)
//
// AI-Constraints: 函数的约束条件
// AI-Alternatives: 替代函数或处理方式
func FunctionName(param1 Type1, param2 Type2) (ReturnType, error) {
    // 实现代码
}
```

## 7. 代码示例

### 7.1 Java类示例

```java
package com.example.project.util;

import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.Date;

/**
 * 日期处理工具类，提供日期格式化、解析和操作功能。
 * <p>
 * 本工具类封装了Java标准日期API的常用操作，提供更简洁的调用方式，
 * 支持常见的日期格式处理，适用于日期显示和存储场景。
 * </p>
 *
 * @author 开发团队
 * @version 1.2.0
 * @since 1.0.0
 * @apiNote 该类中的所有方法均为线程安全的静态方法
 * @see java.util.Date
 * @see java.text.SimpleDateFormat
 * @ai.context 基础工具集，用于应用层日期处理
 * @ai.useCases 
 *     用户输入日期解析
 *     数据库日期格式转换
 *     日期显示格式化
 */
public final class DateUtils {

    /**
     * 默认日期格式: yyyy-MM-dd HH:mm:ss。
     * <p>
     * 系统全局通用的标准日期时间格式，在未指定格式时使用。
     * </p>
     * 
     * @since 1.0.0
     * @ai.default "yyyy-MM-dd HH:mm:ss"
     */
    public static final String DEFAULT_DATE_FORMAT = "yyyy-MM-dd HH:mm:ss";

    /**
     * 私有构造函数，防止实例化。
     * <p>
     * 该类仅提供静态方法，不应被实例化。
     * </p>
     */
    private DateUtils() {
        // 防止实例化
    }

    /**
     * 将日期对象格式化为指定格式的字符串。
     * <p>
     * 将Java日期对象转换为可读性强的字符串表示形式，
     * 根据提供的格式模式进行格式化。
     * </p>
     *
     * @param date   日期对象，不能为null
     * @param format 日期格式模式，例如 "yyyy-MM-dd"
     * @return 格式化后的日期字符串，如果输入参数为null则返回null
     * @apiNote 如果需要频繁调用，建议缓存SimpleDateFormat实例提高性能
     * @example
     * <pre>{@code
     *   Date now = new Date();
     *   String formatted = DateUtils.formatDate(now, "yyyy-MM-dd");
     *   // 返回如 "2024-05-15"
     * }</pre>
     * @ai.constraints 非线程安全，内部创建新的SimpleDateFormat实例
     * @ai.alternatives 考虑使用java.time.DateTimeFormatter替代
     */
    public static String formatDate(Date date, String format) {
        if (date == null || format == null || format.isEmpty()) {
            return null;
        }
        SimpleDateFormat sdf = new SimpleDateFormat(format);
        return sdf.format(date);
    }

    // 其他方法...
}
```

### 7.2 REST API控制器示例

```java
package com.example.project.controller;

import com.example.project.dto.UserDTO;
import com.example.project.exception.UserNotFoundException;
import com.example.project.service.UserService;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.Parameter;
import io.swagger.v3.oas.annotations.media.Content;
import io.swagger.v3.oas.annotations.media.Schema;
import io.swagger.v3.oas.annotations.responses.ApiResponse;
import io.swagger.v3.oas.annotations.tags.Tag;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

/**
 * 用户资源管理控制器，提供用户信息的CRUD操作。
 * <p>
 * 该控制器处理所有与用户相关的HTTP请求，包括查询、创建、
 * 更新和删除用户信息。所有操作均需要适当的认证和授权。
 * </p>
 *
 * @author API团队
 * @version 2.0.0
 * @since 1.0.0
 * @ai.context 用户管理子系统的HTTP接口层
 * @ai.useCases 
 *     用户管理界面
 *     用户自助服务
 *     第三方系统集成
 */
@Tag(name = "用户管理", description = "用户资源的CRUD操作")
@RestController
@RequestMapping("/api/v1/users")
public class UserController {

    private final UserService userService;

    /**
     * 构造函数，通过依赖注入获取用户服务。
     *
     * @param userService 用户服务实例
     */
    public UserController(UserService userService) {
        this.userService = userService;
    }

    /**
     * 根据ID获取用户信息。
     * <p>
     * 查询指定ID的用户详细信息，如果用户不存在则返回404错误。
     * </p>
     *
     * @param id 用户ID
     * @return 包含用户信息的HTTP响应
     * @throws UserNotFoundException 当指定ID的用户不存在时抛出
     * @ai.constraints 需要USER_READ权限
     */
    @Operation(
        summary = "获取用户信息",
        description = "根据用户ID查询用户的详细信息",
        responses = {
            @ApiResponse(
                responseCode = "200",
                description = "成功获取用户信息",
                content = @Content(
                    mediaType = "application/json",
                    schema = @Schema(implementation = UserDTO.class)
                )
            ),
            @ApiResponse(
                responseCode = "404",
                description = "用户不存在"
            )
        }
    )
    @GetMapping("/{id}")
    public ResponseEntity<UserDTO> getUser(
            @Parameter(description = "用户的唯一标识符", example = "123")
            @PathVariable Long id) {
        UserDTO user = userService.getUserById(id);
        return ResponseEntity.ok(user);
    }

    // 其他CRUD操作...
}
```

## 8. 自动化验证

### 8.1 Checkstyle配置

使用以下Checkstyle规则验证Javadoc注解的完整性：

```xml
<module name="JavadocType">
    <property name="severity" value="error"/>
    <property name="scope" value="public"/>
    <property name="allowMissingParamTags" value="false"/>
    <property name="allowUnknownTags" value="true"/>
    <property name="tokens" value="INTERFACE_DEF, CLASS_DEF, ENUM_DEF, ANNOTATION_DEF"/>
</module>
<module name="JavadocMethod">
    <property name="severity" value="error"/>
    <property name="scope" value="public"/>
    <property name="allowMissingParamTags" value="false"/>
    <property name="allowMissingReturnTag" value="false"/>
</module>
```

### 8.2 OpenAPI验证

使用SpringDoc注解验证器确保API文档的完整性：

```java
@Bean
public OpenApiCustomiser openApiCustomiser() {
    return openApi -> {
        // 验证所有操作都有描述
        openApi.getPaths().values().stream()
            .flatMap(pathItem -> pathItem.readOperations().stream())
            .filter(operation -> StringUtils.isEmpty(operation.getDescription()))
            .forEach(operation -> {
                // 记录缺少描述的操作
                log.warn("Operation {} is missing description", operation.getOperationId());
            });
    };
}
``` 