package locate

import (
	"Data/rabbitmq"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

var objects = make(map[string]int)
var mutex sync.Mutex

func Locate(hash string) bool {
	mutex.Lock()
	_, ok := objects[hash]
	mutex.Unlock()
	return ok
}

/* 3.0
func Locate(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
*/
func Add(hash string) {
	mutex.Lock()
	objects[hash] = 1
	mutex.Unlock()
}

func Del(hash string) {
	mutex.Lock()
	delete(objects, hash)
	mutex.Unlock()
}

func StartLocate() {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer q.Close()
	q.Bind("dataServers")
	c := q.Consume()
	for msg := range c {

		hash, e := strconv.Unquote(string(msg.Body))
		if e != nil {
			panic(e)
		}
		exist := Locate(hash)
		if exist {
			q.Send(msg.ReplyTo, os.Getenv("LISTEN_ADDRESS"))
		}

		/* 3.0
		object, e := strconv.Unquote(string(msg.Body))
		fmt.Println(object, os.Getenv("LISTEN_ADDRESS"))
		if e != nil {
			panic(e)
		}
		if Locate(os.Getenv("STORAGE_ROOT") + "/objects/" + object) {
			q.Send(msg.ReplyTo, os.Getenv("LISTEN_ADDRESS"))
		}
		*/
	}
}

func CollectObjects() {
	files, _ := filepath.Glob(os.Getenv("STORAGE_ROOT") + "/objects/*")
	for i := range files {
		hash := filepath.Base(files[i])
		objects[hash] = 1
	}
	for key, value := range objects {
		fmt.Println("key:", key, " value:", value)
	}
}
