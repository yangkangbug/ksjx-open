package com.ksjx.gateway.controller;

import com.ksjx.gateway.service.ServiceDiscoveryService;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.MockitoAnnotations;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;

import java.util.Arrays;
import java.util.Collections;
import java.util.List;
import java.util.Map;

import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.Mockito.*;

/**
 * Unit tests for GatewayController
 */
class GatewayControllerTest {

    @Mock
    private ServiceDiscoveryService serviceDiscoveryService;

    @InjectMocks
    private GatewayController gatewayController;

    @BeforeEach
    void setUp() {
        MockitoAnnotations.openMocks(this);
    }

    @Test
    void testGetServiceInstance_Success() {
        // Arrange
        String serviceName = "test-service";
        List<String> instances = Arrays.asList("192.168.1.1", "192.168.1.2");
        when(serviceDiscoveryService.getServiceInstances(serviceName)).thenReturn(instances);

        // Act
        ResponseEntity<Map<String, Object>> response = gatewayController.getServiceInstance(serviceName);

        // Assert
        assertEquals(HttpStatus.OK, response.getStatusCode());
        assertNotNull(response.getBody());
        assertEquals("192.168.1.1", response.getBody().get("target"));
        assertEquals(serviceName, response.getBody().get("service"));
        assertEquals(2, response.getBody().get("totalInstances"));
        verify(serviceDiscoveryService, times(1)).getServiceInstances(serviceName);
    }

    @Test
    void testGetServiceInstance_NoInstances() {
        // Arrange
        String serviceName = "non-existent-service";
        when(serviceDiscoveryService.getServiceInstances(serviceName)).thenReturn(Collections.emptyList());

        // Act
        ResponseEntity<Map<String, Object>> response = gatewayController.getServiceInstance(serviceName);

        // Assert
        assertEquals(HttpStatus.INTERNAL_SERVER_ERROR, response.getStatusCode());
        assertNotNull(response.getBody());
        assertEquals("No available instances", response.getBody().get("error"));
        verify(serviceDiscoveryService, times(1)).getServiceInstances(serviceName);
    }

    @Test
    void testGetServiceInstance_Exception() {
        // Arrange
        String serviceName = "error-service";
        when(serviceDiscoveryService.getServiceInstances(serviceName))
            .thenThrow(new RuntimeException("Nacos connection failed"));

        // Act
        ResponseEntity<Map<String, Object>> response = gatewayController.getServiceInstance(serviceName);

        // Assert
        assertEquals(HttpStatus.INTERNAL_SERVER_ERROR, response.getStatusCode());
        assertNotNull(response.getBody());
        assertTrue(response.getBody().get("error").toString().contains("Failed to get service instances"));
        verify(serviceDiscoveryService, times(1)).getServiceInstances(serviceName);
    }

    @Test
    void testHealth() {
        // Act
        ResponseEntity<Map<String, Object>> response = gatewayController.health();

        // Assert
        assertEquals(HttpStatus.OK, response.getStatusCode());
        assertNotNull(response.getBody());
        assertEquals("UP", response.getBody().get("status"));
        assertEquals("ksjx-api-gateway", response.getBody().get("service"));
    }
}