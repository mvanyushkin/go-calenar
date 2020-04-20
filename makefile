include .env

init:
	@echo "[x] Loading dependencies..."
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u golang.org/x/lint/golint
	go get -u github.com/pressly/goose/cmd/goose
	go mod download
	@echo  "-> Done."

build:

	@echo "[x] Running the formatting tool"
	go fmt ./...
	@echo  "-> Done."

	@echo "[x] Running the basic linter GO VET"
	go vet ./...
	@echo  "-> Done. But we do NOT trust in GO VET, so we are gonna use GOLINT"

	@echo "[x] Running GOLINT..."
	golint
	@echo  "-> Done."

	@echo "[x] Running unit tests..."
	go test ./...
	@echo  "-> Done."

	@echo  "[x] Building the project..."
	go build -o build/server/server cmd/server/main.go
	go build -o build/client/client cmd/client/main.go
	go build -o build/reminder/reminder cmd/reminder/main.go
	go build -o build/sender/sender cmd/sender/main.go
	@echo  "-> Done"

	@echo "[x] Copying configuration files to the their own binaries"
	cp configs/server/local_config.json build/server/local_config.json
	cp configs/server/local_config.json build/reminder/local_config.json
	cp configs/server/local_config.json build/sender/local_config.json
	@echo  "-> Fuck Yeah! Nice job, Man!"


run:
	goose -dir ./scripts/migrations postgres "user=$(PG_USERNAME) password=$(PG_PASSWORD) dbname=$(DB_NAME) sslmode=disable" up
	./build/server/server& ./build/sender/sender &./build/reminder/reminder &./build/client/client

clean:
	@echo "[x] Cleaning up the previous build result"
	rm -r -f ./build
	@echo  "-> Done."



