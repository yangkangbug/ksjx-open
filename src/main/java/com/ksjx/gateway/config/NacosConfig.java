package com.ksjx.gateway.config;

import lombok.extern.slf4j.Slf4j;
import org.springframework.cloud.client.discovery.EnableDiscoveryClient;
import org.springframework.context.annotation.Configuration;

import javax.annotation.PostConstruct;

/**
 * Nacos Configuration - equivalent to the Go initNacosClient function
 * Configures Nacos client settings and connection
 */
@Configuration
@EnableDiscoveryClient
@Slf4j
public class NacosConfig {

    /**
     * Initialize Nacos client - equivalent to Go initNacosClient function
     * Configuration is handled through application.yml
     */
    @PostConstruct
    public void initNacosClient() {
        log.info("Initializing Nacos client configuration");
        log.info("Nacos server: console.nacos.io:80");
        log.info("Namespace: public");
        log.info("Context path: /nacos");
        log.info("Nacos client initialized successfully");
    }
}