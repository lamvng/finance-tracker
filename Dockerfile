# Build stage
FROM golang:1.22.4-alpine3.20 AS build
WORKDIR /app
COPY go.mod go.sum .
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o finance-tracker .

# Production stage
FROM alpine:3.20.2 AS deploy
WORKDIR /app
USER finance-tracker:finance-tracker
EXPOSE 8080
COPY --from=build /app/finance-tracker .
ENTRYPOINT ["./finance-tracker"]
