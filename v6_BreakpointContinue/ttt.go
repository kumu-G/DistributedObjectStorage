
acceptGzip := false
encoding := r.Header["Accept-Encoding"]
for i := range encoding {
	if encoding[i] == "gzip" {
		acceptGzip = true
		break
	}
}
if acceptGzip {
	w.Header().Set("content-encoding", "gzip")
	w2 := gzip.NewWriter(w)
	io.Copy(w2, stream)
	w2.Close()
} else {
	io.Copy(w, stream)
}
stream.Close()

io.Copy(w, stream)
stream.Close()



********

func commitTempObject(datFile string, tempinfo *tempInfo) {
	f, _ := os.Open(datFile)
	defer f.Close()
	d := url.PathEscape(utils.CalculateHash(f))
	f.Seek(0, io.SeekStart)
	w, _ := os.Create(os.Getenv("STORAGE_ROOT") + "/objects/" + tempinfo.Name + "." + d)
	w2 := gzip.NewWriter(w)
	io.Copy(w2, f)
	w2.Close()
	os.Remove(datFile)
	locate.Add(tempinfo.hash(), tempinfo.id())
}

*****

func sendFile(w io.Writer, file string) {
	f, e := os.Open(file)
	if e != nil {
		log.Println(e)
		return
	}
	defer f.Close()
	gzipStream, e := gzip.NewReader(f)
	if e != nil {
		log.Println(e)
		return
	}
	io.Copy(w, gzipStream)
	gzipStream.Close()
}

