FROM alpine:latest

WORKDIR /app

COPY ./releases/example_spider .
COPY ./example.yml .

RUN chmod +x ./example_spider

ENTRYPOINT ["./example_spider"]
CMD ["-c", "example.yml"]