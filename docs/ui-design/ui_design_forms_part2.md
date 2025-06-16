# 表单组件设计指南 - 第2部分

## 3. 选择类组件

### 3.1 下拉选择器
下拉选择器组件允许用户从列表中选择单个选项。

```typescript
interface SelectProps {
  id: string;
  name: string;
  options: Array<{value: string; label: string}>;
  value?: string;
  onChange: (value: string) => void;
  label?: string;
  placeholder?: string;
  disabled?: boolean;
  required?: boolean;
  error?: string;
  className?: string;
}
```

**设计规范：**
- 高度：40px（与文本输入框一致）
- 边框：1px solid #E2E8F0（灰色-300）
- 圆角：4px
- 背景：白色
- 内边距：8px 12px
- 下拉图标：向下箭头（16x16px）
- 状态：
  - 默认：浅灰色边框
  - 悬停：边框颜色加深至灰色-400
  - 聚焦：蓝色边框（primary-500）带2px轮廓
  - 禁用：浅灰色背景，无悬停效果
  - 错误：红色边框（error-500）

**Tailwind 实现：**
```jsx
<div className="relative">
  {label && (
    <label htmlFor={id} className="block text-sm font-medium text-gray-700 mb-1">
      {label}{required && <span className="text-error-500 ml-1">*</span>}
    </label>
  )}
  <select
    id={id}
    name={name}
    value={value}
    onChange={(e) => onChange(e.target.value)}
    disabled={disabled}
    className={`
      block w-full h-10 pl-3 pr-10 py-2
      bg-white border border-gray-300 rounded-md
      text-sm text-gray-900
      focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500
      disabled:bg-gray-100 disabled:text-gray-500 disabled:cursor-not-allowed
      ${error ? 'border-error-500 focus:border-error-500 focus:ring-error-500' : ''}
      ${className || ''}
    `}
  >
    {placeholder && (
      <option value="" disabled>
        {placeholder}
      </option>
    )}
    {options.map((option) => (
      <option key={option.value} value={option.value}>
        {option.label}
      </option>
    ))}
  </select>
  <div className="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700">
    <svg className="h-4 w-4 fill-current" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20">
      <path d="M9.293 12.95l.707.707L15.657 8l-1.414-1.414L10 10.828 5.757 6.586 4.343 8z" />
    </svg>
  </div>
  {error && (
    <p className="mt-1 text-sm text-error-500">{error}</p>
  )}
</div>
```

### 3.2 多选选择器
多选选择器组件允许用户从列表中选择多个选项。

```typescript
interface MultiSelectProps {
  id: string;
  name: string;
  options: Array<{value: string; label: string}>;
  value: string[];
  onChange: (value: string[]) => void;
  label?: string;
  placeholder?: string;
  disabled?: boolean;
  required?: boolean;
  error?: string;
  className?: string;
  maxSelections?: number;
}
```

**设计规范：**
- 基础高度：40px（随着选择项的增加而扩展）
- 选中的项目显示为内部的标签/芯片
- 每个标签带有"X"图标可删除所选项
- 视觉样式与下拉选择器匹配

**Tailwind 实现：**
```jsx
// 这是一个简化实现 - 在生产环境中，考虑使用如react-select等库
<div className="relative">
  {label && (
    <label htmlFor={id} className="block text-sm font-medium text-gray-700 mb-1">
      {label}{required && <span className="text-error-500 ml-1">*</span>}
    </label>
  )}
  <div 
    className={`
      min-h-10 px-3 py-2 flex flex-wrap gap-1 items-center
      border border-gray-300 rounded-md
      focus-within:ring-2 focus-within:ring-primary-500 focus-within:border-primary-500
      ${disabled ? 'bg-gray-100' : 'bg-white'}
      ${error ? 'border-error-500' : ''}
    `}
  >
    {value.length > 0 ? (
      value.map((item) => (
        <span key={item} className="bg-primary-100 text-primary-800 text-xs rounded px-2 py-1 flex items-center">
          {options.find(option => option.value === item)?.label}
          <button 
            type="button"
            onClick={() => onChange(value.filter(val => val !== item))}
            className="ml-1 text-gray-500 hover:text-gray-700"
            disabled={disabled}
          >
            <svg className="h-3 w-3 fill-current" viewBox="0 0 20 20">
              <path d="M10 8.586l3.293-3.293a1 1 0 011.414 1.414L11.414 10l3.293 3.293a1 1 0 01-1.414 1.414L10 11.414l-3.293 3.293a1 1 0 01-1.414-1.414L8.586 10 5.293 6.707a1 1 0 011.414-1.414L10 8.586z" />
            </svg>
          </button>
        </span>
      ))
    ) : (
      <span className="text-gray-500">{placeholder}</span>
    )}
  </div>
  {error && (
    <p className="mt-1 text-sm text-error-500">{error}</p>
  )}
</div>
```

### 3.3 自动完成
自动完成组件结合了文本输入与根据用户输入过滤的建议列表。

```typescript
interface AutocompleteProps {
  id: string;
  name: string;
  options: Array<{value: string; label: string}>;
  value: string;
  onChange: (value: string) => void;
  onSelect: (option: {value: string; label: string}) => void;
  label?: string;
  placeholder?: string;
  disabled?: boolean;
  required?: boolean;
  error?: string;
  className?: string;
  maxSuggestions?: number;
}
```

**设计规范：**
- 基础样式与文本输入框一致
- 输入时在输入框下方显示下拉面板
- 悬停时突出显示建议项
- 可使用方向键浏览建议
- 回车键选择高亮的建议项
- Esc键关闭建议面板

**Tailwind 实现：**
```jsx
// 简化实现 - 在生产环境中，考虑使用库或无头UI组件
<div className="relative">
  {label && (
    <label htmlFor={id} className="block text-sm font-medium text-gray-700 mb-1">
      {label}{required && <span className="text-error-500 ml-1">*</span>}
    </label>
  )}
  <input
    id={id}
    name={name}
    type="text"
    value={value}
    onChange={(e) => onChange(e.target.value)}
    disabled={disabled}
    placeholder={placeholder}
    className={`
      block w-full h-10 px-3 py-2
      bg-white border border-gray-300 rounded-md
      text-sm text-gray-900
      focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500
      disabled:bg-gray-100 disabled:text-gray-500 disabled:cursor-not-allowed
      ${error ? 'border-error-500 focus:border-error-500 focus:ring-error-500' : ''}
      ${className || ''}
    `}
  />
  
  {/* 建议下拉菜单 - 根据状态有条件地显示 */}
  <div className="absolute z-10 w-full mt-1 bg-white shadow-lg max-h-60 rounded-md py-1 text-sm overflow-auto focus:outline-none">
    {filteredOptions.map((option, index) => (
      <div
        key={option.value}
        className={`
          cursor-default select-none relative py-2 pl-3 pr-9
          ${highlightedIndex === index ? 'bg-primary-100 text-primary-900' : 'text-gray-900'}
        `}
        onClick={() => onSelect(option)}
      >
        {option.label}
      </div>
    ))}
  </div>
  
  {error && (
    <p className="mt-1 text-sm text-error-500">{error}</p>
  )}
</div>
```

## 4. 表单布局模式

### 4.1 基础表单布局
标准表单布局遵循以下准则：

**设计规范：**
- 标签位于输入字段上方
- 字段间距：表单组之间16px（1rem）
- 必填字段使用星号（*）标记
- 错误信息直接显示在字段下方
- 表单操作（提交/取消按钮）根据上下文右对齐或分散对齐

**Tailwind 实现：**
```jsx
<form className="space-y-4">
  <div>
    <label htmlFor="name" className="block text-sm font-medium text-gray-700 mb-1">
      姓名<span className="text-error-500 ml-1">*</span>
    </label>
    <input 
      id="name" 
      name="name" 
      type="text" 
      required 
      className="block w-full h-10 px-3 py-2 border border-gray-300 rounded-md text-sm" 
    />
  </div>
  
  <div>
    <label htmlFor="email" className="block text-sm font-medium text-gray-700 mb-1">
      邮箱<span className="text-error-500 ml-1">*</span>
    </label>
    <input 
      id="email" 
      name="email" 
      type="email" 
      required 
      className="block w-full h-10 px-3 py-2 border border-gray-300 rounded-md text-sm" 
    />
  </div>
  
  <div className="flex justify-end space-x-3 pt-4">
    <button 
      type="button" 
      className="px-4 py-2 text-sm text-gray-700 hover:text-gray-500"
    >
      取消
    </button>
    <button 
      type="submit" 
      className="px-4 py-2 text-sm bg-primary-600 text-white rounded-md hover:bg-primary-700"
    >
      提交
    </button>
  </div>
</form>
```

### 4.2 内联表单布局
用于需要在单行显示的紧凑型表单：

**设计规范：**
- 字段和标签水平排列
- 响应式：在移动设备上转为堆叠布局
- 元素之间的紧凑间距（8px）

**Tailwind 实现：**
```jsx
<form className="sm:flex sm:items-center sm:space-x-4">
  <div className="w-full sm:w-auto">
    <label htmlFor="inline-email" className="sr-only">邮箱</label>
    <input 
      id="inline-email" 
      name="email" 
      type="email" 
      placeholder="输入您的邮箱" 
      className="block w-full h-10 px-3 py-2 border border-gray-300 rounded-md text-sm" 
    />
  </div>
  
  <button 
    type="submit" 
    className="mt-3 sm:mt-0 w-full sm:w-auto px-4 py-2 text-sm bg-primary-600 text-white rounded-md hover:bg-primary-700"
  >
    订阅
  </button>
</form>
```

### 4.3 多列布局
适用于复杂的多列表单：

**设计规范：**
- 桌面端有两列或更多
- 移动端缩减为单列
- 默认等宽列布局
- 可选自定义列宽

**Tailwind 实现：**
```jsx
<form className="space-y-6">
  <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
    <div>
      <label htmlFor="first-name" className="block text-sm font-medium text-gray-700 mb-1">
        名字<span className="text-error-500 ml-1">*</span>
      </label>
      <input 
        id="first-name" 
        name="first-name" 
        type="text" 
        required 
        className="block w-full h-10 px-3 py-2 border border-gray-300 rounded-md text-sm" 
      />
    </div>
    
    <div>
      <label htmlFor="last-name" className="block text-sm font-medium text-gray-700 mb-1">
        姓氏<span className="text-error-500 ml-1">*</span>
      </label>
      <input 
        id="last-name" 
        name="last-name" 
        type="text" 
        required 
        className="block w-full h-10 px-3 py-2 border border-gray-300 rounded-md text-sm" 
      />
    </div>
    
    <div className="md:col-span-2">
      <label htmlFor="address" className="block text-sm font-medium text-gray-700 mb-1">
        地址
      </label>
      <input 
        id="address" 
        name="address" 
        type="text" 
        className="block w-full h-10 px-3 py-2 border border-gray-300 rounded-md text-sm" 
      />
    </div>
  </div>
  
  <div className="flex justify-end space-x-3">
    <button 
      type="button" 
      className="px-4 py-2 text-sm text-gray-700 hover:text-gray-500"
    >
      取消
    </button>
    <button 
      type="submit" 
      className="px-4 py-2 text-sm bg-primary-600 text-white rounded-md hover:bg-primary-700"
    >
      提交
    </button>
  </div>
</form>
```

## 5. 表单验证和错误状态

### 5.1 验证类型
表单支持以下验证模式：

**客户端验证：**
- 用户输入时的实时验证
- 字段失焦时验证（当用户离开字段时）
- 表单提交时验证

**验证指示器：**
- 颜色编码：错误使用红色（error-500）
- 图标指示：错误使用感叹号图标
- 成功状态：可选的绿色对勾表示有效字段

### 5.2 错误消息

**设计规范：**
- 位置：直接位于输入字段下方
- 文本颜色：错误红色（error-500）
- 字体大小：小（text-sm）
- 图标：可选的感叹号图标

**Tailwind 实现：**
```jsx
<div>
  <label htmlFor="email" className="block text-sm font-medium text-gray-700 mb-1">
    邮箱<span className="text-error-500 ml-1">*</span>
  </label>
  <input 
    id="email" 
    name="email" 
    type="email" 
    aria-invalid="true"
    aria-describedby="email-error"
    className="block w-full h-10 px-3 py-2 border border-error-500 rounded-md text-sm focus:ring-error-500 focus:border-error-500" 
  />
  <p className="mt-1 text-sm text-error-500 flex items-center" id="email-error">
    <svg className="h-4 w-4 mr-1" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
      <path fillRule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clipRule="evenodd" />
    </svg>
    请输入有效的邮箱地址
  </p>
</div>
```

## 6. 可访问性指南

### 6.1 表单可访问性要求

**键盘导航：**
- 所有表单控件必须可通过键盘访问
- Tab顺序遵循逻辑顺序
- 聚焦状态清晰可见

**屏幕阅读器支持：**
- 所有表单控件有适当的标签
- 错误消息由屏幕阅读器宣读
- 必填字段适当标识
- 在必要处使用合适的ARIA属性

**表单的ARIA属性：**
- 必填字段使用`aria-required="true"`
- 有错误的字段使用`aria-invalid="true"`
- 使用`aria-describedby`将错误消息与字段关联

**Tailwind 实现：**
```jsx
<div>
  <label htmlFor="name" className="block text-sm font-medium text-gray-700 mb-1" id="name-label">
    姓名<span className="text-error-500 ml-1">*</span>
  </label>
  <input 
    id="name" 
    name="name" 
    type="text" 
    required
    aria-required="true"
    aria-labelledby="name-label"
    aria-invalid={!!error}
    aria-describedby={error ? "name-error" : undefined}
    className="block w-full h-10 px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-primary-500 focus:border-primary-500" 
  />
  {error && (
    <p className="mt-1 text-sm text-error-500" id="name-error">
      {error}
    </p>
  )}
</div>
```

### 6.2 颜色和对比度
- 颜色不是传达信息的唯一手段（图标和文本与颜色一起使用）
- 所有文本满足WCAG 2.1 AA对比度要求（普通文本4.5:1，大文本3:1）
- 焦点指示器与背景有足够的对比度

## 7. 表单状态和交互

### 7.1 加载状态
表单提交数据时：

**设计规范：**
- 提交按钮显示加载旋转图标
- 提交过程中表单控件禁用
- 可选的整个表单加载覆盖层

**Tailwind 实现：**
```jsx
<button 
  type="submit" 
  disabled={isSubmitting}
  className="px-4 py-2 text-sm bg-primary-600 text-white rounded-md hover:bg-primary-700 disabled:opacity-70 flex items-center justify-center"
>
  {isSubmitting ? (
    <>
      <svg className="animate-spin -ml-1 mr-2 h-4 w-4 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
        <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"></circle>
        <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
      </svg>
      提交中...
    </>
  ) : '提交'}
</button>
```

### 7.2 成功状态
表单成功提交后：

**设计规范：**
- 显示成功消息（吐司消息、内联消息或重定向）
- 成功指示器使用绿色（success-500）
- 成功图标（对勾）

**Tailwind 实现：**
```jsx
<div className="rounded-md bg-success-50 p-4 mb-4">
  <div className="flex">
    <div className="flex-shrink-0">
      <svg className="h-5 w-5 text-success-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
        <path fillRule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clipRule="evenodd" />
      </svg>
    </div>
    <div className="ml-3">
      <p className="text-sm font-medium text-success-800">
        表单提交成功！
      </p>
    </div>
  </div>
</div>
```

这完成了表单组件设计指南的第2部分。在下一部分中，我们将介绍动态表单、表单向导模式和性能优化等高级主题。 