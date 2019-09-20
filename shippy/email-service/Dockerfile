FROM alpine:latest

RUN mkdir /app

WORKDIR /app

ADD email-service /app/

ENTRYPOINT ["./email-service"]