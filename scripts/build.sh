
echo '======= Preconfiguration ======='
sh init.sh

pushd ..
echo '======== Lintering ========='
#echo pwd
#echo 'Running go fmt...'
#go fmt
#ret=$?
#if [ $ret -ne 0 ]
#then
#  echo 'an error has happened'
#  exit 1
#fi
#echo 'done.'

echo 'Running go vet...'
go vet
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
go test ./...

echo '======== Building ========='
go build -o cmd/server/server server.go
cp config/local_config.json cmd/server/local_config.json

echo '======== Setting up the access rights ========='
chmod +x cmd/server/server