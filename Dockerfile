FROM golang:1.15

ENV APP_ENVIRONMENT=dev
ENV PORT_NUMBER=8080
ENV REDIS_ADDRESS=rate-app-redis:6379

ENV MAX_QUOTAS=100
ENV WINDOW_SECONDS=3600
ENV RATE_LIMITER_NAME=FixedWindow

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
