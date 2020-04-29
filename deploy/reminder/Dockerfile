FROM golang:latest

WORKDIR /usr/src/app
COPY . .
RUN go build -o reminder cmd/reminder/main.go
RUN cp configs/reminder/local_config.json local_config.json
RUN cp configs/reminder/local_config.json local_config.json

EXPOSE 8888

CMD ["./reminder"]




