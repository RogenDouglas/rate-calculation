FROM golang:1.19

WORKDIR /go/goapp

RUN apt-get update && apt-get install -y librdkafka-dev

CMD ["tail", "-f", "/dev/null"]