# Telemetry Simulator

### 1. Service Overview

The `telemetry-simulator` acts as the initial data source in the platform architecture. It is responsible for generating and sending JSON-formatted telemetry events containing simulated metrics (such as temperature, pressure, and vibration) and metadata for imaginary IoT devices.

### 2. Responsibilities

* Generate realistic synthetic telemetry data simulating varying device conditions.
* Send generated events directly to the `ingestion-service` over the network.
* Simulate different load profiles to test the platform's throughput and resilience.

### 3. Real-World Challenges

To focus strictly on the stream processing architecture, this project simplifies the data origin layer:
* This service generates and pushes events directly via HTTP/network.
* In real-world production environments, this layer would typically involve:
  * Physical IoT devices and sensors.
  * **IoT Gateways** aggregating field device traffic.
  * **Edge Computing** layers performing local filtering and aggregation.
  * MQTT brokers and robust hardware authentication (e.g., mTLS).
