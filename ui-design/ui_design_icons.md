# 图标系统设计指南

## 概述

图标是用户界面中重要的视觉元素，它们能直观地传达信息、引导用户操作并增强界面的可识别性。本文档详细定义了基于React、TypeScript和Tailwind CSS的图标系统，包括图标库选择、尺寸规范、颜色使用和实现方法。

## 图标库选择

为确保一致性和高质量的图标体验，我们推荐使用以下图标库之一：

### 1. Heroicons

- **特点**: 简洁、现代、精心设计的svg图标集，由Tailwind CSS团队维护
- **样式**: 提供outline和solid两种风格
- **官网**: [heroicons.com](https://heroicons.com/)
- **安装**: `npm install @heroicons/react`
- **适用场景**: 适合与Tailwind CSS配合使用的项目，简洁现代的界面

### 2. Phosphor Icons

- **特点**: 灵活、一致、可定制的图标库，拥有广泛的图标集合
- **样式**: 提供多种风格（Regular、Bold、Thin、Light、Fill、Duotone）
- **官网**: [phosphoricons.com](https://phosphoricons.com/)
- **安装**: `npm install phosphor-react`
- **适用场景**: 需要多种样式变体和丰富图标选择的项目

### 3. Lucide

- **特点**: Feather图标的社区延续版本，简洁优雅的线性图标
- **样式**: 统一的线型风格
- **官网**: [lucide.dev](https://lucide.dev/)
- **安装**: `npm install lucide-react`
- **适用场景**: 喜欢极简风格的项目

## 图标命名约定

为保持一致性并简化使用，我们采用以下命名约定：

1. **使用PascalCase命名图标组件**:
   - `UserIcon` 而非 `userIcon` 或 `user-icon`
   - `ChevronRightIcon` 而非 `chevron_right_icon`

2. **使用描述性名称**:
   - 使用 `ArrowUpIcon` 而非 `UpIcon`
   - 使用 `DocumentTextIcon` 而非 `FileIcon`（当需要区分不同类型文件时）

3. **成对图标保持一致**:
   - `ChevronLeftIcon` 和 `ChevronRightIcon`
   - `PlusIcon` 和 `MinusIcon`

4. **状态图标使用明确的描述**:
   - `CheckCircleIcon` 表示成功
   - `XCircleIcon` 表示错误
   - `ExclamationCircleIcon` 表示警告

## 图标尺寸系统

为确保整个应用中图标尺寸的一致性，我们定义了以下标准尺寸：

| 尺寸名称 | 像素值 | Tailwind类 | 使用场景 |
|---------|------|------------|---------|
| 超小 (xs) | 12px | w-3 h-3 | 紧凑型UI、小标签、指示器 |
| 小 (sm) | 16px | w-4 h-4 | 按钮内图标、紧凑导航、辅助说明 |
| 中 (md) | 20px | w-5 h-5 | 标准UI元素、表单控件、常规导航 |
| 大 (lg) | 24px | w-6 h-6 | 主要导航、强调元素、特性图标 |
| 超大 (xl) | 32px | w-8 h-8 | 页面模块、文档类型、功能区块 |
| 巨大 (2xl) | 48px | w-12 h-12 | 空状态、欢迎页、特殊场景 |

这些尺寸应根据上下文和布局灵活使用，但同一级别的UI元素应保持一致的图标尺寸。

## 图标颜色系统

图标的颜色应遵循整体设计系统的色彩规范，同时考虑以下几点：

### 颜色继承

默认情况下，SVG图标应继承其父元素的文本颜色（通过设置`fill="currentColor"`或`stroke="currentColor"`）。这样可以确保图标颜色与周围文本保持一致，并随主题变化而自动调整。

### 语义颜色

对于具有特定含义的图标，应使用语义色：

| 语义 | 颜色类 | 使用场景 |
|------|-------|---------|
| 信息 | text-info-500 | 提示、帮助、信息 |
| 成功 | text-success-500 | 完成、确认、验证通过 |
| 警告 | text-warning-500 | 需要注意、可能存在问题 |
| 错误 | text-error-500 | 错误、失败、被拒绝 |
| 禁用 | text-gray-300 | 不可用功能、禁用状态 |

### 对比度要求

确保图标与其背景之间有足够的对比度，遵循WCAG 2.1 AA标准：

- 图形界面组件的对比度至少为3:1
- 纯装饰性图标可以不受此限制
- 带文本标签的图标可以使用更低对比度

## TypeScript接口定义

```typescript
type IconSize = 'xs' | 'sm' | 'md' | 'lg' | 'xl' | '2xl';

interface IconProps {
  /** 图标尺寸 */
  size?: IconSize;
  /** 图标颜色，可使用语义色或自定义颜色 */
  color?: string;
  /** 额外的类名 */
  className?: string;
  /** 是否对屏幕阅读器隐藏图标 */
  'aria-hidden'?: boolean;
  /** 图标的可访问性描述，当图标单独使用时必须提供 */
  'aria-label'?: string;
}
```

## 图标组件实现

### 基础图标包装器

使用以下包装器组件来标准化图标使用：

```tsx
// Icon.tsx
import React from 'react';
import { classNames } from '../utils';

export const iconSizes = {
  xs: 'w-3 h-3',
  sm: 'w-4 h-4',
  md: 'w-5 h-5',
  lg: 'w-6 h-6', 
  xl: 'w-8 h-8',
  '2xl': 'w-12 h-12'
};

export const Icon: React.FC<IconProps> = ({ 
  size = 'md',
  color,
  className,
  'aria-hidden': ariaHidden = true,
  'aria-label': ariaLabel,
  children 
}) => {
  // 如果指定了颜色, 我们使用tailwind类
  const colorClass = color ? `text-${color}` : '';
  
  // 设置适当的可访问性属性
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
      xmlns="http://www.w3.org/2000/svg"
      fill="none"
      viewBox="0 0 24 24"
      stroke="currentColor"
      {...accessibilityProps}
    >
      {children}
    </svg>
  );
};
```

### 特定图标实现示例

使用图标包装器创建特定图标：

```tsx
// icons/CheckIcon.tsx
import React from 'react';
import { Icon, IconProps } from './Icon';

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

// icons/XIcon.tsx
import React from 'react';
import { Icon, IconProps } from './Icon';

export const XIcon: React.FC<IconProps> = (props) => (
  <Icon {...props}>
    <path 
      strokeLinecap="round" 
      strokeLinejoin="round" 
      strokeWidth={2} 
      d="M6 18L18 6M6 6l12 12" 
    />
  </Icon>
);
```

### 使用第三方图标库

如果使用第三方库如Heroicons，可以创建一个包装器以保持一致的API：

```tsx
// icons/HeroIcon.tsx
import React from 'react';
import { classNames } from '../utils';
import { iconSizes } from './Icon';

interface HeroIconProps extends IconProps {
  Icon: React.ComponentType<React.SVGProps<SVGSVGElement>>;
}

export const HeroIcon: React.FC<HeroIconProps> = ({
  Icon,
  size = 'md',
  color,
  className,
  'aria-hidden': ariaHidden = true,
  'aria-label': ariaLabel,
}) => {
  const colorClass = color ? `text-${color}` : '';
  
  const accessibilityProps = ariaLabel
    ? { 'aria-label': ariaLabel, 'aria-hidden': false }
    : { 'aria-hidden': ariaHidden };
  
  return (
    <Icon
      className={classNames(
        iconSizes[size],
        colorClass,
        className
      )}
      {...accessibilityProps}
    />
  );
};

// 使用示例
import { UserIcon as HeroUserIcon } from '@heroicons/react/24/outline';

export const UserIcon: React.FC<IconProps> = (props) => (
  <HeroIcon Icon={HeroUserIcon} {...props} />
);
```

## 图标集合组织

为了有效管理和使用图标，我们建议以下组织结构：

```
src/
└── components/
    └── icons/
        ├── index.ts        // 导出所有图标
        ├── Icon.tsx        // 基础图标组件
        ├── HeroIcon.tsx    // 第三方库包装器
        ├── UiIcons/        // UI界面图标
        │   ├── index.ts
        │   ├── HomeIcon.tsx
        │   ├── SearchIcon.tsx
        │   └── ...
        ├── ActionIcons/    // 动作图标
        │   ├── index.ts
        │   ├── EditIcon.tsx
        │   ├── DeleteIcon.tsx
        │   └── ...
        └── StatusIcons/    // 状态图标
            ├── index.ts
            ├── CheckCircleIcon.tsx
            ├── XCircleIcon.tsx
            └── ...
```

## 图标使用规范

### 图标 + 文本组合

当图标与文本一起使用时：

```tsx
<button className="flex items-center">
  <PlusIcon size="sm" className="mr-1.5" aria-hidden="true" />
  <span>添加项目</span>
</button>
```

1. 图标应设置`aria-hidden="true"`，让屏幕阅读器忽略
2. 保持一致的间距（图标和文本之间通常使用`mr-1.5`或`mr-2`）
3. 图标大小通常比文本小一级（如文本正常大小，图标使用`sm`尺寸）

### 单独使用的图标

当图标单独使用时（如图标按钮）：

```tsx
<button 
  className="p-2 rounded-full hover:bg-gray-100"
  aria-label="删除项目"
>
  <TrashIcon size="sm" aria-hidden="true" />
</button>
```

1. 必须为容器提供`aria-label`描述功能
2. 应确保点击区域足够大（通常至少40x40px）
3. 提供明确的视觉反馈（如hover、focus状态）

### 装饰性图标

纯装饰用途的图标：

```tsx
<div className="relative">
  <input type="text" className="pl-8" />
  <SearchIcon 
    size="sm" 
    className="absolute left-2 top-1/2 transform -translate-y-1/2 text-gray-400" 
    aria-hidden="true" 
  />
</div>
```

1. 标记为`aria-hidden="true"`
2. 使用较低对比度的颜色（如`text-gray-400`）
3. 不干扰主要内容的阅读

## 图标密度建议

为保持界面的清爽和可读性，建议遵循以下图标密度原则：

1. **相似图标保持间距**: 相似或相关的图标之间应有足够的间距（至少8px）
2. **避免过度使用**: 一个视窗中的图标数量不宜过多，避免视觉混乱
3. **保持一致性**: 同一界面区域使用一致的图标尺寸和样式
4. **分组使用**: 相关功能的图标应视觉上分组在一起

## 响应式图标

在不同设备尺寸上调整图标尺寸：

```tsx
<MenuIcon className="w-4 h-4 md:w-5 md:h-5 lg:w-6 lg:h-6" />
```

或使用我们的尺寸系统：

```tsx
<MenuIcon size="sm" className="md:w-5 md:h-5 lg:w-6 lg:h-6" />
```

## 自定义图标集成

如需整合自定义SVG图标：

1. **优化SVG文件**: 
   - 使用[SVGO](https://github.com/svg/svgo)删除多余元素
   - 确保设置`viewBox`属性
   - 移除固定的`width`和`height`

2. **转换为React组件**:
   ```tsx
   // 从.svg文件转换
   import React from 'react';
   import { Icon, IconProps } from './Icon';
   
   export const CustomIcon: React.FC<IconProps> = (props) => (
     <Icon viewBox="0 0 24 24" {...props}>
       {/* SVG路径内容 */}
       <path d="..." />
     </Icon>
   );
   ```

3. **添加到图标系统**:
   - 放置在适当的分类目录
   - 在相应的index.ts中导出
   - 在图标文档中注册

## 无障碍性考虑

1. **文本替代**: 
   - 为单独使用的图标提供描述性的`aria-label`
   - 配合文本使用的图标设置`aria-hidden="true"`

2. **焦点状态**:
   - 可交互图标（如图标按钮）必须有明显的焦点状态
   - 确保使用`focus:ring`或类似视觉指示

3. **色彩对比度**:
   - 功能性图标与背景的对比度至少为3:1
   - 使用工具测试对比度（如[Contrast Checker](https://webaim.org/resources/contrastchecker/)）

4. **不依赖颜色**:
   - 不仅依靠颜色传达信息，同时使用形状和标签
   - 考虑色盲用户的体验

## 图标动效

适当的动效可以增强用户体验：

```tsx
// 旋转加载图标
<RefreshIcon className="animate-spin h-5 w-5 text-primary-500" />

// 脉冲提示图标
<BellIcon className="animate-pulse h-5 w-5 text-warning-500" />
```

常见图标动效：

1. **旋转**: 用于加载状态 - `animate-spin`
2. **脉冲**: 用于提示注意 - `animate-pulse`
3. **弹跳**: 用于引导方向 - `animate-bounce`
4. **淡入/淡出**: 用于状态变化 - 使用CSS过渡

## 图标最佳实践

1. **保持一致性**: 在整个应用中使用一致的图标风格和尺寸

2. **简洁明了**: 选择简单、直观的图标，避免复杂或模糊的设计

3. **目的明确**: 图标应清晰表达其功能或含义，不要使用晦涩的隐喻

4. **适度使用**: 避免过度使用图标导致界面混乱

5. **配合文本**: 在复杂功能上，图标应与文本标签配合使用

6. **性能考虑**: 
   - 对于经常重用的图标，考虑使用SVG sprite
   - 避免过大的SVG文件
   - 按需加载图标以减少初始加载时间

## 管理和维护

1. **图标清单**: 
   - 维护一个图标清单文档，记录所有可用图标
   - 包括名称、用途和示例

2. **定期审查**:
   - 定期审查图标使用情况，移除未使用的图标
   - 确保新添加的图标符合设计规范

3. **版本控制**:
   - 跟踪图标库的变更
   - 在重大更新时通知团队成员 

## 图标选择空间设计指南

图标选择界面是允许用户从一组可用图标中选择所需图标的UI组件。这种界面在内容管理系统、设计工具、富文本编辑器等场景中非常常见。本节提供了设计和实现图标选择界面的完整指南。

### 1. 图标选择界面类型

根据使用场景和复杂度，图标选择界面可以有以下几种类型：

| 类型 | 描述 | 适用场景 |
|------|------|---------|
| 简单选择器 | 显示固定集合的图标网格 | 有限图标集，简单场景 |
| 分类选择器 | 按类别组织的图标选择界面 | 大量图标，需要分类管理 |
| 搜索式选择器 | 支持搜索和筛选的图标选择器 | 图标库较大，用户知道要找什么 |
| 高级选择器 | 集成搜索、分类、最近使用等功能 | 复杂应用，频繁使用图标选择 |

### 2. 图标选择器组件接口

```typescript
interface IconPickerProps {
  // 基本属性
  selectedIcon: string | null;
  onSelect: (iconName: string) => void;
  icons: IconDefinition[];
  
  // 布局与显示
  layout?: 'grid' | 'list';
  gridColumns?: number | { [breakpoint: string]: number };
  size?: 'sm' | 'md' | 'lg';
  showNames?: boolean;
  maxHeight?: string | number;
  
  // 功能性选项
  categories?: IconCategory[];
  searchable?: boolean;
  recentIconsCount?: number;
  
  // 自定义渲染
  renderIcon?: (icon: IconDefinition) => React.ReactNode;
  emptyState?: React.ReactNode;
  
  // 状态
  disabled?: boolean;
  loading?: boolean;
  
  // 样式和布局
  className?: string;
  containerClassName?: string;
  iconClassName?: string;
}

interface IconCategory {
  id: string;
  name: string;
  icons: string[]; // 该类别中的图标名称
}

interface IconDefinition {
  name: string;
  path: string; // SVG路径或组件引用
  tags?: string[]; // 用于搜索的标签
  category?: string;
}
```

### 3. 图标选择器布局与间距

图标选择器的布局需要考虑以下几点：

#### 3.1 网格布局

- **标准网格间距**: 16px (1rem, gap-4)
- **紧凑网格间距**: 8px (0.5rem, gap-2)
- **响应式列数**:
  - 移动设备: 4-6列
  - 平板设备: 6-8列
  - 桌面设备: 8-12列

#### 3.2 图标单元格

- **标准尺寸**: 48px × 48px (w-12 h-12)
- **可点击区域**: 至少44px × 44px，符合WCAG可访问性要求
- **图标与单元格比例**: 图标大小应为单元格的50%-75%
- **选中状态边距**: 保留2-4px的边距用于显示选中态

#### 3.3 图标选择器弹出层

- **最大高度**: 不超过屏幕高度的70%
- **最小宽度**: 320px
- **理想宽度**: 400-500px
- **边距**: 内边距16px (p-4)

### 4. 图标选择器交互设计

#### 4.1 选择状态

- **悬停状态**: 轻微背景色变化、提示文本
- **选中状态**: 明显的视觉标记（边框、背景色、对勾标记）
- **焦点状态**: 清晰的键盘焦点指示器

#### 4.2 搜索与筛选

- **顶部搜索框**: 始终可见的搜索输入框
- **实时筛选**: 输入时立即更新结果
- **空结果状态**: 清晰的空结果提示和建议
- **搜索关键词高亮**: 可选功能，高亮匹配的关键词

#### 4.3 分组与分类

- **分类标签页**: 使用标签页在主要类别间切换
- **分类标题**: 每个类别有清晰的标题
- **分类滚动**: 支持在类别内垂直滚动
- **展开/折叠**: 选择性支持类别展开折叠

### 5. 图标选择器实现示例

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
import React, { useState, useEffect } from 'react';
import { classNames } from '../utils';

export const IconPicker: React.FC<IconPickerProps> = ({
  selectedIcon,
  onSelect,
  icons,
  layout = 'grid',
  gridColumns = { default: 6, sm: 4, md: 6, lg: 8 },
  size = 'md',
  showNames = false,
  maxHeight = 400,
  categories = [],
  searchable = true,
  recentIconsCount = 0,
  renderIcon,
  emptyState,
  disabled = false,
  loading = false,
  className,
  containerClassName,
  iconClassName,
}) => {
  const [searchQuery, setSearchQuery] = useState('');
  const [activeCategory, setActiveCategory] = useState<string | null>(
    categories.length > 0 ? categories[0].id : null
  );
  const [recentIcons, setRecentIcons] = useState<string[]>([]);
  
  // 处理图标选择
  const handleIconSelect = (iconName: string) => {
    if (disabled) return;
    
    onSelect(iconName);
    
    // 更新最近使用的图标
    if (recentIconsCount > 0) {
      setRecentIcons(prev => {
        const newRecent = prev.filter(name => name !== iconName);
        newRecent.unshift(iconName);
        return newRecent.slice(0, recentIconsCount);
      });
    }
  };
  
  // 筛选图标
  const filteredIcons = icons.filter(icon => {
    // 搜索筛选
    if (searchQuery) {
      const query = searchQuery.toLowerCase();
      return (
        icon.name.toLowerCase().includes(query) ||
        icon.tags?.some(tag => tag.toLowerCase().includes(query))
      );
    }
    
    // 分类筛选
    if (activeCategory === 'recent') {
      return recentIcons.includes(icon.name);
    }
    
    if (activeCategory) {
      return icon.category === activeCategory;
    }
    
    return true;
  });
  
  // 尺寸映射
  const iconSizeClasses = {
    sm: 'w-8 h-8',
    md: 'w-10 h-10',
    lg: 'w-12 h-12',
  };
  
  // 网格列数映射
  const getGridColsClass = () => {
    if (typeof gridColumns === 'number') {
      return `grid-cols-${gridColumns}`;
    }
    
    // 响应式列数
    return Object.entries(gridColumns)
      .map(([breakpoint, cols]) => {
        if (breakpoint === 'default') {
          return `grid-cols-${cols}`;
        }
        return `${breakpoint}:grid-cols-${cols}`;
      })
      .join(' ');
  };
  
  return (
    <div className={classNames('icon-picker', className)}>
      {/* 搜索和分类区域 */}
      <div className="mb-4 space-y-3">
        {searchable && (
          <div className="relative">
            <input
              type="text"
              placeholder="搜索图标..."
              value={searchQuery}
              onChange={(e) => setSearchQuery(e.target.value)}
              disabled={disabled || loading}
              className="w-full px-3 py-2 pl-9 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-primary-500 focus:border-primary-500"
            />
            <svg
              className="absolute left-3 top-1/2 transform -translate-y-1/2 w-4 h-4 text-gray-400"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 20 20"
              fill="currentColor"
            >
              <path
                fillRule="evenodd"
                d="M8 4a4 4 0 100 8 4 4 0 000-8zM2 8a6 6 0 1110.89 3.476l4.817 4.817a1 1 0 01-1.414 1.414l-4.816-4.816A6 6 0 012 8z"
                clipRule="evenodd"
              />
            </svg>
            {searchQuery && (
              <button
                onClick={() => setSearchQuery('')}
                className="absolute right-3 top-1/2 transform -translate-y-1/2 w-4 h-4 text-gray-400 hover:text-gray-600"
                disabled={disabled || loading}
              >
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                  <path
                    fillRule="evenodd"
                    d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z"
                    clipRule="evenodd"
                  />
                </svg>
              </button>
            )}
          </div>
        )}
        
        {categories.length > 0 && (
          <div className="flex space-x-1 overflow-x-auto pb-1">
            {recentIconsCount > 0 && recentIcons.length > 0 && (
              <button
                className={classNames(
                  'px-3 py-1 text-sm rounded-md whitespace-nowrap',
                  activeCategory === 'recent'
                    ? 'bg-primary-100 text-primary-800 font-medium'
                    : 'bg-gray-100 text-gray-700 hover:bg-gray-200'
                )}
                onClick={() => setActiveCategory('recent')}
                disabled={disabled || loading}
              >
                最近使用
              </button>
            )}
            
            {categories.map((category) => (
              <button
                key={category.id}
                className={classNames(
                  'px-3 py-1 text-sm rounded-md whitespace-nowrap',
                  activeCategory === category.id
                    ? 'bg-primary-100 text-primary-800 font-medium'
                    : 'bg-gray-100 text-gray-700 hover:bg-gray-200'
                )}
                onClick={() => setActiveCategory(category.id)}
                disabled={disabled || loading}
              >
                {category.name}
              </button>
            ))}
            
            {activeCategory && (
              <button
                className="px-3 py-1 text-sm rounded-md whitespace-nowrap bg-gray-100 text-gray-700 hover:bg-gray-200"
                onClick={() => setActiveCategory(null)}
                disabled={disabled || loading}
              >
                全部
              </button>
            )}
          </div>
        )}
      </div>
      
      {/* 图标网格 */}
      <div
        className={classNames(
          'overflow-y-auto',
          containerClassName
        )}
        style={{ maxHeight }}
      >
        {loading ? (
          <div className="flex items-center justify-center py-8">
            <svg
              className="animate-spin w-8 h-8 text-primary-500"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
            >
              <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"></circle>
              <path
                className="opacity-75"
                fill="currentColor"
                d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
              ></path>
            </svg>
          </div>
        ) : filteredIcons.length === 0 ? (
          <div className="flex flex-col items-center justify-center py-8 text-gray-500">
            {emptyState || (
              <>
                <svg
                  className="w-16 h-16 mb-2 text-gray-300"
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M9.172 16.172a4 4 0 015.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
                  />
                </svg>
                <p className="text-center">未找到图标</p>
                {searchQuery && <p className="text-sm">尝试不同的搜索词或浏览分类</p>}
              </>
            )}
          </div>
        ) : (
          <div
            className={classNames(
              layout === 'grid'
                ? `grid ${getGridColsClass()} gap-3`
                : 'flex flex-col space-y-2'
            )}
          >
            {filteredIcons.map((icon) => (
              <button
                key={icon.name}
                onClick={() => handleIconSelect(icon.name)}
                className={classNames(
                  'group flex flex-col items-center justify-center p-2 rounded-md transition-colors',
                  selectedIcon === icon.name
                    ? 'ring-2 ring-primary-500 bg-primary-50'
                    : 'hover:bg-gray-100',
                  disabled ? 'opacity-50 cursor-not-allowed' : 'cursor-pointer',
                  layout === 'list' ? 'flex-row justify-start space-x-3' : '',
                  iconClassName
                )}
                disabled={disabled}
                aria-label={`选择图标: ${icon.name}`}
              >
                {renderIcon ? (
                  renderIcon(icon)
                ) : (
                  <svg
                    className={classNames(
                      'text-gray-700',
                      iconSizeClasses[size],
                      selectedIcon === icon.name ? 'text-primary-600' : 'group-hover:text-gray-900'
                    )}
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                    dangerouslySetInnerHTML={{ __html: icon.path }}
                  />
                )}
                
                {showNames && (
                  <span
                    className={classNames(
                      'mt-1 text-xs truncate max-w-full',
                      selectedIcon === icon.name ? 'text-primary-600' : 'text-gray-600',
                      layout === 'list' ? 'mt-0 text-sm' : ''
                    )}
                  >
                    {icon.name}
                  </span>
                )}
              </button>
            ))}
          </div>
        )}
      </div>
    </div>
  );
};
// [AI-BLOCK-END]

### 6. 图标选择器可访问性考虑

#### 6.1 键盘导航

- **方向键导航**: 使用方向键在图标网格中移动
- **Enter键选择**: 按Enter键选择当前聚焦的图标
- **Tab键序列**: 合理的Tab序列，包括搜索框和分类按钮
- **Escape键关闭**: 按Escape键关闭弹出式选择器

#### 6.2 屏幕阅读器支持

- **图标名称朗读**: 每个图标有明确的`aria-label`属性
- **当前选中状态**: 使用`aria-selected`表示当前选中的图标
- **搜索结果反馈**: 宣告搜索结果数量变化
- **分组标记**: 使用适当的ARIA角色标记分组

#### 6.3 可访问性最佳实践

- **颜色对比度**: 图标与背景的对比度至少为3:1
- **图标名称显示**: 考虑默认显示图标名称以增强可用性
- **错误状态反馈**: 搜索无结果时提供清晰的反馈
- **缩放支持**: 确保组件在页面缩放时仍可用

### 7. 响应式设计考虑

- **移动设备优化**:
  - 更大的点击目标（至少48×48像素）
  - 减少列数以增大图标尺寸
  - 考虑全屏模式下的选择器

- **平板设备调整**:
  - 平衡的列数与图标尺寸
  - 横向滚动的分类标签

- **桌面布局**:
  - 更多列以利用宽屏空间
  - 并排显示搜索和分类

### 8. 集成与使用场景

#### 8.1 表单集成

```tsx
<form>
  <div className="mb-4">
    <label className="block text-sm font-medium text-gray-700 mb-1">
      选择图标
    </label>
    <div className="relative">
      <button
        type="button"
        className="flex items-center w-full px-3 py-2 border border-gray-300 rounded-md"
        onClick={() => setShowIconPicker(true)}
      >
        {selectedIcon ? (
          <>
            <span className="w-5 h-5 mr-2">
              {/* 显示所选图标 */}
            </span>
            <span>{selectedIcon}</span>
          </>
        ) : (
          <span className="text-gray-500">点击选择图标</span>
        )}
      </button>
      
      {showIconPicker && (
        <div className="absolute z-10 w-full mt-1 bg-white rounded-md shadow-lg border border-gray-200">
          <IconPicker
            selectedIcon={selectedIcon}
            onSelect={(icon) => {
              setSelectedIcon(icon);
              setShowIconPicker(false);
            }}
            icons={iconList}
            searchable
            showNames
          />
          <div className="p-2 border-t flex justify-end">
            <button
              type="button"
              className="px-3 py-1 text-sm text-gray-700 hover:bg-gray-100 rounded"
              onClick={() => setShowIconPicker(false)}
            >
              取消
            </button>
          </div>
        </div>
      )}
    </div>
  </div>
</form>
```

#### 8.2 内容编辑器集成

```tsx
<div className="rich-text-editor">
  <div className="toolbar flex items-center space-x-1 p-1 border-b">
    <button
      type="button"
      className="p-1 rounded hover:bg-gray-100"
      onClick={() => setShowIconPicker(true)}
      title="插入图标"
    >
      <svg className="w-5 h-5" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
      </svg>
    </button>
    {/* 其他工具栏按钮 */}
    
    {showIconPicker && (
      <div className="absolute z-10 w-80 mt-1 bg-white rounded-md shadow-lg border border-gray-200">
        <IconPicker
          selectedIcon={null}
          onSelect={(icon) => {
            insertIconAtCursor(icon);
            setShowIconPicker(false);
          }}
          icons={iconList}
          searchable
          categories={iconCategories}
        />
      </div>
    )}
  </div>
  
  <div className="editor-content p-4">
    {/* 编辑器内容 */}
  </div>
</div>
```

## 管理和维护

1. **图标清单**: 
   - 维护一个图标清单文档，记录所有可用图标
   - 包括名称、用途和示例

2. **定期审查**:
   - 定期审查图标使用情况，移除未使用的图标
   - 确保新添加的图标符合设计规范

3. **版本控制**:
   - 跟踪图标库的变更
   - 在重大更新时通知团队成员 