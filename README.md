## Build
# 1.Build File main
    GOOS=linux GOARCH=amd64 go build -o main .
# 2.Build Docker Image
    docker build -t go-app .


