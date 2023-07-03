package main

import (
	"Api/heartbeat"
	"Api/locate"
	"Api/objects"
	"Api/temp"
	"Api/versions"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	http.HandleFunc("/versions/", versions.Handler)
	fmt.Println("Api Server starting at port: ", os.Getenv("LISTEN_ADDRESS"))
	fmt.Println("Api Rabbitmq_Server: ", os.Getenv("RABBITMQ_SERVER"))
	fmt.Println("Api ES_Server: ", os.Getenv("ES_SERVER"))
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
