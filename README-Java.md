# KSJX API Gateway - Spring Boot Version

这是基于Go原型的Spring Boot版本API网关实现。

## 功能特性

- **服务发现**: 基于Nacos的服务发现和注册
- **动态路由**: `/service/{name}` 端点进行服务实例查询
- **负载均衡**: 简单的首实例选择策略（与Go原型保持一致）
- **健康检查**: `/service/health` 健康检查端点
- **可观测性**: 集成Spring Boot Actuator

## 技术栈

- **Java**: JDK 8+
- **框架**: Spring Boot 2.7.18
- **服务发现**: Alibaba Nacos
- **构建工具**: Maven
- **测试**: JUnit 5 + Mockito

## 项目结构

```
src/main/java/com/ksjx/gateway/
├── ApiGatewayApplication.java     # 主启动类
├── controller/
│   └── GatewayController.java     # 网关控制器（对应Go main.go的路由逻辑）
├── service/
│   └── ServiceDiscoveryService.java # 服务发现服务（对应Go getServiceInstances函数）
└── config/
    └── NacosConfig.java           # Nacos配置（对应Go initNacosClient函数）
```

## 配置说明

### application.yml
```yaml
server:
  port: 8080

spring:
  application:
    name: ksjx-api-gateway
  cloud:
    nacos:
      discovery:
        server-addr: console.nacos.io:80
        namespace: public
        context-path: /nacos
```

## API端点

### 服务发现
```
GET /service/{serviceName}
```

**响应示例**:
```json
{
  "target": "192.168.1.1",
  "service": "user-service", 
  "totalInstances": 2
}
```

### 健康检查
```
GET /service/health
```

**响应示例**:
```json
{
  "status": "UP",
  "service": "ksjx-api-gateway"
}
```

## 构建和运行

### 构建项目
```bash
mvn clean compile
```

### 运行测试
```bash
mvn test
```

### 启动应用
```bash
mvn spring-boot:run
```

或者构建jar包运行：
```bash
mvn package
java -jar target/api-gateway-1.0.0.jar
```

## 与Go原型的对应关系

| Go函数/结构 | Java实现 | 说明 |
|------------|----------|------|
| `main()` | `ApiGatewayApplication.main()` | 应用启动入口 |
| `initNacosClient()` | `NacosConfig.initNacosClient()` | Nacos客户端初始化 |
| `getServiceInstances()` | `ServiceDiscoveryService.getServiceInstances()` | 获取服务实例 |
| `createHandler()` | `GatewayController.getServiceInstance()` | 请求处理逻辑 |
| Gin路由 | Spring MVC Controller | HTTP路由处理 |

## 扩展功能

基于README.md中的架构设计，可以进一步扩展：

1. **认证鉴权**: 添加JWT认证中间件
2. **限流控制**: 集成Redis实现令牌桶算法
3. **负载均衡**: 实现轮询、权重等策略
4. **监控指标**: 集成Prometheus和Micrometer
5. **配置中心**: 动态路由配置支持

## 部署说明

支持容器化部署：
```dockerfile
FROM openjdk:8-jre-slim
COPY target/api-gateway-1.0.0.jar app.jar
EXPOSE 8080
ENTRYPOINT ["java", "-jar", "/app.jar"]
```