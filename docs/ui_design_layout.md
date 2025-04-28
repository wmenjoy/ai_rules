# 布局系统设计指南

## 1. 概述

布局系统是设计系统的基础，它提供了组织和排列页面元素的结构。本指南详细说明在 React + TypeScript + Tailwind CSS 项目中如何设计和实现一个全面、灵活且响应式的布局系统。

## 2. 设计原则

布局系统遵循以下设计原则：

- **一致性**：提供一致的间距和对齐方式
- **灵活性**：支持各种内容和屏幕尺寸
- **响应式**：在不同设备上提供最佳体验
- **组合性**：允许组件自由组合创建复杂布局
- **可预测性**：布局行为应当直观且可预测

## 3. 网格系统 (Grid)

网格系统提供了一种基于列的结构，帮助设计师和开发者创建对齐的布局。

### 3.1 网格系统特点

| 特性 | 说明 |
|------|------|
| 12列系统 | 基于12列的网格，便于划分不同比例的内容区域 |
| 响应式列数 | 在不同断点可以定义不同的列数 |
| 间隙控制 | 可自定义列间距和行间距 |
| 自动均分 | 支持自动均分列宽 |
| 显式定位 | 支持通过起始位置和跨度控制元素位置 |

### 3.2 网格系统 TypeScript 接口

```typescript
interface GridProps {
  columns?: number | { [key: string]: number }; // 响应式列数
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

### 3.3 网格系统实现示例

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
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
  // 处理响应式列数
  const getColumnsClass = () => {
    if (typeof columns === 'number') {
      return `grid-cols-${columns}`;
    }
    
    // 响应式对象，如 { sm: 2, md: 3, lg: 4 }
    return Object.entries(columns)
      .map(([breakpoint, value]) => {
        return breakpoint === 'xs' 
          ? `grid-cols-${value}` 
          : `${breakpoint}:grid-cols-${value}`;
      })
      .join(' ');
  };
  
  // 处理间隙
  const getGapClass = () => {
    if (!gap) return '';
    
    if (typeof gap === 'number' || typeof gap === 'string') {
      return `gap-${gap}`;
    }
    
    // 响应式对象
    return Object.entries(gap)
      .map(([breakpoint, value]) => {
        return breakpoint === 'xs' 
          ? `gap-${value}` 
          : `${breakpoint}:gap-${value}`;
      })
      .join(' ');
  };
  
  // 行间隙
  const getRowGapClass = () => {
    if (!rowGap) return '';
    return `gap-y-${rowGap}`;
  };
  
  // 列间隙
  const getColumnGapClass = () => {
    if (!columnGap) return '';
    return `gap-x-${columnGap}`;
  };
  
  // 自动流向
  const getAutoFlowClass = () => {
    return `grid-flow-${autoFlow.replace(' ', '-')}`;
  };
  
  return (
    <Component
      className={classNames(
        'grid',
        getColumnsClass(),
        getGapClass(),
        getRowGapClass(),
        getColumnGapClass(),
        getAutoFlowClass(),
        className
      )}
    >
      {children}
    </Component>
  );
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
  // 处理列跨度
  const getColSpanClass = () => {
    if (!colSpan) return '';
    
    if (typeof colSpan === 'number') {
      return `col-span-${colSpan}`;
    }
    
    // 响应式对象
    return Object.entries(colSpan)
      .map(([breakpoint, value]) => {
        return breakpoint === 'xs' 
          ? `col-span-${value}` 
          : `${breakpoint}:col-span-${value}`;
      })
      .join(' ');
  };
  
  // 处理行跨度
  const getRowSpanClass = () => {
    if (!rowSpan) return '';
    
    if (typeof rowSpan === 'number') {
      return `row-span-${rowSpan}`;
    }
    
    // 响应式对象
    return Object.entries(rowSpan)
      .map(([breakpoint, value]) => {
        return breakpoint === 'xs' 
          ? `row-span-${value}` 
          : `${breakpoint}:row-span-${value}`;
      })
      .join(' ');
  };
  
  // 处理列开始位置
  const getColStartClass = () => {
    if (!colStart) return '';
    
    if (typeof colStart === 'number') {
      return `col-start-${colStart}`;
    }
    
    // 响应式对象
    return Object.entries(colStart)
      .map(([breakpoint, value]) => {
        return breakpoint === 'xs' 
          ? `col-start-${value}` 
          : `${breakpoint}:col-start-${value}`;
      })
      .join(' ');
  };
  
  // 处理行开始位置
  const getRowStartClass = () => {
    if (!rowStart) return '';
    
    if (typeof rowStart === 'number') {
      return `row-start-${rowStart}`;
    }
    
    // 响应式对象
    return Object.entries(rowStart)
      .map(([breakpoint, value]) => {
        return breakpoint === 'xs' 
          ? `row-start-${value}` 
          : `${breakpoint}:row-start-${value}`;
      })
      .join(' ');
  };
  
  return (
    <Component
      className={classNames(
        getColSpanClass(),
        getRowSpanClass(),
        getColStartClass(),
        getRowStartClass(),
        className
      )}
    >
      {children}
    </Component>
  );
};
// [AI-BLOCK-END]
```

### 3.4 网格系统使用示例

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
import { Grid, GridItem } from './components';

// 基本网格示例
const BasicGridExample = () => {
  return (
    <Grid columns={3} gap={4}>
      <div className="bg-blue-100 p-4 rounded">项目 1</div>
      <div className="bg-blue-100 p-4 rounded">项目 2</div>
      <div className="bg-blue-100 p-4 rounded">项目 3</div>
      <div className="bg-blue-100 p-4 rounded">项目 4</div>
      <div className="bg-blue-100 p-4 rounded">项目 5</div>
      <div className="bg-blue-100 p-4 rounded">项目 6</div>
    </Grid>
  );
};

// 响应式网格示例
const ResponsiveGridExample = () => {
  return (
    <Grid columns={{ xs: 1, sm: 2, md: 3, lg: 4 }} gap={4}>
      <div className="bg-blue-100 p-4 rounded">项目 1</div>
      <div className="bg-blue-100 p-4 rounded">项目 2</div>
      <div className="bg-blue-100 p-4 rounded">项目 3</div>
      <div className="bg-blue-100 p-4 rounded">项目 4</div>
    </Grid>
  );
};

// 使用GridItem控制布局示例
const ComplexGridExample = () => {
  return (
    <Grid columns={12} gap={4}>
      <GridItem colSpan={12} className="bg-blue-500 p-4 text-white rounded">
        头部区域 (12列)
      </GridItem>
      
      <GridItem colSpan={3} className="bg-blue-300 p-4 rounded">
        侧边栏 (3列)
      </GridItem>
      
      <GridItem colSpan={9} className="bg-blue-200 p-4 rounded">
        主内容区域 (9列)
      </GridItem>
      
      <GridItem colSpan={{ xs: 12, md: 6 }} className="bg-blue-100 p-4 rounded">
        底部区域左侧 (响应式)
      </GridItem>
      
      <GridItem colSpan={{ xs: 12, md: 6 }} className="bg-blue-100 p-4 rounded">
        底部区域右侧 (响应式)
      </GridItem>
    </Grid>
  );
};

// 显式位置示例
const PositionedGridExample = () => {
  return (
    <Grid columns={3} gap={4}>
      <GridItem colStart={1} colSpan={2} className="bg-red-100 p-4 rounded">
        跨越1-2列
      </GridItem>
      
      <GridItem colStart={3} colSpan={1} className="bg-blue-100 p-4 rounded">
        第3列
      </GridItem>
      
      <GridItem colSpan={3} className="bg-green-100 p-4 rounded">
        跨越所有列
      </GridItem>
    </Grid>
  );
};
// [AI-BLOCK-END]
```

## 4. 栈组件 (Stack)

栈组件是一种灵活的布局工具，用于在水平或垂直方向上均匀排列元素。

### 4.1 栈组件特点

| 特性 | 说明 |
|------|------|
| 方向控制 | 支持水平或垂直排列 |
| 间距管理 | 提供统一的元素间距 |
| 对齐方式 | 支持在主轴和交叉轴上进行对齐 |
| 响应式行为 | 可在不同断点改变排列方向 |
| 嵌套能力 | 可以嵌套使用，创建复杂布局 |

### 4.2 栈组件 TypeScript 接口

```typescript
type StackDirection = 'row' | 'column';
type StackSpacing = 0 | 0.5 | 1 | 1.5 | 2 | 2.5 | 3 | 3.5 | 4 | 5 | 6 | 7 | 8 | 9 | 10 | 11 | 12 | 14 | 16 | 20 | 24 | 28 | 32 | 36 | 40 | 44 | 48 | 52 | 56 | 60 | 64 | 72 | 80 | 96;

interface StackProps {
  direction?: StackDirection | { [key: string]: StackDirection };
  spacing?: StackSpacing | { [key: string]: StackSpacing };
  align?: 'start' | 'end' | 'center' | 'baseline' | 'stretch';
  justify?: 'start' | 'end' | 'center' | 'between' | 'around' | 'evenly';
  wrap?: boolean;
  as?: React.ElementType;
  className?: string;
  children: React.ReactNode;
}
```

### 4.3 栈组件实现示例

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
import React from 'react';
import { classNames } from '../utils';

export const Stack: React.FC<StackProps> = ({
  direction = 'column',
  spacing = 4,
  align,
  justify,
  wrap = false,
  as: Component = 'div',
  className,
  children,
}) => {
  // 处理响应式方向
  const getDirectionClass = () => {
    if (typeof direction === 'string') {
      return direction === 'row' ? 'flex-row' : 'flex-col';
    }
    
    // 响应式对象，如 { sm: 'row', lg: 'column' }
    return Object.entries(direction)
      .map(([breakpoint, value]) => {
        const dirClass = value === 'row' ? 'flex-row' : 'flex-col';
        return breakpoint === 'xs' 
          ? dirClass 
          : `${breakpoint}:${dirClass}`;
      })
      .join(' ');
  };
  
  // 处理响应式间距
  const getSpacingClass = () => {
    const getSpaceValue = (dir: StackDirection, space: StackSpacing) => {
      if (space === 0) return '';
      return dir === 'row' ? `space-x-${space}` : `space-y-${space}`;
    };
    
    if (typeof spacing === 'number') {
      if (typeof direction === 'string') {
        return getSpaceValue(direction, spacing);
      }
      
      // 简化处理，对于真正响应式的间距和方向需要更复杂的实现
      return getSpaceValue('column', spacing);
    }
    
    // 响应式对象
    return Object.entries(spacing)
      .map(([breakpoint, value]) => {
        const dir = typeof direction === 'string' ? direction : direction[breakpoint] || 'column';
        const spaceClass = getSpaceValue(dir, value);
        
        return breakpoint === 'xs' 
          ? spaceClass 
          : `${breakpoint}:${spaceClass}`;
      })
      .join(' ');
  };
  
  // 对齐方式
  const getAlignClass = () => {
    if (!align) return '';
    
    const alignments: Record<string, string> = {
      start: 'items-start',
      end: 'items-end',
      center: 'items-center',
      baseline: 'items-baseline',
      stretch: 'items-stretch',
    };
    
    return alignments[align];
  };
  
  // 主轴对齐
  const getJustifyClass = () => {
    if (!justify) return '';
    
    const justifications: Record<string, string> = {
      start: 'justify-start',
      end: 'justify-end',
      center: 'justify-center',
      between: 'justify-between',
      around: 'justify-around',
      evenly: 'justify-evenly',
    };
    
    return justifications[justify];
  };
  
  // 换行
  const getWrapClass = () => {
    return wrap ? 'flex-wrap' : 'flex-nowrap';
  };
  
  return (
    <Component
      className={classNames(
        'flex',
        getDirectionClass(),
        getSpacingClass(),
        getAlignClass(),
        getJustifyClass(),
        getWrapClass(),
        className
      )}
    >
      {children}
    </Component>
  );
};
// [AI-BLOCK-END]
```

### 4.4 栈组件使用示例

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
import { Stack } from './components';

// 垂直栈示例
const VerticalStackExample = () => {
  return (
    <Stack spacing={4}>
      <div className="bg-blue-100 p-4 rounded">项目 1</div>
      <div className="bg-blue-100 p-4 rounded">项目 2</div>
      <div className="bg-blue-100 p-4 rounded">项目 3</div>
    </Stack>
  );
};

// 水平栈示例
const HorizontalStackExample = () => {
  return (
    <Stack direction="row" spacing={4} align="center">
      <div className="bg-blue-100 p-4 rounded">项目 1</div>
      <div className="bg-blue-100 p-4 h-24 rounded">高度更大的项目</div>
      <div className="bg-blue-100 p-4 rounded">项目 3</div>
    </Stack>
  );
};

// 响应式栈示例
const ResponsiveStackExample = () => {
  return (
    <Stack 
      direction={{ xs: 'column', md: 'row' }} 
      spacing={{ xs: 2, md: 4 }}
      align="center"
    >
      <div className="bg-blue-100 p-4 rounded">在移动设备上垂直堆叠</div>
      <div className="bg-blue-100 p-4 rounded">在中等屏幕上水平排列</div>
      <div className="bg-blue-100 p-4 rounded">间距也会根据屏幕大小调整</div>
    </Stack>
  );
};

// 嵌套栈示例
const NestedStackExample = () => {
  return (
    <Stack spacing={8}>
      <div className="bg-gray-100 p-4 rounded">
        <h3 className="font-medium">第一部分</h3>
      </div>
      
      <Stack direction="row" spacing={4} wrap>
        <div className="bg-blue-100 p-4 rounded">嵌套项目 1</div>
        <div className="bg-blue-100 p-4 rounded">嵌套项目 2</div>
        <div className="bg-blue-100 p-4 rounded">嵌套项目 3</div>
      </Stack>
      
      <div className="bg-gray-100 p-4 rounded">
        <h3 className="font-medium">第二部分</h3>
      </div>
    </Stack>
  );
};
// [AI-BLOCK-END]
```

## 5. 分隔线组件 (Divider)

分隔线提供视觉分隔，帮助组织内容和创建层次结构。

### 5.1 分隔线 TypeScript 接口

```typescript
interface DividerProps {
  orientation?: 'horizontal' | 'vertical';
  variant?: 'solid' | 'dashed' | 'dotted';
  color?: string;
  thickness?: 'thin' | 'medium' | 'thick';
  className?: string;
}
```

### 5.2 分隔线实现示例

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
import React from 'react';
import { classNames } from '../utils';

export const Divider: React.FC<DividerProps> = ({
  orientation = 'horizontal',
  variant = 'solid',
  color = 'gray-200',
  thickness = 'thin',
  className,
}) => {
  // 边框样式
  const getVariantClass = () => {
    switch (variant) {
      case 'dashed':
        return 'border-dashed';
      case 'dotted':
        return 'border-dotted';
      case 'solid':
      default:
        return 'border-solid';
    }
  };
  
  // 边框厚度
  const getThicknessClass = () => {
    switch (thickness) {
      case 'medium':
        return orientation === 'horizontal' ? 'border-t-2' : 'border-l-2';
      case 'thick':
        return orientation === 'horizontal' ? 'border-t-4' : 'border-l-4';
      case 'thin':
      default:
        return orientation === 'horizontal' ? 'border-t' : 'border-l';
    }
  };
  
  // 方向特定类
  const getOrientationClass = () => {
    return orientation === 'horizontal'
      ? 'w-full'
      : 'h-full self-stretch';
  };
  
  return (
    <hr
      className={classNames(
        getVariantClass(),
        getThicknessClass(),
        getOrientationClass(),
        `border-${color}`,
        className
      )}
      aria-orientation={orientation}
    />
  );
};
// [AI-BLOCK-END]
```

## 6. 间距系统

间距系统基于一致的比例提供间距值，用于设置元素的边距和内边距。

### 6.1 间距刻度

Tailwind CSS 默认使用 0.25rem (4px) 作为基本单位，形成如下间距刻度：

| 刻度值 | 像素值 | 用途 |
|--------|-------|------|
| 0 | 0px | 无间距 |
| 1 | 4px | 极小间距，紧密元素 |
| 2 | 8px | 小间距，如图标与文本间 |
| 3 | 12px | 中小间距 |
| 4 | 16px | 标准间距，常用默认值 |
| 5 | 20px | 中等间距 |
| 6 | 24px | 中大间距，如段落间 |
| 8 | 32px | 大间距，如卡片间 |
| 10 | 40px | 更大间距，如组件间 |
| 12 | 48px | 区块间大间距 |
| 16 | 64px | 主要区域间距 |
| 20 | 80px | 大区域间距 |
| 24 | 96px | 部分间最大间距 |

### 6.2 间距使用场景

| 场景 | 推荐间距 | 类名示例 |
|------|---------|---------|
| 图标与文本间距 | 2 (8px) | `gap-2`, `mr-2` |
| 相关表单项之间 | 4 (16px) | `space-y-4` |
| 卡片内部内间距 | 4-6 (16-24px) | `p-4`, `p-6` |
| 卡片之间的间距 | 6-8 (24-32px) | `gap-6`, `gap-8` |
| 主要部分间距 | 10-16 (40-64px) | `mt-10`, `mb-16` |

## 7. 响应式系统

响应式系统基于断点控制布局在不同屏幕尺寸上的行为。

### 7.1 断点定义

| 断点名称 | 像素值 | 针对设备 |
|---------|-------|---------|
| xs | 默认 | 移动设备 - 小型手机 |
| sm | 640px | 移动设备 - 手机横屏及大型手机 |
| md | 768px | 平板设备 - 纵向 |
| lg | 1024px | 平板设备 - 横向、小型笔记本 |
| xl | 1280px | 桌面设备 - 标准显示器 |
| 2xl | 1536px | 桌面设备 - 大型显示器 |

### 7.2 响应式前缀用法

在 Tailwind 中，可以使用断点前缀控制响应式行为：

```html
<div class="w-full md:w-2/3 lg:w-1/2 p-4 md:p-6 lg:p-8">
  响应式内容：在移动设备上全宽，中等屏幕上 2/3 宽，大屏幕上 1/2 宽，
  且内边距随屏幕尺寸增加而增加
</div>
```

## 8. 布局模式

### 8.1 常见布局模式

| 布局模式 | 实现方式 | 适用场景 |
|---------|---------|---------|
| 单栏布局 | `Stack` 组件 | 简单文本页面、登录页 |
| 两栏布局 | `Grid` 组件 (colSpan 分配) | 边栏+内容、面板+详情 |
| 三栏布局 | `Grid` 组件 (12栏分为3份) | 复杂管理界面 |
| 卡片网格 | `Grid` 组件 (响应式列数) | 产品列表、图片集 |
| 混合布局 | 嵌套 `Grid` 和 `Stack` | 复杂应用界面 |

### 8.2 实现示例：响应式两栏布局

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
import { Grid, GridItem } from './components';

const TwoColumnLayout = ({ sidebar, content }) => {
  return (
    <Grid columns={12} gap={6}>
      {/* 侧边栏：移动设备上全宽，中等屏幕上占3列 */}
      <GridItem colSpan={{ xs: 12, md: 3 }} className="bg-gray-50 p-4 rounded">
        {sidebar}
      </GridItem>
      
      {/* 主内容：移动设备上全宽，中等屏幕上占9列 */}
      <GridItem colSpan={{ xs: 12, md: 9 }} className="p-4">
        {content}
      </GridItem>
    </Grid>
  );
};
// [AI-BLOCK-END]
```

### 8.3 实现示例：圣杯布局

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
import { Grid, GridItem } from './components';

const HolyGrailLayout = ({ header, footer, nav, content, ads }) => {
  return (
    <Grid columns={12} gap={4} className="min-h-screen">
      {/* 头部 */}
      <GridItem colSpan={12} className="bg-blue-100 p-4">
        {header}
      </GridItem>
      
      {/* 导航区 - 移动设备上位于内容上方 */}
      <GridItem colSpan={{ xs: 12, md: 2 }} className="bg-blue-50 p-4">
        {nav}
      </GridItem>
      
      {/* 主内容区 */}
      <GridItem colSpan={{ xs: 12, md: 8 }} className="bg-white p-4">
        {content}
      </GridItem>
      
      {/* 广告/附加区 - 移动设备上位于内容下方 */}
      <GridItem colSpan={{ xs: 12, md: 2 }} className="bg-blue-50 p-4">
        {ads}
      </GridItem>
      
      {/* 底部 */}
      <GridItem colSpan={12} className="bg-blue-100 p-4 mt-auto">
        {footer}
      </GridItem>
    </Grid>
  );
};
// [AI-BLOCK-END]
```

## 9. 布局系统最佳实践

### 9.1 性能考虑

- 避免过度嵌套布局组件，以减少DOM复杂度
- 使用 `will-change` 和 CSS 硬件加速处理复杂动画
- 考虑大型列表的虚拟化渲染
- 优先使用 CSS Grid 而非复杂的嵌套 Flex 布局
- 考虑使用 `React.memo` 优化布局组件重渲染

### 9.2 可访问性考虑

- 使用语义化 HTML 元素（通过 `as` 属性）
- 确保内容的逻辑顺序与视觉顺序一致
- 在移动设备上确保触摸目标足够大（至少44×44像素）
- 维护合理的内容宽度，提高可读性（通常最大宽度为65-75个字符）
- 在使用 Grid 和绝对定位时注意键盘导航顺序

### 9.3 响应式设计策略

- 采用移动优先的设计方法
- 使用相对单位（rem, em）而非固定像素值
- 在断点变化处测试布局，确保平滑过渡
- 注意内容密度和触摸友好性
- 使用媒体查询调整关键UI元素的大小和位置

### 9.4 扩展与定制

布局系统可通过以下方式扩展：
- 添加自定义断点以适应特定项目需求
- 扩展间距系统以包含特殊值
- 创建项目特定的布局模式组件
- 实现容器查询的包装组件（当浏览器支持时）

## 10. 图标栅格系统设计 (Icon Grid System Design)

图标栅格系统是确保整个应用程序中图标一致性和专业外观的关键。本节详细说明如何在 React + TypeScript + Tailwind CSS 项目中设计和实现一个全面、灵活的图标系统。

### 10.1 图标栅格基础 (Icon Grid Basics)

图标栅格是设计图标的基础框架，它确保所有图标在视觉上保持一致的比例和对齐方式。

#### 10.1.1 基础栅格尺寸

| 栅格尺寸 | 像素尺寸 | 适用场景 |
|---------|---------|---------|
| 微小 | 16×16 | 极小UI元素、文本内图标、密集列表 |
| 小型 | 20×20 | 小型控件、紧凑UI、辅助操作 |
| 标准 | 24×24 | 标准UI控件、工具栏、导航元素 |
| 中型 | 32×32 | 主要操作按钮、突出显示的控件 |
| 大型 | 48×48 | 移动触摸目标、特征图标、营销材料 |
| 超大 | 64×64 | 特性图标、横幅、空状态插图 |

#### 10.1.2 图标设计关键线

关键线是图标设计中用于确保视觉一致性的参考线。

| 栅格尺寸 | 中心区域 | 边距/填充 | 推荐线条粗细 | 推荐圆角 |
|---------|---------|----------|------------|---------|
| 16×16 | 14×14 | 1px | 1px | 1px |
| 20×20 | 18×18 | 1px | 1.5px | 1.5px |
| 24×24 | 20×20 | 2px | 1.5px | 2px |
| 32×32 | 28×28 | 2px | 2px | 2px |
| 48×48 | 40×40 | 4px | 2px | 3px |
| 64×64 | 56×56 | 4px | 2.5px | 4px |

#### 10.1.3 图标栅格对齐原则

1. **像素对齐** - 所有路径和形状应当对齐到整数像素点，避免模糊边缘
2. **视觉重心** - 图标应在其边界框内视觉居中，考虑视觉重量而非几何中心
3. **一致笔画** - 同一组图标应使用一致的笔画粗细和风格
4. **形状简化** - 避免过于复杂的形状，专注于清晰、易识别的轮廓
5. **视觉平衡** - 图标在各个方向上应保持视觉平衡

### 10.2 图标组件 TypeScript 接口 (Icon Component TypeScript Interface)

```typescript
// 图标尺寸选项
type IconSize = 'xs' | 'sm' | 'md' | 'lg' | 'xl' | '2xl';

// 图标风格变体
type IconVariant = 'filled' | 'outlined' | 'duotone' | 'thin';

// 图标颜色类型
type IconColor = 'inherit' | 'current' | string;

// 基础图标属性
interface IconBaseProps {
  size?: IconSize;
  color?: IconColor;
  variant?: IconVariant;
  className?: string;
  ariaHidden?: boolean;
  title?: string;
  testId?: string;
}

// SVG图标属性
interface SvgIconProps extends IconBaseProps {
  path: string;
  viewBox?: string;
}

// 字体图标属性
interface FontIconProps extends IconBaseProps {
  name: string;
  fontFamily?: string;
}

// 统一图标属性
interface IconProps extends IconBaseProps {
  name: string;
  type?: 'svg' | 'font';
}

// 图标按钮属性
interface IconButtonProps {
  icon: string;
  size?: 'sm' | 'md' | 'lg';
  variant?: 'primary' | 'secondary' | 'tertiary' | 'ghost';
  ariaLabel: string;
  onClick?: () => void;
  disabled?: boolean;
  className?: string;
  children?: React.ReactNode;
}
```

### 10.3 图标栅格实现示例 (Icon Grid Implementation Examples)

#### 10.3.1 基础SVG图标组件

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
import React from 'react';
import { classNames } from '../utils';

export const SvgIcon: React.FC<SvgIconProps> = ({
  path,
  size = 'md',
  color = 'currentColor',
  variant = 'outlined',
  viewBox = '0 0 24 24',
  className,
  ariaHidden = true,
  title,
  testId,
  ...props
}) => {
  // 尺寸映射到实际像素尺寸
  const sizeMap: Record<IconSize, string> = {
    'xs': '16',
    'sm': '20',
    'md': '24',
    'lg': '32',
    'xl': '48',
    '2xl': '64'
  };
  
  // 获取尺寸类名
  const getSizeClass = () => {
    const sizePx = sizeMap[size];
    return `w-[${sizePx}px] h-[${sizePx}px]`;
  };
  
  // 颜色处理
  const getColorClass = () => {
    if (color === 'inherit') return 'text-inherit';
    if (color === 'current') return 'text-current';
    if (color.startsWith('text-')) return color;
    return `text-${color}`;
  };
  
  // 变体处理
  const getVariantClasses = () => {
    switch (variant) {
      case 'filled':
        return 'fill-current stroke-none';
      case 'outlined':
        return 'fill-none stroke-current stroke-[1.5px]';
      case 'duotone':
        return 'fill-current fill-opacity-20 stroke-current stroke-[1.5px]';
      case 'thin':
        return 'fill-none stroke-current stroke-[1px]';
      default:
        return 'fill-none stroke-current stroke-[1.5px]';
    }
  };
  
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      viewBox={viewBox}
      className={classNames(
        getSizeClass(),
        getColorClass(),
        getVariantClasses(),
        className
      )}
      aria-hidden={ariaHidden}
      data-testid={testId}
      {...props}
    >
      {title && <title>{title}</title>}
      <path d={path} />
    </svg>
  );
};
// [AI-BLOCK-END]
```

#### 10.3.2 统一图标组件

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
import React from 'react';
import { SvgIcon } from './SvgIcon';
import { FontIcon } from './FontIcon';
import { iconRegistry } from '../utils/iconRegistry';

export const Icon: React.FC<IconProps> = ({
  name,
  type,
  size = 'md',
  color = 'currentColor',
  variant = 'outlined',
  className,
  ariaHidden = true,
  title,
  testId,
  ...props
}) => {
  // 从图标注册表获取图标数据
  const iconData = iconRegistry[name];
  
  if (!iconData) {
    console.warn(`Icon "${name}" not found in icon registry`);
    return null;
  }
  
  // 确定图标类型
  const iconType = type || iconData.type || 'svg';
  
  if (iconType === 'svg') {
    return (
      <SvgIcon
        path={iconData.path}
        viewBox={iconData.viewBox || '0 0 24 24'}
        size={size}
        color={color}
        variant={variant}
        className={className}
        ariaHidden={ariaHidden}
        title={title}
        testId={testId}
        {...props}
      />
    );
  }
  
  if (iconType === 'font') {
    return (
      <FontIcon
        name={name}
        fontFamily={iconData.fontFamily}
        size={size}
        color={color}
        variant={variant}
        className={className}
        ariaHidden={ariaHidden}
        title={title}
        testId={testId}
        {...props}
      />
    );
  }
  
  return null;
};
// [AI-BLOCK-END]
```

#### 10.3.3 图标注册工具

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
// iconRegistry.ts
interface IconData {
  path?: string;
  viewBox?: string;
  type?: 'svg' | 'font';
  fontFamily?: string;
  unicode?: string;
}

type IconRegistry = Record<string, IconData>;

// 图标注册表
export const iconRegistry: IconRegistry = {
  // 基础UI图标
  'arrow-right': {
    type: 'svg',
    path: 'M13.5 4.5L21 12m0 0l-7.5 7.5M21 12H3',
    viewBox: '0 0 24 24'
  },
  'check': {
    type: 'svg',
    path: 'M4.5 12.75l6 6 9-13.5',
    viewBox: '0 0 24 24'
  },
  'close': {
    type: 'svg',
    path: 'M6 18L18 6M6 6l12 12',
    viewBox: '0 0 24 24'
  },
  
  // 更多图标...
};

// 注册新图标
export function registerIcon(name: string, data: IconData): void {
  iconRegistry[name] = data;
}

// 批量注册图标
export function registerIcons(icons: Record<string, IconData>): void {
  Object.entries(icons).forEach(([name, data]) => {
    registerIcon(name, data);
  });
}
// [AI-BLOCK-END]
```

### 10.4 图标使用指南 (Icon Usage Guidelines)

#### 10.4.1 图标可访问性

1. **提供替代文本** - 对于具有语义意义的图标，应提供适当的aria-label或title
2. **装饰性图标** - 纯装饰性图标应设置`aria-hidden="true"`
3. **交互式图标** - 可点击图标应当是真正的按钮或链接元素，具有适当的键盘焦点状态
4. **图标大小** - 确保交互式图标的触摸目标至少为44×44像素
5. **颜色对比度** - 图标颜色应当满足WCAG 2.1 AA标准的4.5:1对比度要求

#### 10.4.2 图标用法最佳实践

| 最佳实践 | 说明 | 示例 |
|---------|------|------|
| 语义一致性 | 图标应一致地表示相同的概念 | 在整个应用中，"保存"始终使用相同的图标 |
| 简洁性 | 使用最简单的形式传达概念 | 使用简单的"心形"而非复杂的"心形加星形" |
| 搭配文本标签 | 单独的图标可能含义不明，应与文本标签配合 | 菜单项同时显示图标和文本 |
| 避免过度使用 | 过多图标会导致视觉干扰 | 每个区域保持适度的图标数量 |
| 图标间距 | 图标与其相关文本保持适当间距 | 图标和文本之间保持8px间距 |

#### 10.4.3 常见误用示例

❌ **风格混搭** - 在同一界面中混合使用不同风格的图标（如填充式与轮廓式）
❌ **尺寸不一** - 相似场景中使用不一致的图标尺寸
❌ **密度过高** - 在有限空间中堆积过多图标
❌ **缺少标签** - 使用含义不明确的图标而没有文本说明
❌ **过度复杂** - 在小尺寸下使用过于复杂的图标导致识别困难

### 10.5 图标样式变体 (Icon Style Variants)

#### 10.5.1 样式变体定义

| 变体类型 | 描述 | 适用场景 |
|---------|------|---------|
| 轮廓型 (Outlined) | 只有线条，无填充 | 默认样式、大多数UI场景、次要操作 |
| 填充型 (Filled) | 完全填充，边界清晰 | 强调、选中状态、主要操作 |
| 双色型 (Duotone) | 轮廓线加半透明填充 | 信息图表、特殊强调、分层信息 |
| 细线型 (Thin) | 细线条，轻量视觉效果 | 密集信息、次要界面元素、背景元素 |

#### 10.5.2 样式变体实现

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
import React from 'react';
import { Icon } from './Icon';

interface IconExamplesProps {
  name: string;
  label?: string;
}

export const IconVariants: React.FC<IconExamplesProps> = ({
  name,
  label
}) => {
  return (
    <div className="flex flex-col space-y-4">
      {label && <h4 className="text-sm font-medium text-gray-700">{label}</h4>}
      
      <div className="flex space-x-6">
        <div className="flex flex-col items-center">
          <Icon name={name} variant="outlined" size="md" />
          <span className="mt-2 text-xs text-gray-500">轮廓型</span>
        </div>
        
        <div className="flex flex-col items-center">
          <Icon name={name} variant="filled" size="md" />
          <span className="mt-2 text-xs text-gray-500">填充型</span>
        </div>
        
        <div className="flex flex-col items-center">
          <Icon name={name} variant="duotone" size="md" />
          <span className="mt-2 text-xs text-gray-500">双色型</span>
        </div>
        
        <div className="flex flex-col items-center">
          <Icon name={name} variant="thin" size="md" />
          <span className="mt-2 text-xs text-gray-500">细线型</span>
        </div>
      </div>
    </div>
  );
};
// [AI-BLOCK-END]
```

#### 10.5.3 变体选择指南

1. **界面层次** - 使用填充型图标强调主要操作，轮廓型用于次要操作
2. **视觉密度** - 在信息密集区域使用轮廓型或细线型，避免视觉过载
3. **状态表示** - 可以使用不同变体表示不同状态（如选中状态使用填充型）
4. **一致性** - 在功能相似的区域保持变体一致性
5. **品牌调性** - 根据品牌风格选择主要变体类型

### 10.6 响应式图标 (Responsive Icons)

#### 10.6.1 响应式尺寸调整

| 断点 | 推荐尺寸调整 | Tailwind 类示例 |
|------|------------|---------------|
| xs (默认) | 基础尺寸，通常较小 | `w-5 h-5` |
| sm (640px+) | 与默认相同或略大 | `sm:w-5 sm:h-5` |
| md (768px+) | 略大，增强可见性 | `md:w-6 md:h-6` |
| lg (1024px+) | 更大，适合桌面显示 | `lg:w-6 lg:h-6` |
| xl (1280px+) | 视情况可增大或保持 | `xl:w-7 xl:h-7` |
| 2xl (1536px+) | 最大尺寸，大型显示器 | `2xl:w-8 2xl:h-8` |

#### 10.6.2 响应式图标组件实现

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
import React from 'react';
import { classNames } from '../utils';

interface ResponsiveIconProps extends IconProps {
  sizes?: {
    xs?: IconSize;
    sm?: IconSize;
    md?: IconSize;
    lg?: IconSize;
    xl?: IconSize;
    '2xl'?: IconSize;
  };
}

export const ResponsiveIcon: React.FC<ResponsiveIconProps> = ({
  name,
  size = 'md',
  sizes,
  ...props
}) => {
  // 根据尺寸获取像素值
  const getSizePx = (sizeValue: IconSize): number => {
    const sizeMap: Record<IconSize, number> = {
      'xs': 16,
      'sm': 20,
      'md': 24,
      'lg': 32,
      'xl': 48,
      '2xl': 64
    };
    return sizeMap[sizeValue];
  };
  
  // 计算响应式类名
  const getResponsiveClasses = () => {
    if (!sizes) {
      return `w-[${getSizePx(size)}px] h-[${getSizePx(size)}px]`;
    }
    
    const baseSize = sizes.xs || size;
    let classes = `w-[${getSizePx(baseSize)}px] h-[${getSizePx(baseSize)}px]`;
    
    if (sizes.sm) {
      classes += ` sm:w-[${getSizePx(sizes.sm)}px] sm:h-[${getSizePx(sizes.sm)}px]`;
    }
    
    if (sizes.md) {
      classes += ` md:w-[${getSizePx(sizes.md)}px] md:h-[${getSizePx(sizes.md)}px]`;
    }
    
    if (sizes.lg) {
      classes += ` lg:w-[${getSizePx(sizes.lg)}px] lg:h-[${getSizePx(sizes.lg)}px]`;
    }
    
    if (sizes.xl) {
      classes += ` xl:w-[${getSizePx(sizes.xl)}px] xl:h-[${getSizePx(sizes.xl)}px]`;
    }
    
    if (sizes['2xl']) {
      classes += ` 2xl:w-[${getSizePx(sizes['2xl'])}px] 2xl:h-[${getSizePx(sizes['2xl'])}px]`;
    }
    
    return classes;
  };
  
  return (
    <Icon 
      name={name}
      className={classNames(getResponsiveClasses(), props.className)}
      {...props}
    />
  );
};
// [AI-BLOCK-END]
```

#### 10.6.3 使用示例

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
import { ResponsiveIcon } from './ResponsiveIcon';

const ResponsiveIconExample = () => {
  return (
    <header className="flex items-center justify-between p-4 bg-white shadow">
      <div className="flex items-center">
        {/* 响应式图标：移动设备上小，桌面大 */}
        <ResponsiveIcon
          name="logo"
          sizes={{
            xs: 'md',  // 移动设备上 24px
            md: 'lg',  // 平板上 32px
            lg: 'xl'   // 桌面上 48px
          }}
          color="primary"
          className="mr-2"
        />
        <h1 className="text-lg md:text-xl font-bold">应用标题</h1>
      </div>
      
      <nav className="flex items-center space-x-2 md:space-x-4">
        {/* 导航图标，在不同屏幕上有适当大小 */}
        <ResponsiveIcon
          name="home"
          sizes={{
            xs: 'sm',  // 移动设备上 20px
            md: 'md'   // 大屏幕上 24px
          }}
          className="text-gray-600 hover:text-gray-900"
        />
        
        <ResponsiveIcon
          name="settings"
          sizes={{
            xs: 'sm',
            md: 'md'
          }}
          className="text-gray-600 hover:text-gray-900"
        />
        
        <ResponsiveIcon
          name="notifications"
          sizes={{
            xs: 'sm',
            md: 'md'
          }}
          className="text-gray-600 hover:text-gray-900"
        />
      </nav>
    </header>
  );
};
// [AI-BLOCK-END]
```

### 10.7 图标按钮组件 (Icon Button Component)

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
import React from 'react';
import { Icon } from './Icon';
import { classNames } from '../utils';

export const IconButton: React.FC<IconButtonProps> = ({
  icon,
  size = 'md',
  variant = 'primary',
  ariaLabel,
  onClick,
  disabled = false,
  className,
  children,
  ...props
}) => {
  // 尺寸映射
  const sizeClasses = {
    sm: 'p-1.5',
    md: 'p-2',
    lg: 'p-2.5'
  };
  
  // 图标尺寸映射
  const iconSizeMap: Record<string, IconSize> = {
    sm: 'xs',   // 16px icon in small button
    md: 'sm',   // 20px icon in medium button
    lg: 'md'    // 24px icon in large button
  };
  
  // 变体样式
  const variantClasses = {
    primary: 'bg-primary-500 text-white hover:bg-primary-600 focus:ring-primary-500',
    secondary: 'bg-gray-100 text-gray-700 hover:bg-gray-200 focus:ring-gray-500',
    tertiary: 'bg-transparent text-gray-700 hover:bg-gray-100 focus:ring-gray-500',
    ghost: 'bg-transparent text-gray-500 hover:text-gray-700 hover:bg-gray-50 focus:ring-gray-500'
  };
  
  return (
    <button
      type="button"
      onClick={onClick}
      disabled={disabled}
      aria-label={ariaLabel}
      className={classNames(
        'inline-flex items-center justify-center rounded-full focus:outline-none focus:ring-2 focus:ring-offset-2',
        sizeClasses[size],
        variantClasses[variant],
        disabled && 'opacity-50 cursor-not-allowed',
        className
      )}
      {...props}
    >
      <Icon
        name={icon}
        size={iconSizeMap[size]}
        aria-hidden="true"
      />
      {children}
    </button>
  );
};
// [AI-BLOCK-END]
```

### 10.8 最佳实践和总结

1. **一致性优先** - 在整个应用中保持图标样式、尺寸和用法一致
2. **语义重要性** - 图标应当增强而非替代文本，确保用户可以理解图标含义
3. **可访问性** - 始终考虑视觉障碍用户，提供适当的替代文本和足够的视觉对比度
4. **性能考虑** - 使用SVG图标并考虑捆绑策略，或实现图标懒加载机制
5. **分组相关图标** - 将功能相关的图标在视觉上和代码结构上组织在一起
6. **定制与扩展** - 建立可扩展的图标系统，允许添加自定义图标而不破坏一致性
7. **版本控制** - 对图标库实施版本控制，避免界面意外变化

图标系统是用户界面设计中的重要元素，它能够以紧凑的形式传达丰富的信息，并通过一致的视觉语言增强用户体验。通过遵循本指南中的原则和实践，可以创建出专业、一致且富有表现力的图标系统。

## 空间设计指南

空间设计是用户界面设计中的关键方面，它通过合理使用间距、边距和留白来创建视觉层次、引导视线流向、提高可读性，并增强整体用户体验。本节详细介绍如何在 React + TypeScript + Tailwind CSS 项目中实现一致且有效的空间设计。

### 1. 间距比例系统

间距系统基于 4px 作为基本单位，创建一个一致的比例系统。这种系统确保界面中的所有间距遵循相同的视觉节奏。

#### 1.1 基础间距比例

| 尺寸级别 | 像素值 | Tailwind 类 | 使用场景 |
|---------|-------|------------|---------|
| 极小 (2xs) | 2px | p-0.5, m-0.5, gap-0.5 | 紧凑型视觉元素内部间距 |
| 超小 (xs) | 4px | p-1, m-1, gap-1 | 极紧凑型界面元素间距 |
| 小 (sm) | 8px | p-2, m-2, gap-2 | 相关元素内部间距 |
| 中 (md) | 16px | p-4, m-4, gap-4 | 标准元素间距，默认值 |
| 大 (lg) | 24px | p-6, m-6, gap-6 | 分组元素间隔 |
| 超大 (xl) | 32px | p-8, m-8, gap-8 | 主要区域分隔 |
| 巨大 (2xl) | 48px | p-12, m-12, gap-12 | 布局部分分隔 |
| 特大 (3xl) | 64px | p-16, m-16, gap-16 | 主要布局区块间距 |

#### 1.2 自定义间距规模扩展

项目中可以在 Tailwind 配置中扩展间距比例：

```javascript
// tailwind.config.js
module.exports = {
  theme: {
    extend: {
      spacing: {
        // 标准比例扩展
        '18': '4.5rem',  // 72px
        '22': '5.5rem',  // 88px
        
        // 特殊用途间距
        'section': '5rem',      // 80px，区段间距
        'hero': '10rem',        // 160px，大型横幅区间距
        'content-block': '3rem' // 48px，内容块间距
      }
    }
  }
}
```

### 2. 组件密度设计

密度设计通过调整间距和尺寸来控制界面的紧凑程度，适应不同用户需求和使用场景。

#### 2.1 密度级别定义

| 密度级别 | 定义 | 适用场景 |
|---------|------|---------|
| 紧凑型 (Compact) | 减少内部填充和间距，增加视野内信息量 | 数据密集型应用、专业工具、有经验用户 |
| 舒适型 (Comfortable) | 平衡的间距，接近设计系统默认值 | 大多数应用场景，为默认设置 |
| 宽松型 (Loose) | 增加内部填充和间距，提高清晰度 | 入门级应用、触摸优化界面、可访问性需求 |

#### 2.2 密度系统实现

使用 CSS 变量实现可切换的密度系统：

```css
:root {
  /* 舒适型（默认）*/
  --spacing-inset-xs: 0.25rem;  /* 4px */
  --spacing-inset-sm: 0.5rem;   /* 8px */
  --spacing-inset-md: 1rem;     /* 16px */
  --spacing-inset-lg: 1.5rem;   /* 24px */
  --spacing-stack-xs: 0.25rem;  /* 垂直间距 4px */
  --spacing-stack-sm: 0.5rem;   /* 垂直间距 8px */
  --spacing-stack-md: 1rem;     /* 垂直间距 16px */
  --spacing-stack-lg: 1.5rem;   /* 垂直间距 24px */
  --spacing-inline-xs: 0.25rem; /* 水平间距 4px */
  --spacing-inline-sm: 0.5rem;  /* 水平间距 8px */
  --spacing-inline-md: 1rem;    /* 水平间距 16px */
  --spacing-inline-lg: 1.5rem;  /* 水平间距 24px */
  
  /* 控件尺寸 */
  --control-height-sm: 1.75rem; /* 28px */
  --control-height-md: 2.25rem; /* 36px */
  --control-height-lg: 2.75rem; /* 44px */
}

/* 紧凑型 */
.density-compact {
  --spacing-inset-xs: 0.125rem; /* 2px */
  --spacing-inset-sm: 0.25rem;  /* 4px */
  --spacing-inset-md: 0.5rem;   /* 8px */
  --spacing-inset-lg: 1rem;     /* 16px */
  --spacing-stack-xs: 0.125rem; /* 2px */
  --spacing-stack-sm: 0.25rem;  /* 4px */
  --spacing-stack-md: 0.5rem;   /* 8px */
  --spacing-stack-lg: 1rem;     /* 16px */
  --spacing-inline-xs: 0.125rem;/* 2px */
  --spacing-inline-sm: 0.25rem; /* 4px */
  --spacing-inline-md: 0.5rem;  /* 8px */
  --spacing-inline-lg: 1rem;    /* 16px */
  
  --control-height-sm: 1.5rem;  /* 24px */
  --control-height-md: 2rem;    /* 32px */
  --control-height-lg: 2.5rem;  /* 40px */
}

/* 宽松型 */
.density-loose {
  --spacing-inset-xs: 0.5rem;   /* 8px */
  --spacing-inset-sm: 0.75rem;  /* 12px */
  --spacing-inset-md: 1.5rem;   /* 24px */
  --spacing-inset-lg: 2rem;     /* 32px */
  --spacing-stack-xs: 0.5rem;   /* 8px */
  --spacing-stack-sm: 0.75rem;  /* 12px */
  --spacing-stack-md: 1.5rem;   /* 24px */
  --spacing-stack-lg: 2rem;     /* 32px */
  --spacing-inline-xs: 0.5rem;  /* 8px */
  --spacing-inline-sm: 0.75rem; /* 12px */
  --spacing-inline-md: 1.5rem;  /* 24px */
  --spacing-inline-lg: 2rem;    /* 32px */
  
  --control-height-sm: 2rem;    /* 32px */
  --control-height-md: 2.5rem;  /* 40px */
  --control-height-lg: 3rem;    /* 48px */
}
```

#### 2.3 密度系统 React 组件示例

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
import React from 'react';
import { classNames } from '../utils';

type Density = 'compact' | 'comfortable' | 'loose';

interface DensityProviderProps {
  density: Density;
  children: React.ReactNode;
}

export const DensityContext = React.createContext<Density>('comfortable');

export const DensityProvider: React.FC<DensityProviderProps> = ({ 
  density = 'comfortable',
  children 
}) => {
  return (
    <DensityContext.Provider value={density}>
      <div className={`density-${density}`}>
        {children}
      </div>
    </DensityContext.Provider>
  );
};

export const useDensity = () => {
  return React.useContext(DensityContext);
};

// 使用密度上下文的按钮组件示例
export const Button: React.FC<{
  size?: 'sm' | 'md' | 'lg';
  children: React.ReactNode;
  className?: string;
}> = ({ size = 'md', children, className }) => {
  const density = useDensity();
  
  // 基于密度和尺寸获取适当的间距类
  const getPaddingClass = () => {
    const baseHorizontal = {
      sm: 'px-2',
      md: 'px-3',
      lg: 'px-4'
    };
    
    const densityModifiers = {
      compact: -1,
      comfortable: 0,
      loose: 1
    };
    
    // 根据密度调整尺寸
    const horizontalSizeIndex = Math.max(
      0, 
      Object.keys(baseHorizontal).indexOf(size) + densityModifiers[density]
    );
    const horizontalSize = Object.keys(baseHorizontal)[horizontalSizeIndex] || size;
    
    return baseHorizontal[horizontalSize as keyof typeof baseHorizontal];
  };
  
  // 基于密度和尺寸获取适当的高度类
  const getHeightClass = () => {
    const baseHeight = {
      sm: 'h-8',
      md: 'h-10',
      lg: 'h-12'
    };
    
    const densityHeight = {
      compact: {
        sm: 'h-6',
        md: 'h-8',
        lg: 'h-10'
      },
      comfortable: baseHeight,
      loose: {
        sm: 'h-10',
        md: 'h-12',
        lg: 'h-14'
      }
    };
    
    return densityHeight[density][size as keyof typeof baseHeight];
  };
  
  return (
    <button
      className={classNames(
        'inline-flex items-center justify-center rounded-md font-medium focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500',
        getPaddingClass(),
        getHeightClass(),
        className
      )}
    >
      {children}
    </button>
  );
};
// [AI-BLOCK-END]
```

### 3. 空白空间使用指南

空白空间（也称为留白）是指界面中没有内容的区域。正确使用空白空间可以提高可读性、引导用户注意力，并创建清晰的视觉层次。

#### 3.1 空白空间原则

| 原则 | 说明 | 实现方法 |
|------|------|---------|
| 一致性 | 相似元素使用一致的间距 | 使用间距比例系统中的固定值 |
| 意图性 | 有意识地使用间距创建关系 | 相关元素间距小，无关元素间距大 |
| 层次感 | 使用间距创建视觉层次 | 标题与内容使用更大间距，段落间使用中等间距 |
| 呼吸空间 | 为复杂内容提供足够空间 | 在内容密集区域周围增加边距 |
| 对比 | 使用空间创造对比，强调重要内容 | 重要信息周围留出更多空白 |

#### 3.2 空白空间的类型

1. **微观空白** - 文字、图标、小组件内部和周围的小空间
   - 行高：`leading-tight`、`leading-normal`、`leading-relaxed`
   - 字间距：`tracking-tight`、`tracking-normal`、`tracking-wide`
   - 内部填充：`p-1`、`p-2`、`p-3`

2. **宏观空白** - 主要布局元素之间的大空间
   - 区段间距：`my-8`、`my-12`、`my-16`
   - 布局边距：`container mx-auto px-4 md:px-6 lg:px-8`
   - 分组间距：`space-y-6`、`space-y-8`、`space-y-10`

#### 3.3 响应式空白空间

空白空间应根据屏幕尺寸调整：

```tsx
<section className="py-6 md:py-8 lg:py-12">
  <div className="container mx-auto px-4 md:px-6 lg:px-8">
    <h2 className="text-2xl font-bold mb-4 md:mb-6 lg:mb-8">
      标题
    </h2>
    <div className="space-y-4 md:space-y-6">
      <p>内容段落</p>
      <p>内容段落</p>
    </div>
  </div>
</section>
```

### 4. 容器边距系统

容器边距用于在不同视口宽度下控制内容的水平边距，确保最佳的可读性和美观性。

#### 4.1 基本容器模式

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
import React from 'react';
import { classNames } from '../utils';

interface ContainerProps {
  maxWidth?: 'xs' | 'sm' | 'md' | 'lg' | 'xl' | '2xl' | 'full' | 'none';
  padding?: boolean;
  className?: string;
  children: React.ReactNode;
}

export const Container: React.FC<ContainerProps> = ({
  maxWidth = 'lg',
  padding = true,
  className,
  children
}) => {
  const maxWidthClasses = {
    xs: 'max-w-xs',        // 320px
    sm: 'max-w-sm',        // 384px
    md: 'max-w-md',        // 448px
    lg: 'max-w-lg',        // 512px
    xl: 'max-w-xl',        // 576px
    '2xl': 'max-w-2xl',    // 672px
    full: 'max-w-full',    // 100%
    none: ''               // 无限制
  };
  
  return (
    <div
      className={classNames(
        'mx-auto w-full',
        maxWidthClasses[maxWidth],
        padding && 'px-4 sm:px-6 lg:px-8',
        className
      )}
    >
      {children}
    </div>
  );
};
// [AI-BLOCK-END]
```

#### 4.2 内容宽度指南

| 内容类型 | 最大宽度 | 解释 |
|---------|----------|------|
| 纯文本内容 | `max-w-prose` (65ch) | 单行约65个字符的宽度，最佳可读性 |
| 表单 | `max-w-md` (28rem) | 适中的表单宽度，减少视线移动 |
| 卡片网格 | `max-w-7xl` (80rem) | 宽内容但有边界，适合卡片网格布局 |
| 全幅布局 | `max-w-full` (100%) | 占满全宽，适合复杂数据表格 |

#### 4.3 多列布局的间距

```tsx
<div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-x-4 md:gap-x-6 lg:gap-x-8 gap-y-6 md:gap-y-8 lg:gap-y-12">
  {/* 列内容 */}
</div>
```

### 5. 空间模式和常见用例

#### 5.1 堆叠模式 (Stack)

用于在垂直方向上排列元素，带有一致的间距。

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
import React from 'react';
import { classNames } from '../utils';

interface StackProps {
  space?: 'none' | 'xs' | 'sm' | 'md' | 'lg' | 'xl';
  dividers?: boolean;
  className?: string;
  children: React.ReactNode;
}

export const Stack: React.FC<StackProps> = ({
  space = 'md',
  dividers = false,
  className,
  children
}) => {
  const spaceClasses = {
    none: 'space-y-0',
    xs: 'space-y-1',
    sm: 'space-y-2',
    md: 'space-y-4',
    lg: 'space-y-6',
    xl: 'space-y-8'
  };
  
  return (
    <div
      className={classNames(
        spaceClasses[space],
        dividers && 'divide-y divide-gray-200',
        className
      )}
    >
      {children}
    </div>
  );
};
// [AI-BLOCK-END]
```

#### 5.2 内联组模式 (Inline Group)

用于在水平方向上排列元素，带有一致的间距。

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
import React from 'react';
import { classNames } from '../utils';

interface InlineGroupProps {
  space?: 'none' | 'xs' | 'sm' | 'md' | 'lg';
  wrap?: boolean;
  align?: 'start' | 'center' | 'end' | 'baseline';
  className?: string;
  children: React.ReactNode;
}

export const InlineGroup: React.FC<InlineGroupProps> = ({
  space = 'md',
  wrap = true,
  align = 'center',
  className,
  children
}) => {
  const spaceClasses = {
    none: 'space-x-0',
    xs: 'space-x-1',
    sm: 'space-x-2',
    md: 'space-x-4',
    lg: 'space-x-6'
  };
  
  const alignClasses = {
    start: 'items-start',
    center: 'items-center',
    end: 'items-end',
    baseline: 'items-baseline'
  };
  
  return (
    <div
      className={classNames(
        'flex',
        wrap ? 'flex-wrap' : 'flex-nowrap',
        spaceClasses[space],
        alignClasses[align],
        className
      )}
    >
      {children}
    </div>
  );
};
// [AI-BLOCK-END]
```

#### 5.3 表单间距模式

表单元素的一致间距处理。

```tsx
<form className="space-y-6">
  <div className="space-y-1">
    <label className="block text-sm font-medium text-gray-700">
      姓名
    </label>
    <input type="text" className="mt-1 block w-full rounded-md border-gray-300" />
  </div>

  <div className="space-y-1">
    <label className="block text-sm font-medium text-gray-700">
      电子邮件
    </label>
    <input type="email" className="mt-1 block w-full rounded-md border-gray-300" />
    <p className="mt-1 text-sm text-gray-500">
      我们不会分享您的电子邮件
    </p>
  </div>
  
  <div className="pt-4">
    <button type="submit" className="px-4 py-2 bg-blue-600 text-white rounded-md">
      提交
    </button>
  </div>
</form>
```

### 6. 视觉层次与空间

通过空间创建视觉层次，引导用户视线和关注点。

#### 6.1 空间层次原则

1. **分组亲近性** - 相关元素靠近，无关元素距离更远
2. **垂直节奏** - 一致的垂直间距创造可预测的布局流
3. **焦点分离** - 重要元素周围的更多空间吸引注意力
4. **分层逻辑** - 顺序间距转换表达内容的层次结构

#### 6.2 层次间距实现示例

```tsx
<article className="max-w-prose mx-auto">
  {/* 大间距隔开主要区块 */}
  <header className="mb-8">
    <h1 className="text-3xl font-bold mb-2">文章标题</h1>
    <p className="text-gray-600">发布于 2023年4月12日</p>
  </header>

  <div className="space-y-6">
    {/* 中等间距隔开部分 */}
    <section className="space-y-4">
      <h2 className="text-2xl font-semibold">第一部分</h2>
      {/* 小间距隔开段落 */}
      <p className="mb-4">段落内容...</p>
      <p>段落内容...</p>
    </section>

    <section className="space-y-4">
      <h2 className="text-2xl font-semibold">第二部分</h2>
      <p className="mb-4">段落内容...</p>
      <p>段落内容...</p>
    </section>
  </div>
  
  {/* 大间距隔开主要区块 */}
  <footer className="mt-12 pt-6 border-t border-gray-200">
    <p className="text-gray-500">作者信息</p>
  </footer>
</article>
```

### 7. 响应式空间设计

在不同屏幕尺寸上调整间距，确保最佳用户体验。

#### 7.1 响应式间距规则

1. **移动优先** - 从最小屏幕开始设计，然后扩展
2. **比例缩放** - 大屏幕上增加间距，小屏幕上减少间距
3. **保持可读性** - 内容周围的间距应确保最佳可读性
4. **优先保留垂直空间** - 在小屏幕上，优先保留垂直间距

#### 7.2 响应式间距实现

```tsx
<section>
  {/* 响应式容器边距 */}
  <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
    {/* 响应式部分间距 */}
    <div className="py-6 md:py-8 lg:py-12">
      {/* 响应式标题间距 */}
      <h2 className="text-2xl font-bold mb-4 md:mb-6 lg:mb-8">
        标题
      </h2>
      
      {/* 响应式卡片网格 */}
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4 md:gap-6 lg:gap-8">
        {/* 卡片 */}
        <div className="bg-white p-4 md:p-6 rounded-md shadow">
          卡片内容
        </div>
        {/* 更多卡片... */}
      </div>
    </div>
  </div>
</section>
``` 