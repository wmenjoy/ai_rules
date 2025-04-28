# 表单组件设计指南 (Part 1)

## 概述

表单组件是用户与应用程序交互的主要方式，它们允许用户输入、提交数据并接收反馈。本文档详细定义了基于React、TypeScript和Tailwind CSS的表单组件系统，包括输入字段、选择控件、布局和验证方式。

## 输入字段

### 文本输入 (Text Input)

最基本的输入字段，用于收集单行文本。

#### 设计规范

- **尺寸**: 高度为40px (默认)
- **内边距**: 水平12px，垂直8px
- **圆角**: 6px (rounded-md)
- **边框**: 1px solid，颜色为gray-300
- **焦点状态**: 边框为primary-500，有2px环形阴影
- **禁用状态**: 背景为gray-100，边框为gray-200
- **错误状态**: 边框为error-500，有2px环形阴影
- **成功状态**: 边框为success-500 (可选)

#### TypeScript接口

```typescript
interface InputProps extends React.InputHTMLAttributes<HTMLInputElement> {
  /** 输入框ID，必须提供，用于关联标签 */
  id: string;
  /** 标签文本 */
  label?: string;
  /** 帮助/提示文本 */
  helpText?: string;
  /** 错误信息 */
  error?: string;
  /** 是否处于错误状态 */
  isInvalid?: boolean;
  /** 是否处于有效状态 */
  isValid?: boolean;
  /** 输入框左侧附加内容 (图标、文本等) */
  leftAddon?: React.ReactNode;
  /** 输入框右侧附加内容 (图标、按钮等) */
  rightAddon?: React.ReactNode;
  /** 自定义类名 */
  className?: string;
}
```

#### Tailwind实现

```tsx
// components/Input.tsx
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
    ? 'border-error-300 text-error-900 placeholder-error-300 focus:ring-error-500 focus:border-error-500'
    : isValid
      ? 'border-success-300 text-gray-900 focus:ring-success-500 focus:border-success-500'
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
        <p className="mt-1 text-sm text-error-600">
          {error}
        </p>
      )}
    </div>
  );
};
```

#### 使用示例

```tsx
// 基本使用
<Input
  id="email"
  label="电子邮箱"
  type="email"
  placeholder="example@example.com"
  helpText="我们不会公开您的邮箱"
/>

// 带图标
<Input
  id="search"
  placeholder="搜索..."
  leftAddon={<SearchIcon className="h-5 w-5" />}
/>

// 错误状态
<Input
  id="username"
  label="用户名"
  isInvalid
  error="用户名已被占用"
/>

// 禁用状态
<Input
  id="username"
  label="用户名"
  disabled
  value="当前用户"
/>
```

### 密码输入 (Password Input)

用于输入敏感信息，带有显示/隐藏切换功能。

#### 额外功能

在基本输入字段的基础上，密码字段应添加以下功能：

- 显示/隐藏密码的切换按钮
- 可选的密码强度指示器

#### Tailwind实现

```tsx
// components/PasswordInput.tsx
import React, { useState } from 'react';
import { Input, InputProps } from './Input';
import { EyeIcon, EyeOffIcon } from '../icons';

interface PasswordInputProps extends Omit<InputProps, 'type'> {
  /** 是否显示密码强度指示器 */
  showStrengthIndicator?: boolean;
}

export const PasswordInput: React.FC<PasswordInputProps> = ({
  showStrengthIndicator = false,
  ...props
}) => {
  const [showPassword, setShowPassword] = useState(false);
  
  const togglePassword = () => setShowPassword(prev => !prev);
  
  return (
    <div className="space-y-1">
      <Input
        type={showPassword ? 'text' : 'password'}
        rightAddon={
          <button
            type="button"
            onClick={togglePassword}
            className="focus:outline-none"
            aria-label={showPassword ? '隐藏密码' : '显示密码'}
          >
            {showPassword ? (
              <EyeOffIcon className="h-5 w-5" />
            ) : (
              <EyeIcon className="h-5 w-5" />
            )}
          </button>
        }
        {...props}
      />
      
      {showStrengthIndicator && props.value && (
        <div className="mt-1">
          <div className="h-1 w-full bg-gray-200 rounded-full overflow-hidden">
            <div 
              className="h-full bg-success-500 rounded-full"
              style={{ width: `${calculatePasswordStrength(props.value?.toString() || '')}%` }}
            />
          </div>
          <p className="text-xs text-gray-500 mt-1">密码强度: {getPasswordStrengthLabel(props.value?.toString() || '')}</p>
        </div>
      )}
    </div>
  );
};

// 简化的密码强度计算函数
const calculatePasswordStrength = (password: string): number => {
  if (!password) return 0;
  
  let strength = 0;
  
  // 长度分数
  strength += Math.min(password.length * 2, 25);
  
  // 字符类型分数
  if (/[A-Z]/.test(password)) strength += 15;
  if (/[a-z]/.test(password)) strength += 10;
  if (/[0-9]/.test(password)) strength += 15;
  if (/[^A-Za-z0-9]/.test(password)) strength += 20;
  
  // 附加分数，基于长度和多样性
  const uniqueChars = new Set(password).size;
  strength += Math.min(uniqueChars * 1.5, 15);
  
  return Math.min(strength, 100);
};

const getPasswordStrengthLabel = (password: string): string => {
  const strength = calculatePasswordStrength(password);
  
  if (strength < 30) return '非常弱';
  if (strength < 50) return '弱';
  if (strength < 70) return '中等';
  if (strength < 90) return '强';
  return '非常强';
};
```

### 文本区域 (Textarea)

用于多行文本输入。

#### 设计规范

基本样式与文本输入相同，但有以下差异：

- **高度**: 默认为3行，可自定义
- **可调整大小**: 支持用户调整大小或固定大小

#### Tailwind实现

```tsx
// components/Textarea.tsx
import React from 'react';
import { classNames } from '../utils';

interface TextareaProps extends React.TextareaHTMLAttributes<HTMLTextAreaElement> {
  id: string;
  label?: string;
  helpText?: string;
  error?: string;
  isInvalid?: boolean;
  isValid?: boolean;
  resize?: 'none' | 'vertical' | 'horizontal' | 'both';
  className?: string;
}

export const Textarea: React.FC<TextareaProps> = ({
  id,
  label,
  helpText,
  error,
  isInvalid = false,
  isValid = false,
  resize = 'vertical',
  className,
  rows = 3,
  ...props
}) => {
  // 基础类
  const baseClasses = 'block w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-offset-0 focus:ring-opacity-50 transition-colors';
  
  // 状态特定类
  const stateClasses = isInvalid
    ? 'border-error-300 text-error-900 placeholder-error-300 focus:ring-error-500 focus:border-error-500'
    : isValid
      ? 'border-success-300 text-gray-900 focus:ring-success-500 focus:border-success-500'
      : 'border-gray-300 text-gray-900 placeholder-gray-400 focus:ring-primary-500 focus:border-primary-500';
  
  // 禁用状态
  const disabledClasses = props.disabled ? 'bg-gray-100 cursor-not-allowed' : 'bg-white';
  
  // 调整大小
  const resizeClass = {
    'none': 'resize-none',
    'vertical': 'resize-y',
    'horizontal': 'resize-x',
    'both': 'resize'
  }[resize];
  
  return (
    <div className="w-full">
      {label && (
        <label htmlFor={id} className="block text-sm font-medium text-gray-700 mb-1">
          {label}
        </label>
      )}
      
      <textarea
        id={id}
        rows={rows}
        className={classNames(
          baseClasses,
          stateClasses,
          disabledClasses,
          resizeClass,
          className
        )}
        aria-invalid={isInvalid ? 'true' : 'false'}
        aria-describedby={helpText ? `${id}-description` : undefined}
        {...props}
      />
      
      {helpText && !error && (
        <p id={`${id}-description`} className="mt-1 text-sm text-gray-500">
          {helpText}
        </p>
      )}
      
      {error && (
        <p className="mt-1 text-sm text-error-600">
          {error}
        </p>
      )}
    </div>
  );
};
```

### 数字输入 (Number Input)

专门用于数字输入，带有增减控制按钮。

#### 设计规范

在基本输入字段的基础上：

- 添加增减按钮
- 可设置最小值、最大值和步进值
- 可显示单位标签

#### Tailwind实现

```tsx
// components/NumberInput.tsx
import React from 'react';
import { Input, InputProps } from './Input';
import { PlusIcon, MinusIcon } from '../icons';

interface NumberInputProps extends Omit<InputProps, 'type'> {
  /** 最小值 */
  min?: number;
  /** 最大值 */
  max?: number;
  /** 步进值 */
  step?: number;
  /** 紧凑模式 (按钮更小) */
  isCompact?: boolean;
  /** 初始值 */
  value?: number | string;
  /** 值变化回调 */
  onChange?: (e: React.ChangeEvent<HTMLInputElement>) => void;
}

export const NumberInput: React.FC<NumberInputProps> = ({
  min,
  max,
  step = 1,
  isCompact = false,
  value,
  onChange,
  ...props
}) => {
  // 增加值
  const increment = () => {
    const currentValue = Number(value) || 0;
    const newValue = currentValue + step;
    
    if (max !== undefined && newValue > max) return;
    
    if (onChange) {
      const event = {
        target: { value: String(newValue) }
      } as React.ChangeEvent<HTMLInputElement>;
      
      onChange(event);
    }
  };
  
  // 减少值
  const decrement = () => {
    const currentValue = Number(value) || 0;
    const newValue = currentValue - step;
    
    if (min !== undefined && newValue < min) return;
    
    if (onChange) {
      const event = {
        target: { value: String(newValue) }
      } as React.ChangeEvent<HTMLInputElement>;
      
      onChange(event);
    }
  };
  
  // 按钮尺寸和样式
  const buttonClasses = isCompact
    ? 'p-1'
    : 'p-2';
  
  // 输入框右侧的控制按钮
  const controls = (
    <div className="flex flex-col border-l border-gray-300 divide-y divide-gray-300">
      <button
        type="button"
        className={`${buttonClasses} hover:bg-gray-100 focus:outline-none`}
        onClick={increment}
        aria-label="增加"
        tabIndex={-1}
      >
        <PlusIcon className="h-3 w-3" />
      </button>
      <button
        type="button"
        className={`${buttonClasses} hover:bg-gray-100 focus:outline-none`}
        onClick={decrement}
        aria-label="减少"
        tabIndex={-1}
      >
        <MinusIcon className="h-3 w-3" />
      </button>
    </div>
  );
  
  return (
    <Input
      type="number"
      min={min}
      max={max}
      step={step}
      value={value}
      onChange={onChange}
      rightAddon={controls}
      {...props}
    />
  );
};
```

### 搜索输入 (Search Input)

特定用于搜索功能的输入框。

#### 设计规范

- 左侧带有搜索图标
- 可选右侧清除按钮
- 圆角可以稍大一些（rounded-full选项）

#### Tailwind实现

```tsx
// components/SearchInput.tsx
import React, { useState } from 'react';
import { Input, InputProps } from './Input';
import { SearchIcon, XIcon } from '../icons';

interface SearchInputProps extends Omit<InputProps, 'type' | 'leftAddon' | 'rightAddon'> {
  /** 是否显示清除按钮 */
  showClearButton?: boolean;
  /** 是否使用完全圆角 */
  rounded?: boolean;
  /** 搜索提交回调 */
  onSearch?: (value: string) => void;
}

export const SearchInput: React.FC<SearchInputProps> = ({
  showClearButton = true,
  rounded = false,
  className = '',
  value,
  onChange,
  onSearch,
  ...props
}) => {
  const [inputValue, setInputValue] = useState(value || '');
  
  // 处理输入变化
  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setInputValue(e.target.value);
    if (onChange) onChange(e);
  };
  
  // 清除输入
  const clearInput = () => {
    setInputValue('');
    
    if (onChange) {
      const event = {
        target: { value: '' }
      } as React.ChangeEvent<HTMLInputElement>;
      
      onChange(event);
    }
  };
  
  // 提交搜索
  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === 'Enter' && onSearch) {
      onSearch(inputValue as string);
    }
    
    if (props.onKeyDown) {
      props.onKeyDown(e);
    }
  };
  
  // 搜索图标
  const searchAddon = <SearchIcon className="h-5 w-5 text-gray-400" />;
  
  // 清除按钮
  const clearAddon = showClearButton && inputValue ? (
    <button
      type="button"
      onClick={clearInput}
      className="focus:outline-none"
      aria-label="清除搜索"
    >
      <XIcon className="h-5 w-5 text-gray-400 hover:text-gray-500" />
    </button>
  ) : null;
  
  return (
    <Input
      type="search"
      value={inputValue}
      onChange={handleChange}
      onKeyDown={handleKeyDown}
      leftAddon={searchAddon}
      rightAddon={clearAddon}
      className={`${rounded ? 'rounded-full' : ''} ${className}`}
      {...props}
    />
  );
};
```

## 选择控件

### 复选框 (Checkbox)

用于多选场景。

#### 设计规范

- **尺寸**: 16px x 16px (默认)
- **边框**: 1px solid，颜色为gray-300
- **圆角**: 4px (rounded)
- **选中状态**: 背景色为primary-600，显示白色勾选图标
- **禁用状态**: 背景色为gray-100，边框为gray-200
- **间距**: 标签与框之间为8px

#### Tailwind实现

```tsx
// components/Checkbox.tsx
import React, { forwardRef } from 'react';
import { classNames } from '../utils';

export interface CheckboxProps extends Omit<React.InputHTMLAttributes<HTMLInputElement>, 'type'> {
  /** 复选框ID */
  id: string;
  /** 标签文本 */
  label?: React.ReactNode;
  /** 描述文本 */
  description?: React.ReactNode;
  /** 是否处于错误状态 */
  isInvalid?: boolean;
  /** 自定义类名 */
  className?: string;
}

export const Checkbox = forwardRef<HTMLInputElement, CheckboxProps>(
  ({ id, label, description, isInvalid = false, className = '', ...props }, ref) => {
    return (
      <div className="flex items-start">
        <div className="flex items-center h-5">
          <input
            id={id}
            type="checkbox"
            ref={ref}
            className={classNames(
              'h-4 w-4 rounded border-gray-300 text-primary-600 focus:ring-primary-500',
              isInvalid && 'border-error-300 focus:ring-error-500',
              props.disabled && 'opacity-50 cursor-not-allowed',
              className
            )}
            aria-invalid={isInvalid ? 'true' : 'false'}
            {...props}
          />
        </div>
        
        {(label || description) && (
          <div className="ml-2 text-sm">
            {label && (
              <label
                htmlFor={id}
                className={classNames(
                  'font-medium',
                  props.disabled ? 'text-gray-400' : 'text-gray-700',
                  isInvalid && 'text-error-700'
                )}
              >
                {label}
              </label>
            )}
            
            {description && (
              <p className="text-gray-500">{description}</p>
            )}
          </div>
        )}
      </div>
    );
  }
);

Checkbox.displayName = 'Checkbox';
```

### 单选按钮 (Radio)

用于单选场景。

#### 设计规范

与复选框类似，但有以下差异：

- 形状为圆形
- 选中状态显示内部圆点而非勾选图标

#### Tailwind实现

```tsx
// components/Radio.tsx
import React, { forwardRef } from 'react';
import { classNames } from '../utils';

export interface RadioProps extends Omit<React.InputHTMLAttributes<HTMLInputElement>, 'type'> {
  /** 单选按钮ID */
  id: string;
  /** 标签文本 */
  label?: React.ReactNode;
  /** 描述文本 */
  description?: React.ReactNode;
  /** 是否处于错误状态 */
  isInvalid?: boolean;
  /** 自定义类名 */
  className?: string;
}

export const Radio = forwardRef<HTMLInputElement, RadioProps>(
  ({ id, label, description, isInvalid = false, className = '', ...props }, ref) => {
    return (
      <div className="flex items-start">
        <div className="flex items-center h-5">
          <input
            id={id}
            type="radio"
            ref={ref}
            className={classNames(
              'h-4 w-4 border-gray-300 text-primary-600 focus:ring-primary-500',
              isInvalid && 'border-error-300 focus:ring-error-500',
              props.disabled && 'opacity-50 cursor-not-allowed',
              className
            )}
            aria-invalid={isInvalid ? 'true' : 'false'}
            {...props}
          />
        </div>
        
        {(label || description) && (
          <div className="ml-2 text-sm">
            {label && (
              <label
                htmlFor={id}
                className={classNames(
                  'font-medium',
                  props.disabled ? 'text-gray-400' : 'text-gray-700',
                  isInvalid && 'text-error-700'
                )}
              >
                {label}
              </label>
            )}
            
            {description && (
              <p className="text-gray-500">{description}</p>
            )}
          </div>
        )}
      </div>
    );
  }
);

Radio.displayName = 'Radio';
```

### 单选按钮组 (Radio Group)

方便管理一组相关的单选按钮。

#### Tailwind实现

```tsx
// components/RadioGroup.tsx
import React from 'react';
import { Radio, RadioProps } from './Radio';

interface RadioOption {
  id: string;
  value: string;
  label: React.ReactNode;
  description?: React.ReactNode;
  disabled?: boolean;
}

interface RadioGroupProps {
  /** 单选组名称 */
  name: string;
  /** 标签文本 */
  label?: string;
  /** 帮助文本 */
  helpText?: string;
  /** 错误信息 */
  errorMessage?: string;
  /** 单选项配置 */
  options: RadioOption[];
  /** 当前选中值 */
  value?: string;
  /** 值变化回调 */
  onChange?: (value: string) => void;
  /** 是否水平布局 */
  inline?: boolean;
  /** 自定义类名 */
  className?: string;
}

export const RadioGroup: React.FC<RadioGroupProps> = ({
  name,
  label,
  helpText,
  errorMessage,
  options,
  value,
  onChange,
  inline = false,
  className,
}) => {
  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (onChange) {
      onChange(e.target.value);
    }
  };
  
  const isInvalid = !!errorMessage;
  const groupId = `radio-group-${name}`;
  
  return (
    <div className={className}>
      {label && (
        <label className="block text-sm font-medium text-gray-700 mb-1">
          {label}
        </label>
      )}
      
      <div
        role="radiogroup"
        aria-labelledby={`${groupId}-label`}
        aria-describedby={
          helpText
            ? `${groupId}-description`
            : errorMessage
            ? `${groupId}-error`
            : undefined
        }
        className={`space-${inline ? 'x' : 'y'}-4 ${inline ? 'flex flex-wrap items-center' : ''}`}
      >
        {options.map((option) => (
          <div
            key={option.id}
            className={inline ? 'mr-4 mb-2' : 'mb-2'}
          >
            <Radio
              id={option.id}
              name={name}
              value={option.value}
              checked={value === option.value}
              onChange={handleChange}
              label={option.label}
              description={option.description}
              disabled={option.disabled}
              isInvalid={isInvalid}
            />
          </div>
        ))}
      </div>
      
      {helpText && !errorMessage && (
        <p id={`${groupId}-description`} className="mt-1 text-sm text-gray-500">
          {helpText}
        </p>
      )}
      
      {errorMessage && (
        <p id={`${groupId}-error`} className="mt-1 text-sm text-error-600">
          {errorMessage}
        </p>
      )}
    </div>
  );
};
```

### 切换开关 (Toggle)

适用于开/关、启用/禁用等二元选择场景。

#### 设计规范

- **尺寸**: 默认宽度为36px，高度为20px
- **滑块尺寸**: 直径为16px
- **关闭状态**: 背景为gray-200，滑块靠左
- **打开状态**: 背景为primary-600，滑块靠右
- **禁用状态**: 降低透明度，鼠标为not-allowed
- **过渡动画**: 滑块移动有平滑过渡效果

#### Tailwind实现

```tsx
// components/Toggle.tsx
import React from 'react';
import { classNames } from '../utils';

interface ToggleProps {
  /** 切换开关ID */
  id: string;
  /** 标签文本 */
  label?: React.ReactNode;
  /** 描述文本 */
  description?: React.ReactNode;
  /** 是否打开 */
  checked?: boolean;
  /** 是否禁用 */
  disabled?: boolean;
  /** 尺寸 */
  size?: 'sm' | 'md' | 'lg';
  /** 自定义类名 */
  className?: string;
  /** 值变化回调 */
  onChange?: (checked: boolean) => void;
}

export const Toggle: React.FC<ToggleProps> = ({
  id,
  label,
  description,
  checked = false,
  disabled = false,
  size = 'md',
  className,
  onChange,
}) => {
  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (onChange) {
      onChange(e.target.checked);
    }
  };
  
  // 尺寸类
  const sizeClasses = {
    sm: {
      toggle: 'w-8 h-4',
      dot: 'h-3 w-3',
      translate: 'translate-x-4',
    },
    md: {
      toggle: 'w-11 h-6',
      dot: 'h-5 w-5',
      translate: 'translate-x-5',
    },
    lg: {
      toggle: 'w-14 h-7',
      dot: 'h-6 w-6',
      translate: 'translate-x-7',
    },
  };
  
  const { toggle, dot, translate } = sizeClasses[size];
  
  return (
    <div className={classNames('flex items-start', className)}>
      <div className="flex items-center h-5">
        <input
          id={id}
          type="checkbox"
          className="sr-only"
          checked={checked}
          disabled={disabled}
          onChange={handleChange}
          aria-labelledby={`${id}-label`}
          aria-describedby={description ? `${id}-description` : undefined}
        />
        
        <button
          type="button"
          onClick={() => !disabled && onChange?.(!checked)}
          className={classNames(
            'relative inline-flex flex-shrink-0 border-2 border-transparent rounded-full cursor-pointer transition-colors ease-in-out duration-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500',
            checked ? 'bg-primary-600' : 'bg-gray-200',
            disabled && 'opacity-50 cursor-not-allowed',
            toggle
          )}
          aria-hidden="true"
        >
          <span
            className={classNames(
              'pointer-events-none inline-block rounded-full bg-white shadow transform ring-0 transition ease-in-out duration-200',
              checked ? translate : 'translate-x-0',
              dot
            )}
          />
        </button>
      </div>
      
      {(label || description) && (
        <div className="ml-3 text-sm">
          {label && (
            <label
              id={`${id}-label`}
              htmlFor={id}
              className={classNames(
                'font-medium text-gray-700',
                disabled && 'text-gray-400'
              )}
            >
              {label}
            </label>
          )}
          
          {description && (
            <p id={`${id}-description`} className="text-gray-500">
              {description}
            </p>
          )}
        </div>
      )}
    </div>
  );
};
```

## 下一部分

这是表单组件设计指南的第一部分，涵盖了基本输入字段和控件。在第二部分中，我们将继续探讨选择框、自动完成、表单布局、验证和交互等更复杂的组件和模式。 