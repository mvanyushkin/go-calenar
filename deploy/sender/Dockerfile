FROM golang:latest

WORKDIR /usr/src/app
COPY . .
RUN go build -o sender cmd/sender/main.go
RUN cp configs/sender/local_config.json local_config.json

EXPOSE 8888

CMD ["./sender"]




