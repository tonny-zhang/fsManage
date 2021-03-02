package api

import (
	"goutils/network"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// FileList 文件信息
type FileList struct {
	Name   string `json:"name"`
	IsDir  bool   `json:"isDir"`
	IsLink bool   `json:"isLink"`
	Size   int64  `json:"size"`
}

// ListDataRes res data for List
type ListDataRes struct {
	List []FileList `json:"list"`
	Nav  string     `json:"nav"`
}

// List 得把指定目录下的文件列表
func List(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	dir := r.FormValue("dir")
	if "" == dir {
		dir = dirCurrent
	} else {
		dir, _ = url.QueryUnescape(dir)

		if !strings.HasPrefix(dir, "/") {
			dir = filepath.Join(dirCurrent, dir)
		}
	}
	dir = filepath.Clean(dir)
	w.Header().Add("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	list, e := ioutil.ReadDir(dir)

	if nil == e {
		var data ListDataRes
		var flist = make([]FileList, 0)
		for _, f := range list {
			s, e := os.Stat(filepath.Join(dir, f.Name()))
			// 可能软链接失效
			if e != nil {
				continue
			}
			info := FileList{
				Name:  f.Name(),
				IsDir: s.IsDir(),
				Size:  f.Size(),
			}

			if !f.IsDir() && info.IsDir {
				info.IsLink = true
			}
			flist = append(flist, info)
		}
		data.Nav = dir
		data.List = flist
		network.JSON(w, 200, "success["+dir+"]", data)
	} else {
		network.JSON(w, 500, e.Error(), nil)
	}
}
