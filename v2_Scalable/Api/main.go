package main

import (
	"Api/heartbeat"
	"Api/locate"
	"Api/objects"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	fmt.Println("Api Server starting at port: ", os.Getenv("LISTEN_ADDRESS"))
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
