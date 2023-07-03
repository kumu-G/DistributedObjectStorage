package main

import (
	"Api/es"
	"Api/utils"
	"log"
	"os"
	"path/filepath"
	"strings"

	"Api/objects"
)

func main() {
	files, _ := filepath.Glob(os.Getenv("STORAGE_ROOT") + "/objects/*")
	for i := range files {
		hash := strings.Split(filepath.Base(files[i]), ".")[0]
		vertify(hash)
	}
}

func vertify(hash string) {
	log.Println("verify", hash)
	size, e := es.SearchHashSize(hash)
	if e != nil {
		log.Println(e)
		return
	}
	// fmt.Println("size:", size)
	// fmt.Println("hash", hash)
	stream, e := objects.GetStream(hash, size)
	if e != nil {
		log.Println(e)
		return
	}
	d := utils.CalculateHash(stream)
	// fmt.Println("d", d)
	if d != hash {
		log.Printf("object hash mismatch, calculated=%s, requested=%s", d, hash)
	}
	stream.Close()
}
