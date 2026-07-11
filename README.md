# wvp-go - GB28181 视频监控平台

<div align="center">
  <img src="https://img.shields.io/badge/golang-1.22-blue" />
  <img src="https://img.shields.io/badge/gin-1.9.1-lightBlue" />
  <img src="https://img.shields.io/badge/vue-3.3.4-brightgreen" />
  <img src="https://img.shields.io/badge/gorm-1.25.2-red" />
  <img src="https://img.shields.io/badge/status-development-yellow" />
</div>

<p align="center">
  基于 gin-vue-admin 框架的 GB28181 视频监控平台 Go 语言实现
</p>

---

## 项目介绍

wvp-go 是 [wvp-GB28181-pro](https://github.com/panjjo/wvp-GB28181-pro) 的 Go 语言移植版本，基于 [gin-vue-admin](https://github.com/flipped-aurora/gin-vue-admin) 框架构建，提供了完整的 GB28181 视频监控功能。

### 核心功能

- **SIP 协议栈**: 设备注册、心跳、目录查询、云台控制
- **视频点播**: INVITE/SDP 协商、实时视频播放、BYE 停止
- **PTZ 控制**: 上下左右、变焦、聚焦、光圈、预置位
- **平台级联**: 向上级注册、通道同步、心跳维护
- **报警处理**: 报警消息接收、存储、处理
- **移动位置**: GPS 信息上报、轨迹查询
- **录像服务**: 录像计划管理、录像存储(本地/S3)、回放、下载
- **ZLMediaKit 对接**: 流媒体服务器 HTTP API 集成

### 技术栈

| 组件 | 技术 | 用途 |
|------|------|------|
| 后端框架 | Gin | RESTful API |
| 数据库 | MySQL/GORM | 数据持久化 |
| 缓存 | Redis | JWT、会话管理 |
| 前端 | Vue 3 + Element Plus | 管理界面 |
| 流媒体 | ZLMediaKit | 视频推拉流 |
| SIP | 自定义协议栈 | GB28181 信令 |
| 日志 | Zap | 结构化日志 |

---

## 快速开始

### 环境要求

- Go >= 1.22
- Node.js >= 18
- MySQL >= 5.7
- Redis
- ZLMediaKit

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

---

## 项目结构

```
wvp-go/
├── server/                    # 后端 Go 代码
│   ├── api/v1/gb28181/       # API 处理器
│   │   ├── device.go         # 设备管理 API
│   │   ├── video.go          # 视频点播 API
│   │   ├── platform.go       # 平台级联 API
│   │   ├── alarm.go          # 报警管理 API
│   │   ├── position.go       # 移动位置 API
│   │   └── record.go         # 录像管理 API
│   ├── config/               # 配置管理
│   │   ├── config.yaml       # 配置文件模板
│   │   └── wvp_config.go     # WVP 配置结构体
│   ├── model/system/         # 数据模型
│   │   ├── device.go         # 设备/通道模型
│   │   ├── platform.go       # 平台/通道模型
│   │   ├── alarm.go          # 报警模型
│   │   ├── position.go       # 位置模型
│   │   └── record.go         # 录像模型
│   ├── pkg/sip/              # SIP 协议栈
│   │   ├── parser.go         # 消息解析器
│   │   ├── builder.go        # 消息构建器
│   │   ├── transport.go      # UDP/TCP 传输
│   │   └── transaction.go    # 事务管理
│   ├── internal/
│   │   ├── gb28181/          # GB28181 核心逻辑
│   │   │   ├── sip/          # SIP 服务器/客户端
│   │   │   ├── message/      # 消息处理器(注册/心跳/目录/PTZ/报警/位置)
│   │   │   ├── session/      # 会话管理(INVITE)
│   │   │   └── platform/     # 平台级联
│   │   ├── media/zlm/        # ZLMediaKit 对接
│   │   └── record/           # 录像服务
│   ├── router/gb28181/       # 路由定义
│   └── utils/                # 工具函数
│       ├── response/         # 统一响应
│       └── wvplog/           # 结构化日志
├── web/                      # 前端 Vue 代码
│   ├── src/api/              # 前端 API 模块
│   ├── src/view/             # 页面组件
│   └── src/router/           # 路由配置
├── migrations/               # 数据库迁移
└── deploy/                   # 部署配置
    ├── docker/               # Dockerfile
    └── docker-compose/       # Docker Compose
```

---

## API 文档

启动后端后访问 Swagger 文档:

```
http://localhost:8888/swagger/index.html
```

### 生成 Swagger 文档

```bash
cd server
swag init
```

### API 概览

| 模块 | 方法 | 路径 | 说明 |
|------|------|------|------|
| 设备 | GET | /api/device/list | 设备列表 |
| 设备 | POST | /api/device | 添加设备 |
| 设备 | PUT | /api/device/:id | 更新设备 |
| 设备 | DELETE | /api/device/:id | 删除设备 |
| 设备 | GET | /api/device/:id/channels | 通道列表 |
| 设备 | POST | /api/device/:id/catalog | 查询目录 |
| 视频 | POST | /api/video/play | 视频点播 |
| 视频 | POST | /api/video/stop/:id | 停止播放 |
| 视频 | POST | /api/video/ptz/:id | 云台控制 |
| 平台 | GET | /api/platform/list | 平台列表 |
| 平台 | POST | /api/platform | 添加平台 |
| 平台 | POST | /api/platform/:id/sync | 同步目录 |
| 平台 | POST | /api/platform/:id/start | 启动级联 |
| 报警 | GET | /api/alarm/list | 报警列表 |
| 报警 | POST | /api/alarm/:id/handle | 处理报警 |
| 位置 | GET | /api/position/list | 位置列表 |
| 位置 | GET | /api/position/:id/track | 位置轨迹 |
| 录像 | GET | /api/record/plan/list | 录像计划列表 |
| 录像 | POST | /api/record/plan | 创建计划 |
| 录像 | POST | /api/record/playback | 录像回放 |
| 录像 | GET | /api/record/file/download/:id | 文件下载 |

---

## 部署

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

### 配置 Nginx

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

---

## 许可证

本项目基于 BSL 1.1 许可证。
