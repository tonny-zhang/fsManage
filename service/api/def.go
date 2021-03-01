package api

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

// 当前执行文件所在目录
var dirCurrent string

// 是否锁定删除启动目录下的文件
var isLockRemove = true

// 是否锁定创建目录
var isLockCreateFolder = true

// 是否锁定上传文件
var isLockUpload = true

func init() {
	if len(os.Args) > 1 {
		dirCurrent = os.Args[1]
	} else {
		var filep, _ = filepath.Abs(os.Args[0])
		dirCurrent = path.Dir(filep)
	}
}

// 是否在锁定的目录下
func isInLockDir(dir string) bool {
	return strings.HasPrefix(dir, dirCurrent+"/")
}
