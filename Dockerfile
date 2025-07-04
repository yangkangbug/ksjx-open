# Multi-stage build for optimized image size
FROM openjdk:8-jre-alpine

# Add metadata
LABEL maintainer="KSJX Team"
LABEL description="KSJX API Gateway - Spring Boot implementation"
LABEL version="1.0.0"

# Create app directory
WORKDIR /app

# Copy the JAR file
COPY target/api-gateway-1.0.0.jar app.jar

# Expose port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/actuator/health || exit 1

# Set JVM options for containerized environment
ENV JAVA_OPTS="-Xmx512m -Xms256m -Djava.security.egd=file:/dev/./urandom"

# Run the application
ENTRYPOINT ["sh", "-c", "java $JAVA_OPTS -jar /app/app.jar"]