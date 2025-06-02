# APISIX Backend

Backend service for APISIX Custom Dashboard built with Go and GoFiber.

## Features

- RESTful API to manage APISIX routes
- API key authentication
- Connects to MariaDB for data storage

## Requirements

- Go 1.20+
- MariaDB server

## Setup

1. Clone repo:
```bash
git clone https://github.com/6410685066/apisix-backend.git
cd apisix-backend
```

2. Install dependencies:
```bash
go mod tidy
```

3. Set environment variables (via `.env` or system):
```
DB_USER=
DB_PASSWORD=
DB_HOST=
DB_NAME=
API_PREFIX=
ALLOW_ORIGINS=
```

1. Run the server:
```bash
go run main.go
```

## Build
1. Build File main
```bash
go build -o apisix-backend
```

2. Build Docker Image
```bash
docker build -t go-app .
```


