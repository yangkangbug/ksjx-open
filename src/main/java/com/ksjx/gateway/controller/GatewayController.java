package com.ksjx.gateway.controller;

import com.ksjx.gateway.service.ServiceDiscoveryService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * Gateway Controller - equivalent to the Go main.go routing logic
 * Handles service discovery and basic load balancing
 */
@RestController
@RequestMapping("/service")
@Slf4j
public class GatewayController {

    @Autowired
    private ServiceDiscoveryService serviceDiscoveryService;

    /**
     * Service discovery endpoint - equivalent to the Go createHandler function
     * GET /service/{name} returns the target instance for the service
     */
    @GetMapping("/{name}")
    public ResponseEntity<Map<String, Object>> getServiceInstance(@PathVariable String name) {
        log.info("Getting service instances for: {}", name);
        
        try {
            List<String> instances = serviceDiscoveryService.getServiceInstances(name);
            
            if (instances == null || instances.isEmpty()) {
                log.warn("No available instances for service: {}", name);
                Map<String, Object> errorResponse = new HashMap<>();
                errorResponse.put("error", "No available instances");
                return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR).body(errorResponse);
            }
            
            // Simple load balancing strategy - use first instance (same as Go prototype)
            String target = instances.get(0);
            log.info("Selected target instance: {} for service: {}", target, name);
            
            Map<String, Object> response = new HashMap<>();
            response.put("target", target);
            response.put("service", name);
            response.put("totalInstances", instances.size());
            
            return ResponseEntity.ok(response);
            
        } catch (Exception e) {
            log.error("Error getting service instances for {}: {}", name, e.getMessage(), e);
            Map<String, Object> errorResponse = new HashMap<>();
            errorResponse.put("error", "Failed to get service instances: " + e.getMessage());
            return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR).body(errorResponse);
        }
    }

    /**
     * Health check endpoint
     */
    @GetMapping("/health")
    public ResponseEntity<Map<String, Object>> health() {
        Map<String, Object> response = new HashMap<>();
        response.put("status", "UP");
        response.put("service", "ksjx-api-gateway");
        return ResponseEntity.ok(response);
    }
}