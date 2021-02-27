package main

import "fsManage/service"

func main() {
	// b, e := ui.ReadFile("ui/index.html")
	// fmt.Println(string(b), e)

	service.Start(4000, "")
}
