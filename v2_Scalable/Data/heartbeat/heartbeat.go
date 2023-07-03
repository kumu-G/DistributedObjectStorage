package heartbeat

import (
	"Data/rabbitmq"
	"fmt"
	"os"
	"time"
)

func StarHeartbeat() {
	fmt.Println("RABBITMQ_SERVER:", os.Getenv("RABBITMQ_SERVER"))
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer q.Close()
	for {
		q.Publish("apiServers", os.Getenv("LISTEN_ADDRESS"))
		time.Sleep(5 * time.Second)
	}
}
