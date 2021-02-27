package api

import (
	"goutils/network"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

// FileList 文件信息
type FileList struct {
	Name  string `json:"name"`
	IsDir bool   `json:"isDir"`
	Size  int64  `json:"size"`
}

// List 得把指定目录下的文件列表
func List(w http.ResponseWriter, r *http.Request) {
	dir := r.FormValue("dir")
	if "" == dir {
		dir = dirCurrent
	} else {
		dir = filepath.Join(dirCurrent, dir)
	}

	list, e := ioutil.ReadDir(dir)

	if nil == e {
		var flist = make([]FileList, 0)
		for _, f := range list {
			flist = append(flist, FileList{
				Name:  f.Name(),
				IsDir: f.IsDir(),
				Size:  f.Size(),
			})
		}
		network.JSON(w, 200, "success["+dir+"]", flist)
	} else {
		network.JSON(w, 500, "error", e.Error())
	}
}
