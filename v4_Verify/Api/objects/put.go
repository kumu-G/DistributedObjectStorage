package objects

import (
	"Api/es"
	"Api/heartbeat"
	"Api/locate"
	"Api/objectstream"
	"Api/utils"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func put(w http.ResponseWriter, r *http.Request) {
	hash := utils.GetHashFromHeader(r.Header)
	if hash == "" {
		log.Println("missing object hash in digest header")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	size := utils.GetSizeFromHeader(r.Header)
	c, e := storeObject(r.Body, hash, size) // 4.0修改入参
	if e != nil {
		log.Println(e)
		w.WriteHeader(c)
		return
	}
	if c != http.StatusOK {
		w.WriteHeader(c)
		return
	}
	name := strings.Split(r.URL.EscapedPath(), "/")[2]
	e = es.AddVersion(name, hash, size)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

//  sotreObject 4.0 修改入参(r io.Reader, object string)->(r io.Reader, hash string, size int64)
func storeObject(r io.Reader, hash string, size int64) (int, error) {
	// b, _ := ioutil.ReadAll(r)
	// fmt.Println("put.go.storeObject.49:", string(b))
	if locate.Exist(url.PathEscape(hash)) {
		return http.StatusOK, nil
	}

	stream, e := putStream(url.PathEscape(hash), size) // 4.0修改入参
	if e != nil {
		return http.StatusServiceUnavailable, e
	}

	// 4.0 使用 TeeReader同步 r读和 stream写
	reader := io.TeeReader(r, stream)
	// 4.0 计算哈希值，从reader中读取hash
	d := utils.CalculateHash(reader)
	if d != hash {
		stream.Commit(false)
		return http.StatusBadRequest, fmt.Errorf("object hash mismatch, calculated=%s, requested=%s", d, hash)
	}
	stream.Commit(true)
	return http.StatusOK, nil
}

// putStream 4.0 修改入参(object string)->(hash string, size int64)
func putStream(hash string, size int64) (*objectstream.TempPutStream, error) {
	server := heartbeat.ChooseRandomDataServer()
	if server == "" {
		return nil, fmt.Errorf("cannot find any dataServer")
	}
	return objectstream.NewTempPutStream(server, hash, size) // NewPutStream->NewTempPutStream
}
