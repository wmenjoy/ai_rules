# 色彩系统设计指南

## 概述

色彩系统是整个设计系统的基础，提供一致的视觉语言和品牌识别。本文档详细定义了基于React + TypeScript + Tailwind CSS项目的色彩系统，包括色彩定义、使用规范和Tailwind配置方法。

## 色彩类别

### 主色系统

主色代表品牌的核心色彩，用于主要界面元素、操作按钮和强调区域。

| 色彩名称 | 色值 | Tailwind类 | 使用场景 |
|---------|------|------------|---------|
| 主色-50 | #f0f9ff | bg-primary-50 text-primary-50 | 极浅背景、hover状态 |
| 主色-100 | #e0f2fe | bg-primary-100 text-primary-100 | 浅色背景、选中状态 |
| 主色-200 | #bae6fd | bg-primary-200 text-primary-200 | 浅色边框、分隔元素 |
| 主色-300 | #7dd3fc | bg-primary-300 text-primary-300 | 次要元素、禁用状态 |
| 主色-400 | #38bdf8 | bg-primary-400 text-primary-400 | 次要按钮、图标 |
| 主色-500 | #0ea5e9 | bg-primary-500 text-primary-500 | 主色基准、链接 |
| 主色-600 | #0284c7 | bg-primary-600 text-primary-600 | 主要按钮、强调元素 |
| 主色-700 | #0369a1 | bg-primary-700 text-primary-700 | 主按钮hover状态、深色文本 |
| 主色-800 | #075985 | bg-primary-800 text-primary-800 | 深色强调、按钮active状态 |
| 主色-900 | #0c4a6e | bg-primary-900 text-primary-900 | 最深色元素、特殊强调 |
| 主色-950 | #082f49 | bg-primary-950 text-primary-950 | 最深色背景 |

### 辅助色系统

辅助色用于补充主色，创建视觉层次和区分不同的信息或界面区域。

| 色彩名称 | 色值 | Tailwind类 | 使用场景 |
|---------|------|------------|---------|
| 辅助色-50 | #f5f3ff | bg-secondary-50 text-secondary-50 | 极浅背景 |
| 辅助色-100 | #ede9fe | bg-secondary-100 text-secondary-100 | 浅色背景 |
| 辅助色-200 | #ddd6fe | bg-secondary-200 text-secondary-200 | 浅色边框 |
| 辅助色-300 | #c4b5fd | bg-secondary-300 text-secondary-300 | 次要元素 |
| 辅助色-400 | #a78bfa | bg-secondary-400 text-secondary-400 | 次要按钮 |
| 辅助色-500 | #8b5cf6 | bg-secondary-500 text-secondary-500 | 辅助色基准 |
| 辅助色-600 | #7c3aed | bg-secondary-600 text-secondary-600 | 辅助按钮 |
| 辅助色-700 | #6d28d9 | bg-secondary-700 text-secondary-700 | 按钮hover状态 |
| 辅助色-800 | #5b21b6 | bg-secondary-800 text-secondary-800 | 深色强调 |
| 辅助色-900 | #4c1d95 | bg-secondary-900 text-secondary-900 | 最深色元素 |
| 辅助色-950 | #2e1065 | bg-secondary-950 text-secondary-950 | 最深色背景 |

### 强调色系统

用于特定场景需要额外吸引注意力或区分的元素。

| 色彩名称 | 色值 | Tailwind类 | 使用场景 |
|---------|------|------------|---------|
| 强调色-50 | #fff7ed | bg-accent-50 text-accent-50 | 极浅背景 |
| 强调色-100 | #ffedd5 | bg-accent-100 text-accent-100 | 浅色背景 |
| 强调色-200 | #fed7aa | bg-accent-200 text-accent-200 | 浅色边框 |
| 强调色-300 | #fdba74 | bg-accent-300 text-accent-300 | 次要元素 |
| 强调色-400 | #fb923c | bg-accent-400 text-accent-400 | 次要按钮 |
| 强调色-500 | #f97316 | bg-accent-500 text-accent-500 | 强调色基准 |
| 强调色-600 | #ea580c | bg-accent-600 text-accent-600 | 强调按钮 |
| 强调色-700 | #c2410c | bg-accent-700 text-accent-700 | 按钮hover状态 |
| 强调色-800 | #9a3412 | bg-accent-800 text-accent-800 | 深色强调 |
| 强调色-900 | #7c2d12 | bg-accent-900 text-accent-900 | 最深色元素 |
| 强调色-950 | #431407 | bg-accent-950 text-accent-950 | 最深色背景 |

### 中性/灰度色阶

用于文本、背景、边框等常规界面元素。

| 色彩名称 | 色值 | Tailwind类 | 使用场景 |
|---------|------|------------|---------|
| 灰色-50 | #f9fafb | bg-gray-50 text-gray-50 | 最浅背景、卡片背景 |
| 灰色-100 | #f3f4f6 | bg-gray-100 text-gray-100 | 浅色背景、分隔区域 |
| 灰色-200 | #e5e7eb | bg-gray-200 text-gray-200 | 边框、分割线 |
| 灰色-300 | #d1d5db | bg-gray-300 text-gray-300 | 深色边框、禁用状态 |
| 灰色-400 | #9ca3af | bg-gray-400 text-gray-400 | 禁用文本、次要图标 |
| 灰色-500 | #6b7280 | bg-gray-500 text-gray-500 | 次要文本、说明文字 |
| 灰色-600 | #4b5563 | bg-gray-600 text-gray-600 | 次要标题、重要说明 |
| 灰色-700 | #374151 | bg-gray-700 text-gray-700 | 正文文本 |
| 灰色-800 | #1f2937 | bg-gray-800 text-gray-800 | 主要标题、深色文本 |
| 灰色-900 | #111827 | bg-gray-900 text-gray-900 | 强调标题、最深文本 |
| 灰色-950 | #030712 | bg-gray-950 text-gray-950 | 最深背景 |

### 语义色系统

用于传达特定的信息状态、反馈和交互。

| 色彩名称 | 色值 | Tailwind类 | 使用场景 |
|---------|------|------------|---------|
| 成功色-50 | #f0fdf4 | bg-success-50 text-success-50 | 成功状态浅背景 |
| 成功色-500 | #22c55e | bg-success-500 text-success-500 | 成功图标、文本 |
| 成功色-600 | #16a34a | bg-success-600 text-success-600 | 成功按钮、强调 |
| 警告色-50 | #fffbeb | bg-warning-50 text-warning-50 | 警告状态浅背景 |
| 警告色-500 | #eab308 | bg-warning-500 text-warning-500 | 警告图标、文本 |
| 警告色-600 | #ca8a04 | bg-warning-600 text-warning-600 | 警告按钮、强调 |
| 错误色-50 | #fef2f2 | bg-error-50 text-error-50 | 错误状态浅背景 |
| 错误色-500 | #ef4444 | bg-error-500 text-error-500 | 错误图标、文本 |
| 错误色-600 | #dc2626 | bg-error-600 text-error-600 | 错误按钮、强调 |
| 信息色-50 | #eff6ff | bg-info-50 text-info-50 | 信息状态浅背景 |
| 信息色-500 | #3b82f6 | bg-info-500 text-info-500 | 信息图标、文本 |
| 信息色-600 | #2563eb | bg-info-600 text-info-600 | 信息按钮、强调 |

## 暗色模式映射

为支持暗色模式，定义了以下颜色映射关系：

| 亮色模式 | 暗色模式 |
|---------|---------|
| 灰色-50 | 灰色-950 |
| 灰色-100 | 灰色-900 |
| 灰色-200 | 灰色-800 |
| 灰色-300 | 灰色-700 |
| 灰色-400 | 灰色-600 |
| 灰色-500 | 灰色-500 |
| 灰色-600 | 灰色-400 |
| 灰色-700 | 灰色-300 |
| 灰色-800 | 灰色-200 |
| 灰色-900 | 灰色-100 |
| 灰色-950 | 灰色-50 |

## 文本与背景组合指南

为确保足够的对比度和可访问性，以下是推荐的文本和背景色组合：

| 背景色 | 文本色 | 对比度比率 | WCAG级别 |
|---------|------|------------|---------|
| 白色/灰色-50 | 灰色-900 | 16:1 | AAA |
| 白色/灰色-50 | 灰色-800 | 13.5:1 | AAA |
| 白色/灰色-50 | 灰色-700 | 10:1 | AAA |
| 白色/灰色-50 | 灰色-600 | 7:1 | AAA |
| 白色/灰色-50 | 主色-700 | 8:1 | AAA |
| 主色-600 | 白色 | 4.5:1 | AA |
| 主色-700 | 白色 | 7:1 | AAA |
| 辅助色-600 | 白色 | 4.5:1 | AA |
| 错误色-600 | 白色 | 4.5:1 | AA |

## Tailwind 配置实现

在`tailwind.config.js`中添加以下配置来实现自定义色彩系统：

```javascript
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
        secondary: {
          50: '#f5f3ff',
          100: '#ede9fe',
          200: '#ddd6fe',
          300: '#c4b5fd',
          400: '#a78bfa',
          500: '#8b5cf6',
          600: '#7c3aed',
          700: '#6d28d9',
          800: '#5b21b6',
          900: '#4c1d95',
          950: '#2e1065',
        },
        accent: {
          50: '#fff7ed',
          100: '#ffedd5',
          200: '#fed7aa',
          300: '#fdba74',
          400: '#fb923c',
          500: '#f97316',
          600: '#ea580c',
          700: '#c2410c',
          800: '#9a3412',
          900: '#7c2d12',
          950: '#431407',
        },
        success: {
          50: '#f0fdf4',
          100: '#dcfce7',
          200: '#bbf7d0',
          300: '#86efac',
          400: '#4ade80',
          500: '#22c55e',
          600: '#16a34a',
          700: '#15803d',
          800: '#166534',
          900: '#14532d',
          950: '#052e16',
        },
        warning: {
          50: '#fffbeb',
          100: '#fef3c7',
          200: '#fde68a',
          300: '#fcd34d',
          400: '#fbbf24',
          500: '#eab308',
          600: '#ca8a04',
          700: '#a16207',
          800: '#854d0e',
          900: '#713f12',
          950: '#422006',
        },
        error: {
          50: '#fef2f2',
          100: '#fee2e2',
          200: '#fecaca',
          300: '#fca5a5',
          400: '#f87171',
          500: '#ef4444',
          600: '#dc2626',
          700: '#b91c1c',
          800: '#991b1b',
          900: '#7f1d1d',
          950: '#450a0a',
        },
        info: {
          50: '#eff6ff',
          100: '#dbeafe',
          200: '#bfdbfe',
          300: '#93c5fd',
          400: '#60a5fa',
          500: '#3b82f6',
          600: '#2563eb',
          700: '#1d4ed8',
          800: '#1e40af',
          900: '#1e3a8a',
          950: '#172554',
        },
      }
    }
  },
  // 启用暗色模式
  darkMode: 'class',
}
```

## 色彩使用规范

### 按钮色彩

- 主要按钮：bg-primary-600 hover:bg-primary-700 text-white
- 次要按钮：bg-gray-600 hover:bg-gray-700 text-white
- 轮廓按钮：border border-gray-300 text-gray-700 hover:bg-gray-50
- 危险按钮：bg-error-600 hover:bg-error-700 text-white
- 成功按钮：bg-success-600 hover:bg-success-700 text-white

### 表单元素色彩

- 输入框边框（默认）：border-gray-300
- 输入框边框（聚焦）：border-primary-500 ring-2 ring-primary-500/50
- 输入框边框（错误）：border-error-500 ring-2 ring-error-500/50
- 输入框边框（成功）：border-success-500 ring-2 ring-success-500/50

### 文本色彩

- 主要标题：text-gray-900 dark:text-white
- 次要标题：text-gray-700 dark:text-gray-300
- 正文文本：text-gray-700 dark:text-gray-300
- 次要文本：text-gray-500 dark:text-gray-400
- 链接文本：text-primary-600 hover:text-primary-700 dark:text-primary-400 dark:hover:text-primary-300

### 背景色彩

- 页面背景：bg-white dark:bg-gray-900
- 卡片背景：bg-white dark:bg-gray-800
- 表格条纹：bg-gray-50 dark:bg-gray-800/50

### 边框和分隔线

- 主要边框：border-gray-200 dark:border-gray-700
- 次要边框：border-gray-100 dark:border-gray-800
- 分隔线：border-gray-200 dark:border-gray-700

## 颜色主题切换

为支持在亮色和暗色模式之间切换，请使用以下实用工具：

```tsx
// 切换主题的React Hook示例
import { useState, useEffect } from 'react';

export function useDarkMode() {
  const [darkMode, setDarkMode] = useState<boolean>(false);

  // 初始化时检查用户偏好
  useEffect(() => {
    const isDark = localStorage.getItem('darkMode') === 'true' || 
                  (window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches);
    setDarkMode(isDark);
    updateTheme(isDark);
  }, []);

  // 更新DOM和localStorage
  const updateTheme = (isDark: boolean) => {
    const root = window.document.documentElement;
    
    if (isDark) {
      root.classList.add('dark');
    } else {
      root.classList.remove('dark');
    }
    
    localStorage.setItem('darkMode', isDark.toString());
  };

  // 切换主题
  const toggleDarkMode = () => {
    setDarkMode(!darkMode);
    updateTheme(!darkMode);
  };

  return { darkMode, toggleDarkMode };
}
```

## 色彩可访问性检查

为确保色彩符合WCAG标准，应使用以下工具检查所有文本和背景组合：

1. WebAIM对比度检查器: https://webaim.org/resources/contrastchecker/
2. Stark Contrast Checker: https://www.getstark.co/
3. Colorable: https://colorable.jxnblk.com/

所有文本应至少符合WCAG 2.1 AA标准的对比度要求：
- 常规文本（小于18px）：对比度至少4.5:1
- 大文本（18px以上或14px粗体）：对比度至少3:1
- 界面组件和图形对象：对比度至少3:1 