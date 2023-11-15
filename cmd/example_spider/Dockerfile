FROM alpine:latest

WORKDIR /app

COPY ./releases/example_spider .
COPY ./example.yml .

ENTRYPOINT ["./example_spider"]
CMD ["-c", "example.yml"]