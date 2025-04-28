# 排版系统设计指南

## 概述

排版系统是设计系统的核心组成部分，它定义了文本的视觉层次结构和信息架构。本文档详细说明了基于React + TypeScript + Tailwind CSS项目的排版系统，包括字体族、字重、字号、行高和使用规范。

## 字体族系统

### 主要字体（Sans-serif）

用于UI界面主要文本、标题、按钮等元素。

```css
font-family: 'Inter', system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
```

Tailwind类：`font-sans`

### 显示字体（Display）

用于较大的标题、首屏文案、特殊强调等元素。

```css
font-family: 'Lexend', system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
```

Tailwind类：`font-display`

### 等宽字体（Monospace）

用于代码展示、技术指标、数据表格等需要等宽字符的场景。

```css
font-family: 'Fira Code', SFMono-Regular, Menlo, Monaco, Consolas, 'Liberation Mono', 'Courier New', monospace;
```

Tailwind类：`font-mono`

## 字重系统

| 字重名称 | 数值 | Tailwind类 | 使用场景 |
|---------|------|------------|---------|
| 超细体 (Thin) | 100 | font-thin | 极大型标题、装饰性文本 |
| 超轻体 (Extra Light) | 200 | font-extralight | 大号标题、特殊强调 |
| 轻体 (Light) | 300 | font-light | 较大标题、次要文本 |
| 常规体 (Regular) | 400 | font-normal | 正文文本、描述文本 |
| 中等体 (Medium) | 500 | font-medium | 小标题、按钮文本、强调文本 |
| 半粗体 (Semi Bold) | 600 | font-semibold | 标题、导航项、重要文本 |
| 粗体 (Bold) | 700 | font-bold | 主要标题、特别强调 |
| 特粗体 (Extra Bold) | 800 | font-extrabold | 超大标题、极度强调 |
| 黑体 (Black) | 900 | font-black | 特殊场景、极大型标题 |

## 字号与行高系统

字号系统采用可缩放的方式，确保在不同屏幕尺寸上都有良好的可读性。

| 名称 | 字号 | 行高 | Tailwind类 | 使用场景 |
|------|------|------|------------|---------|
| xs | 0.75rem (12px) | 1rem (16px) | text-xs | 辅助文本、标签、角标 |
| sm | 0.875rem (14px) | 1.25rem (20px) | text-sm | 次要文本、表单标签、说明文本 |
| base | 1rem (16px) | 1.5rem (24px) | text-base | 正文文本、按钮文本、导航项 |
| lg | 1.125rem (18px) | 1.75rem (28px) | text-lg | 小标题、强调文本、卡片标题 |
| xl | 1.25rem (20px) | 1.75rem (28px) | text-xl | 中标题、部分标题、面板标题 |
| 2xl | 1.5rem (24px) | 2rem (32px) | text-2xl | 大标题、页面标题 |
| 3xl | 1.875rem (30px) | 2.25rem (36px) | text-3xl | 特大标题、首屏标题 |
| 4xl | 2.25rem (36px) | 2.5rem (40px) | text-4xl | 超大标题、登陆页标题 |
| 5xl | 3rem (48px) | 1 (48px) | text-5xl | 英雄区标题、特殊场景 |
| 6xl | 3.75rem (60px) | 1 (60px) | text-6xl | 重点强调、超大展示 |
| 7xl | 4.5rem (72px) | 1 (72px) | text-7xl | 页面横幅、主视觉 |
| 8xl | 6rem (96px) | 1 (96px) | text-8xl | 装饰性标题、营销场景 |
| 9xl | 8rem (128px) | 1 (128px) | text-9xl | 特殊设计场景 |

## 字间距系统

定义字母之间的间距，影响文本的紧凑程度和可读性。

| 名称 | 值 | Tailwind类 | 使用场景 |
|------|------|------------|---------|
| 紧凑 | -0.05em | tracking-tighter | 大号标题、特殊设计 |
| 紧凑 | -0.025em | tracking-tight | 标题、紧凑排版 |
| 正常 | 0em | tracking-normal | 正文文本、默认设置 |
| 宽松 | 0.025em | tracking-wide | 小号文本、强调文本 |
| 更宽松 | 0.05em | tracking-wider | 全大写文本、按钮文本 |
| 最宽松 | 0.1em | tracking-widest | 装饰性文本、徽章 |

## 排版变体系统

为了便于使用统一的排版样式，我们定义了常用的排版变体，每个变体包含完整的字体、字重、字号、行高和颜色设置。

### TypeScript 实现

```typescript
type TypographyVariant = 
  | 'h1' | 'h2' | 'h3' | 'h4' | 'h5' | 'h6'
  | 'subtitle1' | 'subtitle2'
  | 'body1' | 'body2'
  | 'button' | 'caption' | 'overline' | 'code';

const getTypographyClasses = (variant: TypographyVariant): string => {
  switch (variant) {
    case 'h1':
      return 'font-display text-4xl font-bold text-gray-900 dark:text-white leading-tight tracking-tight';
    case 'h2':
      return 'font-display text-3xl font-semibold text-gray-900 dark:text-white leading-tight tracking-tight';
    case 'h3':
      return 'font-display text-2xl font-semibold text-gray-900 dark:text-white leading-tight';
    case 'h4':
      return 'font-sans text-xl font-semibold text-gray-900 dark:text-white leading-snug';
    case 'h5':
      return 'font-sans text-lg font-medium text-gray-900 dark:text-white leading-snug';
    case 'h6':
      return 'font-sans text-base font-medium text-gray-900 dark:text-white leading-normal';
    case 'subtitle1':
      return 'font-sans text-lg font-normal text-gray-700 dark:text-gray-300 leading-normal';
    case 'subtitle2':
      return 'font-sans text-base font-medium text-gray-700 dark:text-gray-300 leading-normal';
    case 'body1':
      return 'font-sans text-base font-normal text-gray-700 dark:text-gray-300 leading-relaxed';
    case 'body2':
      return 'font-sans text-sm font-normal text-gray-700 dark:text-gray-300 leading-relaxed';
    case 'button':
      return 'font-sans text-sm font-medium text-current uppercase tracking-wide';
    case 'caption':
      return 'font-sans text-xs font-normal text-gray-500 dark:text-gray-400 leading-normal';
    case 'overline':
      return 'font-sans text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider';
    case 'code':
      return 'font-mono text-sm font-normal text-gray-800 dark:text-gray-200 bg-gray-100 dark:bg-gray-800 px-1 py-0.5 rounded';
    default:
      return 'font-sans text-base font-normal text-gray-700 dark:text-gray-300';
  }
};

// 使用示例
const Typography: React.FC<{
  variant?: TypographyVariant;
  className?: string;
  children: React.ReactNode;
}> = ({ variant = 'body1', className = '', children }) => {
  const Element = getElementFromVariant(variant);
  return (
    <Element className={`${getTypographyClasses(variant)} ${className}`}>
      {children}
    </Element>
  );
};

// 根据变体返回合适的HTML元素
const getElementFromVariant = (variant: TypographyVariant): keyof JSX.IntrinsicElements => {
  switch (variant) {
    case 'h1': return 'h1';
    case 'h2': return 'h2';
    case 'h3': return 'h3';
    case 'h4': return 'h4';
    case 'h5': return 'h5';
    case 'h6': return 'h6';
    case 'subtitle1': return 'h6';
    case 'subtitle2': return 'h6';
    case 'body1': return 'p';
    case 'body2': return 'p';
    case 'button': return 'span';
    case 'caption': return 'span';
    case 'overline': return 'span';
    case 'code': return 'code';
    default: return 'p';
  }
};
```

## 排版变体规范

### 标题排版

标题使用从h1到h6的6个级别，每个级别有特定的样式：

| 变体 | Tailwind类组合 | 使用场景 |
|------|------------|---------|
| h1 | font-display text-4xl font-bold | 页面主标题、登陆页标题 |
| h2 | font-display text-3xl font-semibold | 主要区域标题、部分标题 |
| h3 | font-display text-2xl font-semibold | 卡片标题、主要部分标题 |
| h4 | font-sans text-xl font-semibold | 内容区域标题、次级标题 |
| h5 | font-sans text-lg font-medium | 段落标题、子部分标题 |
| h6 | font-sans text-base font-medium | 小标题、列表标题 |

### 正文排版

| 变体 | Tailwind类组合 | 使用场景 |
|------|------------|---------|
| subtitle1 | font-sans text-lg font-normal | 大段落介绍、特性描述 |
| subtitle2 | font-sans text-base font-medium | 小段落介绍、卡片描述 |
| body1 | font-sans text-base font-normal | 主要正文、长文本内容 |
| body2 | font-sans text-sm font-normal | 次要正文、辅助内容 |

### 功能性排版

| 变体 | Tailwind类组合 | 使用场景 |
|------|------------|---------|
| button | font-sans text-sm font-medium uppercase | 按钮文本、交互元素 |
| caption | font-sans text-xs font-normal | 注释、附加说明 |
| overline | font-sans text-xs font-medium uppercase | 小标签、分类标识 |
| code | font-mono text-sm | 代码片段、技术内容 |

## Tailwind 配置实现

在`tailwind.config.js`中添加以下配置来实现自定义排版系统：

```javascript
module.exports = {
  theme: {
    extend: {
      fontFamily: {
        sans: ['"Inter"', 'system-ui', '-apple-system', 'BlinkMacSystemFont', '"Segoe UI"', 'Roboto', '"Helvetica Neue"', 'Arial', 'sans-serif'],
        display: ['"Lexend"', 'system-ui', '-apple-system', 'BlinkMacSystemFont', '"Segoe UI"', 'Roboto', '"Helvetica Neue"', 'Arial', 'sans-serif'],
        mono: ['"Fira Code"', 'SFMono-Regular', 'Menlo', 'Monaco', 'Consolas', '"Liberation Mono"', '"Courier New"', 'monospace'],
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
        '7xl': ['4.5rem', { lineHeight: '1' }],
        '8xl': ['6rem', { lineHeight: '1' }],
        '9xl': ['8rem', { lineHeight: '1' }],
      },
      lineHeight: {
        tighter: '1.1',
        snug: '1.375',
        normal: '1.5',
        relaxed: '1.625',
        loose: '2',
      },
    },
  },
}
```

## 字体设置指南

### 引入字体

为了确保字体可用，应使用以下方式引入字体：

#### 使用Google Fonts（CDN方式）：

在`index.html`中添加：

```html
<link rel="preconnect" href="https://fonts.googleapis.com">
<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
<link href="https://fonts.googleapis.com/css2?family=Inter:wght@100;200;300;400;500;600;700;800;900&family=Lexend:wght@300;400;500;600;700&family=Fira+Code:wght@300;400;500;600;700&display=swap" rel="stylesheet">
```

#### 使用本地字体（性能优化）：

```css
/* styles/fonts.css */
@font-face {
  font-family: 'Inter';
  font-style: normal;
  font-weight: 400;
  font-display: swap;
  src: url('/fonts/Inter-Regular.woff2') format('woff2'),
       url('/fonts/Inter-Regular.woff') format('woff');
}

@font-face {
  font-family: 'Inter';
  font-style: normal;
  font-weight: 500;
  font-display: swap;
  src: url('/fonts/Inter-Medium.woff2') format('woff2'),
       url('/fonts/Inter-Medium.woff') format('woff');
}

/* 添加其他字重和字体族 */
```

## 排版最佳实践

### 可读性建议

- 正文文本保持行高为1.5-1.7，确保良好可读性
- 短行文本（标题等）可以使用较紧凑的行高（1.1-1.3）
- 移动设备上字号不应小于14px
- 确保文本与背景有足够对比度（符合WCAG AA标准）

### 响应式排版

Tailwind提供响应式修饰符来实现不同屏幕尺寸的排版调整：

```html
<h1 class="text-2xl md:text-3xl lg:text-4xl font-bold">响应式标题</h1>
<p class="text-sm md:text-base lg:text-lg">响应式段落文本</p>
```

### 文本截断

对于可能溢出容器的文本，使用以下功能类：

- 单行文本截断：`truncate`
- 多行文本截断：
  ```css
  .line-clamp-2 {
    display: -webkit-box;
    -webkit-line-clamp: 2;
    -webkit-box-orient: vertical;
    overflow: hidden;
  }
  ```

### 特殊文本样式

| 文本样式 | Tailwind类 | 使用场景 |
|---------|------------|---------|
| 下划线 | underline | 链接、强调 |
| 删除线 | line-through | 删除内容、折扣价格 |
| 斜体 | italic | 引述、强调 |
| 全大写 | uppercase | 按钮、标签、短标题 |
| 全小写 | lowercase | 特定品牌、风格化文本 |
| 首字母大写 | capitalize | 标题、专有名词 |

## 无障碍性考虑

确保文本符合无障碍标准：

1. 文本与背景的对比度符合WCAG 2.1 AA标准：
   - 正常文本：至少4.5:1
   - 大文本：至少3:1

2. 避免使用过小的字体，基本文本至少为16px（1rem）

3. 适当的行高和字间距，提高可读性

4. 不仅依赖颜色传达信息（为图标添加文本说明等）

5. 确保文本可以缩放，不会破坏布局

## 排版检查工具

1. Contrast Checker: https://webaim.org/resources/contrastchecker/
2. Typography Checker: https://www.checkmycolours.com/
3. Readable: https://readable.com/ 