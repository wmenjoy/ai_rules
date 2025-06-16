# 反馈组件设计指南

## 1. 概述

反馈组件是用户界面中的关键元素，用于向用户传达系统状态、操作结果或请求用户输入。本指南详细说明了在 React + TypeScript + Tailwind CSS 项目中如何设计和实现这些组件，确保提供清晰、一致和及时的用户反馈。

## 2. 设计原则

反馈组件遵循以下设计原则：

- **及时性**：反馈应当在用户操作后立即呈现
- **清晰性**：信息应当简洁明了，使用户容易理解
- **相关性**：反馈内容应与触发操作直接相关
- **一致性**：反馈组件在整个应用中应保持一致的外观和行为
- **非侵入性**：反馈不应过度打断用户工作流程

## 3. 警告组件 (Alert)

警告组件用于显示重要信息，如成功消息、警告或错误通知。

### 3.1 警告变体

| 变体 | 用途 | 颜色方案 |
|------|------|---------|
| 信息 | 一般性提示 | 蓝色 |
| 成功 | 操作成功反馈 | 绿色 |
| 警告 | 需要注意的提示 | 黄色 |
| 错误 | 错误或危险操作提示 | 红色 |

### 3.2 警告组件 TypeScript 接口

```typescript
type AlertVariant = 'info' | 'success' | 'warning' | 'error';

interface AlertProps {
  variant?: AlertVariant;
  title?: React.ReactNode;
  icon?: React.ReactNode;
  action?: React.ReactNode;
  onClose?: () => void;
  isDismissible?: boolean;
  className?: string;
  children: React.ReactNode;
}
```

### 3.3 警告组件实现示例

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
import React from 'react';
import { classNames } from '../utils';
import { XIcon } from './icons';

export const Alert: React.FC<AlertProps> = ({
  variant = 'info',
  title,
  icon,
  action,
  onClose,
  isDismissible = false,
  className,
  children,
}) => {
  // 变体特定类
  const variantClasses = {
    info: 'bg-blue-50 text-blue-800',
    success: 'bg-green-50 text-green-800',
    warning: 'bg-yellow-50 text-yellow-800',
    error: 'bg-red-50 text-red-800',
  };
  
  // 图标颜色（按变体）
  const iconColors = {
    info: 'text-blue-500',
    success: 'text-green-500',
    warning: 'text-yellow-500',
    error: 'text-red-500',
  };
  
  // 边框颜色（按变体）
  const borderColors = {
    info: 'border-blue-300',
    success: 'border-green-300',
    warning: 'border-yellow-300',
    error: 'border-red-300',
  };
  
  return (
    <div
      className={classNames(
        'rounded-md border p-4',
        variantClasses[variant],
        borderColors[variant],
        className
      )}
      role="alert"
    >
      <div className="flex items-start">
        {icon && (
          <div className={classNames('flex-shrink-0 mr-3', iconColors[variant])}>
            {icon}
          </div>
        )}
        
        <div className="flex-1">
          {title && (
            <h3 className="text-sm font-medium mb-1">{title}</h3>
          )}
          
          <div className="text-sm">{children}</div>
          
          {action && (
            <div className="mt-3">{action}</div>
          )}
        </div>
        
        {isDismissible && onClose && (
          <button
            type="button"
            className={classNames(
              'ml-auto -mx-1.5 -my-1.5 rounded-md p-1.5 inline-flex focus:outline-none focus:ring-2 focus:ring-offset-2',
              `bg-${variant}-50 ${iconColors[variant]} hover:bg-${variant}-100 focus:ring-${variant}-500`
            )}
            onClick={onClose}
            aria-label="关闭"
          >
            <XIcon className="h-5 w-5" />
          </button>
        )}
      </div>
    </div>
  );
};
// [AI-BLOCK-END]
```

### 3.4 警告组件使用示例

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
import { Alert } from './components';
import { InformationCircleIcon, CheckCircleIcon, ExclamationIcon, XCircleIcon } from './icons';

const AlertsExample = () => {
  return (
    <div className="space-y-4">
      <Alert 
        variant="info" 
        title="信息提示" 
        icon={<InformationCircleIcon className="h-5 w-5" />}
        isDismissible
        onClose={() => console.log('关闭信息提示')}
      >
        这是一条普通信息提示，用于展示非关键性通知。
      </Alert>
      
      <Alert 
        variant="success" 
        title="操作成功" 
        icon={<CheckCircleIcon className="h-5 w-5" />}
      >
        您的设置已成功保存！系统将在下次登录时应用新设置。
      </Alert>
      
      <Alert 
        variant="warning" 
        title="请注意" 
        icon={<ExclamationIcon className="h-5 w-5" />}
        action={
          <button className="text-sm font-medium text-yellow-800 hover:text-yellow-700">
            查看详情
          </button>
        }
      >
        您的账户存储空间即将用尽，请考虑升级或清理不必要的文件。
      </Alert>
      
      <Alert 
        variant="error" 
        title="发生错误" 
        icon={<XCircleIcon className="h-5 w-5" />}
        isDismissible
        onClose={() => console.log('关闭错误提示')}
      >
        无法连接到服务器，请检查您的网络连接并重试。
      </Alert>
    </div>
  );
};
// [AI-BLOCK-END]
```

## 4. 提示消息组件 (Toast)

提示消息是临时性通知，通常显示在屏幕边缘，在短时间后自动消失。

### 4.1 提示消息变体

| 变体 | 用途 | 特点 |
|------|------|------|
| 信息 | 一般性通知 | 蓝色、简短、非阻断 |
| 成功 | 操作成功通知 | 绿色、短时显示 |
| 警告 | 需要关注的通知 | 黄色、持续时间较长 |
| 错误 | 操作失败通知 | 红色、持续时间更长 |

### 4.2 提示消息 TypeScript 接口

```typescript
type ToastVariant = 'info' | 'success' | 'warning' | 'error';

type ToastPosition = 
  | 'top-left'
  | 'top-center'
  | 'top-right'
  | 'bottom-left'
  | 'bottom-center'
  | 'bottom-right';

interface ToastProps {
  id: string;
  variant?: ToastVariant;
  title?: React.ReactNode;
  icon?: React.ReactNode;
  autoClose?: boolean;
  duration?: number;
  onClose: (id: string) => void;
  className?: string;
  children: React.ReactNode;
}

interface ToastContainerProps {
  position?: ToastPosition;
  className?: string;
  children: React.ReactNode;
}
```

### 4.3 提示消息实现示例

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
import React, { useEffect } from 'react';
import { classNames } from '../utils';
import { XIcon } from './icons';

export const Toast: React.FC<ToastProps> = ({
  id,
  variant = 'info',
  title,
  icon,
  autoClose = true,
  duration = 5000, // 5秒
  onClose,
  className,
  children,
}) => {
  // 变体特定类
  const variantClasses = {
    info: 'bg-blue-50 text-blue-800 border-blue-300',
    success: 'bg-green-50 text-green-800 border-green-300',
    warning: 'bg-yellow-50 text-yellow-800 border-yellow-300',
    error: 'bg-red-50 text-red-800 border-red-300',
  };
  
  // 图标颜色（按变体）
  const iconColors = {
    info: 'text-blue-500',
    success: 'text-green-500',
    warning: 'text-yellow-500',
    error: 'text-red-500',
  };
  
  // 自动关闭效果
  useEffect(() => {
    if (autoClose) {
      const timer = setTimeout(() => {
        onClose(id);
      }, duration);
      
      return () => clearTimeout(timer);
    }
  }, [autoClose, duration, id, onClose]);
  
  return (
    <div
      className={classNames(
        'max-w-sm w-full rounded-md border shadow-lg pointer-events-auto overflow-hidden',
        variantClasses[variant],
        className
      )}
      role="alert"
      aria-live="assertive"
    >
      <div className="p-4">
        <div className="flex items-start">
          {icon && (
            <div className={classNames('flex-shrink-0 mr-3', iconColors[variant])}>
              {icon}
            </div>
          )}
          
          <div className="flex-1">
            {title && (
              <h3 className="text-sm font-medium mb-1">{title}</h3>
            )}
            
            <div className="text-sm">{children}</div>
          </div>
          
          <button
            type="button"
            className={classNames(
              'ml-4 flex-shrink-0 rounded-md inline-flex text-gray-400 hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-gray-500'
            )}
            onClick={() => onClose(id)}
            aria-label="关闭"
          >
            <XIcon className="h-5 w-5" />
          </button>
        </div>
      </div>
      
      {autoClose && (
        <div 
          className={classNames(
            'h-1 bg-opacity-30',
            iconColors[variant]
          )}
          style={{
            animationName: 'toast-progress',
            animationDuration: `${duration}ms`,
            animationTimingFunction: 'linear',
            animationFillMode: 'forwards'
          }}
        />
      )}
    </div>
  );
};

// 提示消息容器组件，用于定位
export const ToastContainer: React.FC<ToastContainerProps> = ({
  position = 'bottom-right',
  className,
  children,
}) => {
  // 位置类
  const positionClasses = {
    'top-left': 'top-0 left-0',
    'top-center': 'top-0 left-1/2 transform -translate-x-1/2',
    'top-right': 'top-0 right-0',
    'bottom-left': 'bottom-0 left-0',
    'bottom-center': 'bottom-0 left-1/2 transform -translate-x-1/2',
    'bottom-right': 'bottom-0 right-0',
  };
  
  return (
    <div
      className={classNames(
        'fixed z-50 p-4 flex flex-col space-y-4 pointer-events-none',
        positionClasses[position],
        className
      )}
      aria-live="polite"
    >
      {children}
    </div>
  );
};
// [AI-BLOCK-END]
```

### 4.4 提示消息使用示例

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
import React, { useState, useEffect } from 'react';
import { Toast, ToastContainer } from './components';
import { CheckCircleIcon, XCircleIcon } from './icons';

interface ToastItem {
  id: string;
  variant: 'info' | 'success' | 'warning' | 'error';
  title: string;
  message: string;
  icon?: React.ReactNode;
}

const ToastExample = () => {
  const [toasts, setToasts] = useState<ToastItem[]>([]);
  
  // 添加提示消息
  const addToast = (toast: Omit<ToastItem, 'id'>) => {
    const id = Date.now().toString();
    setToasts(prev => [...prev, { ...toast, id }]);
  };
  
  // 移除提示消息
  const removeToast = (id: string) => {
    setToasts(prev => prev.filter(toast => toast.id !== id));
  };
  
  // 示例：添加不同类型的提示消息
  useEffect(() => {
    // 2秒后显示成功提示
    const successTimer = setTimeout(() => {
      addToast({
        variant: 'success',
        title: '保存成功',
        message: '您的文档已成功保存',
        icon: <CheckCircleIcon className="h-5 w-5" />
      });
    }, 2000);
    
    // 4秒后显示错误提示
    const errorTimer = setTimeout(() => {
      addToast({
        variant: 'error',
        title: '连接失败',
        message: '无法连接到服务器，请稍后重试',
        icon: <XCircleIcon className="h-5 w-5" />
      });
    }, 4000);
    
    return () => {
      clearTimeout(successTimer);
      clearTimeout(errorTimer);
    };
  }, []);
  
  return (
    <div>
      <div className="mb-4 space-y-2">
        <button
          onClick={() => addToast({
            variant: 'info',
            title: '信息提示',
            message: '这是一条普通信息提示'
          })}
          className="px-4 py-2 bg-blue-600 text-white rounded-md"
        >
          显示信息提示
        </button>
        
        <button
          onClick={() => addToast({
            variant: 'success',
            title: '操作成功',
            message: '您的操作已成功完成',
            icon: <CheckCircleIcon className="h-5 w-5" />
          })}
          className="px-4 py-2 bg-green-600 text-white rounded-md ml-2"
        >
          显示成功提示
        </button>
      </div>
      
      <ToastContainer position="bottom-right">
        {toasts.map(toast => (
          <Toast
            key={toast.id}
            id={toast.id}
            variant={toast.variant}
            title={toast.title}
            icon={toast.icon}
            onClose={removeToast}
            duration={toast.variant === 'error' ? 8000 : 5000}
          >
            {toast.message}
          </Toast>
        ))}
      </ToastContainer>
    </div>
  );
};
// [AI-BLOCK-END]
```

## 5. 模态框组件 (Modal)

模态框用于显示需要用户关注的内容或需要用户输入的表单，通常会阻断对背景内容的交互。

### 5.1 模态框变体

| 变体 | 用途 | 特点 |
|------|------|------|
| 信息模态框 | 显示详细信息 | 简单布局，通常只有确认按钮 |
| 确认模态框 | 请求用户确认操作 | 有确认和取消两个操作 |
| 表单模态框 | 收集用户输入 | 包含表单元素和提交按钮 |
| 全屏模态框 | 展示大量内容 | 在移动设备上占据全屏 |

### 5.2 模态框 TypeScript 接口

```typescript
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

interface ModalFooterProps {
  className?: string;
  children: React.ReactNode;
}
```

### 5.3 模态框实现示例

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
import React, { useEffect } from 'react';
import { createPortal } from 'react-dom';
import { classNames } from '../utils';
import { XIcon } from './icons';

const modalSizes = {
  sm: 'max-w-sm',
  md: 'max-w-md',
  lg: 'max-w-lg',
  xl: 'max-w-xl',
  full: 'max-w-full mx-4',
};

export const Modal: React.FC<ModalProps> = ({
  isOpen,
  onClose,
  title,
  size = 'md',
  closeOnOverlayClick = true,
  closeOnEsc = true,
  initialFocus,
  className,
  children,
}) => {
  // 处理 ESC 键按下
  useEffect(() => {
    const handleKeyDown = (event: KeyboardEvent) => {
      if (closeOnEsc && isOpen && event.key === 'Escape') {
        onClose();
      }
    };
    
    document.addEventListener('keydown', handleKeyDown);
    
    return () => {
      document.removeEventListener('keydown', handleKeyDown);
    };
  }, [closeOnEsc, isOpen, onClose]);
  
  // 焦点捕获
  useEffect(() => {
    if (isOpen) {
      // 锁定body滚动
      document.body.style.overflow = 'hidden';
      
      // 将焦点设置到模态框或指定元素
      if (initialFocus?.current) {
        initialFocus.current.focus();
      }
      
      return () => {
        document.body.style.overflow = '';
      };
    }
  }, [isOpen, initialFocus]);
  
  if (!isOpen) return null;
  
  return createPortal(
    <div className="fixed inset-0 z-50 overflow-y-auto">
      <div
        className="fixed inset-0 bg-black bg-opacity-50 transition-opacity"
        onClick={closeOnOverlayClick ? onClose : undefined}
        aria-hidden="true"
      />
      
      <div className="flex min-h-screen items-center justify-center p-4">
        <div
          className={classNames(
            'bg-white rounded-lg overflow-hidden shadow-xl transform transition-all w-full',
            modalSizes[size],
            className
          )}
          role="dialog"
          aria-modal="true"
          aria-labelledby="modal-title"
          onClick={(e) => e.stopPropagation()}
        >
          {title && (
            <div className="px-6 py-4 border-b border-gray-200 flex items-center justify-between">
              <h3 id="modal-title" className="font-medium text-lg text-gray-900">
                {title}
              </h3>
              <button
                type="button"
                className="rounded-md text-gray-400 hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-primary-500"
                onClick={onClose}
                aria-label="关闭"
              >
                <XIcon className="h-6 w-6" />
              </button>
            </div>
          )}
          
          <div className="p-6">{children}</div>
        </div>
      </div>
    </div>,
    document.body
  );
};

// 模态框底部组件（用于操作按钮）
export const ModalFooter: React.FC<ModalFooterProps> = ({
  className,
  children,
}) => {
  return (
    <div
      className={classNames(
        'px-6 py-4 border-t border-gray-200 flex justify-end space-x-3',
        className
      )}
    >
      {children}
    </div>
  );
};
// [AI-BLOCK-END]
```

### 5.4 模态框使用示例

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
import React, { useState } from 'react';
import { Modal, ModalFooter } from './components';

const ModalExample = () => {
  const [isOpen, setIsOpen] = useState(false);
  
  return (
    <div>
      <button
        onClick={() => setIsOpen(true)}
        className="px-4 py-2 bg-primary-600 text-white rounded-md"
      >
        打开模态框
      </button>
      
      <Modal
        isOpen={isOpen}
        onClose={() => setIsOpen(false)}
        title="确认操作"
        size="md"
      >
        <div>
          <p className="text-gray-700">
            您确定要执行此操作吗？此操作无法撤销。
          </p>
          
          <ModalFooter>
            <button
              type="button"
              className="px-4 py-2 border border-gray-300 rounded-md text-gray-700 hover:bg-gray-50"
              onClick={() => setIsOpen(false)}
            >
              取消
            </button>
            <button
              type="button"
              className="px-4 py-2 bg-primary-600 text-white rounded-md hover:bg-primary-700"
              onClick={() => {
                console.log('确认操作');
                setIsOpen(false);
              }}
            >
              确认
            </button>
          </ModalFooter>
        </div>
      </Modal>
    </div>
  );
};

// 表单模态框示例
const FormModalExample = () => {
  const [isOpen, setIsOpen] = useState(false);
  const [formData, setFormData] = useState({ name: '', email: '' });
  
  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData(prev => ({ ...prev, [name]: value }));
  };
  
  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    console.log('提交表单数据:', formData);
    setIsOpen(false);
  };
  
  return (
    <div>
      <button
        onClick={() => setIsOpen(true)}
        className="px-4 py-2 bg-primary-600 text-white rounded-md"
      >
        打开表单模态框
      </button>
      
      <Modal
        isOpen={isOpen}
        onClose={() => setIsOpen(false)}
        title="用户信息"
        size="md"
      >
        <form onSubmit={handleSubmit}>
          <div className="space-y-4">
            <div>
              <label className="block text-sm font-medium text-gray-700">姓名</label>
              <input
                type="text"
                name="name"
                value={formData.name}
                onChange={handleChange}
                className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md"
              />
            </div>
            
            <div>
              <label className="block text-sm font-medium text-gray-700">邮箱</label>
              <input
                type="email"
                name="email"
                value={formData.email}
                onChange={handleChange}
                className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md"
              />
            </div>
          </div>
          
          <ModalFooter>
            <button
              type="button"
              className="px-4 py-2 border border-gray-300 rounded-md text-gray-700 hover:bg-gray-50"
              onClick={() => setIsOpen(false)}
            >
              取消
            </button>
            <button
              type="submit"
              className="px-4 py-2 bg-primary-600 text-white rounded-md hover:bg-primary-700"
            >
              保存
            </button>
          </ModalFooter>
        </form>
      </Modal>
    </div>
  );
};
// [AI-BLOCK-END]
```

## 6. 反馈组件最佳实践

### 6.1 组件选择指南

| 场景 | 推荐组件 | 理由 |
|------|---------|------|
| 页面内静态通知 | Alert | 持久显示，不会自动消失 |
| 操作结果反馈 | Toast | 临时显示，不打断工作流 |
| 需要用户确认的操作 | Modal | 强制用户关注并做出决定 |
| 表单提交错误 | Alert (inline) | 直接在表单附近显示错误信息 |
| 系统状态变化 | Toast | 非阻断式通知用户状态变化 |

### 6.2 可访问性考虑

- 使用适当的ARIA角色和属性（`role="alert"`、`aria-live`等）
- 确保模态框可通过键盘操作（Tab导航、Esc关闭）
- 为所有提示消息提供足够的显示时间
- 使用足够的颜色对比度
- 避免仅依赖颜色传达信息（使用图标和文本辅助）

### 6.3 响应式设计

- 在移动设备上，提示消息应考虑较小的屏幕尺寸
- 模态框在小屏幕上可考虑全屏显示
- 确保触摸目标足够大（至少44×44像素）
- 多层反馈组件（如Toast）应避免覆盖重要UI元素

### 6.4 性能优化

- 避免同时显示过多Toast消息
- 使用React Portal确保模态框正确渲染在DOM层次结构中
- 考虑延迟加载模态框内容
- 使用CSS动画而非JavaScript动画以获得更好性能

### 6.5 用户体验增强

- 提供适当的动画和过渡效果
- 确保反馈信息简洁明了
- 提供撤销操作的机会（特别是对于Toast通知）
- 避免过度使用模态框打断用户工作流
- 提供清晰的错误信息和恢复建议 