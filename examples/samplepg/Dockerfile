FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
ENTRYPOINT ["./app"]
