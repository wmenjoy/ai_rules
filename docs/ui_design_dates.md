# 日期组件设计指南

## 1. 概述

日期组件是用户界面中的关键元素，允许用户以各种方式查看、选择和输入日期与时间信息。本指南详细说明了在 React + TypeScript + Tailwind CSS 项目中如何设计和实现日期相关组件，确保一致性、可用性和可访问性。

## 2. 设计原则

日期组件遵循以下设计原则：

- **直观性**：用户应能轻松理解如何选择和修改日期
- **灵活性**：支持多种日期格式和选择模式（单日期、日期范围等）
- **精确性**：提供足够精度，同时避免输入错误
- **本地化**：支持不同语言和日期格式的国际化
- **可访问性**：确保所有用户，包括使用辅助技术的用户，都能使用日期组件

## 3. 日期选择器组件 (Date Picker)

日期选择器允许用户从日历界面选择单个日期。

### 3.1 日期选择器变体

| 变体 | 特点 | 使用场景 |
|------|------|---------|
| 下拉日期选择器 | 点击后显示日历弹出层 | 表单中的日期字段 |
| 内联日期选择器 | 直接显示在页面上的日历 | 需要频繁日期选择的界面 |
| 迷你日期选择器 | 紧凑版本的选择器 | 空间有限的界面 |
| 带时间的日期选择器 | 允许选择日期和时间 | 需要精确到时间的场景 |

### 3.2 日期选择器 TypeScript 接口

```typescript
interface DatePickerProps {
  // 基本属性
  value: Date | null;
  onChange: (date: Date | null) => void;
  
  // 限制属性
  minDate?: Date;
  maxDate?: Date;
  disabledDates?: Date[] | ((date: Date) => boolean);
  
  // 显示属性
  format?: string; // 显示格式，如 'yyyy-MM-dd'
  placeholder?: string;
  label?: string;
  variant?: 'dropdown' | 'inline' | 'mini';
  size?: 'sm' | 'md' | 'lg';
  
  // 国际化
  locale?: string;
  firstDayOfWeek?: 0 | 1 | 2 | 3 | 4 | 5 | 6; // 0 = 周日, 1 = 周一, 等
  
  // 状态
  disabled?: boolean;
  readOnly?: boolean;
  error?: string;
  
  // 通用属性
  id?: string;
  name?: string;
  className?: string;
  
  // 高级功能
  showWeekNumbers?: boolean;
  showToday?: boolean;
  clearable?: boolean;
  renderDay?: (date: Date) => React.ReactNode;
}
```

### 3.3 日期选择器实现示例

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
import React, { useState } from 'react';
import { format, isValid, isAfter, isBefore, startOfDay } from 'date-fns';
import { zhCN } from 'date-fns/locale';
import { classNames } from '../utils';

export const DatePicker: React.FC<DatePickerProps> = ({
  value,
  onChange,
  minDate,
  maxDate,
  format: dateFormat = 'yyyy-MM-dd',
  placeholder = '请选择日期',
  label,
  variant = 'dropdown',
  size = 'md',
  locale = 'zhCN',
  disabled = false,
  error,
  className,
  ...props
}) => {
  const [isOpen, setIsOpen] = useState(false);
  const [currentMonth, setCurrentMonth] = useState(value || new Date());
  const localeObj = locale === 'zhCN' ? zhCN : undefined;
  
  const handleDateSelect = (date: Date) => {
    // 检查日期是否在允许范围内
    if (minDate && isBefore(date, startOfDay(minDate))) return;
    if (maxDate && isAfter(date, startOfDay(maxDate))) return;
    
    onChange(date);
    if (variant === 'dropdown') setIsOpen(false);
  };
  
  const sizeClasses = {
    sm: 'text-sm',
    md: 'text-base',
    lg: 'text-lg',
  };
  
  // 日期选择器的实现...
  
  return (
    <div className={classNames('relative', className)}>
      {label && (
        <label className="block text-sm font-medium text-gray-700 mb-1">
          {label}
        </label>
      )}
      
      <div className="relative">
        <input
          type="text"
          readOnly
          placeholder={placeholder}
          value={value && isValid(value) ? format(value, dateFormat, { locale: localeObj }) : ''}
          onClick={() => !disabled && setIsOpen(true)}
          className={classNames(
            'w-full px-3 py-2 border rounded-md',
            'focus:ring-2 focus:ring-primary-500 focus:border-primary-500',
            disabled ? 'bg-gray-100 cursor-not-allowed' : 'cursor-pointer',
            error ? 'border-error-500' : 'border-gray-300',
            sizeClasses[size]
          )}
          disabled={disabled}
          {...props}
        />
        
        <div className="absolute inset-y-0 right-0 flex items-center pr-3 pointer-events-none">
          <svg
            className="h-5 w-5 text-gray-400"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 20 20"
            fill="currentColor"
          >
            <path
              fillRule="evenodd"
              d="M6 2a1 1 0 00-1 1v1H4a2 2 0 00-2 2v10a2 2 0 002 2h12a2 2 0 002-2V6a2 2 0 00-2-2h-1V3a1 1 0 10-2 0v1H7V3a1 1 0 00-1-1zm0 5a1 1 0 000 2h8a1 1 0 100-2H6z"
            />
          </svg>
        </div>
      </div>
      
      {error && <p className="mt-1 text-sm text-error-500">{error}</p>}
      
      {/* 下拉日历的实现 */}
      {isOpen && variant === 'dropdown' && (
        <div className="absolute z-10 mt-1 bg-white rounded-md shadow-lg p-4 border border-gray-200">
          {/* 日历组件实现 */}
        </div>
      )}
      
      {/* 内联日历的实现 */}
      {variant === 'inline' && (
        <div className="mt-2 bg-white rounded-md border border-gray-200 p-4">
          {/* 日历组件实现 */}
        </div>
      )}
    </div>
  );
};
// [AI-BLOCK-END]
```

### 3.4 可访问性考虑

- 使用适当的 ARIA 属性（`aria-label`, `aria-expanded`, `aria-controls`）
- 确保键盘可导航（Tab、方向键、Enter、Escape）
- 支持屏幕阅读器宣告选定日期
- 确保颜色对比度符合 WCAG 标准
- 提供清晰的错误消息和验证

## 4. 日期范围选择器 (Date Range Picker)

日期范围选择器允许用户选择日期的开始和结束，定义一个时间段。

### 4.1 日期范围选择器变体

| 变体 | 特点 | 使用场景 |
|------|------|---------|
| 单日历范围选择器 | 在同一个日历上选择范围 | 短期范围选择 |
| 双日历范围选择器 | 两个并排的日历选择起止日期 | 更灵活的范围选择 |
| 预设范围选择器 | 包含常用预设时间段的选择器 | 报表和分析筛选 |

### 4.2 日期范围选择器 TypeScript 接口

```typescript
interface DateRangePickerProps {
  // 基本属性
  startDate: Date | null;
  endDate: Date | null;
  onRangeChange: (range: [Date | null, Date | null]) => void;
  
  // 限制属性
  minDate?: Date;
  maxDate?: Date;
  minDuration?: number; // 最短持续时间（天）
  maxDuration?: number; // 最长持续时间（天）
  disabledDates?: Date[] | ((date: Date) => boolean);
  
  // 显示属性
  format?: string;
  startPlaceholder?: string;
  endPlaceholder?: string;
  label?: string;
  variant?: 'single' | 'double' | 'presets';
  size?: 'sm' | 'md' | 'lg';
  
  // 预设配置
  presets?: {
    label: string;
    range: [Date, Date];
  }[];
  
  // 国际化
  locale?: string;
  firstDayOfWeek?: 0 | 1 | 2 | 3 | 4 | 5 | 6;
  
  // 状态
  disabled?: boolean;
  readOnly?: boolean;
  error?: string;
  
  // 通用属性
  className?: string;
}
```

### 4.3 日期范围选择器示例

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
import React, { useState } from 'react';
import { format, isValid, addMonths } from 'date-fns';
import { zhCN } from 'date-fns/locale';
import { classNames } from '../utils';

export const DateRangePicker: React.FC<DateRangePickerProps> = ({
  startDate,
  endDate,
  onRangeChange,
  format: dateFormat = 'yyyy-MM-dd',
  startPlaceholder = '开始日期',
  endPlaceholder = '结束日期',
  label,
  variant = 'double',
  size = 'md',
  presets = [],
  locale = 'zhCN',
  disabled = false,
  error,
  className,
}) => {
  const [isOpen, setIsOpen] = useState(false);
  const [hoverDate, setHoverDate] = useState<Date | null>(null);
  const [leftMonth, setLeftMonth] = useState(startDate || new Date());
  const localeObj = locale === 'zhCN' ? zhCN : undefined;
  
  const handleSelect = (date: Date) => {
    if (!startDate || (startDate && endDate)) {
      // 开始新的选择
      onRangeChange([date, null]);
    } else {
      // 完成范围选择
      if (date < startDate) {
        onRangeChange([date, startDate]);
      } else {
        onRangeChange([startDate, date]);
      }
      
      if (variant !== 'single') {
        setIsOpen(false);
      }
    }
  };
  
  const handlePresetSelect = (start: Date, end: Date) => {
    onRangeChange([start, end]);
    setIsOpen(false);
  };
  
  const sizeClasses = {
    sm: 'text-sm py-1',
    md: 'text-base py-2',
    lg: 'text-lg py-3',
  };
  
  return (
    <div className={classNames('relative', className)}>
      {label && (
        <label className="block text-sm font-medium text-gray-700 mb-1">
          {label}
        </label>
      )}
      
      <div className="flex items-center space-x-2">
        {/* 开始日期输入 */}
        <div className="relative flex-1">
          <input
            type="text"
            readOnly
            placeholder={startPlaceholder}
            value={startDate && isValid(startDate) ? format(startDate, dateFormat, { locale: localeObj }) : ''}
            onClick={() => !disabled && setIsOpen(true)}
            className={classNames(
              'w-full px-3 border rounded-md',
              'focus:ring-2 focus:ring-primary-500 focus:border-primary-500',
              disabled ? 'bg-gray-100 cursor-not-allowed' : 'cursor-pointer',
              error ? 'border-error-500' : 'border-gray-300',
              sizeClasses[size]
            )}
            disabled={disabled}
          />
          <div className="absolute inset-y-0 right-0 flex items-center pr-3 pointer-events-none">
            <svg
              className="h-5 w-5 text-gray-400"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 20 20"
              fill="currentColor"
            >
              <path fillRule="evenodd" d="M6 2a1 1 0 00-1 1v1H4a2 2 0 00-2 2v10a2 2 0 002 2h12a2 2 0 002-2V6a2 2 0 00-2-2h-1V3a1 1 0 10-2 0v1H7V3a1 1 0 00-1-1zm0 5a1 1 0 000 2h8a1 1 0 100-2H6z" />
            </svg>
          </div>
        </div>
        
        <span className="text-gray-500">至</span>
        
        {/* 结束日期输入 */}
        <div className="relative flex-1">
          <input
            type="text"
            readOnly
            placeholder={endPlaceholder}
            value={endDate && isValid(endDate) ? format(endDate, dateFormat, { locale: localeObj }) : ''}
            onClick={() => !disabled && setIsOpen(true)}
            className={classNames(
              'w-full px-3 border rounded-md',
              'focus:ring-2 focus:ring-primary-500 focus:border-primary-500',
              disabled ? 'bg-gray-100 cursor-not-allowed' : 'cursor-pointer',
              error ? 'border-error-500' : 'border-gray-300',
              sizeClasses[size]
            )}
            disabled={disabled}
          />
          <div className="absolute inset-y-0 right-0 flex items-center pr-3 pointer-events-none">
            <svg
              className="h-5 w-5 text-gray-400"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 20 20"
              fill="currentColor"
            >
              <path fillRule="evenodd" d="M6 2a1 1 0 00-1 1v1H4a2 2 0 00-2 2v10a2 2 0 002 2h12a2 2 0 002-2V6a2 2 0 00-2-2h-1V3a1 1 0 10-2 0v1H7V3a1 1 0 00-1-1zm0 5a1 1 0 000 2h8a1 1 0 100-2H6z" />
            </svg>
          </div>
        </div>
      </div>
      
      {error && <p className="mt-1 text-sm text-error-500">{error}</p>}
      
      {/* 下拉日历部分 */}
      {isOpen && (
        <div className="absolute z-10 mt-1 bg-white rounded-md shadow-lg border border-gray-200 p-4">
          <div className="flex">
            {/* 预设选项 */}
            {variant === 'presets' && presets.length > 0 && (
              <div className="border-r border-gray-200 pr-4 mr-4">
                <p className="font-medium text-sm mb-2">常用日期范围</p>
                <div className="space-y-1">
                  {presets.map((preset, index) => (
                    <button
                      key={index}
                      className="block w-full text-left px-3 py-1 text-sm rounded hover:bg-gray-100"
                      onClick={() => handlePresetSelect(preset.range[0], preset.range[1])}
                    >
                      {preset.label}
                    </button>
                  ))}
                </div>
              </div>
            )}
            
            {/* 日历部分 */}
            <div>
              <div className="flex space-x-4">
                {/* 左侧日历 */}
                <div>
                  {/* 日历导航和日期网格 */}
                </div>
                
                {/* 右侧日历（双日历模式） */}
                {variant === 'double' && (
                  <div>
                    {/* 第二个日历导航和日期网格 */}
                  </div>
                )}
              </div>
              
              <div className="mt-4 flex justify-end">
                <button
                  type="button"
                  className="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-md hover:bg-gray-50 mr-2"
                  onClick={() => setIsOpen(false)}
                >
                  取消
                </button>
                <button
                  type="button"
                  className="px-4 py-2 text-sm font-medium text-white bg-primary-600 border border-transparent rounded-md hover:bg-primary-700"
                  onClick={() => setIsOpen(false)}
                  disabled={!startDate || !endDate}
                >
                  确定
                </button>
              </div>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};
// [AI-BLOCK-END]
```

## 5. 时间选择器 (Time Picker)

时间选择器允许用户选择特定的时间点。

### 5.1 时间选择器变体

| 变体 | 特点 | 使用场景 |
|------|------|---------|
| 24小时制 | 使用24小时制显示时间 | 默认选择，适合大多数场景 |
| 12小时制 | 使用AM/PM标记的12小时制 | 需要12小时制显示的场景 |
| 时分选择器 | 只选择小时和分钟 | 不需要精确秒数的场景 |
| 时分秒选择器 | 选择时、分、秒 | 需要高精度时间的场景 |

### 5.2 时间选择器 TypeScript 接口

```typescript
interface TimePickerProps {
  // 基本属性
  value: Date | null;
  onChange: (time: Date | null) => void;
  
  // 格式选项
  format?: string; // 'HH:mm', 'HH:mm:ss', 'hh:mm a'
  use24Hours?: boolean;
  showSeconds?: boolean;
  
  // 步进值
  hourStep?: number; // 默认1
  minuteStep?: number; // 默认1
  secondStep?: number; // 默认1
  
  // 显示属性
  placeholder?: string;
  label?: string;
  size?: 'sm' | 'md' | 'lg';
  
  // 限制
  disabledHours?: () => number[];
  disabledMinutes?: (hour: number) => number[];
  disabledSeconds?: (hour: number, minute: number) => number[];
  minTime?: Date;
  maxTime?: Date;
  
  // 状态
  disabled?: boolean;
  readOnly?: boolean;
  error?: string;
  
  // 国际化
  locale?: string;
  
  // 通用属性
  className?: string;
}
```

### 5.3 移动设备考虑

移动设备上的日期选择有其特殊性，应遵循以下指南：

1. **利用原生选择器**
   - 在移动设备上，考虑使用原生日期选择器（`<input type="date">`, `<input type="time">`）
   - 原生选择器提供更好的用户体验和性能

2. **触摸友好设计**
   - 较大的点击区域（推荐至少48×48像素）
   - 适当的间距避免误触
   - 简化的界面，减少干扰元素

3. **垂直滑动选择器**
   - 对于自定义实现，考虑使用垂直滚动选择器而非传统日历
   - 分离年、月、日选择，提供更好的触摸体验

4. **响应式调整**
   - 单日历而非双日历布局
   - 全屏模式的日期选择器
   - 优化的键盘输入支持

## 6. 日期本地化与国际化

日期组件的国际化是至关重要的，应考虑以下几点：

### 6.1 日期格式

针对不同地区使用适当的日期格式：
- 美国: MM/DD/YYYY
- 欧洲大部分地区: DD/MM/YYYY
- 中国: YYYY-MM-DD
- 支持自定义格式以适应不同需求

### 6.2 语言本地化

- 使用 date-fns 或 Intl API 进行日期本地化
- 月份名称、星期几的正确翻译
- 支持多语言切换

### 6.3 一周起始日

不同文化中一周的第一天不同：
- 北美: 周日
- 欧洲大部分地区: 周一
- 中东地区: 周六
- 允许自定义设置

## 7. 日期组件的错误状态与验证

### 7.1 错误显示

- 清晰的错误消息，直接显示在组件下方
- 边框颜色变更（通常为红色）突出错误状态
- 错误图标和辅助文本提供额外上下文

### 7.2 常见验证规则

- 必填验证
- 日期范围验证（最小/最大日期）
- 日期格式验证
- 日期关系验证（结束日期必须晚于开始日期）

## 8. 最佳实践

### 8.1 日期选择器使用建议

- 为日期选择器提供明确的标签
- 在可能的情况下提供示例或格式提示
- 允许用户通过键盘输入日期，而不只是选择
- 使用适当的占位符文本示例预期格式
- 在有合理默认值时预填充

### 8.2 性能考虑

- 延迟加载日历组件，直到需要时
- 使用缓存减少日期计算
- 避免不必要的重渲染
- 考虑使用基于原生日期输入的回退方案

### 8.3 可访问性核对清单

- 支持键盘导航（Tab, 方向键, Enter, Escape）
- 适当的 ARIA 标签和角色
- 高对比度模式支持
- 屏幕阅读器支持，包括状态变化通知
- 支持放大和文本大小调整 