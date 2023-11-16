FROM --platform=${BUILDPLATFORM} golang:1.21 AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=1 CGO_LDFLAGS="-static" go build -ldflags "-s -w -X $(module)/pkg/logger.name=example_spider" -o ./releases/example_spider ./cmd/example_spider


FROM alpine:latest

WORKDIR /app

COPY --from=builder  /app/releases/example_spider .
COPY --from=builder  /app/example.yml .

RUN chmod +x ./example_spider

ENTRYPOINT ["./example_spider"]
CMD ["-c", "example.yml"]