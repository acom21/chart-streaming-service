#  Chart Streaming Service

Real-time streaming service that connects to Binance WebSocket API, aggregates trade ticks into OHLC candles, stores them in PostgreSQL, and optionally streams ticks via gRPC.

##  Features

- Connects to Binance aggregate trade WebSocket (`trade`)
- Aggregates trades into OHLC candles per symbol and interval
- Stores candles in PostgreSQL 
- Streams ticks to clients over gRPC

# Clone the project
git clone https://github.com/acom21/chart-streaming-service.git
cd chart-streaming-service

# Run service + PostgreSQL + migrations
docker compose up --wait
