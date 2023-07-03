package temp

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func post(w http.ResponseWriter, r *http.Request) {
	output, _ := exec.Command("uuidgen").Output()
	uuid := strings.TrimSuffix(string(output), "\n")
	name := strings.Split(r.URL.EscapedPath(), "/")[2]
	size, e := strconv.ParseInt(r.Header.Get("size"), 0, 64)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	t := tempInfo{uuid, name, size}
	e = t.writeToFile()
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, e = os.Create(os.Getenv("STORAGE_ROOT") + "/temp/" + t.Uuid + ".dat")
	if e != nil {
		log.Println("temp.go os.Create .dat fail:", e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte(uuid))
}
