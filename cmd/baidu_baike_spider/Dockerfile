FROM golang:alpine as builder

RUN apk add git
RUN apk add make

WORKDIR /app

RUN git clone -b main --depth 1 https://github.com/lizongying/go-crawler-example.git crawler
RUN cd crawler && make baidu_baike_spider

FROM alpine:latest as prod

WORKDIR /app

COPY --from=builder /app/crawler/releases/baidu_baike_spider /usr/bin/spider
COPY --from=builder /app/crawler/example.yml .

#ENTRYPOINT ["spider"]