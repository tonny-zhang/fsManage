package api

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Download 下载指定文件
func Download(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	file := r.FormValue("file")
	if "" != file {
		file, _ = url.QueryUnescape(file)

		if !strings.HasPrefix(file, "/") {
			file = filepath.Join(dirCurrent, file)
		}
	}
	file = filepath.Clean(file)
	f, e := os.Open(file)
	if e != nil {
		w.WriteHeader(404)
		w.Write([]byte("404 not found"))
		return
	}
	defer f.Close()

	s, _ := f.Stat()
	if s.IsDir() {
		w.Write([]byte("[" + file + "] is folder"))
		return
	}
	w.Header().Add("Content-Length", strconv.Itoa(int(s.Size())))
	w.Header().Add("Content-Disposition", "attachment; filename=\""+filepath.Base(file)+"\";")
	io.Copy(w, f)
}
