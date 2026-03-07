# IoT Telemetry Stream Platform

## Overview

The **IoT Telemetry Stream Platform** is a real-time IoT telemetry streaming platform built with Go microservices, Kafka event streaming, and Flink-based stream processing.

---

Input
↓
Ingestion
↓
Streaming Processing
↓
Storage
↓
Dashboards (Output)

---

## Input:
1 - Device Telemetry schema:
```json
{
  "event_id": "string",
  "device_id": "string",
  "timestamp": "int64",
  "event_type": "string",
  "schema_version": "int",

  "metrics": {
    "temperature": "float",
    "pressure": "float",
    "vibration": "float"
  },

  "metadata": {
    "location": "string",
    "firmware_version": "string"
  }
}
```

---

## Output - Dashboards
1 - Ingestion Health Dashboard:
    - event-generator -> ingestion-service -> kafka
    - Metrics: 
        - Events per second
        - Ingestion latency
        - Publish success rate
        - Broker throughput
        - Queue backlog

2 - Aggregation Analytics Dashboard:
    kafka -> aggregator-streaming -> clickhouse -> dashboard
    - Metrics: 
        - Mean temperature per device
        - Vibration trend over time
        - Events per minute

3 - Anomaly Detection Dashboard:
    kafka -> anomaly-detector-streaming -> clickhouse
    - Metrics: 
        - Anomaly score distribution
        - Alert frequency
        - Detection latency
        - False positive signals
