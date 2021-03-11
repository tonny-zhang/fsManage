# fsManage 为了方便快速在电脑上搭建一个可以基于web的本地文件管理系统，以方便让其它电脑或设备对其进行管理

## 用法
1. 直接运行，将对二进制所在目录进行权限开放，其它目录将没有上传、删除等敏感的权限
1. 指定目录运行 `fsm /doc/` 将对 `/doc/` 目录进行权限开放

## 编译
* 从源码编译 `go build -ldflags "-s -w" -o fsm`

## 依赖包处理
应用里使用了自身维护的一个工具包`goutils`, 这个可以在`go.mod`里查看，也可以把 [goutils](https://github.com/tonny-zhang/goutils)源码下载到和项目同级目录下即可。