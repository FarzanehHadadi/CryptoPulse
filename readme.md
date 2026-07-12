# CryptoPulse

> A real-time cryptocurrency market monitoring, analysis, and alert platform built with Go.

## Overview

CryptoPulse is a backend-focused project designed to monitor cryptocurrency markets in real time, collect market data from exchanges, analyze technical indicators, and generate intelligent alerts.

The primary goal of this project is **not automated trading**. Instead, it aims to build a production-grade market intelligence platform while exploring modern Go backend development practices.

---

# Goals

- Collect real-time market data
- Store historical market information
- Analyze technical indicators
- Generate trading signals
- Send alerts
- Provide APIs for dashboards

---

# Project Scope

### Included

- Market Data Collection
- Historical Data Storage
- Technical Analysis
- Signal Generation
- Alert Engine
- REST API
- Dashboard APIs
- Scheduler
- Concurrent Workers

### Not Included (V1)

- Order execution
- Portfolio management
- Exchange authentication
- Automated trading

---

# Supported Exchange

- Binance Spot Market

Future versions:

- Bybit
- OKX
- KuCoin

---

# Supported Assets

Initial version:

- BTCUSDT
- ETHUSDT
- SOLUSDT
- BNBUSDT

Future versions will support dynamic watchlists.

---

# Tech Stack

## Backend

- Go
- Gin
- PostgreSQL
- Redis
- sqlc + pgx (planned)
- Docker
- Swagger
- Validator

## Infrastructure

- Docker Compose
- Redis
- Background Workers

## Frontend

- Next.js
- Highcharts

---

# High Level Architecture

```
Exchange APIs
        │
        ▼
 Market Collector
        │
        ▼
 PostgreSQL
        │
        ▼
 Analysis Engine
        │
        ▼
 Signal Engine
        │
        ▼
 Alert Service
        │
        ▼
 REST API
        │
        ▼
 Dashboard
```

---

# Planned Features

- Concurrent data collectors
- Scheduler
- Technical indicators
- Alert rules
- Telegram notifications
- Historical analytics

---

# Roadmap

## V1

- Project setup
- Market data collector
- PostgreSQL storage
- Scheduler
- REST API

---

## V2

- Technical indicators
- Moving Average
- RSI
- MACD
- Signal Engine
- Telegram alerts

---

## V3

- WebSocket market stream
- Redis caching
- Dashboard improvements
- Advanced analytics

---

## V4

- gRPC
- Service separation
- Distributed architecture

---

## V5

- Kafka/NATS
- Event-driven architecture
- Horizontal scaling

---

## V6

- Kubernetes
- CI/CD
- Monitoring
- Prometheus
- Grafana

---

## V7

- Backtesting Engine
- Strategy evaluation
- Historical simulations

---

## V8

- AI-powered analysis
- LLM market explanation
- Natural language insights

---

# Learning Objectives

This project focuses on learning:

- Production Go
- Concurrency
- Worker Pools
- Context
- Distributed Systems
- System Design
- Market Analysis
- Clean Architecture
