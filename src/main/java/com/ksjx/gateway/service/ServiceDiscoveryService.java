package com.ksjx.gateway.service;

import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.cloud.client.ServiceInstance;
import org.springframework.cloud.client.discovery.DiscoveryClient;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.List;

/**
 * Service Discovery Service - equivalent to the Go getServiceInstances function
 * Handles Nacos service discovery integration
 */
@Service
@Slf4j
public class ServiceDiscoveryService {

    @Autowired
    private DiscoveryClient discoveryClient;

    /**
     * Get service instances from Nacos - equivalent to Go getServiceInstances function
     * @param serviceName the name of the service to discover
     * @return list of instance IP addresses
     */
    public List<String> getServiceInstances(String serviceName) {
        try {
            log.debug("Querying Nacos for service instances of: {}", serviceName);
            
            List<ServiceInstance> instances = discoveryClient.getInstances(serviceName);
            List<String> instanceAddresses = new ArrayList<>();
            
            if (instances != null) {
                for (ServiceInstance instance : instances) {
                    String address = instance.getHost();
                    instanceAddresses.add(address);
                    log.debug("Found instance: {} for service: {}", address, serviceName);
                }
            }
            
            log.info("Found {} instances for service: {}", instanceAddresses.size(), serviceName);
            return instanceAddresses;
            
        } catch (Exception e) {
            log.error("Failed to get service instances from Nacos for service {}: {}", serviceName, e.getMessage(), e);
            return new ArrayList<>(); // Return empty list on error
        }
    }

    /**
     * Get all available services from Nacos
     * @return list of service names
     */
    public List<String> getAllServices() {
        try {
            return discoveryClient.getServices();
        } catch (Exception e) {
            log.error("Failed to get all services from Nacos: {}", e.getMessage(), e);
            return new ArrayList<>();
        }
    }
}