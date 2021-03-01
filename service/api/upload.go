package api

import (
	"fmt"
	"fsManage/utils"
	"goutils/network"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// 检测文件存在时重命名，防止文件误覆盖
func getFileSave(dir, filename string) string {
	ext := filepath.Ext(filename)
	name := strings.TrimRight(filename, ext)
	for i := 0; i < 10; i++ {
		fileSave := filepath.Join(dir, filename)
		if _, e := os.Stat(fileSave); e == nil {
			filename = name + "(" + strconv.Itoa(i) + ")" + ext
		} else {
			return fileSave
		}
	}
	return filepath.Join(dir, fmt.Sprintf("%s(%d)%s", name, time.Now().UnixNano(), ext))
}

// Upload upload file
func Upload(w http.ResponseWriter, r *http.Request) {
	utils.InitFormValue(r)
	dir := r.FormValue("dir")

	if "" == dir {
		network.JSON(w, 500, "param [dir] is empty", nil)
		return
	}
	if isLockUpload && !isInLockDir(dir) {
		network.JSON(w, 500, "No permission to upload", nil)
		return
	}
	file, fH, eFile := r.FormFile("file")

	if nil != eFile {
		network.JSON(w, 500, eFile.Error(), nil)
		return
	}
	fileSave := getFileSave(dir, fH.Filename)
	dstFile, eCreate := os.Create(fileSave)
	if nil != eCreate {
		network.JSON(w, 500, eCreate.Error(), nil)
		return
	}
	defer dstFile.Close()

	io.Copy(dstFile, file)
	network.JSON(w, 200, "success upload file ["+fileSave+"]", nil)
}
