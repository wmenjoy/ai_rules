# React + TypeScript + Tailwind UI/UX 设计指南

## 概述

本文档提供了一个基于 React + TypeScript + Tailwind CSS 的全面 UI/UX 设计指南。该指南结构化地提供了所有 UI 组件和设计模式的详细规范，以创建一致、可访问和可维护的设计系统。

## 思考过程总结

在设计这个 UI/UX 指南时，我考虑了以下几个关键方面：

1. **组件的精确定义** - 每个组件都需要明确的类型定义、变体、状态和尺寸
2. **技术栈的特点** - 利用 React 的组件化架构、TypeScript 的类型安全和 Tailwind 的实用优先方法
3. **设计标记系统** - 建立色彩、排版、间距等基础设计标记
4. **可访问性要求** - 确保组件符合 WCAG 标准
5. **响应式设计** - 确保所有组件在不同设备上都有良好的表现

基于以上考虑，我创建了一个全面的设计系统，包括设计标记、核心 UI 组件和布局系统。

## 待办事项列表

### 1. 定义色彩系统和 Tailwind 配置

- 创建完整的色彩系统，包括：
  - 主色（带明/暗变体）
  - 辅助色
  - 强调色
  - 中性/灰度色阶
  - 语义色（成功、警告、错误、信息）
  - 背景色
  - 文本色

- 对每种颜色：
  - 定义十六进制值
  - 创建 Tailwind 类映射
  - 编写使用指南
  - 确保文本/背景组合符合 WCAG 2.1 AA 标准

- 在 Tailwind 配置中实现：
```typescript
// Tailwind 配置示例
module.exports = {
  theme: {
    extend: {
      colors: {
        primary: {
          50: '#f0f9ff',
          100: '#e0f2fe',
          200: '#bae6fd',
          300: '#7dd3fc',
          400: '#38bdf8',
          500: '#0ea5e9',
          600: '#0284c7',
          700: '#0369a1',
          800: '#075985',
          900: '#0c4a6e',
          950: '#082f49',
        },
        // 更多颜色定义
      }
    }
  }
}
```

### 2. 建立排版系统

- 定义完整的排版系统：
  - 字体族（主要字体、次要字体、等宽字体）
  - 字重（轻、常规、中等、半粗、粗）
  - 字体大小及缩放（xs、sm、base、lg、xl、2xl 等）
  - 不同情境的行高
  - 字母间距要求
  - 文本颜色使用规范
  - 标题样式（h1-h6）
  - 段落样式
  - 特殊文本元素（引用、代码等）

- 在 Tailwind 配置中实现并创建 TypeScript 工具：
```typescript
// Tailwind 排版配置示例
module.exports = {
  theme: {
    extend: {
      fontFamily: {
        sans: ['"Inter"', 'sans-serif'],
        display: ['"Lexend"', 'sans-serif'],
        mono: ['"Fira Code"', 'monospace'],
      },
      fontSize: {
        xs: ['0.75rem', { lineHeight: '1rem' }],
        sm: ['0.875rem', { lineHeight: '1.25rem' }],
        base: ['1rem', { lineHeight: '1.5rem' }],
        lg: ['1.125rem', { lineHeight: '1.75rem' }],
        xl: ['1.25rem', { lineHeight: '1.75rem' }],
        '2xl': ['1.5rem', { lineHeight: '2rem' }],
        '3xl': ['1.875rem', { lineHeight: '2.25rem' }],
        '4xl': ['2.25rem', { lineHeight: '2.5rem' }],
        '5xl': ['3rem', { lineHeight: '1' }],
        '6xl': ['3.75rem', { lineHeight: '1' }],
      },
    }
  }
}

// TypeScript 排版工具
type TypographyVariant = 
  | 'h1' | 'h2' | 'h3' | 'h4' | 'h5' | 'h6'
  | 'body1' | 'body2' | 'caption' | 'overline';

const getTypographyClasses = (variant: TypographyVariant): string => {
  switch (variant) {
    case 'h1':
      return 'font-display text-4xl font-bold text-gray-900 dark:text-white';
    case 'h2':
      return 'font-display text-3xl font-semibold text-gray-900 dark:text-white';
    // 更多变体...
  }
};
```

### 3. 设计按钮组件系统

- 创建全面的按钮组件系统：

1. 按钮变体：
   - 主要
   - 次要
   - 轮廓
   - 幽灵/文本
   - 危险
   - 成功

2. 按钮尺寸：
   - XS（图标按钮）
   - 小
   - 中（默认）
   - 大

3. 按钮状态：
   - 默认
   - 悬停
   - 激活/按下
   - 聚焦
   - 禁用
   - 加载中

4. TypeScript 接口：
```typescript
type ButtonVariant = 'primary' | 'secondary' | 'outline' | 'ghost' | 'destructive' | 'success';
type ButtonSize = 'xs' | 'sm' | 'md' | 'lg';

interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: ButtonVariant;
  size?: ButtonSize;
  isLoading?: boolean;
  leftIcon?: React.ReactNode;
  rightIcon?: React.ReactNode;
  fullWidth?: boolean;
}
```

5. 使用 Tailwind 实现：
```tsx
// Button.tsx
import React from 'react';
import { Spinner } from './Spinner';
import { classNames } from '../utils';

export const Button: React.FC<ButtonProps> = ({ 
  children,
  variant = 'primary',
  size = 'md',
  isLoading = false,
  leftIcon,
  rightIcon,
  fullWidth = false,
  disabled,
  className,
  ...props
}) => {
  // 基础按钮类
  const baseClasses = 'inline-flex items-center justify-center font-medium rounded transition-colors focus:outline-none focus:ring-2 focus:ring-offset-2';
  
  // 尺寸类
  const sizeClasses = {
    xs: 'text-xs px-2.5 py-1.5',
    sm: 'text-sm px-3 py-2',
    md: 'text-sm px-4 py-2',
    lg: 'text-base px-6 py-3',
  };
  
  // 变体类
  const variantClasses = {
    primary: 'bg-primary-600 text-white hover:bg-primary-700 active:bg-primary-800 focus:ring-primary-500 disabled:bg-primary-300',
    secondary: 'bg-gray-600 text-white hover:bg-gray-700 active:bg-gray-800 focus:ring-gray-500 disabled:bg-gray-300',
    outline: 'border border-gray-300 text-gray-700 bg-transparent hover:bg-gray-50 active:bg-gray-100 focus:ring-gray-500 disabled:text-gray-300',
    ghost: 'text-gray-700 bg-transparent hover:bg-gray-100 active:bg-gray-200 focus:ring-gray-500 disabled:text-gray-300',
    destructive: 'bg-red-600 text-white hover:bg-red-700 active:bg-red-800 focus:ring-red-500 disabled:bg-red-300',
    success: 'bg-green-600 text-white hover:bg-green-700 active:bg-green-800 focus:ring-green-500 disabled:bg-green-300',
  };
  
  // 宽度类
  const widthClass = fullWidth ? 'w-full' : '';
  
  return (
    <button
      className={classNames(
        baseClasses,
        sizeClasses[size],
        variantClasses[variant],
        widthClass,
        isLoading && 'opacity-90 cursor-wait',
        disabled && 'cursor-not-allowed opacity-60',
        className
      )}
      disabled={disabled || isLoading}
      {...props}
    >
      {isLoading && (
        <Spinner className="mr-2 h-4 w-4" />
      )}
      {!isLoading && leftIcon && (
        <span className="mr-2">{leftIcon}</span>
      )}
      {children}
      {!isLoading && rightIcon && (
        <span className="ml-2">{rightIcon}</span>
      )}
    </button>
  );
};
```

6. 使用文档：
   - 示例使用场景
   - 最佳实践（何时使用每种变体）
   - 可访问性考虑（对比度、焦点状态）
   - 响应式行为

### 4. 开发图标系统

- 创建全面的图标系统：

1. 图标库选择：
   - 选择图标库（如 Heroicons、Phosphor Icons、Lucide）
   - 记录图标命名约定
   - 建立图标导入模式

2. 图标尺寸：
   - XS（12px）
   - 小（16px）
   - 中（20px）
   - 大（24px）
   - XL（32px）

3. 图标颜色：
   - 继承文本颜色
   - 特定语义颜色
   - 颜色继承规则

4. TypeScript 接口：
```typescript
type IconSize = 'xs' | 'sm' | 'md' | 'lg' | 'xl';

interface IconProps {
  size?: IconSize;
  color?: string;
  className?: string;
  'aria-hidden'?: boolean;
  'aria-label'?: string;
}
```

5. 使用 TypeScript 和 Tailwind 实现：
```tsx
// Icon.tsx
import React from 'react';
import { classNames } from '../utils';

export const iconSizes = {
  xs: 'w-3 h-3',
  sm: 'w-4 h-4',
  md: 'w-5 h-5',
  lg: 'w-6 h-6', 
  xl: 'w-8 h-8'
};

export const Icon: React.FC<IconProps> = ({ 
  size = 'md',
  color,
  className,
  'aria-hidden': ariaHidden = true,
  'aria-label': ariaLabel,
  children 
}) => {
  const colorClass = color ? `text-${color}` : '';
  
  const accessibilityProps = ariaLabel
    ? { 'aria-label': ariaLabel, 'aria-hidden': false }
    : { 'aria-hidden': ariaHidden };
  
  return (
    <svg
      className={classNames(
        'inline-block flex-shrink-0',
        iconSizes[size],
        colorClass,
        className
      )}
      fill="none"
      viewBox="0 0 24 24"
      stroke="currentColor"
      {...accessibilityProps}
    >
      {children}
    </svg>
  );
};

// 特定图标使用示例
export const CheckIcon: React.FC<IconProps> = (props) => (
  <Icon {...props}>
    <path 
      strokeLinecap="round" 
      strokeLinejoin="round" 
      strokeWidth={2} 
      d="M5 13l4 4L19 7" 
    />
  </Icon>
);
```

6. 图标使用指南：
   - 何时使用图标
   - 图标 + 文本组合
   - 图标密度建议
   - 仅图标按钮的可访问性要求
   - 一致性指南

7. 图标管理：
   - 自定义图标的目录结构
   - SVG 优化流程
   - 图标组件生成工作流
   - 文档中的图标搜索/浏览

### 5. 设计表单组件

- 创建全面的表单组件系统：

1. 输入字段：
   - 文本输入
   - 数字输入
   - 电子邮件输入
   - 密码输入
   - 文本区域
   - 搜索输入

2. 选择控件：
   - 复选框
   - 单选按钮
   - 切换/开关
   - 选择/下拉菜单
   - 多选
   - 自动完成

3. 表单布局组件：
   - 表单组
   - 表单部分
   - 表单网格
   - 标签
   - 帮助文本
   - 错误消息

4. 表单字段状态：
   - 默认
   - 聚焦
   - 已填充
   - 禁用
   - 只读
   - 错误
   - 成功

5. 输入的 TypeScript 接口：
```typescript
interface InputProps extends React.InputHTMLAttributes<HTMLInputElement> {
  id: string;
  label?: string;
  helpText?: string;
  error?: string;
  isInvalid?: boolean;
  isValid?: boolean;
  leftAddon?: React.ReactNode;
  rightAddon?: React.ReactNode;
  className?: string;
}
```

6. 使用 Tailwind 实现：
```tsx
// Input.tsx
import React from 'react';
import { classNames } from '../utils';

export const Input: React.FC<InputProps> = ({
  id,
  label,
  helpText,
  error,
  isInvalid = false,
  isValid = false,
  leftAddon,
  rightAddon,
  className,
  ...props
}) => {
  // 基础输入类
  const baseInputClasses = 'block w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-offset-0 focus:ring-opacity-50 transition-colors';
  
  // 状态特定类
  const stateClasses = isInvalid
    ? 'border-red-300 text-red-900 placeholder-red-300 focus:ring-red-500 focus:border-red-500'
    : isValid
      ? 'border-green-300 text-gray-900 focus:ring-green-500 focus:border-green-500'
      : 'border-gray-300 text-gray-900 placeholder-gray-400 focus:ring-primary-500 focus:border-primary-500';
  
  // 禁用状态
  const disabledClasses = props.disabled ? 'bg-gray-100 cursor-not-allowed' : 'bg-white';
  
  // 带附加组件的输入类
  const inputWithAddonClasses = leftAddon ? 'rounded-l-none' : '';
  const inputWithRightAddonClasses = rightAddon ? 'rounded-r-none' : '';
  
  return (
    <div className="w-full">
      {label && (
        <label htmlFor={id} className="block text-sm font-medium text-gray-700 mb-1">
          {label}
        </label>
      )}
      
      <div className="relative flex rounded-md shadow-sm">
        {leftAddon && (
          <span className="inline-flex items-center px-3 border border-r-0 border-gray-300 bg-gray-50 text-gray-500 rounded-l-md">
            {leftAddon}
          </span>
        )}
        
        <input
          id={id}
          className={classNames(
            baseInputClasses,
            stateClasses,
            disabledClasses,
            inputWithAddonClasses,
            inputWithRightAddonClasses,
            className
          )}
          aria-invalid={isInvalid ? 'true' : 'false'}
          aria-describedby={helpText ? `${id}-description` : undefined}
          {...props}
        />
        
        {rightAddon && (
          <span className="inline-flex items-center px-3 border border-l-0 border-gray-300 bg-gray-50 text-gray-500 rounded-r-md">
            {rightAddon}
          </span>
        )}
      </div>
      
      {helpText && !error && (
        <p id={`${id}-description`} className="mt-1 text-sm text-gray-500">
          {helpText}
        </p>
      )}
      
      {error && (
        <p className="mt-1 text-sm text-red-600">
          {error}
        </p>
      )}
    </div>
  );
};
```

7. 表单验证：
   - 客户端验证模式
   - 使用 React Hook Form 的表单状态管理
   - 错误处理和显示
   - 内联验证

8. 表单可访问性：
   - ARIA 标签和关系
   - 键盘导航
   - 错误公告
   - 焦点管理

### 6. 创建卡片和容器组件

- 设计卡片和容器组件系统：

1. 卡片变体：
   - 基础卡片
   - 交互式卡片
   - 媒体卡片
   - 统计卡片
   - 个人资料卡片
   - 操作卡片

2. 卡片属性：
   - 边框样式（宽度、半径、颜色）
   - 阴影级别
   - 内边距/间距系统
   - 内容组织

3. 容器类型：
   - 内容容器
   - 侧边栏容器
   - 全出血容器
   - 网格容器
   - 面板容器

4. 卡片的 TypeScript 接口：
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

5. 使用 Tailwind 实现：
```tsx
// Card.tsx
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
  
  // 活动状态
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
```

6. 容器实现：
```tsx
// Container.tsx
import React from 'react';
import { classNames } from '../utils';

interface ContainerProps {
  maxWidth?: 'xs' | 'sm' | 'md' | 'lg' | 'xl' | '2xl' | 'full' | 'none';
  padding?: boolean;
  centered?: boolean;
  className?: string;
  children: React.ReactNode;
}

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
```

7. 使用指南：
   - 卡片内容结构最佳实践
   - 何时使用不同的卡片变体
   - 容器组合模式
   - 容器的响应式行为
   - 嵌套指南
   - 间距建议

### 7. 开发导航组件

- 创建全面的导航组件系统：

1. 导航组件：
   - 标签页
   - 面包屑
   - 分页
   - 导航菜单
   - 下拉菜单
   - 侧边栏导航
   - 移动导航（汉堡菜单）

2. 标签组件：
   - 水平标签
   - 垂直标签
   - 下划线标签
   - 容器标签
   - 带计数器/徽章的标签

3. 标签的 TypeScript 接口：
```typescript
interface TabItemProps {
  id: string;
  label: React.ReactNode;
  icon?: React.ReactNode;
  disabled?: boolean;
  count?: number;
}

interface TabsProps {
  items: TabItemProps[];
  activeTab: string;
  onChange: (tabId: string) => void;
  variant?: 'underlined' | 'contained' | 'pills';
  size?: 'sm' | 'md' | 'lg';
  fullWidth?: boolean;
  orientation?: 'horizontal' | 'vertical';
  className?: string;
}
```

4. 使用 Tailwind 实现标签：
```tsx
// Tabs.tsx
import React from 'react';
import { classNames } from '../utils';

export const Tabs: React.FC<TabsProps> = ({
  items,
  activeTab,
  onChange,
  variant = 'underlined',
  size = 'md',
  fullWidth = false,
  orientation = 'horizontal',
  className,
}) => {
  // 省略实现细节...
};

// TabPanel 组件
interface TabPanelProps {
  id: string;
  active: boolean;
  children: React.ReactNode;
}

export const TabPanel: React.FC<TabPanelProps> = ({
  id,
  active,
  children,
}) => {
  if (!active) return null;
  
  return (
    <div
      id={`${id}-panel`}
      role="tabpanel"
      aria-labelledby={`${id}-tab`}
    >
      {children}
    </div>
  );
};
```

5. 面包屑实现：
```tsx
// Breadcrumbs.tsx
import React from 'react';
import { classNames } from '../utils';

interface BreadcrumbItemProps {
  label: React.ReactNode;
  href?: string;
  icon?: React.ReactNode;
  isCurrent?: boolean;
}

interface BreadcrumbsProps {
  items: BreadcrumbItemProps[];
  separator?: React.ReactNode;
  className?: string;
}

export const Breadcrumbs: React.FC<BreadcrumbsProps> = ({
  items,
  separator = '/',
  className,
}) => {
  // 省略实现细节...
};
```

6. 分页实现：
```tsx
// Pagination.tsx
import React from 'react';
import { classNames } from '../utils';

interface PaginationProps {
  currentPage: number;
  totalPages: number;
  onPageChange: (page: number) => void;
  siblingCount?: number;
  size?: 'sm' | 'md' | 'lg';
  className?: string;
}

export const Pagination: React.FC<PaginationProps> = ({
  currentPage,
  totalPages,
  onPageChange,
  siblingCount = 1,
  size = 'md',
  className,
}) => {
  // 省略实现细节...
};
```

7. 使用指南：
   - 导航层次结构最佳实践
   - 移动友好的导航模式
   - 导航的可访问性要求
   - 何时使用不同的导航组件
   - 响应式导航行为

### 8. 设计反馈组件

- 创建反馈组件系统：

1. 反馈组件类型：
   - 提醒/通知
   - 弹窗消息
   - 模态框/对话框
   - 工具提示
   - 弹出框
   - 进度指示器
   - 加载器/旋转器

2. 提醒组件：
   - 提醒变体（信息、成功、警告、错误）
   - 可关闭提醒
   - 带操作的提醒
   - 带图标的提醒

3. 模态组件：
```tsx
// Modal.tsx
import React, { useEffect } from 'react';
import { createPortal } from 'react-dom';
import { classNames } from '../utils';
import { XIcon } from './icons';

interface ModalProps {
  isOpen: boolean;
  onClose: () => void;
  title?: React.ReactNode;
  size?: 'sm' | 'md' | 'lg' | 'xl' | 'full';
  closeOnOverlayClick?: boolean;
  closeOnEsc?: boolean;
  initialFocus?: React.RefObject<HTMLElement>;
  className?: string;
  children: React.ReactNode;
}

// 省略实现细节...
```

4. 工具提示组件：
```tsx
// Tooltip.tsx
import React, { useState } from 'react';
import { classNames } from '../utils';

type TooltipPlacement = 'top' | 'right' | 'bottom' | 'left';

interface TooltipProps {
  content: React.ReactNode;
  placement?: TooltipPlacement;
  delay?: number;
  className?: string;
  children: React.ReactNode;
}

// 省略实现细节...
```

5. 使用指南：
   - 何时使用每种类型的反馈组件
   - 错误消息和通知的最佳实践
   - 模态使用模式
   - 弹窗通知模式
   - 加载状态模式

### 9. 建立布局系统

- 创建全面的布局系统：

1. 布局组件：
   - 网格系统
   - Flex 布局工具
   - 间距系统
   - 容器布局
   - 响应式断点
   - 堆栈（垂直/水平）
   - 分隔线

2. 网格系统：
   - 固定列网格（12列）
   - 自动响应网格
   - 网格间隙控制
   - 响应式网格行为

3. 间距系统：
   - 一致的间距比例
   - 外边距和内边距工具
   - 响应式间距
   - 自动布局间距

4. 网格的 TypeScript 接口：
```typescript
interface GridProps {
  columns?: number | { [key: string]: number }; // 响应式列计数
  gap?: number | string | { [key: string]: number | string };
  rowGap?: number | string;
  columnGap?: number | string;
  as?: React.ElementType;
  autoFlow?: 'row' | 'column' | 'dense' | 'row dense' | 'column dense';
  className?: string;
  children: React.ReactNode;
}

interface GridItemProps {
  colSpan?: number | { [key: string]: number };
  rowSpan?: number | { [key: string]: number };
  colStart?: number | { [key: string]: number };
  rowStart?: number | { [key: string]: number };
  as?: React.ElementType;
  className?: string;
  children: React.ReactNode;
}
```

5. 使用 Tailwind 实现网格：
```tsx
// Grid.tsx
import React from 'react';
import { classNames } from '../utils';

export const Grid: React.FC<GridProps> = ({
  columns = 12,
  gap,
  rowGap,
  columnGap,
  as: Component = 'div',
  autoFlow = 'row',
  className,
  children,
}) => {
  // 省略实现细节...
};

export const GridItem: React.FC<GridItemProps> = ({
  colSpan,
  rowSpan,
  colStart,
  rowStart,
  as: Component = 'div',
  className,
  children,
}) => {
  // 省略实现细节...
};
```

6. 布局系统文档：
   - 断点：
     - xs: 0px（默认）
     - sm: 640px
     - md: 768px
     - lg: 1024px
     - xl: 1280px
     - 2xl: 1536px
   
   - 容器最大宽度：
     - sm: 640px
     - md: 768px
     - lg: 1024px
     - xl: 1280px
     - 2xl: 1536px
     - full: 100%
   
   - 间距比例：
     - 基于 4px (0.25rem) 单位
     - 0: 0px
     - 1: 0.25rem (4px)
     - 2: 0.5rem (8px)
     - 3: 0.75rem (12px)
     - 4: 1rem (16px)
     - 5: 1.25rem (20px)
     - 6: 1.5rem (24px)
     - 等等
   
   - 网格系统：
     - 默认 12 列网格
     - 响应式网格配置
     - 自动流选项
   
   - 布局模式：
     - 侧边栏布局
     - 分割布局
     - 卡片网格
     - 内容-侧边栏布局
     - 仪表盘布局
     - 多级导航布局

7. 使用指南：
   - 移动优先开发
   - 何时使用 Grid vs Flex 布局
   - 处理响应式断点
   - 布局组合模式
   - 布局的性能考虑

## 实施方法

每个组件都详细提供了：
- TypeScript 接口以确保类型安全
- Tailwind CSS 实现
- 变体和自定义选项
- 可访问性考虑
- 使用指南和最佳实践

## 开发方法

设计系统将遵循原子设计方法论：
1. 原子（基本构建块，如按钮、输入）
2. 分子（原子的组合，如表单组）
3. 有机体（复杂的 UI 模式，如页眉、导航栏）
4. 模板（页面布局和结构）
5. 页面（模板的特定实例）

## 未来增强

- 主题切换功能（明/暗模式）
- 动画和过渡库
- 扩展组件变体
- 额外的可访问性改进
- 性能优化

## 下一步

1. 在 Tailwind 配置中实现设计标记
2. 开发核心组件库
3. 创建带示例的文档站点
4. 建立与现有代码库的集成
5. 为组件开发测试策略 