# Demo2 - Go Web Server

这是一个使用 Go 标准库 net/http 创建的简单 Web 服务器项目。

## 项目结构

```
demo2/
├── main.go           # 服务器入口文件
├── go.mod            # Go 模块文件
├── static/           # 静态文件目录
│   └── index.html    # 主页面
├── .gitignore        # Git 忽略文件
└── README.md         # 项目说明文档
```

## 功能特点

- 使用 Go 标准库 net/http 创建 HTTP 服务器
- 支持静态文件服务
- HTTP 路由处理
- 响应式 HTML 页面设计
- 简洁的代码结构

## 运行项目

### 1. 安装依赖

```bash
go mod tidy
```

### 2. 运行服务器

```bash
go run main.go
```

### 3. 访问应用

在浏览器中打开：[http://localhost:8080](http://localhost:8080)

## 技术细节

- **Go 版本**: 1.21
- **HTTP 端口**: 8080
- **主要包**: net/http

## 代码说明

### main.go

- `homeHandler()`: 处理根路径请求，返回 HTML 页面
- `main()`: 初始化 HTTP 路由，启动服务器

### HTTP 路由

- `/`: 主页面
- `/static/`: 静态文件服务

## 扩展建议

- 添加更多路由处理器
- 实现动态内容渲染
- 添加 JSON API 支持
- 集成数据库访问
- 添加中间件支持

## 许可证

MIT License
