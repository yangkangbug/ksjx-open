#!/bin/bash

# KSJX API Gateway Demo Script
# This script demonstrates how to run and test the Spring Boot API Gateway

echo "=== KSJX API Gateway Demo ==="
echo

echo "1. Building the application..."
mvn clean package -DskipTests
echo

echo "2. Application details:"
echo "   - JAR file: target/api-gateway-1.0.0.jar"
echo "   - Size: $(du -h target/api-gateway-1.0.0.jar | cut -f1)"
echo "   - Main class: com.ksjx.gateway.ApiGatewayApplication"
echo

echo "3. Available endpoints:"
echo "   - Service Discovery: GET /service/{serviceName}"
echo "   - Health Check: GET /service/health"
echo "   - Actuator Health: GET /actuator/health"
echo

echo "4. To run the application:"
echo "   java -jar target/api-gateway-1.0.0.jar"
echo

echo "5. Example API calls (when running):"
echo "   curl http://localhost:8080/service/health"
echo "   curl http://localhost:8080/service/user-service"
echo "   curl http://localhost:8080/actuator/health"
echo

echo "6. Configuration:"
echo "   - Nacos Server: console.nacos.io:80"
echo "   - Namespace: public"
echo "   - Port: 8080"
echo

echo "7. Architecture mapping from Go prototype:"
echo "   Go Function              -> Java Implementation"
echo "   main()                   -> ApiGatewayApplication.main()"
echo "   initNacosClient()        -> NacosConfig.initNacosClient()"
echo "   getServiceInstances()    -> ServiceDiscoveryService.getServiceInstances()"
echo "   createHandler()          -> GatewayController.getServiceInstance()"
echo "   Gin routing              -> Spring MVC Controller"
echo

echo "Demo completed successfully!"