package main

import (
	"Api/heartbeat"
	"Api/locate"
	"Api/objects"
	"Api/versions"
	"fmt"
	"log"
	"net/http"
	"os"
)

/*
	数据服务层没有变化：元数据服务的交互完全在接口层实现。
*/

func main() {
	go heartbeat.ListenHeartbeat()                // 心跳-不变
	http.HandleFunc("/objects/", objects.Handler) // 通过es操作
	http.HandleFunc("/locate/", locate.Handler)   // 定位-不变
	http.HandleFunc("/versions/", versions.Handler)
	fmt.Println("Api Server starting at port: ", os.Getenv("LISTEN_ADDRESS"))
	fmt.Println("Api Rabbitmq_Server: ", os.Getenv("RABBITMQ_SERVER"))
	fmt.Println("Api ES_Server: ", os.Getenv("ES_SERVER"))
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
