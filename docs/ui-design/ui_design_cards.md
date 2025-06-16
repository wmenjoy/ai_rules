# 卡片与容器组件设计指南

## 1. 概述

卡片和容器组件是现代界面设计中极其重要的元素，它们提供了视觉上的组织结构，帮助用户理解内容之间的关系。本指南详细说明了在 React + TypeScript + Tailwind CSS 项目中如何设计和实现这些组件。

## 2. 设计原则

卡片和容器组件遵循以下设计原则：

- **层次清晰**：通过阴影、边框和背景色创建清晰的视觉层次
- **内容聚焦**：容器应当突出内容，而不是容器本身
- **一致性**：保持边距、圆角和阴影的一致，创建统一的视觉语言
- **适应性**：容器应当适应不同的内容量和屏幕尺寸
- **互动反馈**：可交互的卡片应当提供明确的状态反馈

## 3. 卡片组件系统

### 3.1 卡片变体

| 变体 | 用途 | 特点 |
|------|------|------|
| 基础卡片 | 通用信息展示 | 简洁的边框和阴影 |
| 交互式卡片 | 可点击的卡片 | 悬停和点击状态 |
| 媒体卡片 | 包含图像/视频的卡片 | 优化的媒体内容展示 |
| 统计卡片 | 数据与统计信息 | 突出数字和图表 |
| 档案卡片 | 用户或产品档案 | 居中布局，突出头像/图片 |
| 操作卡片 | 包含主要操作按钮 | 强调操作区域 |

### 3.2 卡片属性

**边框样式**：
- 默认：`border border-gray-200`
- 强调：`border-2 border-primary-500`
- 无边框：仅使用阴影区分

**阴影等级**：
- 无阴影：`shadow-none`
- 轻微阴影：`shadow-sm`
- 中等阴影：`shadow-md`
- 明显阴影：`shadow-lg`

**圆角设置**：
- 无圆角：`rounded-none`
- 轻微圆角：`rounded-sm`
- 默认圆角：`rounded-md`
- 明显圆角：`rounded-lg`
- 全圆角：`rounded-xl`

**内边距系统**：
- 紧凑：`p-2`
- 标准：`p-4`
- 宽松：`p-6`
- 头部/底部：`px-4 py-3`

### 3.3 卡片组件 TypeScript 接口

```typescript
interface CardProps {
  variant?: 'basic' | 'interactive' | 'media' | 'stat' | 'profile' | 'action';
  elevation?: 'none' | 'sm' | 'md' | 'lg';
  isHoverable?: boolean;
  isSelectable?: boolean;
  isActive?: boolean;
  onClick?: () => void;
  className?: string;
  children: React.ReactNode;
}

interface CardHeaderProps {
  title?: React.ReactNode;
  subtitle?: React.ReactNode;
  action?: React.ReactNode;
  className?: string;
  children?: React.ReactNode;
}

interface CardBodyProps {
  className?: string;
  children: React.ReactNode;
}

interface CardFooterProps {
  className?: string;
  children: React.ReactNode;
}
```

### 3.4 卡片实现示例

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
import React from 'react';
import { classNames } from '../utils';

export const cardElevations = {
  none: '',
  sm: 'shadow-sm',
  md: 'shadow-md',
  lg: 'shadow-lg',
};

export const Card: React.FC<CardProps> = ({
  variant = 'basic',
  elevation = 'sm',
  isHoverable = false,
  isSelectable = false,
  isActive = false,
  onClick,
  className,
  children,
}) => {
  // 基础类
  const baseClasses = 'bg-white rounded-lg border border-gray-200 overflow-hidden';
  
  // 阴影类
  const elevationClass = cardElevations[elevation];
  
  // 交互类
  const interactiveClasses = onClick || isHoverable
    ? 'transition-all duration-200 hover:shadow-md'
    : '';
  
  // 可选择类
  const selectableClasses = isSelectable
    ? 'cursor-pointer hover:border-primary-300'
    : '';
  
  // 激活状态
  const activeClasses = isActive
    ? 'ring-2 ring-primary-500 ring-opacity-50'
    : '';
  
  // 变体特定类
  const variantClasses = {
    basic: '',
    interactive: 'cursor-pointer hover:shadow-md',
    media: '',
    stat: 'text-center p-6',
    profile: 'text-center',
    action: 'text-center p-4',
  };
  
  return (
    <div
      className={classNames(
        baseClasses,
        elevationClass,
        interactiveClasses,
        selectableClasses,
        activeClasses,
        variantClasses[variant],
        className
      )}
      onClick={onClick}
      role={onClick ? 'button' : undefined}
      tabIndex={onClick ? 0 : undefined}
    >
      {children}
    </div>
  );
};

export const CardHeader: React.FC<CardHeaderProps> = ({
  title,
  subtitle,
  action,
  className,
  children,
}) => {
  return (
    <div className={classNames('px-4 py-4 border-b border-gray-200', className)}>
      {children ? (
        children
      ) : (
        <div className="flex items-center justify-between">
          <div>
            {title && (
              <h3 className="text-lg font-medium text-gray-900">
                {title}
              </h3>
            )}
            {subtitle && (
              <p className="mt-1 text-sm text-gray-500">
                {subtitle}
              </p>
            )}
          </div>
          {action && (
            <div className="ml-4">{action}</div>
          )}
        </div>
      )}
    </div>
  );
};

export const CardBody: React.FC<CardBodyProps> = ({
  className,
  children,
}) => {
  return (
    <div className={classNames('px-4 py-4', className)}>
      {children}
    </div>
  );
};

export const CardFooter: React.FC<CardFooterProps> = ({
  className,
  children,
}) => {
  return (
    <div className={classNames('px-4 py-3 bg-gray-50 border-t border-gray-200', className)}>
      {children}
    </div>
  );
};
// [AI-BLOCK-END]
```

### 3.5 使用示例

**基础卡片**

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
<Card>
  <CardHeader 
    title="基础卡片" 
    subtitle="这是一个基础卡片示例" 
  />
  <CardBody>
    <p>卡片内容放置在这里。卡片是灵活的容器，可以包含各种类型的内容。</p>
  </CardBody>
  <CardFooter>
    <p className="text-sm text-gray-500">卡片底部信息</p>
  </CardFooter>
</Card>
// [AI-BLOCK-END]
```

**交互式卡片**

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
<Card 
  variant="interactive" 
  onClick={() => console.log('卡片被点击')}
  isHoverable
>
  <CardBody>
    <div className="flex items-center">
      <div className="flex-shrink-0">
        <svg className="h-12 w-12 text-primary-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
        </svg>
      </div>
      <div className="ml-4">
        <h4 className="text-lg font-medium">文档模板</h4>
        <p className="text-sm text-gray-600">点击查看所有可用的文档模板</p>
      </div>
    </div>
  </CardBody>
</Card>
// [AI-BLOCK-END]
```

**媒体卡片**

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
<Card variant="media">
  <img 
    src="https://images.unsplash.com/photo-1579546929518-9e396f3cc809" 
    alt="抽象渐变" 
    className="w-full h-48 object-cover"
  />
  <CardBody>
    <h3 className="font-medium text-lg">媒体标题</h3>
    <p className="mt-2 text-gray-600">这是一个包含图像的媒体卡片示例。图像显示在卡片顶部。</p>
  </CardBody>
  <CardFooter>
    <div className="flex justify-between items-center">
      <span className="text-sm text-gray-500">2023年10月15日</span>
      <button className="text-primary-600 hover:text-primary-700 text-sm font-medium">
        查看详情
      </button>
    </div>
  </CardFooter>
</Card>
// [AI-BLOCK-END]
```

**统计卡片**

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
<Card variant="stat">
  <CardBody>
    <div className="text-center">
      <p className="text-sm font-medium text-gray-500 uppercase">总销售额</p>
      <p className="mt-2 text-3xl font-bold text-gray-900">¥23,456</p>
      <div className="mt-1">
        <span className="text-green-500 text-sm font-medium">+12.5%</span>
        <span className="text-gray-500 text-sm ml-1">相比上月</span>
      </div>
    </div>
  </CardBody>
</Card>
// [AI-BLOCK-END]
```

## 4. 容器组件系统

### 4.1 容器类型

| 类型 | 用途 | 特点 |
|------|------|------|
| 内容容器 | 常规内容展示 | 居中、固定最大宽度 |
| 侧边栏容器 | 侧边栏内容 | 固定或可折叠宽度 |
| 全出血容器 | 横跨全宽的内容 | 无水平内边距限制 |
| 网格容器 | 网格布局内容 | 定义列数和间隙 |
| 面板容器 | 应用程序UI面板 | 可嵌套、固定高度 |

### 4.2 容器组件 TypeScript 接口

```typescript
interface ContainerProps {
  maxWidth?: 'xs' | 'sm' | 'md' | 'lg' | 'xl' | '2xl' | 'full' | 'none';
  padding?: boolean;
  centered?: boolean;
  className?: string;
  children: React.ReactNode;
}

// 容器最大宽度值
const maxWidthValues = {
  xs: '20rem',      // 320px
  sm: '24rem',      // 384px
  md: '28rem',      // 448px
  lg: '32rem',      // 512px
  xl: '36rem',      // 576px
  '2xl': '42rem',   // 672px
  full: '100%',
  none: 'none'
};
```

### 4.3 容器实现示例

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
import React from 'react';
import { classNames } from '../utils';

const maxWidthClasses = {
  xs: 'max-w-xs',
  sm: 'max-w-sm',
  md: 'max-w-md',
  lg: 'max-w-lg',
  xl: 'max-w-xl',
  '2xl': 'max-w-2xl',
  full: 'max-w-full',
  none: '',
};

export const Container: React.FC<ContainerProps> = ({
  maxWidth = 'lg',
  padding = true,
  centered = true,
  className,
  children,
}) => {
  return (
    <div
      className={classNames(
        maxWidthClasses[maxWidth],
        padding ? 'px-4 sm:px-6 md:px-8' : '',
        centered ? 'mx-auto' : '',
        className
      )}
    >
      {children}
    </div>
  );
};

interface SectionProps {
  className?: string;
  children: React.ReactNode;
  background?: 'white' | 'light' | 'dark' | 'primary';
  paddingY?: 'none' | 'sm' | 'md' | 'lg';
}

export const Section: React.FC<SectionProps> = ({
  className,
  children,
  background = 'white',
  paddingY = 'md',
}) => {
  // 背景类
  const backgroundClasses = {
    white: 'bg-white',
    light: 'bg-gray-50',
    dark: 'bg-gray-800 text-white',
    primary: 'bg-primary-500 text-white',
  };
  
  // 垂直内边距类
  const paddingClasses = {
    none: '',
    sm: 'py-4',
    md: 'py-8',
    lg: 'py-16',
  };
  
  return (
    <section
      className={classNames(
        backgroundClasses[background],
        paddingClasses[paddingY],
        className
      )}
    >
      {children}
    </section>
  );
};

interface PanelProps {
  className?: string;
  children: React.ReactNode;
  header?: React.ReactNode;
  footer?: React.ReactNode;
  bordered?: boolean;
  shadowed?: boolean;
}

export const Panel: React.FC<PanelProps> = ({
  className,
  children,
  header,
  footer,
  bordered = true,
  shadowed = false,
}) => {
  return (
    <div
      className={classNames(
        'rounded-lg overflow-hidden',
        bordered ? 'border border-gray-200' : '',
        shadowed ? 'shadow-md' : '',
        className
      )}
    >
      {header && (
        <div className="px-4 py-3 border-b border-gray-200 bg-gray-50">
          {header}
        </div>
      )}
      
      <div className="p-4">
        {children}
      </div>
      
      {footer && (
        <div className="px-4 py-3 border-t border-gray-200 bg-gray-50">
          {footer}
        </div>
      )}
    </div>
  );
};
// [AI-BLOCK-END]
```

### 4.4 使用示例

**基础容器**

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
<Container maxWidth="md">
  <h1 className="text-2xl font-bold">内容标题</h1>
  <p className="mt-4">
    此内容被放置在中等宽度的容器中，在较大屏幕上会有最大宽度限制，
    而在较小屏幕上则会有适当的内边距。
  </p>
</Container>
// [AI-BLOCK-END]
```

**分段结构**

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
<>
  <Section background="light">
    <Container>
      <h2 className="text-2xl font-bold text-center">浅色背景区块</h2>
      <p className="mt-4 text-center">这个区块使用浅灰色背景和标准内边距。</p>
    </Container>
  </Section>
  
  <Section background="white" paddingY="lg">
    <Container>
      <h2 className="text-2xl font-bold text-center">白色背景区块</h2>
      <p className="mt-4 text-center">这个区块使用白色背景和较大内边距。</p>
    </Container>
  </Section>
  
  <Section background="primary">
    <Container>
      <h2 className="text-2xl font-bold text-center">主色背景区块</h2>
      <p className="mt-4 text-center">这个区块使用主题色背景。</p>
    </Container>
  </Section>
</>
// [AI-BLOCK-END]
```

**面板容器**

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
<Panel
  header={<h3 className="font-medium">设置面板</h3>}
  footer={
    <div className="flex justify-end">
      <button className="px-4 py-2 bg-primary-500 text-white rounded-md">
        保存设置
      </button>
    </div>
  }
  shadowed
>
  <div className="space-y-4">
    <div>
      <label className="block text-sm font-medium text-gray-700">用户名</label>
      <input
        type="text"
        className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md"
      />
    </div>
    
    <div>
      <label className="block text-sm font-medium text-gray-700">邮箱通知</label>
      <select className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md">
        <option>每天</option>
        <option>每周</option>
        <option>从不</option>
      </select>
    </div>
  </div>
</Panel>
// [AI-BLOCK-END]
```

## 5. 布局最佳实践

### 5.1 卡片内容布局

- **内容层次**：重要信息应放在卡片顶部
- **视觉平衡**：保持内容视觉重量均衡
- **空白使用**：使用足够的内边距增强可读性
- **一致性**：相似卡片之间保持布局一致性
- **响应式考虑**：卡片内容应在各种尺寸上合理展示

### 5.2 容器嵌套规则

- **嵌套深度限制**：避免超过3层嵌套
- **内边距叠加**：嵌套容器时注意内边距叠加效应
- **响应式调整**：随着屏幕尺寸减小，减少嵌套层级
- **性能考虑**：过多的嵌套容器会影响渲染性能

### 5.3 响应式调整

| 断点 | 卡片调整 | 容器调整 |
|------|---------|---------|
| 移动端 (<640px) | 单列布局，减少内边距 | 最小内边距，最大宽度 100% |
| 平板 (640px-1024px) | 双列或三列布局 | 居中布局，适中内边距 |
| 桌面 (>1024px) | 多列布局，完整内容 | 固定最大宽度，标准内边距 |

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
// 响应式卡片网格示例
<div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
  <Card>
    <CardBody>卡片内容 1</CardBody>
  </Card>
  <Card>
    <CardBody>卡片内容 2</CardBody>
  </Card>
  <Card>
    <CardBody>卡片内容 3</CardBody>
  </Card>
</div>
// [AI-BLOCK-END]
```

### 5.4 可访问性考虑

- 交互式卡片必须有适当的键盘焦点样式
- 使用适当的 ARIA 角色和状态
- 确保颜色对比度符合 WCAG 标准
- 可点击卡片应提供清晰的互动提示

## 6. 卡片与容器组件使用场景

### 6.1 常见卡片使用场景

- **仪表盘**：统计数据卡片
- **产品列表**：产品卡片
- **用户档案**：个人信息卡片
- **内容摘要**：博客文章或新闻卡片
- **特性展示**：产品特性卡片

### 6.2 容器使用场景

- **页面布局**：使用 Container 控制内容宽度
- **内容分组**：使用 Section 划分不同内容区域
- **功能面板**：使用 Panel 展示应用程序界面

## 7. 总结与最佳实践

### 设计建议

- 保持卡片和容器样式的一致性
- 为交互式卡片提供明确的视觉反馈
- 使用卡片层次展示内容重要性
- 响应式布局中适当调整容器内边距
- 避免在容器中使用固定宽度

### 性能考虑

- 减少不必要的嵌套以优化渲染性能
- 大量卡片时考虑虚拟滚动或分页
- 使用合适的图片尺寸避免不必要的加载

### 扩展性

卡片和容器系统可通过以下方式扩展：
- 添加新的变体满足特定需求
- 实现主题化支持
- 增加动画和过渡效果
- 实现容器查询以增强响应式能力 