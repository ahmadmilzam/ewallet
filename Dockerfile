# Builder stage
FROM golang:1.22.2-alpine3.19 AS builder
WORKDIR /app
COPY . .
RUN mkdir -p bin && go build -o bin/main cmd/main.go

# Runner stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/bin/main .
COPY migrations ./migrations
# COPY config/config.yaml ./config/config.yaml
RUN apk update && \
  apk add --no-cache tzdata
ENV TZ=Asia/Jakarta
EXPOSE 3000
ENTRYPOINT [ "/app/main", "start" ]
