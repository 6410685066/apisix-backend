FROM alpine:3.20
WORKDIR /app
COPY main .
EXPOSE 8080
CMD ["./main"]
