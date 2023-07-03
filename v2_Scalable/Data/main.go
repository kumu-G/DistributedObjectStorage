package main

import (
	"Data/heartbeat"
	"Data/locate"
	"Data/objects"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	go heartbeat.StarHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	fmt.Println("Data Server starting at port: ", os.Getenv("LISTEN_ADDRESS"))
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
