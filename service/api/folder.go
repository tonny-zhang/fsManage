package api

import (
	"fsManage/utils"
	"goutils/network"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

// CreateFolder 创建指定目录
func CreateFolder(w http.ResponseWriter, r *http.Request) {
	utils.InitFormValue(r)
	dir := r.FormValue("dir")
	if "" != dir {
		dir, _ = url.QueryUnescape(dir)
		dir = filepath.Clean(dir)
	} else {
		network.JSON(w, 500, "param [dir] is empty", nil)
		return
	}
	if !filepath.IsAbs(dir) {
		dir = filepath.Join(dirCurrent, dir)
	}

	e := os.MkdirAll(dir, 0755)
	if nil == e {
		network.JSON(w, 200, "success create ["+dir+"]", nil)
	} else {
		network.JSON(w, 500, e.Error(), nil)
	}
}
