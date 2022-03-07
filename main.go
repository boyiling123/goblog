package main

import (
	"fmt"
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Hello, 这里是 goblog</h1>")
	} else if r.URL.Path == "/about" {
		fmt.Fprint(w, "这是关于页面<a href='https://baidu.com'>访问百度</a>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>404</h1>")
	}

}

func main() {
	http.HandleFunc("/", handleFunc)
	http.ListenAndServe(":3000", nil)
}
