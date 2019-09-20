FROM alpine:latest

RUN mkdir /app

WORKDIR /app

ADD consignment.json /app/
ADD consignment-cli /app/

ENTRYPOINT ["./consignment-cli"]