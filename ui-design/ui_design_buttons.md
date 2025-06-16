# 按钮组件设计指南

## 概述

按钮是用户界面中最常见和最重要的交互元素之一，它们表明用户可以采取的行动并直接影响用户体验。本文档详细定义了基于React、TypeScript和Tailwind CSS的按钮组件系统，包括变体、尺寸、状态和实现方式。

## 按钮变体

按钮变体根据其视觉重要性和应用场景进行分类，确保界面中的视觉层次一致。

### 主要按钮 (Primary)

用于页面中最主要或最推荐的操作，如"提交"、"保存"、"确认"等。

- 外观：填充背景色、白色文字
- 使用场景：表单提交、主要操作、引导流程中的下一步
- Tailwind类: `bg-primary-600 hover:bg-primary-700 active:bg-primary-800 text-white focus:ring-2 focus:ring-primary-500 focus:ring-offset-2`

### 次要按钮 (Secondary)

用于支持主要操作但不需要特别强调的操作，如"取消"、"上一步"等。

- 外观：灰色填充背景、白色文字
- 使用场景：次要操作、辅助功能
- Tailwind类: `bg-gray-600 hover:bg-gray-700 active:bg-gray-800 text-white focus:ring-2 focus:ring-gray-500 focus:ring-offset-2`

### 轮廓按钮 (Outline)

用于需要视觉轻量但依然重要的操作。

- 外观：透明背景、有边框、文字颜色同边框
- 使用场景：过滤器、分类选择、次要操作
- Tailwind类: `border border-gray-300 text-gray-700 bg-transparent hover:bg-gray-50 active:bg-gray-100 focus:ring-2 focus:ring-gray-500 focus:ring-offset-2`

### 幽灵按钮 (Ghost)

最轻量级的按钮，几乎没有视觉元素，只在交互时显示反馈。

- 外观：透明背景、无边框、只有文字
- 使用场景：工具栏、紧凑界面、辅助操作
- Tailwind类: `text-gray-700 bg-transparent hover:bg-gray-100 active:bg-gray-200 focus:ring-2 focus:ring-gray-500 focus:ring-offset-2`

### 危险按钮 (Destructive)

用于警告用户执行破坏性操作，如删除、清空等。

- 外观：红色背景、白色文字
- 使用场景：删除、移除、清空数据等危险操作
- Tailwind类: `bg-red-600 hover:bg-red-700 active:bg-red-800 text-white focus:ring-2 focus:ring-red-500 focus:ring-offset-2`

### 成功按钮 (Success)

用于表示操作成功或确认积极行为。

- 外观：绿色背景、白色文字
- 使用场景：确认、完成、提交等积极操作
- Tailwind类: `bg-green-600 hover:bg-green-700 active:bg-green-800 text-white focus:ring-2 focus:ring-green-500 focus:ring-offset-2`

## 按钮尺寸

提供多种尺寸适应不同UI场景需求。

### 超小型 (XS)

- 高度: 24px (1.5rem)
- 内边距: 横向10px (px-2.5)，纵向6px (py-1.5)
- 字体大小: 12px (text-xs)
- 使用场景: 图标按钮、紧凑型UI、表格内的按钮
- Tailwind类: `text-xs px-2.5 py-1.5`

### 小型 (SM)

- 高度: 32px (2rem)
- 内边距: 横向12px (px-3)，纵向8px (py-2)
- 字体大小: 14px (text-sm)
- 使用场景: 表单内操作、卡片内按钮、辅助按钮
- Tailwind类: `text-sm px-3 py-2`

### 中型 (MD) - 默认

- 高度: 40px (2.5rem)
- 内边距: 横向16px (px-4)，纵向8px (py-2)
- 字体大小: 14px (text-sm)
- 使用场景: 大多数标准按钮、表单提交按钮
- Tailwind类: `text-sm px-4 py-2`

### 大型 (LG)

- 高度: 48px (3rem)
- 内边距: 横向24px (px-6)，纵向12px (py-3)
- 字体大小: 16px (text-base)
- 使用场景: 主要行动点、营销页面按钮、突出显示的按钮
- Tailwind类: `text-base px-6 py-3`

## 按钮状态

每个按钮都具有以下状态，以提供用户交互反馈。

### 默认状态

按钮的基本状态，没有任何用户交互。

### 悬停状态 (Hover)

当用户鼠标悬停在按钮上时的状态，通常加深背景色。

- Tailwind类: `hover:bg-primary-700` (基于变体有不同实现)

### 激活状态 (Active/Pressed)

当用户点击按钮时的状态，通常更加深背景色。

- Tailwind类: `active:bg-primary-800` (基于变体有不同实现)

### 聚焦状态 (Focus)

当按钮通过键盘导航获得焦点时的状态，通常添加环形高亮。

- Tailwind类: `focus:outline-none focus:ring-2 focus:ring-primary-500 focus:ring-offset-2`

### 禁用状态 (Disabled)

当按钮不可用或操作不可执行时的状态，通常降低透明度并改变鼠标光标。

- Tailwind类: `disabled:opacity-50 disabled:cursor-not-allowed disabled:bg-opacity-80`

### 加载状态 (Loading)

当按钮触发的操作正在进行中，通常显示加载指示器并禁用按钮。

- 实现方式: 添加旋转器图标并设置`disabled`属性

## TypeScript接口

```typescript
type ButtonVariant = 'primary' | 'secondary' | 'outline' | 'ghost' | 'destructive' | 'success';
type ButtonSize = 'xs' | 'sm' | 'md' | 'lg';

interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  /** 按钮变体类型 */
  variant?: ButtonVariant;
  /** 按钮尺寸 */
  size?: ButtonSize;
  /** 是否处于加载状态 */
  isLoading?: boolean;
  /** 左侧图标 */
  leftIcon?: React.ReactNode;
  /** 右侧图标 */
  rightIcon?: React.ReactNode;
  /** 是否铺满容器宽度 */
  fullWidth?: boolean;
  /** 自定义类名 */
  className?: string;
  /** 子元素 */
  children: React.ReactNode;
}
```

## Tailwind实现

以下是结合React和TypeScript的完整Button组件实现：

```tsx
// Button.tsx
import React from 'react';
import { Spinner } from './Spinner'; // 加载指示器组件
import { classNames } from '../utils'; // 工具函数合并className

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
    xs: 'text-xs px-2.5 py-1.5 rounded',
    sm: 'text-sm px-3 py-2 rounded',
    md: 'text-sm px-4 py-2 rounded-md',
    lg: 'text-base px-6 py-3 rounded-md',
  };
  
  // 变体类
  const variantClasses = {
    primary: 'bg-primary-600 text-white hover:bg-primary-700 active:bg-primary-800 focus:ring-primary-500 disabled:bg-primary-300',
    secondary: 'bg-gray-600 text-white hover:bg-gray-700 active:bg-gray-800 focus:ring-gray-500 disabled:bg-gray-300',
    outline: 'border border-gray-300 text-gray-700 bg-transparent hover:bg-gray-50 active:bg-gray-100 focus:ring-gray-500 disabled:text-gray-300 disabled:border-gray-200',
    ghost: 'text-gray-700 bg-transparent hover:bg-gray-100 active:bg-gray-200 focus:ring-gray-500 disabled:text-gray-300',
    destructive: 'bg-red-600 text-white hover:bg-red-700 active:bg-red-800 focus:ring-red-500 disabled:bg-red-300',
    success: 'bg-green-600 text-white hover:bg-green-700 active:bg-green-800 focus:ring-green-500 disabled:bg-green-300',
  };
  
  // 宽度类
  const widthClass = fullWidth ? 'w-full' : '';
  
  // 加载和禁用状态
  const isDisabled = disabled || isLoading;
  
  return (
    <button
      className={classNames(
        baseClasses,
        sizeClasses[size],
        variantClasses[variant],
        widthClass,
        isLoading && 'opacity-90 cursor-wait',
        isDisabled && 'cursor-not-allowed opacity-60',
        className
      )}
      disabled={isDisabled}
      aria-disabled={isDisabled}
      {...props}
    >
      {isLoading && (
        <Spinner className="mr-2 h-4 w-4" aria-hidden="true" />
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

## 辅助组件: 旋转器 (Spinner)

用于在按钮加载状态显示的加载指示器：

```tsx
// Spinner.tsx
import React from 'react';
import { classNames } from '../utils';

interface SpinnerProps {
  className?: string;
  size?: 'sm' | 'md' | 'lg';
}

export const Spinner: React.FC<SpinnerProps> = ({ 
  className = '',
  size = 'md'
}) => {
  const sizeClasses = {
    sm: 'h-4 w-4',
    md: 'h-5 w-5',
    lg: 'h-6 w-6',
  };

  return (
    <svg
      className={classNames(
        'animate-spin text-current',
        sizeClasses[size],
        className
      )}
      xmlns="http://www.w3.org/2000/svg"
      fill="none"
      viewBox="0 0 24 24"
      aria-hidden="true"
    >
      <circle
        className="opacity-25"
        cx="12"
        cy="12"
        r="10"
        stroke="currentColor"
        strokeWidth="4"
      ></circle>
      <path
        className="opacity-75"
        fill="currentColor"
        d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
      ></path>
    </svg>
  );
};
```

## 常见按钮组合

### 图标按钮

```tsx
<Button variant="primary" leftIcon={<SearchIcon />}>
  搜索
</Button>

<Button variant="outline" rightIcon={<ArrowRightIcon />}>
  下一步
</Button>
```

### 仅图标按钮

```tsx
<Button
  variant="ghost"
  size="xs"
  aria-label="关闭"
  className="p-1" // 覆盖默认内边距，确保图标居中
>
  <CloseIcon className="h-4 w-4" />
</Button>
```

### 加载状态按钮

```tsx
<Button
  variant="primary"
  isLoading
>
  保存中
</Button>
```

### 按钮组

```tsx
<div className="inline-flex rounded-md shadow-sm">
  <Button
    variant="outline"
    className="rounded-r-none border-r-0"
  >
    上一步
  </Button>
  <Button
    variant="outline"
    className="rounded-l-none"
  >
    下一步
  </Button>
</div>
```

## 无障碍性考虑

为确保按钮组件符合无障碍性标准，请遵循以下原则：

1. **文字对比度**: 确保按钮文本与背景的对比度符合WCAG 2.1 AA标准（最低4.5:1）

2. **焦点状态**: 所有按钮必须有明显的焦点指示器（通过`focus:ring-2`实现）

3. **ARIA属性**: 
   - 对于图标按钮，添加`aria-label`属性描述按钮功能
   - 对于禁用按钮，确保设置`aria-disabled="true"`
   - 对于加载状态，使用`aria-busy="true"`

4. **键盘可访问性**: 确保所有按钮可通过Tab键聚焦

5. **标签清晰**: 按钮文本应清晰描述其行为，避免使用"点击这里"等模糊表述

## 性能优化

1. **通过条件渲染减少DOM**: 只在需要时渲染图标和加载指示器

2. **使用正确的语义元素**: 如果不是提交表单，考虑使用`type="button"`防止默认提交行为

3. **避免不必要的重渲染**: 考虑使用React.memo包装按钮组件以优化性能

## 响应式设计

针对不同设备尺寸，可以调整按钮的大小：

```jsx
<Button
  variant="primary"
  className="text-sm px-3 py-1.5 md:text-base md:px-4 md:py-2 lg:px-6 lg:py-3"
>
  响应式按钮
</Button>
```

或使用不同尺寸的工具类：

```jsx
<Button
  variant="primary"
  size="sm"
  className="md:text-base md:px-4 md:py-2 lg:px-6 lg:py-3"
>
  响应式按钮
</Button>
```

## 最佳实践

1. **一致性**: 在整个应用中保持按钮样式和行为的一致性，不要创建太多自定义按钮

2. **视觉层次**: 每个页面只使用一个主要按钮，避免"按钮竞争"

3. **反馈清晰**: 所有按钮交互应提供视觉反馈（悬停、点击、加载等状态）

4. **命名简洁**: 按钮文字应该简短、明确，表达用户期望的操作

5. **位置固定**: 在相似的模式中保持按钮位置的一致性（如表单提交按钮总是在右下角）

## 设计系统集成

1. **主题切换**: 确保按钮样式在暗色模式下仍然可用：

```tsx
// 暗色模式下的按钮样式扩展
const darkModeVariantClasses = {
  primary: 'dark:bg-primary-500 dark:hover:bg-primary-600 dark:focus:ring-primary-400',
  secondary: 'dark:bg-gray-700 dark:hover:bg-gray-600',
  outline: 'dark:border-gray-600 dark:text-gray-200 dark:hover:bg-gray-700',
  ghost: 'dark:text-gray-300 dark:hover:bg-gray-800',
  // 其他变体...
};
```

2. **扩展性**: 设计组件以便于扩展添加新的变体或尺寸

3. **兼容性**: 确保按钮样式在所有支持的浏览器中正常显示
``` 