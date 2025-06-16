# 表单组件设计指南 - 第三部分

## 8. 高级表单模式

### 8.1 动态表单

动态表单允许用户根据需求添加或删除表单元素，通常用于收集数量不确定的数据项。

```typescript
interface DynamicFieldProps {
  fields: Array<{
    id: string;
    value: string;
  }>;
  onAdd: () => void;
  onRemove: (id: string) => void;
  onChange: (id: string, value: string) => void;
  label: string;
  addButtonText?: string;
  maxFields?: number;
  minFields?: number;
}
```

**设计规范：**
- 每个字段包含输入框和删除按钮
- 底部有"添加"按钮，点击后生成新字段
- 当达到最大字段数量时禁用添加按钮
- 当剩余最小字段数量时禁用删除按钮

**Tailwind 实现：**
```jsx
<div className="space-y-4">
  <label className="block text-sm font-medium text-gray-700">{label}</label>
  
  {fields.map((field) => (
    <div key={field.id} className="flex items-center space-x-2">
      <input
        type="text"
        value={field.value}
        onChange={(e) => onChange(field.id, e.target.value)}
        className="block w-full h-10 px-3 py-2 border border-gray-300 rounded-md text-sm"
      />
      <button
        type="button"
        onClick={() => onRemove(field.id)}
        disabled={fields.length <= (minFields || 1)}
        className="p-2 text-gray-500 hover:text-gray-700 disabled:opacity-50 disabled:cursor-not-allowed"
        aria-label="删除字段"
      >
        <svg className="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
        </svg>
      </button>
    </div>
  ))}
  
  <button
    type="button"
    onClick={onAdd}
    disabled={maxFields !== undefined && fields.length >= maxFields}
    className="flex items-center text-sm text-primary-600 hover:text-primary-700 disabled:opacity-50 disabled:cursor-not-allowed"
  >
    <svg className="h-4 w-4 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
    </svg>
    {addButtonText || "添加新字段"}
  </button>
</div>
```

### 8.2 表单向导 (Form Wizard)

表单向导将复杂表单分解为多个步骤，引导用户逐步完成。

```typescript
interface FormWizardProps {
  steps: Array<{
    id: string;
    title: string;
    description?: string;
    component: React.ReactNode;
  }>;
  currentStep: number;
  onNext: () => void;
  onPrevious: () => void;
  onComplete: () => void;
  isStepValid?: (stepIndex: number) => boolean;
  isSubmitting?: boolean;
}
```

**设计规范：**
- 顶部显示步骤指示器，突出显示当前步骤
- 内容区域显示当前步骤的表单
- 底部有导航按钮（上一步、下一步、完成）
- 支持步骤验证，防止用户在填写无效表单时继续

**Tailwind 实现：**
```jsx
<div className="max-w-3xl mx-auto">
  {/* 步骤指示器 */}
  <nav aria-label="进度" className="mb-8">
    <ol className="flex items-center">
      {steps.map((step, index) => (
        <li key={step.id} className={`relative ${index !== steps.length - 1 ? 'pr-8 sm:pr-20' : ''}`}>
          {/* 连接线 */}
          {index !== steps.length - 1 && (
            <div className="absolute top-4 left-4 -ml-px mt-0.5 h-0.5 w-full sm:w-20 bg-gray-300">
              <div 
                className="h-0.5 bg-primary-600 transition-all duration-300 ease-in-out" 
                style={{ width: currentStep > index ? '100%' : '0%' }} 
              />
            </div>
          )}
          
          {/* 步骤圆点 */}
          <div className="group relative flex items-start">
            <span className="flex items-center h-9">
              <span 
                className={`relative z-10 flex h-8 w-8 items-center justify-center rounded-full ${
                  currentStep > index 
                    ? 'bg-primary-600' 
                    : currentStep === index 
                      ? 'border-2 border-primary-600 bg-white' 
                      : 'border-2 border-gray-300 bg-white'
                }`}
              >
                {currentStep > index ? (
                  <svg className="h-5 w-5 text-white" viewBox="0 0 20 20" fill="currentColor">
                    <path fillRule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clipRule="evenodd" />
                  </svg>
                ) : (
                  <span className={`${currentStep === index ? 'text-primary-600' : 'text-gray-500'}`}>
                    {index + 1}
                  </span>
                )}
              </span>
            </span>
            
            {/* 步骤文本 */}
            <span className="ml-3 mt-0.5">
              <span className={`block text-sm font-medium ${
                currentStep >= index ? 'text-primary-600' : 'text-gray-500'
              }`}>
                {step.title}
              </span>
              {step.description && (
                <span className="block text-xs text-gray-500">
                  {step.description}
                </span>
              )}
            </span>
          </div>
        </li>
      ))}
    </ol>
  </nav>
  
  {/* 当前步骤内容 */}
  <div className="bg-white shadow-sm rounded-md p-6 mb-6">
    {steps[currentStep].component}
  </div>
  
  {/* 导航按钮 */}
  <div className="flex justify-between">
    <button
      type="button"
      onClick={onPrevious}
      disabled={currentStep === 0}
      className="px-4 py-2 text-sm bg-white border border-gray-300 rounded-md text-gray-700 hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
    >
      上一步
    </button>
    
    {currentStep < steps.length - 1 ? (
      <button
        type="button"
        onClick={onNext}
        disabled={isStepValid && !isStepValid(currentStep)}
        className="px-4 py-2 text-sm bg-primary-600 text-white rounded-md hover:bg-primary-700 disabled:opacity-50 disabled:cursor-not-allowed"
      >
        下一步
      </button>
    ) : (
      <button
        type="button"
        onClick={onComplete}
        disabled={isSubmitting || (isStepValid && !isStepValid(currentStep))}
        className="px-4 py-2 text-sm bg-primary-600 text-white rounded-md hover:bg-primary-700 disabled:opacity-50 disabled:cursor-not-allowed"
      >
        {isSubmitting ? (
          <>
            <svg className="animate-spin -ml-1 mr-2 h-4 w-4 text-white inline-block" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
              <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"></circle>
              <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            提交中...
          </>
        ) : '完成'}
      </button>
    )}
  </div>
</div>
```

### 8.3 条件表单字段

根据用户的选择或输入，动态显示或隐藏特定表单字段。

**设计规范：**
- 平滑过渡效果，避免突兀的显示/隐藏
- 条件字段应在逻辑上与触发字段关联
- 提供清晰的视觉提示表明字段间的关系

**Tailwind 实现：**
```jsx
<div className="space-y-4">
  {/* 触发字段 */}
  <div>
    <div className="flex items-center">
      <input
        id="show-additional"
        type="checkbox"
        checked={showAdditional}
        onChange={(e) => setShowAdditional(e.target.checked)}
        className="h-4 w-4 text-primary-600 focus:ring-primary-500 border-gray-300 rounded"
      />
      <label htmlFor="show-additional" className="ml-2 block text-sm text-gray-700">
        需要添加更多信息？
      </label>
    </div>
  </div>
  
  {/* 条件字段 - 使用过渡效果 */}
  <div className={`space-y-4 overflow-hidden transition-all duration-300 ease-in-out ${
    showAdditional ? 'max-h-96 opacity-100' : 'max-h-0 opacity-0'
  }`}>
    {/* 只有当showAdditional为true时，这些字段才会渲染在DOM中 */}
    {showAdditional && (
      <>
        <div>
          <label htmlFor="additional-info" className="block text-sm font-medium text-gray-700">
            附加信息
          </label>
          <input
            id="additional-info"
            type="text"
            className="mt-1 block w-full h-10 px-3 py-2 border border-gray-300 rounded-md text-sm"
          />
        </div>
        
        <div>
          <label htmlFor="additional-details" className="block text-sm font-medium text-gray-700">
            详细说明
          </label>
          <textarea
            id="additional-details"
            rows={3}
            className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md text-sm"
          ></textarea>
        </div>
      </>
    )}
  </div>
</div>
```

## 9. 表单数据处理

### 9.1 表单状态管理

#### 使用React Hook Form

```tsx
import { useForm } from 'react-hook-form';
import { yupResolver } from '@hookform/resolvers/yup';
import * as yup from 'yup';

const schema = yup.object({
  name: yup.string().required('姓名是必填项'),
  email: yup.string().email('请输入有效的邮箱地址').required('邮箱是必填项'),
  age: yup.number().positive('年龄必须是正数').integer('年龄必须是整数').required('年龄是必填项'),
}).required();

export const UserForm = () => {
  const { register, handleSubmit, formState: { errors } } = useForm({
    resolver: yupResolver(schema)
  });
  
  const onSubmit = (data) => {
    console.log(data);
  };
  
  return (
    <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
      <div>
        <label htmlFor="name" className="block text-sm font-medium text-gray-700">姓名</label>
        <input
          id="name"
          {...register('name')}
          className={`mt-1 block w-full h-10 px-3 py-2 border ${errors.name ? 'border-error-500' : 'border-gray-300'} rounded-md text-sm`}
        />
        {errors.name && <p className="mt-1 text-sm text-error-500">{errors.name.message}</p>}
      </div>
      
      <div>
        <label htmlFor="email" className="block text-sm font-medium text-gray-700">邮箱</label>
        <input
          id="email"
          type="email"
          {...register('email')}
          className={`mt-1 block w-full h-10 px-3 py-2 border ${errors.email ? 'border-error-500' : 'border-gray-300'} rounded-md text-sm`}
        />
        {errors.email && <p className="mt-1 text-sm text-error-500">{errors.email.message}</p>}
      </div>
      
      <div>
        <label htmlFor="age" className="block text-sm font-medium text-gray-700">年龄</label>
        <input
          id="age"
          type="number"
          {...register('age')}
          className={`mt-1 block w-full h-10 px-3 py-2 border ${errors.age ? 'border-error-500' : 'border-gray-300'} rounded-md text-sm`}
        />
        {errors.age && <p className="mt-1 text-sm text-error-500">{errors.age.message}</p>}
      </div>
      
      <button
        type="submit"
        className="px-4 py-2 text-sm bg-primary-600 text-white rounded-md hover:bg-primary-700"
      >
        提交
      </button>
    </form>
  );
};
```

### 9.2 表单数据转换

```typescript
interface FormDataTransformer<T, U> {
  // 将API数据转换为表单数据
  fromApi: (apiData: T) => U;
  // 将表单数据转换为API数据
  toApi: (formData: U) => T;
}

// 示例：用户资料表单
interface UserApiData {
  user_id: number;
  first_name: string;
  last_name: string;
  birth_date: string; // ISO格式: "1990-01-01"
  address: {
    street: string;
    city: string;
    postal_code: string;
    country: string;
  };
}

interface UserFormData {
  id: number;
  fullName: string;
  birthDate: Date;
  streetAddress: string;
  city: string;
  postalCode: string;
  country: string;
}

// 数据转换器
const userDataTransformer: FormDataTransformer<UserApiData, UserFormData> = {
  fromApi: (apiData) => ({
    id: apiData.user_id,
    fullName: `${apiData.first_name} ${apiData.last_name}`,
    birthDate: new Date(apiData.birth_date),
    streetAddress: apiData.address.street,
    city: apiData.address.city,
    postalCode: apiData.address.postal_code,
    country: apiData.address.country,
  }),
  
  toApi: (formData) => {
    // 分割全名为姓和名
    const nameParts = formData.fullName.split(' ');
    const lastName = nameParts.pop() || '';
    const firstName = nameParts.join(' ');
    
    return {
      user_id: formData.id,
      first_name: firstName,
      last_name: lastName,
      birth_date: formData.birthDate.toISOString().split('T')[0],
      address: {
        street: formData.streetAddress,
        city: formData.city,
        postal_code: formData.postalCode,
        country: formData.country,
      }
    };
  }
};
```

## 10. 表单性能优化

### 10.1 表单渲染优化

**设计规范：**
- 避免不必要的重新渲染
- 使用记忆化技术减少渲染开销
- 延迟加载复杂表单组件

**实现示例：**
```tsx
import React, { useState, useMemo, useCallback } from 'react';

// 优化表单组件
const OptimizedForm = () => {
  const [formData, setFormData] = useState({
    name: '',
    email: '',
    message: ''
  });
  
  // 使用useCallback优化事件处理函数
  const handleChange = useCallback((e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  }, []);
  
  // 使用useMemo优化表单验证
  const formErrors = useMemo(() => {
    const errors: Record<string, string> = {};
    
    if (!formData.name) {
      errors.name = '请输入姓名';
    }
    
    if (!formData.email) {
      errors.email = '请输入邮箱';
    } else if (!/\S+@\S+\.\S+/.test(formData.email)) {
      errors.email = '邮箱格式不正确';
    }
    
    return errors;
  }, [formData.name, formData.email]);
  
  // 优化提交按钮状态计算
  const isFormValid = useMemo(() => {
    return Object.keys(formErrors).length === 0 && 
           formData.name !== '' && 
           formData.email !== '';
  }, [formErrors, formData.name, formData.email]);
  
  return (
    <form className="space-y-4">
      <div>
        <label htmlFor="name" className="block text-sm font-medium text-gray-700">姓名</label>
        <input
          id="name"
          name="name"
          value={formData.name}
          onChange={handleChange}
          className={`mt-1 block w-full h-10 px-3 py-2 border ${formErrors.name ? 'border-error-500' : 'border-gray-300'} rounded-md text-sm`}
        />
        {formErrors.name && <p className="mt-1 text-sm text-error-500">{formErrors.name}</p>}
      </div>
      
      <div>
        <label htmlFor="email" className="block text-sm font-medium text-gray-700">邮箱</label>
        <input
          id="email"
          name="email"
          type="email"
          value={formData.email}
          onChange={handleChange}
          className={`mt-1 block w-full h-10 px-3 py-2 border ${formErrors.email ? 'border-error-500' : 'border-gray-300'} rounded-md text-sm`}
        />
        {formErrors.email && <p className="mt-1 text-sm text-error-500">{formErrors.email}</p>}
      </div>
      
      <div>
        <label htmlFor="message" className="block text-sm font-medium text-gray-700">留言</label>
        <textarea
          id="message"
          name="message"
          rows={4}
          value={formData.message}
          onChange={handleChange}
          className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md text-sm"
        />
      </div>
      
      <button
        type="submit"
        disabled={!isFormValid}
        className="px-4 py-2 text-sm bg-primary-600 text-white rounded-md hover:bg-primary-700 disabled:opacity-50 disabled:cursor-not-allowed"
      >
        提交
      </button>
    </form>
  );
};
```

### 10.2 大型表单处理策略

**设计规范：**
- 拆分为多个小型表单或步骤
- 实现表单分段加载
- 使用虚拟滚动处理长列表选项

**实现示例：**
```tsx
import React, { Suspense, lazy } from 'react';

// 懒加载表单部分
const BasicInfoSection = lazy(() => import('./BasicInfoSection'));
const AddressSection = lazy(() => import('./AddressSection'));
const PreferencesSection = lazy(() => import('./PreferencesSection'));

const LargeForm = () => {
  const [activeSection, setActiveSection] = useState('basic');
  
  return (
    <div className="space-y-6">
      {/* 分段导航 */}
      <div className="flex border-b border-gray-200">
        <button
          type="button"
          className={`py-4 px-6 focus:outline-none ${
            activeSection === 'basic' 
              ? 'border-b-2 border-primary-500 text-primary-600' 
              : 'text-gray-500 hover:text-gray-700'
          }`}
          onClick={() => setActiveSection('basic')}
        >
          基本信息
        </button>
        <button
          type="button"
          className={`py-4 px-6 focus:outline-none ${
            activeSection === 'address' 
              ? 'border-b-2 border-primary-500 text-primary-600' 
              : 'text-gray-500 hover:text-gray-700'
          }`}
          onClick={() => setActiveSection('address')}
        >
          地址信息
        </button>
        <button
          type="button"
          className={`py-4 px-6 focus:outline-none ${
            activeSection === 'preferences' 
              ? 'border-b-2 border-primary-500 text-primary-600' 
              : 'text-gray-500 hover:text-gray-700'
          }`}
          onClick={() => setActiveSection('preferences')}
        >
          个人偏好
        </button>
      </div>
      
      {/* 表单部分 - 使用Suspense实现懒加载 */}
      <div className="py-4">
        <Suspense fallback={<div className="text-center py-4">加载中...</div>}>
          {activeSection === 'basic' && <BasicInfoSection />}
          {activeSection === 'address' && <AddressSection />}
          {activeSection === 'preferences' && <PreferencesSection />}
        </Suspense>
      </div>
      
      {/* 导航按钮 */}
      <div className="flex justify-between">
        <button
          type="button"
          onClick={() => {
            if (activeSection === 'address') setActiveSection('basic');
            else if (activeSection === 'preferences') setActiveSection('address');
          }}
          disabled={activeSection === 'basic'}
          className="px-4 py-2 text-sm bg-white border border-gray-300 rounded-md text-gray-700 hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
        >
          上一步
        </button>
        
        <button
          type="button"
          onClick={() => {
            if (activeSection === 'basic') setActiveSection('address');
            else if (activeSection === 'address') setActiveSection('preferences');
          }}
          disabled={activeSection === 'preferences'}
          className="px-4 py-2 text-sm bg-primary-600 text-white rounded-md hover:bg-primary-700 disabled:opacity-50 disabled:cursor-not-allowed"
        >
          下一步
        </button>
      </div>
    </div>
  );
};
```

## 11. 表单实践建议

### 11.1 表单可用性最佳实践

1. **适当的标签位置**
   - 垂直表单使用顶部标签
   - 水平紧凑表单可使用左侧标签
   - 确保标签与输入字段的关联清晰

2. **合理的字段分组**
   - 相关字段应该归类在一起
   - 使用视觉分隔符（卡片、分割线）区分不同分组
   - 为每个分组提供描述性标题

3. **简化表单**
   - 仅收集必要信息
   - 减少必填字段数量
   - 合并可合并的字段（如"全名"替代"姓"和"名"）

4. **清晰的动作按钮**
   - 主要操作按钮应视觉突出
   - 危险操作应有确认步骤
   - 按钮文本应明确表达操作（"保存"、"提交申请"而非"确定"）

### 11.2 表单测试清单

1. **功能性测试**
   - 所有表单控件是否正常工作
   - 验证规则是否正确执行
   - 提交处理是否正常

2. **可访问性测试**
   - 使用键盘可以访问所有控件
   - 屏幕阅读器能正确朗读表单元素
   - 错误消息是否能被辅助技术识别

3. **响应式测试**
   - 在各种屏幕尺寸下表单布局是否合理
   - 触摸设备上表单控件尺寸是否足够大
   - 虚拟键盘弹出时表单是否正确滚动

4. **性能测试**
   - 加载时间是否可接受
   - 输入响应是否流畅
   - 大型表单是否有性能问题

### 11.3 用户体验提升技巧

1. **即时反馈**
   - 输入时提供实时验证反馈
   - 使用图标、颜色和微动画增强反馈
   - 成功提交后提供清晰确认

2. **智能默认值**
   - 预填用户可能选择的值
   - 使用上下文信息设置默认值（如当前位置）
   - 记住用户之前的选择

3. **渐进式表单**
   - 先展示简单字段，逐步引入复杂选项
   - 使用条件逻辑隐藏不相关的字段
   - 允许用户保存进度并稍后继续

4. **减少认知负荷**
   - 一次只展示相关表单字段
   - 提供清晰的说明和示例
   - 使用自动完成减少输入量

## 12. 结语

本设计指南提供了构建高质量表单组件的全面框架。通过遵循这些设计规范和最佳实践，可以创建出既美观又易用的表单，有效提升用户体验。

表单设计应该关注以下核心原则：
- 简单易用：减少用户认知负担
- 响应迅速：提供即时反馈
- 容错设计：帮助用户避免和修正错误
- 灵活适应：适应不同设备和用户需求
- 可访问性：确保所有用户都能使用

随着设计系统的发展，本指南也将不断更新，融入新的设计思路和技术实践，持续改进表单组件体验。 