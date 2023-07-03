package locate

import (
	"Data/rabbitmq"
	"Data/types"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

var objects = make(map[string]int)
var mutex sync.Mutex

// 5.0		return bool -> return int
func Locate(hash string) int {
	// fmt.Println("Objects:", objects)
	// fmt.Println("Objects hash:", hash)
	// fmt.Println("Objects k-v:", objects["p8SB0DNymMtu%252F9rWrY7Vm93NgNkq9eGXP9eLAZ6+q68="])
	mutex.Lock()
	id, ok := objects[hash]
	// fmt.Println("Objects id:", id)
	// fmt.Println("Objects ok:", ok)
	mutex.Unlock()
	if !ok {
		return -1
	}
	return id
}

// 5.0 modify
func Add(hash string, id int) {
	mutex.Lock()
	// objects[hash] = 1
	objects[hash] = id
	mutex.Unlock()
}

func Del(hash string) {
	mutex.Lock()
	delete(objects, hash)
	mutex.Unlock()
}

// 5.0 modify
func StartLocate() {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer q.Close()
	q.Bind("dataServers")
	c := q.Consume()
	for msg := range c {
		// fmt.Println("msg:", string(msg.Body))
		hash, e := strconv.Unquote(string(msg.Body))
		// fmt.Println("Data Locate hash:", hash)
		if e != nil {
			panic(e)
		}
		id := Locate(hash)
		// fmt.Println(os.Getenv("LISTEN_ADDRESS"), ",ID:", id) // id = -1
		if id != -1 {
			q.Send(msg.ReplyTo, types.LocateMessage{Addr: os.Getenv("LISTEN_ADDRESS"), Id: id})
		}
	}
}

func CollectObjects() {
	filename := os.Getenv("STORAGE_ROOT") + "/objects/*"
	files, e := filepath.Glob(filename)
	if e != nil {
		log.Println("Open " + filename + " fail!")
	}
	fmt.Println("CollectObjects", filename)
	for i := range files {
		file := strings.Split(filepath.Base(files[i]), ".")
		if len(file) != 3 {
			panic(files[i])
		}
		hash := file[0]
		id, e := strconv.Atoi(file[1])
		fmt.Println("CollectObjects() id: ", id)
		fmt.Println("CollectObjects() hash: ", hash)
		if e != nil {
			panic(e)
		}
		objects[hash] = id
	}
}
