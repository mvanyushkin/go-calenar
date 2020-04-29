FROM golang:latest

WORKDIR /usr/src/app
COPY . .
RUN go build -o server cmd/server/main.go
RUN cp configs/server/local_config.json local_config.json

EXPOSE 8888
CMD ["./server"]




