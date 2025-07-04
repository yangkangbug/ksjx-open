# KSJX API Gateway - API Examples

## Service Discovery Endpoint

### Request
```http
GET /service/user-service HTTP/1.1
Host: localhost:8080
```

### Response (Success)
```json
{
  "target": "192.168.1.10",
  "service": "user-service",
  "totalInstances": 3
}
```

### Response (No Instances)
```json
{
  "error": "No available instances"
}
```

## Health Check Endpoint

### Request
```http
GET /service/health HTTP/1.1
Host: localhost:8080
```

### Response
```json
{
  "status": "UP",
  "service": "ksjx-api-gateway"
}
```

## Actuator Health Endpoint

### Request
```http
GET /actuator/health HTTP/1.1
Host: localhost:8080
```

### Response
```json
{
  "status": "UP",
  "components": {
    "diskSpace": {
      "status": "UP",
      "details": {
        "total": 124554051584,
        "free": 112460136448,
        "threshold": 10485760
      }
    },
    "nacos": {
      "status": "UP",
      "details": {
        "connectedServices": 5
      }
    }
  }
}
```

## Error Responses

### Service Error
```json
{
  "error": "Failed to get service instances: Connection timeout"
}
```

## Architecture Comparison

| Feature | Go Implementation | Spring Boot Implementation |
|---------|------------------|---------------------------|
| Framework | Gin | Spring Boot + Spring MVC |
| Service Discovery | Nacos SDK Go | Spring Cloud Alibaba |
| Configuration | Hard-coded | application.yml |
| Logging | log package | SLF4J + Logback |
| Health Checks | Custom | Spring Actuator |
| Testing | Manual | JUnit 5 + Mockito |
| Packaging | Binary | Executable JAR |
| Dependencies | Go modules | Maven |