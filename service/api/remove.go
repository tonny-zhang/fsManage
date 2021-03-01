package api

import (
	"fsManage/utils"
	"goutils/network"
	"net/http"
	"os"
	"path/filepath"
)

// Remove 删除指定文件（暂时不支持删除目录）
func Remove(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	utils.InitFormValue(r)

	file := r.FormValue("file")

	if "" == file {
		network.JSON(w, 500, "param [file] is empty", nil)
	} else {
		file = filepath.Clean(file)
		if isLockRemove && !isInLockDir(file) {
			network.JSON(w, 500, "No permission to delete", nil)
			return
		}

		s, e := os.Stat(file)
		if nil != e {
			network.JSON(w, 500, e.Error(), nil)
			return
		} else if s.IsDir() {
			network.JSON(w, 500, "not support delete dir", nil)
			return
		}

		e = os.Remove(file)
		if nil == e {
			network.JSON(w, 200, "success", nil)
		} else {
			network.JSON(w, 500, e.Error(), nil)
		}
	}
}
