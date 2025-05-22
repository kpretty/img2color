# img2color

一个基于Go语言的图片主色调提取工具，可以从URL获取图片并分析其主色调。本项目是基于 [anzhiyu-c/img2color-go](https://github.com/anzhiyu-c/img2color-go) 的二次开发。

## 功能特点

- 支持从URL获取图片并提取主色调
- 支持多种图片格式，包括WebP
- 提供RESTful API接口
- 支持跨域请求
- 支持Referer白名单控制

## 安装

### 前置条件

- Go 1.24.3 或更高版本

### 安装步骤

```bash
# 克隆仓库
git clone https://github.com/yourusername/img2color.git
cd img2color

# 安装依赖
go mod download

# 编译
go build -o img2color
```

## 使用方法

### 启动服务

```bash
./img2color
```

服务默认在8080端口启动。

### API使用

#### 获取图片主色调

```
GET /api?img=图片URL
```

**请求参数：**

- `img`: 图片的URL地址（必填）

**响应示例：**

```json
{
  "RGB": "#3a8fb7"
}
```

**响应说明：**

- `RGB`: 图片的主色调，以十六进制颜色代码表示

## 实现原理

1. 接收包含图片URL的HTTP请求
2. 下载图片并解码为图像对象
3. 将图像缩小以提高处理效率
4. 计算所有像素的RGB平均值
5. 将RGB值转换为十六进制颜色代码
6. 返回颜色代码作为JSON响应
