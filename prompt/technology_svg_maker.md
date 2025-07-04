
生成技术图解的SVG提示语

``` 提示语
请为`content`创建一系列技术图解SVG，要求如下：

## 基础规范
- 画布尺寸：1200x800
- 背景色：使用#f8f9fa、#fafafa、#f9fafb等浅灰色系
- 每个SVG需包含清晰的注释标记（如：<!-- ch1_overview.svg -->）

## 视觉风格
- 标题：居中显示，font-size="28"，font-weight="bold"，颜色#1a1a1a
- 主要模块：使用圆角矩形（rx="15"），配合柔和色彩+opacity="0.9"
- 配色方案：
  - 蓝色系：#A5B4FC, #93C5FD, #A7D7C5  (淡雅的天蓝和薄荷青)
  - 绿色系：#A7F3D0, #BBF7D0, #D1FAE5  (清新的薄荷绿)
  - 紫色系：#C4B5FD, #D8B4FE, #E9D5FF  (温柔的薰衣草紫)
  - 橙色系：#FDE68A, #FED7AA, #FFE4E1  (温暖的杏色和淡粉)
  - 红色系：#FECACA, #FBCFE8, #F9A8D4  (柔和的珊瑚粉和樱花粉)
  - 中性色：#94A3B8, #A1A1AA, #D1D5DB  (比原来更浅的灰色，与淡彩更搭)


## 文字层级
- 模块标题：font-size="16-20"，font-weight="bold"，白色或深色
- 主要内容：font-size="12-14"
- 次要说明：font-size="11"
- 注释文字：font-size="10"
- 文字方向：尽可能采用横向，除了连线中间的文字可以视情况采用跟随连线方向的文字朝向

## 布局原则
- 采用模块化网格布局，充分利用1200x800空间
- 主要内容区域留出上边距80px（标题占用）
- 模块化原则
  - 同层模块margin、padding、align尽可能保持一致性
  - 为了避免模块纷杂，需对模块进行分层分类，比如1、2、3模块同属于某个分类A，则可以用一个大的模块A包裹这些模块以提高信息表达效率
  - 对于一些非主体重点内容，可以额外用辅助性模块额外旁路去展示，避免因辅助性模块引入导致整体视图混乱。
- 连线原则
  - 使用连接线、箭头展示关系（定义marker和path）
  - 对于复杂连线的场景，需先评估连线的样式跟位置，避免连线混乱交叉导致不易理解
  - 连线中间可以增加辅助文本节点进一步表达，但需考虑文字不要跟其他模块重叠
- 重叠交叉二次校验
  - 需对是否出现重叠交叉的情况进行二次修正
  - 文字是否重叠交叉
  - 模块是否重叠交叉
  - 连线是否重叠交叉
  
## 内容呈现决策指南

### 信息类型与可视化映射
根据`content`的不同信息类型，选择合适的呈现方式，对于非常复杂的场景，可能需结合多个呈现方式

1. **整体概览/生态系统类**
   - 采用中心辐射式布局
   - 核心概念置于中心，相关维度环绕
   - 适合展示：技术栈全景、系统架构概览、概念关系网

2. **时间演进/发展历程类**
   - 横向时间轴为主线
   - 关键节点用色块标注，配合简短说明
   - 适合展示：技术迭代、版本演进、里程碑事件

3. **多方案对比/选型类**
   - 并列式卡片布局，每个方案独立色系
   - 底部配置特性对比表格
   - 适合展示：技术选型、方案比较、优劣分析

4. **流程步骤/工作流类**
   - 箭头连接的线性或分支流程
   - 关键节点详细展开说明
   - 适合展示：处理流程、方法步骤、数据流向

5. **层次结构/依赖关系类**
   - 垂直分层或树状结构
   - 用连线表示依赖和交互
   - 适合展示：系统分层、组件依赖、继承关系

6. **矩阵分类/多维度类**
   - 网格化布局，分组展示
   - 同类用相近色系，便于识别
   - 适合展示：技术分类、功能矩阵、特性集合

### 内容组织原则

1. **信息密度控制**
   - 主要观点：大字体、醒目色块
   - 支撑细节：小字体、次要位置
   - 数据指标：独立框体展示

2. **视觉层次设计**
   - 用颜色深浅区分重要程度
   - 用大小变化体现主次关系
   - 用位置布局引导阅读顺序

3. **关联性表达**
   - 相关内容：使用相同色系
   - 对比内容：使用对比色
   - 流程关系：使用箭头连线
   - 包含关系：使用嵌套布局

### 呈现角度建议

基于`content`的特点，灵活选择和组合以下呈现角度：
- 宏观视角：全局概览、体系架构
- 时间视角：演进历程、发展趋势
- 比较视角：方案对比、优劣分析
- 深入视角：核心原理、技术细节
- 实践视角：最佳实践、应用案例
- 前瞻视角：发展方向、未来展望

每个SVG应该有明确的信息重点，避免面面俱到。根据内容特性选择最合适的可视化方式，确保信息传达的有效性和美观性。

## 特殊元素
- 使用<g>标签组织相关元素，添加id标识
- 数据表格：用线条+文字构建，表头加粗
- 流程箭头：定义marker，使用path绘制
- 强调框：使用带stroke的圆角矩形
- 图标效果：圆形配数字表示步骤，圆点表示列表项

## 信息密度
- 每个SVG应包含丰富信息，避免空旷
- 合理使用子标题、说明文字、数据标注
- 重要概念用视觉层次突出
- 保持整体的专业性和技术深度

## 设计原则总结
- 形式服务于内容：先分析信息类型，再选择呈现方式
- 保持视觉一致性：统一的色彩体系和布局规范
- 注重信息层级：通过视觉手段建立清晰的阅读顺序
- 平衡美观与实用：既要视觉吸引力，更要信息传达效率
```