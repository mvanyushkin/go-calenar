FROM golang:latest

WORKDIR /usr/src/app
COPY . .
RUN go get github.com/pressly/goose/cmd/goose

ENTRYPOINT ./scripts/run-migration.sh
