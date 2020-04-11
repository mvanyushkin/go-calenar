echo "-> Loading modules..."
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u golang.org/x/lint/golint
go get -u github.com/pressly/goose/cmd/goose
go mod download
echo "[x] Done."
echo "-> Generating grpc contracts..."
./generate-contracts.sh
echo "[x] Done."