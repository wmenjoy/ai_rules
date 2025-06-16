# 导航组件设计指南

## 1. 概述

导航组件是用户界面的核心元素，它们帮助用户在应用程序中定位自己并在不同部分之间移动。本指南详细说明了在 React + TypeScript + Tailwind CSS 项目中如何设计和实现常见的导航组件。

## 2. 设计原则

导航组件遵循以下设计原则：

- **清晰性**：用户应始终知道自己在应用中的位置
- **一致性**：导航元素应在整个应用中保持一致的外观和行为
- **简洁性**：避免过度复杂的导航结构，保持层次清晰
- **可访问性**：确保所有导航元素都可通过键盘操作和辅助技术访问
- **响应式**：导航组件应适应不同屏幕尺寸，提供良好的移动体验

## 3. 标签页组件 (Tabs)

标签页用于在同一视图中切换不同类别的内容，保持上下文的连续性。

### 3.1 标签页变体

| 变体 | 特点 | 使用场景 |
|------|------|---------|
| 下划线标签页 | 简洁，强调当前选项通过下划线 | 内容密集的界面 |
| 容器式标签页 | 使用背景色和阴影区分 | 需要更强视觉分隔的场景 |
| 胶囊标签页 | 圆角矩形样式 | 强调分类或筛选操作 |

### 3.2 标签页 TypeScript 接口

```typescript
interface TabItemProps {
  id: string;
  label: React.ReactNode;
  icon?: React.ReactNode;
  disabled?: boolean;
  count?: number;
}

interface TabsProps {
  items: TabItemProps[];
  activeTab: string;
  onChange: (tabId: string) => void;
  variant?: 'underlined' | 'contained' | 'pills';
  size?: 'sm' | 'md' | 'lg';
  fullWidth?: boolean;
  orientation?: 'horizontal' | 'vertical';
  className?: string;
}
```

### 3.3 标签页实现示例

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
import React from 'react';
import { classNames } from '../utils';

export const Tabs: React.FC<TabsProps> = ({
  items,
  activeTab,
  onChange,
  variant = 'underlined',
  size = 'md',
  fullWidth = false,
  orientation = 'horizontal',
  className,
}) => {
  // 尺寸类
  const sizeClasses = {
    sm: 'text-sm',
    md: 'text-base',
    lg: 'text-lg',
  };
  
  // 变体特定类
  const variantTabListClasses = {
    underlined: 'border-b border-gray-200',
    contained: 'bg-gray-100 p-1 rounded-lg',
    pills: 'space-x-2',
  };
  
  // 激活标签类
  const getActiveTabClasses = (active: boolean) => {
    if (!active) return '';
    
    switch (variant) {
      case 'underlined':
        return 'border-b-2 border-primary-500 text-primary-600';
      case 'contained':
        return 'bg-white shadow text-primary-600';
      case 'pills':
        return 'bg-primary-500 text-white';
      default:
        return '';
    }
  };
  
  // 未激活标签类
  const getInactiveTabClasses = () => {
    switch (variant) {
      case 'underlined':
        return 'border-b-2 border-transparent hover:border-gray-300 hover:text-gray-700';
      case 'contained':
        return 'hover:text-gray-700';
      case 'pills':
        return 'hover:bg-gray-100';
      default:
        return '';
    }
  };
  
  // 禁用标签类
  const getDisabledTabClasses = () => {
    return 'opacity-50 cursor-not-allowed';
  };
  
  // 方向特定类
  const orientationClasses = {
    horizontal: 'flex-row',
    vertical: 'flex-col',
  };
  
  return (
    <div className={classNames(orientation === 'vertical' ? 'flex' : '', className)}>
      <div
        className={classNames(
          'flex',
          orientationClasses[orientation],
          variantTabListClasses[variant],
          fullWidth && orientation === 'horizontal' ? 'w-full' : '',
          orientation === 'vertical' ? 'flex-shrink-0' : ''
        )}
        role="tablist"
      >
        {items.map((item) => (
          <button
            key={item.id}
            role="tab"
            aria-selected={activeTab === item.id}
            aria-controls={`${item.id}-panel`}
            id={`${item.id}-tab`}
            disabled={item.disabled}
            className={classNames(
              'px-4 py-2 font-medium transition-colors focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500',
              sizeClasses[size],
              fullWidth && orientation === 'horizontal' ? 'flex-1' : '',
              orientation === 'vertical' ? 'text-left' : 'text-center',
              item.disabled ? getDisabledTabClasses() : activeTab === item.id ? getActiveTabClasses(true) : getInactiveTabClasses()
            )}
            onClick={() => {
              if (!item.disabled) {
                onChange(item.id);
              }
            }}
          >
            <div className="flex items-center justify-center">
              {item.icon && <span className="mr-2">{item.icon}</span>}
              {item.label}
              {item.count !== undefined && (
                <span className={classNames(
                  'ml-2 px-2 py-0.5 text-xs font-medium rounded-full',
                  activeTab === item.id 
                    ? variant === 'pills' ? 'bg-white bg-opacity-30 text-white' : 'bg-primary-100 text-primary-700' 
                    : 'bg-gray-100 text-gray-700'
                )}>
                  {item.count}
                </span>
              )}
            </div>
          </button>
        ))}
      </div>
    </div>
  );
};

// TabPanel 组件
interface TabPanelProps {
  id: string;
  active: boolean;
  children: React.ReactNode;
}

export const TabPanel: React.FC<TabPanelProps> = ({
  id,
  active,
  children,
}) => {
  if (!active) return null;
  
  return (
    <div
      id={`${id}-panel`}
      role="tabpanel"
      aria-labelledby={`${id}-tab`}
    >
      {children}
    </div>
  );
};
// [AI-BLOCK-END]
```

### 3.4 标签页使用示例

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
import React, { useState } from 'react';
import { Tabs, TabPanel } from './components';

const TabsExample = () => {
  const [activeTab, setActiveTab] = useState('details');
  
  const tabItems = [
    { id: 'details', label: '基本信息' },
    { id: 'settings', label: '设置', count: 3 },
    { id: 'security', label: '安全', icon: <LockIcon /> },
    { id: 'notifications', label: '通知', disabled: true }
  ];
  
  return (
    <div className="space-y-4">
      <Tabs
        items={tabItems}
        activeTab={activeTab}
        onChange={setActiveTab}
        variant="underlined"
      />
      
      <TabPanel id="details" active={activeTab === 'details'}>
        <div className="py-4">
          <h3 className="text-lg font-medium">基本信息内容</h3>
          <p className="mt-2 text-gray-600">这里显示用户的基本信息...</p>
        </div>
      </TabPanel>
      
      <TabPanel id="settings" active={activeTab === 'settings'}>
        <div className="py-4">
          <h3 className="text-lg font-medium">设置内容</h3>
          <p className="mt-2 text-gray-600">这里显示用户的设置选项...</p>
        </div>
      </TabPanel>
      
      <TabPanel id="security" active={activeTab === 'security'}>
        <div className="py-4">
          <h3 className="text-lg font-medium">安全内容</h3>
          <p className="mt-2 text-gray-600">这里显示安全相关的设置...</p>
        </div>
      </TabPanel>
      
      <TabPanel id="notifications" active={activeTab === 'notifications'}>
        <div className="py-4">
          <h3 className="text-lg font-medium">通知内容</h3>
          <p className="mt-2 text-gray-600">这里显示通知相关的设置...</p>
        </div>
      </TabPanel>
    </div>
  );
};
// [AI-BLOCK-END]
```

## 4. 面包屑组件 (Breadcrumbs)

面包屑导航提供用户在应用层次结构中的位置路径，帮助用户理解当前页面与整体架构的关系。

### 4.1 面包屑 TypeScript 接口

```typescript
interface BreadcrumbItemProps {
  label: React.ReactNode;
  href?: string;
  icon?: React.ReactNode;
  isCurrent?: boolean;
}

interface BreadcrumbsProps {
  items: BreadcrumbItemProps[];
  separator?: React.ReactNode;
  className?: string;
}
```

### 4.2 面包屑实现示例

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
import React from 'react';
import { classNames } from '../utils';

export const Breadcrumbs: React.FC<BreadcrumbsProps> = ({
  items,
  separator = '/',
  className,
}) => {
  return (
    <nav className={classNames('flex', className)} aria-label="Breadcrumb">
      <ol className="flex items-center space-x-2">
        {items.map((item, index) => (
          <li key={index} className="flex items-center">
            {index > 0 && (
              <span className="mx-2 text-gray-400">{separator}</span>
            )}
            
            {item.isCurrent ? (
              <span
                className="text-gray-700 font-medium"
                aria-current="page"
              >
                {item.icon && <span className="mr-1">{item.icon}</span>}
                {item.label}
              </span>
            ) : (
              <a
                href={item.href}
                className="text-gray-500 hover:text-gray-700 hover:underline"
              >
                {item.icon && <span className="mr-1">{item.icon}</span>}
                {item.label}
              </a>
            )}
          </li>
        ))}
      </ol>
    </nav>
  );
};
// [AI-BLOCK-END]
```

### 4.3 面包屑使用示例

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
import { Breadcrumbs } from './components';
import { HomeIcon, FolderIcon, DocumentIcon } from './icons';

const BreadcrumbsExample = () => {
  const breadcrumbItems = [
    { 
      label: '首页', 
      href: '/', 
      icon: <HomeIcon className="w-4 h-4" />
    },
    {
      label: '项目文档',
      href: '/docs',
      icon: <FolderIcon className="w-4 h-4" />
    },
    {
      label: '设计指南',
      isCurrent: true,
      icon: <DocumentIcon className="w-4 h-4" />
    }
  ];
  
  return (
    <Breadcrumbs items={breadcrumbItems} />
  );
};
// [AI-BLOCK-END]
```

## 5. 分页组件 (Pagination)

分页组件允许用户在大型数据集或长列表中导航，显示当前页码并允许用户切换页面。

### 5.1 分页 TypeScript 接口

```typescript
interface PaginationProps {
  currentPage: number;
  totalPages: number;
  onPageChange: (page: number) => void;
  siblingCount?: number;
  size?: 'sm' | 'md' | 'lg';
  className?: string;
}
```

### 5.2 分页实现示例

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
import React from 'react';
import { classNames } from '../utils';

export const Pagination: React.FC<PaginationProps> = ({
  currentPage,
  totalPages,
  onPageChange,
  siblingCount = 1,
  size = 'md',
  className,
}) => {
  // 生成要显示的页码
  const getPageNumbers = () => {
    const totalPageNumbers = siblingCount * 2 + 3; // 两侧的兄弟页码 + 第一页 + 最后一页 + 当前页
    
    // 如果总页数小于要显示的页码数，显示所有页码
    if (totalPages <= totalPageNumbers) {
      return Array.from({ length: totalPages }, (_, i) => i + 1);
    }
    
    const leftSiblingIndex = Math.max(currentPage - siblingCount, 1);
    const rightSiblingIndex = Math.min(currentPage + siblingCount, totalPages);
    
    const shouldShowLeftDots = leftSiblingIndex > 2;
    const shouldShowRightDots = rightSiblingIndex < totalPages - 1;
    
    // 始终显示第一页和最后一页
    const firstPageIndex = 1;
    const lastPageIndex = totalPages;
    
    // 左侧无省略号
    if (!shouldShowLeftDots && shouldShowRightDots) {
      const leftItemCount = 3 + 2 * siblingCount;
      const leftRange = Array.from({ length: leftItemCount }, (_, i) => i + 1);
      
      return [...leftRange, '...', lastPageIndex];
    }
    
    // 右侧无省略号
    if (shouldShowLeftDots && !shouldShowRightDots) {
      const rightItemCount = 3 + 2 * siblingCount;
      const rightRange = Array.from(
        { length: rightItemCount },
        (_, i) => totalPages - rightItemCount + i + 1
      );
      
      return [firstPageIndex, '...', ...rightRange];
    }
    
    // 两侧都有省略号
    if (shouldShowLeftDots && shouldShowRightDots) {
      const middleRange = Array.from(
        { length: rightSiblingIndex - leftSiblingIndex + 1 },
        (_, i) => leftSiblingIndex + i
      );
      
      return [firstPageIndex, '...', ...middleRange, '...', lastPageIndex];
    }
    
    return [];
  };
  
  const pageNumbers = getPageNumbers();
  
  // 尺寸类
  const sizeClasses = {
    sm: 'h-8 w-8 text-sm',
    md: 'h-10 w-10 text-base',
    lg: 'h-12 w-12 text-lg',
  };
  
  // 按钮基础类
  const buttonBaseClasses = 'inline-flex items-center justify-center rounded-md font-medium focus:outline-none focus:ring-2 focus:ring-primary-500 focus:ring-offset-2';
  
  return (
    <nav
      className={classNames('flex items-center justify-center', className)}
      aria-label="Pagination"
    >
      {/* 上一页按钮 */}
      <button
        onClick={() => onPageChange(currentPage - 1)}
        disabled={currentPage === 1}
        className={classNames(
          buttonBaseClasses,
          sizeClasses[size],
          currentPage === 1 
            ? 'cursor-not-allowed text-gray-300' 
            : 'text-gray-700 hover:bg-gray-100'
        )}
        aria-label="上一页"
      >
        &lt;
      </button>
      
      {/* 页码 */}
      <div className="flex items-center mx-2">
        {pageNumbers.map((pageNumber, index) => (
          <React.Fragment key={index}>
            {pageNumber === '...' ? (
              <span className="mx-1 text-gray-400">...</span>
            ) : (
              <button
                onClick={() => onPageChange(Number(pageNumber))}
                className={classNames(
                  buttonBaseClasses,
                  sizeClasses[size],
                  'mx-0.5',
                  Number(pageNumber) === currentPage
                    ? 'bg-primary-500 text-white'
                    : 'text-gray-700 hover:bg-gray-100'
                )}
                aria-current={Number(pageNumber) === currentPage ? 'page' : undefined}
              >
                {pageNumber}
              </button>
            )}
          </React.Fragment>
        ))}
      </div>
      
      {/* 下一页按钮 */}
      <button
        onClick={() => onPageChange(currentPage + 1)}
        disabled={currentPage === totalPages}
        className={classNames(
          buttonBaseClasses,
          sizeClasses[size],
          currentPage === totalPages 
            ? 'cursor-not-allowed text-gray-300' 
            : 'text-gray-700 hover:bg-gray-100'
        )}
        aria-label="下一页"
      >
        &gt;
      </button>
    </nav>
  );
};
// [AI-BLOCK-END]
```

### 5.3 分页使用示例

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
import React, { useState } from 'react';
import { Pagination } from './components';

const PaginationExample = () => {
  const [currentPage, setCurrentPage] = useState(1);
  const totalPages = 10;
  
  return (
    <div className="space-y-4">
      <p className="text-center">当前显示第 {currentPage} 页，共 {totalPages} 页</p>
      <Pagination
        currentPage={currentPage}
        totalPages={totalPages}
        onPageChange={setCurrentPage}
      />
    </div>
  );
};
// [AI-BLOCK-END]
```

## 6. 导航菜单组件 (NavMenu)

导航菜单用于应用程序的主要导航，提供应用程序不同部分之间的链接。

### 6.1 导航菜单 TypeScript 接口

```typescript
interface NavItemProps {
  label: React.ReactNode;
  href: string;
  icon?: React.ReactNode;
  isActive?: boolean;
  hasSubItems?: boolean;
  subItems?: Omit<NavItemProps, 'subItems'>[];
}

interface NavMenuProps {
  items: NavItemProps[];
  orientation?: 'horizontal' | 'vertical';
  variant?: 'light' | 'dark' | 'transparent';
  fullWidth?: boolean;
  className?: string;
}
```

### 6.2 导航菜单实现示例

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
import React, { useState } from 'react';
import { classNames } from '../utils';

export const NavMenu: React.FC<NavMenuProps> = ({
  items,
  orientation = 'horizontal',
  variant = 'light',
  fullWidth = false,
  className,
}) => {
  const [openSubMenus, setOpenSubMenus] = useState<Record<string, boolean>>({});
  
  // 切换子菜单的开/关状态
  const toggleSubMenu = (itemHref: string) => {
    setOpenSubMenus(prev => ({
      ...prev,
      [itemHref]: !prev[itemHref],
    }));
  };
  
  // 变体类
  const variantClasses = {
    light: 'bg-white text-gray-700',
    dark: 'bg-gray-800 text-white',
    transparent: 'bg-transparent text-current',
  };
  
  // 方向类
  const orientationClasses = {
    horizontal: 'flex-row',
    vertical: 'flex-col',
  };
  
  // 导航项目渲染
  const renderNavItem = (item: NavItemProps, isSubItem = false) => {
    // 适用于所有项目的基础类
    const baseItemClasses = classNames(
      'flex items-center px-4 py-2 font-medium rounded-md transition-colors',
      !isSubItem && orientation === 'horizontal' && 'h-16',
      isSubItem ? 'text-sm pl-8' : '',
      item.isActive 
        ? 'text-primary-600 bg-primary-50' 
        : 'hover:bg-gray-100',
      item.hasSubItems ? 'cursor-pointer' : ''
    );
    
    // 内容包装器
    const contentWrapper = (
      <>
        {item.icon && <span className="mr-3">{item.icon}</span>}
        <span>{item.label}</span>
        {item.hasSubItems && (
          <svg 
            className={classNames(
              'ml-2 h-4 w-4 transition-transform',
              openSubMenus[item.href] ? 'transform rotate-180' : ''
            )} 
            fill="none" 
            viewBox="0 0 24 24" 
            stroke="currentColor"
          >
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 9l-7 7-7-7" />
          </svg>
        )}
      </>
    );
    
    return (
      <li key={item.href} className="relative">
        {item.hasSubItems ? (
          <div
            className={baseItemClasses}
            onClick={() => toggleSubMenu(item.href)}
          >
            {contentWrapper}
          </div>
        ) : (
          <a
            href={item.href}
            className={baseItemClasses}
            aria-current={item.isActive ? 'page' : undefined}
          >
            {contentWrapper}
          </a>
        )}
        
        {/* 子菜单 */}
        {item.hasSubItems && item.subItems && openSubMenus[item.href] && (
          <ul className="mt-1 ml-2 space-y-1">
            {item.subItems.map(subItem => renderNavItem(subItem, true))}
          </ul>
        )}
      </li>
    );
  };
  
  return (
    <nav
      className={classNames(
        variantClasses[variant],
        orientation === 'horizontal' && 'border-b border-gray-200',
        className
      )}
    >
      <ul
        className={classNames(
          'flex',
          orientationClasses[orientation],
          orientation === 'horizontal' ? 'items-center space-x-2' : 'space-y-1',
          fullWidth && orientation === 'horizontal' ? 'w-full' : ''
        )}
      >
        {items.map(item => renderNavItem(item))}
      </ul>
    </nav>
  );
};
// [AI-BLOCK-END]
```

### 6.3 导航菜单使用示例

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
import { NavMenu } from './components';
import { HomeIcon, UserIcon, DocumentIcon, CogIcon } from './icons';

const NavMenuExample = () => {
  const navItems = [
    {
      label: '首页',
      href: '/',
      icon: <HomeIcon className="w-5 h-5" />,
      isActive: true
    },
    {
      label: '文档',
      href: '/docs',
      icon: <DocumentIcon className="w-5 h-5" />,
      hasSubItems: true,
      subItems: [
        { label: '入门指南', href: '/docs/getting-started' },
        { label: '组件', href: '/docs/components' },
        { label: 'API 参考', href: '/docs/api' }
      ]
    },
    {
      label: '用户',
      href: '/users',
      icon: <UserIcon className="w-5 h-5" />
    },
    {
      label: '设置',
      href: '/settings',
      icon: <CogIcon className="w-5 h-5" />
    }
  ];
  
  return (
    <NavMenu items={navItems} />
  );
};
// [AI-BLOCK-END]
```

## 7. 移动导航组件 (MobileNav)

移动导航组件专为小屏幕设备设计，通常采用汉堡菜单样式和折叠导航。

### 7.1 移动导航 TypeScript 接口

```typescript
interface MobileNavProps {
  items: NavItemProps[];
  logo?: React.ReactNode;
  className?: string;
}
```

### 7.2 移动导航实现示例

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
import React, { useState } from 'react';
import { classNames } from '../utils';

export const MobileNav: React.FC<MobileNavProps> = ({
  items,
  logo,
  className,
}) => {
  const [isOpen, setIsOpen] = useState(false);
  
  return (
    <nav className={classNames('bg-white', className)}>
      {/* 导航栏 */}
      <div className="px-4 flex items-center justify-between h-16 border-b border-gray-200">
        {/* Logo */}
        <div className="flex-shrink-0">
          {logo}
        </div>
        
        {/* 汉堡菜单按钮 */}
        <button
          type="button"
          className="p-2 rounded-md text-gray-500 hover:text-gray-700 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-primary-500"
          onClick={() => setIsOpen(!isOpen)}
          aria-expanded={isOpen}
        >
          <span className="sr-only">{isOpen ? '关闭菜单' : '打开菜单'}</span>
          <svg
            className="h-6 w-6"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
            aria-hidden="true"
          >
            {isOpen ? (
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M6 18L18 6M6 6l12 12" />
            ) : (
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 6h16M4 12h16M4 18h16" />
            )}
          </svg>
        </button>
      </div>
      
      {/* 移动菜单 */}
      <div
        className={classNames(
          'fixed inset-0 z-40 bg-white transform transition-transform ease-in-out duration-300',
          isOpen ? 'translate-x-0' : '-translate-x-full',
          'pt-16' // 为顶部导航栏留出空间
        )}
      >
        <div className="h-full overflow-y-auto">
          <ul className="px-2 pt-2 pb-3 space-y-1">
            {items.map((item) => (
              <li key={item.href}>
                <a
                  href={item.href}
                  className={classNames(
                    'block px-3 py-2 rounded-md text-base font-medium',
                    item.isActive
                      ? 'bg-primary-50 text-primary-600'
                      : 'text-gray-700 hover:bg-gray-100'
                  )}
                  aria-current={item.isActive ? 'page' : undefined}
                  onClick={() => {
                    if (!item.hasSubItems) {
                      setIsOpen(false);
                    }
                  }}
                >
                  <div className="flex items-center">
                    {item.icon && <span className="mr-3">{item.icon}</span>}
                    {item.label}
                  </div>
                </a>
                
                {/* 子项目 */}
                {item.hasSubItems && item.subItems && (
                  <ul className="mt-1 ml-4 space-y-1">
                    {item.subItems.map((subItem) => (
                      <li key={subItem.href}>
                        <a
                          href={subItem.href}
                          className={classNames(
                            'block px-3 py-2 rounded-md text-sm font-medium',
                            subItem.isActive
                              ? 'bg-primary-50 text-primary-600'
                              : 'text-gray-700 hover:bg-gray-100'
                          )}
                          aria-current={subItem.isActive ? 'page' : undefined}
                          onClick={() => setIsOpen(false)}
                        >
                          {subItem.label}
                        </a>
                      </li>
                    ))}
                  </ul>
                )}
              </li>
            ))}
          </ul>
        </div>
      </div>
      
      {/* 背景遮罩 */}
      {isOpen && (
        <div
          className="fixed inset-0 z-30 bg-gray-500 bg-opacity-50 transition-opacity"
          onClick={() => setIsOpen(false)}
        ></div>
      )}
    </nav>
  );
};
// [AI-BLOCK-END]
```

### 7.3 移动导航使用示例

```tsx
// [AI-BLOCK-START] - 生成工具: Claude 3.7 Sonnet
import { MobileNav } from './components';
import { HomeIcon, UserIcon, DocumentIcon, CogIcon } from './icons';

const MobileNavExample = () => {
  const navItems = [
    {
      label: '首页',
      href: '/',
      icon: <HomeIcon className="w-5 h-5" />,
      isActive: true
    },
    {
      label: '文档',
      href: '/docs',
      icon: <DocumentIcon className="w-5 h-5" />,
      hasSubItems: true,
      subItems: [
        { label: '入门指南', href: '/docs/getting-started' },
        { label: '组件', href: '/docs/components' },
        { label: 'API 参考', href: '/docs/api' }
      ]
    },
    {
      label: '用户',
      href: '/users',
      icon: <UserIcon className="w-5 h-5" />
    },
    {
      label: '设置',
      href: '/settings',
      icon: <CogIcon className="w-5 h-5" />
    }
  ];
  
  return (
    <MobileNav
      items={navItems}
      logo={<img src="/logo.svg" alt="Logo" className="h-8 w-auto" />}
    />
  );
};
// [AI-BLOCK-END]
```

## 8. 导航组件使用场景与最佳实践

### 8.1 使用场景选择

| 导航组件 | 适用场景 | 注意事项 |
|---------|---------|---------|
| 标签页 | 在同一视图中显示相关内容分类 | 避免过多标签，建议不超过7个 |
| 面包屑 | 深层次的信息层级展示 | 保持层级清晰，避免过深的结构 |
| 分页 | 大型列表和数据集 | 确保当前页码清晰可见 |
| 导航菜单 | 主要应用导航 | 突出显示当前位置 |
| 移动导航 | 响应式设计中的小屏幕导航 | 简化结构，确保可触摸区域足够大 |

### 8.2 可访问性考虑

- 所有导航元素应支持键盘导航
- 使用适当的 ARIA 属性（如 `aria-current`、`aria-expanded`）
- 确保足够的颜色对比度
- 提供清晰的焦点状态
- 对于移动导航，确保足够大的点击区域

### 8.3 响应式设计策略

- **小屏幕**：
  - 标签页可折叠为下拉菜单
  - 面包屑可简化为仅显示上一级和当前级
  - 分页可简化为仅显示上一页/下一页按钮
  - 导航菜单转换为汉堡菜单

- **中等屏幕**：
  - 标签页可滚动而非折叠
  - 面包屑显示完整但使用较紧凑的样式
  - 分页显示有限数量的页码

- **大屏幕**：
  - 显示完整导航结构
  - 更宽松的间距和布局

### 8.4 性能优化

- 导航组件应轻量化，避免不必要的复杂计算
- 考虑为大型导航结构使用代码分割
- 使用 `React.memo` 优化导航组件渲染

### 8.5 扩展与定制

导航组件系统可通过以下方式扩展：
- 添加更多变体以满足特定设计需求
- 实现动画和过渡效果增强用户体验
- 集成主题切换支持
- 添加国际化支持
- 实现上下文感知导航（基于用户权限或历史记录） 