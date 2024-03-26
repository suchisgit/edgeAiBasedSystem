# Flexible AI Model Serving and GPU Resource Management in Hybrid Clouds

This repository contains a Next.js-based frontend and a Golang-based backend, created to facilitate flexible AI model serving and GPU resource management in hybrid cloud environments.

## Key Features

**GPU Utilization Monitoring**: Designed a system to monitor GPU utilization on individual nodes and present the data on Grafana Dashboards.

**Metrics Visualization**: The Prometheus server retrieves metrics from the DCGM pod and makes them available for Grafana to visualize.

**Log Monitoring**: Integrated Loki to display logs from deployed ML pods on the Grafana dashboard.

**User-friendly Deployment**: Developed an application with a user-friendly interface that streamlines pod deployment. Users can input their DockerHub image, and the application automatically allocates a pod.


## High level architecture.

![Architecture](https://github.com/suchisgit/edgeAiBasedSystem/blob/main/MonitoringSystem-architecture.png)

## How to Use:

### 1) automation bash scripts inside automation-scripts folder has to be run.

### 2) all the pods has to be deployed based on the configuration files provided in the configuration-files folder.

### 3) command to run front-end <br>
`npm run dev -- -p 4000`

### 4) command to run back-end <br>
`run main.go`

