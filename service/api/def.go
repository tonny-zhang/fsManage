package api

import (
	"os"
	"path"
	"path/filepath"
)

// 当前执行文件所在目录
var dirCurrent string

func init() {
	if len(os.Args) > 1 {
		dirCurrent = os.Args[1]
	} else {
		var filep, _ = filepath.Abs(os.Args[0])
		dirCurrent = path.Dir(filep)
	}
}