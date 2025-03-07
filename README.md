# ksjx open platform

本项目为高性能 API 网关构建方案，实现了以下核心功能：
- 动态路由系统（从配置中心加载路由配置）
- JWT 认证及 RBAC 权限校验
- 流量控制：令牌桶限流和 Hystrix 熔断
- 性能监控：基于 Prometheus 的指标采集
- 服务发现与负载均衡（模拟 Nacos 获取实例并随机选择）
- 多级缓存设计（内存缓存 + Redis 连接池）
- 反向代理实现
- gRPC 协议支持（通过 Protobuf 定义）

本项目使用 Go 语言（1.21+）开发，无 Kubernetes 部署，适合直接在服务器上运行。

## 构建与运行

1. 安装依赖：
   - Go 1.21+
   - Redis 服务器（默认地址 localhost:6379）
2. 编译项目：
   ```bash
    go build -o ksjx-open .
   ```
3. 设置环境变量（可选，默认值见 config/config.go）：

   - SERVER_ADDRESS：服务监听地址，如 `:8080`
   - JWT_SECRET：JWT 密钥，如 `your-secret-key`
   - REDIS_ADDR：Redis 地址，如 `localhost:6379`
   - CONFIG_CENTER_URL：配置中心地址（示例值）
   - NACOS_URL：Nacos 服务地址（示例值）

4. 启动程序：

   ```bash
   ./ksjx-open
   ```

5. 运行测试：

   ```bash
   go test ./...
   ```

## 文件结构

```
├── README.md             # 项目描述文件
├── main.go               # 程序入口文件
├── config
│   └── config.go         # 配置加载
├── router
│   └── router.go         # 动态路由初始化
├── middleware
│   ├── auth.go           # JWT 认证及权限校验中间件
│   ├── limit.go          # 令牌桶限流与 Hystrix 熔断中间件
│   └── metrics.go        # Prometheus 指标采集中间件
├── service
│   ├── discovery.go      # 模拟 Nacos 服务发现
│   ├── loadbalancer.go   # 负载均衡策略
│   └── cache.go          # 内存缓存与 Redis 连接池
├── proxy
│   └── proxy.go          # 反向代理实现
├── api
│   └── api.proto         # gRPC 协议定义
└── tests
    └── analysis_test.go  # 单元测试与分析验证
```