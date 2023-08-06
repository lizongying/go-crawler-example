FROM golang:alpine as builder

RUN apk add git
RUN apk add make

WORKDIR /app

RUN git clone -b main --depth 1 https://github.com/lizongying/go-crawler-example.git crawler
RUN cd crawler && make zhihuSpider

FROM alpine:latest as prod

WORKDIR /app

COPY --from=builder /app/crawler/releases/zhihuSpider /usr/bin/spider
COPY --from=builder /app/crawler/example.yml .

#ENTRYPOINT ["spider"]