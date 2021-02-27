package service

import (
	"embed"
	"goutils/logger"
	"io/fs"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
)

// 通配符的默认目录和源文件所在的目录是同一目录，所以我们只能匹配同目录下的文件或目录，不能匹配到父目录
// https://www.cnblogs.com/apocelipes/p/13907858.html

//go:embed ui
var assets embed.FS

type fsFunc func(name string) (fs.File, error)

func (f fsFunc) Open(name string) (fs.File, error) {
	return f(name)
}

func assetHandler(prefix string, assets embed.FS, root string) http.Handler {
	handler := fsFunc(func(name string) (fs.File, error) {
		assetPath := path.Join("ui", name)

		// If we can't find the asset, fs can handle the error
		file, err := assets.Open(assetPath)
		if err != nil {
			return nil, err
		}

		// Otherwise assume this is a legitimate request routed correctly
		return file, err
	})

	return http.StripPrefix(prefix, http.FileServer(http.FS(handler)))
}

// Start 开启服务
// 只开启本地服务时host使用 127.0.0.1
func Start(port int, host string) {
	mux := http.NewServeMux()
	mux.Handle("/", assetHandler("/", assets, "ui"))

	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test"))
	})

	var server = &http.Server{
		Addr:         host + ":" + strconv.Itoa(port),
		WriteTimeout: time.Second * 60,
		Handler:      mux,
	}
	logger.PrintInfof("http service at %s", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		logger.PrintError(err)
		os.Exit(0)
	} else {
		logger.PrintInfof("http service at %s", server.Addr)
	}
}
