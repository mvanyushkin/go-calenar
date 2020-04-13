
echo '======= Preconfiguration ======='
./init.sh

echo '===== Regenerate contracts ====='
./generate-contracts.sh

echo '======== Lintering ========='

echo 'Running go fmt...'
go fmt  ../...
ret=$?
if [ $ret -ne 0 ]
then
  echo 'an error has happened'
  exit 1
fi
echo 'done.'

echo 'Running go vet...'
go vet ../...
ret=$?
if [ $ret -ne 0 ]
then
  echo 'an error has happened'
  exit 1
fi
echo 'done.'

echo 'Running golint...'
golint
ret=$?
if [ $ret -ne 0 ]
then
  echo 'an error has happened'
  exit 1
fi
echo 'done.'

echo '======== Testing ========='
go test ../...

echo '======== Building ========='

go build -o ../build/server/server ../cmd/server/main.go
go build -o ../build/client/client ../cmd/client/main.go
go build -o ../build/reminder/reminder ../cmd/reminder/main.go
go build -o ../build/sender/sender ../cmd/sender/main.go
cp ../local_config.json ../build/server/local_config.json
cp ../local_config.json ../build/reminder/local_config.json
cp ../local_config.json ../build/sender/local_config.json

echo '======== Setting up the access rights ========='
chmod +x ../build/server/server
chmod +x ../build/client/client
chmod +x ../build/sender/sender
chmod +x ../build/reminder/reminder

