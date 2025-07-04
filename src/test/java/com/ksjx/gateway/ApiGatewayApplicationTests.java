package com.ksjx.gateway;

import org.junit.jupiter.api.Test;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.TestPropertySource;

/**
 * Basic integration test for the API Gateway application
 */
@SpringBootTest
@TestPropertySource(properties = {
    "spring.cloud.nacos.discovery.enabled=false",
    "spring.cloud.nacos.config.enabled=false"
})
class ApiGatewayApplicationTests {

    @Test
    void contextLoads() {
        // Test that the Spring context loads successfully
    }
}