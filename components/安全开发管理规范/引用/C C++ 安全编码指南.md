> 编 号：
>
> 版 本 号 ：
>
> 受控状态：
>
> 密 级：
>
> **【组织名称】**
>
> **C/C++ 安全编码指南**
>
> **版权声明和保密须知**
>
> 本文件中出现的任何文字叙述、文档格式、插图、照片、方法、过程等内容，除另有特别注明，版权均属【组织名称】所有，受到有关产权及版权法保护。任何单位和个人未经【组织名称】的书面授权许可，不得复制或引用本文件的任何片断，无论通过电子形式或非电子形式。
>
> **Copyright © 2022【组织名称】 版权所有**
>
> 文档信息
>
> 文档编号：
>
> 文档分类：
>
> 编写：
>
> 审核：
>
> 批准：
>
> 初次发布日期：
>
> 生效日期：
>
> 修订日期：
>
> 版本记录
>
> **版本号**
>
> **版本日期**
>
> **修改**
>
> **审批人**
>
> **修改履历**
>
> 目 录
>
> TOC \\o \"1-3\" \\h \\u HYPERLINK \\l \"\_Toc207802014\"
> [1、前言]{.underline} PAGEREF \_Toc207802014 \\h 5
>
> HYPERLINK \\l \"\_Toc207802015\" [1.1目的]{.underline} PAGEREF
> \_Toc207802015 \\h 5
>
> HYPERLINK \\l \"\_Toc207802016\" [1.2适用范围]{.underline} PAGEREF
> \_Toc207802016 \\h 5
>
> HYPERLINK \\l \"\_Toc207802017\" [2、通用安全编码规范]{.underline}
> PAGEREF \_Toc207802017 \\h 5
>
> HYPERLINK \\l \"\_Toc207802018\" [2.1内存安全基础]{.underline} PAGEREF
> \_Toc207802018 \\h 5
>
> HYPERLINK \\l \"\_Toc207802019\"
> [2.1.1禁止使用未初始化变量]{.underline} PAGEREF \_Toc207802019 \\h 5
>
> HYPERLINK \\l \"\_Toc207802020\"
> [2.1.2强制启用安全编译选项]{.underline} PAGEREF \_Toc207802020 \\h 6
>
> HYPERLINK \\l \"\_Toc207802021\" [2.2类型安全规范]{.underline} PAGEREF
> \_Toc207802021 \\h 6
>
> HYPERLINK \\l \"\_Toc207802022\" [2.2.1禁止随意类型转换]{.underline}
> PAGEREF \_Toc207802022 \\h 6
>
> HYPERLINK \\l \"\_Toc207802023\"
> [2.2.2使用固定大小数据类型]{.underline} PAGEREF \_Toc207802023 \\h 7
>
> HYPERLINK \\l \"\_Toc207802024\" [2.3宏定义安全]{.underline} PAGEREF
> \_Toc207802024 \\h 7
>
> HYPERLINK \\l \"\_Toc207802025\"
> [2.3.1宏参数加括号避免展开副作用]{.underline} PAGEREF \_Toc207802025
> \\h 7
>
> HYPERLINK \\l \"\_Toc207802026\"
> [2.3.2禁止宏定义多语句无代码块]{.underline} PAGEREF \_Toc207802026 \\h
> 8
>
> HYPERLINK \\l \"\_Toc207802027\" [3、内存安全专项规范]{.underline}
> PAGEREF \_Toc207802027 \\h 9
>
> HYPERLINK \\l \"\_Toc207802028\" [3.1动态内存管理]{.underline} PAGEREF
> \_Toc207802028 \\h 9
>
> HYPERLINK \\l \"\_Toc207802029\" [3.1.1 malloc/free
> 配对，禁止重复释放或野指针释放]{.underline} PAGEREF \_Toc207802029 \\h
> 9
>
> HYPERLINK \\l \"\_Toc207802030\" [3.1.2 C++
> 优先使用智能指针，禁用裸指针管理动态内存]{.underline} PAGEREF
> \_Toc207802030 \\h 9
>
> HYPERLINK \\l \"\_Toc207802031\" [3.2栈缓冲区溢出防护]{.underline}
> PAGEREF \_Toc207802031 \\h 10
>
> HYPERLINK \\l \"\_Toc207802032\"
> [3.2.1禁止使用无边界检查的字符串函数]{.underline} PAGEREF
> \_Toc207802032 \\h 10
>
> HYPERLINK \\l \"\_Toc207802033\" [3.2.2启用栈保护编译选项]{.underline}
> PAGEREF \_Toc207802033 \\h 11
>
> HYPERLINK \\l \"\_Toc207802034\" [3.3指针安全操作]{.underline} PAGEREF
> \_Toc207802034 \\h 11
>
> HYPERLINK \\l \"\_Toc207802035\" [3.3.1禁止空指针解引用]{.underline}
> PAGEREF \_Toc207802035 \\h 11
>
> HYPERLINK \\l \"\_Toc207802036\" [3.3.2禁止指针越界访问]{.underline}
> PAGEREF \_Toc207802036 \\h 12
>
> HYPERLINK \\l \"\_Toc207802037\" [4、输入验证与过滤规范]{.underline}
> PAGEREF \_Toc207802037 \\h 13
>
> HYPERLINK \\l \"\_Toc207802038\" [4.1]{.underline}
> [命令行参数验证]{.underline} PAGEREF \_Toc207802038 \\h 13
>
> HYPERLINK \\l \"\_Toc207802039\" [4.2]{.underline}
> [文件输入验证]{.underline} PAGEREF \_Toc207802039 \\h 15
>
> HYPERLINK \\l \"\_Toc207802040\" [5、数据类型与溢出防护]{.underline}
> PAGEREF \_Toc207802040 \\h 16
>
> HYPERLINK \\l \"\_Toc207802041\" [5.1]{.underline}
> [整数溢出防护]{.underline} PAGEREF \_Toc207802041 \\h 16
>
> HYPERLINK \\l \"\_Toc207802042\" [5.2]{.underline}
> [枚举类型安全使用]{.underline} PAGEREF \_Toc207802042 \\h 17
>
> HYPERLINK \\l \"\_Toc207802043\" [6、风险与应对措施]{.underline}
> PAGEREF \_Toc207802043 \\h 18
>
> HYPERLINK \\l \"\_Toc207802044\" [6.1共享数据加锁保护]{.underline}
> PAGEREF \_Toc207802044 \\h 18
>
> HYPERLINK \\l \"\_Toc207802045\" [6.2禁止线程退出时持有锁]{.underline}
> PAGEREF \_Toc207802045 \\h 19
>
> HYPERLINK \\l \"\_Toc207802046\" [7、执行与监督]{.underline} PAGEREF
> \_Toc207802046 \\h 21
>
> HYPERLINK \\l \"\_Toc207802047\" [7．1代码审查重点]{.underline}
> PAGEREF \_Toc207802047 \\h 21
>
> HYPERLINK \\l \"\_Toc207802048\" [7.2安全测试工具链]{.underline}
> PAGEREF \_Toc207802048 \\h 21
>
> HYPERLINK \\l \"\_Toc207802049\" [7.3培训与更新]{.underline} PAGEREF
> \_Toc207802049 \\h 21
>
> HYPERLINK \\l \"\_Toc207802050\" [8、附录]{.underline} PAGEREF
> \_Toc207802050 \\h 22
>
> HYPERLINK \\l \"\_Toc207802051\" [8.1 常用 C/C++ 安全库]{.underline}
> PAGEREF \_Toc207802051 \\h 22
>
> HYPERLINK \\l \"\_Toc207802052\" [8.2 安全编译选项汇总]{.underline}
> PAGEREF \_Toc207802052 \\h 22
>
> **前言**
>
> **1.1目的**
>
> 针对 C/C++
> 语言无自动内存管理、指针操作灵活等特性，规范开发人员编码行为，防范缓冲区溢出、整数溢出、空指针解引用、内存泄漏等典型安全漏洞，保障嵌入式系统、操作系统内核、高性能服务等
> C/C++ 项目的稳定性、安全性及合规性（如满足 ISO/IEC 17961、CERT C/C++
> 等标准）。
>
> **1.2适用范围**
>
> 适用于公司所有 C/C++
> 开发项目（含嵌入式软件、服务器端程序、桌面应用、驱动程序等）；覆盖所有
> C/C++ 开发人员（初级 / 中级 /
> 高级工程师）、代码审查人员、测试工程师及架构设计人员；支持主流编译环境（GCC、Clang、MSVC）。
>
> **2、通用安全编码规范**
>
> **2.1内存安全基础**
>
> **2.1.1禁止使用未初始化变量**
>
> 未初始化变量的值为随机值，可能导致程序逻辑异常或信息泄露，所有变量（含栈变量、全局变量）必须显式初始化。
>
> **2.1.2强制启用安全编译选项**
>
> 通过编译器选项增强内存安全检查，不同编译环境推荐配置如下：

编译环境

安全选项

作用

GCC/Clang

-Wall -Wextra -Werror

开启所有警告，将警告视为错误（避免忽略潜在风险）

GCC/Clang

-fstack-protector-strong

启用栈保护（防范栈缓冲区溢出）

GCC/Clang

-fPIE -pie

生成位置无关可执行文件（防范地址泄漏导致的攻击）

MSVC

/W4 /WX

开启最高级别警告，将警告视为错误

MSVC

/GS

启用栈安全检查

> **2.2类型安全规范**
>
> **2.2.1禁止随意类型转换**
>
> 避免无意义的类型转换（尤其是指针类型转换），防止内存访问越界；必须转换时，使用static_cast（编译时检查）或dynamic_cast（运行时检查，仅
> C++），禁止使用reinterpret_cast（无类型检查）。
>
> **2.2.2使用固定大小数据类型**
>
> 避免使用int、long等依赖平台的变量类型（如 32 位系统int为 4 字节，16
> 位系统为 2
> 字节），优先使用\<stdint.h\>（C）/\<cstdint\>（C++）中的固定大小类型（如int32_t、uint64_t），确保跨平台一致性。
>
> **2.3宏定义安全**
>
> **2.3.1宏参数加括号避免展开副作用**
>
> 宏展开为文本替换，若参数含表达式，未加括号会导致运算优先级错误，所有宏参数及整体表达式必须加括号。
>
> **2.3.2禁止宏定义多语句无代码块**
>
> **3、内存安全专项规范**
>
> **3.1动态内存管理**
>
> **3.1.1 malloc/free 配对，禁止重复释放或野指针释放**
>
> 动态内存必须遵循 "申请 - 使用 - 释放"
> 闭环，malloc后检查是否为NULL（防止内存分配失败），free后将指针置为NULL（避免悬垂指针），禁止释放非动态内存或已释放内存。
>
> **3.1.2 C++ 优先使用智能指针，禁用裸指针管理动态内存**
>
> C++
> 项目中，使用std::unique_ptr（独占所有权）、std::shared_ptr（共享所有权）等智能指针自动管理内存，避免手动new/delete导致的内存泄漏或重复释放。
>
> **3.2栈缓冲区溢出防护**
>
> **3.2.1禁止使用无边界检查的字符串函数**
>
> strcpy、strcat、gets等函数无缓冲区边界检查，易导致栈溢出，替换为带长度限制的安全函数（如strncpy、strncat、fgets），且需显式检查长度。
>
> **3.2.2启用栈保护编译选项**
>
> 编译时启用栈保护（GCC/Clang：-fstack-protector-strong，MSVC：/GS），当栈缓冲区溢出时，程序会触发崩溃（而非继续执行恶意代码），配合核心转储可定位漏洞。
>
> **3.3指针安全操作**
>
> **3.3.1禁止空指针解引用**
>
> 所有指针使用前必须检查是否为NULL，尤其是函数返回的指针（如fopen、malloc），避免空指针解引用导致程序崩溃。
>
> **3.3.2禁止指针越界访问**
>
> 指针访问数组或缓冲区时，必须确保索引在合法范围内（0 ≤ 索引 \<
> 长度），禁止通过指针算术运算越界。
>
> **4、输入验证与过滤规范**
>
> **命令行参数验证**
>
> 所有命令行参数（argc/argv）必须验证，包括参数数量、参数格式（如是否为数字、路径是否合法），禁止直接使用未验证的参数。
>
> **文件输入验证**
>
> 读取文件时，需验证文件路径合法性（防止路径遍历攻击）、文件大小（避免超大文件导致内存耗尽）、文件内容格式（如是否为预期的二进制结构或文本格式）。
>
> **5、数据类型与溢出防护**
>
> **整数溢出防护**
>
> 整数运算前必须检查是否会溢出（如加法前检查a + b \>
> INT_MAX），或使用安全算术库（如libsafec），禁止忽略溢出导致的逻辑错误或内存越界。
>
> **枚举类型安全使用**
>
> C++ 中使用enum class（强类型枚举）替代 C
> 风格枚举，避免枚举值隐式转换为整数导致的逻辑错误；C
> 语言中需显式检查枚举值合法性。
>
> **6、风险与应对措施**
>
> **6.1共享数据加锁保护**
>
> 多线程访问共享数据（如全局变量、堆内存）时，必须使用互斥锁（pthread_mutex_t（C）、std::mutex（C++））或读写锁，禁止无保护的并发访问导致数据竞争。
>
> **6.2禁止线程退出时持有锁**
>
> 线程退出（如pthread_exit、函数返回）前必须释放已持有的锁，避免其他线程永久阻塞（死锁）。
>
> **7、执行与监督**
>
> **7.1代码审查重点**
>
> 代码审查需聚焦 C/C++ 特有安全风险，审查清单包括：
>
> 内存管理：malloc/free/new/delete是否配对，智能指针使用是否正确；
>
> 指针操作：是否存在空指针解引用、越界访问；
>
> 缓冲区操作：是否使用无边界检查函数（如strcpy），长度计算是否正确；
>
> 整数运算：是否检查溢出，固定大小类型使用是否规范；
>
> 线程：共享数据是否加锁，锁是否正确释放；
>
> 编译选项：是否启用栈保护、警告即错误等安全选项。
>
> **7.2安全测试工具链**
>
> 工具类型
>
> 工具名称
>
> 用途
>
> 使用示例

静态代码分析

Cppcheck

检测内存泄漏、空指针、缓冲区溢出等

cppcheck \--enable=all app.c

静态代码分析

Clang Static Analyzer

编译器内置分析，检测未定义行为

scan-build gcc -o app app.c

动态内存检测

Valgrind (Memcheck)

检测内存泄漏、使用已释放内存、越界访问

valgrind \--leak-check=full ./app

运行时安全检测

AddressSanitizer

快速检测内存错误（比 Valgrind 快）

gcc -o app app.c -fsanitize=address -g

代码合规检查

CERT C Checker

检查是否符合 CERT C 安全标准

集成到 Jenkins 流水线，自动扫描代码

> **7.3培训与更新**
>
> 每季度组织 C/C++
> 安全编码培训，重点讲解内存安全、多线程安全等高频漏洞案例（如
> Heartbleed 漏洞的缓冲区溢出问题）；
>
> 技术部门跟踪 C/C++ 标准更新（如 C23、C++23
> 的安全特性）及编译器安全选项升级，及时更新本指南；
>
> 每年对存量项目进行安全审计，使用工具扫描并修复历史代码中的安全漏洞。
>
> **8、附录**
>
> **8.1 常用 C/C++ 安全库**
>
> 库名称
>
> 用途
>
> 适用语言
>
> 安装命令（Linux）

libsafec

提供安全的字符串 / 内存操作函数（如 strlcpy）

> C

sudo apt install libsafec-dev

Safe C++C++ 安全容器与算法库（防溢出、越界）C++源码编译：HYPERLINK
\"https://github.com/dcleblanc/SafeC++\"[https://github.com/dcleblanc/SafeC++]{.underline}

ThreadSanitizer

检测多线程数据竞争

> C/C++

编译器内置（GCC/Clang），编译时加-fsanitize=thread

> **8.2 安全编译选项汇总**
>
> 编译环境
>
> 选项
>
> 作用
>
> GCC/Clang

-Wall -Wextra -Werror

开启所有警告，将警告视为错误

> GCC/Clang

-fstack-protector-strong

启用栈保护，防范栈缓冲区溢出

> GCC/Clang

-fsanitize=address

运行时检测内存错误（开发环境用）

> GCC/Clang

-fPIE -pie

生成位置无关可执行文件，防范地址泄漏

> MSVC

/GS

启用栈安全检查

> MSVC

/W4 /WX

最高级别警告，警告视为错误

> 所有环境

-D_FORTIFY_SOURCE=2

增强标准库函数（如 strcpy）的边界检查

> PAGE 1
>
> 34
>
> 第
>
> 共
>
> 页
>
> PAGE 24
>
> 34
>
> 第
>
> 页
>
> 共
>
> 页
