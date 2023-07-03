LISTEN_ADDRESS=172.29.1.1:12345 STORAGE_ROOT=/tmp/1 go run Data/main.go &
LISTEN_ADDRESS=172.29.1.2:12345 STORAGE_ROOT=/tmp/1 go run Data/main.go &
LISTEN_ADDRESS=172.29.1.3:12345 STORAGE_ROOT=/tmp/1 go run Data/main.go &
LISTEN_ADDRESS=172.29.1.4:12345 STORAGE_ROOT=/tmp/1 go run Data/main.go &
LISTEN_ADDRESS=172.29.1.5:12345 STORAGE_ROOT=/tmp/1 go run Data/main.go &
LISTEN_ADDRESS=172.29.1.6:12345 STORAGE_ROOT=/tmp/1 go run Data/main.go &

LISTEN_ADDRESS=172.29.2.1:12345 go run Api/main.go &
LISTEN_ADDRESS=172.29.2.2:12345 go run Api/main.go &
echo "Start"
                      
