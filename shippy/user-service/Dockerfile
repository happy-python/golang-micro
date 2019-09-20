FROM alpine:latest

RUN mkdir /app

WORKDIR /app

ADD user-service /app/

ENTRYPOINT ["./user-service"]