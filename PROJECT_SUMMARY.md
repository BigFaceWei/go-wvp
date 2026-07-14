# wvp-go 项目总结

## 项目概述

wvp-go 是一个基于 **gin-vue-admin** 框架的 **GB28181 视频监控平台**，使用 Go 语言实现。它是 [wvp-GB28181-pro](https://github.com/panjjo/wvp-GB28181-pro) 的 Go 语言移植版本，提供了完整的 GB28181 视频监控功能。

## 核心功能

### 1. SIP 协议栈
- 设备注册与认证
- 心跳管理
- 目录查询
- 云台控制 (PTZ)

### 2. 视频点播
- INVITE/SDP 协商
- 实时视频播放
- BYE 停止播放
- 多流支持

### 3. PTZ 控制
- 上下左右移动
- 变焦、聚焦、光圈控制
- 预置位管理

### 4. 平台级联
- 向上级平台注册
- 通道同步
- 心跳维护
- 多级平台支持

### 5. 报警处理
- 报警消息接收
- 报警存储
- 报警处理流程

### 6. 移动位置
- GPS 信息上报
- 轨迹查询
- 位置历史记录

### 7. 录像服务
- 录像计划管理
- 录像存储 (本地/S3)
- 录像回放
- 文件下载

### 8. ZLMediaKit 对接
- 流媒体服务器 HTTP API 集成
- 视频推拉流支持

## 技术栈

### 后端技术

| 组件 | 技术 | 用途 |
|------|------|------|
| 编程语言 | Go 1.24+ | 后端开发 |
| Web 框架 | Gin | RESTful API |
| 数据库 ORM | GORM | 数据持久化 |
| 数据库 | MySQL/PostgreSQL/SQL Server | 数据存储 |
| 缓存 | Redis | JWT、会话管理 |
| 日志 | Zap | 结构化日志 |
| 配置管理 | Viper | 配置文件管理 |
| API 文档 | Swagger | API 文档生成 |
| 权限管理 | Casbin | 访问控制 |
| 定时任务 | robfig/cron | 定时任务调度 |

### 前端技术

| 组件 | 技术 | 用途 |
|------|------|------|
| 框架 | Vue 3.5+ | 前端开发 |
| UI 库 | Element Plus | 界面组件 |
| 构建工具 | Vite 8+ | 项目构建 |
| 状态管理 | Pinia | 状态管理 |
| 路由 | Vue Router | 路由管理 |
| HTTP 客户端 | Axios | API 调用 |
| 图表 | ECharts | 数据可视化 |
| CSS 框架 | UnoCSS | 样式工具 |

### 流媒体技术

| 组件 | 技术 | 用途 |
|------|------|------|
| 流媒体服务器 | ZLMediaKit | 视频推拉流 |
| 协议 | GB28181 | 视频监控协议 |
| 信令 | SIP | 会话初始化协议 |

## 项目结构

```
wvp-go/
├── server/                    # 后端 Go 代码
│   ├── api/v1/gb28181/       # API 处理器
│   ├── config/               # 配置管理
│   ├── core/                 # 核心功能
│   ├── global/               # 全局变量
│   ├── initialize/           # 系统初始化
│   ├── internal/             # 内部实现
│   │   ├── gb28181/          # GB28181 核心逻辑
│   │   ├── media/            # 媒体处理
│   │   └── record/           # 录像服务
│   ├── middleware/            # 中间件
│   ├── model/                # 数据模型
│   ├── pkg/                  # 工具包
│   │   └── sip/              # SIP 协议栈
│   ├── router/               # 路由定义
│   ├── service/              # 业务逻辑
│   └── utils/                # 工具函数
├── web/                      # 前端 Vue 代码
│   ├── src/
│   │   ├── api/              # 前端 API 模块
│   │   ├── view/             # 页面组件
│   │   └── router/           # 路由配置
│   └── package.json
├── migrations/               # 数据库迁移
└── deploy/                   # 部署配置
    ├── docker/               # Dockerfile
    └── docker-compose/       # Docker Compose
```

## 环境要求

### 开发环境

- **Go**: >= 1.22
- **Node.js**: >= 18
- **MySQL**: >= 5.7
- **Redis**: 最新版
- **ZLMediaKit**: 最新版

### 运行环境

- **操作系统**: Linux/Windows/macOS
- **数据库**: MySQL 5.7+ / PostgreSQL / SQL Server
- **缓存**: Redis
- **流媒体**: ZLMediaKit

## 快速开始

### 1. 数据库初始化

```bash
mysql -u root -p < migrations/001_init.sql
mysql -u root -p < migrations/002_seed.sql
```

### 2. 配置

编辑 `server/config/config.yaml`:

```yaml
sip:
  sip_ip: "0.0.0.0"
  sip_port: 5060
  domain: "3402000000"
  server_id: "34020000002000000001"

media:
  zlm_host: "127.0.0.1"
  zlm_port: 8080
  zlm_secret: "your_secret"

mysql:
  host: "127.0.0.1"
  port: 3306
  dbname: "wvp"
  username: "root"
  password: "12345678"
```

### 3. 启动后端

```bash
cd server
go generate
go run .
```

### 4. 启动前端

```bash
cd web
npm install
npm run serve
```

访问 `http://localhost:8080` 进入管理界面。

## API 文档

启动后端后访问 Swagger 文档:

```
http://localhost:8888/swagger/index.html
```

### 主要 API 模块

| 模块 | 方法 | 路径 | 说明 |
|------|------|------|------|
| 设备 | GET | /api/device/list | 设备列表 |
| 设备 | POST | /api/device | 添加设备 |
| 设备 | PUT | /api/device/:id | 更新设备 |
| 设备 | DELETE | /api/device/:id | 删除设备 |
| 视频 | POST | /api/video/play | 视频点播 |
| 视频 | POST | /api/video/stop/:id | 停止播放 |
| 视频 | POST | /api/video/ptz/:id | 云台控制 |
| 平台 | GET | /api/platform/list | 平台列表 |
| 平台 | POST | /api/platform | 添加平台 |
| 平台 | POST | /api/platform/:id/sync | 同步目录 |
| 报警 | GET | /api/alarm/list | 报警列表 |
| 报警 | POST | /api/alarm/:id/handle | 处理报警 |
| 位置 | GET | /api/position/list | 位置列表 |
| 录像 | GET | /api/record/plan/list | 录像计划列表 |
| 录像 | POST | /api/record/playback | 录像回放 |

## 部署方式

### Docker Compose

```bash
cd deploy/docker-compose
docker-compose up -d
```

### 裸机部署

1. 安装 MySQL、Redis、ZLMediaKit
2. 配置 `config.yaml`
3. 运行 `go run .` 启动后端
4. 运行 `npm run build && npm run serve` 启动前端

### Nginx 配置

```nginx
server {
    listen 80;
    server_name your-domain.com;

    location /api/ {
        proxy_pass http://127.0.0.1:8888;
    }

    location / {
        root /path/to/web/dist;
        try_files $uri $uri/ /index.html;
    }
}
```

## 开发指南

### 代码规范

- 遵循 Go 语言标准规范
- 使用 gofmt 格式化代码
- 使用 goimports 管理导入
- 使用 golangci-lint 进行代码检查

### 测试

```bash
cd server
go test ./...
```

### 构建

```bash
# 后端构建
cd server
go build -o wvp-go .

# 前端构建
cd web
npm run build
```

## 特性亮点

1. **高性能**: Go 语言编写，性能优异
2. **易扩展**: 基于 gin-vue-admin 框架，易于扩展
3. **完整功能**: 覆盖 GB28181 标准的主要功能
4. **多数据库支持**: 支持 MySQL、PostgreSQL、SQL Server
5. **云存储**: 支持阿里云 OSS、AWS S3、腾讯 COS 等
6. **容器化部署**: 提供 Docker 和 Docker Compose 配置
7. **API 文档**: 自动生成 Swagger API 文档
8. **权限管理**: 基于 Casbin 的细粒度权限控制

## 许可证

本项目基于 BSL 1.1 许可证。

## 相关链接

- [gin-vue-admin 框架](https://github.com/flipped-aurora/gin-vue-admin)
- [wvp-GB28181-pro 原版](https://github.com/panjjo/wvp-GB28181-pro)
- [ZLMediaKit 流媒体服务器](https://github.com/ZLMediaKit/ZLMediaKit)
- [GB28181 标准文档](https://baike.baidu.com/item/GB/T 28181-2016)
