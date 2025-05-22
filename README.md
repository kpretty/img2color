# img2color

一个基于Go语言的图片主色调提取工具，可以从URL获取图片并分析其主色调。本项目是基于 [anzhiyu-c/img2color-go](https://github.com/anzhiyu-c/img2color-go) 的二次开发。

## 功能特点

- 支持从URL获取图片并提取主色调
- 支持多种图片格式，包括WebP
- 提供RESTful API接口
- 支持跨域请求
- 支持Referer白名单控制
- 支持本地缓存和Redis缓存，提高性能

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

### 环境变量配置

可以通过环境变量自定义服务的行为：

- `CACHE_TYPE`: 缓存类型，可选值为 "local"（默认）或 "redis"
- `LOCAL_CACHE_CAPACITY`: 本地缓存容量，默认为 "100000"
- `REDIS_ADDR`: Redis服务器地址，默认为 "127.0.0.1:6379"
- `REDIS_PASSWORD`: Redis密码，默认为空
- `REDIS_DB`: Redis数据库编号，默认为 "0"
- `ALLOWED_REFERERS`: 允许的Referer列表，以逗号分隔，默认为 "*"（允许所有）

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

## 项目依赖

- [disintegration/imaging](https://github.com/disintegration/imaging): 图像处理库
- [lucasb-eyer/go-colorful](https://github.com/lucasb-eyer/go-colorful): 颜色处理库
- [nfnt/resize](https://github.com/nfnt/resize): 图像缩放库
- [bluele/gcache](https://github.com/bluele/gcache): 本地缓存库
- [redis/go-redis](https://github.com/redis/go-redis): Redis客户端
- [golang.org/x/image](https://golang.org/x/image): 支持WebP等格式

## 实现原理

1. 接收包含图片URL的HTTP请求
2. 下载图片并解码为图像对象
3. 将图像缩小以提高处理效率
4. 计算所有像素的RGB平均值
5. 将RGB值转换为十六进制颜色代码
6. 缓存结果并返回颜色代码作为JSON响应
