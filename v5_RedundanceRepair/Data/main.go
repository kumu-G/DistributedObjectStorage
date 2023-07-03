package main

import (
	"Data/heartbeat"
	"Data/locate"
	"Data/objects"
	"Data/temp"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	locate.CollectObjects() // 4.0 新增预读取本地对象
	go heartbeat.StarHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	fmt.Println("Data Server starting at port: ", os.Getenv("LISTEN_ADDRESS"))
	fmt.Println("Data Rabbitmq_Server: ", os.Getenv("RABBITMQ_SERVER"))
	fmt.Println("Data ES_Server: ", os.Getenv("ES_SERVER"))
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
